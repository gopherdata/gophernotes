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
 * ast_node.go
 *
 *  Created on Feb 25, 2017
 *      Author Massimiliano Ghilardi
 */

package ast2

import (
	"go/ast"
	"go/token"
)

func asInterface(x interface{}, isnil bool) interface{} {
	if isnil {
		return nil
	}
	return x
}

func asNode(x ast.Node, isnil bool) ast.Node {
	if isnil {
		return nil
	}
	return x
}

//
// .................. functions Interface() interface{}
//
func (x ArrayType) Interface() interface{}      { return asInterface(x.X, x.X == nil) }
func (x AssignStmt) Interface() interface{}     { return asInterface(x.X, x.X == nil) }
func (x BadDecl) Interface() interface{}        { return asInterface(x.X, x.X == nil) }
func (x BadExpr) Interface() interface{}        { return asInterface(x.X, x.X == nil) }
func (x BadStmt) Interface() interface{}        { return asInterface(x.X, x.X == nil) }
func (x BasicLit) Interface() interface{}       { return asInterface(x.X, x.X == nil) }
func (x BinaryExpr) Interface() interface{}     { return asInterface(x.X, x.X == nil) }
func (x BranchStmt) Interface() interface{}     { return asInterface(x.X, x.X == nil) }
func (x CallExpr) Interface() interface{}       { return asInterface(x.X, x.X == nil) }
func (x CaseClause) Interface() interface{}     { return asInterface(x.X, x.X == nil) }
func (x ChanType) Interface() interface{}       { return asInterface(x.X, x.X == nil) }
func (x CommClause) Interface() interface{}     { return asInterface(x.X, x.X == nil) }
func (x CompositeLit) Interface() interface{}   { return asInterface(x.X, x.X == nil) }
func (x DeclStmt) Interface() interface{}       { return asInterface(x.X, x.X == nil) }
func (x DeferStmt) Interface() interface{}      { return asInterface(x.X, x.X == nil) }
func (x Ellipsis) Interface() interface{}       { return asInterface(x.X, x.X == nil) }
func (x EmptyStmt) Interface() interface{}      { return asInterface(x.X, x.X == nil) }
func (x ExprStmt) Interface() interface{}       { return asInterface(x.X, x.X == nil) }
func (x Field) Interface() interface{}          { return asInterface(x.X, x.X == nil) }
func (x ForStmt) Interface() interface{}        { return asInterface(x.X, x.X == nil) }
func (x FuncDecl) Interface() interface{}       { return asInterface(x.X, x.X == nil) }
func (x FuncLit) Interface() interface{}        { return asInterface(x.X, x.X == nil) }
func (x FuncType) Interface() interface{}       { return asInterface(x.X, x.X == nil) }
func (x GoStmt) Interface() interface{}         { return asInterface(x.X, x.X == nil) }
func (x Ident) Interface() interface{}          { return asInterface(x.X, x.X == nil) }
func (x IfStmt) Interface() interface{}         { return asInterface(x.X, x.X == nil) }
func (x ImportSpec) Interface() interface{}     { return asInterface(x.X, x.X == nil) }
func (x IncDecStmt) Interface() interface{}     { return asInterface(x.X, x.X == nil) }
func (x IndexExpr) Interface() interface{}      { return asInterface(x.X, x.X == nil) }
func (x InterfaceType) Interface() interface{}  { return asInterface(x.X, x.X == nil) }
func (x KeyValueExpr) Interface() interface{}   { return asInterface(x.X, x.X == nil) }
func (x LabeledStmt) Interface() interface{}    { return asInterface(x.X, x.X == nil) }
func (x MapType) Interface() interface{}        { return asInterface(x.X, x.X == nil) }
func (x Package) Interface() interface{}        { return asInterface(x.X, x.X == nil) }
func (x ParenExpr) Interface() interface{}      { return asInterface(x.X, x.X == nil) }
func (x RangeStmt) Interface() interface{}      { return asInterface(x.X, x.X == nil) }
func (x SelectStmt) Interface() interface{}     { return asInterface(x.X, x.X == nil) }
func (x SelectorExpr) Interface() interface{}   { return asInterface(x.X, x.X == nil) }
func (x SendStmt) Interface() interface{}       { return asInterface(x.X, x.X == nil) }
func (x SliceExpr) Interface() interface{}      { return asInterface(x.X, x.X == nil) }
func (x StarExpr) Interface() interface{}       { return asInterface(x.X, x.X == nil) }
func (x StructType) Interface() interface{}     { return asInterface(x.X, x.X == nil) }
func (x SwitchStmt) Interface() interface{}     { return asInterface(x.X, x.X == nil) }
func (x TypeAssertExpr) Interface() interface{} { return asInterface(x.X, x.X == nil) }
func (x TypeSpec) Interface() interface{}       { return asInterface(x.X, x.X == nil) }
func (x TypeSwitchStmt) Interface() interface{} { return asInterface(x.X, x.X == nil) }
func (x UnaryExpr) Interface() interface{}      { return asInterface(x.X, x.X == nil) }
func (x ValueSpec) Interface() interface{}      { return asInterface(x.X, x.X == nil) }

