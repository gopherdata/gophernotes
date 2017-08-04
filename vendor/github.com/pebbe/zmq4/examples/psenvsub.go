//
//  Pubsub envelope subscriber.
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
)

func main() {
	//  Prepare our subscriber
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5563")
	subscriber.SetSubscribe("B")

	for {
		//  Read envelope with address
		address, _ := subscriber.Recv(0)
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		fmt.Printf("[%s] %s\n", address, contents)
	}
}
