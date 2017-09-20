// +build gomacro_fast_compact

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
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *
 * call1ret1compact.go
 *
 *  Created on Apr 15, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
)

func (c *Comp) call1ret1(call *Call, maxdepth int) I {
	expr := call.Fun
	exprfun := expr.AsX1()

	arg := call.Args[0]
	argfun := arg.AsX1()
	var ret I

	switch expr.Type.Out(0).Kind() {
	case r.Bool:
		ret = func(env *Env) bool {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return retv.Bool()
		}

	case r.Int:
		ret = func(env *Env) int {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return int(retv.Int())
		}

	case r.Int8:
		ret = func(env *Env) int8 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return int8(retv.Int())
		}

	case r.Int16:
		ret = func(env *Env) int16 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return int16(retv.Int())
		}

	case r.Int32:
		ret = func(env *Env) int32 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return int32(retv.Int())
		}

	case r.Int64:
		ret = func(env *Env) int64 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return retv.Int()
		}

	case r.Uint:
		ret = func(env *Env) uint {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return uint(retv.Uint())
		}

	case r.Uint8:
		ret = func(env *Env) uint8 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return uint8(retv.Uint())
		}

	case r.Uint16:
		ret = func(env *Env) uint16 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return uint16(retv.Uint())
		}

	case r.Uint32:
		ret = func(env *Env) uint32 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return uint32(retv.Uint())
		}

	case r.Uint64:
		ret = func(env *Env) uint64 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return retv.Uint()
		}

	case r.Uintptr:
		ret = func(env *Env) uintptr {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return uintptr(retv.Uint())
		}

	case r.Float32:
		ret = func(env *Env) float32 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return float32(retv.Float())
		}

	case r.Float64:
		ret = func(env *Env) float64 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return retv.Float()
		}

	case r.Complex64:
		ret = func(env *Env) complex64 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return complex64(retv.Complex())
		}

	case r.Complex128:
		ret = func(env *Env) complex128 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return retv.Complex()
		}

	case r.String:
		ret = func(env *Env) string {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)[0]
			return retv.String()
		}
	default:
		ret = func(env *Env) r.Value {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			return funv.Call(argv)[0]
		}
	}
	return ret
}
