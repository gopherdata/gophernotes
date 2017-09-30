package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/cosmos72/gomacro/ast2"
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
	HBSocket      *zmq.Socket
	Key           []byte
}

// KernelLanguageInfo holds information about the language that this kernel executes code in.
type kernelLanguageInfo struct {
	Name              string `json:"name"`
	Version           string `json:"version"`
	MIMEType          string `json:"mimetype"`
	FileExtension     string `json:"file_extension"`
	PygmentsLexer     string `json:"pygments_lexer"`
	CodeMirrorMode    string `json:"codemirror_mode"`
	NBConvertExporter string `json:"nbconvert_exporter"`
}

// HelpLink stores data to be displayed in the help menu of the notebook.
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

// shutdownReply encodes a boolean indication of stutdown/restart.
type shutdownReply struct {
	Restart bool `json:"restart"`
}

const (
	kernelStarting = "starting"
	kernelBusy     = "busy"
	kernelIdle     = "idle"
)

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

	// channelsWG waits for all channel handlers to shutdown.
	var channelsWG sync.WaitGroup

	// Start up the heartbeat handler.
	shutdownHeartbeat := runHeartbeat(sockets.HBSocket, &channelsWG)

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

	// Request that the heartbeat channel handler be shutdown.
	shutdownHeartbeat()

	// Wait for the channel handlers to finish shutting down.
	channelsWG.Wait()
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

	sg.HBSocket, err = context.NewSocket(zmq.REP)
	if err != nil {
		return sg, err
	}

	// Bind the sockets.
	address := fmt.Sprintf("%v://%v:%%v", connInfo.Transport, connInfo.IP)
	sg.ShellSocket.Bind(fmt.Sprintf(address, connInfo.ShellPort))
	sg.ControlSocket.Bind(fmt.Sprintf(address, connInfo.ControlPort))
	sg.StdinSocket.Bind(fmt.Sprintf(address, connInfo.StdinPort))
	sg.IOPubSocket.Bind(fmt.Sprintf(address, connInfo.IOPubPort))
	sg.HBSocket.Bind(fmt.Sprintf(address, connInfo.HBPort))

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
			ProtocolVersion:       ProtocolVersion,
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
	// Extract the data from the request.
	reqcontent := receipt.Msg.Content.(map[string]interface{})
	code := reqcontent["code"].(string)
	silent := reqcontent["silent"].(bool)

	if !silent {
		ExecCounter++
	}

	// Prepare the map that will hold the reply content.
	content := make(map[string]interface{})
	content["execution_count"] = ExecCounter

	// Tell the front-end that the kernel is working and when finished notify the
	// front-end that the kernel is idle again.
	if err := receipt.PublishKernelStatus(kernelBusy); err != nil {
		log.Printf("Error publishing kernel status 'busy': %v\n", err)
	}
	defer func() {
		if err := receipt.PublishKernelStatus(kernelIdle); err != nil {
			log.Printf("Error publishing kernel status 'idle': %v\n", err)
		}
	}()

	// Tell the front-end what the kernel is about to execute.
	if err := receipt.PublishExecutionInput(ExecCounter, code); err != nil {
		log.Printf("Error publishing execution input: %v\n", err)
	}

	// Redirect the standard out from the REPL.
	oldStdout := os.Stdout
	rOut, wOut, err := os.Pipe()
	if err != nil {
		return err
	}
	os.Stdout = wOut

	// Redirect the standard error from the REPL.
	oldStderr := os.Stderr
	rErr, wErr, err := os.Pipe()
	if err != nil {
		return err
	}
	os.Stderr = wErr

	var writersWG sync.WaitGroup
	writersWG.Add(2)

	// Forward all data written to stdout/stderr to the front-end.
	go func() {
		defer writersWG.Done()
		jupyterStdOut := JupyterStreamWriter{StreamStdout, &receipt}
		io.Copy(&jupyterStdOut, rOut)
	}()

	go func() {
		defer writersWG.Done()
		jupyterStdErr := JupyterStreamWriter{StreamStderr, &receipt}
		io.Copy(&jupyterStdErr, rErr)
	}()

	val, executionErr := doEval(ir, code)

	//TODO if value is a certain type like image then display it instead

	// Close and restore the streams.
	wOut.Close()
	os.Stdout = oldStdout

	wErr.Close()
	os.Stderr = oldStderr

	// Wait for the writers to finish forwarding the data.
	writersWG.Wait()

	if executionErr == nil {
		content["status"] = "ok"
		content["user_expressions"] = make(map[string]string)

		if !silent && val != nil {
			// Publish the result of the execution.
			if err := receipt.PublishExecutionResult(ExecCounter, fmt.Sprint(val)); err != nil {
				log.Printf("Error publishing execution result: %v\n", err)
			}
		}
	} else {
		content["status"] = "error"
		content["ename"] = "ERROR"
		content["evalue"] = executionErr.Error()
		content["traceback"] = nil

		if err := receipt.PublishExecutionError(executionErr.Error(), []string{executionErr.Error()}); err != nil {
			log.Printf("Error publishing execution error: %v\n", err)
		}
	}

	// Send the output back to the notebook.
	return receipt.Reply("execute_reply", content)
}

