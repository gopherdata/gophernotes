/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
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

	. "github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
)

func (env *Env) Eval(src interface{}) (r.Value, []r.Value) {
	return env.EvalAst(env.Parse(src))
}

func (env *Env) Eval1(src interface{}) r.Value {
	return env.EvalAst1(env.Parse(src))
}

func (env *Env) EvalAst1(in Ast) r.Value {
	value, extraValues := env.EvalAst(in)
	if len(extraValues) > 1 {
		env.WarnExtraValues(extraValues)
	}
	return value
}

func (env *Env) EvalAst(in Ast) (r.Value, []r.Value) {
	switch in := in.(type) {
	case AstWithNode:
		if in != nil {
			return env.EvalNode(ToNode(in))
		}
	case AstWithSlice:
		if in != nil {
			var ret r.Value
			var rets []r.Value
			n := in.Size()
			for i := 0; i < n; i++ {
				ret, rets = env.EvalNode(ToNode(in.Get(i)))
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
		return env.evalDecl(node)
	case ast.Expr:
		return env.evalExpr(node)
	case ast.Stmt:
		return env.evalStatement(node)
	case *ast.File:
		return env.evalFile(node)
	default:
		return env.Errorf("unimplemented Eval for %v <%v>", node, r.TypeOf(node))
	}
}

func (env *Env) EvalNode1(node ast.Node) r.Value {
	value, extraValues := env.EvalNode(node)
	if len(extraValues) > 1 {
		env.WarnExtraValues(extraValues)
	}
	return value
}

// parse, without macroexpansion
func (env *Env) ParseOnly(src interface{}) Ast {
	var form Ast
	switch src := src.(type) {
	case Ast:
		form = src
	case ast.Node:
		form = ToAst(src)
	default:
		bytes := ReadBytes(src)
		nodes := env.ParseBytes(bytes)

		if env.Options&OptShowParse != 0 {
			env.Debugf("after parse: %v", nodes)
		}
		switch len(nodes) {
		case 0:
			form = nil
		case 1:
			form = ToAst(nodes[0])
		default:
			form = NodeSlice{X: nodes}
		}
	}
	return form
}

// Parse, with macroexpansion
func (env *Env) Parse(src interface{}) Ast {
	form := env.ParseOnly(src)

	// macroexpansion phase.
	form, _ = env.MacroExpandAstCodewalk(form)

	if env.Options&OptShowMacroExpand != 0 {
		env.Debugf("after macroexpansion: %v", form.Interface())
	}
	if env.Options&(OptCollectDeclarations|OptCollectStatements) != 0 {
		env.CollectAst(form)
	}
	return form
}
