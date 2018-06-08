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
 * convert.go
 *
 *  Created on Apr 30, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// Convert compiles a type conversion expression
func (c *Comp) Convert(node ast.Expr, t xr.Type) *Expr {
	e := c.Expr1(node, nil)

	return c.convert(e, t, node)
}

// Convert compiles a type conversion expression
func (c *Comp) convert(e *Expr, t xr.Type, nodeOpt ast.Expr) *Expr {
	if e.Untyped() {
		e.ConstTo(t)
	}

	if e.Type != nil && e.Type.IdenticalTo(t) {
		return e
	} else if e.Type != nil && e.Type.ReflectType() == t.ReflectType() {
		if e.Const() {
			return c.exprValue(t, e.Value)
		} else {
			return exprFun(t, e.Fun)
		}
	} else if e.Type == nil && IsNillableKind(t.Kind()) {
		e.Type = t
		e.Value = xr.Zero(t).Interface()
	} else if e.Type != nil && e.Type.ConvertibleTo(t) {
	} else {
		c.Errorf("cannot convert %v to %v: %v", e.Type, t, nodeOpt)
		return nil
	}
	rtype := t.ReflectType()
	if e.Const() {
		val := r.ValueOf(e.Value).Convert(rtype).Interface()
		return c.exprValue(t, val)
	}
	fun := e.AsX1()
	var ret I
	switch t.Kind() {
	case r.Bool:
		ret = func(env *Env) bool {
			val := fun(env).Convert(rtype)
			return val.Bool()
		}
	case r.Int:
		ret = func(env *Env) int {
			val := fun(env).Convert(rtype)
			return int(val.Int())
		}
	case r.Int8:
		ret = func(env *Env) int8 {
			val := fun(env).Convert(rtype)
			return int8(val.Int())
		}
	case r.Int16:
		ret = func(env *Env) int16 {
			val := fun(env).Convert(rtype)
			return int16(val.Int())
		}
	case r.Int32:
		ret = func(env *Env) int32 {
			val := fun(env).Convert(rtype)
			return int32(val.Int())
		}
	case r.Int64:
		ret = func(env *Env) int64 {
			val := fun(env).Convert(rtype)
			return val.Int()
		}
	case r.Uint:
		ret = func(env *Env) uint {
			val := fun(env).Convert(rtype)
			return uint(val.Uint())
		}
	case r.Uint8:
		ret = func(env *Env) uint8 {
			val := fun(env).Convert(rtype)
			return uint8(val.Uint())
		}
	case r.Uint16:
		ret = func(env *Env) uint16 {
			val := fun(env).Convert(rtype)
			return uint16(val.Uint())
		}
	case r.Uint32:
		ret = func(env *Env) uint32 {
			val := fun(env).Convert(rtype)
			return uint32(val.Uint())
		}
	case r.Uint64:
		ret = func(env *Env) uint64 {
			val := fun(env).Convert(rtype)
			return val.Uint()
		}
	case r.Uintptr:
		ret = func(env *Env) uintptr {
			val := fun(env).Convert(rtype)
			return uintptr(val.Uint())
		}
	case r.Float32:
		ret = func(env *Env) float32 {
			val := fun(env).Convert(rtype)
			return float32(val.Float())
		}
	case r.Float64:
		ret = func(env *Env) float64 {
			val := fun(env).Convert(rtype)
			return val.Float()
		}
	case r.Complex64:
		ret = func(env *Env) complex64 {
			val := fun(env).Convert(rtype)
			return complex64(val.Complex())
		}
	case r.Complex128:
		ret = func(env *Env) complex128 {
			val := fun(env).Convert(rtype)
			return val.Complex()
		}
	case r.String:
		ret = func(env *Env) string {
			val := fun(env).Convert(rtype)
			return val.String()
		}
	default:
		if conv := c.Converter(e.Type, t); conv != nil {
			ret = func(env *Env) r.Value {
				return conv(fun(env))
			}
		} else {
			ret = func(env *Env) r.Value {
				return fun(env)
			}
		}
	}
	eret := exprFun(t, ret)
	if e.Const() {
		eret.EvalConst(COptKeepUntyped)
	}
	return eret
}

