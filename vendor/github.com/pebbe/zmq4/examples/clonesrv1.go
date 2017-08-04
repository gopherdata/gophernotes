//
//  Clone server Model One
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"github.com/pebbe/zmq4/examples/kvsimple"

	"fmt"
	"math/rand"
	"time"
)

func main() {
	//  Prepare our context and publisher socket
	publisher, _ := zmq.NewSocket(zmq.PUB)
	publisher.Bind("tcp://*:5556")
	time.Sleep(200 * time.Millisecond)

	kvmap := make(map[string]*kvsimple.Kvmsg)
	rand.Seed(time.Now().UnixNano())

	sequence := int64(1)
	for ; true; sequence++ {
		//  Distribute as key-value message
		kvmsg := kvsimple.NewKvmsg(sequence)
		kvmsg.SetKey(fmt.Sprint(rand.Intn(10000)))
		kvmsg.SetBody(fmt.Sprint(rand.Intn(1000000)))
		err := kvmsg.Send(publisher)
		kvmsg.Store(kvmap)
		if err != nil {
			break
		}
	}
	fmt.Printf("Interrupted\n%d messages out\n", sequence)
}
