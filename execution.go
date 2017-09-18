package main

import (
	"fmt"
	"go/token"

	repl "github.com/gopherds/gophernotes/internal/repl"
	"strings"
)

var (
	// REPLSession manages the I/O to/from the notebook.
	REPLSession *repl.Session
	fset        *token.FileSet
)

// ExecCounter is incremented each time we run user code in the notebook.
var ExecCounter int

// SetupExecutionEnvironment initializes the REPL session and set of tmp files.
func SetupExecutionEnvironment() {

	var err error
	REPLSession, err = repl.NewSession()
	if err != nil {
		panic(err)
	}

	fset = token.NewFileSet()
}

type KernelStatusMsg struct {
	ExecState string `json:"execution_state"`
}

func setBusy(receipt MsgReceipt) {
	var status KernelStatusMsg

	msg := NewMsg("status", receipt.Msg)
	status.ExecState = "busy"
	msg.Content = status

	receipt.SendResponse(receipt.Sockets.IOPubSocket, msg)
}

func setIdle(receipt MsgReceipt) {
	var status KernelStatusMsg

	msg := NewMsg("status", receipt.Msg)
	status.ExecState = "idle"
	msg.Content = status

	receipt.SendResponse(receipt.Sockets.IOPubSocket, msg)
}


type ExecuteInputMsg struct {
	ExecCount int    `json:"execution_count"`
	Code      string `json:"code"`
}

func publishExecuteInput(receipt MsgReceipt, execCount int, code string) {
	var pubExecInputContent ExecuteInputMsg
	pubExecInput := NewMsg("execute_input", receipt.Msg)
	pubExecInputContent.ExecCount = execCount
	pubExecInputContent.Code = code
	receipt.SendResponse(receipt.Sockets.IOPubSocket, pubExecInput)
}

type StreamMsg struct {
	Stream string `json:"name"`
	Data string `json:"text"`
}
func publishWriteStream(receipt MsgReceipt, stdout bool, text string) {
	var pubWriteStreamContent StreamMsg
	pubWriteStream := NewMsg("stream", receipt.Msg)

	if stdout {
		pubWriteStreamContent.Stream = "stdout"
	} else {
		pubWriteStreamContent.Stream = "stderr"
	}
	pubWriteStreamContent.Data = text

	receipt.SendResponse(receipt.Sockets.IOPubSocket, pubWriteStream)
}

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

type ExecuteResultMsg struct {
	Execcount int                    `json:"execution_count"`
	Data      map[string]string      `json:"data"`
	Metadata  map[string]interface{} `json:"metadata"`
}

// HandleExecuteRequest runs code from an execute_request method, and sends the various
// reply messages.
func HandleExecuteRequest(receipt MsgReceipt) {

	reply := NewMsg("execute_reply", receipt.Msg)
	replyContent := make(map[string]interface{})

	reqContent := receipt.Msg.Content.(map[string]interface{})
	code := reqContent["code"].(string)
	silent := reqContent["silent"].(bool)

	if !silent {
		ExecCounter++
	}
	replyContent["execution_count"] = ExecCounter

	// Tell the front-end the kernel is working and when finished publish idle status
    setBusy(receipt)
	defer setIdle(receipt)
    
    // Tell the front-end what is being executed
    publishExecuteInput(receipt, ExecCounter, code)
    
	// Do the compilation/execution magic.
	val, stderr, err := REPLSession.Eval(code)

	if err == nil {
		replyContent["status"] = "ok"
		replyContent["payload"] = make([]map[string]interface{}, 0)
		replyContent["user_variables"] = make(map[string]string)
		replyContent["user_expressions"] = make(map[string]string)
        
		if !silent && len(val) > 0 {
			var outContent OutputMsg
			out := NewMsg("execute_result", receipt.Msg)
            
			outContent.Execcount = ExecCounter
			outContent.Data = make(map[string]string)
			outContent.Data["text/plain"] = fmt.Sprintf("%s", val)
			outContent.Metadata = make(map[string]interface{})
            
			out.Content = outContent
            
			receipt.SendResponse(receipt.Sockets.IOPubSocket, out)
		}

		if !silent && stderr.String() != "" {
			publishWriteStream(receipt, false, stderr.String())
		}
	} else {
		replyContent["status"] = "error"
		replyContent["ename"] = "ERROR"
		replyContent["evalue"] = err.Error()
		replyContent["traceback"] = []string{strings.TrimSpace(stderr.String())}

		errMsg := NewMsg("error", receipt.Msg)
		errMsg.Content = ErrMsg{"Error", err.Error(), []string{strings.TrimSpace(stderr.String())}}
		receipt.SendResponse(receipt.Sockets.IOPubSocket, errMsg)
	}

	// send the output back to the notebook
	reply.Content = replyContent
	receipt.SendResponse(receipt.Sockets.ShellSocket, reply)
}
