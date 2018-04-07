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

func (imp Import) String() string {
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
	out := c.Stdout
	stringer := typestringer(c.Path)
	if binds := c.Binds; len(binds) > 0 {
		showPackageHeader(out, c.Name, c.Path, "binds")

		keys := make([]string, len(binds))
		i := 0
		for k := range binds {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			bind := binds[k]
			if bind == nil {
				continue
			}
			if bind.Const() {
				showValue(out, k, bind.ConstValue(), bind.Type, stringer)
				continue
			}
			expr := c.Symbol(bind.AsSymbol(0))
			showValue(out, k, ir.RunExpr1(expr), expr.Type, stringer)
		}
		fmt.Fprintln(out)
	}
	showTypes(out, c.Name, c.Path, c.Types, stringer)
}

func (ir *Interp) ShowImportedPackage(name string) {
	var imp Import
	var ok bool
	if bind := ir.Comp.Binds[name]; bind != nil && bind.Const() && bind.Type != nil && bind.Type.ReflectType() == rtypeOfImport {
		imp, ok = bind.Value.(Import)
	}
	if !ok {
		ir.Comp.Warnf("not an imported package: %q", name)
		return
	}
	c := ir.Comp
	out := c.Stdout
	stringer := typestringer(imp.Path)
	if binds := imp.Binds; len(binds) > 0 {
		showPackageHeader(out, imp.Name, imp.Path, "binds")

		keys := make([]string, len(binds))
		i := 0
		for k := range binds {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			showValue(out, k, binds[k], imp.BindTypes[k], stringer)
		}
		fmt.Fprintln(out)
	}
	showTypes(out, imp.Name, imp.Path, imp.Types, stringer)
}

func showTypes(out io.Writer, name string, path string, types map[string]xr.Type, stringer func(xr.Type) string) {
	if len(types) > 0 {
		showPackageHeader(out, name, path, "types")

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

func showPackageHeader(out io.Writer, name string, path string, kind string) {
	if name == path {
		fmt.Fprintf(out, "// ----- %s %s -----\n", name, kind)
	} else if name == base.FileName(path) {
		fmt.Fprintf(out, "// ----- %q %s -----\n", path, kind)
	} else {
		fmt.Fprintf(out, "// ----- %s %q %s -----\n", name, path, kind)
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
