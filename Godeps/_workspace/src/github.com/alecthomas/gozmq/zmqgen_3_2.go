// +build zmq_3_x
//

package gozmq

import (
	"time"
)

// This file was generated automatically.  Changes made here will be lost.

// Socket Option Getters

// ZMQ_TYPE: Retrieve socket type.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc3
//
func (s *Socket) Type() (SocketType, error) {
	value, err := s.GetSockOptUInt64(TYPE)
	return SocketType(value), err
}

// ZMQ_RCVMORE: More message data parts to follow.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc4
//
func (s *Socket) RcvMore() (bool, error) {
	value, err := s.GetSockOptInt(RCVMORE)
	return value != 0, err
}

// ZMQ_SNDHWM: Retrieves high water mark for outbound messages.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc5
//
func (s *Socket) SndHWM() (int, error) {
	return s.GetSockOptInt(SNDHWM)
}

// ZMQ_RCVHWM: Retrieve high water mark for inbound messages.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc6
//
func (s *Socket) RcvHWM() (int, error) {
	return s.GetSockOptInt(RCVHWM)
}

// ZMQ_AFFINITY: Retrieve I/O thread affinity.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc7
//
func (s *Socket) Affinity() (uint64, error) {
	return s.GetSockOptUInt64(AFFINITY)
}

// ZMQ_IDENTITY: Set socket identity.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc8
//
func (s *Socket) Identity() (string, error) {
	return s.GetSockOptString(IDENTITY)
}

// ZMQ_RATE: Retrieve multicast data rate.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc9
//
func (s *Socket) Rate() (int64, error) {
	return s.GetSockOptInt64(RATE)
}

