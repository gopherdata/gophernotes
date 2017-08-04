//
//  Simple message queuing broker.
//  Same as request-reply broker but using QUEUE device
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"log"
)

func main() {
	var err error

	//  Socket facing clients
	frontend, _ := zmq.NewSocket(zmq.ROUTER)
	defer frontend.Close()
	err = frontend.Bind("tcp://*:5559")
	if err != nil {
		log.Fatalln("Binding frontend:", err)
	}

	//  Socket facing services
	backend, _ := zmq.NewSocket(zmq.DEALER)
	defer backend.Close()
	err = backend.Bind("tcp://*:5560")
	if err != nil {
		log.Fatalln("Binding backend:", err)
	}

	//  Start the proxy
	err = zmq.Proxy(frontend, backend, nil)
	log.Fatalln("Proxy interrupted:", err)
}
