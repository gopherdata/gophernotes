package main

import (
	"encoding/json"
	"fmt"
	zmq "github.com/alecthomas/gozmq"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var logger *log.Logger

// ConnectionInfo stores the contents of the kernel connection file created by Jupyter.
type ConnectionInfo struct {
	Signature_scheme string
	Transport        string
	Stdin_port       int
	Control_port     int
	IOPub_port       int
	HB_port          int
	Shell_port       int
	Key              string
	IP               string
}

// SocketGroup holds the sockets needed to communicate with the kernel, and
// the key for message signing.
type SocketGroup struct {
	Shell_socket   *zmq.Socket
	Control_socket *zmq.Socket
	Stdin_socket   *zmq.Socket
	IOPub_socket   *zmq.Socket
	Key            []byte
}

// PrepareSockets sets up the ZMQ sockets through which the kernel will communicate.
func PrepareSockets(conn_info ConnectionInfo) (sg SocketGroup) {

	context, _ := zmq.NewContext()
	sg.Shell_socket, _ = context.NewSocket(zmq.ROUTER)
	sg.Control_socket, _ = context.NewSocket(zmq.ROUTER)
	sg.Stdin_socket, _ = context.NewSocket(zmq.ROUTER)
	sg.IOPub_socket, _ = context.NewSocket(zmq.PUB)

	address := fmt.Sprintf("%v://%v:%%v", conn_info.Transport, conn_info.IP)

	sg.Shell_socket.Bind(fmt.Sprintf(address, conn_info.Shell_port))
	sg.Control_socket.Bind(fmt.Sprintf(address, conn_info.Control_port))
	sg.Stdin_socket.Bind(fmt.Sprintf(address, conn_info.Stdin_port))
	sg.IOPub_socket.Bind(fmt.Sprintf(address, conn_info.IOPub_port))

	// Message signing key
	sg.Key = []byte(conn_info.Key)

	// Start the heartbeat device
	HB_socket, _ := context.NewSocket(zmq.REP)
	HB_socket.Bind(fmt.Sprintf(address, conn_info.HB_port))
	go zmq.Device(zmq.FORWARDER, HB_socket, HB_socket)

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
	Protocol_version []int  `json:"protocol_version"`
	Language         string `json:"language"`
}

// KernelStatus holds a kernel state, for status broadcast messages.
type KernelStatus struct {
	ExecutionState string `json:"execution_state"`
}

// SendKernelInfo sends a kernel_info_reply message.
func SendKernelInfo(receipt MsgReceipt) {
	reply := NewMsg("kernel_info_reply", receipt.Msg)
	reply.Content = KernelInfo{[]int{4, 0}, "go"}
	receipt.SendResponse(receipt.Sockets.Shell_socket, reply)
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
	receipt.SendResponse(receipt.Sockets.Shell_socket, reply)
	logger.Println("Shutting down in response to shutdown_request")
	os.Exit(0)
}

// RunKernel is the main entry point to start the kernel. This is what is called by the
// gophernotes executable.
func RunKernel(connection_file string, logwriter io.Writer) {

	logger = log.New(logwriter, "gophernotes ", log.LstdFlags)

	// set up the "Session" with the replpkg
	SetupExecutionEnvironment()

	var conn_info ConnectionInfo
	bs, err := ioutil.ReadFile(connection_file)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(bs, &conn_info)
	if err != nil {
		log.Fatalln(err)
	}
	logger.Printf("%+v\n", conn_info)

	// Set up the ZMQ sockets through which the kernel will communicate
	sockets := PrepareSockets(conn_info)

	pi := zmq.PollItems{
		zmq.PollItem{Socket: sockets.Shell_socket, Events: zmq.POLLIN},
		zmq.PollItem{Socket: sockets.Stdin_socket, Events: zmq.POLLIN},
		zmq.PollItem{Socket: sockets.Control_socket, Events: zmq.POLLIN},
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
