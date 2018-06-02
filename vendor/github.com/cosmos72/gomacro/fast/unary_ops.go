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
 * unary_plus.go
 *
 *  Created on Apr 07, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) UnaryPlus(node *ast.UnaryExpr, xe *Expr) *Expr {
	if !IsCategory(xe.Type.Kind(), r.Int, r.Uint, r.Float64, r.Complex128) {
		return c.invalidUnaryExpr(node, xe)
	}
	return xe
}

func (c *Comp) UnaryMinus(node *ast.UnaryExpr, xe *Expr) *Expr {
	// if xe is constant, UnaryExpr will invoke EvalConst()
	// on our return value. no need to optimize that.
	x := xe.Fun
	var fun I
	switch x := x.(type) {
	case func(env *Env) int:
		fun = func(env *Env) int {
			return -x(env)
		}
	case func(env *Env) int8:
		fun = func(env *Env) int8 {
			return -x(env)
		}
	case func(env *Env) int16:
		fun = func(env *Env) int16 {
			return -x(env)
		}
	case func(env *Env) int32:
		fun = func(env *Env) int32 {
			return -x(env)
		}
	case func(env *Env) int64:
		fun = func(env *Env) int64 {
			return -x(env)
		}
	case func(env *Env) uint:
		fun = func(env *Env) uint {
			return -x(env)
		}
	case func(env *Env) uint8:
		fun = func(env *Env) uint8 {
			return -x(env)
		}
	case func(env *Env) uint16:
		fun = func(env *Env) uint16 {
			return -x(env)
		}
	case func(env *Env) uint32:
		fun = func(env *Env) uint32 {
			return -x(env)
		}
	case func(env *Env) uint64:
		fun = func(env *Env) uint64 {
			return -x(env)
		}
	case func(env *Env) uintptr:
		fun = func(env *Env) uintptr {
			return -x(env)
		}
	case func(env *Env) float32:
		fun = func(env *Env) float32 {
			return -x(env)
		}
	case func(env *Env) float64:
		fun = func(env *Env) float64 {
			return -x(env)
		}
	case func(env *Env) complex64:
		fun = func(env *Env) complex64 {
			return -x(env)
		}
	case func(env *Env) complex128:
		fun = func(env *Env) complex128 {
			return -x(env)
		}
	default:
		return c.invalidUnaryExpr(node, xe)
	}
	return exprFun(xe.Type, fun)
}

func (c *Comp) UnaryXor(node *ast.UnaryExpr, xe *Expr) *Expr {
	// if xe is constant, UnaryExpr will invoke EvalConst()
	// on our return value. no need to optimize that.
	x := xe.Fun
	var fun I
	switch x := x.(type) {
	case func(env *Env) int:
		fun = func(env *Env) int {
			return ^x(env)
		}
	case func(env *Env) int8:
		fun = func(env *Env) int8 {
			return ^x(env)
		}
	case func(env *Env) int16:
		fun = func(env *Env) int16 {
			return ^x(env)
		}
	case func(env *Env) int32:
		fun = func(env *Env) int32 {
			return ^x(env)
		}
	case func(env *Env) int64:
		fun = func(env *Env) int64 {
			return ^x(env)
		}
	case func(env *Env) uint:
		fun = func(env *Env) uint {
			return ^x(env)
		}
	case func(env *Env) uint8:
		fun = func(env *Env) uint8 {
			return ^x(env)
		}
	case func(env *Env) uint16:
		fun = func(env *Env) uint16 {
			return ^x(env)
		}
	case func(env *Env) uint32:
		fun = func(env *Env) uint32 {
			return ^x(env)
		}
	case func(env *Env) uint64:
		fun = func(env *Env) uint64 {
			return ^x(env)
		}
	case func(env *Env) uintptr:
		fun = func(env *Env) uintptr {
			return ^x(env)
		}
	default:
		return c.invalidUnaryExpr(node, xe)
	}
	return exprFun(xe.Type, fun)
}

func (c *Comp) UnaryNot(node *ast.UnaryExpr, xe *Expr) *Expr {
	// if xe is constant, UnaryExpr will invoke EvalConst()
	// on our return value. no need to optimize that.
	x := xe.Fun
	var fun I
	switch x := x.(type) {
	case func(env *Env) bool:
		fun = func(env *Env) bool {
			return !x(env)
		}
	default:
		return c.invalidUnaryExpr(node, xe)
	}
	return exprFun(xe.Type, fun)
}

// StarExpr compiles unary operator * i.e. pointer dereference
func (c *Comp) StarExpr(node *ast.StarExpr) *Expr {
	expr := node.X
	for {
		switch e := expr.(type) {
		case *ast.ParenExpr:
			expr = e.X
			continue
		case *ast.UnaryExpr:
			if e.Op == token.AND {
				// optimize * & x -> x, but check that x is addressable
				c.placeOrAddress(e.X, PlaceAddress, nil)
				return c.Expr1(e.X, nil)
			}
		}
		break
	}
	addr := c.Expr1(expr, nil) // panics if addr returns zero values, warns if returns multiple values
	taddr := addr.Type
	if taddr.Kind() != r.Ptr {
		c.Errorf("unary operation * on non-pointer <%v>: %v", taddr, node)
	}
	return c.Deref(addr)
}

