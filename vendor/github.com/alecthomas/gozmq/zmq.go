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

// Go (golang) Bindings for 0mq (zmq, zeromq)
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
	"sync"
	"syscall"
	"time"
	"unsafe"
)

type SocketType int

type IntSocketOption int
type Int64SocketOption int
type UInt64SocketOption int
type StringSocketOption int
type BoolSocketOption int

type MessageOption int
type SendRecvOption int

const (
	// NewSocket types
	PAIR   = SocketType(C.ZMQ_PAIR)
	PUB    = SocketType(C.ZMQ_PUB)
	SUB    = SocketType(C.ZMQ_SUB)
	REQ    = SocketType(C.ZMQ_REQ)
	REP    = SocketType(C.ZMQ_REP)
	DEALER = SocketType(C.ZMQ_DEALER)
	ROUTER = SocketType(C.ZMQ_ROUTER)
	PULL   = SocketType(C.ZMQ_PULL)
	PUSH   = SocketType(C.ZMQ_PUSH)
	XPUB   = SocketType(C.ZMQ_XPUB)
	XSUB   = SocketType(C.ZMQ_XSUB)

	// Deprecated aliases
	XREQ       = DEALER
	XREP       = ROUTER
	UPSTREAM   = PULL
	DOWNSTREAM = PUSH

	// NewSocket options
	AFFINITY          = UInt64SocketOption(C.ZMQ_AFFINITY)
	IDENTITY          = StringSocketOption(C.ZMQ_IDENTITY)
	SUBSCRIBE         = StringSocketOption(C.ZMQ_SUBSCRIBE)
	UNSUBSCRIBE       = StringSocketOption(C.ZMQ_UNSUBSCRIBE)
	RATE              = Int64SocketOption(C.ZMQ_RATE)
	RECOVERY_IVL      = Int64SocketOption(C.ZMQ_RECOVERY_IVL)
	SNDBUF            = UInt64SocketOption(C.ZMQ_SNDBUF)
	RCVBUF            = UInt64SocketOption(C.ZMQ_RCVBUF)
	FD                = Int64SocketOption(C.ZMQ_FD)
	EVENTS            = UInt64SocketOption(C.ZMQ_EVENTS)
	TYPE              = UInt64SocketOption(C.ZMQ_TYPE)
	LINGER            = IntSocketOption(C.ZMQ_LINGER)
	RECONNECT_IVL     = IntSocketOption(C.ZMQ_RECONNECT_IVL)
	RECONNECT_IVL_MAX = IntSocketOption(C.ZMQ_RECONNECT_IVL_MAX)
	BACKLOG           = IntSocketOption(C.ZMQ_BACKLOG)

	// Send/recv options
	SNDMORE = SendRecvOption(C.ZMQ_SNDMORE)
)

type zmqErrno syscall.Errno

var (
	// Additional ZMQ errors
	ENOTSOCK       error = zmqErrno(C.ENOTSOCK)
	EFSM           error = zmqErrno(C.EFSM)
	EINVAL         error = syscall.EINVAL
	ENOCOMPATPROTO error = zmqErrno(C.ENOCOMPATPROTO)
	ETERM          error = zmqErrno(C.ETERM)
	EMTHREAD       error = zmqErrno(C.EMTHREAD)
)

type PollEvents C.short

const (
	POLLIN  = PollEvents(C.ZMQ_POLLIN)
	POLLOUT = PollEvents(C.ZMQ_POLLOUT)
	POLLERR = PollEvents(C.ZMQ_POLLERR)
)

type DeviceType int

const (
	STREAMER  = DeviceType(C.ZMQ_STREAMER)
	FORWARDER = DeviceType(C.ZMQ_FORWARDER)
	QUEUE     = DeviceType(C.ZMQ_QUEUE)
)

var (
	pollunit time.Duration
)

func init() {
	if v, _, _ := Version(); v < 3 {
		pollunit = time.Microsecond
	} else {
		pollunit = time.Millisecond
	}
}

// void zmq_version (int *major, int *minor, int *patch);
func Version() (int, int, int) {
	var major, minor, patch C.int
	C.zmq_version(&major, &minor, &patch)
	return int(major), int(minor), int(patch)
}

func (e zmqErrno) Error() string {
	return C.GoString(C.zmq_strerror(C.int(e)))
}

// If possible, convert a syscall.Errno to a zmqErrno.
func casterr(fromcgo error) error {
	errno, ok := fromcgo.(syscall.Errno)
	if !ok {
		return fromcgo
	}
	zmqerrno := zmqErrno(errno)
	switch zmqerrno {
	case ENOTSOCK:
		return zmqerrno
	}
	if zmqerrno >= C.ZMQ_HAUSNUMERO {
		return zmqerrno
	}
	return errno
}

func getErrorForTesting() error {
	return zmqErrno(C.EFSM)
}

/*
 * A context handles socket creation and asynchronous message delivery.
 * There should generally be one context per application.
 */
