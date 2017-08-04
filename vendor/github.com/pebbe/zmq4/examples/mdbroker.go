//
//  Majordomo Protocol broker.
//  A minimal Go implementation of the Majordomo Protocol as defined in
//  http://rfc.zeromq.org/spec:7 and http://rfc.zeromq.org/spec:8.
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"github.com/pebbe/zmq4/examples/mdapi"

	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

const (
	//  We'd normally pull these from config data

	HEARTBEAT_LIVENESS = 3                       //  3-5 is reasonable
	HEARTBEAT_INTERVAL = 2500 * time.Millisecond //  msecs
	HEARTBEAT_EXPIRY   = HEARTBEAT_INTERVAL * HEARTBEAT_LIVENESS
)

//  The broker class defines a single broker instance:

type Broker struct {
	socket       *zmq.Socket         //  Socket for clients & workers
	verbose      bool                //  Print activity to stdout
	endpoint     string              //  Broker binds to this endpoint
	services     map[string]*Service //  Hash of known services
	workers      map[string]*Worker  //  Hash of known workers
	waiting      []*Worker           //  List of waiting workers
	heartbeat_at time.Time           //  When to send HEARTBEAT
}

//  The service class defines a single service instance:

type Service struct {
	broker   *Broker    //  Broker instance
	name     string     //  Service name
	requests [][]string //  List of client requests
	waiting  []*Worker  //  List of waiting workers
}

//  The worker class defines a single worker, idle or active:

type Worker struct {
	broker    *Broker   //  Broker instance
	id_string string    //  Identity of worker as string
	identity  string    //  Identity frame for routing
	service   *Service  //  Owning service, if known
	expiry    time.Time //  Expires at unless heartbeat
}

//  Here are the constructor and destructor for the broker:

func NewBroker(verbose bool) (broker *Broker, err error) {

	//  Initialize broker state
	broker = &Broker{
		verbose:      verbose,
		services:     make(map[string]*Service),
		workers:      make(map[string]*Worker),
		waiting:      make([]*Worker, 0),
		heartbeat_at: time.Now().Add(HEARTBEAT_INTERVAL),
	}
	broker.socket, err = zmq.NewSocket(zmq.ROUTER)

	broker.socket.SetRcvhwm(500000) // or example mdclient2 won't work

	runtime.SetFinalizer(broker, (*Broker).Close)
	return
}

func (broker *Broker) Close() (err error) {
	if broker.socket != nil {
		err = broker.socket.Close()
		broker.socket = nil
	}
	return
}

//  The bind method binds the broker instance to an endpoint. We can call
//  this multiple times. Note that MDP uses a single socket for both clients
//  and workers:

func (broker *Broker) Bind(endpoint string) (err error) {
	err = broker.socket.Bind(endpoint)
	if err != nil {
		log.Println("E: MDP broker/0.2.0 failed to bind at", endpoint)
		return
	}
	log.Println("I: MDP broker/0.2.0 is active at", endpoint)
	return
}

//  The WorkerMsg method processes one READY, REPLY, HEARTBEAT or
//  DISCONNECT message sent to the broker by a worker:

func (broker *Broker) WorkerMsg(sender string, msg []string) {
	//  At least, command
	if len(msg) == 0 {
		panic("len(msg) == 0")
	}

	command, msg := popStr(msg)
	id_string := fmt.Sprintf("%q", sender)
	_, worker_ready := broker.workers[id_string]
	worker := broker.WorkerRequire(sender)

	switch command {
	case mdapi.MDPW_READY:
		if worker_ready { //  Not first command in session
			worker.Delete(true)
		} else if len(sender) >= 4 /*  Reserved service name */ && sender[:4] == "mmi." {
			worker.Delete(true)
		} else {
			//  Attach worker to service and mark as idle
			worker.service = broker.ServiceRequire(msg[0])
			worker.Waiting()
		}
	case mdapi.MDPW_REPLY:
		if worker_ready {
			//  Remove & save client return envelope and insert the
			//  protocol header and service name, then rewrap envelope.
			client, msg := unwrap(msg)
			broker.socket.SendMessage(client, "", mdapi.MDPC_CLIENT, worker.service.name, msg)
			worker.Waiting()
		} else {
			worker.Delete(true)
		}
	case mdapi.MDPW_HEARTBEAT:
		if worker_ready {
			worker.expiry = time.Now().Add(HEARTBEAT_EXPIRY)
		} else {
			worker.Delete(true)
		}
	case mdapi.MDPW_DISCONNECT:
		worker.Delete(false)
	default:
		log.Printf("E: invalid input message %q\n", msg)
	}
}

