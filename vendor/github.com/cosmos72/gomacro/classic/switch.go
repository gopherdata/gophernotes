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

func (env *Env) evalSwitch(node *ast.SwitchStmt) (ret r.Value, rets []r.Value) {
	if node.Init != nil {
		// the scope of variables defined in the init statement of a switch
		// is the switch itself
		env = NewEnv(env, "switch")
		env.evalStatement(node.Init)
	}
	var tag r.Value
	if node.Tag == nil {
		tag = True
	} else {
		tag = env.evalExpr1(node.Tag)
	}
	if node.Body == nil || len(node.Body.List) == 0 {
		return None, nil
	}
	isFallthrough := false
	cases := node.Body.List
	n := len(cases)
	default_i := n
	for i := 0; i < n; i++ {
		case_ := cases[i].(*ast.CaseClause)
		if !isFallthrough && case_.List == nil {
			// default will be executed later, if no case matches
			default_i = i
		} else if isFallthrough || env.caseMatches(tag, case_.List) {
			ret, rets, isFallthrough = env.evalCaseBody(i == default_i, case_)
			if !isFallthrough {
				return ret, rets
			}
		}
	}
	// even "default:" can end with fallthrough...
	for i := default_i; i < n; i++ {
		case_ := cases[i].(*ast.CaseClause)
		ret, rets, isFallthrough = env.evalCaseBody(i == default_i, case_)
		if !isFallthrough {
			return ret, rets
		}
	}
	return None, nil
}

func (env *Env) caseMatches(tag r.Value, list []ast.Expr) bool {
	var i interface{}
	var t r.Type = nil
	if tag != None && tag != Nil {
		i = tag.Interface()
		t = tag.Type()
	}
	for _, expr := range list {
		v := env.evalExpr1(expr)
		if t == nil {
			if v == Nil || v == None {
				return true
			}
		} else {
			v = env.valueToType(v, t)
			// https://golang.org/pkg/reflect
			// "To compare two Values, compare the results of the Interface method"
			if v.Interface() == i {
				return true
			}
		}
	}
	return false
}

func (env *Env) evalCaseBody(isDefault bool, case_ *ast.CaseClause) (ret r.Value, rets []r.Value, isFallthrough bool) {
	if case_ == nil || len(case_.Body) == 0 {
		return None, nil, false
	}
	body := case_.Body
	n := len(body)
	// implement fallthrough
	if last, ok := body[n-1].(*ast.BranchStmt); ok {
		if last.Tok == token.FALLTHROUGH {
			isFallthrough = true
			body = body[0 : n-1]
		}
	}

	// each case body has its own environment
	label := "case:"
	if isDefault {
		label = "default:"
	}
	panicking := true
	defer func() {
		if panicking {
			switch pan := recover().(type) {
			case eBreak:
				ret, rets, isFallthrough = None, nil, false
			default:
				panic(pan)
			}
		}
	}()
	env = NewEnv(env, label)
	ret, rets = env.evalStatements(body)
	panicking = false
	return
}
