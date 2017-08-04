//  flcliapi - Freelance Pattern agent class.
// Implements the Freelance Protocol at http://rfc.zeromq.org/spec:10.
package flcliapi

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"strconv"
	"time"
)

const (
	//  If no server replies within this time, abandon request
	GLOBAL_TIMEOUT = 3000 * time.Millisecond
	//  PING interval for servers we think are alive
	PING_INTERVAL = 2000 * time.Millisecond
	//  Server considered dead if silent for this long
	SERVER_TTL = 6000 * time.Millisecond
)

//  This API works in two halves, a common pattern for APIs that need to
//  run in the background. One half is an front-end object our application
//  creates and works with; the other half is a back-end "agent" that runs
//  in a background thread. The front-end talks to the back-end over an
//  inproc pipe socket:

//  ---------------------------------------------------------------------
//  Structure of our front-end class

type Flcliapi struct {
	pipe *zmq.Socket //  Pipe through to flcliapi agent
}

//  This is the thread that handles our real flcliapi class

//  ---------------------------------------------------------------------
//  Constructor

func New() (flcliapi *Flcliapi) {
	flcliapi = &Flcliapi{}
	flcliapi.pipe, _ = zmq.NewSocket(zmq.PAIR)
	flcliapi.pipe.Bind("inproc://pipe")
	go flcliapi_agent()
	return
}

//  To implement the connect method, the front-end object sends a multi-part
//  message to the back-end agent. The first part is a string "CONNECT", and
//  the second part is the endpoint. It waits 100msec for the connection to
//  come up, which isn't pretty, but saves us from sending all requests to a
//  single server, at start-up time:

func (flcliapi *Flcliapi) Connect(endpoint string) {
	flcliapi.pipe.SendMessage("CONNECT", endpoint)
	time.Sleep(100 * time.Millisecond) //  Allow connection to come up
}

//  To implement the request method, the front-end object sends a message
//  to the back-end, specifying a command "REQUEST" and the request message:

func (flcliapi *Flcliapi) Request(request []string) (reply []string, err error) {
	flcliapi.pipe.SendMessage("REQUEST", request)
	reply, err = flcliapi.pipe.RecvMessage(0)
	if err == nil {
		status := reply[0]
		reply = reply[1:]
		if status == "FAILED" {
			reply = reply[0:0]
		}
	}
	return
}

//  Here we see the back-end agent. It runs as an attached thread, talking
//  to its parent over a pipe socket. It is a fairly complex piece of work
//  so we'll break it down into pieces. First, the agent manages a set of
//  servers, using our familiar class approach:

//  ---------------------------------------------------------------------
//  Simple class for one server we talk to

type server_t struct {
	endpoint string    //  Server identity/endpoint
	alive    bool      //  true if known to be alive
	ping_at  time.Time //  Next ping at this time
	expires  time.Time //  Expires at this time
}

func server_new(endpoint string) (server *server_t) {
	server = &server_t{
		endpoint: endpoint,
		alive:    false,
		ping_at:  time.Now().Add(PING_INTERVAL),
		expires:  time.Now().Add(SERVER_TTL),
	}
	return
}

func (server *server_t) ping(socket *zmq.Socket) {
	if time.Now().After(server.ping_at) {
		socket.SendMessage(server.endpoint, "PING")
		server.ping_at = time.Now().Add(PING_INTERVAL)
	}
}

func (server *server_t) tickless(t time.Time) time.Time {
	if t.After(server.ping_at) {
		return server.ping_at
	}
	return t
}

//  We build the agent as a class that's capable of processing messages
//  coming in from its various sockets:

//  ---------------------------------------------------------------------
//  Simple class for one background agent

type agent_t struct {
	pipe     *zmq.Socket          //  Socket to talk back to application
	router   *zmq.Socket          //  Socket to talk to servers
	servers  map[string]*server_t //  Servers we've connected to
	actives  []*server_t          //  Servers we know are alive
	sequence int                  //  Number of requests ever sent
	request  []string             //  Current request if any
	reply    []string             //  Current reply if any
	expires  time.Time            //  Timeout for request/reply
}

