//
//  Majordomo Protocol worker example.
//  Uses the mdwrk API to hide all MDP aspects
//

package main

import (
	"github.com/pebbe/zmq4/examples/mdapi"

	"log"
	"os"
)

func main() {
	var verbose bool
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		verbose = true
	}
	session, _ := mdapi.NewMdwrk("tcp://localhost:5555", "echo", verbose)

	var err error
	var request, reply []string
	for {
		request, err = session.Recv(reply)
		if err != nil {
			break //  Worker was interrupted
		}
		reply = request //  Echo is complex... :-)
	}
	log.Println(err)
}
