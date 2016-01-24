// +build zmq_4_x
//

package gozmq

import (
	"time"
)

// This file was generated automatically.  Changes made here will be lost.

// Socket Option Getters

// ZMQ_TYPE: Retrieve socket type.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc3
//
func (s *Socket) Type() (SocketType, error) {
	value, err := s.GetSockOptUInt64(TYPE)
	return SocketType(value), err
}

// ZMQ_RCVMORE: More message data parts to follow.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc4
//
func (s *Socket) RcvMore() (bool, error) {
	value, err := s.GetSockOptInt(RCVMORE)
	return value != 0, err
}

// ZMQ_SNDHWM: Retrieves high water mark for outbound messages.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc5
//
func (s *Socket) SndHWM() (int, error) {
	return s.GetSockOptInt(SNDHWM)
}

// ZMQ_RCVHWM: Retrieve high water mark for inbound messages.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc6
//
func (s *Socket) RcvHWM() (int, error) {
	return s.GetSockOptInt(RCVHWM)
}

// ZMQ_AFFINITY: Retrieve I/O thread affinity.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc7
//
func (s *Socket) Affinity() (uint64, error) {
	return s.GetSockOptUInt64(AFFINITY)
}

// ZMQ_IDENTITY: Retrieve socket identity.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc8
//
func (s *Socket) Identity() (string, error) {
	return s.GetSockOptString(IDENTITY)
}

// ZMQ_RATE: Retrieve multicast data rate.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc9
//
func (s *Socket) Rate() (int64, error) {
	return s.GetSockOptInt64(RATE)
}

