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
 * call2ret1.go
 *
 *  Created on Apr 15, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) call2ret1(call *Call, maxdepth int) I {

	expr := call.Fun
	if expr.Sym != nil && expr.Sym.Desc.Index() == NoIndex {
		Errorf("internal error: call2ret1() invoked for constant function %#v. use call_builtin() instead", expr)
	}
	t := expr.Type
	rtout := t.Out(0).ReflectType()
	kout := rtout.Kind()

	exprfun := expr.AsX1()
	argfunsX1 := call.MakeArgfunsX1()
	argfuns := [2]func(*Env) r.Value{
		argfunsX1[0],
		argfunsX1[1],
	}

	var ret I
	switch kout {
	case r.Bool:
		ret = func(env *Env) bool {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return ret0.Bool()
		}
	case r.Int:
		ret = func(env *Env) int {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return int(ret0.Int())
		}
	case r.Int8:
		ret = func(env *Env) int8 {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return int8(ret0.Int())
		}
	case r.Int16:
		ret = func(env *Env) int16 {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return int16(ret0.Int())
		}
	case r.Int32:
		ret = func(env *Env) int32 {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return int32(ret0.Int())
		}
	case r.Int64:
		ret = func(env *Env) int64 {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return ret0.Int()
		}
	case r.Uint:
		ret = func(env *Env) uint {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return uint(ret0.Uint())
		}
	case r.Uint8:
		ret = func(env *Env) uint8 {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return uint8(ret0.Uint())
		}
	case r.Uint16:
		ret = func(env *Env) uint16 {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return uint16(ret0.Uint())
		}
	case r.Uint32:
		ret = func(env *Env) uint32 {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return uint32(ret0.Uint())
		}
	case r.Uint64:
		ret = func(env *Env) uint64 {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return ret0.Uint()
		}
	case r.Uintptr:
		ret = func(env *Env) uintptr {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return uintptr(ret0.Uint())
		}
	case r.Float32:
		ret = func(env *Env) float32 {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return float32(ret0.Float())
		}
	case r.Float64:
		ret = func(env *Env) float64 {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return ret0.Float()
		}
	case r.Complex64:
		ret = func(env *Env) complex64 {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return complex64(ret0.Complex())
		}
	case r.Complex128:
		ret = func(env *Env) complex128 {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return ret0.Complex()
		}
	case r.String:
		ret = func(env *Env) string {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			ret0 := callxr(funv, argv)[0]
			return ret0.String()
		}
	default:
		ret = func(env *Env) r.Value {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			return callxr(funv, argv)[0]
		}
	}
	return ret
}
