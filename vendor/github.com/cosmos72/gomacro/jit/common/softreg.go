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
 * softreg.go
 *
 *  Created on Feb 09, 2019
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"fmt"
)

// soft register id. 24 bits wide.
type SoftRegId uint32

const (
	FirstSoftRegId = 0x000000
	LastSoftRegId  = 0x7FFFFF
	FirstTempRegId = 0x800000
	LastTempRegId  = 0xFFFFFF
)

func (s SoftRegId) Validate() {
}

func (s SoftRegId) String() string {
	if s >= FirstTempRegId {
		return fmt.Sprintf("T%d", uint32(s-FirstTempRegId))
	}
	return fmt.Sprintf("S%d", uint32(s))
}

// soft register.
// may be mapped by assembler to an actual
// machine register or to a memory location
type SoftReg uint32

func MakeSoftReg(id SoftRegId, kind Kind) SoftReg {
	return SoftReg(id)<<8 | SoftReg(kind)
}

func (s SoftReg) Kind() Kind {
	return Kind(s & 0xFF)
}

func (s SoftReg) Id() SoftRegId {
	return SoftRegId(s >> 8)
}

var sizeSuffix = map[Size]string{
	0: ".?",
	1: ".b",
	2: ".h",
	4: ".w",
}

func (s SoftReg) String() string {
	id := s.Id()
	suffix := sizeSuffix[s.Kind().Size()]

	if id >= FirstTempRegId {
		return fmt.Sprintf("t%d%s", uint32(id-FirstTempRegId), suffix)
	}
	return fmt.Sprintf("s%d%s", uint32(id), suffix)
}

// use Asm to convert softreg to Reg or Mem
func (s SoftReg) Arg(asm *Asm) Arg {
	return asm.softRegs.Arg(s)
}

func (s SoftReg) RegId(asm *Asm) RegId {
	return asm.softRegs.RegId(s)
}

// implement AsmCode interface
func (s SoftReg) asmcode() {
}

// implement jit.Expr interface
func (s SoftReg) Const() bool {
	return false
}

func (s SoftReg) Valid() bool {
	return s.Kind() != Invalid
}

func (s SoftReg) IsTemp() bool {
	return s.Id() >= FirstTempRegId
}

func (s SoftReg) Validate() {
	if !s.Valid() {
		errorf("invalid SoftReg: %v", s)
	}
}

// ===================================

type regIdOrMem struct {
	off   int32
	regId RegId
	ismem bool
}

type SoftRegs map[SoftRegId]regIdOrMem

func (softRegs SoftRegs) Arg(s SoftReg) Arg {
	rm, ok := softRegs[s.Id()]
	if !ok {
		errorf("unallocated SoftReg: %v", s)
		return nil
	} else if rm.ismem {
		return MakeMem(rm.off, rm.regId, s.Kind())
	} else {
		return MakeReg(rm.regId, s.Kind())
	}
}

func (softRegs SoftRegs) RegId(s SoftReg) RegId {
	rm, ok := softRegs[s.Id()]
	if !ok {
		errorf("unallocated SoftReg: %v", s)
	}
	return rm.regId
}
