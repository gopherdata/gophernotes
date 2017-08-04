//
//  Task sink - design 2.
//  Adds pub-sub flow to send kill signal to workers
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"time"
)

func main() {
	//  Socket to receive messages on
	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Bind("tcp://*:5558")

	//  Socket for worker control
	controller, _ := zmq.NewSocket(zmq.PUB)
	defer controller.Close()
	controller.Bind("tcp://*:5559")

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
	fmt.Println("\nTotal elapsed time:", time.Since(start_time))

	//  Send kill signal to workers
	controller.Send("KILL", 0)

	//  Finished
	time.Sleep(time.Second) //  Give 0MQ time to deliver
}
