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
 * api.go
 *
 *  Created on: May 03, 2018
 *      Author: Massimiliano Ghilardi
 */

package dep

import (
	"fmt"
	"go/ast"
	"go/token"
)

// Support for out-of-order code

type Kind int

const (
	Unknown Kind = iota
	Const
	Expr
	Func
	Import
	Macro
	Method
	Package
	Stmt
	Type
	TypeFwd
	Var
	VarMulti
)

var kinds = map[Kind]string{
	Unknown:  "Unknown",
	Const:    "Const",
	Expr:     "Expr",
	Func:     "Func",
	Import:   "Import",
	Macro:    "Macro",
	Method:   "Method",
	Package:  "Package",
	Stmt:     "Stmt",
	Type:     "Type",
	TypeFwd:  "TypeFwd", // forward type declaration
	Var:      "Var",
	VarMulti: "VarMulti", // variables initialized with multi-value expression
}

func (k Kind) String() string {
	name, ok := kinds[k]
	if ok {
		return name
	}
	return fmt.Sprintf("Kind%d", int(k))
}

// for multiple const or var declarations in a single *ast.ValueSpec
type Extra struct {
	Ident *ast.Ident
	Type  ast.Expr
	Value ast.Expr
	Iota  int // for constants, value of iota to use
}

// convert *Extra to ast.Spec
func (extra *Extra) Spec() *ast.ValueSpec {
	spec := &ast.ValueSpec{
		Names: []*ast.Ident{extra.Ident},
		Type:  extra.Type,
	}
	if extra.Value != nil {
		spec.Values = []ast.Expr{extra.Value}
	}
	return spec
}

type Decl struct {
	Kind  Kind
	Name  string
	Node  ast.Node // nil for multiple const or var declarations in a single *ast.ValueSpec - in such case, see Extra
	Deps  []string // names of types, constants and variables used in Node's declaration
	Pos   token.Pos
	Extra *Extra
}

type DeclList []*Decl

func NewDecl(kind Kind, name string, node ast.Node, pos token.Pos, deps []string) *Decl {
	return &Decl{Kind: kind, Name: name, Node: node, Deps: sort_unique_inplace(deps), Pos: pos}
}

type DeclMap map[string]DeclList

type Scope struct {
	Decls  DeclMap
	Outer  *Scope
	Gensym int
}

func NewScope(outer *Scope) *Scope {
	return &Scope{
		Decls: make(DeclMap),
		Outer: outer,
	}
}

type Sorter struct {
	scope Scope
	queue []ast.Node
}

func NewSorter() *Sorter {
	return &Sorter{
		scope: Scope{
			Decls: make(DeclMap),
		},
	}
}

// Sorter resolves top-level constant, type, function and var
// declaration order by analyzing their dependencies.
//
// also resolves top-level var initialization order
// analyzing their dependencies.
