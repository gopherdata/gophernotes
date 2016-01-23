package main

// This file deals with executing user code and sending output messages.

import (
    "fmt"
    //eval "github.com/sbinet/go-eval"
    gore "github.com/motemen/gore"
    "go/token"
)

// // World holds the user namespace for the REPL.
// var World *eval.World
// var fset *token.FileSet
// // ExecCounter is incremented each time we run user code.
// var ExecCounter int = 0

func SetupExecutionEnvironment() {

    s, err := gore.NewSession()
    if err != nil {
        panic(err)
    }

    // World = eval.NewWorld()
    // fset = token.NewFileSet()
}

// RunCode runs the given user code, returning the expression value and/or an error.
func RunCode(text string) (val interface{}, err error) {
    var code eval.Code
    code, err = World.Compile(fset, text)
    if err != nil {
        return nil, err
    }
    val, err = code.Run()
    return
}

// OutputMsg holds the data for a pyout message.
type OutputMsg struct {
    Execcount int `json:"execution_count"`
    Data map[string]string `json:"data"`
    Metadata map[string]interface{} `json:"metadata"`
}

type ErrMsg struct {
    EName string `json:"ename"`
    EValue string `json:"evalue"`
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
    val, err := RunCode(code)
    if err == nil {
        content["status"] = "ok"
        content["payload"] = make([]map[string]interface{}, 0)
        content["user_variables"] = make(map[string]string)
        content["user_expressions"] = make(map[string]string)
        if (val != nil) && !silent {
            var out_content OutputMsg
            out := NewMsg("pyout", receipt.Msg)
            out_content.Execcount = ExecCounter
            out_content.Data = make(map[string]string)
            out_content.Data["text/plain"] = fmt.Sprint(val)
            out_content.Metadata = make(map[string]interface{})
            out.Content = out_content
            receipt.SendResponse(receipt.Sockets.IOPub_socket, out)
        }
    } else {
        content["status"] = "error"
        content["ename"] = "ERROR"
        content["evalue"] = err.Error()
        content["traceback"] = []string{err.Error()}
        errormsg := NewMsg("pyerr", receipt.Msg)
        errormsg.Content = ErrMsg{"Error", err.Error(), []string{err.Error()}}
        receipt.SendResponse(receipt.Sockets.IOPub_socket, errormsg)
    }
    reply.Content = content
    receipt.SendResponse(receipt.Sockets.Shell_socket, reply)
    idle := NewMsg("status", receipt.Msg)
    idle.Content = KernelStatus{"idle"}
    receipt.SendResponse(receipt.Sockets.IOPub_socket, idle)
}