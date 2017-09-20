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
 * call_multivalue.go
 *
 *  Created on May 29, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
)

// call_multivalue compiles foo(bar()) where bar() returns multiple values
func call_multivalue(call *Call, maxdepth int) I {
	// no need to special case variadic functions here
	expr := call.Fun
	exprfun := expr.AsX1()
	argfun := call.Args[0].AsXV(OptDefaults)
	nout := len(call.OutTypes)
	var ret I
	switch nout {
	case 0:
		if call.Ellipsis {
			ret = func(env *Env) {
				funv := exprfun(env)
				_, argv := argfun(env)
				funv.CallSlice(argv)
			}
		} else {
			ret = func(env *Env) {
				funv := exprfun(env)
				_, argv := argfun(env)
				funv.Call(argv)
			}
		}
	case 1:
		if call.Ellipsis {
			ret = call_multivalue_ellipsis_ret1(call, maxdepth)
		} else {
			ret = call_multivalue_ret1(call, maxdepth)
		}
	default:
		if call.Ellipsis {
			ret = func(env *Env) (r.Value, []r.Value) {
				funv := exprfun(env)
				_, argv := argfun(env)
				rets := funv.CallSlice(argv)
				return rets[0], rets
			}
		} else {
			ret = func(env *Env) (r.Value, []r.Value) {
				funv := exprfun(env)
				_, argv := argfun(env)
				rets := funv.Call(argv)
				return rets[0], rets
			}
		}
	}
	return ret
}

// mandatory optimization: fast_interpreter ASSUMES that expressions
// returning bool, int, uint, float, complex, string do NOT wrap them in reflect.Value
func call_multivalue_ret1(call *Call, maxdepth int) I {
	exprfun := call.Fun.AsX1()
	argfun := call.Args[0].AsXV(OptDefaults)
	kout := call.OutTypes[0].Kind()
	var ret I
	switch kout {
	case r.Bool:
		ret = func(env *Env) bool {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return retv.Bool()
		}
	case r.Int:
		ret = func(env *Env) int {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return int(retv.Int())
		}
	case r.Int8:
		ret = func(env *Env) int8 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return int8(retv.Int())
		}
	case r.Int16:
		ret = func(env *Env) int16 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return int16(retv.Int())
		}
	case r.Int32:
		ret = func(env *Env) int32 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return int32(retv.Int())
		}
	case r.Int64:
		ret = func(env *Env) int64 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return retv.Int()
		}
	case r.Uint:
		ret = func(env *Env) uint {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return uint(retv.Uint())
		}
	case r.Uint8:
		ret = func(env *Env) uint8 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return uint8(retv.Uint())
		}
	case r.Uint16:
		ret = func(env *Env) uint16 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return uint16(retv.Uint())
		}
	case r.Uint32:
		ret = func(env *Env) uint32 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return uint32(retv.Uint())
		}
	case r.Uint64:
		ret = func(env *Env) uint64 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return retv.Uint()
		}
	case r.Uintptr:
		ret = func(env *Env) uintptr {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return uintptr(retv.Uint())
		}
	case r.Float32:
		ret = func(env *Env) float32 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return float32(retv.Float())
		}
	case r.Float64:
		ret = func(env *Env) float64 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return retv.Float()
		}
	case r.Complex64:
		ret = func(env *Env) complex64 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return complex64(retv.Complex())
		}
	case r.Complex128:
		ret = func(env *Env) complex128 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return retv.Complex()
		}
	case r.String:
		ret = func(env *Env) string {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.Call(argv)[0]
			return retv.String()
		}
	default:
		ret = func(env *Env) r.Value {
			funv := exprfun(env)
			_, argv := argfun(env)
			return funv.Call(argv)[0]
		}
	}
	return ret
}

// mandatory optimization: fast_interpreter ASSUMES that expressions
// returning bool, int, uint, float, complex, string do NOT wrap them in reflect.Value
func call_multivalue_ellipsis_ret1(call *Call, maxdepth int) I {
	exprfun := call.Fun.AsX1()
	argfun := call.Args[0].AsXV(OptDefaults)
	kout := call.OutTypes[0].Kind()
	var ret I
	switch kout {
	case r.Bool:
		ret = func(env *Env) bool {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return retv.Bool()
		}
	case r.Int:
		ret = func(env *Env) int {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return int(retv.Int())
		}
	case r.Int8:
		ret = func(env *Env) int8 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return int8(retv.Int())
		}
	case r.Int16:
		ret = func(env *Env) int16 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return int16(retv.Int())
		}
	case r.Int32:
		ret = func(env *Env) int32 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return int32(retv.Int())
		}
	case r.Int64:
		ret = func(env *Env) int64 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return retv.Int()
		}
	case r.Uint:
		ret = func(env *Env) uint {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return uint(retv.Uint())
		}
	case r.Uint8:
		ret = func(env *Env) uint8 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return uint8(retv.Uint())
		}
	case r.Uint16:
		ret = func(env *Env) uint16 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return uint16(retv.Uint())
		}
	case r.Uint32:
		ret = func(env *Env) uint32 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return uint32(retv.Uint())
		}
	case r.Uint64:
		ret = func(env *Env) uint64 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return retv.Uint()
		}
	case r.Uintptr:
		ret = func(env *Env) uintptr {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return uintptr(retv.Uint())
		}
	case r.Float32:
		ret = func(env *Env) float32 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return float32(retv.Float())
		}
	case r.Float64:
		ret = func(env *Env) float64 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return retv.Float()
		}
	case r.Complex64:
		ret = func(env *Env) complex64 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return complex64(retv.Complex())
		}
	case r.Complex128:
		ret = func(env *Env) complex128 {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return retv.Complex()
		}
	case r.String:
		ret = func(env *Env) string {
			funv := exprfun(env)
			_, argv := argfun(env)
			retv := funv.CallSlice(argv)[0]
			return retv.String()
		}
	default:
		ret = func(env *Env) r.Value {
			funv := exprfun(env)
			_, argv := argfun(env)
			return funv.CallSlice(argv)[0]
		}
	}
	return ret
}
