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
 * wrap.go
 *
 *  Created on: Feb 25, 2017
 *      Author: Massimiliano Ghilardi
 */

package ast2

import (
	"fmt"
	"go/ast"
	"go/token"
	r "reflect"
)

// unused
/*
func CloneAst(in Ast) Ast {
	switch in := in.(type) {
	case AstWithNode:
		return CloneAstWithNode(in)
	case AstWithSlice:
		return CloneAstWithSlice(in)
	default:
		errorf("CloneAst: unsupported argument type, expecting AstWithNode or AstWithSlice: %v // %T", in, in)
		return nil
	}
}

func CloneAstWithNode(in AstWithNode) AstWithNode {
	form := in.New().(AstWithNode)
	n := in.Size()
	for i := 0; i < n; i++ {
		form.Set(i, CloneAst(in.Get(i)))
	}
	return form
}

func CloneAstWithSlice(in AstWithSlice) AstWithSlice {
	form := in.New().(AstWithSlice)
	n := in.Size()
	for i := 0; i < n; i++ {
		form = form.Append(CloneAst(in.Get(i)))
	}
	return form
}
*/

func AnyToAstWithNode(any interface{}, caller interface{}) AstWithNode {
	node := AnyToAst(any, caller)
	switch node := node.(type) {
	case AstWithNode:
		return node
	default:
		errorf("%s: cannot convert to ast.Node: %v <%v>", caller, any, r.TypeOf(any))
		return nil
	}
}

func AnyToAstWithSlice(any interface{}, caller interface{}) AstWithSlice {
	node := AnyToAst(any, caller)
	switch node := node.(type) {
	case nil:
		return NodeSlice{}
	case AstWithSlice:
		return node
	default:
		errorf("%s: cannot convert to slice of ast.Node: %v <%v>", caller, any, r.TypeOf(any))
		return nil
	}
}

func AnyToAst(any interface{}, caller interface{}) Ast {
	var str string
	var tok token.Token
	switch node := any.(type) {
	case nil:
		return nil
	case Ast:
		return node
	case ast.Node:
		return ToAst(node)
	case []Ast:
		return AstSlice{X: node}
	case []ast.Node:
		return NodeSlice{X: node}
	case []*ast.Field:
		return FieldSlice{X: node}
	case []ast.Decl:
		return DeclSlice{X: node}
	case []ast.Expr:
		return ExprSlice{X: node}
	case []*ast.Ident:
		return IdentSlice{X: node}
	case []ast.Stmt:
		return StmtSlice{X: node}
	case []ast.Spec:
		return SpecSlice{X: node}
	case bool:
		if node {
			str = "true"
		} else {
			str = "false"
		}
		return Ident{X: &ast.Ident{Name: str}}
	/*
		case rune: // Go cannot currently distinguish rune from int32
			tok = token.CHAR
			str = fmt.Sprintf("%q", node)
	*/
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr:
		tok = token.INT
		str = fmt.Sprintf("%d", node)
	case float32, float64:
		tok = token.FLOAT
		str = fmt.Sprintf("%g", node)
	case complex64, complex128:
		errorf("%s: unimplemented conversion of %T to ast.Node: %v", caller, any, any)
		return nil
	case string:
		tok = token.STRING
		str = fmt.Sprintf("%q", node)
	default:
		errorf("%s: cannot convert to ast.Node: %v // %T", caller, any, any)
		return nil
	}
	return BasicLit{X: &ast.BasicLit{Kind: tok, Value: str}}

}

// ToAst2 returns either n0 (if i == 0) or n1, converted to Ast
func ToAst1(i int, node ast.Node) AstWithNode {
	if i == 0 {
		return ToAst(node)
	} else {
		return badIndex(i, 1)
	}
}

// ToAst2 returns either n0 (if i == 0) or n1, converted to Ast
func ToAst2(i int, n0 ast.Node, n1 ast.Node) AstWithNode {
	var n ast.Node
	switch i {
	case 0:
		n = n0
	case 1:
		n = n1
	default:
		return badIndex(i, 2)
	}
	return ToAst(n)
}

func ToAst3(i int, n0 ast.Node, n1 ast.Node, n2 *ast.BlockStmt) AstWithNode {
	var n ast.Node
	switch i {
	case 0:
		n = n0
	case 1:
		n = n1
	case 2:
		if n2 == nil {
			return nil
		}
		return BlockStmt{n2}
	default:
		return badIndex(i, 3)
	}
	return ToAst(n)
}

