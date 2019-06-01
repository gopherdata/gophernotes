/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * machine.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

// ============================================================================
// register
const (
	noregid = RegId(AMD64-1)<<8 + iota
	RAX
	RCX
	RDX
	RBX
	RSP
	RBP
	RSI
	RDI
	R8
	R9
	R10
	R11
	R12
	R13
	R14
	R15
	RLo RegId = RAX
	RHi RegId = R15
	// suggested register to point to local variables
	RVAR = RSI
)

var regName1 = [...]string{
	RAX - RLo: "%al",
	RCX - RLo: "%cl",
	RDX - RLo: "%dl",
	RBX - RLo: "%bl",
	RSP - RLo: "%spl",
	RBP - RLo: "%bpl",
	RSI - RLo: "%sil",
	RDI - RLo: "%dil",
	R8 - RLo:  "%r8b",
	R9 - RLo:  "%r9b",
	R10 - RLo: "%r10b",
	R11 - RLo: "%r11b",
	R12 - RLo: "%r12b",
	R13 - RLo: "%r13b",
	R14 - RLo: "%r14b",
	R15 - RLo: "%r15b",
}
var regName2 = [...]string{
	RAX - RLo: "%ax",
	RCX - RLo: "%cx",
	RDX - RLo: "%dx",
	RBX - RLo: "%bx",
	RSP - RLo: "%sp",
	RBP - RLo: "%bp",
	RSI - RLo: "%si",
	RDI - RLo: "%di",
	R8 - RLo:  "%r8w",
	R9 - RLo:  "%r9w",
	R10 - RLo: "%r10w",
	R11 - RLo: "%r11w",
	R12 - RLo: "%r12w",
	R13 - RLo: "%r13w",
	R14 - RLo: "%r14w",
	R15 - RLo: "%r15w",
}
var regName4 = [...]string{
	RAX - RLo: "%eax",
	RCX - RLo: "%ecx",
	RDX - RLo: "%edx",
	RBX - RLo: "%ebx",
	RSP - RLo: "%esp",
	RBP - RLo: "%ebp",
	RSI - RLo: "%esi",
	RDI - RLo: "%edi",
	R8 - RLo:  "%r8d",
	R9 - RLo:  "%r9d",
	R10 - RLo: "%r10d",
	R11 - RLo: "%r11d",
	R12 - RLo: "%r12d",
	R13 - RLo: "%r13d",
	R14 - RLo: "%r14d",
	R15 - RLo: "%r15d",
}
var regName8 = [...]string{
	RAX - RLo: "%rax",
	RCX - RLo: "%rcx",
	RDX - RLo: "%rdx",
	RBX - RLo: "%rbx",
	RSP - RLo: "%rsp",
	RBP - RLo: "%rbp",
	RSI - RLo: "%rsi",
	RDI - RLo: "%rdi",
	R8 - RLo:  "%r8",
	R9 - RLo:  "%r9",
	R10 - RLo: "%r10",
	R11 - RLo: "%r11",
	R12 - RLo: "%r12",
	R13 - RLo: "%r13",
	R14 - RLo: "%r14",
	R15 - RLo: "%r15",
}

func bits(id RegId) uint8 {
	id.Validate()
	return uint8(id - RLo)
}

func lohiId(id RegId) (uint8, uint8) {
	bits := bits(id)
	return bits & 0x7, (bits & 0x8) >> 3
}

func lohi(r Reg) (uint8, uint8) {
	return lohiId(r.RegId())
}

// return number of assembler bytes needed to encode m.off
func offlen(m Mem, id RegId) (offlen uint8, offbit uint8) {
	moffset := m.Offset()
	switch {
	// (%rbp) and (%r13) registers must use 1-byte offset even if m.off == 0
	case moffset == 0 && id != RBP && id != R13:
		return 0, 0
	case moffset == int32(int8(moffset)):
		return 1, 0x40
	default:
		return 4, 0x80
	}
}

func quirk24(asm *Asm, id RegId) *Asm {
	if id == RSP || id == R12 {
		asm.Bytes(0x24) // amd64 quirk
	}
	return asm
}
