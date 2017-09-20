//
//  Clone server Model Two
//
//  In the original C example, the client misses updates between snapshot
//  and further updates. Sometimes, it even misses the END message of
//  the snapshot, so it waits for it forever.
//  This Go implementation has some modifications to improve this, but it
//  is still not fully reliable.

package main

import (
	zmq "github.com/pebbe/zmq4"
	"github.com/pebbe/zmq4/examples/kvsimple"

	"fmt"
	"math/rand"
	"time"
)

func main() {
	//  Prepare our context and sockets
	publisher, _ := zmq.NewSocket(zmq.PUB)
	publisher.Bind("tcp://*:5557")

	sequence := int64(0)
	rand.Seed(time.Now().UnixNano())

	//  Start state manager and wait for synchronization signal
	updates, _ := zmq.NewSocket(zmq.PAIR)
	updates.Bind("inproc://pipe")
	go state_manager()
	updates.RecvMessage(0) // "READY"

	for {
		//  Distribute as key-value message
		sequence++
		kvmsg := kvsimple.NewKvmsg(sequence)
		kvmsg.SetKey(fmt.Sprint(rand.Intn(10000)))
		kvmsg.SetBody(fmt.Sprint(rand.Intn(1000000)))
		if kvmsg.Send(publisher) != nil {
			break
		}
		if kvmsg.Send(updates) != nil {
			break
		}
	}
	fmt.Printf("Interrupted\n%d messages out\n", sequence)
}

//  The state manager task maintains the state and handles requests from
//  clients for snapshots:

func state_manager() {
	kvmap := make(map[string]*kvsimple.Kvmsg)

	pipe, _ := zmq.NewSocket(zmq.PAIR)
	pipe.Connect("inproc://pipe")
	pipe.SendMessage("READY")
	snapshot, _ := zmq.NewSocket(zmq.ROUTER)
	snapshot.Bind("tcp://*:5556")

	poller := zmq.NewPoller()
	poller.Add(pipe, zmq.POLLIN)
	poller.Add(snapshot, zmq.POLLIN)
	sequence := int64(0) //  Current snapshot version number
LOOP:
	for {
		polled, err := poller.Poll(-1)
		if err != nil {
			break //  Context has been shut down
		}
		for _, item := range polled {
			switch socket := item.Socket; socket {
			case pipe:
				//  Apply state update from main thread
				kvmsg, err := kvsimple.RecvKvmsg(pipe)
				if err != nil {
					break LOOP //  Interrupted
				}
				sequence, _ = kvmsg.GetSequence()
				kvmsg.Store(kvmap)
			case snapshot:
				//  Execute state snapshot request
				msg, err := snapshot.RecvMessage(0)
				if err != nil {
					break LOOP //  Interrupted
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

				// Give client some time to deal with it.
				// This reduces the risk that the client won't see
				// the END message, but it doesn't eliminate the risk.
				time.Sleep(100 * time.Millisecond)

				//  Now send END message with sequence number
				fmt.Printf("Sending state shapshot=%d\n", sequence)
				snapshot.Send(identity, zmq.SNDMORE)
				kvmsg := kvsimple.NewKvmsg(sequence)
				kvmsg.SetKey("KTHXBAI")
				kvmsg.SetBody("")
				kvmsg.Send(snapshot)
			}
		}
	}
}
