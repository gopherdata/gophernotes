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
 * env.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"bufio"
	"fmt"
	"io"
	r "reflect"

	. "github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/imports"
)

type ThreadGlobals struct {
	Globals
	AllMethods map[r.Type]Methods // methods implemented by interpreted code
	FastInterp interface{}        // *fast.Interp // temporary...
}

func NewThreadGlobals() *ThreadGlobals {
	tg := &ThreadGlobals{
		AllMethods: make(map[r.Type]Methods),
	}
	tg.Globals.Init()
	return tg
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
		Name:       path,
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
	fenv := env.FileEnv()
	currpath := fenv.ThreadGlobals.PackagePath
	if path == currpath {
		return env
	}
	fenv.AsPackage().SaveToPackages(currpath)

	nenv := NewEnv(fenv.TopEnv(), path)
	nenv.MergePackage(imports.Packages[path])
	nenv.ThreadGlobals.PackagePath = path

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

func (env *Env) ReadMultiline(in *bufio.Reader, opts ReadOptions) (str string, firstToken int) {
	str, firstToken, err := ReadMultiline(in, opts, env.Stdout, "gomacro> ")
	if err != nil && err != io.EOF {
		fmt.Fprintf(env.Stderr, "// read error: %s\n", err)
	}
	return str, firstToken
}

// macroexpand + collect + eval
func (env *Env) classicEval(form Ast) (r.Value, []r.Value) {
	// macroexpansion phase.
	form, _ = env.MacroExpandAstCodewalk(form)

	if env.Options&OptShowMacroExpand != 0 {
		env.Debugf("after macroexpansion: %v", form.Interface())
	}

	// collect phase
	if env.Options&(OptCollectDeclarations|OptCollectStatements) != 0 {
		env.CollectAst(form)
	}

	// eval phase
	if env.Options&OptMacroExpandOnly != 0 {
		return r.ValueOf(form.Interface()), nil
	} else {
		return env.EvalAst(form)
	}
}
