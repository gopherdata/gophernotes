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
 * binary_eql.go
 *
 *  Created on Apr 02, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/token"
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) Eql(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	if xe.IsNil() {
		if ye.IsNil() {
			return c.invalidBinaryExpr(node, xe, ye)
		} else {
			return c.eqlneqNil(node, xe, ye)
		}
	} else if ye.IsNil() {
		return c.eqlneqNil(node, xe, ye)
	}

	if !xe.Type.Comparable() || !xe.Type.Comparable() {
		return c.invalidBinaryExpr(node, xe, ye)
	}

	xc, yc := xe.Const(), ye.Const()
	if xe.Type.Kind() != r.Interface && ye.Type.Kind() != r.Interface {
		c.toSameFuncType(node, xe, ye)
	}

	k := xe.Type.Kind()
	yk := ye.Type.Kind()

	var fun func(env *Env) bool
	if k != yk {

	} else if xc == yc {
		x, y := xe.Fun, ye.Fun
		switch k {
		case r.Bool:
			{
				x := x.(func(*Env) bool)
				y := y.(func(*Env) bool)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}
		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}
		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}
		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}
		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}
		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}
		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}

		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := y.(func(*Env) complex64)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}

		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := y.(func(*Env) complex128)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}

		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(func(*Env) string)
				fun = func(env *Env) bool { return x(env) == y(env) }

			}

		}

	} else if yc {
		x := xe.Fun
		yv := r.ValueOf(ye.Value)
		if k == r.Bool && yv.Bool() {
			return xe
		}
		switch k {
		case r.Bool:

			{
				x := x.(func(*Env) bool)
				y := yv.Bool()
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(yv.Int())
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(yv.Int())
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(yv.Int())
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(yv.Int())
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := yv.Int()
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(yv.Uint())
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(yv.Uint())
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(yv.Uint())
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(yv.Uint())
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := yv.Uint()
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(yv.Uint())
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.Float32:

			{
				x := x.(func(*Env) float32)
				y :=

					float32(yv.Float())
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.Float64:

			{
				x := x.(func(*Env) float64)
				y := yv.Float()
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.Complex64:

			{
				x := x.(func(*Env) complex64)
				y :=

					complex64(yv.Complex())
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.Complex128:

			{
				x := x.(func(*Env) complex128)
				y := yv.Complex()
				fun = func(env *Env) bool { return x(env) == y }

			}
		case r.String:

			{
				x := x.(func(*Env) string)
				y := yv.String()
				fun = func(env *Env) bool { return x(env) == y }

			}
		}

	} else {
		xv := r.ValueOf(xe.Value)
		y := ye.Fun
		if k == r.Bool && xv.Bool() {
			return ye
		}
		switch k {
		case r.Bool:

			{
				x := xv.Bool()

				y := y.(func(*Env) bool)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.Int:

			{
				x := int(

					xv.Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.Int8:

			{
				x := int8(

					xv.Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.Int16:

			{
				x := int16(

					xv.Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.Int32:

			{
				x := int32(

					xv.Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.Int64:

			{
				x := xv.Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.Uint:

			{
				x := uint(

					xv.Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.Uint8:

			{
				x := uint8(

					xv.Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.Uint16:

			{
				x := uint16(

					xv.Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.Uint32:

			{
				x := uint32(

					xv.Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.Uint64:

			{
				x := xv.Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.Uintptr:

			{
				x := uintptr(

					xv.Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.Float32:

			{
				x :=

					float32(

						xv.Float())

				y := y.(func(*Env) float32)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.Float64:

			{
				x := xv.Float()

				y := y.(func(*Env) float64)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.Complex64:

			{
				x :=

					complex64(

						xv.Complex())

				y := y.(func(*Env) complex64)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.Complex128:

			{
				x := xv.Complex()

				y := y.(func(*Env) complex128)
				fun = func(env *Env) bool { return x == y(env) }

			}
		case r.String:

			{
				x := xv.String()

				y := y.(func(*Env) string)
				fun = func(env *Env) bool { return x == y(env) }

			}
		}

	}
	if fun != nil {
		return c.exprBool(fun)
	}
	return c.eqlneqMisc(node, xe, ye)
}
func (c *Comp) Neq(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	if xe.IsNil() {
		if ye.IsNil() {
			return c.invalidBinaryExpr(node, xe, ye)
		} else {
			return c.eqlneqNil(node, xe, ye)
		}
	} else if ye.IsNil() {
		return c.eqlneqNil(node, xe, ye)
	}

	if !xe.Type.Comparable() || !xe.Type.Comparable() {
		return c.invalidBinaryExpr(node, xe, ye)
	}

	xc, yc := xe.Const(), ye.Const()
	if xe.Type.Kind() != r.Interface && ye.Type.Kind() != r.Interface {
		c.toSameFuncType(node, xe, ye)
	}

	k := xe.Type.Kind()
	yk := ye.Type.Kind()

	var fun func(env *Env) bool
	if k != yk {

	} else if xc == yc {
		x, y := xe.Fun, ye.Fun
		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}
		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}
		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}
		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}
		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}
		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}
		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}

		case r.Complex64:
			{
				x := x.(func(*Env) complex64)
				y := y.(func(*Env) complex64)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}

		case r.Complex128:
			{
				x := x.(func(*Env) complex128)
				y := y.(func(*Env) complex128)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}

		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(func(*Env) string)
				fun = func(env *Env) bool { return x(env) != y(env) }

			}

		}

	} else if yc {
		x := xe.Fun
		yv := r.ValueOf(ye.Value)
		if k == r.Bool && !yv.Bool() {
			return xe
		}
		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := int(yv.Int())
				fun = func(env *Env) bool { return x(env) != y }

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := int8(yv.Int())
				fun = func(env *Env) bool { return x(env) != y }

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := int16(yv.Int())
				fun = func(env *Env) bool { return x(env) != y }

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := int32(yv.Int())
				fun = func(env *Env) bool { return x(env) != y }

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := yv.Int()
				fun = func(env *Env) bool { return x(env) != y }

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				y := uint(yv.Uint())
				fun = func(env *Env) bool { return x(env) != y }

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				y := uint8(yv.Uint())
				fun = func(env *Env) bool { return x(env) != y }

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				y := uint16(yv.Uint())
				fun = func(env *Env) bool { return x(env) != y }

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				y := uint32(yv.Uint())
				fun = func(env *Env) bool { return x(env) != y }

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				y := yv.Uint()
				fun = func(env *Env) bool { return x(env) != y }

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				y := uintptr(yv.Uint())
				fun = func(env *Env) bool { return x(env) != y }

			}
		case r.Float32:

			{
				x := x.(func(*Env) float32)
				y :=

					float32(yv.Float())
				fun = func(env *Env) bool { return x(env) != y }

			}
		case r.Float64:

			{
				x := x.(func(*Env) float64)
				y := yv.Float()
				fun = func(env *Env) bool { return x(env) != y }

			}
		case r.Complex64:

			{
				x := x.(func(*Env) complex64)
				y :=

					complex64(yv.Complex())
				fun = func(env *Env) bool { return x(env) != y }

			}
		case r.Complex128:

			{
				x := x.(func(*Env) complex128)
				y := yv.Complex()
				fun = func(env *Env) bool { return x(env) != y }

			}
		case r.String:

			{
				x := x.(func(*Env) string)
				y := yv.String()
				fun = func(env *Env) bool { return x(env) != y }

			}
		}

	} else {
		xv := r.ValueOf(xe.Value)
		y := ye.Fun
		if k == r.Bool && !xv.Bool() {
			return ye
		}
		switch k {
		case r.Int:

			{
				x := int(

					xv.Int())

				y := y.(func(*Env) int)
				fun = func(env *Env) bool { return x != y(env) }

			}
		case r.Int8:

			{
				x := int8(

					xv.Int())

				y := y.(func(*Env) int8)
				fun = func(env *Env) bool { return x != y(env) }

			}
		case r.Int16:

			{
				x := int16(

					xv.Int())

				y := y.(func(*Env) int16)
				fun = func(env *Env) bool { return x != y(env) }

			}
		case r.Int32:

			{
				x := int32(

					xv.Int())

				y := y.(func(*Env) int32)
				fun = func(env *Env) bool { return x != y(env) }

			}
		case r.Int64:

			{
				x := xv.Int()

				y := y.(func(*Env) int64)
				fun = func(env *Env) bool { return x != y(env) }

			}
		case r.Uint:

			{
				x := uint(

					xv.Uint())

				y := y.(func(*Env) uint)
				fun = func(env *Env) bool { return x != y(env) }

			}
		case r.Uint8:

			{
				x := uint8(

					xv.Uint())

				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool { return x != y(env) }

			}
		case r.Uint16:

			{
				x := uint16(

					xv.Uint())

				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool { return x != y(env) }

			}
		case r.Uint32:

			{
				x := uint32(

					xv.Uint())

				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool { return x != y(env) }

			}
		case r.Uint64:

			{
				x := xv.Uint()

				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool { return x != y(env) }

			}
		case r.Uintptr:

			{
				x := uintptr(

					xv.Uint())

				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool { return x != y(env) }

			}
		case r.Float32:

			{
				x :=

					float32(

						xv.Float())

				y := y.(func(*Env) float32)
				fun = func(env *Env) bool { return x != y(env) }

			}
		case r.Float64:

			{
				x := xv.Float()

				y := y.(func(*Env) float64)
				fun = func(env *Env) bool { return x != y(env) }

			}
		case r.Complex64:

			{
				x :=

					complex64(

						xv.Complex())

				y := y.(func(*Env) complex64)
				fun = func(env *Env) bool { return x != y(env) }

			}
		case r.Complex128:

			{
				x := xv.Complex()

				y := y.(func(*Env) complex128)
				fun = func(env *Env) bool { return x != y(env) }

			}
		case r.String:

			{
				x := xv.String()

				y := y.(func(*Env) string)
				fun = func(env *Env) bool { return x != y(env) }

			}
		}

	}
	if fun != nil {
		return c.exprBool(fun)
	}
	return c.eqlneqMisc(node, xe, ye)
}
func (c *Comp) eqlneqMisc(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	var fun func(*Env) bool

	x := xe.AsX1()
	y := ye.AsX1()
	t1 := xe.Type
	t2 := ye.Type
	extractor1 := c.extractor(t1)
	extractor2 := c.extractor(t2)

	if node.Op == token.EQL {
		fun = func(env *Env) bool {
			v1 := x(env)
			v2 := y(env)
			if v1 == Nil || v2 == Nil {
				return v1 == v2
			}

			t1, t2 := t1, t2
			if extractor1 != nil {
				v1, t1 = extractor1(v1)
			}

			if extractor2 != nil {
				v2, t2 = extractor2(v2)
			}

			if v1 == Nil || v2 == Nil {
				return v1 == v2
			}
			return v1.Interface() == v2.Interface() &&
				(t1 == nil || t2 == nil || t1.IdenticalTo(t2))
		}
	} else {
		fun = func(env *Env) bool {
			v1 := x(env)
			v2 := y(env)
			if v1 == Nil || v2 == Nil {
				return v1 != v2
			}

			t1, t2 := t1, t2
			if extractor1 != nil {
				v1, t1 = extractor1(v1)
			}

			if extractor2 != nil {
				v2, t2 = extractor2(v2)
			}

			if v1 == Nil || v2 == Nil {
				return v1 != v2
			}
			return v1.Interface() != v2.Interface() ||
				t1 != nil && t2 != nil && !t1.IdenticalTo(t2)
		}
	}
	return c.exprBool(fun)
}
func (c *Comp) eqlneqNil(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	var e *Expr
	if ye.IsNil() {
		e = xe
	} else {
		e = ye
	}

	if !IsNillableKind(e.Type.Kind()) {
		return c.invalidBinaryExpr(node, xe, ye)
	}

	var fun func(env *Env) bool
	if f, ok := e.Fun.(func(env *Env) (r.Value, []r.Value)); ok {
		e.CheckX1()
		if node.Op == token.EQL {
			fun = func(env *Env) bool {
				v, _ := f(env)
				vnil := v == Nil || IsNillableKind(v.Kind()) && v.IsNil()
				return vnil
			}
		} else {
			fun = func(env *Env) bool {
				v, _ := f(env)
				vnil := v == Nil || IsNillableKind(v.Kind()) && v.IsNil()
				return !vnil
			}
		}

	} else {
		f := e.AsX1()
		if node.Op == token.EQL {
			fun = func(env *Env) bool {
				v := f(env)
				vnil := v == Nil || IsNillableKind(v.Kind()) && v.IsNil()
				return vnil
			}
		} else {
			fun = func(env *Env) bool {
				v := f(env)
				vnil := v == Nil || IsNillableKind(v.Kind()) && v.IsNil()
				return !vnil
			}
		}

	}
	return c.exprBool(fun)
}
