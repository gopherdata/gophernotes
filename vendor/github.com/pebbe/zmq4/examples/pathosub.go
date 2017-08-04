//
//  Pathological subscriber
//  Subscribes to one random topic and prints received messages
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	if len(os.Args) == 2 {
		subscriber.Connect(os.Args[1])
	} else {
		subscriber.Connect("tcp://localhost:5556")
	}

	rand.Seed(time.Now().UnixNano())
	subscription := fmt.Sprintf("%03d", rand.Intn(1000))
	subscriber.SetSubscribe(subscription)

	for {
		msg, err := subscriber.RecvMessage(0)
		if err != nil {
			break
		}
		topic := msg[0]
		data := msg[1]
		if topic != subscription {
			panic("topic != subscription")
		}
		fmt.Println(data)
	}
}
