/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * expr1.go
 *
 *  Created on Apr 03, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/constant"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) litValue(value I) Lit {
	return Lit{Type: c.TypeOf(value), Value: value}
}

func (c *Comp) exprUntypedLit(kind r.Kind, obj constant.Value) *Expr {
	return &Expr{Lit: Lit{Type: c.TypeOfUntypedLit(), Value: MakeUntypedLit(kind, obj, &c.Universe.BasicTypes)}}
}

func (c *Comp) exprValue(t xr.Type, value I) *Expr {
	if t == nil {
		t = c.TypeOf(value)
	}
	return exprValue(t, value)
}

func exprValue(t xr.Type, value I) *Expr {
	if t == nil {
		base.Errorf("internal error! exprValue() value = %v invoked with type = nil", value)
	}
	return &Expr{Lit: Lit{Type: t, Value: value}, EFlags: EFlag4Value(value)}
}

func exprLit(lit Lit, sym *Symbol) *Expr {
	return &Expr{Lit: lit, Sym: sym, EFlags: EFlag4Value(lit.Value)}
}

func exprFun(t xr.Type, fun I) *Expr {
	return &Expr{Lit: Lit{Type: t}, Fun: fun}
}

func exprX1(t xr.Type, fun func(env *Env) r.Value) *Expr {
	return &Expr{Lit: Lit{Type: t}, Fun: fun}
}

func exprXV(types []xr.Type, fun func(env *Env) (r.Value, []r.Value)) *Expr {
	if len(types) == 1 {
		return &Expr{Lit: Lit{Type: types[0]}, Fun: fun}
	} else {
		return &Expr{Lit: Lit{Type: types[0]}, Types: types, Fun: fun}
	}
}

func expr0(fun func(env *Env)) *Expr {
	return &Expr{Types: zeroTypes, Fun: fun}
}

func (c *Comp) exprBool(fun func(env *Env) bool) *Expr {
	return &Expr{Lit: Lit{Type: c.TypeOfBool()}, Fun: fun}
}

func (c *Comp) exprUint8(fun func(env *Env) uint8) *Expr {
	return &Expr{Lit: Lit{Type: c.TypeOfUint8()}, Fun: fun}
}

func (c *Comp) exprString(fun func(env *Env) string) *Expr {
	return &Expr{Lit: Lit{Type: c.TypeOfString()}, Fun: fun}
}

func (expr *Expr) EvalConst(opts CompileOptions) I {
	if expr == nil {
		return nil
	}
	if expr.Const() {
		if opts == COptDefaults && expr.Untyped() {
			return expr.ConstTo(expr.DefaultType())
		}
		return expr.Value
	}
	ret := expr.AsX1()(nil)
	if ret == base.None {
		base.Errorf("constant should evaluate to a single value, found no values at all")
		return nil
	}
	var value I
	if ret != base.Nil {
		value = ret.Interface()
	}
	expr.Value = value
	expr.EFlags = EFlag4Value(value)
	expr.Fun = nil // no longer needed, will be recreated if needed as a wrapper for the computed value
	return value
}
