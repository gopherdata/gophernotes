package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	zmq "github.com/pebbe/zmq4"
)

const (
	failure = "\u2717"
	success = "\u2713"
)

const (
	connectionKey = "a0436f6c-1916-498b-8eb9-e81ab9368e84"
	sessionID     = "ba65a05c-106a-4799-9a94-7f5631bbe216"
	transport     = "tcp"
	ip            = "127.0.0.1"
	shellPort     = 57503
	iopubPort     = 40885
)

type testJupyterClient struct {
	shellSocket *zmq.Socket
	ioSocket    *zmq.Socket
}

// newTestJupyterClient creates and connects a fresh client to the kernel. Upon error, newTestJupyterClient
// will Fail the test.
func newTestJupyterClient(t *testing.T) (testJupyterClient, func()) {
	addrShell := fmt.Sprintf("%s://%s:%d", transport, ip, shellPort)
	addrIO := fmt.Sprintf("%s://%s:%d", transport, ip, iopubPort)

	// Prepare the shell socket.
	shell, err := zmq.NewSocket(zmq.REQ)
	if err != nil {
		t.Fatal("NewSocket:", err)
	}

	if err = shell.Connect(addrShell); err != nil {
		t.Fatal("shell.Connect:", err)
	}

	// Prepare the IOPub socket.
	iopub, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		t.Fatal("NewSocket:", err)
	}

	if err = iopub.Connect(addrIO); err != nil {
		t.Fatal("iopub.Connect:", err)
	}

	if err = iopub.SetSubscribe(""); err != nil {
		t.Fatal("iopub.SetSubscribe", err)
	}

	// wait for a second to give the tcp connection time to complete to avoid missing the early pub messages
	time.Sleep(1 * time.Second)

	return testJupyterClient{shell, iopub}, func() {
		if err := shell.Close(); err != nil {
			t.Fatal("shell.Close", err)
		}
		if err = iopub.Close(); err != nil {
			t.Fatal("iopub.Close", err)
		}
	}
}

// sendShellRequest sends a message to the kernel over the shell channel. Upon error, sendShellRequest
// will Fail the test.
func (client *testJupyterClient) sendShellRequest(t *testing.T, request ComposedMsg) {
	if _, err := client.shellSocket.Send("<IDS|MSG>", zmq.SNDMORE); err != nil {
		t.Fatal("shellSocket.Send:", err)
	}

	reqMsgParts, err := request.ToWireMsg([]byte(connectionKey))
	if err != nil {
		t.Fatal("request.ToWireMsg:", err)
	}

	if _, err = client.shellSocket.SendMessage(reqMsgParts); err != nil {
		t.Fatal("shellSocket.SendMessage:", err)
	}
}

// recvShellReply tries to read a reply message from the shell channel. It will timeout after the given
// timeout delay. Upon error or timeout, recvShellReply will Fail the test.
func (client *testJupyterClient) recvShellReply(t *testing.T, timeout time.Duration) (reply ComposedMsg) {
	ch := make(chan ComposedMsg)

	go func() {
		repMsgParts, err := client.shellSocket.RecvMessageBytes(0)
		if err != nil {
			t.Fatal("Shell socket RecvMessageBytes:", err)
		}

		msgParsed, _, err := WireMsgToComposedMsg(repMsgParts, []byte(connectionKey))
		if err != nil {
			t.Fatal("Could not parse wire message:", err)
		}

		ch <- msgParsed
	}()

	select {
	case reply = <-ch:
	case <-time.After(timeout):
		t.Fatal("recvShellReply timed out")
	}

	return
}

// recvIOSub tries to read a published message from the IOPub channel. It will timeout after the given
// timeout delay. Upon error or timeout, recvIOSub will Fail the test.
func (client *testJupyterClient) recvIOSub(t *testing.T, timeout time.Duration) (sub ComposedMsg) {
	ch := make(chan ComposedMsg)

	go func() {
		repMsgParts, err := client.ioSocket.RecvMessageBytes(0)
		if err != nil {
			t.Fatal("IOPub socket RecvMessageBytes:", err)
		}

		msgParsed, _, err := WireMsgToComposedMsg(repMsgParts, []byte(connectionKey))
		if err != nil {
			t.Fatal("Could not parse wire message:", err)
		}

		ch <- msgParsed
	}()

	select {
	case sub = <-ch:
	case <-time.After(timeout):
		t.Fatal("recvIOSub timed out")
	}

	return
}

