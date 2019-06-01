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
 * switch_type.go
 *
 *  Created on: Mar 25, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"go/ast"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (env *Env) evalTypeSwitch(node *ast.TypeSwitchStmt) (ret r.Value, rets []r.Value) {
	// the scope of variables defined in the init and assign statements of a type switch
	// is the type switch itself
	if node.Init != nil {
		env = NewEnv(env, "type switch")
		env.evalStatement(node.Init)
	}
	varname, expr := env.mustBeTypeSwitchStatement(node.Assign)
	v := env.evalExpr1(expr)
	if node.Body == nil || len(node.Body.List) == 0 {
		return None, nil
	}
	var vt r.Type = nil
	if v != None && v != Nil {
		// go through interface{} to obtain actual concrete type
		val := v.Interface()
		v = r.ValueOf(val)
		if val != nil {
			vt = v.Type()
		}
	}
	var default_ *ast.CaseClause
	for _, stmt := range node.Body.List {
		case_ := stmt.(*ast.CaseClause)
		if case_.List == nil {
			// default will be executed later, if no case matches
			default_ = case_
		} else if t, ok := env.typecaseMatches(vt, case_.List); ok {
			return env.evalTypecaseBody(varname, t, v, case_, false)
		}
	}
	if default_ != nil {
		return env.evalTypecaseBody(varname, TypeOfInterface, v, default_, true)
	}
	return None, nil
}

func (env *Env) mustBeTypeSwitchStatement(node ast.Stmt) (*ast.Ident, ast.Expr) {
	switch stmt := node.(type) {
	case *ast.ExprStmt:
		// x.(type)
		return env.mustBeTypeSwitchAssert(node, stmt.X)
	case *ast.AssignStmt:
		// v := x.(type)
		if len(stmt.Lhs) == 1 && len(stmt.Rhs) == 1 && stmt.Tok == token.DEFINE {
			l := stmt.Lhs[0]
			if lhs, ok := l.(*ast.Ident); ok {
				r := stmt.Rhs[0]
				_, rhs := env.mustBeTypeSwitchAssert(node, r)
				return lhs, rhs
			}
		}
	}
	return env.badTypeSwitchStatement(node)
}

func (env *Env) mustBeTypeSwitchAssert(s ast.Stmt, x ast.Expr) (*ast.Ident, ast.Expr) {
	e, ok := x.(*ast.TypeAssertExpr)
	if !ok || e.Type != nil {
		return env.badTypeSwitchStatement(s)
	}
	return nil, e.X
}

func (env *Env) badTypeSwitchStatement(s ast.Stmt) (*ast.Ident, ast.Expr) {
	env.Errorf("invalid type switch expression, expecting x.(type) or v := x.(type), found: %v <%v>",
		s, r.TypeOf(s))
	return nil, nil
}

func (env *Env) typecaseMatches(vt r.Type, list []ast.Expr) (r.Type, bool) {
	for _, expr := range list {
		t := env.evalTypeOrNil(expr)
		if t == nil {
			if vt == nil {
				return TypeOfInterface, true
			}
		} else if vt.AssignableTo(t) {
			return t, true
		}
	}
	return nil, false
}

func (env *Env) evalTypecaseBody(varname *ast.Ident, t r.Type, val r.Value, case_ *ast.CaseClause, isDefault bool) (ret r.Value, rets []r.Value) {
	if case_ == nil || len(case_.Body) == 0 {
		return None, nil
	}
	panicking := true
	defer func() {
		if panicking {
			switch pan := recover().(type) {
			case eBreak:
				ret, rets = None, nil
			default:
				panic(pan)
			}
		}
	}()
	// each case body has its own environment
	label := "case:"
	if isDefault {
		label = "default:"
	}
	env = NewEnv(env, label)
	if varname != nil {
		env.DefineVar(varname.Name, t, val)
	}
	ret, rets = env.evalStatements(case_.Body)
	panicking = false
	return
}
