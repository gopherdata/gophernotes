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
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"testing"

	"github.com/cosmos72/gomacro/jit/asm"
)

const (
	t0 SoftReg = SoftReg(FirstTempRegId+iota)<<8 | SoftReg(Uint64)
	t1
)

func CompareCode(actual Code, expected Code) int {

	if n1, n2 := len(actual), len(expected); n1 != n2 {
		if n1 < n2 {
			return n1
		}
		return n2
	}
	for i := range actual {
		if actual[i] != expected[i] {
			return i
		}
	}
	return -1
}

func TestExpr1(t *testing.T) {
	var c Comp
	c.Init()
	r := MakeReg(c.RLo, Uint64)
	e := NewExpr1(
		NEG, NewExpr1(NOT, r),
	)
	c.Expr(e)
	t.Logf("expr: %v", e)
	actual := c.code
	expected := Code{
		asm.ALLOC, t0,
		asm.NOT2, r, t0,
		asm.NEG2, t0, t0,
	}

	if i := CompareCode(actual, expected); i >= 0 {
		t.Errorf("miscompiled code at index %d:\n\texpected %v\n\tactual   %v",
			i, expected, actual)
	} else {
		t.Log("compiled to", actual)
	}
}

func TestExpr2(t *testing.T) {
	var c Comp
	c7 := MakeConst(7, Uint64)
	c9 := MakeConst(9, Uint64)
	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)
		r1 := MakeReg(c.RLo, Uint64)
		r2 := MakeReg(c.RLo+1, Uint64)
		e := NewExpr2(
			ADD, NewExpr2(MUL, c7, r1), NewExpr2(SUB, c9, r2),
		)
		c.Expr(e)
		t.Logf("expr2: %v", e)

		actual := c.code
		expected := Code{
			asm.ALLOC, t0,
			asm.MUL3, c7, r1, t0,
			asm.ALLOC, t1,
			asm.SUB3, c9, r2, t1,
			asm.ADD3, t0, t1, t0,
			asm.FREE, t1,
		}

		if i := CompareCode(actual, expected); i >= 0 {
			t.Errorf("miscompiled code at index %d:\n\texpected %v\n\tactual   %v",
				i, expected, actual)
		} else {
			t.Log("compiled to", archId, actual)
		}
	}
}

func TestCompileStmt1(t *testing.T) {
	var c Comp
	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)

		m1 := c.MakeVar(0, Uint64)
		m2 := c.MakeVar(1, Uint32)
		m3 := c.MakeVar(2, Uint8)
		m3w := c.MakeVar(2, Uint16)
		m4w := c.MakeVar(3, Uint16)

		source := Source{
			INC, m1, // m1++
			DEC, m2, // m2--
			ZERO, m3, // m3 = 0
			ASSIGN, m3w, NewExpr1(UINT16, m3), // m3w = uint16(m3)
			NOP, m4w, // _ = m4w
			ASSIGN, m4w, m3w, // m4w = m3w
		}
		c.Compile(source)
		actual := c.Code()
		t.Logf("source: %v", source)
		expected := Code{
			asm.INC, m1,
			asm.DEC, m2,
			asm.ZERO, m3,
			asm.CAST, m3, m3w,
			// asm.NOP, m4, // NOP is optimized away
			asm.MOV, m3w, m4w,
		}

		if i := CompareCode(actual, expected); i >= 0 {
			t.Errorf("miscompiled code at index %d:\n\texpected %v\n\tactual   %v",
				i, expected, actual)
		} else {
			t.Log("compiled to", archId, actual)
		}
	}
}

func TestCompileStmt2(t *testing.T) {
	var c Comp
	_7 := MakeConst(7, Uint64)
	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)
		s0 := c.NewSoftReg(Uint64)
		s1 := c.NewSoftReg(Uint64)

		stmt := NewStmt2(ASSIGN, s0,
			NewExpr1(NEG,
				NewExpr2(MUL, s1, _7)),
		)
		c.Stmt(stmt)
		actual := c.Code()
		t.Logf("stmt: %v", stmt)
		expected := Code{
			asm.ALLOC, s0,
			asm.ALLOC, s1,
			asm.MUL3, s1, _7, s0,
			asm.NEG2, s0, s0,
		}

		if i := CompareCode(actual, expected); i >= 0 {
			t.Errorf("miscompiled code at index %d:\n\texpected %v\n\tactual   %v",
				i, expected, actual)
		} else {
			t.Log("compiled to", archId, actual)
		}
	}
}
