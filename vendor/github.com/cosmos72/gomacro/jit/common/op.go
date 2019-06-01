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
 * op.go
 *
 *  Created on Feb 11, 2019
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"fmt"
	"go/token"
)

// ============================================================================
// no-arg instruction
type Op0 uint8

const (
	BAD = Op0(token.ILLEGAL)   // invalid instruction, guaranteed to signal exception
	NOP = Op0(token.SEMICOLON) // somewhat arbitrary choice
	RET = Op0(token.RETURN)
)

var op0Name = map[Op0]string{
	BAD: "BAD",
	RET: "RET",
	NOP: "NOP",
}

func (op Op0) String() string {
	s, ok := op0Name[op]
	if !ok {
		s = fmt.Sprintf("Op0(%d)", uint8(op))
	}
	return s
}

// implement AsmCode interface
func (op Op0) asmcode() {
}

// ============================================================================
// one-arg instruction
type Op1 uint8

const (
	ZERO = Op1(token.DEFAULT) // somewhat arbitrary choice
	INC  = Op1(token.INC)     // ++
	DEC  = Op1(token.DEC)     // --
	NEG1 = Op1(token.VAR + 1) // - // avoid conflict between NEG2 and SUB2
	NOT1 = Op1(token.VAR + 2) // ^ // avoid conflict between NOT2 and XOR2
	JMP  = Op1(token.GOTO)
)

var op1Name = map[Op1]string{
	ZERO: "ZERO",
	INC:  "INC",
	DEC:  "DEC",
	NOT1: "NOT1",
	NEG1: "NEG1",
	/* JMP:
	 * if argument is signed integer constant N
	 *    => jump forward or backward N assembly instructions
	 *       (each instruction includes its arguments)
	 * if argument is pointer constant
	 *    => jump to absolute address
	 */
	JMP: "JMP",
}

func (op Op1) String() string {
	s, ok := op1Name[op]
	if !ok {
		s = fmt.Sprintf("Op1(%d)", uint8(op))
	}
	return s
}

// implement AsmCode interface
func (op Op1) asmcode() {
}

// ============================================================================
// two-arg instruction
type Op2 uint8

const (
	ADD2 = Op2(token.ADD)
	SUB2 = Op2(token.SUB)
	ADC2 = Op2(token.ADD + token.VAR) // add with carry
	SBB2 = Op2(token.SUB + token.VAR) // subtract with borrow
	MUL2 = Op2(token.MUL)
	DIV2 = Op2(token.QUO) // divide
	QUO2 = DIV2           // alias for DIV
	REM2 = Op2(token.REM) // remainder

	AND2     = Op2(token.AND)
	OR2      = Op2(token.OR)
	XOR2     = Op2(token.XOR)
	SHL2     = Op2(token.SHL)
	SHR2     = Op2(token.SHR)
	AND_NOT2 = Op2(token.AND_NOT)
	LAND2    = Op2(token.LAND) // &&
	LOR2     = Op2(token.LOR)  // ||

	MOV  = Op2(token.ASSIGN) // =
	CAST = Op2(token.TYPE)   // somewhat arbitrary choice

	LEA2 = Op2(token.ARROW) // amd64 only. somewhat arbitrary choice
	// XCHG = ??
	// two-arg versions of NOT1, NEG1 above
	NEG2  = Op2(NEG1)
	NOT2  = Op2(NOT1)
	JMPIF = Op2(JMP)
)

var op2Name = map[Op2]string{
	ADD2: "ADD2",
	SUB2: "SUB2",
	ADC2: "ADC2",
	SBB2: "SBB2",
	MUL2: "MUL2",
	DIV2: "DIV2",
	REM2: "REM2",

	AND2:     "AND2",
	OR2:      "OR2",
	XOR2:     "XOR2",
	SHL2:     "SHL2",
	SHR2:     "SHR2",
	AND_NOT2: "AND_NOT2",
	LAND2:    "LAND2",
	LOR2:     "LOR2",

	MOV:  "MOV",
	CAST: "CAST",
	LEA2: "LEA2",
	// XCHG: "XCHG",
	NEG2:  "NEG2",
	NOT2:  "NOT2",
	JMPIF: "JMPIF",
}

func (op Op2) String() string {
	s, ok := op2Name[op]
	if !ok {
		s = fmt.Sprintf("Op2(%d)", int(op))
	}
	return s
}

// implement AsmCode interface
func (op Op2) asmcode() {
}

// ============================================================================
// three-arg instruction
type Op3 uint8

const (
	ADD3 = Op3(ADD2)
	SUB3 = Op3(SUB2)
	ADC3 = Op3(ADC2)
	SBB3 = Op3(SBB2)
	MUL3 = Op3(MUL2)
	DIV3 = Op3(DIV2)
	REM3 = Op3(REM2)

	AND3     = Op3(AND2)
	OR3      = Op3(OR2)
	XOR3     = Op3(XOR2)
	SHL3     = Op3(SHL2)
	SHR3     = Op3(SHR2)
	AND_NOT3 = Op3(AND_NOT2)
	LAND3    = Op3(LAND2)
	LOR3     = Op3(LOR2)

	GETIDX = Op3(token.LBRACK) // a[b] -> val
	SETIDX = Op3(token.RBRACK) // a[b] <- val
)

var op3Name = map[Op3]string{
	ADD3: "ADD3",
	SUB3: "SUB3",
	SBB3: "SBB3",
	ADC3: "ADC3",
	MUL3: "MUL3",
	DIV3: "DIV3",
	REM3: "REM3",

	AND3:     "AND3",
	OR3:      "OR3",
	XOR3:     "XOR3",
	SHL3:     "SHL3",
	SHR3:     "SHR3",
	AND_NOT3: "AND_NOT3",
	LAND3:    "LAND3",
	LOR3:     "LOR3",

	GETIDX: "GETIDX",
	SETIDX: "SETIDX",
}

func (op Op3) String() string {
	s, ok := op3Name[op]
	if !ok {
		s = fmt.Sprintf("Op3(%d)", int(op))
	}
	return s
}

// implement AsmCode interface
func (op Op3) asmcode() {
}

// ============================================================================
// four-arg instruction
type Op4 uint8

const (
	LEA4 = Op4(LEA2) // amd64 only
)

var op4Name = map[Op4]string{
	LEA4: "LEA4",
}

func (op Op4) String() string {
	s, ok := op4Name[op]
	if !ok {
		s = fmt.Sprintf("Op4(%d)", int(op))
	}
	return s
}

// implement AsmCode interface
func (op Op4) asmcode() {
}
