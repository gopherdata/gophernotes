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
 * ast_convert.go
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

	mt "github.com/cosmos72/gomacro/token"
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
		errorf("CloneAst: unsupported argument type, expecting AstWithNode or AstWithSlice: %v <%v>", in, r.TypeOf(in))
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
		errorf("%s: unimplemented conversion of <%v> to ast.Node: %v", r.TypeOf(any), caller, any)
		return nil
	case string:
		tok = token.STRING
		str = fmt.Sprintf("%q", node)
	default:
		errorf("%s: cannot convert to ast.Node: %v <%v>", caller, any, r.TypeOf(any))
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
		errorf("unsupported node type <%v>", r.TypeOf(node))
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
		errorf("cannot convert to ast.Node: %v <%v>", y, r.TypeOf(y))
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
		errorf("cannot convert to *ast.BasicLit: %v <%v>", y, r.TypeOf(y))
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
		errorf("cannot convert to *ast.CallExpr: %v <%v>", y, r.TypeOf(y))
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
		errorf("cannot convert to ast.Decl: %v <%v>", y, r.TypeOf(y))
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
	return &ast.UnaryExpr{Op: mt.MACRO, X: fun}
}
