/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * unwrap.go
 *
 *  Created on: May 06, 2018
 *      Author: Massimiliano Ghilardi
 */

package ast2

import (
	"go/ast"
	"go/token"
	r "reflect"

	"github.com/cosmos72/gomacro/go/etoken"
)

// ToNode converts Ast back ast.Node, or panics on failure
// (it fails if the argument is not AstWithNode)
func ToNode(x Ast) ast.Node {
	switch x := x.(type) {
	case nil:
		return nil
	case AstWithNode:
		return x.Node()
	default:
		y := x.Interface()
		errorf("cannot convert to ast.Node: %v // %T", y, y)
		return nil
	}
}

func ToBasicLit(x Ast) *ast.BasicLit {
	switch x := x.(type) {
	case nil:
		break
	case BasicLit:
		return x.X
	default:
		y := x.Interface()
		errorf("cannot convert to *ast.BasicLit: %v // %T", y, y)
	}
	return nil
}

func ToBlockStmt(x Ast) *ast.BlockStmt {
	switch x := x.(type) {
	case nil:
		break
	case BlockStmt:
		return x.X
	default:
		stmt := ToStmt(x)
		return &ast.BlockStmt{Lbrace: stmt.Pos(), List: []ast.Stmt{stmt}, Rbrace: stmt.End()}
	}
	return nil
}

func ToCallExpr(x Ast) *ast.CallExpr {
	switch x := x.(type) {
	case nil:
		break
	case CallExpr:
		return x.X
	default:
		y := x.Interface()
		errorf("cannot convert to *ast.CallExpr: %v // %T", y, y)
	}
	return nil
}

func ToDecl(x Ast) ast.Decl {
	switch node := ToNode(x).(type) {
	case ast.Decl:
		return node
	case nil:
	default:
		y := x.Interface()
		errorf("cannot convert to ast.Decl: %v // %T", y, y)
	}
	return nil
}

func ToExpr(x Ast) ast.Expr {
	switch node := ToNode(x).(type) {
	case nil:
		break
	case ast.Expr:
		return node
	case *ast.BlockStmt:
		return BlockStmtToExpr(node)
	case *ast.EmptyStmt:
		return &ast.Ident{NamePos: node.Semicolon, Name: "nil"}
	case *ast.ExprStmt:
		return node.X
	case ast.Stmt:
		list := []ast.Stmt{node}
		block := &ast.BlockStmt{List: list}
		return BlockStmtToExpr(block)
	default:
		errorf("unimplemented conversion from %v to ast.Expr: %v <%v>",
			r.TypeOf(node), node, r.TypeOf(node))
	}
	return nil
}

func ToExprSlice(x Ast) []ast.Expr {
	switch x := x.(type) {
	case nil:
		break
	case ExprSlice:
		return x.X
	case AstWithSlice:
		n := x.Size()
		ret := make([]ast.Expr, n)
		for i := 0; i < n; i++ {
			ret[i] = ToExpr(x.Get(i))
		}
		return ret
	default:
		errorf("unimplemented conversion from %v <%v> to []ast.Expr", x, r.TypeOf(x))
	}
	return nil
}

func ToField(x Ast) *ast.Field {
	switch node := ToNode(x).(type) {
	case nil:
		break
	case *ast.Field:
		return node
	default:
		errorf("cannot convert %v <%v> to *ast.Field", node, r.TypeOf(node))
	}
	return nil
}

func ToFile(x Ast) *ast.File {
	switch node := ToNode(x).(type) {
	case nil:
		break
	case *ast.File:
		return node
	default:
		errorf("cannot convert %v <%v> to *ast.File", node, r.TypeOf(node))
	}
	return nil
}

func ToFieldList(x Ast) *ast.FieldList {
	switch node := ToNode(x).(type) {
	case nil:
		break
	case *ast.FieldList:
		return node
	case *ast.Field:
		return &ast.FieldList{Opening: node.Pos(), List: []*ast.Field{node}, Closing: node.End()}
	default:
		errorf("cannot convert %v <%v> to *ast.Field", node, r.TypeOf(node))
	}
	return nil
}

