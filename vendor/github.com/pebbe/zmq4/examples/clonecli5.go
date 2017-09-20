//
//  Clone client Model Five
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"github.com/pebbe/zmq4/examples/kvmsg"

	"fmt"
	"math/rand"
	"time"
)

const (
	SUBTREE = "/client/"
)

func main() {
	snapshot, _ := zmq.NewSocket(zmq.DEALER)
	snapshot.Connect("tcp://localhost:5556")
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	subscriber.SetSubscribe(SUBTREE)
	subscriber.Connect("tcp://localhost:5557")
	publisher, _ := zmq.NewSocket(zmq.PUSH)
	publisher.Connect("tcp://localhost:5558")

	kvmap := make(map[string]*kvmsg.Kvmsg)
	rand.Seed(time.Now().UnixNano())

	//  We first request a state snapshot:
	sequence := int64(0)
	snapshot.SendMessage("ICANHAZ?", SUBTREE)
	for {
		kvmsg, err := kvmsg.RecvKvmsg(snapshot)
		if err != nil {
			break //  Interrupted
		}
		if key, _ := kvmsg.GetKey(); key == "KTHXBAI" {
			sequence, _ := kvmsg.GetSequence()
			fmt.Println("I: received snapshot =", sequence)
			break //  Done
		}
		kvmsg.Store(kvmap)
	}
	snapshot.Close()

	poller := zmq.NewPoller()
	poller.Add(subscriber, zmq.POLLIN)
	alarm := time.Now().Add(1000 * time.Millisecond)
	for {
		tickless := alarm.Sub(time.Now())
		if tickless < 0 {
			tickless = 0
		}
		polled, err := poller.Poll(tickless)
		if err != nil {
			break //  Context has been shut down
		}
		if len(polled) == 1 {
			kvmsg, err := kvmsg.RecvKvmsg(subscriber)
			if err != nil {
				break //  Interrupted
			}

			//  Discard out-of-sequence kvmsgs, incl. heartbeats
			if seq, _ := kvmsg.GetSequence(); seq > sequence {
				sequence = seq
				kvmsg.Store(kvmap)
				fmt.Println("I: received update =", sequence)
			}
		}
		//  If we timed-out, generate a random kvmsg
		if time.Now().After(alarm) {
			kvmsg := kvmsg.NewKvmsg(0)
			kvmsg.SetKey(fmt.Sprintf("%s%d", SUBTREE, rand.Intn(10000)))
			kvmsg.SetBody(fmt.Sprint(rand.Intn(1000000)))
			kvmsg.SetProp("ttl", fmt.Sprintf("%d", rand.Intn((30)))) // seconds
			kvmsg.Send(publisher)
			alarm = time.Now().Add(1000 * time.Millisecond)
		}
	}
	fmt.Printf("Interrupted\n%d messages in\n", sequence)
}
