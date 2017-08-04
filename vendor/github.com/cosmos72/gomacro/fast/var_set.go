// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

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
 * var_set.go
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

func (c *Comp) varSetZero(upn int, index int, t xr.Type) {
	zero := xr.Zero(t).Interface()
	c.varSetConst(upn, index, t, zero)
}
func (c *Comp) varSetConst(upn int, index int, t xr.Type, val I) {
	v := r.ValueOf(val)
	rt := t.ReflectType()
	if ValueType(v) == nil {
		v = r.Zero(rt)
	} else {
		v = v.Convert(rt)
	}

	var ret func(env *Env) (Stmt, *Env)
	switch upn {
	case 0:
		switch t.Kind() {
		case r.Bool:

			{
				val := v.Bool()

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int:

			{
				val := int(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:
			{
				val := int8(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:
			{
				val := int16(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			{
				val := int32(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			{
				val := v.Int()

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			{
				val :=

					uint(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			{
				val :=

					uint8(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint16:
			{
				val :=

					uint16(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint32:
			{
				val :=

					uint32(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint64:
			{
				val := v.Uint()

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uintptr:
			{
				val :=

					uintptr(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float32:
			{
				val :=

					float32(v.Float())

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float64:
			{
				val := v.Float()

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex64:
			{
				val :=

					complex64(v.Complex())

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex128:
			{
				val := v.Complex()

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Binds[index].SetComplex(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.String:
			{
				val := v.String()

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Binds[index].SetString(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case 1:
		switch t.Kind() {
		case r.Bool:

			{
				val := v.Bool()

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int:

			{
				val := int(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:
			{
				val := int8(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:
			{
				val := int16(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			{
				val := int32(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			{
				val := v.Int()

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			{
				val :=

					uint(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			{
				val :=

					uint8(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint16:
			{
				val :=

					uint16(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint32:
			{
				val :=

					uint32(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint64:
			{
				val := v.Uint()

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uintptr:
			{
				val :=

					uintptr(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float32:
			{
				val :=

					float32(v.Float())

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float64:
			{
				val := v.Float()

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex64:
			{
				val :=

					complex64(v.Complex())

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex128:
			{
				val := v.Complex()

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Binds[index].SetComplex(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.String:
			{
				val := v.String()

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Binds[index].SetString(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case 2:
		switch t.Kind() {
		case r.Bool:

			{
				val := v.Bool()

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int:

			{
				val := int(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:
			{
				val := int8(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:
			{
				val := int16(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			{
				val := int32(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			{
				val := v.Int()

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			{
				val :=

					uint(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			{
				val :=

					uint8(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint16:
			{
				val :=

					uint16(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint32:
			{
				val :=

					uint32(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint64:
			{
				val := v.Uint()

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uintptr:
			{
				val :=

					uintptr(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float32:
			{
				val :=

					float32(v.Float())

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float64:
			{
				val := v.Float()

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex64:
			{
				val :=

					complex64(v.Complex())

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex128:
			{
				val := v.Complex()

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Binds[index].SetComplex(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.String:
			{
				val := v.String()

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Binds[index].SetString(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		switch t.Kind() {
		case r.Bool:

			{
				val := v.Bool()

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*bool)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int:

			{
				val := int(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:
			{
				val := int8(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int8)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:
			{
				val := int16(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int16)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			{
				val := int32(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int32)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			{
				val := v.Int()

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int64)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			{
				val :=

					uint(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			{
				val :=

					uint8(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint8)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint16:
			{
				val :=

					uint16(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint16)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint32:
			{
				val :=

					uint32(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint32)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint64:
			{
				val := v.Uint()

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						IntBinds[index] = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uintptr:
			{
				val :=

					uintptr(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uintptr)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float32:
			{
				val :=

					float32(v.Float())

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float32)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float64:
			{
				val := v.Float()

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float64)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex64:
			{
				val :=

					complex64(v.Complex())

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*complex64)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex128:
			{
				val := v.Complex()

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Binds[index].SetComplex(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.String:
			{
				val := v.String()

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Binds[index].SetString(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case c.Depth - 1:
		switch t.Kind() {
		case r.Bool:

			{
				val := v.Bool()

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int:

			{
				val := int(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:
			{
				val := int8(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:
			{
				val := int16(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			{
				val := int32(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			{
				val := v.Int()

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			{
				val :=

					uint(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			{
				val :=

					uint8(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint16:
			{
				val :=

					uint16(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint32:
			{
				val :=

					uint32(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint64:
			{
				val := v.Uint()

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uintptr:
			{
				val :=

					uintptr(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float32:
			{
				val :=

					float32(v.Float())

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float64:
			{
				val := v.Float()

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex64:
			{
				val :=

					complex64(v.Complex())

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex128:
			{
				val := v.Complex()

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						Binds[index].SetComplex(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.String:
			{
				val := v.String()

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						Binds[index].SetString(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	}
	c.append(ret)
}
func (c *Comp) varSetExpr(upn int, index int, t xr.Type, e *Expr) {
	rt := t.ReflectType()
	fun := e.Fun
	var ret func(env *Env) (Stmt, *Env)
	switch upn {
	case 0:
		switch t.Kind() {
		case r.Bool:

			{
				fun := fun.(func(*Env) bool)

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int:

			{
				fun := fun.(func(*Env) int)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:

			{
				fun := fun.(func(*Env) int8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:

			{
				fun := fun.(func(*Env) int16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			{
				fun := fun.(func(*Env) int32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			{
				fun := fun.(func(*Env) int64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			{
				fun := fun.(func(*Env) uint)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			{
				fun := fun.(func(*Env) uint8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint16:
			{
				fun := fun.(func(*Env) uint16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint32:
			{
				fun := fun.(func(*Env) uint32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint64:
			{
				fun := fun.(func(*Env) uint64)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uintptr:
			{
				fun := fun.(func(*Env) uintptr)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float32:
			{
				fun := fun.(func(*Env) float32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float64:
			{
				fun := fun.(func(*Env) float64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex64:
			{
				fun := fun.(func(*Env) complex64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex128:
			{
				fun := fun.(func(*Env) complex128)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Binds[index].SetComplex(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.String:
			{
				fun := fun.(func(*Env) string)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Binds[index].SetString(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:
			{
				fun := e.AsX1()

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Binds[index].Set(fun(env).Convert(rt),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
	case 1:
		switch t.Kind() {
		case r.Bool:

			{
				fun := fun.(func(*Env) bool)

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int:

			{
				fun := fun.(func(*Env) int)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:

			{
				fun := fun.(func(*Env) int8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:

			{
				fun := fun.(func(*Env) int16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			{
				fun := fun.(func(*Env) int32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			{
				fun := fun.(func(*Env) int64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			{
				fun := fun.(func(*Env) uint)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			{
				fun := fun.(func(*Env) uint8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint16:
			{
				fun := fun.(func(*Env) uint16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint32:
			{
				fun := fun.(func(*Env) uint32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint64:
			{
				fun := fun.(func(*Env) uint64)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uintptr:
			{
				fun := fun.(func(*Env) uintptr)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float32:
			{
				fun := fun.(func(*Env) float32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float64:
			{
				fun := fun.(func(*Env) float64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex64:
			{
				fun := fun.(func(*Env) complex64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex128:
			{
				fun := fun.(func(*Env) complex128)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Binds[index].SetComplex(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.String:
			{
				fun := fun.(func(*Env) string)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Binds[index].SetString(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:
			{
				fun := e.AsX1()

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Binds[index].Set(fun(env).Convert(rt),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
	case 2:
		switch t.Kind() {
		case r.Bool:

			{
				fun := fun.(func(*Env) bool)

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int:

			{
				fun := fun.(func(*Env) int)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:

			{
				fun := fun.(func(*Env) int8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:

			{
				fun := fun.(func(*Env) int16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			{
				fun := fun.(func(*Env) int32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			{
				fun := fun.(func(*Env) int64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			{
				fun := fun.(func(*Env) uint)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			{
				fun := fun.(func(*Env) uint8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint16:
			{
				fun := fun.(func(*Env) uint16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint32:
			{
				fun := fun.(func(*Env) uint32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint64:
			{
				fun := fun.(func(*Env) uint64)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uintptr:
			{
				fun := fun.(func(*Env) uintptr)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float32:
			{
				fun := fun.(func(*Env) float32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float64:
			{
				fun := fun.(func(*Env) float64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex64:
			{
				fun := fun.(func(*Env) complex64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex128:
			{
				fun := fun.(func(*Env) complex128)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Binds[index].SetComplex(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.String:
			{
				fun := fun.(func(*Env) string)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Binds[index].SetString(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:
			{
				fun := e.AsX1()

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Binds[index].Set(fun(env).Convert(rt),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
	default:
		switch t.Kind() {
		case r.Bool:

			{
				fun := fun.(func(*Env) bool)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*bool)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int:

			{
				fun := fun.(func(*Env) int)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:

			{
				fun := fun.(func(*Env) int8)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int8)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:

			{
				fun := fun.(func(*Env) int16)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int16)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			{
				fun := fun.(func(*Env) int32)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int32)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			{
				fun := fun.(func(*Env) int64)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int64)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			{
				fun := fun.(func(*Env) uint)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			{
				fun := fun.(func(*Env) uint8)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint8)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint16:
			{
				fun := fun.(func(*Env) uint16)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint16)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint32:
			{
				fun := fun.(func(*Env) uint32)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint32)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint64:
			{
				fun := fun.(func(*Env) uint64)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						IntBinds[index] = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uintptr:
			{
				fun := fun.(func(*Env) uintptr)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uintptr)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float32:
			{
				fun := fun.(func(*Env) float32)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float32)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float64:
			{
				fun := fun.(func(*Env) float64)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float64)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex64:
			{
				fun := fun.(func(*Env) complex64)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*complex64)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex128:
			{
				fun := fun.(func(*Env) complex128)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Binds[index].SetComplex(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.String:
			{
				fun := fun.(func(*Env) string)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Binds[index].SetString(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:
			{
				fun := e.AsX1()

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Binds[index].Set(fun(env).Convert(rt),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
	case c.Depth - 1:
		switch t.Kind() {
		case r.Bool:

			{
				fun := fun.(func(*Env) bool)

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int:

			{
				fun := fun.(func(*Env) int)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:

			{
				fun := fun.(func(*Env) int8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:

			{
				fun := fun.(func(*Env) int16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			{
				fun := fun.(func(*Env) int32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			{
				fun := fun.(func(*Env) int64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			{
				fun := fun.(func(*Env) uint)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			{
				fun := fun.(func(*Env) uint8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint16:
			{
				fun := fun.(func(*Env) uint16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint32:
			{
				fun := fun.(func(*Env) uint32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uint64:
			{
				fun := fun.(func(*Env) uint64)

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Uintptr:
			{
				fun := fun.(func(*Env) uintptr)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float32:
			{
				fun := fun.(func(*Env) float32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Float64:
			{
				fun := fun.(func(*Env) float64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex64:
			{
				fun := fun.(func(*Env) complex64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.Complex128:
			{
				fun := fun.(func(*Env) complex128)

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						Binds[index].SetComplex(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case r.String:
			{
				fun := fun.(func(*Env) string)

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						Binds[index].SetString(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:
			{
				fun := e.AsX1()

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						Binds[index].Set(fun(env).Convert(rt),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
	}
	c.append(ret)
}
