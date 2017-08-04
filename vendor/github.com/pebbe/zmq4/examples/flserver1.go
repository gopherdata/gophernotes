//
//  Freelance server - Model 1.
//  Trivial echo service
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("I: syntax: %s <endpoint>\n", os.Args[0])
		return
	}
	server, _ := zmq.NewSocket(zmq.REP)
	server.Bind(os.Args[1])

	fmt.Println("I: echo service is ready at", os.Args[1])
	for {
		msg, err := server.RecvMessage(0)
		if err != nil {
			break //  Interrupted
		}
		server.SendMessage(msg)
	}
	fmt.Println("W: interrupted")
}
