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
	if xe.IsNil {
		if ye.IsNil {
			return c.invalidBinaryExpr(node, xe, ye)
		} else {
			return c.eqlneqNil(node, xe, ye)
		}
	} else if ye.IsNil {
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
	if xe.IsNil {
		if ye.IsNil {
			return c.invalidBinaryExpr(node, xe, ye)
		} else {
			return c.eqlneqNil(node, xe, ye)
		}
	} else if ye.IsNil {
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

	if xe.Type.Kind() == r.Interface || ye.Type.Kind() == r.Interface {

		xe.CheckX1()
		ye.CheckX1()
	}

	switch x := xe.Fun.(type) {
	case func(*Env) (r.Value, []r.Value):
		switch y := ye.Fun.(type) {
		case func(*Env) (r.Value, []r.Value):
			if node.Op == token.EQL {
				fun = func(env *Env) bool {
					v1, _ := x(env)
					v2, _ := y(env)
					if v1 == Nil || v2 == Nil {
						return v1 == v2
					} else {
						return v1.Interface() == v2.Interface()
					}

				}
			} else {
				fun = func(env *Env) bool {
					v1, _ := x(env)
					v2, _ := y(env)
					if v1 == Nil || v2 == Nil {
						return v1 != v2
					} else {
						return v1.Interface() != v2.Interface()
					}

				}
			}

		default:
			y1 := ye.AsX1()
			if node.Op == token.EQL {
				fun = func(env *Env) bool {
					v1, _ := x(env)
					v2 := y1(env)
					if v1 == Nil || v2 == Nil {
						return v1 == v2
					} else {
						return v1.Interface() == v2.Interface()
					}

				}
			} else {
				fun = func(env *Env) bool {
					v1, _ := x(env)
					v2 := y1(env)
					if v1 == Nil || v2 == Nil {
						return v1 != v2
					} else {
						return v1.Interface() != v2.Interface()
					}

				}
			}

		}
	default:
		x1 := xe.AsX1()

		switch y := ye.Fun.(type) {
		case func(*Env) (r.Value, []r.Value):
			if node.Op == token.EQL {
				fun = func(env *Env) bool {
					v1 := x1(env)
					v2, _ := y(env)
					if v1 == Nil || v2 == Nil {
						return v1 == v2
					} else {
						return v1.Interface() == v2.Interface()
					}

				}
			} else {
				fun = func(env *Env) bool {
					v1 := x1(env)
					v2, _ := y(env)
					if v1 == Nil || v2 == Nil {
						return v1 != v2
					} else {
						return v1.Interface() != v2.Interface()
					}

				}
			}

		default:
			y1 := ye.AsX1()
			if node.Op == token.EQL {
				fun = func(env *Env) bool {
					v1 := x1(env)
					v2 := y1(env)
					if v1 == Nil || v2 == Nil {
						return v1 == v2
					} else {
						return v1.Interface() == v2.Interface()
					}

				}
			} else {
				fun = func(env *Env) bool {
					v1 := x1(env)
					v2 := y1(env)
					if v1 == Nil || v2 == Nil {
						return v1 != v2
					} else {
						return v1.Interface() != v2.Interface()
					}

				}
			}

		}
	}
	return c.exprBool(fun)
}
func (c *Comp) eqlneqNil(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	var e *Expr
	if ye.IsNil {
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
