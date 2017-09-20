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
 * init.go
 *
 *  Created on May 19, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"reflect"
	"unsafe"

	"go/types"
)

var rbasictypes = []reflect.Type{
	reflect.Bool:          reflect.TypeOf(bool(false)),
	reflect.Int:           reflect.TypeOf(int(0)),
	reflect.Int8:          reflect.TypeOf(int8(0)),
	reflect.Int16:         reflect.TypeOf(int16(0)),
	reflect.Int32:         reflect.TypeOf(int32(0)),
	reflect.Int64:         reflect.TypeOf(int64(0)),
	reflect.Uint:          reflect.TypeOf(uint(0)),
	reflect.Uint8:         reflect.TypeOf(uint8(0)),
	reflect.Uint16:        reflect.TypeOf(uint16(0)),
	reflect.Uint32:        reflect.TypeOf(uint32(0)),
	reflect.Uint64:        reflect.TypeOf(uint64(0)),
	reflect.Uintptr:       reflect.TypeOf(uintptr(0)),
	reflect.Float32:       reflect.TypeOf(float32(0)),
	reflect.Float64:       reflect.TypeOf(float64(0)),
	reflect.Complex64:     reflect.TypeOf(complex64(0)),
	reflect.Complex128:    reflect.TypeOf(complex128(0)),
	reflect.String:        reflect.TypeOf(string("")),
	reflect.UnsafePointer: reflect.TypeOf(unsafe.Pointer(nil)),
}

func (v *Universe) makebasictypes() []Type {
	m := make([]Type, len(rbasictypes))
	for gkind := types.Bool; gkind <= types.UnsafePointer; gkind++ {
		kind := ToReflectKind(gkind)
		gtype := types.Typ[gkind]
		rtype := rbasictypes[kind]
		if gtype == nil || rtype == nil {
			continue
		}
		t := wrap(&xtype{kind: kind, gtype: gtype, rtype: rtype, universe: v})
		v.add(t)
		m[kind] = t
	}
	return m
}

func (v *Universe) makeerror() Type {
	t := wrap(&xtype{
		kind:     reflect.Interface,
		gtype:    types.Universe.Lookup("error").Type(),
		rtype:    reflect.TypeOf((*error)(nil)).Elem(),
		universe: v,
	})
	v.add(t)
	return t
}

func (v *Universe) makeinterface() Type {
	t := wrap(&xtype{
		kind:     reflect.Interface,
		gtype:    types.NewInterface(nil, nil).Complete(),
		rtype:    reflect.TypeOf((*interface{})(nil)).Elem(),
		universe: v,
	})
	v.add(t)
	return t
}

func (v *Universe) Init() *Universe {
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return v.init()
}

func (v *Universe) init() *Universe {
	v.BasicTypes = v.makebasictypes()
	v.TypeOfError = v.makeerror()
	v.TypeOfInterface = v.makeinterface()
	// critical! trying to rebuild "error" type creates a non-indentical copy... lots of conversions would fail
	v.cache(v.TypeOfError.ReflectType(), v.TypeOfError)
	v.cache(v.TypeOfInterface.ReflectType(), v.TypeOfInterface)
	return v
}

func NewUniverse() *Universe {
	v := &Universe{}
	return v.init()
}

const MaxDepth = int(^uint(0) >> 1)

var (
	universe = (&Universe{ThreadSafe: true}).Init()

	reflectTypeOfInterfaceHeader = reflect.TypeOf(InterfaceHeader{})
)

func DefaultUniverse() *Universe {
	return universe
}

// Bits returns the size of the type in bits.
// It panics if the type's Kind is not one of the
// sized or unsized Int, Uint, Float, or Complex kinds.
func (t *xtype) Bits() int {
	return t.rtype.Bits()
}

// Align returns the alignment in bytes of a value of
// this type when allocated in memory.
func (t *xtype) Align() int {
	return t.rtype.Align()
}

// FieldAlign returns the alignment in bytes of a value of
// this type when used as a field in a struct.
func (t *xtype) FieldAlign() int {
	return t.rtype.FieldAlign()
}
