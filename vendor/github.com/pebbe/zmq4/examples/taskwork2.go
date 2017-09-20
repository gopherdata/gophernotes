//
//  Task worker - design 2.
//  Adds pub-sub flow to receive and respond to kill signal
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"strconv"
	"time"
)

func main() {
	//  Socket to receive messages on
	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5557")

	//  Socket to send messages to
	sender, _ := zmq.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Connect("tcp://localhost:5558")

	//  Socket for control input
	controller, _ := zmq.NewSocket(zmq.SUB)
	defer controller.Close()
	controller.Connect("tcp://localhost:5559")
	controller.SetSubscribe("")

	//  Process messages from receiver and controller
	poller := zmq.NewPoller()
	poller.Add(receiver, zmq.POLLIN)
	poller.Add(controller, zmq.POLLIN)
	//  Process messages from both sockets
LOOP:
	for {
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			switch s := socket.Socket; s {
			case receiver:
				msg, _ := s.Recv(0)

				//  Do the work
				t, _ := strconv.Atoi(msg)
				time.Sleep(time.Duration(t) * time.Millisecond)

				//  Send results to sink
				sender.Send(msg, 0)

				//  Simple progress indicator for the viewer
				fmt.Printf(".")
			case controller:
				//  Any controller command acts as 'KILL'
				break LOOP //  Exit loop
			}
		}
	}
	fmt.Println()
}
