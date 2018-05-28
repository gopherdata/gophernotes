/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
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

package base

import (
	"go/ast"
	"go/token"

	. "github.com/cosmos72/gomacro/ast2"
	mp "github.com/cosmos72/gomacro/parser"
	mt "github.com/cosmos72/gomacro/token"
)

// SimplifyNodeForQuote unwraps ast.BlockStmt, ast.ExprStmt, ast.ParenExpr and ast.DeclStmt
// and returns their contents.
// used to implement classic.Env.evalQuote() and classic.Env.evalQuasiQuote(), be extra careful if you patch it!
func SimplifyNodeForQuote(in ast.Node, unwrapTrivialBlocks bool) ast.Node {
	// unwrap expressions... they fit in more places and make the life easier to MacroExpand and evalQuasiquote
	// also, if unwrapTrivialBlocks is true, unwrap a single-statement block { foo } to foo
	for {
		switch node := in.(type) {
		case *ast.BlockStmt:
			if unwrapTrivialBlocks {
				switch len(node.List) {
				case 0:
					return &ast.EmptyStmt{Semicolon: node.End(), Implicit: false}
				case 1:
					in = node.List[0]
					unwrapTrivialBlocks = false
					continue
				}
			}
			return node
		case *ast.ExprStmt:
			return node.X
		case *ast.ParenExpr:
			return node.X
		case *ast.DeclStmt:
			return node.Decl
		}
		return in
	}
}

// SimplifyAstForQuote unwraps ast2.BlockStmt, ast2.ExprStmt, ast2.ParenExpr and ast2.DeclStmt
// and returns their contents.
// used to implement fast.Comp.QuasiQuote(), be extra careful if you patch it!
func SimplifyAstForQuote(in Ast, unwrapTrivialBlocks bool) Ast {
	// unwrap expressions... they fit in more places and make the life easier to MacroExpand and evalQuasiquote
	// also, if unwrapTrivialBlocks is true, unwrap a single-statement block { foo } to foo
	for {
		switch form := in.(type) {
		case BlockStmt:
			if unwrapTrivialBlocks {
				switch form.Size() {
				case 0:
					return EmptyStmt{&ast.EmptyStmt{Semicolon: form.X.List[0].End(), Implicit: false}}
				case 1:
					in = form.Get(0)
					unwrapTrivialBlocks = false
					continue
				}
			}
			return form
		case ExprStmt, ParenExpr, DeclStmt:
			return in.Get(0)
		}
		return in
	}
}

// restricted version of UnwrapTrivialAst
func UnwrapTrivialNode(node ast.Node) ast.Node {
	in := ToAst(node)
	out := unwrapTrivialAst2(in, true)
	return ToNode(out)
}

// unwrapTrivialAst extract the content from ParenExpr, ExprStmt, DeclStmt:
// such nodes are trivial wrappers for their contents
func UnwrapTrivialAst(in Ast) Ast {
	return unwrapTrivialAst2(in, true)
}

func UnwrapTrivialAstKeepBlocks(in Ast) Ast {
	return unwrapTrivialAst2(in, false)
}

func unwrapTrivialAst2(in Ast, unwrapTrivialBlockStmt bool) Ast {
	for {
		switch form := in.(type) {
		case BlockStmt:
			if !unwrapTrivialBlockStmt || form.Size() != 1 {
				return form
			}
			// a one-element block is trivial UNLESS it contains a declaration.
			// reason: the declaration alters its scope with new bindings.
			// unwrapping it would alters the OUTER scope.
			// i.e. { var x = foo() } and var x = foo() give different scopes
			// to the variable 'x' so they are not equivalent.
			//
			// same reasoning for { x := foo() } versus x := foo()
			child := form.Get(0)
			switch child := child.(type) {
			case DeclStmt:
				return in
			case AssignStmt:
				if child.Op() == token.DEFINE {
					return in
				}
			}
			// Debugf("unwrapTrivialAst(block) unwrapping %#v <%T>\n\tto %#v <%T>\n", form.Interface(), form.Interface(), child.Interface(), child.Interface())
			in = child
		case ParenExpr, ExprStmt, DeclStmt:
			child := form.Get(0)
			// Debugf("unwrapTrivialAst(1) unwrapped %#v <%T>\n\tto %#v <%T>\n", form.Interface(), form.Interface(), child.Interface(), child.Interface())
			in = child
		default:
			return in
		}
	}
}

// MakeQuote invokes parser.MakeQuote() and wraps the resulting ast.Node,
// which represents quote{<form>}, into an Ast struct
func MakeQuote(form UnaryExpr) (UnaryExpr, BlockStmt) {
	expr, block := mp.MakeQuote(nil, form.X.Op, form.X.OpPos, nil)
	return UnaryExpr{expr}, BlockStmt{block}
}

