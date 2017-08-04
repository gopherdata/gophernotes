//
//  Asynchronous client-to-server (DEALER to ROUTER).
//
//  While this example runs in a single process, that is just to make
//  it easier to start and stop the example. Each task has its own
//  context and conceptually acts as a separate process.

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

//  ---------------------------------------------------------------------
//  This is our client task
//  It connects to the server, and then sends a request once per second
//  It collects responses as they arrive, and it prints them out. We will
//  run several client tasks in parallel, each with a different random ID.

func client_task() {
	var mu sync.Mutex

	client, _ := zmq.NewSocket(zmq.DEALER)
	defer client.Close()

	//  Set random identity to make tracing easier
	set_id(client)
	client.Connect("tcp://localhost:5570")

	go func() {
		for request_nbr := 1; true; request_nbr++ {
			time.Sleep(time.Second)
			mu.Lock()
			client.SendMessage(fmt.Sprintf("request #%d", request_nbr))
			mu.Unlock()
		}
	}()

	for {
		time.Sleep(10 * time.Millisecond)
		mu.Lock()
		msg, err := client.RecvMessage(zmq.DONTWAIT)
		if err == nil {
			id, _ := client.GetIdentity()
			fmt.Println(msg[0], id)
		}
		mu.Unlock()
	}
}

//  This is our server task.
//  It uses the multithreaded server model to deal requests out to a pool
//  of workers and route replies back to clients. One worker can handle
//  one request at a time but one client can talk to multiple workers at
//  once.

func server_task() {

	//  Frontend socket talks to clients over TCP
	frontend, _ := zmq.NewSocket(zmq.ROUTER)
	defer frontend.Close()
	frontend.Bind("tcp://*:5570")

	//  Backend socket talks to workers over inproc
	backend, _ := zmq.NewSocket(zmq.DEALER)
	defer backend.Close()
	backend.Bind("inproc://backend")

	//  Launch pool of worker threads, precise number is not critical
	for i := 0; i < 5; i++ {
		go server_worker()
	}

	//  Connect backend to frontend via a proxy
	err := zmq.Proxy(frontend, backend, nil)
	log.Fatalln("Proxy interrupted:", err)
}

//  Each worker task works on one request at a time and sends a random number
//  of replies back, with random delays between replies:

func server_worker() {

	worker, _ := zmq.NewSocket(zmq.DEALER)
	defer worker.Close()
	worker.Connect("inproc://backend")

	for {
		//  The DEALER socket gives us the reply envelope and message
		msg, _ := worker.RecvMessage(0)
		identity, content := pop(msg)

		//  Send 0..4 replies back
		replies := rand.Intn(5)
		for reply := 0; reply < replies; reply++ {
			//  Sleep for some fraction of a second
			time.Sleep(time.Duration(rand.Intn(1000)+1) * time.Millisecond)
			worker.SendMessage(identity, content)
		}
	}
}

//  The main thread simply starts several clients, and a server, and then
//  waits for the server to finish.

func main() {
	rand.Seed(time.Now().UnixNano())

	go client_task()
	go client_task()
	go client_task()
	go server_task()

	//  Run for 5 seconds then quit
	time.Sleep(5 * time.Second)
}

func set_id(soc *zmq.Socket) {
	identity := fmt.Sprintf("%04X-%04X", rand.Intn(0x10000), rand.Intn(0x10000))
	soc.SetIdentity(identity)
}

func pop(msg []string) (head, tail []string) {
	if msg[1] == "" {
		head = msg[:2]
		tail = msg[2:]
	} else {
		head = msg[:1]
		tail = msg[1:]
	}
	return
}
