package kvmsg

import (
	zmq "github.com/pebbe/zmq4"

	"os"
	"testing"
)

//  The test is the same as in kvsimple with added support
//  for the uuid and property features of kvmsg

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
	kvmsg.SetUuid()
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

	//  Test send and receive of message with properties
	kvmsg = NewKvmsg(2)
	err = kvmsg.SetProp("prop1", "value1")
	if err != nil {
		t.Error(err)
	}
	kvmsg.SetProp("prop2", "value1")
	kvmsg.SetProp("prop2", "value2")
	kvmsg.SetKey("key")
	kvmsg.SetUuid()
	kvmsg.SetBody("body")
	if val, err := kvmsg.GetProp("prop2"); err != nil || val != "value2" {
		if err != nil {
			t.Error(err)
		}
		t.Error("Expected \"prop2\" = \"value2\", got \"" + val + "\"")
	}
	kvmsg.Dump()
	err = kvmsg.Send(output)

	kvmsg, err = RecvKvmsg(input)
	if err != nil {
		t.Error(err)
	}
	kvmsg.Dump()
	key, err = kvmsg.GetKey()
	if err != nil {
		t.Error(err)
	}
	if key != "key" {
		t.Error("Expected \"key\", got \"" + key + "\"")
	}
	prop, err := kvmsg.GetProp("prop2")
	if err != nil {
		t.Error(err)
	}
	if prop != "value2" {
		t.Error("Expected property \"value2\", got \"" + key + "\"")
	}

	input.Close()
	output.Close()
	os.Remove("kvmsg_selftest.ipc")
}
