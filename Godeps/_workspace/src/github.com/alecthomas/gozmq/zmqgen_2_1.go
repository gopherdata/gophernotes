// +build zmq_2_1
//

package gozmq

import (
	"time"
)

// This file was generated automatically.  Changes made here will be lost.

// Socket Option Getters

// ZMQ_TYPE: Retrieve socket type.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc3
//
func (s *Socket) Type() (SocketType, error) {
	value, err := s.GetSockOptUInt64(TYPE)
	return SocketType(value), err
}

// ZMQ_RCVMORE: More message parts to follow.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc4
//
func (s *Socket) RcvMore() (bool, error) {
	value, err := s.GetSockOptUInt64(RCVMORE)
	return value != 0, err
}

// ZMQ_HWM: Retrieve high water mark.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc5
//
func (s *Socket) HWM() (uint64, error) {
	return s.GetSockOptUInt64(HWM)
}

// ZMQ_SWAP: Retrieve disk offload size.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc6
//
func (s *Socket) Swap() (int64, error) {
	return s.GetSockOptInt64(SWAP)
}

// ZMQ_AFFINITY: Retrieve I/O thread affinity.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc7
//
func (s *Socket) Affinity() (uint64, error) {
	return s.GetSockOptUInt64(AFFINITY)
}

// ZMQ_IDENTITY: Retrieve socket identity.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc8
//
func (s *Socket) Identity() (string, error) {
	return s.GetSockOptString(IDENTITY)
}

// ZMQ_RATE: Retrieve multicast data rate.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc9
//
func (s *Socket) Rate() (int64, error) {
	return s.GetSockOptInt64(RATE)
}

