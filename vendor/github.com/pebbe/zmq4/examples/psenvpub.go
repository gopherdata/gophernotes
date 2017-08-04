//
//  Pubsub envelope publisher.
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"time"
)

func main() {
	//  Prepare our publisher
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Bind("tcp://*:5563")

	for {
		//  Write two messages, each with an envelope and content
		publisher.Send("A", zmq.SNDMORE)
		publisher.Send("We don't want to see this", 0)
		publisher.Send("B", zmq.SNDMORE)
		publisher.Send("We would like to see this", 0)
		time.Sleep(time.Second)
	}
}
