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

//  Structure of our class
//  We access these properties only via class methods

// Majordomo Protocol Client API.
type Mdcli struct {
	broker  string
	client  *zmq.Socket   //  Socket to broker
	verbose bool          //  Print activity to stdout
	timeout time.Duration //  Request timeout
	retries int           //  Request retries
	poller  *zmq.Poller
}

//  ---------------------------------------------------------------------

//  Connect or reconnect to broker.
func (mdcli *Mdcli) ConnectToBroker() (err error) {
	if mdcli.client != nil {
		mdcli.client.Close()
		mdcli.client = nil
	}
	mdcli.client, err = zmq.NewSocket(zmq.REQ)
	if err != nil {
		if mdcli.verbose {
			log.Println("E: ConnectToBroker() creating socket failed")
		}
		return
	}
	mdcli.poller = zmq.NewPoller()
	mdcli.poller.Add(mdcli.client, zmq.POLLIN)

	if mdcli.verbose {
		log.Printf("I: connecting to broker at %s...", mdcli.broker)
	}
	err = mdcli.client.Connect(mdcli.broker)
	if err != nil && mdcli.verbose {
		log.Println("E: ConnectToBroker() failed to connect to broker", mdcli.broker)
	}

	return
}

//  Here we have the constructor and destructor for our mdcli class:

//  ---------------------------------------------------------------------
//  Constructor

func NewMdcli(broker string, verbose bool) (mdcli *Mdcli, err error) {

	mdcli = &Mdcli{
		broker:  broker,
		verbose: verbose,
		timeout: time.Duration(2500 * time.Millisecond),
		retries: 3, //  Before we abandon
	}
	err = mdcli.ConnectToBroker()
	runtime.SetFinalizer(mdcli, (*Mdcli).Close)
	return
}

//  ---------------------------------------------------------------------
//  Destructor

func (mdcli *Mdcli) Close() (err error) {
	if mdcli.client != nil {
		err = mdcli.client.Close()
		mdcli.client = nil
	}
	return
}

//  These are the class methods. We can set the request timeout and number
//  of retry attempts, before sending requests:

//  ---------------------------------------------------------------------

//  Set request timeout.
func (mdcli *Mdcli) SetTimeout(timeout time.Duration) {
	mdcli.timeout = timeout
}

//  ---------------------------------------------------------------------

//  Set request retries.
func (mdcli *Mdcli) SetRetries(retries int) {
	mdcli.retries = retries
}

//  Here is the send method. It sends a request to the broker and gets a
//  reply even if it has to retry several times. It returns the reply
//  message, or error if there was no reply after multiple attempts:
func (mdcli *Mdcli) Send(service string, request ...string) (reply []string, err error) {
	//  Prefix request with protocol frames
	//  Frame 1: "MDPCxy" (six bytes, MDP/Client x.y)
	//  Frame 2: Service name (printable string)

	req := make([]string, 2, len(request)+2)
	req = append(req, request...)
	req[1] = service
	req[0] = MDPC_CLIENT
	if mdcli.verbose {
		log.Printf("I: send request to '%s' service: %q\n", service, req)
	}
	for retries_left := mdcli.retries; retries_left > 0; retries_left-- {
		_, err = mdcli.client.SendMessage(req)
		if err != nil {
			break
		}

		//  On any blocking call, libzmq will return -1 if there was
		//  an error; we could in theory check for different error codes
		//  but in practice it's OK to assume it was EINTR (Ctrl-C):

		var polled []zmq.Polled
		polled, err = mdcli.poller.Poll(mdcli.timeout)
		if err != nil {
			break //  Interrupted
		}

		//  If we got a reply, process it
		if len(polled) > 0 {
			var msg []string
			msg, err = mdcli.client.RecvMessage(0)
			if err != nil {
				break
			}
			if mdcli.verbose {
				log.Printf("I: received reply: %q\n", msg)
			}
			//  We would handle malformed replies better in real code
			if len(msg) < 3 {
				panic("len(msg) < 3")
			}

			if msg[0] != MDPC_CLIENT {
				panic("msg[0] != MDPC_CLIENT")
			}

			if msg[1] != service {
				panic("msg[1] != service")
			}

			reply = msg[2:]
			return //  Success
		} else {
			if mdcli.verbose {
				log.Println("W: no reply, reconnecting...")
			}
			mdcli.ConnectToBroker()
		}
	}
	if err == nil {
		err = errors.New("permanent error")
	}
	if mdcli.verbose {
		log.Println("W: permanent error, abandoning")
	}
	return
}
