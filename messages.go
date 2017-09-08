package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/nu7hatch/gouuid"
	zmq "github.com/pebbe/zmq4"
)

// MsgHeader encodes header info for ZMQ messages.
type MsgHeader struct {
	MsgID    string `json:"msg_id"`
	Username string `json:"username"`
	Session  string `json:"session"`
	MsgType  string `json:"msg_type"`
}

// ComposedMsg represents an entire message in a high-level structure.
type ComposedMsg struct {
	Header       MsgHeader
	ParentHeader MsgHeader
	Metadata     map[string]interface{}
	Content      interface{}
}

// msgReceipt represents a received message, its return identities, and
// the sockets for communication.
type msgReceipt struct {
	Msg        ComposedMsg
	Identities [][]byte
	Sockets    SocketGroup
}

// OutputMsg holds the data for a pyout message.
type OutputMsg struct {
	Execcount int                    `json:"execution_count"`
	Data      map[string]string      `json:"data"`
	Metadata  map[string]interface{} `json:"metadata"`
}

// ErrMsg encodes the traceback of errors output to the notebook.
type ErrMsg struct {
	EName     string   `json:"ename"`
	EValue    string   `json:"evalue"`
	Traceback []string `json:"traceback"`
}

// InvalidSignatureError is returned when the signature on a received message does not
// validate.
type InvalidSignatureError struct{}

func (e *InvalidSignatureError) Error() string {
	return "A message had an invalid signature"
}

// WireMsgToComposedMsg translates a multipart ZMQ messages received from a socket into
// a ComposedMsg struct and a slice of return identities. This includes verifying the
// message signature.
func WireMsgToComposedMsg(msgparts [][]byte, signkey []byte) (ComposedMsg, [][]byte, error) {

	i := 0
	for string(msgparts[i]) != "<IDS|MSG>" {
		i++
	}
	identities := msgparts[:i]

	// Validate signature.
	var msg ComposedMsg
	if len(signkey) != 0 {
		mac := hmac.New(sha256.New, signkey)
		for _, msgpart := range msgparts[i+2 : i+6] {
			mac.Write(msgpart)
		}
		signature := make([]byte, hex.DecodedLen(len(msgparts[i+1])))
		hex.Decode(signature, msgparts[i+1])
		if !hmac.Equal(mac.Sum(nil), signature) {
			return msg, nil, &InvalidSignatureError{}
		}
	}

	// Unmarshal contents.
	json.Unmarshal(msgparts[i+2], &msg.Header)
	json.Unmarshal(msgparts[i+3], &msg.ParentHeader)
	json.Unmarshal(msgparts[i+4], &msg.Metadata)
	json.Unmarshal(msgparts[i+5], &msg.Content)
	return msg, identities, nil
}

// ToWireMsg translates a ComposedMsg into a multipart ZMQ message ready to send, and
// signs it. This does not add the return identities or the delimiter.
func (msg ComposedMsg) ToWireMsg(signkey []byte) ([][]byte, error) {

	msgparts := make([][]byte, 5)

	header, err := json.Marshal(msg.Header)
	if err != nil {
		return msgparts, err
	}
	msgparts[1] = header

	parentHeader, err := json.Marshal(msg.ParentHeader)
	if err != nil {
		return msgparts, err
	}
	msgparts[2] = parentHeader

	if msg.Metadata == nil {
		msg.Metadata = make(map[string]interface{})
	}

	metadata, err := json.Marshal(msg.Metadata)
	if err != nil {
		return msgparts, err
	}
	msgparts[3] = metadata

	content, err := json.Marshal(msg.Content)
	if err != nil {
		return msgparts, err
	}
	msgparts[4] = content

	// Sign the message.
	if len(signkey) != 0 {
		mac := hmac.New(sha256.New, signkey)
		for _, msgpart := range msgparts[1:] {
			mac.Write(msgpart)
		}
		msgparts[0] = make([]byte, hex.EncodedLen(mac.Size()))
		hex.Encode(msgparts[0], mac.Sum(nil))
	}

	return msgparts, nil
}

// SendResponse sends a message back to return identites of the received message.
func (receipt *msgReceipt) SendResponse(socket *zmq.Socket, msg ComposedMsg) error {

	for _, idt := range receipt.Identities {
		_, err := socket.Send(string(idt), zmq.SNDMORE)
		if err != nil {
			return err
		}
	}

	_, err := socket.Send("<IDS|MSG>", zmq.SNDMORE)
	if err != nil {
		return err
	}

	msgParts, err := msg.ToWireMsg(receipt.Sockets.Key)
	if err != nil {
		return err
	}

	_, err = socket.SendMessage(msgParts)
	if err != nil {
		return err
	}

	return nil
}

