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
 * output.go
 *
 *  Created on: Mar 30, 2018
 *      Author: Massimiliano Ghilardi
 */

package fast

import (
	"fmt"
	"go/types"
	"io"
	r "reflect"
	"sort"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (b Builtin) String() string {
	return fmt.Sprintf("%p", b.Compile)
}

func (imp *Import) String() string {
	return fmt.Sprintf("{%s %q, %d binds, %d types}", imp.Name, imp.Path, len(imp.Binds), len(imp.Types))
}

func typestringer(path string) func(xr.Type) string {
	name := base.FileName(path)
	if name == path {
		return xr.Type.String
	}
	qualifier := func(pkg *types.Package) string {
		pkgpath := pkg.Path()
		if pkgpath == path {
			// base.Debugf("replaced package path %q -> %s", path, name)
			return name
		}
		// base.Debugf("keep package path %q, does not match %q", pkgpath, path)
		return pkgpath
	}
	return func(t xr.Type) string {
		return types.TypeString(t.GoType(), qualifier)
	}
}

func (ir *Interp) ShowPackage(name string) {
	if len(name) != 0 {
		ir.ShowImportedPackage(name)
		return
	}
	// show current package and its outer scopes
	stack := make([]*Interp, 0)
	interp := ir
	for {
		stack = append(stack, interp)
		c := interp.Comp
		env := interp.env
		for i := 0; i < c.UpCost && env != nil; i++ {
			env = env.Outer
		}
		c = c.Outer
		if env == nil || c == nil {
			break
		}
		interp = &Interp{c, env}
	}
	for i := len(stack) - 1; i >= 0; i-- {
		stack[i].ShowAsPackage()
	}
}

func (ir *Interp) ShowAsPackage() {
	c := ir.Comp
	env := ir.PrepareEnv()
	out := c.Globals.Stdout
	stringer := typestringer(c.Path)
	if binds := c.Binds; len(binds) > 0 {
		base.ShowPackageHeader(out, c.Name, c.Path, "binds")

		keys := make([]string, len(binds))
		i := 0
		for k := range binds {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			if bind := binds[k]; bind != nil {
				v := bind.RuntimeValue(env)
				showValue(out, k, v, bind.Type, stringer)
			}
		}
		fmt.Fprintln(out)
	}
	showTypes(out, c.Name, c.Path, c.Types, stringer)
}

func (ir *Interp) ShowImportedPackage(name string) {
	var imp *Import
	var ok bool
	if bind := ir.Comp.Binds[name]; bind != nil && bind.Const() && bind.Type != nil && bind.Type.ReflectType() == rtypeOfPtrImport {
		imp, ok = bind.Value.(*Import)
	}
	if !ok {
		ir.Comp.Warnf("not an imported package: %q", name)
		return
	}
	imp.Show(ir.Comp.CompGlobals)
}

func (imp *Import) Show(g *CompGlobals) {
	stringer := typestringer(imp.Path)
	out := g.Stdout
	if binds := imp.Binds; len(binds) > 0 {
		base.ShowPackageHeader(out, imp.Name, imp.Path, "binds")

		keys := make([]string, len(binds))
		i := 0
		for k := range binds {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		env := imp.env
		for _, k := range keys {
			bind := imp.Binds[k]
			v := bind.RuntimeValue(env)
			showValue(out, k, v, bind.Type, stringer)
		}
		fmt.Fprintln(out)
	}
	showTypes(out, imp.Name, imp.Path, imp.Types, stringer)
}

func showTypes(out io.Writer, name string, path string, types map[string]xr.Type, stringer func(xr.Type) string) {
	if len(types) > 0 {
		base.ShowPackageHeader(out, name, path, "types")

		keys := make([]string, len(types))
		i := 0
		for k := range types {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			t := types[k]
			if t != nil {
				showType(out, k, t, stringer)
			}
		}
		fmt.Fprintln(out)
	}
}

const spaces15 = "               "

func showValue(out io.Writer, name string, v r.Value, t xr.Type, stringer func(xr.Type) string) {
	n := len(name) & 15
	str := stringer(t)
	if v == base.Nil || v == base.None {
		fmt.Fprintf(out, "%s%s = nil\t// %s\n", name, spaces15[n:], str)
	} else {
		fmt.Fprintf(out, "%s%s = %v\t// %s\n", name, spaces15[n:], v, str)
	}
}

func showType(out io.Writer, name string, t xr.Type, stringer func(xr.Type) string) {
	n := len(name) & 15
	fmt.Fprintf(out, "%s%s = %v\t// %v\n", name, spaces15[n:], stringer(t), t.Kind())
}
