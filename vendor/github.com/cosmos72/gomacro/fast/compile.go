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
 * compile.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/dep"
	"github.com/cosmos72/gomacro/gls"
)

func NewComp(outer *Comp, code *Code) *Comp {
	if outer == nil {
		return &Comp{UpCost: 1}
	}
	c := Comp{
		UpCost:      1,
		Depth:       outer.Depth + 1,
		Outer:       outer,
		CompGlobals: outer.CompGlobals,
	}
	// Debugf("NewComp(%p->%p) %s", outer, &c, debug.Stack())
	if code != nil {
		c.Code = *code
	}
	return &c
}

func (c *Comp) TopComp() *Comp {
	for ; c != nil; c = c.Outer {
		if c.Outer == nil {
			break
		}
	}
	return c
}

func (c *Comp) FileComp() *Comp {
	for ; c != nil; c = c.Outer {
		outer := c.Outer
		if outer == nil || outer.Outer == nil {
			break
		}
	}
	return c
}

func NewIrGlobals() *IrGlobals {
	return &IrGlobals{
		gls:     make(map[uintptr]*Run),
		Globals: *NewGlobals(),
	}
}

func (g *IrGlobals) glsGet(goid uintptr) *Run {
	g.lock.Lock()
	ret := g.gls[goid]
	g.lock.Unlock()
	return ret
}

func (run *Run) getRun4Goid(goid uintptr) *Run {
	g := run.IrGlobals
	ret := g.glsGet(goid)
	if ret == nil {
		ret = run.new(goid)
		ret.glsStore()
	}
	return ret
}

func (tg *Run) glsStore() {
	g := tg.IrGlobals
	goid := tg.goid
	g.lock.Lock()
	g.gls[goid] = tg
	g.lock.Unlock()
}

func (tg *Run) glsDel() {
	g := tg.IrGlobals
	goid := tg.goid
	g.lock.Lock()
	delete(g.gls, goid)
	g.lock.Unlock()
}

func (run *Run) new(goid uintptr) *Run {
	return &Run{
		IrGlobals: run.IrGlobals,
		goid:      goid,
		// Interrupt, Signal, PoolSize and Pool are zero-initialized, fine with that
	}
}

// common part between NewEnv() and newEnv4Func()
func newEnv(run *Run, outer *Env, nbind int, nintbind int) *Env {
	pool := &run.Pool // pool is an array, do NOT copy it!
	index := run.PoolSize - 1
	var env *Env
	if index >= 0 {
		run.PoolSize = index
		env = pool[index]
		pool[index] = nil
	} else {
		env = &Env{}
	}
	if cap(env.Vals) >= nbind {
		env.Vals = env.Vals[0:nbind]
	} else {
		env.Vals = make([]r.Value, nbind)
	}
	if cap(env.Ints) >= nintbind {
		env.Ints = env.Ints[0:nintbind]
	} else {
		env.Ints = make([]uint64, nintbind)
	}
	env.Outer = outer
	env.Run = run
	env.FileEnv = outer.FileEnv
	return env
}

// return a new, nested Env with given number of binds and intbinds
func NewEnv(outer *Env, nbind int, nintbind int) *Env {
	run := outer.Run

	// manually inline
	// env := newEnv(run, outer, nbind, nintbind)
	var env *Env
	{
		pool := &run.Pool // pool is an array, do NOT copy it!
		index := run.PoolSize - 1
		if index >= 0 {
			run.PoolSize = index
			env = pool[index]
			pool[index] = nil
		} else {
			env = &Env{}
		}
		if cap(env.Vals) >= nbind {
			env.Vals = env.Vals[0:nbind]
		} else {
			env.Vals = make([]r.Value, nbind)
		}
		if cap(env.Ints) >= nintbind {
			env.Ints = env.Ints[0:nintbind]
		} else {
			env.Ints = make([]uint64, nintbind)
		}
		env.Outer = outer
		env.Run = run
		env.FileEnv = outer.FileEnv
	}
	env.IP = outer.IP
	env.Code = outer.Code
	env.DebugPos = outer.DebugPos
	env.CallDepth = outer.CallDepth
	// this is a nested *Env, not a function body: to obtain the caller function,
	// follow env.Outer.Outer... chain until you find an *Env with non-nil Caller
	// env.Caller = nil
	// DebugCallStack Debugf("NewEnv(%p->%p) nbind=%d nintbind=%d calldepth: %d->%d", outer, env, nbind, nintbind, outer.CallDepth, env.CallDepth)
	run.CurrEnv = env
	return env
}

