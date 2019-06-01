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
 * arch.go
 *
 *  Created on May 26, 2018
 *      Author Massimiliano Ghilardi
 */

package arm64

import (
	"fmt"

	"github.com/cosmos72/gomacro/jit/common"
)

type Arm64 struct {
}

func init() {
	common.Archs[ARM64] = Arm64{}
}

// implement Arch interface
func (Arm64) Id() ArchId {
	return ARM64
}

func (Arm64) String() string {
	return NAME
}

func (Arm64) RegIdConfig() RegIdConfig {
	return RegIdConfig{
		RLo:         RLo,
		RHi:         RHi,
		RSP:         RSP,
		RVAR:        RVAR,
		RAllocFirst: RLo,
	}
}

func (Arm64) RegIdValid(id RegId) bool {
	return id >= RLo && id < RHi // XZR/XSP is valid only in few, hand-checked cases
}

func (Arm64) RegIdString(id RegId) string {
	var s string
	if id >= RLo && id <= RHi {
		s = regName8[id]
	}
	if s == "" {
		s = fmt.Sprintf("unknown_reg(%#x)", uint8(id))
	}
	return s
}

func (Arm64) RegValid(r Reg) bool {
	// XZR/XSP is valid only in few, hand-checked cases
	return r.RegId().Valid() && r.Kind().Size() != 0
}

func (Arm64) RegString(r Reg) string {
	var s string
	id := r.RegId()
	if id >= RLo && id <= RHi {
		switch r.Kind().Size() {
		case 1, 2, 4:
			s = regName4[id]
		case 8:
			s = regName8[id]
		}
	}
	if s == "" {
		s = fmt.Sprintf("unknown_reg(%#x,%v)", uint8(id), r.Kind())
	}
	return s
}

func (arch Arm64) MemString(m Mem) string {
	var regstr string
	regid := m.RegId()
	if regid == XZR {
		regstr = "xsp"
	} else {
		regstr = arch.RegIdString(regid)
	}
	return fmt.Sprintf("%v@{%s+%v}", m.Kind(), regstr, m.Offset())

}

// print arm64 machine code as sequence of 4-byte instructions
func (Arm64) CodeString(code MachineCode) string {
	const hexdigit string = "0123456789abcdef"
	bytes := code.Bytes
	i, j, n := 0, 0, len(bytes)
	buf := make([]byte, (n+3)/4*9)
	for ; i+4 <= n; i += 4 {
		buf[j+0] = hexdigit[bytes[i+3]>>4]
		buf[j+1] = hexdigit[bytes[i+3]&0xF]
		buf[j+2] = hexdigit[bytes[i+2]>>4]
		buf[j+3] = hexdigit[bytes[i+2]&0xF]
		buf[j+4] = hexdigit[bytes[i+1]>>4]
		buf[j+5] = hexdigit[bytes[i+1]&0xF]
		buf[j+6] = hexdigit[bytes[i+0]>>4]
		buf[j+7] = hexdigit[bytes[i+0]&0xF]
		buf[j+8] = ' '
		j += 9
	}
	for k := n - 1; k >= i; k-- {
		buf[j+0] = hexdigit[bytes[k]>>4]
		buf[j+1] = hexdigit[bytes[k]&0xF]
		j += 2
	}
	return string(buf[:j])
}

// Prologue used to add the following instruction to generated code,
// but now it does nothing, because adding ANY code is the user's responsibility:
//   ldr x29, [sp, #8]
// equivalent to:
// asm.Asm(MOV, MakeMem(8, XSP, Uint64), MakeReg(X29, Uint64))
func (Arm64) Prologue(asm *Asm) *Asm {
	// return asm.Uint32(0xf94007fd)
	// equivalent:
	// return asm.Asm(MOV, MakeMem(8, XSP, Uint64), MakeReg(X29, Uint64))
	return asm
}

func (arch Arm64) Epilogue(asm *Asm) *Asm {
	return arch.Op0(asm, RET)
}

func (Arm64) Init(asm *Asm, start SaveSlot, end SaveSlot) *Asm {
	asm.RegIncUse(X28) // pointer to goroutine-local data
	asm.RegIncUse(X30) // return address register
	asm.RegIncUse(XZR) // zero register / stack pointer
	return asm
}