// ZMQ_RECOVERY_IVL: Get multicast recovery interval.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc10
//
func (s *Socket) RecoveryIvl() (time.Duration, error) {
	ms, err := s.GetSockOptInt64(RECOVERY_IVL)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_SNDBUF: Retrieve kernel transmit buffer size.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc11
//
func (s *Socket) SndBuf() (uint64, error) {
	return s.GetSockOptUInt64(SNDBUF)
}

// ZMQ_RCVBUF: Retrieve kernel receive buffer size.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc12
//
func (s *Socket) RcvBuf() (uint64, error) {
	return s.GetSockOptUInt64(RCVBUF)
}

// ZMQ_LINGER: Retrieve linger period for socket shutdown.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc13
//
func (s *Socket) Linger() (time.Duration, error) {
	ms, err := s.GetSockOptInt(LINGER)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_RECONNECT_IVL: Retrieve reconnection interval.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc14
//
func (s *Socket) ReconnectIvl() (time.Duration, error) {
	ms, err := s.GetSockOptInt(RECONNECT_IVL)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_RECONNECT_IVL_MAX: Retrieve maximum reconnection interval.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc15
//
func (s *Socket) ReconnectIvlMax() (time.Duration, error) {
	ms, err := s.GetSockOptInt(RECONNECT_IVL_MAX)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_BACKLOG: Retrieve maximum length of the queue of outstanding connections.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc16
//
func (s *Socket) Backlog() (int, error) {
	return s.GetSockOptInt(BACKLOG)
}

// ZMQ_MAXMSGSIZE: Maximum acceptable inbound message size.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc17
//
func (s *Socket) MaxMsgSize() (int64, error) {
	return s.GetSockOptInt64(MAXMSGSIZE)
}

// ZMQ_RCVTIMEO: Maximum time before a socket operation returns with EAGAIN.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc19
//
func (s *Socket) RcvTimeout() (time.Duration, error) {
	ms, err := s.GetSockOptInt(RCVTIMEO)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_SNDTIMEO: Maximum time before a socket operation returns with EAGAIN.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc20
//
func (s *Socket) SndTimeout() (time.Duration, error) {
	ms, err := s.GetSockOptInt(SNDTIMEO)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_IPV4ONLY: Retrieve IPv4-only socket override status.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc21
//
func (s *Socket) IPv4Only() (bool, error) {
	value, err := s.GetSockOptInt(IPV4ONLY)
	return value != 0, err
}

// ZMQ_DELAY_ATTACH_ON_CONNECT: Retrieve attach-on-connect value.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc22
//
func (s *Socket) DelayAttachOnConnect() (bool, error) {
	value, err := s.GetSockOptInt(DELAY_ATTACH_ON_CONNECT)
	return value != 0, err
}

// ZMQ_EVENTS: Retrieve socket event state.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc24
//
func (s *Socket) Events() (uint64, error) {
	return s.GetSockOptUInt64(EVENTS)
}

// ZMQ_TCP_KEEPALIVE: Override SO_KEEPALIVE socket option.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc26
//
func (s *Socket) TCPKeepalive() (int, error) {
	return s.GetSockOptInt(TCP_KEEPALIVE)
}

// ZMQ_TCP_KEEPALIVE_IDLE: Override TCP_KEEPCNT(or TCP_KEEPALIVE on some OS).
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc27
//
func (s *Socket) TCPKeepaliveIdle() (int, error) {
	return s.GetSockOptInt(TCP_KEEPALIVE_IDLE)
}

// ZMQ_TCP_KEEPALIVE_CNT: Override TCP_KEEPCNT socket option.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc28
//
func (s *Socket) TCPKeepaliveCnt() (int, error) {
	return s.GetSockOptInt(TCP_KEEPALIVE_CNT)
}

// ZMQ_TCP_KEEPALIVE_INTVL: Override TCP_KEEPINTVL socket option.
//
// See: http://api.zeromq.org/3.2:zmq-getsockopt#toc29
//
func (s *Socket) TCPKeepaliveIntvl() (int, error) {
	return s.GetSockOptInt(TCP_KEEPALIVE_INTVL)
}

// Socket Option Setters

// ZMQ_SNDHWM: Set high water mark for outbound messages.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc3
//
func (s *Socket) SetSndHWM(value int) error {
	return s.SetSockOptInt(SNDHWM, value)
}

// ZMQ_RCVHWM: Set high water mark for inbound messages.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc4
//
func (s *Socket) SetRcvHWM(value int) error {
	return s.SetSockOptInt(RCVHWM, value)
}

// ZMQ_AFFINITY: Set I/O thread affinity.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc5
//
func (s *Socket) SetAffinity(value uint64) error {
	return s.SetSockOptUInt64(AFFINITY, value)
}

// ZMQ_SUBSCRIBE: Establish message filter.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc6
//
func (s *Socket) SetSubscribe(value string) error {
	return s.SetSockOptString(SUBSCRIBE, value)
}

// ZMQ_UNSUBSCRIBE: Remove message filter.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc7
//
func (s *Socket) SetUnsubscribe(value string) error {
	return s.SetSockOptString(UNSUBSCRIBE, value)
}

// ZMQ_IDENTITY: Set socket identity.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc8
//
func (s *Socket) SetIdentity(value string) error {
	return s.SetSockOptString(IDENTITY, value)
}

// ZMQ_RATE: Set multicast data rate.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc9
//
func (s *Socket) SetRate(value int64) error {
	return s.SetSockOptInt64(RATE, value)
}

// ZMQ_RECOVERY_IVL: Set multicast recovery interval.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc10
//
func (s *Socket) SetRecoveryIvl(value time.Duration) error {
	return s.SetSockOptInt64(RECOVERY_IVL, int64(value/time.Millisecond))
}

// ZMQ_SNDBUF: Set kernel transmit buffer size.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc11
//
func (s *Socket) SetSndBuf(value uint64) error {
	return s.SetSockOptUInt64(SNDBUF, value)
}

// ZMQ_RCVBUF: Set kernel receive buffer size.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc12
//
func (s *Socket) SetRcvBuf(value uint64) error {
	return s.SetSockOptUInt64(RCVBUF, value)
}

// ZMQ_LINGER: Set linger period for socket shutdown.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc13
//
func (s *Socket) SetLinger(value time.Duration) error {
	return s.SetSockOptInt(LINGER, int(value/time.Millisecond))
}

// ZMQ_RECONNECT_IVL: Set reconnection interval.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc14
//
func (s *Socket) SetReconnectIvl(value time.Duration) error {
	return s.SetSockOptInt(RECONNECT_IVL, int(value/time.Millisecond))
}

// ZMQ_RECONNECT_IVL_MAX: Set maximum reconnection interval.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc15
//
func (s *Socket) SetReconnectIvlMax(value time.Duration) error {
	return s.SetSockOptInt(RECONNECT_IVL_MAX, int(value/time.Millisecond))
}

// ZMQ_BACKLOG: Set maximum length of the queue of outstanding connections.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc16
//
func (s *Socket) SetBacklog(value int) error {
	return s.SetSockOptInt(BACKLOG, value)
}

// ZMQ_MAXMSGSIZE: Maximum acceptable inbound message size.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc17
//
func (s *Socket) SetMaxMsgSize(value int64) error {
	return s.SetSockOptInt64(MAXMSGSIZE, value)
}

// ZMQ_RCVTIMEO: Maximum time before a recv operation returns with EAGAIN.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc19
//
func (s *Socket) SetRcvTimeout(value time.Duration) error {
	return s.SetSockOptInt(RCVTIMEO, int(value/time.Millisecond))
}

// ZMQ_SNDTIMEO: Maximum time before a send operation returns with EAGAIN.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc20
//
func (s *Socket) SetSndTimeout(value time.Duration) error {
	return s.SetSockOptInt(SNDTIMEO, int(value/time.Millisecond))
}

// ZMQ_IPV4ONLY: Use IPv4-only sockets.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc21
//
func (s *Socket) SetIPv4Only(value bool) error {
	if value {
		return s.SetSockOptInt(IPV4ONLY, 1)
	}
	return s.SetSockOptInt(IPV4ONLY, 0)
}

// ZMQ_DELAY_ATTACH_ON_CONNECT: Accept messages only when connections are made.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc22
//
func (s *Socket) SetDelayAttachOnConnect(value bool) error {
	if value {
		return s.SetSockOptInt(DELAY_ATTACH_ON_CONNECT, 1)
	}
	return s.SetSockOptInt(DELAY_ATTACH_ON_CONNECT, 0)
}

// ZMQ_ROUTER_MANDATORY: accept only routable messages on ROUTER sockets.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc23
//
func (s *Socket) SetROUTERMandatory(value bool) error {
	if value {
		return s.SetSockOptInt(ROUTER_MANDATORY, 1)
	}
	return s.SetSockOptInt(ROUTER_MANDATORY, 0)
}

// ZMQ_XPUB_VERBOSE: provide all subscription messages on XPUB sockets.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc24
//
func (s *Socket) SetXPUBVerbose(value bool) error {
	if value {
		return s.SetSockOptInt(XPUB_VERBOSE, 1)
	}
	return s.SetSockOptInt(XPUB_VERBOSE, 0)
}

// ZMQ_TCP_KEEPALIVE: Override SO_KEEPALIVE socket option.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc25
//
func (s *Socket) SetTCPKeepalive(value int) error {
	return s.SetSockOptInt(TCP_KEEPALIVE, value)
}

// ZMQ_TCP_KEEPALIVE_IDLE: Override TCP_KEEPCNT(or TCP_KEEPALIVE on some OS).
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc26
//
func (s *Socket) SetTCPKeepaliveIdle(value int) error {
	return s.SetSockOptInt(TCP_KEEPALIVE_IDLE, value)
}

// ZMQ_TCP_KEEPALIVE_CNT: Override TCP_KEEPCNT socket option.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc27
//
func (s *Socket) SetTCPKeepaliveCnt(value int) error {
	return s.SetSockOptInt(TCP_KEEPALIVE_CNT, value)
}

// ZMQ_TCP_KEEPALIVE_INTVL: Override TCP_KEEPINTVL socket option.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc28
//
func (s *Socket) SetTCPKeepaliveIntvl(value int) error {
	return s.SetSockOptInt(TCP_KEEPALIVE_INTVL, value)
}

// ZMQ_TCP_ACCEPT_FILTER: Assign filters to allow new TCP connections.
//
// See: http://api.zeromq.org/3.2:zmq-setsockopt#toc29
//
func (s *Socket) SetTCPAcceptFilter(value string) error {
	return s.SetSockOptString(TCP_ACCEPT_FILTER, value)
}
