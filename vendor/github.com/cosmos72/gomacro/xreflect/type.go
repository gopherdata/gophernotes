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
 * type.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	r "reflect"

	"github.com/cosmos72/gomacro/go/types"

	"github.com/cosmos72/gomacro/go/etoken"
	"github.com/cosmos72/gomacro/go/typeutil"
)

func identicalType(t, u Type) bool {
	xt := unwrap(t)
	yt := unwrap(u)
	xnil := xt == nil
	ynil := yt == nil
	if xnil || ynil {
		return xnil == ynil
	}
	return xt == yt || xt.identicalTo(yt)
}

func debugOnMismatchCache(m *typeutil.Map, gtype types.Type, rtype r.Type, cached Type) {
	debugf("overwriting mismatched reflect.Type found in cache for type %v (hash 0x%x):\n\tnew reflect.Type: %v\n\told reflect.Type: %v",
		typeutil.String(gtype), m.Hasher().Hash(gtype), rtype, cached.ReflectType()) //, debug.Stack())
}

func (t *xtype) warnOnSuspiciousCache() {
	// reflect cannot create new interface types or new named types: accept whatever we have.
	// also, it cannot create unnamed structs containing unexported fields. again, accept whatever we have.
	// instead complain on mismatch for non-interface, non-named types
	rt := t.rtype
	if !t.Named() && len(rt.Name()) != 0 && rt.Kind() != r.Interface && rt.Kind() != r.Struct {
		xerrorf(t, "caching suspicious type %v => %v", t.gtype, rt)
	}
}

func (m *Types) clear() {
	*m = Types{}
}

func (m *Types) add(t Type) {
	xt := unwrap(t)

	if xt.rtype == rTypeOfForward {
		if m.gmap.At(xt.gtype) != nil {
			// debugf("not adding again type to cache: %v <%v> reflect type: <%v>\n%s", xt.kind, xt.gtype, xt.rtype)
			return
		}
	} else {
		xt.warnOnSuspiciousCache()
	}
	switch xt.kind {
	case r.Func:
		// even function types can be named => they need SetUnderlying() before being complete
		if !xt.needSetUnderlying() {
			xt.NumIn() // check consistency
		}
	case r.Interface:
		rtype := t.ReflectType()
		rkind := rtype.Kind()
		if rkind != r.Interface && (rkind != r.Ptr || rtype.Elem().Kind() != r.Struct) {
			errorf(t, "bug! inconsistent type <%v>: has kind = %s but its Type.Reflect() is %s\n\tinstead of interface or pointer-to-struct: <%v>",
				t, t.Kind(), rtype.Kind(), t.ReflectType())
		}
	}
	m.gmap.Set(xt.gtype, t)
	// debugf("added type to cache: %v <%v> reflect type: <%v>", xt.kind, xt.gtype, xt.rtype)
}

// all unexported methods assume lock is already held
func (v *Universe) maketype3(kind r.Kind, gtype types.Type, rtype r.Type) Type {
	if gtype == nil {
		errorf(nil, "MakeType of nil types.Type")
	} else if rtype == nil {
		errorf(nil, "MakeType of nil reflect.Type")
	}
	ret := v.Types.gmap.At(gtype)
	if ret != nil {
		t := ret.(Type)
		switch t.ReflectType() {
		case rtype:
			return t
		case rTypeOfForward:
			// update t, do not create a new Type
			t.UnsafeForceReflectType(rtype)
			return t
		}
		if v.debug() {
			debugOnMismatchCache(&v.Types.gmap, gtype, rtype, t)
		}
	}
	xt := &xtype{kind: kind, gtype: gtype, rtype: rtype, universe: v}
	t := wrap(xt)
	v.add(t)
	v.addTypeMethodsCTI(xt)
	return t
}

func (v *Universe) maketype(gtype types.Type, rtype r.Type) Type {
	return v.maketype3(gtypeToKind(nil, gtype), gtype, rtype)
}

func (v *Universe) MakeType(gtype types.Type, rtype r.Type) Type {
	kind := gtypeToKind(nil, gtype)
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return v.maketype3(kind, gtype, rtype)
}

// GoType returns the go/types.Type corresponding to the type.
func (t *xtype) GoType() types.Type {
	return t.gtype
}

// ReflectType returns a best-effort reflect.Type that approximates the type.
// It may be inexact for the following reasons:
// 1) missing reflect.NamedOf(): no way to programmatically create named types, or to access the underlying type of a named type
// 2) missing reflect.InterfaceOf(): interface types created at runtime will be approximated by structs
// 3) missing reflect.MethodOf(): method types created at runtime will be approximated by functions
//    whose first parameter is the receiver
// 4) reflect.StructOf() does not support embedded or unexported fields
// 5) go/reflect lacks the ability to create self-referencing types:
//    references to the type itself will be replaced by interface{}.
//
// Examples:
//    after invoking at runtime type2.NewStruct() and type2.NewNamed()
//    to create the following type:
//        type List struct { Elem int; Rest *List }
//    ReflectType will return a reflect.Type equivalent to:
//        struct { Elem int; Rest interface{} }
//    i.e. the type name will be missing due to limitation 1 above,
//    and the field 'Rest' will have type interface{} instead of *List due to limitation 5.
func (t *xtype) ReflectType() r.Type {
	return t.rtype
}

func (t *xtype) UnsafeForceReflectType(rtype r.Type) {
	t.rtype = rtype
}

func (t *xtype) Universe() *Universe {
	return t.universe
}

