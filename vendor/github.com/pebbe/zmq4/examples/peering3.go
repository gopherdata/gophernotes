//
//  Broker peering simulation (part 3).
//  Prototypes the full flow of status and tasks
//

/*

One of the differences between peering2 and peering3 is that
peering2 always uses Poll() and then uses a helper function socketInPolled()
to check if a specific socket returned a result, while peering3 uses PollAll()
and checks the event state of the socket in a specific index in the list.

*/

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	NBR_CLIENTS  = 10
	NBR_WORKERS  = 5
	WORKER_READY = "**READY**" //  Signals worker is ready
)

var (
	//  Our own name; in practice this would be configured per node
	self string
)

//  This is the client task. It issues a burst of requests and then
//  sleeps for a few seconds. This simulates sporadic activity; when
//  a number of clients are active at once, the local workers should
//  be overloaded. The client uses a REQ socket for requests and also
//  pushes statistics to the monitor socket:

func client_task(i int) {
	client, _ := zmq.NewSocket(zmq.REQ)
	defer client.Close()
	client.Connect("ipc://" + self + "-localfe.ipc")
	monitor, _ := zmq.NewSocket(zmq.PUSH)
	defer monitor.Close()
	monitor.Connect("ipc://" + self + "-monitor.ipc")

	poller := zmq.NewPoller()
	poller.Add(client, zmq.POLLIN)
	for {
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		for burst := rand.Intn(15); burst > 0; burst-- {
			task_id := fmt.Sprintf("%04X-%s-%d", rand.Intn(0x10000), self, i)

			//  Send request with random hex ID
			client.Send(task_id, 0)

			//  Wait max ten seconds for a reply, then complain
			sockets, err := poller.Poll(10 * time.Second)
			if err != nil {
				break //  Interrupted
			}

			if len(sockets) == 1 {
				reply, err := client.Recv(0)
				if err != nil {
					break //  Interrupted
				}
				//  Worker is supposed to answer us with our task id
				id := strings.Fields(reply)[0]
				if id != task_id {
					panic("id != task_id")
				}
				monitor.Send(reply, 0)
			} else {
				monitor.Send("E: CLIENT EXIT - lost task "+task_id, 0)
				return
			}
		}
	}
}

//  This is the worker task, which uses a REQ socket to plug into the
//  load-balancer. It's the same stub worker task you've seen in other
//  examples:

func worker_task(i int) {
	worker, _ := zmq.NewSocket(zmq.REQ)
	defer worker.Close()
	worker.Connect("ipc://" + self + "-localbe.ipc")

	//  Tell broker we're ready for work
	worker.SendMessage(WORKER_READY)

	//  Process messages as they arrive
	for {
		msg, err := worker.RecvMessage(0)
		if err != nil {
			break //  Interrupted
		}

		//  Workers are busy for 0/1 seconds
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		n := len(msg) - 1
		worker.SendMessage(msg[:n], fmt.Sprintf("%s %s-%d", msg[n], self, i))
	}
}

//  The main task begins by setting-up all its sockets. The local frontend
//  talks to clients, and our local backend talks to workers. The cloud
//  frontend talks to peer brokers as if they were clients, and the cloud
//  backend talks to peer brokers as if they were workers. The state
//  backend publishes regular state messages, and the state frontend
//  subscribes to all state backends to collect these messages. Finally,
//  we use a PULL monitor socket to collect printable messages from tasks:

