//
//  Task sink.
//  Binds PULL socket to tcp://localhost:5558
//  Collects results from workers via that socket
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"time"
)

func main() {
	//  Prepare our socket
	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Bind("tcp://*:5558")

	//  Wait for start of batch
	receiver.Recv(0)

	//  Start our clock now
	start_time := time.Now()

	//  Process 100 confirmations
	for task_nbr := 0; task_nbr < 100; task_nbr++ {
		receiver.Recv(0)
		if task_nbr%10 == 0 {
			fmt.Print(":")
		} else {
			fmt.Print(".")
		}
	}

	//  Calculate and report duration of batch
	fmt.Println("\nTotal elapsed time:", time.Since(start_time))
}
