//
//  Freelance server - Model 2.
//  Does some work, replies OK, with message sequencing
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

	fmt.Println("I: service is ready at", os.Args[1])
	for {
		request, err := server.RecvMessage(0)
		if err != nil {
			break //  Interrupted
		}
		//  Fail nastily if run against wrong client
		if len(request) != 2 {
			panic("len(request) != 2")
		}

		identity := request[0]

		server.SendMessage(identity, "OK")
	}
	fmt.Println("W: interrupted")
}