// ZMQ_RECOVERY_IVL_MSEC: Get multicast recovery interval in milliseconds.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc11
//
func (s *Socket) RecoveryIvl() (time.Duration, error) {
	ms, err := s.GetSockOptInt64(RECOVERY_IVL_MSEC)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_MCAST_LOOP: Control multicast loop-back.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc12
//
func (s *Socket) McastLoop() (bool, error) {
	value, err := s.GetSockOptInt64(MCAST_LOOP)
	return value != 0, err
}

// ZMQ_SNDBUF: Retrieve kernel transmit buffer size.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc13
//
func (s *Socket) SndBuf() (uint64, error) {
	return s.GetSockOptUInt64(SNDBUF)
}

// ZMQ_RCVBUF: Retrieve kernel receive buffer size.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc14
//
func (s *Socket) RcvBuf() (uint64, error) {
	return s.GetSockOptUInt64(RCVBUF)
}

// ZMQ_LINGER: Retrieve linger period for socket shutdown.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc15
//
func (s *Socket) Linger() (time.Duration, error) {
	ms, err := s.GetSockOptInt(LINGER)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_RECONNECT_IVL: Retrieve reconnection interval.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc16
//
func (s *Socket) ReconnectIvl() (time.Duration, error) {
	ms, err := s.GetSockOptInt(RECONNECT_IVL)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_RECONNECT_IVL_MAX: Retrieve maximum reconnection interval.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc17
//
func (s *Socket) ReconnectIvlMax() (time.Duration, error) {
	ms, err := s.GetSockOptInt(RECONNECT_IVL_MAX)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_BACKLOG: Retrieve maximum length of the queue of outstanding connections.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc18
//
func (s *Socket) Backlog() (int, error) {
	return s.GetSockOptInt(BACKLOG)
}

// ZMQ_EVENTS: Retrieve socket event state.
//
// See: http://api.zeromq.org/2.1:zmq-getsockopt#toc20
//
func (s *Socket) Events() (uint64, error) {
	return s.GetSockOptUInt64(EVENTS)
}

// Socket Option Setters

// ZMQ_HWM: Set high water mark.
//
// See: http://api.zeromq.org/2.1:zmq-setsockopt#toc3
//
func (s *Socket) SetHWM(value uint64) error {
	return s.SetSockOptUInt64(HWM, value)
}

// ZMQ_SWAP: Set disk offload size.
//
// See: http://api.zeromq.org/2.1:zmq-setsockopt#toc4
//
func (s *Socket) SetSwap(value int64) error {
	return s.SetSockOptInt64(SWAP, value)
}

// ZMQ_AFFINITY: Set I/O thread affinity.
//
// See: http://api.zeromq.org/2.1:zmq-setsockopt#toc5
//
func (s *Socket) SetAffinity(value uint64) error {
	return s.SetSockOptUInt64(AFFINITY, value)
}

// ZMQ_IDENTITY: Set socket identity.
//
// See: http://api.zeromq.org/2.1:zmq-setsockopt#toc6
//
func (s *Socket) SetIdentity(value string) error {
	return s.SetSockOptString(IDENTITY, value)
}

// ZMQ_SUBSCRIBE: Establish message filter.
//
// See: http://api.zeromq.org/2.1:zmq-setsockopt#toc7
//
func (s *Socket) SetSubscribe(value string) error {
	return s.SetSockOptString(SUBSCRIBE, value)
}

// ZMQ_UNSUBSCRIBE: Remove message filter.
//
// See: http://api.zeromq.org/2.1:zmq-setsockopt#toc8
//
func (s *Socket) SetUnsubscribe(value string) error {
	return s.SetSockOptString(UNSUBSCRIBE, value)
}

// ZMQ_RATE: Set multicast data rate.
//
// See: http://api.zeromq.org/2.1:zmq-setsockopt#toc9
//
func (s *Socket) SetRate(value int64) error {
	return s.SetSockOptInt64(RATE, value)
}

// ZMQ_RECOVERY_IVL_MSEC: Set multicast recovery interval in milliseconds.
//
// See: http://api.zeromq.org/2.1:zmq-setsockopt#toc11
//
func (s *Socket) SetRecoveryIvl(value time.Duration) error {
	return s.SetSockOptInt64(RECOVERY_IVL_MSEC, int64(value/time.Millisecond))
}

// ZMQ_MCAST_LOOP: Control multicast loop-back.
//
// See: http://api.zeromq.org/2.1:zmq-setsockopt#toc12
//
func (s *Socket) SetMcastLoop(value bool) error {
	if value {
		return s.SetSockOptInt64(MCAST_LOOP, 1)
	}
	return s.SetSockOptInt64(MCAST_LOOP, 0)
}

// ZMQ_SNDBUF: Set kernel transmit buffer size.
//
// See: http://api.zeromq.org/2.1:zmq-setsockopt#toc13
//
func (s *Socket) SetSndBuf(value uint64) error {
	return s.SetSockOptUInt64(SNDBUF, value)
}

// ZMQ_RCVBUF: Set kernel receive buffer size.
//
// See: http://api.zeromq.org/2.1:zmq-setsockopt#toc14
//
func (s *Socket) SetRcvBuf(value uint64) error {
	return s.SetSockOptUInt64(RCVBUF, value)
}

// ZMQ_LINGER: Set linger period for socket shutdown.
//
// See: http://api.zeromq.org/2.1:zmq-setsockopt#toc15
//
func (s *Socket) SetLinger(value time.Duration) error {
	return s.SetSockOptInt(LINGER, int(value/time.Millisecond))
}

// ZMQ_RECONNECT_IVL: Set reconnection interval.
//
// See: http://api.zeromq.org/2.1:zmq-setsockopt#toc16
//
func (s *Socket) SetReconnectIvl(value time.Duration) error {
	return s.SetSockOptInt(RECONNECT_IVL, int(value/time.Millisecond))
}

// ZMQ_RECONNECT_IVL_MAX: Set maximum reconnection interval.
//
// See: http://api.zeromq.org/2.1:zmq-setsockopt#toc17
//
func (s *Socket) SetReconnectIvlMax(value time.Duration) error {
	return s.SetSockOptInt(RECONNECT_IVL_MAX, int(value/time.Millisecond))
}

// ZMQ_BACKLOG: Set maximum length of the queue of outstanding connections.
//
// See: http://api.zeromq.org/2.1:zmq-setsockopt#toc18
//
func (s *Socket) SetBacklog(value int) error {
	return s.SetSockOptInt(BACKLOG, value)
}
