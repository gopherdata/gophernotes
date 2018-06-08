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
 * util.go
 *
 *  Created on: May 03, 2018
 *      Author: Massimiliano Ghilardi
 */

package dep

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"os"
	"sort"

	"github.com/cosmos72/gomacro/base"
)

// ===================== DeclMap =====================

func (m DeclMap) Dup() DeclMap {
	ret := make(DeclMap, len(m))
	for name, decl := range m {
		ret[name] = decl
	}
	return ret
}

func (m DeclMap) List() DeclList {
	list := make(DeclList, len(m))
	i := 0
	for _, e := range m {
		list[i] = e
		i++
	}
	return list
}

// remove all dependencies that cannot be resolved, i.e. not present among m
func (m DeclMap) RemoveUnresolvableDeps() {
	for _, decl := range m {
		decl.RemoveUnresolvableDeps(m)
	}
}

func (m DeclMap) Print() {
	m.List().SortByPos().Print()
}

func (m DeclMap) depMap() depMap {
	ret := make(depMap, len(m))
	for _, decl := range m {
		ret[decl.Name] = decl.depSet()
	}
	return ret
}

// ===================== DeclList ====================

func (list DeclList) Map() DeclMap {
	m := make(DeclMap, len(list))
	for _, e := range list {
		m[e.Name] = e
	}
	return m
}

func (list DeclList) SortByPos() DeclList {
	sort.Slice(list, func(i, j int) bool {
		a, b := list[i], list[j]
		return a.Pos < b.Pos
	})
	return list
}

func (list DeclList) Reverse() DeclList {
	n := len(list)
	for i := 0; i < n/2; i++ {
		temp := list[i]
		j := n - i - 1
		list[i] = list[j]
		list[j] = temp
	}
	return list
}

func (list DeclList) Print() {
	for _, decl := range list {
		decl.Print()
	}
}

// ======================= Decl ======================

func NewDeclImport(spec ast.Spec, counter *int) *Decl {
	node, ok := spec.(*ast.ImportSpec)
	if !ok {
		base.Errorf("NewDeclImport(): unsupported import: expecting *ast.ImportSpec, found: %v // %T", spec, spec)
	}

	var name string
	if ident := node.Name; ident != nil {
		if ident.Name != "." {
			name = ident.Name
		}
	} else {
		name = basename(unquote(node.Path.Value))
	}
	if len(name) == 0 {
		name = fmt.Sprintf("<import%d>", *counter)
		*counter++
	}
	return NewDecl(Import, name, node, node.Pos(), nil)
}

func NewDeclPackage(spec ast.Spec, counter *int) *Decl {
	node, ok := spec.(*ast.ValueSpec)
	if !ok {
		base.Errorf("NewDeclPackage(): unsupported package: expecting *ast.ValueSpec, found: %v // %T", spec, spec)
	}

	var pos token.Pos
	if len(node.Names) != 0 {
		pos = node.Names[0].Pos()
	} else if len(node.Values) != 0 {
		pos = node.Values[0].Pos()
	}
	name := fmt.Sprintf("<package%d>", *counter)
	*counter++
	return NewDecl(Package, name, node, pos, nil)
}

func NewDeclExpr(node ast.Expr, counter *int) *Decl {
	name := fmt.Sprintf("<expr%d>", *counter)
	*counter++
	return NewDecl(Expr, name, node, node.Pos(), nil)
}

func NewDeclFunc(kind Kind, name string, node *ast.FuncDecl, deps []string) *Decl {
	return NewDecl(kind, name, node, node.Name.Pos(), deps)
}

func NewDeclStmt(node ast.Stmt, counter *int) *Decl {
	name := fmt.Sprintf("<stmt%d>", *counter)
	*counter++
	return NewDecl(Stmt, name, node, node.Pos(), nil)
}

func NewDeclType(node *ast.TypeSpec, deps []string) *Decl {
	name := node.Name.Name
	deps = sort_unique_inplace(deps)
	// support self-referencing types, as for example: type List struct { First int; Rest *List }
	deps = remove_item_inplace(name, deps)

	return &Decl{Kind: Type, Name: name, Node: node, Deps: deps, Pos: node.Name.Pos()}
}

func NewDeclVar(ident *ast.Ident, node ast.Spec, typ ast.Expr, value ast.Expr, deps []string) *Decl {
	decl := NewDecl(Var, ident.Name, node, ident.Pos(), deps)
	decl.Extra = &Extra{
		Ident: ident,
		Type:  typ,
		Value: value,
	}
	return decl
}

func NewDeclVarMulti(ident *ast.Ident, node *ast.ValueSpec, deps []string) *Decl {
	return NewDecl(VarMulti, ident.Name, node, ident.Pos(), deps)
}

func (decl *Decl) depSet() set {
	ret := make(set, len(decl.Deps))
	for _, dep := range decl.Deps {
		ret[dep] = void{}
	}
	return ret
}

// remove all dependencies that cannot be resolved, i.e. not present among m
func (decl *Decl) RemoveUnresolvableDeps(m DeclMap) {
	decl.Deps = filter_if_inplace(decl.Deps, func(name string) bool {
		return m[name] != nil
	})
}

func (decl *Decl) Fprint(out io.Writer) {
	fmt.Fprintf(out, "%s%s%s\t%v\n", decl.Name, spaces(decl.Name), decl.Kind, decl)
}

func (decl *Decl) Print() {
	decl.Fprint(os.Stdout)
}

const _spaces = "                                "

func spaces(name string) string {
	return _spaces[len(name)%32:]
}
