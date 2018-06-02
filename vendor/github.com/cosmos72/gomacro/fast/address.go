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
 * address.go
 *
 *  Created on Apr 05, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"
	"unsafe"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) AddressOf(node *ast.UnaryExpr) *Expr { return c.addressOf(node.X, nil) }
func (c *Comp) addressOf(expr ast.Expr, t xr.Type) *Expr {
	for {
		switch e := expr.(type) {
		case *ast.ParenExpr:
			expr = e.X
			continue
		case *ast.StarExpr:

			if t != nil {
				t = t.Elem()
			}

			ret := c.Expr1(e.X, t)
			if ret.Type.Kind() != r.Ptr {
				c.Errorf("unary operation * on non-pointer <%v>: %v", ret.Type, e)
			}

		}
		break
	}
	place := c.placeOrAddress(expr, PlaceAddress, t)

	if place.IsVar() {
		va := place.Var
		return va.Address(c.Depth)
	} else if place.Addr == nil {
		c.Errorf("cannot take the address of %v <%v>", expr, place.Type)
		return nil
	} else {

		t := c.Universe.PtrTo(place.Type)
		return exprX1(t, place.Addr)
	}
}
func (c *Comp) AddressOfVar(name string) *Expr {
	sym := c.Resolve(name)
	va := sym.AsVar(PlaceAddress)
	return va.Address(c.Depth)
}
func (va *Var) Address(maxdepth int) *Expr {
	upn := va.Upn
	k := va.Type.Kind()
	index := va.Desc.Index()
	if index == NoIndex {
		base.Errorf("cannot take the address of %s: _", va.Desc.Class())
		return nil
	}
	var ret I
	intbinds := va.Desc.Class() == IntBind
	switch upn {
	case 0:
		switch k {
		case r.Bool:

			if intbinds {
				ret = func(env *Env) *bool {
					env.IntAddressTaken = true
					return (*bool)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *bool {
					return env.Vals[index].Addr().Interface().(*bool)
				}
			}
		case r.Int:

			if intbinds {
				ret = func(env *Env) *int {
					env.IntAddressTaken = true
					return (*int)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int {
					return env.Vals[index].Addr().Interface().(*int)
				}
			}
		case r.Int8:

			if intbinds {
				ret = func(env *Env) *int8 {
					env.IntAddressTaken = true
					return (*int8)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int8 {
					return env.Vals[index].Addr().Interface().(*int8)
				}
			}
		case r.Int16:

			if intbinds {
				ret = func(env *Env) *int16 {
					env.IntAddressTaken = true
					return (*int16)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int16 {
					return env.Vals[index].Addr().Interface().(*int16)
				}
			}
		case r.Int32:

			if intbinds {
				ret = func(env *Env) *int32 {
					env.IntAddressTaken = true
					return (*int32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int32 {
					return env.Vals[index].Addr().Interface().(*int32)
				}
			}
		case r.Int64:

			if intbinds {
				ret = func(env *Env) *int64 {
					env.IntAddressTaken = true
					return (*int64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int64 {
					return env.Vals[index].Addr().Interface().(*int64)
				}
			}
		case r.Uint:

			if intbinds {
				ret = func(env *Env) *uint {
					env.IntAddressTaken = true
					return (*uint)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint {
					return env.Vals[index].Addr().Interface().(*uint)
				}
			}
		case r.Uint8:

			if intbinds {
				ret = func(env *Env) *uint8 {
					env.IntAddressTaken = true
					return (*uint8)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint8 {
					return env.Vals[index].Addr().Interface().(*uint8)
				}
			}
		case r.Uint16:

			if intbinds {
				ret = func(env *Env) *uint16 {
					env.IntAddressTaken = true
					return (*uint16)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint16 {
					return env.Vals[index].Addr().Interface().(*uint16)
				}
			}
		case r.Uint32:

			if intbinds {
				ret = func(env *Env) *uint32 {
					env.IntAddressTaken = true
					return (*uint32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint32 {
					return env.Vals[index].Addr().Interface().(*uint32)
				}
			}
		case r.Uint64:

			if intbinds {
				ret = func(env *Env) *uint64 {
					env.IntAddressTaken = true
					return &env.Ints[index]

				}
			} else {
				ret = func(env *Env) *uint64 {
					return env.Vals[index].Addr().Interface().(*uint64)
				}
			}
		case r.Uintptr:

			if intbinds {
				ret = func(env *Env) *uintptr {
					env.IntAddressTaken = true
					return (*uintptr)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uintptr {
					return env.Vals[index].Addr().Interface().(*uintptr)
				}
			}
		case r.Float32:

			if intbinds {
				ret = func(env *Env) *float32 {
					env.IntAddressTaken = true
					return (*float32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *float32 {
					return env.Vals[index].Addr().Interface().(*float32)
				}
			}
		case r.Float64:

			if intbinds {
				ret = func(env *Env) *float64 {
					env.IntAddressTaken = true
					return (*float64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *float64 {
					return env.Vals[index].Addr().Interface().(*float64)
				}
			}
		case r.Complex64:

			if intbinds {
				ret = func(env *Env) *complex64 {
					env.IntAddressTaken = true
					return (*complex64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *complex64 {
					return env.Vals[index].Addr().Interface().(*complex64)
				}
			}
		case r.Complex128:

			if intbinds {
				ret = func(env *Env) *complex128 {
					env.IntAddressTaken = true
					return (*complex128)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *complex128 {
					return env.Vals[index].Addr().Interface().(*complex128)
				}
			}
		default:

			ret = func(env *Env) r.Value {
				return env.Vals[index].Addr()
			}
		}
	case 1:
		switch k {
		case r.Bool:

			if intbinds {
				ret = func(env *Env) *bool {
					env = env.
						Outer

					env.IntAddressTaken = true
					return (*bool)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *bool {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*bool)
				}
			}
		case r.Int:

			if intbinds {
				ret = func(env *Env) *int {
					env = env.
						Outer

					env.IntAddressTaken = true
					return (*int)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*int)
				}
			}
		case r.Int8:

			if intbinds {
				ret = func(env *Env) *int8 {
					env = env.
						Outer

					env.IntAddressTaken = true
					return (*int8)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int8 {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*int8)
				}
			}
		case r.Int16:

			if intbinds {
				ret = func(env *Env) *int16 {
					env = env.
						Outer

					env.IntAddressTaken = true
					return (*int16)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int16 {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*int16)
				}
			}
		case r.Int32:

			if intbinds {
				ret = func(env *Env) *int32 {
					env = env.
						Outer

					env.IntAddressTaken = true
					return (*int32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int32 {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*int32)
				}
			}
		case r.Int64:

			if intbinds {
				ret = func(env *Env) *int64 {
					env = env.
						Outer

					env.IntAddressTaken = true
					return (*int64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int64 {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*int64)
				}
			}
		case r.Uint:

			if intbinds {
				ret = func(env *Env) *uint {
					env = env.
						Outer

					env.IntAddressTaken = true
					return (*uint)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*uint)
				}
			}
		case r.Uint8:

			if intbinds {
				ret = func(env *Env) *uint8 {
					env = env.
						Outer

					env.IntAddressTaken = true
					return (*uint8)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint8 {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*uint8)
				}
			}
		case r.Uint16:

			if intbinds {
				ret = func(env *Env) *uint16 {
					env = env.
						Outer

					env.IntAddressTaken = true
					return (*uint16)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint16 {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*uint16)
				}
			}
		case r.Uint32:

			if intbinds {
				ret = func(env *Env) *uint32 {
					env = env.
						Outer

					env.IntAddressTaken = true
					return (*uint32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint32 {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*uint32)
				}
			}
		case r.Uint64:

			if intbinds {
				ret = func(env *Env) *uint64 {
					env = env.
						Outer

					env.IntAddressTaken = true
					return &env.Ints[index]

				}
			} else {
				ret = func(env *Env) *uint64 {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*uint64)
				}
			}
		case r.Uintptr:

			if intbinds {
				ret = func(env *Env) *uintptr {
					env = env.
						Outer

					env.IntAddressTaken = true
					return (*uintptr)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uintptr {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*uintptr)
				}
			}
		case r.Float32:

			if intbinds {
				ret = func(env *Env) *float32 {
					env = env.
						Outer

					env.IntAddressTaken = true
					return (*float32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *float32 {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*float32)
				}
			}
		case r.Float64:

			if intbinds {
				ret = func(env *Env) *float64 {
					env = env.
						Outer

					env.IntAddressTaken = true
					return (*float64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *float64 {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*float64)
				}
			}
		case r.Complex64:

			if intbinds {
				ret = func(env *Env) *complex64 {
					env = env.
						Outer

					env.IntAddressTaken = true
					return (*complex64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *complex64 {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*complex64)
				}
			}
		case r.Complex128:

			if intbinds {
				ret = func(env *Env) *complex128 {
					env = env.
						Outer

					env.IntAddressTaken = true
					return (*complex128)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *complex128 {
					env = env.
						Outer
					return env.Vals[index].Addr().Interface().(*complex128)
				}
			}
		default:

			ret = func(env *Env) r.Value {
				env = env.
					Outer
				return env.Vals[index].Addr()

			}
		}
	case 2:
		switch k {
		case r.Bool:

			if intbinds {
				ret = func(env *Env) *bool {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return (*bool)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *bool {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*bool)
				}
			}
		case r.Int:

			if intbinds {
				ret = func(env *Env) *int {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return (*int)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*int)
				}
			}
		case r.Int8:

			if intbinds {
				ret = func(env *Env) *int8 {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return (*int8)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int8 {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*int8)
				}
			}
		case r.Int16:

			if intbinds {
				ret = func(env *Env) *int16 {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return (*int16)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int16 {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*int16)
				}
			}
		case r.Int32:

			if intbinds {
				ret = func(env *Env) *int32 {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return (*int32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int32 {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*int32)
				}
			}
		case r.Int64:

			if intbinds {
				ret = func(env *Env) *int64 {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return (*int64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int64 {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*int64)
				}
			}
		case r.Uint:

			if intbinds {
				ret = func(env *Env) *uint {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return (*uint)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*uint)
				}
			}
		case r.Uint8:

			if intbinds {
				ret = func(env *Env) *uint8 {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return (*uint8)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint8 {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*uint8)
				}
			}
		case r.Uint16:

			if intbinds {
				ret = func(env *Env) *uint16 {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return (*uint16)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint16 {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*uint16)
				}
			}
		case r.Uint32:

			if intbinds {
				ret = func(env *Env) *uint32 {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return (*uint32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint32 {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*uint32)
				}
			}
		case r.Uint64:

			if intbinds {
				ret = func(env *Env) *uint64 {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return &env.Ints[index]

				}
			} else {
				ret = func(env *Env) *uint64 {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*uint64)
				}
			}
		case r.Uintptr:

			if intbinds {
				ret = func(env *Env) *uintptr {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return (*uintptr)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uintptr {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*uintptr)
				}
			}
		case r.Float32:

			if intbinds {
				ret = func(env *Env) *float32 {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return (*float32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *float32 {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*float32)
				}
			}
		case r.Float64:

			if intbinds {
				ret = func(env *Env) *float64 {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return (*float64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *float64 {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*float64)
				}
			}
		case r.Complex64:

			if intbinds {
				ret = func(env *Env) *complex64 {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return (*complex64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *complex64 {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*complex64)
				}
			}
		case r.Complex128:

			if intbinds {
				ret = func(env *Env) *complex128 {
					env = env.
						Outer.Outer

					env.IntAddressTaken = true
					return (*complex128)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *complex128 {
					env = env.
						Outer.Outer
					return env.Vals[index].Addr().Interface().(*complex128)
				}
			}
		default:

			ret = func(env *Env) r.Value {
				env = env.
					Outer.Outer
				return env.Vals[index].Addr()

			}
		}
	default:
		switch k {
		case r.Bool:

			if intbinds {
				ret = func(env *Env) *bool {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return (*bool)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *bool {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*bool)
				}
			}
		case r.Int:

			if intbinds {
				ret = func(env *Env) *int {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return (*int)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*int)
				}
			}
		case r.Int8:

			if intbinds {
				ret = func(env *Env) *int8 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return (*int8)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int8 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*int8)
				}
			}
		case r.Int16:

			if intbinds {
				ret = func(env *Env) *int16 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return (*int16)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int16 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*int16)
				}
			}
		case r.Int32:

			if intbinds {
				ret = func(env *Env) *int32 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return (*int32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int32 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*int32)
				}
			}
		case r.Int64:

			if intbinds {
				ret = func(env *Env) *int64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return (*int64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*int64)
				}
			}
		case r.Uint:

			if intbinds {
				ret = func(env *Env) *uint {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return (*uint)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*uint)
				}
			}
		case r.Uint8:

			if intbinds {
				ret = func(env *Env) *uint8 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return (*uint8)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint8 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*uint8)
				}
			}
		case r.Uint16:

			if intbinds {
				ret = func(env *Env) *uint16 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return (*uint16)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint16 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*uint16)
				}
			}
		case r.Uint32:

			if intbinds {
				ret = func(env *Env) *uint32 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return (*uint32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint32 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*uint32)
				}
			}
		case r.Uint64:

			if intbinds {
				ret = func(env *Env) *uint64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return &env.Ints[index]

				}
			} else {
				ret = func(env *Env) *uint64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*uint64)
				}
			}
		case r.Uintptr:

			if intbinds {
				ret = func(env *Env) *uintptr {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return (*uintptr)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uintptr {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*uintptr)
				}
			}
		case r.Float32:

			if intbinds {
				ret = func(env *Env) *float32 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return (*float32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *float32 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*float32)
				}
			}
		case r.Float64:

			if intbinds {
				ret = func(env *Env) *float64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return (*float64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *float64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*float64)
				}
			}
		case r.Complex64:

			if intbinds {
				ret = func(env *Env) *complex64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return (*complex64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *complex64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*complex64)
				}
			}
		case r.Complex128:

			if intbinds {
				ret = func(env *Env) *complex128 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.IntAddressTaken = true
					return (*complex128)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *complex128 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Vals[index].Addr().Interface().(*complex128)
				}
			}
		default:

			ret = func(env *Env) r.Value {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Vals[index].Addr()

			}
		}
	case maxdepth - 1:
		switch k {
		case r.Bool:

			if intbinds {
				ret = func(env *Env) *bool {
					env = env.FileEnv

					env.IntAddressTaken = true
					return (*bool)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *bool {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*bool)
				}
			}
		case r.Int:

			if intbinds {
				ret = func(env *Env) *int {
					env = env.FileEnv

					env.IntAddressTaken = true
					return (*int)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*int)
				}
			}
		case r.Int8:

			if intbinds {
				ret = func(env *Env) *int8 {
					env = env.FileEnv

					env.IntAddressTaken = true
					return (*int8)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int8 {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*int8)
				}
			}
		case r.Int16:

			if intbinds {
				ret = func(env *Env) *int16 {
					env = env.FileEnv

					env.IntAddressTaken = true
					return (*int16)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int16 {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*int16)
				}
			}
		case r.Int32:

			if intbinds {
				ret = func(env *Env) *int32 {
					env = env.FileEnv

					env.IntAddressTaken = true
					return (*int32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int32 {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*int32)
				}
			}
		case r.Int64:

			if intbinds {
				ret = func(env *Env) *int64 {
					env = env.FileEnv

					env.IntAddressTaken = true
					return (*int64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int64 {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*int64)
				}
			}
		case r.Uint:

			if intbinds {
				ret = func(env *Env) *uint {
					env = env.FileEnv

					env.IntAddressTaken = true
					return (*uint)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*uint)
				}
			}
		case r.Uint8:

			if intbinds {
				ret = func(env *Env) *uint8 {
					env = env.FileEnv

					env.IntAddressTaken = true
					return (*uint8)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint8 {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*uint8)
				}
			}
		case r.Uint16:

			if intbinds {
				ret = func(env *Env) *uint16 {
					env = env.FileEnv

					env.IntAddressTaken = true
					return (*uint16)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint16 {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*uint16)
				}
			}
		case r.Uint32:

			if intbinds {
				ret = func(env *Env) *uint32 {
					env = env.FileEnv

					env.IntAddressTaken = true
					return (*uint32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint32 {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*uint32)
				}
			}
		case r.Uint64:

			if intbinds {
				ret = func(env *Env) *uint64 {
					env = env.FileEnv

					env.IntAddressTaken = true
					return &env.Ints[index]

				}
			} else {
				ret = func(env *Env) *uint64 {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*uint64)
				}
			}
		case r.Uintptr:

			if intbinds {
				ret = func(env *Env) *uintptr {
					env = env.FileEnv

					env.IntAddressTaken = true
					return (*uintptr)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uintptr {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*uintptr)
				}
			}
		case r.Float32:

			if intbinds {
				ret = func(env *Env) *float32 {
					env = env.FileEnv

					env.IntAddressTaken = true
					return (*float32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *float32 {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*float32)
				}
			}
		case r.Float64:

			if intbinds {
				ret = func(env *Env) *float64 {
					env = env.FileEnv

					env.IntAddressTaken = true
					return (*float64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *float64 {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*float64)
				}
			}
		case r.Complex64:

			if intbinds {
				ret = func(env *Env) *complex64 {
					env = env.FileEnv

					env.IntAddressTaken = true
					return (*complex64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *complex64 {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*complex64)
				}
			}
		case r.Complex128:

			if intbinds {
				ret = func(env *Env) *complex128 {
					env = env.FileEnv

					env.IntAddressTaken = true
					return (*complex128)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *complex128 {
					env = env.FileEnv
					return env.Vals[index].Addr().Interface().(*complex128)
				}
			}
		default:

			ret = func(env *Env) r.Value {
				env = env.FileEnv
				return env.Vals[index].Addr()

			}
		}
	case maxdepth:
		switch k {
		case r.Bool:

			if intbinds {
				ret = func(env *Env) *bool {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return (*bool)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *bool {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*bool)
				}
			}
		case r.Int:

			if intbinds {
				ret = func(env *Env) *int {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return (*int)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*int)
				}
			}
		case r.Int8:

			if intbinds {
				ret = func(env *Env) *int8 {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return (*int8)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int8 {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*int8)
				}
			}
		case r.Int16:

			if intbinds {
				ret = func(env *Env) *int16 {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return (*int16)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int16 {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*int16)
				}
			}
		case r.Int32:

			if intbinds {
				ret = func(env *Env) *int32 {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return (*int32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int32 {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*int32)
				}
			}
		case r.Int64:

			if intbinds {
				ret = func(env *Env) *int64 {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return (*int64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *int64 {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*int64)
				}
			}
		case r.Uint:

			if intbinds {
				ret = func(env *Env) *uint {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return (*uint)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*uint)
				}
			}
		case r.Uint8:

			if intbinds {
				ret = func(env *Env) *uint8 {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return (*uint8)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint8 {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*uint8)
				}
			}
		case r.Uint16:

			if intbinds {
				ret = func(env *Env) *uint16 {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return (*uint16)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint16 {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*uint16)
				}
			}
		case r.Uint32:

			if intbinds {
				ret = func(env *Env) *uint32 {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return (*uint32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uint32 {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*uint32)
				}
			}
		case r.Uint64:

			if intbinds {
				ret = func(env *Env) *uint64 {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return &env.Ints[index]

				}
			} else {
				ret = func(env *Env) *uint64 {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*uint64)
				}
			}
		case r.Uintptr:

			if intbinds {
				ret = func(env *Env) *uintptr {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return (*uintptr)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *uintptr {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*uintptr)
				}
			}
		case r.Float32:

			if intbinds {
				ret = func(env *Env) *float32 {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return (*float32)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *float32 {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*float32)
				}
			}
		case r.Float64:

			if intbinds {
				ret = func(env *Env) *float64 {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return (*float64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *float64 {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*float64)
				}
			}
		case r.Complex64:

			if intbinds {
				ret = func(env *Env) *complex64 {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return (*complex64)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *complex64 {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*complex64)
				}
			}
		case r.Complex128:

			if intbinds {
				ret = func(env *Env) *complex128 {
					env = env.FileEnv.Outer

					env.IntAddressTaken = true
					return (*complex128)(unsafe.Pointer(&env.Ints[index]))

				}
			} else {
				ret = func(env *Env) *complex128 {
					env = env.FileEnv.Outer
					return env.Vals[index].Addr().Interface().(*complex128)
				}
			}
		default:

			ret = func(env *Env) r.Value {
				env = env.FileEnv.Outer
				return env.Vals[index].Addr()

			}
		}
	}
	u := va.Type.Universe()
	return &Expr{Lit: Lit{Type: u.PtrTo(va.Type)}, Fun: ret}
}
