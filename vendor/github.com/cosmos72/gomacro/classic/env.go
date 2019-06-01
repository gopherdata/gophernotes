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
 * env.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"go/ast"
	r "reflect"

	"github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/paths"
	"github.com/cosmos72/gomacro/imports"
)

type ThreadGlobals struct {
	*Globals
	AllMethods map[r.Type]Methods // methods implemented by interpreted code
	currOpt    CmdOpt
}

func NewThreadGlobals() *ThreadGlobals {
	return &ThreadGlobals{
		Globals:    NewGlobals(),
		AllMethods: make(map[r.Type]Methods),
	}
}

type Env struct {
	*ThreadGlobals
	Binds      BindMap
	Types      TypeMap
	Proxies    TypeMap
	Outer      *Env
	CallStack  *CallStack
	iotaOffset int
	Name, Path string
}

func NewEnv(outer *Env, path string) *Env {
	env := &Env{
		iotaOffset: 1,
		Outer:      outer,
		Name:       paths.FileName(path),
		Path:       path,
	}
	if outer == nil {
		env.ThreadGlobals = NewThreadGlobals()
		env.CallStack = &CallStack{Frames: []CallFrame{CallFrame{}}}
		env.addBuiltins()
		env.addInterpretedBuiltins()
	} else {
		env.ThreadGlobals = outer.ThreadGlobals
		env.CallStack = outer.CallStack
	}
	return env
}

func (env *Env) TopEnv() *Env {
	for ; env != nil; env = env.Outer {
		if env.Outer == nil {
			break
		}
	}
	return env
}

func (env *Env) FileEnv() *Env {
	for ; env != nil; env = env.Outer {
		outer := env.Outer
		if outer == nil || outer.Outer == nil {
			break
		}
	}
	return env
}

func (env *Env) AsPackage() imports.Package {
	return imports.Package{
		Binds:   env.Binds.AsMap(),
		Types:   env.Types.AsMap(),
		Proxies: env.Proxies.AsMap(),
	}
}

func (env *Env) MergePackage(pkg imports.Package) {
	env.Binds.Ensure().Merge(pkg.Binds)
	env.Types.Ensure().Merge(pkg.Types)
	env.Proxies.Ensure().Merge(pkg.Proxies)
}

func (env *Env) ChangePackage(path string) *Env {
	g := env.ThreadGlobals
	currpath := g.PackagePath
	if path == currpath {
		return env
	}
	fenv := env.FileEnv()
	if fenv.ThreadGlobals != g {
		env.Warnf("ChangePackage: env.ThreadGlobals = %#v\n\tenv.FileEnv().ThreadGlobals = %#v", g, fenv.ThreadGlobals)
	}

	// FIXME really store into imports.Packages fenv's interpreted functions, types, variable and constants ?
	// We need a way to find fenv by name later, but storing it in imports.Packages seems excessive.
	imports.Packages.MergePackage(currpath, fenv.AsPackage())

	nenv := NewEnv(fenv.TopEnv(), path)
	nenv.MergePackage(imports.Packages[path])
	nenv.ThreadGlobals = env.ThreadGlobals
	nenv.ThreadGlobals.PackagePath = path

	if env.Globals.Options&OptShowPrompt != 0 {
		env.Debugf("switched to package %q\n%s", path)
	}

	return nenv
}

// CurrentFrame returns the CallFrame representing the current function call
func (env *Env) CurrentFrame() *CallFrame {
	if env != nil {
		frames := env.CallStack.Frames
		if n := len(frames); n > 0 {
			return &frames[n-1]
		}
	}
	return nil
}

// CallerFrame returns the CallFrame representing the caller's function.
// needed by recover()
func (env *Env) CallerFrame() *CallFrame {
	if env != nil {
		frames := env.CallStack.Frames
		if n := len(frames); n > 1 {
			return &frames[n-2]
		}
	}
	return nil
}

// ValueOf returns the value of a constant, function or variable.
// for variables, the returned reflect.Value is settable and addressable
// returns the zero reflect.Value if not found
func (env *Env) ValueOf(name string) (value r.Value) {
	found := false
	for e := env; e != nil; e = e.Outer {
		if value, found = e.Binds.Get(name); found {
			break
		}
	}
	return
}

// parse, without macroexpansion
func (env *Env) ParseOnly(src interface{}) ast2.Ast {
	var form ast2.Ast
	switch src := src.(type) {
	case ast2.Ast:
		form = src
	case ast.Node:
		form = ast2.ToAst(src)
	default:
		bytes := ReadBytes(src)
		nodes := env.ParseBytes(bytes)

		if env.Options&OptShowParse != 0 {
			env.Debugf("after parse: %v", nodes)
		}
		switch len(nodes) {
		case 0:
			form = nil
		case 1:
			form = ast2.ToAst(nodes[0])
		default:
			form = ast2.NodeSlice{X: nodes}
		}
	}
	return form
}

// Parse, with macroexpansion
func (env *Env) Parse(src interface{}) ast2.Ast {
	form := env.ParseOnly(src)

	// macroexpansion phase.
	form, _ = env.MacroExpandAstCodewalk(form)

	if env.Options&OptShowMacroExpand != 0 {
		env.Debugf("after macroexpansion: %v", form.Interface())
	}
	if env.Options&(OptCollectDeclarations|OptCollectStatements) != 0 {
		env.CollectAst(form)
	}
	return form
}
