//
//  Freelance client - Model 1.
//  Uses REQ socket to query one or more services
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"errors"
	"fmt"
	"os"
	"time"
)

const (
	REQUEST_TIMEOUT = 1000 * time.Millisecond
	MAX_RETRIES     = 3 //  Before we abandon
)

func try_request(endpoint string, request []string) (reply []string, err error) {
	fmt.Printf("I: trying echo service at %s...\n", endpoint)
	client, _ := zmq.NewSocket(zmq.REQ)
	client.Connect(endpoint)

	//  Send request, wait safely for reply
	client.SendMessage(request)
	poller := zmq.NewPoller()
	poller.Add(client, zmq.POLLIN)
	polled, err := poller.Poll(REQUEST_TIMEOUT)
	reply = []string{}
	if len(polled) == 1 {
		reply, err = client.RecvMessage(0)
	} else {
		err = errors.New("Time out")
	}
	return
}

//  The client uses a Lazy Pirate strategy if it only has one server to talk
//  to. If it has 2 or more servers to talk to, it will try each server just
//  once:

func main() {
	request := []string{"Hello world"}
	reply := []string{}
	var err error

	endpoints := len(os.Args) - 1
	if endpoints == 0 {
		fmt.Printf("I: syntax: %s <endpoint> ...\n", os.Args[0])
	} else if endpoints == 1 {
		//  For one endpoint, we retry N times
		for retries := 0; retries < MAX_RETRIES; retries++ {
			endpoint := os.Args[1]
			reply, err = try_request(endpoint, request)
			if err == nil {
				break //  Successful
			}
			fmt.Printf("W: no response from %s, retrying...\n", endpoint)
		}
	} else {
		//  For multiple endpoints, try each at most once
		for endpoint_nbr := 0; endpoint_nbr < endpoints; endpoint_nbr++ {
			endpoint := os.Args[endpoint_nbr+1]
			reply, err = try_request(endpoint, request)
			if err == nil {
				break //  Successful
			}
			fmt.Println("W: no response from", endpoint)
		}
	}
	if len(reply) > 0 {
		fmt.Printf("Service is running OK: %q\n", reply)
	}
}
