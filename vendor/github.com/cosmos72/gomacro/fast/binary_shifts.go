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
 * binary_shifts.go
 *
 *  Created on Apr 08, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"
)

func (c *Comp) Shl(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	if ze := c.prepareShift(node, xe, ye); ze != nil {
		return ze
	}

	xc, yc := xe.Const(), ye.Const()
	xk, yk := xe.Type.Kind(), ye.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch xk {
		case r.Int:

			{
				x := x.(func(*Env) int)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int { return x(env) << y(env) }

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int { return x(env) << y(env) }

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int { return x(env) << y(env) }

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int { return x(env) << y(env) }

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int { return x(env) << y(env) }

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int { return x(env) << y(env) }

					}

				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int8 { return x(env) << y(env) }

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int8 { return x(env) << y(env) }

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int8 { return x(env) << y(env) }

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int8 { return x(env) << y(env) }

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int8 { return x(env) << y(env) }

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int8 { return x(env) << y(env) }

					}

				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int16 { return x(env) << y(env) }

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int16 { return x(env) << y(env) }

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int16 { return x(env) << y(env) }

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int16 { return x(env) << y(env) }

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int16 { return x(env) << y(env) }

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int16 { return x(env) << y(env) }

					}

				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int32 { return x(env) << y(env) }

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int32 { return x(env) << y(env) }

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int32 { return x(env) << y(env) }

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int32 { return x(env) << y(env) }

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int32 { return x(env) << y(env) }

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int32 { return x(env) << y(env) }

					}

				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int64 { return x(env) << y(env) }

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int64 { return x(env) << y(env) }

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int64 { return x(env) << y(env) }

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int64 { return x(env) << y(env) }

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int64 { return x(env) << y(env) }

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int64 { return x(env) << y(env) }

					}

				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint {
							return x(env) << y(env)
						}

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint {
							return x(env) << y(env)
						}

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint {
							return x(env) << y(env)
						}

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint {
							return x(env) << y(env)
						}

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint {
							return x(env) << y(env)
						}

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint {
							return x(env) << y(env)
						}

					}

				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint8 {
							return x(env) << y(env)
						}

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint8 {
							return x(env) << y(env)
						}

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint8 {
							return x(env) << y(env)
						}

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint8 {
							return x(env) << y(env)
						}

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint8 {
							return x(env) << y(env)
						}

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint8 {
							return x(env) << y(env)
						}

					}

				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint16 {
							return x(env) << y(env)
						}

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint16 {
							return x(env) << y(env)
						}

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint16 {
							return x(env) << y(env)
						}

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint16 {
							return x(env) << y(env)
						}

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint16 {
							return x(env) << y(env)
						}

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint16 {
							return x(env) << y(env)
						}

					}

				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint32 {
							return x(env) << y(env)
						}

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint32 {
							return x(env) << y(env)
						}

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint32 {
							return x(env) << y(env)
						}

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint32 {
							return x(env) << y(env)
						}

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint32 {
							return x(env) << y(env)
						}

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint32 {
							return x(env) << y(env)
						}

					}

				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint64 {
							return x(env) << y(env)
						}

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint64 {
							return x(env) << y(env)
						}

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint64 {
							return x(env) << y(env)
						}

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint64 {
							return x(env) << y(env)
						}

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint64 {
							return x(env) << y(env)
						}

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint64 {
							return x(env) << y(env)
						}

					}

				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uintptr {
							return x(env) << y(env)
						}

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uintptr {
							return x(env) << y(env)
						}

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uintptr {
							return x(env) << y(env)
						}

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uintptr {
							return x(env) << y(env)
						}

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uintptr {
							return x(env) << y(env)
						}

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uintptr {
							return x(env) << y(env)
						}

					}

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

		yv := r.ValueOf(y)

		switch xk {
		case r.Int:

			{
				x := x.(func(*Env) int)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) int {
							return x(env) << y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) int {
							return x(env) << y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) int {
							return x(env) << y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) int {
							return x(env) << y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) int {
							return x(env) << y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) int {
							return x(env) << y
						}

					}
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) int8 {
							return x(env) << y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) int8 {
							return x(env) << y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) int8 {
							return x(env) << y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) int8 {
							return x(env) << y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) int8 {
							return x(env) << y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) int8 {
							return x(env) << y
						}

					}
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) int16 {
							return x(env) << y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) int16 {
							return x(env) << y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) int16 {
							return x(env) << y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) int16 {
							return x(env) << y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) int16 {
							return x(env) << y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) int16 {
							return x(env) << y
						}

					}
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) int32 {
							return x(env) << y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) int32 {
							return x(env) << y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) int32 {
							return x(env) << y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) int32 {
							return x(env) << y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) int32 {
							return x(env) << y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) int32 {
							return x(env) << y
						}

					}
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) int64 {
							return x(env) << y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) int64 {
							return x(env) << y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) int64 {
							return x(env) << y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) int64 {
							return x(env) << y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) int64 {
							return x(env) << y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) int64 {
							return x(env) << y
						}

					}
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) uint {
							return x(env) << y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) uint {
							return x(env) << y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) uint {
							return x(env) << y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) uint {
							return x(env) << y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) uint {
							return x(env) << y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) uint {
							return x(env) << y
						}

					}
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) uint8 {
							return x(env) << y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) uint8 {
							return x(env) << y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) uint8 {
							return x(env) << y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) uint8 {
							return x(env) << y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) uint8 {
							return x(env) << y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) uint8 {
							return x(env) << y
						}

					}
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) uint16 {
							return x(env) << y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) uint16 {
							return x(env) << y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) uint16 {
							return x(env) << y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) uint16 {
							return x(env) << y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) uint16 {
							return x(env) << y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) uint16 {
							return x(env) << y
						}

					}
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) uint32 {
							return x(env) << y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) uint32 {
							return x(env) << y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) uint32 {
							return x(env) << y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) uint32 {
							return x(env) << y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) uint32 {
							return x(env) << y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) uint32 {
							return x(env) << y
						}

					}
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) uint64 {
							return x(env) << y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) uint64 {
							return x(env) << y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) uint64 {
							return x(env) << y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) uint64 {
							return x(env) << y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) uint64 {
							return x(env) << y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) uint64 {
							return x(env) << y
						}

					}
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) uintptr {
							return x(env) << y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) uintptr {
							return x(env) << y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) uintptr {
							return x(env) << y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) uintptr {
							return x(env) << y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) uintptr {
							return x(env) << y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) uintptr {
							return x(env) << y
						}

					}
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		xv := r.ValueOf(xe.Value)
		y := ye.Fun

		switch xk {
		case r.Int:

			{
				x := int(xv.Int())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int {
							return x << y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int {
							return x << y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int {
							return x << y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int {
							return x << y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int {
							return x << y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int {
							return x << y(env)
						}

					}
				}

			}
		case r.Int8:

			{
				x := int8(xv.Int())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int8 {
							return x << y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int8 {
							return x << y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int8 {
							return x << y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int8 {
							return x << y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int8 {
							return x << y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int8 {
							return x << y(env)
						}

					}
				}

			}
		case r.Int16:

			{
				x := int16(xv.Int())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int16 {
							return x << y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int16 {
							return x << y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int16 {
							return x << y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int16 {
							return x << y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int16 {
							return x << y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int16 {
							return x << y(env)
						}

					}
				}

			}
		case r.Int32:

			{
				x := int32(xv.Int())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int32 {
							return x << y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int32 {
							return x << y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int32 {
							return x << y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int32 {
							return x << y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int32 {
							return x << y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int32 {
							return x << y(env)
						}

					}
				}

			}
		case r.Int64:

			{
				x := xv.Int()
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int64 {
							return x << y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int64 {
							return x << y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int64 {
							return x << y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int64 {
							return x << y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int64 {
							return x << y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int64 {
							return x << y(env)
						}

					}
				}

			}
		case r.Uint:

			{
				x := uint(xv.Uint())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint {
							return x << y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint {
							return x << y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint {
							return x << y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint {
							return x << y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint {
							return x << y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint {
							return x << y(env)
						}

					}
				}

			}
		case r.Uint8:

			{
				x := uint8(xv.Uint())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint8 {
							return x << y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint8 {
							return x << y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint8 {
							return x << y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint8 {
							return x << y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint8 {
							return x << y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint8 {
							return x << y(env)
						}

					}
				}

			}
		case r.Uint16:

			{
				x := uint16(xv.Uint())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint16 {
							return x << y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint16 {
							return x << y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint16 {
							return x << y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint16 {
							return x << y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint16 {
							return x << y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint16 {
							return x << y(env)
						}

					}
				}

			}
		case r.Uint32:

			{
				x := uint32(xv.Uint())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint32 {
							return x << y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint32 {
							return x << y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint32 {
							return x << y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint32 {
							return x << y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint32 {
							return x << y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint32 {
							return x << y(env)
						}

					}
				}

			}
		case r.Uint64:

			{
				x := xv.Uint()
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint64 {
							return x << y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint64 {
							return x << y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint64 {
							return x << y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint64 {
							return x << y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint64 {
							return x << y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint64 {
							return x << y(env)
						}

					}
				}

			}
		case r.Uintptr:

			{
				x := uintptr(xv.Uint())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uintptr {
							return x << y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uintptr {
							return x << y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uintptr {
							return x << y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uintptr {
							return x << y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uintptr {
							return x << y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uintptr {
							return x << y(env)
						}

					}
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Shr(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	if ze := c.prepareShift(node, xe, ye); ze != nil {
		return ze
	}

	xc, yc := xe.Const(), ye.Const()
	xk, yk := xe.Type.Kind(), ye.Type.Kind()

	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch xk {
		case r.Int:

			{
				x := x.(func(*Env) int)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int { return x(env) >> y(env) }

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int { return x(env) >> y(env) }

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int { return x(env) >> y(env) }

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int { return x(env) >> y(env) }

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int { return x(env) >> y(env) }

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int { return x(env) >> y(env) }

					}

				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int8 { return x(env) >> y(env) }

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int8 { return x(env) >> y(env) }

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int8 { return x(env) >> y(env) }

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int8 { return x(env) >> y(env) }

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int8 { return x(env) >> y(env) }

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int8 { return x(env) >> y(env) }

					}

				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int16 { return x(env) >> y(env) }

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int16 { return x(env) >> y(env) }

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int16 { return x(env) >> y(env) }

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int16 { return x(env) >> y(env) }

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int16 { return x(env) >> y(env) }

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int16 { return x(env) >> y(env) }

					}

				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int32 { return x(env) >> y(env) }

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int32 { return x(env) >> y(env) }

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int32 { return x(env) >> y(env) }

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int32 { return x(env) >> y(env) }

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int32 { return x(env) >> y(env) }

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int32 { return x(env) >> y(env) }

					}

				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int64 { return x(env) >> y(env) }

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int64 { return x(env) >> y(env) }

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int64 { return x(env) >> y(env) }

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int64 { return x(env) >> y(env) }

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int64 { return x(env) >> y(env) }

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int64 { return x(env) >> y(env) }

					}

				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint {
							return x(env) >> y(env)
						}

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint {
							return x(env) >> y(env)
						}

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint {
							return x(env) >> y(env)
						}

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint {
							return x(env) >> y(env)
						}

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint {
							return x(env) >> y(env)
						}

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint {
							return x(env) >> y(env)
						}

					}

				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint8 {
							return x(env) >> y(env)
						}

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint8 {
							return x(env) >> y(env)
						}

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint8 {
							return x(env) >> y(env)
						}

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint8 {
							return x(env) >> y(env)
						}

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint8 {
							return x(env) >> y(env)
						}

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint8 {
							return x(env) >> y(env)
						}

					}

				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint16 {
							return x(env) >> y(env)
						}

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint16 {
							return x(env) >> y(env)
						}

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint16 {
							return x(env) >> y(env)
						}

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint16 {
							return x(env) >> y(env)
						}

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint16 {
							return x(env) >> y(env)
						}

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint16 {
							return x(env) >> y(env)
						}

					}

				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint32 {
							return x(env) >> y(env)
						}

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint32 {
							return x(env) >> y(env)
						}

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint32 {
							return x(env) >> y(env)
						}

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint32 {
							return x(env) >> y(env)
						}

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint32 {
							return x(env) >> y(env)
						}

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint32 {
							return x(env) >> y(env)
						}

					}

				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint64 {
							return x(env) >> y(env)
						}

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint64 {
							return x(env) >> y(env)
						}

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint64 {
							return x(env) >> y(env)
						}

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint64 {
							return x(env) >> y(env)
						}

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint64 {
							return x(env) >> y(env)
						}

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint64 {
							return x(env) >> y(env)
						}

					}

				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				switch yk {
				case r.Uint:
					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uintptr {
							return x(env) >> y(env)
						}

					}

				case r.Uint8:
					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uintptr {
							return x(env) >> y(env)
						}

					}

				case r.Uint16:
					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uintptr {
							return x(env) >> y(env)
						}

					}

				case r.Uint32:
					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uintptr {
							return x(env) >> y(env)
						}

					}

				case r.Uint64:
					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uintptr {
							return x(env) >> y(env)
						}

					}

				case r.Uintptr:
					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uintptr {
							return x(env) >> y(env)
						}

					}

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

		yv := r.ValueOf(y)

		switch xk {
		case r.Int:

			{
				x := x.(func(*Env) int)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) int {
							return x(env) >> y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) int {
							return x(env) >> y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) int {
							return x(env) >> y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) int {
							return x(env) >> y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) int {
							return x(env) >> y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) int {
							return x(env) >> y
						}

					}
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) int8 {
							return x(env) >> y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) int8 {
							return x(env) >> y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) int8 {
							return x(env) >> y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) int8 {
							return x(env) >> y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) int8 {
							return x(env) >> y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) int8 {
							return x(env) >> y
						}

					}
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) int16 {
							return x(env) >> y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) int16 {
							return x(env) >> y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) int16 {
							return x(env) >> y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) int16 {
							return x(env) >> y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) int16 {
							return x(env) >> y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) int16 {
							return x(env) >> y
						}

					}
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) int32 {
							return x(env) >> y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) int32 {
							return x(env) >> y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) int32 {
							return x(env) >> y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) int32 {
							return x(env) >> y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) int32 {
							return x(env) >> y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) int32 {
							return x(env) >> y
						}

					}
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) int64 {
							return x(env) >> y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) int64 {
							return x(env) >> y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) int64 {
							return x(env) >> y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) int64 {
							return x(env) >> y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) int64 {
							return x(env) >> y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) int64 {
							return x(env) >> y
						}

					}
				}

			}
		case r.Uint:

			{
				x := x.(func(*Env) uint)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) uint {
							return x(env) >> y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) uint {
							return x(env) >> y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) uint {
							return x(env) >> y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) uint {
							return x(env) >> y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) uint {
							return x(env) >> y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) uint {
							return x(env) >> y
						}

					}
				}

			}
		case r.Uint8:

			{
				x := x.(func(*Env) uint8)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) uint8 {
							return x(env) >> y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) uint8 {
							return x(env) >> y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) uint8 {
							return x(env) >> y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) uint8 {
							return x(env) >> y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) uint8 {
							return x(env) >> y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) uint8 {
							return x(env) >> y
						}

					}
				}

			}
		case r.Uint16:

			{
				x := x.(func(*Env) uint16)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) uint16 {
							return x(env) >> y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) uint16 {
							return x(env) >> y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) uint16 {
							return x(env) >> y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) uint16 {
							return x(env) >> y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) uint16 {
							return x(env) >> y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) uint16 {
							return x(env) >> y
						}

					}
				}

			}
		case r.Uint32:

			{
				x := x.(func(*Env) uint32)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) uint32 {
							return x(env) >> y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) uint32 {
							return x(env) >> y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) uint32 {
							return x(env) >> y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) uint32 {
							return x(env) >> y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) uint32 {
							return x(env) >> y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) uint32 {
							return x(env) >> y
						}

					}
				}

			}
		case r.Uint64:

			{
				x := x.(func(*Env) uint64)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) uint64 {
							return x(env) >> y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) uint64 {
							return x(env) >> y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) uint64 {
							return x(env) >> y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) uint64 {
							return x(env) >> y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) uint64 {
							return x(env) >> y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) uint64 {
							return x(env) >> y
						}

					}
				}

			}
		case r.Uintptr:

			{
				x := x.(func(*Env) uintptr)
				switch yk {
				case r.Uint:

					{
						y := uint(yv.Uint())
						fun = func(env *Env) uintptr {
							return x(env) >> y
						}

					}
				case r.Uint8:

					{
						y := uint8(yv.Uint())
						fun = func(env *Env) uintptr {
							return x(env) >> y
						}

					}
				case r.Uint16:

					{
						y := uint16(yv.Uint())
						fun = func(env *Env) uintptr {
							return x(env) >> y
						}

					}
				case r.Uint32:

					{
						y := uint32(yv.Uint())
						fun = func(env *Env) uintptr {
							return x(env) >> y
						}

					}
				case r.Uint64:

					{
						y := yv.Uint()

						fun = func(env *Env) uintptr {
							return x(env) >> y
						}

					}
				case r.Uintptr:

					{
						y := uintptr(yv.Uint())
						fun = func(env *Env) uintptr {
							return x(env) >> y
						}

					}
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		xv := r.ValueOf(xe.Value)
		y := ye.Fun

		switch xk {
		case r.Int:

			{
				x := int(xv.Int())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int {
							return x >> y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int {
							return x >> y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int {
							return x >> y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int {
							return x >> y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int {
							return x >> y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int {
							return x >> y(env)
						}

					}
				}

			}
		case r.Int8:

			{
				x := int8(xv.Int())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int8 {
							return x >> y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int8 {
							return x >> y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int8 {
							return x >> y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int8 {
							return x >> y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int8 {
							return x >> y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int8 {
							return x >> y(env)
						}

					}
				}

			}
		case r.Int16:

			{
				x := int16(xv.Int())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int16 {
							return x >> y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int16 {
							return x >> y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int16 {
							return x >> y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int16 {
							return x >> y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int16 {
							return x >> y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int16 {
							return x >> y(env)
						}

					}
				}

			}
		case r.Int32:

			{
				x := int32(xv.Int())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int32 {
							return x >> y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int32 {
							return x >> y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int32 {
							return x >> y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int32 {
							return x >> y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int32 {
							return x >> y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int32 {
							return x >> y(env)
						}

					}
				}

			}
		case r.Int64:

			{
				x := xv.Int()
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) int64 {
							return x >> y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) int64 {
							return x >> y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) int64 {
							return x >> y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) int64 {
							return x >> y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) int64 {
							return x >> y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) int64 {
							return x >> y(env)
						}

					}
				}

			}
		case r.Uint:

			{
				x := uint(xv.Uint())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint {
							return x >> y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint {
							return x >> y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint {
							return x >> y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint {
							return x >> y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint {
							return x >> y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint {
							return x >> y(env)
						}

					}
				}

			}
		case r.Uint8:

			{
				x := uint8(xv.Uint())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint8 {
							return x >> y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint8 {
							return x >> y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint8 {
							return x >> y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint8 {
							return x >> y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint8 {
							return x >> y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint8 {
							return x >> y(env)
						}

					}
				}

			}
		case r.Uint16:

			{
				x := uint16(xv.Uint())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint16 {
							return x >> y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint16 {
							return x >> y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint16 {
							return x >> y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint16 {
							return x >> y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint16 {
							return x >> y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint16 {
							return x >> y(env)
						}

					}
				}

			}
		case r.Uint32:

			{
				x := uint32(xv.Uint())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint32 {
							return x >> y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint32 {
							return x >> y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint32 {
							return x >> y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint32 {
							return x >> y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint32 {
							return x >> y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint32 {
							return x >> y(env)
						}

					}
				}

			}
		case r.Uint64:

			{
				x := xv.Uint()
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uint64 {
							return x >> y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uint64 {
							return x >> y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uint64 {
							return x >> y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uint64 {
							return x >> y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uint64 {
							return x >> y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uint64 {
							return x >> y(env)
						}

					}
				}

			}
		case r.Uintptr:

			{
				x := uintptr(xv.Uint())
				switch yk {
				case r.Uint:

					{
						y := y.(func(*Env) uint)
						fun = func(env *Env) uintptr {
							return x >> y(env)
						}

					}
				case r.Uint8:

					{
						y := y.(func(*Env) uint8)
						fun = func(env *Env) uintptr {
							return x >> y(env)
						}

					}
				case r.Uint16:

					{
						y := y.(func(*Env) uint16)
						fun = func(env *Env) uintptr {
							return x >> y(env)
						}

					}
				case r.Uint32:

					{
						y := y.(func(*Env) uint32)
						fun = func(env *Env) uintptr {
							return x >> y(env)
						}

					}
				case r.Uint64:

					{
						y := y.(func(*Env) uint64)
						fun = func(env *Env) uintptr {
							return x >> y(env)
						}

					}
				case r.Uintptr:

					{
						y := y.(func(*Env) uintptr)
						fun = func(env *Env) uintptr {
							return x >> y(env)
						}

					}
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
