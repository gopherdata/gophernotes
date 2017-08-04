//
//  ROUTER-to-DEALER example.
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"math/rand"
	"time"
)

const (
	NBR_WORKERS = 10
)

func worker_task() {
	worker, _ := zmq.NewSocket(zmq.DEALER)
	defer worker.Close()
	set_id(worker) //  Set a printable identity
	worker.Connect("tcp://localhost:5671")

	total := 0
	for {
		//  Tell the broker we're ready for work
		worker.Send("", zmq.SNDMORE)
		worker.Send("Hi Boss", 0)

		//  Get workload from broker, until finished
		worker.Recv(0) //  Envelope delimiter
		workload, _ := worker.Recv(0)
		if workload == "Fired!" {
			fmt.Printf("Completed: %d tasks\n", total)
			break
		}
		total++

		//  Do some random work
		time.Sleep(time.Duration(rand.Intn(500)+1) * time.Millisecond)
	}
}

func main() {
	broker, _ := zmq.NewSocket(zmq.ROUTER)
	defer broker.Close()

	broker.Bind("tcp://*:5671")
	rand.Seed(time.Now().UnixNano())

	for worker_nbr := 0; worker_nbr < NBR_WORKERS; worker_nbr++ {
		go worker_task()
	}
	//  Run for five seconds and then tell workers to end
	start_time := time.Now()
	workers_fired := 0
	for {
		//  Next message gives us least recently used worker
		identity, _ := broker.Recv(0)
		broker.Send(identity, zmq.SNDMORE)
		broker.Recv(0) //  Envelope delimiter
		broker.Recv(0) //  Response from worker
		broker.Send("", zmq.SNDMORE)

		//  Encourage workers until it's time to fire them
		if time.Since(start_time) < 5*time.Second {
			broker.Send("Work harder", 0)
		} else {
			broker.Send("Fired!", 0)
			workers_fired++
			if workers_fired == NBR_WORKERS {
				break
			}
		}
	}

	time.Sleep(time.Second)
}

func set_id(soc *zmq.Socket) {
	identity := fmt.Sprintf("%04X-%04X", rand.Intn(0x10000), rand.Intn(0x10000))
	soc.SetIdentity(identity)
}
