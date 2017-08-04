//
//  Espresso Pattern
//  This shows how to capture data using a pub-sub proxy
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"math/rand"
	"time"
)

//  The subscriber thread requests messages starting with
//  A and B, then reads and counts incoming messages.

func subscriber_thread() {
	//  Subscribe to "A" and "B"
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	subscriber.Connect("tcp://localhost:6001")
	subscriber.SetSubscribe("A")
	subscriber.SetSubscribe("B")
	defer subscriber.Close() // cancel subscribe

	for count := 0; count < 5; count++ {
		_, err := subscriber.RecvMessage(0)
		if err != nil {
			break //  Interrupted
		}
	}
}

//  The publisher sends random messages starting with A-J:

func publisher_thread() {
	publisher, _ := zmq.NewSocket(zmq.PUB)
	publisher.Bind("tcp://*:6000")

	for {
		s := fmt.Sprintf("%c-%05d", rand.Intn(10)+'A', rand.Intn(100000))
		_, err := publisher.SendMessage(s)
		if err != nil {
			break //  Interrupted
		}
		time.Sleep(100 * time.Millisecond) //  Wait for 1/10th second
	}
}

//  The listener receives all messages flowing through the proxy, on its
//  pipe. In CZMQ, the pipe is a pair of ZMQ_PAIR sockets that connects
//  attached child threads. In other languages your mileage may vary:

func listener_thread() {
	pipe, _ := zmq.NewSocket(zmq.PAIR)
	pipe.Bind("inproc://pipe")

	//  Print everything that arrives on pipe
	for {
		msg, err := pipe.RecvMessage(0)
		if err != nil {
			break //  Interrupted
		}
		fmt.Printf("%q\n", msg)
	}
}

//  The main task starts the subscriber and publisher, and then sets
//  itself up as a listening proxy. The listener runs as a child thread:

func main() {
	//  Start child threads
	go publisher_thread()
	go subscriber_thread()
	go listener_thread()

	time.Sleep(100 * time.Millisecond)

	subscriber, _ := zmq.NewSocket(zmq.XSUB)
	subscriber.Connect("tcp://localhost:6000")
	publisher, _ := zmq.NewSocket(zmq.XPUB)
	publisher.Bind("tcp://*:6001")
	listener, _ := zmq.NewSocket(zmq.PAIR)
	listener.Connect("inproc://pipe")
	zmq.Proxy(subscriber, publisher, listener)

	fmt.Println("interrupted")
}
