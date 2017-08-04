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
 * var_setter.go
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
		index := desc.Index()
		if index == NoIndex {
			// assigning a value to _ has no effect at all
			return nil
		}
		zero := xr.Zero(t)
		switch upn {
		case 0:
			switch t.Kind() {
			case r.Complex128:
				ret = func(env *Env, v r.Value) {
					env.Binds[index].SetComplex(v.Complex())
				}
			case r.String:
				ret = func(env *Env, v r.Value) {
					if v.Kind() != r.String {
						v = v.Convert(TypeOfString)
					}
					env.Binds[index].SetString(v.String())
				}
			case r.Chan, r.Interface, r.Map, r.Ptr, r.Slice:
				ret = func(env *Env, v r.Value) {
					if v == Nil || v == None {
						v = zero
					} else if v.Type() != rt {
						v = v.Convert(rt)
					}
					env.Binds[index].Set(v)
				}
			default:
				ret = func(env *Env, v r.Value) {
					if v.Type() != rt {
						v = v.Convert(rt)
					}
					env.Binds[index].Set(v)
				}
			}
		case 1:
			switch t.Kind() {
			case r.Complex128:
				ret = func(env *Env, v r.Value) {
					env.Outer.Binds[index].SetComplex(v.Complex())
				}
			case r.String:
				ret = func(env *Env, v r.Value) {
					if v.Kind() != r.String {
						v = v.Convert(TypeOfString)
					}
					env.Outer.Binds[index].SetString(v.String())
				}
			case r.Chan, r.Interface, r.Map, r.Ptr, r.Slice:
				ret = func(env *Env, v r.Value) {
					if v == Nil || v == None {
						v = zero
					} else if v.Type() != rt {
						v = v.Convert(rt)
					}
					env.Outer.Binds[index].Set(v)
				}
			default:
				ret = func(env *Env, v r.Value) {
					if v.Type() != rt {
						v = v.Convert(rt)
					}
					env.Outer.Binds[index].Set(v)
				}
			}
		case 2:
			switch t.Kind() {
			case r.Complex128:
				ret = func(env *Env, v r.Value) {
					env.Outer.Outer.Binds[index].SetComplex(v.Complex())
				}
			case r.String:
				ret = func(env *Env, v r.Value) {
					if v.Kind() != r.String {
						v = v.Convert(TypeOfString)
					}
					env.Outer.Outer.Binds[index].SetString(v.String())
				}
			case r.Chan, r.Interface, r.Map, r.Ptr, r.Slice:
				ret = func(env *Env, v r.Value) {
					if v == Nil || v == None {
						v = zero
					} else if v.Type() != rt {
						v = v.Convert(rt)
					}
					env.Outer.Outer.Binds[index].Set(v)
				}
			default:
				ret = func(env *Env, v r.Value) {
					if v.Type() != rt {
						v = v.Convert(rt)
					}
					env.Outer.Outer.Binds[index].Set(v)
				}
			}
		default:
			switch t.Kind() {
			case r.Complex128:
				ret = func(env *Env, v r.Value) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					o.Binds[index].SetComplex(v.Complex())
				}
			case r.String:
				ret = func(env *Env, v r.Value) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					if v.Kind() != r.String {
						v = v.Convert(TypeOfString)
					}
					o.Binds[index].SetString(v.String())
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
						v = v.Convert(rt)
					}
					o.Binds[index].Set(v)
				}
			default:
				ret = func(env *Env, v r.Value) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					if v.Type() != rt {
						v = v.Convert(rt)
					}
					o.Binds[index].Set(v)
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
					*(*bool)(unsafe.Pointer(&env.IntBinds[index])) = v.Bool()
				}
			case r.Int:
				ret = func(env *Env, v r.Value) {
					*(*int)(unsafe.Pointer(&env.IntBinds[index])) = int(v.Int())
				}
			case r.Int8:
				ret = func(env *Env, v r.Value) {
					*(*int8)(unsafe.Pointer(&env.IntBinds[index])) = int8(v.Int())
				}
			case r.Int16:
				ret = func(env *Env, v r.Value) {
					*(*int16)(unsafe.Pointer(&env.IntBinds[index])) = int16(v.Int())
				}
			case r.Int32:
				ret = func(env *Env, v r.Value) {
					*(*int32)(unsafe.Pointer(&env.IntBinds[index])) = int32(v.Int())
				}
			case r.Int64:
				ret = func(env *Env, v r.Value) {
					*(*int64)(unsafe.Pointer(&env.IntBinds[index])) = v.Int()
				}
			case r.Uint:
				ret = func(env *Env, v r.Value) {
					*(*uint)(unsafe.Pointer(&env.IntBinds[index])) = uint(v.Uint())
				}
			case r.Uint8:
				ret = func(env *Env, v r.Value) {
					*(*uint8)(unsafe.Pointer(&env.IntBinds[index])) = uint8(v.Uint())
				}
			case r.Uint16:
				ret = func(env *Env, v r.Value) {
					*(*uint16)(unsafe.Pointer(&env.IntBinds[index])) = uint16(v.Uint())
				}
			case r.Uint32:
				ret = func(env *Env, v r.Value) {
					*(*uint32)(unsafe.Pointer(&env.IntBinds[index])) = uint32(v.Uint())
				}
			case r.Uint64:
				ret = func(env *Env, v r.Value) {
					env.IntBinds[index] = v.Uint()
				}
			case r.Uintptr:
				ret = func(env *Env, v r.Value) {
					*(*uintptr)(unsafe.Pointer(&env.IntBinds[index])) = uintptr(v.Uint())
				}
			case r.Float32:
				ret = func(env *Env, v r.Value) {
					*(*float32)(unsafe.Pointer(&env.IntBinds[index])) = float32(v.Float())
				}
			case r.Float64:
				ret = func(env *Env, v r.Value) {
					*(*float64)(unsafe.Pointer(&env.IntBinds[index])) = v.Float()
				}
			case r.Complex64:
				ret = func(env *Env, v r.Value) {
					*(*complex64)(unsafe.Pointer(&env.IntBinds[index])) = complex64(v.Complex())
				}
			default:
				c.Errorf("unsupported type, cannot use for optimized assignment: %s <%v>", va.Name, t)
				return nil
			}
		case 1:
			switch t.Kind() {
			case r.Bool:
				ret = func(env *Env, v r.Value) {
					*(*bool)(unsafe.Pointer(&env.Outer.IntBinds[index])) = v.Bool()
				}
			case r.Int:
				ret = func(env *Env, v r.Value) {
					*(*int)(unsafe.Pointer(&env.Outer.IntBinds[index])) = int(v.Int())
				}
			case r.Int8:
				ret = func(env *Env, v r.Value) {
					*(*int8)(unsafe.Pointer(&env.Outer.IntBinds[index])) = int8(v.Int())
				}
			case r.Int16:
				ret = func(env *Env, v r.Value) {
					*(*int16)(unsafe.Pointer(&env.Outer.IntBinds[index])) = int16(v.Int())
				}
			case r.Int32:
				ret = func(env *Env, v r.Value) {
					*(*int32)(unsafe.Pointer(&env.Outer.IntBinds[index])) = int32(v.Int())
				}
			case r.Int64:
				ret = func(env *Env, v r.Value) {
					*(*int64)(unsafe.Pointer(&env.Outer.IntBinds[index])) = v.Int()
				}
			case r.Uint:
				ret = func(env *Env, v r.Value) {
					*(*uint)(unsafe.Pointer(&env.Outer.IntBinds[index])) = uint(v.Uint())
				}
			case r.Uint8:
				ret = func(env *Env, v r.Value) {
					*(*uint8)(unsafe.Pointer(&env.Outer.IntBinds[index])) = uint8(v.Uint())
				}
			case r.Uint16:
				ret = func(env *Env, v r.Value) {
					*(*uint16)(unsafe.Pointer(&env.Outer.IntBinds[index])) = uint16(v.Uint())
				}
			case r.Uint32:
				ret = func(env *Env, v r.Value) {
					*(*uint32)(unsafe.Pointer(&env.Outer.IntBinds[index])) = uint32(v.Uint())
				}
			case r.Uint64:
				ret = func(env *Env, v r.Value) {
					env.Outer.IntBinds[index] = v.Uint()
				}
			case r.Uintptr:
				ret = func(env *Env, v r.Value) {
					*(*uintptr)(unsafe.Pointer(&env.Outer.IntBinds[index])) = uintptr(v.Uint())
				}
			case r.Float32:
				ret = func(env *Env, v r.Value) {
					*(*float32)(unsafe.Pointer(&env.Outer.IntBinds[index])) = float32(v.Float())
				}
			case r.Float64:
				ret = func(env *Env, v r.Value) {
					*(*float64)(unsafe.Pointer(&env.Outer.IntBinds[index])) = v.Float()
				}
			case r.Complex64:
				ret = func(env *Env, v r.Value) {
					*(*complex64)(unsafe.Pointer(&env.Outer.IntBinds[index])) = complex64(v.Complex())
				}
			default:
				c.Errorf("unsupported type, cannot use for optimized assignment: %s <%v>", va.Name, t)
				return nil
			}
		case 2:
			switch t.Kind() {
			case r.Bool:
				ret = func(env *Env, v r.Value) {
					*(*bool)(unsafe.Pointer(&env.Outer.Outer.IntBinds[index])) = v.Bool()
				}
			case r.Int:
				ret = func(env *Env, v r.Value) {
					*(*int)(unsafe.Pointer(&env.Outer.Outer.IntBinds[index])) = int(v.Int())
				}
			case r.Int8:
				ret = func(env *Env, v r.Value) {
					*(*int8)(unsafe.Pointer(&env.Outer.Outer.IntBinds[index])) = int8(v.Int())
				}
			case r.Int16:
				ret = func(env *Env, v r.Value) {
					*(*int16)(unsafe.Pointer(&env.Outer.Outer.IntBinds[index])) = int16(v.Int())
				}
			case r.Int32:
				ret = func(env *Env, v r.Value) {
					*(*int32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[index])) = int32(v.Int())
				}
			case r.Int64:
				ret = func(env *Env, v r.Value) {
					*(*int64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[index])) = v.Int()
				}
			case r.Uint:
				ret = func(env *Env, v r.Value) {
					*(*uint)(unsafe.Pointer(&env.Outer.Outer.IntBinds[index])) = uint(v.Uint())
				}
			case r.Uint8:
				ret = func(env *Env, v r.Value) {
					*(*uint8)(unsafe.Pointer(&env.Outer.Outer.IntBinds[index])) = uint8(v.Uint())
				}
			case r.Uint16:
				ret = func(env *Env, v r.Value) {
					*(*uint16)(unsafe.Pointer(&env.Outer.Outer.IntBinds[index])) = uint16(v.Uint())
				}
			case r.Uint32:
				ret = func(env *Env, v r.Value) {
					*(*uint32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[index])) = uint32(v.Uint())
				}
			case r.Uint64:
				ret = func(env *Env, v r.Value) {
					env.Outer.Outer.IntBinds[index] = v.Uint()
				}
			case r.Uintptr:
				ret = func(env *Env, v r.Value) {
					*(*uintptr)(unsafe.Pointer(&env.Outer.Outer.IntBinds[index])) = uintptr(v.Uint())
				}
			case r.Float32:
				ret = func(env *Env, v r.Value) {
					*(*float32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[index])) = float32(v.Float())
				}
			case r.Float64:
				ret = func(env *Env, v r.Value) {
					*(*float64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[index])) = v.Float()
				}
			case r.Complex64:
				ret = func(env *Env, v r.Value) {
					*(*complex64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[index])) = complex64(v.Complex())
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
					*(*bool)(unsafe.Pointer(&env.Outer.Outer.Outer.IntBinds[index])) = v.Bool()
				}
			case r.Int:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*int)(unsafe.Pointer(&env.Outer.Outer.Outer.IntBinds[index])) = int(v.Int())
				}
			case r.Int8:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*int8)(unsafe.Pointer(&env.Outer.Outer.Outer.IntBinds[index])) = int8(v.Int())
				}
			case r.Int16:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*int16)(unsafe.Pointer(&env.Outer.Outer.Outer.IntBinds[index])) = int16(v.Int())
				}
			case r.Int32:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*int32)(unsafe.Pointer(&env.Outer.Outer.Outer.IntBinds[index])) = int32(v.Int())
				}
			case r.Int64:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*int64)(unsafe.Pointer(&env.Outer.Outer.Outer.IntBinds[index])) = v.Int()
				}
			case r.Uint:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*uint)(unsafe.Pointer(&env.Outer.Outer.Outer.IntBinds[index])) = uint(v.Uint())
				}
			case r.Uint8:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*uint8)(unsafe.Pointer(&env.Outer.Outer.Outer.IntBinds[index])) = uint8(v.Uint())
				}
			case r.Uint16:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*uint16)(unsafe.Pointer(&env.Outer.Outer.Outer.IntBinds[index])) = uint16(v.Uint())
				}
			case r.Uint32:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*uint32)(unsafe.Pointer(&env.Outer.Outer.Outer.IntBinds[index])) = uint32(v.Uint())
				}
			case r.Uint64:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					env.Outer.Outer.Outer.IntBinds[index] = v.Uint()
				}
			case r.Uintptr:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*uintptr)(unsafe.Pointer(&env.Outer.Outer.Outer.IntBinds[index])) = uintptr(v.Uint())
				}
			case r.Float32:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*float32)(unsafe.Pointer(&env.Outer.Outer.Outer.IntBinds[index])) = float32(v.Float())
				}
			case r.Float64:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*float64)(unsafe.Pointer(&env.Outer.Outer.Outer.IntBinds[index])) = v.Float()
				}
			case r.Complex64:
				ret = func(env *Env, v r.Value) {
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					*(*complex64)(unsafe.Pointer(&env.Outer.Outer.Outer.IntBinds[index])) = complex64(v.Complex())
				}
			default:
				c.Errorf("unsupported type, cannot use for optimized assignment: %s <%v>", va.Name, t)
				return nil
			}
		}
	}
	return ret
}