//
// .................. functions Node() ast.Node
//
func (x ArrayType) Node() ast.Node      { return asNode(x.X, x.X == nil) }
func (x AssignStmt) Node() ast.Node     { return asNode(x.X, x.X == nil) }
func (x BadDecl) Node() ast.Node        { return asNode(x.X, x.X == nil) }
func (x BadExpr) Node() ast.Node        { return asNode(x.X, x.X == nil) }
func (x BadStmt) Node() ast.Node        { return asNode(x.X, x.X == nil) }
func (x BasicLit) Node() ast.Node       { return asNode(x.X, x.X == nil) }
func (x BinaryExpr) Node() ast.Node     { return asNode(x.X, x.X == nil) }
func (x BranchStmt) Node() ast.Node     { return asNode(x.X, x.X == nil) }
func (x CallExpr) Node() ast.Node       { return asNode(x.X, x.X == nil) }
func (x CaseClause) Node() ast.Node     { return asNode(x.X, x.X == nil) }
func (x ChanType) Node() ast.Node       { return asNode(x.X, x.X == nil) }
func (x CommClause) Node() ast.Node     { return asNode(x.X, x.X == nil) }
func (x CompositeLit) Node() ast.Node   { return asNode(x.X, x.X == nil) }
func (x DeclStmt) Node() ast.Node       { return asNode(x.X, x.X == nil) }
func (x DeferStmt) Node() ast.Node      { return asNode(x.X, x.X == nil) }
func (x Ellipsis) Node() ast.Node       { return asNode(x.X, x.X == nil) }
func (x EmptyStmt) Node() ast.Node      { return asNode(x.X, x.X == nil) }
func (x ExprStmt) Node() ast.Node       { return asNode(x.X, x.X == nil) }
func (x Field) Node() ast.Node          { return asNode(x.X, x.X == nil) }
func (x ForStmt) Node() ast.Node        { return asNode(x.X, x.X == nil) }
func (x FuncDecl) Node() ast.Node       { return asNode(x.X, x.X == nil) }
func (x FuncLit) Node() ast.Node        { return asNode(x.X, x.X == nil) }
func (x FuncType) Node() ast.Node       { return asNode(x.X, x.X == nil) }
func (x GoStmt) Node() ast.Node         { return asNode(x.X, x.X == nil) }
func (x Ident) Node() ast.Node          { return asNode(x.X, x.X == nil) }
func (x IfStmt) Node() ast.Node         { return asNode(x.X, x.X == nil) }
func (x ImportSpec) Node() ast.Node     { return asNode(x.X, x.X == nil) }
func (x IncDecStmt) Node() ast.Node     { return asNode(x.X, x.X == nil) }
func (x IndexExpr) Node() ast.Node      { return asNode(x.X, x.X == nil) }
func (x InterfaceType) Node() ast.Node  { return asNode(x.X, x.X == nil) }
func (x KeyValueExpr) Node() ast.Node   { return asNode(x.X, x.X == nil) }
func (x LabeledStmt) Node() ast.Node    { return asNode(x.X, x.X == nil) }
func (x MapType) Node() ast.Node        { return asNode(x.X, x.X == nil) }
func (x Package) Node() ast.Node        { return asNode(x.X, x.X == nil) }
func (x ParenExpr) Node() ast.Node      { return asNode(x.X, x.X == nil) }
func (x RangeStmt) Node() ast.Node      { return asNode(x.X, x.X == nil) }
func (x SelectStmt) Node() ast.Node     { return asNode(x.X, x.X == nil) }
func (x SelectorExpr) Node() ast.Node   { return asNode(x.X, x.X == nil) }
func (x SendStmt) Node() ast.Node       { return asNode(x.X, x.X == nil) }
func (x SliceExpr) Node() ast.Node      { return asNode(x.X, x.X == nil) }
func (x StarExpr) Node() ast.Node       { return asNode(x.X, x.X == nil) }
func (x StructType) Node() ast.Node     { return asNode(x.X, x.X == nil) }
func (x SwitchStmt) Node() ast.Node     { return asNode(x.X, x.X == nil) }
func (x TypeAssertExpr) Node() ast.Node { return asNode(x.X, x.X == nil) }
func (x TypeSpec) Node() ast.Node       { return asNode(x.X, x.X == nil) }
func (x TypeSwitchStmt) Node() ast.Node { return asNode(x.X, x.X == nil) }
func (x UnaryExpr) Node() ast.Node      { return asNode(x.X, x.X == nil) }
func (x ValueSpec) Node() ast.Node      { return asNode(x.X, x.X == nil) }