func newEnv4Func(outer *Env, nbind int, nintbind int, debugComp *Comp) *Env {
	goid := gls.GoID()
	run := outer.Run
	if run.goid != goid {
		// no luck... get the correct ThreadGlobals for goid
		run = run.getRun4Goid(goid)
	}
	// manually inline
	// env := newEnv(run, outer, nbind, nintbind)
	var env *Env
	{
		pool := &run.Pool // pool is an array, do NOT copy it!
		index := run.PoolSize - 1
		if index >= 0 {
			run.PoolSize = index
			env = pool[index]
			pool[index] = nil
		} else {
			env = &Env{}
		}
		if cap(env.Vals) >= nbind {
			env.Vals = env.Vals[0:nbind]
		} else {
			env.Vals = make([]r.Value, nbind)
		}
		if cap(env.Ints) >= nintbind {
			env.Ints = env.Ints[0:nintbind]
		} else {
			env.Ints = make([]uint64, nintbind)
		}
		env.Outer = outer
		env.Run = run
		env.FileEnv = outer.FileEnv
	}
	env.DebugComp = debugComp
	caller := run.CurrEnv
	env.Caller = caller
	if caller == nil {
		env.CallDepth = 1
	} else {
		env.CallDepth = caller.CallDepth + 1
	}
	// DebugCallStack Debugf("newEnv4Func(%p->%p) nbind=%d nintbind=%d calldepth: %d->%d", caller, env, nbind, nintbind, env.CallDepth-1, env.CallDepth)
	run.CurrEnv = env
	return env
}

func (env *Env) MarkUsedByClosure() {
	for ; env != nil && !env.UsedByClosure; env = env.Outer {
		env.UsedByClosure = true
	}
}

// FreeEnv tells the interpreter that given nested *Env is no longer needed.
func (env *Env) FreeEnv() {
	run := env.Run
	run.CurrEnv = env.Outer
	env.freeEnv(run)
}

// freeEnv4Func tells the interpreter that given function body *Env is no longer needed.
func (env *Env) freeEnv4Func() {
	run := env.Run
	run.CurrEnv = env.Caller
	env.freeEnv(run)
}

func (env *Env) freeEnv(run *Run) {
	// DebugCallStack Debugf("FreeEnv(%p->%p), calldepth: %d->%d", env, caller, env.CallDepth, caller.CallDepth)
	if env.UsedByClosure {
		// in use, cannot recycle
		return
	}
	n := run.PoolSize
	if n >= poolCapacity {
		return
	}
	if env.IntAddressTaken {
		env.Ints = nil
		env.IntAddressTaken = false
	}
	env.Outer = nil
	env.Code = nil
	env.DebugPos = nil
	env.DebugComp = nil
	env.Caller = nil
	env.Run = nil
	env.FileEnv = nil
	run.Pool[n] = env // pool is an array, be careful NOT to copy it!
	run.PoolSize = n + 1
}

func (env *Env) Top() *Env {
	if env == nil {
		return nil
	}
	if file := env.FileEnv; file != nil {
		if top := file.Outer; top != nil && top.Outer == nil {
			return top
		}
	}
	for o := env.Outer; o != nil; o = o.Outer {
		env = o
	}
	return env
}

func (env *Env) Up(n int) *Env {
	for ; n >= 3; n -= 3 {
		env = env.Outer.Outer.Outer
	}
	switch n {
	case 2:
		env = env.Outer
		fallthrough
	case 1:
		env = env.Outer
	}
	return env
}

// combined Parse + MacroExpandCodeWalk
func (c *Comp) Parse(src string) Ast {
	// do NOT set c.Globals.Line = 0
	// caller can do it manually if needed
	nodes := c.ParseBytes([]byte(src))
	forms := anyToAst(nodes, "Parse")

	forms, _ = c.MacroExpandCodewalk(forms)
	if c.Options&OptShowMacroExpand != 0 {
		c.Debugf("after macroexpansion: %v", forms.Interface())
	}
	return forms
}

