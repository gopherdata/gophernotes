/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * kind.go
 *
 *  Created on Jan 24, 2019
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"reflect"
)

type Kind uint8 // narrow version of reflect.Kind

const (
	Invalid = Kind(reflect.Invalid)
	Bool    = Kind(reflect.Bool)
	Int     = Kind(reflect.Int)
	Int8    = Kind(reflect.Int8)
	Int16   = Kind(reflect.Int16)
	Int32   = Kind(reflect.Int32)
	Int64   = Kind(reflect.Int64)
	Uint    = Kind(reflect.Uint)
	Uint8   = Kind(reflect.Uint8)
	Uint16  = Kind(reflect.Uint16)
	Uint32  = Kind(reflect.Uint32)
	Uint64  = Kind(reflect.Uint64)
	Uintptr = Kind(reflect.Uintptr)
	Float32 = Kind(reflect.Float32)
	Float64 = Kind(reflect.Float64)
	Ptr     = Kind(reflect.Ptr)
	KLo     = Bool
	KHi     = Ptr
)

var ksize = [...]Size{
	Bool:    1,
	Int:     Size(reflect.TypeOf(int(0)).Size()),
	Int8:    1,
	Int16:   2,
	Int32:   4,
	Int64:   8,
	Uint:    Size(reflect.TypeOf(uint(0)).Size()),
	Uint8:   1,
	Uint16:  2,
	Uint32:  4,
	Uint64:  8,
	Uintptr: Size(reflect.TypeOf(uintptr(0)).Size()),
	Float32: 4,
	Float64: 8,
	Ptr:     Size(reflect.TypeOf((*int)(nil)).Size()),
}

func (k Kind) Size() Size {
	if k >= KLo && k <= KHi {
		return ksize[k]
	}
	return 0
}

func (k Kind) Signed() bool {
	switch k {
	case Bool, Uint, Uint8, Uint16, Uint32, Uint64, Uintptr, Ptr:
		return false
	default:
		return true
	}
}

func (k Kind) IsFloat() bool {
	return k == Float32 || k == Float64
}

func (k Kind) String() string {
	return reflect.Kind(k).String()
}

// implement AsmCode interface
func (k Kind) asmcode() {
}
