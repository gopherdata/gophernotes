//
//  Paranoid Pirate worker.
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"math/rand"
	"time"
)

const (
	HEARTBEAT_LIVENESS = 3                        //  3-5 is reasonable
	HEARTBEAT_INTERVAL = 1000 * time.Millisecond  //  msecs
	INTERVAL_INIT      = 1000 * time.Millisecond  //  Initial reconnect
	INTERVAL_MAX       = 32000 * time.Millisecond //  After exponential backoff

	//  Paranoid Pirate Protocol constants
	PPP_READY     = "\001" //  Signals worker is ready
	PPP_HEARTBEAT = "\002" //  Signals worker heartbeat
)

//  Helper function that returns a new configured socket
//  connected to the Paranoid Pirate queue

func s_worker_socket() (*zmq.Socket, *zmq.Poller) {
	worker, _ := zmq.NewSocket(zmq.DEALER)
	worker.Connect("tcp://localhost:5556")

	//  Tell queue we're ready for work
	fmt.Println("I: worker ready")
	worker.Send(PPP_READY, 0)

	poller := zmq.NewPoller()
	poller.Add(worker, zmq.POLLIN)

	return worker, poller
}

//  We have a single task, which implements the worker side of the
//  Paranoid Pirate Protocol (PPP). The interesting parts here are
//  the heartbeating, which lets the worker detect if the queue has
//  died, and vice-versa:

func main() {
	worker, poller := s_worker_socket()

	//  If liveness hits zero, queue is considered disconnected
	liveness := HEARTBEAT_LIVENESS
	interval := INTERVAL_INIT

	//  Send out heartbeats at regular intervals
	heartbeat_at := time.Tick(HEARTBEAT_INTERVAL)

	rand.Seed(time.Now().UnixNano())
	for cycles := 0; true; {
		sockets, err := poller.Poll(HEARTBEAT_INTERVAL)
		if err != nil {
			break //  Interrupted
		}

		if len(sockets) == 1 {
			//  Get message
			//  - 3-part envelope + content -> request
			//  - 1-part HEARTBEAT -> heartbeat
			msg, err := worker.RecvMessage(0)
			if err != nil {
				break //  Interrupted
			}

			//  To test the robustness of the queue implementation we //
			//  simulate various typical problems, such as the worker
			//  crashing, or running very slowly. We do this after a few
			//  cycles so that the architecture can get up and running
			//  first:
			if len(msg) == 3 {
				cycles++
				if cycles > 3 && rand.Intn(5) == 0 {
					fmt.Println("I: simulating a crash")
					break
				} else if cycles > 3 && rand.Intn(5) == 0 {
					fmt.Println("I: simulating CPU overload")
					time.Sleep(3 * time.Second)
				}
				fmt.Println("I: normal reply")
				worker.SendMessage(msg)
				liveness = HEARTBEAT_LIVENESS
				time.Sleep(time.Second) //  Do some heavy work
			} else if len(msg) == 1 {
				//  When we get a heartbeat message from the queue, it means the
				//  queue was (recently) alive, so reset our liveness indicator:
				if msg[0] == PPP_HEARTBEAT {
					liveness = HEARTBEAT_LIVENESS
				} else {
					fmt.Printf("E: invalid message: %q\n", msg)
				}
			} else {
				fmt.Printf("E: invalid message: %q\n", msg)
			}
			interval = INTERVAL_INIT
		} else {
			//  If the queue hasn't sent us heartbeats in a while, destroy the
			//  socket and reconnect. This is the simplest most brutal way of
			//  discarding any messages we might have sent in the meantime://
			liveness--
			if liveness == 0 {
				fmt.Println("W: heartbeat failure, can't reach queue")
				fmt.Println("W: reconnecting in", interval)
				time.Sleep(interval)

				if interval < INTERVAL_MAX {
					interval = 2 * interval
				}
				worker, poller = s_worker_socket()
				liveness = HEARTBEAT_LIVENESS
			}
		}

		//  Send heartbeat to queue if it's time
		select {
		case <-heartbeat_at:
			fmt.Println("I: worker heartbeat")
			worker.Send(PPP_HEARTBEAT, 0)
		default:
		}
	}
}
