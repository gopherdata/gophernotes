//
//  Binary Star server, using bstar reactor.
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"github.com/pebbe/zmq4/examples/bstar"

	"fmt"
	"os"
)

//  Echo service
func echo(socket *zmq.Socket) (err error) {
	msg, err := socket.RecvMessage(0)
	if err != nil {
		return
	}
	_, err = socket.SendMessage(msg)
	return
}

func main() {
	//  Arguments can be either of:
	//      -p  primary server, at tcp://localhost:5001
	//      -b  backup server, at tcp://localhost:5002
	var bst *bstar.Bstar
	if len(os.Args) == 2 && os.Args[1] == "-p" {
		fmt.Println("I: Primary active, waiting for backup (passive)")
		bst, _ = bstar.New(bstar.PRIMARY, "tcp://*:5003", "tcp://localhost:5004")
		bst.Voter("tcp://*:5001", zmq.ROUTER, echo)
	} else if len(os.Args) == 2 && os.Args[1] == "-b" {
		fmt.Println("I: Backup passive, waiting for primary (active)")
		bst, _ = bstar.New(bstar.BACKUP, "tcp://*:5004", "tcp://localhost:5003")
		bst.Voter("tcp://*:5002", zmq.ROUTER, echo)
	} else {
		fmt.Println("Usage: bstarsrvs { -p | -b }")
		return
	}
	bst.Start()
}
