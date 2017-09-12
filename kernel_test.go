package main

import (
	"errors"
	"os"
	"sync"
	"testing"
	"time"

	zmq "github.com/pebbe/zmq4"
)

const (
	failure = "\u2717"
	success = "\u2713"
)

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
		Input  string
		Output string
	}{
		{"import \"fmt\"\na := 1\nfmt.Println(a)", "1\n"},
		{"a = 2\nfmt.Println(a)", "2\n"},
		{"func myFunc(x int) int {\nreturn x+1\n}\nfmt.Println(\"func defined\")", "func dfined\n"},
		{"b := myFunc(1)\nfmt.Println(b)", "2\n"},
	}

	t.Logf("Should be able to evaluate valid code in notebook cells.")

	for k, tc := range cases {

		// Give a progress report.
		t.Logf("  Evaluating code snippet %d/%d.", k+1, len(cases))

		// Get the result.
		result := testEvaluate(t, tc.Input, k)

		// Compare the result.
		if result != tc.Output {
			t.Errorf("\t%s Test case produced unexpected results.", failure)
			continue
		}
		t.Logf("\t%s Should return the correct cell output.", success)
	}
}

// testEvaluate evaluates a cell.
func testEvaluate(t *testing.T, codeIn string, testCaseIndex int) string {

	// Define the shell socket.
	addrShell := "tcp://127.0.0.1:57503"
	addrIO := "tcp://127.0.0.1:40885"

	// Create a message.
	msg, err := NewMsg("execute_request", ComposedMsg{})
	if err != nil {
		t.Fatal("Create New Message:", err)
	}

	// Fill in remaining header information.
	msg.Header.Session = "ba65a05c-106a-4799-9a94-7f5631bbe216"
	msg.Header.Username = "blah"

	// Fill in Metadata.
	msg.Metadata = make(map[string]interface{})

	// Fill in content.
	content := make(map[string]interface{})
	content["code"] = codeIn
	content["silent"] = false
	msg.Content = content

	// Prepare the shell socket.
	sock, err := zmq.NewSocket(zmq.REQ)
	if err != nil {
		t.Fatal("NewSocket:", err)
	}
	defer sock.Close()

	if err = sock.Connect(addrShell); err != nil {
		t.Fatal("sock.Connect:", err)
	}

	// Prepare the IOPub subscriber.
	sockIO, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		t.Fatal("NewSocket:", err)
	}
	defer sockIO.Close()

	if err = sockIO.Connect(addrIO); err != nil {
		t.Fatal("sockIO.Connect:", err)
	}

	sockIO.SetSubscribe("")

	// Start the subscriber.
	quit := make(chan struct{})
	var result string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {

			case <-quit:
				return

			default:
				msgParts, err := sockIO.RecvMessageBytes(0)
				if err != nil {
					t.Fatal("sockIO.RecvMessageBytes:", err)
				}

				msgParsed, _, err := WireMsgToComposedMsg(msgParts, []byte("a0436f6c-1916-498b-8eb9-e81ab9368e84"))
				if err != nil {
					t.Fatal("WireMsgToComposedMsg:", err)
				}

				if msgParsed.Header.MsgType == "execute_result" {
					content, ok := msgParsed.Content.(map[string]interface{})
					if !ok {
						t.Fatal("msgParsed.Content.(map[string]interface{})", errors.New("Could not cast type"))
					}
					data, ok := content["data"]
					if !ok {
						t.Fatal("content[\"data\"]", errors.New("Data field not present"))
					}
					dataMap, ok := data.(map[string]interface{})
					if !ok {
						t.Fatal("data.(map[string]string)", errors.New("Could not cast type"))
					}
					rawResult, ok := dataMap["text/plain"]
					if !ok {
						t.Fatal("dataMap[\"text/plain\"]", errors.New("text/plain field not present"))
					}
					result, ok = rawResult.(string)
					if !ok {
						t.Fatal("rawResult.(string)", errors.New("Could not cast result as string"))
					}
					return
				}
			}
		}
	}()

	time.Sleep(1 * time.Second)

	// Send the execute request.
	if _, err := sock.Send("<IDS|MSG>", zmq.SNDMORE); err != nil {
		t.Fatal("sock.Send:", err)
	}

	msgParts, err := msg.ToWireMsg([]byte("a0436f6c-1916-498b-8eb9-e81ab9368e84"))
	if err != nil {
		t.Fatal("msg.ToWireMsg:", err)
	}

	if _, err = sock.SendMessage(msgParts); err != nil {
		t.Fatal("sock.SendMessage:", err)
	}

	// Wait for the result.  If we timeout, kill the subscriber.
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	// Compare the result to the expect and clean up.
	select {
	case <-done:
		return result
	case <-time.After(10 * time.Second):
		close(quit)
		t.Fatalf("[test case %d] Evaution timed out!", testCaseIndex+1)
	}

	return ""
}
