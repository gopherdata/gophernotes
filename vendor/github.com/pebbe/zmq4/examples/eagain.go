//
//  Shows how to provoke EAGAIN when reaching HWM
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
)

func main() {

	mailbox, _ := zmq.NewSocket(zmq.DEALER)
	mailbox.SetSndhwm(4)
	mailbox.SetSndtimeo(0)
	mailbox.Connect("tcp://localhost:9876")

	for count := 0; count < 10; count++ {
		fmt.Println("Sending message", count)
		_, err := mailbox.SendMessage(fmt.Sprint("message ", count))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