func main() {
	//  First argument is this broker's name
	//  Other arguments are our peers' names
	//
	if len(os.Args) < 2 {
		fmt.Println("syntax: peering1 me {you}...")
		os.Exit(1)
	}
	self = os.Args[1]
	fmt.Printf("I: preparing broker at %s...\n", self)
	rand.Seed(time.Now().UnixNano())

	//  Prepare local frontend and backend
	localfe, _ := zmq.NewSocket(zmq.ROUTER)
	defer localfe.Close()
	localfe.Bind("ipc://" + self + "-localfe.ipc")

	localbe, _ := zmq.NewSocket(zmq.ROUTER)
	defer localbe.Close()
	localbe.Bind("ipc://" + self + "-localbe.ipc")

	//  Bind cloud frontend to endpoint
	cloudfe, _ := zmq.NewSocket(zmq.ROUTER)
	defer cloudfe.Close()
	cloudfe.SetIdentity(self)
	cloudfe.Bind("ipc://" + self + "-cloud.ipc")

	//  Connect cloud backend to all peers
	cloudbe, _ := zmq.NewSocket(zmq.ROUTER)
	defer cloudbe.Close()
	cloudbe.SetIdentity(self)
	for _, peer := range os.Args[2:] {
		fmt.Printf("I: connecting to cloud frontend at '%s'\n", peer)
		cloudbe.Connect("ipc://" + peer + "-cloud.ipc")
	}
	//  Bind state backend to endpoint
	statebe, _ := zmq.NewSocket(zmq.PUB)
	defer statebe.Close()
	statebe.Bind("ipc://" + self + "-state.ipc")

	//  Connect state frontend to all peers
	statefe, _ := zmq.NewSocket(zmq.SUB)
	defer statefe.Close()
	statefe.SetSubscribe("")
	for _, peer := range os.Args[2:] {
		fmt.Printf("I: connecting to state backend at '%s'\n", peer)
		statefe.Connect("ipc://" + peer + "-state.ipc")
	}
	//  Prepare monitor socket
	monitor, _ := zmq.NewSocket(zmq.PULL)
	defer monitor.Close()
	monitor.Bind("ipc://" + self + "-monitor.ipc")

	//  After binding and connecting all our sockets, we start our child
	//  tasks - workers and clients:

	for worker_nbr := 0; worker_nbr < NBR_WORKERS; worker_nbr++ {
		go worker_task(worker_nbr)
	}

	//  Start local clients
	for client_nbr := 0; client_nbr < NBR_CLIENTS; client_nbr++ {
		go client_task(client_nbr)
	}

	//  Queue of available workers
	local_capacity := 0
	cloud_capacity := 0
	workers := make([]string, 0)

	primary := zmq.NewPoller()
	primary.Add(localbe, zmq.POLLIN)
	primary.Add(cloudbe, zmq.POLLIN)
	primary.Add(statefe, zmq.POLLIN)
	primary.Add(monitor, zmq.POLLIN)

	secondary1 := zmq.NewPoller()
	secondary1.Add(localfe, zmq.POLLIN)
	secondary2 := zmq.NewPoller()
	secondary2.Add(localfe, zmq.POLLIN)
	secondary2.Add(cloudfe, zmq.POLLIN)

	msg := make([]string, 0)
	for {

		//  If we have no workers ready, wait indefinitely
		timeout := time.Duration(time.Second)
		if local_capacity == 0 {
			timeout = -1
		}
		sockets, err := primary.PollAll(timeout)
		if err != nil {
			break //  Interrupted
		}

		//  Track if capacity changes during this iteration
		previous := local_capacity

		//  Handle reply from local worker
		msg = msg[0:0]

		if sockets[0].Events&zmq.POLLIN != 0 { // 0 == localbe
			msg, err = localbe.RecvMessage(0)
			if err != nil {
				break //  Interrupted
			}
			var identity string
			identity, msg = unwrap(msg)
			workers = append(workers, identity)
			local_capacity++

			//  If it's READY, don't route the message any further
			if msg[0] == WORKER_READY {
				msg = msg[0:0]
			}
		} else if sockets[1].Events&zmq.POLLIN != 0 { // 1 == cloudbe
			//  Or handle reply from peer broker
			msg, err = cloudbe.RecvMessage(0)
			if err != nil {
				break //  Interrupted
			}
			//  We don't use peer broker identity for anything
			_, msg = unwrap(msg)
		}

		if len(msg) > 0 {

			//  Route reply to cloud if it's addressed to a broker
			to_broker := false
			for _, peer := range os.Args[2:] {
				if peer == msg[0] {
					to_broker = true
					break
				}
			}
			if to_broker {
				cloudfe.SendMessage(msg)
			} else {
				localfe.SendMessage(msg)
			}
		}

		//  If we have input messages on our statefe or monitor sockets we
		//  can process these immediately:

		if sockets[2].Events&zmq.POLLIN != 0 { // 2 == statefe
			var status string
			m, _ := statefe.RecvMessage(0)
			_, m = unwrap(m) // peer
			status, _ = unwrap(m)
			cloud_capacity, _ = strconv.Atoi(status)
		}
		if sockets[3].Events&zmq.POLLIN != 0 { // 3 == monitor
			status, _ := monitor.Recv(0)
			fmt.Println(status)
		}
		//  Now route as many clients requests as we can handle. If we have
		//  local capacity we poll both localfe and cloudfe. If we have cloud
		//  capacity only, we poll just localfe. We route any request locally
		//  if we can, else we route to the cloud.

		for local_capacity+cloud_capacity > 0 {
			var sockets []zmq.Polled
			var err error
			if local_capacity > 0 {
				sockets, err = secondary2.PollAll(0)
			} else {
				sockets, err = secondary1.PollAll(0)
			}
			if err != nil {
				panic(err)
			}

			if sockets[0].Events&zmq.POLLIN != 0 { // 0 == localfe
				msg, _ = localfe.RecvMessage(0)
			} else if len(sockets) > 1 && sockets[1].Events&zmq.POLLIN != 0 { // 1 == cloudfe
				msg, _ = cloudfe.RecvMessage(0)
			} else {
				break //  No work, go back to primary
			}

			if local_capacity > 0 {
				localbe.SendMessage(workers[0], "", msg)
				workers = workers[1:]
				local_capacity--
			} else {
				//  Route to random broker peer
				random_peer := rand.Intn(len(os.Args)-2) + 2
				cloudbe.SendMessage(os.Args[random_peer], "", msg)
			}
		}
		//  We broadcast capacity messages to other peers; to reduce chatter
		//  we do this only if our capacity changed.

		if local_capacity != previous {
			//  We stick our own identity onto the envelope
			//  Broadcast new capacity
			statebe.SendMessage(self, "", local_capacity)
		}
	}
}

//  Pops frame off front of message and returns it as 'head'
//  If next frame is empty, pops that empty frame.
//  Return remaining frames of message as 'tail'
func unwrap(msg []string) (head string, tail []string) {
	head = msg[0]
	if len(msg) > 1 && msg[1] == "" {
		tail = msg[2:]
	} else {
		tail = msg[1:]
	}
	return
}