// compile code. support out-of-order declarations
func (c *Comp) Compile(in Ast) *Expr {
	if in == nil {
		return nil
	}
	switch node := in.Interface().(type) {
	case *ast.File, ast.Decl, *ast.ValueSpec:
		// complicated, use general technique below
	case ast.Node:
		// shortcut
		return c.compileNode(node, dep.Unknown)
	}
	// order declarations by topological sort on their dependencies
	sorter := dep.NewSorter()
	sorter.LoadAst(in)

	decls := sorter.All()

	switch n := len(decls); n {
	case 0:
		return nil
	case 1:
		return c.compileDecl(decls[0])
	default:
		exprs := make([]*Expr, 0, n)
		for _, decl := range decls {
			e := c.compileDecl(decl)
			if e != nil {
				exprs = append(exprs, e)
			}
		}
		return exprList(exprs, c.CompileOptions())
	}
	return nil
}

// compile code. support out-of-order declarations too
func (c *Comp) CompileNode(node ast.Node) *Expr {
	return c.Compile(ToAst(node))
}

func (c *Comp) compileDecl(decl *dep.Decl) *Expr {
	if decl == nil {
		return nil
	}
	if extra := decl.Extra; extra != nil {
		// decl.Node may declare multiple constants or variables:
		// do not use it!
		// instead get the single const or var declaration from Extra
		switch decl.Kind {
		case dep.Const:
			// see Comp.GenDecl() in declaration.go for a discussion
			// on the scope where to declare iota, and what to do
			// with any previous declaration of iota in the same scope
			top := c.TopComp()
			defer top.endIota(top.beginIota())
			top.setIota(extra.Iota)

			c.DeclConsts(extra.Spec(), nil, nil)
			return c.Code.AsExpr()
		case dep.Var:
			c.DeclVars(extra.Spec())
			return c.Code.AsExpr()
		}
	}
	if node := decl.Node; node != nil {
		return c.compileNode(node, decl.Kind)
	}
	// may happen for second and later variables in VarMulti,
	// which CANNOT be declared individually
	return nil
}

// compileExpr is a wrapper for Compile
// that guarantees Code does not get clobbered/cleared.
// Used by Comp.Quasiquote
func (c *Comp) compileExpr(in Ast) *Expr {
	cf := NewComp(c, nil)
	cf.UpCost = 0
	cf.Depth--
	return cf.Compile(in)
}

// common backend for Compile, CompileNode, File, compileDecl.
// does NOT support out-of-order declarations
func (c *Comp) compileNode(node ast.Node, kind dep.Kind) *Expr {
	if n := c.Code.Len(); n != 0 {
		c.Warnf("Compile: discarding %d previously compiled statements from code buffer", n)
	}
	if node == nil {
		return nil
	}
	c.Code.Clear()
	c.Loop = nil
	c.Func = nil
	c.Labels = nil
	c.FuncMaker = nil
	c.Pos = node.Pos()
	switch node := node.(type) {
	case ast.Decl:
		c.Decl(node)
	case ast.Expr:
		return c.Expr(node, nil)
	case *ast.ImportSpec:
		// dep.Sorter.Some() returns naked *ast.ImportSpec,
		// instead of *ast.GenDecl containing one or more *ast.ImportSpec as parser does
		c.Import(node)
	case *ast.TypeSpec:
		// dep.Sorter.Some() returns naked *ast.TypeSpec,
		// instead of *ast.GenDecl containing one or more *ast.TypeSpec as parser does
		if kind == dep.TypeFwd {
			// forward type declaration
			c.DeclNamedType(node.Name.Name)
		} else {
			c.DeclType(node)
		}
	case *ast.ExprStmt:
		// special case of statement
		return c.Expr(node.X, nil)
	case ast.Stmt:
		c.Stmt(node)
	case *ast.File:
		// not c.File(node): unnecessary and risks an infinite recursion
		for _, decl := range node.Decls {
			c.Decl(decl)
		}
	default:
		c.Errorf("unsupported node type, expecting <ast.Decl>, <ast.Expr>, <ast.Stmt> or <*ast.File>, found %v <%v>", node, r.TypeOf(node))
		return nil
	}
	return c.Code.AsExpr()
}

// compile file. support out-of-order declarations too
func (c *Comp) File(node *ast.File) {
	if node != nil {
		c.Name = node.Name.Name
		c.Compile(File{node})
	}
}

func (c *Comp) Append(stmt Stmt, pos token.Pos) {
	c.Code.Append(stmt, pos)
}

func (c *Comp) append(stmt Stmt) {
	c.Code.Append(stmt, c.Pos)
}