// Named returns whether the type is named.
// It returns false for unnamed types.
func (t *xtype) Named() bool {
	if t != nil {
		switch t.gtype.(type) {
		case *types.Basic, *types.Named:
			return true
		}
	}
	return false
}

// Name returns the type's name within its package.
// It returns an empty string for unnamed types.
func (t *xtype) Name() string {
	if t != nil {
		switch gtype := t.gtype.(type) {
		case *types.Basic:
			return gtype.Name()
		case *types.Named:
			return gtype.Obj().Name()
		}
	}
	return ""
}

// Pkg returns a named type's package, that is, the package where it was defined.
// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
// Pkg will return nil.
func (t *xtype) Pkg() *Package {
	switch gtype := t.gtype.(type) {
	case *types.Named:
		return (*Package)(gtype.Obj().Pkg())
	default:
		return nil
	}
}

// PkgName returns a named type's package name, that is,
// the default name that the package provides when imported.
// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
// the package name will be the empty string.
func (t *xtype) PkgName() string {
	if gtype, ok := t.gtype.(*types.Named); ok {
		pkg := gtype.Obj().Pkg()
		// pkg may be nil for builtin named types, as for example 'error'
		if pkg != nil {
			return pkg.Name()
		}
	}
	return ""
}

// PkgPath returns a named type's package path, that is, the import path
// that uniquely identifies the package, such as "encoding/base64".
// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
// the package path will be the empty string.
func (t *xtype) PkgPath() string {
	if gtype, ok := t.gtype.(*types.Named); ok {
		pkg := gtype.Obj().Pkg()
		// pkg may be nil for builtin named types, as for example 'error'
		if pkg != nil {
			return pkg.Path()
		}
	}
	return ""
}

// Size returns the number of bytes needed to store
// a value of the given type; it is analogous to unsafe.Sizeof.
func (t *xtype) Size() uintptr {
	return t.rtype.Size()
}

// String returns a string representation of a type.
func (t *xtype) String() string {
	if t == nil {
		return "<nil>"
	}
	return typeutil.String(t.gtype)
}

func (t *xtype) string2(name string) string {
	if t == nil {
		return "<nil>"
	}
	return typeutil.String2(name, t.gtype)
}

/*
// Underlying returns the underlying type of a type.
func (t *xtype) Underlying() Type {
	return Type{t.underlying}
}
*/

func (t *xtype) gunderlying() types.Type {
	return t.gtype.Underlying()
}

// best-effort implementation of missing reflect.Type.Underlying()
func (t *xtype) runderlying() r.Type {
	return ReflectUnderlying(t.rtype)
}

// Kind returns the specific kind of the type.
func (t *xtype) Kind() r.Kind {
	if t == nil {
		return r.Invalid
	}
	return t.kind
}

// Implements reports whether the type implements the interface type u.
// It panics if u's Kind is not Interface
func (t *xtype) Implements(u Type) bool {
	if u.Kind() != r.Interface {
		xerrorf(t, "Type.Implements of non-interface type: %v", u)
	}
	xu := unwrap(u)
	return t.gtype == xu.gtype ||
		(types.Implements(t.gtype, xu.gtype.Underlying().(*types.Interface)) &&
			matchReceiverType(t, xu))
}

func matchReceiverType(t, u *xtype) bool {
	t = getConstrainedReceiverType(t, true)
	u = getConstrainedReceiverType(u, false)
	return t == nil || u == nil || t.identicalTo(u)
}

func getConstrainedReceiverType(t *xtype, lhs bool) *xtype {
	if etoken.GENERICS_V2_CTI {
		if lhs && t.kind != r.Interface {
			return t
		}
		if recv, ok := t.GetUserData(ConstrainedInterfaceReceiverType); ok {
			if trecv, ok := recv.(Type); ok {
				return unwrap(trecv)
			}
		}
	}
	return nil
}

// IdenticalTo reports whether the type is identical to type u.
func (t *xtype) IdenticalTo(u Type) bool {
	xu := unwrap(u)
	return t == xu || t.identicalTo(xu)
}

func (t *xtype) identicalTo(u *xtype) bool {
	return typeutil.Identical(t.GoType(), u.GoType())
}

// AssignableTo reports whether a value of the type is assignable to type u.
func (t *xtype) AssignableTo(u Type) bool {
	// debugf("AssignableTo: <%v> <%v>", t, u)
	return t.gtype == u.GoType() ||
		(types.AssignableTo(t.gtype, u.GoType()) &&
			matchReceiverType(t, unwrap(u)))
}

// ConvertibleTo reports whether a value of the type is convertible to type u.
func (t *xtype) ConvertibleTo(u Type) bool {
	return t.gtype == u.GoType() ||
		(types.ConvertibleTo(t.gtype, u.GoType()) &&
			matchReceiverType(t, unwrap(u)))
}

// Comparable reports whether values of this type are comparable.
func (t *xtype) Comparable() bool {
	return types.Comparable(t.gtype)
}

// GetUserData returns the user-supplied data of the type.
func (t *xtype) GetUserData(key interface{}) (interface{}, bool) {
	if t == nil {
		return nil, false
	}
	val, ok := t.userdata[key]
	return val, ok
}

// SetUserData sets the user-supplied data of the type.
func (t *xtype) SetUserData(key, value interface{}) {
	if t.userdata == nil {
		t.userdata = map[interface{}]interface{}{
			key: value,
		}
	} else {
		t.userdata[key] = value
	}
}

// Zero returns a Value representing the zero value for the specified type.
func Zero(t Type) r.Value {
	return r.Zero(t.ReflectType())
}
