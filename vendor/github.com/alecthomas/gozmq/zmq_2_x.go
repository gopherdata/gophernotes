// +build !zmq_3_x,!zmq_4_x

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
import "unsafe"

const (
	RCVMORE           = UInt64SocketOption(C.ZMQ_RCVMORE)
	RECOVERY_IVL_MSEC = Int64SocketOption(C.ZMQ_RECOVERY_IVL_MSEC)
	SWAP              = Int64SocketOption(C.ZMQ_SWAP)
	MCAST_LOOP        = Int64SocketOption(C.ZMQ_MCAST_LOOP)
	HWM               = UInt64SocketOption(C.ZMQ_HWM)
	NOBLOCK           = SendRecvOption(C.ZMQ_NOBLOCK)

	// Forwards-compatible aliases:
	DONTWAIT = NOBLOCK
)

// Get a context option.
func (c *Context) IOThreads() (int, error) {
	return c.iothreads, nil
}

// Set a context option.
func (c *Context) SetIOThreads(value int) error {
	c.iothreads = value
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

	if rc, err := C.zmq_send(s.s, &m, C.int(flags)); rc != 0 {
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
	if rc, err = C.zmq_recv(s.s, &m, C.int(flags)); rc != 0 {
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

// Portability helper
func (s *Socket) getRcvmore() (more bool, err error) {
	value, err := s.GetSockOptUInt64(RCVMORE)
	more = value != 0
	return
}
