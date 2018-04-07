/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * interpreter.go
 *
 *  Created on: Apr 02, 2017
 *      Author: Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	"github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// Interp is the fast interpreter.
// It contains both the tree-of-closures builder Comp
// and the interpreter's runtime environment Env
type Interp struct {
	Comp *Comp
	env  *Env // not exported. to access it, call CompEnv.PrepareEnv()
}

func (ir *Interp) RunExpr1(e *Expr) r.Value {
	if e == nil {
		return None
	}
	env := ir.PrepareEnv()
	fun := e.AsX1()
	return fun(env)
}

func (ir *Interp) RunExpr(e *Expr) (r.Value, []r.Value) {
	if e == nil {
		return None, nil
	}
	env := ir.PrepareEnv()
	fun := e.AsXV(ir.Comp.CompileOptions)
	return fun(env)
}

func (ir *Interp) Parse(src string) ast2.Ast {
	return ir.Comp.Parse(src)
}

// combined Parse + Compile
func (ir *Interp) Compile(src string) *Expr {
	c := ir.Comp
	return c.Compile(c.Parse(src))
}

func (ir *Interp) CompileNode(node ast.Node) *Expr {
	return ir.Comp.CompileNode(node)
}

func (ir *Interp) CompileAst(form ast2.Ast) *Expr {
	return ir.Comp.Compile(form)
}

// combined Parse + Compile + RunExpr
func (ir *Interp) Eval(src string) (r.Value, []r.Value) {
	c := ir.Comp
	return ir.RunExpr(c.Compile(c.Parse(src)))
}

func (ir *Interp) ChangePackage(name, path string) {
	if len(path) == 0 {
		path = name
	} else {
		name = FileName(path)
	}
	c := ir.Comp
	currpath := c.Path
	if path == currpath {
		return
	}
}

// DeclConst compiles a constant declaration
func (ir *Interp) DeclConst(name string, t xr.Type, value I) {
	ir.Comp.DeclConst0(name, t, value)
}

// DeclFunc compiles a function declaration
func (ir *Interp) DeclFunc(name string, fun I) {
	ir.Comp.DeclFunc0(name, fun)
	ir.apply()
}

// DeclBuiltin compiles a builtin function declaration
func (ir *Interp) DeclBuiltin(name string, builtin Builtin) {
	ir.Comp.DeclBuiltin0(name, builtin)
}

// DeclEnvFunc compiles a function declaration that accesses interpreter's *CompEnv
func (ir *Interp) DeclEnvFunc(name string, function Function) {
	ir.Comp.DeclEnvFunc0(name, function)
	ir.apply()
}

// DeclType declares a type
func (ir *Interp) DeclType(t xr.Type) {
	ir.Comp.DeclType0(t)
}

// DeclType declares a type alias
func (ir *Interp) DeclTypeAlias(alias string, t xr.Type) {
	ir.Comp.DeclTypeAlias0(alias, t)
}

// DeclVar compiles a variable declaration
func (ir *Interp) DeclVar(name string, t xr.Type, value I) {
	if t == nil {
		t = ir.Comp.TypeOf(value)
	}
	ir.Comp.DeclVar0(name, t, ir.Comp.exprValue(t, value))
	ir.apply()
}

// apply executes the compiled declarations, statements and expressions,
// then clears the compiled buffer
func (ir *Interp) apply() {
	exec := ir.Comp.Code.Exec()
	if exec != nil {
		exec(ir.PrepareEnv())
	}
}

// AddressOfVar compiles the expression &name, then executes it
// returns the zero value if name is not found or is not addressable
func (ir *Interp) AddressOfVar(name string) (addr r.Value) {
	c := ir.Comp
	sym := c.TryResolve(name)
	if sym != nil {
		switch sym.Desc.Class() {
		case VarBind, IntBind:
			va := sym.AsVar(PlaceAddress)
			expr := va.Address(c.Depth)
			return ir.RunExpr1(expr)
		}
	}
	return Nil
}

// replacement of reflect.TypeOf() that uses xreflect.TypeOf()
func (ir *Interp) TypeOf(val interface{}) xr.Type {
	return ir.Comp.TypeOf(val)
}

// ValueOf retrieves the value of a constant, function or variable
// The returned value is settable and addressable only for variables
// returns the zero value if name is not found
func (ir *Interp) ValueOf(name string) (value r.Value) {
	sym := ir.Comp.TryResolve(name)
	if sym == nil {
		return Nil
	}
	switch sym.Desc.Class() {
	case ConstBind:
		return sym.Bind.ConstValue()
	case IntBind:
		value = ir.AddressOfVar(name)
		if value != Nil {
			value = value.Elem() // dereference
		}
		return value
	case VarBind:
		env := ir.PrepareEnv()
		for i := 0; i < sym.Upn; i++ {
			env = env.Outer
		}
		return env.Binds[sym.Desc.Index()]
	default:
		expr := ir.Comp.Symbol(sym)
		return ir.RunExpr1(expr)
	}
}

func (ir *Interp) PrepareEnv() *Env {
	return ir.prepareEnv(128)
}

func (ir *Interp) prepareEnv(minDelta int) *Env {
	c := ir.Comp
	env := ir.env
	// usually we know at Env creation how many slots are needed in c.Env.Binds
	// but here we are modifying an existing Env...
	if minDelta < 0 {
		minDelta = 0
	}
	capacity, min := cap(env.Binds), c.BindNum
	// c.Debugf("prepareEnv() before: c.BindNum = %v, minDelta = %v, len(env.Binds) = %v, cap(env.Binds) = %v, env = %p", c.BindNum, minDelta, len(env.Binds), cap(env.Binds), env)

	if capacity < min {
		if capacity <= min/2 {
			capacity = min
		} else {
			capacity *= 2
		}
		if capacity-min < minDelta {
			capacity = min + minDelta
		}
		binds := make([]r.Value, min, capacity)
		copy(binds, env.Binds)
		env.Binds = binds
	}
	if len(env.Binds) < min {
		env.Binds = env.Binds[0:min:cap(env.Binds)]
	}
	// c.Debugf("prepareEnv() after:  c.BindNum = %v, minDelta = %v, len(env.Binds) = %v, cap(env.Binds) = %v, env = %p", c.BindNum, minDelta, len(env.Binds), cap(env.Binds), env)

	capacity, min = cap(env.IntBinds), c.IntBindNum
	if capacity < min {
		if capacity <= min/2 {
			capacity = min
		} else {
			capacity *= 2
		}
		if capacity-min < minDelta {
			capacity = min + minDelta
		}
		binds := make([]uint64, min, capacity)
		copy(binds, env.IntBinds)
		env.IntBinds = binds
	}
	if len(env.IntBinds) < min {
		env.IntBinds = env.IntBinds[0:min:cap(env.IntBinds)]
	}
	return env
}
