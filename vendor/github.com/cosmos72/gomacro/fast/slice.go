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
 * slice.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"
)

// SliceExpr compiles slice[lo:hi] and slice[lo:hi:max]
func (c *Comp) SliceExpr(node *ast.SliceExpr) *Expr {
	e := c.Expr1(node.X, nil)
	if e.Const() {
		e.ConstTo(e.DefaultType())
	}
	if e.Type.Kind() == r.Array {
		c.sliceArrayMustBeAddressable(node, e)
	}
	lo := c.sliceIndex(node.Low)
	hi := c.sliceIndex(node.High)
	max := c.sliceIndex(node.Max)
	var ret *Expr
	if node.Slice3 {
		ret = c.slice3(node, e, lo, hi, max)
	} else {
		ret = c.slice2(node, e, lo, hi)
	}
	// constant propagation
	if e.Const() && (lo == nil || lo.Const()) && (hi == nil || hi.Const()) && (max == nil || max.Const()) {
		ret.EvalConst(COptKeepUntyped)
	}
	return ret
}

func (c *Comp) sliceIndex(node ast.Expr) *Expr {
	if node == nil {
		return nil
	}
	idx := c.Expr1(node, nil)
	if idx.Const() {
		idx.ConstTo(c.TypeOfInt())
		if idx.Value.(int) < 0 {
			c.Errorf("negative slice index: %v == %v", node, idx)
		}
	} else if idx.Type == nil || !idx.Type.AssignableTo(c.TypeOfInt()) {
		c.Errorf("invalid slice index: expecting integer, found: %v <%v>", idx.Type, node)
	}
	return idx
}

// slice2 compiles slice[lo:hi]
func (c *Comp) slice2(node *ast.SliceExpr, e, lo, hi *Expr) *Expr {
	t := e.Type
	switch t.Kind() {
	case r.String:
		return c.sliceString(e, lo, hi)
	case r.Ptr:
		if t.Elem().Kind() != r.Array {
			break
		}
		fallthrough
	case r.Slice, r.Array:
		if t.Kind() == r.Ptr {
			t = t.Elem()
			objfun := e.AsX1()
			e = exprX1(t, func(env *Env) r.Value {
				return objfun(env).Elem()
			})
		}
		objfun := e.AsX1()
		if lo == nil {
			lo = c.exprValue(c.TypeOfInt(), 0)
		}
		var fun func(env *Env) r.Value
		if lo.Const() {
			lo := lo.Value.(int)
			if hi == nil {
				fun = func(env *Env) r.Value {
					obj := objfun(env)
					return obj.Slice(lo, obj.Len())
				}
			} else if hi.Const() {
				hi := hi.Value.(int)
				fun = func(env *Env) r.Value {
					obj := objfun(env)
					return obj.Slice(lo, hi)
				}
			} else {
				hifun := hi.WithFun().(func(*Env) int)
				fun = func(env *Env) r.Value {
					obj := objfun(env)
					hi := hifun(env)
					return obj.Slice(lo, hi)
				}
			}
		} else {
			lofun := lo.WithFun().(func(*Env) int)
			if hi == nil {
				fun = func(env *Env) r.Value {
					obj := objfun(env)
					lo := lofun(env)
					return obj.Slice(lo, obj.Len())
				}
			} else if hi.Const() {
				hi := hi.Value.(int)
				fun = func(env *Env) r.Value {
					obj := objfun(env)
					lo := lofun(env)
					return obj.Slice(lo, hi)
				}
			} else {
				hifun := hi.WithFun().(func(*Env) int)
				fun = func(env *Env) r.Value {
					obj := objfun(env)
					lo := lofun(env)
					hi := hifun(env)
					return obj.Slice(lo, hi)
				}
			}
		}
		tout := c.Universe.SliceOf(t.Elem())
		return exprX1(tout, fun)
	}
	c.Errorf("cannot slice %v: %v", t, node)
	return nil
}

