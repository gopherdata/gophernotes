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
 * zcompile_test.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package disasm

import (
	"testing"

	. "github.com/cosmos72/gomacro/jit"
	"github.com/cosmos72/gomacro/jit/asm"
)

const (
	t0 = SoftReg(FirstTempRegId+iota)<<8 | SoftReg(Uint64)
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

func TestCompileExpr1(t *testing.T) {
	var c Comp
	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)
		r := MakeReg(c.RLo, Uint64)
		e := NewExpr1(
			NEG, NewExpr1(NOT, r),
		)
		c.Expr(e)
		actual := c.Code()

		t.Log("expr: ", e)

		expected := Code{
			asm.ALLOC, t0,
			asm.NOT2, r, t0,
			asm.NEG2, t0, t0,
		}

		if i := CompareCode(actual, expected); i >= 0 {
			t.Errorf("miscompiled code at index %d:\n\texpected %v\n\tactual   %v",
				i, expected, actual)
		} else {
			t.Log("compiled to:", actual)
		}

		c.Epilogue()
		PrintDisasm(t, c.Assemble())
	}
}

func TestCompileExpr2(t *testing.T) {
	var c Comp
	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)
		a := c.Asm()

		c7 := MakeConst(7, Uint64)
		c9 := MakeConst(9, Uint64)
		r1 := a.RegAlloc(Uint64)
		r2 := a.RegAlloc(Uint64)
		// compile
		e := NewExpr2(
			ADD, NewExpr2(MUL, c7, r1), NewExpr2(SUB, c9, r2),
		)
		c.Expr(e)
		actual := c.Code()

		t.Log("expr: ", e)

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
			t.Log("compiled to:", actual)
		}

		c.Epilogue()
		PrintDisasm(t, c.Assemble())
	}

}

func TestCompileExpr3(t *testing.T) {
	var c Comp

	t0 := MakeSoftReg(FirstTempRegId, Int64)

	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)

		c_2 := MakeConst(-2, Int64)
		m := c.MakeVar(0, Int64)
		// compile
		e := NewExpr2(
			AND_NOT,
			NewExpr2(DIV, m, c_2),
			m,
		)
		c.Expr(e)
		actual := c.Code()

		t.Log("expr: ", e)

		expected := Code{
			asm.ALLOC, t0,
			asm.DIV3, m, c_2, t0,
			asm.AND_NOT3, t0, m, t0,
		}

		if i := CompareCode(actual, expected); i >= 0 {
			t.Errorf("miscompiled code at index %d:\n\texpected %v\n\tactual   %v",
				i, expected, actual)
		} else {
			t.Log("compiled to:", actual)
		}

		c.Epilogue()
		PrintDisasm(t, c.Assemble())
	}
}

func TestCompileStmt1(t *testing.T) {
	var c Comp
	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)

		m1 := c.MakeVar(0, Uint64)
		m2 := c.MakeVar(1, Uint32)
		m3w := c.MakeVar(2, Uint16)
		m3 := c.MakeVar(2, Uint8)
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
			// asm.NOP, m4w, // NOP is optimized away
			asm.MOV, m3w, m4w,
		}

		if i := CompareCode(actual, expected); i >= 0 {
			t.Errorf("miscompiled code at index %d:\n\texpected %v\n\tactual   %v",
				i, expected, actual)
		} else {
			t.Log("compiled to:", actual)
		}

		c.Epilogue()
		PrintDisasm(t, c.Assemble())
	}
}

