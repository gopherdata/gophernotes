package main

import (
	interp "github.com/cosmos72/gomacro/fast"
)

type Completion struct {
	class,
	name,
	typ string
}

type CompletionResponse struct {
	partial     int
	completions []Completion
}

/************************************************************
* entry function
************************************************************/
func handleCompleteRequest(ir *interp.Interp, receipt msgReceipt) error {
	// Extract the data from the request.
	reqcontent := receipt.Msg.Content.(map[string]interface{})
	code := reqcontent["code"].(string)
	cursorPos := int(reqcontent["cursor_pos"].(float64))

	// autocomplete the code at the cursor position
	prefix, matches, _ := ir.CompleteWords(code, cursorPos)

	// prepare the reply
	content := make(map[string]interface{})

	if len(matches) == 0 {
		content["ename"] = "ERROR"
		content["evalue"] = "no completions found"
		content["traceback"] = nil
		content["status"] = "error"
	} else {
		partialWord := interp.TailIdentifier(prefix)
		content["cursor_start"] = float64(len(prefix) - len(partialWord))
		content["cursor_end"] = float64(cursorPos)
		content["matches"] = matches
		content["status"] = "ok"
	}

	return receipt.Reply("complete_reply", content)
}
