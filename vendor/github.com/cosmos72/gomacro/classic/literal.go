/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * literal.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"go/ast"
	"go/token"
	r "reflect"
	"strconv"
	"strings"

	. "github.com/cosmos72/gomacro/base"
	bstrings "github.com/cosmos72/gomacro/base/strings"
)

func (env *Env) evalLiteral0(node *ast.BasicLit) interface{} {
	kind := node.Kind
	str := node.Value
	var ret interface{}

	switch kind {

	case token.INT:
		if strings.HasPrefix(str, "-") {
			i64, err := strconv.ParseInt(str, 0, 64)
			if err != nil {
				return env.Error(err)
			}
			// prefer int to int64. reason: in compiled Go,
			// type inference deduces int for all constants representable by an int
			i := int(i64)
			if int64(i) == i64 {
				return i
			}
			return i64
		} else {
			u64, err := strconv.ParseUint(str, 0, 64)
			if err != nil {
				return env.Error(err)
			}
			// prefer, in order: int, int64, uint, uint64. reason: in compiled Go,
			// type inference deduces int for all constants representable by an int
			i := int(u64)
			if i >= 0 && uint64(i) == u64 {
				return i
			}
			i64 := int64(u64)
			if i64 >= 0 && uint64(i64) == u64 {
				return i64
			}
			u := uint(u64)
			if uint64(u) == u64 {
				return u
			}
			return u64
		}

	case token.FLOAT:
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return env.Error(err)
		}
		ret = f

	case token.IMAG:
		if strings.HasSuffix(str, "i") {
			str = str[:len(str)-1]
		}
		im, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return env.Error(err)
		}
		ret = complex(0.0, im)
		// env.Debugf("evalLiteral(): parsed IMAG %s -> %T %#v -> %T %#v", str, im, im, ret, ret)

	case token.CHAR:
		ch, err := bstrings.UnescapeChar(str)
		if err != nil {
			env.Errorf("%v: invalid char literal: %s", err, str)
			return nil
		}
		return ch

	case token.STRING:
		return bstrings.UnescapeString(str)

	default:
		env.Errorf("unimplemented basic literal: %v", node)
		ret = nil
	}
	return ret
}

func (env *Env) evalCompositeLiteral(node *ast.CompositeLit) (r.Value, []r.Value) {
	t, ellipsis := env.evalType2(node.Type, false)
	if t == nil {
		env.Errorf("unimplemented type inference in composite literal: %v", node)
	}
	obj := Nil
	switch t.Kind() {
	case r.Map:
		obj = r.MakeMap(t)
		kt := t.Key()
		vt := t.Elem()
		for _, elt := range node.Elts {
			switch elt := elt.(type) {
			case *ast.KeyValueExpr:
				key := env.valueToType(env.evalExpr1(elt.Key), kt)
				val := env.valueToType(env.evalExpr1(elt.Value), vt)
				obj.SetMapIndex(key, val)
			default:
				env.Errorf("map literal: invalid element, expecting <*ast.KeyValueExpr>, found: %v <%v>", elt, r.TypeOf(elt))
			}
		}
		// in compiled Go, map literals are addressable
		place := r.New(t).Elem()
		place.Set(obj)
		obj = place
	case r.Array, r.Slice:
		vt := t.Elem()
		idx := -1
		val := Nil
		zero := Nil
		if t.Kind() == r.Array {
			obj = r.New(t).Elem()
		} else {
			zero = r.Zero(vt)
			obj = r.MakeSlice(t, 0, len(node.Elts))
		}
		for _, elt := range node.Elts {
			switch elt := elt.(type) {
			case *ast.KeyValueExpr:
				idx = int(env.valueToType(env.evalExpr1(elt.Key), TypeOfInt).Int())
				val = env.valueToType(env.evalExpr1(elt.Value), vt)
			default:
				// golang specs:
				// "An element without a key uses the previous element's index plus one.
				// If the first element has no key, its index is zero."
				idx++
				val = env.valueToType(env.evalExpr1(elt), vt)
			}
			if zero != Nil { // is slice, or array with unknown size [...]T{}
				for obj.Len() <= idx {
					obj = r.Append(obj, zero)
				}
			}
			obj.Index(idx).Set(val)
		}
		if ellipsis {
			// we can finally compute the actual array length...
			// too bad we have to copy the elements
			n := obj.Len()
			t = r.ArrayOf(n, vt)
			array := r.New(t).Elem()
			for i := 0; i < n; i++ {
				array.Index(i).Set(obj.Index(i))
			}
			obj = array
		} else if t.Kind() == r.Slice {
			// in compiled Go, slice literals are addressable
			place := r.New(t).Elem()
			place.Set(obj)
			obj = place
		}
	case r.Struct:
		obj = r.New(t).Elem()
		var pairs, elts bool
		var field r.Value
		var expr ast.Expr
		for idx, elt := range node.Elts {
			switch elt := elt.(type) {
			case *ast.KeyValueExpr:
				if elts {
					return env.Errorf("cannot mix keyed and non-keyed initializers in struct composite literal: %v", node)
				}
				pairs = true
				name := elt.Key.(*ast.Ident).Name
				field = obj.FieldByName(name)
				expr = elt.Value
			default:
				if pairs {
					return env.Errorf("cannot mix keyed and non-keyed initializers in struct composite literal: %v", node)
				}
				elts = true
				field = obj.Field(idx)
				expr = elt
			}
			val := env.valueToType(env.evalExpr1(expr), field.Type())
			field.Set(val)
		}
	default:
		env.Errorf("unexpected composite literal: %v", node)
	}
	return obj, nil
}
