//
//  Simple Pirate worker.
//  Connects REQ socket to tcp://*:5556
//  Implements worker part of load-balancing
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"math/rand"
	"time"
)

const (
	WORKER_READY = "\001" //  Signals worker is ready
)

func main() {
	worker, _ := zmq.NewSocket(zmq.REQ)
	defer worker.Close()

	//  Set random identity to make tracing easier
	rand.Seed(time.Now().UnixNano())
	identity := fmt.Sprintf("%04X-%04X", rand.Intn(0x10000), rand.Intn(0x10000))
	worker.SetIdentity(identity)
	worker.Connect("tcp://localhost:5556")

	//  Tell broker we're ready for work
	fmt.Printf("I: (%s) worker ready\n", identity)
	worker.Send(WORKER_READY, 0)

	for cycles := 0; true; {
		msg, err := worker.RecvMessage(0)
		if err != nil {
			break //  Interrupted
		}

		//  Simulate various problems, after a few cycles
		cycles++
		if cycles > 3 && rand.Intn(5) == 0 {
			fmt.Printf("I: (%s) simulating a crash\n", identity)
			break
		} else if cycles > 3 && rand.Intn(5) == 0 {
			fmt.Printf("I: (%s) simulating CPU overload\n", identity)
			time.Sleep(3 * time.Second)
		}

		fmt.Printf("I: (%s) normal reply\n", identity)
		time.Sleep(time.Second) //  Do some heavy work
		worker.SendMessage(msg)
	}
}
