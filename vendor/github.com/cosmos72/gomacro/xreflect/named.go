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
 * named.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/token"
	"go/types"
	"reflect"
	"sort"
	"unsafe"
)

// NamedOf returns a new named type for the given type name and package.
// Initially, the underlying type may be set to interface{} - use SetUnderlying to change it.
// These two steps are separate to allow creating self-referencing types,
// as for example type List struct { Elem int; Rest *List }
func (v *Universe) NamedOf(name, pkgpath string, kind reflect.Kind) Type {
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return v.namedOf(name, pkgpath, kind)
}

func (v *Universe) namedOf(name, pkgpath string, kind reflect.Kind) Type {
	underlying := v.BasicTypes[kind]
	if underlying == nil {
		underlying = v.TypeOfForward
	}
	return v.reflectNamedOf(name, pkgpath, kind, underlying.ReflectType())
}

// alternate version of namedOf(), to be used when reflect.Type is known
func (v *Universe) reflectNamedOf(name, pkgpath string, kind reflect.Kind, rtype reflect.Type) Type {
	underlying := v.BasicTypes[kind]
	if underlying == nil {
		underlying = v.TypeOfInterface
	}
	pkg := v.loadPackage(pkgpath)
	typename := types.NewTypeName(token.NoPos, (*types.Package)(pkg), name, nil)
	return v.maketype3(
		// kind may be inaccurate or reflect.Invalid;
		// underlying.GoType() will often be inaccurate and equal to interface{};
		// rtype will often be inaccurate and equal to Incomplete.
		// All these issues will be fixed by Type.SetUnderlying()
		kind,
		types.NewNamed(typename, underlying.GoType(), nil),
		rtype,
	)
}

// SetUnderlying sets the underlying type of a named type and marks t as complete.
// It panics if the type is unnamed, or if the underlying type is named,
// or if SetUnderlying() was already invoked on the named type.
func (t *xtype) SetUnderlying(underlying Type) {
	switch gtype := t.gtype.(type) {
	case *types.Named:
		v := t.universe
		if t.kind != reflect.Invalid || gtype.Underlying() != v.TypeOfInterface.GoType() || t.rtype != v.TypeOfInterface.ReflectType() {
			// redefined type. try really hard to support it.
			v.InvalidateCache()
			// xerrorf(t, "SetUnderlying invoked multiple times on named type %v", t)
		}
		tunderlying := unwrap(underlying)
		gunderlying := tunderlying.gtype.Underlying() // in case underlying is named
		t.kind = gtypeToKind(t, gunderlying)
		gtype.SetUnderlying(gunderlying)
		// debugf("SetUnderlying: updated <%v> reflect Type from <%v> to <%v>", gtype, t.rtype, underlying.ReflectType())
		t.rtype = underlying.ReflectType()
		if t.kind == reflect.Interface {
			// propagate methodvalues from underlying interface to named type
			t.methodvalues = tunderlying.methodvalues
			t.methodcache = nil
			t.fieldcache = nil
		}
	default:
		xerrorf(t, "SetUnderlying of unnamed type %v", t)
	}
}