//
// .................. functions Op() token.Token
//
func (x ArrayType) Op() token.Token  { return token.LBRACK }
func (x AssignStmt) Op() token.Token { return x.X.Tok }
func (x BadDecl) Op() token.Token    { return token.ILLEGAL }
func (x BadExpr) Op() token.Token    { return token.ILLEGAL }
func (x BadStmt) Op() token.Token    { return token.ILLEGAL }
func (x BasicLit) Op() token.Token   { return x.X.Kind }
func (x BinaryExpr) Op() token.Token { return x.X.Op }
func (x BranchStmt) Op() token.Token { return x.X.Tok }
func (x CallExpr) Op() token.Token   { return token.RPAREN }
func (x CaseClause) Op() token.Token {
	if len(x.X.List) != 0 {
		return token.CASE
	} else {
		return token.DEFAULT
	}
}
func (x ChanType) Op() token.Token { return token.CHAN }
func (x CommClause) Op() token.Token {
	if x.X.Comm != nil {
		return token.CASE
	} else {
		return token.DEFAULT
	}
}
func (x CompositeLit) Op() token.Token   { return token.RBRACE }
func (x DeclStmt) Op() token.Token       { return x.X.Decl.(*ast.GenDecl).Tok }
func (x DeferStmt) Op() token.Token      { return token.DEFER }
func (x Ellipsis) Op() token.Token       { return token.ELLIPSIS }
func (x EmptyStmt) Op() token.Token      { return token.SEMICOLON }
func (x ExprStmt) Op() token.Token       { return token.ELSE } // FIXME
func (x Field) Op() token.Token          { return token.PERIOD }
func (x ForStmt) Op() token.Token        { return token.FOR }
func (x FuncDecl) Op() token.Token       { return token.FUNC }
func (x FuncLit) Op() token.Token        { return token.FUNC }
func (x FuncType) Op() token.Token       { return token.FUNC }
func (x GoStmt) Op() token.Token         { return token.GO }
func (x Ident) Op() token.Token          { return token.IDENT }
func (x IfStmt) Op() token.Token         { return token.IF }
func (x ImportSpec) Op() token.Token     { return token.IMPORT }
func (x IncDecStmt) Op() token.Token     { return x.X.Tok }
func (x IndexExpr) Op() token.Token      { return token.LBRACK }
func (x InterfaceType) Op() token.Token  { return token.INTERFACE }
func (x KeyValueExpr) Op() token.Token   { return token.COLON } // FIXME
func (x LabeledStmt) Op() token.Token    { return token.COLON } // FIXME
func (x MapType) Op() token.Token        { return token.MAP }
func (x Package) Op() token.Token        { return token.PACKAGE }
func (x ParenExpr) Op() token.Token      { return token.LPAREN }
func (x RangeStmt) Op() token.Token      { return token.RANGE }
func (x SelectStmt) Op() token.Token     { return token.SELECT }
func (x SelectorExpr) Op() token.Token   { return token.CASE }
func (x SendStmt) Op() token.Token       { return token.CHAN }   // FIXME
func (x SliceExpr) Op() token.Token      { return token.RBRACK } // FIXME
func (x StarExpr) Op() token.Token       { return token.MUL }
func (x StructType) Op() token.Token     { return token.STRUCT }
func (x SwitchStmt) Op() token.Token     { return token.SWITCH }
func (x TypeAssertExpr) Op() token.Token { return token.TYPE } // FIXME
func (x TypeSpec) Op() token.Token       { return token.TYPE }
func (x TypeSwitchStmt) Op() token.Token { return token.SWITCH } // FIXME
func (x UnaryExpr) Op() token.Token      { return x.X.Op }
func (x ValueSpec) Op() token.Token      { return token.VAR } // can be VAR or CONST

