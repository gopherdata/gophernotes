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
 * import_scanner.go
 *
 *  Created on Mar 06, 2017
 *      Author Massimiliano Ghilardi
 */

package base

import (
	"fmt"
	"go/types"
	r "reflect"
	"sort"
)

type TypeVisitor func(name string, t types.Type) bool

// implemented by *types.Pointer, *types.Array, *types.Slice, *types.Chan
type typeWithElem interface {
	Elem() types.Type
}

var depth int = 0

func (o *Output) trace(msg ...interface{}) {
	const dots = ". . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . "
	const n = len(dots)
	i := 2 * depth
	for i > n {
		fmt.Fprint(o.Stdout, dots)
		i -= n
	}
	// i <= n
	fmt.Fprint(o.Stdout, dots[0:i])
	fmt.Fprintln(o.Stdout, msg...)
}

func trace(o *Output, caller string, name string, x interface{}) *Output {
	o.trace(caller, "(", name, x)
	depth++
	return o
}

func un(o *Output) {
	depth--
	o.trace(")")
}

func (o *Output) traverseType(name string, in types.Type, visitor TypeVisitor) {
	for {
		// defer un(trace(o, "traverseType", name, r.TypeOf(in)))

		if !visitor(name, in) {
			return
		}
		switch t := in.(type) {
		case *types.Basic:
			break
		case *types.Named:
			u := t.Underlying()
			if in != u {
				name = t.Obj().Name()
				in = u
				continue
			}
		case *types.Signature:
			if recv := t.Recv(); recv != nil {
				u := recv.Type()
				// the receiver is often the interface containing this signature...
				// avoid infinite recursion!
				if in != u {
					if _, ok := u.(*types.Interface); !ok {
						o.traverseType(recv.Name(), u, visitor)
					}
				}
			}
			tuples := []*types.Tuple{t.Params(), t.Results()}
			for _, tuple := range tuples {
				n := tuple.Len()
				for i := 0; i < n; i++ {
					v := tuple.At(i)
					o.traverseType(v.Name(), v.Type(), visitor)
				}
			}
		case *types.Interface:
			n := t.NumMethods()
			for i := 0; i < n; i++ {
				method := t.Method(i)
				o.traverseType(method.Name(), method.Type(), visitor)
			}
		case *types.Struct:
			n := t.NumFields()
			for i := 0; i < n; i++ {
				field := t.Field(i)
				o.traverseType(field.Name(), field.Type(), visitor)
			}
		case *types.Map:
			o.traverseType("", t.Key(), visitor)
			name = ""
			in = t.Elem()
			continue
		case typeWithElem: // *types.Pointer, *types.Array, *types.Slice, *types.Chan
			name = ""
			in = t.Elem()
			continue
		default:
			o.Warnf("traverseType: unimplemented %#v <%v>", t, r.TypeOf(t))
		}
		break
	}
}

type importExtractor struct {
	imports map[string]bool
	seen    map[types.Type]bool
	o       *Output
}

func (ie *importExtractor) visitPackage(pkg *types.Package, requireAllInterfaceMethodsExported bool) {
	scope := pkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		t := extractInterface(obj, requireAllInterfaceMethodsExported)
		if t != nil {
			ie.o.traverseType("", t, ie.visitType)
		}
	}
}

func (ie *importExtractor) visitType(name string, t types.Type) bool {
	if ie.seen[t] {
		return false
	}
	switch t := t.(type) {
	case *types.Named:
		if obj := t.Obj(); obj != nil {
			if pkg := obj.Pkg(); pkg != nil {
				ie.imports[pkg.Path()] = true
			}
		}
		// no need to visit the definition of a named type
		return false
	default:
		return true
	}
}

func extractInterface(obj types.Object, requireAllMethodsExported bool) *types.Interface {
	if obj == nil || !obj.Exported() {
		return nil
	}
	switch obj.(type) {
	case *types.TypeName:
		u := obj.Type().Underlying()
		if u, ok := u.(*types.Interface); ok {
			if !requireAllMethodsExported || allMethodsExported(u) {
				return u
			}
		}
	}
	return nil
}

func allMethodsExported(intf *types.Interface) bool {
	n := intf.NumMethods()
	for i := 0; i < n; i++ {
		if !intf.Method(i).Exported() {
			return false
		}
	}
	return true
}

// we need to collect only the imports that actually appear in package's interfaces methods
// because Go rejects programs with unused imports
func (o *Output) CollectPackageImports(pkg *types.Package, requireAllInterfaceMethodsExported bool) []string {
	ie := importExtractor{
		// we always need to import the package itself
		imports: map[string]bool{pkg.Path(): true},
		o:       o,
	}
	ie.visitPackage(pkg, requireAllInterfaceMethodsExported)

	strings := make([]string, len(ie.imports))
	i := 0
	for str := range ie.imports {
		strings[i] = str
		i++
	}
	sort.Strings(strings)
	return strings
}
