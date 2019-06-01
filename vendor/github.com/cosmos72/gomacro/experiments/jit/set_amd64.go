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
 * set_amd64.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

func (asm *Asm) load(dst hwReg, src Arg) *Asm {
	switch a := src.(type) {
	case Reg:
		return asm.mov(dst, asm.reg(a))
	case *Const:
		return asm.loadConst(dst, a.val)
	case *Var:
		lo, hi := dst.lohi()
		return asm.Bytes(0x48|hi*4, 0x8b, 0x87|lo*8).Idx(a) //   movq   src(%rdi),%reg
	default:
		errorf("invalid src type: %#v // %T", a, a)
		return nil
	}
}

func (asm *Asm) loadConst(dst hwReg, val int64) *Asm {
	lo, hi := dst.lohi()
	if val == int64(uint32(val)) {
		if hi != 0 {
			asm.Bytes(0x41)
		}
		return asm.Bytes(0xb8 + lo).Uint32(uint32(val)) //            movl   $val,%regl // zero extend
	} else if val == int64(int32(val)) {
		return asm.Bytes(0x48|hi, 0xc7, 0xc0|lo).Int32(int32(val)) // movq   $val,%reg  // sign extend
	} else {
		return asm.Bytes(0x48|hi, 0xb8+lo).Int64(val) //              movabs $val,%reg
	}
}

func (asm *Asm) mov(dst hwReg, src hwReg) *Asm {
	if dst == src {
		return asm
	}
	slo, shi := src.lohi()
	dlo, dhi := dst.lohi()
	return asm.Bytes(0x48|dhi|shi*4, 0x89, 0xc0+dlo+slo*8) //  movq   %reg_src,%reg_dst
}

func (asm *Asm) store(dst *Var, src Arg) *Asm {
	switch a := src.(type) {
	case *Const:
		if val := a.val; val == int64(int32(val)) {
			return asm.Bytes(0x48, 0xc7, 0x87).Idx(dst).Int32(int32(val)) //  movq   $val,z(%rdi)
		}
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
	lo, hi := src.lohi()
	return asm.Bytes(0x48|hi*4, 0x89, 0x87|lo*8).Idx(dst) //   movq   %reg,dst(%rdi)
}