type Context struct {
	c         unsafe.Pointer
	mutex     sync.Mutex // ensure init is only called once
	init      func()     // func that calls zmq_init
	err       error      // error returned from zmq_init
	iothreads int        // hold the iothreads option until zmq_init time
}

// Create a new context.
func NewContext() (*Context, error) {
	c := &Context{iothreads: 1}
	c.init = func() {
		c.mutex.Lock()
		defer c.mutex.Unlock()
		if c.c == nil && c.err == nil {
			// C.NULL is correct but causes a runtime failure on darwin at present
			if ptr, err := C.zmq_init(C.int(c.iothreads)); ptr != nil /*C.NULL*/ {
				c.c = ptr
			} else {
				c.err = casterr(err)
			}
		}
	}
	return c, nil
}

func (c *Context) Close() {
	// C.NULL is correct but causes a runtime failure on darwin at present
	if c.c != nil /*C.NULL*/ {
		C.zmq_term(c.c)
	}
}

// Create a new socket.
// void *zmq_socket (void *context, int type);
func (c *Context) NewSocket(t SocketType) (*Socket, error) {
	if c.init(); c.err != nil {
		return nil, c.err
	}
	s, err := C.zmq_socket(c.c, C.int(t))
	// C.NULL is correct but causes a runtime failure on darwin at present
	if s != nil /*C.NULL*/ {
		return &Socket{c: c, s: s}, nil
	}
	return nil, casterr(err)
}

type Socket struct {
	// XXX Ensure the zmq context doesn't get destroyed underneath us.
	c *Context
	s unsafe.Pointer
}

// Shutdown the socket.
// int zmq_close (void *s);
func (s *Socket) Close() error {
	if s.c == nil {
		return ENOTSOCK
	}
	if rc, err := C.zmq_close(s.s); rc != 0 {
		return casterr(err)
	}
	s.c = nil
	return nil
}

// Set an int option on the socket.
// int zmq_setsockopt (void *s, int option, const void *optval, size_t optvallen);
func (s *Socket) SetSockOptInt(option IntSocketOption, value int) error {
	val := C.int(value)
	if rc, err := C.zmq_setsockopt(s.s, C.int(option), unsafe.Pointer(&val), C.size_t(unsafe.Sizeof(val))); rc != 0 {
		return casterr(err)
	}
	return nil
}

// Set an int64 option on the socket.
// int zmq_setsockopt (void *s, int option, const void *optval, size_t optvallen);
func (s *Socket) SetSockOptInt64(option Int64SocketOption, value int64) error {
	if rc, err := C.zmq_setsockopt(s.s, C.int(option), unsafe.Pointer(&value), C.size_t(unsafe.Sizeof(value))); rc != 0 {
		return casterr(err)
	}
	return nil
}

// Set a uint64 option on the socket.
// int zmq_setsockopt (void *s, int option, const void *optval, size_t optvallen);
func (s *Socket) SetSockOptUInt64(option UInt64SocketOption, value uint64) error {
	if rc, err := C.zmq_setsockopt(s.s, C.int(option), unsafe.Pointer(&value), C.size_t(unsafe.Sizeof(value))); rc != 0 {
		return casterr(err)
	}
	return nil
}

// Set a string option on the socket.
// int zmq_setsockopt (void *s, int option, const void *optval, size_t optvallen);
func (s *Socket) SetSockOptString(option StringSocketOption, value string) error {
	v := C.CString(value)
	defer C.free(unsafe.Pointer(v))
	if rc, err := C.zmq_setsockopt(s.s, C.int(option), unsafe.Pointer(v), C.size_t(len(value))); rc != 0 {
		return casterr(err)
	}
	return nil
}

// Set a string option on the socket to nil.
// int zmq_setsockopt (void *s, int option, const void *optval, size_t optvallen);
func (s *Socket) SetSockOptStringNil(option StringSocketOption) error {
	if rc, err := C.zmq_setsockopt(s.s, C.int(option), nil, 0); rc != 0 {
		return casterr(err)
	}
	return nil
}

// Get an int option from the socket.
// int zmq_getsockopt (void *s, int option, void *optval, size_t *optvallen);
func (s *Socket) GetSockOptInt(option IntSocketOption) (value int, err error) {
	size := C.size_t(unsafe.Sizeof(value))
	var rc C.int
	if rc, err = C.zmq_getsockopt(s.s, C.int(option), unsafe.Pointer(&value), &size); rc != 0 {
		err = casterr(err)
		return
	}
	return
}

// Get an int64 option from the socket.
// int zmq_getsockopt (void *s, int option, void *optval, size_t *optvallen);
func (s *Socket) GetSockOptInt64(option Int64SocketOption) (value int64, err error) {
	size := C.size_t(unsafe.Sizeof(value))
	var rc C.int
	if rc, err = C.zmq_getsockopt(s.s, C.int(option), unsafe.Pointer(&value), &size); rc != 0 {
		err = casterr(err)
		return
	}
	return
}

