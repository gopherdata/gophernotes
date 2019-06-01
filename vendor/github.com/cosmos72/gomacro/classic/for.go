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
 * for.go
 *
 *  Created on: Feb 15, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"go/ast"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (env *Env) evalFor(node *ast.ForStmt) (r.Value, []r.Value) {
	// Debugf("evalFor() init = %#v, cond = %#v, post = %#v, body = %#v", node.Init, node.Cond, node.Post, node.Body)

	if node.Init != nil {
		env = NewEnv(env, "for {}")
		env.evalStatement(node.Init)
	}
	for {
		if node.Cond != nil {
			cond := env.evalExpr1(node.Cond)
			if cond.Kind() != r.Bool {
				cf := cond.Interface()
				return env.Errorf("for: invalid condition type <%T> %#v, expecting <bool>", cf, cf)
			}
			if !cond.Bool() {
				break
			}
		}
		if !env.evalForBodyOnce(node.Body) {
			break
		}
		if node.Post != nil {
			env.evalStatement(node.Post)
		}
	}
	return None, nil
}

func (env *Env) evalForRange(node *ast.RangeStmt) (r.Value, []r.Value) {
	// Debugf("evalForRange() init = %#v, cond = %#v, post = %#v, body = %#v", node.Init, node.Cond, node.Post, node.Body)

	container := env.evalExpr1(node.X)
	if container == Nil || container == None {
		return env.Errorf("invalid for range: cannot iterate on nil: %v evaluated to %v", node.X, container)
	}

	switch container.Kind() {
	case r.Chan:
		return env.evalForRangeChannel(container, node)
	case r.Map:
		return env.evalForRangeMap(container, node)
	case r.Slice, r.Array:
		return env.evalForRangeSlice(container, node)
	case r.String:
		// Golang specs https://golang.org/ref/spec#RangeClause
		// "For a string value, the "range" clause iterates over the Unicode code points in the string"
		return env.evalForRangeString(container.String(), node)
	case r.Ptr:
		if container.Elem().Kind() == r.Array {
			return env.evalForRangeSlice(container.Elem(), node)
		}
	}
	return env.Errorf("invalid for range: expecting array, channel, map, slice, string, or pointer to array, found: %v <%v>",
		container, typeOf(container))
}

