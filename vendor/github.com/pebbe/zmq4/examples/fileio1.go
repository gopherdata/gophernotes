//  File Transfer model #1
//
//  In which the server sends the entire file to the client in
//  large chunks with no attempt at flow control.

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"io"
	"os"
)

const (
	CHUNK_SIZE = 250000
)

func client_thread(pipe chan<- string) {
	dealer, _ := zmq.NewSocket(zmq.DEALER)
	dealer.Connect("tcp://127.0.0.1:6000")

	dealer.Send("fetch", 0)
	total := 0  //  Total bytes received
	chunks := 0 //  Total chunks received

	for {
		frame, err := dealer.RecvBytes(0)
		if err != nil {
			break //  Shutting down, quit
		}
		chunks++
		size := len(frame)
		total += size
		if size == 0 {
			break //  Whole file received
		}
	}
	fmt.Printf("%v chunks received, %v bytes\n", chunks, total)
	pipe <- "OK"
}

//  The server thread reads the file from disk in chunks, and sends
//  each chunk to the client as a separate message. We only have one
//  test file, so open that once and then serve it out as needed:

func server_thread() {
	file, err := os.Open("testdata")
	if err != nil {
		panic(err)
	}

	router, _ := zmq.NewSocket(zmq.ROUTER)
	//  Default HWM is 1000, which will drop messages here
	//  since we send more than 1,000 chunks of test data,
	//  so set an infinite HWM as a simple, stupid solution:
	router.SetRcvhwm(0)
	router.SetSndhwm(0)
	router.Bind("tcp://*:6000")
	for {
		//  First frame in each message is the sender identity
		identity, err := router.Recv(0)
		if err != nil {
			break //  Shutting down, quit
		}

		//  Second frame is "fetch" command
		command, _ := router.Recv(0)
		if command != "fetch" {
			panic("command != \"fetch\"")
		}

		chunk := make([]byte, CHUNK_SIZE)
		for {
			n, _ := io.ReadFull(file, chunk)
			router.SendMessage(identity, chunk[:n])
			if n == 0 {
				break //  Always end with a zero-size frame
			}
		}
	}
	file.Close()
}

//  The main task starts the client and server threads; it's easier
//  to test this as a single process with threads, than as multiple
//  processes:

func main() {
	pipe := make(chan string)

	//  Start child threads
	go server_thread()
	go client_thread(pipe)
	//  Loop until client tells us it's done
	<-pipe
}
