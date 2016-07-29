package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	zmq "github.com/alecthomas/gozmq"
)

var logger *log.Logger

// ConnectionInfo stores the contents of the kernel connection file created by Jupyter.
type ConnectionInfo struct {
	SignatureScheme string `json:"signature_scheme"`
	Transport       string `json:"transport"`
	StdinPort       int    `json:"stdin_port"`
	ControlPort     int    `json:"control_port"`
	IOPubPort       int    `json:"iopub_port"`
	HBPort          int    `json:"hb_port"`
	ShellPort       int    `json:"shell_port"`
	Key             string `json:"key"`
	IP              string `json:"ip"`
}

// SocketGroup holds the sockets needed to communicate with the kernel, and
// the key for message signing.
type SocketGroup struct {
	ShellSocket   *zmq.Socket
	ControlSocket *zmq.Socket
	StdinSocket   *zmq.Socket
	IOPubSocket   *zmq.Socket
	Key           []byte
}

// PrepareSockets sets up the ZMQ sockets through which the kernel will communicate.
func PrepareSockets(connInfo ConnectionInfo) (sg SocketGroup) {

	// TODO handle errors.
	context, _ := zmq.NewContext()
	sg.ShellSocket, _ = context.NewSocket(zmq.ROUTER)
	sg.ControlSocket, _ = context.NewSocket(zmq.ROUTER)
	sg.StdinSocket, _ = context.NewSocket(zmq.ROUTER)
	sg.IOPubSocket, _ = context.NewSocket(zmq.PUB)

	address := fmt.Sprintf("%v://%v:%%v", connInfo.Transport, connInfo.IP)

	sg.ShellSocket.Bind(fmt.Sprintf(address, connInfo.ShellPort))
	sg.ControlSocket.Bind(fmt.Sprintf(address, connInfo.ControlPort))
	sg.StdinSocket.Bind(fmt.Sprintf(address, connInfo.StdinPort))
	sg.IOPubSocket.Bind(fmt.Sprintf(address, connInfo.IOPubPort))

	// Message signing key
	sg.Key = []byte(connInfo.Key)

	// Start the heartbeat device
	HBSocket, _ := context.NewSocket(zmq.REP)
	HBSocket.Bind(fmt.Sprintf(address, connInfo.HBPort))
	go zmq.Device(zmq.FORWARDER, HBSocket, HBSocket)

	return
}

// HandleShellMsg responds to a message on the shell ROUTER socket.
func HandleShellMsg(receipt MsgReceipt) {
	switch receipt.Msg.Header.MsgType {
	case "kernel_info_request":
		SendKernelInfo(receipt)
	case "execute_request":
		HandleExecuteRequest(receipt)
	case "shutdown_request":
		HandleShutdownRequest(receipt)
	default:
		logger.Println("Unhandled shell message:", receipt.Msg.Header.MsgType)
	}
}

// KernelInfo holds information about the igo kernel, for kernel_info_reply messages.
type KernelInfo struct {
	ProtocolVersion []int  `json:"protocol_version"`
	Language        string `json:"language"`
}

// KernelStatus holds a kernel state, for status broadcast messages.
type KernelStatus struct {
	ExecutionState string `json:"execution_state"`
}

// SendKernelInfo sends a kernel_info_reply message.
func SendKernelInfo(receipt MsgReceipt) {
	reply := NewMsg("kernel_info_reply", receipt.Msg)
	reply.Content = KernelInfo{[]int{4, 0}, "go"}
	receipt.SendResponse(receipt.Sockets.ShellSocket, reply)
}

// ShutdownReply encodes a boolean indication of stutdown/restart
type ShutdownReply struct {
	Restart bool `json:"restart"`
}

// HandleShutdownRequest sends a "shutdown" message
func HandleShutdownRequest(receipt MsgReceipt) {
	reply := NewMsg("shutdown_reply", receipt.Msg)
	content := receipt.Msg.Content.(map[string]interface{})
	restart := content["restart"].(bool)
	reply.Content = ShutdownReply{restart}
	receipt.SendResponse(receipt.Sockets.ShellSocket, reply)
	logger.Println("Shutting down in response to shutdown_request")
	os.Exit(0)
}

// RunKernel is the main entry point to start the kernel. This is what is called by the
// gophernotes executable.
func RunKernel(connectionFile string, logwriter io.Writer) {

	logger = log.New(logwriter, "gophernotes ", log.LstdFlags)

	// set up the "Session" with the replpkg
	SetupExecutionEnvironment()

	var connInfo ConnectionInfo
	bs, err := ioutil.ReadFile(connectionFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(bs, &connInfo)
	if err != nil {
		log.Fatalln(err)
	}
	logger.Printf("%+v\n", connInfo)

	// Set up the ZMQ sockets through which the kernel will communicate
	sockets := PrepareSockets(connInfo)

	pi := zmq.PollItems{
		zmq.PollItem{Socket: sockets.ShellSocket, Events: zmq.POLLIN},
		zmq.PollItem{Socket: sockets.StdinSocket, Events: zmq.POLLIN},
		zmq.PollItem{Socket: sockets.ControlSocket, Events: zmq.POLLIN},
	}

	var msgparts [][]byte
	// Message receiving loop:
	for {
		_, err = zmq.Poll(pi, -1)
		if err != nil {
			log.Fatalln(err)
		}
		switch {
		case pi[0].REvents&zmq.POLLIN != 0: // shell socket
			msgparts, _ = pi[0].Socket.RecvMultipart(0)
			msg, ids, err := WireMsgToComposedMsg(msgparts, sockets.Key)
			if err != nil {
				fmt.Println(err)
				return
			}
			HandleShellMsg(MsgReceipt{msg, ids, sockets})
		case pi[1].REvents&zmq.POLLIN != 0: // stdin socket - not implemented.
			pi[1].Socket.RecvMultipart(0)
		case pi[2].REvents&zmq.POLLIN != 0: // control socket - treat like shell socket.
			msgparts, _ = pi[2].Socket.RecvMultipart(0)
			msg, ids, err := WireMsgToComposedMsg(msgparts, sockets.Key)
			if err != nil {
				fmt.Println(err)
				return
			}
			HandleShellMsg(MsgReceipt{msg, ids, sockets})
		}
	}
}
