//
//  Simple Pirate broker.
//  This is identical to load-balancing pattern, with no reliability
//  mechanisms. It depends on the client for recovery. Runs forever.
//

package main

import (
	zmq "github.com/pebbe/zmq4"
)

const (
	WORKER_READY = "\001" //  Signals worker is ready
)

func main() {
	frontend, _ := zmq.NewSocket(zmq.ROUTER)
	backend, _ := zmq.NewSocket(zmq.ROUTER)
	defer frontend.Close()
	defer backend.Close()
	frontend.Bind("tcp://*:5555") //  For clients
	backend.Bind("tcp://*:5556")  //  For workers

	//  Queue of available workers
	workers := make([]string, 0)

	poller1 := zmq.NewPoller()
	poller1.Add(backend, zmq.POLLIN)
	poller2 := zmq.NewPoller()
	poller2.Add(backend, zmq.POLLIN)
	poller2.Add(frontend, zmq.POLLIN)

	//  The body of this example is exactly the same as lbbroker2.
LOOP:
	for {
		//  Poll frontend only if we have available workers
		var sockets []zmq.Polled
		var err error
		if len(workers) > 0 {
			sockets, err = poller2.Poll(-1)
		} else {
			sockets, err = poller1.Poll(-1)
		}
		if err != nil {
			break //  Interrupted
		}
		for _, socket := range sockets {
			switch s := socket.Socket; s {
			case backend: //  Handle worker activity on backend
				//  Use worker identity for load-balancing
				msg, err := s.RecvMessage(0)
				if err != nil {
					break LOOP //  Interrupted
				}
				var identity string
				identity, msg = unwrap(msg)
				workers = append(workers, identity)

				//  Forward message to client if it's not a READY
				if msg[0] != WORKER_READY {
					frontend.SendMessage(msg)
				}

			case frontend:
				//  Get client request, route to first available worker
				msg, err := s.RecvMessage(0)
				if err == nil {
					backend.SendMessage(workers[0], "", msg)
					workers = workers[1:]
				}
			}
		}
	}
}

//  Pops frame off front of message and returns it as 'head'
//  If next frame is empty, pops that empty frame.
//  Return remaining frames of message as 'tail'
func unwrap(msg []string) (head string, tail []string) {
	head = msg[0]
	if len(msg) > 1 && msg[1] == "" {
		tail = msg[2:]
	} else {
		tail = msg[1:]
	}
	return
}
