//
//  Clone server Model Three
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"github.com/pebbe/zmq4/examples/kvsimple"

	"fmt"
	"time"
)

func main() {
	snapshot, _ := zmq.NewSocket(zmq.ROUTER)
	snapshot.Bind("tcp://*:5556")
	publisher, _ := zmq.NewSocket(zmq.PUB)
	publisher.Bind("tcp://*:5557")
	collector, _ := zmq.NewSocket(zmq.PULL)
	collector.Bind("tcp://*:5558")

	//  The body of the main task collects updates from clients and
	//  publishes them back out to clients:

	sequence := int64(0)
	kvmap := make(map[string]*kvsimple.Kvmsg)

	poller := zmq.NewPoller()
	poller.Add(collector, zmq.POLLIN)
	poller.Add(snapshot, zmq.POLLIN)
LOOP:
	for {
		polled, err := poller.Poll(1000 * time.Millisecond)
		if err != nil {
			break
		}
		for _, item := range polled {
			switch socket := item.Socket; socket {
			case collector:
				//  Apply state update sent from client
				kvmsg, err := kvsimple.RecvKvmsg(collector)
				if err != nil {
					break LOOP //  Interrupted
				}
				sequence++
				kvmsg.SetSequence(sequence)
				kvmsg.Send(publisher)
				kvmsg.Store(kvmap)
				fmt.Println("I: publishing update", sequence)
			case snapshot:
				//  Execute state snapshot request
				msg, err := snapshot.RecvMessage(0)
				if err != nil {
					break LOOP
				}
				identity := msg[0]

				//  Request is in second frame of message
				request := msg[1]
				if request != "ICANHAZ?" {
					fmt.Println("E: bad request, aborting")
					break LOOP
				}
				//  Send state snapshot to client

				//  For each entry in kvmap, send kvmsg to client
				for _, kvmsg := range kvmap {
					snapshot.Send(identity, zmq.SNDMORE)
					kvmsg.Send(snapshot)
				}

				//  Now send END message with sequence number
				fmt.Println("I: sending shapshot =", sequence)
				snapshot.Send(identity, zmq.SNDMORE)
				kvmsg := kvsimple.NewKvmsg(sequence)
				kvmsg.SetKey("KTHXBAI")
				kvmsg.SetBody("")
				kvmsg.Send(snapshot)
			}
		}
	}
	fmt.Printf("Interrupted\n%d messages handled\n", sequence)
}
