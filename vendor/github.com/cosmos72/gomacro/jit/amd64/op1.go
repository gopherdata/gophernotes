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
 * op1.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package amd64

// ============================================================================
// one-arg instruction

var op1Val = [256]uint8{
	ZERO: 0x08,
	NOT1: 0x10,
	NEG1: 0x18,
	INC:  0x20,
	DEC:  0x28,
}

func op1lohi(op Op1) (uint8, uint8) {
	val := op1Val[op]
	if val == 0 {
		errorf("unknown Op1 instruction: %v", op)
	}
	return val & 0x18, val >> 2
}

// ============================================================================
func (arch Amd64) Op1(asm *Asm, op Op1, a Arg) *Asm {
	arch.op1(asm, op, a)
	return asm
}

func (arch Amd64) op1(asm *Asm, op Op1, a Arg) Amd64 {
	if op == JMP {
		return arch.jmp(asm, a)
	}
	switch a := a.(type) {
	case Reg:
		arch.op1Reg(asm, op, a)
	case Mem:
		arch.op1Mem(asm, op, a)
	case Const:
		errorf("destination cannot be a constant: %v %v", op, a)
	default:
		errorf("unknown destination type %T, expecting Reg or Mem: %v %v", a, op, a)
	}
	return arch
}

// OP %reg_dst
func (arch Amd64) op1Reg(asm *Asm, op Op1, r Reg) Amd64 {
	if op == ZERO {
		return arch.zeroReg(asm, r)
	}
	rlo, rhi := lohi(r)
	oplo, ophi := op1lohi(op)

	switch SizeOf(r) {
	case 1:
		if r.RegId() >= RSP {
			asm.Byte(0x40 | rhi)
		}
		asm.Bytes(0xF6|ophi, 0xC0|oplo|rlo)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if rhi != 0 {
			asm.Byte(0x41)
		}
		asm.Bytes(0xF7|ophi, 0xC0|oplo|rlo)
	case 8:
		asm.Bytes(0x48|rhi, 0xF7|ophi, 0xC0|oplo|rlo)
	default:
		errorf("unsupported register size %v, expecting 1,2,4 or 8 bytes: %v %v", SizeOf(r), op, r)
	}
	return arch
}

// OP off_m(%reg_m)
func (arch Amd64) op1Mem(asm *Asm, op Op1, m Mem) Amd64 {
	if op == ZERO {
		return arch.zeroMem(asm, m)
	}
	regid := m.RegId()
	rlo, dhi := lohiId(regid)
	oplo, ophi := op1lohi(op)

	offlen, offbit := offlen(m, regid)

	switch SizeOf(m) {
	case 1:
		if dhi != 0 {
			asm.Byte(0x41)
		}
		asm.Bytes(0xF6|ophi, offbit|oplo|rlo)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi != 0 {
			asm.Byte(0x41)
		}
		asm.Bytes(0xF7|ophi, offbit|oplo|rlo)
	case 8:
		asm.Bytes(0x48|dhi, 0xF7|ophi, offbit|oplo|rlo)
	default:
		errorf("unsupported memory size %v, expecting 1,2,4 or 8 bytes: %v %v", SizeOf(m), op, m)
	}
	quirk24(asm, regid)
	switch offlen {
	case 1:
		asm.Int8(int8(m.Offset()))
	case 4:
		asm.Int32(m.Offset())
	}
	return arch
}

// zero a register or memory location
func (arch Amd64) zero(asm *Asm, dst Arg) Amd64 {
	switch dst := dst.(type) {
	case Reg:
		arch.zeroReg(asm, dst)
	case Mem:
		arch.zeroMem(asm, dst)
	case Const:
		errorf("destination cannot be a constant: %v %v", ZERO, dst)
	default:
		errorf("unknown destination type %T, expecting Reg or Mem: %v %v", dst, ZERO, dst)
	}
	return arch
}

func (arch Amd64) zeroReg(asm *Asm, dst Reg) Amd64 {
	dlo, dhi := lohi(dst)
	if dhi == 0 {
		asm.Bytes(0x31, 0xC0|dlo|dlo<<3)
	} else {
		asm.Bytes(0x48|dhi<<1|dhi<<2, 0x31, 0xC0|dlo|dlo<<3)
	}
	return arch
}

// zero a memory location
func (arch Amd64) zeroMem(asm *Asm, dst Mem) Amd64 {
	return arch.movConstMem(asm, MakeConst(0, dst.Kind()), dst)
}

func (arch Amd64) jmp(asm *Asm, dst Arg) Amd64 {
	c, ok := dst.(Const)
	if !ok {
		errorf("JMP destination must be a constant: %v %v", JMP, dst)
	}
	if dst.Kind() != Ptr {
		errorf("unimplemented: relative jump: %v %v", JMP, dst)
	}
	asm.Byte(0xE9)
	asm.Int32(0) // adjusted later by Asm.link()
	asm.AddJump(uintptr(c.Val()))
	return arch
}
