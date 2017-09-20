//
//  Round-trip demonstrator.
//
//  While this example runs in a single process, that is just to make
//  it easier to start and stop the example. The client task signals to
//  main when it's ready.
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"time"
)

func ClientTask(pipe chan<- bool) {
	client, _ := zmq.NewSocket(zmq.DEALER)
	client.Connect("tcp://localhost:5555")
	fmt.Println("Setting up test...")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Synchronous round-trip test...")
	start := time.Now()
	var requests int
	for requests = 0; requests < 10000; requests++ {
		client.Send("hello", 0)
		client.Recv(0)
	}
	fmt.Println(requests, "calls in", time.Since(start))

	fmt.Println("Asynchronous round-trip test...")
	start = time.Now()
	for requests = 0; requests < 100000; requests++ {
		client.Send("hello", 0)
	}
	for requests = 0; requests < 100000; requests++ {
		client.Recv(0)
	}
	fmt.Println(requests, "calls in", time.Since(start))
	pipe <- true
}

//  Here is the worker task. All it does is receive a message, and
//  bounce it back the way it came:

func WorkerTask() {
	worker, _ := zmq.NewSocket(zmq.DEALER)
	worker.Connect("tcp://localhost:5556")

	for {
		msg, _ := worker.RecvMessage(0)
		worker.SendMessage(msg)
	}
}

//  Here is the broker task. It uses the zmq_proxy function to switch
//  messages between frontend and backend:

func BrokerTask() {
	//  Prepare our sockets
	frontend, _ := zmq.NewSocket(zmq.DEALER)
	frontend.Bind("tcp://*:5555")
	backend, _ := zmq.NewSocket(zmq.DEALER)
	backend.Bind("tcp://*:5556")
	zmq.Proxy(frontend, backend, nil)
}

//  Finally, here's the main task, which starts the client, worker, and
//  broker, and then runs until the client signals it to stop:

func main() {
	//  Create threads
	pipe := make(chan bool)
	go ClientTask(pipe)
	go WorkerTask()
	go BrokerTask()

	//  Wait for signal on client pipe
	<-pipe
}
