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
 * stmt.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"fmt"
)

type Stmt interface {
	stmt()
}

// unary statement X Inst,
// for example X++
type Stmt1 struct {
	Dst  Expr
	Inst Inst1
}

// binary statement X Inst Y,
// for example X += Y
type Stmt2 struct {
	Dst  Expr
	Src  Expr
	Inst Inst2
}

// ternary statement X Y Inst Z,
// for example X[Y] = Z
type Stmt3 struct {
	Dst  Expr
	DArg Expr
	Src  Expr
	Inst Inst3
}

// N-ary assignment
// a,b,c... = p,q,r...
type StmtN struct {
	Dst []Expr
	Src []Expr
}

func NewStmt1(inst Inst1, dst Expr) *Stmt1 {
	return &Stmt1{dst, inst}
}

func NewStmt2(inst Inst2, dst Expr, src Expr) *Stmt2 {
	return &Stmt2{dst, src, inst}
}

func NewStmt3(inst Inst3, dst Expr, darg Expr, src Expr) *Stmt3 {
	return &Stmt3{dst, darg, src, inst}
}

func NewStmtN(dst []Expr, src []Expr) *StmtN {
	return &StmtN{dst, src}
}

// implement Stmt interface
func (t *Stmt1) stmt() {
}

func (t *Stmt3) stmt() {
}

func (t *Stmt2) stmt() {
}

func (t *StmtN) stmt() {
}

func (t *Stmt1) String() string {
	switch t.Inst {
	case NOP:
		return fmt.Sprintf("_ = %v;", t.Dst)
	default:
		return fmt.Sprintf("%v%v;", t.Dst, t.Inst)
	}
}

func (t *Stmt2) String() string {
	return fmt.Sprintf("%v %v %v;", t.Dst, t.Inst, t.Src)
}

func (t *Stmt3) String() string {
	if t.Inst == IDX_ASSIGN {
		return fmt.Sprintf("%v[%v] = %v;", t.Dst, t.DArg, t.Src)
	}
	return fmt.Sprintf("%v %v %v %v;", t.Dst, t.DArg, t.Inst, t.Src)
}

// compile statement
func (c *Comp) Stmt(t Stmt) {
	switch t := t.(type) {
	case *Stmt1:
		c.Stmt1(t.Inst, t.Dst)
	case *Stmt2:
		c.Stmt2(t.Inst, t.Dst, t.Src)
	case *Stmt3:
		c.Stmt3(t.Inst, t.Dst, t.DArg, t.Src)
	case *StmtN:
		c.StmtN(t.Dst, t.Src)
	default:
		errorf("unknown Stmt type %T: %v", t, t)
	}
}

// compile unary statement
func (c *Comp) Stmt1(inst Inst1, tdst Expr) {
	dst, soft := c.Expr(tdst)
	if inst != NOP {
		checkAssignable(dst)
	}
	c.code.Inst1(inst, dst)
	c.freeTempReg(soft)
}

// compile binary statement
func (c *Comp) Stmt2(inst Inst2, tdst Expr, tsrc Expr) {
	// evaluate left-hand side first
	dst, dsoft := c.Expr(tdst)
	checkAssignable(dst)
	var dto Expr
	if inst == ASSIGN {
		// we can overwrite dst early
		// only if it's a plain ASSIGN
		dto = dst
	}
	src, ssoft := c.expr(tsrc, dto)
	c.code.Inst2(inst, dst, src)
	c.freeTempReg(dsoft)
	if ssoft.Id() != dsoft.Id() {
		c.freeTempReg(ssoft)
	}
}

// compile ternary statement
func (c *Comp) Stmt3(inst Inst3, tdst Expr, tdarg Expr, tsrc Expr) {
	// evaluate left-hand side first
	dst, dsoft := c.Expr(tdst)
	darg, dasoft := c.Expr(tdarg)
	src, ssoft := c.Expr(tsrc)
	c.code.Inst3(inst, dst, darg, src)
	c.freeTempReg(dsoft)
	c.freeTempReg(dasoft)
	c.freeTempReg(ssoft)
}

// compile n-ary statement
func (c *Comp) StmtN(tdst []Expr, tsrc []Expr) {
	n := len(tdst)
	if n != len(tsrc) {
		errorf("assignment mismatch: %d variables but %d values: %v = %v", n, len(tsrc), tdst, tsrc)
	}
	dst := make([]Expr, n)
	src := make([]Expr, n)
	// evaluate left-hand side first
	for i, x := range tdst {
		e, _ := c.Expr(x)
		checkAssignable(e)
		dst[i] = e
	}
	for i, x := range tsrc {
		e, soft := c.Expr(x)
		if _, ok := e.(Mem); ok && !soft.Valid() {
			// source is a local variable. we must evaluate it,
			// in case it also appears in left-hand side
			soft = c.newTempReg(e.Kind())
			c.code.Inst2(ASSIGN, soft, e)
			e = soft
		}
		src[i] = e
	}
	for i := range src {
		c.code.Inst2(ASSIGN, dst[i], src[i])
	}
	for i := n - 1; i >= 0; i-- {
		if soft, ok := src[i].(SoftReg); ok && soft.Valid() {
			c.freeTempReg(soft)
		}
	}
	for i := n - 1; i >= 0; i-- {
		if soft, ok := dst[i].(SoftReg); ok && soft.Valid() {
			c.freeTempReg(soft)
		}
	}
}