// Get a uint64 option from the socket.
// int zmq_getsockopt (void *s, int option, void *optval, size_t *optvallen);
func (s *Socket) GetSockOptUInt64(option UInt64SocketOption) (value uint64, err error) {
	size := C.size_t(unsafe.Sizeof(value))
	var rc C.int
	if rc, err = C.zmq_getsockopt(s.s, C.int(option), unsafe.Pointer(&value), &size); rc != 0 {
		println("GetSockOptUInt64:", err.Error())
		err = casterr(err)
		return
	}
	return
}

// Get a string option from the socket.
// int zmq_getsockopt (void *s, int option, void *optval, size_t *optvallen);
func (s *Socket) GetSockOptString(option StringSocketOption) (value string, err error) {
	var buffer [1024]byte
	var size C.size_t = 1024
	var rc C.int
	if rc, err = C.zmq_getsockopt(s.s, C.int(option), unsafe.Pointer(&buffer), &size); rc != 0 {
		err = casterr(err)
		return
	}
	value = string(buffer[:size])
	return
}

func (s *Socket) GetSockOptBool(option BoolSocketOption) (value bool, err error) {
	size := C.size_t(unsafe.Sizeof(value))
	var rc C.int
	if rc, err = C.zmq_getsockopt(s.s, C.int(option), unsafe.Pointer(&value), &size); rc != 0 {
		err = casterr(err)
		return
	}
	return
}

// Bind the socket to a listening address.
// int zmq_bind (void *s, const char *addr);
func (s *Socket) Bind(address string) error {
	a := C.CString(address)
	defer C.free(unsafe.Pointer(a))
	if rc, err := C.zmq_bind(s.s, a); rc != 0 {
		return casterr(err)
	}
	return nil
}

// Connect the socket to an address.
// int zmq_connect (void *s, const char *addr);
func (s *Socket) Connect(address string) error {
	if s.c == nil {
		return ENOTSOCK
	}
	a := C.CString(address)
	defer C.free(unsafe.Pointer(a))
	if rc, err := C.zmq_connect(s.s, a); rc != 0 {
		return casterr(err)
	}
	return nil
}

// Send a multipart message.
func (s *Socket) SendMultipart(parts [][]byte, flags SendRecvOption) (err error) {
	for i := 0; i < len(parts)-1; i++ {
		if err = s.Send(parts[i], SNDMORE|flags); err != nil {
			return
		}
	}
	err = s.Send(parts[(len(parts)-1)], flags)
	return
}

// Receive a multipart message.
func (s *Socket) RecvMultipart(flags SendRecvOption) (parts [][]byte, err error) {
	parts = make([][]byte, 0)
	for {
		var data []byte
		var more bool

		data, err = s.Recv(flags)
		if err != nil {
			return
		}
		parts = append(parts, data)
		more, err = s.getRcvmore()
		if err != nil {
			return
		}
		if !more {
			break
		}
	}
	return
}

// return the
func (s *Socket) apiSocket() unsafe.Pointer {
	return s.s
}

// Item to poll for read/write events on, either a *Socket or a file descriptor
type PollItem struct {
	Socket  *Socket         // socket to poll for events on
	Fd      ZmqOsSocketType // fd to poll for events on as returned from os.File.Fd()
	Events  PollEvents      // event set to poll for
	REvents PollEvents      // events that were present
}

// a set of items to poll for events on
type PollItems []PollItem

// Poll ZmqSockets and file descriptors for I/O readiness. Timeout is in
// time.Duration. The smallest possible timeout is time.Millisecond for
// ZeroMQ version 3 and above, and time.Microsecond for earlier versions.
func Poll(items []PollItem, timeout time.Duration) (count int, err error) {
	zitems := make([]C.zmq_pollitem_t, len(items))
	for i, pi := range items {
		zitems[i].socket = pi.Socket.apiSocket()
		zitems[i].fd = pi.Fd.ToRaw()
		zitems[i].events = C.short(pi.Events)
	}
	ztimeout := C.long(-1)
	if timeout >= 0 {
		ztimeout = C.long(uint64(timeout / pollunit))
	}
	rc, err := C.zmq_poll(&zitems[0], C.int(len(zitems)), ztimeout)
	if rc == -1 {
		return 0, casterr(err)
	}

	for i, zi := range zitems {
		items[i].REvents = PollEvents(zi.revents)
	}

	return int(rc), nil
}

// run a zmq_device passing messages between in and out
func Device(t DeviceType, in, out *Socket) error {
	if rc, err := C.zmq_device(C.int(t), in.apiSocket(), out.apiSocket()); rc != 0 {
		return casterr(err)
	}
	return errors.New("zmq_device() returned unexpectedly.")
}

// XXX For now, this library abstracts zmq_msg_t out of the API.
// int zmq_msg_init (zmq_msg_t *msg);
// int zmq_msg_init_size (zmq_msg_t *msg, size_t size);
// int zmq_msg_close (zmq_msg_t *msg);
// size_t zmq_msg_size (zmq_msg_t *msg);
// void *zmq_msg_data (zmq_msg_t *msg);
// int zmq_msg_copy (zmq_msg_t *dest, zmq_msg_t *src);
// int zmq_msg_move (zmq_msg_t *dest, zmq_msg_t *src);
