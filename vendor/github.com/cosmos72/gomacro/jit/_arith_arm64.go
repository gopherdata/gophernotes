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
 * arith_arm64.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

// xz += a
func (asm *Asm) Add(z Reg, a Arg) *Asm {
	if a.Const() {
		val := a.(*Const).val
		if asm.add_const(z, val) || asm.sub_const(z, -val) {
			return asm
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Uint32(0x8b<<24 | tmp.lo()<<16 | asm.lo(z)*0x21) //  add  xz, xz, xtmp
	asm.hwFree(tmp, alloc)
	return asm
}

// xz -= a
func (asm *Asm) Sub(z Reg, a Arg) *Asm {
	if a.Const() {
		val := a.(*Const).val
		if asm.sub_const(z, val) || asm.add_const(z, -val) {
			return asm
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Uint32(0xcb<<24 | tmp.lo()<<16 | asm.lo(z)*0x21) //  sub  xz, xz, xtmp
	asm.hwFree(tmp, alloc)
	return asm
}

func (asm *Asm) add_const(z Reg, val int64) bool {
	if val == 0 {
		return true
	} else if uint64(val) < 4096 {
		asm.Uint32(0x91<<24 | uint32(val)<<10 | asm.lo(z)*0x21) // add  xz, xz, $val
		return true
	}
	return false
}

func (asm *Asm) sub_const(z Reg, val int64) bool {
	if val == 0 {
		return true
	} else if uint64(val) < 4096 {
		asm.Uint32(0xd1<<24 | uint32(val)<<10 | asm.lo(z)*0x21) // sub  xz, xz, $val
		return true
	}
	return false
}

// xz *= a
func (asm *Asm) Mul(z Reg, a Arg) *Asm {
	if a.Const() {
		val := a.(*Const).val
		if val == 0 {
			return asm.LoadConst(z, 0)
		} else if val == 1 {
			return asm
		} else if val == 2 {
			return asm.Add(z, z)
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Uint32(0x9b007c00 | tmp.lo()<<16 | asm.lo(z)*0x21) //  mul  xz, xz, xtmp
	asm.hwFree(tmp, alloc)
	return asm
}

// xz /= a       signed division
func (asm *Asm) SDiv(z Reg, a Arg) *Asm {
	return asm.div(z, a, signed)
}

// xz /= a       unsigned division
func (asm *Asm) UDiv(z Reg, a Arg) *Asm {
	return asm.div(z, a, unsigned)
}

// xz %= a       signed remainder
func (asm *Asm) SRem(z Reg, a Arg) *Asm {
	return asm.rem(z, a, signed)
}

// xz %= a       unsigned remainder
func (asm *Asm) URem(z Reg, a Arg) *Asm {
	return asm.rem(z, a, unsigned)
}

func (asm *Asm) div(z Reg, a Arg, k divkind) *Asm {
	if a.Const() {
		val := a.(*Const).val
		if val == 0 {
			// cause a runtime fault by clearing x29 then dereferencing it
			return asm.loadConst(x29, 0).storeReg(&Var{}, x29)
		} else if val == 1 {
			return asm
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	var flag uint32
	if k&unsigned == 0 {
		flag = 0x400
	}
	asm.Uint32(0x9ac00800 | flag | tmp.lo()<<16 | asm.lo(z)*0x21) //  {s,u}div  xz, xz, xtmp
	asm.hwFree(tmp, alloc)
	return asm
}

func (asm *Asm) rem(z Reg, a Arg, k divkind) *Asm {
	if a.Const() {
		c := a.(*Const)
		val := c.val
		if val == 0 {
			// cause a runtime fault by clearing x29 then dereferencing it
			return asm.loadConst(x29, 0).storeReg(&Var{}, x29)
		} else if val&(val-1) == 0 {
			// transform xz %= power-of-two
			// into      zx &= power-of-two - 1
			return asm.And(z, &Const{c.kind, val - 1})
		}
	}
	den, alloc := asm.hwAlloc(a) //                                       // den = a
	quo := asm.hwRegs.Alloc()
	var flag uint32
	if k&unsigned == 0 {
		flag = 0x400
	}
	asm.Uint32(0x9ac08000 | flag | den.lo()<<16 | asm.lo(z)<<5 | quo.lo()) // {s,u}div  quo, xz, den  // quo = xz / den
	asm.Uint32(0x9b008000 | den.lo()<<16 | quo.lo()<<5 | asm.lo(z)*0x401)  // msub  xz, quo, den, xz  // xz  = xz - quo * den
	asm.hwFree(quo, true)
	asm.hwFree(den, alloc)
	return asm
}

// xz = - xz
func (asm *Asm) Neg(z Reg) *Asm {
	return asm.Uint32(0xcb0003e0 | asm.lo(z)*0x10001) // neg xz, xz
}
