# _NOTE:_ These gozmq bindings are in maintenance mode. Only critical bugs will be fixed. Henceforth I would suggest using [@pebbe's](https://github.com/pebbe) actively maintained bindings for [zmq2](https://github.com/pebbe/zmq2), [zmq3](https://github.com/pebbe/zmq3) and [zmq4](https://github.com/pebbe/zmq4).

## Go (golang) Bindings for 0mq (zmq, zeromq)

[![Build Status](https://travis-ci.org/alecthomas/gozmq.png)](https://travis-ci.org/alecthomas/gozmq)

This package implements [Go](http://golang.org) (golang) bindings for
the [0mq](http://zeromq.org) C API.

GoZMQ [does not](#zero-copy) support zero-copy.

A full list of examples is included in the [zguide](https://github.com/imatix/zguide/tree/master/examples/Go).

Note that this is *not* the same as [this
implementation](http://github.com/boggle/gozero) or [this
implementation](http://code.google.com/p/gozmq/).

## Upgrading

GoZMQ has made some public changes that will break old code.  Fortunately, we've also written a tool based on `go fix` that will upgrade your code for you!  Here's how to run it over your source (after making a backup of course):

    go get github.com/alecthomas/gozmq/gozmqfix
    cd $YOUR_SOURCE_DIR
    gozmqfix .

## Installing

GoZMQ currently supports ZMQ 2.1.x, 2.2.x, 3.x and 4.x. Following are instructions on how to compile against these versions.

For ZeroMQ 2.2.x install with:

    go get github.com/alecthomas/gozmq

For 2.1.x install with:

    go get -tags zmq_2_1 github.com/alecthomas/gozmq

For 3.x install with:

    go get -tags zmq_3_x github.com/alecthomas/gozmq
    
For 4.x install with:

    go get -tags zmq_4_x github.com/alecthomas/gozmq

### Troubleshooting

#### Go can't find ZMQ

If the go tool can't find zmq and you know it is installed, you may need to override the C compiler/linker flags.

eg. If you installed zmq into `/opt/zmq` you might try:

  CGO_CFLAGS=-I/opt/zmq/include CGO_LDFLAGS=-L/opt/zmq/lib \
    go get github.com/alecthomas/gozmq

#### Mismatch in version of ZMQ

If you get errors like this with 'go get' or 'go build':

    1: error: 'ZMQ_FOO' undeclared (first use in this function)

There are two possibilities:

1. Your version of zmq is *very* old. In this case you will need to download and build zmq yourself.
2. You are building gozmq against the wrong version of zmq. See the [installation](#installation) instructions for details on how to target the correct version.

## Differences from the C API

The API implemented by this package does not attempt to expose
`zmq_msg_t` at all. Instead, `Recv()` and `Send()` both operate on byte
slices, allocating and freeing the memory automatically. Currently this
requires copying to/from C malloced memory, but a future implementation
may be able to avoid this to a certain extent.

All major features are supported: contexts, sockets, devices, and polls.

## Example

Here are direct translations of some of the examples from [this blog
post](http://nichol.as/zeromq-an-introduction).

A simple echo server:

```go
package main

import zmq "github.com/alecthomas/gozmq"

func main() {
  context, _ := zmq.NewContext()
  socket, _ := context.NewSocket(zmq.REP)
  socket.Bind("tcp://127.0.0.1:5000")
  socket.Bind("tcp://127.0.0.1:6000")

  for {
    msg, _ := socket.Recv(0)
    println("Got", string(msg))
    socket.Send(msg, 0)
  }
}
```

A simple client for the above server:

```go
package main

import "fmt"
import zmq "github.com/alecthomas/gozmq"

func main() {
  context, _ := zmq.NewContext()
  socket, _ := context.NewSocket(zmq.REQ)
  socket.Connect("tcp://127.0.0.1:5000")
  socket.Connect("tcp://127.0.0.1:6000")

  for i := 0; i < 10; i++ {
    msg := fmt.Sprintf("msg %d", i)
    socket.Send([]byte(msg), 0)
    println("Sending", msg)
    socket.Recv(0)
  }
}
```

## Caveats

### Zero-copy

GoZMQ does not support zero-copy.

GoZMQ does not attempt to expose `zmq_msg_t` at all. Instead, `Recv()` and `Send()`
both operate on byte slices, allocating and freeing the memory automatically.
Currently this requires copying to/from C malloced memory, but a future
implementation may be able to avoid this to a certain extent.


### Memory management

It's not entirely clear from the 0mq documentation how memory for
`zmq_msg_t` and packet data is managed once 0mq takes ownership. After
digging into the source a little, this package operates under the
following (educated) assumptions:

-   References to `zmq_msg_t` structures are not held by the C API
    beyond the duration of any function call.
-   Packet data is reference counted internally by the C API. The count
    is incremented when a packet is queued for delivery to a destination
    (the inference being that for delivery to N destinations, the
    reference count will be incremented N times) and decremented once
    the packet has either been delivered or errored.

## Usage

```go
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
```

```go
const (
  POLLIN  = PollEvents(C.ZMQ_POLLIN)
  POLLOUT = PollEvents(C.ZMQ_POLLOUT)
  POLLERR = PollEvents(C.ZMQ_POLLERR)
)
```

```go
const (
  STREAMER  = DeviceType(C.ZMQ_STREAMER)
  FORWARDER = DeviceType(C.ZMQ_FORWARDER)
  QUEUE     = DeviceType(C.ZMQ_QUEUE)
)
```

```go
const (
  RCVTIMEO = IntSocketOption(C.ZMQ_RCVTIMEO)
  SNDTIMEO = IntSocketOption(C.ZMQ_SNDTIMEO)
)
```

```go
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
```

```go
const (
  RCVMORE = IntSocketOption(C.ZMQ_RCVMORE)
  SNDHWM  = IntSocketOption(C.ZMQ_SNDHWM)
  RCVHWM  = IntSocketOption(C.ZMQ_RCVHWM)

  // TODO Not documented in the man page...
  //LAST_ENDPOINT       = UInt64SocketOption(C.ZMQ_LAST_ENDPOINT)
  FAIL_UNROUTABLE     = BoolSocketOption(C.ZMQ_FAIL_UNROUTABLE)
  TCP_KEEPALIVE       = IntSocketOption(C.ZMQ_TCP_KEEPALIVE)
  TCP_KEEPALIVE_CNT   = IntSocketOption(C.ZMQ_TCP_KEEPALIVE_CNT)
  TCP_KEEPALIVE_IDLE  = IntSocketOption(C.ZMQ_TCP_KEEPALIVE_IDLE)
  TCP_KEEPALIVE_INTVL = IntSocketOption(C.ZMQ_TCP_KEEPALIVE_INTVL)
  TCP_ACCEPT_FILTER   = StringSocketOption(C.ZMQ_TCP_ACCEPT_FILTER)

  // Message options
  MORE = MessageOption(C.ZMQ_MORE)

  // Send/recv options
  DONTWAIT = SendRecvOption(C.ZMQ_DONTWAIT)

  // Deprecated aliases
  NOBLOCK = DONTWAIT
)
```

```go
var (
  // Additional ZMQ errors
  ENOTSOCK       error = zmqErrno(C.ENOTSOCK)
  EFSM           error = zmqErrno(C.EFSM)
  EINVAL         error = zmqErrno(C.EINVAL)
  ENOCOMPATPROTO error = zmqErrno(C.ENOCOMPATPROTO)
  ETERM          error = zmqErrno(C.ETERM)
  EMTHREAD       error = zmqErrno(C.EMTHREAD)
)
```

#### func  Device

```go
func Device(t DeviceType, in, out *Socket) error
```
run a zmq_device passing messages between in and out

#### func  Poll

```go
func Poll(items []PollItem, timeout time.Duration) (count int, err error)
```
Poll ZmqSockets and file descriptors for I/O readiness. Timeout is in
time.Duration. The smallest possible timeout is time.Millisecond for ZeroMQ
version 3 and above, and time.Microsecond for earlier versions.

#### func  Proxy

```go
func Proxy(in, out, capture *Socket) error
```
run a zmq_proxy with in, out and capture sockets

#### func  Version

```go
func Version() (int, int, int)
```
void zmq_version (int *major, int *minor, int *patch);

#### type BoolSocketOption

```go
type BoolSocketOption int
```


#### type Context

```go
type Context struct {
}
```

* A context handles socket creation and asynchronous message delivery. * There
should generally be one context per application.

#### func  NewContext

```go
func NewContext() (*Context, error)
```
Create a new context.

#### func (*Context) Close

```go
func (c *Context) Close()
```

#### func (*Context) IOThreads

```go
func (c *Context) IOThreads() (int, error)
```
Get a context option.

#### func (*Context) MaxSockets

```go
func (c *Context) MaxSockets() (int, error)
```

#### func (*Context) NewSocket

```go
func (c *Context) NewSocket(t SocketType) (*Socket, error)
```
Create a new socket. void *zmq_socket (void *context, int type);

#### func (*Context) SetIOThreads

```go
func (c *Context) SetIOThreads(value int) error
```
Set a context option.

#### func (*Context) SetMaxSockets

```go
func (c *Context) SetMaxSockets(value int) error
```

#### type DeviceType

```go
type DeviceType int
```


#### type Int64SocketOption

```go
type Int64SocketOption int
```


#### type IntSocketOption

```go
type IntSocketOption int
```


#### type MessageOption

```go
type MessageOption int
```


#### type PollEvents

```go
type PollEvents C.short
```


#### type PollItem

```go
type PollItem struct {
  Socket  *Socket         // socket to poll for events on
  Fd      ZmqOsSocketType // fd to poll for events on as returned from os.File.Fd()
  Events  PollEvents      // event set to poll for
  REvents PollEvents      // events that were present
}
```

Item to poll for read/write events on, either a *Socket or a file descriptor

#### type PollItems

```go
type PollItems []PollItem
```

a set of items to poll for events on

#### type SendRecvOption

```go
type SendRecvOption int
```


#### type Socket

```go
type Socket struct {
}
```


#### func (*Socket) Affinity

```go
func (s *Socket) Affinity() (uint64, error)
```
ZMQ_AFFINITY: Retrieve I/O thread affinity.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc7

#### func (*Socket) Backlog

```go
func (s *Socket) Backlog() (int, error)
```
ZMQ_BACKLOG: Retrieve maximum length of the queue of outstanding connections.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc18

#### func (*Socket) Bind

```go
func (s *Socket) Bind(address string) error
```
Bind the socket to a listening address. int zmq_bind (void *s, const char
*addr);

#### func (*Socket) Close

```go
func (s *Socket) Close() error
```
Shutdown the socket. int zmq_close (void *s);

#### func (*Socket) Connect

```go
func (s *Socket) Connect(address string) error
```
Connect the socket to an address. int zmq_connect (void *s, const char *addr);

#### func (*Socket) Events

```go
func (s *Socket) Events() (uint64, error)
```
ZMQ_EVENTS: Retrieve socket event state.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc20

#### func (*Socket) GetSockOptBool

```go
func (s *Socket) GetSockOptBool(option BoolSocketOption) (value bool, err error)
```

#### func (*Socket) GetSockOptInt

```go
func (s *Socket) GetSockOptInt(option IntSocketOption) (value int, err error)
```
Get an int option from the socket. int zmq_getsockopt (void *s, int option, void
*optval, size_t *optvallen);

#### func (*Socket) GetSockOptInt64

```go
func (s *Socket) GetSockOptInt64(option Int64SocketOption) (value int64, err error)
```
Get an int64 option from the socket. int zmq_getsockopt (void *s, int option,
void *optval, size_t *optvallen);

#### func (*Socket) GetSockOptString

```go
func (s *Socket) GetSockOptString(option StringSocketOption) (value string, err error)
```
Get a string option from the socket. int zmq_getsockopt (void *s, int option,
void *optval, size_t *optvallen);

#### func (*Socket) GetSockOptUInt64

```go
func (s *Socket) GetSockOptUInt64(option UInt64SocketOption) (value uint64, err error)
```
Get a uint64 option from the socket. int zmq_getsockopt (void *s, int option,
void *optval, size_t *optvallen);

#### func (*Socket) HWM

```go
func (s *Socket) HWM() (uint64, error)
```
ZMQ_HWM: Retrieve high water mark.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc5

#### func (*Socket) Identity

```go
func (s *Socket) Identity() (string, error)
```
ZMQ_IDENTITY: Retrieve socket identity.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc8

#### func (*Socket) Linger

```go
func (s *Socket) Linger() (time.Duration, error)
```
ZMQ_LINGER: Retrieve linger period for socket shutdown.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc15

#### func (*Socket) McastLoop

```go
func (s *Socket) McastLoop() (bool, error)
```
ZMQ_MCAST_LOOP: Control multicast loop-back.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc12

#### func (*Socket) Rate

```go
func (s *Socket) Rate() (int64, error)
```
ZMQ_RATE: Retrieve multicast data rate.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc9

#### func (*Socket) RcvBuf

```go
func (s *Socket) RcvBuf() (uint64, error)
```
ZMQ_RCVBUF: Retrieve kernel receive buffer size.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc14

#### func (*Socket) RcvHWM

```go
func (s *Socket) RcvHWM() (int, error)
```
ZMQ_RCVHWM: Retrieve high water mark for inbound messages.

See: http://api.zeromq.org/3.2:zmq-getsockopt#toc6

#### func (*Socket) RcvMore

```go
func (s *Socket) RcvMore() (bool, error)
```
ZMQ_RCVMORE: More message parts to follow.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc4

#### func (*Socket) RcvTimeout

```go
func (s *Socket) RcvTimeout() (time.Duration, error)
```
ZMQ_RCVTIMEO: Maximum time before a socket operation returns with EAGAIN.

See: http://api.zeromq.org/2.2:zmq-getsockopt#toc6

#### func (*Socket) ReconnectIvl

```go
func (s *Socket) ReconnectIvl() (time.Duration, error)
```
ZMQ_RECONNECT_IVL: Retrieve reconnection interval.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc16

#### func (*Socket) ReconnectIvlMax

```go
func (s *Socket) ReconnectIvlMax() (time.Duration, error)
```
ZMQ_RECONNECT_IVL_MAX: Retrieve maximum reconnection interval.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc17

#### func (*Socket) RecoveryIvl

```go
func (s *Socket) RecoveryIvl() (time.Duration, error)
```
ZMQ_RECOVERY_IVL_MSEC: Get multicast recovery interval in milliseconds.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc11

#### func (*Socket) Recv

```go
func (s *Socket) Recv(flags SendRecvOption) (data []byte, err error)
```
Receive a message from the socket. int zmq_recv (void *s, zmq_msg_t *msg, int
flags);

#### func (*Socket) RecvMultipart

```go
func (s *Socket) RecvMultipart(flags SendRecvOption) (parts [][]byte, err error)
```
Receive a multipart message.

#### func (*Socket) Send

```go
func (s *Socket) Send(data []byte, flags SendRecvOption) error
```
Send a message to the socket. int zmq_send (void *s, zmq_msg_t *msg, int flags);

#### func (*Socket) SendMultipart

```go
func (s *Socket) SendMultipart(parts [][]byte, flags SendRecvOption) (err error)
```
Send a multipart message.

#### func (*Socket) SetAffinity

```go
func (s *Socket) SetAffinity(value uint64) error
```
ZMQ_AFFINITY: Set I/O thread affinity.

See: http://api.zeromq.org/2.1:zmq-setsockopt#toc5

#### func (*Socket) SetBacklog

```go
func (s *Socket) SetBacklog(value int) error
```
ZMQ_BACKLOG: Set maximum length of the queue of outstanding connections.

See: http://api.zeromq.org/2.1:zmq-setsockopt#toc18

#### func (*Socket) SetHWM

```go
func (s *Socket) SetHWM(value uint64) error
```
ZMQ_HWM: Set high water mark.

See: http://api.zeromq.org/2.1:zmq-setsockopt#toc3

#### func (*Socket) SetIdentity

```go
func (s *Socket) SetIdentity(value string) error
```
ZMQ_IDENTITY: Set socket identity.

See: http://api.zeromq.org/2.1:zmq-setsockopt#toc6

#### func (*Socket) SetLinger

```go
func (s *Socket) SetLinger(value time.Duration) error
```
ZMQ_LINGER: Set linger period for socket shutdown.

See: http://api.zeromq.org/2.1:zmq-setsockopt#toc15

#### func (*Socket) SetMcastLoop

```go
func (s *Socket) SetMcastLoop(value bool) error
```
ZMQ_MCAST_LOOP: Control multicast loop-back.

See: http://api.zeromq.org/2.1:zmq-setsockopt#toc12

#### func (*Socket) SetRate

```go
func (s *Socket) SetRate(value int64) error
```
ZMQ_RATE: Set multicast data rate.

See: http://api.zeromq.org/2.1:zmq-setsockopt#toc9

#### func (*Socket) SetRcvBuf

```go
func (s *Socket) SetRcvBuf(value uint64) error
```
ZMQ_RCVBUF: Set kernel receive buffer size.

See: http://api.zeromq.org/2.1:zmq-setsockopt#toc14

#### func (*Socket) SetRcvHWM

```go
func (s *Socket) SetRcvHWM(value int) error
```
ZMQ_RCVHWM: Set high water mark for inbound messages.

See: http://api.zeromq.org/3.2:zmq-setsockopt#toc4

#### func (*Socket) SetRcvTimeout

```go
func (s *Socket) SetRcvTimeout(value time.Duration) error
```
ZMQ_RCVTIMEO: Maximum time before a recv operation returns with EAGAIN.

See: http://api.zeromq.org/2.2:zmq-setsockopt#toc9

#### func (*Socket) SetReconnectIvl

```go
func (s *Socket) SetReconnectIvl(value time.Duration) error
```
ZMQ_RECONNECT_IVL: Set reconnection interval.

See: http://api.zeromq.org/2.1:zmq-setsockopt#toc16

#### func (*Socket) SetReconnectIvlMax

```go
func (s *Socket) SetReconnectIvlMax(value time.Duration) error
```
ZMQ_RECONNECT_IVL_MAX: Set maximum reconnection interval.

See: http://api.zeromq.org/2.1:zmq-setsockopt#toc17

#### func (*Socket) SetRecoveryIvl

```go
func (s *Socket) SetRecoveryIvl(value time.Duration) error
```
ZMQ_RECOVERY_IVL_MSEC: Set multicast recovery interval in milliseconds.

See: http://api.zeromq.org/2.1:zmq-setsockopt#toc11

#### func (*Socket) SetSndBuf

```go
func (s *Socket) SetSndBuf(value uint64) error
```
ZMQ_SNDBUF: Set kernel transmit buffer size.

See: http://api.zeromq.org/2.1:zmq-setsockopt#toc13

#### func (*Socket) SetSndHWM

```go
func (s *Socket) SetSndHWM(value int) error
```
ZMQ_SNDHWM: Set high water mark for outbound messages.

See: http://api.zeromq.org/3.2:zmq-setsockopt#toc3

#### func (*Socket) SetSndTimeout

```go
func (s *Socket) SetSndTimeout(value time.Duration) error
```
ZMQ_SNDTIMEO: Maximum time before a send operation returns with EAGAIN.

See: http://api.zeromq.org/2.2:zmq-setsockopt#toc10

#### func (*Socket) SetSockOptInt

```go
func (s *Socket) SetSockOptInt(option IntSocketOption, value int) error
```
Set an int option on the socket. int zmq_setsockopt (void *s, int option, const
void *optval, size_t optvallen);

#### func (*Socket) SetSockOptInt64

```go
func (s *Socket) SetSockOptInt64(option Int64SocketOption, value int64) error
```
Set an int64 option on the socket. int zmq_setsockopt (void *s, int option,
const void *optval, size_t optvallen);

#### func (*Socket) SetSockOptString

```go
func (s *Socket) SetSockOptString(option StringSocketOption, value string) error
```
Set a string option on the socket. int zmq_setsockopt (void *s, int option,
const void *optval, size_t optvallen);

#### func (*Socket) SetSockOptStringNil

```go
func (s *Socket) SetSockOptStringNil(option StringSocketOption) error
```
Set a string option on the socket to nil. int zmq_setsockopt (void *s, int
option, const void *optval, size_t optvallen);

#### func (*Socket) SetSockOptUInt64

```go
func (s *Socket) SetSockOptUInt64(option UInt64SocketOption, value uint64) error
```
Set a uint64 option on the socket. int zmq_setsockopt (void *s, int option,
const void *optval, size_t optvallen);

#### func (*Socket) SetSubscribe

```go
func (s *Socket) SetSubscribe(value string) error
```
ZMQ_SUBSCRIBE: Establish message filter.

See: http://api.zeromq.org/2.1:zmq-setsockopt#toc7

#### func (*Socket) SetSwap

```go
func (s *Socket) SetSwap(value int64) error
```
ZMQ_SWAP: Set disk offload size.

See: http://api.zeromq.org/2.1:zmq-setsockopt#toc4

#### func (*Socket) SetTCPKeepalive

```go
func (s *Socket) SetTCPKeepalive(value int) error
```
ZMQ_TCP_KEEPALIVE: Override SO_KEEPALIVE socket option.

See: http://api.zeromq.org/3.2:zmq-setsockopt#toc25

#### func (*Socket) SetTCPKeepaliveCnt

```go
func (s *Socket) SetTCPKeepaliveCnt(value int) error
```
ZMQ_TCP_KEEPALIVE_CNT: Override TCP_KEEPCNT socket option.

See: http://api.zeromq.org/3.2:zmq-setsockopt#toc27

#### func (*Socket) SetTCPKeepaliveIdle

```go
func (s *Socket) SetTCPKeepaliveIdle(value int) error
```
ZMQ_TCP_KEEPALIVE_IDLE: Override TCP_KEEPCNT(or TCP_KEEPALIVE on some OS).

See: http://api.zeromq.org/3.2:zmq-setsockopt#toc26

#### func (*Socket) SetTCPKeepaliveIntvl

```go
func (s *Socket) SetTCPKeepaliveIntvl(value int) error
```
ZMQ_TCP_KEEPALIVE_INTVL: Override TCP_KEEPINTVL socket option.

See: http://api.zeromq.org/3.2:zmq-setsockopt#toc28

#### func (*Socket) SetUnsubscribe

```go
func (s *Socket) SetUnsubscribe(value string) error
```
ZMQ_UNSUBSCRIBE: Remove message filter.

See: http://api.zeromq.org/2.1:zmq-setsockopt#toc8

#### func (*Socket) SndBuf

```go
func (s *Socket) SndBuf() (uint64, error)
```
ZMQ_SNDBUF: Retrieve kernel transmit buffer size.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc13

#### func (*Socket) SndHWM

```go
func (s *Socket) SndHWM() (int, error)
```
ZMQ_SNDHWM: Retrieves high water mark for outbound messages.

See: http://api.zeromq.org/3.2:zmq-getsockopt#toc5

#### func (*Socket) SndTimeout

```go
func (s *Socket) SndTimeout() (time.Duration, error)
```
ZMQ_SNDTIMEO: Maximum time before a socket operation returns with EAGAIN.

See: http://api.zeromq.org/2.2:zmq-getsockopt#toc7

#### func (*Socket) Swap

```go
func (s *Socket) Swap() (int64, error)
```
ZMQ_SWAP: Retrieve disk offload size.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc6

#### func (*Socket) TCPKeepalive

```go
func (s *Socket) TCPKeepalive() (int, error)
```
ZMQ_TCP_KEEPALIVE: Override SO_KEEPALIVE socket option.

See: http://api.zeromq.org/3.2:zmq-getsockopt#toc26

#### func (*Socket) TCPKeepaliveCnt

```go
func (s *Socket) TCPKeepaliveCnt() (int, error)
```
ZMQ_TCP_KEEPALIVE_CNT: Override TCP_KEEPCNT socket option.

See: http://api.zeromq.org/3.2:zmq-getsockopt#toc28

#### func (*Socket) TCPKeepaliveIdle

```go
func (s *Socket) TCPKeepaliveIdle() (int, error)
```
ZMQ_TCP_KEEPALIVE_IDLE: Override TCP_KEEPCNT(or TCP_KEEPALIVE on some OS).

See: http://api.zeromq.org/3.2:zmq-getsockopt#toc27

#### func (*Socket) TCPKeepaliveIntvl

```go
func (s *Socket) TCPKeepaliveIntvl() (int, error)
```
ZMQ_TCP_KEEPALIVE_INTVL: Override TCP_KEEPINTVL socket option.

See: http://api.zeromq.org/3.2:zmq-getsockopt#toc29

#### func (*Socket) Type

```go
func (s *Socket) Type() (SocketType, error)
```
ZMQ_TYPE: Retrieve socket type.

See: http://api.zeromq.org/2.1:zmq-getsockopt#toc3

#### type SocketType

```go
type SocketType int
```


#### type StringSocketOption

```go
type StringSocketOption int
```


#### type UInt64SocketOption

```go
type UInt64SocketOption int
```


#### type ZmqOsSocketType

```go
type ZmqOsSocketType C.SOCKET
```


#### func (ZmqOsSocketType) ToRaw

```go
func (self ZmqOsSocketType) ToRaw() C.SOCKET
```

*(generated from .[godocdown](https://github.com/robertkrimen/godocdown).md with `godocdown github.com/alecthomas/gozmq > README.md`)*
