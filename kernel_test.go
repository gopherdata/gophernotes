package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-zeromq/zmq4"
)

const (
	failure = "\u2717"
	success = "\u2713"
)

const (
	connectionFile = "fixtures/connection_file.json"
	sessionID      = "ba65a05c-106a-4799-9a94-7f5631bbe216"
)

var (
	connectionKey string
	transport     string
	ip            string
	shellPort     int
	iopubPort     int
)

//==============================================================================

func TestMain(m *testing.M) {
	os.Exit(runTest(m))
}

// runTest initializes the environment for the tests and allows for
// the proper exit if the test fails or succeeds.
func runTest(m *testing.M) int {
	// Parse the connection info.
	var connInfo ConnectionInfo

	connData, err := ioutil.ReadFile(connectionFile)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(connData, &connInfo); err != nil {
		log.Fatal(err)
	}

	// Store the connection parameters globally for use by the test client.
	connectionKey = connInfo.Key
	transport = connInfo.Transport
	ip = connInfo.IP
	shellPort = connInfo.ShellPort
	iopubPort = connInfo.IOPubPort

	// Start the kernel.
	go runKernel(connectionFile)

	return m.Run()
}

//==============================================================================

// TestEvaluate tests the evaluation of consecutive cells.
func TestEvaluate(t *testing.T) {
	cases := []struct {
		Input  []string
		Output string
	}{
		{[]string{
			"a := 1",
			"a",
		}, "1"},
		{[]string{
			"a = 2",
			"a + 3",
		}, "5"},
		{[]string{
			"func myFunc(x int) int {",
			"    return x+1",
			"}",
			"myFunc(1)",
		}, "2"},
		{[]string{
			"b := myFunc(1)",
		}, ""},
		{[]string{
			"type Rect struct {",
			"    Width, Height int",
			"}",
			"Rect{10, 30}",
		}, "{10 30}"},
		{[]string{
			"type Rect struct {",
			"    Width, Height int",
			"}",
			"&Rect{10, 30}",
		}, "&{10 30}"},
		{[]string{
			"func a(b int) (int, int) {",
			"    return 2 + b, b",
			"}",
			"a(10)",
		}, "12 10"},
		{[]string{
			`import "errors"`,
			"func a() (interface{}, error) {",
			`    return nil, errors.New("To err is human")`,
			"}",
			"a()",
		}, "<nil> To err is human"},
		{[]string{
			`c := []string{"gophernotes", "is", "super", "bad"}`,
			"c[:3]",
		}, "[gophernotes is super]"},
		{[]string{
			"m := map[string]int{",
			`    "a": 10,`,
			`    "c": 30,`,
			"}",
			`m["c"]`,
		}, "30 true"},
		{[]string{
			"if 1 < 2 {",
			"    3",
			"}",
		}, ""},
		{[]string{
			"d := 10",
			"d++",
		}, ""},
		{[]string{
			"out := make(chan int)",
			"go func() {",
			"    out <- 123",
			"}()",
			"<-out",
		}, "123 true"},
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

	content, pub := client.executeCode(t, codeIn)

	status := getString(t, "content", content, "status")

	if status != "ok" {
		t.Fatalf("\t%s Execution encountered error [%s]: %s", failure, content["ename"], content["evalue"])
	}

	for _, pubMsg := range pub {
		if pubMsg.Header.MsgType == "execute_result" {
			content = getMsgContentAsJSONObject(t, pubMsg)

			bundledMIMEData := getJSONObject(t, "content", content, "data")
			textRep := getString(t, `content["data"]`, bundledMIMEData, "text/plain")

			return textRep
		}
	}

	return ""
}

// TestPanicGeneratesError tests that executing code with an un-recovered panic properly generates both
// an error "execute_reply" and publishes an "error" message.
func TestPanicGeneratesError(t *testing.T) {
	client, closeClient := newTestJupyterClient(t)
	defer closeClient()

	content, pub := client.executeCode(t, `panic("error")`)

	status := getString(t, "content", content, "status")

	if status != "error" {
		t.Fatalf("\t%s Execution did not raise expected error", failure)
	}

	var foundPublishedError bool
	for _, pubMsg := range pub {
		if pubMsg.Header.MsgType == "error" {
			foundPublishedError = true
			break
		}
	}

	if !foundPublishedError {
		t.Fatalf("\t%s Execution did not publish an expected \"error\" message", failure)
	}
}

// TestPrintStdout tests that data written to stdout publishes the same data in a "stdout" "stream" message.
func TestPrintStdout(t *testing.T) {
	cases := []struct {
		Input  []string
		Output []string
	}{
		{[]string{
			`import "fmt"`,
			"a := 1",
			"fmt.Println(a)",
		}, []string{"1\n"}},
		{[]string{
			"a = 2",
			"fmt.Print(a)",
		}, []string{"2"}},
		{[]string{
			`import "os"`,
			`os.Stdout.WriteString("3")`,
		}, []string{"3"}},
		{[]string{
			`fmt.Fprintf(os.Stdout, "%d\n", 4)`,
		}, []string{"4\n"}},
		{[]string{
			`import "time"`,
			"for i := 0; i < 3; i++ {",
			"    fmt.Println(i)",
			"    time.Sleep(500 * time.Millisecond)", // Stall to prevent prints from buffering into single message.
			"}",
		}, []string{"0\n", "1\n", "2\n"}},
	}

	t.Logf("Should produce stdout stream messages when writing to stdout")

cases:
	for k, tc := range cases {
		// Give a progress report.
		t.Logf("  Evaluating code snippet %d/%d.", k+1, len(cases))

		// Get the result.
		stdout, _ := testOutputStream(t, strings.Join(tc.Input, "\n"))

		// Compare the result.
		if len(stdout) != len(tc.Output) {
			t.Errorf("\t%s Test case expected %d message(s) on stdout but got %d.", failure, len(tc.Output), len(stdout))
			continue
		}
		for i, expected := range tc.Output {
			if stdout[i] != expected {
				t.Errorf("\t%s Test case returned unexpected messages on stdout.", failure)
				continue cases
			}
		}
		t.Logf("\t%s Returned the expected messages on stdout.", success)
	}
}

// TestPrintStderr tests that data written to stderr publishes the same data in a "stderr" "stream" message.
func TestPrintStderr(t *testing.T) {
	cases := []struct {
		Input  []string
		Output []string
	}{
		{[]string{
			`import "fmt"`,
			`import "os"`,
			"a := 1",
			"fmt.Fprintln(os.Stderr, a)",
		}, []string{"1\n"}},
		{[]string{
			`os.Stderr.WriteString("2")`,
		}, []string{"2"}},
		{[]string{
			`import "time"`,
			"for i := 0; i < 3; i++ {",
			"    fmt.Fprintln(os.Stderr, i)",
			"    time.Sleep(500 * time.Millisecond)", // Stall to prevent prints from buffering into single message.
			"}",
		}, []string{"0\n", "1\n", "2\n"}},
	}

	t.Logf("Should produce stderr stream messages when writing to stderr")

cases:
	for k, tc := range cases {
		// Give a progress report.
		t.Logf("  Evaluating code snippet %d/%d.", k+1, len(cases))

		// Get the result.
		_, stderr := testOutputStream(t, strings.Join(tc.Input, "\n"))

		// Compare the result.
		if len(stderr) != len(tc.Output) {
			t.Errorf("\t%s Test case expected %d message(s) on stderr but got %d.", failure, len(tc.Output), len(stderr))
			continue
		}
		for i, expected := range tc.Output {
			if stderr[i] != expected {
				t.Errorf("\t%s Test case returned unexpected messages on stderr.", failure)
				continue cases
			}
		}
		t.Logf("\t%s Returned the expected messages on stderr.", success)
	}
}

//==============================================================================

// testJupyterClient holds references to the 2 sockets it uses to communicate with the kernel.
type testJupyterClient struct {
	shellSocket zmq4.Socket
	ioSocket    zmq4.Socket
}

// newTestJupyterClient creates and connects a fresh client to the kernel. Upon error, newTestJupyterClient
// will Fail the test.
func newTestJupyterClient(t *testing.T) (testJupyterClient, func()) {
	t.Helper()

	var (
		err       error
		ctx       = context.Background()
		addrShell = fmt.Sprintf("%s://%s:%d", transport, ip, shellPort)
		addrIO    = fmt.Sprintf("%s://%s:%d", transport, ip, iopubPort)
	)

	// Prepare the shell socket.
	shell := zmq4.NewReq(ctx)
	if err = shell.Dial(addrShell); err != nil {
		t.Fatalf("\t%s shell.Connect: %s", failure, err)
	}

	// Prepare the IOPub socket.
	iopub := zmq4.NewSub(ctx)
	if err = iopub.Dial(addrIO); err != nil {
		t.Fatalf("\t%s iopub.Connect: %s", failure, err)
	}

	if err = iopub.SetOption(zmq4.OptionSubscribe, ""); err != nil {
		t.Fatalf("\t%s iopub.SetSubscribe: %s", failure, err)
	}

	// Wait for a second to give the tcp connection time to complete to avoid missing the early pub messages.
	time.Sleep(1 * time.Second)

	return testJupyterClient{shell, iopub}, func() {
		if err := shell.Close(); err != nil {
			t.Errorf("\t%s shell.Close: %s", failure, err)
		}
		if err = iopub.Close(); err != nil {
			t.Errorf("\t%s iopub.Close: %s", failure, err)
		}
	}
}

// sendShellRequest sends a message to the kernel over the shell channel. Upon error, sendShellRequest
// will Fail the test.
func (client *testJupyterClient) sendShellRequest(t *testing.T, request ComposedMsg) {
	t.Helper()

	var (
		frames [][]byte
		err    error
	)

	frames = append(frames, []byte("<IDS|MSG>"))

	reqMsgParts, err := request.ToWireMsg([]byte(connectionKey))
	if err != nil {
		t.Fatalf("\t%s request.ToWireMsg: %s", failure, err)
	}
	frames = append(frames, reqMsgParts...)

	if err = client.shellSocket.SendMulti(zmq4.NewMsgFrom(frames...)); err != nil {
		t.Fatalf("\t%s shellSocket.SendMessage: %s", failure, err)
	}
}

// recvShellReply tries to read a reply message from the shell channel. It will timeout after the given
// timeout delay. Upon error or timeout, recvShellReply will Fail the test.
func (client *testJupyterClient) recvShellReply(t *testing.T, timeout time.Duration) ComposedMsg {
	t.Helper()

	ch := make(chan ComposedMsg)

	go func() {
		repMsgParts, err := client.shellSocket.Recv()
		if err != nil {
			t.Fatalf("\t%s Shell socket RecvMessageBytes: %s", failure, err)
		}

		msgParsed, _, err := WireMsgToComposedMsg(repMsgParts.Frames, []byte(connectionKey))
		if err != nil {
			t.Fatalf("\t%s Could not parse wire message: %s", failure, err)
		}

		ch <- msgParsed
	}()

	var reply ComposedMsg

	select {
	case reply = <-ch:
		return reply
	case <-time.After(timeout):
		t.Fatalf("\t%s recvShellReply timed out", failure)
	}

	return reply
}

// recvIOSub tries to read a published message from the IOPub channel. It will timeout after the given
// timeout delay. Upon error or timeout, recvIOSub will Fail the test.
func (client *testJupyterClient) recvIOSub(t *testing.T, timeout time.Duration) ComposedMsg {
	t.Helper()

	ch := make(chan ComposedMsg)

	go func() {
		repMsgParts, err := client.ioSocket.Recv()
		if err != nil {
			t.Fatalf("\t%s IOPub socket RecvMessageBytes: %s", failure, err)
		}

		msgParsed, _, err := WireMsgToComposedMsg(repMsgParts.Frames, []byte(connectionKey))
		if err != nil {
			t.Fatalf("\t%s Could not parse wire message: %s", failure, err)
		}

		ch <- msgParsed
	}()

	var sub ComposedMsg
	select {
	case sub = <-ch:
	case <-time.After(timeout):
		t.Fatalf("\t%s recvIOSub timed out", failure)
	}

	return sub
}

// performJupyterRequest preforms a request and awaits a reply on the shell channel. Additionally all messages on the
// IOPub channel between the opening 'busy' messages and closing 'idle' message are captured and returned. The request
// will timeout after the given timeout delay. Upon error or timeout, request will Fail the test.
func (client *testJupyterClient) performJupyterRequest(t *testing.T, request ComposedMsg, timeout time.Duration) (ComposedMsg, []ComposedMsg) {
	t.Helper()

	client.sendShellRequest(t, request)
	reply := client.recvShellReply(t, timeout)

	// Read the expected 'busy' message and ensure it is in fact, a 'busy' message.
	subMsg := client.recvIOSub(t, 1*time.Second)
	assertMsgTypeEquals(t, subMsg, "status")

	subData := getMsgContentAsJSONObject(t, subMsg)
	execState := getString(t, "content", subData, "execution_state")

	if execState != kernelBusy {
		t.Fatalf("\t%s Expected a 'busy' status message but got '%s'", failure, execState)
	}

	var pub []ComposedMsg

	// Read messages from the IOPub channel until an 'idle' message is received.
	for {
		subMsg = client.recvIOSub(t, 100*time.Millisecond)

		// If the message is a 'status' message, ensure it is an 'idle' status.
		if subMsg.Header.MsgType == "status" {
			subData = getMsgContentAsJSONObject(t, subMsg)
			execState = getString(t, "content", subData, "execution_state")

			if execState != kernelIdle {
				t.Fatalf("\t%s Expected a 'idle' status message but got '%s'", failure, execState)
			}

			// Break from the loop as we don't expect any other IOPub messages after the 'idle'.
			break
		}

		// Add the message to the pub collection.
		pub = append(pub, subMsg)
	}

	return reply, pub
}

// executeCode creates an execute request for the given code and preforms the request. It returns the content of the
// reply as well as all of the messages captured from the IOPub channel during the execution.
func (client *testJupyterClient) executeCode(t *testing.T, code string) (map[string]interface{}, []ComposedMsg) {
	t.Helper()

	// Create a message.
	request, err := NewMsg("execute_request", ComposedMsg{})
	if err != nil {
		t.Fatalf("\t%s NewMsg: %s", failure, err)
	}

	// Fill in remaining header information.
	request.Header.Session = sessionID
	request.Header.Username = "KernelTester"

	// Fill in Metadata.
	request.Metadata = make(map[string]interface{})

	// Fill in content.
	content := make(map[string]interface{})
	content["code"] = code
	content["silent"] = false
	request.Content = content

	// Make the request.
	reply, pub := client.performJupyterRequest(t, request, 10*time.Second)

	// Ensure the reply is an execute_reply and extract the content from the reply.
	assertMsgTypeEquals(t, reply, "execute_reply")
	content = getMsgContentAsJSONObject(t, reply)

	return content, pub
}

// assertMsgTypeEquals is a test helper that fails the test if the message header's MsgType is not the
// expectedType.
func assertMsgTypeEquals(t *testing.T, msg ComposedMsg, expectedType string) {
	t.Helper()

	if msg.Header.MsgType != expectedType {
		t.Fatalf("\t%s Expected message of type '%s' but was '%s'", failure, expectedType, msg.Header.MsgType)
	}
}

// getMsgContentAsJSONObject is a test helper that fails the rest if the message content is not a
// map[string]interface{} and returns the content as a map[string]interface{} if it is of the correct type.
func getMsgContentAsJSONObject(t *testing.T, msg ComposedMsg) map[string]interface{} {
	t.Helper()

	content, ok := msg.Content.(map[string]interface{})
	if !ok {
		t.Fatalf("\t%s Message content is not a JSON object", failure)
	}

	return content
}

// getString is a test helper that retrieves a value as a string from the content at the given key. If the key
// does not exist in the content map or the value is not a string this will fail the test. The jsonObjectName
// parameter is a string used to name the content for more helpful fail messages.
func getString(t *testing.T, jsonObjectName string, content map[string]interface{}, key string) string {
	t.Helper()

	raw, ok := content[key]
	if !ok {
		t.Fatalf("\t%s %s[\"%s\"] field not present", failure, jsonObjectName, key)
	}

	value, ok := raw.(string)
	if !ok {
		t.Fatalf("\t%s %s[\"%s\"] is not a string", failure, jsonObjectName, key)
	}

	return value
}

// getJSONObject is a test helper that retrieves a value as a map[string]interface{} from the content at the given key.
// If the key  does not exist in the content map or the value is not a map[string]interface{} this will fail the test.
// The jsonObjectName parameter is a string used to name the content for more helpful fail messages.
func getJSONObject(t *testing.T, jsonObjectName string, content map[string]interface{}, key string) map[string]interface{} {
	t.Helper()

	raw, ok := content[key]
	if !ok {
		t.Fatalf("\t%s %s[\"%s\"] field not present", failure, jsonObjectName, key)
	}

	value, ok := raw.(map[string]interface{})
	if !ok {
		t.Fatalf("\t%s %s[\"%s\"] is not a JSON object", failure, jsonObjectName, key)
	}

	return value
}

// testOutputStream is a test helper that collects "stream" messages upon executing the codeIn.
func testOutputStream(t *testing.T, codeIn string) ([]string, []string) {
	t.Helper()

	client, closeClient := newTestJupyterClient(t)
	defer closeClient()

	_, pub := client.executeCode(t, codeIn)

	var stdout, stderr []string
	for _, pubMsg := range pub {
		if pubMsg.Header.MsgType == "stream" {
			content := getMsgContentAsJSONObject(t, pubMsg)
			streamType := getString(t, "content", content, "name")
			streamData := getString(t, "content", content, "text")

			switch streamType {
			case StreamStdout:
				stdout = append(stdout, streamData)
			case StreamStderr:
				stderr = append(stderr, streamData)
			default:
				t.Fatalf("Unknown stream type '%s'", streamType)
			}
		}
	}

	return stdout, stderr
}
