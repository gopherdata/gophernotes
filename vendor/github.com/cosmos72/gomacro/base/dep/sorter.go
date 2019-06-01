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
 * sorter.go
 *
 *  Created on: May 03, 2018
 *      Author: Massimiliano Ghilardi
 */

package dep

import (
	"go/ast"
	"go/token"

	"github.com/cosmos72/gomacro/ast2"
)

func (s *Sorter) LoadNode(node ast.Node) {
	s.LoadAst(ast2.ToAst(node))
}

func (s *Sorter) LoadNodes(nodes []ast.Node) {
	s.LoadAst(ast2.NodeSlice{nodes})
}

func (s *Sorter) LoadAst(form ast2.Ast) {
	s.queue = ast2.ToNodesAppend(s.queue, form)
}

// return one of:
// * a list of imports
// * a list of declarations
// * a list of expressions and statements
func (s *Sorter) Some() DeclList {

	decls := s.popPackages()
	if len(decls) == 0 {
		decls = s.popImports()
	}
	if len(decls) == 0 {
		decls = s.popDecls()
	}
	if len(decls) == 0 {
		decls = s.popStmts()
	}
	return decls
}

func (s *Sorter) All() DeclList {
	var all DeclList

	for {
		decls := s.Some()
		if len(decls) == 0 {
			break
		}
		all = append(all, decls...)
	}
	return all
}

func (s *Sorter) popPackages() []*Decl {
	var list DeclList
	i, n := 0, len(s.queue)
loop:
	for ; i < n; i++ {
		node := s.queue[i]
		switch node := node.(type) {
		case nil:
			continue
		case *ast.GenDecl:
			if node != nil && node.Tok == token.PACKAGE {
				for _, spec := range node.Specs {
					list = append(list, NewDeclPackage(spec, &s.scope.Gensym))
				}
				continue
			}
		}
		// /*DELETEME*/ fmt.Printf("popPackages stopping at node: %v %T\n", node, node)
		break loop
	}
	if i > 0 {
		s.queue = s.queue[i:]
	}
	if len(list) == 0 {
		return nil
	}
	return list.SortByPos()
}

func (s *Sorter) popImports() []*Decl {
	var list DeclList
	i, n := 0, len(s.queue)
loop:
	for ; i < n; i++ {
		node := s.queue[i]
		switch node := node.(type) {
		case nil:
			continue
		case *ast.GenDecl:
			if node != nil && node.Tok == token.IMPORT {
				for _, spec := range node.Specs {
					list = append(list, NewDeclImport(spec, &s.scope.Gensym))
				}
				continue
			}
		}
		// /*DELETEME*/ fmt.Printf("popImports stopping at node: %v %T\n", node, node)
		break loop
	}
	if i > 0 {
		s.queue = s.queue[i:]
	}
	if len(list) == 0 {
		return nil
	}
	return list.SortByPos()
}

func (s *Sorter) popDecls() []*Decl {
	var nodes []ast.Node
	i, n := 0, len(s.queue)
loop:
	for ; i < n; i++ {
		node := s.queue[i]
		switch node := node.(type) {
		case nil:
			continue
		case *ast.GenDecl:
			if node != nil && node.Tok != token.IMPORT && node.Tok != token.PACKAGE {
				nodes = append(nodes, node)
				continue
			}
		case ast.Decl:
			if node != nil {
				nodes = append(nodes, node)
				continue
			}
		}
		// /*DELETEME*/ fmt.Printf("popDecls stopping at node: %v %T\n", node, node)
		break loop
	}
	if i > 0 {
		s.queue = s.queue[i:]
	}
	if len(nodes) == 0 {
		return nil
	}
	s.scope.Decls = make(DeclMap)

	s.scope.Nodes(nodes)
	s.scope.Decls.RemoveUnresolvableDeps()
	m := s.scope.Decls.Dup()

	s.scope.Decls = nil

	g := graph{
		Nodes: m,
		Edges: m.depMap(),
	}
	return g.Sort()
}

func (s *Sorter) popStmts() []*Decl {
	var list DeclList
	i, n := 0, len(s.queue)
loop:
	for ; i < n; i++ {
		node := s.queue[i]
		switch node := node.(type) {
		case nil:
			continue
		case ast.Expr:
			if node != nil {
				list = append(list, NewDeclExpr(node, &s.scope.Gensym))
				continue
			}
		case ast.Stmt:
			if node != nil {
				list = append(list, NewDeclStmt(node, &s.scope.Gensym))
				continue
			}
		}
		// /*DELETEME*/ fmt.Printf("popStmts stopping at node: %v %T\n", node, node)
		break loop
	}
	if i > 0 {
		s.queue = s.queue[i:]
	}
	if len(list) == 0 {
		return nil
	}
	return list.SortByPos()
}
