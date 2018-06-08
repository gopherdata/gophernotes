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
 * binary_relops.go
 *
 *  Created on Apr 12, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"
)

func (c *Comp) Lss(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun func(*Env) bool
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) bool { return x(env) < y(env) }

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) bool { return x(env) < y(env) }

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) bool { return x(env) < y(env) }

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) bool { return x(env) < y(env) }

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) bool { return x(env) < y(env) }

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) bool { return x(env) < y(env) }

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool { return x(env) < y(env) }

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool { return x(env) < y(env) }

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool { return x(env) < y(env) }

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool { return x(env) < y(env) }

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool { return x(env) < y(env) }

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) bool { return x(env) < y(env) }

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) bool { return x(env) < y(env) }

			}

		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(func(*Env) string)
				fun = func(env *Env) bool { return x(env) < y(env) }

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		yv := r.ValueOf(ye.Value)

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(yv.Int())
				fun = func(env *Env) bool { return x(env) < y }

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(yv.Int())
				fun = func(env *Env) bool { return x(env) < y }

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(yv.Int())
				fun = func(env *Env) bool { return x(env) < y }

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(yv.Int())
				fun = func(env *Env) bool { return x(env) < y }

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := yv.Int()
				fun = func(env *Env) bool { return x(env) < y }

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(yv.Uint())
				fun = func(env *Env) bool { return x(env) < y }

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(yv.Uint())
				fun = func(env *Env) bool { return x(env) < y }

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(yv.Uint())
				fun = func(env *Env) bool { return x(env) < y }

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(yv.Uint())
				fun = func(env *Env) bool { return x(env) < y }

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := yv.Uint()
				fun = func(env *Env) bool { return x(env) < y }

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(yv.Uint())
				fun = func(env *Env) bool { return x(env) < y }

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := float32(yv.Float())
				fun = func(env *Env) bool { return x(env) < y }

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := yv.Float()
				fun = func(env *Env) bool { return x(env) < y }

			}
		case r.String:
			{
				x := x.(func(*Env) string)
				y := yv.String()
				fun = func(env *Env) bool { return x(env) < y }

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		xv := r.ValueOf(xe.Value)
		y := ye.Fun

		switch k {
		case r.Int:

			{
				x := int(

					xv.Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) bool { return x < y(env) }

			}
		case r.Int8:

			{
				x := int8(

					xv.Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) bool { return x < y(env) }

			}
		case r.Int16:

			{
				x := int16(

					xv.Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) bool { return x < y(env) }

			}
		case r.Int32:

			{
				x := int32(

					xv.Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) bool { return x < y(env) }

			}
		case r.Int64:

			{
				x := xv.Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) bool { return x < y(env) }

			}
		case r.Uint:

			{
				x := uint(

					xv.Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) bool { return x < y(env) }

			}
		case r.Uint8:

			{
				x := uint8(

					xv.Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool { return x < y(env) }

			}
		case r.Uint16:

			{
				x := uint16(

					xv.Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool { return x < y(env) }

			}
		case r.Uint32:

			{
				x := uint32(

					xv.Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool { return x < y(env) }

			}
		case r.Uint64:

			{
				x := xv.Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool { return x < y(env) }

			}
		case r.Uintptr:

			{
				x := uintptr(

					xv.Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool { return x < y(env) }

			}
		case r.Float32:

			{
				x := float32(

					xv.Float())

				y := y.(func(*Env) float32)
				fun = func(env *Env) bool { return x < y(env) }

			}
		case r.Float64:

			{
				x := xv.Float()

				y := y.(func(*Env) float64)
				fun = func(env *Env) bool { return x < y(env) }

			}
		case r.String:

			{
				x := xv.String()

				y := y.(func(*Env) string)
				fun = func(env *Env) bool { return x < y(env) }

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return c.exprBool(fun)
}
func (c *Comp) Gtr(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun func(*Env) bool
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) bool { return x(env) > y(env) }

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) bool { return x(env) > y(env) }

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) bool { return x(env) > y(env) }

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) bool { return x(env) > y(env) }

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) bool { return x(env) > y(env) }

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) bool { return x(env) > y(env) }

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool { return x(env) > y(env) }

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool { return x(env) > y(env) }

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool { return x(env) > y(env) }

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool { return x(env) > y(env) }

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool { return x(env) > y(env) }

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) bool { return x(env) > y(env) }

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) bool { return x(env) > y(env) }

			}

		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(func(*Env) string)
				fun = func(env *Env) bool { return x(env) > y(env) }

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		yv := r.ValueOf(ye.Value)

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(yv.Int())
				fun = func(env *Env) bool { return x(env) > y }

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(yv.Int())
				fun = func(env *Env) bool { return x(env) > y }

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(yv.Int())
				fun = func(env *Env) bool { return x(env) > y }

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(yv.Int())
				fun = func(env *Env) bool { return x(env) > y }

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := yv.Int()
				fun = func(env *Env) bool { return x(env) > y }

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(yv.Uint())
				fun = func(env *Env) bool { return x(env) > y }

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(yv.Uint())
				fun = func(env *Env) bool { return x(env) > y }

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(yv.Uint())
				fun = func(env *Env) bool { return x(env) > y }

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(yv.Uint())
				fun = func(env *Env) bool { return x(env) > y }

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := yv.Uint()
				fun = func(env *Env) bool { return x(env) > y }

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(yv.Uint())
				fun = func(env *Env) bool { return x(env) > y }

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := float32(yv.Float())
				fun = func(env *Env) bool { return x(env) > y }

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := yv.Float()
				fun = func(env *Env) bool { return x(env) > y }

			}
		case r.String:
			{
				x := x.(func(*Env) string)
				y := yv.String()
				fun = func(env *Env) bool { return x(env) > y }

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		xv := r.ValueOf(xe.Value)
		y := ye.Fun

		switch k {
		case r.Int:

			{
				x := int(

					xv.Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) bool { return x > y(env) }

			}
		case r.Int8:

			{
				x := int8(

					xv.Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) bool { return x > y(env) }

			}
		case r.Int16:

			{
				x := int16(

					xv.Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) bool { return x > y(env) }

			}
		case r.Int32:

			{
				x := int32(

					xv.Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) bool { return x > y(env) }

			}
		case r.Int64:

			{
				x := xv.Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) bool { return x > y(env) }

			}
		case r.Uint:

			{
				x := uint(

					xv.Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) bool { return x > y(env) }

			}
		case r.Uint8:

			{
				x := uint8(

					xv.Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool { return x > y(env) }

			}
		case r.Uint16:

			{
				x := uint16(

					xv.Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool { return x > y(env) }

			}
		case r.Uint32:

			{
				x := uint32(

					xv.Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool { return x > y(env) }

			}
		case r.Uint64:

			{
				x := xv.Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool { return x > y(env) }

			}
		case r.Uintptr:

			{
				x := uintptr(

					xv.Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool { return x > y(env) }

			}
		case r.Float32:

			{
				x := float32(

					xv.Float())

				y := y.(func(*Env) float32)
				fun = func(env *Env) bool { return x > y(env) }

			}
		case r.Float64:

			{
				x := xv.Float()

				y := y.(func(*Env) float64)
				fun = func(env *Env) bool { return x > y(env) }

			}
		case r.String:

			{
				x := xv.String()

				y := y.(func(*Env) string)
				fun = func(env *Env) bool { return x > y(env) }

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return c.exprBool(fun)
}
func (c *Comp) Leq(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun func(*Env) bool
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) bool { return x(env) <= y(env) }

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) bool { return x(env) <= y(env) }

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) bool { return x(env) <= y(env) }

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) bool { return x(env) <= y(env) }

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) bool { return x(env) <= y(env) }

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) bool { return x(env) <= y(env) }

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool { return x(env) <= y(env) }

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool { return x(env) <= y(env) }

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool { return x(env) <= y(env) }

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool { return x(env) <= y(env) }

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool { return x(env) <= y(env) }

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) bool { return x(env) <= y(env) }

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) bool { return x(env) <= y(env) }

			}

		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(func(*Env) string)
				fun = func(env *Env) bool { return x(env) <= y(env) }

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		yv := r.ValueOf(ye.Value)

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(yv.Int())
				fun = func(env *Env) bool { return x(env) <= y }

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(yv.Int())
				fun = func(env *Env) bool { return x(env) <= y }

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(yv.Int())
				fun = func(env *Env) bool { return x(env) <= y }

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(yv.Int())
				fun = func(env *Env) bool { return x(env) <= y }

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := yv.Int()
				fun = func(env *Env) bool { return x(env) <= y }

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(yv.Uint())
				fun = func(env *Env) bool { return x(env) <= y }

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(yv.Uint())
				fun = func(env *Env) bool { return x(env) <= y }

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(yv.Uint())
				fun = func(env *Env) bool { return x(env) <= y }

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(yv.Uint())
				fun = func(env *Env) bool { return x(env) <= y }

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := yv.Uint()
				fun = func(env *Env) bool { return x(env) <= y }

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(yv.Uint())
				fun = func(env *Env) bool { return x(env) <= y }

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := float32(yv.Float())
				fun = func(env *Env) bool { return x(env) <= y }

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := yv.Float()
				fun = func(env *Env) bool { return x(env) <= y }

			}
		case r.String:
			{
				x := x.(func(*Env) string)
				y := yv.String()
				fun = func(env *Env) bool { return x(env) <= y }

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		xv := r.ValueOf(xe.Value)
		y := ye.Fun

		switch k {
		case r.Int:

			{
				x := int(

					xv.Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) bool { return x <= y(env) }

			}
		case r.Int8:

			{
				x := int8(

					xv.Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) bool { return x <= y(env) }

			}
		case r.Int16:

			{
				x := int16(

					xv.Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) bool { return x <= y(env) }

			}
		case r.Int32:

			{
				x := int32(

					xv.Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) bool { return x <= y(env) }

			}
		case r.Int64:

			{
				x := xv.Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) bool { return x <= y(env) }

			}
		case r.Uint:

			{
				x := uint(

					xv.Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) bool { return x <= y(env) }

			}
		case r.Uint8:

			{
				x := uint8(

					xv.Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool { return x <= y(env) }

			}
		case r.Uint16:

			{
				x := uint16(

					xv.Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool { return x <= y(env) }

			}
		case r.Uint32:

			{
				x := uint32(

					xv.Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool { return x <= y(env) }

			}
		case r.Uint64:

			{
				x := xv.Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool { return x <= y(env) }

			}
		case r.Uintptr:

			{
				x := uintptr(

					xv.Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool { return x <= y(env) }

			}
		case r.Float32:

			{
				x := float32(

					xv.Float())

				y := y.(func(*Env) float32)
				fun = func(env *Env) bool { return x <= y(env) }

			}
		case r.Float64:

			{
				x := xv.Float()

				y := y.(func(*Env) float64)
				fun = func(env *Env) bool { return x <= y(env) }

			}
		case r.String:

			{
				x := xv.String()

				y := y.(func(*Env) string)
				fun = func(env *Env) bool { return x <= y(env) }

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return c.exprBool(fun)
}
func (c *Comp) Geq(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun func(*Env) bool
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) bool { return x(env) >= y(env) }

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) bool { return x(env) >= y(env) }

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) bool { return x(env) >= y(env) }

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) bool { return x(env) >= y(env) }

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) bool { return x(env) >= y(env) }

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) bool { return x(env) >= y(env) }

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool { return x(env) >= y(env) }

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool { return x(env) >= y(env) }

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool { return x(env) >= y(env) }

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool { return x(env) >= y(env) }

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool { return x(env) >= y(env) }

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) bool { return x(env) >= y(env) }

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) bool { return x(env) >= y(env) }

			}

		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(func(*Env) string)
				fun = func(env *Env) bool { return x(env) >= y(env) }

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		yv := r.ValueOf(ye.Value)

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(yv.Int())
				fun = func(env *Env) bool { return x(env) >= y }

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(yv.Int())
				fun = func(env *Env) bool { return x(env) >= y }

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(yv.Int())
				fun = func(env *Env) bool { return x(env) >= y }

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(yv.Int())
				fun = func(env *Env) bool { return x(env) >= y }

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := yv.Int()
				fun = func(env *Env) bool { return x(env) >= y }

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(yv.Uint())
				fun = func(env *Env) bool { return x(env) >= y }

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(yv.Uint())
				fun = func(env *Env) bool { return x(env) >= y }

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(yv.Uint())
				fun = func(env *Env) bool { return x(env) >= y }

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(yv.Uint())
				fun = func(env *Env) bool { return x(env) >= y }

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := yv.Uint()
				fun = func(env *Env) bool { return x(env) >= y }

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(yv.Uint())
				fun = func(env *Env) bool { return x(env) >= y }

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := float32(yv.Float())
				fun = func(env *Env) bool { return x(env) >= y }

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := yv.Float()
				fun = func(env *Env) bool { return x(env) >= y }

			}
		case r.String:
			{
				x := x.(func(*Env) string)
				y := yv.String()
				fun = func(env *Env) bool { return x(env) >= y }

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		xv := r.ValueOf(xe.Value)
		y := ye.Fun

		switch k {
		case r.Int:

			{
				x := int(

					xv.Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) bool { return x >= y(env) }

			}
		case r.Int8:

			{
				x := int8(

					xv.Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) bool { return x >= y(env) }

			}
		case r.Int16:

			{
				x := int16(

					xv.Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) bool { return x >= y(env) }

			}
		case r.Int32:

			{
				x := int32(

					xv.Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) bool { return x >= y(env) }

			}
		case r.Int64:

			{
				x := xv.Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) bool { return x >= y(env) }

			}
		case r.Uint:

			{
				x := uint(

					xv.Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) bool { return x >= y(env) }

			}
		case r.Uint8:

			{
				x := uint8(

					xv.Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool { return x >= y(env) }

			}
		case r.Uint16:

			{
				x := uint16(

					xv.Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool { return x >= y(env) }

			}
		case r.Uint32:

			{
				x := uint32(

					xv.Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool { return x >= y(env) }

			}
		case r.Uint64:

			{
				x := xv.Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool { return x >= y(env) }

			}
		case r.Uintptr:

			{
				x := uintptr(

					xv.Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool { return x >= y(env) }

			}
		case r.Float32:

			{
				x := float32(

					xv.Float())

				y := y.(func(*Env) float32)
				fun = func(env *Env) bool { return x >= y(env) }

			}
		case r.Float64:

			{
				x := xv.Float()

				y := y.(func(*Env) float64)
				fun = func(env *Env) bool { return x >= y(env) }

			}
		case r.String:

			{
				x := xv.String()

				y := y.(func(*Env) string)
				fun = func(env *Env) bool { return x >= y(env) }

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return c.exprBool(fun)
}
