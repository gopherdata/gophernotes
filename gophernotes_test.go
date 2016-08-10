package main

import (
	"testing"

	repl "github.com/gopherds/gophernotes/internal/repl"
	"github.com/stretchr/testify/assert"
)

// noError is a helper function for testing
func noError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// TestStub is a canary test to make sure testing and testify are good to go
func TestStub(t *testing.T) {
	assert.True(t, true, "This is good. Canary test passing")
}

// TestRun_import tests importing and use of a package
func TestRun_import(t *testing.T) {
	s, err := repl.NewSession()
	noError(t, err)

	codes := []string{
		":import encoding/json",
		"b, err := json.Marshal(nil)",
		"string(b)",
	}

	for _, code := range codes {
		_, _, err := s.Eval(code)
		noError(t, err)
	}
}

// TestRun_QuickFix_evaluated_but_not_used makes sure errors are not thrown for
// evaluations of variables that aren't used
func TestRun_QuickFix_evaluated_but_not_used(t *testing.T) {
	s, err := repl.NewSession()
	noError(t, err)

	codes := []string{
		`[]byte("")`,
		`make([]int, 0)`,
		`1+1`,
		`func() {}`,
		`(4 & (1 << 1))`,
		`1`,
	}

	for _, code := range codes {
		_, _, err := s.Eval(code)
		noError(t, err)
	}
}

// TestRun_QuickFix_used_as_value tests assignment of values to variables
// and subsequent use
func TestRun_QuickFix_used_as_value(t *testing.T) {
	s, err := repl.NewSession()
	noError(t, err)

	codes := []string{
		`:import log`,
		`a := 1`,
		`log.SetPrefix("")`,
	}

	for _, code := range codes {
		_, _, err := s.Eval(code)
		noError(t, err)
	}
}

// TestRun_Copy tests a copy within the replpkg
func TestRun_Copy(t *testing.T) {
	s, err := repl.NewSession()
	noError(t, err)

	codes := []string{
		`a := []string{"hello", "world"}`,
		`b := []string{"goodbye", "world"}`,
		`copy(a, b)`,
		`if (a[0] != "goodbye") {
			panic("should be copied")
		}`,
	}

	for _, code := range codes {
		_, _, err := s.Eval(code)
		noError(t, err)
	}
}

// TestRun_Const tests a constant within the replpkg
func TestRun_Const(t *testing.T) {
	s, err := repl.NewSession()
	noError(t, err)

	codes := []string{
		`const ( a = iota; b )`,
		`a`,
		`b`,
	}

	for _, code := range codes {
		_, _, err := s.Eval(code)
		noError(t, err)
	}
}
