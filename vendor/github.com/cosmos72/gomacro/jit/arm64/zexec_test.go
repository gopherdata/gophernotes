// +build arm64

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
 * z_exec_test.go
 *
 *  Created on Feb 07, 2019
 *      Author Massimiliano Ghilardi
 */

package arm64

import (
	"testing"
)

func Param(offset int32, kind Kind) Mem {
	return MakeMem(offset, XSP, kind)
}

func VarKind(idx int64, kind Kind) Mem {
	return MakeMem(int32(idx)*8, X29, kind)
}

func Var(idx int64) Mem {
	return MakeMem(int32(idx)*8, X29, Int64)
}

func Init(asm *Asm) *Asm {
	asm.InitArch(Arm64{})
	asm.RegIncUse(X29)
	asm.Load(Param(8, Uint64), MakeReg(X29, Uint64))
	return asm
}

func TestExecNop(t *testing.T) {
	var f func()
	var asm Asm
	asm.InitArch(Arm64{}).Func(&f)
	f()
}

func TestExecZero(t *testing.T) {
	var f func() uint64
	var asm Asm
	asm.InitArch(Arm64{})

	asm.Assemble( //
		ZERO, Param(8, Uint64),
	).Func(&f)

	actual := f()
	expected := uint64(0)
	if actual != expected {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestExecConst(t *testing.T) {
	var f func() uint64
	var asm Asm
	var expected uint64 = 7

	asm.InitArch(Arm64{})
	asm.Assemble( //
		MOV, ConstUint64(expected), Param(8, Uint64),
	).Func(&f)

	actual := f()
	if actual != expected {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestExecLoadStore(t *testing.T) {
	var f func() uint64
	var asm Asm
	var expected uint64 = 0x12345678abcdef0

	r := asm.InitArch(Arm64{}).RegAlloc(Uint64)
	asm.Assemble( //
		MOV, ConstUint64(expected), r,
		MOV, r, Param(8, Uint64),
	).Func(&f)

	actual := f()
	if actual != expected {
		t.Errorf("expected 0x%x, actual 0x%x", expected, actual)
	}
}

func TestExecUnary(t *testing.T) {
	var c uint64 = 0x64776657f7754abc
	binds := [...]uint64{c}

	var asm Asm
	r := Init(&asm).RegAlloc(Uint64)
	v := VarKind(0, Uint64)

	var f func(*uint64)
	asm.Assemble( //
		MOV, v, r,
		NEG1, r,
		NOT1, r,
		MOV, r, v,
	).Func(&f)
	f(&binds[0])

	expected := ^-c
	actual := binds[0]

	if actual != expected {
		t.Errorf("expected 0x%x, actual 0x%x", expected, actual)
	}
}

func TestExecDiv(t *testing.T) {
	var f func(*int64)
	var asm Asm
	v0, v1, v2 := Var(0), Var(1), Var(2)

	Init(&asm)
	asm.Assemble(DIV3, v0, v1, v2).Func(&f)

	for a := int64(-5); a < 5; a++ {
		for b := int64(-5); b < 5; b++ {
			callDiv(t, a, b, f)
		}
	}
	const maxint64 = int64(^uint64(0) >> 1)
	const minint64 = ^maxint64

	for a := int64(-5); a < 5; a++ {
		for b := int64(-5); b < 5; b++ {
			callDiv(t, a+maxint64, b, f)
		}
	}
}

func callDiv(t *testing.T, a int64, b int64, f func(*int64)) {
	ints := [3]int64{a, b, ^int64(0)}
	f(&ints[0])
	var c int64
	if b != 0 {
		c = a / b
	}
	if ints[2] != c {
		t.Errorf("Div %v %v returned %v, expecting %d", a, b, ints[2], c)
	} else {
		t.Logf("%v / %v = %v", a, b, c)
	}
}
