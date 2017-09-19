package zmq4_test

import (
	zmq "github.com/pebbe/zmq4/draft"

	"testing"
	"time"
)

func TestUdp(t *testing.T) {

	if _, minor, _ := zmq.Version(); minor < 2 {
		t.Skip("Sockets RADIO and DISH need ZeroMQ 4.2 with draft enabled")
	}

	ctx, err := zmq.NewContext()
	if err != nil {
		t.Fatal("NewContext:", err)
	}
	defer ctx.Term()

	radio, err := ctx.NewSocket(zmq.RADIO)
	if err != nil {
		t.Fatal("NewSocket RADIO:", err)
	}
	defer radio.Close()
	dish, err := ctx.NewSocket(zmq.DISH)
	if err != nil {
		t.Fatal("NewSocket DISH:", err)
	}
	defer dish.Close()

	//  Connecting dish should fail
	err = dish.Connect("udp://127.0.0.1:5556")
	if err == nil {
		t.Fatal("Expected fail on dish.Connect")
	}

	err = dish.Bind("udp://*:5556")
	if err != nil {
		t.Fatal("dish.Bind:", err)
	}

	//  Bind radio should fail
	err = radio.Bind("udp://*:5556")
	if err == nil {
		t.Fatal("Expected fail on radio.Bind")
	}

	err = radio.Connect("udp://127.0.0.1:5556")
	if err != nil {
		t.Fatal("radio.Connect:", err)
	}

	time.Sleep(300 * time.Millisecond)

	err = dish.Join("TV")
	if err != nil {
		t.Fatal("dish.Join:", err)
	}

	_, err = radio.Send("Friends", 0, zmq.OptGroup("TV"))
	if err != nil {
		t.Fatal("radio.SendMessage:", err)
	}

	msg, opt, err := dish.RecvWithOpts(0, zmq.OptGroup(""))
	if err != nil {
		t.Fatal("dish.RecvWithOpt:", err)
	}
	if len(opt) != 1 {
		t.Fatal("dish.RecvWithOpt: wrong number off options")
	}
	if string(opt[0].(zmq.OptGroup)) != "TV" {
		t.Fatal("dish.RecvWithOpt: wrong group: %v", string(opt[0].(zmq.OptGroup)))
	}
	if msg != "Friends" {
		t.Fatal("dish.RecvWithOpt: wrong message: %q", msg)
	}
}

func TestClientServer(t *testing.T) {

	if _, minor, _ := zmq.Version(); minor < 2 {
		t.Skip("Sockets CLIENT and SERVER need ZeroMQ 4.2 with draft enabled")
	}

	ctx, err := zmq.NewContext()
	if err != nil {
		t.Fatal("NewContext:", err)
	}
	defer ctx.Term()

	server, err := ctx.NewSocket(zmq.SERVER)
	if err != nil {
		t.Fatal("NewSocket SERVER:", err)
	}
	defer server.Close()
	client, err := ctx.NewSocket(zmq.CLIENT)
	if err != nil {
		t.Fatal("NewSocket CLIENT:", err)
	}
	defer client.Close()

	addr := "tcp://127.0.0.1:9797"
	err = server.Bind(addr)
	if err != nil {
		t.Fatal("server.Bind:", err)
	}
	err = client.Connect(addr)
	if err != nil {
		t.Fatal("client.Connect:", err)
	}

	content := "12345678ABCDEFGH12345678abcdefgh"
	rc, err := client.Send(content, zmq.DONTWAIT)
	if err != nil {
		t.Fatal("client.Send DONTWAIT: ", err)
	}
	if rc != 32 {
		t.Fatal("client.Send DONTWAIT: ", err32)
	}

	msg, opts, err := server.RecvWithOpts(0, zmq.OptRoutingId(0))
	if err != nil {
		t.Fatal("server.Recv: ", err)
	}
	//  Check that message is still the same
	if msg != content {
		t.Fatal("server.Recv: %q != %q", msg, content)
	}

	rc, err = server.Send(content, 0, opts[0])
	if err != nil {
		t.Fatal("server.Send:", err)
	}
	if rc != 32 {
		t.Fatal("server.Send: ", err32)
	}

	//  Receive message at client side
	msg, err = client.Recv(0)
	if err != nil {
		t.Fatal("client.Recv: ", err)
	}

	//  Check that message is still the same
	if msg != content {
		t.Fatalf("client.Recv: %q != %q", msg, content)
	}

}
