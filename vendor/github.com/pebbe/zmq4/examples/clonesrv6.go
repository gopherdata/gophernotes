//
//  Clone server Model Six
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"github.com/pebbe/zmq4/examples/bstar"
	"github.com/pebbe/zmq4/examples/kvmsg"

	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

//  Our server is defined by these properties
type clonesrv_t struct {
	kvmap      map[string]*kvmsg.Kvmsg //  Key-value store
	kvmap_init bool
	bstar      *bstar.Bstar   //  Bstar reactor core
	sequence   int64          //  How many updates we're at
	port       int            //  Main port we're working on
	peer       int            //  Main port of our peer
	publisher  *zmq.Socket    //  Publish updates and hugz
	collector  *zmq.Socket    //  Collect updates from clients
	subscriber *zmq.Socket    //  Get updates from peer
	pending    []*kvmsg.Kvmsg //  Pending updates from clients
	primary    bool           //  TRUE if we're primary
	active     bool           //  TRUE if we're active
	passive    bool           //  TRUE if we're passive
}

//  The main task parses the command line to decide whether to start
//  as primary or backup server. We're using the Binary Star pattern
//  for reliability. This interconnects the two servers so they can
//  agree on which is primary, and which is backup. To allow the two
//  servers to run on the same box, we use different ports for primary
//  and backup. Ports 5003/5004 are used to interconnect the servers.
//  Ports 5556/5566 are used to receive voting events (snapshot requests
//  in the clone pattern). Ports 5557/5567 are used by the publisher,
//  and ports 5558/5568 by the collector:

func main() {
	var err error

	srv := &clonesrv_t{}

	if len(os.Args) == 2 && os.Args[1] == "-p" {
		log.Println("I: primary active, waiting for backup (passive)")
		srv.bstar, err = bstar.New(bstar.PRIMARY, "tcp://*:5003", "tcp://localhost:5004")
		if err != nil {
			log.Println(err)
			return
		}
		srv.bstar.Voter("tcp://*:5556", zmq.ROUTER, func(soc *zmq.Socket) error { return snapshots(soc, srv) })
		srv.port = 5556
		srv.peer = 5566
		srv.primary = true
	} else if len(os.Args) == 2 && os.Args[1] == "-b" {
		log.Println("I: backup passive, waiting for primary (active)")
		srv.bstar, err = bstar.New(bstar.BACKUP, "tcp://*:5004", "tcp://localhost:5003")
		srv.bstar.Voter("tcp://*:5566", zmq.ROUTER, func(soc *zmq.Socket) error { return snapshots(soc, srv) })
		srv.port = 5566
		srv.peer = 5556
		srv.primary = false
	} else {
		fmt.Println("Usage: clonesrv4 { -p | -b }")
		return
	}
	//  Primary server will become first active
	if srv.primary {
		srv.kvmap = make(map[string]*kvmsg.Kvmsg, 0)
		srv.kvmap_init = true
	}

	srv.pending = make([]*kvmsg.Kvmsg, 0)
	srv.bstar.SetVerbose(true)

	//  Set up our clone server sockets
	srv.publisher, _ = zmq.NewSocket(zmq.PUB)
	srv.collector, _ = zmq.NewSocket(zmq.SUB)
	srv.collector.SetSubscribe("")
	srv.publisher.Bind(fmt.Sprint("tcp://*:", srv.port+1))
	srv.collector.Bind(fmt.Sprint("tcp://*:", srv.port+2))

	//  Set up our own clone client interface to peer
	srv.subscriber, _ = zmq.NewSocket(zmq.SUB)
	srv.subscriber.SetSubscribe("")
	srv.subscriber.Connect(fmt.Sprint("tcp://localhost:", srv.peer+1))

	//  After we've set-up our sockets we register our binary star
	//  event handlers, and then start the bstar reactor. This finishes
	//  when the user presses Ctrl-C, or the process receives a SIGINT
	//  interrupt:

	//  Register state change handlers
	srv.bstar.NewActive(func() error { return new_active(srv) })
	srv.bstar.NewPassive(func() error { return new_passive(srv) })

	//  Register our other handlers with the bstar reactor
	srv.bstar.Reactor.AddSocket(srv.collector, zmq.POLLIN,
		func(e zmq.State) error { return collector(srv) })
	srv.bstar.Reactor.AddChannelTime(time.Tick(1000*time.Millisecond), 1,
		func(i interface{}) error {
			if e := flush_ttl(srv); e != nil {
				return e
			}
			return send_hugz(srv)
		})

	err = srv.bstar.Start()
	log.Println(err)
}

func snapshots(socket *zmq.Socket, srv *clonesrv_t) (err error) {

	msg, err := socket.RecvMessage(0)
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
			socket.Send(identity, zmq.SNDMORE)
			kvmsg.Send(socket)
		}
	}

	//  Now send END message with sequence number
	log.Println("I: sending shapshot =", srv.sequence)
	socket.Send(identity, zmq.SNDMORE)
	kvmsg := kvmsg.NewKvmsg(srv.sequence)
	kvmsg.SetKey("KTHXBAI")
	kvmsg.SetBody(subtree)
	kvmsg.Send(socket)

	return
}

