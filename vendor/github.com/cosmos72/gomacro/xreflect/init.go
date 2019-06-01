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
 * init.go
 *
 *  Created on May 19, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	r "reflect"
	"unsafe"

	"github.com/cosmos72/gomacro/go/types"
)

var rbasictypes = []r.Type{
	r.Bool:          r.TypeOf(bool(false)),
	r.Int:           r.TypeOf(int(0)),
	r.Int8:          r.TypeOf(int8(0)),
	r.Int16:         r.TypeOf(int16(0)),
	r.Int32:         r.TypeOf(int32(0)),
	r.Int64:         r.TypeOf(int64(0)),
	r.Uint:          r.TypeOf(uint(0)),
	r.Uint8:         r.TypeOf(uint8(0)),
	r.Uint16:        r.TypeOf(uint16(0)),
	r.Uint32:        r.TypeOf(uint32(0)),
	r.Uint64:        r.TypeOf(uint64(0)),
	r.Uintptr:       r.TypeOf(uintptr(0)),
	r.Float32:       r.TypeOf(float32(0)),
	r.Float64:       r.TypeOf(float64(0)),
	r.Complex64:     r.TypeOf(complex64(0)),
	r.Complex128:    r.TypeOf(complex128(0)),
	r.String:        r.TypeOf(string("")),
	r.UnsafePointer: r.TypeOf(unsafe.Pointer(nil)),
}

var ReflectBasicTypes = rbasictypes

func (v *Universe) makeBasicTypes() []Type {
	m := make([]Type, len(rbasictypes))
	for gkind := types.Bool; gkind <= types.UnsafePointer; gkind++ {
		kind := ToReflectKind(gkind)
		rtype := rbasictypes[kind]
		gtype := types.Typ[gkind]
		if rtype == nil || gtype == nil {
			continue
		}
		t := wrap(&xtype{kind: kind, gtype: gtype, rtype: rtype, universe: v})
		v.add(t)
		m[kind] = t
	}
	return m
}

func (v *Universe) makeError() Type {
	t := wrap(&xtype{
		kind:     r.Interface,
		gtype:    types.Universe.Lookup("error").Type(),
		rtype:    r.TypeOf((*error)(nil)).Elem(),
		universe: v,
	})
	v.add(t)
	return t
}

func (v *Universe) makeInterface() Type {
	t := wrap(&xtype{
		kind:     r.Interface,
		gtype:    types.NewInterface(nil, nil).Complete(),
		rtype:    rTypeOfInterface,
		universe: v,
	})
	v.add(t)
	return t
}

func (v *Universe) makeForward() Type {
	t := wrap(&xtype{
		kind:     r.Invalid,
		gtype:    types.NewInterface(nil, nil).Complete(),
		rtype:    rTypeOfForward,
		universe: v,
	})
	v.add(t)
	return t
}

func NewUniverse() *Universe {
	v := &Universe{}
	v.BasicTypes = v.makeBasicTypes()
	v.addBasicTypesMethodsCTI()
	v.TypeOfForward = v.makeForward()
	v.TypeOfInterface = v.makeInterface()
	v.TypeOfError = v.makeError()
	// critical! trying to rebuild "error" type creates a non-identical copy... lots of conversions would fail
	v.cache(v.TypeOfError.ReflectType(), v.TypeOfError)
	v.cache(v.TypeOfInterface.ReflectType(), v.TypeOfInterface)
	return v
}

const MaxDepth = int(^uint(0) >> 1)

var (
	rTypeOfInterface       = r.TypeOf((*interface{})(nil)).Elem()
	rTypeOfInterfaceHeader = r.TypeOf(InterfaceHeader{})
	rTypeOfForward         = r.TypeOf((*Forward)(nil)).Elem()
)

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