// Converter returns a function that converts reflect.Value from tin to tout
// also supports conversion from interpreted types to interfaces
func (c *Comp) Converter(tin, tout xr.Type) func(r.Value) r.Value {
	if !tin.ConvertibleTo(tout) {
		c.Errorf("cannot convert from <%v> to <%v>", tin, tout)
	}
	rtin := tin.ReflectType()
	rtout := tout.ReflectType()
	switch {
	case rtin == rtout:
		return nil
	case rtin.ConvertibleTo(rtout):
		// most conversions, including from compiled type to compiled interface
		if rtin.Kind() != r.Interface {
			return func(obj r.Value) r.Value {
				return obj.Convert(rtout)
			}
		}
		// extract objects wrapped in proxies (if any)
		g := c.CompGlobals
		return func(obj r.Value) r.Value {
			obj, _ = g.extractFromProxy(obj)
			if obj.IsValid() {
				return obj.Convert(rtout)
			} else {
				return r.Zero(rtout)
			}
		}
	case xr.IsEmulatedInterface(tout):
		// conversion from type to emulated interface
		return c.converterToEmulatedInterface(tin, tout)
	case rtin == c.Universe.TypeOfForward.ReflectType():
		// conversion from forward-declared type
		return c.converterFromForward(tin, tout)
	case rtout.Kind() == r.Interface:
		// conversion from interpreted type to compiled interface.
		// must use a proxy that pre-implement compiled interfaces.
		return c.converterToProxy(tin, tout)
	case rtin.Kind() == r.Func && rtout.Kind() == r.Func:
		// conversion between func() and self-referencing named func type,
		// as for example type F func(F)
		return c.converterFunc(tin, tout)
	default:
		c.Errorf("unimplemented conversion from <%v> to <%v> with reflect.Type <%v> to <%v>",
			tin, tout, rtin, rtout)
		return nil
	}
}

// conversion from forward-declared type
func (c *Comp) converterFromForward(tin, tout xr.Type) func(r.Value) r.Value {
	rtout := tout.ReflectType()
	return func(val r.Value) r.Value {
		val = val.Elem()
		if val.Type() != rtout {
			val = val.Convert(rtout)
		}
		return val
	}
}

// conversion between func() and self-referencing named func type,
// as for example type F func(F)
func (c *Comp) converterFunc(tin, tout xr.Type) func(r.Value) r.Value {
	rtin := tin.ReflectType()
	rtout := tout.ReflectType()
	nin := rtin.NumIn()
	nout := rtin.NumOut()
	if nin != rtout.NumIn() || nout != rtout.NumOut() || rtin.IsVariadic() != rtout.IsVariadic() {

		c.Errorf("unimplemented conversion from <%v> to <%v> with reflect.Type <%v> to <%v>",
			tin, tout, rtin, rtout)
	}
	convarg := make([]func(r.Value) r.Value, nin)
	for i := 0; i < nin; i++ {
		// arguments must be adapted to actual func type: rtin
		convarg[i] = c.Converter(tout.In(i), tin.In(i))
	}
	convret := make([]func(r.Value) r.Value, nout)
	for i := 0; i < nout; i++ {
		// results must be adapted to expected func type: rtout
		convret[i] = c.Converter(tin.Out(i), tout.Out(i))
	}
	return func(f r.Value) r.Value {
		return r.MakeFunc(rtout, func(args []r.Value) []r.Value {
			for i, conv := range convarg {
				if conv != nil {
					args[i] = conv(args[i])
				}
			}
			rets := f.Call(args)
			for i, conv := range convret {
				if conv != nil {
					rets[i] = conv(rets[i])
				}
			}
			return rets
		})
	}
}