// MakeQuote2 invokes parser.MakeQuote() and wraps the resulting ast.Node,
// which represents quote{<form>}, into an Ast struct
func MakeQuote2(form UnaryExpr, toQuote AstWithNode) UnaryExpr {
	var node ast.Node
	if toQuote != nil {
		node = toQuote.Node()
	}
	// Debugf("node   = %#v\n", node)
	// Debugf("form   = %#v\n", form)
	// Debugf("form.X = %#v\n", form.X)
	expr, _ := mp.MakeQuote(nil, form.X.Op, form.X.OpPos, node)
	return UnaryExpr{expr}
}

// MakeNestedQuote invokes parser.MakeQuote() multiple times, passing op=toks[i] at each call
func MakeNestedQuote(form AstWithNode, toks []token.Token, pos []token.Pos) AstWithNode {
	for i := len(toks) - 1; i >= 0; i-- {
		expr, _ := mp.MakeQuote(nil, toks[i], pos[i], form.Node())
		form = UnaryExpr{expr}
	}
	return form
}

// DuplicateNestedUnquotes is a support function to handle the following complication:
// in Common Lisp, the right-most unquote pairs with the left-most comma!
// we implement the same mechanics, so we must drill down to the last unquote/unquote_splice
// and, for unquote_splice, create a copy of the unquote/unquote_splice stack for each result.
// Example:
//   x:=quote{7; 8}
//   quasiquote{quasiquote{1; unquote{2}; unquote{unquote_splice{x}}}}
// must return
//   quasiquote{1; unquote{2}; unquote{7}; unquote{8}}
func DuplicateNestedUnquotes(src UnaryExpr, depth int, toappend Ast) Ast {
	if depth == 0 {
		return toappend
	}
	head, tail := MakeQuote(src)
	var form Ast = src

	for ; depth > 1; depth-- {
		form = form.Get(0).Get(1)
		form = UnwrapTrivialAst(form)

		src = form.(UnaryExpr)
		expr, newTail := MakeQuote(src)
		// cheat: we know that BlockStmt.Append() always returns the receiver unmodified
		tail.Append(expr)
		tail = newTail
	}
	// cheat: we know that BlockStmt.Append() always returns the receiver unmodified
	if toappend != nil {
		tail.Append(toappend)
	}
	return head
}

// return the expression inside nested mt.UNQUOTE and/or mt.UNQUOTE_SPLICE contained in 'unquote'
func DescendNestedUnquotes(unquote UnaryExpr) (lastUnquote UnaryExpr, depth int) {
	depth = 1
	for {
		form := unquote.Get(0).Get(1)
		// do NOT UnwrapTrivialAst(form): we want the BlockStmt

		// Debugf("DescendNestedUnquotes: %v // %T", UnwrapTrivialAst(form).Interface(), UnwrapTrivialAst(form).Interface())

		if form != nil && form.Size() == 1 {
			if block, ok := form.(BlockStmt); ok {
				form = UnwrapTrivialAst(block.Get(0))
				if form != nil && form.Size() == 1 {
					if expr, ok := form.(UnaryExpr); ok {
						if op := expr.Op(); op == mt.UNQUOTE || op == mt.UNQUOTE_SPLICE {
							unquote = expr
							depth++
							continue
						}
					}
				}
			}
		}
		// Debugf("DescendNestedUnquotes: returning depth = %d, %v // %T", depth, unquote.Interface(), unquote.Interface())
		return unquote, depth
	}
}

// return the sequence of nested mt.UNQUOTE and/or mt.UNQUOTE_SPLICE contained in 'unquote'
func CollectNestedUnquotes(unquote UnaryExpr) ([]token.Token, []token.Pos) {
	// Debugf("CollectNestedUnquotes: %v // %T", unquote.X, unquote.X)

	var toks []token.Token
	var pos []token.Pos
	for {
		unary := unquote.X
		toks = append(toks, unary.Op)
		pos = append(pos, unary.OpPos)

		form := unquote.Get(0).Get(1)
		// do NOT UnwrapTrivialAst(form): we want the BlockStmt

		if form != nil && form.Size() == 1 {
			if block, ok := form.(BlockStmt); ok {
				form = UnwrapTrivialAst(block.Get(0))
				if form != nil && form.Size() == 1 {
					if expr, ok := form.(UnaryExpr); ok {
						if op := expr.X.Op; op == mt.UNQUOTE || op == mt.UNQUOTE_SPLICE {
							unquote = expr
							continue
						}
					}
				}
			}
		}
		// Debugf("CollectNestedUnquotes: returning toks = %v, pos = %v // %T", toks, pos)
		return toks, pos
	}
}
