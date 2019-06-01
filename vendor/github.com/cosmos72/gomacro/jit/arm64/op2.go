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
 * op2.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arm64

// ============================================================================
// two-arg instruction

func op2val(op Op2) uint32 {
	var val uint32
	switch op {
	case NEG2:
		val = 0x4B0003E0
	case NOT2:
		val = 0x2A2003E0
	default:
		errorf("unknown Op2 instruction: %v", op)
	}
	return val
}

// ============================================================================
func (arch Arm64) Op2(asm *Asm, op Op2, src Arg, dst Arg) *Asm {
	arch.op2(asm, op, src, dst)
	return asm
}

func (arch Arm64) op2(asm *Asm, op Op2, src Arg, dst Arg) Arm64 {
	switch op {
	case CAST:
		if SizeOf(src) != SizeOf(dst) {
			return arch.cast(asm, src, dst)
		}
		fallthrough
	case MOV:
		return arch.mov(asm, src, dst)
	case NEG2, NOT2:
		break
	default:
		// dst OP= src
		//    translates to
		// dst = dst OP src
		//    note the argument order
		return arch.op3(asm, Op3(op), dst, src, dst)
	}

	op2val(op) // validate op

	assert(src.Kind() == dst.Kind())
	if dst.Const() {
		errorf("destination cannot be a constant: %v %v, %v", op, src, dst)
	}

	switch src := src.(type) {
	case Reg:
		switch dst := dst.(type) {
		case Reg:
			arch.op2RegReg(asm, op, src, dst)
		case Mem:
			r := asm.RegAlloc(dst.Kind())
			arch.op2RegReg(asm, op, src, r).store(asm, r, dst)
			asm.RegFree(r)
		default:
			errorf("unknown destination type %T, expecting Reg or Mem: %v %v, %v", dst, op, src, dst)
		}
	case Mem:
		switch dst := dst.(type) {
		case Reg:
			arch.load(asm, src, dst).op2RegReg(asm, op, dst, dst)
		case Mem:
			r := asm.RegAlloc(dst.Kind())
			arch.load(asm, src, r).op2RegReg(asm, op, r, r).store(asm, r, dst)
			asm.RegFree(r)
		default:
			errorf("unknown destination type %T, expecting Reg or Mem: %v %v, %v", dst, op, src, dst)
		}
	case Const:
		var c Const
		if op == NEG2 {
			c = MakeConst(-src.Val(), src.Kind())
		} else {
			c = MakeConst(^src.Val(), src.Kind())
		}
		return arch.mov(asm, c, dst)
	default:
		errorf("unknown argument type %T, expecting Const, Reg or Mem: %v %v, %v", src, op, src, dst)
	}
	return arch
}

func (arch Arm64) op2RegReg(asm *Asm, op Op2, src Reg, dst Reg) Arm64 {
	asm.Uint32(kbit(dst) | op2val(op) | val(src)<<16 | val(dst))
	return arch
}
