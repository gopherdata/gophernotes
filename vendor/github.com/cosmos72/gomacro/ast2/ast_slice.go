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
 * ast_slice.go
 *
 *  Created on Feb 25, 2017
 *      Author Massimiliano Ghilardi
 */

package ast2

import (
	"go/ast"
	"go/token"
)

// Ast wrappers for variable-length slices of ast.Nodes - they are not full-blown ast.Node

func (x AstSlice) Interface() interface{}   { return asInterface(x.X, x.X == nil) }
func (x NodeSlice) Interface() interface{}  { return asInterface(x.X, x.X == nil) }
func (x ExprSlice) Interface() interface{}  { return asInterface(x.X, x.X == nil) }
func (x FieldSlice) Interface() interface{} { return asInterface(x.X, x.X == nil) }
func (x DeclSlice) Interface() interface{}  { return asInterface(x.X, x.X == nil) }
func (x IdentSlice) Interface() interface{} { return asInterface(x.X, x.X == nil) }
func (x SpecSlice) Interface() interface{}  { return asInterface(x.X, x.X == nil) }
func (x StmtSlice) Interface() interface{}  { return asInterface(x.X, x.X == nil) }

func (x AstSlice) Op() token.Token   { return token.COMMA }     // FIXME
func (x NodeSlice) Op() token.Token  { return token.COMMA }     // FIXME
func (x ExprSlice) Op() token.Token  { return token.COMMA }     // FIXME
func (x FieldSlice) Op() token.Token { return token.SEMICOLON } // FIXME
func (x DeclSlice) Op() token.Token  { return token.SEMICOLON } // FIXME
func (x IdentSlice) Op() token.Token { return token.COMMA }     // FIXME
func (x SpecSlice) Op() token.Token  { return token.SEMICOLON } // FIXME
func (x StmtSlice) Op() token.Token  { return token.SEMICOLON } // FIXME

func (x AstSlice) New() Ast   { return AstSlice{} }
func (x NodeSlice) New() Ast  { return NodeSlice{} }
func (x ExprSlice) New() Ast  { return ExprSlice{} }
func (x FieldSlice) New() Ast { return FieldSlice{} }
func (x DeclSlice) New() Ast  { return DeclSlice{} }
func (x IdentSlice) New() Ast { return IdentSlice{} }
func (x SpecSlice) New() Ast  { return SpecSlice{} }
func (x StmtSlice) New() Ast  { return StmtSlice{} }

func (x AstSlice) Size() int   { return len(x.X) }
func (x NodeSlice) Size() int  { return len(x.X) }
func (x ExprSlice) Size() int  { return len(x.X) }
func (x FieldSlice) Size() int { return len(x.X) }
func (x DeclSlice) Size() int  { return len(x.X) }
func (x IdentSlice) Size() int { return len(x.X) }
func (x SpecSlice) Size() int  { return len(x.X) }
func (x StmtSlice) Size() int  { return len(x.X) }

func (x AstSlice) Get(i int) Ast   { return x.X[i] }
func (x NodeSlice) Get(i int) Ast  { return ToAst(x.X[i]) }
func (x ExprSlice) Get(i int) Ast  { return ToAst(x.X[i]) }
func (x FieldSlice) Get(i int) Ast { return ToAst(x.X[i]) }
func (x DeclSlice) Get(i int) Ast  { return ToAst(x.X[i]) }
func (x IdentSlice) Get(i int) Ast { return ToAst(x.X[i]) }
func (x SpecSlice) Get(i int) Ast  { return ToAst(x.X[i]) }
func (x StmtSlice) Get(i int) Ast  { return ToAst(x.X[i]) }

func (x AstSlice) Set(i int, child Ast)   { x.X[i] = child }
func (x NodeSlice) Set(i int, child Ast)  { x.X[i] = ToNode(child) }
func (x ExprSlice) Set(i int, child Ast)  { x.X[i] = ToExpr(child) }
func (x FieldSlice) Set(i int, child Ast) { x.X[i] = ToField(child) }
func (x DeclSlice) Set(i int, child Ast)  { x.X[i] = ToDecl(child) }
func (x IdentSlice) Set(i int, child Ast) { x.X[i] = ToIdent(child) }
func (x SpecSlice) Set(i int, child Ast)  { x.X[i] = ToSpec(child) }
func (x StmtSlice) Set(i int, child Ast)  { x.X[i] = ToStmt(child) }

func (x AstSlice) Slice(lo, hi int) AstWithSlice   { x.X = x.X[lo:hi]; return x }
func (x NodeSlice) Slice(lo, hi int) AstWithSlice  { x.X = x.X[lo:hi]; return x }
func (x ExprSlice) Slice(lo, hi int) AstWithSlice  { x.X = x.X[lo:hi]; return x }
func (x FieldSlice) Slice(lo, hi int) AstWithSlice { x.X = x.X[lo:hi]; return x }
func (x DeclSlice) Slice(lo, hi int) AstWithSlice  { x.X = x.X[lo:hi]; return x }
func (x IdentSlice) Slice(lo, hi int) AstWithSlice { x.X = x.X[lo:hi]; return x }
func (x SpecSlice) Slice(lo, hi int) AstWithSlice  { x.X = x.X[lo:hi]; return x }
func (x StmtSlice) Slice(lo, hi int) AstWithSlice  { x.X = x.X[lo:hi]; return x }

