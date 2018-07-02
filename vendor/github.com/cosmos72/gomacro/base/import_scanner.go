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
	"strings"
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
			// do not export proxies for empty interfaces:
			// using reflect.Value.Convert() at runtime is enough
			if u.NumMethods() != 0 && (!requireAllMethodsExported || allMethodsExported(u)) {
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

// return the string after last '/' in path
func FileName(path string) string {
	return path[1+strings.LastIndexByte(path, '/'):]
}

// return the string up to (and including) last '/' in path
func DirName(path string) string {
	return path[0 : 1+strings.LastIndexByte(path, '/')]
}

// remove last byte from string
func RemoveLastByte(s string) string {
	if n := len(s); n != 0 {
		s = s[:n-1]
	}
	return s
}

// we need to collect only the imports that actually appear in package's interfaces methods
// because Go rejects programs with unused imports.
//
// To avoid naming conflicts when importing two different packages
// that end with the same name, as for example image/draw and golang.org/x/image/draw,
// we rename conflicting packages and return a map[path]renamed
func (o *Output) CollectPackageImportsWithRename(pkg *types.Package, requireAllInterfaceMethodsExported bool) map[string]string {
	ie := importExtractor{
		// we always need to import the package itself
		imports: map[string]bool{pkg.Path(): true},
		o:       o,
	}
	ie.visitPackage(pkg, requireAllInterfaceMethodsExported)

	// for deterministic renaming, use a sorted []string instead of a map[string]bool
	paths := getKeys(ie.imports)
	sort.Strings(paths)

	nametopath := renamePackages(paths)
	pathtoname := transposeKeyValue(nametopath)

	// do NOT rename the package we are scanning!
	path := pkg.Path()
	name := sanitizeIdentifier(FileName(path))
	if name2 := pathtoname[path]; name2 != name {
		// some *other* path may be associated to name.
		// in case, swap the names of the two packages
		if path2, ok := nametopath[name]; ok {
			pathtoname[path2] = name2
		}
		pathtoname[path] = name
	}
	return pathtoname
}

// given a slice []path, return a map[name]path where all paths
// that end in the same name have been assigned unique names
func renamePackages(in []string) map[string]string {
	out := make(map[string]string)
	for _, path := range in {
		name := renamePackage(path, out)
		out[name] = path
	}
	return out
}

// given a path and a map[name]path, extract the path last name.
// Change it (if needed) to a value that is NOT in map and return it.
func renamePackage(path string, out map[string]string) string {
	name := sanitizeIdentifier(FileName(path))
	if _, exists := out[name]; !exists {
		return name
	}
	n := len(name)
	for n != 0 && isDigit(name[n-1]) {
		n--
	}
	name = name[:n]
	for i := uint64(0); i < ^uint64(0); i++ {
		namei := fmt.Sprintf("%s%d", name, i)
		if _, exists := out[namei]; !exists {
			return namei
		}
	}
	Errorf("failed to find a non-conflicting rename for package %q", path)
	return "???"
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

// given a map k -> v, return an *unsorted* slice of its keys
func getKeys(in map[string]bool) []string {
	keys := make([]string, len(in))
	i := 0
	for key := range in {
		keys[i] = key
		i++
	}
	return keys
}

// given a map k -> v, return a map v -> k
func transposeKeyValue(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[v] = k
	}
	return out
}
