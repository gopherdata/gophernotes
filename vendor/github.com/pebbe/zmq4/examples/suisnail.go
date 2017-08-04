//
//  Suicidal Snail
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

//  This is our subscriber. It connects to the publisher and subscribes to
//  everything. It sleeps for a short time between messages to simulate doing
//  too much work. If a message is more than 1 second late, it croaks:

const (
	MAX_ALLOWED_DELAY = 1000 * time.Millisecond
)

func subscriber(pipe chan<- string) {
	//  Subscribe to everything
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	subscriber.SetSubscribe("")
	subscriber.Connect("tcp://localhost:5556")
	defer subscriber.Close()

	//  Get and process messages
	for {
		msg, _ := subscriber.RecvMessage(0)
		i, _ := strconv.Atoi(msg[0])
		clock := time.Unix(int64(i), 0)
		fmt.Println(clock)

		//  Suicide snail logic
		if time.Now().After(clock.Add(MAX_ALLOWED_DELAY)) {
			log.Println("E: subscriber cannot keep up, aborting")
			break
		}
		//  Work for 1 msec plus some random additional time
		time.Sleep(time.Duration(1 + rand.Intn(2)))
	}
	pipe <- "gone and died"
}

//  This is our publisher task. It publishes a time-stamped message to its
//  PUB socket every 1 msec:

func publisher(pipe <-chan string) {
	//  Prepare publisher
	publisher, _ := zmq.NewSocket(zmq.PUB)
	publisher.Bind("tcp://*:5556")
	defer publisher.Close()

LOOP:
	for {
		//  Send current clock (msecs) to subscribers
		publisher.SendMessage(time.Now().Unix())
		select {
		case <-pipe:
			break LOOP
		default:
		}
		time.Sleep(time.Millisecond)
	}
}

//  The main task simply starts a client, and a server, and then
//  waits for the client to signal that it has died:

func main() {
	pubpipe := make(chan string)
	subpipe := make(chan string)
	go publisher(pubpipe)
	go subscriber(subpipe)
	<-subpipe
	pubpipe <- "break"
	time.Sleep(100 * time.Millisecond)
}
