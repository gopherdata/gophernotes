//
//  Synchronized subscriber
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"log"
	"time"
)

func main() {

	//  First, connect our subscriber socket
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5561")
	subscriber.SetSubscribe("")

	//  0MQ is so fast, we need to wait a while...
	time.Sleep(time.Second)

	//  Second, synchronize with publisher
	syncclient, _ := zmq.NewSocket(zmq.REQ)
	defer syncclient.Close()
	syncclient.Connect("tcp://localhost:5562")

	//  - send a synchronization request
	syncclient.Send("", 0)

	//  - wait for synchronization reply
	syncclient.Recv(0)

	//  Third, get our updates and report how many we got
	update_nbr := 0
	for {
		msg, e := subscriber.Recv(0)
		if e != nil {
			log.Println(e)
			break
		}
		if msg == "END" {
			break
		}
		update_nbr++
	}
	fmt.Printf("Received %d updates\n", update_nbr)
}
