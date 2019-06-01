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
 * z_test.go
 *
 *  Created on Feb 07, 2019
 *      Author Massimiliano Ghilardi
 */

package arm64

import (
	"testing"
)

func MakeCode(instr ...uint32) MachineCode {
	bytes := make([]uint8, len(instr)*4)
	for i, inst := range instr {
		bytes[4*i+0] = byte(inst >> 0)
		bytes[4*i+1] = byte(inst >> 8)
		bytes[4*i+2] = byte(inst >> 16)
		bytes[4*i+3] = byte(inst >> 24)
	}
	return MachineCode{ARM64, bytes}
}

func TestSample(t *testing.T) {
	var asm Asm
	asm.InitArch(Arm64{})

	id := RLo
	x := MakeReg(id+0, Uint64)
	y := MakeReg(id+1, Uint64)
	z := MakeReg(id+2, Uint64)
	m := MakeMem(8, id, Uint64)
	c := ConstUint64(0xFFF)
	asm.RegIncUse(id)
	asm.RegIncUse(id + 1)
	asm.RegIncUse(id + 2)
	asm.Assemble( //
		MOV, MakeMem(8, RSP, Uint64), MakeReg(RVAR, Uint64),
		MOV, c, x, //
		MOV, c, m, //
		MOV, m, x, //
		NOP,           //
		ADD3, x, y, z, //
		SUB3, x, y, z, //
		AND3, x, y, z, //
		OR3, x, y, z, //
		XOR3, x, y, z, //
		SHL3, x, y, z, //
		SHR3, x, y, z, //
		NOP,           //
		ADD3, c, x, z, // test commutativity optimization
		SUB3, x, c, z, //
		AND3, c, x, z, //
		OR3, c, x, z, //
		XOR3, x, c, z, //
	).Epilogue()
	asm.RegDecUse(id)
	asm.RegDecUse(id + 1)
	asm.RegDecUse(id + 2)

	actual := asm.Code()
	expected := MakeCode(
		0xf94007fd, //	ldr	x29, [sp, #8]
		0xd281ffe0, //	mov	x0, #0xfff
		0xd281ffe3, //	mov	x3, #0xfff
		0xf9000403, //	str	x3, [x0, #8]
		0xf9400400, //	ldr	x0, [x0, #8]
		0xd503201f, //	nop
		0x8b010002, //	add	x2, x0, x1
		0xcb010002, //	sub	x2, x0, x1
		0x8a010002, //	and	x2, x0, x1
		0xaa010002, //	orr	x2, x0, x1
		0xca010002, //	eor	x2, x0, x1
		0x9ac12002, //	lsl	x2, x0, x1
		0x9ac12402, //	lsr	x2, x0, x1
		0xd503201f, //	nop
		0x913ffc02, //	add	x2, x0, #0xfff
		0xd13ffc02, //	sub	x2, x0, #0xfff
		0x92402c02, //	and	x2, x0, #0xfff
		0xb2402c02, //	orr	x2, x0, #0xfff
		0xd2402c02, //	eor	x2, x0, #0xfff
		0xd65f03c0, //	ret
	)

	if !actual.Equal(expected) {
		t.Errorf("bad assembled code:\n\texpected %s\n\tactual   %s",
			expected, actual)
	}
}

func TestCast(t *testing.T) {
	var asm Asm
	asm.InitArch(Arm64{})

	id := RLo

	for _, skind := range [...]Kind{
		Int8, Int16, Int32, Int64,
		Uint8, Uint16, Uint32, Uint64,
	} {
		src := MakeReg(id, skind)
		for _, dkind := range [...]Kind{Uint8, Uint16, Uint32, Uint64} {
			dst := MakeReg(id, dkind)
			asm.Assemble(CAST, src, dst)
		}
	}

	actual := asm.Code()
	expected := MakeCode(
		0x13001c00, // sxtb	w0, w0
		0x13001c00, // sxtb	w0, w0
		0x93401c00, // sxtb	x0, w0
		0x13003c00, // sxth	w0, w0
		0x93403c00, // sxth	x0, w0
		0x93407c00, // sxtw	x0, w0
		0x12001c00, // and	w0, w0, #0xff
		0x12001c00, // and	w0, w0, #0xff
		0x92401c00, // and	x0, x0, #0xff
		0x12003c00, // and	w0, w0, #0xffff
		0x92403c00, // and	x0, x0, #0xffff
		0x2a0003e0, // mov	w0, w0
	)

	if !actual.Equal(expected) {
		t.Errorf("bad assembled code:\n\texpected %s\n\tactual   %s",
			expected, actual)
	}
}

