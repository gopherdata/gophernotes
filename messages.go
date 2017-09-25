package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/nu7hatch/gouuid"
	zmq "github.com/pebbe/zmq4"
)

// MsgHeader encodes header info for ZMQ messages.
type MsgHeader struct {
	MsgID           string `json:"msg_id"`
	Username        string `json:"username"`
	Session         string `json:"session"`
	MsgType         string `json:"msg_type"`
	ProtocolVersion string `json:"version"`
	Timestamp       string `json:"date"`
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

// bundledMIMEData holds data that can be presented in multiple formats. The keys are MIME types
// and the values are the data formatted with respect to it's MIME type. All bundles should contain
// at least a "text/plain" representation with a string value.
type bundledMIMEData map[string]interface{}

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
	msg.Header.ProtocolVersion = ProtocolVersion
	msg.Header.Timestamp = time.Now().UTC().Format(time.RFC3339)

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

// newTextMIMEDataBundle creates a bundledMIMEData that only contains a text representation described
// by the value parameter.
func newTextBundledMIMEData(value string) bundledMIMEData {
	return bundledMIMEData{
		"text/plain": value,
	}
}

// PublishKernelStatus publishes a status message notifying front-ends of the state the kernel is in. Supports
// states "starting", "busy", and "idle".
func (receipt *msgReceipt) PublishKernelStatus(status string) error {
	return receipt.Publish("status",
		struct {
			ExecutionState string `json:"execution_state"`
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
			ExecCount int             `json:"execution_count"`
			Data      bundledMIMEData `json:"data"`
			Metadata  bundledMIMEData `json:"metadata"`
		}{
			ExecCount: execCount,
			Data:      newTextBundledMIMEData(output),
			Metadata:  make(bundledMIMEData),
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

const (
	// StreamStdout defines the stream name for standard out on the front-end. It
	// is used in `PublishWriteStream` to specify the stream to write to.
	StreamStdout = "stdout"

	// StreamStderr defines the stream name for standard error on the front-end. It
	// is used in `PublishWriteStream` to specify the stream to write to.
	StreamStderr = "stderr"
)

// PublishWriteStream prints the data string to a stream on the front-end. This is
// either `StreamStdout` or `StreamStderr`.
func (receipt *msgReceipt) PublishWriteStream(stream string, data string) error {
	return receipt.Publish("stream",
		struct {
			Stream string `json:"name"`
			Data   string `json:"text"`
		}{
			Stream: stream,
			Data:   data,
		},
	)
}

// JupyterStreamWriter is an `io.Writer` implementation that writes the data to the notebook
// front-end.
type JupyterStreamWriter struct {
	stream  string
	receipt *msgReceipt
}

// Write implements `io.Writer.Write` by publishing the data via `PublishWriteStream`
func (writer *JupyterStreamWriter) Write(p []byte) (n int, err error) {
	data := string(p)
	n = len(p)

	err = writer.receipt.PublishWriteStream(writer.stream, data)
	if err != nil {
		n = 0
	}

	return
}
