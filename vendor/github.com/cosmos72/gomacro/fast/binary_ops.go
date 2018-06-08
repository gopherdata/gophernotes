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
 * binary_ops.go
 *
 *  Created on Apr 12, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) Add(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) + y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) + y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) + y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) + y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) + y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) + y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) + y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) + y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) + y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) + y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) + y(env)
				}

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x(env) + y(env)
				}

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x(env) + y(env)
				}

			}

		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x(env) + y(env)
				}

			}

		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x(env) + y(env)
				}

			}

		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(func(*Env) string)
				fun = func(env *Env) string {
					return x(env) + y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		if isLiteralNumber(y, 0) || y == "" {
			return xe
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) + y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) + y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) + y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) + y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) + y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) + y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) + y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) + y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) + y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) + y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) + y
				}

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := float32(r.ValueOf(y).Float())
				fun = func(env *Env) float32 {
					return x(env) + y
				}

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := r.ValueOf(y).Float()
				fun = func(env *Env) float64 {
					return x(env) + y
				}

			}
		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := complex64(r.ValueOf(y).Complex())
				fun = func(env *Env) complex64 {
					return x(env) + y
				}

			}
		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := r.ValueOf(y).Complex()
				fun = func(env *Env) complex128 {
					return x(env) + y
				}

			}
		case r.String:
			{
				x := x.(func(*Env) string)
				y := r.ValueOf(y).String()
				fun = func(env *Env) string {
					return x(env) + y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun
		if isLiteralNumber(x, 0) || x == "" {
			return ye
		}

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x + y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x + y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x + y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x + y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x + y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x + y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x + y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x + y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x + y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x + y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x + y(env)
				}

			}
		case r.Float32:

			{
				x := float32(

					r.ValueOf(x).Float())

				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x + y(env)
				}

			}
		case r.Float64:

			{
				x := r.ValueOf(x).Float()

				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x + y(env)
				}

			}
		case r.Complex64:

			{
				x := complex64(

					r.ValueOf(x).Complex())

				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x + y(env)
				}

			}
		case r.Complex128:

			{
				x := r.ValueOf(x).Complex()

				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x + y(env)
				}

			}
		case r.String:

			{
				x := r.ValueOf(x).String()

				y := y.(func(*Env) string)
				fun = func(env *Env) string {
					return x + y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Sub(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) - y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) - y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) - y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) - y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) - y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) - y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) - y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) - y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) - y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) - y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) - y(env)
				}

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x(env) - y(env)
				}

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x(env) - y(env)
				}

			}

		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x(env) - y(env)
				}

			}

		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x(env) - y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		if isLiteralNumber(y, 0) {
			return xe
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) - y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) - y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) - y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) - y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) - y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) - y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) - y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) - y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) - y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) - y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) - y
				}

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := float32(r.ValueOf(y).Float())
				fun = func(env *Env) float32 {
					return x(env) - y
				}

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := r.ValueOf(y).Float()
				fun = func(env *Env) float64 {
					return x(env) - y
				}

			}
		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := complex64(r.ValueOf(y).Complex())
				fun = func(env *Env) complex64 {
					return x(env) - y
				}

			}
		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := r.ValueOf(y).Complex()
				fun = func(env *Env) complex128 {
					return x(env) - y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x - y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x - y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x - y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x - y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x - y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x - y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x - y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x - y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x - y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x - y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x - y(env)
				}

			}
		case r.Float32:

			{
				x := float32(

					r.ValueOf(x).Float())

				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x - y(env)
				}

			}
		case r.Float64:

			{
				x := r.ValueOf(x).Float()

				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x - y(env)
				}

			}
		case r.Complex64:

			{
				x := complex64(

					r.ValueOf(x).Complex())

				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x - y(env)
				}

			}
		case r.Complex128:

			{
				x := r.ValueOf(x).Complex()

				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x - y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Mul(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) * y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) * y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) * y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) * y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) * y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) * y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) * y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) * y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) * y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) * y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) * y(env)
				}

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x(env) * y(env)
				}

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x(env) * y(env)
				}

			}

		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x(env) * y(env)
				}

			}

		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x(env) * y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		if ze := c.mulPow2(node, xe, ye); ze != nil {
			return ze
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) * y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) * y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) * y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) * y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) * y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) * y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) * y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) * y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) * y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) * y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) * y
				}

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := float32(r.ValueOf(y).Float())
				fun = func(env *Env) float32 {
					return x(env) * y
				}

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := r.ValueOf(y).Float()
				fun = func(env *Env) float64 {
					return x(env) * y
				}

			}
		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := complex64(r.ValueOf(y).Complex())
				fun = func(env *Env) complex64 {
					return x(env) * y
				}

			}
		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := r.ValueOf(y).Complex()
				fun = func(env *Env) complex128 {
					return x(env) * y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun
		if ze := c.mulPow2(node, xe, ye); ze != nil {
			return ze
		}

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x * y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x * y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x * y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x * y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x * y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x * y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x * y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x * y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x * y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x * y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x * y(env)
				}

			}
		case r.Float32:

			{
				x := float32(

					r.ValueOf(x).Float())

				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x * y(env)
				}

			}
		case r.Float64:

			{
				x := r.ValueOf(x).Float()

				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x * y(env)
				}

			}
		case r.Complex64:

			{
				x := complex64(

					r.ValueOf(x).Complex())

				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x * y(env)
				}

			}
		case r.Complex128:

			{
				x := r.ValueOf(x).Complex()

				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x * y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Quo(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) / y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) / y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) / y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) / y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) / y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) / y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) / y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) / y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) / y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) / y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) / y(env)
				}

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x(env) / y(env)
				}

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x(env) / y(env)
				}

			}

		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x(env) / y(env)
				}

			}

		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x(env) / y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		if isLiteralNumber(y, 0) {
			c.Errorf("division by zero")
			return nil
		} else if ze := c.quoPow2(node, xe, ye); ze != nil {
			return ze
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) / y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) / y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) / y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) / y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) / y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) / y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) / y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) / y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) / y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) / y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) / y
				}

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := float32(r.ValueOf(y).Float())
				fun = func(env *Env) float32 {
					return x(env) / y
				}

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := r.ValueOf(y).Float()
				fun = func(env *Env) float64 {
					return x(env) / y
				}

			}
		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := complex64(r.ValueOf(y).Complex())
				fun = func(env *Env) complex64 {
					return x(env) / y
				}

			}
		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := r.ValueOf(y).Complex()
				fun = func(env *Env) complex128 {
					return x(env) / y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x / y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x / y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x / y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x / y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x / y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x / y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x / y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x / y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x / y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x / y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x / y(env)
				}

			}
		case r.Float32:

			{
				x := float32(

					r.ValueOf(x).Float())

				y := y.(func(*Env) float32)
				fun = func(env *Env) float32 {
					return x / y(env)
				}

			}
		case r.Float64:

			{
				x := r.ValueOf(x).Float()

				y := y.(func(*Env) float64)
				fun = func(env *Env) float64 {
					return x / y(env)
				}

			}
		case r.Complex64:

			{
				x := complex64(

					r.ValueOf(x).Complex())

				y := y.(func(*Env) complex64)
				fun = func(env *Env) complex64 {
					return x / y(env)
				}

			}
		case r.Complex128:

			{
				x := r.ValueOf(x).Complex()

				y := y.(func(*Env) complex128)
				fun = func(env *Env) complex128 {
					return x / y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Rem(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) % y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) % y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) % y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) % y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) % y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) % y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) % y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) % y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) % y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) % y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) % y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value

		if isLiteralNumber(y, 0) {
			c.Errorf("division by zero")
			return nil
		} else if ze := c.remPow2(node, xe, ye); ze != nil {
			return ze
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) % y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) % y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) % y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) % y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) % y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) % y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) % y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) % y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) % y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) % y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) % y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x % y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x % y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x % y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x % y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x % y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x % y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x % y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x % y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x % y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x % y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x % y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) mulPow2(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {

	if xe.Const() == ye.Const() {
		return nil
	}

	if xe.Const() {
		xe, ye = ye, xe
	}

	if isLiteralNumber(ye.Value, 0) {
		return c.exprZero(xe)
	} else if isLiteralNumber(ye.Value, 1) {
		return xe
	} else if isLiteralNumber(ye.Value, -1) {
		node1 := &ast.UnaryExpr{OpPos: node.OpPos, Op: token.SUB, X: node.X}
		return c.UnaryMinus(node1, xe)
	}
	ypositive := true
	yv := r.ValueOf(ye.Value)
	var y uint64
	switch base.KindToCategory(yv.Kind()) {
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
		return nil
	}
	if !isPowerOfTwo(y) {
		return nil
	}

	shift := integerLen(y) - 1
	x := xe.Fun
	var fun I
	switch xe.Type.Kind() {
	case r.Int:
		{
			x := x.(func(*Env) int)
			if ypositive {
				switch shift {
				case 1:
					fun = func(env *Env) int {
						return x(env) << 1
					}

				case 2:
					fun = func(env *Env) int {
						return x(env) << 2
					}

				case 8:
					fun = func(env *Env) int {
						return x(env) << 8
					}

				default:
					fun = func(env *Env) int {
						return x(env) << shift
					}

				}
			} else {
				fun = func(env *Env) int {
					return -(x(env) << shift)
				}
			}

		}

	case r.Int8:
		{
			x := x.(func(*Env) int8)
			if ypositive {
				switch shift {
				case 1:
					fun = func(env *Env) int8 {
						return x(env) << 1
					}

				case 2:
					fun = func(env *Env) int8 {
						return x(env) << 2
					}

				case 8:
					fun = func(env *Env) int8 {
						return x(env) << 8
					}

				default:
					fun = func(env *Env) int8 {
						return x(env) << shift
					}

				}
			} else {
				fun = func(env *Env) int8 {
					return -(x(env) << shift)
				}
			}

		}

	case r.Int16:
		{
			x := x.(func(*Env) int16)
			if ypositive {
				switch shift {
				case 1:
					fun = func(env *Env) int16 {
						return x(env) << 1
					}

				case 2:
					fun = func(env *Env) int16 {
						return x(env) << 2
					}

				case 8:
					fun = func(env *Env) int16 {
						return x(env) << 8
					}

				default:
					fun = func(env *Env) int16 {
						return x(env) << shift
					}

				}
			} else {
				fun = func(env *Env) int16 {
					return -(x(env) << shift)
				}
			}

		}

	case r.Int32:
		{
			x := x.(func(*Env) int32)
			if ypositive {
				switch shift {
				case 1:
					fun = func(env *Env) int32 {
						return x(env) << 1
					}

				case 2:
					fun = func(env *Env) int32 {
						return x(env) << 2
					}

				case 8:
					fun = func(env *Env) int32 {
						return x(env) << 8
					}

				default:
					fun = func(env *Env) int32 {
						return x(env) << shift
					}

				}
			} else {
				fun = func(env *Env) int32 {
					return -(x(env) << shift)
				}
			}

		}

	case r.Int64:
		{
			x := x.(func(*Env) int64)
			if ypositive {
				switch shift {
				case 1:
					fun = func(env *Env) int64 {
						return x(env) << 1
					}

				case 2:
					fun = func(env *Env) int64 {
						return x(env) << 2
					}

				case 8:
					fun = func(env *Env) int64 {
						return x(env) << 8
					}

				default:
					fun = func(env *Env) int64 {
						return x(env) << shift
					}

				}
			} else {
				fun = func(env *Env) int64 {
					return -(x(env) << shift)
				}
			}

		}

	case r.Uint:
		{
			x := x.(func(*Env) uint)
			switch shift {
			case 1:
				fun = func(env *Env) uint {
					return x(env) << 1
				}

			case 2:
				fun = func(env *Env) uint {
					return x(env) << 2
				}

			case 8:
				fun = func(env *Env) uint {
					return x(env) << 8
				}

			default:
				fun = func(env *Env) uint {
					return x(env) << shift
				}

			}
		}

	case r.Uint8:
		{
			x := x.(func(*Env) uint8)
			switch shift {
			case 1:
				fun = func(env *Env) uint8 {
					return x(env) << 1
				}

			case 2:
				fun = func(env *Env) uint8 {
					return x(env) << 2
				}

			case 8:
				fun = func(env *Env) uint8 {
					return x(env) << 8
				}

			default:
				fun = func(env *Env) uint8 {
					return x(env) << shift
				}

			}
		}

	case r.Uint16:
		{
			x := x.(func(*Env) uint16)
			switch shift {
			case 1:
				fun = func(env *Env) uint16 {
					return x(env) << 1
				}

			case 2:
				fun = func(env *Env) uint16 {
					return x(env) << 2
				}

			case 8:
				fun = func(env *Env) uint16 {
					return x(env) << 8
				}

			default:
				fun = func(env *Env) uint16 {
					return x(env) << shift
				}

			}
		}

	case r.Uint32:
		{
			x := x.(func(*Env) uint32)
			switch shift {
			case 1:
				fun = func(env *Env) uint32 {
					return x(env) << 1
				}

			case 2:
				fun = func(env *Env) uint32 {
					return x(env) << 2
				}

			case 8:
				fun = func(env *Env) uint32 {
					return x(env) << 8
				}

			default:
				fun = func(env *Env) uint32 {
					return x(env) << shift
				}

			}
		}

	case r.Uint64:
		{
			x := x.(func(*Env) uint64)
			switch shift {
			case 1:
				fun = func(env *Env) uint64 {
					return x(env) << 1
				}

			case 2:
				fun = func(env *Env) uint64 {
					return x(env) << 2
				}

			case 8:
				fun = func(env *Env) uint64 {
					return x(env) << 8
				}

			default:
				fun = func(env *Env) uint64 {
					return x(env) << shift
				}

			}
		}

	case r.Uintptr:
		{
			x := x.(func(*Env) uintptr)
			switch shift {
			case 1:
				fun = func(env *Env) uintptr {
					return x(env) << 1
				}

			case 2:
				fun = func(env *Env) uintptr {
					return x(env) << 2
				}

			case 8:
				fun = func(env *Env) uintptr {
					return x(env) << 8
				}

			default:
				fun = func(env *Env) uintptr {
					return x(env) << shift
				}

			}
		}

	default:
		return nil
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) quoPow2(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {

	if xe.Const() || !ye.Const() {
		return nil
	}

	if isLiteralNumber(ye.Value, 0) {
		c.Errorf("division by zero")
		return nil
	} else if isLiteralNumber(ye.Value, 1) {
		return xe
	} else if isLiteralNumber(ye.Value, -1) {
		node1 := &ast.UnaryExpr{OpPos: node.OpPos, Op: token.SUB, X: node.X}
		return c.UnaryMinus(node1, xe)
	}
	ypositive := true
	yv := r.ValueOf(ye.Value)
	var y uint64
	switch base.KindToCategory(yv.Kind()) {
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
		return nil
	}
	if !isPowerOfTwo(y) {
		return nil
	}

	shift := integerLen(y) - 1
	x := xe.Fun
	var fun I
	switch xe.Type.Kind() {
	case r.Int:
		{
			x := x.(func(*Env) int)
			y_1 :=

				int(y - 1)
			if ypositive {
				fun = func(env *Env) int {
					n := x(env)
					if n < 0 {
						n += y_1
					}
					return n >> shift
				}
			} else {
				fun = func(env *Env) int {
					n := x(env)
					if n < 0 {
						n += y_1
					}
					return -(n >> shift)
				}
			}

		}

	case r.Int8:
		{
			x := x.(func(*Env) int8)
			y_1 :=

				int8(y - 1)
			if ypositive {
				fun = func(env *Env) int8 {
					n := x(env)
					if n < 0 {
						n += y_1
					}
					return n >> shift
				}
			} else {
				fun = func(env *Env) int8 {
					n := x(env)
					if n < 0 {
						n += y_1
					}
					return -(n >> shift)
				}
			}

		}

	case r.Int16:
		{
			x := x.(func(*Env) int16)
			y_1 :=

				int16(y - 1)
			if ypositive {
				fun = func(env *Env) int16 {
					n := x(env)
					if n < 0 {
						n += y_1
					}
					return n >> shift
				}
			} else {
				fun = func(env *Env) int16 {
					n := x(env)
					if n < 0 {
						n += y_1
					}
					return -(n >> shift)
				}
			}

		}

	case r.Int32:
		{
			x := x.(func(*Env) int32)
			y_1 :=

				int32(y - 1)
			if ypositive {
				fun = func(env *Env) int32 {
					n := x(env)
					if n < 0 {
						n += y_1
					}
					return n >> shift
				}
			} else {
				fun = func(env *Env) int32 {
					n := x(env)
					if n < 0 {
						n += y_1
					}
					return -(n >> shift)
				}
			}

		}

	case r.Int64:
		{
			x := x.(func(*Env) int64)
			y_1 :=

				int64(y - 1)
			if ypositive {
				fun = func(env *Env) int64 {
					n := x(env)
					if n < 0 {
						n += y_1
					}
					return n >> shift
				}
			} else {
				fun = func(env *Env) int64 {
					n := x(env)
					if n < 0 {
						n += y_1
					}
					return -(n >> shift)
				}
			}

		}

	case r.Uint:
		{
			x := x.(func(*Env) uint)
			fun = func(env *Env) uint {
				return x(env) >> shift
			}

		}

	case r.Uint8:
		{
			x := x.(func(*Env) uint8)
			fun = func(env *Env) uint8 {
				return x(env) >> shift
			}

		}

	case r.Uint16:
		{
			x := x.(func(*Env) uint16)
			fun = func(env *Env) uint16 {
				return x(env) >> shift
			}

		}

	case r.Uint32:
		{
			x := x.(func(*Env) uint32)
			fun = func(env *Env) uint32 {
				return x(env) >> shift
			}

		}

	case r.Uint64:
		{
			x := x.(func(*Env) uint64)
			fun = func(env *Env) uint64 {
				return x(env) >> shift
			}

		}

	case r.Uintptr:
		{
			x := x.(func(*Env) uintptr)
			fun = func(env *Env) uintptr {
				return x(env) >> shift
			}

		}

	default:
		return nil
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) remPow2(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {

	if xe.Const() || !ye.Const() {
		return nil
	}

	if isLiteralNumber(ye.Value, 0) {
		c.Errorf("division by zero")
		return nil
	} else if isLiteralNumber(ye.Value, 1) {
		return c.exprZero(xe)
	}

	yv := r.ValueOf(ye.Value)
	var y uint64
	switch base.KindToCategory(yv.Kind()) {
	case r.Int:
		sy := yv.Int()
		if sy < 0 {
			y = uint64(-sy)
		} else {
			y = uint64(sy)
		}

	case r.Uint:
		y = yv.Uint()
	default:
		return nil
	}
	if !isPowerOfTwo(y) {
		return nil
	}

	x := xe.Fun
	var fun I
	switch xe.Type.Kind() {
	case r.Int:
		{
			x := x.(func(*Env) int)
			y_1 :=

				int(y - 1)
			fun = func(env *Env) int {
				n := x(env)
				if n >= 0 {
					return n & y_1
				}
				return -(-n & y_1)
			}
		}

	case r.Int8:
		{
			x := x.(func(*Env) int8)
			y_1 :=

				int8(y - 1)
			fun = func(env *Env) int8 {
				n := x(env)
				if n >= 0 {
					return n & y_1
				}
				return -(-n & y_1)
			}
		}

	case r.Int16:
		{
			x := x.(func(*Env) int16)
			y_1 :=

				int16(y - 1)
			fun = func(env *Env) int16 {
				n := x(env)
				if n >= 0 {
					return n & y_1
				}
				return -(-n & y_1)
			}
		}

	case r.Int32:
		{
			x := x.(func(*Env) int32)
			y_1 :=

				int32(y - 1)
			fun = func(env *Env) int32 {
				n := x(env)
				if n >= 0 {
					return n & y_1
				}
				return -(-n & y_1)
			}
		}

	case r.Int64:
		{
			x := x.(func(*Env) int64)
			y_1 :=

				int64(y - 1)
			fun = func(env *Env) int64 {
				n := x(env)
				if n >= 0 {
					return n & y_1
				}
				return -(-n & y_1)
			}
		}

	case r.Uint:
		{
			x := x.(func(*Env) uint)
			y_1 :=

				uint(y - 1)
			fun = func(env *Env) uint {
				return x(env) & y_1
			}

		}

	case r.Uint8:
		{
			x := x.(func(*Env) uint8)
			y_1 :=

				uint8(y - 1)
			fun = func(env *Env) uint8 {
				return x(env) & y_1
			}

		}

	case r.Uint16:
		{
			x := x.(func(*Env) uint16)
			y_1 :=

				uint16(y - 1)
			fun = func(env *Env) uint16 {
				return x(env) & y_1
			}

		}

	case r.Uint32:
		{
			x := x.(func(*Env) uint32)
			y_1 :=

				uint32(y - 1)
			fun = func(env *Env) uint32 {
				return x(env) & y_1
			}

		}

	case r.Uint64:
		{
			x := x.(func(*Env) uint64)
			y_1 :=

				uint64(y - 1)
			fun = func(env *Env) uint64 {
				return x(env) & y_1
			}

		}

	case r.Uintptr:
		{
			x := x.(func(*Env) uintptr)
			y_1 :=

				uintptr(y - 1)
			fun = func(env *Env) uintptr {
				return x(env) & y_1
			}

		}

	default:
		return nil
	}
	return exprFun(xe.Type, fun)
}
func isPowerOfTwo(n uint64) bool { return n != 0 && n&(n-1) == 0 }
func integerLen(n uint64) uint8 {
	var l uint8
	for n > 0xff {
		l += 8
		n >>= 8
	}
	for n != 0 {
		l++
		n >>= 1
	}
	return l
}
func (c *Comp) And(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) & y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) & y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) & y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) & y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) & y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) & y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) & y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) & y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) & y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) & y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) & y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		if isLiteralNumber(y, 0) {
			return c.exprZero(xe)
		} else if isLiteralNumber(y, -1) {
			return xe
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) & y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) & y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) & y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) & y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) & y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) & y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) & y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) & y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) & y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) & y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) & y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun
		if isLiteralNumber(x, 0) {
			return c.exprZero(ye)
		} else if isLiteralNumber(x, -1) {
			return ye
		}

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x & y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x & y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x & y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x & y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x & y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x & y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x & y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x & y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x & y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x & y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x & y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Or(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) | y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) | y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) | y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) | y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) | y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) | y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) | y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) | y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) | y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) | y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) | y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value

		if isLiteralNumber(y, 0) {
			return xe
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) | y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) | y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) | y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) | y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) | y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) | y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) | y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) | y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) | y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) | y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) | y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun

		if isLiteralNumber(x, 0) {
			return ye
		}

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x | y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x | y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x | y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x | y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x | y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x | y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x | y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x | y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x | y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x | y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x | y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Xor(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) ^ y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) ^ y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) ^ y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) ^ y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) ^ y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) ^ y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) ^ y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) ^ y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) ^ y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) ^ y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) ^ y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		if isLiteralNumber(y, 0) {
			return xe
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) ^ y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) ^ y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) ^ y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) ^ y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) ^ y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) ^ y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) ^ y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) ^ y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) ^ y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) ^ y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) ^ y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun
		if isLiteralNumber(x, 0) {
			return ye
		}

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x ^ y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x ^ y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x ^ y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x ^ y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x ^ y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x ^ y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x ^ y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x ^ y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x ^ y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x ^ y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x ^ y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Andnot(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x(env) &^ y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x(env) &^ y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x(env) &^ y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x(env) &^ y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x(env) &^ y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x(env) &^ y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x(env) &^ y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x(env) &^ y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x(env) &^ y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x(env) &^ y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x(env) &^ y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		if isLiteralNumber(y, -1) {
			return c.exprZero(xe)
		} else if isLiteralNumber(y, 0) {
			return xe
		}

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(r.ValueOf(y).Int())
				fun = func(env *Env) int {
					return x(env) &^ y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(r.ValueOf(y).Int())
				fun = func(env *Env) int8 {
					return x(env) &^ y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(r.ValueOf(y).Int())
				fun = func(env *Env) int16 {
					return x(env) &^ y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(r.ValueOf(y).Int())
				fun = func(env *Env) int32 {
					return x(env) &^ y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := r.ValueOf(y).Int()
				fun = func(env *Env) int64 {
					return x(env) &^ y
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(r.ValueOf(y).Uint())
				fun = func(env *Env) uint {
					return x(env) &^ y
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(r.ValueOf(y).Uint())
				fun = func(env *Env) uint8 {
					return x(env) &^ y
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(r.ValueOf(y).Uint())
				fun = func(env *Env) uint16 {
					return x(env) &^ y
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(r.ValueOf(y).Uint())
				fun = func(env *Env) uint32 {
					return x(env) &^ y
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := r.ValueOf(y).Uint()
				fun = func(env *Env) uint64 {
					return x(env) &^ y
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(r.ValueOf(y).Uint())
				fun = func(env *Env) uintptr {
					return x(env) &^ y
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun
		if isLiteralNumber(x, 0) {
			return c.exprZero(ye)
		}

		switch k {
		case r.Int:

			{
				x := int(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) int {
					return x &^ y(env)
				}

			}
		case r.Int8:

			{
				x := int8(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) int8 {
					return x &^ y(env)
				}

			}
		case r.Int16:

			{
				x := int16(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) int16 {
					return x &^ y(env)
				}

			}
		case r.Int32:

			{
				x := int32(

					r.ValueOf(x).Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) int32 {
					return x &^ y(env)
				}

			}
		case r.Int64:

			{
				x := r.ValueOf(x).Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) int64 {
					return x &^ y(env)
				}

			}
		case r.Uint:

			{
				x := uint(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) uint {
					return x &^ y(env)
				}

			}
		case r.Uint8:

			{
				x := uint8(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) uint8 {
					return x &^ y(env)
				}

			}
		case r.Uint16:

			{
				x := uint16(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) uint16 {
					return x &^ y(env)
				}

			}
		case r.Uint32:

			{
				x := uint32(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) uint32 {
					return x &^ y(env)
				}

			}
		case r.Uint64:

			{
				x := r.ValueOf(x).Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) uint64 {
					return x &^ y(env)
				}

			}
		case r.Uintptr:

			{
				x := uintptr(

					r.ValueOf(x).Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) uintptr {
					return x &^ y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) exprZero(xe *Expr) *Expr {
	if xe.Const() {
		xe.ConstTo(xe.DefaultType())
		return c.exprValue(xe.Type, xr.Zero(xe.Type).Interface())
	}
	t := xe.Type
	k := t.Kind()
	x := xe.Fun
	var fun I

	switch k {
	case r.Bool:
		{
			x := x.(func(*Env) bool)
			fun = func(env *Env) (zero bool,

			) {
				x(env)
				return

			}
		}
	case r.Int:
		{
			x := x.(func(*Env) int)
			fun = func(env *Env) (zero int,

			) {
				x(env)
				return

			}
		}
	case r.Int8:
		{
			x := x.(func(*Env) int8)
			fun = func(env *Env) (zero int8,

			) {
				x(env)
				return

			}
		}
	case r.Int16:
		{
			x := x.(func(*Env) int16)
			fun = func(env *Env) (zero int16,

			) {
				x(env)
				return

			}
		}
	case r.Int32:
		{
			x := x.(func(*Env) int32)
			fun = func(env *Env) (zero int32,

			) {
				x(env)
				return

			}
		}
	case r.Int64:
		{
			x := x.(func(*Env) int64)
			fun = func(env *Env) (zero int64,

			) {
				x(env)
				return

			}
		}
	case r.Uint:
		{
			x := x.(func(*Env) uint)
			fun = func(env *Env) (zero uint) {
				x(env)
				return

			}
		}
	case r.Uint8:
		{
			x := x.(func(*Env) uint8)
			fun = func(env *Env) (zero uint8) {
				x(env)
				return

			}
		}
	case r.Uint16:
		{
			x := x.(func(*Env) uint16)
			fun = func(env *Env) (zero uint16) {
				x(env)
				return

			}
		}
	case r.Uint32:
		{
			x := x.(func(*Env) uint32)
			fun = func(env *Env) (zero uint32) {
				x(env)
				return

			}
		}
	case r.Uint64:
		{
			x := x.(func(*Env) uint64)
			fun = func(env *Env) (zero uint64) {
				x(env)
				return

			}
		}
	case r.Uintptr:
		{
			x := x.(func(*Env) uintptr)
			fun = func(env *Env) (zero uintptr) {
				x(env)
				return

			}
		}

	case r.Float32:
		{
			x := x.(func(*Env) float32)
			fun = func(env *Env) (zero float32,

			) {
				x(env)
				return

			}
		}

	case r.Float64:
		{
			x := x.(func(*Env) float64)
			fun = func(env *Env) (zero float64,

			) {
				x(env)
				return

			}
		}

	case r.Complex64:
		{
			x := x.(func(*Env) complex64)
			fun = func(env *Env) (zero complex64,

			) {
				x(env)
				return

			}
		}

	case r.Complex128:
		{
			x := x.(func(*Env) complex128)
			fun = func(env *Env) (zero complex128,

			) {
				x(env)
				return

			}
		}

	case r.String:
		{
			x := x.(func(*Env) string)
			fun = func(env *Env) (zero string,

			) {
				x(env)
				return

			}
		}

	default:
		{
			zero := xr.Zero(t)
			x := funAsX1(x, nil)
			fun = func(env *Env) r.Value {
				x(env)
				return zero
			}
		}

	}
	return exprFun(t, fun)
}
