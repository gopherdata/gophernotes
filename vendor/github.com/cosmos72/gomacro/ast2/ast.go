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
 * ast.go
 *
 *  Created on Feb 24, 2017
 *      Author Massimiliano Ghilardi
 */

package ast2

import (
	"go/ast"
	"go/token"
)

type (
	Ast interface {
		Interface() interface{}
		Op() token.Token
		Size() int
		Get(i int) Ast
		Set(i int, child Ast)
		New() Ast // returns a copy of Ast. the children are not copied
	}
	AstWithNode interface {
		Ast
		Node() ast.Node
	}
	AstWithSlice interface {
		Ast
		Slice(lo, hi int) AstWithSlice
		Append(child Ast) AstWithSlice
	}

	AstSlice   struct{ X []Ast }
	NodeSlice  struct{ X []ast.Node }
	ExprSlice  struct{ X []ast.Expr }
	FieldSlice struct{ X []*ast.Field }
	DeclSlice  struct{ X []ast.Decl }
	IdentSlice struct{ X []*ast.Ident }
	StmtSlice  struct{ X []ast.Stmt }
	SpecSlice  struct{ X []ast.Spec }

	ArrayType      struct{ X *ast.ArrayType }
	AssignStmt     struct{ X *ast.AssignStmt }
	BadDecl        struct{ X *ast.BadDecl }
	BadExpr        struct{ X *ast.BadExpr }
	BadStmt        struct{ X *ast.BadStmt }
	BasicLit       struct{ X *ast.BasicLit }
	BinaryExpr     struct{ X *ast.BinaryExpr }
	BlockStmt      struct{ X *ast.BlockStmt }
	BranchStmt     struct{ X *ast.BranchStmt }
	CallExpr       struct{ X *ast.CallExpr }
	CaseClause     struct{ X *ast.CaseClause }
	ChanType       struct{ X *ast.ChanType }
	CommClause     struct{ X *ast.CommClause }
	CompositeLit   struct{ X *ast.CompositeLit }
	DeclStmt       struct{ X *ast.DeclStmt }
	DeferStmt      struct{ X *ast.DeferStmt }
	Ellipsis       struct{ X *ast.Ellipsis }
	EmptyStmt      struct{ X *ast.EmptyStmt }
	ExprStmt       struct{ X *ast.ExprStmt }
	Field          struct{ X *ast.Field }
	FieldList      struct{ X *ast.FieldList }
	File           struct{ X *ast.File }
	ForStmt        struct{ X *ast.ForStmt }
	FuncDecl       struct{ X *ast.FuncDecl }
	FuncLit        struct{ X *ast.FuncLit }
	FuncType       struct{ X *ast.FuncType }
	GenDecl        struct{ X *ast.GenDecl }
	GoStmt         struct{ X *ast.GoStmt }
	Ident          struct{ X *ast.Ident }
	IfStmt         struct{ X *ast.IfStmt }
	ImportSpec     struct{ X *ast.ImportSpec }
	IncDecStmt     struct{ X *ast.IncDecStmt }
	IndexExpr      struct{ X *ast.IndexExpr }
	InterfaceType  struct{ X *ast.InterfaceType }
	KeyValueExpr   struct{ X *ast.KeyValueExpr }
	LabeledStmt    struct{ X *ast.LabeledStmt }
	MapType        struct{ X *ast.MapType }
	Package        struct{ X *ast.Package }
	ParenExpr      struct{ X *ast.ParenExpr }
	RangeStmt      struct{ X *ast.RangeStmt }
	ReturnStmt     struct{ X *ast.ReturnStmt }
	SelectStmt     struct{ X *ast.SelectStmt }
	SelectorExpr   struct{ X *ast.SelectorExpr }
	SendStmt       struct{ X *ast.SendStmt }
	SliceExpr      struct{ X *ast.SliceExpr }
	StarExpr       struct{ X *ast.StarExpr }
	StructType     struct{ X *ast.StructType }
	SwitchStmt     struct{ X *ast.SwitchStmt }
	TypeAssertExpr struct{ X *ast.TypeAssertExpr }
	TypeSpec       struct{ X *ast.TypeSpec }
	TypeSwitchStmt struct{ X *ast.TypeSwitchStmt }
	UnaryExpr      struct{ X *ast.UnaryExpr }
	ValueSpec      struct{ X *ast.ValueSpec }
)
