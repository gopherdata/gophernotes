//
//  Hello World client.
//  Connects REQ socket to tcp://localhost:5555
//  Sends "Hello" to server, expects "World" back
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
)

func main() {
	//  Socket to talk to server
	fmt.Println("Connecting to hello world server...")
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect("tcp://localhost:5555")

	for request_nbr := 0; request_nbr != 10; request_nbr++ {
		// send hello
		msg := fmt.Sprintf("Hello %d", request_nbr)
		fmt.Println("Sending ", msg)
		requester.Send(msg, 0)

		// Wait for reply:
		reply, _ := requester.Recv(0)
		fmt.Println("Received ", reply)
	}
}
