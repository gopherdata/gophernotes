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
		return t.universe.maketype(gtype.Elem(), rtype.Elem())
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
	return v.MakeType(
		types.NewPointer(elem.GoType()),
		reflect.PtrTo(elem.ReflectType()))
}

func (v *Universe) SliceOf(elem Type) Type {
	return v.MakeType(
		types.NewSlice(elem.GoType()),
		reflect.SliceOf(elem.ReflectType()))
}
