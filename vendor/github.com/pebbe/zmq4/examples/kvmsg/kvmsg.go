//  kvmsg class - key-value message class for example applications
package kvmsg

import (
	zmq "github.com/pebbe/zmq4"

	"github.com/pborman/uuid"

	"errors"
	"fmt"
	"os"
	"strings"
)

//  Message is formatted on wire as 4 frames:
//  frame 0: key (0MQ string)
//  frame 1: sequence (8 bytes, network order)
//  frame 2: uuid (blob, 16 bytes)
//  frame 3: properties (0MQ string)
//  frame 4: body (blob)
const (
	frame_KEY    = 0
	frame_SEQ    = 1
	frame_UUID   = 2
	frame_PROPS  = 3
	frame_BODY   = 4
	kvmsg_FRAMES = 5
)

//  Structure of our class
type Kvmsg struct {
	//  Presence indicators for each frame
	present []bool
	//  Corresponding 0MQ message frames, if any
	frame []string
	//  List of properties, as name=value strings
	props []string
}

//  These two helpers serialize a list of properties to and from a
//  message frame:

func (kvmsg *Kvmsg) encode_props() {
	kvmsg.frame[frame_PROPS] = strings.Join(kvmsg.props, "\n") + "\n"
	kvmsg.present[frame_PROPS] = true
}

func (kvmsg *Kvmsg) decode_props() {
	kvmsg.props = strings.Split(kvmsg.frame[frame_PROPS], "\n")
	if ln := len(kvmsg.props); ln > 0 && kvmsg.props[ln-1] == "" {
		kvmsg.props = kvmsg.props[:ln-1]
	}
}

//  Constructor, takes a sequence number for the new Kvmsg instance.
func NewKvmsg(sequence int64) (kvmsg *Kvmsg) {
	kvmsg = &Kvmsg{
		present: make([]bool, kvmsg_FRAMES),
		frame:   make([]string, kvmsg_FRAMES),
		props:   make([]string, 0),
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
	kvmsg.decode_props()
	return
}

//  Send key-value message to socket; any empty frames are sent as such.
func (kvmsg *Kvmsg) Send(socket *zmq.Socket) (err error) {
	//fmt.Printf("Send to %s: %q\n", socket, kvmsg.frame)
	kvmsg.encode_props()
	_, err = socket.SendMessage(kvmsg.frame)
	return
}

//  The Dup method duplicates a kvmsg instance, returns the new instance.
func (kvmsg *Kvmsg) Dup() (dup *Kvmsg) {
	dup = &Kvmsg{
		present: make([]bool, kvmsg_FRAMES),
		frame:   make([]string, kvmsg_FRAMES),
		props:   make([]string, len(kvmsg.props)),
	}
	copy(dup.present, kvmsg.present)
	copy(dup.frame, kvmsg.frame)
	copy(dup.props, kvmsg.props)
	return
}

//  Return key from last read message, if any, else NULL
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

func (kvmsg *Kvmsg) GetUuid() (uuid string, err error) {
	if !kvmsg.present[frame_UUID] {
		err = errors.New("Uuid not set")
		return
	}
	uuid = kvmsg.frame[frame_UUID]
	return
}

//  Sets the UUID to a random generated value
func (kvmsg *Kvmsg) SetUuid() {
	kvmsg.frame[frame_UUID] = string(uuid.NewRandom()) // raw 16 bytes
	kvmsg.present[frame_UUID] = true

}

// Get message property, return error if no such property is defined.
func (kvmsg *Kvmsg) GetProp(name string) (value string, err error) {
	if !kvmsg.present[frame_PROPS] {
		err = errors.New("No properties set")
		return
	}
	f := name + "="
	for _, prop := range kvmsg.props {
		if strings.HasPrefix(prop, f) {
			value = prop[len(f):]
			return
		}
	}
	err = errors.New("Property not set")
	return
}

//  Set message property. Property name cannot contain '='.
func (kvmsg *Kvmsg) SetProp(name, value string) (err error) {
	if strings.Index(name, "=") >= 0 {
		err = errors.New("No '=' allowed in property name")
		return
	}
	p := name + "="
	for i, prop := range kvmsg.props {
		if strings.HasPrefix(prop, p) {
			kvmsg.props = append(kvmsg.props[:i], kvmsg.props[i+1:]...)
			break
		}
	}
	kvmsg.props = append(kvmsg.props, name+"="+value)
	kvmsg.present[frame_PROPS] = true
	return
}

//  The store method stores the key-value message into a hash map, unless
//  the key is nil.
func (kvmsg *Kvmsg) Store(kvmap map[string]*Kvmsg) {
	if kvmsg.present[frame_KEY] {
		if kvmsg.present[frame_BODY] && kvmsg.frame[frame_BODY] != "" {
			kvmap[kvmsg.frame[frame_KEY]] = kvmsg
		} else {
			delete(kvmap, kvmsg.frame[frame_KEY])
		}
	}
}

//  The dump method extends the kvsimple implementation with support for
//  message properties.
func (kvmsg *Kvmsg) Dump() {
	size := kvmsg.Size()
	body, _ := kvmsg.GetBody()
	seq, _ := kvmsg.GetSequence()
	key, _ := kvmsg.GetKey()
	fmt.Fprintf(os.Stderr, "[seq:%v][key:%v][size:%v] ", seq, key, size)
	p := "["
	for _, prop := range kvmsg.props {
		fmt.Fprint(os.Stderr, p, prop)
		p = ";"
	}
	if p == ";" {
		fmt.Fprint(os.Stderr, "]")
	}
	for char_nbr := 0; char_nbr < size; char_nbr++ {
		fmt.Fprintf(os.Stderr, "%02X", body[char_nbr])
	}
	fmt.Fprintln(os.Stderr)
}

// The test function is in kvmsg_test.go