//  Process a request coming from a client. We implement MMI requests
//  directly here (at present, we implement only the mmi.service request):

func (broker *Broker) ClientMsg(sender string, msg []string) {
	//  Service name + body
	if len(msg) < 2 {
		panic("len(msg) < 2")
	}

	service_frame, msg := popStr(msg)
	service := broker.ServiceRequire(service_frame)

	//  Set reply return identity to client sender
	m := []string{sender, ""}
	msg = append(m, msg...)

	//  If we got a MMI service request, process that internally
	if len(service_frame) >= 4 && service_frame[:4] == "mmi." {
		var return_code string
		if service_frame == "mmi.service" {
			name := msg[len(msg)-1]
			service, ok := broker.services[name]
			if ok && len(service.waiting) > 0 {
				return_code = "200"
			} else {
				return_code = "404"
			}
		} else {
			return_code = "501"
		}

		msg[len(msg)-1] = return_code

		//  Remove & save client return envelope and insert the
		//  protocol header and service name, then rewrap envelope.
		client, msg := unwrap(msg)
		broker.socket.SendMessage(client, "", mdapi.MDPC_CLIENT, service_frame, msg)
	} else {
		//  Else dispatch the message to the requested service
		service.Dispatch(msg)
	}
}

//  The purge method deletes any idle workers that haven't pinged us in a
//  while. We hold workers from oldest to most recent, so we can stop
//  scanning whenever we find a live worker. This means we'll mainly stop
//  at the first worker, which is essential when we have large numbers of
//  workers (since we call this method in our critical path):

func (broker *Broker) Purge() {
	now := time.Now()
	for len(broker.waiting) > 0 {
		if broker.waiting[0].expiry.After(now) {
			break //  Worker is alive, we're done here
		}
		if broker.verbose {
			log.Println("I: deleting expired worker:", broker.waiting[0].id_string)
		}
		broker.waiting[0].Delete(false)
	}
}

//  Here is the implementation of the methods that work on a service:

//  Lazy constructor that locates a service by name, or creates a new
//  service if there is no service already with that name.

func (broker *Broker) ServiceRequire(service_frame string) (service *Service) {
	name := service_frame
	service, ok := broker.services[name]
	if !ok {
		service = &Service{
			broker:   broker,
			name:     name,
			requests: make([][]string, 0),
			waiting:  make([]*Worker, 0),
		}
		broker.services[name] = service
		if broker.verbose {
			log.Println("I: added service:", name)
		}
	}
	return
}

//  The dispatch method sends requests to waiting workers:

func (service *Service) Dispatch(msg []string) {

	if len(msg) > 0 {
		//  Queue message if any
		service.requests = append(service.requests, msg)
	}

	service.broker.Purge()
	for len(service.waiting) > 0 && len(service.requests) > 0 {
		var worker *Worker
		worker, service.waiting = popWorker(service.waiting)
		service.broker.waiting = delWorker(service.broker.waiting, worker)
		msg, service.requests = popMsg(service.requests)
		worker.Send(mdapi.MDPW_REQUEST, "", msg)
	}
}

//  Here is the implementation of the methods that work on a worker:

//  Lazy constructor that locates a worker by identity, or creates a new
//  worker if there is no worker already with that identity.

func (broker *Broker) WorkerRequire(identity string) (worker *Worker) {

	//  broker.workers is keyed off worker identity
	id_string := fmt.Sprintf("%q", identity)
	worker, ok := broker.workers[id_string]
	if !ok {
		worker = &Worker{
			broker:    broker,
			id_string: id_string,
			identity:  identity,
		}
		broker.workers[id_string] = worker
		if broker.verbose {
			log.Printf("I: registering new worker: %s\n", id_string)
		}
	}
	return
}