// ZMQ_RECOVERY_IVL: Get multicast recovery interval.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc10
//
func (s *Socket) RecoveryIvl() (time.Duration, error) {
	ms, err := s.GetSockOptInt64(RECOVERY_IVL)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_SNDBUF: Retrieve kernel transmit buffer size.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc11
//
func (s *Socket) SndBuf() (uint64, error) {
	return s.GetSockOptUInt64(SNDBUF)
}

// ZMQ_RCVBUF: Retrieve kernel receive buffer size.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc12
//
func (s *Socket) RcvBuf() (uint64, error) {
	return s.GetSockOptUInt64(RCVBUF)
}

// ZMQ_LINGER: Retrieve linger period for socket shutdown.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc13
//
func (s *Socket) Linger() (time.Duration, error) {
	ms, err := s.GetSockOptInt(LINGER)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_RECONNECT_IVL: Retrieve reconnection interval.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc14
//
func (s *Socket) ReconnectIvl() (time.Duration, error) {
	ms, err := s.GetSockOptInt(RECONNECT_IVL)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_RECONNECT_IVL_MAX: Retrieve maximum reconnection interval.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc15
//
func (s *Socket) ReconnectIvlMax() (time.Duration, error) {
	ms, err := s.GetSockOptInt(RECONNECT_IVL_MAX)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_BACKLOG: Retrieve maximum length of the queue of outstanding connections.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc16
//
func (s *Socket) Backlog() (int, error) {
	return s.GetSockOptInt(BACKLOG)
}

// ZMQ_MAXMSGSIZE: Maximum acceptable inbound message size.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc17
//
func (s *Socket) MaxMsgSize() (int64, error) {
	return s.GetSockOptInt64(MAXMSGSIZE)
}

// ZMQ_RCVTIMEO: Maximum time before a socket operation returns with EAGAIN.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc19
//
func (s *Socket) RcvTimeout() (time.Duration, error) {
	ms, err := s.GetSockOptInt(RCVTIMEO)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_SNDTIMEO: Maximum time before a socket operation returns with EAGAIN.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc20
//
func (s *Socket) SndTimeout() (time.Duration, error) {
	ms, err := s.GetSockOptInt(SNDTIMEO)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_IPV6: Retrieve IPv6 socket status.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc21
//
func (s *Socket) Ipv6() (bool, error) {
	value, err := s.GetSockOptInt(IPV6)
	return value != 0, err
}

// ZMQ_IPV4ONLY: Retrieve IPv4-only socket override status.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc22
//
func (s *Socket) IPv4Only() (bool, error) {
	value, err := s.GetSockOptInt(IPV4ONLY)
	return value != 0, err
}

// ZMQ_IMMEDIATE: Retrieve attach-on-connect value.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc23
//
func (s *Socket) Immediate() (bool, error) {
	value, err := s.GetSockOptInt(IMMEDIATE)
	return value != 0, err
}

// ZMQ_EVENTS: Retrieve socket event state.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc25
//
func (s *Socket) Events() (uint64, error) {
	return s.GetSockOptUInt64(EVENTS)
}

// ZMQ_TCP_KEEPALIVE: Override SO_KEEPALIVE socket option.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc27
//
func (s *Socket) TCPKeepalive() (int, error) {
	return s.GetSockOptInt(TCP_KEEPALIVE)
}

// ZMQ_TCP_KEEPALIVE_IDLE: Override TCP_KEEPCNT(or TCP_KEEPALIVE on some OS).
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc28
//
func (s *Socket) TCPKeepaliveIdle() (int, error) {
	return s.GetSockOptInt(TCP_KEEPALIVE_IDLE)
}

// ZMQ_TCP_KEEPALIVE_CNT: Override TCP_KEEPCNT socket option.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc29
//
func (s *Socket) TCPKeepaliveCnt() (int, error) {
	return s.GetSockOptInt(TCP_KEEPALIVE_CNT)
}

// ZMQ_TCP_KEEPALIVE_INTVL: Override TCP_KEEPINTVL socket option.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc30
//
func (s *Socket) TCPKeepaliveIntvl() (int, error) {
	return s.GetSockOptInt(TCP_KEEPALIVE_INTVL)
}

// ZMQ_MECHANISM: Retrieve current security mechanism.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc31
//
func (s *Socket) Mechanism() (int, error) {
	return s.GetSockOptInt(MECHANISM)
}

// ZMQ_PLAIN_SERVER: Retrieve current PLAIN server role.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc32
//
func (s *Socket) PlainServer() (int, error) {
	return s.GetSockOptInt(PLAIN_SERVER)
}

// ZMQ_PLAIN_USERNAME: Retrieve current PLAIN username.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc33
//
func (s *Socket) PlainUsername() (string, error) {
	return s.GetSockOptString(PLAIN_USERNAME)
}

// ZMQ_PLAIN_PASSWORD: Retrieve current password.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc34
//
func (s *Socket) PlainPassword() (string, error) {
	return s.GetSockOptString(PLAIN_PASSWORD)
}

// ZMQ_CURVE_PUBLICKEY: Retrieve current CURVE public key.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc35
//
func (s *Socket) CurvePublickey() (string, error) {
	return s.GetSockOptString(CURVE_PUBLICKEY)
}

// ZMQ_CURVE_SECRETKEY: Retrieve current CURVE secret key.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc36
//
func (s *Socket) CurveSecretkey() (string, error) {
	return s.GetSockOptString(CURVE_SECRETKEY)
}

// ZMQ_CURVE_SERVERKEY: Retrieve current CURVE server key.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc37
//
func (s *Socket) CurveServerkey() (string, error) {
	return s.GetSockOptString(CURVE_SERVERKEY)
}

// ZMQ_ZAP_DOMAIN: Retrieve RFC 27 authentication domain.
//
// See: http://api.zeromq.org/4.0:zmq-getsockopt#toc38
//
func (s *Socket) ZapDomain() (string, error) {
	return s.GetSockOptString(ZAP_DOMAIN)
}

// Socket Option Setters

// ZMQ_SNDHWM: Set high water mark for outbound messages.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc3
//
func (s *Socket) SetSndHWM(value int) error {
	return s.SetSockOptInt(SNDHWM, value)
}

// ZMQ_RCVHWM: Set high water mark for inbound messages.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc4
//
func (s *Socket) SetRcvHWM(value int) error {
	return s.SetSockOptInt(RCVHWM, value)
}

// ZMQ_AFFINITY: Set I/O thread affinity.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc5
//
func (s *Socket) SetAffinity(value uint64) error {
	return s.SetSockOptUInt64(AFFINITY, value)
}

// ZMQ_SUBSCRIBE: Establish message filter.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc6
//
func (s *Socket) SetSubscribe(value string) error {
	return s.SetSockOptString(SUBSCRIBE, value)
}

// ZMQ_UNSUBSCRIBE: Remove message filter.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc7
//
func (s *Socket) SetUnsubscribe(value string) error {
	return s.SetSockOptString(UNSUBSCRIBE, value)
}

// ZMQ_IDENTITY: Set socket identity.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc8
//
func (s *Socket) SetIdentity(value string) error {
	return s.SetSockOptString(IDENTITY, value)
}

// ZMQ_RATE: Set multicast data rate.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc9
//
func (s *Socket) SetRate(value int64) error {
	return s.SetSockOptInt64(RATE, value)
}

// ZMQ_RECOVERY_IVL: Set multicast recovery interval.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc10
//
func (s *Socket) SetRecoveryIvl(value time.Duration) error {
	return s.SetSockOptInt64(RECOVERY_IVL, int64(value/time.Millisecond))
}

// ZMQ_SNDBUF: Set kernel transmit buffer size.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc11
//
func (s *Socket) SetSndBuf(value uint64) error {
	return s.SetSockOptUInt64(SNDBUF, value)
}

// ZMQ_RCVBUF: Set kernel receive buffer size.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc12
//
func (s *Socket) SetRcvBuf(value uint64) error {
	return s.SetSockOptUInt64(RCVBUF, value)
}

// ZMQ_LINGER: Set linger period for socket shutdown.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc13
//
func (s *Socket) SetLinger(value time.Duration) error {
	return s.SetSockOptInt(LINGER, int(value/time.Millisecond))
}

// ZMQ_RECONNECT_IVL: Set reconnection interval.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc14
//
func (s *Socket) SetReconnectIvl(value time.Duration) error {
	return s.SetSockOptInt(RECONNECT_IVL, int(value/time.Millisecond))
}

// ZMQ_RECONNECT_IVL_MAX: Set maximum reconnection interval.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc15
//
func (s *Socket) SetReconnectIvlMax(value time.Duration) error {
	return s.SetSockOptInt(RECONNECT_IVL_MAX, int(value/time.Millisecond))
}

// ZMQ_BACKLOG: Set maximum length of the queue of outstanding connections.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc16
//
func (s *Socket) SetBacklog(value int) error {
	return s.SetSockOptInt(BACKLOG, value)
}

// ZMQ_MAXMSGSIZE: Maximum acceptable inbound message size.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc17
//
func (s *Socket) SetMaxMsgSize(value int64) error {
	return s.SetSockOptInt64(MAXMSGSIZE, value)
}

// ZMQ_RCVTIMEO: Maximum time before a recv operation returns with EAGAIN.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc19
//
func (s *Socket) SetRcvTimeout(value time.Duration) error {
	return s.SetSockOptInt(RCVTIMEO, int(value/time.Millisecond))
}

// ZMQ_SNDTIMEO: Maximum time before a send operation returns with EAGAIN.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc20
//
func (s *Socket) SetSndTimeout(value time.Duration) error {
	return s.SetSockOptInt(SNDTIMEO, int(value/time.Millisecond))
}

// ZMQ_IPV6: Enable IPv6 on socket.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc21
//
func (s *Socket) SetIpv6(value bool) error {
	if value {
		return s.SetSockOptInt(IPV6, 1)
	}
	return s.SetSockOptInt(IPV6, 0)
}

// ZMQ_IPV4ONLY: Use IPv4-only on socket.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc22
//
func (s *Socket) SetIPv4Only(value bool) error {
	if value {
		return s.SetSockOptInt(IPV4ONLY, 1)
	}
	return s.SetSockOptInt(IPV4ONLY, 0)
}

// ZMQ_IMMEDIATE: Queue messages only to completed connections.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc23
//
func (s *Socket) SetImmediate(value bool) error {
	if value {
		return s.SetSockOptInt(IMMEDIATE, 1)
	}
	return s.SetSockOptInt(IMMEDIATE, 0)
}

// ZMQ_ROUTER_MANDATORY: accept only routable messages on ROUTER sockets.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc24
//
func (s *Socket) SetROUTERMandatory(value bool) error {
	if value {
		return s.SetSockOptInt(ROUTER_MANDATORY, 1)
	}
	return s.SetSockOptInt(ROUTER_MANDATORY, 0)
}

// ZMQ_ROUTER_RAW: switch ROUTER socket to raw mode.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc25
//
func (s *Socket) SetROUTERRaw(value int) error {
	return s.SetSockOptInt(ROUTER_RAW, value)
}

// ZMQ_PROBE_ROUTER: bootstrap connections to ROUTER sockets.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc26
//
func (s *Socket) SetProbeROUTER(value int) error {
	return s.SetSockOptInt(PROBE_ROUTER, value)
}

// ZMQ_XPUB_VERBOSE: provide all subscription messages on XPUB sockets.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc27
//
func (s *Socket) SetXPUBVerbose(value bool) error {
	if value {
		return s.SetSockOptInt(XPUB_VERBOSE, 1)
	}
	return s.SetSockOptInt(XPUB_VERBOSE, 0)
}

// ZMQ_REQ_CORRELATE: match replies with requests.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc28
//
func (s *Socket) SetReqCorrelate(value int) error {
	return s.SetSockOptInt(REQ_CORRELATE, value)
}

// ZMQ_REQ_RELAXED: relax strict alternation between request and reply.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc29
//
func (s *Socket) SetReqRelaxed(value int) error {
	return s.SetSockOptInt(REQ_RELAXED, value)
}

// ZMQ_TCP_KEEPALIVE: Override SO_KEEPALIVE socket option.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc30
//
func (s *Socket) SetTCPKeepalive(value int) error {
	return s.SetSockOptInt(TCP_KEEPALIVE, value)
}

// ZMQ_TCP_KEEPALIVE_IDLE: Override TCP_KEEPCNT (or TCP_KEEPALIVE on some OS).
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc31
//
func (s *Socket) SetTCPKeepaliveIdle(value int) error {
	return s.SetSockOptInt(TCP_KEEPALIVE_IDLE, value)
}

// ZMQ_TCP_KEEPALIVE_CNT: Override TCP_KEEPCNT socket option.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc32
//
func (s *Socket) SetTCPKeepaliveCnt(value int) error {
	return s.SetSockOptInt(TCP_KEEPALIVE_CNT, value)
}

// ZMQ_TCP_KEEPALIVE_INTVL: Override TCP_KEEPINTVL socket option.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc33
//
func (s *Socket) SetTCPKeepaliveIntvl(value int) error {
	return s.SetSockOptInt(TCP_KEEPALIVE_INTVL, value)
}

// ZMQ_TCP_ACCEPT_FILTER: Assign filters to allow new TCP connections.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc34
//
func (s *Socket) SetTCPAcceptFilter(value string) error {
	return s.SetSockOptString(TCP_ACCEPT_FILTER, value)
}

// ZMQ_PLAIN_SERVER: Set PLAIN server role.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc35
//
func (s *Socket) SetPlainServer(value int) error {
	return s.SetSockOptInt(PLAIN_SERVER, value)
}

// ZMQ_PLAIN_USERNAME: Set PLAIN security username.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc36
//
func (s *Socket) SetPlainUsername(value string) error {
	return s.SetSockOptString(PLAIN_USERNAME, value)
}

// ZMQ_PLAIN_PASSWORD: Set PLAIN security password.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc37
//
func (s *Socket) SetPlainPassword(value string) error {
	return s.SetSockOptString(PLAIN_PASSWORD, value)
}

// ZMQ_CURVE_SERVER: Set CURVE server role.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc38
//
func (s *Socket) SetCurveServer(value int) error {
	return s.SetSockOptInt(CURVE_SERVER, value)
}

// ZMQ_CURVE_PUBLICKEY: Set CURVE public key.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc39
//
func (s *Socket) SetCurvePublickey(value string) error {
	return s.SetSockOptString(CURVE_PUBLICKEY, value)
}

// ZMQ_CURVE_SECRETKEY: Set CURVE secret key.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc40
//
func (s *Socket) SetCurveSecretkey(value string) error {
	return s.SetSockOptString(CURVE_SECRETKEY, value)
}

// ZMQ_CURVE_SERVERKEY: Set CURVE server key.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc41
//
func (s *Socket) SetCurveServerkey(value string) error {
	return s.SetSockOptString(CURVE_SERVERKEY, value)
}

// ZMQ_ZAP_DOMAIN: Set RFC 27 authentication domain.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc42
//
func (s *Socket) SetZapDomain(value string) error {
	return s.SetSockOptString(ZAP_DOMAIN, value)
}

// ZMQ_CONFLATE: Keep only last message.
//
// See: http://api.zeromq.org/4.0:zmq-setsockopt#toc43
//
func (s *Socket) SetConflate(value bool) error {
	if value {
		return s.SetSockOptInt(CONFLATE, 1)
	}
	return s.SetSockOptInt(CONFLATE, 0)
}
