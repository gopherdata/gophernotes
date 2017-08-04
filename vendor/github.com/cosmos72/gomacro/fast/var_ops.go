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
 * var_setops.go
 *
 *  Created on Apr 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/token"
	r "reflect"
	"unsafe"

	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) varAddConst(upn int, index int, t xr.Type, val I) {
	if isLiteralNumber(val, 0) || val == "" {
		return
	}

	{
		var ret Stmt
		switch t.Kind() {
		case r.Int:
			val := int(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:
			val := int8(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:
			val := int16(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			val := int32(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			val := r.ValueOf(val).Int()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			val := uint(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			val := uint8(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint16:
			val := uint16(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint32:
			val := uint32(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint64:
			val := r.ValueOf(val).Uint()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.IntBinds[index] += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uintptr:
			val := uintptr(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Float32:
			val :=

				float32(r.ValueOf(val).Float())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float32)(unsafe.Pointer(&o.IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Float64:
			val := r.ValueOf(val).Float()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float64)(unsafe.Pointer(&o.IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Complex64:
			val :=

				complex64(r.ValueOf(val).Complex())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex64)(unsafe.Pointer(&o.IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Complex128:
			val := r.ValueOf(val).Complex()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Binds[index]
						lhs.SetComplex(lhs.Complex() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Binds[index]
						lhs.SetComplex(lhs.Complex() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs :=

							env.ThreadGlobals.FileEnv.Binds[index]
						lhs.SetComplex(lhs.Complex() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.String:
			val := r.ValueOf(val).String()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Binds[index]
						lhs.SetString(lhs.String() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Binds[index]
						lhs.SetString(lhs.String() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Binds[index]
						lhs.SetString(lhs.String() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Binds[index]
						lhs.SetString(lhs.String() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.ThreadGlobals.FileEnv.
							Binds[index]
						lhs.SetString(lhs.String() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.ADD, t)

		}
		c.append(ret)
	}
}
func (c *Comp) varAddExpr(upn int, index int, t xr.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int)(unsafe.Pointer(&o.IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int8)(unsafe.Pointer(&o.IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int16)(unsafe.Pointer(&o.IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int32)(unsafe.Pointer(&o.IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int64)(unsafe.Pointer(&o.IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint)(unsafe.Pointer(&o.IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.IntBinds[index] += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*float32)(unsafe.Pointer(&o.IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*float64)(unsafe.Pointer(&o.IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*complex64)(unsafe.Pointer(&o.IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex128:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				{
					lhs :=

						o.Binds[index]
					lhs.SetComplex(lhs.Complex() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs :=

						env.ThreadGlobals.FileEnv.Binds[index]
					lhs.SetComplex(lhs.Complex() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) string:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				{
					lhs :=

						o.Binds[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.ThreadGlobals.FileEnv.
						Binds[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.ADD, t)

	}
	c.append(ret)
}
func (c *Comp) varSubConst(upn int, index int, t xr.Type, val I) {
	if isLiteralNumber(val, 0) {
		return
	}

	{
		var ret Stmt
		switch t.Kind() {
		case r.Int:
			val := int(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:
			val := int8(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:
			val := int16(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			val := int32(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			val := r.ValueOf(val).Int()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			val := uint(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			val := uint8(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint16:
			val := uint16(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint32:
			val := uint32(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint64:
			val := r.ValueOf(val).Uint()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.IntBinds[index] -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uintptr:
			val := uintptr(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Float32:
			val :=

				float32(r.ValueOf(val).Float())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float32)(unsafe.Pointer(&o.IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Float64:
			val := r.ValueOf(val).Float()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float64)(unsafe.Pointer(&o.IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Complex64:
			val :=

				complex64(r.ValueOf(val).Complex())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex64)(unsafe.Pointer(&o.IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Complex128:
			val := r.ValueOf(val).Complex()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Binds[index]
						lhs.SetComplex(lhs.Complex() -
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() -
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() -
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Binds[index]
						lhs.SetComplex(lhs.Complex() -
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs :=

							env.ThreadGlobals.FileEnv.Binds[index]
						lhs.SetComplex(lhs.Complex() -
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.SUB, t)

		}
		c.append(ret)
	}
}
func (c *Comp) varSubExpr(upn int, index int, t xr.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int)(unsafe.Pointer(&o.IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int8)(unsafe.Pointer(&o.IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int16)(unsafe.Pointer(&o.IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int32)(unsafe.Pointer(&o.IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int64)(unsafe.Pointer(&o.IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint)(unsafe.Pointer(&o.IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.IntBinds[index] -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*float32)(unsafe.Pointer(&o.IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*float64)(unsafe.Pointer(&o.IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*complex64)(unsafe.Pointer(&o.IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex128:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				{
					lhs :=

						o.Binds[index]
					lhs.SetComplex(lhs.Complex() -
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs :=

						env.ThreadGlobals.FileEnv.Binds[index]
					lhs.SetComplex(lhs.Complex() -
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.SUB, t)

	}
	c.append(ret)
}
func (c *Comp) varMulConst(upn int, index int, t xr.Type, val I) {
	if isLiteralNumber(val, 0) {

		c.varSetZero(upn, index, t)
		return
	} else if isLiteralNumber(val, 1) {
		return
	}

	{
		var ret Stmt
		switch t.Kind() {
		case r.Int:
			val := int(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:
			val := int8(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:
			val := int16(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			val := int32(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			val := r.ValueOf(val).Int()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			val := uint(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			val := uint8(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint16:
			val := uint16(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint32:
			val := uint32(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint64:
			val := r.ValueOf(val).Uint()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.IntBinds[index] *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uintptr:
			val := uintptr(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Float32:
			val :=

				float32(r.ValueOf(val).Float())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float32)(unsafe.Pointer(&o.IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Float64:
			val := r.ValueOf(val).Float()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float64)(unsafe.Pointer(&o.IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Complex64:
			val :=

				complex64(r.ValueOf(val).Complex())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex64)(unsafe.Pointer(&o.IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Complex128:
			val := r.ValueOf(val).Complex()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Binds[index]
						lhs.SetComplex(lhs.Complex() *
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() *
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() *
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Binds[index]
						lhs.SetComplex(lhs.Complex() *
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs :=

							env.ThreadGlobals.FileEnv.Binds[index]
						lhs.SetComplex(lhs.Complex() *
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.MUL, t)

		}
		c.append(ret)
	}
}
func (c *Comp) varMulExpr(upn int, index int, t xr.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int)(unsafe.Pointer(&o.IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int8)(unsafe.Pointer(&o.IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int16)(unsafe.Pointer(&o.IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int32)(unsafe.Pointer(&o.IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int64)(unsafe.Pointer(&o.IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint)(unsafe.Pointer(&o.IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.IntBinds[index] *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*float32)(unsafe.Pointer(&o.IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*float64)(unsafe.Pointer(&o.IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*complex64)(unsafe.Pointer(&o.IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex128:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				{
					lhs :=

						o.Binds[index]
					lhs.SetComplex(lhs.Complex() *
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs :=

						env.ThreadGlobals.FileEnv.Binds[index]
					lhs.SetComplex(lhs.Complex() *
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.MUL, t)

	}
	c.append(ret)
}
func (c *Comp) varQuoPow2(upn int, index int, t xr.Type, val I) bool {
	if isLiteralNumber(val, 0) {
		c.Errorf("division by %v <%v>", val, t)
		return false
	} else if isLiteralNumber(val, 1) {
		return true
	}

	ypositive := true
	yv := r.ValueOf(val)
	var y uint64
	switch KindToCategory(yv.Kind()) {
	case r.Int:
		sy := yv.Int()
		if sy < 0 {
			ypositive = false
			y = uint64(-sy)
		} else {
			y = uint64(sy)
		}

	case r.Uint:
		y = yv.Uint()
	default:
		return false
	}
	if !isPowerOfTwo(y) {
		return false
	}

	shift := integerLen(y) - 1
	var ret Stmt

	switch t.Kind() {
	case r.Int:
		switch upn {
		case 0:

			{
				y_1 :=

					int(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int)(unsafe.Pointer(&env.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int)(unsafe.Pointer(&env.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		case 1:

			{
				y_1 :=

					int(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int)(unsafe.Pointer(&env.
							Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int)(unsafe.Pointer(&env.
							Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		case 2:
			{
				y_1 :=

					int(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int)(unsafe.Pointer(&env.
							Outer.Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int)(unsafe.Pointer(&env.
							Outer.Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		default:
			{
				y_1 :=

					int(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						addr := (*int)(unsafe.Pointer(&o.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						addr := (*int)(unsafe.Pointer(&o.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		case c.Depth - 1:
			{
				y_1 :=

					int(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		}

	case r.Int8:
		switch upn {
		case 0:

			{
				y_1 :=

					int8(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int8)(unsafe.Pointer(&env.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int8)(unsafe.Pointer(&env.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		case 1:

			{
				y_1 :=

					int8(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int8)(unsafe.Pointer(&env.
							Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int8)(unsafe.Pointer(&env.
							Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		case 2:
			{
				y_1 :=

					int8(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int8)(unsafe.Pointer(&env.
							Outer.Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int8)(unsafe.Pointer(&env.
							Outer.Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		default:
			{
				y_1 :=

					int8(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						addr := (*int8)(unsafe.Pointer(&o.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						addr := (*int8)(unsafe.Pointer(&o.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		case c.Depth - 1:
			{
				y_1 :=

					int8(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		}

	case r.Int16:
		switch upn {
		case 0:

			{
				y_1 :=

					int16(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int16)(unsafe.Pointer(&env.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int16)(unsafe.Pointer(&env.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		case 1:

			{
				y_1 :=

					int16(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int16)(unsafe.Pointer(&env.
							Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int16)(unsafe.Pointer(&env.
							Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		case 2:
			{
				y_1 :=

					int16(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int16)(unsafe.Pointer(&env.
							Outer.Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int16)(unsafe.Pointer(&env.
							Outer.Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		default:
			{
				y_1 :=

					int16(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						addr := (*int16)(unsafe.Pointer(&o.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						addr := (*int16)(unsafe.Pointer(&o.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		case c.Depth - 1:
			{
				y_1 :=

					int16(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		}

	case r.Int32:
		switch upn {
		case 0:

			{
				y_1 :=

					int32(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int32)(unsafe.Pointer(&env.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int32)(unsafe.Pointer(&env.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		case 1:

			{
				y_1 :=

					int32(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int32)(unsafe.Pointer(&env.
							Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int32)(unsafe.Pointer(&env.
							Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		case 2:
			{
				y_1 :=

					int32(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int32)(unsafe.Pointer(&env.
							Outer.Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int32)(unsafe.Pointer(&env.
							Outer.Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		default:
			{
				y_1 :=

					int32(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						addr := (*int32)(unsafe.Pointer(&o.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						addr := (*int32)(unsafe.Pointer(&o.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		case c.Depth - 1:
			{
				y_1 :=

					int32(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		}

	case r.Int64:
		switch upn {
		case 0:

			{
				y_1 :=

					int64(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int64)(unsafe.Pointer(&env.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int64)(unsafe.Pointer(&env.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		case 1:

			{
				y_1 :=

					int64(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int64)(unsafe.Pointer(&env.
							Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int64)(unsafe.Pointer(&env.
							Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		case 2:
			{
				y_1 :=

					int64(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int64)(unsafe.Pointer(&env.
							Outer.Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int64)(unsafe.Pointer(&env.
							Outer.Outer.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		default:
			{
				y_1 :=

					int64(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						addr := (*int64)(unsafe.Pointer(&o.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						addr := (*int64)(unsafe.Pointer(&o.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		case c.Depth - 1:
			{
				y_1 :=

					int64(y - 1)
				if ypositive {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = n >> shift
						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {

						addr := (*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.IntBinds[index]))

						n := *addr
						if n < 0 {
							n += y_1
						}

						*addr = -(n >> shift)
						env.IP++
						return env.Code[env.IP], env
					}
				}

			}
		}

	case r.Uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {

				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint)(unsafe.Pointer(&o.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case r.Uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {

				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case r.Uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {

				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case r.Uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {

				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case r.Uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.IntBinds[index] >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.IntBinds[index] >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.IntBinds[index] >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {

				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.IntBinds[index] >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {

				env.ThreadGlobals.FileEnv.IntBinds[index] >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case r.Uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {

				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.IntBinds[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		}

	}
	if ret == nil {
		return false
	}

	c.append(ret)
	return true
}
func (c *Comp) varQuoConst(upn int, index int, t xr.Type, val I) {
	if c.varQuoPow2(upn, index, t, val) {
		return
	}

	{
		var ret Stmt
		switch t.Kind() {
		case r.Int:
			val := int(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:
			val := int8(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:
			val := int16(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			val := int32(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			val := r.ValueOf(val).Int()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			val := uint(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			val := uint8(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint16:
			val := uint16(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint32:
			val := uint32(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint64:
			val := r.ValueOf(val).Uint()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.IntBinds[index] /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uintptr:
			val := uintptr(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Float32:
			val :=

				float32(r.ValueOf(val).Float())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float32)(unsafe.Pointer(&o.IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Float64:
			val := r.ValueOf(val).Float()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float64)(unsafe.Pointer(&o.IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Complex64:
			val :=

				complex64(r.ValueOf(val).Complex())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex64)(unsafe.Pointer(&o.IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Complex128:
			val := r.ValueOf(val).Complex()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Binds[index]
						lhs.SetComplex(lhs.Complex() /
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() /
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() /
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Binds[index]
						lhs.SetComplex(lhs.Complex() /
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs :=

							env.ThreadGlobals.FileEnv.Binds[index]
						lhs.SetComplex(lhs.Complex() /
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.QUO, t)

		}
		c.append(ret)
	}
}
func (c *Comp) varQuoExpr(upn int, index int, t xr.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int)(unsafe.Pointer(&o.IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int8)(unsafe.Pointer(&o.IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int16)(unsafe.Pointer(&o.IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int32)(unsafe.Pointer(&o.IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int64)(unsafe.Pointer(&o.IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint)(unsafe.Pointer(&o.IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.IntBinds[index] /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*float32)(unsafe.Pointer(&o.IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*float64)(unsafe.Pointer(&o.IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*complex64)(unsafe.Pointer(&o.IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex128:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				{
					lhs :=

						o.Binds[index]
					lhs.SetComplex(lhs.Complex() /
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs :=

						env.ThreadGlobals.FileEnv.Binds[index]
					lhs.SetComplex(lhs.Complex() /
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.QUO, t)

	}
	c.append(ret)
}
func (c *Comp) varRemConst(upn int, index int, t xr.Type, val I) {
	if IsCategory(t.Kind(), r.Int, r.Uint) {
		if isLiteralNumber(val, 0) {
			c.Errorf("division by %v <%v>", val, t)
			return
		} else if isLiteralNumber(val, 1) {

			c.varSetZero(upn, index, t)
			return
		}
	}

	{
		var ret Stmt
		switch t.Kind() {
		case r.Int:
			val := int(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:
			val := int8(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:
			val := int16(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			val := int32(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			val := r.ValueOf(val).Int()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			val := uint(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			val := uint8(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint16:
			val := uint16(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint32:
			val := uint32(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint64:
			val := r.ValueOf(val).Uint()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.IntBinds[index] %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uintptr:
			val := uintptr(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.REM, t)

		}
		c.append(ret)
	}
}
func (c *Comp) varRemExpr(upn int, index int, t xr.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int)(unsafe.Pointer(&o.IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int8)(unsafe.Pointer(&o.IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int16)(unsafe.Pointer(&o.IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int32)(unsafe.Pointer(&o.IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int64)(unsafe.Pointer(&o.IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint)(unsafe.Pointer(&o.IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.IntBinds[index] %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.REM, t)

	}
	c.append(ret)
}
func (c *Comp) varAndConst(upn int, index int, t xr.Type, val I) {
	if IsCategory(t.Kind(), r.Int, r.Uint) {
		if isLiteralNumber(val, -1) {
			return
		} else if isLiteralNumber(val, 0) {

			c.varSetZero(upn, index, t)
			return
		}
	}

	{
		var ret Stmt
		switch t.Kind() {
		case r.Int:
			val := int(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:
			val := int8(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:
			val := int16(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			val := int32(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			val := r.ValueOf(val).Int()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			val := uint(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			val := uint8(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint16:
			val := uint16(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint32:
			val := uint32(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint64:
			val := r.ValueOf(val).Uint()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.IntBinds[index] &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uintptr:
			val := uintptr(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.AND, t)

		}
		c.append(ret)
	}
}
func (c *Comp) varAndExpr(upn int, index int, t xr.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int)(unsafe.Pointer(&o.IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int8)(unsafe.Pointer(&o.IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int16)(unsafe.Pointer(&o.IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int32)(unsafe.Pointer(&o.IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int64)(unsafe.Pointer(&o.IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint)(unsafe.Pointer(&o.IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.IntBinds[index] &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.AND, t)

	}
	c.append(ret)
}
func (c *Comp) varOrConst(upn int, index int, t xr.Type, val I) {
	if IsCategory(t.Kind(), r.Int, r.Uint) && isLiteralNumber(val, 0) {
		return
	}

	{
		var ret Stmt
		switch t.Kind() {
		case r.Int:
			val := int(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:
			val := int8(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:
			val := int16(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			val := int32(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			val := r.ValueOf(val).Int()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			val := uint(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			val := uint8(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint16:
			val := uint16(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint32:
			val := uint32(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint64:
			val := r.ValueOf(val).Uint()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.IntBinds[index] |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uintptr:
			val := uintptr(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.OR, t)

		}
		c.append(ret)
	}
}
func (c *Comp) varOrExpr(upn int, index int, t xr.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int)(unsafe.Pointer(&o.IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int8)(unsafe.Pointer(&o.IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int16)(unsafe.Pointer(&o.IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int32)(unsafe.Pointer(&o.IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int64)(unsafe.Pointer(&o.IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint)(unsafe.Pointer(&o.IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.IntBinds[index] |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.OR, t)

	}
	c.append(ret)
}
func (c *Comp) varXorConst(upn int, index int, t xr.Type, val I) {
	if IsCategory(t.Kind(), r.Int, r.Uint) && isLiteralNumber(val, 0) {
		return
	}

	{
		var ret Stmt
		switch t.Kind() {
		case r.Int:
			val := int(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:
			val := int8(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:
			val := int16(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			val := int32(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			val := r.ValueOf(val).Int()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			val := uint(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			val := uint8(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint16:
			val := uint16(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint32:
			val := uint32(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint64:
			val := r.ValueOf(val).Uint()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.IntBinds[index] ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uintptr:
			val := uintptr(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.XOR, t)

		}
		c.append(ret)
	}
}
func (c *Comp) varXorExpr(upn int, index int, t xr.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int)(unsafe.Pointer(&o.IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int8)(unsafe.Pointer(&o.IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int16)(unsafe.Pointer(&o.IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int32)(unsafe.Pointer(&o.IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int64)(unsafe.Pointer(&o.IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint)(unsafe.Pointer(&o.IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.XOR, t)

	}
	c.append(ret)
}
func (c *Comp) varAndnotConst(upn int, index int, t xr.Type, val I) {
	if IsCategory(t.Kind(), r.Int, r.Uint) {
		if isLiteralNumber(val, -1) {

			c.varSetZero(upn, index, t)
			return
		} else if isLiteralNumber(val, 0) {
			return
		}
	}

	{
		var ret Stmt
		switch t.Kind() {
		case r.Int:
			val := int(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int8:
			val := int8(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int16:
			val := int16(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int32:
			val := int32(r.ValueOf(val).Int())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Int64:
			val := r.ValueOf(val).Int()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			val := uint(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint8:
			val := uint8(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint16:
			val := uint16(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint32:
			val := uint32(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint64:
			val := r.ValueOf(val).Uint()
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.IntBinds[index] &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uintptr:
			val := uintptr(r.ValueOf(val).Uint())
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.AND_NOT, t)

		}
		c.append(ret)
	}
}
func (c *Comp) varAndnotExpr(upn int, index int, t xr.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int)(unsafe.Pointer(&o.IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int8)(unsafe.Pointer(&o.IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int16)(unsafe.Pointer(&o.IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int32)(unsafe.Pointer(&o.IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*int64)(unsafe.Pointer(&o.IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint)(unsafe.Pointer(&o.IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint8)(unsafe.Pointer(&o.IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint16)(unsafe.Pointer(&o.IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint32)(unsafe.Pointer(&o.IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.IntBinds[index] &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uintptr)(unsafe.Pointer(&o.IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.AND_NOT, t)

	}
	c.append(ret)
}
func (c *Comp) SetVar(va *Var, op token.Token, init *Expr) {
	t := va.Type
	var ok, shift bool
	switch op {
	case token.SHL, token.SHL_ASSIGN, token.SHR, token.SHR_ASSIGN:
		shift = true
		if init.Untyped() {
			init.ConstTo(c.TypeOfUint64())
			ok = true
		} else {
			ok = init.Type != nil && KindToCategory(init.Type.Kind()) == r.Uint
		}

	default:
		if init.Const() {
			init.ConstTo(t)
			ok = true
		} else if init.Type == nil {
			ok = op == token.ASSIGN && IsNillableKind(t.Kind())
		} else {
			ok = init.Type.AssignableTo(t)
		}

	}
	if !ok {
		c.Errorf("incompatible types in assignment: %v %s %v", t, op, init.Type)
		return
	}
	class := va.Desc.Class()
	if class != VarBind && class != IntBind {
		c.Errorf("invalid operator %s on %v", op, class)
		return
	}
	upn := va.Upn
	index := va.Desc.Index()
	if index == NoIndex {
		if op != token.ASSIGN {
			c.Errorf("invalid operator %s on _", op)
		}

		if !init.Const() {
			c.append(init.AsStmt())
		}

		return
	}
	if init.Const() {
		rt := t.ReflectType()
		val := init.Value
		v := r.ValueOf(val)
		if v == None || v == Nil {
			v = r.Zero(rt)
			val = v.Interface()
		} else if v.Type() != rt && !shift {
			v = v.Convert(rt)
			val = v.Interface()
		}
		switch op {
		case token.ASSIGN:
			c.varSetConst(upn, index, t, val)
		case token.ADD, token.ADD_ASSIGN:
			c.varAddConst(upn, index, t, val)
		case token.SUB, token.SUB_ASSIGN:
			c.varSubConst(upn, index, t, val)
		case token.MUL, token.MUL_ASSIGN:
			c.varMulConst(upn, index, t, val)
		case token.QUO, token.QUO_ASSIGN:
			c.varQuoConst(upn, index, t, val)
		case token.REM, token.REM_ASSIGN:
			c.varRemConst(upn, index, t, val)
		case token.AND, token.AND_ASSIGN:
			c.varAndConst(upn, index, t, val)
		case token.OR, token.OR_ASSIGN:
			c.varOrConst(upn, index, t, val)
		case token.XOR, token.XOR_ASSIGN:
			c.varXorConst(upn, index, t, val)
		case token.SHL, token.SHL_ASSIGN:
			c.varShlConst(upn, index, t, val)
		case token.SHR, token.SHR_ASSIGN:
			c.varShrConst(upn, index, t, val)
		case token.AND_NOT, token.AND_NOT_ASSIGN:
			c.varAndnotConst(upn, index, t, val)
		default:
			c.Errorf("invalid operator %s", op)
		}
	} else {
		fun := init.Fun
		switch op {
		case token.ASSIGN:
			c.varSetExpr(upn, index, t, init)
		case token.ADD, token.ADD_ASSIGN:
			c.varAddExpr(upn, index, t, fun)
		case token.SUB, token.SUB_ASSIGN:
			c.varSubExpr(upn, index, t, fun)
		case token.MUL, token.MUL_ASSIGN:
			c.varMulExpr(upn, index, t, fun)
		case token.QUO, token.QUO_ASSIGN:
			c.varQuoExpr(upn, index, t, fun)
		case token.REM, token.REM_ASSIGN:
			c.varRemExpr(upn, index, t, fun)
		case token.AND, token.AND_ASSIGN:
			c.varAndExpr(upn, index, t, fun)
		case token.OR, token.OR_ASSIGN:
			c.varOrExpr(upn, index, t, fun)
		case token.XOR, token.XOR_ASSIGN:
			c.varXorExpr(upn, index, t, fun)
		case token.SHL, token.SHL_ASSIGN:
			c.varShlExpr(upn, index, t, fun)
		case token.SHR, token.SHR_ASSIGN:
			c.varShrExpr(upn, index, t, fun)
		case token.AND_NOT, token.AND_NOT_ASSIGN:
			c.varAndnotExpr(upn, index, t, fun)
		default:
			c.Errorf("invalid operator %s", op)
		}
	}
}