func TestMem(t *testing.T) {
	var asm Asm
	asm.InitArch(Arm64{})

	id := RLo
	for _, skind := range [...]Kind{
		Int8, Int16, Int32, Int64,
		Uint8, Uint16, Uint32, Uint64,
	} {
		asm.RegIncUse(id)

		s := MakeMem(0, id, skind)
		c := MakeConst(0xFF, skind)
		for _, dkind := range [...]Kind{Uint8, Uint16, Uint32, Uint64} {

			d := MakeMem(8, id, dkind)
			if skind == dkind {
				asm.Assemble(ADD3, s, c, d)
			} else {
				asm.Assemble(CAST, s, d)
			}
		}
		asm.Assemble(NOP)
	}
	asm.Epilogue()

	actual := asm.Code()
	expected := MakeCode(
		0x39400001, // ldrb	w1, [x0]
		0x39002001, // strb	w1, [x0, #8]
		0x39400001, // ldrb	w1, [x0]
		0x13001c21, // sxtb	w1, w1
		0x79001001, // strh	w1, [x0, #8]
		0x39400001, // ldrb	w1, [x0]
		0x13001c21, // sxtb	w1, w1
		0xb9000801, // str	w1, [x0, #8]
		0x39400001, // ldrb	w1, [x0]
		0x93401c21, // sxtb	x1, w1
		0xf9000401, // str	x1, [x0, #8]
		0xd503201f, // nop
		0x79400001, // ldrh	w1, [x0]
		0x39002001, // strb	w1, [x0, #8]
		0x79400001, // ldrh	w1, [x0]
		0x79001001, // strh	w1, [x0, #8]
		0x79400001, // ldrh	w1, [x0]
		0x13003c21, // sxth	w1, w1
		0xb9000801, // str	w1, [x0, #8]
		0x79400001, // ldrh	w1, [x0]
		0x93403c21, // sxth	x1, w1
		0xf9000401, // str	x1, [x0, #8]
		0xd503201f, // nop
		0xb9400001, // ldr	w1, [x0]
		0x39002001, // strb	w1, [x0, #8]
		0xb9400001, // ldr	w1, [x0]
		0x79001001, // strh	w1, [x0, #8]
		0xb9400001, // ldr	w1, [x0]
		0xb9000801, // str	w1, [x0, #8]
		0xb9400001, // ldr	w1, [x0]
		0x93407c21, // sxtw	x1, w1
		0xf9000401, // str	x1, [x0, #8]
		0xd503201f, // nop
		0xf9400001, // ldr	x1, [x0]
		0x39002001, // strb	w1, [x0, #8]
		0xf9400001, // ldr	x1, [x0]
		0x79001001, // strh	w1, [x0, #8]
		0xf9400001, // ldr	x1, [x0]
		0xb9000801, // str	w1, [x0, #8]
		0xf9400001, // ldr	x1, [x0]
		0xf9000401, // str	x1, [x0, #8]
		0xd503201f, // nop
		0x39400001, // ldrb	w1, [x0]
		0x1103fc21, // add	w1, w1, #0xff
		0x39002001, // strb	w1, [x0, #8]
		0x39400001, // ldrb	w1, [x0]
		0x12001c21, // and	w1, w1, #0xff
		0x79001001, // strh	w1, [x0, #8]
		0x39400001, // ldrb	w1, [x0]
		0x12001c21, // and	w1, w1, #0xff
		0xb9000801, // str	w1, [x0, #8]
		0x39400001, // ldrb	w1, [x0]
		0x92401c21, // and	x1, x1, #0xff
		0xf9000401, // str	x1, [x0, #8]
		0xd503201f, // nop
		0x79400001, // ldrh	w1, [x0]
		0x39002001, // strb	w1, [x0, #8]
		0x79400001, // ldrh	w1, [x0]
		0x1103fc21, // add	w1, w1, #0xff
		0x79001001, // strh	w1, [x0, #8]
		0x79400001, // ldrh	w1, [x0]
		0x12003c21, // and	w1, w1, #0xffff
		0xb9000801, // str	w1, [x0, #8]
		0x79400001, // ldrh	w1, [x0]
		0x92403c21, // and	x1, x1, #0xffff
		0xf9000401, // str	x1, [x0, #8]
		0xd503201f, // nop
		0xb9400001, // ldr	w1, [x0]
		0x39002001, // strb	w1, [x0, #8]
		0xb9400001, // ldr	w1, [x0]
		0x79001001, // strh	w1, [x0, #8]
		0xb9400001, // ldr	w1, [x0]
		0x1103fc21, // add	w1, w1, #0xff
		0xb9000801, // str	w1, [x0, #8]
		0xb9400001, // ldr	w1, [x0]
		0x2a0103e1, // mov	w1, w1
		0xf9000401, // str	x1, [x0, #8]
		0xd503201f, // nop
		0xf9400001, // ldr	x1, [x0]
		0x39002001, // strb	w1, [x0, #8]
		0xf9400001, // ldr	x1, [x0]
		0x79001001, // strh	w1, [x0, #8]
		0xf9400001, // ldr	x1, [x0]
		0xb9000801, // str	w1, [x0, #8]
		0xf9400001, // ldr	x1, [x0]
		0x9103fc21, // add	x1, x1, #0xff
		0xf9000401, // str	x1, [x0, #8]
		0xd503201f, // nop
		0xd65f03c0, // ret
	)

	if !actual.Equal(expected) {
		t.Errorf("bad assembled code:\n\texpected %s\n\tactual   %s",
			expected, actual)
	}
}

func TestSoftRegId(t *testing.T) {
	var asm Asm
	asm.InitArch(Arm64{})

	var a, b, c SoftReg = MakeSoftReg(0, Uint64), MakeSoftReg(1, Uint64), MakeSoftReg(2, Uint64)
	asm.Assemble(
		ALLOC, a,
		ALLOC, b,
		ALLOC, c,
		MOV, ConstUint64(1), a,
		MOV, ConstUint64(2), b,
		ADD3, a, b, c,
		FREE, a,
		FREE, b,
		FREE, c,
	).Epilogue()

	actual := asm.Code()
	expected := MakeCode(
		0xd2800020, // movz	x0, #0x1
		0xd2800041, // movz	x1, #0x2
		0x8b010002, // add	x2, x0, x1
		0xd65f03c0, // ret
	)

	if !actual.Equal(expected) {
		t.Errorf("miscompiled code:\n\texpected %s\n\tactual   %s",
			expected, actual)
	}
}