//  The delete method deletes the current worker.

func (worker *Worker) Delete(disconnect bool) {
	if disconnect {
		worker.Send(mdapi.MDPW_DISCONNECT, "", []string{})
	}

	if worker.service != nil {
		worker.service.waiting = delWorker(worker.service.waiting, worker)
	}
	worker.broker.waiting = delWorker(worker.broker.waiting, worker)
	delete(worker.broker.workers, worker.id_string)
}

//  The send method formats and sends a command to a worker. The caller may
//  also provide a command option, and a message payload:

func (worker *Worker) Send(command, option string, msg []string) (err error) {
	n := 4
	if option != "" {
		n++
	}
	m := make([]string, n, n+len(msg))
	m = append(m, msg...)

	//  Stack protocol envelope to start of message
	if option != "" {
		m[4] = option
	}
	m[3] = command
	m[2] = mdapi.MDPW_WORKER

	//  Stack routing envelope to start of message
	m[1] = ""
	m[0] = worker.identity

	if worker.broker.verbose {
		log.Printf("I: sending %s to worker %q\n", mdapi.MDPS_COMMANDS[command], m)
	}
	_, err = worker.broker.socket.SendMessage(m)
	return
}

//  This worker is now waiting for work

func (worker *Worker) Waiting() {
	//  Queue to broker and service waiting lists
	worker.broker.waiting = append(worker.broker.waiting, worker)
	worker.service.waiting = append(worker.service.waiting, worker)
	worker.expiry = time.Now().Add(HEARTBEAT_EXPIRY)
	worker.service.Dispatch([]string{})
}

//  Finally here is the main task. We create a new broker instance and
//  then processes messages on the broker socket:

func main() {
	verbose := false
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		verbose = true
	}

	broker, _ := NewBroker(verbose)
	broker.Bind("tcp://*:5555")

	poller := zmq.NewPoller()
	poller.Add(broker.socket, zmq.POLLIN)

	//  Get and process messages forever or until interrupted
	for {
		polled, err := poller.Poll(HEARTBEAT_INTERVAL)
		if err != nil {
			break //  Interrupted
		}

		//  Process next input message, if any
		if len(polled) > 0 {
			msg, err := broker.socket.RecvMessage(0)
			if err != nil {
				break //  Interrupted
			}
			if broker.verbose {
				log.Printf("I: received message: %q\n", msg)
			}
			sender, msg := popStr(msg)
			_, msg = popStr(msg)
			header, msg := popStr(msg)

			switch header {
			case mdapi.MDPC_CLIENT:
				broker.ClientMsg(sender, msg)
			case mdapi.MDPW_WORKER:
				broker.WorkerMsg(sender, msg)
			default:
				log.Printf("E: invalid message: %q\n", msg)
			}
		}
		//  Disconnect and delete any expired workers
		//  Send heartbeats to idle workers if needed
		if time.Now().After(broker.heartbeat_at) {
			broker.Purge()
			for _, worker := range broker.waiting {
				worker.Send(mdapi.MDPW_HEARTBEAT, "", []string{})
			}
			broker.heartbeat_at = time.Now().Add(HEARTBEAT_INTERVAL)
		}
	}
	log.Println("W: interrupt received, shutting down...")
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

func popStr(ss []string) (s string, ss2 []string) {
	s = ss[0]
	ss2 = ss[1:]
	return
}

func popMsg(msgs [][]string) (msg []string, msgs2 [][]string) {
	msg = msgs[0]
	msgs2 = msgs[1:]
	return
}

func popWorker(workers []*Worker) (worker *Worker, workers2 []*Worker) {
	worker = workers[0]
	workers2 = workers[1:]
	return
}

func delWorker(workers []*Worker, worker *Worker) []*Worker {
	for i := 0; i < len(workers); i++ {
		if workers[i] == worker {
			workers = append(workers[:i], workers[i+1:]...)
			i--
		}
	}
	return workers
}
