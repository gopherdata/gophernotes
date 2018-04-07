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
	"github.com/cosmos72/gomacro/imports"
)

var (
	nilEnv *Env
	NilEnv = []r.Value{r.ValueOf(nilEnv)}
)

func (ir *ThreadGlobals) ShowHelp(out io.Writer) {
	fmt.Fprint(out, `// type Go code to execute it. example: func add(x, y int) int { return x + y }

// interpreter commands:
:classic CODE   execute CODE using the classic interpreter
:env [name]     show available functions, variables and constants
                in current package, or from imported package "name"
:fast CODE      execute CODE using the fast interpreter (default)
:help           show this help
:inspect EXPR   inspect expression interactively
:options [OPTS] show or toggle interpreter options
:quit           quit the interpreter
:write [FILE]   write collected declarations and/or statements to standard output or to FILE
                use :o Declarations and/or :o Statements to start collecting them
`)
}

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
			env.Debugf("%d:\t     %v, runningDefers = %v, panic is nil", i, name, frame.runningDefers)
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
			env.showPackage(e.Name, &pkg)
		}
		return
	}
	bind, ok := env.resolveIdentifier(&ast.Ident{Name: packageName})
	if !ok {
		env.Warnf("not an imported package: %q", packageName)
		return
	}
	val, ok := bind.Interface().(*PackageRef)
	if !ok {
		env.Warnf("not an imported package: %q = %v <%v>", packageName, val, typeOf(bind))
		return
	}
	env.showPackage(val.Name, &val.Package)
}

func (env *Env) showPackage(name string, pkg *imports.Package) {
	const spaces15 = "               "
	out := env.Stdout
	binds := pkg.Binds
	if len(binds) > 0 {
		fmt.Fprintf(out, "// ----- %s binds -----\n", name)

		keys := make([]string, len(binds))
		i := 0
		for k := range binds {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			n := len(k) & 15
			fmt.Fprintf(out, "%s%s = ", k, spaces15[n:])
			bind := binds[k]
			if bind != Nil {
				switch bind := bind.Interface().(type) {
				case *Env:
					fmt.Fprintf(out, "%p // %v\n", bind, r.TypeOf(bind))
					continue
				}
			}
			env.Fprintf(out, "%v // %v\n", bind, r.TypeOf(bind))
		}
		fmt.Fprintln(out)
	}
	types := pkg.Types
	if len(types) > 0 {
		fmt.Fprintf(out, "// ----- %s types -----\n", name)

		keys := make([]string, len(types))
		i := 0
		for k := range types {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			n := len(k) & 15
			t := types[k]
			fmt.Fprintf(out, "%s%s = %v\t// %v\n", k, spaces15[n:], t, t.Kind())
		}
		fmt.Fprintln(out)
	}
}
