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
 * assignment.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"go/ast"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/reflect"
)

type placeType struct {
	obj    r.Value // the map to modify, or a settable r.Value
	mapkey r.Value // the map key to set, or Nil
}

// dummy place for assignment to _
var _Place = placeType{
	obj: r.ValueOf(struct{}{}),
}

func (env *Env) evalAssignments(node *ast.AssignStmt) (r.Value, []r.Value) {
	left := node.Lhs
	right := node.Rhs
	op := node.Tok
	nleft := len(left)
	nright := len(right)

	if nright != 1 && nleft != nright {
		return env.Errorf("value count mismatch: cannot assign %d values to %d places: %v", nright, nleft, node)
	}

	// side effects happen left to right, with some unspecified cases,
	// so first Eval() all node.Lhs, then Eval() all node.Rhs
	// https://golang.org/ref/spec#Order_of_evaluation

	if op == token.DEFINE {
		names := make([]string, nleft)
		for i := 0; i < nleft; i++ {
			ident, ok := left[i].(*ast.Ident)
			if !ok {
				return env.Errorf("variable declaration: invalid identifier: %v", left[i])
			}
			names[i] = ident.Name
		}
		values := env.evalExprsMultipleValues(right, nleft)
		return env.defineConstsVarsOrFuncs(names, nil, values, false)

	} else {
		places := env.evalPlaces(left)
		values := env.evalExprsMultipleValues(right, nleft)
		return env.assignPlaces(places, op, values)
	}
}

func (env *Env) evalPlaces(node []ast.Expr) []placeType {
	n := len(node)
	places := make([]placeType, n)
	for i := 0; i < n; i++ {
		places[i] = env.evalPlace(node[i])
	}
	return places
}

func (env *Env) evalPlace(node ast.Expr) placeType {
	obj := Nil
	// ignore parenthesis: (expr) = value is the same as expr = value
	for {
		if paren, ok := node.(*ast.ParenExpr); ok {
			node = paren.X
		} else {
			break
		}
	}
	switch node := node.(type) {
	case *ast.IndexExpr:
		obj = env.evalExpr1(node.X)
		index := env.evalExpr1(node.Index)

		switch obj.Kind() {
		case r.Map:
			// make a copy of obj and index, to protect against "evil assignment" m, i, m[i] = nil, 1, 2  where m is a map
			if obj != Nil && obj.CanSet() {
				obj = obj.Convert(obj.Type())
			}
			if index != Nil && index.CanSet() {
				index = index.Convert(index.Type())
			}
			return placeType{obj, index}
		default:
			if obj.Kind() != r.Ptr || obj.Elem().Kind() != r.Array {
				env.Errorf("unsupported index operation: %v [ %v ]. not an array, map, slice or string: %v <%v>",
					node.X, index, obj, typeOf(obj))
				return placeType{}
			}
			obj = obj.Elem()
			fallthrough
		case r.Array, r.Slice, r.String:
			i, ok := env.toInt(index)
			if !ok {
				env.Errorf("invalid index, expecting an int: %v <%v>", index, typeOf(index))
				return placeType{}
			}
			obj = obj.Index(int(i))
		}
	case *ast.Ident:
		if node.Name == "_" {
			return _Place
		}
		obj = env.evalExpr1(node)

	default:
		obj = env.evalExpr1(node)
	}
	if !obj.CanSet() {
		env.Errorf("cannot assign to read-only location: %v", node)
		return placeType{}
	}
	return placeType{obj, Nil}
}

func (env *Env) assignPlaces(places []placeType, op token.Token, values []r.Value) (r.Value, []r.Value) {
	n := len(places)
	if n == 1 {
		return env.assignPlace(places[0], op, values[0]), nil
	}
	// the naive loop
	//   for i := range places { env.assignPlace(places[i], op, values[i]) }
	// is bugged. It breaks, among others, the common Go idiom to swap two values: a,b = b,a
	//
	// More in general, Go guarantees that all assignments happen *as if*
	// the rhs values, and all lhs operands of indexing, dereferencing and struct field access,
	// were copied to temporary locations before the assignments.
	// That's exactly what we must do.
	for i := 0; i < n; i++ {
		p := &places[i]
		v := p.mapkey
		if v != Nil && v.CanSet() {
			p.mapkey = v.Convert(v.Type()) // r.Value.Convert() makes a copy
		}
		v = values[i]
		if v != Nil && v.CanSet() {
			values[i] = v.Convert(v.Type()) // r.Value.Convert() makes a copy
		}
	}
	for i := 0; i < n; i++ {
		values[i] = env.assignPlace(places[i], op, values[i])
	}
	return reflect.UnpackValues(values)
}

func (env *Env) assignPlace(place placeType, op token.Token, value r.Value) r.Value {
	obj := place.obj
	if obj == _Place.obj {
		return value
	}
	key := place.mapkey
	if key == Nil {
		t := typeOf(obj)
		value = env.valueToType(value, t)
		if op != token.ASSIGN {
			value = env.evalBinaryExpr(obj, op, value)
		}
		obj.Set(value)
		return value
	}
	// map[key] OP value
	key = env.valueToType(key, obj.Type().Key())

	// env.Debugf("setting map[key]: %v <%v> [%v <%v>] %s %v <%v>", obj, TypeOf(obj), key, TypeOf(key), op, value, TypeOf(value))

	currValue, _, t := env.mapIndex(obj, key)
	value = env.valueToType(value, t)
	if op != token.ASSIGN {
		value = env.evalBinaryExpr(currValue, op, value)
		value = env.valueToType(value, t) // in case evalBinaryExpr() converted it
	}
	obj.SetMapIndex(key, value)
	return value
}
