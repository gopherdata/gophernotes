//
//  Broker peering simulation (part 2).
//  Prototypes the request-reply flow
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	NBR_CLIENTS  = 10
	NBR_WORKERS  = 3
	WORKER_READY = "**READY**" //  Signals worker is ready
)

var (
	peers = make(map[string]bool)
)

//  The client task does a request-reply dialog using a standard
//  synchronous REQ socket:

func client_task(name string, i int) {
	clientname := fmt.Sprintf("Client-%s-%d", name, i)

	client, _ := zmq.NewSocket(zmq.REQ)
	defer client.Close()
	client.SetIdentity(clientname)
	client.Connect("ipc://" + name + "-localfe.ipc")

	for {
		//  Send request, get reply
		client.Send("HELLO from "+clientname, 0)
		reply, err := client.Recv(0)
		if err != nil {
			fmt.Println("client_task interrupted", name)
			break //  Interrupted
		}
		fmt.Printf("%s: %s\n", clientname, reply)
		time.Sleep(time.Duration(500+rand.Intn(1000)) * time.Millisecond)
	}
}

//  The worker task plugs into the load-balancer using a REQ
//  socket:

func worker_task(name string, i int) {
	workername := fmt.Sprintf("Worker-%s-%d", name, i)

	worker, _ := zmq.NewSocket(zmq.REQ)
	defer worker.Close()
	worker.SetIdentity(workername)
	worker.Connect("ipc://" + name + "-localbe.ipc")

	//  Tell broker we're ready for work
	worker.SendMessage(WORKER_READY)

	//  Process messages as they arrive
	for {
		msg, err := worker.RecvMessage(0)
		if err != nil {
			fmt.Println("worker_task interrupted", name)
			break //  Interrupted
		}

		i := len(msg) - 1
		fmt.Printf("%s: %s\n", workername, msg[i])
		worker.SendMessage(msg[:i], "OK from "+workername)
	}
}

//  The main task begins by setting-up its frontend and backend sockets
//  and then starting its client and worker tasks:

func main() {
	//  First argument is this broker's name
	//  Other arguments are our peers' names
	//
	if len(os.Args) < 2 {
		fmt.Println("syntax: peering2 me {you}...")
		os.Exit(1)
	}
	for _, peer := range os.Args[2:] {
		peers[peer] = true
	}

	self := os.Args[1]
	fmt.Println("I: preparing broker at", self)
	rand.Seed(time.Now().UnixNano())

	//  Bind cloud frontend to endpoint
	cloudfe, _ := zmq.NewSocket(zmq.ROUTER)
	defer cloudfe.Close()
	cloudfe.SetIdentity(self)
	cloudfe.Bind("ipc://" + self + "-cloud.ipc")

	//  Connect cloud backend to all peers
	cloudbe, _ := zmq.NewSocket(zmq.ROUTER)
	defer cloudbe.Close()
	cloudbe.SetIdentity(self)
	for _, peer := range os.Args[2:] {
		fmt.Println("I: connecting to cloud frontend at", peer)
		cloudbe.Connect("ipc://" + peer + "-cloud.ipc")
	}
	//  Prepare local frontend and backend
	localfe, _ := zmq.NewSocket(zmq.ROUTER)
	defer localfe.Close()
	localfe.Bind("ipc://" + self + "-localfe.ipc")
	localbe, _ := zmq.NewSocket(zmq.ROUTER)
	defer localbe.Close()
	localbe.Bind("ipc://" + self + "-localbe.ipc")

	//  Get user to tell us when we can start...
	fmt.Print("Press Enter when all brokers are started: ")
	var line string
	fmt.Scanln(&line)

	//  Start local workers
	for worker_nbr := 0; worker_nbr < NBR_WORKERS; worker_nbr++ {
		go worker_task(self, worker_nbr)
	}

	//  Start local clients
	for client_nbr := 0; client_nbr < NBR_CLIENTS; client_nbr++ {
		go client_task(self, client_nbr)
	}

	//  Here we handle the request-reply flow. We're using load-balancing
	//  to poll workers at all times, and clients only when there are one or
	//  more workers available.

	//  Least recently used queue of available workers
	workers := make([]string, 0)

	backends := zmq.NewPoller()
	backends.Add(localbe, zmq.POLLIN)
	backends.Add(cloudbe, zmq.POLLIN)
	frontends := zmq.NewPoller()
	frontends.Add(localfe, zmq.POLLIN)
	frontends.Add(cloudfe, zmq.POLLIN)

	msg := []string{}
	number_of_peers := len(os.Args) - 2

	for {
		//  First, route any waiting replies from workers
		//  If we have no workers anyhow, wait indefinitely
		timeout := time.Second
		if len(workers) == 0 {
			timeout = -1
		}
		sockets, err := backends.Poll(timeout)
		if err != nil {
			log.Println(err)
			break //  Interrupted
		}

		msg = msg[:]
		if socketInPolled(localbe, sockets) {
			//  Handle reply from local worker
			msg, err = localbe.RecvMessage(0)
			if err != nil {
				log.Println(err)
				break //  Interrupted
			}
			var identity string
			identity, msg = unwrap(msg)
			workers = append(workers, identity)

			//  If it's READY, don't route the message any further
			if msg[0] == WORKER_READY {
				msg = msg[0:0]
			}
		} else if socketInPolled(cloudbe, sockets) {
			//  Or handle reply from peer broker
			msg, err = cloudbe.RecvMessage(0)
			if err != nil {
				log.Println(err)
				break //  Interrupted
			}

			//  We don't use peer broker identity for anything
			_, msg = unwrap(msg)
		}

		if len(msg) > 0 {
			//  Route reply to cloud if it's addressed to a broker
			if peers[msg[0]] {
				cloudfe.SendMessage(msg)
			} else {
				localfe.SendMessage(msg)
			}
		}

		//  Now we route as many client requests as we have worker capacity
		//  for. We may reroute requests from our local frontend, but not from
		//  the cloud frontend. We reroute randomly now, just to test things
		//  out. In the next version we'll do this properly by calculating
		//  cloud capacity:

		for len(workers) > 0 {
			sockets, err := frontends.Poll(0)
			if err != nil {
				log.Println(err)
				break //  Interrupted
			}
			var reroutable bool
			//  We'll do peer brokers first, to prevent starvation
			if socketInPolled(cloudfe, sockets) {
				msg, _ = cloudfe.RecvMessage(0)
				reroutable = false
			} else if socketInPolled(localfe, sockets) {
				msg, _ = localfe.RecvMessage(0)
				reroutable = true
			} else {
				break //  No work, go back to backends
			}

			//  If reroutable, send to cloud 20% of the time
			//  Here we'd normally use cloud status information
			//
			if reroutable && number_of_peers > 0 && rand.Intn(5) == 0 {
				//  Route to random broker peer
				random_peer := os.Args[2+rand.Intn(number_of_peers)]
				cloudbe.SendMessage(random_peer, "", msg)
			} else {
				localbe.SendMessage(workers[0], "", msg)
				workers = workers[1:]
			}
		}
	}
	fmt.Println("Exit")
}

//  Pops frame off front of message and returns it as 'head'
//  If next frame is empty, pops that empty frame.
//  Return remaining frames of message as 'tail'
func unwrap(msg []string) (head string, tail []string) {
	head = msg[0]
	if len(msg) > 1 && msg[1] == "" {
		tail = msg[2:]
	} else {
		tail = msg[1:]
	}
	return
}

// Returns true if *Socket is in []Polled
func socketInPolled(s *zmq.Socket, p []zmq.Polled) bool {
	for _, pp := range p {
		if pp.Socket == s {
			return true
		}
	}
	return false
}
