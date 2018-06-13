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
 * call.go
 *
 *  Created on Apr 15, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type Call struct {
	Fun      *Expr
	Args     []*Expr
	OutTypes []xr.Type
	Builtin  bool // if true, call is a builtin function
	Const    bool // if true, call has no side effects and always returns the same result => it can be invoked at compile time
	Ellipsis bool // if true, must use reflect.Value.CallSlice or equivalent to invoke the function
}

func newCall1(fun *Expr, arg *Expr, isconst bool, outtypes ...xr.Type) *Call {
	return &Call{
		Fun:      fun,
		Args:     []*Expr{arg},
		OutTypes: outtypes,
		Const:    isconst,
	}
}

func (call *Call) MakeArgfunsX1() []func(*Env) r.Value {
	args := call.Args
	argfuns := make([]func(*Env) r.Value, len(args))
	for i, arg := range args {
		argfuns[i] = arg.AsX1()
	}
	return argfuns
}

// CallExpr compiles a function call or a type conversion
func (c *Comp) CallExpr(node *ast.CallExpr) *Expr {
	var fun *Expr
	if len(node.Args) == 1 {
		var t xr.Type
		fun, t = c.Expr1OrType(node.Fun)
		if t != nil {
			return c.Convert(node.Args[0], t)
		}
	}
	call := c.prepareCall(node, fun)
	return c.call_any(call)
}

// callExpr compiles the common part between CallExpr and Go statement
func (c *Comp) prepareCall(node *ast.CallExpr, fun *Expr) *Call {
	if fun == nil {
		fun = c.Expr1(node.Fun, nil)
	}
	t := fun.Type
	var builtin bool
	var lastarg *Expr
	if t.IdenticalTo(c.TypeOfBuiltin()) {
		return c.callBuiltin(node, fun)
	} else if t.IdenticalTo(c.TypeOfFunction()) {
		fun, lastarg = c.callFunction(node, fun)
		t = fun.Type
		builtin = true
	}
	if t.Kind() != r.Func {
		c.Errorf("call of non-function: %v <%v>", node.Fun, t)
		return nil
	}
	var args []*Expr
	if len(node.Args) == 1 {
		// support foo(bar()) where bar() returns multiple values
		arg := c.Expr(node.Args[0], nil)
		if arg.NumOut() == 0 {
			c.Errorf("function argument returns zero values: %v ", node.Args[0])
		}
		args = []*Expr{arg}
	} else {
		args = c.Exprs(node.Args)
	}
	if lastarg != nil {
		args = append(args, lastarg)
	}
	ellipsis := node.Ellipsis != token.NoPos
	c.checkCallArgs(node, t, args, ellipsis)

	outn := t.NumOut()
	outtypes := make([]xr.Type, outn)
	for i := 0; i < outn; i++ {
		outtypes[i] = t.Out(i)
	}
	return &Call{Fun: fun, Args: args, OutTypes: outtypes, Builtin: builtin, Ellipsis: ellipsis}
}

// call_any emits a compiled function call
func (c *Comp) call_any(call *Call) *Expr {
	expr := &Expr{}
	tout := call.OutTypes
	nout := len(tout)
	expr.SetTypes(tout)

	maxdepth := c.Depth
	// functions imported from other packages are constant too...
	// but call_builtin does not know about them
	if call.Fun.Const() {
		if call.Builtin {
			fun := c.call_builtin(call)
			if _, untyped := fun.(UntypedLit); untyped {
				// complex(), real(), imag() of untyped constants produce an untyped constant, not a function
				expr.Value = fun
				return expr
			} else {
				expr.Fun = fun
			}
		} else {
			// normal calls do not expect function to be a constant.
			call.Fun.WithFun()
		}
	}

	if expr.Fun != nil {
		// done already
	} else if len(call.Args) == 1 && call.Args[0].NumOut() > 1 {
		// support foo(bar()) where bar() returns multiple values.
		//
		// do NOT use this case for calls like fmt.Printf("foo") where the function
		// formally expects two args but is variadic => accepts one arg too:
		// fixes gophernotes issue 118
		expr.Fun = call_multivalue(call, maxdepth)
	} else if nout == 0 {
		expr.Fun = c.call_ret0(call, maxdepth)
	} else if nout == 1 {
		expr.Fun = c.call_ret1(call, maxdepth)
	} else {
		expr.Fun = c.call_ret2plus(call, maxdepth)
	}
	// constant propagation - only if function returns a single value
	if call.Const && len(call.OutTypes) == 1 {
		expr.EvalConst(COptDefaults)
		// c.Debugf("pre-computed result of constant call %v: %v <%v>", call, expr.Value, TypeOf(expr.Value))
	}
	return expr
}

