// Clone client API stack (multithreaded).
package clone

import (
	zmq "github.com/pebbe/zmq4"
	"github.com/pebbe/zmq4/examples/kvmsg"

	"fmt"
	"strconv"
	"time"
)

//  =====================================================================
//  Synchronous part, works in our application thread

//  ---------------------------------------------------------------------
//  Structure of our class

var (
	pipe_nmb uint64
)

type Clone struct {
	pipe *zmq.Socket //  Pipe through to clone agent
}

//  This is the thread that handles our real clone class

//  Constructor for the clone class. Note that we create
//  the pipe that connects our frontend to the
//  backend agent:

func New() (clone *Clone) {
	clone = &Clone{}
	clone.pipe, _ = zmq.NewSocket(zmq.PAIR)
	pipename := fmt.Sprint("inproc://pipe", pipe_nmb)
	pipe_nmb++
	clone.pipe.Bind(pipename)
	go clone_agent(pipename)
	return
}

//  Specify subtree for snapshot and updates, which we must do before
//  connecting to a server since the subtree specification is sent as
//  first command to the server. Sends a [SUBTREE][subtree] command to
//  the agent:

func (clone *Clone) Subtree(subtree string) {
	clone.pipe.SendMessage("SUBTREE", subtree)
}

//  Connect to a new server endpoint. We can connect to at most two
//  servers. Sends [CONNECT][endpoint][service] to the agent:

func (clone *Clone) Connect(address, service string) {
	clone.pipe.SendMessage("CONNECT", address, service)
}

//  Set a new value in the shared hashmap. Sends a [SET][key][value][ttl]
//  command through to the agent which does the actual work:

func (clone *Clone) Set(key, value string, ttl int) {
	clone.pipe.SendMessage("SET", key, value, ttl)
}

//  Look-up value in distributed hash table. Sends [GET][key] to the agent and
//  waits for a value response. If there is no value available, will eventually
//  return error:

func (clone *Clone) Get(key string) (value string, err error) {

	clone.pipe.SendMessage("GET", key)

	reply, e := clone.pipe.RecvMessage(0)
	if e != nil {
		err = e
		return
	}
	value = reply[0]
	return
}

//  The back-end agent manages a set of servers, which we implement using
//  our simple class model:

type server_t struct {
	address    string      //  Server address
	port       int         //  Server port
	snapshot   *zmq.Socket //  Snapshot socket
	subscriber *zmq.Socket //  Incoming updates
	expiry     time.Time   //  When server expires
	requests   int64       //  How many snapshot requests made?
}

func server_new(address string, port int, subtree string) (server *server_t) {
	server = &server_t{}

	fmt.Printf("I: adding server %s:%d...\n", address, port)
	server.address = address
	server.port = port

	server.snapshot, _ = zmq.NewSocket(zmq.DEALER)
	server.snapshot.Connect(fmt.Sprintf("%s:%d", address, port))
	server.subscriber, _ = zmq.NewSocket(zmq.SUB)
	server.subscriber.Connect(fmt.Sprintf("%s:%d", address, port+1))
	server.subscriber.SetSubscribe(subtree)
	return
}

//  Here is the implementation of the back-end agent itself:

const (
	//  Number of servers we will talk to
	server_MAX = 2

	//  Server considered dead if silent for this long
	server_TTL = 5000 * time.Millisecond
)

const (
	//  States we can be in
	state_INITIAL = iota //  Before asking server for state
	state_SYNCING        //  Getting state from server
	state_ACTIVE         //  Getting new updates from server
)

type agent_t struct {
	pipe        *zmq.Socket             //  Pipe back to application
	kvmap       map[string]*kvmsg.Kvmsg //  Actual key/value table
	subtree     string                  //  Subtree specification, if any
	server      [server_MAX]*server_t
	nbr_servers int         //  0 to SERVER_MAX
	state       int         //  Current state
	cur_server  int         //  If active, server 0 or 1
	sequence    int64       //  Last kvmsg processed
	publisher   *zmq.Socket //  Outgoing updates
}

func agent_new(pipe *zmq.Socket) (agent *agent_t) {
	agent = &agent_t{}
	agent.pipe = pipe
	agent.kvmap = make(map[string]*kvmsg.Kvmsg)
	agent.subtree = ""
	agent.state = state_INITIAL
	agent.publisher, _ = zmq.NewSocket(zmq.PUB)
	return
}

//  Here we handle the different control messages from the front-end;
//  SUBTREE, CONNECT, SET, and GET:

func (agent *agent_t) control_message() (err error) {
	msg, e := agent.pipe.RecvMessage(0)
	if e != nil {
		return e
	}
	command := msg[0]
	msg = msg[1:]

	switch command {
	case "SUBTREE":
		agent.subtree = msg[0]
	case "CONNECT":
		address := msg[0]
		service := msg[1]
		if agent.nbr_servers < server_MAX {
			serv, _ := strconv.Atoi(service)
			agent.server[agent.nbr_servers] = server_new(address, serv, agent.subtree)
			agent.nbr_servers++
			//  We broadcast updates to all known servers
			agent.publisher.Connect(fmt.Sprintf("%s:%d", address, serv+2))
		} else {
			fmt.Printf("E: too many servers (max. %d)\n", server_MAX)
		}
	case "SET":
		//  When we set a property, we push the new key-value pair onto
		//  all our connected servers:
		key := msg[0]
		value := msg[1]
		ttl := msg[2]

		//  Send key-value pair on to server
		kvmsg := kvmsg.NewKvmsg(0)
		kvmsg.SetKey(key)
		kvmsg.SetUuid()
		kvmsg.SetBody(value)
		kvmsg.SetProp("ttl", ttl)
		kvmsg.Store(agent.kvmap)
		kvmsg.Send(agent.publisher)
	case "GET":
		key := msg[0]
		value := ""
		if kvmsg, ok := agent.kvmap[key]; ok {
			value, _ = kvmsg.GetBody()
		}
		agent.pipe.SendMessage(value)
	}
	return
}

//  The asynchronous agent manages a server pool and handles the
//  request/reply dialog when the application asks for it:

func clone_agent(pipename string) {

	pipe, _ := zmq.NewSocket(zmq.PAIR)
	pipe.Connect(pipename)

	agent := agent_new(pipe)

LOOP:
	for {
		poller := zmq.NewPoller()
		poller.Add(pipe, zmq.POLLIN)
		server := agent.server[agent.cur_server]
		switch agent.state {
		case state_INITIAL:
			//  In this state we ask the server for a snapshot,
			//  if we have a server to talk to...
			if agent.nbr_servers > 0 {
				fmt.Printf("I: waiting for server at %s:%d...\n", server.address, server.port)
				if server.requests < 2 {
					server.snapshot.SendMessage("ICANHAZ?", agent.subtree)
					server.requests++
				}
				server.expiry = time.Now().Add(server_TTL)
				agent.state = state_SYNCING
				poller.Add(server.snapshot, zmq.POLLIN)
			}

		case state_SYNCING:
			//  In this state we read from snapshot and we expect
			//  the server to respond, else we fail over.
			poller.Add(server.snapshot, zmq.POLLIN)

		case state_ACTIVE:
			//  In this state we read from subscriber and we expect
			//  the server to give hugz, else we fail over.
			poller.Add(server.subscriber, zmq.POLLIN)
			break
		}
		poll_timer := time.Duration(-1)
		if server != nil {
			poll_timer = server.expiry.Sub(time.Now())
			if poll_timer < 0 {
				poll_timer = 0
			}
		}
		//  We're ready to process incoming messages; if nothing at all
		//  comes from our server within the timeout, that means the
		//  server is dead:

		polled, err := poller.Poll(poll_timer)
		if err != nil {
			break
		}

		if len(polled) > 0 {
			for _, item := range polled {
				switch socket := item.Socket; socket {
				case pipe:

					err = agent.control_message()
					if err != nil {
						break LOOP
					}

				default:
					kvmsg, e := kvmsg.RecvKvmsg(socket)
					if e != nil {
						err = e
						break LOOP
					}

					//  Anything from server resets its expiry time
					server.expiry = time.Now().Add(server_TTL)
					if agent.state == state_SYNCING {
						//  Store in snapshot until we're finished
						server.requests = 0
						if key, _ := kvmsg.GetKey(); key == "KTHXBAI" {
							agent.sequence, _ = kvmsg.GetSequence()
							agent.state = state_ACTIVE
							fmt.Printf("I: received from %s:%d snapshot=%d\n", server.address, server.port, agent.sequence)
						} else {
							kvmsg.Store(agent.kvmap)
						}
					} else if agent.state == state_ACTIVE {
						//  Discard out-of-sequence updates, incl. hugz
						if seq, _ := kvmsg.GetSequence(); seq > agent.sequence {
							agent.sequence = seq
							kvmsg.Store(agent.kvmap)
							fmt.Printf("I: received from %s:%d update=%d\n", server.address, server.port, agent.sequence)
						}
					}
				}
			}
		} else {
			//  Server has died, failover to next
			fmt.Printf("I: server at %s:%d didn't give hugz\n", server.address, server.port)
			agent.cur_server = (agent.cur_server + 1) % agent.nbr_servers
			agent.state = state_INITIAL
		}
	}
}
