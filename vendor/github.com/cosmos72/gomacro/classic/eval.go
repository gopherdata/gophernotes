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
 * eval.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"go/ast"
	r "reflect"

	"github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/reflect"
)

func (env *Env) Eval(src interface{}) (r.Value, []r.Value) {
	return env.EvalAst(env.Parse(src))
}

func (env *Env) Eval1(src interface{}) r.Value {
	return env.EvalAst1(env.Parse(src))
}

func (env *Env) EvalAst1(in ast2.Ast) r.Value {
	value, extraValues := env.EvalAst(in)
	if len(extraValues) > 1 {
		env.WarnExtraValues(extraValues)
	}
	return value
}

func (env *Env) EvalAst(in ast2.Ast) (r.Value, []r.Value) {
	switch in := in.(type) {
	case ast2.AstWithNode:
		if in != nil {
			return env.EvalNode(ast2.ToNode(in))
		}
	case ast2.AstWithSlice:
		if in != nil {
			var ret r.Value
			var rets []r.Value
			n := in.Size()
			for i := 0; i < n; i++ {
				ret, rets = env.EvalNode(ast2.ToNode(in.Get(i)))
			}
			return ret, rets
		}
	case nil:
		return None, nil
	default:
		return env.Errorf("EvalAst(): expecting <AstWithNode> or <AstWithSlice>, found: %v <%v>",
			in, r.TypeOf(in))
	}
	return env.Errorf("EvalAst(): expecting <AstWithNode> or <AstWithSlice>, found: nil")
}

func (env *Env) EvalNode(node ast.Node) (r.Value, []r.Value) {
	switch node := node.(type) {
	case ast.Decl:
		env.evalDecl(node)
	case ast.Expr:
		// Go expressions *DO* return values
		return env.evalExpr(node)
	case ast.Stmt:
		env.evalStatement(node)
	case *ast.File:
		env.evalFile(node)
	default:
		return env.Errorf("unimplemented Eval for %v <%v>", node, r.TypeOf(node))
	}
	// Go declarations, statements and files do not return values
	return None, nil
}

func (env *Env) EvalNode1(node ast.Node) r.Value {
	value, extraValues := env.EvalNode(node)
	if len(extraValues) > 1 {
		env.WarnExtraValues(extraValues)
	}
	return value
}

// macroexpand + collect + eval
func (env *Env) classicEval(form ast2.Ast) []r.Value {
	// macroexpansion phase.
	form, _ = env.MacroExpandAstCodewalk(form)

	if env.Options&OptShowMacroExpand != 0 {
		env.Debugf("after macroexpansion: %v", form.Interface())
	}

	// collect phase
	if env.Options&(OptCollectDeclarations|OptCollectStatements) != 0 {
		env.CollectAst(form)
	}

	// eval phase
	if env.Options&OptMacroExpandOnly != 0 {
		return reflect.PackValues(r.ValueOf(form.Interface()), nil)
	} else {
		return reflect.PackValues(env.EvalAst(form))
	}
}
