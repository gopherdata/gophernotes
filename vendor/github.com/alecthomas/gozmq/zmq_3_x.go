// +build zmq_3_x zmq_4_x

/*
  Copyright 2010-2012 Alec Thomas

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

package gozmq

/*
#cgo pkg-config: libzmq
#include <zmq.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

const (
	RCVMORE = IntSocketOption(C.ZMQ_RCVMORE)
	SNDHWM  = IntSocketOption(C.ZMQ_SNDHWM)
	RCVHWM  = IntSocketOption(C.ZMQ_RCVHWM)

	// TODO Not documented in the man page...
	//LAST_ENDPOINT       = UInt64SocketOption(C.ZMQ_LAST_ENDPOINT)
	DELAY_ATTACH_ON_CONNECT = IntSocketOption(C.ZMQ_DELAY_ATTACH_ON_CONNECT)
	FAIL_UNROUTABLE         = BoolSocketOption(C.ZMQ_FAIL_UNROUTABLE)
	IPV4ONLY                = IntSocketOption(C.ZMQ_IPV4ONLY)
	MAXMSGSIZE              = Int64SocketOption(C.ZMQ_MAXMSGSIZE)
	ROUTER_MANDATORY        = IntSocketOption(C.ZMQ_ROUTER_MANDATORY)
	TCP_KEEPALIVE           = IntSocketOption(C.ZMQ_TCP_KEEPALIVE)
	TCP_KEEPALIVE_CNT       = IntSocketOption(C.ZMQ_TCP_KEEPALIVE_CNT)
	TCP_KEEPALIVE_IDLE      = IntSocketOption(C.ZMQ_TCP_KEEPALIVE_IDLE)
	TCP_KEEPALIVE_INTVL     = IntSocketOption(C.ZMQ_TCP_KEEPALIVE_INTVL)
	TCP_ACCEPT_FILTER       = StringSocketOption(C.ZMQ_TCP_ACCEPT_FILTER)
	XPUB_VERBOSE            = IntSocketOption(C.ZMQ_XPUB_VERBOSE)

	// Message options
	MORE = MessageOption(C.ZMQ_MORE)

	// Send/recv options
	DONTWAIT = SendRecvOption(C.ZMQ_DONTWAIT)

	// Deprecated aliases
	NOBLOCK = DONTWAIT
)

// Socket transport events
type Event int

const (
	EVENT_CONNECTED       = Event(C.ZMQ_EVENT_CONNECTED)
	EVENT_CONNECT_DELAYED = Event(C.ZMQ_EVENT_CONNECT_DELAYED)
	EVENT_CONNECT_RETRIED = Event(C.ZMQ_EVENT_CONNECT_RETRIED)

	EVENT_LISTENING   = Event(C.ZMQ_EVENT_LISTENING)
	EVENT_BIND_FAILED = Event(C.ZMQ_EVENT_BIND_FAILED)

	EVENT_ACCEPTED      = Event(C.ZMQ_EVENT_ACCEPTED)
	EVENT_ACCEPT_FAILED = Event(C.ZMQ_EVENT_ACCEPT_FAILED)

	EVENT_CLOSED       = Event(C.ZMQ_EVENT_CLOSED)
	EVENT_CLOSE_FAILED = Event(C.ZMQ_EVENT_CLOSE_FAILED)
	EVENT_DISCONNECTED = Event(C.ZMQ_EVENT_DISCONNECTED)

	EVENT_ALL = EVENT_CONNECTED | EVENT_CONNECT_DELAYED |
		EVENT_CONNECT_RETRIED | EVENT_LISTENING | EVENT_BIND_FAILED |
		EVENT_ACCEPTED | EVENT_ACCEPT_FAILED | EVENT_CLOSED |
		EVENT_CLOSE_FAILED | EVENT_DISCONNECTED
)

// Get a context option.
// int zmq_ctx_get (void *c, int);
func (c *Context) get(option C.int) (int, error) {
	if c.init(); c.err != nil {
		return -1, c.err
	}
	var value C.int
	var err error
	if value, err = C.zmq_ctx_get(c.c, option); err != nil {
		return -1, casterr(err)
	}
	return int(value), nil
}

// Set a context option.
// int zmq_ctx_set (void *c, int, int);
func (c *Context) set(option C.int, value int) error {
	if c.init(); c.err != nil {
		return c.err
	}
	if rc, err := C.zmq_ctx_set(c.c, option, C.int(value)); rc == -1 {
		return casterr(err)
	}
	return nil
}

func (c *Context) IOThreads() (int, error) {
	return c.get(C.ZMQ_IO_THREADS)
}

func (c *Context) MaxSockets() (int, error) {
	return c.get(C.ZMQ_MAX_SOCKETS)
}

func (c *Context) SetIOThreads(value int) error {
	return c.set(C.ZMQ_IO_THREADS, value)
}

func (c *Context) SetMaxSockets(value int) error {
	return c.set(C.ZMQ_MAX_SOCKETS, value)
}

func (s *Socket) SetHWM(value int) error {
	snd := s.SetSndHWM(value)
	rcv := s.SetRcvHWM(value)
	if snd != nil {
		return snd
	}
	return rcv
}

func (s *Socket) SetTCPAcceptFilterNil() error {
	return s.SetSockOptStringNil(TCP_ACCEPT_FILTER)
}

// Disconnect the socket from the address.
// int zmq_disconnect (void *s, const char *addr);
func (s *Socket) Disconnect(address string) error {
	if s.c == nil {
		return ENOTSOCK
	}
	a := C.CString(address)
	defer C.free(unsafe.Pointer(a))
	if rc, err := C.zmq_disconnect(s.s, a); rc != 0 {
		return casterr(err)
	}
	return nil
}

// Send a message to the socket.
// int zmq_send (void *s, zmq_msg_t *msg, int flags);
func (s *Socket) Send(data []byte, flags SendRecvOption) error {
	var m C.zmq_msg_t
	// Copy data array into C-allocated buffer.
	size := C.size_t(len(data))

	if rc, err := C.zmq_msg_init_size(&m, size); rc != 0 {
		return casterr(err)
	}

	if size > 0 {
		// FIXME Ideally this wouldn't require a copy.
		C.memcpy(C.zmq_msg_data(&m), unsafe.Pointer(&data[0]), size) // XXX I hope this works...(seems to)
	}

	if rc, err := C.zmq_sendmsg(s.s, &m, C.int(flags)); rc == -1 {
		// zmq_send did not take ownership, free message
		C.zmq_msg_close(&m)
		return casterr(err)
	}
	return nil
}

// Receive a message from the socket.
// int zmq_recv (void *s, zmq_msg_t *msg, int flags);
func (s *Socket) Recv(flags SendRecvOption) (data []byte, err error) {
	// Allocate and initialise a new zmq_msg_t
	var m C.zmq_msg_t
	var rc C.int
	if rc, err = C.zmq_msg_init(&m); rc != 0 {
		err = casterr(err)
		return
	}
	defer C.zmq_msg_close(&m)
	// Receive into message
	if rc, err = C.zmq_recvmsg(s.s, &m, C.int(flags)); rc == -1 {
		err = casterr(err)
		return
	}
	err = nil
	// Copy message data into a byte array
	// FIXME Ideally this wouldn't require a copy.
	size := C.zmq_msg_size(&m)
	if size > 0 {
		data = C.GoBytes(C.zmq_msg_data(&m), C.int(size))
	} else {
		data = nil
	}
	return
}

// Register a monitoring callback endpoint.
// int zmq_socket_monitor (void *s, const char *addr, int events);
func (s *Socket) Monitor(address string, events Event) error {
	a := C.CString(address)
	defer C.free(unsafe.Pointer(a))

	rc, err := C.zmq_socket_monitor(s.apiSocket(), a, C.int(events))
	if rc == -1 {
		return casterr(err)
	}
	return nil
}

// Portability helper
func (s *Socket) getRcvmore() (more bool, err error) {
	value, err := s.GetSockOptInt(RCVMORE)
	more = value != 0
	return
}

// run a zmq_proxy with in, out and capture sockets
func Proxy(in, out, capture *Socket) error {
	var c unsafe.Pointer
	if capture != nil {
		c = capture.apiSocket()
	}
	if rc, err := C.zmq_proxy(in.apiSocket(), out.apiSocket(), c); rc != 0 {
		return casterr(err)
	}
	return errors.New("zmq_proxy() returned unexpectedly.")
}
