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
 * select.go
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

type selectLhsExpr struct {
	lhs [2]ast.Expr
	tok token.Token
}

func (env *Env) evalSelect(node *ast.SelectStmt) (ret r.Value, rets []r.Value) {
	if node.Body == nil || len(node.Body.List) == 0 {
		return None, nil
	}
	list := node.Body.List
	n := len(list)
	lhs := make([]selectLhsExpr, n)
	ops := make([]r.SelectCase, n)

	for i, stmt := range list {
		case_ := stmt.(*ast.CommClause)
		comm := case_.Comm
		if comm == nil {
			// default
			ops[i].Dir = r.SelectDefault
		} else {
			env.mustBeSelectStatement(comm, &lhs[i], &ops[i])
		}
	}
	i, recv, recvOk := r.Select(ops)
	case_ := list[i].(*ast.CommClause)
	return env.evalSelectBody(lhs[i], [2]r.Value{recv, r.ValueOf(recvOk)}, case_)
}

func (env *Env) mustBeSelectStatement(stmt ast.Stmt, lhs *selectLhsExpr, op *r.SelectCase) {
	switch node := stmt.(type) {
	case *ast.ExprStmt:
		// <-ch
		op.Dir = r.SelectRecv
		op.Chan = env.mustBeSelectRecv(stmt, node.X)
		return
	case *ast.AssignStmt:
		// v := <-ch or v = <-ch
		llist := node.Lhs
		lnum := len(llist)
		if (lnum == 1 || lnum == 2) && len(node.Rhs) == 1 {
			l0 := llist[0]
			var l1 ast.Expr
			if lnum == 2 {
				l1 = llist[1]
			}
			r0 := node.Rhs[0]
			switch node.Tok {
			case token.DEFINE:
				if _, ok := l0.(*ast.Ident); ok || l0 == nil {
					if _, ok := l1.(*ast.Ident); ok || l1 == nil {
						op.Dir = r.SelectRecv
						op.Chan = env.mustBeSelectRecv(node, r0)
						lhs.lhs[0] = l0
						lhs.lhs[1] = l1
						lhs.tok = node.Tok
						return
					}
				}
			case token.ASSIGN:
				op.Dir = r.SelectRecv
				op.Chan = env.mustBeSelectRecv(node, r0)
				lhs.lhs[0] = l0
				lhs.lhs[1] = l1
				lhs.tok = node.Tok
				return
			}
		}
	case *ast.SendStmt:
		// ch <- v
		op.Dir = r.SelectSend
		op.Chan = env.evalExpr1(node.Chan)
		op.Send = env.evalExpr1(node.Value)
		return
	}
	env.badSelectStatement(stmt)
}

func (env *Env) mustBeSelectRecv(stmt ast.Stmt, node ast.Expr) r.Value {
	for {
		switch expr := node.(type) {
		case *ast.ParenExpr:
			node = expr.X
			continue
		case *ast.UnaryExpr:
			if expr.Op == token.ARROW {
				return env.evalExpr1(expr.X)
			}
		}
		break
	}
	return env.badSelectStatement(stmt)
}

func (env *Env) badSelectStatement(stmt ast.Stmt) r.Value {
	env.Errorf("invalid select case, expecting [ch <- val] or [<-ch] or [var := <-ch] or [place = <-ch], found: %v <%v>",
		stmt, r.TypeOf(stmt))
	return None
}

func (env *Env) evalSelectBody(lhs selectLhsExpr, val [2]r.Value, case_ *ast.CommClause) (ret r.Value, rets []r.Value) {
	if case_ == nil || len(case_.Body) == 0 {
		// apply lhs side effects even without body
		if lhs.tok == token.ASSIGN {
			for i := 0; i < 2; i++ {
				if expr := lhs.lhs[i]; expr != nil {
					place := env.evalPlace(expr)
					env.assignPlace(place, token.ASSIGN, val[i])
				}
			}
		}
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
	if case_.Comm == nil {
		label = "default:"
	}
	env2 := NewEnv(env, label)
	for i := 0; i < 2; i++ {
		if expr := lhs.lhs[i]; expr != nil {
			if lhs.tok == token.DEFINE {
				env2.DefineVar(expr.(*ast.Ident).Name, nil, val[i])
			} else {
				place := env.evalPlace(expr)
				env.assignPlace(place, token.ASSIGN, val[i])
			}
		}
	}
	ret, rets = env2.evalStatements(case_.Body)
	panicking = false
	return
}
