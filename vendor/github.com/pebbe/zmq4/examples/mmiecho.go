//
//  MMI echo query example.
//

package main

import (
	"github.com/pebbe/zmq4/examples/mdapi"

	"fmt"
	"os"
)

func main() {
	var verbose bool
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		verbose = true
	}
	session, _ := mdapi.NewMdcli("tcp://localhost:5555", verbose)

	//  This is the service we want to look up
	request := "echo"

	//  This is the service we send our request to
	reply, err := session.Send("mmi.service", request)

	if err == nil {
		fmt.Println("Lookup echo service:", reply[0])
	} else {
		fmt.Println("E: no response from broker, make sure it's running")
	}
}
