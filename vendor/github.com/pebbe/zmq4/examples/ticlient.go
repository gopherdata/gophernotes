//
//  Titanic client example.
//  Implements client side of http://rfc.zeromq.org/spec:9

package main

import (
	"github.com/pebbe/zmq4/examples/mdapi"

	"errors"
	"fmt"
	"os"
	"time"
)

//  Calls a TSP service
//  Returns response if successful (status code 200 OK), else NULL
//
func ServiceCall(session *mdapi.Mdcli, service string, request ...string) (reply []string, err error) {
	reply = []string{}
	msg, err := session.Send(service, request...)
	if err == nil {
		switch status := msg[0]; status {
		case "200":
			reply = msg[1:]
			return
		case "400":
			fmt.Println("E: client fatal error, aborting")
			os.Exit(1)
		case "500":
			fmt.Println("E: server fatal error, aborting")
			os.Exit(1)
		}
	} else {
		fmt.Println("E: " + err.Error())
		os.Exit(0)
	}

	err = errors.New("Didn't succeed")
	return //  Didn't succeed, don't care why not
}

//  The main task tests our service call by sending an echo request:

func main() {
	var verbose bool
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		verbose = true
	}
	session, _ := mdapi.NewMdcli("tcp://localhost:5555", verbose)

	//  1. Send 'echo' request to Titanic
	reply, err := ServiceCall(session, "titanic.request", "echo", "Hello world")
	if err != nil {
		fmt.Println(err)
		return
	}

	var uuid string
	if err == nil {
		uuid = reply[0]
		fmt.Println("I: request UUID", uuid)
	}

	time.Sleep(100 * time.Millisecond)

	//  2. Wait until we get a reply
	for {
		reply, err := ServiceCall(session, "titanic.reply", uuid)
		if err == nil {
			fmt.Println("Reply:", reply[0])

			//  3. Close request
			ServiceCall(session, "titanic.close", uuid)
			break
		} else {
			fmt.Println("I: no reply yet, trying again...")
			time.Sleep(5 * time.Second) //  Try again in 5 seconds
		}
	}
}