func (env *Env) evalForRangeMap(obj r.Value, node *ast.RangeStmt) (r.Value, []r.Value) {
	knode := nilIfIdentUnderscore(node.Key)
	vnode := nilIfIdentUnderscore(node.Value)
	tok := node.Tok
	switch tok {
	case token.DEFINE:
		env = NewEnv(env, "range map {}")
		t := obj.Type()
		k := env.defineForIterVar(knode, t.Key())
		v := env.defineForIterVar(vnode, t.Elem())

		for _, key := range obj.MapKeys() {
			if k != Nil {
				k.Set(key)
			}
			if v != Nil {
				v.Set(obj.MapIndex(key))
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	case token.ASSIGN:
		for _, key := range obj.MapKeys() {
			// Golang specs https://golang.org/ref/spec#RangeClause
			// "Function calls on the left are evaluated once per iteration"
			//
			// we actually evaluate once per iteration the full expressions on the left
			if knode != nil {
				kplace := env.evalPlace(knode)
				env.assignPlace(kplace, tok, key)
			}
			if vnode != nil {
				vplace := env.evalPlace(vnode)
				env.assignPlace(vplace, tok, obj.MapIndex(key))
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	}
	return None, nil
}

func (env *Env) evalForRangeChannel(obj r.Value, node *ast.RangeStmt) (r.Value, []r.Value) {
	knode := nilIfIdentUnderscore(node.Key)
	if node.Value != nil {
		return env.Errorf("range expression is a channel: expecting at most one iteration variable, found two: %v %v", node.Key, node.Value)
	}

	tok := node.Tok
	switch tok {
	case token.DEFINE:
		env = NewEnv(env, "range channel {}")
		k := env.defineForIterVar(knode, obj.Type().Elem())

		for {
			recv, ok := obj.Recv()
			if !ok {
				break
			}
			if k != Nil {
				k.Set(recv)
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	case token.ASSIGN:
		for {
			recv, ok := obj.Recv()
			if !ok {
				break
			}
			// Golang specs https://golang.org/ref/spec#RangeClause
			// "Function calls on the left are evaluated once per iteration"
			//
			// we actually evaluate once per iteration the full expressions on the left
			if knode != nil {
				kplace := env.evalPlace(knode)
				env.assignPlace(kplace, tok, recv)
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	}
	return None, nil
}

func (env *Env) evalForRangeString(str string, node *ast.RangeStmt) (r.Value, []r.Value) {
	knode := nilIfIdentUnderscore(node.Key)
	vnode := nilIfIdentUnderscore(node.Value)
	tok := node.Tok
	switch tok {
	case token.DEFINE:
		env = NewEnv(env, "range string {}")
		k := env.defineForIterVar(knode, TypeOfInt)
		v := env.defineForIterVar(vnode, TypeOfRune)

		for i, rune := range str {
			if k != Nil {
				k.Set(r.ValueOf(i))
			}
			if v != Nil {
				v.Set(r.ValueOf(rune))
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	case token.ASSIGN:
		for i, rune := range str {
			// Golang specs https://golang.org/ref/spec#RangeClause
			// "Function calls on the left are evaluated once per iteration"
			//
			// we actually evaluate once per iteration the full expressions on the left
			if knode != nil {
				kplace := env.evalPlace(knode)
				env.assignPlace(kplace, tok, r.ValueOf(i))
			}
			if vnode != nil {
				vplace := env.evalPlace(vnode)
				env.assignPlace(vplace, tok, r.ValueOf(rune))
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	}
	return None, nil
}

func (env *Env) evalForRangeSlice(obj r.Value, node *ast.RangeStmt) (r.Value, []r.Value) {
	knode := nilIfIdentUnderscore(node.Key)
	vnode := nilIfIdentUnderscore(node.Value)
	tok := node.Tok
	switch tok {
	case token.DEFINE:
		env = NewEnv(env, "range slice/array {}")
		k := env.defineForIterVar(knode, TypeOfInt)
		v := env.defineForIterVar(vnode, obj.Type().Elem())

		n := obj.Len()
		for i := 0; i < n; i++ {
			if k != Nil {
				k.Set(r.ValueOf(i))
			}
			if v != Nil {
				v.Set(obj.Index(i))
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	case token.ASSIGN:
		n := obj.Len()
		for i := 0; i < n; i++ {
			// Golang specs https://golang.org/ref/spec#RangeClause
			// "Function calls on the left are evaluated once per iteration"
			//
			// we actually evaluate once per iteration the full expressions on the left
			if knode != nil {
				kplace := env.evalPlace(knode)
				env.assignPlace(kplace, tok, r.ValueOf(i))
			}
			if vnode != nil {
				vplace := env.evalPlace(vnode)
				env.assignPlace(vplace, tok, obj.Index(i))
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	}
	return None, nil
}

func (env *Env) evalForBodyOnce(node *ast.BlockStmt) (cont bool) {
	defer func() {
		if rec := recover(); rec != nil {
			switch rec := rec.(type) {
			case eBreak:
				cont = false
			case eContinue:
				cont = true
			default:
				panic(rec)
			}
		}
	}()
	env.evalBlock(node)
	return true
}

func (env *Env) defineForIterVar(node ast.Expr, t r.Type) r.Value {
	if node == nil || t == nil {
		return Nil
	}
	name := node.(*ast.Ident).Name
	env.DefineVar(name, t, r.Zero(t))
	return env.Binds.Get1(name)
}

func nilIfIdentUnderscore(node ast.Expr) ast.Expr {
	if ident, ok := node.(*ast.Ident); ok {
		if ident.Name == "_" {
			return nil
		}
	}
	return node
}
