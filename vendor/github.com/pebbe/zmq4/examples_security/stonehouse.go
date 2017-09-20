//  The Stonehouse Pattern
//
//  Where we allow any clients to connect, but we promise clients
//  that we are who we claim to be, and our conversations won't be
//  tampered with or modified, or spied on.

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

	//  Tell the authenticator to allow any CURVE requests for this domain
	zmq.AuthCurveAdd("domain1", zmq.CURVE_ALLOW_ANY)

	//  We need two certificates, one for the client and one for
	//  the server. The client must know the server's public key
	//  to make a CURVE connection.
	client_public, client_secret, err := zmq.NewCurveKeypair()
	checkErr(err)
	server_public, server_secret, err := zmq.NewCurveKeypair()
	checkErr(err)

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

	fmt.Println("Stonehouse test OK")

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
