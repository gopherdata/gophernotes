// Majordomo Protocol Worker API.
// Implements the MDP/Worker spec at http://rfc.zeromq.org/spec:7.

package mdapi

import (
	zmq "github.com/pebbe/zmq4"

	"log"
	"runtime"
	"time"
)

const (
	heartbeat_liveness = 3 //  3-5 is reasonable
)

//  This is the structure of a worker API instance. We use a pseudo-OO
//  approach in a lot of the C examples, as well as the CZMQ binding:

//  Structure of our class
//  We access these properties only via class methods

// Majordomo Protocol Worker API.
type Mdwrk struct {
	broker  string
	service string
	worker  *zmq.Socket //  Socket to broker
	poller  *zmq.Poller
	verbose bool //  Print activity to stdout

	//  Heartbeat management
	heartbeat_at time.Time     //  When to send HEARTBEAT
	liveness     int           //  How many attempts left
	heartbeat    time.Duration //  Heartbeat delay, msecs
	reconnect    time.Duration //  Reconnect delay, msecs

	expect_reply bool   //  False only at start
	reply_to     string //  Return identity, if any
}

//  We have two utility functions; to send a message to the broker and
//  to (re-)connect to the broker:

//  ---------------------------------------------------------------------

//  Send message to broker.
func (mdwrk *Mdwrk) SendToBroker(command string, option string, msg []string) (err error) {

	n := 3
	if option != "" {
		n++
	}
	m := make([]string, n, n+len(msg))
	m = append(m, msg...)

	//  Stack protocol envelope to start of message
	if option != "" {
		m[3] = option
	}
	m[2] = command
	m[1] = MDPW_WORKER
	m[0] = ""

	if mdwrk.verbose {
		log.Printf("I: sending %s to broker %q\n", MDPS_COMMANDS[command], m)
	}
	_, err = mdwrk.worker.SendMessage(m)
	return
}

//  ---------------------------------------------------------------------

//  Connect or reconnect to broker.
func (mdwrk *Mdwrk) ConnectToBroker() (err error) {
	if mdwrk.worker != nil {
		mdwrk.worker.Close()
		mdwrk.worker = nil
	}
	mdwrk.worker, err = zmq.NewSocket(zmq.DEALER)
	err = mdwrk.worker.Connect(mdwrk.broker)
	if mdwrk.verbose {
		log.Printf("I: connecting to broker at %s...\n", mdwrk.broker)
	}
	mdwrk.poller = zmq.NewPoller()
	mdwrk.poller.Add(mdwrk.worker, zmq.POLLIN)

	//  Register service with broker
	err = mdwrk.SendToBroker(MDPW_READY, mdwrk.service, []string{})

	//  If liveness hits zero, queue is considered disconnected
	mdwrk.liveness = heartbeat_liveness
	mdwrk.heartbeat_at = time.Now().Add(mdwrk.heartbeat)

	return
}

//  Here we have the constructor and destructor for our mdwrk class:

//  ---------------------------------------------------------------------
//  Constructor

func NewMdwrk(broker, service string, verbose bool) (mdwrk *Mdwrk, err error) {

	mdwrk = &Mdwrk{
		broker:    broker,
		service:   service,
		verbose:   verbose,
		heartbeat: 2500 * time.Millisecond,
		reconnect: 2500 * time.Millisecond,
	}

	err = mdwrk.ConnectToBroker()

	runtime.SetFinalizer(mdwrk, (*Mdwrk).Close)

	return
}

//  ---------------------------------------------------------------------
//  Destructor

func (mdwrk *Mdwrk) Close() {
	if mdwrk.worker != nil {
		mdwrk.worker.Close()
		mdwrk.worker = nil
	}
}

//  We provide two methods to configure the worker API. You can set the
//  heartbeat interval and retries to match the expected network performance.

//  ---------------------------------------------------------------------

//  Set heartbeat delay.
func (mdwrk *Mdwrk) SetHeartbeat(heartbeat time.Duration) {
	mdwrk.heartbeat = heartbeat
}

//  ---------------------------------------------------------------------

//  Set reconnect delay.
func (mdwrk *Mdwrk) SetReconnect(reconnect time.Duration) {
	mdwrk.reconnect = reconnect
}

//  This is the recv method; it's a little misnamed since it first sends
//  any reply and then waits for a new request. If you have a better name
//  for this, let me know:

//  ---------------------------------------------------------------------

//  Send reply, if any, to broker and wait for next request.
func (mdwrk *Mdwrk) Recv(reply []string) (msg []string, err error) {
	//  Format and send the reply if we were provided one
	if len(reply) == 0 && mdwrk.expect_reply {
		panic("No reply, expected")
	}
	if len(reply) > 0 {
		if mdwrk.reply_to == "" {
			panic("mdwrk.reply_to == \"\"")
		}
		m := make([]string, 2, 2+len(reply))
		m = append(m, reply...)
		m[0] = mdwrk.reply_to
		m[1] = ""
		err = mdwrk.SendToBroker(MDPW_REPLY, "", m)
	}
	mdwrk.expect_reply = true

	for {
		var polled []zmq.Polled
		polled, err = mdwrk.poller.Poll(mdwrk.heartbeat)
		if err != nil {
			break //  Interrupted
		}

		if len(polled) > 0 {
			msg, err = mdwrk.worker.RecvMessage(0)
			if err != nil {
				break //  Interrupted
			}
			if mdwrk.verbose {
				log.Printf("I: received message from broker: %q\n", msg)
			}
			mdwrk.liveness = heartbeat_liveness

			//  Don't try to handle errors, just assert noisily
			if len(msg) < 3 {
				panic("len(msg) < 3")
			}

			if msg[0] != "" {
				panic("msg[0] != \"\"")
			}

			if msg[1] != MDPW_WORKER {
				panic("msg[1] != MDPW_WORKER")
			}

			command := msg[2]
			msg = msg[3:]
			switch command {
			case MDPW_REQUEST:
				//  We should pop and save as many addresses as there are
				//  up to a null part, but for now, just save one...
				mdwrk.reply_to, msg = unwrap(msg)
				//  Here is where we actually have a message to process; we
				//  return it to the caller application:
				return //  We have a request to process
			case MDPW_HEARTBEAT:
				//  Do nothing for heartbeats
			case MDPW_DISCONNECT:
				mdwrk.ConnectToBroker()
			default:
				log.Printf("E: invalid input message %q\n", msg)
			}
		} else {
			mdwrk.liveness--
			if mdwrk.liveness == 0 {
				if mdwrk.verbose {
					log.Println("W: disconnected from broker - retrying...")
				}
				time.Sleep(mdwrk.reconnect)
				mdwrk.ConnectToBroker()
			}
		}
		//  Send HEARTBEAT if it's time
		if time.Now().After(mdwrk.heartbeat_at) {
			mdwrk.SendToBroker(MDPW_HEARTBEAT, "", []string{})
			mdwrk.heartbeat_at = time.Now().Add(mdwrk.heartbeat)
		}
	}
	return
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
