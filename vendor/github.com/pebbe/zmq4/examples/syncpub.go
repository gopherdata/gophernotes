//
//  Synchronized publisher.
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
)

const (
	//  We wait for 10 subscribers
	SUBSCRIBERS_EXPECTED = 10
)

func main() {

	ctx, _ := zmq.NewContext()
	defer ctx.Term()

	//  Socket to talk to clients
	publisher, _ := ctx.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.SetSndhwm(1100000)
	publisher.Bind("tcp://*:5561")

	//  Socket to receive signals
	syncservice, _ := ctx.NewSocket(zmq.REP)
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
	}

	publisher.Send("END", 0)

}
