//
//  Freelance client - Model 3.
//  Uses flcliapi class to encapsulate Freelance pattern
//

package main

import (
	"github.com/pebbe/zmq4/examples/flcliapi"

	"fmt"
	"time"
)

func main() {
	//  Create new freelance client object
	client := flcliapi.New()

	//  Connect to several endpoints
	client.Connect("tcp://localhost:5555")
	client.Connect("tcp://localhost:5556")
	client.Connect("tcp://localhost:5557")

	//  Send a bunch of name resolution 'requests', measure time
	start := time.Now()
	req := []string{"random name"}
	for requests := 1000; requests > 0; requests-- {
		_, err := client.Request(req)
		if err != nil {
			fmt.Println("E: name service not available, aborting")
			break
		}
	}
	fmt.Println("Average round trip cost:", time.Now().Sub(start)/1000)
}
