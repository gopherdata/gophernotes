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
 * op2.go
 *
 *  Created on Feb 16, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"fmt"
	"go/token"

	"github.com/cosmos72/gomacro/jit/common"
)

type Op2 uint8 // binary expression operator

const (
	ADD     = Op2(common.ADD3)
	SUB     = Op2(common.SUB3)
	MUL     = Op2(common.MUL3)
	DIV     = Op2(common.DIV3)
	QUO     = DIV
	REM     = Op2(common.REM3)
	AND     = Op2(common.AND3)
	OR      = Op2(common.OR3)
	XOR     = Op2(common.XOR3)
	SHL     = Op2(common.SHL3)
	SHR     = Op2(common.SHR3)
	AND_NOT = Op2(common.AND_NOT3) // &^
	LAND    = Op2(common.LAND3)    // &&
	LOR     = Op2(common.LOR3)     // ||
	/*
		EQL     = Op2(common.EQL3)
		LSS     = Op2(common.LSS3)
		GTR     = Op2(common.GTR3)
		NEQ     = Op2(common.NEQ3)
		LEQ     = Op2(common.LEQ3)
		GEQ     = Op2(common.GEQ3)
	*/
	IDX = Op2(common.GETIDX) // a[b]
)

var op2name = map[Op2]string{
	ADD:     "+",
	SUB:     "-",
	MUL:     "*",
	QUO:     "/",
	REM:     "%",
	AND:     "&",
	OR:      "|",
	XOR:     "^",
	SHL:     "<<",
	SHR:     ">>",
	AND_NOT: "&^",
	LAND:    "&&",
	LOR:     "||",
	/*
		EQL    :"==",
		LSS    :"<",
		GTR    :"<",
		NEQ    :"!=",
		LEQ    :"<=",
		GEQ    :">=",
	*/
	IDX: "[]",
}

// =======================================================

func (op Op2) Valid() bool {
	_, ok := op2name[op]
	return ok
}

func (op Op2) Validate() {
	if !op.Valid() {
		errorf("unknown Op2: %v", op)
	}
}

// convert to common.Op3
func (op Op2) Asm() common.Op3 {
	op.Validate()
	return common.Op3(op)
}

var tokenNoOp2 = fmt.Errorf("failed to convert token.Token to jit.Op2")

// convert token.Token to Op2
func TokenOp2(tok token.Token) (Op2, error) {
	op := Op2(tok)
	if !op.Valid() {
		return 0, tokenNoOp2
	}
	return op, nil
}

func (op Op2) String() string {
	s, ok := op2name[op]
	if !ok {
		s = fmt.Sprintf("Op2(%d)", uint8(op))
	}
	return s
}

func (op Op2) IsCommutative() bool {
	return op.Asm().IsCommutative()
}