func TestCompileStmt2(t *testing.T) {
	var c Comp

	_7 := MakeConst(7, Int64)
	_5 := MakeConst(5, Int64)
	t0 := MakeSoftReg(FirstTempRegId, Int64)

	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)
		s0 := c.NewSoftReg(Int64)
		s1 := c.NewSoftReg(Int64)

		source := Source{
			ASSIGN, s0,
			NewExpr2(SUB,
				NewExpr2(MUL, s1, _7),
				NewExpr2(DIV, s1, _5),
			),
		}
		c.Compile(source)
		actual := c.Code()

		t.Log("source:", source)

		expected := Code{
			asm.ALLOC, s0,
			asm.ALLOC, s1,
			asm.MUL3, s1, _7, s0,
			asm.ALLOC, t0,
			asm.DIV3, s1, _5, t0,
			asm.SUB3, s0, t0, s0,
			asm.FREE, t0,
		}

		if i := CompareCode(actual, expected); i >= 0 {
			t.Errorf("miscompiled code at index %d:\n\texpected %v\n\tactual   %v",
				i, expected, actual)
		} else {
			t.Log("compiled to:", actual)
		}

		c.Epilogue()
		PrintDisasm(t, c.Assemble())
	}
}

func TestCompileGetidx(t *testing.T) {
	var c Comp

	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)

		stack_in := c.MakeParam(8, Uint64)
		stack_out := c.MakeParam(16, Uint8)
		_42 := ConstUint8(42)
		idx := ConstUint64(3)

		c.Asm().RegIncUse(c.RegIdConfig.RVAR)
		// on amd64 and arm64, in a func(env *Env) ...
		// the parameter env is on the stack at [RSP+8]
		rvar := MakeReg(c.RegIdConfig.RVAR, Uint64)
		// env = stack[env_param]
		c.Stmt2(ASSIGN, rvar, stack_in)
		// rvar = env.Ints equivalent to rvar = &env.Ints[0]
		c.Stmt2(ASSIGN, rvar, NewExprIdx(rvar, idx, Uint64))
		// compile accumulated jit expression and copy result to stack.
		e := _42
		// on amd64 and arm64, in a func(env *Env) ...
		// the return value is on the stack at [RSP+16]
		c.Stmt2(ASSIGN, stack_out, e)

		actual := c.Code()

		expected := Code{
			asm.MOV, stack_in, rvar,
			asm.GETIDX, rvar, idx, rvar,
			asm.MOV, _42, stack_out,
		}

		if i := CompareCode(actual, expected); i >= 0 {
			t.Errorf("miscompiled code at index %d:\n\texpected %v\n\tactual   %v",
				i, expected, actual)
		} else {
			t.Log("compiled to:", actual)
		}

		c.Epilogue()
		PrintDisasm(t, c.Assemble())
	}
}

func TestCompileInterpStmtNop(t *testing.T) {
	var c Comp
	envIP := ConstUint64(7)
	envCode := ConstUint64(8)
	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)
		renv := MakeSoftReg(0, Uint64)
		rip := MakeSoftReg(1, Uint64)
		rcode := MakeSoftReg(2, Uint64)
		source := Source{
			ALLOC, renv,
			ALLOC, rip,
			ALLOC, rcode,
			// on amd64 and arm64, in a func(env *Env) ...
			// the parameter env is on the stack at [RSP+8]
			// renv = stack[env_param]
			ASSIGN, renv, c.MakeParam(8, Uint64),
			// rip = env.IP
			ASSIGN, rip, NewExprIdx(renv, envIP, Uint64),
			// rip++
			INC, rip,
			// env.IP = rip
			IDX_ASSIGN, renv, envIP, rip,
			// s = env.Code
			ASSIGN, rcode, NewExprIdx(renv, envCode, Uint64),
			// s = s[rip] i.e. s = env.Code[rip] i.e. s = env.Code[env.IP+1]
			ASSIGN, rcode, NewExprIdx(rcode, rip, Uint64),
			// stack[env_result] = renv
			ASSIGN, c.MakeParam(24, Uint64), renv,
			// stack[stmt_result] = s, with s == env.Code[env.IP+1]
			ASSIGN, c.MakeParam(16, Uint64), rcode,
			FREE, renv,
			FREE, rip,
			FREE, rcode,
		}
		c.Compile(source)
		c.Epilogue()

		t.Log("source:", source)
		t.Log("compiled to:", c.Code())

		PrintDisasm(t, c.Assemble())
	}
}