// doEval evaluates the code in the interpreter. This function captures an uncaught panic
// as well as the value of the last statement/expression.
func doEval(ir *classic.Interp, code string) (_ interface{}, err error) {
	// Capture a panic from the evaluation if one occurs and store it in the `err` return parameter.
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if err, ok = r.(error); !ok {
				err = errors.New(fmt.Sprint(r))
			}
		}
	}()

	// Prepare and perform the multiline evaluation.
	env := ir.Env
	env.Options &^= base.OptShowPrompt
	env.Options &^= base.OptTrapPanic
	env.Line = 0

	// Parse the input code (and don't preform gomacro's macroexpansion).
	src := ir.ParseOnly(code)

	if src == nil {
		return nil, nil
	}

	// Check if the last node is an expression.
	var srcEndsWithExpr bool

	// If the parsed ast is a single node, check if the node implements `ast.Expr`. Otherwise if the is multiple
	// nodes then just check if the last one is an expression.
	if srcAstWithNode, ok := src.(ast2.AstWithNode); ok {
		_, srcEndsWithExpr = srcAstWithNode.Node().(ast.Expr)
	} else if srcNodeSlice, ok := src.(ast2.NodeSlice); ok {
		nodes := srcNodeSlice.X
		_, srcEndsWithExpr = nodes[len(nodes)-1].(ast.Expr)
	}

	// Evaluate the code.
	result, results := ir.EvalAst(src)

	// If the source ends with an expression, then the result of the execution is the value of the expression. In the
	// case of multiple return values (from a function call for example), the first non-nil value is the result.
	if srcEndsWithExpr {
		// `len(results) == 0` implies a single result stored in `result`.
		if len(results) == 0 {
			return base.ValueInterface(result), nil
		}

		// Set `val` to be the first non-nil result.
		for _, result := range results {
			val := base.ValueInterface(result)
			if val != nil {
				return val, nil
			}
		}
	}

	return nil, nil
}

// handleShutdownRequest sends a "shutdown" message.
func handleShutdownRequest(receipt msgReceipt) {
	content := receipt.Msg.Content.(map[string]interface{})
	restart := content["restart"].(bool)

	reply := shutdownReply{
		Restart: restart,
	}

	if err := receipt.Reply("shutdown_reply", reply); err != nil {
		log.Fatal(err)
	}

	log.Println("Shutting down in response to shutdown_request")
	os.Exit(0)
}

// runHeartbeat starts a go-routine for handling heartbeat ping messages sent over the given `hbSocket`. The `wg`'s
// `Done` method is invoked after the thread is completely shutdown. To request a shutdown the returned `func()` can
// be called.
func runHeartbeat(hbSocket *zmq.Socket, wg *sync.WaitGroup) func() {
	quit := make(chan bool)

	// Start the handler that will echo any received messages back to the sender.
	wg.Add(1)
	go func() {
		defer wg.Done()

		// Create a `Poller` to check for incoming messages.
		poller := zmq.NewPoller()
		poller.Add(hbSocket, zmq.POLLIN)

		for {
			select {
			case <-quit:
				return
			default:
				// Check for received messages waiting at most 500ms for once to arrive.
				pingEvents, err := poller.Poll(500 * time.Millisecond)
				if err != nil {
					log.Fatalf("Error polling heartbeat channel: %v\n", err)
				}

				// If there is at least 1 message waiting then echo it.
				if len(pingEvents) > 0 {
					// Read a message from the heartbeat channel as a simple byte string.
					pingMsg, err := hbSocket.RecvBytes(0)
					if err != nil {
						log.Fatalf("Error reading heartbeat ping bytes: %v\n", err)
					}

					// Send the received byte string back to let the front-end know that the kernel is alive.
					_, err = hbSocket.SendBytes(pingMsg, 0)
					if err != nil {
						log.Printf("Error sending heartbeat pong bytes: %b\n", err)
					}
				}
			}
		}
	}()

	// Wrap the quit channel in a function that writes `true` to the channel to shutdown the handler.
	return func() {
		quit <- true
	}
}
