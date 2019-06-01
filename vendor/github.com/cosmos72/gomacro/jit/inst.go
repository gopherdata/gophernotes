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
 * inst.go
 *
 *  Created on Feb 15, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"fmt"
	"go/token"

	"github.com/cosmos72/gomacro/jit/asm"
)

type Inst1 uint8 // unary statement operator
type Inst2 uint8 // binary statement operator
type Inst3 uint8 // ternary statement operator

type Inst1Misc uint8 // miscellaneous statement operator

const (
	INC  = Inst1(asm.INC)  // ++
	DEC  = Inst1(asm.DEC)  // --
	ZERO = Inst1(asm.ZERO) // = 0
	NOP  = Inst1(asm.NOP)  // used to wrap an expression into a statement

	ASSIGN         = Inst2(asm.MOV)
	ADD_ASSIGN     = Inst2(asm.ADD2)
	SUB_ASSIGN     = Inst2(asm.SUB2)
	MUL_ASSIGN     = Inst2(asm.MUL2)
	QUO_ASSIGN     = Inst2(asm.DIV2)
	REM_ASSIGN     = Inst2(asm.REM2)
	AND_ASSIGN     = Inst2(asm.AND2)
	OR_ASSIGN      = Inst2(asm.OR2)
	XOR_ASSIGN     = Inst2(asm.XOR2)
	SHL_ASSIGN     = Inst2(asm.SHL2)
	SHR_ASSIGN     = Inst2(asm.SHR2)
	AND_NOT_ASSIGN = Inst2(asm.AND_NOT2)
	LAND_ASSIGN    = Inst2(asm.LAND2)
	LOR_ASSIGN     = Inst2(asm.LOR2)

	IDX_ASSIGN = Inst3(asm.SETIDX) // a[b] = val

	// allocate / free soft register
	ALLOC = Inst1Misc(asm.ALLOC)
	FREE  = Inst1Misc(asm.FREE)
)

var inst1name = map[Inst1]string{
	INC:  "++",
	DEC:  "--",
	ZERO: " = 0",
	NOP:  "",
}

var inst2name = map[Inst2]string{
	ASSIGN:         "=",
	ADD_ASSIGN:     "+=",
	SUB_ASSIGN:     "-=",
	MUL_ASSIGN:     "*=",
	QUO_ASSIGN:     "/=",
	REM_ASSIGN:     "%=",
	AND_ASSIGN:     "&=",
	OR_ASSIGN:      "|=",
	XOR_ASSIGN:     "^=",
	SHL_ASSIGN:     "<<=",
	SHR_ASSIGN:     ">>=",
	AND_NOT_ASSIGN: "&^=",
	LAND_ASSIGN:    "&&=",
	LOR_ASSIGN:     "||=",
}

var tokenToInst2 = map[token.Token]Inst2{
	token.ASSIGN:         ASSIGN,
	token.ADD_ASSIGN:     ADD_ASSIGN,
	token.SUB_ASSIGN:     SUB_ASSIGN,
	token.MUL_ASSIGN:     MUL_ASSIGN,
	token.QUO_ASSIGN:     QUO_ASSIGN,
	token.REM_ASSIGN:     REM_ASSIGN,
	token.AND_ASSIGN:     AND_ASSIGN,
	token.OR_ASSIGN:      OR_ASSIGN,
	token.XOR_ASSIGN:     XOR_ASSIGN,
	token.SHL_ASSIGN:     SHL_ASSIGN,
	token.SHR_ASSIGN:     SHR_ASSIGN,
	token.AND_NOT_ASSIGN: AND_NOT_ASSIGN,
}

var inst3name = map[Inst3]string{
	IDX_ASSIGN: "[]=",
}

var misc1name = map[Inst1Misc]string{
	ALLOC: "ALLOC",
	FREE:  "FREE",
}

// =======================================================

func (inst Inst1) Valid() bool {
	_, ok := inst1name[inst]
	return ok
}

func (inst Inst1) Validate() {
	if !inst.Valid() {
		errorf("unknown Inst1: %v", inst)
	}
}

// convert to asm.Op1
func (inst Inst1) Asm() asm.Op1 {
	inst.Validate()
	return asm.Op1(inst)
}

func (inst Inst1) String() string {
	s, ok := inst1name[inst]
	if !ok {
		s = fmt.Sprintf("Inst1(%d)", uint8(inst))
	}
	return s
}

// =======================================================

func (inst Inst2) Valid() bool {
	_, ok := inst2name[inst]
	return ok
}

func (inst Inst2) Validate() {
	if !inst.Valid() {
		errorf("unknown Inst2: %v", inst)
	}
}

// convert to asm.Op2
func (inst Inst2) Asm() asm.Op2 {
	inst.Validate()
	return asm.Op2(inst)
}

var tokenNoInst2 = fmt.Errorf("failed to convert token.Token to jit.Inst2")

// convert token.Token to Inst2
func TokenInst2(tok token.Token) (Inst2, error) {
	inst, ok := tokenToInst2[tok]
	if !ok {
		return 0, tokenNoInst2
	}
	return inst, nil
}
func (inst Inst2) String() string {
	s, ok := inst2name[inst]
	if !ok {
		s = fmt.Sprintf("Inst2(%d)", uint8(inst))
	}
	return s
}

// =======================================================

func (inst Inst3) Valid() bool {
	_, ok := inst3name[inst]
	return ok
}

func (inst Inst3) Validate() {
	if !inst.Valid() {
		errorf("unknown Inst3: %v", inst)
	}
}

// convert to asm.Op3
func (inst Inst3) Asm() asm.Op3 {
	inst.Validate()
	return asm.Op3(inst)
}

func (inst Inst3) String() string {
	s, ok := inst3name[inst]
	if !ok {
		s = fmt.Sprintf("Inst3(%d)", uint8(inst))
	}
	return s
}

// =======================================================

func (inst Inst1Misc) Valid() bool {
	_, ok := misc1name[inst]
	return ok
}

func (inst Inst1Misc) Validate() {
	if !inst.Valid() {
		errorf("unknown Inst1Misc: %v", inst)
	}
}

// convert to asm.Op1Misc
func (inst Inst1Misc) Asm() asm.Op1Misc {
	inst.Validate()
	return asm.Op1Misc(inst)
}

func (inst Inst1Misc) String() string {
	s, ok := misc1name[inst]
	if !ok {
		s = fmt.Sprintf("Inst1Misc(%d)", uint8(inst))
	}
	return s
}
