//
//  Clone client Model Two
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
	"time"
)

func main() {
	snapshot, _ := zmq.NewSocket(zmq.DEALER)
	snapshot.Connect("tcp://localhost:5556")

	subscriber, _ := zmq.NewSocket(zmq.SUB)
	subscriber.SetRcvhwm(100000) // or messages between snapshot and next are lost
	subscriber.SetSubscribe("")
	subscriber.Connect("tcp://localhost:5557")

	time.Sleep(time.Second) // or messages between snapshot and next are lost

	kvmap := make(map[string]*kvsimple.Kvmsg)

	//  Get state snapshot
	sequence := int64(0)
	snapshot.SendMessage("ICANHAZ?")
	for {
		kvmsg, err := kvsimple.RecvKvmsg(snapshot)
		if err != nil {
			fmt.Println(err)
			break //  Interrupted
		}
		if key, _ := kvmsg.GetKey(); key == "KTHXBAI" {
			sequence, _ = kvmsg.GetSequence()
			fmt.Printf("Received snapshot=%d\n", sequence)
			break //  Done
		}
		kvmsg.Store(kvmap)
	}
	snapshot.Close()

	first := true
	//  Now apply pending updates, discard out-of-sequence messages
	for {
		kvmsg, err := kvsimple.RecvKvmsg(subscriber)
		if err != nil {
			fmt.Println(err)
			break //  Interrupted
		}
		if seq, _ := kvmsg.GetSequence(); seq > sequence {
			sequence, _ = kvmsg.GetSequence()
			kvmsg.Store(kvmap)
			if first {
				// Show what the first regular update is after the snapshot,
				// to see if we missed updates.
				first = false
				fmt.Println("Next:", sequence)
			}
		}
	}
}
