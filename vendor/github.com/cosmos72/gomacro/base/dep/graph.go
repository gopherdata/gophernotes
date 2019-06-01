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
 * graph.go
 *
 *  Created on: May 03, 2018
 *      Author: Massimiliano Ghilardi
 */

package dep

import (
	"bytes"
	"fmt"
	"go/token"

	"github.com/cosmos72/gomacro/base/output"
)

func (f *fwdDeclList) add(decl *Decl) {
	name := decl.Name
	if _, ok := f.Set[name]; ok {
		return
	}
	fwd := *decl
	fwd.Kind = TypeFwd
	f.List = append(f.List, &fwd)
	f.Set[name] = void{}
}

const DEBUG_GRAPH = false

func (g *graph) Sort() DeclList {
	g.RemoveUnresolvableDeps()

	if DEBUG_GRAPH {
		fmt.Print("---- all decls ----\n")
		g.Nodes.Print()
	}

	sorted := make(DeclList, 0, len(g.Nodes))
	fwd := fwdDeclList{Set: make(set)}

	for len(g.Nodes) != 0 {
		buf := g.RemoveNodesNoDeps()
		if len(buf) == 0 {
			buf = g.RemoveTypeFwd()
			if len(buf) == 0 {
				g.circularDependencyError()
			}
		}
		g.RemoveUnresolvableDeps()
		sorted = append(sorted, buf.SortByPos()...)
	}

	if len(fwd.List) != 0 {
		sorted = append(fwd.List, sorted...)
	}
	return sorted
}

// remove from g.Nodes the nodes that have no dependencies and return them.
// Implementation choice: remove at most a single name -> better preserves source code ordering
func (g *graph) RemoveNodesNoDeps() DeclList {
	var ret DeclList
	var pos token.Pos
	var retname string
	for name, list := range g.Nodes {
		if len(g.Edges[name]) == 0 {
			for _, decl := range list {
				// among nodes with no dependencies, choose the one with smallest Pos
				if ret == nil || decl.Pos < pos {
					ret = list
					pos = decl.Pos
					retname = name
					break
				}
			}
		}
	}
	if ret == nil {
		return nil
	}
	delete(g.Edges, retname)
	delete(g.Nodes, retname)
	return ret
}

// remove from g.Edges dependencies that are not in g.Nodes
func (g *graph) RemoveUnresolvableDeps() {
	for name := range g.Nodes {
		if edges, ok := g.Edges[name]; ok {
			for edge := range edges {
				if _, ok := g.Nodes[edge]; !ok {
					// node not in g.Nodes, drop the edge
					delete(edges, edge)
				}
			}
		}
	}
}

// remove from g.Edges dependencies that are in m
func (g *graph) RemoveDeps(m DeclMap) {
	for name := range g.Nodes {
		if edges, ok := g.Edges[name]; ok {
			for edge := range edges {
				if _, ok := m[edge]; ok {
					// node in m, drop the edge
					delete(edges, edge)
				}
			}
		}
	}
}

// for nodes with Kind 'k', remove from g.Edges dependencies that are in m
func (g *graph) RemoveDepsFor(k Kind, m DeclMap) {
	for name, list := range g.Nodes {
		for _, decl := range list {
			if decl.Kind != k {
				continue
			}
			if edges, ok := g.Edges[name]; ok {
				for edge := range edges {
					if _, ok := m[edge]; ok {
						// node in m, drop the edge
						delete(edges, edge)
					}
				}
			}
		}
	}
}

// return forward declarations for some types that hopefully break
// as many circular dependencies as possible
func (g *graph) RemoveTypeFwd() DeclList {
	ctx := visitCtx{
		visiting: make(map[string]int),
		visited:  make(map[string]int),
		cycleFunc: func(node *Decl, ctx *visitCtx) {
			ctx.visiting[node.Name]++
		},
	}
	for _, node := range g.Nodes.List().SortByPos().Reverse() {
		if len(ctx.visited) == len(g.Nodes) {
			break
		}
		g.visit(node, &ctx)
	}
	var list DeclList
	most := 1
	for name, count := range ctx.visited {
		for _, decl := range g.Nodes[name] {
			if decl == nil || decl.Kind != Type || count < most {
				continue
			}
			if count > most {
				list = nil // discard  types collected so far
			}
			most = count
			list = append(list, decl)
		}
	}
	if len(list) == 0 {
		return nil
	}
	// change Kind of returned Decls to TypeFwd
	for i, e := range list {
		fwd := *e
		fwd.Kind = TypeFwd
		list[i] = &fwd
	}
	g.RemoveDepsFor(Type, list.Map())
	return list
}

func (g *graph) visit(node *Decl, ctx *visitCtx) {
	name := node.Name
	if _, ok := ctx.visited[name]; ok {
		return
	}
	if _, ok := ctx.visiting[name]; ok {
		if fun := ctx.cycleFunc; fun != nil {
			fun(node, ctx)
		}
		return
	}
	if fun := ctx.beforeFunc; fun != nil {
		fun(node, ctx)
	}
	ctx.visiting[name] = 0
	for name := range g.Edges[name] {
		for _, node := range g.Nodes[name] {
			g.visit(node, ctx)
		}
	}
	ctx.visited[name] = ctx.visiting[name]
	delete(ctx.visiting, name)
	if fun := ctx.afterFunc; fun != nil {
		fun(node, ctx)
	}
}

func (g *graph) circularDependencyError() {
	var path, cycle []string

	ctx := visitCtx{
		visiting: make(map[string]int),
		visited:  make(map[string]int),
		beforeFunc: func(node *Decl, ctx *visitCtx) {
			path = append(path, node.Name)
		},
		afterFunc: func(node *Decl, ctx *visitCtx) {
			path = path[:len(path)-1]
		},
		cycleFunc: func(node *Decl, ctx *visitCtx) {
			// collect the shortest cycle
			name := node.Name
			temp := dup(append(path, name))
			for len(temp) != 0 {
				if temp[0] == name {
					break
				}
				temp = temp[1:]
			}
			if len(cycle) != 0 && len(temp) >= len(cycle) {
				return
			}
			cycle = temp
		},
	}
	for _, node := range g.Nodes.List().SortByPos() {
		if len(ctx.visited) == len(g.Nodes) {
			break
		}
		g.visit(node, &ctx)
	}

	var buf bytes.Buffer // strings.Builder requires Go >= 1.10

	buf.WriteString("declaration loop\n")

	if len(cycle) != 0 {
		for i, name := range cycle[1:] {
			fmt.Fprintf(&buf, "\t%s uses %s\n", cycle[i], name)
		}
	}
	output.Errorf("%s", buf.String())
}
