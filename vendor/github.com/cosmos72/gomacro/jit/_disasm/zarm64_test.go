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
 * zarm64_test.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package disasm

import (
	"testing"

	. "github.com/cosmos72/gomacro/jit/arm64"
)

func TestArm64Sample(T *testing.T) {
	var asm Asm

	for id := RLo; id+2 <= RHi; id++ {
		asm.InitArch(Arm64{})
		if asm.RegIsUsed(id) || asm.RegIsUsed(id+1) || asm.RegIsUsed(id+2) {
			continue
		}
		r := MakeReg(id+0, Int64)
		s := MakeReg(id+1, Int64)
		t := MakeReg(id+2, Int64)
		m := MakeMem(8, id, Int64)
		c := ConstInt64(0xFFF)
		one := ConstUint8(1)
		ur := MakeReg(id+0, Uint64)
		us := MakeReg(id+1, Uint64)
		ut := MakeReg(id+2, Uint64)
		br := MakeReg(id+0, Uint8)
		bt := MakeReg(id+2, Uint8)
		asm.RegIncUse(id)
		asm.RegIncUse(id + 1)
		asm.RegIncUse(id + 2)
		asm.Assemble(MOV, c, r, //
			MOV, c, m, //
			MOV, m, r, //
			NOP,           //
			ADD3, r, s, t, //
			SUB3, r, s, t, //
			AND3, r, s, t, //
			OR3, r, s, t, //
			XOR3, r, s, t, //
			SHL3, r, us, t, //
			SHR3, ur, us, ut, //
			SHR3, r, us, t, //
			NOP, //
			// test commutativity optimization
			ADD3, c, r, t, //
			SUB3, r, c, t, //
			AND3, c, r, t, //
			OR3, c, r, t, //
			XOR3, r, c, t, //
			SHL3, r, one, t, //
			SHR3, ur, one, ut, //
			SHR3, r, one, t, //
			NOP, //
			NOP, //
			// test 8-bit registers
			ADD3, one, br, bt, //
			SUB3, br, one, bt, //
			AND3, one, br, bt, //
			OR3, one, br, bt, //
			XOR3, br, one, bt, //
			SHL3, br, one, bt, //
			SHR3, br, one, bt, //
		).Epilogue()
		asm.RegDecUse(id)
		asm.RegDecUse(id + 1)
		asm.RegDecUse(id + 2)

		if id == RLo || id == RHi {
			PrintDisasm(T, asm.Code())
		}
	}
}

func TestArm64Zero(t *testing.T) {
	r := MakeReg(RLo, Uint64)
	xzr := MakeReg(XZR, Uint64)
	m := MakeMem(8, XSP, Uint64)

	var asm Asm
	asm.InitArch(Arm64{})
	asm.Assemble(
		ZERO, r,
		MOV, xzr, r,
		ZERO, m,
		RET)

	PrintDisasm(t, asm.Code())
}

func TestArm64Cast(t *testing.T) {
	var asm Asm
	asm.InitArch(Arm64{})
	for _, skind := range [...]Kind{
		Int8, Int16, Int32, Int64,
		Uint8, Uint16, Uint32, Uint64,
	} {

		src := MakeReg(RLo, skind)
		for _, dkind := range [...]Kind{Uint8, Uint16, Uint32, Uint64} {
			dst := MakeReg(RLo, dkind)
			asm.Assemble(CAST, src, dst)
		}
	}
	PrintDisasm(t, asm.Code())
}

func TestArm64Mem(t *testing.T) {
	var asm Asm
	asm.InitArch(Arm64{})

	id := RLo
	asm.RegIncUse(id)

	for _, skind := range [...]Kind{
		Int8, Int16, Int32, Int64,
		Uint8, Uint16, Uint32, Uint64,
	} {
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
	PrintDisasm(t, asm.Code())
}

func TestArm64Unary(t *testing.T) {

	var asm Asm
	asm.InitArch(Arm64{})
	r := MakeReg(X27, Uint64)
	s := MakeReg(X28, Uint64)
	v := MakeMem(0, X29, Uint64)

	asm.Assemble( //
		MOV, v, r,
		NEG2, r, s,
		NOT2, s, r,
		MOV, r, v,
	)
	asm.Epilogue()
	PrintDisasm(t, asm.Code())
}

func TestArm64Index(t *testing.T) {
	r := MakeReg(X0, Uint64)
	s := MakeReg(X1, Uint64)

	var asm Asm
	asm.InitArch(Arm64{})

	asm.RegIncUse(X0)
	asm.RegIncUse(X1)
	asm.RegIncUse(X2)

	for _, k := range [...]Kind{Uint8, Uint16, Uint32, Uint64} {
		v := MakeReg(X2, k)
		c := MakeConst(0x33, k)
		zero := MakeConst(0, k)
		asm.Assemble(
			GETIDX, r, s, v,
			SETIDX, r, s, v,
			GETIDX, r, c, v,
			SETIDX, r, c, v,
			SETIDX, r, s, zero,
			SETIDX, r, c, zero,
			NOP,
		)
	}
	PrintDisasm(t, asm.Code())
}

func TestArm64SoftReg(t *testing.T) {
	var asm Asm
	asm.InitArch(Arm64{})

	a := MakeSoftReg(0, Uint64)
	b := MakeSoftReg(1, Uint64)
	c := MakeSoftReg(2, Uint64)

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
	PrintDisasm(t, asm.Code())
}