func (c *Comp) checkCallArgs(node *ast.CallExpr, t xr.Type, args []*Expr, ellipsis bool) {
	variadic := t.IsVariadic()
	if ellipsis {
		if variadic {
			// a variadic function invoked as fun(x, y...)
			// behaves exactly as a non-variadic function call:
			// number and type of arguments must match
			variadic = false
		} else {
			c.Errorf("invalid use of ... in call to non-variadic function <%v>: %v", t, node)
			return
		}
	}
	n := t.NumIn()
	narg := len(args)
	if narg == 1 {
		// support foo(bar()) where bar() returns multiple values
		narg = args[0].NumOut()
	}
	if narg < n-1 || (!variadic && narg != n) {
		c.badCallArgNum(node.Fun, t, args)
		return
	}
	var ti, tlast xr.Type
	if variadic {
		tlast = t.In(n - 1).Elem()
	}
	var convs []func(r.Value) r.Value
	needconvs := false
	multivalue := len(args) != narg
	if multivalue {
		convs = make([]func(r.Value) r.Value, narg)
	}
	for i := 0; i < narg; i++ {
		if variadic && i >= n-1 {
			ti = tlast
		} else {
			ti = t.In(i)
		}
		if multivalue {
			// support foo(bar()) where bar() returns multiple values
			targ := args[0].Out(i)
			if targ == nil || !targ.AssignableTo(ti) {
				c.Errorf("cannot use <%v> as <%v> in argument to %v", targ, ti, node.Fun)
			} else if conv := c.Converter(targ, ti); conv != nil {
				convs[i] = conv
				args[0].Types[i] = ti
				needconvs = true
			}
			continue
		}
		// one argument per parameter: foo(arg1, arg2 /*...*/)
		arg := args[i]
		if arg.Const() {
			arg.ConstTo(ti)
		} else if arg.Type == nil || !arg.Type.AssignableTo(ti) {
			c.Errorf("cannot use <%v> as <%v> in argument to %v", arg.Type, ti, node.Fun)
		} else {
			arg.To(c, ti)
		}
	}
	if !multivalue || !needconvs {
		return
	}
	f := args[0].AsXV(COptDefaults)
	args[0].Fun = func(env *Env) (r.Value, []r.Value) {
		_, vs := f(env)
		for i, conv := range convs {
			if conv != nil {
				vs[i] = conv(vs[i])
			}
		}
		return vs[0], vs
	}
}

// mandatory optimization: fast_interpreter ASSUMES that expressions
// returning bool, int, uint, float, complex, string do NOT wrap them in reflect.Value
func (c *Comp) call_ret0(call *Call, maxdepth int) func(env *Env) {
	if call.Ellipsis {
		return call_ellipsis_ret0(call, maxdepth)
	} else if call.Fun.Type.IsVariadic() {
		return call_variadic_ret0(call, maxdepth)
	}
	// optimize fun(t1, t2)
	exprfun := call.Fun.AsX1()
	var ret func(*Env)
	switch len(call.Args) {
	case 0:
		ret = c.call0ret0(call, maxdepth)
	case 1:
		ret = c.call1ret0(call, maxdepth)
	case 2:
		ret = c.call2ret0(call, maxdepth)
	case 3:
		argfunsX1 := call.MakeArgfunsX1()
		argfuns := [3]func(*Env) r.Value{
			argfunsX1[0],
			argfunsX1[1],
			argfunsX1[2],
		}
		ret = func(env *Env) {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
				argfuns[2](env),
			}
			callxr(funv, argv)
		}
	}
	if ret == nil {
		argfunsX1 := call.MakeArgfunsX1()
		ret = func(env *Env) {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfunsX1))
			for i, argfun := range argfunsX1 {
				argv[i] = argfun(env)
			}
			callxr(funv, argv)
		}
	}
	return ret
}

