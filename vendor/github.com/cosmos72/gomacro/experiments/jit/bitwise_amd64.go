// +build amd64

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
 * bitwise_amd64.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

// %reg_z &= a
func (asm *Asm) And(z Reg, a Arg) *Asm {
	lo, hi := asm.lohi(z)
	if a.Const() {
		val := a.(*Const).val
		if val == 0 {
			return asm.LoadConst(z, 0)
		} else if val == -1 {
			return asm
		} else if val == int64(uint32(val)) {
			if hi != 0 {
				asm.Bytes(0x41)
			}
			return asm.Bytes(0x81, 0xe0+lo).Uint32(uint32(val)) //        andl  $val,%reg_z // zero extend
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48+hi, 0x81, 0xe0+lo).Int32(int32(val)) // andq  $val,%reg_z // sign extend
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Bytes(0x48+hi+tmp.hi()*4, 0x21, 0xc0+lo+tmp.lo()*8) //      and  %reg_tmp,%reg_z
	asm.hwFree(tmp, alloc)
	return asm
}

// %reg_z |= a
func (asm *Asm) Or(z Reg, a Arg) *Asm {
	lo, hi := asm.lohi(z)
	if a.Const() {
		val := a.(*Const).val
		if val == 0 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48+hi, 0x81, 0xc8+lo).Int32(int32(val)) // orq   $val,%reg_z // sign extend
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Bytes(0x48+hi+tmp.hi()*4, 0x09, 0xc0+lo+tmp.lo()*8) //      or   %reg_tmp,%reg_z
	asm.hwFree(tmp, alloc)
	return asm
}

// %reg_z ^= a
func (asm *Asm) Xor(z Reg, a Arg) *Asm {
	lo, hi := asm.lohi(z)
	if a.Const() {
		val := a.(*Const).val
		if val == 0 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48+hi, 0x81, 0xf0+lo).Int32(int32(val)) // xorq  $val,%reg_z // sign extend
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Bytes(0x48+hi+tmp.hi()*4, 0x31, 0xc0+lo+tmp.lo()*8) //      xor  %reg_tmp,%reg_z
	asm.hwFree(tmp, alloc)
	return asm
}

// %reg_z &^= a
func (asm *Asm) Andnot(z Reg, a Arg) *Asm {
	lo, hi := asm.lohi(z)
	var tmp hwReg
	if a.Const() {
		val := ^a.(*Const).val // negate val!
		if val == 0 {
			return asm.LoadConst(z, 0)
		} else if val == -1 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48+hi, 0x81, 0xe0+lo).Int32(int32(val)) // andq  $val,%reg_z // sign extend
		}
		tmp = asm.hwAllocConst(val)
	} else {
		// always allocate a register, because we need to complement it
		tmp = asm.hwRegs.Alloc()
		asm.load(tmp, a)
		asm.Bytes(0x48|tmp.hi(), 0xf7, 0xd0|tmp.lo()) //            not    %reg_tmp
	}
	asm.Bytes(0x48+hi+tmp.hi()*4, 0x21, 0xc0+lo+tmp.lo()*8) //      and  %reg_tmp,%reg_z
	asm.hwFree(tmp, true)
	return asm
}

// %reg_z = ^ %reg_z
func (asm *Asm) Not(z Reg) *Asm {
	lo, hi := asm.lohi(z)
	asm.Bytes(0x48+hi, 0xf7, 0xd0+lo) //  not    %reg_z
	return asm
}
