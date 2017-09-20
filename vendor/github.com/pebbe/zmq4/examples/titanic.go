//
//  Titanic service.
//
//  Implements server side of http://rfc.zeromq.org/spec:9

package main

import (
	"github.com/pebbe/zmq4/examples/mdapi"

	"github.com/pborman/uuid"

	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

//  Returns freshly allocated request filename for given UUID

const (
	TITANIC_DIR = ".titanic"
)

func RequestFilename(uuid string) string {
	return TITANIC_DIR + "/" + uuid + "req"
}

//  Returns freshly allocated reply filename for given UUID

func ReplyFilename(uuid string) string {
	return TITANIC_DIR + "/" + uuid + "rep"
}

//  The "titanic.request" task waits for requests to this service. It writes
//  each request to disk and returns a UUID to the client. The client picks
//  up the reply asynchronously using the "titanic.reply" service:

func TitanicRequest(chRequest chan<- string) {
	worker, _ := mdapi.NewMdwrk("tcp://localhost:5555", "titanic.request", false)

	reply := []string{}
	for {
		//  Send reply if it's not null
		//  And then get next request from broker
		request, err := worker.Recv(reply)
		if err != nil {
			break //  Interrupted, exit
		}

		//  Ensure message directory exists
		os.MkdirAll(TITANIC_DIR, 0700)

		//  Generate UUID and save message to disk
		uuid := uuid.New()
		file, err := os.Create(RequestFilename(uuid))
		fmt.Fprint(file, strings.Join(request, "\n"))
		file.Close()

		//  Send UUID through to message queue
		chRequest <- uuid

		//  Now send UUID back to client
		//  Done by the mdwrk_recv() at the top of the loop
		reply = []string{"200", uuid}
	}
}

//  The "titanic.reply" task checks if there's a reply for the specified
//  request (by UUID), and returns a 200 OK, 300 Pending, or 400 Unknown
//  accordingly:

func TitanicReply() {
	worker, _ := mdapi.NewMdwrk("tcp://localhost:5555", "titanic.reply", false)

	pending := []string{"300"}
	unknown := []string{"400"}
	reply := []string{}
	for {
		request, err := worker.Recv(reply)
		if err != nil {
			break //  Interrupted, exit
		}

		uuid := request[0]
		req_filename := RequestFilename(uuid)
		rep_filename := ReplyFilename(uuid)
		data, err := ioutil.ReadFile(rep_filename)
		if err == nil {
			reply = strings.Split("200\n"+string(data), "\n")
		} else {
			_, err := os.Stat(req_filename)
			if err == nil {
				reply = pending
			} else {
				reply = unknown
			}
		}
	}
}

//  The "titanic.close" task removes any waiting replies for the request
//  (specified by UUID). It's idempotent, so safe to call more than once
//  in a row:

func TitanicClose() {
	worker, _ := mdapi.NewMdwrk("tcp://localhost:5555", "titanic.close", false)

	ok := []string{"200"}
	reply := []string{}
	for {
		request, err := worker.Recv(reply)
		if err != nil {
			break //  Interrupted, exit
		}

		uuid := request[0]
		os.Remove(RequestFilename(uuid))
		os.Remove(ReplyFilename(uuid))

		reply = ok
	}

}

//  This is the main thread for the Titanic worker. It starts three child
//  threads; for the request, reply, and close services. It then dispatches
//  requests to workers using a simple brute-force disk queue. It receives
//  request UUIDs from the titanic.request service, saves these to a disk
//  file, and then throws each request at MDP workers until it gets a
//  response:

func main() {
	var verbose bool
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		verbose = true
	}

	chRequest := make(chan string)
	go TitanicRequest(chRequest)
	go TitanicReply()
	go TitanicClose()

	//  Ensure message directory exists
	os.MkdirAll(TITANIC_DIR, 0700)

	// Fill the queue
	queue := make([]string, 0)
	files, err := ioutil.ReadDir(TITANIC_DIR)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if strings.HasSuffix(name, "req") {
				uuid := name[:len(name)-3]
				_, err := os.Stat(ReplyFilename(uuid))
				if err != nil {
					queue = append(queue, uuid)
				}
			}
		}
	}

	//  Main dispatcher loop
	for {
		//  We'll dispatch once per second, if there's no activity
		select {
		case <-time.After(time.Second):
		case uuid := <-chRequest:
			//  Append UUID to queue
			queue = append(queue, uuid)
		}

		//  Brute-force dispatcher
		queue2 := make([]string, 0, len(queue))
		for _, entry := range queue {
			if verbose {
				fmt.Println("I: processing request", entry)
			}
			if !ServiceSuccess(entry) {
				queue2 = append(queue2, entry)
			}
		}
		queue = queue2
	}
}

//  Here we first check if the requested MDP service is defined or not,
//  using a MMI lookup to the Majordomo broker. If the service exists
//  we send a request and wait for a reply using the conventional MDP
//  client API. This is not meant to be fast, just very simple:

func ServiceSuccess(uuid string) bool {
	// If reply already exists, treat as successful
	_, err := os.Stat(ReplyFilename(uuid))
	if err == nil {
		return true
	}

	//  Load request message, service will be first frame
	data, err := ioutil.ReadFile(RequestFilename(uuid))

	//  If the client already closed request, treat as successful
	if err != nil {
		return true
	}

	request := strings.Split(string(data), "\n")

	service_name := request[0]
	request = request[1:]

	//  Create MDP client session with short timeout
	client, err := mdapi.NewMdcli("tcp://localhost:5555", false)
	client.SetTimeout(time.Second) //  1 sec
	client.SetRetries(1)           //  only 1 retry

	//  Use MMI protocol to check if service is available
	mmi_reply, err := client.Send("mmi.service", service_name)
	if err != nil || mmi_reply[0] != "200" {
		return false
	}

	reply, err := client.Send(service_name, request...)
	if err != nil {
		return false
	}

	file, err := os.Create(ReplyFilename(uuid))
	fmt.Fprint(file, strings.Join(reply, "\n"))
	file.Close()

	return true

}
