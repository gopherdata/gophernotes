//
//  Multithreaded Hello World server.
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"log"
	"time"
)

func worker_routine() {
	//  Socket to talk to dispatcher
	receiver, _ := zmq.NewSocket(zmq.REP)
	defer receiver.Close()
	receiver.Connect("inproc://workers")

	for {
		msg, e := receiver.Recv(0)
		if e != nil {
			break
		}
		fmt.Println("Received request: [" + msg + "]")

		//  Do some 'work'
		time.Sleep(time.Second)

		//  Send reply back to client
		receiver.Send("World", 0)
	}
}

func main() {
	//  Socket to talk to clients
	clients, _ := zmq.NewSocket(zmq.ROUTER)
	defer clients.Close()
	clients.Bind("tcp://*:5555")

	//  Socket to talk to workers
	workers, _ := zmq.NewSocket(zmq.DEALER)
	defer workers.Close()
	workers.Bind("inproc://workers")

	//  Launch pool of worker goroutines
	for thread_nbr := 0; thread_nbr < 5; thread_nbr++ {
		go worker_routine()
	}
	//  Connect work threads to client threads via a queue proxy
	err := zmq.Proxy(clients, workers, nil)
	log.Fatalln("Proxy interrupted:", err)
}
