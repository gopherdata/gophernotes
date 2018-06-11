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
 * place_ops.go
 *
 *  Created on Apr 25, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) placeAddConst(place *Place, val I) {
	if isLiteralNumber(val, 0) || val == "" {
		c.placeForSideEffects(place)
		return
	}

	{
		var ret Stmt
		lhsfun := place.Fun
		keyfun := place.MapKey
		v := r.ValueOf(val)

		if keyfun == nil {
			switch KindToCategory(place.Type.Kind()) {
			case r.Int:
				val := v.Int()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetInt(lhs.Int() + int64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := v.Uint()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetUint(lhs.Uint() + uint64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Float64:
				val := v.Float()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetFloat(lhs.Float() +
						val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Complex128:
				val := v.Complex()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetComplex(lhs.Complex() +
						val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.String:
				val := v.String()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetString(lhs.String() +
						val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.ADD, place.Type)

			}
		} else {
			switch place.Type.Kind() {
			case r.Int:
				val := int(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int

					if v := lhs.MapIndex(key); v != Nil {
						result = int(v.Int())
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int8:
				val := int8(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int8

					if v := lhs.MapIndex(key); v != Nil {
						result = int8(v.Int())
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int16:
				val := int16(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int16

					if v := lhs.MapIndex(key); v != Nil {
						result = int16(v.Int())
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int32:
				val := int32(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int32

					if v := lhs.MapIndex(key); v != Nil {
						result = int32(v.Int())
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int64:
				val := int64(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Int()
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := uint(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint

					if v := lhs.MapIndex(key); v != Nil {
						result = uint(v.Uint())
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint8:
				val := uint8(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint8

					if v := lhs.MapIndex(key); v != Nil {
						result = uint8(v.Uint())
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint16:
				val := uint16(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint16

					if v := lhs.MapIndex(key); v != Nil {
						result = uint16(v.Uint())
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint32:
				val := uint32(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint32

					if v := lhs.MapIndex(key); v != Nil {
						result = uint32(v.Uint())
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint64:
				val := uint64(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Uint()
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uintptr:
				val := uintptr(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uintptr

					if v := lhs.MapIndex(key); v != Nil {
						result = uintptr(v.Uint())
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Float32:
				val := float32(v.Float())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result float32

					if v := lhs.MapIndex(key); v != Nil {
						result = float32(v.Float())
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Float64:
				val := float64(v.Float())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result float64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Float()
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Complex64:
				val := complex64(v.Complex())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result complex64

					if v := lhs.MapIndex(key); v != Nil {
						result = complex64(v.Complex())
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Complex128:
				val := complex128(v.Complex())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result complex128

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Complex()
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.String:
				val := string(v.String())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result string

					if v := lhs.MapIndex(key); v != Nil {
						result = v.String()
					}
					result += val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.ADD, place.Type)

			}
		}

		c.append(ret)
	}
}
func (c *Comp) placeAddExpr(place *Place, fun I) {
	var ret Stmt
	lhsfun := place.Fun
	keyfun := place.MapKey
	if keyfun == nil {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() + int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() + int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() + int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() + int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() +
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() + uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() + uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() + uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() + uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() +
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() + uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() + float64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() +
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() + complex128(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() +
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) string:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetString(lhs.String() +
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.ADD, place.Type, funTypeOut(fun))

		}
	} else {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int

				if v := lhs.MapIndex(key); v != Nil {
					result = int(v.Int())
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int8

				if v := lhs.MapIndex(key); v != Nil {
					result = int8(v.Int())
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int16

				if v := lhs.MapIndex(key); v != Nil {
					result = int16(v.Int())
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int32

				if v := lhs.MapIndex(key); v != Nil {
					result = int32(v.Int())
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Int()
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint

				if v := lhs.MapIndex(key); v != Nil {
					result = uint(v.Uint())
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint8

				if v := lhs.MapIndex(key); v != Nil {
					result = uint8(v.Uint())
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint16

				if v := lhs.MapIndex(key); v != Nil {
					result = uint16(v.Uint())
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint32

				if v := lhs.MapIndex(key); v != Nil {
					result = uint32(v.Uint())
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Uint()
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uintptr

				if v := lhs.MapIndex(key); v != Nil {
					result = uintptr(v.Uint())
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result float32

				if v := lhs.MapIndex(key); v != Nil {
					result = float32(v.Float())
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result float64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Float()
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result complex64

				if v := lhs.MapIndex(key); v != Nil {
					result = complex64(v.Complex())
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result complex128

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Complex()
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) string:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result string

				if v := lhs.MapIndex(key); v != Nil {
					result = v.String()
				}
				result += fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.ADD, place.Type, funTypeOut(fun))

		}
	}

	c.append(ret)
}
func (c *Comp) placeSubConst(place *Place, val I) {
	if isLiteralNumber(val, 0) {
		c.placeForSideEffects(place)
		return
	}

	{
		var ret Stmt
		lhsfun := place.Fun
		keyfun := place.MapKey
		v := r.ValueOf(val)

		if keyfun == nil {
			switch KindToCategory(place.Type.Kind()) {
			case r.Int:
				val := v.Int()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetInt(lhs.Int() - int64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := v.Uint()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetUint(lhs.Uint() - uint64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Float64:
				val := v.Float()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetFloat(lhs.Float() -
						val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Complex128:
				val := v.Complex()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetComplex(lhs.Complex() -
						val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.SUB, place.Type)

			}
		} else {
			switch place.Type.Kind() {
			case r.Int:
				val := int(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int

					if v := lhs.MapIndex(key); v != Nil {
						result = int(v.Int())
					}
					result -= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int8:
				val := int8(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int8

					if v := lhs.MapIndex(key); v != Nil {
						result = int8(v.Int())
					}
					result -= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int16:
				val := int16(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int16

					if v := lhs.MapIndex(key); v != Nil {
						result = int16(v.Int())
					}
					result -= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int32:
				val := int32(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int32

					if v := lhs.MapIndex(key); v != Nil {
						result = int32(v.Int())
					}
					result -= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int64:
				val := int64(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Int()
					}
					result -= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := uint(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint

					if v := lhs.MapIndex(key); v != Nil {
						result = uint(v.Uint())
					}
					result -= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint8:
				val := uint8(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint8

					if v := lhs.MapIndex(key); v != Nil {
						result = uint8(v.Uint())
					}
					result -= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint16:
				val := uint16(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint16

					if v := lhs.MapIndex(key); v != Nil {
						result = uint16(v.Uint())
					}
					result -= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint32:
				val := uint32(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint32

					if v := lhs.MapIndex(key); v != Nil {
						result = uint32(v.Uint())
					}
					result -= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint64:
				val := uint64(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Uint()
					}
					result -= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uintptr:
				val := uintptr(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uintptr

					if v := lhs.MapIndex(key); v != Nil {
						result = uintptr(v.Uint())
					}
					result -= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Float32:
				val := float32(v.Float())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result float32

					if v := lhs.MapIndex(key); v != Nil {
						result = float32(v.Float())
					}
					result -= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Float64:
				val := float64(v.Float())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result float64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Float()
					}
					result -= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Complex64:
				val := complex64(v.Complex())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result complex64

					if v := lhs.MapIndex(key); v != Nil {
						result = complex64(v.Complex())
					}
					result -= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Complex128:
				val := complex128(v.Complex())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result complex128

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Complex()
					}
					result -= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.SUB, place.Type)

			}
		}

		c.append(ret)
	}
}
func (c *Comp) placeSubExpr(place *Place, fun I) {
	var ret Stmt
	lhsfun := place.Fun
	keyfun := place.MapKey
	if keyfun == nil {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() - int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() - int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() - int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() - int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() -
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() - uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() - uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() - uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() - uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() -
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() - uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() - float64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() -
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() - complex128(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() -
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SUB, place.Type, funTypeOut(fun))

		}
	} else {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int

				if v := lhs.MapIndex(key); v != Nil {
					result = int(v.Int())
				}
				result -= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int8

				if v := lhs.MapIndex(key); v != Nil {
					result = int8(v.Int())
				}
				result -= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int16

				if v := lhs.MapIndex(key); v != Nil {
					result = int16(v.Int())
				}
				result -= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int32

				if v := lhs.MapIndex(key); v != Nil {
					result = int32(v.Int())
				}
				result -= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Int()
				}
				result -= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint

				if v := lhs.MapIndex(key); v != Nil {
					result = uint(v.Uint())
				}
				result -= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint8

				if v := lhs.MapIndex(key); v != Nil {
					result = uint8(v.Uint())
				}
				result -= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint16

				if v := lhs.MapIndex(key); v != Nil {
					result = uint16(v.Uint())
				}
				result -= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint32

				if v := lhs.MapIndex(key); v != Nil {
					result = uint32(v.Uint())
				}
				result -= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Uint()
				}
				result -= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uintptr

				if v := lhs.MapIndex(key); v != Nil {
					result = uintptr(v.Uint())
				}
				result -= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result float32

				if v := lhs.MapIndex(key); v != Nil {
					result = float32(v.Float())
				}
				result -= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result float64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Float()
				}
				result -= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result complex64

				if v := lhs.MapIndex(key); v != Nil {
					result = complex64(v.Complex())
				}
				result -= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result complex128

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Complex()
				}
				result -= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SUB, place.Type, funTypeOut(fun))

		}
	}

	c.append(ret)
}
func (c *Comp) placeMulConst(place *Place, val I) {
	if isLiteralNumber(val, 0) {

		c.placeSetZero(place)
		return
	} else if isLiteralNumber(val, 1) {
		c.placeForSideEffects(place)
		return
	}

	{
		var ret Stmt
		lhsfun := place.Fun
		keyfun := place.MapKey
		v := r.ValueOf(val)

		if keyfun == nil {
			switch KindToCategory(place.Type.Kind()) {
			case r.Int:
				val := v.Int()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetInt(lhs.Int() * int64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := v.Uint()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetUint(lhs.Uint() * uint64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Float64:
				val := v.Float()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetFloat(lhs.Float() *
						val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Complex128:
				val := v.Complex()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetComplex(lhs.Complex() *
						val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.MUL, place.Type)

			}
		} else {
			switch place.Type.Kind() {
			case r.Int:
				val := int(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int

					if v := lhs.MapIndex(key); v != Nil {
						result = int(v.Int())
					}
					result *= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int8:
				val := int8(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int8

					if v := lhs.MapIndex(key); v != Nil {
						result = int8(v.Int())
					}
					result *= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int16:
				val := int16(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int16

					if v := lhs.MapIndex(key); v != Nil {
						result = int16(v.Int())
					}
					result *= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int32:
				val := int32(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int32

					if v := lhs.MapIndex(key); v != Nil {
						result = int32(v.Int())
					}
					result *= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int64:
				val := int64(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Int()
					}
					result *= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := uint(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint

					if v := lhs.MapIndex(key); v != Nil {
						result = uint(v.Uint())
					}
					result *= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint8:
				val := uint8(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint8

					if v := lhs.MapIndex(key); v != Nil {
						result = uint8(v.Uint())
					}
					result *= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint16:
				val := uint16(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint16

					if v := lhs.MapIndex(key); v != Nil {
						result = uint16(v.Uint())
					}
					result *= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint32:
				val := uint32(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint32

					if v := lhs.MapIndex(key); v != Nil {
						result = uint32(v.Uint())
					}
					result *= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint64:
				val := uint64(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Uint()
					}
					result *= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uintptr:
				val := uintptr(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uintptr

					if v := lhs.MapIndex(key); v != Nil {
						result = uintptr(v.Uint())
					}
					result *= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Float32:
				val := float32(v.Float())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result float32

					if v := lhs.MapIndex(key); v != Nil {
						result = float32(v.Float())
					}
					result *= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Float64:
				val := float64(v.Float())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result float64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Float()
					}
					result *= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Complex64:
				val := complex64(v.Complex())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result complex64

					if v := lhs.MapIndex(key); v != Nil {
						result = complex64(v.Complex())
					}
					result *= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Complex128:
				val := complex128(v.Complex())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result complex128

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Complex()
					}
					result *= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.MUL, place.Type)

			}
		}

		c.append(ret)
	}
}
func (c *Comp) placeMulExpr(place *Place, fun I) {
	var ret Stmt
	lhsfun := place.Fun
	keyfun := place.MapKey
	if keyfun == nil {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() * int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() * int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() * int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() * int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() *
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() * uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() * uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() * uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() * uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() *
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() * uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() * float64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() *
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() * complex128(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() *
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.MUL, place.Type, funTypeOut(fun))

		}
	} else {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int

				if v := lhs.MapIndex(key); v != Nil {
					result = int(v.Int())
				}
				result *= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int8

				if v := lhs.MapIndex(key); v != Nil {
					result = int8(v.Int())
				}
				result *= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int16

				if v := lhs.MapIndex(key); v != Nil {
					result = int16(v.Int())
				}
				result *= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int32

				if v := lhs.MapIndex(key); v != Nil {
					result = int32(v.Int())
				}
				result *= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Int()
				}
				result *= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint

				if v := lhs.MapIndex(key); v != Nil {
					result = uint(v.Uint())
				}
				result *= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint8

				if v := lhs.MapIndex(key); v != Nil {
					result = uint8(v.Uint())
				}
				result *= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint16

				if v := lhs.MapIndex(key); v != Nil {
					result = uint16(v.Uint())
				}
				result *= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint32

				if v := lhs.MapIndex(key); v != Nil {
					result = uint32(v.Uint())
				}
				result *= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Uint()
				}
				result *= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uintptr

				if v := lhs.MapIndex(key); v != Nil {
					result = uintptr(v.Uint())
				}
				result *= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result float32

				if v := lhs.MapIndex(key); v != Nil {
					result = float32(v.Float())
				}
				result *= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result float64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Float()
				}
				result *= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result complex64

				if v := lhs.MapIndex(key); v != Nil {
					result = complex64(v.Complex())
				}
				result *= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result complex128

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Complex()
				}
				result *= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.MUL, place.Type, funTypeOut(fun))

		}
	}

	c.append(ret)
}
func (c *Comp) placeQuoConst(place *Place, val I) {
	if c.placeQuoPow2(place, val) {
		return
	}

	{
		var ret Stmt
		lhsfun := place.Fun
		keyfun := place.MapKey
		v := r.ValueOf(val)

		if keyfun == nil {
			switch KindToCategory(place.Type.Kind()) {
			case r.Int:
				val := v.Int()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetInt(lhs.Int() / int64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := v.Uint()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetUint(lhs.Uint() / uint64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Float64:
				val := v.Float()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetFloat(lhs.Float() /
						val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Complex128:
				val := v.Complex()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetComplex(lhs.Complex() /
						val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.QUO, place.Type)

			}
		} else {
			switch place.Type.Kind() {
			case r.Int:
				val := int(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int

					if v := lhs.MapIndex(key); v != Nil {
						result = int(v.Int())
					}
					result /= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int8:
				val := int8(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int8

					if v := lhs.MapIndex(key); v != Nil {
						result = int8(v.Int())
					}
					result /= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int16:
				val := int16(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int16

					if v := lhs.MapIndex(key); v != Nil {
						result = int16(v.Int())
					}
					result /= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int32:
				val := int32(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int32

					if v := lhs.MapIndex(key); v != Nil {
						result = int32(v.Int())
					}
					result /= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int64:
				val := int64(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Int()
					}
					result /= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := uint(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint

					if v := lhs.MapIndex(key); v != Nil {
						result = uint(v.Uint())
					}
					result /= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint8:
				val := uint8(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint8

					if v := lhs.MapIndex(key); v != Nil {
						result = uint8(v.Uint())
					}
					result /= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint16:
				val := uint16(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint16

					if v := lhs.MapIndex(key); v != Nil {
						result = uint16(v.Uint())
					}
					result /= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint32:
				val := uint32(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint32

					if v := lhs.MapIndex(key); v != Nil {
						result = uint32(v.Uint())
					}
					result /= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint64:
				val := uint64(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Uint()
					}
					result /= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uintptr:
				val := uintptr(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uintptr

					if v := lhs.MapIndex(key); v != Nil {
						result = uintptr(v.Uint())
					}
					result /= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Float32:
				val := float32(v.Float())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result float32

					if v := lhs.MapIndex(key); v != Nil {
						result = float32(v.Float())
					}
					result /= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Float64:
				val := float64(v.Float())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result float64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Float()
					}
					result /= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Complex64:
				val := complex64(v.Complex())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result complex64

					if v := lhs.MapIndex(key); v != Nil {
						result = complex64(v.Complex())
					}
					result /= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Complex128:
				val := complex128(v.Complex())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result complex128

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Complex()
					}
					result /= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.QUO, place.Type)

			}
		}

		c.append(ret)
	}
}
func (c *Comp) placeQuoExpr(place *Place, fun I) {
	var ret Stmt
	lhsfun := place.Fun
	keyfun := place.MapKey
	if keyfun == nil {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() / int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() / int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() / int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() / int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() /
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() / uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() / uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() / uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() / uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() /
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() / uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() / float64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() /
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() / complex128(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() /
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.QUO, place.Type, funTypeOut(fun))

		}
	} else {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int

				if v := lhs.MapIndex(key); v != Nil {
					result = int(v.Int())
				}
				result /= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int8

				if v := lhs.MapIndex(key); v != Nil {
					result = int8(v.Int())
				}
				result /= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int16

				if v := lhs.MapIndex(key); v != Nil {
					result = int16(v.Int())
				}
				result /= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int32

				if v := lhs.MapIndex(key); v != Nil {
					result = int32(v.Int())
				}
				result /= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Int()
				}
				result /= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint

				if v := lhs.MapIndex(key); v != Nil {
					result = uint(v.Uint())
				}
				result /= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint8

				if v := lhs.MapIndex(key); v != Nil {
					result = uint8(v.Uint())
				}
				result /= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint16

				if v := lhs.MapIndex(key); v != Nil {
					result = uint16(v.Uint())
				}
				result /= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint32

				if v := lhs.MapIndex(key); v != Nil {
					result = uint32(v.Uint())
				}
				result /= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Uint()
				}
				result /= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uintptr

				if v := lhs.MapIndex(key); v != Nil {
					result = uintptr(v.Uint())
				}
				result /= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result float32

				if v := lhs.MapIndex(key); v != Nil {
					result = float32(v.Float())
				}
				result /= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result float64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Float()
				}
				result /= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result complex64

				if v := lhs.MapIndex(key); v != Nil {
					result = complex64(v.Complex())
				}
				result /= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result complex128

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Complex()
				}
				result /= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.QUO, place.Type, funTypeOut(fun))

		}
	}

	c.append(ret)
}
func (c *Comp) placeRemConst(place *Place, val I) {
	if IsCategory(place.Type.Kind(), r.Int, r.Uint) {
		if isLiteralNumber(val, 0) {
			c.Errorf("division by %v <%v>", val, place.Type)
			return
		} else if isLiteralNumber(val, 1) {

			c.placeSetZero(place)
			return
		}
	}

	{
		var ret Stmt
		lhsfun := place.Fun
		keyfun := place.MapKey
		v := r.ValueOf(val)

		if keyfun == nil {
			switch KindToCategory(place.Type.Kind()) {
			case r.Int:
				val := v.Int()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetInt(lhs.Int() % int64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := v.Uint()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetUint(lhs.Uint() % uint64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.REM, place.Type)

			}
		} else {
			switch place.Type.Kind() {
			case r.Int:
				val := int(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int

					if v := lhs.MapIndex(key); v != Nil {
						result = int(v.Int())
					}
					result %= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int8:
				val := int8(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int8

					if v := lhs.MapIndex(key); v != Nil {
						result = int8(v.Int())
					}
					result %= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int16:
				val := int16(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int16

					if v := lhs.MapIndex(key); v != Nil {
						result = int16(v.Int())
					}
					result %= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int32:
				val := int32(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int32

					if v := lhs.MapIndex(key); v != Nil {
						result = int32(v.Int())
					}
					result %= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int64:
				val := int64(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Int()
					}
					result %= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := uint(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint

					if v := lhs.MapIndex(key); v != Nil {
						result = uint(v.Uint())
					}
					result %= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint8:
				val := uint8(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint8

					if v := lhs.MapIndex(key); v != Nil {
						result = uint8(v.Uint())
					}
					result %= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint16:
				val := uint16(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint16

					if v := lhs.MapIndex(key); v != Nil {
						result = uint16(v.Uint())
					}
					result %= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint32:
				val := uint32(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint32

					if v := lhs.MapIndex(key); v != Nil {
						result = uint32(v.Uint())
					}
					result %= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint64:
				val := uint64(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Uint()
					}
					result %= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uintptr:
				val := uintptr(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uintptr

					if v := lhs.MapIndex(key); v != Nil {
						result = uintptr(v.Uint())
					}
					result %= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.REM, place.Type)

			}
		}

		c.append(ret)
	}
}
func (c *Comp) placeRemExpr(place *Place, fun I) {
	var ret Stmt
	lhsfun := place.Fun
	keyfun := place.MapKey
	if keyfun == nil {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() % int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() % int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() % int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() % int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() %
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() % uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() % uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() % uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() % uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() %
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() % uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.REM, place.Type, funTypeOut(fun))

		}
	} else {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int

				if v := lhs.MapIndex(key); v != Nil {
					result = int(v.Int())
				}
				result %= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int8

				if v := lhs.MapIndex(key); v != Nil {
					result = int8(v.Int())
				}
				result %= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int16

				if v := lhs.MapIndex(key); v != Nil {
					result = int16(v.Int())
				}
				result %= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int32

				if v := lhs.MapIndex(key); v != Nil {
					result = int32(v.Int())
				}
				result %= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Int()
				}
				result %= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint

				if v := lhs.MapIndex(key); v != Nil {
					result = uint(v.Uint())
				}
				result %= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint8

				if v := lhs.MapIndex(key); v != Nil {
					result = uint8(v.Uint())
				}
				result %= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint16

				if v := lhs.MapIndex(key); v != Nil {
					result = uint16(v.Uint())
				}
				result %= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint32

				if v := lhs.MapIndex(key); v != Nil {
					result = uint32(v.Uint())
				}
				result %= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Uint()
				}
				result %= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uintptr

				if v := lhs.MapIndex(key); v != Nil {
					result = uintptr(v.Uint())
				}
				result %= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.REM, place.Type, funTypeOut(fun))

		}
	}

	c.append(ret)
}
func (c *Comp) placeAndConst(place *Place, val I) {
	if IsCategory(place.Type.Kind(), r.Int, r.Uint) {
		if isLiteralNumber(val, -1) {
			c.placeForSideEffects(place)
			return
		} else if isLiteralNumber(val, 0) {

			c.placeSetZero(place)
			return
		}
	}

	{
		var ret Stmt
		lhsfun := place.Fun
		keyfun := place.MapKey
		v := r.ValueOf(val)

		if keyfun == nil {
			switch KindToCategory(place.Type.Kind()) {
			case r.Int:
				val := v.Int()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetInt(lhs.Int() & int64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := v.Uint()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetUint(lhs.Uint() & uint64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.AND, place.Type)

			}
		} else {
			switch place.Type.Kind() {
			case r.Int:
				val := int(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int

					if v := lhs.MapIndex(key); v != Nil {
						result = int(v.Int())
					}
					result &= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int8:
				val := int8(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int8

					if v := lhs.MapIndex(key); v != Nil {
						result = int8(v.Int())
					}
					result &= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int16:
				val := int16(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int16

					if v := lhs.MapIndex(key); v != Nil {
						result = int16(v.Int())
					}
					result &= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int32:
				val := int32(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int32

					if v := lhs.MapIndex(key); v != Nil {
						result = int32(v.Int())
					}
					result &= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int64:
				val := int64(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Int()
					}
					result &= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := uint(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint

					if v := lhs.MapIndex(key); v != Nil {
						result = uint(v.Uint())
					}
					result &= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint8:
				val := uint8(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint8

					if v := lhs.MapIndex(key); v != Nil {
						result = uint8(v.Uint())
					}
					result &= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint16:
				val := uint16(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint16

					if v := lhs.MapIndex(key); v != Nil {
						result = uint16(v.Uint())
					}
					result &= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint32:
				val := uint32(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint32

					if v := lhs.MapIndex(key); v != Nil {
						result = uint32(v.Uint())
					}
					result &= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint64:
				val := uint64(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Uint()
					}
					result &= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uintptr:
				val := uintptr(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uintptr

					if v := lhs.MapIndex(key); v != Nil {
						result = uintptr(v.Uint())
					}
					result &= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.AND, place.Type)

			}
		}

		c.append(ret)
	}
}
func (c *Comp) placeAndExpr(place *Place, fun I) {
	var ret Stmt
	lhsfun := place.Fun
	keyfun := place.MapKey
	if keyfun == nil {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() & int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() & int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() & int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() & int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() &
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() & uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() & uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() & uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() & uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() &
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() & uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.AND, place.Type, funTypeOut(fun))

		}
	} else {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int

				if v := lhs.MapIndex(key); v != Nil {
					result = int(v.Int())
				}
				result &= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int8

				if v := lhs.MapIndex(key); v != Nil {
					result = int8(v.Int())
				}
				result &= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int16

				if v := lhs.MapIndex(key); v != Nil {
					result = int16(v.Int())
				}
				result &= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int32

				if v := lhs.MapIndex(key); v != Nil {
					result = int32(v.Int())
				}
				result &= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Int()
				}
				result &= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint

				if v := lhs.MapIndex(key); v != Nil {
					result = uint(v.Uint())
				}
				result &= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint8

				if v := lhs.MapIndex(key); v != Nil {
					result = uint8(v.Uint())
				}
				result &= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint16

				if v := lhs.MapIndex(key); v != Nil {
					result = uint16(v.Uint())
				}
				result &= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint32

				if v := lhs.MapIndex(key); v != Nil {
					result = uint32(v.Uint())
				}
				result &= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Uint()
				}
				result &= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uintptr

				if v := lhs.MapIndex(key); v != Nil {
					result = uintptr(v.Uint())
				}
				result &= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.AND, place.Type, funTypeOut(fun))

		}
	}

	c.append(ret)
}
func (c *Comp) placeOrConst(place *Place, val I) {
	if IsCategory(place.Type.Kind(), r.Int, r.Uint) && isLiteralNumber(val, 0) {
		c.placeForSideEffects(place)
		return
	}

	{
		var ret Stmt
		lhsfun := place.Fun
		keyfun := place.MapKey
		v := r.ValueOf(val)

		if keyfun == nil {
			switch KindToCategory(place.Type.Kind()) {
			case r.Int:
				val := v.Int()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetInt(lhs.Int() | int64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := v.Uint()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetUint(lhs.Uint() | uint64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.OR, place.Type)

			}
		} else {
			switch place.Type.Kind() {
			case r.Int:
				val := int(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int

					if v := lhs.MapIndex(key); v != Nil {
						result = int(v.Int())
					}
					result |= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int8:
				val := int8(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int8

					if v := lhs.MapIndex(key); v != Nil {
						result = int8(v.Int())
					}
					result |= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int16:
				val := int16(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int16

					if v := lhs.MapIndex(key); v != Nil {
						result = int16(v.Int())
					}
					result |= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int32:
				val := int32(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int32

					if v := lhs.MapIndex(key); v != Nil {
						result = int32(v.Int())
					}
					result |= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int64:
				val := int64(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Int()
					}
					result |= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := uint(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint

					if v := lhs.MapIndex(key); v != Nil {
						result = uint(v.Uint())
					}
					result |= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint8:
				val := uint8(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint8

					if v := lhs.MapIndex(key); v != Nil {
						result = uint8(v.Uint())
					}
					result |= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint16:
				val := uint16(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint16

					if v := lhs.MapIndex(key); v != Nil {
						result = uint16(v.Uint())
					}
					result |= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint32:
				val := uint32(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint32

					if v := lhs.MapIndex(key); v != Nil {
						result = uint32(v.Uint())
					}
					result |= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint64:
				val := uint64(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Uint()
					}
					result |= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uintptr:
				val := uintptr(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uintptr

					if v := lhs.MapIndex(key); v != Nil {
						result = uintptr(v.Uint())
					}
					result |= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.OR, place.Type)

			}
		}

		c.append(ret)
	}
}
func (c *Comp) placeOrExpr(place *Place, fun I) {
	var ret Stmt
	lhsfun := place.Fun
	keyfun := place.MapKey
	if keyfun == nil {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() | int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() | int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() | int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() | int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() |
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() | uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() | uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() | uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() | uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() |
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() | uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.OR, place.Type, funTypeOut(fun))

		}
	} else {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int

				if v := lhs.MapIndex(key); v != Nil {
					result = int(v.Int())
				}
				result |= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int8

				if v := lhs.MapIndex(key); v != Nil {
					result = int8(v.Int())
				}
				result |= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int16

				if v := lhs.MapIndex(key); v != Nil {
					result = int16(v.Int())
				}
				result |= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int32

				if v := lhs.MapIndex(key); v != Nil {
					result = int32(v.Int())
				}
				result |= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Int()
				}
				result |= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint

				if v := lhs.MapIndex(key); v != Nil {
					result = uint(v.Uint())
				}
				result |= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint8

				if v := lhs.MapIndex(key); v != Nil {
					result = uint8(v.Uint())
				}
				result |= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint16

				if v := lhs.MapIndex(key); v != Nil {
					result = uint16(v.Uint())
				}
				result |= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint32

				if v := lhs.MapIndex(key); v != Nil {
					result = uint32(v.Uint())
				}
				result |= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Uint()
				}
				result |= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uintptr

				if v := lhs.MapIndex(key); v != Nil {
					result = uintptr(v.Uint())
				}
				result |= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.OR, place.Type, funTypeOut(fun))

		}
	}

	c.append(ret)
}
func (c *Comp) placeXorConst(place *Place, val I) {
	if IsCategory(place.Type.Kind(), r.Int, r.Uint) && isLiteralNumber(val, 0) {
		c.placeForSideEffects(place)
		return
	}

	{
		var ret Stmt
		lhsfun := place.Fun
		keyfun := place.MapKey
		v := r.ValueOf(val)

		if keyfun == nil {
			switch KindToCategory(place.Type.Kind()) {
			case r.Int:
				val := v.Int()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetInt(lhs.Int() ^ int64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := v.Uint()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetUint(lhs.Uint() ^ uint64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.XOR, place.Type)

			}
		} else {
			switch place.Type.Kind() {
			case r.Int:
				val := int(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int

					if v := lhs.MapIndex(key); v != Nil {
						result = int(v.Int())
					}
					result ^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int8:
				val := int8(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int8

					if v := lhs.MapIndex(key); v != Nil {
						result = int8(v.Int())
					}
					result ^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int16:
				val := int16(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int16

					if v := lhs.MapIndex(key); v != Nil {
						result = int16(v.Int())
					}
					result ^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int32:
				val := int32(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int32

					if v := lhs.MapIndex(key); v != Nil {
						result = int32(v.Int())
					}
					result ^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int64:
				val := int64(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Int()
					}
					result ^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := uint(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint

					if v := lhs.MapIndex(key); v != Nil {
						result = uint(v.Uint())
					}
					result ^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint8:
				val := uint8(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint8

					if v := lhs.MapIndex(key); v != Nil {
						result = uint8(v.Uint())
					}
					result ^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint16:
				val := uint16(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint16

					if v := lhs.MapIndex(key); v != Nil {
						result = uint16(v.Uint())
					}
					result ^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint32:
				val := uint32(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint32

					if v := lhs.MapIndex(key); v != Nil {
						result = uint32(v.Uint())
					}
					result ^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint64:
				val := uint64(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Uint()
					}
					result ^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uintptr:
				val := uintptr(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uintptr

					if v := lhs.MapIndex(key); v != Nil {
						result = uintptr(v.Uint())
					}
					result ^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.XOR, place.Type)

			}
		}

		c.append(ret)
	}
}
func (c *Comp) placeXorExpr(place *Place, fun I) {
	var ret Stmt
	lhsfun := place.Fun
	keyfun := place.MapKey
	if keyfun == nil {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() ^ int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() ^ int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() ^ int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() ^ int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() ^
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() ^ uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() ^ uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() ^ uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() ^ uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() ^
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() ^ uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.XOR, place.Type, funTypeOut(fun))

		}
	} else {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int

				if v := lhs.MapIndex(key); v != Nil {
					result = int(v.Int())
				}
				result ^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int8

				if v := lhs.MapIndex(key); v != Nil {
					result = int8(v.Int())
				}
				result ^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int16

				if v := lhs.MapIndex(key); v != Nil {
					result = int16(v.Int())
				}
				result ^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int32

				if v := lhs.MapIndex(key); v != Nil {
					result = int32(v.Int())
				}
				result ^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Int()
				}
				result ^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint

				if v := lhs.MapIndex(key); v != Nil {
					result = uint(v.Uint())
				}
				result ^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint8

				if v := lhs.MapIndex(key); v != Nil {
					result = uint8(v.Uint())
				}
				result ^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint16

				if v := lhs.MapIndex(key); v != Nil {
					result = uint16(v.Uint())
				}
				result ^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint32

				if v := lhs.MapIndex(key); v != Nil {
					result = uint32(v.Uint())
				}
				result ^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Uint()
				}
				result ^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uintptr

				if v := lhs.MapIndex(key); v != Nil {
					result = uintptr(v.Uint())
				}
				result ^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.XOR, place.Type, funTypeOut(fun))

		}
	}

	c.append(ret)
}
func (c *Comp) placeAndnotConst(place *Place, val I) {
	if IsCategory(place.Type.Kind(), r.Int, r.Uint) {
		if isLiteralNumber(val, -1) {

			c.placeSetZero(place)
			return
		} else if isLiteralNumber(val, 0) {
			c.placeForSideEffects(place)
			return
		}
	}

	{
		var ret Stmt
		lhsfun := place.Fun
		keyfun := place.MapKey
		v := r.ValueOf(val)

		if keyfun == nil {
			switch KindToCategory(place.Type.Kind()) {
			case r.Int:
				val := v.Int()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetInt(lhs.Int() &^ int64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := v.Uint()

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					lhs.SetUint(lhs.Uint() &^ uint64(val,
					),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.AND_NOT, place.Type)

			}
		} else {
			switch place.Type.Kind() {
			case r.Int:
				val := int(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int

					if v := lhs.MapIndex(key); v != Nil {
						result = int(v.Int())
					}
					result &^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int8:
				val := int8(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int8

					if v := lhs.MapIndex(key); v != Nil {
						result = int8(v.Int())
					}
					result &^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int16:
				val := int16(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int16

					if v := lhs.MapIndex(key); v != Nil {
						result = int16(v.Int())
					}
					result &^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int32:
				val := int32(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int32

					if v := lhs.MapIndex(key); v != Nil {
						result = int32(v.Int())
					}
					result &^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Int64:
				val := int64(v.Int())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result int64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Int()
					}
					result &^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:
				val := uint(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint

					if v := lhs.MapIndex(key); v != Nil {
						result = uint(v.Uint())
					}
					result &^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint8:
				val := uint8(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint8

					if v := lhs.MapIndex(key); v != Nil {
						result = uint8(v.Uint())
					}
					result &^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint16:
				val := uint16(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint16

					if v := lhs.MapIndex(key); v != Nil {
						result = uint16(v.Uint())
					}
					result &^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint32:
				val := uint32(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint32

					if v := lhs.MapIndex(key); v != Nil {
						result = uint32(v.Uint())
					}
					result &^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint64:
				val := uint64(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uint64

					if v := lhs.MapIndex(key); v != Nil {
						result = v.Uint()
					}
					result &^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uintptr:
				val := uintptr(v.Uint())

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					var result uintptr

					if v := lhs.MapIndex(key); v != Nil {
						result = uintptr(v.Uint())
					}
					result &^= val

					lhs.SetMapIndex(key, r.ValueOf(result))
					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf(`invalid operator %s= on <%v>`, token.AND_NOT, place.Type)

			}
		}

		c.append(ret)
	}
}
func (c *Comp) placeAndnotExpr(place *Place, fun I) {
	var ret Stmt
	lhsfun := place.Fun
	keyfun := place.MapKey
	if keyfun == nil {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() &^ int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() &^ int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() &^ int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() &^ int64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() &^
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() &^ uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() &^ uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() &^ uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() &^ uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() &^
					fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() &^ uint64(fun(env),
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.AND_NOT, place.Type, funTypeOut(fun))

		}
	} else {
		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int

				if v := lhs.MapIndex(key); v != Nil {
					result = int(v.Int())
				}
				result &^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int8

				if v := lhs.MapIndex(key); v != Nil {
					result = int8(v.Int())
				}
				result &^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int16

				if v := lhs.MapIndex(key); v != Nil {
					result = int16(v.Int())
				}
				result &^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int32

				if v := lhs.MapIndex(key); v != Nil {
					result = int32(v.Int())
				}
				result &^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result int64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Int()
				}
				result &^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint

				if v := lhs.MapIndex(key); v != Nil {
					result = uint(v.Uint())
				}
				result &^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint8

				if v := lhs.MapIndex(key); v != Nil {
					result = uint8(v.Uint())
				}
				result &^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint16

				if v := lhs.MapIndex(key); v != Nil {
					result = uint16(v.Uint())
				}
				result &^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint32

				if v := lhs.MapIndex(key); v != Nil {
					result = uint32(v.Uint())
				}
				result &^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uint64

				if v := lhs.MapIndex(key); v != Nil {
					result = v.Uint()
				}
				result &^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				key := keyfun(env)
				var result uintptr

				if v := lhs.MapIndex(key); v != Nil {
					result = uintptr(v.Uint())
				}
				result &^= fun(env)

				lhs.SetMapIndex(key, r.ValueOf(result))
				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.AND_NOT, place.Type, funTypeOut(fun))

		}
	}

	c.append(ret)
}
func (c *Comp) SetPlace(place *Place, op token.Token, init *Expr) {
	if place.IsVar() {
		c.SetVar(&place.Var, op, init)
		return
	}
	t := place.Type
	if init.Const() {
		init.ConstTo(t)
	} else if init.Type == nil || !init.Type.AssignableTo(t) {
		c.Errorf("incompatible types in assignment: <%v> %s <%v>", t, op, init.Type)
		return
	}
	rt := t.ReflectType()
	if init.Const() {
		val := init.Value
		v := r.ValueOf(val)
		if v == None || v == Nil {
			v = r.Zero(rt)
			val = v.Interface()
		} else if v.Type() != rt {
			v = convert(v, rt)
			val = v.Interface()
		}
		switch op {
		case token.ASSIGN:
			c.placeSetConst(place, val)
		case token.ADD, token.ADD_ASSIGN:
			c.placeAddConst(place, val)
		case token.SUB, token.SUB_ASSIGN:
			c.placeSubConst(place, val)
		case token.MUL, token.MUL_ASSIGN:
			c.placeMulConst(place, val)
		case token.QUO, token.QUO_ASSIGN:
			c.placeQuoConst(place, val)
		case token.REM, token.REM_ASSIGN:
			c.placeRemConst(place, val)
		case token.AND, token.AND_ASSIGN:
			c.placeAndConst(place, val)
		case token.OR, token.OR_ASSIGN:
			c.placeOrConst(place, val)
		case token.XOR, token.XOR_ASSIGN:
			c.placeAndConst(place, val)
		case token.AND_NOT, token.AND_NOT_ASSIGN:
			c.placeAndnotConst(place, val)
		default:
			c.Errorf("operator %s is not implemented", op)
		}
	} else {
		fun := init.Fun
		switch op {
		case token.ASSIGN:
			c.placeSetExpr(place, fun)
		case token.ADD, token.ADD_ASSIGN:
			c.placeAddExpr(place, fun)
		case token.SUB, token.SUB_ASSIGN:
			c.placeSubExpr(place, fun)
		case token.MUL, token.MUL_ASSIGN:
			c.placeMulExpr(place, fun)
		case token.QUO, token.QUO_ASSIGN:
			c.placeQuoExpr(place, fun)
		case token.REM, token.REM_ASSIGN:
			c.placeRemExpr(place, fun)
		case token.AND, token.AND_ASSIGN:
			c.placeAndExpr(place, fun)
		case token.OR, token.OR_ASSIGN:
			c.placeOrExpr(place, fun)
		case token.XOR, token.XOR_ASSIGN:
			c.placeAndExpr(place, fun)
		case token.AND_NOT, token.AND_NOT_ASSIGN:
			c.placeAndnotExpr(place, fun)
		default:
			c.Errorf("operator %s is not implemented", op)
		}
	}
}