//  The collector is more complex than in the clonesrv5 example since how
//  process updates depends on whether we're active or passive. The active
//  applies them immediately to its kvmap, whereas the passive queues them
//  as pending:

//  If message was already on pending list, remove it and return TRUE,
//  else return FALSE.

func (srv *clonesrv_t) was_pending(kvmsg *kvmsg.Kvmsg) bool {
	uuid1, _ := kvmsg.GetUuid()
	for i, msg := range srv.pending {
		if uuid2, _ := msg.GetUuid(); uuid1 == uuid2 {
			srv.pending = append(srv.pending[:i], srv.pending[i+1:]...)
			return true
		}
	}
	return false
}

func collector(srv *clonesrv_t) (err error) {

	kvmsg, err := kvmsg.RecvKvmsg(srv.collector)
	if err != nil {
		return
	}

	if srv.active {
		srv.sequence++
		kvmsg.SetSequence(srv.sequence)
		kvmsg.Send(srv.publisher)
		if ttls, e := kvmsg.GetProp("ttl"); e == nil {
			ttl, e := strconv.ParseInt(ttls, 10, 64)
			if e != nil {
				err = e
				return
			}
			kvmsg.SetProp("ttl", fmt.Sprint(time.Now().Add(time.Duration(ttl)*time.Second).Unix()))
		}
		kvmsg.Store(srv.kvmap)
		log.Println("I: publishing update =", srv.sequence)
	} else {
		//  If we already got message from active, drop it, else
		//  hold on pending list
		if !srv.was_pending(kvmsg) {
			srv.pending = append(srv.pending, kvmsg)
		}
	}
	return
}

//  We purge ephemeral values using exactly the same code as in
//  the previous clonesrv5 example.
//  If key-value pair has expired, delete it and publish the
//  fact to listening clients.

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

// //  We send a HUGZ message once a second to all subscribers so that they
// //  can detect if our server dies. They'll then switch over to the backup
// //  server, which will become active:

func send_hugz(srv *clonesrv_t) (err error) {

	kvmsg := kvmsg.NewKvmsg(srv.sequence)
	kvmsg.SetKey("HUGZ")
	kvmsg.SetBody("")
	err = kvmsg.Send(srv.publisher)
	return
}

//  When we switch from passive to active, we apply our pending list so that
//  our kvmap is up-to-date. When we switch to passive, we wipe our kvmap
//  and grab a new snapshot from the active:

func new_active(srv *clonesrv_t) (err error) {

	srv.active = true
	srv.passive = false

	//  Stop subscribing to updates
	srv.bstar.Reactor.RemoveSocket(srv.subscriber)

	//  Apply pending list to own hash table
	for _, msg := range srv.pending {
		srv.sequence++
		msg.SetSequence(srv.sequence)
		msg.Send(srv.publisher)
		msg.Store(srv.kvmap)
		fmt.Println("I: publishing pending =", srv.sequence)
	}
	srv.pending = srv.pending[0:0]
	return
}

func new_passive(srv *clonesrv_t) (err error) {

	srv.kvmap = make(map[string]*kvmsg.Kvmsg)
	srv.kvmap_init = false
	srv.active = false
	srv.passive = true

	//  Start subscribing to updates
	srv.bstar.Reactor.AddSocket(srv.subscriber, zmq.POLLIN,
		func(e zmq.State) error { return subscriber(srv) })

	return
}

//  When we get an update, we create a new kvmap if necessary, and then
//  add our update to our kvmap. We're always passive in this case:

func subscriber(srv *clonesrv_t) (err error) {
	//  Get state snapshot if necessary
	if !srv.kvmap_init {
		srv.kvmap_init = true
		snapshot, _ := zmq.NewSocket(zmq.DEALER)
		snapshot.Connect(fmt.Sprint("tcp://localhost:", srv.peer))
		fmt.Printf("I: asking for snapshot from: tcp://localhost:%v\n", srv.peer)
		snapshot.SendMessage("ICANHAZ?", "") // blank subtree to get all
		for {
			kvmsg, e := kvmsg.RecvKvmsg(snapshot)
			if e != nil {
				err = e
				break
			}
			if key, _ := kvmsg.GetKey(); key == "KTHXBAI" {
				srv.sequence, _ = kvmsg.GetSequence()
				break //  Done
			}
			kvmsg.Store(srv.kvmap)
		}
		fmt.Println("I: received snapshot =", srv.sequence)
	}
	//  Find and remove update off pending list
	kvmsg, e := kvmsg.RecvKvmsg(srv.subscriber)
	if e != nil {
		err = e
		return
	}

	if key, _ := kvmsg.GetKey(); key != "HUGZ" {
		if !srv.was_pending(kvmsg) {
			//  If active update came before client update, flip it
			//  around, store active update (with sequence) on pending
			//  list and use to clear client update when it comes later
			srv.pending = append(srv.pending, kvmsg)
		}
		//  If update is more recent than our kvmap, apply it
		if seq, _ := kvmsg.GetSequence(); seq > srv.sequence {
			srv.sequence = seq
			kvmsg.Store(srv.kvmap)
			fmt.Println("I: received update =", srv.sequence)
		}
	}
	return
}
