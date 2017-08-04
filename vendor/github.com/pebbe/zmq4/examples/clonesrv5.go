//
//  Clone server Model Five
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"github.com/pebbe/zmq4/examples/kvmsg"

	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

//  Our server is defined by these properties
type clonesrv_t struct {
	kvmap     map[string]*kvmsg.Kvmsg //  Key-value store
	port      int                     //  Main port we're working on
	sequence  int64                   //  How many updates we're at
	snapshot  *zmq.Socket             //  Handle snapshot requests
	publisher *zmq.Socket             //  Publish updates to clients
	collector *zmq.Socket             //  Collect updates from clients
}

func main() {

	srv := &clonesrv_t{
		port:  5556,
		kvmap: make(map[string]*kvmsg.Kvmsg),
	}

	//  Set up our clone server sockets
	srv.snapshot, _ = zmq.NewSocket(zmq.ROUTER)
	srv.snapshot.Bind(fmt.Sprint("tcp://*:", srv.port))
	srv.publisher, _ = zmq.NewSocket(zmq.PUB)
	srv.publisher.Bind(fmt.Sprint("tcp://*:", srv.port+1))
	srv.collector, _ = zmq.NewSocket(zmq.PULL)
	srv.collector.Bind(fmt.Sprint("tcp://*:", srv.port+2))

	//  Register our handlers with reactor
	reactor := zmq.NewReactor()
	reactor.AddSocket(srv.snapshot, zmq.POLLIN,
		func(e zmq.State) error { return snapshots(srv) })
	reactor.AddSocket(srv.collector, zmq.POLLIN,
		func(e zmq.State) error { return collector(srv) })
	reactor.AddChannelTime(time.Tick(1000*time.Millisecond), 1,
		func(v interface{}) error { return flush_ttl(srv) })

	log.Println(reactor.Run(100 * time.Millisecond)) // precision: .1 seconds
}

//  This is the reactor handler for the snapshot socket; it accepts
//  just the ICANHAZ? request and replies with a state snapshot ending
//  with a KTHXBAI message:

func snapshots(srv *clonesrv_t) (err error) {

	msg, err := srv.snapshot.RecvMessage(0)
	if err != nil {
		return
	}
	identity := msg[0]

	//  Request is in second frame of message
	request := msg[1]
	if request != "ICANHAZ?" {
		err = errors.New("E: bad request, aborting")
		return
	}
	subtree := msg[2]

	//  Send state socket to client
	for _, kvmsg := range srv.kvmap {
		if key, _ := kvmsg.GetKey(); strings.HasPrefix(key, subtree) {
			srv.snapshot.Send(identity, zmq.SNDMORE)
			kvmsg.Send(srv.snapshot)
		}
	}

	//  Now send END message with sequence number
	log.Println("I: sending shapshot =", srv.sequence)
	srv.snapshot.Send(identity, zmq.SNDMORE)
	kvmsg := kvmsg.NewKvmsg(srv.sequence)
	kvmsg.SetKey("KTHXBAI")
	kvmsg.SetBody(subtree)
	kvmsg.Send(srv.snapshot)

	return
}

//  We store each update with a new sequence number, and if necessary, a
//  time-to-live. We publish updates immediately on our publisher socket:

func collector(srv *clonesrv_t) (err error) {
	kvmsg, err := kvmsg.RecvKvmsg(srv.collector)
	if err != nil {
		return
	}

	srv.sequence++
	kvmsg.SetSequence(srv.sequence)
	kvmsg.Send(srv.publisher)
	if ttls, e := kvmsg.GetProp("ttl"); e == nil {
		// change duration into specific time, using the same property: ugly!
		ttl, e := strconv.ParseInt(ttls, 10, 64)
		if e != nil {
			err = e
			return
		}
		kvmsg.SetProp("ttl", fmt.Sprint(time.Now().Add(time.Duration(ttl)*time.Second).Unix()))
	}
	kvmsg.Store(srv.kvmap)
	log.Println("I: publishing update =", srv.sequence)

	return
}

//  At regular intervals we flush ephemeral values that have expired. This
//  could be slow on very large data sets:

func flush_ttl(srv *clonesrv_t) (err error) {

	for _, kvmsg := range srv.kvmap {

		//  If key-value pair has expired, delete it and publish the
		//  fact to listening clients.

		if ttls, e := kvmsg.GetProp("ttl"); e == nil {
			ttl, e := strconv.ParseInt(ttls, 10, 64)
			if e != nil {
				err = e
				continue
			}
			if time.Now().After(time.Unix(ttl, 0)) {
				srv.sequence++
				kvmsg.SetSequence(srv.sequence)
				kvmsg.SetBody("")
				e = kvmsg.Send(srv.publisher)
				if e != nil {
					err = e
				}
				kvmsg.Store(srv.kvmap)
				log.Println("I: publishing delete =", srv.sequence)
			}
		}
	}
	return
}
