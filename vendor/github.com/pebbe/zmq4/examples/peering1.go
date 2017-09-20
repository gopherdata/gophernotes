//
//  Broker peering simulation (part 1).
//  Prototypes the state flow
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	//  First argument is this broker's name
	//  Other arguments are our peers' names
	//
	if len(os.Args) < 2 {
		fmt.Println("syntax: peering1 me {you}...")
		os.Exit(1)
	}
	self := os.Args[1]
	fmt.Printf("I: preparing broker at %s...\n", self)
	rand.Seed(time.Now().UnixNano())

	//  Bind state backend to endpoint
	statebe, _ := zmq.NewSocket(zmq.PUB)
	defer statebe.Close()
	statebe.Bind("ipc://" + self + "-state.ipc")

	//  Connect statefe to all peers
	statefe, _ := zmq.NewSocket(zmq.SUB)
	defer statefe.Close()
	statefe.SetSubscribe("")
	for _, peer := range os.Args[2:] {
		fmt.Printf("I: connecting to state backend at '%s'\n", peer)
		statefe.Connect("ipc://" + peer + "-state.ipc")
	}

	//  The main loop sends out status messages to peers, and collects
	//  status messages back from peers. The zmq_poll timeout defines
	//  our own heartbeat:

	poller := zmq.NewPoller()
	poller.Add(statefe, zmq.POLLIN)
	for {
		//  Poll for activity, or 1 second timeout
		sockets, err := poller.Poll(time.Second)
		if err != nil {
			break
		}

		//  Handle incoming status messages
		if len(sockets) == 1 {
			msg, _ := statefe.RecvMessage(0)
			peer_name := msg[0]
			available := msg[1]
			fmt.Printf("%s - %s workers free\n", peer_name, available)
		} else {
			statebe.SendMessage(self, rand.Intn(10))
		}
	}
}
