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
 * arch.go
 *
 *  Created on Feb 13, 2019
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"fmt"
)

type ArchId uint8

const (
	NOARCH ArchId = iota
	ARM64
	AMD64
)

type Arch interface {
	Id() ArchId
	String() string
	RegIdConfig() RegIdConfig
	RegIdValid(id RegId) bool
	RegIdString(id RegId) string // RegId -> string
	RegValid(r Reg) bool
	RegString(r Reg) string          // Reg -> string
	MemString(m Mem) string          // Mem -> string
	CodeString(c MachineCode) string // Code -> string

	Init(asm *Asm, saveStart, saveEnd SaveSlot) *Asm
	Prologue(asm *Asm) *Asm
	Epilogue(asm *Asm) *Asm

	Op0(asm *Asm, op Op0) *Asm
	Op1(asm *Asm, op Op1, dst Arg) *Asm
	Op2(asm *Asm, op Op2, src Arg, dst Arg) *Asm
	Op3(asm *Asm, op Op3, a Arg, b Arg, dst Arg) *Asm
	Op4(asm *Asm, op Op4, a Arg, b Arg, c Arg, dst Arg) *Asm
}

var Archs = make(map[ArchId]Arch) // {ARM64:Arm64{}, AMD64:Amd64{}}

func (archId ArchId) String() string {
	arch := Archs[archId]
	if arch != nil {
		return arch.String()
	}
	return fmt.Sprintf("ArchId(%d)", uint8(archId))
}

func (code MachineCode) String() string {
	arch := Archs[code.ArchId]
	if arch != nil {
		return arch.CodeString(code)
	}
	return fmt.Sprintf("%x", code.Bytes)
}
