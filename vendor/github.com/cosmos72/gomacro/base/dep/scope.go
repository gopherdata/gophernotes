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
 * loader.go
 *
 *  Created on: May 03, 2018
 *      Author: Massimiliano Ghilardi
 */

package dep

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
	"strings"

	"github.com/cosmos72/gomacro/base/output"

	"github.com/cosmos72/gomacro/ast2"
)

func (s *Scope) Ast(form ast2.Ast) []string {
	var deps []string
	switch form := form.(type) {
	case nil:
	case ast2.AstWithNode:
		deps = s.Node(form.Node())
	case ast2.AstWithSlice:
		n := form.Size()
		for i := 0; i < n; i++ {
			deps = append(deps, s.Ast(form.Get(i))...)
		}
	default:
		output.Errorf("Scope.Ast(): unsupported ast2.Ast node, expecting ast2.AstWithNode or ast2.AstWithSlice, found %v // %T", form, form)
	}
	return deps
}

func (s *Scope) Nodes(nodes []ast.Node) {
	s.Ast(ast2.NodeSlice{nodes})
}

func (s *Scope) Node(node ast.Node) []string {
	var deps []string
	switch node := node.(type) {
	case nil:
	case ast.Decl:
		deps = s.Decl(node)
	case ast.Expr:
		s.add(NewDeclExpr(node, &s.Gensym))
	case ast.Stmt:
		s.add(NewDeclStmt(node, &s.Gensym))
	case *ast.File:
		deps = s.File(node)
	default:
		output.Errorf("Scope.Ast(): unsupported node type, expecting ast.Decl, ast.Expr, ast.Stmt or *ast.File, found %v // %T", node, node)
	}
	return sort_unique_inplace(deps)
}

func (s *Scope) Decl(node ast.Node) []string {
	var deps []string
	switch node := node.(type) {
	case nil:
	case *ast.GenDecl:
		deps = s.GenDecl(node)
	case *ast.FuncDecl:
		deps = s.Func(node)
	default:
		output.Errorf("Scope.Decl(): unsupported declaration, expecting *ast.GenDecl or *ast.FuncDecl, found: %v // %T", node, node)
	}
	return deps
}

func (s *Scope) File(node *ast.File) []string {
	var deps []string
	if node != nil {
		for _, decl := range node.Decls {
			deps = append(deps, s.Decl(decl)...)
		}
	}
	return deps
}

// for consts that inherit type and initializers from a previous *ast.ValueSpec
type ConstDeps struct {
	Type      ast.Expr
	TypeDeps  []string
	Values    []ast.Expr
	ValueDeps [][]string
}

func (s *Scope) GenDecl(node *ast.GenDecl) []string {
	var deps []string
	switch node.Tok {
	case token.CONST:
		var defaults ConstDeps
		iota := 0
		for _, spec := range node.Specs {
			deps = append(deps, s.Consts(spec, iota, &defaults)...)
			iota++
		}
	case token.IMPORT:
		for _, spec := range node.Specs {
			s.Import(spec)
		}
	case token.PACKAGE:
		for _, spec := range node.Specs {
			s.Package(spec)
		}
	case token.TYPE:
		for _, spec := range node.Specs {
			deps = append(deps, s.Type(spec)...)
		}
	case token.VAR:
		for _, spec := range node.Specs {
			deps = append(deps, s.Vars(spec)...)
		}
	default:
		output.Errorf("Scope.GenDecl(): unsupported declaration kind, expecting token.IMPORT, token.PACKAGE, token.CONST, token.TYPE or token.VAR, found %v: %v // %T",
			node.Tok, node, node)
	}
	return deps
}

