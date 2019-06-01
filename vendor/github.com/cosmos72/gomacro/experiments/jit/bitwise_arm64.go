// +build arm64

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
 * bitwise_arm64.go
 *
 *  Created on May 27, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

// xz &= a
func (asm *Asm) And(z Reg, a Arg) *Asm {
	if a.Const() {
		if asm.and_const(z, a.(*Const).val) {
			return asm
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Uint32(0x8a<<24 | tmp.lo()<<16 | asm.lo(z)*0x21) //  and  xz, xz, xtmp
	asm.hwFree(tmp, alloc)
	return asm
}

// xz |= a
func (asm *Asm) Or(z Reg, a Arg) *Asm {
	if a.Const() {
		if asm.or_const(z, a.(*Const).val) {
			return asm
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Uint32(0xaa<<24 | tmp.lo()<<16 | asm.lo(z)*0x21) //  orr  xz, xz, xtmp
	asm.hwFree(tmp, alloc)
	return asm
}

// xz ^= a
func (asm *Asm) Xor(z Reg, a Arg) *Asm {
	if a.Const() {
		if asm.xor_const(z, a.(*Const).val) {
			return asm
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Uint32(0xca<<24 | tmp.lo()<<16 | asm.lo(z)*0x21) //  eor  xz, xz, xtmp
	asm.hwFree(tmp, alloc)
	return asm
}

// xz &^= a
func (asm *Asm) Andnot(z Reg, a Arg) *Asm {
	var tmp hwReg
	if a.Const() {
		val := ^a.(*Const).val // complement val
		if asm.and_const(z, val) {
			return asm
		}
		tmp = asm.hwAllocConst(val)
	} else {
		// always allocate a register, because we need to complement it
		tmp = asm.hwRegs.Alloc()
		asm.load(tmp, a)
		asm.Uint32(0xaa2003e0 | tmp.lo()*0x10001) //  mvn  xtmp, xtmp
	}
	asm.Uint32(0x8a<<24 | tmp.lo()<<16 | asm.lo(z)*0x21) //  and  xz, xz, xtmp
	asm.hwFree(tmp, true)
	return asm
}

// xz = ^ xz
func (asm *Asm) Not(z Reg) *Asm {
	return asm.Uint32(0xaa2003e0 | asm.lo(z)*0x10001) //  mvn  xz, xz
}

func (asm *Asm) and_const(z Reg, val int64) bool {
	if val == 0 {
		asm.LoadConst(z, 0)
		return true
	} else if val == -1 {
		return true
	} else if bitmask, ok := bitmask_imm[uint64(val)]; ok {
		asm.Uint32(0x92<<24 | uint32(bitmask)<<10 | asm.lo(z)*0x21)
		return true
	}
	return false
}

func (asm *Asm) or_const(z Reg, val int64) bool {
	if val == 0 {
		return true
	} else if val == -1 {
		asm.LoadConst(z, -1)
		return true
	} else if bitmask, ok := bitmask_imm[uint64(val)]; ok {
		asm.Uint32(0xb2<<24 | uint32(bitmask)<<10 | asm.lo(z)*0x21)
		return true
	}
	return false
}

func (asm *Asm) xor_const(z Reg, val int64) bool {
	if val == 0 {
		return true
	} else if val == -1 {
		asm.Not(z)
		return true
	} else if bitmask, ok := bitmask_imm[uint64(val)]; ok {
		asm.Uint32(0xd2<<24 | uint32(bitmask)<<10 | asm.lo(z)*0x21)
		return true
	}
	return false
}

// the possible immediate constants for bitwise operations are quite complicated:
// see https://dinfuehr.github.io/blog/encoding-of-immediate-values-on-aarch64/
// and https://stackoverflow.com/questions/30904718/range-of-immediate-values-in-armv8-a64-assembly/33265035#33265035
//
// solution: generate them and store in a map for fast lookup
var bitmask_imm = make(map[uint64]uint16)

func init() {
	for size := uint16(2); size <= 64; size *= 2 {
		var n, imms uint16 = 0, 0x1e * size
		if size == 64 {
			n, imms = 1, 0
		}
		for length := uint16(1); length < size; length++ {
			val := ^uint64(0) >> (64 - length)
			for e := size; e < 64; e *= 2 {
				val |= val << e
			}
			for rotation := uint16(0); rotation < size; rotation++ {
				bitmask_imm[val] = n<<12 | rotation<<6 | imms | (length - 1)
				val = val>>1 | val<<63
			}
		}
	}
}
