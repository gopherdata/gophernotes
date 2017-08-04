//  The Woodhouse Pattern
//
//  It may keep some malicious people out but all it takes is a bit
//  of network sniffing, and they'll be able to fake their way in.

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"log"
	"runtime"
)

func main() {

	//  Start authentication engine
	zmq.AuthSetVerbose(true)
	zmq.AuthStart()
	zmq.AuthAllow("domain1", "127.0.0.1")

	//  Tell the authenticator how to handle PLAIN requests
	zmq.AuthPlainAdd("domain1", "admin", "secret")

	//  Create and bind server socket
	server, _ := zmq.NewSocket(zmq.PUSH)
	server.ServerAuthPlain("domain1")
	server.Bind("tcp://*:9000")

	//  Create and connect client socket
	client, _ := zmq.NewSocket(zmq.PULL)
	client.SetPlainUsername("admin")
	client.SetPlainPassword("secret")
	client.Connect("tcp://127.0.0.1:9000")

	//  Send a single message from server to client
	_, err := server.Send("Hello", 0)
	checkErr(err)
	message, err := client.Recv(0)
	checkErr(err)
	if message != "Hello" {
		log.Fatalln(message, "!= Hello")
	}

	zmq.AuthStop()

	fmt.Println("Woodhouse test OK")

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
