//
//  Clone client Model One
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"github.com/pebbe/zmq4/examples/kvsimple"

	"fmt"
)

func main() {
	//  Prepare our context and updates socket
	updates, _ := zmq.NewSocket(zmq.SUB)
	updates.SetSubscribe("")
	updates.Connect("tcp://localhost:5556")

	kvmap := make(map[string]*kvsimple.Kvmsg)

	sequence := int64(0)
	for ; true; sequence++ {
		kvmsg, err := kvsimple.RecvKvmsg(updates)
		if err != nil {
			break //  Interrupted
		}
		kvmsg.Store(kvmap)
	}
	fmt.Printf("Interrupted\n%d messages in\n", sequence)
}