// mandatory optimization: fast_interpreter ASSUMES that expressions
// returning bool, int, uint, float, complex, string do NOT wrap them in reflect.Value
func (c *Comp) call_ret1(call *Call, maxdepth int) I {
	if call.Ellipsis {
		return call_ellipsis_ret1(call, maxdepth)
	} else if call.Fun.Type.IsVariadic() {
		return call_variadic_ret1(call, maxdepth)
	}
	var ret I
	switch len(call.Args) {
	case 0:
		ret = c.call0ret1(call, maxdepth)
	case 1:
		ret = c.call1ret1(call, maxdepth)
	case 2:
		ret = c.call2ret1(call, maxdepth)
	default:
		ret = c.callnret1(call, maxdepth)
	}
	return ret
}

// cannot optimize much here... fast_interpreter ASSUMES that expressions
// returning multiple values actually return (reflect.Value, []reflect.Value)
func (c *Comp) call_ret2plus(call *Call, maxdepth int) func(env *Env) (r.Value, []r.Value) {
	if call.Ellipsis {
		return call_ellipsis_ret2plus(call, maxdepth)
	}
	// no need to special case variadic functions here
	expr := call.Fun
	exprfun := expr.AsX1()
	argfunsX1 := call.MakeArgfunsX1()
	var ret func(*Env) (r.Value, []r.Value)
	// slightly optimize fun() (tret0, tret1)
	switch len(call.Args) {
	case 0:
		ret = func(env *Env) (r.Value, []r.Value) {
			funv := exprfun(env)
			retv := callxr(funv, base.ZeroValues)
			return retv[0], retv
		}
	case 1:
		argfun := argfunsX1[0]
		ret = func(env *Env) (r.Value, []r.Value) {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)
			return retv[0], retv
		}
	case 2:
		argfuns := [2]func(*Env) r.Value{
			argfunsX1[0],
			argfunsX1[1],
		}
		ret = func(env *Env) (r.Value, []r.Value) {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			retv := callxr(funv, argv)
			return retv[0], retv
		}
	case 3:
		argfuns := [3]func(*Env) r.Value{
			argfunsX1[0],
			argfunsX1[1],
			argfunsX1[2],
		}
		ret = func(env *Env) (r.Value, []r.Value) {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
				argfuns[2](env),
			}
			retv := callxr(funv, argv)
			return retv[0], retv
		}
	default:
		// general case
		ret = func(env *Env) (r.Value, []r.Value) {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfunsX1))
			for i, argfun := range argfunsX1 {
				argv[i] = argfun(env)
			}
			retv := callxr(funv, argv)
			return retv[0], retv
		}
	}
	return ret
}

// replacement for reflect.Value.Call() that correctly handles
// functions wrapped in xr.Forward
func callxr(fun r.Value, args []r.Value) []r.Value {
	if fun.Kind() == r.Interface {
		fun = fun.Elem()
	}
	return fun.Call(args)
}

func callslicexr(fun r.Value, args []r.Value) []r.Value {
	if fun.Kind() == r.Interface {
		fun = fun.Elem()
	}
	return fun.CallSlice(args)
}

func (c *Comp) badCallArgNum(fun ast.Expr, t xr.Type, args []*Expr) *Call {
	prefix := "not enough"
	n := t.NumIn()
	nargs := len(args)
	if nargs > n {
		prefix = "too many"
	}
	have := bytes.Buffer{}
	for i, arg := range args {
		if i == 0 {
			fmt.Fprintf(&have, "%v", arg.Type)
		} else {
			fmt.Fprintf(&have, ", %v", arg.Type)
		}
	}
	want := bytes.Buffer{}
	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Fprintf(&want, "%v", t.In(i))
		} else {
			fmt.Fprintf(&want, ", %v", t.In(i))
		}
	}
	c.Errorf("%s arguments in call to %v:\n\thave (%s)\n\twant (%s)", prefix, fun, have.Bytes(), want.Bytes())
	return nil
}
