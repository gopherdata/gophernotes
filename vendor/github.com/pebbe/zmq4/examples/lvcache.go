//
//  Last value cache
//  Uses XPUB subscription messages to re-send data
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"time"
)

func main() {
	frontend, _ := zmq.NewSocket(zmq.SUB)
	frontend.Bind("tcp://*:5557")
	backend, _ := zmq.NewSocket(zmq.XPUB)
	backend.Bind("tcp://*:5558")

	//  Subscribe to every single topic from publisher
	frontend.SetSubscribe("")

	//  Store last instance of each topic in a cache
	cache := make(map[string]string)

	//  We route topic updates from frontend to backend, and
	//  we handle subscriptions by sending whatever we cached,
	//  if anything:
	poller := zmq.NewPoller()
	poller.Add(frontend, zmq.POLLIN)
	poller.Add(backend, zmq.POLLIN)
LOOP:
	for {
		polled, err := poller.Poll(1000 * time.Millisecond)
		if err != nil {
			break //  Interrupted
		}

		for _, item := range polled {
			switch socket := item.Socket; socket {
			case frontend:
				//  Any new topic data we cache and then forward
				msg, err := frontend.RecvMessage(0)
				if err != nil {
					break LOOP
				}
				cache[msg[0]] = msg[1]
				backend.SendMessage(msg)
			case backend:
				//  When we get a new subscription we pull data from the cache:
				msg, err := backend.RecvMessage(0)
				if err != nil {
					break LOOP
				}
				frame := msg[0]
				//  Event is one byte 0=unsub or 1=sub, followed by topic
				if frame[0] == 1 {
					topic := frame[1:]
					fmt.Println("Sending cached topic", topic)
					previous, ok := cache[topic]
					if ok {
						backend.SendMessage(topic, previous)
					}
				}
			}
		}
	}
}
