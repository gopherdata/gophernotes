//
//  Lazy Pirate server.
//  Binds REQ socket to tcp://*:5555
//  Like hwserver except:
//   - echoes request as-is
//   - randomly runs slowly, or exits to simulate a crash.
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	server, _ := zmq.NewSocket(zmq.REP)
	defer server.Close()
	server.Bind("tcp://*:5555")

	for cycles := 0; true; {
		request, _ := server.RecvMessage(0)
		cycles++

		//  Simulate various problems, after a few cycles
		if cycles > 3 && rand.Intn(3) == 0 {
			fmt.Println("I: simulating a crash")
			break
		} else if cycles > 3 && rand.Intn(3) == 0 {
			fmt.Println("I: simulating CPU overload")
			time.Sleep(2 * time.Second)
		}
		fmt.Printf("I: normal request (%s)\n", request)
		time.Sleep(time.Second) //  Do some heavy work
		server.SendMessage(request)
	}
}
