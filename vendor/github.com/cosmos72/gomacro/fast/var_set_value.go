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
 * var_set_value.go
 *
 *  Created on Apr 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
	"unsafe"

	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// varSetValue compiles 'name = value' where value is a reflect.Value passed at runtime.
// Used to assign variables with the result of multi-valued expressions,
// and to implement multiple assignment var1, var2... = expr1, expr2...
func (c *Comp) varSetValue(va *Var) func(*Env, r.Value) {
	t := va.Type
	rt := t.ReflectType()
	upn := va.Upn
	desc := va.Desc
	var ret func(env *Env, v r.Value)

	switch desc.Class() {
	default:
		c.Errorf("cannot assign to %v %s", desc.Class(), va.Name)
		return nil
	case VarBind:
		// if current package is at least partially compiled, also variables
		// with kind = Bool, Int*, Uint*, Float*, Complex* may have class == VarBind

		index := desc.Index()
		if index == NoIndex {
			// assigning a value to _ has no effect at all
			return nil
		}
		zero := xr.Zero(t)
		switch upn {
		case 0:
			switch t.Kind() {
			case r.Bool:
				ret = func(env *Env, v r.Value) {
					env.Vals[index].SetBool(v.Bool())
				}
			case r.Int, r.Int8, r.Int32, r.Int64:
				ret = func(env *Env, v r.Value) {
					env.Vals[index].SetInt(v.Int())
				}
			case r.Uint, r.Uint8, r.Uint32, r.Uint64, r.Uintptr:
				ret = func(env *Env, v r.Value) {
					env.Vals[index].SetUint(v.Uint())
				}
			case r.Float32, r.Float64:
				ret = func(env *Env, v r.Value) {
					env.Vals[index].SetFloat(v.Float())
				}
			case r.Complex64, r.Complex128:
				ret = func(env *Env, v r.Value) {
					env.Vals[index].SetComplex(v.Complex())
				}
			case r.String:
				ret = func(env *Env, v r.Value) {
					if v.Kind() != r.String {
						v = convert(v, TypeOfString)
					}
					env.Vals[index].SetString(v.String())
				}
			case r.Chan, r.Interface, r.Map, r.Ptr, r.Slice:
				ret = func(env *Env, v r.Value) {
					if v == Nil || v == None {
						v = zero
					} else if v.Type() != rt {
						v = convert(v, rt)
					}
					env.Vals[index].Set(v)
				}
			default:
				ret = func(env *Env, v r.Value) {
					if v.Type() != rt {
						v = convert(v, rt)
					}
					env.Vals[index].Set(v)
				}
			}
		case 1:
			switch t.Kind() {
			case r.Bool:
				ret = func(env *Env, v r.Value) {
					env.Outer.Vals[index].SetBool(v.Bool())
				}
			case r.Int, r.Int8, r.Int32, r.Int64:
				ret = func(env *Env, v r.Value) {
					env.Outer.Vals[index].SetInt(v.Int())
				}
			case r.Uint, r.Uint8, r.Uint32, r.Uint64, r.Uintptr:
				ret = func(env *Env, v r.Value) {
					env.Outer.Vals[index].SetUint(v.Uint())
				}
			case r.Float32, r.Float64:
				ret = func(env *Env, v r.Value) {
					env.Outer.Vals[index].SetFloat(v.Float())
				}
			case r.Complex64, r.Complex128:
				ret = func(env *Env, v r.Value) {
					env.Outer.Vals[index].SetComplex(v.Complex())
				}
			case r.String:
				ret = func(env *Env, v r.Value) {
					if v.Kind() != r.String {
						v = convert(v, TypeOfString)
					}
					env.Outer.Vals[index].SetString(v.String())
				}
			case r.Chan, r.Interface, r.Map, r.Ptr, r.Slice:
				ret = func(env *Env, v r.Value) {
					if v == Nil || v == None {
						v = zero
					} else if v.Type() != rt {
						v = convert(v, rt)
					}
					env.Outer.Vals[index].Set(v)
				}
			default:
				ret = func(env *Env, v r.Value) {
					if v.Type() != rt {
						v = convert(v, rt)
					}
					env.Outer.Vals[index].Set(v)
				}
			}
		case 2:
			switch t.Kind() {
			case r.Bool:
				ret = func(env *Env, v r.Value) {
					env.Outer.Outer.Vals[index].SetBool(v.Bool())
				}
			case r.Int, r.Int8, r.Int32, r.Int64:
				ret = func(env *Env, v r.Value) {
					env.Outer.Outer.Vals[index].SetInt(v.Int())
				}
			case r.Uint, r.Uint8, r.Uint32, r.Uint64, r.Uintptr:
				ret = func(env *Env, v r.Value) {
					env.Outer.Outer.Vals[index].SetUint(v.Uint())
				}
			case r.Float32, r.Float64:
				ret = func(env *Env, v r.Value) {
					env.Outer.Outer.Vals[index].SetFloat(v.Float())
				}
			case r.Complex64, r.Complex128:
				ret = func(env *Env, v r.Value) {
					env.Outer.Outer.Vals[index].SetComplex(v.Complex())
				}
			case r.String:
				ret = func(env *Env, v r.Value) {
					if v.Kind() != r.String {
						v = convert(v, TypeOfString)
					}
					env.Outer.Outer.Vals[index].SetString(v.String())
				}
			case r.Chan, r.Interface, r.Map, r.Ptr, r.Slice:
				ret = func(env *Env, v r.Value) {
					if v == Nil || v == None {
						v = zero
					} else if v.Type() != rt {
						v = convert(v, rt)
					}
					env.Outer.Outer.Vals[index].Set(v)
				}
			default:
				ret = func(env *Env, v r.Value) {
					if v.Type() != rt {
						v = convert(v, rt)
					}
					env.Outer.Outer.Vals[index].Set(v)
				}
			}
		case c.Depth - 1:
			switch t.Kind() {
			case r.Bool:
				ret = func(env *Env, v r.Value) {
					env.FileEnv.Vals[index].SetBool(v.Bool())
				}
			case r.Int, r.Int8, r.Int32, r.Int64:
				ret = func(env *Env, v r.Value) {
					env.FileEnv.Vals[index].SetInt(v.Int())
				}
			case r.Uint, r.Uint8, r.Uint32, r.Uint64, r.Uintptr:
				ret = func(env *Env, v r.Value) {
					env.FileEnv.Vals[index].SetUint(v.Uint())
				}
			case r.Float32, r.Float64:
				ret = func(env *Env, v r.Value) {
					env.FileEnv.Vals[index].SetFloat(v.Float())
				}
			case r.Complex64, r.Complex128:
				ret = func(env *Env, v r.Value) {
					env.FileEnv.Vals[index].SetComplex(v.Complex())
				}
			case r.String:
				ret = func(env *Env, v r.Value) {
					if v.Kind() != r.String {
						v = convert(v, TypeOfString)
					}
					env.FileEnv.Vals[index].SetString(v.String())
				}
			case r.Chan, r.Interface, r.Map, r.Ptr, r.Slice:
				ret = func(env *Env, v r.Value) {
					if v == Nil || v == None {
						v = zero
					} else if v.Type() != rt {
						v = convert(v, rt)
					}
					env.FileEnv.Vals[index].Set(v)
				}
			default:
				ret = func(env *Env, v r.Value) {
					if v.Type() != rt {
						v = convert(v, rt)
					}
					env.FileEnv.Vals[index].Set(v)
				}
			}
		default:
			switch t.Kind() {
			case r.Bool:
				ret = func(env *Env, v r.Value) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					o.Vals[index].SetBool(v.Bool())
				}
			case r.Int, r.Int8, r.Int32, r.Int64:
				ret = func(env *Env, v r.Value) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					o.Vals[index].SetInt(v.Int())
				}
			case r.Uint, r.Uint8, r.Uint32, r.Uint64, r.Uintptr:
				ret = func(env *Env, v r.Value) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					o.Vals[index].SetUint(v.Uint())
				}
			case r.Float32, r.Float64:
				ret = func(env *Env, v r.Value) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					o.Vals[index].SetFloat(v.Float())
				}
			case r.Complex64, r.Complex128:
				ret = func(env *Env, v r.Value) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					o.Vals[index].SetComplex(v.Complex())
				}
			case r.String:
				ret = func(env *Env, v r.Value) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					if v.Kind() != r.String {
						v = convert(v, TypeOfString)
					}
					o.Vals[index].SetString(v.String())
				}
			case r.Chan, r.Interface, r.Map, r.Ptr, r.Slice:
				ret = func(env *Env, v r.Value) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					if v == Nil || v == None {
						v = zero
					} else if v.Type() != rt {
						v = convert(v, rt)
					}
					o.Vals[index].Set(v)
				}
			default:
				ret = func(env *Env, v r.Value) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					if v.Type() != rt {
						v = convert(v, rt)
					}
					o.Vals[index].Set(v)
				}
			}
		}
	case IntBind:
		index := desc.Index()
		if index == NoIndex {
			// assigning a value to _ has no effect at all
			return nil
		}
		switch upn {
		case 0:
			switch t.Kind() {
			case r.Bool:
				ret = func(env *Env, v r.Value) {
					*(*bool)(unsafe.Pointer(&env.Ints[index])) = v.Bool()
				}
			case r.Int:
				ret = func(env *Env, v r.Value) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) = int(v.Int())
				}
			case r.Int8:
				ret = func(env *Env, v r.Value) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) = int8(v.Int())
				}
			case r.Int16:
				ret = func(env *Env, v r.Value) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) = int16(v.Int())
				}
			case r.Int32:
				ret = func(env *Env, v r.Value) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) = int32(v.Int())
				}
			case r.Int64:
				ret = func(env *Env, v r.Value) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) = v.Int()
				}
			case r.Uint:
				ret = func(env *Env, v r.Value) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) = uint(v.Uint())
				}
			case r.Uint8:
				ret = func(env *Env, v r.Value) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) = uint8(v.Uint())
				}
			case r.Uint16:
				ret = func(env *Env, v r.Value) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) = uint16(v.Uint())
				}
			case r.Uint32:
				ret = func(env *Env, v r.Value) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) = uint32(v.Uint())
				}
			case r.Uint64:
				ret = func(env *Env, v r.Value) {
					env.Ints[index] = v.Uint()
				}
			case r.Uintptr:
				ret = func(env *Env, v r.Value) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) = uintptr(v.Uint())
				}
			case r.Float32:
				ret = func(env *Env, v r.Value) {
					*(*float32)(unsafe.Pointer(&env.Ints[index])) = float32(v.Float())
				}
			case r.Float64:
				ret = func(env *Env, v r.Value) {
					*(*float64)(unsafe.Pointer(&env.Ints[index])) = v.Float()
				}
			case r.Complex64:
				ret = func(env *Env, v r.Value) {
					*(*complex64)(unsafe.Pointer(&env.Ints[index])) = complex64(v.Complex())
				}
			case r.Complex128:
				ret = func(env *Env, v r.Value) {
					*(*complex128)(unsafe.Pointer(&env.Ints[index])) = v.Complex()
				}
			default:
				c.Errorf("unsupported type, cannot use for optimized assignment: %s <%v>", va.Name, t)
				return nil
			}
		case 1:
			switch t.Kind() {
			case r.Bool:
				ret = func(env *Env, v r.Value) {
					*(*bool)(unsafe.Pointer(&env.Outer.Ints[index])) = v.Bool()
				}
			case r.Int:
				ret = func(env *Env, v r.Value) {
					*(*int)(unsafe.Pointer(&env.Outer.Ints[index])) = int(v.Int())
				}
			case r.Int8:
				ret = func(env *Env, v r.Value) {
					*(*int8)(unsafe.Pointer(&env.Outer.Ints[index])) = int8(v.Int())
				}
			case r.Int16:
				ret = func(env *Env, v r.Value) {
					*(*int16)(unsafe.Pointer(&env.Outer.Ints[index])) = int16(v.Int())
				}
			case r.Int32:
				ret = func(env *Env, v r.Value) {
					*(*int32)(unsafe.Pointer(&env.Outer.Ints[index])) = int32(v.Int())
				}
			case r.Int64:
				ret = func(env *Env, v r.Value) {
					*(*int64)(unsafe.Pointer(&env.Outer.Ints[index])) = v.Int()
				}
			case r.Uint:
				ret = func(env *Env, v r.Value) {
					*(*uint)(unsafe.Pointer(&env.Outer.Ints[index])) = uint(v.Uint())
				}
			case r.Uint8:
				ret = func(env *Env, v r.Value) {
					*(*uint8)(unsafe.Pointer(&env.Outer.Ints[index])) = uint8(v.Uint())
				}
			case r.Uint16:
				ret = func(env *Env, v r.Value) {
					*(*uint16)(unsafe.Pointer(&env.Outer.Ints[index])) = uint16(v.Uint())
				}
			case r.Uint32:
				ret = func(env *Env, v r.Value) {
					*(*uint32)(unsafe.Pointer(&env.Outer.Ints[index])) = uint32(v.Uint())
				}
			case r.Uint64:
				ret = func(env *Env, v r.Value) {
					env.Outer.Ints[index] = v.Uint()
				}
			case r.Uintptr:
				ret = func(env *Env, v r.Value) {
					*(*uintptr)(unsafe.Pointer(&env.Outer.Ints[index])) = uintptr(v.Uint())
				}
			case r.Float32:
				ret = func(env *Env, v r.Value) {
					*(*float32)(unsafe.Pointer(&env.Outer.Ints[index])) = float32(v.Float())
				}
			case r.Float64:
				ret = func(env *Env, v r.Value) {
					*(*float64)(unsafe.Pointer(&env.Outer.Ints[index])) = v.Float()
				}
			case r.Complex64:
				ret = func(env *Env, v r.Value) {
					*(*complex64)(unsafe.Pointer(&env.Outer.Ints[index])) = complex64(v.Complex())
				}
			case r.Complex128:
				ret = func(env *Env, v r.Value) {
					*(*complex128)(unsafe.Pointer(&env.Outer.Ints[index])) = v.Complex()
				}
			default:
				c.Errorf("unsupported type, cannot use for optimized assignment: %s <%v>", va.Name, t)
				return nil
			}
		case 2:
			switch t.Kind() {
			case r.Bool:
				ret = func(env *Env, v r.Value) {
					*(*bool)(unsafe.Pointer(&env.Outer.Outer.Ints[index])) = v.Bool()
				}
			case r.Int:
				ret = func(env *Env, v r.Value) {
					*(*int)(unsafe.Pointer(&env.Outer.Outer.Ints[index])) = int(v.Int())
				}
			case r.Int8:
				ret = func(env *Env, v r.Value) {
					*(*int8)(unsafe.Pointer(&env.Outer.Outer.Ints[index])) = int8(v.Int())
				}
			case r.Int16:
				ret = func(env *Env, v r.Value) {
					*(*int16)(unsafe.Pointer(&env.Outer.Outer.Ints[index])) = int16(v.Int())
				}
			case r.Int32:
				ret = func(env *Env, v r.Value) {
					*(*int32)(unsafe.Pointer(&env.Outer.Outer.Ints[index])) = int32(v.Int())
				}
			case r.Int64:
				ret = func(env *Env, v r.Value) {
					*(*int64)(unsafe.Pointer(&env.Outer.Outer.Ints[index])) = v.Int()
				}
			case r.Uint:
				ret = func(env *Env, v r.Value) {
					*(*uint)(unsafe.Pointer(&env.Outer.Outer.Ints[index])) = uint(v.Uint())
				}
			case r.Uint8:
				ret = func(env *Env, v r.Value) {
					*(*uint8)(unsafe.Pointer(&env.Outer.Outer.Ints[index])) = uint8(v.Uint())
				}
			case r.Uint16:
				ret = func(env *Env, v r.Value) {
					*(*uint16)(unsafe.Pointer(&env.Outer.Outer.Ints[index])) = uint16(v.Uint())
				}
			case r.Uint32:
				ret = func(env *Env, v r.Value) {
					*(*uint32)(unsafe.Pointer(&env.Outer.Outer.Ints[index])) = uint32(v.Uint())
				}
			case r.Uint64:
				ret = func(env *Env, v r.Value) {
					env.Outer.Outer.Ints[index] = v.Uint()
				}
			case r.Uintptr:
				ret = func(env *Env, v r.Value) {
					*(*uintptr)(unsafe.Pointer(&env.Outer.Outer.Ints[index])) = uintptr(v.Uint())
				}
			case r.Float32:
				ret = func(env *Env, v r.Value) {
					*(*float32)(unsafe.Pointer(&env.Outer.Outer.Ints[index])) = float32(v.Float())
				}
			case r.Float64:
				ret = func(env *Env, v r.Value) {
					*(*float64)(unsafe.Pointer(&env.Outer.Outer.Ints[index])) = v.Float()
				}
			case r.Complex64:
				ret = func(env *Env, v r.Value) {
					*(*complex64)(unsafe.Pointer(&env.Outer.Outer.Ints[index])) = complex64(v.Complex())
				}
			case r.Complex128:
				ret = func(env *Env, v r.Value) {
					*(*complex128)(unsafe.Pointer(&env.Outer.Outer.Ints[index])) = v.Complex()
				}
			default:
				c.Errorf("unsupported type, cannot use for optimized assignment: %s <%v>", va.Name, t)
				return nil
			}
		case c.Depth - 1:
			switch t.Kind() {
			case r.Bool:
				ret = func(env *Env, v r.Value) {
					*(*bool)(unsafe.Pointer(&env.FileEnv.Ints[index])) = v.Bool()
				}
			case r.Int:
				ret = func(env *Env, v r.Value) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) = int(v.Int())
				}
			case r.Int8:
				ret = func(env *Env, v r.Value) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) = int8(v.Int())
				}
			case r.Int16:
				ret = func(env *Env, v r.Value) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) = int16(v.Int())
				}
			case r.Int32:
				ret = func(env *Env, v r.Value) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) = int32(v.Int())
				}
			case r.Int64:
				ret = func(env *Env, v r.Value) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) = v.Int()
				}
			case r.Uint:
				ret = func(env *Env, v r.Value) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) = uint(v.Uint())
				}
			case r.Uint8:
				ret = func(env *Env, v r.Value) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) = uint8(v.Uint())
				}
			case r.Uint16:
				ret = func(env *Env, v r.Value) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) = uint16(v.Uint())
				}
			case r.Uint32:
				ret = func(env *Env, v r.Value) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) = uint32(v.Uint())
				}
			case r.Uint64:
				ret = func(env *Env, v r.Value) {
					env.FileEnv.Ints[index] = v.Uint()
				}
			case r.Uintptr:
				ret = func(env *Env, v r.Value) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) = uintptr(v.Uint())
				}
			case r.Float32:
				ret = func(env *Env, v r.Value) {
					*(*float32)(unsafe.Pointer(&env.FileEnv.Ints[index])) = float32(v.Float())
				}
			case r.Float64:
				ret = func(env *Env, v r.Value) {
					*(*float64)(unsafe.Pointer(&env.FileEnv.Ints[index])) = v.Float()
				}
			case r.Complex64:
				ret = func(env *Env, v r.Value) {
					*(*complex64)(unsafe.Pointer(&env.FileEnv.Ints[index])) = complex64(v.Complex())
				}
			case r.Complex128:
				ret = func(env *Env, v r.Value) {
					*(*complex128)(unsafe.Pointer(&env.FileEnv.Ints[index])) = v.Complex()
				}
			default:
				c.Errorf("unsupported type, cannot use for optimized assignment: %s <%v>", va.Name, t)
				return nil
			}
		default:
			switch t.Kind() {
			case r.Bool:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*bool)(unsafe.Pointer(&env.Outer.Outer.Outer.Ints[index])) = v.Bool()
				}
			case r.Int:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*int)(unsafe.Pointer(&env.Outer.Outer.Outer.Ints[index])) = int(v.Int())
				}
			case r.Int8:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*int8)(unsafe.Pointer(&env.Outer.Outer.Outer.Ints[index])) = int8(v.Int())
				}
			case r.Int16:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*int16)(unsafe.Pointer(&env.Outer.Outer.Outer.Ints[index])) = int16(v.Int())
				}
			case r.Int32:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*int32)(unsafe.Pointer(&env.Outer.Outer.Outer.Ints[index])) = int32(v.Int())
				}
			case r.Int64:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*int64)(unsafe.Pointer(&env.Outer.Outer.Outer.Ints[index])) = v.Int()
				}
			case r.Uint:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*uint)(unsafe.Pointer(&env.Outer.Outer.Outer.Ints[index])) = uint(v.Uint())
				}
			case r.Uint8:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*uint8)(unsafe.Pointer(&env.Outer.Outer.Outer.Ints[index])) = uint8(v.Uint())
				}
			case r.Uint16:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*uint16)(unsafe.Pointer(&env.Outer.Outer.Outer.Ints[index])) = uint16(v.Uint())
				}
			case r.Uint32:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*uint32)(unsafe.Pointer(&env.Outer.Outer.Outer.Ints[index])) = uint32(v.Uint())
				}
			case r.Uint64:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					env.Outer.Outer.Outer.Ints[index] = v.Uint()
				}
			case r.Uintptr:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*uintptr)(unsafe.Pointer(&env.Outer.Outer.Outer.Ints[index])) = uintptr(v.Uint())
				}
			case r.Float32:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*float32)(unsafe.Pointer(&env.Outer.Outer.Outer.Ints[index])) = float32(v.Float())
				}
			case r.Float64:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*float64)(unsafe.Pointer(&env.Outer.Outer.Outer.Ints[index])) = v.Float()
				}
			case r.Complex64:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*complex64)(unsafe.Pointer(&env.Outer.Outer.Outer.Ints[index])) = complex64(v.Complex())
				}
			case r.Complex128:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*complex128)(unsafe.Pointer(&env.Outer.Outer.Outer.Ints[index])) = v.Complex()
				}
			default:
				c.Errorf("unsupported type, cannot use for optimized assignment: %s <%v>", va.Name, t)
				return nil
			}
		}
	}
	return ret
}
