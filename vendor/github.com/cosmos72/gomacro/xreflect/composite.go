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
 * composite.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/types"
	"reflect"
)

// ChanDir returns a channel type's direction.
// It panics if the type's Kind is not Chan.
func (t *xtype) ChanDir() reflect.ChanDir {
	if t.Kind() != reflect.Chan {
		xerrorf(t, "ChanDir of non-chan type %v", t)
	}
	return t.rtype.ChanDir()
}

// Elem returns a type's element type.
// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
func (t *xtype) Elem() Type {
	v := t.universe
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return t.elem()
}

func (t *xtype) elem() Type {
	gtype := t.gunderlying()
	rtype := t.rtype
	switch gtype := gtype.(type) {
	case *types.Array:
		return t.universe.maketype(gtype.Elem(), rtype.Elem())
	case *types.Chan:
		return t.universe.maketype(gtype.Elem(), rtype.Elem())
	case *types.Map:
		return t.universe.maketype(gtype.Elem(), rtype.Elem())
	case *types.Pointer:
		// if reflect type is xreflect.Forward due to contagion,
		// we do not know the element type -> return xreflect.Forward
		if rtype != rTypeOfForward {
			rtype = rtype.Elem()
		}
		return t.universe.maketype(gtype.Elem(), rtype)
	case *types.Slice:
		return t.universe.maketype(gtype.Elem(), rtype.Elem())
	default:
		xerrorf(t, "Elem of invalid type %v", t)
		return nil
	}
}

// Key returns a map type's key type.
// It panics if the type's Kind is not Map.
func (t *xtype) Key() Type {
	if t.Kind() != reflect.Map {
		xerrorf(t, "Key of non-map type %v", t)
	}
	gtype := t.gunderlying().(*types.Map)
	return t.universe.MakeType(gtype.Key(), t.rtype.Key())
}

// Len returns an array type's length.
// It panics if the type's Kind is not Array.
func (t *xtype) Len() int {
	if t.Kind() != reflect.Array {
		xerrorf(t, "Len of non-array type %v", t)
	}
	return t.rtype.Len()
}

func (v *Universe) ArrayOf(count int, elem Type) Type {
	return v.MakeType(
		types.NewArray(elem.GoType(), int64(count)),
		reflect.ArrayOf(count, elem.ReflectType()))
}

func (v *Universe) ChanOf(dir reflect.ChanDir, elem Type) Type {
	gdir := dirToGdir(dir)
	return v.MakeType(
		types.NewChan(gdir, elem.GoType()),
		reflect.ChanOf(dir, elem.ReflectType()))
}

func (v *Universe) MapOf(key, elem Type) Type {
	return v.MakeType(
		types.NewMap(key.GoType(), elem.GoType()),
		reflect.MapOf(key.ReflectType(), elem.ReflectType()))
}

func (v *Universe) PtrTo(elem Type) Type {
	rtyp := elem.ReflectType()

	// do not create the reflect type *xreflect.Forward
	// because it hurts the implementation of recursive types.
	// Instead, consider xreflect.Forward as slightly contagious.
	if rtyp != rTypeOfForward {
		rtyp = reflect.PtrTo(rtyp)
	}

	return v.MakeType(
		types.NewPointer(elem.GoType()),
		rtyp)
}

func (v *Universe) SliceOf(elem Type) Type {
	return v.MakeType(
		types.NewSlice(elem.GoType()),
		reflect.SliceOf(elem.ReflectType()))
}
