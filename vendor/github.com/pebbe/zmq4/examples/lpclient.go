//
//  Lazy Pirate client.
//  Use zmq_poll to do a safe request-reply
//  To run, start lpserver and then randomly kill/restart it
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"strconv"
	"time"
)

const (
	REQUEST_TIMEOUT = 2500 * time.Millisecond //  msecs, (> 1000!)
	REQUEST_RETRIES = 3                       //  Before we abandon
	SERVER_ENDPOINT = "tcp://localhost:5555"
)

func main() {
	fmt.Println("I: connecting to server...")
	client, err := zmq.NewSocket(zmq.REQ)
	if err != nil {
		panic(err)
	}
	client.Connect(SERVER_ENDPOINT)

	poller := zmq.NewPoller()
	poller.Add(client, zmq.POLLIN)

	sequence := 0
	retries_left := REQUEST_RETRIES
	for retries_left > 0 {
		//  We send a request, then we work to get a reply
		sequence++
		client.SendMessage(sequence)

		for expect_reply := true; expect_reply; {
			//  Poll socket for a reply, with timeout
			sockets, err := poller.Poll(REQUEST_TIMEOUT)
			if err != nil {
				break //  Interrupted
			}

			//  Here we process a server reply and exit our loop if the
			//  reply is valid. If we didn't a reply we close the client
			//  socket and resend the request. We try a number of times
			//  before finally abandoning:

			if len(sockets) > 0 {
				//  We got a reply from the server, must match sequence
				reply, err := client.RecvMessage(0)
				if err != nil {
					break //  Interrupted
				}
				seq, _ := strconv.Atoi(reply[0])
				if seq == sequence {
					fmt.Printf("I: server replied OK (%s)\n", reply[0])
					retries_left = REQUEST_RETRIES
					expect_reply = false
				} else {
					fmt.Printf("E: malformed reply from server: %s\n", reply)
				}
			} else {
				retries_left--
				if retries_left == 0 {
					fmt.Println("E: server seems to be offline, abandoning")
					break
				} else {
					fmt.Println("W: no response from server, retrying...")
					//  Old socket is confused; close it and open a new one
					client.Close()
					client, _ = zmq.NewSocket(zmq.REQ)
					client.Connect(SERVER_ENDPOINT)
					// Recreate poller for new client
					poller = zmq.NewPoller()
					poller.Add(client, zmq.POLLIN)
					//  Send request again, on new socket
					client.SendMessage(sequence)
				}
			}
		}
	}
	client.Close()
}
