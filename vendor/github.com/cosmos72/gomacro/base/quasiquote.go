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

// DuplicateNestedUnquotes is a support function to handle the following complication:
// in Common Lisp, the right-most unquote pairs with the left-most comma!
// we implement the same mechanics, so we must drill down to the last unquote/unquote_splice
// and, for unquote_splice, create a copy of the unquote/unquote_splice stack for each result.
// Example:
//   x:=quote{7; 8}
//   quasiquote{quasiquote{1; unquote{2}; unquote{unquote_splice{x}}}}
// must return
//   quasiquote{1; unquote{2}; unquote{7}; unquote{8}}
func DuplicateNestedUnquotes(src UnaryExpr, depth int, content Ast) Ast {
	if depth == 0 {
		return content
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
	if content != nil {
		tail.Append(content)
	}
	return head
}

func DescendNestedUnquotes(unquote UnaryExpr) (lastUnquote UnaryExpr, depth int) {
	depth = 1
	for {
		form := unquote.Get(0).Get(1)
		form = UnwrapTrivialAst(form)

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
		return unquote, depth
	}
}