//
// .................. functions New() Ast
//
func (x ArrayType) New() Ast  { return ArrayType{&ast.ArrayType{Lbrack: x.X.Lbrack}} }
func (x AssignStmt) New() Ast { return AssignStmt{&ast.AssignStmt{TokPos: x.X.TokPos, Tok: x.X.Tok}} }
func (x BadDecl) New() Ast    { return BadDecl{&ast.BadDecl{From: x.X.From, To: x.X.To}} }
func (x BadExpr) New() Ast    { return BadExpr{&ast.BadExpr{From: x.X.From, To: x.X.To}} }
func (x BadStmt) New() Ast    { return BadStmt{&ast.BadStmt{From: x.X.From, To: x.X.To}} }
func (x BasicLit) New() Ast {
	return BasicLit{&ast.BasicLit{ValuePos: x.X.ValuePos, Value: x.X.Value, Kind: x.X.Kind}}
}
func (x BinaryExpr) New() Ast { return BinaryExpr{&ast.BinaryExpr{OpPos: x.X.OpPos, Op: x.X.Op}} }
func (x BranchStmt) New() Ast { return BranchStmt{&ast.BranchStmt{TokPos: x.X.TokPos, Tok: x.X.Tok}} }
func (x CallExpr) New() Ast {
	return CallExpr{&ast.CallExpr{Lparen: x.X.Lparen, Ellipsis: x.X.Ellipsis, Rparen: x.X.Rparen}}
}
func (x CaseClause) New() Ast { return CaseClause{&ast.CaseClause{Case: x.X.Case, Colon: x.X.Colon}} }
func (x ChanType) New() Ast {
	return ChanType{&ast.ChanType{Begin: x.X.Begin, Arrow: x.X.Arrow, Dir: x.X.Dir}}
}
func (x CommClause) New() Ast { return CommClause{&ast.CommClause{Case: x.X.Case, Colon: x.X.Colon}} }
func (x CompositeLit) New() Ast {
	return CompositeLit{&ast.CompositeLit{Lbrace: x.X.Lbrace, Rbrace: x.X.Rbrace}}
}
func (x DeclStmt) New() Ast  { return DeclStmt{&ast.DeclStmt{}} }
func (x DeferStmt) New() Ast { return DeferStmt{&ast.DeferStmt{Defer: x.X.Defer}} }
func (x Ellipsis) New() Ast  { return Ellipsis{&ast.Ellipsis{Ellipsis: x.X.Ellipsis}} }
func (x EmptyStmt) New() Ast {
	return EmptyStmt{&ast.EmptyStmt{Semicolon: x.X.Semicolon, Implicit: x.X.Implicit}}
}
func (x ExprStmt) New() Ast { return ExprStmt{&ast.ExprStmt{}} }
func (x Field) New() Ast    { return Field{&ast.Field{Doc: x.X.Doc, Comment: x.X.Comment}} }
func (x ForStmt) New() Ast  { return ForStmt{&ast.ForStmt{For: x.X.For}} }
func (x FuncDecl) New() Ast { return FuncDecl{&ast.FuncDecl{Doc: x.X.Doc}} }
func (x FuncLit) New() Ast  { return FuncLit{&ast.FuncLit{}} }
func (x FuncType) New() Ast { return FuncType{&ast.FuncType{Func: x.X.Func}} }
func (x GoStmt) New() Ast   { return GoStmt{&ast.GoStmt{Go: x.X.Go}} }
func (x Ident) New() Ast    { return Ident{&ast.Ident{NamePos: x.X.NamePos, Name: x.X.Name}} }
func (x IfStmt) New() Ast   { return IfStmt{&ast.IfStmt{If: x.X.If}} }
func (x ImportSpec) New() Ast {
	return ImportSpec{&ast.ImportSpec{Doc: x.X.Doc, Comment: x.X.Comment, EndPos: x.X.EndPos}}
}
func (x IncDecStmt) New() Ast { return IncDecStmt{&ast.IncDecStmt{TokPos: x.X.TokPos, Tok: x.X.Tok}} }
func (x IndexExpr) New() Ast  { return IndexExpr{&ast.IndexExpr{Lbrack: x.X.Lbrack, Rbrack: x.X.Rbrack}} }
func (x InterfaceType) New() Ast {
	return InterfaceType{&ast.InterfaceType{Interface: x.X.Interface, Incomplete: x.X.Incomplete}}
}
func (x KeyValueExpr) New() Ast { return KeyValueExpr{&ast.KeyValueExpr{Colon: x.X.Colon}} }
func (x LabeledStmt) New() Ast  { return LabeledStmt{&ast.LabeledStmt{Colon: x.X.Colon}} }
func (x MapType) New() Ast      { return MapType{&ast.MapType{Map: x.X.Map}} }
func (x Package) New() Ast {
	return Package{&ast.Package{Name: x.X.Name, Scope: x.X.Scope, Imports: x.X.Imports}}
}
func (x ParenExpr) New() Ast { return ParenExpr{&ast.ParenExpr{Lparen: x.X.Lparen, Rparen: x.X.Rparen}} }
func (x RangeStmt) New() Ast {
	return RangeStmt{&ast.RangeStmt{For: x.X.For, TokPos: x.X.TokPos, Tok: x.X.Tok}}
}
func (x SelectStmt) New() Ast   { return SelectStmt{&ast.SelectStmt{Select: x.X.Select}} }
func (x SelectorExpr) New() Ast { return SelectorExpr{&ast.SelectorExpr{}} }
func (x SendStmt) New() Ast     { return SendStmt{&ast.SendStmt{Arrow: x.X.Arrow}} }
func (x SliceExpr) New() Ast    { return SliceExpr{&ast.SliceExpr{Lbrack: x.X.Lbrack, Rbrack: x.X.Rbrack}} }
func (x StarExpr) New() Ast     { return StarExpr{&ast.StarExpr{Star: x.X.Star}} }
func (x StructType) New() Ast   { return StructType{&ast.StructType{Incomplete: x.X.Incomplete}} }
func (x SwitchStmt) New() Ast   { return SwitchStmt{&ast.SwitchStmt{Switch: x.X.Switch}} }
func (x TypeAssertExpr) New() Ast {
	return TypeAssertExpr{&ast.TypeAssertExpr{Lparen: x.X.Lparen, Rparen: x.X.Rparen}}
}
func (x TypeSpec) New() Ast {
	return TypeSpec{&ast.TypeSpec{Doc: x.X.Doc, Assign: x.X.Assign, Comment: x.X.Comment}}
}
func (x TypeSwitchStmt) New() Ast { return TypeSwitchStmt{&ast.TypeSwitchStmt{Switch: x.X.Switch}} }
func (x UnaryExpr) New() Ast      { return UnaryExpr{&ast.UnaryExpr{OpPos: x.X.OpPos, Op: x.X.Op}} }
func (x ValueSpec) New() Ast      { return ValueSpec{&ast.ValueSpec{Doc: x.X.Doc, Comment: x.X.Comment}} }

