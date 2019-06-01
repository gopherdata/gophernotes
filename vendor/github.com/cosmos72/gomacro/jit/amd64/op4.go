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
 * op4.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package amd64

// ============================================================================
// four-arg instruction

func op4val(op Op4) uint8 {
	if op != LEA4 {
		errorf("unknown Op4 instruction: %v", op)
	}
	return 0x8D
}

// ============================================================================
func (arch Amd64) Op4(asm *Asm, op Op4, a Arg, b Arg, c Arg, dst Arg) *Asm {
	arch.op4(asm, op, a, b, c, dst)
	return asm
}

func (arch Amd64) op4(asm *Asm, op Op4, a Arg, b Arg, c Arg, dst Arg) Amd64 {
	assert(op == LEA4)

	src_m := a.(Mem)
	var reg Reg
	var scale int64
	if b != nil {
		reg = b.(Reg)
	}
	if c != nil {
		scale = c.(Const).Val()
	}
	dreg := dst.(Reg)

	if reg.RegId() == NoRegId || scale == 0 {
		return arch.op2MemReg(asm, LEA2, src_m, dreg)
	} else if src_m.RegId() == NoRegId && scale == 1 {
		return arch.op2MemReg(asm, LEA2, MakeMem(src_m.Offset(), reg.RegId(), src_m.Kind()), dreg)
	}
	return arch.lea4(asm, src_m, reg, scale, dreg)
}

func (arch Amd64) lea4(asm *Asm, m Mem, reg Reg, scale int64, dst Reg) Amd64 {
	op := LEA4
	assert(SizeOf(dst) == 8)
	assert(SizeOf(m) == 8)
	assert(SizeOf(reg) == 8)
	var scalebit uint8
	switch scale {
	case 1:
		scalebit = 0
	case 2:
		scalebit = 0x40
	case 4:
		scalebit = 0x80
	case 8:
		scalebit = 0xC0
	default:
		errorf("LEA: unsupported scale %v, expecting 1,2,4 or 8: %v %v, %v, %v, %v",
			scale, op, m, reg, scale, dst)
	}
	dlo, dhi := lohi(dst)
	var mlo, mhi uint8
	var mofflen, offbit uint8

	if mregid := m.RegId(); mregid.Valid() {
		mlo, mhi = lohiId(mregid)
		mofflen, offbit = offlen(m, mregid)
	} else {
		// no mem register
		mofflen = 4
		scalebit |= 0x05
	}
	if reg.RegId() == RSP {
		errorf("LEA: register RSP cannot be scaled: %v %v, %v, %v, %v",
			op, m, reg, scale, dst)
	}
	rlo, rhi := lohi(reg)

	asm.Bytes(0x48|dhi<<2|rhi<<1|mhi, op4val(op), offbit|0x04|dlo<<3, scalebit|rlo<<3|mlo)
	switch mofflen {
	case 1:
		asm.Int8(int8(m.Offset()))
	case 4:
		asm.Int32(m.Offset())
	}
	return arch
}
