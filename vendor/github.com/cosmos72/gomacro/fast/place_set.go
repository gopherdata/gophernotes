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
 * place_set.go
 *
 *  Created on Apr 25, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) placeSetZero(place *Place) {
	rt := place.Type.ReflectType()
	zero := r.Zero(rt).Interface()
	c.placeSetConst(place, zero)
}
func (c *Comp) placeSetConst(place *Place, val I) {
	rt := place.Type.ReflectType()
	v := r.ValueOf(val)
	if ValueType(v) == nil {
		v = r.Zero(rt)
	} else {
		v = convert(v, rt)
	}

	lhs := place.Fun
	var ret Stmt
	if mapkey := place.MapKey; mapkey != nil {
		ret = func(env *Env) (Stmt, *Env) {

			obj := lhs(env)
			key := mapkey(env)
			obj.SetMapIndex(key, v)
			env.IP++
			return env.Code[env.IP], env
		}
		c.append(ret)
		return
	}
	switch KindToCategory(rt.Kind()) {
	case r.Bool:

		{
			val := v.Bool()

			ret = func(env *Env) (Stmt, *Env) {
				lhs(env).SetBool(val)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Int:

		{
			val := v.Int()

			ret = func(env *Env) (Stmt, *Env) {
				lhs(env).SetInt(val)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Uint:

		{
			val := v.Uint()

			ret = func(env *Env) (Stmt, *Env) {
				lhs(env).SetUint(val)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Float64:

		{
			val := v.Float()

			ret = func(env *Env) (Stmt, *Env) {
				lhs(env).SetFloat(val)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Complex128:

		{
			val := v.Complex()

			ret = func(env *Env) (Stmt, *Env) {
				lhs(env).SetComplex(val)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.String:

		{
			val := v.String()

			ret = func(env *Env) (Stmt, *Env) {
				lhs(env).SetString(val)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:

		{
			val := v

			ret = func(env *Env) (Stmt, *Env) {
				lhs(env).Set(val)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	}

	c.append(ret)
}
func (c *Comp) placeSetExpr(place *Place, fun I) {
	rt := place.Type.ReflectType()
	lhs := place.Fun
	var ret Stmt
	if mapkey := place.MapKey; mapkey != nil {
		rhs := funAsX1(fun, nil)
		ret = func(env *Env) (Stmt, *Env) {

			obj := lhs(env)
			key := mapkey(env)
			val := rhs(env)
			if val.Type() != rt {
				val = convert(val, rt)
			}

			obj.SetMapIndex(key, val)
			env.IP++
			return env.Code[env.IP], env
		}
		c.append(ret)
		return
	}
	switch rt.Kind() {
	case r.Bool:

		{
			rhs := fun.(func(*Env) bool)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetBool(value)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Int:

		{
			rhs := fun.(func(*Env) int)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetInt(int64(value))

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Int8:

		{
			rhs := fun.(func(*Env) int8)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetInt(int64(value))

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Int16:

		{
			rhs := fun.(func(*Env) int16)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetInt(int64(value))

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Int32:

		{
			rhs := fun.(func(*Env) int32)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetInt(int64(value))

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Int64:

		{
			rhs := fun.(func(*Env) int64)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetInt(int64(value))

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Uint:

		{
			rhs := fun.(func(*Env) uint)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetUint(uint64(value))

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Uint8:

		{
			rhs := fun.(func(*Env) uint8)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetUint(uint64(value))

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Uint16:

		{
			rhs := fun.(func(*Env) uint16)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetUint(uint64(value))

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Uint32:

		{
			rhs := fun.(func(*Env) uint32)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetUint(uint64(value))

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Uint64:

		{
			rhs := fun.(func(*Env) uint64)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetUint(uint64(value))

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Uintptr:

		{
			rhs := fun.(func(*Env) uintptr)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetUint(uint64(value))

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Float32:

		{
			rhs := fun.(func(*Env) float32)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetFloat(float64(value))

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Float64:

		{
			rhs := fun.(func(*Env) float64)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetFloat(float64(value))

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Complex64:

		{
			rhs := fun.(func(*Env) complex64)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetComplex(complex128(value))

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Complex128:

		{
			rhs := fun.(func(*Env) complex128)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetComplex(complex128(value))

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.String:

		{
			rhs := fun.(func(*Env) string)

			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				place.SetString(value)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		{
			rhs := funAsX1(fun, nil)

			zero := r.Zero(rt)
			ret = func(env *Env) (Stmt, *Env) {
				place := lhs(env)
				value := rhs(env)
				if value == Nil || value == None {
					value = zero
				} else if value.Type() != rt {
					value = convert(value, rt)
				}

				place.Set(value)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	}
	c.append(ret)
}