// sliceString compiles string[lo:hi]
func (c *Comp) sliceString(e, lo, hi *Expr) *Expr {
	objfun := e.WithFun().(func(*Env) string)
	var fun func(env *Env) string
	if lo == nil {
		if hi == nil {
			fun = objfun
		} else if hi.Const() {
			hi := hi.Value.(int)
			fun = func(env *Env) string {
				obj := objfun(env)
				return obj[:hi]
			}
		} else {
			hifun := hi.WithFun().(func(*Env) int)
			fun = func(env *Env) string {
				obj := objfun(env)
				hi := hifun(env)
				return obj[:hi]
			}
		}
	} else if lo.Const() {
		lo := lo.Value.(int)
		if hi == nil {
			fun = func(env *Env) string {
				obj := objfun(env)
				return obj[lo:]
			}
		} else if hi.Const() {
			hi := hi.Value.(int)
			fun = func(env *Env) string {
				obj := objfun(env)
				return obj[lo:hi]
			}
		} else {
			hifun := hi.WithFun().(func(*Env) int)
			fun = func(env *Env) string {
				obj := objfun(env)
				hi := hifun(env)
				return obj[lo:hi]
			}
		}
	} else {
		lofun := lo.WithFun().(func(*Env) int)
		if hi == nil {
			fun = func(env *Env) string {
				obj := objfun(env)
				lo := lofun(env)
				return obj[lo:]
			}
		} else if hi.Const() {
			hi := hi.Value.(int)
			fun = func(env *Env) string {
				obj := objfun(env)
				lo := lofun(env)
				return obj[lo:hi]
			}
		} else {
			hifun := hi.WithFun().(func(*Env) int)
			fun = func(env *Env) string {
				obj := objfun(env)
				lo := lofun(env)
				hi := hifun(env)
				return obj[lo:hi]
			}
		}
	}
	return exprFun(c.TypeOfString(), fun)
}

// slice3 compiles slice[lo:hi:max]
func (c *Comp) slice3(node *ast.SliceExpr, e, lo, hi, max *Expr) *Expr {
	if lo == nil {
		lo = c.exprValue(c.TypeOfInt(), 0)
	}
	if hi == nil {
		c.Errorf("final index required in 3-index slice: %v", node)
	}
	if max == nil {
		c.Errorf("final index required in 3-index slice: %v", node)
	}
	t := e.Type
	switch t.Kind() {
	case r.String:
		c.Errorf("invalid operation %v (3-index slice of string)", node)
		return nil
	case r.Ptr:
		if t.Elem().Kind() != r.Array {
			break
		}
		fallthrough
	case r.Slice, r.Array:
		objfun := e.AsX1()
		lofun := lo.WithFun().(func(*Env) int)
		hifun := hi.WithFun().(func(*Env) int)
		maxfun := max.WithFun().(func(*Env) int)
		var fun func(env *Env) r.Value
		if t.Kind() == r.Ptr {
			t = t.Elem()
			fun = func(env *Env) r.Value {
				obj := objfun(env).Elem()
				lo := lofun(env)
				hi := hifun(env)
				max := maxfun(env)
				return obj.Slice3(lo, hi, max)
			}
		} else {
			fun = func(env *Env) r.Value {
				obj := objfun(env)
				lo := lofun(env)
				hi := hifun(env)
				max := maxfun(env)
				return obj.Slice3(lo, hi, max)
			}
		}
		tout := c.Universe.SliceOf(t.Elem())
		return exprX1(tout, fun)
	}
	c.Errorf("cannot slice %v: %v", t, node)
	return nil
}

func (c *Comp) sliceArrayMustBeAddressable(node *ast.SliceExpr, e *Expr) {
	panicking := true
	defer func() {
		if panicking {
			c.Errorf("cannot slice: array must be addressable: %v <%v>", node, e.Type)
		}
	}()
	c.placeOrAddress(node.X, PlaceAddress, nil)
	panicking = false
}