// constants
func (s *Scope) Consts(node ast.Spec, iota int, defaults *ConstDeps) []string {
	var deps []string

	if node, ok := node.(*ast.ValueSpec); ok {
		if node.Type != nil && node.Values == nil {
			output.Errorf("const declaration cannot have type without expression: %v // %T", node, node)
		}
		// if expressions are omitted, they default to the last ones found (with their type, if any)
		if node.Type != nil || node.Values != nil {
			defaults.Type = node.Type
			defaults.TypeDeps = s.Expr(node.Type)
			deps = append(deps, defaults.TypeDeps...)

			defaults.Values = node.Values
			defaults.ValueDeps = s.Exprs(node.Values)
			for _, list := range defaults.ValueDeps {
				deps = append(deps, list...)
			}
		}
		if len(defaults.Values) != len(node.Names) {
			output.Errorf("%d consts initialized with %d expressions: %v %v = %v",
				len(node.Names), len(defaults.Values), node.Names, defaults.Type, defaults.Values)
		}
		var declNode ast.Spec
		if len(node.Names) == 1 {
			declNode = node
		}
		for i, ident := range node.Names {
			var value ast.Expr
			deps := defaults.TypeDeps
			if i < len(defaults.Values) {
				value = defaults.Values[i]
				deps = append(dup(deps), defaults.ValueDeps[i]...)
			}
			s.Const(ident, declNode, iota, defaults.Type, value, deps)
		}
	} else {
		output.Errorf("unsupported constant declaration: expecting *ast.ValueSpec, found: %v // %T", node, node)
	}
	return deps
}

// constant
func (s *Scope) Const(ident *ast.Ident, node ast.Spec, iota int, typ ast.Expr, value ast.Expr, deps []string) *Decl {
	decl := NewDecl(Const, ident.Name, node, ident.Pos(), deps)
	decl.Extra = &Extra{
		Ident: ident,
		Type:  typ,
		Value: value,
		Iota:  iota,
	}
	return s.add(decl)
}

func unquote(src string) string {
	ret, err := strconv.Unquote(src)
	if err != nil && len(src) >= 2 {
		if ch := src[0]; ch == src[len(src)-1] && (ch == '\'' || ch == '"' || ch == '`') {
			ret = src[1 : len(src)-1]
		} else {
			ret = src
		}
	}
	return ret
}

func basename(path string) string {
	return path[1+strings.LastIndexByte(path, '/'):]
}

// import
func (s *Scope) Import(node ast.Spec) {
	s.add(NewDeclImport(node, &s.Gensym))
}

// package
func (s *Scope) Package(node ast.Spec) {
	s.add(NewDeclPackage(node, &s.Gensym))
}

// variables
func (s *Scope) Vars(node ast.Spec) []string {
	var alldeps []string
	if node, ok := node.(*ast.ValueSpec); ok {
		if len(node.Names) > 1 && len(node.Values) == 1 {
			return s.varsMultiValueExpr(node)
		}
		if len(node.Values) != 0 && len(node.Names) != len(node.Values) {
			output.Errorf("%d vars initialized with %d expressions: %v", len(node.Names), len(node.Values), node)
		}
		typDeps := s.Expr(node.Type)
		alldeps = append(alldeps, typDeps...)
		var declNode ast.Spec
		if len(node.Names) == 1 {
			declNode = node
		}
		for i, ident := range node.Names {
			deps := typDeps
			var value ast.Expr
			if i < len(node.Values) {
				value = node.Values[i]
				valueDeps := s.Expr(value)
				alldeps = append(alldeps, valueDeps...)
				if len(valueDeps) != 0 {
					deps = append(dup(typDeps), valueDeps...)
				}
			}
			s.Var(ident, declNode, node.Type, value, deps)
		}
	} else {
		output.Errorf("Scope.Vars(): unsupported variable declaration: expecting *ast.ValueSpec, found: %v // %T", node, node)
	}
	return alldeps
}

func (s *Scope) varsMultiValueExpr(node *ast.ValueSpec) []string {
	deps := append(s.Expr(node.Type), s.Expr(node.Values[0])...)
	for _, ident := range node.Names {
		s.add(NewDeclVarMulti(ident, node, deps))
		node = nil // store node only in the first VarMulti
	}
	return deps
}

// variable
func (s *Scope) Var(ident *ast.Ident, node ast.Spec, typ ast.Expr, value ast.Expr, deps []string) *Decl {
	return s.add(NewDeclVar(ident, node, typ, value, deps))
}

