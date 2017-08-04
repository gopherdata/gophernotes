//  The Grasslands Pattern
//
//  The Classic ZeroMQ model, plain text with no protection at all.

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"log"
	"runtime"
)

func main() {

	//  Create and bind server socket
	server, err := zmq.NewSocket(zmq.PUSH)
	checkErr(err)
	checkErr(server.Bind("tcp://*:9000"))

	//  Create and connect client socket
	client, err := zmq.NewSocket(zmq.PULL)
	checkErr(err)
	checkErr(client.Connect("tcp://127.0.0.1:9000"))

	//  Send a single message from server to client
	_, err = server.Send("Hello", 0)
	checkErr(err)
	message, err := client.Recv(0)
	checkErr(err)
	if message != "Hello" {
		log.Fatalln(message, "!= Hello")
	}

	fmt.Println("Grasslands test OK")
}

func checkErr(err error) {
	if err != nil {
		log.SetFlags(0)
		_, filename, lineno, ok := runtime.Caller(1)
		if ok {
			log.Fatalf("%v:%v: %v", filename, lineno, err)
		} else {
			log.Fatalln(err)
		}
	}
}
