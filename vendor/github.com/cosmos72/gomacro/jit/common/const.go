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
 * const.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

type Const struct {
	val  int64
	kind Kind
}

func (c Const) String() string {
	return fmt.Sprintf("%v(0x%x)", c.kind, c.val)
}

// implement Arg interface
func (c Const) RegId() RegId {
	return NoRegId
}

func (c Const) Kind() Kind {
	return c.kind
}

func (c Const) Const() bool {
	return true
}

func (c Const) asmcode() {
}

func (c Const) Val() int64 {
	return c.val
}

// convert Const to a different kind
func (c Const) Cast(to Kind) Const {
	val := c.val
	// sign-extend or zero-extend to 64 bits
	switch c.kind {
	case Bool:
		if val != 0 {
			// non-zero means true => convert to 1
			val = 1
		}
	case Int:
		val = int64(int(val))
	case Int8:
		val = int64(int8(val))
	case Int16:
		val = int64(int16(val))
	case Int32:
		val = int64(int32(val))
	case Int64:
		// nothing to do
	case Uint:
		val = int64(uint(val))
	case Uint8:
		val = int64(uint8(val))
	case Uint16:
		val = int64(uint16(val))
	case Uint32:
		val = int64(uint32(val))
	case Uint64:
		val = int64(uint64(val)) // should be a nop
	case Uintptr:
		val = int64(uintptr(val))
	case Float32, Float64:
		errorf("float constants not supported yet")
	default:
		errorf("invalid constant kind: %v", c)
	}
	// let caller truncate val as needed
	return Const{val: val, kind: to}
}

func MakeConst(val int64, kind Kind) Const {
	return Const{val: val, kind: kind}
}

func ConstInt(val int) Const {
	return Const{val: int64(val), kind: Int}
}

func ConstInt8(val int8) Const {
	return Const{val: int64(val), kind: Int8}
}

func ConstInt16(val int16) Const {
	return Const{val: int64(val), kind: Int16}
}

func ConstInt32(val int32) Const {
	return Const{val: int64(val), kind: Int32}
}

func ConstInt64(val int64) Const {
	return Const{val: val, kind: Int64}
}

func ConstUint(val uint) Const {
	return Const{val: int64(val), kind: Uint}
}

func ConstUint8(val uint8) Const {
	return Const{val: int64(val), kind: Uint8}
}

func ConstUint16(val uint16) Const {
	return Const{val: int64(val), kind: Uint16}
}

func ConstUint32(val uint32) Const {
	return Const{val: int64(val), kind: Uint32}
}

func ConstUint64(val uint64) Const {
	return Const{val: int64(val), kind: Uint64}
}

func ConstUintptr(val uintptr) Const {
	return Const{val: int64(val), kind: Uintptr}
}

// guaranteed to work only if val points to non-Go memory,
// as for example C/C++ memory
func ConstPointer(val *uint8) Const {
	return Const{val: int64(uintptr(unsafe.Pointer(val))), kind: Ptr}
}

var constInterfaceFail = fmt.Errorf("unsupported jit constant kind")

func ConstInterface(ival interface{}, t reflect.Type) (Const, error) {
	v := reflect.ValueOf(ival)
	kind := Kind(t.Kind())
	var val int64
	switch kind {
	case Bool:
		if v.Bool() {
			val = 1
		}
	case Int, Int8, Int16, Int32, Int64:
		val = v.Int()
	case Uint, Uint8, Uint16, Uint32, Uint64, Uintptr:
		val = int64(v.Uint())
	case Float32:
		val = int64(math.Float32bits(float32(v.Float())))
	case Float64:
		val = int64(math.Float64bits(v.Float()))
	default:
		return Const{}, constInterfaceFail
	}
	return Const{val: val, kind: kind}, nil
}