// function or method
func (s *Scope) Func(node *ast.FuncDecl) []string {
	inner := NewScope(s)

	name := node.Name.Name
	deps := inner.Expr(node.Type)

	kind := Func
	if node.Recv != nil && len(node.Recv.List) != 0 {
		types := inner.Expr(node.Recv)
		// method names are not global!
		// without this, a method Foo.String would overwrite a func String in s.Decls[]
		//
		// also makes it impossible to depend on a method, but nothing can depend on a method,
		// Except the constant returned by unsafe.Sizeof(type.method),
		// but we do not support unsafe.Sizeof() yet and all methods have the same size anyway
		if len(types) == 1 {
			name = fmt.Sprintf("%s.%s", types[0], name)
		} else {
			name = fmt.Sprintf("%d.%s", s.Gensym, name)
			s.Gensym++
		}

		deps = append(deps, types...)
		kind = Method
	}
	// support recursive functions: forward-declare the function
	// decl := &Decl{Kind: kind, Name: name}
	// s.Decls.add(decl)

	// check function body for global constants, types, variables!
	deps = append(deps, inner.Expr(node.Body)...)

	s.add(NewDeclFunc(kind, name, node, deps))
	return deps
}

// type
func (s *Scope) Type(node ast.Spec) []string {
	var deps []string
	if node, ok := node.(*ast.TypeSpec); ok {
		deps = s.Expr(node.Type)

		s.add(NewDeclType(node, deps))
	} else {
		output.Errorf("Scope.Type(): unexpected declaration type, expecting *ast.TypeSpec, found: %v // %T", node, node)
	}
	return deps
}

func (s *Scope) Expr(node ast.Node) []string {
	if node == nil {
		return nil
	}
	return s.AstExpr(ast2.ToAst(node))
}

func (s *Scope) Exprs(list []ast.Expr) [][]string {
	n := len(list)
	if n == 0 {
		return nil
	}
	deps := make([][]string, n)
	for i, expr := range list {
		deps[i] = s.Expr(expr)
	}
	return deps
}

func (s *Scope) AstExpr(in ast2.Ast) []string {
	if in == nil {
		return nil
	}
	var deps []string
	switch node := in.Interface().(type) {
	case *ast.FuncLit:
		deps = append(deps, s.Expr(node.Type)...)
		in = ast2.BlockStmt{node.Body}
		// open a new scope
		s = NewScope(s)
	case *ast.BlockStmt, *ast.FuncType, *ast.InterfaceType, *ast.StructType:
		// open a new scope
		s = NewScope(s)
	case *ast.KeyValueExpr:
		// ignore the key if it's an ast.Ident
		// FIXME this is correct for struct initializers only
		if _, ok := node.Key.(*ast.Ident); !ok {
			deps = append(deps, s.Expr(node.Key)...)
		}
		in = ast2.ToAst(node.Value)
	case ast.Decl:
		return s.Decl(node)
	case *ast.SelectorExpr:
		return s.selectorExpr(node)
	case *ast.Field:
		// declare field names and compute dependencies for their type
		deps = append(deps, s.Expr(node.Type)...)
		for _, ident := range node.Names {
			s.Var(ident, nil, node.Type, nil, deps)
		}
		return deps
	}
	if form, ok := in.(ast2.Ident); ok && form.X != nil && !s.isLocal(form.X.Name) {
		deps = append(deps, form.X.Name)
	}

	for i, n := 0, in.Size(); i < n; i++ {
		form := in.Get(i)
		if form != nil {
			deps = append(deps, s.AstExpr(form)...)
		}
	}
	return sort_unique_inplace(deps)
}

// return true if name refers to a local declaration
func (s *Scope) isLocal(name string) bool {
	outer := s.Outer
	// outer == nil is top-level scope: not local
	for ; outer != nil; s = outer {
		if _, ok := s.Decls[name]; ok {
			return true
		}
		outer = outer.Outer
	}
	return false
}

// compute dependencies for: package.symbol, type.method, type.field.
// only the part *before* the dot may be a local declaration,
// but dependency from type.method is stronger than dependency from type,
// so keep both
func (s *Scope) selectorExpr(node *ast.SelectorExpr) []string {
	deps := s.Expr(node.X)
	if typ, ok := node.X.(*ast.Ident); ok && typ != nil && !s.isLocal(typ.Name) {
		deps = append(deps, typ.Name+"."+node.Sel.Name)
	}
	return deps
}

func (s *Scope) add(decl *Decl) *Decl {
	return s.Decls.add(decl)
}