//
// .................. functions Size() int
//
func (x ArrayType) Size() int    { return 2 }
func (x AssignStmt) Size() int   { return 2 }
func (x BadDecl) Size() int      { return 0 }
func (x BadExpr) Size() int      { return 0 }
func (x BadStmt) Size() int      { return 0 }
func (x BasicLit) Size() int     { return 0 }
func (x BinaryExpr) Size() int   { return 2 }
func (x BranchStmt) Size() int   { return 1 }
func (x CallExpr) Size() int     { return 2 }
func (x CaseClause) Size() int   { return 2 }
func (x ChanType) Size() int     { return 1 }
func (x CommClause) Size() int   { return 2 }
func (x CompositeLit) Size() int { return 2 }
func (x DeclStmt) Size() int     { return 1 }
func (x DeferStmt) Size() int    { return 1 }
func (x Ellipsis) Size() int     { return 1 }
func (x EmptyStmt) Size() int    { return 0 }
func (x ExprStmt) Size() int     { return 1 }
func (x Field) Size() int {
	// do not crash on nil *ast.Field as first receiver of generic functions
	if x.X == nil {
		return 0
	}
	return 3
}
func (x ForStmt) Size() int        { return 4 }
func (x FuncDecl) Size() int       { return 4 }
func (x FuncLit) Size() int        { return 2 }
func (x FuncType) Size() int       { return 2 }
func (x GoStmt) Size() int         { return 1 }
func (x Ident) Size() int          { return 0 }
func (x IfStmt) Size() int         { return 4 }
func (x ImportSpec) Size() int     { return 2 }
func (x IncDecStmt) Size() int     { return 1 }
func (x IndexExpr) Size() int      { return 2 }
func (x InterfaceType) Size() int  { return 1 }
func (x KeyValueExpr) Size() int   { return 2 }
func (x LabeledStmt) Size() int    { return 2 }
func (x MapType) Size() int        { return 2 }
func (x Package) Size() int        { return 2 }
func (x ParenExpr) Size() int      { return 1 }
func (x RangeStmt) Size() int      { return 4 }
func (x SelectStmt) Size() int     { return 1 }
func (x SelectorExpr) Size() int   { return 2 }
func (x SendStmt) Size() int       { return 2 }
func (x SliceExpr) Size() int      { return 4 }
func (x StarExpr) Size() int       { return 1 }
func (x StructType) Size() int     { return 1 }
func (x SwitchStmt) Size() int     { return 3 }
func (x TypeAssertExpr) Size() int { return 2 }
func (x TypeSpec) Size() int       { return 2 }
func (x TypeSwitchStmt) Size() int { return 3 }
func (x UnaryExpr) Size() int      { return 1 }
func (x ValueSpec) Size() int      { return 3 }

