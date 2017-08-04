//  The Ironhouse Pattern
//
//  Security doesn't get any stronger than this. An attacker is going to
//  have to break into your systems to see data before/after encryption.

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
	zmq.AuthAllow("domain1", "127.0.0.1/8")

	//  We need two certificates, one for the client and one for
	//  the server. The client must know the server's public key
	//  to make a CURVE connection.
	client_public, client_secret, err := zmq.NewCurveKeypair()
	checkErr(err)
	server_public, server_secret, err := zmq.NewCurveKeypair()
	checkErr(err)

	//  Tell authenticator to use this public client key
	zmq.AuthCurveAdd("domain1", client_public)

	//  Create and bind server socket
	server, _ := zmq.NewSocket(zmq.PUSH)
	server.ServerAuthCurve("domain1", server_secret)
	server.Bind("tcp://*:9000")

	//  Create and connect client socket
	client, _ := zmq.NewSocket(zmq.PULL)
	client.ClientAuthCurve(server_public, client_public, client_secret)
	client.Connect("tcp://127.0.0.1:9000")

	//  Send a single message from server to client
	_, err = server.Send("Hello", 0)
	checkErr(err)
	message, err := client.Recv(0)
	checkErr(err)
	if message != "Hello" {
		log.Fatalln(message, "!= Hello")
	}

	zmq.AuthStop()

	fmt.Println("Ironhouse test OK")

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
