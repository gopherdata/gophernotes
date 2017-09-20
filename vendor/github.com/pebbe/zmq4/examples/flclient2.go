//
//  Freelance client - Model 2.
//  Uses DEALER socket to blast one or more services
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (

	//  If not a single service replies within this time, give up
	GLOBAL_TIMEOUT = 2500 * time.Millisecond
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("I: syntax: %s <endpoint> ...\n", os.Args[0])
		return
	}
	//  Create new freelance client object
	client := new_flclient()

	//  Connect to each endpoint
	for argn := 1; argn < len(os.Args); argn++ {
		client.connect(os.Args[argn])
	}

	//  Send a bunch of name resolution 'requests', measure time
	start := time.Now()
	for requests := 10000; requests > 0; requests-- {
		_, err := client.request("random name")
		if err != nil {
			fmt.Println("E: name service not available, aborting")
			break
		}
	}
	fmt.Println("Average round trip cost:", time.Now().Sub(start))
}

//  Here is the flclient class implementation. Each instance has
//  a DEALER socket it uses to talk to the servers, a counter of how many
//  servers it's connected to, and a request sequence number:

type flclient_t struct {
	socket   *zmq.Socket //  DEALER socket talking to servers
	servers  int         //  How many servers we have connected to
	sequence int         //  Number of requests ever sent
}

//  --------------------------------------------------------------------
//  Constructor

func new_flclient() (client *flclient_t) {
	client = &flclient_t{}

	client.socket, _ = zmq.NewSocket(zmq.DEALER)
	return
}

//  --------------------------------------------------------------------
//  Connect to new server endpoint

func (client *flclient_t) connect(endpoint string) {
	client.socket.Connect(endpoint)
	client.servers++
}

//  The request method does the hard work. It sends a request to all
//  connected servers in parallel (for this to work, all connections
//  have to be successful and completed by this time). It then waits
//  for a single successful reply, and returns that to the caller.
//  Any other replies are just dropped:

func (client *flclient_t) request(request ...string) (reply []string, err error) {
	reply = []string{}

	//  Prefix request with sequence number and empty envelope
	client.sequence++

	//  Blast the request to all connected servers
	for server := 0; server < client.servers; server++ {
		client.socket.SendMessage("", client.sequence, request)
	}
	//  Wait for a matching reply to arrive from anywhere
	//  Since we can poll several times, calculate each one
	endtime := time.Now().Add(GLOBAL_TIMEOUT)
	poller := zmq.NewPoller()
	poller.Add(client.socket, zmq.POLLIN)
	for time.Now().Before(endtime) {
		polled, err := poller.Poll(endtime.Sub(time.Now()))
		if err == nil && len(polled) > 0 {
			//  Reply is [empty][sequence][OK]
			reply, _ = client.socket.RecvMessage(0)
			if len(reply) != 3 {
				panic("len(reply) != 3")
			}
			sequence := reply[1]
			reply = reply[2:]
			sequence_nbr, _ := strconv.Atoi(sequence)
			if sequence_nbr == client.sequence {
				break
			}
		}
	}
	if len(reply) == 0 {
		err = errors.New("No reply")
	}
	return
}
