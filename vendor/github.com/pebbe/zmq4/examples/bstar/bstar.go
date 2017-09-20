//  bstar - Binary Star reactor.
package bstar

import (
	zmq "github.com/pebbe/zmq4"

	"errors"
	"log"
	"strconv"
	"time"
)

const (
	PRIMARY = true
	BACKUP  = false
)

//  States we can be in at any point in time
type state_t int

const (
	_             = state_t(iota)
	state_PRIMARY //  Primary, waiting for peer to connect
	state_BACKUP  //  Backup, waiting for peer to connect
	state_ACTIVE  //  Active - accepting connections
	state_PASSIVE //  Passive - not accepting connections
)

//  Events, which start with the states our peer can be in
type event_t int

const (
	_              = event_t(iota)
	peer_PRIMARY   //  HA peer is pending primary
	peer_BACKUP    //  HA peer is pending backup
	peer_ACTIVE    //  HA peer is active
	peer_PASSIVE   //  HA peer is passive
	client_REQUEST //  Client makes request
)

//  Structure of our class

type Bstar struct {
	Reactor     *zmq.Reactor            //  Reactor loop
	statepub    *zmq.Socket             //  State publisher
	statesub    *zmq.Socket             //  State subscriber
	state       state_t                 //  Current state
	event       event_t                 //  Current event
	peer_expiry time.Time               //  When peer is considered 'dead'
	voter_fn    func(*zmq.Socket) error //  Voting socket handler
	active_fn   func() error            //  Call when become active
	passive_fn  func() error            //  Call when become passive
}

//  The finite-state machine is the same as in the proof-of-concept server.
//  To understand this reactor in detail, first read the CZMQ zloop class.

//  We send state information every this often
//  If peer doesn't respond in two heartbeats, it is 'dead'
const (
	bstar_HEARTBEAT = 1000 * time.Millisecond //  In msecs
)

//  ---------------------------------------------------------------------
//  Binary Star finite state machine (applies event to state)
//  Returns error if there was an exception, nil if event was valid.

func (bstar *Bstar) execute_fsm() (exception error) {
	//  Primary server is waiting for peer to connect
	//  Accepts client_REQUEST events in this state
	if bstar.state == state_PRIMARY {
		if bstar.event == peer_BACKUP {
			log.Println("I: connected to backup (passive), ready as active")
			bstar.state = state_ACTIVE
			if bstar.active_fn != nil {
				bstar.active_fn()
			}
		} else if bstar.event == peer_ACTIVE {
			log.Println("I: connected to backup (active), ready as passive")
			bstar.state = state_PASSIVE
			if bstar.passive_fn != nil {
				bstar.passive_fn()
			}
		} else if bstar.event == client_REQUEST {
			// Allow client requests to turn us into the active if we've
			// waited sufficiently long to believe the backup is not
			// currently acting as active (i.e., after a failover)
			if time.Now().After(bstar.peer_expiry) {
				log.Println("I: request from client, ready as active")
				bstar.state = state_ACTIVE
				if bstar.active_fn != nil {
					bstar.active_fn()
				}
			} else {
				// Don't respond to clients yet - it's possible we're
				// performing a failback and the backup is currently active
				exception = errors.New("Exception")
			}
		}
	} else if bstar.state == state_BACKUP {
		//  Backup server is waiting for peer to connect
		//  Rejects client_REQUEST events in this state
		if bstar.event == peer_ACTIVE {
			log.Println("I: connected to primary (active), ready as passive")
			bstar.state = state_PASSIVE
			if bstar.passive_fn != nil {
				bstar.passive_fn()
			}
		} else if bstar.event == client_REQUEST {
			exception = errors.New("Exception")
		}
	} else if bstar.state == state_ACTIVE {
		//  Server is active
		//  Accepts client_REQUEST events in this state
		//  The only way out of ACTIVE is death
		if bstar.event == peer_ACTIVE {
			//  Two actives would mean split-brain
			log.Println("E: fatal error - dual actives, aborting")
			exception = errors.New("Exception")
		}
	} else if bstar.state == state_PASSIVE {
		//  Server is passive
		//  client_REQUEST events can trigger failover if peer looks dead
		if bstar.event == peer_PRIMARY {
			//  Peer is restarting - become active, peer will go passive
			log.Println("I: primary (passive) is restarting, ready as active")
			bstar.state = state_ACTIVE
		} else if bstar.event == peer_BACKUP {
			//  Peer is restarting - become active, peer will go passive
			log.Println("I: backup (passive) is restarting, ready as active")
			bstar.state = state_ACTIVE
		} else if bstar.event == peer_PASSIVE {
			//  Two passives would mean cluster would be non-responsive
			log.Println("E: fatal error - dual passives, aborting")
			exception = errors.New("Exception")
		} else if bstar.event == client_REQUEST {
			//  Peer becomes active if timeout has passed
			//  It's the client request that triggers the failover
			if time.Now().After(bstar.peer_expiry) {
				//  If peer is dead, switch to the active state
				log.Println("I: failover successful, ready as active")
				bstar.state = state_ACTIVE
			} else {
				//  If peer is alive, reject connections
				exception = errors.New("Exception")
			}
		}
		//  Call state change handler if necessary
		if bstar.state == state_ACTIVE && bstar.active_fn != nil {
			bstar.active_fn()
		}
	}
	return
}

