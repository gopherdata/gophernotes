//
//  Reading from multiple sockets.
//  This version uses zmq.Poll()
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
)

func main() {

	//  Connect to task ventilator
	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5557")

	//  Connect to weather server
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5556")
	subscriber.SetSubscribe("10001 ")

	//  Initialize poll set
	poller := zmq.NewPoller()
	poller.Add(receiver, zmq.POLLIN)
	poller.Add(subscriber, zmq.POLLIN)
	//  Process messages from both sockets
	for {
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			switch s := socket.Socket; s {
			case receiver:
				task, _ := s.Recv(0)
				//  Process task
				fmt.Println("Got task:", task)
			case subscriber:
				update, _ := s.Recv(0)
				//  Process weather update
				fmt.Println("Got weather update:", update)
			}
		}
	}
}
