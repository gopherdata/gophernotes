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
 * dsl.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

type Op uint8

const (
	LOAD Op = iota
	STORE

	ALLOC
	FREE

	ADD
	SUB
	MUL
	SDIV // signed   quotient
	UDIV // unsigned quotient
	SREM // signed   remainder
	UREM // unsigned remainder

	AND
	OR
	XOR
	ANDNOT

	NEG
	NOT
)

type divkind int

const (
	signed, unsigned divkind = 0, 1
	div, rem         divkind = 0, 2
)

func (asm *Asm) Asm(args ...interface{}) *Asm {
	n := len(args)
	for i := 0; i < n; i++ {
		op, ok := args[i].(Op)
		if !ok {
			errorf("syntax error: expecting OP [args], found %v", args[i])
		}
		i += asm.Op(op, args[i+1:]...)
	}
	return asm
}

func (asm *Asm) Op(op Op, args ...interface{}) int {
	var n int
	switch op {
	case LOAD, ADD, SUB, MUL, SDIV, UDIV, SREM, UREM, AND, OR, XOR, ANDNOT:
		if len(args) < 2 {
			errorf("syntax error: expecting OP arg1 arg2, found %v", append([]interface{}{op}, args...)...)
		}
		asm.Op2(op, args[0].(Reg), args[1].(Arg))
		n = 2
	case STORE:
		asm.Store(args[0].(*Var), args[1].(Reg))
		n = 2
	case ALLOC:
		asm.Alloc(args[0].(Reg))
		n = 1
	case FREE:
		asm.Free(args[0].(Reg))
		n = 1
	case NEG, NOT:
		if len(args) < 1 {
			errorf("syntax error: expecting OP arg1, found %v", op)
		}
		asm.Op1(op, args[0].(Reg))
		n = 1
	default:
		errorf("unknown operator: %v", op)
	}
	return n
}

func (asm *Asm) Op1(op Op, z Reg) *Asm {
	switch op {
	case NEG:
		asm.Neg(z)
	case NOT:
		asm.Not(z)
	default:
		errorf("unknown unary operator: %v", op)
	}
	return asm
}

func (asm *Asm) Op2(op Op, z Reg, a Arg) *Asm {
	switch op {
	case LOAD:
		asm.Load(z, a)
	case ADD:
		asm.Add(z, a)
	case SUB:
		asm.Sub(z, a)
	case MUL:
		asm.Mul(z, a)
	case SDIV:
		asm.SDiv(z, a)
	case UDIV:
		asm.UDiv(z, a)
	case SREM:
		asm.SRem(z, a)
	case UREM:
		asm.URem(z, a)
	case AND:
		asm.And(z, a)
	case OR:
		asm.Or(z, a)
	case XOR:
		asm.Xor(z, a)
	case ANDNOT:
		asm.Andnot(z, a)
	default:
		errorf("unknown binary operator: %v", op)
	}
	return asm
}