func (bstar *Bstar) update_peer_expiry() {
	bstar.peer_expiry = time.Now().Add(2 * bstar_HEARTBEAT)
}

//  ---------------------------------------------------------------------
//  Reactor event handlers...

//  Publish our state to peer
func (bstar *Bstar) send_state() (err error) {
	_, err = bstar.statepub.SendMessage(int(bstar.state))
	return
}

//  Receive state from peer, execute finite state machine
func (bstar *Bstar) recv_state() (err error) {
	msg, err := bstar.statesub.RecvMessage(0)
	if err == nil {
		e, _ := strconv.Atoi(msg[0])
		bstar.event = event_t(e)
	}
	return bstar.execute_fsm()
}

//  Application wants to speak to us, see if it's possible
func (bstar *Bstar) voter_ready(socket *zmq.Socket) error {
	//  If server can accept input now, call appl handler
	bstar.event = client_REQUEST
	err := bstar.execute_fsm()
	if err == nil {
		bstar.voter_fn(socket)
	} else {
		//  Destroy waiting message, no-one to read it
		socket.RecvMessage(0)
	}
	return nil
}

//  This is the constructor for our bstar class. We have to tell it whether
//  we're primary or backup server, and our local and remote endpoints to
//  bind and connect to:

func New(primary bool, local, remote string) (bstar *Bstar, err error) {

	bstar = &Bstar{}

	//  Initialize the Binary Star
	bstar.Reactor = zmq.NewReactor()
	if primary {
		bstar.state = state_PRIMARY
	} else {
		bstar.state = state_BACKUP
	}

	//  Create publisher for state going to peer
	bstar.statepub, err = zmq.NewSocket(zmq.PUB)
	bstar.statepub.Bind(local)

	//  Create subscriber for state coming from peer
	bstar.statesub, err = zmq.NewSocket(zmq.SUB)
	bstar.statesub.SetSubscribe("")
	bstar.statesub.Connect(remote)

	//  Set-up basic reactor events
	bstar.Reactor.AddChannelTime(time.Tick(bstar_HEARTBEAT), 1,
		func(i interface{}) error { return bstar.send_state() })
	bstar.Reactor.AddSocket(bstar.statesub, zmq.POLLIN,
		func(e zmq.State) error { return bstar.recv_state() })

	return
}

//  The voter method registers a client voter socket. Messages received
//  on this socket provide the client_REQUEST events for the Binary Star
//  FSM and are passed to the provided application handler. We require
//  exactly one voter per bstar instance:

func (bstar *Bstar) Voter(endpoint string, socket_type zmq.Type, handler func(*zmq.Socket) error) {
	//  Hold actual handler so we can call this later
	socket, _ := zmq.NewSocket(socket_type)
	socket.Bind(endpoint)
	if bstar.voter_fn != nil {
		panic("Double voter function")
	}
	bstar.voter_fn = handler
	bstar.Reactor.AddSocket(socket, zmq.POLLIN,
		func(e zmq.State) error { return bstar.voter_ready(socket) })
}

//  Register handlers to be called each time there's a state change:

func (bstar *Bstar) NewActive(handler func() error) {
	if bstar.active_fn != nil {
		panic("Double Active")
	}
	bstar.active_fn = handler
}

func (bstar *Bstar) NewPassive(handler func() error) {
	if bstar.passive_fn != nil {
		panic("Double Passive")
	}
	bstar.passive_fn = handler
}

//  Enable/disable verbose tracing, for debugging:

func (bstar *Bstar) SetVerbose(verbose bool) {
	bstar.Reactor.SetVerbose(verbose)
}

//?  Finally, start the configured reactor. It will end if any handler
//?  returns error to the reactor, or if the process receives SIGINT or SIGTERM:

func (bstar *Bstar) Start() error {
	if bstar.voter_fn == nil {
		panic("Missing voter function")
	}
	bstar.update_peer_expiry()
	return bstar.Reactor.Run(bstar_HEARTBEAT / 5)
}
