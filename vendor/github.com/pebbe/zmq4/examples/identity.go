//
//  Demonstrate identities as used by the request-reply pattern.
//  Run this program by itself.
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"regexp"
)

var (
	all_char = regexp.MustCompile("^[^[:cntrl:]]*$")
)

func main() {
	sink, _ := zmq.NewSocket(zmq.ROUTER)
	defer sink.Close()
	sink.Bind("inproc://example")

	//  First allow 0MQ to set the identity
	anonymous, _ := zmq.NewSocket(zmq.REQ)
	defer anonymous.Close()
	anonymous.Connect("inproc://example")
	anonymous.Send("ROUTER uses a generated UUID", 0)
	dump(sink)

	//  Then set the identity ourselves
	identified, _ := zmq.NewSocket(zmq.REQ)
	defer identified.Close()
	identified.SetIdentity("PEER2")
	identified.Connect("inproc://example")
	identified.Send("ROUTER socket uses REQ's socket identity", 0)
	dump(sink)
}

func dump(soc *zmq.Socket) {
	fmt.Println("----------------------------------------")
	for {
		//  Process all parts of the message
		message, _ := soc.Recv(0)

		//  Dump the message as text or binary
		fmt.Printf("[%03d] ", len(message))
		if all_char.MatchString(message) {
			fmt.Print(message)
		} else {
			for i := 0; i < len(message); i++ {
				fmt.Printf("%02X ", message[i])
			}
		}
		fmt.Println()

		more, _ := soc.GetRcvmore()
		if !more {
			break
		}
	}
}
