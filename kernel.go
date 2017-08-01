package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/cosmos72/gomacro/classic"
	zmq "github.com/pebbe/zmq4"
)

// ExecCounter is incremented each time we run user code in the notebook.
var ExecCounter int

// ConnectionInfo stores the contents of the kernel connection
// file created by Jupyter.
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

// SocketGroup holds the sockets needed to communicate with the kernel,
// and the key for message signing.
type SocketGroup struct {
	ShellSocket   *zmq.Socket
	ControlSocket *zmq.Socket
	StdinSocket   *zmq.Socket
	IOPubSocket   *zmq.Socket
	Key           []byte
}

// kernelInfo holds information about the igo kernel, for
// kernel_info_reply messages.
type kernelInfo struct {
	ProtocolVersion []int  `json:"protocol_version"`
	Language        string `json:"language"`
}

// kernelStatus holds a kernel state, for status broadcast messages.
type kernelStatus struct {
	ExecutionState string `json:"execution_state"`
}

// shutdownReply encodes a boolean indication of stutdown/restart
type shutdownReply struct {
	Restart bool `json:"restart"`
}

// runKernel is the main entry point to start the kernel.
func runKernel(connectionFile string) {

	// Set up the "Session" with the replpkg.
	ir := classic.New()

	// Parse the connection info.
	var connInfo ConnectionInfo

	connData, err := ioutil.ReadFile(connectionFile)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(connData, &connInfo); err != nil {
		log.Fatal(err)
	}

	// Set up the ZMQ sockets through which the kernel will communicate.
	sockets, err := prepareSockets(connInfo)
	if err != nil {
		log.Fatal(err)
	}

	poller := zmq.NewPoller()
	poller.Add(sockets.ShellSocket, zmq.POLLIN)
	poller.Add(sockets.StdinSocket, zmq.POLLIN)
	poller.Add(sockets.ControlSocket, zmq.POLLIN)

	// msgParts will store a received multipart message.
	var msgParts [][]byte

	// Start a message receiving loop.
	for {
		polled, err := poller.Poll(-1)
		if err != nil {
			log.Fatal(err)
		}

		for _, item := range polled {

			// Handle various types of messages.
			switch socket := item.Socket; socket {

			// Handle shell messages.
			case sockets.ShellSocket:
				msgParts, err = sockets.ShellSocket.RecvMessageBytes(0)
				if err != nil {
					log.Println(err)
				}

				msg, ids, err := WireMsgToComposedMsg(msgParts, sockets.Key)
				if err != nil {
					log.Println(err)
					return
				}

				handleShellMsg(ir, msgReceipt{msg, ids, sockets})

			// TODO Handle stdin socket.
			case sockets.StdinSocket:
				sockets.StdinSocket.RecvMessageBytes(0)

			// Handle control messages.
			case sockets.ControlSocket:
				msgParts, err = sockets.ControlSocket.RecvMessageBytes(0)
				if err != nil {
					log.Println(err)
					return
				}

				msg, ids, err := WireMsgToComposedMsg(msgParts, sockets.Key)
				if err != nil {
					log.Println(err)
					return
				}

				handleShellMsg(ir, msgReceipt{msg, ids, sockets})
			}
		}
	}
}

// prepareSockets sets up the ZMQ sockets through which the kernel
// will communicate.
func prepareSockets(connInfo ConnectionInfo) (SocketGroup, error) {

	// Initialize the context.
	context, err := zmq.NewContext()
	if err != nil {
		return SocketGroup{}, err
	}

	// Initialize the socket group.
	var sg SocketGroup

	sg.ShellSocket, err = context.NewSocket(zmq.ROUTER)
	if err != nil {
		return sg, err
	}

	sg.ControlSocket, err = context.NewSocket(zmq.ROUTER)
	if err != nil {
		return sg, err
	}

	sg.StdinSocket, err = context.NewSocket(zmq.ROUTER)
	if err != nil {
		return sg, err
	}

	sg.IOPubSocket, err = context.NewSocket(zmq.PUB)
	if err != nil {
		return sg, err
	}

	// Bind the sockets.
	address := fmt.Sprintf("%v://%v:%%v", connInfo.Transport, connInfo.IP)
	sg.ShellSocket.Bind(fmt.Sprintf(address, connInfo.ShellPort))
	sg.ControlSocket.Bind(fmt.Sprintf(address, connInfo.ControlPort))
	sg.StdinSocket.Bind(fmt.Sprintf(address, connInfo.StdinPort))
	sg.IOPubSocket.Bind(fmt.Sprintf(address, connInfo.IOPubPort))

	// Set the message signing key.
	sg.Key = []byte(connInfo.Key)

	return sg, nil
}

