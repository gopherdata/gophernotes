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
 * interpreter.go
 *
 *  Created on: Apr 02, 2017
 *      Author: Massimiliano Ghilardi
 */

package fast

import (
	"bufio"
	"errors"
	"fmt"
	"go/types"
	"io"
	"os"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/gls"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// Interp is the fast interpreter.
// It contains both the tree-of-closures builder Comp
// and the interpreter's runtime environment Env
type Interp struct {
	Comp *Comp
	env  *Env // not exported. to access it, call Interp.PrepareEnv()
}

func New() *Interp {
	top := newTopInterp("builtin")
	top.env.UsedByClosure = true // do not free this *Env
	file := NewInnerInterp(top, "main", "main")
	file.env.UsedByClosure = true // do not free this *Env
	return file
}

func newTopInterp(path string) *Interp {
	name := FileName(path)

	g := NewIrGlobals()
	universe := xr.NewUniverse()

	cg := &CompGlobals{
		IrGlobals:    g,
		Universe:     universe,
		KnownImports: make(map[string]*Import),
		interf2proxy: make(map[r.Type]r.Type),
		proxy2interf: make(map[r.Type]xr.Type),
		Prompt:       "gomacro> ",
	}
	goid := gls.GoID()
	run := &Run{IrGlobals: g, goid: goid}
	// early register run in goroutine-local data
	g.gls[goid] = run

	ir := &Interp{
		Comp: &Comp{
			CompGlobals: cg,
			CompBinds: CompBinds{
				Name: name,
				Path: path,
			},
			UpCost: 1,
			Depth:  0,
			Outer:  nil,
		},
		env: &Env{
			Outer: nil,
			Run:   run,
		},
	}
	// tell xreflect about our packages "fast" and "main"
	universe.CachePackage(types.NewPackage("fast", "fast"))
	universe.CachePackage(types.NewPackage("main", "main"))

	// no need to scavenge for Builtin, Function,  Macro and UntypedLit fields and methods.
	// actually, making them opaque helps securing against malicious interpreted code.
	for _, rtype := range []r.Type{rtypeOfBuiltin, rtypeOfFunction, rtypeOfPtrImport, rtypeOfMacro} {
		cg.opaqueType(rtype, "fast")
	}
	cg.opaqueType(rtypeOfUntypedLit, "untyped")

	ir.addBuiltins()
	return ir
}

func NewInnerInterp(outer *Interp, name string, path string) *Interp {
	if len(name) == 0 {
		name = FileName(path)
	}

	outerComp := outer.Comp
	outerEnv := outer.env
	run := outerEnv.Run

	env := &Env{
		Outer:     outerEnv,
		Run:       run,
		FileEnv:   outerEnv.FileEnv,
		CallDepth: outerEnv.CallDepth,
	}

	if outerEnv.Outer == nil {
		env.FileEnv = env
	} else {
		env.FileEnv = outerEnv.FileEnv
	}

	// do NOT set g.CurrEnv = ir.Env, it messes up the call stack
	return &Interp{
		&Comp{
			CompGlobals: outerComp.CompGlobals,
			CompBinds: CompBinds{
				Name: name,
				Path: path,
			},
			UpCost: 1,
			Depth:  outerComp.Depth + 1,
			Outer:  outerComp,
		},
		env,
	}
}

func (ir *Interp) SetInspector(inspector Inspector) {
	ir.Comp.Globals.Inspector = inspector
}

func (ir *Interp) SetDebugger(debugger Debugger) {
	ir.env.Run.Debugger = debugger
}

func (ir *Interp) Interrupt(os.Signal) {
	ir.env.Run.interrupt()
}

// ============================================================================

// DeclConst compiles a constant declaration
func (ir *Interp) DeclConst(name string, t xr.Type, value I) {
	ir.Comp.DeclConst0(name, t, value)
}

// DeclFunc compiles a function declaration
func (ir *Interp) DeclFunc(name string, fun I) {
	ir.Comp.DeclFunc0(name, fun)
	ir.apply()
}

// DeclBuiltin compiles a builtin function declaration
func (ir *Interp) DeclBuiltin(name string, builtin Builtin) {
	ir.Comp.DeclBuiltin0(name, builtin)
}

// DeclEnvFunc compiles a function declaration that accesses interpreter's *CompEnv
func (ir *Interp) DeclEnvFunc(name string, function Function) {
	ir.Comp.DeclEnvFunc0(name, function)
	ir.apply()
}

// DeclType declares a type
func (ir *Interp) DeclType(t xr.Type) {
	ir.Comp.DeclType0(t)
}

// DeclType declares a type alias
func (ir *Interp) DeclTypeAlias(alias string, t xr.Type) {
	ir.Comp.DeclTypeAlias0(alias, t)
}

// DeclVar compiles a variable declaration
func (ir *Interp) DeclVar(name string, t xr.Type, value I) {
	if t == nil {
		t = ir.Comp.TypeOf(value)
	}
	ir.Comp.DeclVar0(name, t, ir.Comp.exprValue(t, value))
	ir.apply()
}

// apply executes the compiled declarations, statements and expressions,
// then clears the compiled buffer
func (ir *Interp) apply() {
	exec := ir.Comp.Code.Exec()
	if exec != nil {
		exec(ir.PrepareEnv())
	}
}

// AddressOfVar compiles the expression &name, then executes it
// returns the zero value if name is not found or is not addressable
func (ir *Interp) AddressOfVar(name string) (addr r.Value) {
	c := ir.Comp
	sym := c.TryResolve(name)
	var v r.Value
	if sym != nil {
		switch sym.Desc.Class() {
		case VarBind, IntBind:
			va := sym.AsVar(PlaceAddress)
			expr := va.Address(c.Depth)
			v, _ = ir.RunExpr1(expr)
		}
	}
	return v
}

// replacement of reflect.TypeOf() that uses xreflect.TypeOf()
func (ir *Interp) TypeOf(val interface{}) xr.Type {
	return ir.Comp.TypeOf(val)
}

// ValueOf retrieves the value of a constant, function or variable
// The returned value is settable and addressable only for variables
// returns the zero value if name is not found
func (ir *Interp) ValueOf(name string) (value r.Value) {
	sym := ir.Comp.TryResolve(name)
	if sym == nil {
		return Nil
	}
	switch sym.Desc.Class() {
	case ConstBind:
		return sym.Bind.ConstValue()
	case IntBind:
		value = ir.AddressOfVar(name)
		if value.IsValid() {
			value = value.Elem() // dereference
		}
		return value
	default:
		env := ir.PrepareEnv()
		for i := 0; i < sym.Upn; i++ {
			env = env.Outer
		}
		return env.Vals[sym.Desc.Index()]
	}
}

// ===================== Eval(), EvalFile(), EvalReader() ============================

// combined Parse + Compile + RunExpr1
func (ir *Interp) Eval1(src string) (r.Value, xr.Type) {
	return ir.RunExpr1(ir.Compile(src))
}

// combined Parse + Compile + RunExpr
func (ir *Interp) Eval(src string) ([]r.Value, []xr.Type) {
	return ir.RunExpr(ir.Compile(src))
}

func (ir *Interp) EvalFile(filepath string) (comments string, err error) {
	g := ir.Comp.CompGlobals
	saveFilename := g.Filepath
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer func() {
		f.Close()
		g.Filepath = saveFilename
	}()
	g.Filepath = filepath
	return ir.EvalReader(f)
}

func (ir *Interp) EvalReader(src io.Reader) (comments string, err error) {
	g := ir.Comp.CompGlobals
	savein := g.Readline
	saveopts := g.Options
	g.Line = 0
	in := MakeBufReadline(bufio.NewReader(src), g.Stdout)
	g.Readline = in
	// parsing a file: suppress prompt and printing expression results
	g.Options &^= OptShowPrompt | OptShowEval | OptShowEvalType
	defer func() {
		g.Readline = savein
		g.Options = saveopts
		if rec := recover(); rec != nil {
			switch rec := rec.(type) {
			case error:
				err = rec
			default:
				err = errors.New(fmt.Sprint(rec))
			}
		}
	}()

	// perform the first iteration manually, to collect comments
	str, firstToken := g.ReadMultiline(ReadOptCollectAllComments, g.Prompt)
	if firstToken >= 0 {
		comments = str[0:firstToken]
		if firstToken > 0 {
			str = str[firstToken:]
			g.IncLine(comments)
		}
	}

	if ir.ParseEvalPrint(str) {
		for ir.ReadParseEvalPrint() {
		}
	}
	return comments, nil
}
