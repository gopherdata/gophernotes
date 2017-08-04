//
//  Clone client Model Six
//

package main

import (
	"github.com/pebbe/zmq4/examples/clone"

	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	SUBTREE = "/client/"
)

func main() {
	//  Create distributed hash instance
	clone := clone.New()

	//  Specify configuration
	clone.Subtree(SUBTREE)
	clone.Connect("tcp://localhost", "5556")
	clone.Connect("tcp://localhost", "5566")

	//  Set random tuples into the distributed hash
	for {
		//  Set random value, check it was stored
		key := fmt.Sprintf("%s%d", SUBTREE, rand.Intn(10000))
		value := fmt.Sprint(rand.Intn(1000000))
		clone.Set(key, value, rand.Intn(30))
		v, _ := clone.Get(key)
		if v != value {
			log.Fatalf("Set: %v - Get: %v - Equal: %v\n", value, v, value == v)
		}
		time.Sleep(time.Second)
	}
}
