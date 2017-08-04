//
//  Binary Star server proof-of-concept implementation. This server does no
//  real work; it just demonstrates the Binary Star failover model.

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"os"
	"strconv"
	"time"
)

//  States we can be in at any point in time
type state_t int

const (
	_             = state_t(iota)
	STATE_PRIMARY //  Primary, waiting for peer to connect
	STATE_BACKUP  //  Backup, waiting for peer to connect
	STATE_ACTIVE  //  Active - accepting connections
	STATE_PASSIVE //  Passive - not accepting connections
)

//  Events, which start with the states our peer can be in
type event_t int

const (
	_              = event_t(iota)
	PEER_PRIMARY   //  HA peer is pending primary
	PEER_BACKUP    //  HA peer is pending backup
	PEER_ACTIVE    //  HA peer is active
	PEER_PASSIVE   //  HA peer is passive
	CLIENT_REQUEST //  Client makes request
)

//  Our finite state machine
type bstar_t struct {
	state       state_t   //  Current state
	event       event_t   //  Current event
	peer_expiry time.Time //  When peer is considered 'dead'
}

//  We send state information every this often
//  If peer doesn't respond in two heartbeats, it is 'dead'
const (
	HEARTBEAT = 1000 * time.Millisecond //  In msecs
)

//  The heart of the Binary Star design is its finite-state machine (FSM).
//  The FSM runs one event at a time. We apply an event to the current state,
//  which checks if the event is accepted, and if so sets a new state:

func StateMachine(fsm *bstar_t) (exception bool) {
	//  These are the PRIMARY and BACKUP states; we're waiting to become
	//  ACTIVE or PASSIVE depending on events we get from our peer:
	if fsm.state == STATE_PRIMARY {
		if fsm.event == PEER_BACKUP {
			fmt.Println("I: connected to backup (passive), ready as active")
			fsm.state = STATE_ACTIVE
		} else if fsm.event == PEER_ACTIVE {
			fmt.Println("I: connected to backup (active), ready as passive")
			fsm.state = STATE_PASSIVE
		}
		//  Accept client connections
	} else if fsm.state == STATE_BACKUP {
		if fsm.event == PEER_ACTIVE {
			fmt.Println("I: connected to primary (active), ready as passive")
			fsm.state = STATE_PASSIVE
		} else if fsm.event == CLIENT_REQUEST {
			//  Reject client connections when acting as backup
			exception = true
		}
	} else if fsm.state == STATE_ACTIVE {
		//  These are the ACTIVE and PASSIVE states:
		if fsm.event == PEER_ACTIVE {
			//  Two actives would mean split-brain
			fmt.Println("E: fatal error - dual actives, aborting")
			exception = true
		}
	} else if fsm.state == STATE_PASSIVE {
		//  Server is passive
		//  CLIENT_REQUEST events can trigger failover if peer looks dead
		if fsm.event == PEER_PRIMARY {
			//  Peer is restarting - become active, peer will go passive
			fmt.Println("I: primary (passive) is restarting, ready as active")
			fsm.state = STATE_ACTIVE
		} else if fsm.event == PEER_BACKUP {
			//  Peer is restarting - become active, peer will go passive
			fmt.Println("I: backup (passive) is restarting, ready as active")
			fsm.state = STATE_ACTIVE
		} else if fsm.event == PEER_PASSIVE {
			//  Two passives would mean cluster would be non-responsive
			fmt.Println("E: fatal error - dual passives, aborting")
			exception = true
		} else if fsm.event == CLIENT_REQUEST {
			//  Peer becomes active if timeout has passed
			//  It's the client request that triggers the failover
			if time.Now().After(fsm.peer_expiry) {
				//  If peer is dead, switch to the active state
				fmt.Println("I: failover successful, ready as active")
				fsm.state = STATE_ACTIVE
			} else {
				//  If peer is alive, reject connections
				exception = true
			}
		}
	}
	return
}

//  This is our main task. First we bind/connect our sockets with our
//  peer and make sure we will get state messages correctly. We use
//  three sockets; one to publish state, one to subscribe to state, and
//  one for client requests/replies:

func main() {
	//  Arguments can be either of:
	//      -p  primary server, at tcp://localhost:5001
	//      -b  backup server, at tcp://localhost:5002
	statepub, _ := zmq.NewSocket(zmq.PUB)
	statesub, _ := zmq.NewSocket(zmq.SUB)
	statesub.SetSubscribe("")
	frontend, _ := zmq.NewSocket(zmq.ROUTER)
	fsm := &bstar_t{peer_expiry: time.Now().Add(2 * HEARTBEAT)}

	if len(os.Args) == 2 && os.Args[1] == "-p" {
		fmt.Println("I: Primary active, waiting for backup (passive)")
		frontend.Bind("tcp://*:5001")
		statepub.Bind("tcp://*:5003")
		statesub.Connect("tcp://localhost:5004")
		fsm.state = STATE_PRIMARY
	} else if len(os.Args) == 2 && os.Args[1] == "-b" {
		fmt.Println("I: Backup passive, waiting for primary (active)")
		frontend.Bind("tcp://*:5002")
		statepub.Bind("tcp://*:5004")
		statesub.Connect("tcp://localhost:5003")
		fsm.state = STATE_BACKUP
	} else {
		fmt.Println("Usage: bstarsrv { -p | -b }")
		return
	}
	//  We now process events on our two input sockets, and process these
	//  events one at a time via our finite-state machine. Our "work" for
	//  a client request is simply to echo it back:

	//  Set timer for next outgoing state message
	send_state_at := time.Now().Add(HEARTBEAT)

	poller := zmq.NewPoller()
	poller.Add(frontend, zmq.POLLIN)
	poller.Add(statesub, zmq.POLLIN)

LOOP:
	for {
		time_left := send_state_at.Sub(time.Now())
		if time_left < 0 {
			time_left = 0
		}
		polled, err := poller.Poll(time_left)
		if err != nil {
			break //  Context has been shut down
		}
		for _, socket := range polled {
			switch socket.Socket {
			case frontend:
				//  Have a client request
				msg, _ := frontend.RecvMessage(0)
				fsm.event = CLIENT_REQUEST
				if !StateMachine(fsm) {
					//  Answer client by echoing request back
					frontend.SendMessage(msg)
				}
			case statesub:
				//  Have state from our peer, execute as event
				message, _ := statesub.RecvMessage(0)
				i, _ := strconv.Atoi(message[0])
				fsm.event = event_t(i)
				if StateMachine(fsm) {
					break LOOP //  Error, so exit
				}
				fsm.peer_expiry = time.Now().Add(2 * HEARTBEAT)
			}
		}
		//  If we timed-out, send state to peer
		if time.Now().After(send_state_at) {
			statepub.SendMessage(int(fsm.state))
			send_state_at = time.Now().Add(HEARTBEAT)
		}
	}
	fmt.Println("W: interrupted")
}