func agent_new() (agent *agent_t) {
	agent = &agent_t{
		servers: make(map[string]*server_t),
		actives: make([]*server_t, 0),
		request: make([]string, 0),
		reply:   make([]string, 0),
	}
	agent.pipe, _ = zmq.NewSocket(zmq.PAIR)
	agent.pipe.Connect("inproc://pipe")
	agent.router, _ = zmq.NewSocket(zmq.ROUTER)
	return
}

//  The control_message method processes one message from our front-end
//  class (it's going to be CONNECT or REQUEST):

func (agent *agent_t) control_message() {
	msg, _ := agent.pipe.RecvMessage(0)
	command := msg[0]
	msg = msg[1:]

	switch command {
	case "CONNECT":
		endpoint := msg[0]
		fmt.Printf("I: connecting to %s...\n", endpoint)
		err := agent.router.Connect(endpoint)
		if err != nil {
			panic("agent.router.Connect(endpoint) failed")
		}
		server := server_new(endpoint)
		agent.servers[endpoint] = server
		agent.actives = append(agent.actives, server)
		server.ping_at = time.Now().Add(PING_INTERVAL)
		server.expires = time.Now().Add(SERVER_TTL)
	case "REQUEST":
		if len(agent.request) > 0 {
			panic("len(agent.request) > 0") //  Strict request-reply cycle
		}
		//  Prefix request with sequence number --(and empty envelope)--
		agent.request = make([]string, 1, 1+len(msg))
		agent.sequence++
		agent.request[0] = fmt.Sprint(agent.sequence)
		agent.request = append(agent.request, msg...)
		//  Request expires after global timeout
		agent.expires = time.Now().Add(GLOBAL_TIMEOUT)
	}
}

//  The router_message method processes one message from a connected
//  server:

func (agent *agent_t) router_message() {
	reply, _ := agent.router.RecvMessage(0)

	//  Frame 0 is server that replied
	endpoint := reply[0]
	reply = reply[1:]
	server, ok := agent.servers[endpoint]
	if !ok {
		panic("No server for endpoint")
	}
	if !server.alive {
		agent.actives = append(agent.actives, server)
		server.alive = true
	}
	server.ping_at = time.Now().Add(PING_INTERVAL)
	server.expires = time.Now().Add(SERVER_TTL)

	//  Frame 1 may be sequence number for reply
	sequence, _ := strconv.Atoi(reply[0])
	reply = reply[1:]
	if sequence == agent.sequence {
		agent.pipe.SendMessage("OK", reply)
		agent.request = agent.request[0:0]
	}
}

//  Finally here's the agent task itself, which polls its two sockets
//  and processes incoming messages:

func flcliapi_agent() {

	agent := agent_new()

	poller := zmq.NewPoller()
	poller.Add(agent.pipe, zmq.POLLIN)
	poller.Add(agent.router, zmq.POLLIN)
	for {
		//  Calculate tickless timer, up to 1 hour
		tickless := time.Now().Add(time.Hour)
		if len(agent.request) > 0 && tickless.After(agent.expires) {
			tickless = agent.expires
		}
		for key := range agent.servers {
			tickless = agent.servers[key].tickless(tickless)
		}

		polled, err := poller.Poll(tickless.Sub(time.Now()))
		if err != nil {
			break //  Context has been shut down
		}

		for _, item := range polled {
			switch item.Socket {
			case agent.pipe:
				agent.control_message()
			case agent.router:
				agent.router_message()
			}
		}

		//  If we're processing a request, dispatch to next server
		if len(agent.request) > 0 {
			if time.Now().After(agent.expires) {
				//  Request expired, kill it
				agent.pipe.SendMessage("FAILED")
				agent.request = agent.request[0:0]
			} else {
				//  Find server to talk to, remove any expired ones
				for len(agent.actives) > 0 {
					server := agent.actives[0]
					if time.Now().After(server.expires) {
						agent.actives = agent.actives[1:]
						server.alive = false
					} else {
						agent.router.SendMessage(server.endpoint, agent.request)
						break
					}
				}
			}
		}
		//  --(Disconnect and delete any expired servers)--
		//  Send heartbeats to idle servers if needed
		for key := range agent.servers {
			agent.servers[key].ping(agent.router)
		}
	}
}
