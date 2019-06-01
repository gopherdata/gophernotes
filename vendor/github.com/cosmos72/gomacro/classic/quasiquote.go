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
 * quasiquote.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"fmt"
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	etoken "github.com/cosmos72/gomacro/go/etoken"
)

func (env *Env) evalQuote(node *ast.BlockStmt) ast.Node {
	return SimplifyNodeForQuote(node, true)
}

// evalQuasiquote evaluates the body of a quasiquote{} represented as ast.Node
func (env *Env) evalQuasiquote(node *ast.BlockStmt) ast.Node {
	// we invoke SimplifyNodeForQuote() at the end, not at the beginning.
	// reason: to support quasiquote{unquote_splice ...}
	toUnwrap := node != SimplifyNodeForQuote(node, true)

	in := ToAst(node)
	out := env.evalQuasiquoteAst(in, 1)
	ret := ToNode(out)
	return SimplifyNodeForQuote(ret, toUnwrap)
}

// evalQuasiquoteAst evaluates the body of a quasiquote{} represented as Ast
// use unified API to traverse ast.Node... every other solution is a nightmare
func (env *Env) evalQuasiquoteAst(in Ast, depth int) (out Ast) {
	if in == nil {
		return nil
	}
	inSlice, canSplice := in.(AstWithSlice)
	env.debugQuasiQuote("evaluating", depth, canSplice, in.Interface())
	if !canSplice {
		in = UnwrapTrivialAst(in) // drill through DeclStmt, ExprStmt, ParenExpr, one-element BlockStmt
	}
	if in == nil || in.Size() == 0 {
		return in
	}

	if !canSplice {
		if in, ok := in.(UnaryExpr); ok {
			switch in.Op() {
			case etoken.QUASIQUOTE:
				// equivalent to ToAst(form.p.X.(*ast.FuncLit).Body)
				toexpand := in.Get(0).Get(1)
				env.debugQuasiQuote("recursing inside QUASIQUOTE", depth+1, canSplice, toexpand.Interface())
				expansion := env.evalQuasiquoteAst(toexpand, depth+1)
				return MakeQuote2(in, expansion.(AstWithNode))
			case etoken.UNQUOTE:
				if depth <= 1 {
					y := env.evalUnquote(in)
					return AnyToAst(y, "unquote")
				} else {
					// equivalent to ToAst(form.p.X.(*ast.FuncLit).Body)
					toexpand := in.Get(0).Get(1)
					env.debugQuasiQuote("recursing inside UNQUOTE", depth-1, canSplice, toexpand.Interface())
					expansion := env.evalQuasiquoteAst(toexpand, depth-1)
					return MakeQuote2(in, expansion.(AstWithNode))
				}
			case etoken.UNQUOTE_SPLICE:
				y := in.Interface()
				env.Errorf("quasiquote: cannot splice in single-statement context: %v <%v>", y, r.TypeOf(y))
				return nil
			}
		}

		out := in.New()
		ni := in.Size()
		for i := 0; i < ni; i++ {
			child := in.Get(i)
			if child == nil {
				env.debugQuasiQuote("child is nil", depth, canSplice, child)
			} else {
				env.debugQuasiQuote("general case: recurse on child", depth, canSplice, child.Interface())
				child = env.evalQuasiquoteAst(child, depth)
			}
			out.Set(i, child)
		}
		return out
	}

	outSlice := inSlice.New().(AstWithSlice)
	ni := inSlice.Size()
	for i := 0; i < ni; i++ {
		// drill through DeclStmt, ExprStmt, ParenExpr
		child := UnwrapTrivialAstKeepBlocks(inSlice.Get(i))
		switch child := child.(type) {
		case UnaryExpr:
			switch child.Op() {
			case etoken.QUASIQUOTE:
				// equivalent to ToAst(form.p.X.(*ast.FuncLit).Body)
				toexpand := child.Get(0).Get(1)
				env.debugQuasiQuote("recursing inside QUASIQUOTE", depth+1, canSplice, toexpand.Interface())
				expansion := env.evalQuasiquoteAst(toexpand, depth+1)
				child = MakeQuote2(child, expansion.(AstWithNode))
				outSlice = outSlice.Append(child)
				goto Next
			case etoken.UNQUOTE, etoken.UNQUOTE_SPLICE:
				// complication: in Common Lisp, the right-most unquote pairs with the left-most comma!
				// we implement the same mechanics, so we must drill down to the last unquote/unquote_splice
				// and, for unquote_splice, create a copy of the unquote/unquote_splice stack for each result.
				// Example:
				//   x:=quote{7; 8}
				//   quasiquote{quasiquote{1; unquote{2}; unquote{unquote_splice{x}}}}
				// must return
				//   quasiquote{1; unquote{2}; unquote{7}; unquote{8}}
				lastUnquote, unquoteDepth := DescendNestedUnquotes(child)

				op := lastUnquote.Op()

				env.debugQuasiQuote(fmt.Sprintf("inside %s, lastUnquote is %s (unquoteDepth = %d)",
					etoken.String(child.Op()), etoken.String(op), unquoteDepth), depth, canSplice, child)

				if unquoteDepth > depth {
					env.Errorf("%s not inside quasiquote: %v <%v>", etoken.String(op), lastUnquote, r.TypeOf(lastUnquote))
					return nil
				} else if unquoteDepth < depth {
					toexpand := child.Get(0).Get(1)
					env.debugQuasiQuote(fmt.Sprintf("recursing inside %s, lastUnquote is %s", etoken.String(child.Op()), etoken.String(op)),
						depth-1, canSplice, toexpand.Interface())
					expansion := env.evalQuasiquoteAst(toexpand, depth-1)
					child = MakeQuote2(child, expansion.(AstWithNode))
					outSlice = outSlice.Append(child)
				} else {
					env.debugQuasiQuote("calling unquote on", depth-unquoteDepth, canSplice, lastUnquote.Interface())
					toInsert := AnyToAst(env.evalUnquote(lastUnquote), etoken.String(op))
					if toInsert == nil {
						env.debugQuasiQuote("unquote returned", depth-unquoteDepth, canSplice, toInsert)
					} else {
						env.debugQuasiQuote("unquote returned", depth-unquoteDepth, canSplice, toInsert.Interface())
					}
					if op == etoken.UNQUOTE {
						stack := DuplicateNestedUnquotes(child, unquoteDepth-1, toInsert)
						outSlice = outSlice.Append(stack)
					} else if toInsert != nil {
						toSplice := ToAstWithSlice(toInsert, "unquote_splice")
						nj := toSplice.Size()
						for j := 0; j < nj; j++ {
							stack := DuplicateNestedUnquotes(child, unquoteDepth-1, toSplice.Get(j))
							outSlice = outSlice.Append(stack)
						}
					}
				}
				goto Next
			}
		}
		if child == nil {
			env.debugQuasiQuote("child is nil", depth, canSplice, child)
		} else {
			env.debugQuasiQuote("general case: recurse on child", depth, canSplice, child.Interface())
			child = env.evalQuasiquoteAst(child, depth)
		}
		outSlice = outSlice.Append(child)
	Next:
		env.debugQuasiQuote("accumulated", depth, canSplice, outSlice.Interface())
	}
	return outSlice
}

func (env *Env) debugQuasiQuote(msg string, depth int, canSplice bool, x interface{}) {
	if env.Options&OptDebugQuasiquote != 0 {
		env.Debugf("quasiquote: %s (depth = %d, canSplice = %v)\n%v <%v>", msg, depth, canSplice, x, r.TypeOf(x))
	}
}

// evalUnquote performs expansion inside a QUASIQUOTE
func (env *Env) evalUnquote(inout UnaryExpr) interface{} {
	block := inout.X.X.(*ast.FuncLit).Body

	ret, extraValues := env.evalBlock(block)
	if len(extraValues) > 1 {
		env.Warnf("unquote returned %d values, only the first one will be used: %v", len(extraValues), block)
	}
	if ret == None || ret == Nil {
		return nil
	}
	return ret.Interface()
}
