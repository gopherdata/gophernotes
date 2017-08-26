package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/cosmos72/gomacro/base"
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

// kernelStatus holds a kernel state, for status broadcast messages.
type kernelStatus struct {
	ExecutionState string `json:"execution_state"`
}

// KernelLanguageInfo holds information about the language that this kernel executes code in
type kernelLanguageInfo struct {
	Name              string `json:"name"`
	Version           string `json:"version"`
	MIMEType          string `json:"mimetype"`
	FileExtension     string `json:"file_extension"`
	PygmentsLexer     string `json:"pygments_lexer"`
	CodeMirrorMode    string `json:"codemirror_mode"`
	NBConvertExporter string `json:"nbconvert_exporter"`
}

// HelpLink stores data to be displayed in the help menu of the notebook
type helpLink struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

// KernelInfo holds information about the igo kernel, for kernel_info_reply messages.
type kernelInfo struct {
	ProtocolVersion       string             `json:"protocol_version"`
	Implementation        string             `json:"implementation"`
	ImplementationVersion string             `json:"implementation_version"`
	LanguageInfo          kernelLanguageInfo `json:"language_info"`
	Banner                string             `json:"banner"`
	HelpLinks             []helpLink         `json:"help_links"`
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
	return receipt.Reply("kernel_info_reply",
		kernelInfo{
			ProtocolVersion:       "5.0",
			Implementation:        "gophernotes",
			ImplementationVersion: Version,
			Banner:                fmt.Sprintf("Go kernel: gophernotes - v%s", Version),
			LanguageInfo: kernelLanguageInfo{
				Name:          "go",
				Version:       runtime.Version(),
				FileExtension: ".go",
			},
			HelpLinks: []helpLink{
				{Text: "Go", URL: "https://golang.org/"},
				{Text: "gophernotes", URL: "https://github.com/gopherdata/gophernotes"},
			},
		},
	)
}

// handleExecuteRequest runs code from an execute_request method,
// and sends the various reply messages.
func handleExecuteRequest(ir *classic.Interp, receipt msgReceipt) error {
	// Extract the data from the request
	reqcontent := receipt.Msg.Content.(map[string]interface{})
	code := reqcontent["code"].(string)
	in := bufio.NewReader(strings.NewReader(code))
	silent := reqcontent["silent"].(bool)

	if !silent {
		ExecCounter++
	}

	// Prepare the map that will hold the reply content
	content := make(map[string]interface{})
	content["execution_count"] = ExecCounter

	// Tell the front-end that the kernel is working and when finished notify the
	// front-end that the kernel is idle again
	receipt.PublishKernelBusy()
	defer receipt.PublishKernelIdle()

	// Tell the front-end what the kernel is about to execute
	receipt.PublishExecutionInput(ExecCounter, code)

	// Redirect the standard out from the REPL.
	oldStdout := os.Stdout
	rOut, wOut, err := os.Pipe()
	if err != nil {
		return err
	}
	os.Stdout = wOut

	// Redirect the standard error from the REPL.
	rErr, wErr, err := os.Pipe()
	if err != nil {
		return err
	}
	ir.Stderr = wErr

	// Prepare and perform the multiline evaluation.
	env := ir.Env
	env.Options &^= base.OptShowPrompt
	env.Line = 0

	// Perform the first iteration manually, to collect comments
	var comments string
	str, firstToken := env.ReadMultiline(in, base.ReadOptCollectAllComments)
	if firstToken >= 0 {
		comments = str[0:firstToken]
		if firstToken > 0 {
			str = str[firstToken:]
			env.IncLine(comments)
		}
	}
	if ir.ParseEvalPrint(str, in) {
		ir.Repl(in)
	}

	// Copy the stdout in a separate goroutine to prevent
	// blocking on printing.
	outStdout := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, rOut)
		outStdout <- buf.String()
	}()

	// Return stdout back to normal state.
	wOut.Close()
	os.Stdout = oldStdout
	val := <-outStdout

	// Copy the stderr in a separate goroutine to prevent
	// blocking on printing.
	outStderr := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, rErr)
		outStderr <- buf.String()
	}()

	wErr.Close()
	stdErr := <-outStderr

	// TODO write stdout and stderr to streams rather than publishing as results

	if len(val) > 0 {
		content["status"] = "ok"
		content["user_expressions"] = make(map[string]string)

		if !silent {
			// Publish the result of the execution
			receipt.PublishExecutionResult(ExecCounter, val)
		}
	}

	if len(stdErr) > 0 {
		content["status"] = "error"
		content["ename"] = "ERROR"
		content["evalue"] = stdErr
		content["traceback"] = nil

		receipt.PublishExecutionError(stdErr, stdErr)
	}

	// Send the output back to the notebook.
	return receipt.Reply("execute_reply", content)
}

// handleShutdownRequest sends a "shutdown" message
func handleShutdownRequest(receipt msgReceipt) {
	content := receipt.Msg.Content.(map[string]interface{})
	restart := content["restart"].(bool)

	err := receipt.Reply("shutdown_reply",
		shutdownReply{
			Restart: restart,
		},
	)

	if err != nil {
		log.Fatal(err)
	}


	log.Println("Shutting down in response to shutdown_request")
	os.Exit(0)
}
