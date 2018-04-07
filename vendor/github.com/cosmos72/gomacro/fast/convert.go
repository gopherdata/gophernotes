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
	e := c.Expr1(node)
	if e.Untyped() {
		e.ConstTo(t)
	}

	if e.Type != nil && e.Type.IdenticalTo(t) {
		return e
	} else if e.Type == nil && IsNillableKind(t.Kind()) {
		e.Type = t
		e.Value = xr.Zero(t).Interface()
	} else if e.Type != nil && e.Type.ConvertibleTo(t) {
	} else {
		c.Errorf("cannot convert %v to %v: %v", e.Type, t, node)
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
	return exprFun(t, ret)
}

// Converter returns a function that converts reflect.Value from tin to tout
// also supports conversion from interpreted types to interfaces
func (c *Comp) Converter(tin, tout xr.Type) func(val r.Value) r.Value {
	if !tin.ConvertibleTo(tout) {
		c.Errorf("cannot convert from <%v> to <%v>", tin, tout)
	} else if tin.IdenticalTo(tout) {
		return nil
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
	case rtout.Kind() == r.Interface:
		// conversion from interpreted type to compiled interface.
		// must use a proxy that pre-implement compiled interfaces.
		return c.converterToProxy(tin, tout)
	default:
		c.Errorf("unimplemented conversion from <%v> to <%v>", tin, tout)
		return nil
	}
}
