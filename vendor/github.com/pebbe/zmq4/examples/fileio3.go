//  File Transfer model #3
//
//  In which the client requests each chunk individually, using
//  command pipelining to give us a credit-based flow control.

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"os"
	"strconv"
)

const (
	CHUNK_SIZE = 250000
	PIPELINE   = 10
)

func client_thread(pipe chan<- string) {
	dealer, _ := zmq.NewSocket(zmq.DEALER)
	dealer.Connect("tcp://127.0.0.1:6000")

	//  Up to this many chunks in transit
	credit := PIPELINE

	total := 0  //  Total bytes received
	chunks := 0 //  Total chunks received
	offset := 0 //  Offset of next chunk request

	for {
		for credit > 0 {
			//  Ask for next chunk
			dealer.SendMessage("fetch", offset, CHUNK_SIZE)
			offset += CHUNK_SIZE
			credit--
		}
		chunk, err := dealer.RecvBytes(0)
		if err != nil {
			break //  Shutting down, quit
		}
		chunks++
		credit++
		size := len(chunk)
		total += size
		if size < CHUNK_SIZE {
			break //  Last chunk received; exit
		}
	}
	fmt.Printf("%v chunks received, %v bytes\n", chunks, total)
	pipe <- "OK"
}

//  The rest of the code is exactly the same as in model 2, except
//  that we set the HWM on the server's ROUTER socket to PIPELINE
//  to act as a sanity check.

//  The server thread waits for a chunk request from a client,
//  reads that chunk and sends it back to the client:

func server_thread() {
	file, err := os.Open("testdata")
	if err != nil {
		panic(err)
	}

	router, _ := zmq.NewSocket(zmq.ROUTER)
	router.SetRcvhwm(PIPELINE * 2)
	router.SetSndhwm(PIPELINE * 2)
	router.Bind("tcp://*:6000")
	for {
		msg, err := router.RecvMessage(0)
		if err != nil {
			break //  Shutting down, quit
		}
		//  First frame in each message is the sender identity
		identity := msg[0]

		//  Second frame is "fetch" command
		if msg[1] != "fetch" {
			panic("command != \"fetch\"")
		}

		//  Third frame is chunk offset in file
		offset, _ := strconv.ParseInt(msg[2], 10, 64)

		//  Fourth frame is maximum chunk size
		chunksz, _ := strconv.Atoi(msg[3])

		//  Read chunk of data from file
		chunk := make([]byte, chunksz)
		n, _ := file.ReadAt(chunk, offset)

		//  Send resulting chunk to client
		router.SendMessage(identity, chunk[:n])
	}
	file.Close()
}

//  The main task is just the same as in the first model.

func main() {
	pipe := make(chan string)

	//  Start child threads
	go server_thread()
	go client_thread(pipe)
	//  Loop until client tells us it's done
	<-pipe
}
