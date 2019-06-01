// +build !gomacro_jit

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
 * jit_stub.go
 *
 *  Created on May 05, 2019
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/token"

	xr "github.com/cosmos72/gomacro/xreflect"
)

type jitExpr struct{}

type Jit struct{}

func NewJit() *Jit {
	return nil
}

// if supported, set e.Jit to jit constant == e.Lit.Value
// always returns e.
func (j *Jit) Const(e *Expr) *Expr {
	return e
}

// if supported, set e.Jit to jit expression that will compute xe
// always returns e.
func (j *Jit) Identity(e *Expr, xe *Expr) *Expr {
	return e
}

// if supported, set e.Jit to jit expression that will compute t(xe)
// always returns e.
func (j *Jit) Cast(e *Expr, t xr.Type, xe *Expr) *Expr {
	return e
}

// if supported, set e.Jit to jit expression that will compute *xe
// always returns e.
func (j *Jit) Deref(e *Expr, xe *Expr) *Expr {
	return e
}

// if supported, set e.Jit to jit expression that will compute op xe
// always returns e.
func (j *Jit) UnaryExpr(e *Expr, op token.Token, xe *Expr) *Expr {
	return e
}

// if supported, set e.Jit to jit expression that will compute xe op ye
// always returns e.
func (j *Jit) BinaryExpr(e *Expr, op token.Token, xe *Expr, ye *Expr) *Expr {
	return e
}

// if supported, set e.Jit to jit expression that will read local variable
// always returns e.
func (j *Jit) Symbol(e *Expr) *Expr {
	return e
}

// if supported, return a jit-compiled statement that will perform va OP= init
// return nil on failure
func (j *Jit) SetVar(va *Var, op token.Token, init *Expr) Stmt {
	return nil
}

// if supported, return a jit-compiled Stmt that will evaluate Expr.
// return nil on failure
func (j *Jit) AsStmt(e *Expr) Stmt {
	return nil
}

// if supported, replace e.Fun with a jit-compiled equivalent function.
// always returns e.
func (j *Jit) Fun(e *Expr) *Expr {
	return e
}