// handleShellMsg responds to a message on the shell ROUTER socket.
func handleShellMsg(ir *classic.Interp, receipt msgReceipt) {
	switch receipt.Msg.Header.MsgType {
	case "kernel_info_request":
		if err := sendKernelInfo(receipt); err != nil {
			log.Fatal(err)
		}
	case "execute_request":
		if err := handleExecuteRequest(ir, receipt); err != nil {
			log.Fatal(err)
		}
	case "shutdown_request":
		handleShutdownRequest(receipt)
	default:
		log.Println("Unhandled shell message: ", receipt.Msg.Header.MsgType)
	}
}

// sendKernelInfo sends a kernel_info_reply message.
func sendKernelInfo(receipt msgReceipt) error {
	reply, err := NewMsg("kernel_info_reply", receipt.Msg)
	if err != nil {
		return err
	}

	reply.Content = kernelInfo{[]int{4, 0}, "go"}
	if err := receipt.SendResponse(receipt.Sockets.ShellSocket, reply); err != nil {
		return err
	}

	return nil
}

// handleExecuteRequest runs code from an execute_request method,
// and sends the various reply messages.
func handleExecuteRequest(ir *classic.Interp, receipt msgReceipt) error {

	reply, err := NewMsg("execute_reply", receipt.Msg)
	if err != nil {
		return err
	}

	content := make(map[string]interface{})

	reqcontent := receipt.Msg.Content.(map[string]interface{})
	code := reqcontent["code"].(string)
	silent := reqcontent["silent"].(bool)

	if !silent {
		ExecCounter++
	}

	content["execution_count"] = ExecCounter

	// Do the compilation/execution magic.
	rawVal, rawErr := ir.Eval(code)
	val := fmt.Sprintln(rawVal)

	fmt.Printf("\nThis is val: %s\n\n", val)
	fmt.Printf("This is rawErr: %s\n\n", fmt.Sprintln(rawErr))

	if len(val) > 0 {
		content["status"] = "ok"
		content["payload"] = make([]map[string]interface{}, 0)
		content["user_variables"] = make(map[string]string)
		content["user_expressions"] = make(map[string]string)
		if !silent {
			var outContent OutputMsg

			out, err := NewMsg("pyout", receipt.Msg)
			if err != nil {
				return err
			}

			outContent.Execcount = ExecCounter
			outContent.Data = make(map[string]string)
			outContent.Data["text/plain"] = val
			outContent.Metadata = make(map[string]interface{})
			out.Content = outContent
			receipt.SendResponse(receipt.Sockets.IOPubSocket, out)
		}
	} else {
		content["status"] = "error"
		content["ename"] = "ERROR"
		content["evalue"] = fmt.Sprintln(rawErr)
		content["traceback"] = nil

		errormsg, err := NewMsg("pyerr", receipt.Msg)
		if err != nil {
			return err
		}

		errormsg.Content = ErrMsg{"Error", fmt.Sprintln(rawErr), nil}
		receipt.SendResponse(receipt.Sockets.IOPubSocket, errormsg)
	}

	// Send the output back to the notebook.
	reply.Content = content

	if err := receipt.SendResponse(receipt.Sockets.ShellSocket, reply); err != nil {
		return err
	}

	idle, err := NewMsg("status", receipt.Msg)
	if err != nil {
		return err
	}

	idle.Content = kernelStatus{"idle"}

	if err := receipt.SendResponse(receipt.Sockets.IOPubSocket, idle); err != nil {
		return err
	}

	return nil
}

// handleShutdownRequest sends a "shutdown" message
func handleShutdownRequest(receipt msgReceipt) {
	reply, err := NewMsg("shutdown_reply", receipt.Msg)
	if err != nil {
		log.Fatal(err)
	}

	content := receipt.Msg.Content.(map[string]interface{})
	restart := content["restart"].(bool)
	reply.Content = shutdownReply{restart}

	if err := receipt.SendResponse(receipt.Sockets.ShellSocket, reply); err != nil {
		log.Fatal(err)
	}

	log.Println("Shutting down in response to shutdown_request")
	os.Exit(0)
}
