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
 * api_internal.go
 *
 *  Created on: May 05, 2018
 *      Author: Massimiliano Ghilardi
 */

package dep

type void struct{}

type set map[string]void

type depMap map[string]set

type fwdDeclList struct {
	List DeclList
	Set  set
}

type graph struct {
	Nodes DeclMap
	Edges depMap
}

type visitCtx struct {
	visiting   map[string]int
	visited    map[string]int
	beforeFunc func(node *Decl, ctx *visitCtx) // invoked once for each node, in visit pre-order
	afterFunc  func(node *Decl, ctx *visitCtx) // invoked once for each node, in visit post-order
	cycleFunc  func(node *Decl, ctx *visitCtx) // invoked when ctx.visiting[node.Name] exists already, i.e. for cycles
}
