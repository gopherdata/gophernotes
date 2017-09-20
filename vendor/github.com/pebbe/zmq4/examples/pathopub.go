//
//  Pathological publisher
//  Sends out 1,000 topics and then one random update per second
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
	publisher, _ := zmq.NewSocket(zmq.PUB)
	if len(os.Args) == 2 {
		publisher.Connect(os.Args[1])
	} else {
		publisher.Bind("tcp://*:5556")
	}

	//  Ensure subscriber connection has time to complete
	time.Sleep(time.Second)

	//  Send out all 1,000 topic messages
	for topic_nbr := 0; topic_nbr < 1000; topic_nbr++ {
		_, err := publisher.SendMessage(fmt.Sprintf("%03d", topic_nbr), "Save Roger")
		if err != nil {
			fmt.Println(err)
		}
	}
	//  Send one random update per second
	rand.Seed(time.Now().UnixNano())
	for {
		time.Sleep(time.Second)
		_, err := publisher.SendMessage(fmt.Sprintf("%03d", rand.Intn(1000)), "Off with his head!")
		if err != nil {
			fmt.Println(err)
		}
	}
}