// AddMethod adds method 'name' to type.
// It panics if the type is unnamed, or if the signature is not a function type,
// Returns the method index, or < 0 in case of errors
func (t *xtype) AddMethod(name string, signature Type) int {
	gtype, ok := t.gtype.(*types.Named)
	if !ok {
		xerrorf(t, "AddMethod on unnamed type %v", t)
	}
	kind := gtypeToKind(t, gtype.Underlying())
	if kind == reflect.Ptr || kind == reflect.Interface {
		xerrorf(t, "AddMethod: cannot add methods to named %s type: <%v>", kind, t)
	}
	if signature.Kind() != reflect.Func {
		xerrorf(t, "AddMethod on <%v> of non-function: %v", t, signature)
	}
	gsig := signature.gunderlying().(*types.Signature)
	// accept both signatures "non-nil receiver" and "nil receiver, use the first parameter as receiver"
	grecv := gsig.Recv()
	if grecv == nil && gsig.Params().Len() != 0 {
		grecv = gsig.Params().At(0)
	}
	if grecv == nil {
		xerrorf(t, "AddMethod on <%v> of function with no receiver and no parameters: %v", t, gsig)
	}
	if !types.IdenticalIgnoreTags(grecv.Type(), gtype) &&
		// !types.IdenticalIgnoreTags(grecv.Type(), gtype.Underlying()) &&
		!types.IdenticalIgnoreTags(grecv.Type(), types.NewPointer(gtype)) {

		label := "receiver"
		if gsig.Recv() == nil {
			label = "first parameter"
		}
		xerrorf(t, "AddMethod on <%v> of function <%v> with mismatched %s type: %v", t, gsig, label, grecv.Type())
	}

	gpkg := gtype.Obj().Pkg()
	gfun := types.NewFunc(token.NoPos, gpkg, name, gsig)

	n1 := gtype.NumMethods()
	index := unsafeAddMethod(gtype, gfun)
	n2 := gtype.NumMethods()

	// update the caches... be careful if the method was just redefined
	nilv := reflect.Value{}
	for len(t.methodvalues) < n2 {
		t.methodvalues = append(t.methodvalues, nilv)
	}
	t.methodvalues[index] = nilv
	if n1 == n2 {
		// an existing method was overwritten.
		// it may be cached in some other type's method cache.
		t.universe.InvalidateMethodCache()
	}
	return index
}

// RemoveMethods removes given methods from type.
// It panics if the type is unnamed, or if the signature is not a function type,
func (t *xtype) RemoveMethods(names []string, pkgpath string) {
	gtype, ok := t.gtype.(*types.Named)
	if !ok {
		xerrorf(t, "RemoveMethods on unnamed type %v", t)
	}
	if len(names) == 0 {
		return
	}
	n1 := gtype.NumMethods()
	unsafeRemoveMethods(gtype, names, pkgpath)
	n2 := gtype.NumMethods()
	if n1 != n2 {
		// some existing methods were removed.
		// they may be cached in some other type's method cache.
		t.universe.InvalidateMethodCache()
	}
}

// internal representation of go/types.Named
type unsafeNamed struct {
	obj        *types.TypeName
	underlying types.Type
	methods    []*types.Func
}

// patched version of go/types.Named.AddMethod() that *overwrites* matching methods
// (the original does not)
func unsafeAddMethod(gtype *types.Named, gfun *types.Func) int {
	if gfun.Name() == "_" {
		return -1
	}
	gt := (*unsafeNamed)(unsafe.Pointer(gtype))
	qname := QNameGo(gfun)
	for i, m := range gt.methods {
		if qname == QNameGo(m) {
			gt.methods[i] = gfun
			return i
		}
	}
	gt.methods = append(gt.methods, gfun)
	return len(gt.methods) - 1
}

func unsafeRemoveMethods(gtype *types.Named, names []string, pkgpath string) {
	names = append([]string{}, names...) // make a copy
	sort.Strings(names)                  // and sort it

	gt := (*unsafeNamed)(unsafe.Pointer(gtype))

	n1 := len(gt.methods)
	n2 := n1
	for i, j := 0, 0; i < n1; i++ {
		m := gt.methods[i]
		name := m.Name()
		pos := sort.SearchStrings(names, name)
		if pos < len(names) && names[pos] == name && (m.Exported() || m.Pkg().Path() == pkgpath) {
			// delete this method
			n2--
			continue
		}
		if i != j {
			gt.methods[j] = gt.methods[i]
		}
		j++
	}
	if n1 != n2 {
		gt.methods = gt.methods[:n2]
	}
}

// GetMethods returns the pointer to the method values.
// It panics if the type is unnamed
func (t *xtype) GetMethods() *[]reflect.Value {
	if !t.Named() {
		xerrorf(t, "GetMethods on unnamed type %v", t)
	}
	resizemethodvalues(t)
	return &t.methodvalues
}
