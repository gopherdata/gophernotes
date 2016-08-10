package main

import (
	"fmt"
	"go/token"

	repl "github.com/gopherds/gophernotes/internal/repl"
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

// HandleExecuteRequest runs code from an execute_request method, and sends the various
// reply messages.
func HandleExecuteRequest(receipt MsgReceipt) {

	reply := NewMsg("execute_reply", receipt.Msg)
	content := make(map[string]interface{})
	reqcontent := receipt.Msg.Content.(map[string]interface{})
	code := reqcontent["code"].(string)
	silent := reqcontent["silent"].(bool)
	if !silent {
		ExecCounter++
	}
	content["execution_count"] = ExecCounter

	// Do the compilation/execution magic.
	val, stderr, err := REPLSession.Eval(code)

	if err == nil {
		content["status"] = "ok"
		content["payload"] = make([]map[string]interface{}, 0)
		content["user_variables"] = make(map[string]string)
		content["user_expressions"] = make(map[string]string)
		if len(val) > 0 && !silent {
			var outContent OutputMsg
			out := NewMsg("pyout", receipt.Msg)
			outContent.Execcount = ExecCounter
			outContent.Data = make(map[string]string)
			outContent.Data["text/plain"] = fmt.Sprint(val)
			outContent.Metadata = make(map[string]interface{})
			out.Content = outContent
			receipt.SendResponse(receipt.Sockets.IOPubSocket, out)
		}
	} else {
		content["status"] = "error"
		content["ename"] = "ERROR"
		content["evalue"] = err.Error()
		content["traceback"] = []string{stderr.String()}
		errormsg := NewMsg("pyerr", receipt.Msg)
		errormsg.Content = ErrMsg{"Error", err.Error(), []string{stderr.String()}}
		receipt.SendResponse(receipt.Sockets.IOPubSocket, errormsg)
	}

	// send the output back to the notebook
	reply.Content = content
	receipt.SendResponse(receipt.Sockets.ShellSocket, reply)
	idle := NewMsg("status", receipt.Msg)
	idle.Content = KernelStatus{"idle"}
	receipt.SendResponse(receipt.Sockets.IOPubSocket, idle)
}