func (x AstSlice) Append(child Ast) AstWithSlice   { x.X = append(x.X, child); return x }
func (x NodeSlice) Append(child Ast) AstWithSlice  { x.X = append(x.X, ToNode(child)); return x }
func (x ExprSlice) Append(child Ast) AstWithSlice  { x.X = append(x.X, ToExpr(child)); return x }
func (x FieldSlice) Append(child Ast) AstWithSlice { x.X = append(x.X, ToField(child)); return x }
func (x DeclSlice) Append(child Ast) AstWithSlice  { x.X = append(x.X, ToDecl(child)); return x }
func (x IdentSlice) Append(child Ast) AstWithSlice { x.X = append(x.X, ToIdent(child)); return x }
func (x SpecSlice) Append(child Ast) AstWithSlice  { x.X = append(x.X, ToSpec(child)); return x }
func (x StmtSlice) Append(child Ast) AstWithSlice  { x.X = append(x.X, ToStmt(child)); return x }

// variable-length ast.Nodes

func (x BlockStmt) Interface() interface{}  { return asInterface(x.X, x.X == nil) }
func (x FieldList) Interface() interface{}  { return asInterface(x.X, x.X == nil) }
func (x File) Interface() interface{}       { return asInterface(x.X, x.X == nil) }
func (x GenDecl) Interface() interface{}    { return asInterface(x.X, x.X == nil) }
func (x ReturnStmt) Interface() interface{} { return asInterface(x.X, x.X == nil) }

func (x BlockStmt) Node() ast.Node  { return asNode(x.X, x.X == nil) }
func (x FieldList) Node() ast.Node  { return asNode(x.X, x.X == nil) }
func (x File) Node() ast.Node       { return asNode(x.X, x.X == nil) }
func (x GenDecl) Node() ast.Node    { return asNode(x.X, x.X == nil) }
func (x ReturnStmt) Node() ast.Node { return asNode(x.X, x.X == nil) }

func (x BlockStmt) Op() token.Token  { return token.LBRACE }
func (x FieldList) Op() token.Token  { return token.ELLIPSIS }
func (x File) Op() token.Token       { return token.EOF }
func (x GenDecl) Op() token.Token    { return x.X.Tok }
func (x ReturnStmt) Op() token.Token { return token.RETURN }

func (x BlockStmt) New() Ast { return BlockStmt{&ast.BlockStmt{Lbrace: x.X.Lbrace, Rbrace: x.X.Rbrace}} }
func (x FieldList) New() Ast {
	return FieldList{&ast.FieldList{Opening: x.X.Opening, Closing: x.X.Closing}}
}
func (x File) New() Ast {
	return File{&ast.File{Doc: x.X.Doc, Package: x.X.Package, Name: x.X.Name, Scope: x.X.Scope, Imports: x.X.Imports, Comments: x.X.Comments}}
}
func (x GenDecl) New() Ast {
	return GenDecl{&ast.GenDecl{Doc: x.X.Doc, TokPos: x.X.TokPos, Tok: x.X.Tok, Lparen: x.X.Lparen, Rparen: x.X.Rparen}}
}

// do not copy position of "return" keyword.
// otherwise go/format may insert a newline between "return" and the following expressions
func (x ReturnStmt) New() Ast { return ReturnStmt{&ast.ReturnStmt{}} }

func (x BlockStmt) Size() int  { return len(x.X.List) }
func (x FieldList) Size() int  { return len(x.X.List) }
func (x File) Size() int       { return len(x.X.Decls) }
func (x GenDecl) Size() int    { return len(x.X.Specs) }
func (x ReturnStmt) Size() int { return len(x.X.Results) }

func (x BlockStmt) Get(i int) Ast  { return ToAst(x.X.List[i]) }
func (x FieldList) Get(i int) Ast  { return ToAst(x.X.List[i]) }
func (x File) Get(i int) Ast       { return ToAst(x.X.Decls[i]) }
func (x GenDecl) Get(i int) Ast    { return ToAst(x.X.Specs[i]) }
func (x ReturnStmt) Get(i int) Ast { return ToAst(x.X.Results[i]) }

func (x BlockStmt) Set(i int, child Ast)  { x.X.List[i] = ToStmt(child) }
func (x FieldList) Set(i int, child Ast)  { x.X.List[i] = ToField(child) }
func (x File) Set(i int, child Ast)       { x.X.Decls[i] = ToDecl(child) }
func (x GenDecl) Set(i int, child Ast)    { x.X.Specs[i] = ToSpec(child) }
func (x ReturnStmt) Set(i int, child Ast) { x.X.Results[i] = ToExpr(child) }

func (x BlockStmt) Slice(lo, hi int) AstWithSlice  { x.X.List = x.X.List[lo:hi]; return x }
func (x FieldList) Slice(lo, hi int) AstWithSlice  { x.X.List = x.X.List[lo:hi]; return x }
func (x File) Slice(lo, hi int) AstWithSlice       { x.X.Decls = x.X.Decls[lo:hi]; return x }
func (x GenDecl) Slice(lo, hi int) AstWithSlice    { x.X.Specs = x.X.Specs[lo:hi]; return x }
func (x ReturnStmt) Slice(lo, hi int) AstWithSlice { x.X.Results = x.X.Results[lo:hi]; return x }

func (x BlockStmt) Append(child Ast) AstWithSlice {
	x.X.List = append(x.X.List, ToStmt(child))
	return x
}
func (x FieldList) Append(child Ast) AstWithSlice {
	x.X.List = append(x.X.List, ToField(child))
	return x
}
func (x File) Append(child Ast) AstWithSlice {
	x.X.Decls = append(x.X.Decls, ToDecl(child))
	return x
}
func (x GenDecl) Append(child Ast) AstWithSlice {
	x.X.Specs = append(x.X.Specs, ToSpec(child))
	return x
}
func (x ReturnStmt) Append(child Ast) AstWithSlice {
	x.X.Results = append(x.X.Results, ToExpr(child))
	return x
}