// Deref compiles unary operator * i.e. pointer dereference
func (c *Comp) Deref(addr *Expr) *Expr {
	taddr := addr.Type
	if taddr.Kind() != r.Ptr {
		c.Errorf("unary operation * on non-pointer <%v>", taddr)
	}
	x1 := addr.AsX1() // panics if addr returns zero values, warns if returns multiple values
	t := taddr.Elem()
	x := addr.Fun
	var fun I
	// fast interpreter expects that Exprs returning primitive types or string
	// do NOT wrap them into reflect.Value
	switch x := x.(type) {
	case func(env *Env) *bool:
		fun = func(env *Env) bool {
			return *x(env)
		}
	case func(env *Env) *int:
		fun = func(env *Env) int {
			return *x(env)
		}
	case func(env *Env) *int8:
		fun = func(env *Env) int8 {
			return *x(env)
		}
	case func(env *Env) *int16:
		fun = func(env *Env) int16 {
			return *x(env)
		}
	case func(env *Env) *int32:
		fun = func(env *Env) int32 {
			return *x(env)
		}
	case func(env *Env) *int64:
		fun = func(env *Env) int64 {
			return *x(env)
		}
	case func(env *Env) *uint:
		fun = func(env *Env) uint {
			return *x(env)
		}
	case func(env *Env) *uint8:
		fun = func(env *Env) uint8 {
			return *x(env)
		}
	case func(env *Env) *uint16:
		fun = func(env *Env) uint16 {
			return *x(env)
		}
	case func(env *Env) *uint32:
		fun = func(env *Env) uint32 {
			return *x(env)
		}
	case func(env *Env) *uint64:
		fun = func(env *Env) uint64 {
			return *x(env)
		}
	case func(env *Env) *uintptr:
		fun = func(env *Env) uintptr {
			return *x(env)
		}
	case func(env *Env) *float32:
		fun = func(env *Env) float32 {
			return *x(env)
		}
	case func(env *Env) *float64:
		fun = func(env *Env) float64 {
			return *x(env)
		}
	case func(env *Env) *complex64:
		fun = func(env *Env) complex64 {
			return *x(env)
		}
	default:
		fun = c.derefUnwrap(t, x1)
	}
	return exprFun(t, fun)
}

// deref0Unwrap compiles unary operator * on reflect.Value - unwraps reflect.Value.Elem() if possible
func (c *Comp) derefUnwrap(t xr.Type, x1 func(*Env) r.Value) I {
	var fun I
	switch t.Kind() {
	case r.Bool:
		fun = func(env *Env) bool {
			return x1(env).Elem().Bool()
		}
	case r.Int:
		fun = func(env *Env) int {
			return int(x1(env).Elem().Int())
		}
	case r.Int8:
		fun = func(env *Env) int8 {
			return int8(x1(env).Elem().Int())
		}
	case r.Int16:
		fun = func(env *Env) int16 {
			return int16(x1(env).Elem().Int())
		}
	case r.Int32:
		fun = func(env *Env) int32 {
			return int32(x1(env).Elem().Int())
		}
	case r.Int64:
		fun = func(env *Env) int64 {
			return x1(env).Elem().Int()
		}
	case r.Uint:
		fun = func(env *Env) uint {
			return uint(x1(env).Elem().Uint())
		}
	case r.Uint8:
		fun = func(env *Env) uint8 {
			return uint8(x1(env).Elem().Uint())
		}
	case r.Uint16:
		fun = func(env *Env) uint16 {
			return uint16(x1(env).Elem().Uint())
		}
	case r.Uint32:
		fun = func(env *Env) uint32 {
			return uint32(x1(env).Elem().Uint())
		}
	case r.Uint64:
		fun = func(env *Env) uint64 {
			return x1(env).Elem().Uint()
		}
	case r.Uintptr:
		fun = func(env *Env) uintptr {
			return uintptr(x1(env).Elem().Uint())
		}
	case r.Float32:
		fun = func(env *Env) float32 {
			return float32(x1(env).Elem().Float())
		}
	case r.Float64:
		fun = func(env *Env) float64 {
			return x1(env).Elem().Float()
		}
	case r.Complex64:
		fun = func(env *Env) complex64 {
			return complex64(x1(env).Elem().Complex())
		}
	case r.Complex128:
		fun = func(env *Env) complex128 {
			return x1(env).Elem().Complex()
		}
	case r.String:
		fun = func(env *Env) string {
			return x1(env).Elem().String()
		}
	default:
		fun = func(env *Env) r.Value {
			return x1(env).Elem()
		}
	}
	return fun
}
