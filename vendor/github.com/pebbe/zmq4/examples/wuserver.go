//
//  Weather update server.
//  Binds PUB socket to tcp://*:5556
//  Publishes random weather updates
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"math/rand"
	"time"
)

func main() {

	//  Prepare our publisher
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Bind("tcp://*:5556")
	publisher.Bind("ipc://weather.ipc")

	//  Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// loop for a while apparently
	for {

		//  Get values that will fool the boss
		zipcode := rand.Intn(100000)
		temperature := rand.Intn(215) - 80
		relhumidity := rand.Intn(50) + 10

		//  Send message to all subscribers
		msg := fmt.Sprintf("%05d %d %d", zipcode, temperature, relhumidity)
		publisher.Send(msg, 0)
	}
}
