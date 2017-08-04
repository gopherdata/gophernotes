// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package token defines constants representing the lexical tokens of the Go
// programming language and basic operations on tokens (printing, predicates).
//
package token

import (
	"go/token"
	base "go/token"
)

const (
	QUOTE base.Token = (base.VAR+127)&^127 + iota
	QUASIQUOTE
	UNQUOTE
	UNQUOTE_SPLICE
	SPLICE
	MACRO
	FUNCTION
	LAMBDA
	TYPECASE
)

var tokens map[base.Token]string

var keywords map[string]base.Token

func init() {
	tokens = map[base.Token]string{
		SPLICE:         "~splice",
		QUOTE:          "~quote",
		QUASIQUOTE:     "~quasiquote",
		UNQUOTE:        "~unquote",
		UNQUOTE_SPLICE: "~unquote_splice",
		MACRO:          "~macro",
		FUNCTION:       "~func",
		LAMBDA:         "~lambda",
		TYPECASE:       "~typecase",
	}

	keywords = make(map[string]base.Token)
	for k, v := range tokens {
		keywords[v[1:]] = k // skip ~ in lookup table
	}
}

// Lookup maps a identifier to its keyword token.
func Lookup(lit string) base.Token {
	if lit == "macro" {
		// allow the spelling "macro" because "~macro" is really ugly in source code...
		// especially when writing :~macro
		return MACRO
	}
	return token.Lookup(lit)
}

// LookupSpecial maps a identifier starting with '~' to its keyword token.
func LookupSpecial(lit string) base.Token {
	tok, _ := keywords[lit]
	return tok
}

func String(tok base.Token) string {
	if str, ok := tokens[tok]; ok {
		return str
	}
	return tok.String()
}

// Predicates

// IsLiteral returns true for tokens corresponding to identifiers
// and basic type literals; it returns false otherwise.
//
func IsLiteral(tok base.Token) bool {
	return tok.IsLiteral()
}

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
//
func IsOperator(tok base.Token) bool {
	return tok.IsOperator()
}

// IsKeyword returns true for tokens corresponding to keywords;
// it returns false otherwise.
//
func IsKeyword(tok base.Token) bool {
	return tok.IsKeyword()
}

// IsMacroKeyword returns true for tokens corresponding to macro-related keywords;
// it returns false otherwise.
//
func IsMacroKeyword(tok base.Token) bool {
	_, ok := tokens[tok]
	return ok
}