//
// .................. functions Get(int) Ast
//
func (x ArrayType) Get(i int) Ast { return ToAst2(i, x.X.Len, x.X.Elt) }
func (x AssignStmt) Get(i int) Ast {
	var slice []ast.Expr
	switch i {
	case 0:
		slice = x.X.Lhs
	case 1:
		slice = x.X.Rhs
	default:
		return badIndex(i, 2)
	}
	if slice != nil {
		return ExprSlice{slice}
	}
	return nil
}
func (x BadDecl) Get(i int) Ast    { return badIndex(i, 0) }
func (x BadExpr) Get(i int) Ast    { return badIndex(i, 0) }
func (x BadStmt) Get(i int) Ast    { return badIndex(i, 0) }
func (x BasicLit) Get(i int) Ast   { return badIndex(i, 0) }
func (x BinaryExpr) Get(i int) Ast { return ToAst2(i, x.X.X, x.X.Y) }
func (x BranchStmt) Get(i int) Ast { return Ident{x.X.Label} }
func (x CallExpr) Get(i int) Ast {
	if i == 0 {
		return ToAst(x.X.Fun)
	} else if i == 1 {
		if node := x.X.Args; node != nil {
			return ExprSlice{node}
		}
		return nil
	} else {
		return badIndex(i, 2)
	}
}
func (x CaseClause) Get(i int) Ast {
	if i == 0 {
		if list := x.X.List; list != nil {
			return ExprSlice{list}
		}
		return nil
	} else if i == 1 {
		if list := x.X.Body; list != nil {
			return StmtSlice{list}
		}
		return nil
	} else {
		return badIndex(i, 2)
	}
}
func (x ChanType) Get(i int) Ast { return ToAst1(i, x.X.Value) }
func (x CommClause) Get(i int) Ast {
	if i == 0 {
		return ToAst(x.X.Comm)
	} else if i == 1 {
		if list := x.X.Body; list != nil {
			return StmtSlice{list}
		}
		return nil
	} else {
		return badIndex(i, 2)
	}
}
func (x CompositeLit) Get(i int) Ast {
	if i == 0 {
		return ToAst(x.X.Type)
	} else if i == 1 {
		if x.X.Elts != nil {
			return ExprSlice{x.X.Elts}
		}
		return nil
	} else {
		return badIndex(i, 2)
	}
}
func (x DeclStmt) Get(i int) Ast  { return ToAst1(i, x.X.Decl) }
func (x DeferStmt) Get(i int) Ast { return CallExpr{x.X.Call} }
func (x Ellipsis) Get(i int) Ast  { return ToAst1(i, x.X.Elt) }
func (x EmptyStmt) Get(i int) Ast { return badIndex(i, 0) }
func (x ExprStmt) Get(i int) Ast  { return ToAst1(i, x.X.X) }
func (x Field) Get(i int) Ast {
	if i == 0 {
		if x.X.Names != nil {
			return IdentSlice{x.X.Names}
		}
		return nil
	} else if i == 1 {
		return ToAst(x.X.Type)
	} else if i == 2 {
		return ToAst(x.X.Tag)
	} else {
		return badIndex(i, 3)
	}
}
func (x ForStmt) Get(i int) Ast {
	var node ast.Node
	switch i {
	case 0:
		node = x.X.Init
	case 1:
		node = x.X.Cond
	case 2:
		node = x.X.Post
	case 3:
		node = x.X.Body
	default:
		return badIndex(i, 4)
	}
	return ToAst(node)
}
func (x FuncDecl) Get(i int) Ast   { return ToAst4(i, x.X.Recv, x.X.Name, x.X.Type, x.X.Body) }
func (x FuncLit) Get(i int) Ast    { return ToAst2(i, x.X.Type, x.X.Body) }
func (x FuncType) Get(i int) Ast   { return ToAst2(i, x.X.Params, x.X.Results) }
func (x GoStmt) Get(i int) Ast     { return CallExpr{x.X.Call} }
func (x Ident) Get(i int) Ast      { return badIndex(i, 0) }
func (x IfStmt) Get(i int) Ast     { return ToAst4(i, x.X.Init, x.X.Cond, x.X.Body, x.X.Else) }
func (x ImportSpec) Get(i int) Ast { return ToAst2(i, x.X.Name, x.X.Path) }
func (x IncDecStmt) Get(i int) Ast { return ToAst1(i, x.X.X) }
func (x IndexExpr) Get(i int) Ast  { return ToAst2(i, x.X.X, x.X.Index) }
func (x InterfaceType) Get(i int) Ast {
	if i == 0 {
		if x.X.Methods != nil {
			return FieldList{x.X.Methods}
		}
		return nil
	} else {
		return badIndex(i, 1)
	}
}
func (x KeyValueExpr) Get(i int) Ast   { return ToAst2(i, x.X.Key, x.X.Value) }
func (x LabeledStmt) Get(i int) Ast    { return ToAst2(i, x.X.Label, x.X.Stmt) }
func (x MapType) Get(i int) Ast        { return ToAst2(i, x.X.Key, x.X.Value) }
func (x Package) Get(i int) Ast        { return nil } // TODO
func (x ParenExpr) Get(i int) Ast      { return ToAst1(i, x.X.X) }
func (x RangeStmt) Get(i int) Ast      { return ToAst4(i, x.X.Key, x.X.Value, x.X.X, x.X.Body) }
func (x SelectStmt) Get(i int) Ast     { return ToAst1(i, x.X.Body) }
func (x SelectorExpr) Get(i int) Ast   { return ToAst2(i, x.X.X, x.X.Sel) }
func (x SendStmt) Get(i int) Ast       { return ToAst2(i, x.X.Chan, x.X.Value) }
func (x SliceExpr) Get(i int) Ast      { return ToAst4(i, x.X.X, x.X.Low, x.X.High, x.X.Max) }
func (x StarExpr) Get(i int) Ast       { return ToAst1(i, x.X.X) }
func (x StructType) Get(i int) Ast     { return ToAst1(i, x.X.Fields) }
func (x SwitchStmt) Get(i int) Ast     { return ToAst3(i, x.X.Init, x.X.Tag, x.X.Body) }
func (x TypeAssertExpr) Get(i int) Ast { return ToAst2(i, x.X.X, x.X.Type) }
func (x TypeSpec) Get(i int) Ast       { return ToAst2(i, x.X.Name, x.X.Type) }
func (x TypeSwitchStmt) Get(i int) Ast { return ToAst3(i, x.X.Init, x.X.Assign, x.X.Body) }
func (x UnaryExpr) Get(i int) Ast      { return ToAst1(i, x.X.X) }
func (x ValueSpec) Get(i int) Ast {
	switch i {
	case 0:
		if x.X.Names != nil {
			return IdentSlice{x.X.Names}
		}
	case 1:
		if x.X.Type != nil {
			return ToAst(x.X.Type)
		}
	case 2:
		if x.X.Values != nil {
			return ExprSlice{x.X.Values}
		}
	default:
		return badIndex(i, 3)
	}
	return nil
}

