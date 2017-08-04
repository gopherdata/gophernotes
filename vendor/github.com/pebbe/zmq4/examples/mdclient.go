//
//  Majordomo Protocol client example.
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
	session, _ := mdapi.NewMdcli("tcp://localhost:5555", verbose)

	count := 0
	for ; count < 100000; count++ {
		_, err := session.Send("echo", "Hello world")
		if err != nil {
			log.Println(err)
			break //  Interrupt or failure
		}
	}
	fmt.Printf("%d requests/replies processed\n", count)
}
