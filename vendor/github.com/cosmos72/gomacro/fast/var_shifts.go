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
 * var_shifts.go
 *
 *  Created on May 17, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/token"
	r "reflect"
	"unsafe"

	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) varShlConst(va *Var, val I) {
	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind

	t2 := r.TypeOf(val)
	if t2 == nil || KindToCategory(t2.Kind()) != r.Uint {
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHL, t, t2)
	}

	if isLiteralNumber(val, 0) {
		return
	}
	{
		val := r.ValueOf(val).Uint()
		var ret Stmt
		switch t.Kind() {
		case r.Int:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int)(unsafe.Pointer(&o.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int8:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int8)(unsafe.Pointer(&o.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int16:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int16)(unsafe.Pointer(&o.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int32:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int32)(unsafe.Pointer(&o.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int64:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int64)(unsafe.Pointer(&o.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint)(unsafe.Pointer(&o.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint8:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint8)(unsafe.Pointer(&o.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint16:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint16)(unsafe.Pointer(&o.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint32:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint32)(unsafe.Pointer(&o.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint64:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Ints[index] <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Ints[index] <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Ints[index] <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Ints[index] <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.Ints[index] <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uintptr:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uintptr)(unsafe.Pointer(&o.Ints[index])) <<= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() <<
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHL, t, t2)

		}
		c.append(ret)
	}
}
func (c *Comp) varShlExpr(va *Var, function I) {
	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind

	t2 := funTypeOut(function)
	if t2 == nil || KindToCategory(t2.Kind()) != r.Uint {
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHL, t, t2)
	}

	fun := asFunUint8(function)
	{
		var ret Stmt
		switch t.Kind() {
		case r.Int:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int8:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int8)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int16:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int16)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int32:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int32)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int64:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int64)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint8:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint8)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint16:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint16)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint32:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint32)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint64:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Ints[index] <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Ints[index] <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Ints[index] <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Ints[index] <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.Ints[index] <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uintptr:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.
							Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uintptr)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() <<
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHL, t, t2)

		}
		c.append(ret)
	}
}
func (c *Comp) varShrConst(va *Var, val I) {
	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind

	t2 := r.TypeOf(val)
	if t2 == nil || KindToCategory(t2.Kind()) != r.Uint {
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHR, t, t2)
	}

	if isLiteralNumber(val, 0) {
		return
	}
	{
		val := r.ValueOf(val).Uint()
		var ret Stmt
		switch t.Kind() {
		case r.Int:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int)(unsafe.Pointer(&o.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int8:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int8)(unsafe.Pointer(&o.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int16:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int16)(unsafe.Pointer(&o.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int32:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int32)(unsafe.Pointer(&o.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int64:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int64)(unsafe.Pointer(&o.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint)(unsafe.Pointer(&o.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint8:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint8)(unsafe.Pointer(&o.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint16:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint16)(unsafe.Pointer(&o.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint32:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint32)(unsafe.Pointer(&o.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint64:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Ints[index] >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Ints[index] >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Ints[index] >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Ints[index] >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.Ints[index] >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uintptr:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uintptr)(unsafe.Pointer(&o.Ints[index])) >>= val

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() >>
								val,
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHR, t, t2)

		}
		c.append(ret)
	}
}
func (c *Comp) varShrExpr(va *Var, function I) {
	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind

	t2 := funTypeOut(function)
	if t2 == nil || KindToCategory(t2.Kind()) != r.Uint {
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHR, t, t2)
	}

	fun := asFunUint8(function)
	{
		var ret Stmt
		switch t.Kind() {
		case r.Int:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int8:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int8)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int16:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int16)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int32:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int32)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Int64:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*int64)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetInt(lhs.Int() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint8:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint8)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint16:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint16)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint32:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uint32)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uint64:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Ints[index] >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.
							Ints[index] >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.
							Outer.Outer.
							Ints[index] >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						env.FileEnv.
							Ints[index] >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}

						o.Ints[index] >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		case r.Uintptr:
			switch upn {
			case 0:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.
							Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case 2:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.
							Outer.Outer.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.
								Outer.Outer.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			case c.Depth - 1:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						{
							lhs := env.FileEnv.
								Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			default:

				if intbinds {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						*(*uintptr)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

						env.IP++
						return env.Code[env.IP], env
					}
				} else {
					ret = func(env *Env) (Stmt, *Env) {
						o := env.Outer.Outer.Outer
						for i := 3; i < upn; i++ {
							o = o.Outer
						}
						{
							lhs :=

								o.Vals[index]
							lhs.SetUint(lhs.Uint() >>
								fun(env),
							)
						}

						env.IP++
						return env.Code[env.IP], env
					}
				}
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHR, t, t2)

		}
		c.append(ret)
	}
}
func asFunUint8(fun I) func(*Env) uint8 {
	var f func(*Env) uint8
	switch fun := fun.(type) {
	case func(*Env) uint:
		return func(env *Env) uint8 {
			val := fun(env)
			if val > 255 {
				val = 255
			}
			return uint8(val)
		}
	case func(*Env) uint8:
		f = fun
	case func(*Env) uint16:
		return func(env *Env) uint8 {
			val := fun(env)
			if val > 255 {
				val = 255
			}
			return uint8(val)
		}
	case func(*Env) uint32:
		return func(env *Env) uint8 {
			val := fun(env)
			if val > 255 {
				val = 255
			}
			return uint8(val)
		}

	case func(*Env) uint64:
		return func(env *Env) uint8 {
			val := fun(env)
			if val > 255 {
				val = 255
			}
			return uint8(val)
		}

	case func(*Env) uintptr:
		return func(env *Env) uint8 {
			val := fun(env)
			if val > 255 {
				val = 255
			}
			return uint8(val)
		}

	}
	return f
}