// NewMsg creates a new ComposedMsg to respond to a parent message.
// This includes setting up its headers.
func NewMsg(msgType string, parent ComposedMsg) (ComposedMsg, error) {
	var msg ComposedMsg

	msg.ParentHeader = parent.Header
	msg.Header.Session = parent.Header.Session
	msg.Header.Username = parent.Header.Username
	msg.Header.MsgType = msgType

	u, err := uuid.NewV4()
	if err != nil {
		return msg, err
	}
	msg.Header.MsgID = u.String()

	return msg, nil
}

// Publish creates a new ComposedMsg and sends it back to the return identities over the
// IOPub channel.
func (receipt *msgReceipt) Publish(msgType string, content interface{}) error {
	msg, err := NewMsg(msgType, receipt.Msg)

	if err != nil {
		return err
	}

	msg.Content = content
	return receipt.SendResponse(receipt.Sockets.IOPubSocket, msg)
}

// Reply creates a new ComposedMsg and sends it back to the return identities over the
// Shell channel.
func (receipt *msgReceipt) Reply(msgType string, content interface{}) error {
	msg, err := NewMsg(msgType, receipt.Msg)

	if err != nil {
		return err
	}

	msg.Content = content
	return receipt.SendResponse(receipt.Sockets.ShellSocket, msg)
}

// MIMEDataBundle holds data that can be presented in multiple formats. The keys are MIME types
// and the values are the data formatted with respect to it's MIME type. All bundle should contain
// at least a "text/plain" representation with a string value.
type MIMEDataBundle map[string]interface{}

// NewTextMIMEDataBundle creates a MIMEDataBundle that only contains a text representation described
// the the parameter 'value'.
func NewTextMIMEDataBundle(value string) MIMEDataBundle {
	return MIMEDataBundle{
		"text/plain": value,
	}
}

type KernelStatus string

const (
	KernelStarting KernelStatus = "starting"
	KernelBusy                  = "busy"
	KernelIdle                  = "idle"
)

// PublishKernelStatus publishes a status message notifying front-ends of the state the kernel is in.
func (receipt *msgReceipt) PublishKernelStatus(status KernelStatus) error {
	return receipt.Publish("status",
		struct {
			ExecutionState KernelStatus `json:"execution_state"`
		}{
			ExecutionState: status,
		},
	)
}

// PublishExecutionInput publishes a status message notifying front-ends of what code is
// currently being executed.
func (receipt *msgReceipt) PublishExecutionInput(execCount int, code string) error {
	return receipt.Publish("execute_input",
		struct {
			ExecCount int    `json:"execution_count"`
			Code      string `json:"code"`
		}{
			ExecCount: execCount,
			Code:      code,
		},
	)
}

// PublishExecuteResult publishes the result of the `execCount` execution as a string.
func (receipt *msgReceipt) PublishExecutionResult(execCount int, output string) error {
	return receipt.Publish("execute_result",
		struct {
			ExecCount int            `json:"execution_count"`
			Data      MIMEDataBundle `json:"data"`
			Metadata  MIMEDataBundle `json:"metadata"`
		}{
			ExecCount: execCount,
			Data:      NewTextMIMEDataBundle(output),
			Metadata:  make(MIMEDataBundle),
		},
	)
}

// PublishExecuteResult publishes a serialized error that was encountered during execution.
func (receipt *msgReceipt) PublishExecutionError(err string, trace []string) error {
	return receipt.Publish("error",
		struct {
			Name  string   `json:"ename"`
			Value string   `json:"evalue"`
			Trace []string `json:"traceback"`
		}{
			Name:  "ERROR",
			Value: err,
			Trace: trace,
		},
	)
}

type Stream string

const (
	StreamStdout Stream = "stdout"
	StreamStderr        = "stderr"
)

// PublishWriteStream prints the data string to a stream on the front-end. This is
// either `StreamStdout` or `StreamStderr`.
func (receipt *msgReceipt) PublishWriteStream(stream Stream, data string) error {
	return receipt.Publish("stream",
		struct {
			Stream Stream `json:"name"`
			Data   string `json:"text"`
		}{
			Stream: stream,
			Data:   data,
		},
	)
}
