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
 * arith_amd64.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

func assert(flag bool) {
	if !flag {
		panic("jit internal error, assertion failed")
	}
}

// %reg_z = const
func (asm *Asm) mov_const(z hwReg, c uint64) *Asm {
	if c == 0 {
		return asm.xor_reg_self(z)
	}
	if s := int64(c); s == int64(int32(s)) {
		return asm.op_reg_const(z, hwMOV, int32(s))
	}
	zlo, zhi := z.lohi()
	return asm.Bytes(0x48|zhi, 0xB8|zlo).Uint64(c)
}

// %reg_z ^= %reg_z // compact way to zero a register
func (asm *Asm) xor_reg_self(z hwReg) *Asm {
	zlo, zhi := z.lohi()
	if zhi == 0 {
		return asm.Bytes(0x31, 0xC0|zlo|zlo<<3)
	} else {
		return asm.Bytes(0x48|zhi<<1|zhi<<2, 0x31, 0xC0|zlo|zlo<<3)
	}
}

// %reg_z OP= const
func (asm *Asm) op_reg_const(z hwReg, op hwOp, c int32) *Asm {
	zlo, zhi := z.lohi()

	if op == hwMOV {
		// hwMOV has different encoding and only supports
		// 32-bit signed immediate constants.
		// Use mov_const for 64bit wide constants.
		return asm.Bytes(0x48|zhi, 0xC7, 0xC0|zlo).Int32(c)
	}

	if c == int32(int8(c)) {
		asm.Bytes(0x48|zhi, 0x83, 0xC0|op|zlo, uint8(int8(c)))
	} else if z == rAX {
		asm.Bytes(0x48|zhi, 0x05|op).Int32(c)
	} else {
		asm.Bytes(0x48|zhi, 0x81, 0xC0|op|zlo).Int32(c)
	}
	return asm
}

// %reg_z OP= %reg_r
func (asm *Asm) op_reg_reg(z hwReg, op hwOp, r hwReg) *Asm {
	zlo, zhi := z.lohi()
	rlo, rhi := r.lohi()

	return asm.Bytes(0x48|zhi|rhi<<2, 0x01|op, 0xC0|zlo|rlo<<3)
}

// off_m(%reg_m) OP= %reg_r
func (asm *Asm) op_mem_reg(m hwMem, op hwOp, r hwReg) *Asm {

	assert(m.off < 0x80000000)
	assert(m.siz == 8) // TODO mem access by 1, 2 or 4 bytes

	z := m.reg
	zlo, zhi := z.lohi()
	rlo, rhi := r.lohi()

	// (%rbp) and (%r13) destinations must use 1-byte offset even if m.off == 0
	noOffset := m.off == 0 && z != rBP && z != rR13

	if noOffset {
		asm.Bytes(0x48|zhi|rhi<<2, 0x01|op, zlo|rlo<<3)
	} else if m.off < 0x80 {
		asm.Bytes(0x48|zhi|rhi<<2, 0x01|op, 0x40|zlo|rlo<<3)
	} else {
		asm.Bytes(0x48|zhi|rhi<<2, 0x01|op, 0x80|zlo|rlo<<3)
	}
	if z == rSP || z == rR12 {
		asm.Bytes(0x24) // amd64 quirk
	}
	if noOffset {
		// nothing to do
	} else if m.off < 0x80 {
		asm.Bytes(uint8(m.off))
	} else {
		asm.Uint32(m.off)
	}
	return asm
}

// %reg_z OP= off_m(%reg_m)
func (asm *Asm) op_reg_mem(z hwReg, op hwOp, m hwMem) *Asm {

	assert(m.off < 0x80000000)

	zlo, zhi := z.lohi()
	r := m.reg
	rlo, rhi := r.lohi()

	// (%rbp) and (%r13) sources must use 1-byte offset even if m.off == 0
	noOffset := m.off == 0 && r != rBP && r != rR13

	if noOffset {
		asm.Bytes(0x48|zhi|rhi<<2, 0x03|op, zlo<<3|rlo)
	} else if m.off < 0x80 {
		asm.Bytes(0x48|zhi|rhi<<2, 0x03|op, 0x40|zlo<<3|rlo)
	} else {
		asm.Bytes(0x48|zhi|rhi<<2, 0x03|op, 0x80|zlo<<3|rlo)
	}
	if r == rSP || r == rR12 {
		asm.Bytes(0x24) // amd64 quirk
	}
	if noOffset {
		// nothing to do
	} else if m.off < 0x80 {
		asm.Bytes(uint8(m.off))
	} else {
		asm.Uint32(m.off)
	}
	return asm
}
