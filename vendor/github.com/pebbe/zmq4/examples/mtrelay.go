//
//  Multithreaded relay.
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
)

func step1() {
	//  Connect to step2 and tell it we're ready
	xmitter, _ := zmq.NewSocket(zmq.PAIR)
	defer xmitter.Close()
	xmitter.Connect("inproc://step2")
	fmt.Println("Step 1 ready, signaling step 2")
	xmitter.Send("READY", 0)
}

func step2() {
	//  Bind inproc socket before starting step1
	receiver, _ := zmq.NewSocket(zmq.PAIR)
	defer receiver.Close()
	receiver.Bind("inproc://step2")
	go step1()

	//  Wait for signal and pass it on
	receiver.Recv(0)

	//  Connect to step3 and tell it we're ready
	xmitter, _ := zmq.NewSocket(zmq.PAIR)
	defer xmitter.Close()
	xmitter.Connect("inproc://step3")
	fmt.Println("Step 2 ready, signaling step 3")
	xmitter.Send("READY", 0)
}

func main() {

	//  Bind inproc socket before starting step2
	receiver, _ := zmq.NewSocket(zmq.PAIR)
	defer receiver.Close()
	receiver.Bind("inproc://step3")
	go step2()

	//  Wait for signal
	receiver.Recv(0)

	fmt.Println("Test successful!")
}