func ToAst4(i int, n0 ast.Node, n1 ast.Node, n2 ast.Node, n3 ast.Node) AstWithNode {
	var n ast.Node
	switch i {
	case 0:
		n = n0
	case 1:
		n = n1
	case 2:
		n = n2
	case 3:
		n = n3
	default:
		return badIndex(i, 4)
	}
	return ToAst(n)
}

// ToAst converts an ast.Node to Ast, providing uniform access to the node contents
//
func ToAst(node ast.Node) AstWithNode {
	var x AstWithNode
	switch node := node.(type) {
	case nil:
		return nil
	case *ast.ArrayType:
		x = ArrayType{node}
	case *ast.AssignStmt:
		x = AssignStmt{node}
	case *ast.BadDecl:
		x = BadDecl{node}
	case *ast.BadExpr:
		x = BadExpr{node}
	case *ast.BadStmt:
		x = BadStmt{node}
	case *ast.BasicLit:
		x = BasicLit{node}
	case *ast.BinaryExpr:
		x = BinaryExpr{node}
	case *ast.BlockStmt:
		if node != nil { // we can get typed nil from many places
			x = BlockStmt{node}
		}
	case *ast.BranchStmt:
		x = BranchStmt{node}
	case *ast.CallExpr:
		x = CallExpr{node}
	case *ast.CaseClause:
		x = CaseClause{node}
	case *ast.ChanType:
		x = ChanType{node}
	case *ast.CommClause:
		x = CommClause{node}
	case *ast.CompositeLit:
		x = CompositeLit{node}
	case *ast.DeclStmt:
		x = DeclStmt{node}
	case *ast.DeferStmt:
		x = DeferStmt{node}
	case *ast.Ellipsis:
		x = Ellipsis{node}
	case *ast.EmptyStmt:
		x = EmptyStmt{node}
	case *ast.ExprStmt:
		x = ExprStmt{node}
	case *ast.Field:
		x = Field{node}
	case *ast.FieldList:
		if node != nil { // we can get typed nil from many places
			x = FieldList{node}
		}
	case *ast.File:
		x = File{node}
	case *ast.ForStmt:
		x = ForStmt{node}
	case *ast.FuncDecl:
		x = FuncDecl{node}
	case *ast.FuncLit:
		x = FuncLit{node}
	case *ast.FuncType:
		x = FuncType{node}
	case *ast.GenDecl:
		if node != nil {
			x = GenDecl{node}
		}
	case *ast.GoStmt:
		x = GoStmt{node}
	case *ast.Ident:
		x = Ident{node}
	case *ast.IfStmt:
		x = IfStmt{node}
	case *ast.ImportSpec:
		x = ImportSpec{node}
	case *ast.IncDecStmt:
		x = IncDecStmt{node}
	case *ast.IndexExpr:
		x = IndexExpr{node}
	case *ast.InterfaceType:
		x = InterfaceType{node}
	case *ast.KeyValueExpr:
		x = KeyValueExpr{node}
	case *ast.LabeledStmt:
		x = LabeledStmt{node}
	case *ast.MapType:
		x = MapType{node}
	case *ast.Package:
		x = Package{node}
	case *ast.ParenExpr:
		x = ParenExpr{node}
	case *ast.RangeStmt:
		x = RangeStmt{node}
	case *ast.ReturnStmt:
		x = ReturnStmt{node}
	case *ast.SelectStmt:
		x = SelectStmt{node}
	case *ast.SelectorExpr:
		x = SelectorExpr{node}
	case *ast.SendStmt:
		x = SendStmt{node}
	case *ast.SliceExpr:
		x = SliceExpr{node}
	case *ast.StarExpr:
		x = StarExpr{node}
	case *ast.StructType:
		x = StructType{node}
	case *ast.SwitchStmt:
		x = SwitchStmt{node}
	case *ast.TypeAssertExpr:
		x = TypeAssertExpr{node}
	case *ast.TypeSpec:
		x = TypeSpec{node}
	case *ast.TypeSwitchStmt:
		x = TypeSwitchStmt{node}
	case *ast.UnaryExpr:
		x = UnaryExpr{node}
	case *ast.ValueSpec:
		x = ValueSpec{node}
	default:
		errorf("unsupported node type %T", node)
	}
	return x
}

func ToAstWithSlice(x Ast, caller interface{}) AstWithSlice {
	switch x := x.(type) {
	case AstWithSlice:
		return x
	default:
		y := x.Interface()
		errorf("%s: cannot convert to slice of ast.Node: %v <%v>", caller, y, r.TypeOf(y))
		return nil
	}
}
