//  Interface class for Chapter 8.
//  This implements an "interface" to our network of nodes.
package intface

import (
	zmq "github.com/pebbe/zmq4"

	"github.com/pborman/uuid"

	"bytes"
	"errors"
	"net"
	"time"
)

//  =====================================================================
//  Synchronous part, works in our application thread

//  ---------------------------------------------------------------------
//  Structure of our class

type Intface struct {
	pipe *zmq.Socket //  Pipe through to agent
}

//  This is the thread that handles our real interface class

//  Here is the constructor for the interface class.
//  Note that the class has barely any properties, it is just an excuse
//  to start the background thread, and a wrapper around zmsg_recv():

func New() (iface *Intface) {
	iface = &Intface{}
	var err error
	iface.pipe, err = zmq.NewSocket(zmq.PAIR)
	if err != nil {
		panic(err)
	}
	err = iface.pipe.Bind("inproc://iface")
	if err != nil {
		panic(err)
	}
	go iface.agent()
	time.Sleep(100 * time.Millisecond)
	return
}

//  Here we wait for a message from the interface. This returns
//  us a []string, or error if interrupted:

func (iface *Intface) Recv() (msg []string, err error) {
	msg, err = iface.pipe.RecvMessage(0)
	return
}

//  =====================================================================
// //  Asynchronous part, works in the background

//  This structure defines each peer that we discover and track:

type peer_t struct {
	uuid_bytes  []byte
	uuid_string string
	expires_at  time.Time
}

const (
	PING_PORT_NUMBER = 9999
	PING_INTERVAL    = 1000 * time.Millisecond //  Once per second
	PEER_EXPIRY      = 5000 * time.Millisecond //  Five seconds and it's gone
)

//  We have a constructor for the peer class:

func new_peer(uuid uuid.UUID) (peer *peer_t) {
	peer = &peer_t{
		uuid_bytes:  []byte(uuid),
		uuid_string: uuid.String(),
	}
	return
}

//  Just resets the peers expiry time; we call this method
//  whenever we get any activity from a peer.

func (peer *peer_t) is_alive() {
	peer.expires_at = time.Now().Add(PEER_EXPIRY)
}

//  This structure holds the context for our agent, so we can
//  pass that around cleanly to methods which need it:

type agent_t struct {
	pipe        *zmq.Socket //  Pipe back to application
	udp         *zmq.Socket
	conn        *net.UDPConn
	uuid_bytes  []byte //  Our UUID
	uuid_string string
	peers       map[string]*peer_t //  Hash of known peers, fast lookup
}

//  Now the constructor for our agent. Each interface
//  has one agent object, which implements its background thread:

func new_agent() (agent *agent_t) {

	// push output from udp into zmq socket
	bcast := &net.UDPAddr{Port: PING_PORT_NUMBER, IP: net.IPv4bcast}
	conn, e := net.ListenUDP("udp", bcast)
	if e != nil {
		panic(e)
	}
	go func() {
		buffer := make([]byte, 1024)
		udp, _ := zmq.NewSocket(zmq.PAIR)
		udp.Bind("inproc://udp")
		for {
			if n, _, err := conn.ReadFrom(buffer); err == nil {
				udp.SendBytes(buffer[:n], 0)
			}
		}
	}()
	time.Sleep(100 * time.Millisecond)

	pipe, _ := zmq.NewSocket(zmq.PAIR)
	pipe.Connect("inproc://iface")
	udp, _ := zmq.NewSocket(zmq.PAIR)
	udp.Connect("inproc://udp")

	uuid := uuid.NewRandom()
	agent = &agent_t{
		pipe:        pipe,
		udp:         udp,
		conn:        conn,
		uuid_bytes:  []byte(uuid),
		uuid_string: uuid.String(),
		peers:       make(map[string]*peer_t),
	}

	return
}

//  Here we handle the different control messages from the front-end.

func (agent *agent_t) control_message() (err error) {
	//  Get the whole message off the pipe in one go
	msg, e := agent.pipe.RecvMessage(0)
	if e != nil {
		return e
	}
	command := msg[0]

	//  We don't actually implement any control commands yet
	//  but if we did, this would be where we did it..
	switch command {
	case "EXAMPLE":
	default:
	}

	return
}

//  This is how we handle a beacon coming into our UDP socket;
//  this may be from other peers or an echo of our own broadcast
//  beacon:

func (agent *agent_t) handle_beacon() (err error) {

	msg, err := agent.udp.RecvMessage(0)
	if len(msg[0]) != 16 {
		return errors.New("Not a uuid")
	}

	//  If we got a UUID and it's not our own beacon, we have a peer
	uuid := uuid.UUID(msg[0])
	if bytes.Compare(uuid, agent.uuid_bytes) != 0 {
		//  Find or create peer via its UUID string
		uuid_string := uuid.String()
		peer, ok := agent.peers[uuid_string]
		if !ok {
			peer = new_peer(uuid)
			agent.peers[uuid_string] = peer

			//  Report peer joined the network
			agent.pipe.SendMessage("JOINED", uuid_string)
		}
		//  Any activity from the peer means it's alive
		peer.is_alive()
	}
	return
}

//  This method checks one peer item for expiry; if the peer hasn't
//  sent us anything by now, it's 'dead' and we can delete it:

func (agent *agent_t) reap_peer(peer *peer_t) {
	if time.Now().After(peer.expires_at) {
		//  Report peer left the network
		agent.pipe.SendMessage("LEFT", peer.uuid_string)
		delete(agent.peers, peer.uuid_string)
	}
}

//  This is the main loop for the background agent. It uses zmq_poll
//  to monitor the front-end pipe (commands from the API) and the
//  back-end UDP handle (beacons):

func (iface *Intface) agent() {
	//  Create agent instance to pass around
	agent := new_agent()

	//  Send first beacon immediately
	ping_at := time.Now()

	poller := zmq.NewPoller()
	poller.Add(agent.pipe, zmq.POLLIN)
	poller.Add(agent.udp, zmq.POLLIN)

	bcast := &net.UDPAddr{Port: PING_PORT_NUMBER, IP: net.IPv4bcast}
	for {
		timeout := ping_at.Add(time.Millisecond).Sub(time.Now())
		if timeout < 0 {
			timeout = 0
		}
		polled, err := poller.Poll(timeout)
		if err != nil {
			break
		}

		for _, item := range polled {
			switch socket := item.Socket; socket {
			case agent.pipe:
				//  If we had activity on the pipe, go handle the control
				//  message. Current code never sends control messages.
				agent.control_message()

			case agent.udp:
				//  If we had input on the UDP socket, go process that
				agent.handle_beacon()
			}
		}

		//  If we passed the 1-second mark, broadcast our beacon
		now := time.Now()
		if now.After(ping_at) {
			agent.conn.WriteTo(agent.uuid_bytes, bcast)
			ping_at = now.Add(PING_INTERVAL)
		}
		//  Delete and report any expired peers
		for _, peer := range agent.peers {
			agent.reap_peer(peer)
		}
	}
}