//
// .................. functions Set(int, Ast)
//
func (x ArrayType) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.X.Len = expr
	} else if i == 1 {
		x.X.Elt = expr
	} else {
		badIndex(i, 2)
	}
}
func (x AssignStmt) Set(i int, child Ast) {
	exprs := ToExprSlice(child)
	if i == 0 {
		x.X.Lhs = exprs
	} else if i == 1 {
		x.X.Rhs = exprs
	} else {
		badIndex(i, 2)
	}
}
func (x BadDecl) Set(i int, child Ast)  { badIndex(i, 0) }
func (x BadExpr) Set(i int, child Ast)  { badIndex(i, 0) }
func (x BadStmt) Set(i int, child Ast)  { badIndex(i, 0) }
func (x BasicLit) Set(i int, child Ast) { badIndex(i, 0) }
func (x BinaryExpr) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.X.X = expr
	} else if i == 1 {
		x.X.Y = expr
	} else {
		badIndex(i, 2)
	}
}
func (x BranchStmt) Set(i int, child Ast) {
	if i == 0 {
		x.X.Label = ToIdent(child)
	} else {
		badIndex(i, 1)
	}
}
func (x CallExpr) Set(i int, child Ast) {
	if i == 0 {
		x.X.Fun = ToExpr(child)
	} else if i == 1 {
		x.X.Args = ToExprSlice(child)
	} else {
		badIndex(i, 2)
	}
}
func (x CaseClause) Set(i int, child Ast) {
	if i == 0 {
		x.X.List = ToExprSlice(child)
	} else if i == 1 {
		x.X.Body = ToStmtSlice(child)
	} else {
		badIndex(i, 2)
	}
}
func (x ChanType) Set(i int, child Ast) {
	if i == 0 {
		x.X.Value = ToExpr(child)
	} else {
		badIndex(i, 1)
	}
}
func (x CommClause) Set(i int, child Ast) {
	if i == 0 {
		x.X.Comm = ToStmt(child)
	} else if i == 1 {
		x.X.Body = ToStmtSlice(child)
	} else {
		badIndex(i, 2)
	}
}
func (x CompositeLit) Set(i int, child Ast) {
	if i == 0 {
		x.X.Type = ToExpr(child)
	} else if i == 1 {
		x.X.Elts = ToExprSlice(child)
	} else {
		badIndex(i, 2)
	}
}
func (x DeclStmt) Set(i int, child Ast) {
	if i == 0 {
		x.X.Decl = ToDecl(child)
	} else {
		badIndex(i, 1)
	}
}
func (x DeferStmt) Set(i int, child Ast) {
	if i == 0 {
		x.X.Call = ToCallExpr(child)
	} else {
		badIndex(i, 1)
	}
}
func (x Ellipsis) Set(i int, child Ast) {
	if i == 0 {
		x.X.Elt = ToExpr(child)
	} else {
		badIndex(i, 1)
	}
}
func (x EmptyStmt) Set(i int, child Ast) { badIndex(i, 0) }
func (x ExprStmt) Set(i int, child Ast) {
	if i == 0 {
		x.X.X = ToExpr(child)
	} else {
		badIndex(i, 1)
	}
}
func (x Field) Set(i int, child Ast) {
	if i == 0 {
		x.X.Names = ToIdentSlice(child)
	} else if i == 1 {
		x.X.Type = ToExpr(child)
	} else if i == 2 {
		x.X.Tag = ToBasicLit(child)
	} else {
		badIndex(i, 3)
	}
}
func (x ForStmt) Set(i int, child Ast) {
	switch i {
	case 0:
		x.X.Init = ToStmt(child)
	case 1:
		x.X.Cond = ToExpr(child)
	case 2:
		x.X.Post = ToStmt(child)
	case 3:
		x.X.Body = ToBlockStmt(child)
	default:
		badIndex(i, 4)
	}
}
func (x FuncDecl) Set(i int, child Ast) {
	switch i {
	case 0:
		x.X.Recv = ToFieldList(child)
	case 1:
		x.X.Name = ToIdent(child)
	case 2:
		x.X.Type = ToFuncType(child)
	case 3:
		x.X.Body = ToBlockStmt(child)
	default:
		badIndex(i, 4)
	}
}
func (x FuncLit) Set(i int, child Ast) {
	if i == 0 {
		x.X.Type = ToFuncType(child)
	} else if i == 1 {
		x.X.Body = ToBlockStmt(child)
	} else {
		badIndex(i, 2)
	}
}
func (x FuncType) Set(i int, child Ast) {
	list := ToFieldList(child)
	if i == 0 {
		x.X.Params = list
	} else if i == 1 {
		x.X.Results = list
	} else {
		badIndex(i, 2)
	}
}
func (x GoStmt) Set(i int, child Ast) {
	if i == 0 {
		x.X.Call = ToCallExpr(child)
	} else {
		badIndex(i, 1)
	}
}
func (x Ident) Set(i int, child Ast) { badIndex(i, 0) }
func (x IfStmt) Set(i int, child Ast) {
	switch i {
	case 0:
		x.X.Init = ToStmt(child)
	case 1:
		x.X.Cond = ToExpr(child)
	case 2:
		x.X.Body = ToBlockStmt(child)
	case 3:
		x.X.Else = ToStmt(child)
	default:
		badIndex(i, 4)
	}
}
func (x ImportSpec) Set(i int, child Ast) {
	if i == 0 {
		x.X.Name = ToIdent(child)
	} else if i == 1 {
		x.X.Path = ToBasicLit(child)
	} else {
		badIndex(i, 2)
	}
}
func (x IncDecStmt) Set(i int, child Ast) {
	if i == 0 {
		x.X.X = ToExpr(child)
	} else {
		badIndex(i, 1)
	}
}
func (x IndexExpr) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.X.X = expr
	} else if i == 1 {
		x.X.Index = expr
	} else {
		badIndex(i, 2)
	}
}
func (x InterfaceType) Set(i int, child Ast) {
	if i == 0 {
		x.X.Methods = ToFieldList(child)
	} else {
		badIndex(i, 1)
	}
}
func (x KeyValueExpr) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.X.Key = expr
	} else if i == 1 {
		x.X.Value = expr
	} else {
		badIndex(i, 2)
	}
}
func (x LabeledStmt) Set(i int, child Ast) {
	if i == 0 {
		x.X.Label = ToIdent(child)
	} else if i == 1 {
		x.X.Stmt = ToStmt(child)
	} else {
		badIndex(i, 2)
	}
}
func (x MapType) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.X.Key = expr
	} else if i == 1 {
		x.X.Value = expr
	} else {
		badIndex(i, 2)
	}
}
func (x Package) Set(i int, child Ast) {} // TODO
func (x ParenExpr) Set(i int, child Ast) {
	if i == 0 {
		x.X.X = ToExpr(child)
	} else {
		badIndex(i, 1)
	}
}
func (x RangeStmt) Set(i int, child Ast) {
	switch i {
	case 0:
		x.X.Key = ToExpr(child)
	case 1:
		x.X.Value = ToExpr(child)
	case 2:
		x.X.X = ToExpr(child)
	case 3:
		x.X.Body = ToBlockStmt(child)
	default:
		badIndex(i, 4)
	}
}
func (x SelectStmt) Set(i int, child Ast) {
	if i == 0 {
		x.X.Body = ToBlockStmt(child)
	} else {
		badIndex(i, 1)
	}
}
func (x SelectorExpr) Set(i int, child Ast) {
	if i == 0 {
		x.X.X = ToExpr(child)
	} else if i == 1 {
		x.X.Sel = ToIdent(child)
	} else {
		badIndex(i, 2)
	}
}
func (x SendStmt) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.X.Chan = expr
	} else if i == 1 {
		x.X.Value = expr
	} else {
		badIndex(i, 2)
	}
}
func (x SliceExpr) Set(i int, child Ast) {
	expr := ToExpr(child)
	switch i {
	case 0:
		x.X.X = expr
	case 1:
		x.X.Low = expr
	case 2:
		x.X.High = expr
	case 3:
		x.X.Max = expr
		x.X.Slice3 = expr != nil
	default:
		badIndex(i, 4)
	}
}
func (x StarExpr) Set(i int, child Ast) {
	if i == 0 {
		x.X.X = ToExpr(child)
	} else {
		badIndex(i, 1)
	}
}
func (x StructType) Set(i int, child Ast) {
	if i == 0 {
		x.X.Fields = ToFieldList(child)
	} else {
		badIndex(i, 1)
	}
}
func (x SwitchStmt) Set(i int, child Ast) {
	switch i {
	case 0:
		x.X.Init = ToStmt(child)
	case 1:
		x.X.Tag = ToExpr(child)
	case 2:
		x.X.Body = ToBlockStmt(child)
	default:
		badIndex(i, 3)
	}
}
func (x TypeAssertExpr) Set(i int, child Ast) {
	if i == 0 {
		x.X.X = ToExpr(child)
	} else if i == 1 {
		x.X.Type = ToExpr(child)
	} else {
		badIndex(i, 2)
	}
}
func (x TypeSpec) Set(i int, child Ast) {
	if i == 0 {
		x.X.Name = ToIdent(child)
	} else if i == 1 {
		x.X.Type = ToExpr(child)
	} else {
		badIndex(i, 2)
	}
}
func (x TypeSwitchStmt) Set(i int, child Ast) {
	switch i {
	case 0:
		x.X.Init = ToStmt(child)
	case 1:
		x.X.Assign = ToStmt(child)
	case 2:
		x.X.Body = ToBlockStmt(child)
	default:
		badIndex(i, 3)
	}
}
func (x UnaryExpr) Set(i int, child Ast) {
	if i == 0 {
		x.X.X = ToExpr(child)
	} else {
		badIndex(i, 1)
	}
}
func (x ValueSpec) Set(i int, child Ast) {
	switch i {
	case 0:
		x.X.Names = ToIdentSlice(child)
	case 1:
		x.X.Type = ToExpr(child)
	case 2:
		x.X.Values = ToExprSlice(child)
	default:
		badIndex(i, 3)
	}
}
