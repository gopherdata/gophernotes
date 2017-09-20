// Majordomo Protocol Client API.
// Implements the MDP/Worker spec at http://rfc.zeromq.org/spec:7.

package mdapi

import (
	zmq "github.com/pebbe/zmq4"

	"errors"
	"log"
	"runtime"
	"time"
)

var (
	errPermanent = errors.New("permanent error, abandoning request")
)

//  Structure of our class
//  We access these properties only via class methods

// Majordomo Protocol Client API.
type Mdcli2 struct {
	broker  string
	client  *zmq.Socket   //  Socket to broker
	verbose bool          //  Print activity to stdout
	timeout time.Duration //  Request timeout
	poller  *zmq.Poller
}

//  ---------------------------------------------------------------------

//  Connect or reconnect to broker. In this asynchronous class we use a
//  DEALER socket instead of a REQ socket; this lets us send any number
//  of requests without waiting for a reply.
func (mdcli2 *Mdcli2) ConnectToBroker() (err error) {
	if mdcli2.client != nil {
		mdcli2.client.Close()
		mdcli2.client = nil
	}
	mdcli2.client, err = zmq.NewSocket(zmq.DEALER)
	if err != nil {
		if mdcli2.verbose {
			log.Println("E: ConnectToBroker() creating socket failed")
		}
		return
	}
	mdcli2.poller = zmq.NewPoller()
	mdcli2.poller.Add(mdcli2.client, zmq.POLLIN)

	if mdcli2.verbose {
		log.Printf("I: connecting to broker at %s...", mdcli2.broker)
	}
	err = mdcli2.client.Connect(mdcli2.broker)
	if err != nil && mdcli2.verbose {
		log.Println("E: ConnectToBroker() failed to connect to broker", mdcli2.broker)
	}

	return
}

//  Here we have the constructor and destructor for our mdcli2 class:

//  The constructor and destructor are the same as in mdcliapi, except
//  we don't do retries, so there's no retries property.
//  ---------------------------------------------------------------------
//  Constructor

func NewMdcli2(broker string, verbose bool) (mdcli2 *Mdcli2, err error) {

	mdcli2 = &Mdcli2{
		broker:  broker,
		verbose: verbose,
		timeout: time.Duration(2500 * time.Millisecond),
	}
	err = mdcli2.ConnectToBroker()
	runtime.SetFinalizer(mdcli2, (*Mdcli2).Close)
	return
}

//  ---------------------------------------------------------------------
//  Destructor

func (mdcli2 *Mdcli2) Close() (err error) {
	if mdcli2.client != nil {
		err = mdcli2.client.Close()
		mdcli2.client = nil
	}
	return
}

//  ---------------------------------------------------------------------

//  Set request timeout.
func (mdcli2 *Mdcli2) SetTimeout(timeout time.Duration) {
	mdcli2.timeout = timeout
}

//  The send method now just sends one message, without waiting for a
//  reply. Since we're using a DEALER socket we have to send an empty
//  frame at the start, to create the same envelope that the REQ socket
//  would normally make for us:
func (mdcli2 *Mdcli2) Send(service string, request ...string) (err error) {
	//  Prefix request with protocol frames
	//  Frame 0: empty (REQ emulation)
	//  Frame 1: "MDPCxy" (six bytes, MDP/Client x.y)
	//  Frame 2: Service name (printable string)

	req := make([]string, 3, len(request)+3)
	req = append(req, request...)
	req[2] = service
	req[1] = MDPC_CLIENT
	req[0] = ""
	if mdcli2.verbose {
		log.Printf("I: send request to '%s' service: %q\n", service, req)
	}
	_, err = mdcli2.client.SendMessage(req)
	return
}

//  The recv method waits for a reply message and returns that to the
//  caller.
//  ---------------------------------------------------------------------
//  Returns the reply message or NULL if there was no reply. Does not
//  attempt to recover from a broker failure, this is not possible
//  without storing all unanswered requests and resending them all...

func (mdcli2 *Mdcli2) Recv() (msg []string, err error) {

	msg = []string{}

	//  Poll socket for a reply, with timeout
	polled, err := mdcli2.poller.Poll(mdcli2.timeout)
	if err != nil {
		return //  Interrupted
	}

	//  If we got a reply, process it
	if len(polled) > 0 {
		msg, err = mdcli2.client.RecvMessage(0)
		if err != nil {
			log.Println("W: interrupt received, killing client...")
			return
		}

		if mdcli2.verbose {
			log.Printf("I: received reply: %q\n", msg)
		}
		//  Don't try to handle errors, just assert noisily
		if len(msg) < 4 {
			panic("len(msg) < 4")
		}

		if msg[0] != "" {
			panic("msg[0] != \"\"")
		}

		if msg[1] != MDPC_CLIENT {
			panic("msg[1] != MDPC_CLIENT")
		}

		msg = msg[3:]
		return //  Success
	}

	err = errPermanent
	if mdcli2.verbose {
		log.Println(err)
	}
	return
}
