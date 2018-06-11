// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

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

func (c *Comp) varSetZero(va *Var) {
	zero := xr.Zero(va.Type).Interface()
	c.varSetConst(va, zero)
}
func (c *Comp) varSetConst(va *Var, val I) {
	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()

	v := r.ValueOf(val)
	rt := t.ReflectType()
	if ValueType(v) == nil {
		v = r.Zero(rt)
	} else {
		v = convert(v, rt)
	}

	var ret func(env *Env) (Stmt, *Env)
	intbinds := va.Desc.Class() == IntBind
	switch upn {
	case 0:
		switch t.Kind() {
		case r.Bool:

			{
				val := v.Bool()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*bool)(unsafe.Pointer(&env.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetBool(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int:

			{
				val := int(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int8:
			{
				val := int8(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int16:
			{
				val := int16(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int32:
			{
				val := int32(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int64:
			{
				val := v.Int()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetInt(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint:
			{
				val :=

					uint(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint8:
			{
				val :=

					uint8(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uint16:
			{
				val :=

					uint16(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uint32:
			{
				val :=

					uint32(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uint64:
			{
				val := v.Uint()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Ints[index] = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetUint(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uintptr:
			{
				val :=

					uintptr(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Float32:
			{
				val :=

					float32(v.Float())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float32)(unsafe.Pointer(&env.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetFloat(float64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Float64:
			{
				val := v.Float()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float64)(unsafe.Pointer(&env.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetFloat(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Complex64:
			{
				val :=

					complex64(v.Complex())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex64)(unsafe.Pointer(&env.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetComplex(complex128(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Complex128:
			{
				val := v.Complex()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex128)(unsafe.Pointer(&env.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetComplex(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.String:
			{
				val := v.String()

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Vals[index].SetString(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Vals[index].Set(v,
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

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*bool)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetBool(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int:

			{
				val := int(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int8:
			{
				val := int8(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int16:
			{
				val := int16(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int32:
			{
				val := int32(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int64:
			{
				val := v.Int()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetInt(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint:
			{
				val :=

					uint(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint8:
			{
				val :=

					uint8(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uint16:
			{
				val :=

					uint16(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uint32:
			{
				val :=

					uint32(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uint64:
			{
				val := v.Uint()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Ints[index] = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetUint(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uintptr:
			{
				val :=

					uintptr(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Float32:
			{
				val :=

					float32(v.Float())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float32)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetFloat(float64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Float64:
			{
				val := v.Float()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float64)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetFloat(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Complex64:
			{
				val :=

					complex64(v.Complex())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex64)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetComplex(complex128(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Complex128:
			{
				val := v.Complex()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex128)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetComplex(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.String:
			{
				val := v.String()

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Vals[index].SetString(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					Vals[index].Set(v,
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

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*bool)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetBool(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int:

			{
				val := int(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int8:
			{
				val := int8(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int16:
			{
				val := int16(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int32:
			{
				val := int32(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int64:
			{
				val := v.Int()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetInt(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint:
			{
				val :=

					uint(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint8:
			{
				val :=

					uint8(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uint16:
			{
				val :=

					uint16(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uint32:
			{
				val :=

					uint32(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uint64:
			{
				val := v.Uint()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Ints[index] = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetUint(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uintptr:
			{
				val :=

					uintptr(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Float32:
			{
				val :=

					float32(v.Float())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float32)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetFloat(float64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Float64:
			{
				val := v.Float()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float64)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetFloat(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Complex64:
			{
				val :=

					complex64(v.Complex())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex64)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetComplex(complex128(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Complex128:
			{
				val := v.Complex()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex128)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetComplex(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.String:
			{
				val := v.String()

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Vals[index].SetString(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					Vals[index].Set(v,
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

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*bool)(unsafe.Pointer(&o.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetBool(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int:

			{
				val := int(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*int)(unsafe.Pointer(&o.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int8:
			{
				val := int8(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*int8)(unsafe.Pointer(&o.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int16:
			{
				val := int16(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*int16)(unsafe.Pointer(&o.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int32:
			{
				val := int32(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*int32)(unsafe.Pointer(&o.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int64:
			{
				val := v.Int()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*int64)(unsafe.Pointer(&o.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetInt(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint:
			{
				val :=

					uint(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*uint)(unsafe.Pointer(&o.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint8:
			{
				val :=

					uint8(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*uint8)(unsafe.Pointer(&o.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uint16:
			{
				val :=

					uint16(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*uint16)(unsafe.Pointer(&o.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uint32:
			{
				val :=

					uint32(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*uint32)(unsafe.Pointer(&o.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uint64:
			{
				val := v.Uint()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Ints[index] = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetUint(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uintptr:
			{
				val :=

					uintptr(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*uintptr)(unsafe.Pointer(&o.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Float32:
			{
				val :=

					float32(v.Float())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*float32)(unsafe.Pointer(&o.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetFloat(float64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Float64:
			{
				val := v.Float()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*float64)(unsafe.Pointer(&o.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetFloat(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Complex64:
			{
				val :=

					complex64(v.Complex())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*complex64)(unsafe.Pointer(&o.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetComplex(complex128(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Complex128:
			{
				val := v.Complex()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*complex128)(unsafe.Pointer(&o.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetComplex(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
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
						Vals[index].SetString(val,
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
					Vals[index].Set(v,
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

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*bool)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetBool(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int:

			{
				val := int(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int8:
			{
				val := int8(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int16:
			{
				val := int16(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int32:
			{
				val := int32(v.Int())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetInt(int64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int64:
			{
				val := v.Int()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetInt(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint:
			{
				val :=

					uint(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint8:
			{
				val :=

					uint8(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uint16:
			{
				val :=

					uint16(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uint32:
			{
				val :=

					uint32(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uint64:
			{
				val := v.Uint()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Ints[index] = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetUint(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Uintptr:
			{
				val :=

					uintptr(v.Uint())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetUint(uint64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Float32:
			{
				val :=

					float32(v.Float())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float32)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetFloat(float64(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Float64:
			{
				val := v.Float()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float64)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetFloat(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Complex64:
			{
				val :=

					complex64(v.Complex())

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex64)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetComplex(complex128(val,
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.Complex128:
			{
				val := v.Complex()

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex128)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetComplex(val,
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.String:
			{
				val := v.String()

				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Vals[index].SetString(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.FileEnv.
					Vals[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	}
	c.append(ret)
}
func (c *Comp) varSetExpr(va *Var, e *Expr) {
	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()

	fun := e.Fun
	var ret func(env *Env) (Stmt, *Env)
	intbinds := va.Desc.Class() == IntBind
	switch upn {
	case 0:
		switch t.Kind() {
		case r.Bool:

			{
				fun := fun.(func(*Env) bool)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*bool)(unsafe.Pointer(&env.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetBool(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int:

			{
				fun := fun.(func(*Env) int)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int8:

			{
				fun := fun.(func(*Env) int8)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int16:

			{
				fun := fun.(func(*Env) int16)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int32:

			{
				fun := fun.(func(*Env) int32)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int64:

			{
				fun := fun.(func(*Env) int64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetInt(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint:

			{
				fun := fun.(func(*Env) uint)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint8:

			{
				fun := fun.(func(*Env) uint8)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint16:

			{
				fun := fun.(func(*Env) uint16)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint32:

			{
				fun := fun.(func(*Env) uint32)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint64:
			{
				fun := fun.(func(*Env) uint64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Ints[index] = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetUint(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uintptr:
			{
				fun := fun.(func(*Env) uintptr)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Float32:
			{
				fun := fun.(func(*Env) float32)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float32)(unsafe.Pointer(&env.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetFloat(float64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Float64:
			{
				fun := fun.(func(*Env) float64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float64)(unsafe.Pointer(&env.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetFloat(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Complex64:
			{
				fun := fun.(func(*Env) complex64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex64)(unsafe.Pointer(&env.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetComplex(complex128(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Complex128:
			{
				fun := fun.(func(*Env) complex128)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex128)(unsafe.Pointer(&env.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].SetComplex(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.String:
			{
				fun := fun.(func(*Env) string)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Vals[index].SetString(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:
			{
				fun := e.AsX1()
				if conv := c.Converter(e.Type, t); conv == nil {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].Set(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Vals[index].Set(conv(fun(env)),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		}
	case 1:
		switch t.Kind() {
		case r.Bool:

			{
				fun := fun.(func(*Env) bool)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*bool)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetBool(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int:

			{
				fun := fun.(func(*Env) int)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int8:

			{
				fun := fun.(func(*Env) int8)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int16:

			{
				fun := fun.(func(*Env) int16)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int32:

			{
				fun := fun.(func(*Env) int32)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int64:

			{
				fun := fun.(func(*Env) int64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetInt(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint:

			{
				fun := fun.(func(*Env) uint)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint8:

			{
				fun := fun.(func(*Env) uint8)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint16:

			{
				fun := fun.(func(*Env) uint16)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint32:

			{
				fun := fun.(func(*Env) uint32)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint64:
			{
				fun := fun.(func(*Env) uint64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Ints[index] = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetUint(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uintptr:
			{
				fun := fun.(func(*Env) uintptr)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Float32:
			{
				fun := fun.(func(*Env) float32)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float32)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetFloat(float64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Float64:
			{
				fun := fun.(func(*Env) float64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float64)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetFloat(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Complex64:
			{
				fun := fun.(func(*Env) complex64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex64)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetComplex(complex128(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Complex128:
			{
				fun := fun.(func(*Env) complex128)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex128)(unsafe.Pointer(&env.
							Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].SetComplex(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.String:
			{
				fun := fun.(func(*Env) string)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Vals[index].SetString(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:
			{
				fun := e.AsX1()
				if conv := c.Converter(e.Type, t); conv == nil {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].Set(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Vals[index].Set(conv(fun(env)),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		}
	case 2:
		switch t.Kind() {
		case r.Bool:

			{
				fun := fun.(func(*Env) bool)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*bool)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetBool(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int:

			{
				fun := fun.(func(*Env) int)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int8:

			{
				fun := fun.(func(*Env) int8)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int16:

			{
				fun := fun.(func(*Env) int16)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int32:

			{
				fun := fun.(func(*Env) int32)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int64:

			{
				fun := fun.(func(*Env) int64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetInt(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint:

			{
				fun := fun.(func(*Env) uint)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint8:

			{
				fun := fun.(func(*Env) uint8)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint16:

			{
				fun := fun.(func(*Env) uint16)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint32:

			{
				fun := fun.(func(*Env) uint32)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint64:
			{
				fun := fun.(func(*Env) uint64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Ints[index] = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetUint(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uintptr:
			{
				fun := fun.(func(*Env) uintptr)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Float32:
			{
				fun := fun.(func(*Env) float32)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float32)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetFloat(float64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Float64:
			{
				fun := fun.(func(*Env) float64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float64)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetFloat(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Complex64:
			{
				fun := fun.(func(*Env) complex64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex64)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetComplex(complex128(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Complex128:
			{
				fun := fun.(func(*Env) complex128)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex128)(unsafe.Pointer(&env.
							Outer.Outer.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].SetComplex(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.String:
			{
				fun := fun.(func(*Env) string)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Vals[index].SetString(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:
			{
				fun := e.AsX1()
				if conv := c.Converter(e.Type, t); conv == nil {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].Set(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Vals[index].Set(conv(fun(env)),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		}
	default:
		switch t.Kind() {
		case r.Bool:

			{
				fun := fun.(func(*Env) bool)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*bool)(unsafe.Pointer(&o.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetBool(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int:

			{
				fun := fun.(func(*Env) int)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*int)(unsafe.Pointer(&o.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int8:

			{
				fun := fun.(func(*Env) int8)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*int8)(unsafe.Pointer(&o.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int16:

			{
				fun := fun.(func(*Env) int16)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*int16)(unsafe.Pointer(&o.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int32:

			{
				fun := fun.(func(*Env) int32)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*int32)(unsafe.Pointer(&o.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int64:

			{
				fun := fun.(func(*Env) int64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*int64)(unsafe.Pointer(&o.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetInt(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint:

			{
				fun := fun.(func(*Env) uint)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*uint)(unsafe.Pointer(&o.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint8:

			{
				fun := fun.(func(*Env) uint8)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*uint8)(unsafe.Pointer(&o.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint16:

			{
				fun := fun.(func(*Env) uint16)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*uint16)(unsafe.Pointer(&o.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint32:

			{
				fun := fun.(func(*Env) uint32)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*uint32)(unsafe.Pointer(&o.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint64:
			{
				fun := fun.(func(*Env) uint64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Ints[index] = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetUint(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uintptr:
			{
				fun := fun.(func(*Env) uintptr)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*uintptr)(unsafe.Pointer(&o.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Float32:
			{
				fun := fun.(func(*Env) float32)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*float32)(unsafe.Pointer(&o.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetFloat(float64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Float64:
			{
				fun := fun.(func(*Env) float64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*float64)(unsafe.Pointer(&o.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetFloat(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Complex64:
			{
				fun := fun.(func(*Env) complex64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*complex64)(unsafe.Pointer(&o.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetComplex(complex128(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Complex128:
			{
				fun := fun.(func(*Env) complex128)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						*(*complex128)(unsafe.Pointer(&o.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].SetComplex(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
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
						Vals[index].SetString(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:
			{
				fun := e.AsX1()
				if conv := c.Converter(e.Type, t); conv == nil {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].Set(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.
							Vals[index].Set(conv(fun(env)),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		}
	case c.Depth - 1:
		switch t.Kind() {
		case r.Bool:

			{
				fun := fun.(func(*Env) bool)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*bool)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetBool(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int:

			{
				fun := fun.(func(*Env) int)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int8:

			{
				fun := fun.(func(*Env) int8)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int16:

			{
				fun := fun.(func(*Env) int16)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int32:

			{
				fun := fun.(func(*Env) int32)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetInt(int64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int64:

			{
				fun := fun.(func(*Env) int64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetInt(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint:

			{
				fun := fun.(func(*Env) uint)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint8:

			{
				fun := fun.(func(*Env) uint8)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint16:

			{
				fun := fun.(func(*Env) uint16)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint32:

			{
				fun := fun.(func(*Env) uint32)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint64:
			{
				fun := fun.(func(*Env) uint64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Ints[index] = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetUint(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uintptr:
			{
				fun := fun.(func(*Env) uintptr)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetUint(uint64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Float32:
			{
				fun := fun.(func(*Env) float32)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float32)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetFloat(float64(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Float64:
			{
				fun := fun.(func(*Env) float64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*float64)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetFloat(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Complex64:
			{
				fun := fun.(func(*Env) complex64)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex64)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetComplex(complex128(fun(env),
						))

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Complex128:
			{
				fun := fun.(func(*Env) complex128)

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*complex128)(unsafe.Pointer(&env.FileEnv.
							Ints[index])) = fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].SetComplex(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		case r.String:
			{
				fun := fun.(func(*Env) string)

				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Vals[index].SetString(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:
			{
				fun := e.AsX1()
				if conv := c.Converter(e.Type, t); conv == nil {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].Set(fun(env),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Vals[index].Set(conv(fun(env)),
						)

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}

		}
	}
	c.append(ret)
}
