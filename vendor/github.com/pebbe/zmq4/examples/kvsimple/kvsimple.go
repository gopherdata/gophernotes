//  kvsimple - simple key-value message class for example applications.
//
//  This is a very much unlike typical Go.
package kvsimple

import (
	zmq "github.com/pebbe/zmq4"

	"errors"
	"fmt"
	"os"
)

const (
	frame_KEY    = 0
	frame_SEQ    = 1
	frame_BODY   = 2
	kvmsg_FRAMES = 3
)

//  The Kvmsg type holds a single key-value message consisting of a
//  list of 0 or more frames.
type Kvmsg struct {
	//  Presence indicators for each frame
	present []bool
	//  Corresponding 0MQ message frames, if any
	frame []string
}

//  Constructor, takes a sequence number for the new Kvmsg instance.
func NewKvmsg(sequence int64) (kvmsg *Kvmsg) {
	kvmsg = &Kvmsg{
		present: make([]bool, kvmsg_FRAMES),
		frame:   make([]string, kvmsg_FRAMES),
	}
	kvmsg.SetSequence(sequence)
	return
}

//  The RecvKvmsg function reads a key-value message from socket, and returns a new
//  Kvmsg instance.
func RecvKvmsg(socket *zmq.Socket) (kvmsg *Kvmsg, err error) {
	kvmsg = &Kvmsg{
		present: make([]bool, kvmsg_FRAMES),
		frame:   make([]string, kvmsg_FRAMES),
	}
	msg, err := socket.RecvMessage(0)
	if err != nil {
		return
	}
	//fmt.Printf("Recv from %s: %q\n", socket, msg)
	for i := 0; i < kvmsg_FRAMES && i < len(msg); i++ {
		kvmsg.frame[i] = msg[i]
		kvmsg.present[i] = true
	}
	return
}

//  The send method sends a multi-frame key-value message to a socket.
func (kvmsg *Kvmsg) Send(socket *zmq.Socket) (err error) {
	//fmt.Printf("Send to %s: %q\n", socket, kvmsg.frame)
	_, err = socket.SendMessage(kvmsg.frame)
	return
}

func (kvmsg *Kvmsg) GetKey() (key string, err error) {
	if !kvmsg.present[frame_KEY] {
		err = errors.New("Key not set")
		return
	}
	key = kvmsg.frame[frame_KEY]
	return
}

func (kvmsg *Kvmsg) SetKey(key string) {
	kvmsg.frame[frame_KEY] = key
	kvmsg.present[frame_KEY] = true
}

func (kvmsg *Kvmsg) GetSequence() (sequence int64, err error) {
	if !kvmsg.present[frame_SEQ] {
		err = errors.New("Sequence not set")
		return
	}
	source := kvmsg.frame[frame_SEQ]
	sequence = int64(source[0])<<56 +
		int64(source[1])<<48 +
		int64(source[2])<<40 +
		int64(source[3])<<32 +
		int64(source[4])<<24 +
		int64(source[5])<<16 +
		int64(source[6])<<8 +
		int64(source[7])
	return
}

func (kvmsg *Kvmsg) SetSequence(sequence int64) {

	source := make([]byte, 8)
	source[0] = byte((sequence >> 56) & 255)
	source[1] = byte((sequence >> 48) & 255)
	source[2] = byte((sequence >> 40) & 255)
	source[3] = byte((sequence >> 32) & 255)
	source[4] = byte((sequence >> 24) & 255)
	source[5] = byte((sequence >> 16) & 255)
	source[6] = byte((sequence >> 8) & 255)
	source[7] = byte((sequence) & 255)

	kvmsg.frame[frame_SEQ] = string(source)
	kvmsg.present[frame_SEQ] = true
}

func (kvmsg *Kvmsg) GetBody() (body string, err error) {
	if !kvmsg.present[frame_BODY] {
		err = errors.New("Body not set")
		return
	}
	body = kvmsg.frame[frame_BODY]
	return
}

func (kvmsg *Kvmsg) SetBody(body string) {
	kvmsg.frame[frame_BODY] = body
	kvmsg.present[frame_BODY] = true
}

//  The size method returns the body size of the last-read message, if any.
func (kvmsg *Kvmsg) Size() int {
	if kvmsg.present[frame_BODY] {
		return len(kvmsg.frame[frame_BODY])
	}
	return 0
}

//  The store method stores the key-value message into a hash map, unless
//  the key is nil.
func (kvmsg *Kvmsg) Store(kvmap map[string]*Kvmsg) {
	if kvmsg.present[frame_KEY] {
		kvmap[kvmsg.frame[frame_KEY]] = kvmsg
	}
}

//  The dump method prints the key-value message to stderr,
//  for debugging and tracing.
func (kvmsg *Kvmsg) Dump() {
	size := kvmsg.Size()
	body, _ := kvmsg.GetBody()
	seq, _ := kvmsg.GetSequence()
	key, _ := kvmsg.GetKey()
	fmt.Fprintf(os.Stderr, "[seq:%v][key:%v][size:%v]", seq, key, size)
	for char_nbr := 0; char_nbr < size; char_nbr++ {
		fmt.Fprintf(os.Stderr, "%02X", body[char_nbr])
	}
	fmt.Fprintln(os.Stderr)
}

// The test function is in kvsimple_test.go
