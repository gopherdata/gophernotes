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
 * set_arm64.go
 *
 *  Created on May 27, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

func idx(a *Var) uint32 {
	return uint32(a.idx) * S
}

func (asm *Asm) load(dst hwReg, src Arg) *Asm {
	switch a := src.(type) {
	case Reg:
		return asm.mov(dst, asm.reg(a))
	case *Const:
		return asm.loadConst(dst, a.val)
	case *Var:
		off := idx(a)
		if off <= 32760 && off&7 == 0 {
			return asm.Uint32(0xf94003a0 | off<<7 | dst.lo()) // ldr   xdst, [x29, #src]
		}
		tmp := asm.hwAllocConst(int64(off))
		asm.Uint32(0xf8606ba0 | tmp.lo()<<16 | dst.lo()) //	     ldr   xdst, [x29, xtmp]
		return asm.hwFree(tmp, true)
	default:
		errorf("invalid src type: %#v // %T", a, a)
		return nil
	}
}

func (asm *Asm) loadConst(dst hwReg, val int64) *Asm {
	lo := dst.lo()
	u := uint64(val)
	asm.Uint32(0xd2800000 | uint32(u&0xffff)<<5 | lo) //	     mov   xdst, #val16
	u >>= 16
	for shift := uint32(1); u != 0 && shift <= 3; shift++ {
		if mask := uint32(u & 0xffff); mask != 0 {
			asm.Uint32(0xf2800000 | shift<<21 | mask<<5 | lo) // movk  xdst, #mask, lsl #shift
		}
		u >>= 16
	}
	return asm
}

func (asm *Asm) mov(dst hwReg, src hwReg) *Asm {
	if dst == src {
		return asm
	}
	return asm.Uint32(0xaa0003e0 | src.lo()<<16 | dst.lo()) //  mov   xdst, xsrc
}

func (asm *Asm) store(dst *Var, src Arg) *Asm {
	switch a := src.(type) {
	case *Var:
		if dst.desc == a.desc {
			return asm
		}
	}
	tmp, alloc := asm.hwAlloc(src)
	asm.storeReg(dst, tmp)
	return asm.hwFree(tmp, alloc)
}

func (asm *Asm) storeReg(dst *Var, src hwReg) *Asm {
	off := idx(dst)
	if off <= 32760 && off&7 == 0 {
		return asm.Uint32(0xf90003a0 | off<<7 | src.lo()) // str   xsrc, [x29, #dst]
	}
	tmp := asm.hwAllocConst(int64(off))
	asm.Uint32(0xf8206ba0 | tmp.lo()<<16 | src.lo()) //	      str   xsrc, [x29, xtmp]
	return asm.hwFree(tmp, true)
}
