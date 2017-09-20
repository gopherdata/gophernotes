//
//  Majordomo Protocol client example - asynchronous.
//  Uses the mdcli API to hide all MDP aspects
//

package main

import (
	"github.com/pebbe/zmq4/examples/mdapi"

	"fmt"
	"log"
	"os"
)

func main() {
	var verbose bool
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		verbose = true
	}
	session, _ := mdapi.NewMdcli2("tcp://localhost:5555", verbose)

	var count int
	for count = 0; count < 100000; count++ {
		err := session.Send("echo", "Hello world")
		if err != nil {
			log.Println("Send:", err)
			break
		}
	}
	for count = 0; count < 100000; count++ {
		_, err := session.Recv()
		if err != nil {
			log.Println("Recv:", err)
			break
		}
	}
	fmt.Printf("%d replies received\n", count)
}
