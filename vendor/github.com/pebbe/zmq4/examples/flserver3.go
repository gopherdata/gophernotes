//
//  Freelance server - Model 3.
//  Uses an ROUTER/ROUTER socket but just one thread
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"os"
)

func main() {
	var verbose bool
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		verbose = true
	}

	//  Prepare server socket with predictable identity
	bind_endpoint := "tcp://*:5555"
	connect_endpoint := "tcp://localhost:5555"
	server, _ := zmq.NewSocket(zmq.ROUTER)
	server.SetIdentity(connect_endpoint)
	server.Bind(bind_endpoint)
	fmt.Println("I: service is ready at", bind_endpoint)

	for {
		request, err := server.RecvMessage(0)
		if err != nil {
			break
		}
		if verbose {
			fmt.Printf("%q\n", request)
		}

		//  Frame 0: identity of client
		//  Frame 1: PING, or client control frame
		//  Frame 2: request body
		identity := request[0]
		control := request[1]
		reply := make([]string, 1, 3)
		if control == "PING" {
			reply = append(reply, "PONG")
		} else {
			reply = append(reply, control)
			reply = append(reply, "OK")
		}
		reply[0] = identity
		if verbose {
			fmt.Printf("%q\n", reply)
		}
		server.SendMessage(reply)
	}
	fmt.Println("W: interrupted")
}
