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
 * statement.go
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

type eBreak struct {
	label string
}

type eContinue struct {
	label string
}

func (_ eBreak) Error() string {
	return "break outside for or switch"
}

func (_ eContinue) Error() string {
	return "continue outside for"
}

type eReturn struct {
	results []r.Value
}

func (_ eReturn) Error() string {
	return "return outside function"
}

func (env *Env) evalBlock(block *ast.BlockStmt) (r.Value, []r.Value) {
	if block == nil || len(block.List) == 0 {
		return None, nil
	}
	env = NewEnv(env, "{}")

	return env.evalStatements(block.List)
}

func (env *Env) evalStatements(list []ast.Stmt) (r.Value, []r.Value) {
	ret := None
	var rets []r.Value

	for i := range list {
		ret, rets = env.evalStatement(list[i])
	}
	return ret, rets
}

func (env *Env) evalStatement(stmt ast.Stmt) (r.Value, []r.Value) {
again:
	if stmt != nil {
		env.Pos = stmt.Pos()
	}
	switch node := stmt.(type) {
	case *ast.AssignStmt:
		return env.evalAssignments(node)
	case *ast.BlockStmt:
		return env.evalBlock(node)
	case *ast.BranchStmt:
		return env.evalBranch(node)
	case *ast.CaseClause, *ast.CommClause:
		return env.Errorf("misplaced case: not inside switch or select: %v <%v>", node, r.TypeOf(node))
	case *ast.DeclStmt:
		return env.evalDecl(node.Decl)
	case *ast.DeferStmt:
		return env.evalDefer(node.Call)
	case *ast.EmptyStmt:
		return None, nil
	case *ast.ExprStmt:
		return env.evalExpr(node.X)
	case *ast.ForStmt:
		return env.evalFor(node)
	case *ast.GoStmt:
		return env.evalGo(node)
	case *ast.IfStmt:
		return env.evalIf(node)
	case *ast.IncDecStmt:
		return env.evalIncDec(node)
	case *ast.LabeledStmt:
		stmt = node
		goto again
	case *ast.RangeStmt:
		return env.evalForRange(node)
	case *ast.ReturnStmt:
		return env.evalReturn(node)
	case *ast.SelectStmt:
		return env.evalSelect(node)
	case *ast.SendStmt:
		return env.evalSend(node)
	case *ast.SwitchStmt:
		return env.evalSwitch(node)
	case *ast.TypeSwitchStmt:
		return env.evalTypeSwitch(node)
	default:
		return env.Errorf("unimplemented statement: %v <%v>", node, r.TypeOf(node))
	}
}

func (env *Env) evalBranch(node *ast.BranchStmt) (r.Value, []r.Value) {
	var label string
	if node.Label != nil {
		label = node.Label.Name
	}
	switch node.Tok {
	case token.BREAK:
		panic(eBreak{label})
	case token.CONTINUE:
		panic(eContinue{label})
	case token.GOTO:
		return env.Errorf("unimplemented: goto")
	case token.FALLTHROUGH:
		return env.Errorf("invalid fallthrough: not the last statement in a case")
	default:
		return env.Errorf("unimplemented branch: %v <%v>", node, r.TypeOf(node))
	}
}

func (env *Env) evalGo(stmt *ast.GoStmt) (r.Value, []r.Value) {
	if !MultiThread {
		env.Errorf("cannot create goroutine: %v\n\treason: this copy of gomacro was compiled with build tag 'gomacro_singlethread'", stmt)
	}

	node := stmt.Call
	fun := env.evalExpr1(node.Fun)

	switch fun.Kind() {
	case r.Struct:
		switch fun := fun.Interface().(type) {
		case Constructor:
			// evaluate args in the caller's goroutine
			t, args := env.evalConstructorArgs(fun, node)
			go fun.exec(env, t, args)
		case Function:
			// evaluate args in the caller's goroutine
			args := env.evalFunctionArgs(fun, node)
			go fun.exec(env, args)
		}
	case r.Func:
		// evaluate args in the caller's goroutine
		args := env.evalFuncArgs(fun, node)
		if node.Ellipsis == token.NoPos {
			go fun.Call(args)
		} else {
			go fun.CallSlice(args)
		}
	}
	return None, nil
}

func (env *Env) evalIf(node *ast.IfStmt) (r.Value, []r.Value) {
	if node.Init != nil {
		env = NewEnv(env, "if {}")
		_, _ = env.evalStatement(node.Init)
	}
	cond, _ := env.EvalNode(node.Cond)
	if cond.Kind() != r.Bool {
		cf := cond.Interface()
		return env.Errorf("if: invalid condition type <%T> %#v, expecting <bool>", cf, cf)
	}
	if cond.Bool() {
		return env.evalBlock(node.Body)
	} else if node.Else != nil {
		return env.evalStatement(node.Else)
	} else {
		return Nil, nil
	}
}

func (env *Env) evalIncDec(node *ast.IncDecStmt) (r.Value, []r.Value) {
	var op token.Token
	switch node.Tok {
	case token.INC:
		op = token.ADD_ASSIGN
	case token.DEC:
		op = token.SUB_ASSIGN
	default:
		return env.Errorf("unsupported *ast.IncDecStmt operation, expecting ++ or -- : %v <%v>", node, r.TypeOf(node))
	}
	place := env.evalPlace(node.X)
	return env.assignPlace(place, op, One), nil
}

func (env *Env) evalSend(node *ast.SendStmt) (r.Value, []r.Value) {
	channel := env.evalExpr1(node.Chan)
	if channel.Kind() != r.Chan {
		return env.Errorf("<- invoked on non-channel: %v evaluated to %v <%v>", node.Chan, channel, typeOf(channel))
	}
	value := env.evalExpr1(node.Value)
	channel.Send(value)
	return None, nil
}

func (env *Env) evalReturn(node *ast.ReturnStmt) (r.Value, []r.Value) {
	var rets []r.Value
	if len(node.Results) == 1 {
		// return foo() returns *all* the values returned by foo, not just the first one
		rets = reflect.PackValues(env.evalExpr(node.Results[0]))
	} else {
		rets = env.evalExprs(node.Results)
	}
	panic(eReturn{rets})
}
