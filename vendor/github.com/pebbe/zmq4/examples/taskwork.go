//
//  Task worker.
//  Connects PULL socket to tcp://localhost:5557
//  Collects workloads from ventilator via that socket
//  Connects PUSH socket to tcp://localhost:5558
//  Sends results to sink via that socket
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

	//  Process tasks forever
	for {
		s, _ := receiver.Recv(0)

		//  Simple progress indicator for the viewer
		fmt.Print(s + ".")

		//  Do the work
		msec, _ := strconv.Atoi(s)
		time.Sleep(time.Duration(msec) * time.Millisecond)

		//  Send results to sink
		sender.Send("", 0)
	}
}
