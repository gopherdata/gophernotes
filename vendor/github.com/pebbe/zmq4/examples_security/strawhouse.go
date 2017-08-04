//  The Strawhouse Pattern
//
//  We allow or deny clients according to their IP address. It may keep
//  spammers and idiots away, but won't stop a real attacker for more
//  than a heartbeat.

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"log"
	"runtime"
)

func main() {

	//  Get some indication of what the authenticator is deciding
	zmq.AuthSetVerbose(true)

	//  Start the authentication engine. This engine
	//  allows or denies incoming connections (talking to the libzmq
	//  core over a protocol called ZAP).
	zmq.AuthStart()

	//  Whitelist our address; any other address will be rejected
	zmq.AuthAllow("domain1", "127.0.0.1")

	//  Create and bind server socket
	server, err := zmq.NewSocket(zmq.PUSH)
	checkErr(err)
	server.ServerAuthNull("domain1")
	server.Bind("tcp://*:9000")

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

	zmq.AuthStop()

	fmt.Println("Strawhouse test OK")
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