// request preforms a request and awaits a reply on the shell channel. Additionally all messages on the IOPub channel
// between the opening 'busy' messages and closing 'idle' message are captured and returned. The request will timeout
// after the given timeout delay. Upon error or timeout, request will Fail the test.
func (client *testJupyterClient) request(t *testing.T, request ComposedMsg, timeout time.Duration) (reply ComposedMsg, pub []ComposedMsg) {
	client.sendShellRequest(t, request)
	reply = client.recvShellReply(t, timeout)

	// Read the expected 'busy' message and ensure it is in fact, a 'busy' message
	subMsg := client.recvIOSub(t, 1*time.Second)
	if subMsg.Header.MsgType != "status" {
		t.Fatalf("Expected a 'status' message but received a '%s' message on IOPub", subMsg.Header.MsgType)
	}

	subData, ok := subMsg.Content.(map[string]interface{})
	if !ok {
		t.Fatal("'status' message content is not a json object")
	}

	execState, ok := subData["execution_state"]
	if !ok {
		t.Fatal("'status' message content is missing the 'execution_state' field")
	}

	if execState != kernelBusy {
		t.Fatalf("Expected a 'busy' status message but got '%v'", execState)
	}

	// Read messages from the IOPub channel until an 'idle' message is received
	for {
		subMsg = client.recvIOSub(t, 100*time.Millisecond)

		// If the message is a 'status' message, ensure it is an 'idle' status
		if subMsg.Header.MsgType == "status" {
			subData, ok = subMsg.Content.(map[string]interface{})
			if !ok {
				t.Fatal("'status' message content is not a json object")
			}

			execState, ok = subData["execution_state"]
			if !ok {
				t.Fatal("'status' message content is missing the 'execution_state' field")
			}

			if execState != kernelIdle {
				t.Fatalf("Expected a 'idle' status message but got '%v'", execState)
			}

			// Break from the loop as we don't expect any other IOPub messages after the 'idle'
			break
		}

		// Add the message to the pub collection
		pub = append(pub, subMsg)
	}

	return
}

func TestMain(m *testing.M) {
	os.Exit(runTest(m))
}

// runTest initializes the environment for the tests and allows for
// the proper exit if the test fails or succeeds.
func runTest(m *testing.M) int {

	// Start the kernel.
	go runKernel("fixtures/connection_file.json")

	return m.Run()
}

//==============================================================================

// TestEvaluate tests the evaluation of consecutive cells..
func TestEvaluate(t *testing.T) {
	cases := []struct {
		Input  []string
		Output string
	}{
		{[]string{
			"import \"fmt\"",
			"a := 1",
			"fmt.Println(a)",
		}, "1\n"},
		{[]string{
			"a = 2",
			"fmt.Println(a)",
		}, "2\n"},
		{[]string{
			"func myFunc(x int) int {",
			"    return x+1",
			"}",
			"fmt.Println(\"func defined\")",
		}, "func defined\n"},
		{[]string{
			"b := myFunc(1)",
			"fmt.Println(b)",
		}, "2\n"},
	}

	t.Logf("Should be able to evaluate valid code in notebook cells.")

	for k, tc := range cases {

		// Give a progress report.
		t.Logf("  Evaluating code snippet %d/%d.", k+1, len(cases))

		// Get the result.
		result := testEvaluate(t, strings.Join(tc.Input, "\n"))

		// Compare the result.
		if result != tc.Output {
			t.Errorf("\t%s Test case produced unexpected results.", failure)
			continue
		}
		t.Logf("\t%s Should return the correct cell output.", success)
	}
}

// testEvaluate evaluates a cell.
func testEvaluate(t *testing.T, codeIn string) string {
	client, closeClient := newTestJupyterClient(t)
	defer closeClient()

	// Create a message.
	request, err := NewMsg("execute_request", ComposedMsg{})
	if err != nil {
		t.Fatal("NewMessage:", err)
	}

	// Fill in remaining header information.
	request.Header.Session = sessionID
	request.Header.Username = "KernelTester"

	// Fill in Metadata.
	request.Metadata = make(map[string]interface{})

	// Fill in content.
	content := make(map[string]interface{})
	content["code"] = codeIn
	content["silent"] = false
	request.Content = content

	reply, pub := client.request(t, request, 10*time.Second)

	if reply.Header.MsgType != "execute_reply" {
		t.Fatal("reply.Header.MsgType", errors.New("reply is not an 'execute_reply'"))
	}

	content, ok := reply.Content.(map[string]interface{})
	if !ok {
		t.Fatal("reply.Content.(map[string]interface{})", errors.New("reply content is not a json object"))
	}

	statusRaw, ok := content["status"]
	if !ok {
		t.Fatal("content[\"status\"]", errors.New("status field not present in 'execute_reply'"))
	}

	status, ok := statusRaw.(string)
	if !ok {
		t.Fatal("content[\"status\"]", errors.New("status field value is not a string"))
	}

	if status != "ok" {
		t.Fatalf("Execution encountered error [%s]: %s", content["ename"], content["evalue"])
	}

	for _, pubMsg := range pub {
		if pubMsg.Header.MsgType == "execute_result" {
			content, ok := pubMsg.Content.(map[string]interface{})
			if !ok {
				t.Fatal("pubMsg.Content.(map[string]interface{})", errors.New("pubMsg 'execute_result' content is not a json object"))
			}

			bundledMIMEDataRaw, ok := content["data"]
			if !ok {
				t.Fatal("content[\"data\"]", errors.New("data field not present in 'execute_result'"))
			}

			bundledMIMEData, ok := bundledMIMEDataRaw.(map[string]interface{})
			if !ok {
				t.Fatal("content[\"data\"]", errors.New("data field is not a MIME data bundle in 'execute_result'"))
			}

			textRepRaw, ok := bundledMIMEData["text/plain"]
			if !ok {
				t.Fatal("content[\"data\"]", errors.New("data field doesn't contain a text representation in 'execute_result'"))
			}

			textRep, ok := textRepRaw.(string)
			if !ok {
				t.Fatal("content[\"data\"][\"text/plain\"]", errors.New("text representation is not a string in 'execute_result'"))
			}

			return textRep
		}
	}

	return ""
}
