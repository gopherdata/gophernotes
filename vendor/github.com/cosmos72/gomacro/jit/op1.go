/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * op1.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"fmt"
	"go/token"
	"reflect"

	"github.com/cosmos72/gomacro/jit/common"
)

type Op1 uint8 // unary expression operator

const (
	// instead of a single CAST, we must implement
	// one Op1 per destination type:
	// INT8 ... INT64, UINT8 ... UINT64, etc.
	INT     = Op1(common.Int)
	INT8    = Op1(common.Int8)
	INT16   = Op1(common.Int16)
	INT32   = Op1(common.Int32)
	INT64   = Op1(common.Int64)
	UINT    = Op1(common.Uint)
	UINT8   = Op1(common.Uint8)
	UINT16  = Op1(common.Uint16)
	UINT32  = Op1(common.Uint32)
	UINT64  = Op1(common.Uint64)
	UINTPTR = Op1(common.Uintptr)
	FLOAT32 = Op1(common.Float32)
	FLOAT64 = Op1(common.Float64)
	PTR     = Op1(common.Ptr)
	NEG     = Op1(common.NEG2)
	NOT     = Op1(common.NOT2)
)

var op1name = map[Op1]string{
	INT:     "int",
	INT8:    "int8",
	INT16:   "int16",
	INT32:   "int32",
	INT64:   "int64",
	UINT:    "uint",
	UINT8:   "uint8",
	UINT16:  "uint16",
	UINT32:  "uint32",
	UINT64:  "uint64",
	FLOAT32: "float32",
	FLOAT64: "float64",
	PTR:     "pointer",
	NEG:     "-",
	NOT:     "^",
}

// =======================================================

func (op Op1) Valid() bool {
	_, ok := op1name[op]
	return ok
}

func (op Op1) Validate() {
	if !op.Valid() {
		errorf("unknown Op1: %v", op)
	}
}

func (op Op1) IsCast() bool {
	return op.Valid() && op >= INT && op <= PTR
}

// convert to common.Op2
func (op Op1) Asm() common.Op2 {
	op.Validate()
	if op.IsCast() {
		return common.CAST
	}
	return common.Op2(op)
}

var kindNoOp1 = fmt.Errorf("failed to convert reflect.Kind to jit.Op1")

// convert reflect.Kind to Op1
// used to create cast expressions
func KindOp1(kind reflect.Kind) (Op1, error) {
	if kind >= reflect.Int && kind <= reflect.Complex128 {
		op := Op1(kind)
		if op.Valid() {
			return op, nil
		}
	}
	return 0, kindNoOp1
}

var tokenNoOp1 = fmt.Errorf("failed to convert token.Token to jit.Op1")

// convert token.Token to Op1
func TokenOp1(tok token.Token) (Op1, error) {
	switch tok {
	case token.SUB:
		return NEG, nil
	case token.XOR:
		return NOT, nil
	default:
		return 0, tokenNoOp1
	}
}

func (op Op1) String() string {
	s, ok := op1name[op]
	if !ok {
		s = fmt.Sprintf("Op1(%d)", uint8(op))
	}
	return s
}
