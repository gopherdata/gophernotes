package kvsimple

import (
	zmq "github.com/pebbe/zmq4"

	"os"
	"testing"
)

func TestKvmsg(t *testing.T) {

	//  Prepare our context and sockets
	output, err := zmq.NewSocket(zmq.DEALER)
	if err != nil {
		t.Error(err)
	}

	err = output.Bind("ipc://kvmsg_selftest.ipc")
	if err != nil {
		t.Error(err)
	}

	input, err := zmq.NewSocket(zmq.DEALER)
	if err != nil {
		t.Error(err)
	}

	err = input.Connect("ipc://kvmsg_selftest.ipc")
	if err != nil {
		t.Error(err)
	}

	kvmap := make(map[string]*Kvmsg)

	//  Test send and receive of simple message
	kvmsg := NewKvmsg(1)
	kvmsg.SetKey("key")
	kvmsg.SetBody("body")
	kvmsg.Dump()
	err = kvmsg.Send(output)

	kvmsg.Store(kvmap)
	if err != nil {
		t.Error(err)
	}

	kvmsg, err = RecvKvmsg(input)
	if err != nil {
		t.Error(err)
	}
	kvmsg.Dump()
	key, err := kvmsg.GetKey()
	if err != nil {
		t.Error(err)
	}
	if key != "key" {
		t.Error("Expected \"key\", got \"" + key + "\"")
	}
	kvmsg.Store(kvmap)

	input.Close()
	output.Close()
	os.Remove("kvmsg_selftest.ipc")
}
