//
//  Synchronized publisher.
//
//  This diverts from the C example by introducing time delays.
//  Without these delays, the subscribers won't catch the END message.
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"time"
)

const (
	//  We wait for 10 subscribers
	SUBSCRIBERS_EXPECTED = 10
)

func main() {

	//  Socket to talk to clients
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Bind("tcp://*:5561")

	//  Socket to receive signals
	syncservice, _ := zmq.NewSocket(zmq.REP)
	defer syncservice.Close()
	syncservice.Bind("tcp://*:5562")

	//  Get synchronization from subscribers
	fmt.Println("Waiting for subscribers")
	for subscribers := 0; subscribers < SUBSCRIBERS_EXPECTED; subscribers++ {
		//  - wait for synchronization request
		syncservice.Recv(0)
		//  - send synchronization reply
		syncservice.Send("", 0)
	}
	//  Now broadcast exactly 1M updates followed by END
	fmt.Println("Broadcasting messages")
	for update_nbr := 0; update_nbr < 1000000; update_nbr++ {
		publisher.Send("Rhubarb", 0)
		// subscribers don't get all messages if publisher is too fast
		// a one microsecond pause may still be too short
		time.Sleep(time.Microsecond)
	}

	// a longer pause ensures subscribers are ready to receive this
	time.Sleep(time.Second)
	publisher.Send("END", 0)

	// what's another second?
	time.Sleep(time.Second)
}
