//
//  Binary Star client proof-of-concept implementation. This client does no
//  real work; it just demonstrates the Binary Star failover model.

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"strconv"
	"time"
)

const (
	REQUEST_TIMEOUT = 1000 * time.Millisecond //  msecs
	SETTLE_DELAY    = 2000 * time.Millisecond //  Before failing over
)

func main() {

	server := []string{"tcp://localhost:5001", "tcp://localhost:5002"}
	server_nbr := 0

	fmt.Printf("I: connecting to server at %s...\n", server[server_nbr])
	client, _ := zmq.NewSocket(zmq.REQ)
	client.Connect(server[server_nbr])

	poller := zmq.NewPoller()
	poller.Add(client, zmq.POLLIN)

	sequence := 0
LOOP:
	for {
		//  We send a request, then we work to get a reply
		sequence++
		client.SendMessage(sequence)

		for expect_reply := true; expect_reply; {
			//  Poll socket for a reply, with timeout
			polled, err := poller.Poll(REQUEST_TIMEOUT)
			if err != nil {
				break LOOP //  Interrupted
			}

			//  We use a Lazy Pirate strategy in the client. If there's no
			//  reply within our timeout, we close the socket and try again.
			//  In Binary Star, it's the client vote which decides which
			//  server is primary; the client must therefore try to connect
			//  to each server in turn:

			if len(polled) == 1 {
				//  We got a reply from the server, must match sequence
				reply, _ := client.RecvMessage(0)
				seq, _ := strconv.Atoi(reply[0])
				if seq == sequence {
					fmt.Printf("I: server replied OK (%s)\n", reply[0])
					expect_reply = false
					time.Sleep(time.Second) //  One request per second
				} else {
					fmt.Printf("E: bad reply from server: %q\n", reply)
				}

			} else {
				fmt.Println("W: no response from server, failing over")

				//  Old socket is confused; close it and open a new one
				client.Close()
				server_nbr = 1 - server_nbr
				time.Sleep(SETTLE_DELAY)
				fmt.Printf("I: connecting to server at %s...\n", server[server_nbr])
				client, _ = zmq.NewSocket(zmq.REQ)
				client.Connect(server[server_nbr])

				poller = zmq.NewPoller()
				poller.Add(client, zmq.POLLIN)

				//  Send request again, on new socket
				client.SendMessage(sequence)
			}
		}
	}
}
