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
 * named.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/token"
	r "reflect"
	"sort"
	"unsafe"

	"github.com/cosmos72/gomacro/go/etoken"

	"github.com/cosmos72/gomacro/go/types"
)

// NamedOf returns a new named type for the given type name and package.
// Initially, the underlying type may be set to interface{} - use SetUnderlying to change it.
// These two steps are separate to allow creating self-referencing types,
// as for example type List struct { Elem int; Rest *List }
func (v *Universe) NamedOf(name, pkgpath string) Type {
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return v.namedOf(name, pkgpath)
}

func (v *Universe) namedOf(name, pkgpath string) Type {
	return v.reflectNamedOf(name, pkgpath, v.TypeOfForward.ReflectType())
}

// alternate version of namedOf(), to be used when reflect.Type is known
func (v *Universe) reflectNamedOf(name, pkgpath string, rtype r.Type) Type {
	underlying := v.BasicTypes[rtype.Kind()]
	if underlying == nil {
		underlying = v.TypeOfForward
	}
	pkg := v.loadPackage(pkgpath)
	typename := types.NewTypeName(token.NoPos, (*types.Package)(pkg), name, nil)
	return v.maketype3(
		// kind is reflect.Invalid;
		// underlying.GoType() will often be inaccurate and equal to interface{};
		// rtype will often be inaccurate and equal to TypeOfForward.
		// All these issues will be fixed by Type.SetUnderlying()
		r.Invalid,
		// if etoken.GENERICS_V2_CTI, v.BasicTypes[kind] is a named type
		// wrapping the actual basic type
		types.NewNamed(typename, underlying.GoType().Underlying(), nil),
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
		if t.kind != r.Invalid || gtype.Underlying() != v.TypeOfForward.GoType() || t.rtype != v.TypeOfForward.ReflectType() {
			// redefined type. try really hard to support it.
			v.InvalidateCache()
			// xerrorf(t, "SetUnderlying invoked multiple times on named type %v", t)
		}
		xunderlying := unwrap(underlying)
		gunderlying := xunderlying.gtype.Underlying() // in case underlying is named
		t.kind = gtypeToKind(xunderlying, gunderlying)
		gtype.SetUnderlying(gunderlying)
		// debugf("SetUnderlying: updated <%v> reflect Type from <%v> to <%v>", gtype, t.rtype, underlying.ReflectType())
		t.rtype = underlying.ReflectType()
		if t.kind == r.Interface {
			// propagate methodvalues from underlying interface to named type
			t.methodvalues = xunderlying.methodvalues
			t.methodcache = nil
			t.fieldcache = nil
		} else if etoken.GENERICS_V2_CTI {
			v.addTypeMethodsCTI(t)
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
	if kind == r.Ptr || kind == r.Interface {
		xerrorf(t, "AddMethod: cannot add methods to named %s type: <%v>", kind, t)
	}
	if signature.Kind() != r.Func {
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
	index := gtype.ReplaceMethod(gfun)
	n2 := gtype.NumMethods()

	nilv := r.Value{}
	for len(t.methodvalues) < n2 {
		t.methodvalues = append(t.methodvalues, nilv)
	}
	// store in t.methodvalues[index] a nil function with the correct reflect.Type:
	// needed by Type.GetMethod(int) to retrieve the method's reflect.Type
	//
	// fixes gophernotes issue 174
	t.methodvalues[index] = r.Zero(signature.ReflectType())
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
func (t *xtype) GetMethods() *[]r.Value {
	if !etoken.GENERICS_V2_CTI && !t.Named() {
		xerrorf(t, "GetMethods on unnamed type %v", t)
	}
	resizemethodvalues(t)
	return &t.methodvalues
}
