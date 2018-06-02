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
 * index.go
 *
 *  Created on Apr 23, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) IndexExpr(node *ast.IndexExpr) *Expr { return c.indexExpr(node, true) }
func (c *Comp) IndexExpr1(node *ast.IndexExpr) *Expr { return c.indexExpr(node, false) }
func (c *Comp) indexExpr(node *ast.IndexExpr, multivalued bool) *Expr {
	obj := c.Expr1(node.X, nil)
	idx := c.Expr1(node.Index, nil)
	if obj.Untyped() {
		obj.ConstTo(obj.DefaultType())
	}

	t := obj.Type
	var ret *Expr
	switch t.Kind() {
	case r.Array, r.Slice, r.String:
		ret = c.vectorIndex(node, obj, idx)
	case r.Map:
		if multivalued {
			ret = c.mapIndex(node, obj, idx)
		} else {
			ret = c.mapIndex1(node, obj, idx)
		}

	case r.Ptr:
		if t.Elem().Kind() == r.Array {
			objfun := obj.AsX1()
			deref := exprFun(t.Elem(), func(env *Env) r.Value { return objfun(env).Elem() },
			)
			ret = c.vectorIndex(node, deref, idx)
			break
		}
		fallthrough
	default:
		c.Errorf("invalid operation: %v (type %v does not support indexing)", node, t)
		return nil
	}
	if obj.Const() && idx.Const() {
		ret.EvalConst(COptKeepUntyped)
	}
	return ret
}
func (c *Comp) vectorIndex(node *ast.IndexExpr, obj *Expr, idx *Expr) *Expr {
	k := idx.Type.Kind()
	cat := base.KindToCategory(k)
	if cat == r.Int || cat == r.Uint || idx.Untyped() {
		if !c.TypeOfInt().IdenticalTo(idx.Type) {
			idx = c.convert(idx, c.TypeOfInt(), node.Index)
		}
	} else {
		c.Errorf("non-integer %s index: %v <%v>", k, node.Index, idx.Type)
	}

	t := obj.Type
	if t.Kind() == r.String {
		return c.stringIndex(node, obj, idx)
	}

	t = t.Elem()
	objfun := obj.AsX1()
	var fun I
	if idx.Const() {
		i := idx.Value.(int)
		switch t.Kind() {
		case r.Bool:
			fun = func(env *Env) bool {
				objv := objfun(env)
				return objv.Index(i).Bool()
			}

		case r.Int:
			fun = func(env *Env) int {
				objv := objfun(env)
				return int(objv.Index(i).Int())
			}

		case r.Int8:
			fun = func(env *Env) int8 {
				objv := objfun(env)
				return int8(objv.Index(i).Int())
			}

		case r.Int16:
			fun = func(env *Env) int16 {
				objv := objfun(env)
				return int16(objv.Index(i).Int())
			}

		case r.Int32:
			fun = func(env *Env) int32 {
				objv := objfun(env)
				return int32(objv.Index(i).Int())
			}

		case r.Int64:
			fun = func(env *Env) int64 {
				objv := objfun(env)
				return objv.Index(i).Int()
			}

		case r.Uint:
			fun = func(env *Env) uint {
				objv := objfun(env)
				return uint(objv.Index(i).Uint())
			}

		case r.Uint8:
			fun = func(env *Env) uint8 {
				objv := objfun(env)
				return uint8(objv.Index(i).Uint())
			}

		case r.Uint16:
			fun = func(env *Env) uint16 {
				objv := objfun(env)
				return uint16(objv.Index(i).Uint())
			}

		case r.Uint32:
			fun = func(env *Env) uint32 {
				objv := objfun(env)
				return uint32(objv.Index(i).Uint())
			}

		case r.Uint64:
			fun = func(env *Env) uint64 {
				objv := objfun(env)
				return objv.Index(i).Uint()
			}

		case r.Uintptr:
			fun = func(env *Env) uintptr {
				objv := objfun(env)
				return uintptr(objv.Index(i).Uint())
			}

		case r.Float32:
			fun = func(env *Env) float32 {
				objv := objfun(env)
				return float32(objv.Index(i).Float())
			}

		case r.Float64:
			fun = func(env *Env) float64 {
				objv := objfun(env)
				return objv.Index(i).Float()
			}

		case r.Complex64:
			fun = func(env *Env) complex64 {
				objv := objfun(env)
				return complex64(objv.Index(i).Complex())
			}

		case r.Complex128:
			fun = func(env *Env) complex128 {
				objv := objfun(env)
				return objv.Index(i).Complex()
			}

		case r.String:
			fun = func(env *Env) string {
				objv := objfun(env)
				return objv.Index(i).String()
			}

		default:
			fun = func(env *Env) r.Value {
				objv := objfun(env)
				return objv.Index(i)

			}

		}
	} else {
		idxfun := idx.WithFun().(func(*Env) int)
		switch t.Kind() {
		case r.Bool:
			fun = func(env *Env) bool {
				objv := objfun(env)
				i := idxfun(env)
				return objv.Index(i).Bool()
			}

		case r.Int:
			fun = func(env *Env) int {
				objv := objfun(env)
				i := idxfun(env)
				return int(objv.Index(i).Int())
			}

		case r.Int8:
			fun = func(env *Env) int8 {
				objv := objfun(env)
				i := idxfun(env)
				return int8(objv.Index(i).Int())
			}

		case r.Int16:
			fun = func(env *Env) int16 {
				objv := objfun(env)
				i := idxfun(env)
				return int16(objv.Index(i).Int())
			}

		case r.Int32:
			fun = func(env *Env) int32 {
				objv := objfun(env)
				i := idxfun(env)
				return int32(objv.Index(i).Int())
			}

		case r.Int64:
			fun = func(env *Env) int64 {
				objv := objfun(env)
				i := idxfun(env)
				return objv.Index(i).Int()
			}

		case r.Uint:
			fun = func(env *Env) uint {
				objv := objfun(env)
				i := idxfun(env)
				return uint(objv.Index(i).Uint())
			}

		case r.Uint8:
			fun = func(env *Env) uint8 {
				objv := objfun(env)
				i := idxfun(env)
				return uint8(objv.Index(i).Uint())
			}

		case r.Uint16:
			fun = func(env *Env) uint16 {
				objv := objfun(env)
				i := idxfun(env)
				return uint16(objv.Index(i).Uint())
			}

		case r.Uint32:
			fun = func(env *Env) uint32 {
				objv := objfun(env)
				i := idxfun(env)
				return uint32(objv.Index(i).Uint())
			}

		case r.Uint64:
			fun = func(env *Env) uint64 {
				objv := objfun(env)
				i := idxfun(env)
				return objv.Index(i).Uint()
			}

		case r.Uintptr:
			fun = func(env *Env) uintptr {
				objv := objfun(env)
				i := idxfun(env)
				return uintptr(objv.Index(i).Uint())
			}

		case r.Float32:
			fun = func(env *Env) float32 {
				objv := objfun(env)
				i := idxfun(env)
				return float32(objv.Index(i).Float())
			}

		case r.Float64:
			fun = func(env *Env) float64 {
				objv := objfun(env)
				i := idxfun(env)
				return objv.Index(i).Float()
			}

		case r.Complex64:
			fun = func(env *Env) complex64 {
				objv := objfun(env)
				i := idxfun(env)
				return complex64(objv.Index(i).Complex())
			}

		case r.Complex128:
			fun = func(env *Env) complex128 {
				objv := objfun(env)
				i := idxfun(env)
				return objv.Index(i).Complex()
			}

		case r.String:
			fun = func(env *Env) string {
				objv := objfun(env)
				i := idxfun(env)
				return objv.Index(i).String()
			}

		default:
			fun = func(env *Env) r.Value {
				objv := objfun(env)
				i := idxfun(env)
				return objv.Index(i)

			}

		}
	}
	return exprFun(t, fun)
}
func (c *Comp) stringIndex(node *ast.IndexExpr, obj *Expr, idx *Expr) *Expr {
	idxfun := idx.WithFun().(func(*Env) int)
	objfun := obj.WithFun().(func(*Env) string)
	var fun func(env *Env) uint8
	if obj.Const() {
		str := obj.Value.(string)
		fun = func(env *Env) uint8 {
			i := idxfun(env)
			return str[i]
		}
	} else if idx.Const() {
		i := idx.Value.(int)
		fun = func(env *Env) uint8 {
			str := objfun(env)
			return str[i]
		}
	} else {
		fun = func(env *Env) uint8 {
			str := objfun(env)
			i := idxfun(env)
			return str[i]
		}
	}

	e := c.exprUint8(fun)
	if obj.Const() && idx.Const() {
		panicking := true
		defer func() {
			if panicking {
				recover()
				c.Errorf("string index out of range: %v", node)
			}
		}()
		e.EvalConst(COptKeepUntyped)
		panicking = false
	}
	return e
}
func (c *Comp) mapIndex(node *ast.IndexExpr, obj *Expr, idx *Expr) *Expr {
	t := obj.Type
	tkey := t.Key()
	tval := t.Elem()
	idxconst := idx.Const()
	if idxconst {
		idx.ConstTo(tkey)
	} else if idx.Type == nil || !idx.Type.AssignableTo(tkey) {
		c.Errorf("cannot use %v <%v> as <%v> in map index", node.Index, idx.Type, tkey)
	}

	objfun := obj.AsX1()
	zero := xr.Zero(tval)
	var fun func(env *Env) (r.Value, []r.Value)
	if idxconst {
		key := r.ValueOf(idx.Value)
		fun = func(env *Env) (r.Value, []r.Value) {
			obj := objfun(env)
			val := obj.MapIndex(key)
			var ok r.Value
			if val == base.Nil {
				val = zero
				ok = base.False
			} else {
				ok = base.True
			}
			return val, []r.Value{val, ok}
		}
	} else {
		keyfun := idx.AsX1()
		fun = func(env *Env) (r.Value, []r.Value) {
			obj := objfun(env)
			key := keyfun(env)
			val := obj.MapIndex(key)
			var ok r.Value
			if val == base.Nil {
				val = zero
				ok = base.False
			} else {
				ok = base.True
			}
			return val, []r.Value{val, ok}
		}
	}
	return exprXV([]xr.Type{tval, c.TypeOfBool()}, fun)
}
func (c *Comp) mapIndex1(node *ast.IndexExpr, obj *Expr, idx *Expr) *Expr {
	t := obj.Type
	tkey := t.Key()
	tval := t.Elem()
	idxconst := idx.Const()
	if idxconst {
		idx.ConstTo(tkey)
	} else if idx.Type == nil || !idx.Type.AssignableTo(tkey) {
		c.Errorf("cannot use %v <%v> as <%v> in map index", node.Index, idx.Type, tkey)
	}

	objfun := obj.AsX1()
	var fun I
	if idxconst {
		key := r.ValueOf(idx.Value)
		switch tval.Kind() {
		case r.Bool:
			fun = func(env *Env) bool {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result bool

				if v != base.Nil {
					result = v.Bool()
				}
				return result
			}
		case r.Int:
			fun = func(env *Env) int {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result int

				if v != base.Nil {
					result = int(v.Int())
				}
				return result
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result int8

				if v != base.Nil {
					result = int8(v.Int())
				}
				return result
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result int16

				if v != base.Nil {
					result = int16(v.Int())
				}
				return result
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result int32
				if v != base.Nil {
					result = int32(v.Int())
				}
				return result
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result int64
				if v != base.Nil {
					result = v.Int()
				}
				return result
			}
		case r.Uint:
			fun = func(env *Env) uint {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result uint
				if v != base.Nil {
					result = uint(v.Uint())
				}
				return result
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result uint8
				if v != base.Nil {
					result =
						uint8(v.Uint())
				}
				return result
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result uint16
				if v != base.Nil {
					result =

						uint16(v.Uint())
				}
				return result
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result uint32
				if v != base.Nil {
					result =

						uint32(v.Uint())
				}
				return result
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result uint64
				if v != base.Nil {
					result = v.Uint()
				}
				return result
			}

		case r.Uintptr:
			fun = func(env *Env) uintptr {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result uintptr
				if v != base.Nil {
					result =

						uintptr(v.Uint())
				}
				return result
			}

		case r.Float32:
			fun = func(env *Env) float32 {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result float32
				if v != base.Nil {
					result =

						float32(v.Float())
				}
				return result
			}

		case r.Float64:
			fun = func(env *Env) float64 {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result float64
				if v != base.Nil {
					result = v.Float()
				}
				return result
			}

		case r.Complex64:
			fun = func(env *Env) complex64 {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result complex64
				if v != base.Nil {
					result =

						complex64(v.Complex())
				}
				return result
			}

		case r.Complex128:
			fun = func(env *Env) complex128 {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result complex128
				if v != base.Nil {
					result = v.Complex()
				}
				return result
			}

		case r.String:
			fun = func(env *Env) string {
				obj := objfun(env)
				v := obj.MapIndex(key)
				var result string
				if v != base.Nil {
					result = v.String()
				}
				return result
			}

		default:
			{
				zero := xr.Zero(tval)
				fun = func(env *Env) r.Value {
					obj := objfun(env)
					result := obj.MapIndex(key)
					if result == base.Nil {
						result = zero
					}
					return result
				}
			}

		}
	} else {
		keyfun := idx.AsX1()
		switch tval.Kind() {
		case r.Bool:
			fun = func(env *Env) bool {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result bool
				if v != base.Nil {
					result = v.Bool()
				}
				return result
			}

		case r.Int:
			fun = func(env *Env) int {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result int
				if v != base.Nil {
					result =

						int(v.Int())
				}
				return result
			}

		case r.Int8:
			fun = func(env *Env) int8 {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result int8
				if v != base.Nil {
					result =

						int8(v.Int())
				}
				return result
			}

		case r.Int16:
			fun = func(env *Env) int16 {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result int16
				if v != base.Nil {
					result =

						int16(v.Int())
				}
				return result
			}

		case r.Int32:
			fun = func(env *Env) int32 {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result int32
				if v != base.Nil {
					result =

						int32(v.Int())
				}
				return result
			}

		case r.Int64:
			fun = func(env *Env) int64 {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result int64
				if v != base.Nil {
					result = v.Int()
				}
				return result
			}

		case r.Uint:
			fun = func(env *Env) uint {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result uint
				if v != base.Nil {
					result =

						uint(v.Uint())
				}
				return result
			}

		case r.Uint8:
			fun = func(env *Env) uint8 {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result uint8
				if v != base.Nil {
					result =

						uint8(v.Uint())
				}
				return result
			}

		case r.Uint16:
			fun = func(env *Env) uint16 {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result uint16
				if v != base.Nil {
					result =

						uint16(v.Uint())
				}
				return result
			}

		case r.Uint32:
			fun = func(env *Env) uint32 {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result uint32
				if v != base.Nil {
					result =

						uint32(v.Uint())
				}
				return result
			}

		case r.Uint64:
			fun = func(env *Env) uint64 {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result uint64
				if v != base.Nil {
					result = v.Uint()
				}
				return result
			}

		case r.Uintptr:
			fun = func(env *Env) uintptr {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result uintptr
				if v != base.Nil {
					result =

						uintptr(v.Uint())
				}
				return result
			}

		case r.Float32:
			fun = func(env *Env) float32 {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result float32
				if v != base.Nil {
					result =

						float32(v.Float())
				}
				return result
			}

		case r.Float64:
			fun = func(env *Env) float64 {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result float64
				if v != base.Nil {
					result = v.Float()
				}
				return result
			}

		case r.Complex64:
			fun = func(env *Env) complex64 {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result complex64
				if v != base.Nil {
					result =

						complex64(v.Complex())
				}
				return result
			}

		case r.Complex128:
			fun = func(env *Env) complex128 {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result complex128
				if v != base.Nil {
					result = v.Complex()
				}
				return result
			}

		case r.String:
			fun = func(env *Env) string {
				obj := objfun(env)
				key := keyfun(env)
				v := obj.MapIndex(key)
				var result string
				if v != base.Nil {
					result = v.String()
				}
				return result
			}

		default:
			{
				zero := xr.Zero(tval)
				fun = func(env *Env) r.Value {
					obj := objfun(env)
					key := keyfun(env)
					result := obj.MapIndex(key)
					if result == base.Nil {
						result = zero
					}
					return result
				}
			}

		}
	}
	return exprFun(tval, fun)
}
func (c *Comp) IndexPlace(node *ast.IndexExpr, opt PlaceOption) *Place {
	obj := c.Expr1(node.X, nil)
	idx := c.Expr1(node.Index, nil)
	if obj.Untyped() {
		obj.ConstTo(obj.DefaultType())
	}

	t := obj.Type
	switch t.Kind() {
	case r.Array, r.Slice:
		return c.vectorPlace(node, obj, idx)
	case r.String:

		c.Errorf("%s a byte in a string: %v", opt, node)
		return nil
	case r.Map:

		if opt == PlaceAddress {
			c.Errorf("%s a map element: %v", opt, node)
			return nil
		}
		return c.mapPlace(node, obj, idx)
	case r.Ptr:
		if t.Elem().Kind() == r.Array {
			return c.vectorPtrPlace(node, obj, idx)
		}

		fallthrough
	default:
		c.Errorf("invalid operation: %v (type %v does not support indexing)", node, t)
		return nil
	}
}
func (c *Comp) mapPlace(node *ast.IndexExpr, obj *Expr, idx *Expr) *Place {
	tmap := obj.Type
	tkey := tmap.Key()
	idxconst := idx.Const()
	if idxconst {
		idx.ConstTo(tkey)
	} else if idx.Type == nil || !idx.Type.AssignableTo(tkey) {
		c.Errorf("cannot use %v <%v> as type <%v> in map index: %v", node.Index, idx.Type, tkey, node)
	}
	return &Place{Var: Var{Type: tmap.Elem()}, Fun: obj.AsX1(), MapKey: idx.AsX1(), MapType: tmap}
}
func (c *Comp) vectorPlace(node *ast.IndexExpr, obj *Expr, idx *Expr) *Place {
	idxconst := idx.Const()
	if idxconst {
		idx.ConstTo(c.TypeOfInt())
	} else if idx.Type == nil || !idx.Type.AssignableTo(c.TypeOfInt()) {
		c.Errorf("non-integer %s index: %v <%v>", obj.Type.Kind(), node.Index, idx.Type)
	}

	t := obj.Type.Elem()
	objfun := obj.AsX1()
	var fun, addr func(env *Env) r.Value
	if idxconst {
		i := idx.Value.(int)
		fun = func(env *Env) r.Value {
			objv := objfun(env)
			return objv.Index(i)
		}
		addr = func(env *Env) r.Value {
			objv := objfun(env)
			return objv.Index(i).Addr()
		}
	} else {
		idxfun := idx.WithFun().(func(*Env) int)
		fun = func(env *Env) r.Value {
			objv := objfun(env)
			i := idxfun(env)
			return objv.Index(i)
		}
		addr = func(env *Env) r.Value {
			objv := objfun(env)
			i := idxfun(env)
			return objv.Index(i).Addr()
		}
	}
	return &Place{Var: Var{Type: t}, Fun: fun, Addr: addr}
}
func (c *Comp) vectorPtrPlace(node *ast.IndexExpr, obj *Expr, idx *Expr) *Place {
	idxconst := idx.Const()
	if idxconst {
		idx.ConstTo(c.TypeOfInt())
	} else if idx.Type == nil || !idx.Type.AssignableTo(c.TypeOfInt()) {
		c.Errorf("non-integer %s index: %v <%v>", obj.Type.Kind(), node.Index, idx.Type)
	}

	t := obj.Type.Elem().Elem()
	objfun := obj.AsX1()
	var fun, addr func(env *Env) r.Value
	if idxconst {
		i := idx.Value.(int)
		fun = func(env *Env) r.Value {
			objv := objfun(env).Elem()
			return objv.Index(i)
		}
		addr = func(env *Env) r.Value {
			objv := objfun(env).Elem()
			return objv.Index(i).Addr()
		}
	} else {
		idxfun := idx.WithFun().(func(*Env) int)
		fun = func(env *Env) r.Value {
			objv := objfun(env).Elem()
			i := idxfun(env)
			return objv.Index(i)
		}
		addr = func(env *Env) r.Value {
			objv := objfun(env).Elem()
			i := idxfun(env)
			return objv.Index(i).Addr()
		}
	}
	return &Place{Var: Var{Type: t}, Fun: fun, Addr: addr}
}
