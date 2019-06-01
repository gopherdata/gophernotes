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
 * output.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"fmt"
	"go/ast"
	"io"
	r "reflect"
	"sort"

	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/genimport"
	"github.com/cosmos72/gomacro/base/output"
	"github.com/cosmos72/gomacro/base/reflect"
	"github.com/cosmos72/gomacro/imports"
)

var (
	nilEnv *Env
	NilEnv = []r.Value{r.ValueOf(nilEnv)}
)

func (env *Env) showStack() {
	frames := env.CallStack.Frames
	n := len(frames)
	for i := 1; i < n; i++ {
		frame := &frames[i]
		name := ""
		if frame.FuncEnv != nil {
			name = frame.FuncEnv.Name
		}
		if frame.panicking {
			env.Debugf("%d:\t     %v, runningDefers = %v, panic = %v", i, name, frame.runningDefers, frame.panick)
		} else {
			env.Debugf("%d:\t     %v, runningDefers = %v", i, name, frame.runningDefers)
		}
	}
}

func (env *Env) ShowPackage(packageName string) {
	if len(packageName) == 0 {
		stack := make([]*Env, 0)
		for e := env; e != nil; e = e.Outer {
			stack = append(stack, e)
		}
		for i := len(stack) - 1; i >= 0; i-- {
			e := stack[i]
			pkg := e.AsPackage()
			env.showPackage(e.Name, e.Path, &pkg)
		}
		return
	}
	bind, ok := env.resolveIdentifier(&ast.Ident{Name: packageName})
	if !ok {
		env.Warnf("not an imported package: %q", packageName)
		return
	}
	val, ok := bind.Interface().(*genimport.PackageRef)
	if !ok {
		env.Warnf("not an imported package: %q = %v <%v>", packageName, val, typeOf(bind))
		return
	}
	env.showPackage(val.Name, val.Path, &val.Package)
}

func (env *Env) showPackage(name string, path string, pkg *imports.Package) {
	const spaces15 = "               "
	out := env.Stdout
	binds := pkg.Binds
	if len(binds) > 0 {
		output.ShowPackageHeader(out, name, path, "binds")

		keys := make([]string, len(binds))
		i := 0
		for k := range binds {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			showValue(out, k, binds[k])
		}
		fmt.Fprintln(out)
	}
	types := pkg.Types
	if len(types) > 0 {
		output.ShowPackageHeader(out, name, path, "types")

		keys := make([]string, len(types))
		i := 0
		for k := range types {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			showType(out, k, types[k])
		}
		fmt.Fprintln(out)
	}
}

const spaces15 = "               "

func showValue(out io.Writer, name string, v r.Value) {
	n := len(name) & 15
	if !v.IsValid() || v == None {
		fmt.Fprintf(out, "%s%s = nil\t// nil\n", name, spaces15[n:])
	} else {
		fmt.Fprintf(out, "%s%s = %v\t// %s\n", name, spaces15[n:], v, reflect.Type(v))
	}
}

func showType(out io.Writer, name string, t r.Type) {
	n := len(name) & 15
	fmt.Fprintf(out, "%s%s = %v\t// %v\n", name, spaces15[n:], t, t.Kind())
}