func ToFuncType(x Ast) *ast.FuncType {
	switch node := ToNode(x).(type) {
	case nil:
		break
	case *ast.FuncType:
		return node
	default:
		errorf("cannot convert %v <%v> to *ast.FuncType", node, r.TypeOf(node))
	}
	return nil
}

func ToImportSpec(x Ast) *ast.ImportSpec {
	switch node := ToNode(x).(type) {
	case nil:
		break
	case *ast.ImportSpec:
		return node
	default:
		errorf("cannot convert %v <%v> to *ast.ImportSpec", node, r.TypeOf(node))
	}
	return nil
}

func ToIdent(x Ast) *ast.Ident {
	switch node := ToNode(x).(type) {
	case nil:
		break
	case *ast.Ident:
		return node
	default:
		errorf("cannot convert %v <%v> to *ast.Ident", node, r.TypeOf(node))
	}
	return nil
}

func ToIdentSlice(x Ast) []*ast.Ident {
	switch x := x.(type) {
	case nil:
		break
	case IdentSlice:
		return x.X
	case AstWithSlice:
		n := x.Size()
		ret := make([]*ast.Ident, n)
		for i := 0; i < n; i++ {
			ret[i] = ToIdent(x.Get(i))
		}
		return ret
	default:
		errorf("unimplemented conversion from %v <%v> to []*ast.Ident", x, r.TypeOf(x))
	}
	return nil
}

func ToSpec(x Ast) ast.Spec {
	switch node := ToNode(x).(type) {
	case nil:
		break
	case ast.Spec:
		return node
	default:
		errorf("cannot convert %v <%v> to ast.Spec", node, r.TypeOf(node))
	}
	return nil
}

func ToStmt(x Ast) ast.Stmt {
	switch node := ToNode(x).(type) {
	case ast.Stmt:
		return node
	case ast.Decl:
		return &ast.DeclStmt{Decl: node}
	case ast.Expr:
		return &ast.ExprStmt{X: node}
	case nil:
		break
	default:
		errorf("unimplemented conversion from %v <%v> to ast.Stmt", node, r.TypeOf(node))
	}
	return nil
}

func ToStmtSlice(x Ast) []ast.Stmt {
	switch x := x.(type) {
	case nil:
		break
	case StmtSlice:
		return x.X
	case AstWithSlice:
		n := x.Size()
		ret := make([]ast.Stmt, n)
		for i := 0; i < n; i++ {
			ret[i] = ToStmt(x.Get(i))
		}
		return ret
	default:
		errorf("unimplemented conversion from %v <%v> to []ast.Stmt", x, r.TypeOf(x))
	}
	return nil
}

func BlockStmtToExpr(node *ast.BlockStmt) ast.Expr {
	if node == nil {
		return nil
	}
	list := node.List
	switch len(list) {
	case 0:
		// convert {} to nil, because {} in expression context means "no useful value"
		return &ast.Ident{NamePos: node.Lbrace, Name: "nil"}
	case 1:
		// check if we are lucky...
		switch node := list[0].(type) {
		case *ast.ExprStmt:
			return node.X
		case *ast.EmptyStmt:
			// convert { ; } to nil, because { ; } in expression context means "no useful value"
			return &ast.Ident{NamePos: node.Semicolon, Name: "nil"}
		}
	}

	// due to go/ast strictly typed model, there is only one mechanism
	// to insert a statement inside an expression: use a closure.
	// so we return a unary expression: MACRO (func() { /*block*/ })
	typ := &ast.FuncType{Func: token.NoPos, Params: &ast.FieldList{}}
	fun := &ast.FuncLit{Type: typ, Body: node}
	return &ast.UnaryExpr{Op: etoken.MACRO, X: fun}
}
