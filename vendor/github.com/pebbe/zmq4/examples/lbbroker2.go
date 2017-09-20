//
//  Load-balancing broker.
//  Demonstrates use of higher level functions.
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"strings"
	"time"
)

const (
	NBR_CLIENTS  = 10
	NBR_WORKERS  = 3
	WORKER_READY = "\001" //  Signals worker is ready
)

//  Basic request-reply client using REQ socket
//
func client_task() {
	client, _ := zmq.NewSocket(zmq.REQ)
	defer client.Close()
	client.Connect("ipc://frontend.ipc")

	//  Send request, get reply
	for {
		client.SendMessage("HELLO")
		reply, _ := client.RecvMessage(0)
		if len(reply) == 0 {
			break
		}
		fmt.Println("Client:", strings.Join(reply, "\n\t"))
		time.Sleep(time.Second)
	}
}

//  Worker using REQ socket to do load-balancing
//
func worker_task() {
	worker, _ := zmq.NewSocket(zmq.REQ)
	defer worker.Close()
	worker.Connect("ipc://backend.ipc")

	//  Tell broker we're ready for work
	worker.SendMessage(WORKER_READY)

	//  Process messages as they arrive
	for {
		msg, e := worker.RecvMessage(0)
		if e != nil {
			break //  Interrupted ??
		}
		msg[len(msg)-1] = "OK"
		worker.SendMessage(msg)
	}
}

//  Now we come to the main task. This has the identical functionality to
//  the previous lbbroker example but uses higher level functions to read
//  and send messages:

func main() {
	//  Prepare our sockets
	frontend, _ := zmq.NewSocket(zmq.ROUTER)
	backend, _ := zmq.NewSocket(zmq.ROUTER)
	defer frontend.Close()
	defer backend.Close()
	frontend.Bind("ipc://frontend.ipc")
	backend.Bind("ipc://backend.ipc")

	for client_nbr := 0; client_nbr < NBR_CLIENTS; client_nbr++ {
		go client_task()
	}
	for worker_nbr := 0; worker_nbr < NBR_WORKERS; worker_nbr++ {
		go worker_task()
	}

	//  Queue of available workers
	workers := make([]string, 0, 10)

	poller1 := zmq.NewPoller()
	poller1.Add(backend, zmq.POLLIN)
	poller2 := zmq.NewPoller()
	poller2.Add(backend, zmq.POLLIN)
	poller2.Add(frontend, zmq.POLLIN)

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
			switch socket.Socket {
			case backend:
				//  Handle worker activity on backend

				//  Use worker identity for load-balancing
				msg, err := backend.RecvMessage(0)
				if err != nil {
					break LOOP //  Interrupted
				}
				identity, msg := unwrap(msg)
				workers = append(workers, identity)

				//  Forward message to client if it's not a READY
				if msg[0] != WORKER_READY {
					frontend.SendMessage(msg)
				}

			case frontend:
				//  Get client request, route to first available worker
				msg, err := frontend.RecvMessage(0)
				if err == nil {
					backend.SendMessage(workers[0], "", msg)
					workers = workers[1:]
				}
			}
		}
	}

	time.Sleep(100 * time.Millisecond)
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
