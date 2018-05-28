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
 * constant.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	r "reflect"
)

type none struct{}

// the following constants must match with github.com/cosmos72/gomacro/xreflect/gensym.go
const (
	StrGensymInterface string = "\U0001202A" // name of extra struct field needed by the interpreter when creating interface proxies
	StrGensymPrivate   string = "\U00012038" // prefix to generate names for unexported struct fields
	StrGensymAnonymous string = "\U00012039" // prefix to generate names for anonymous struct fields
	StrGensym          string = "\U00012035" // prefix to generate names in macros
	// the symbols above are chosen somewhat arbitrarily. Reasons:
	// * accepted by Go compiler as identifier names in source code
	// * belong to an ancient language no longer spoken, so hopefully low collision risk
	// * outside Unicode basic plane, so hopefully lower collision risk
	// * relatively simple glyph picture

	MaxUint16 = ^uint16(0)
	MaxUint   = ^uint(0)
	MaxInt    = int(MaxUint >> 1)
	MinInt    = ^MaxInt
)

var (
	Nil = r.Value{}

	None = r.ValueOf(none{}) // used to indicate "no value"

	True  = r.ValueOf(true)
	False = r.ValueOf(false)

	One = r.ValueOf(1)

	TypeOfInt   = r.TypeOf(int(0))
	TypeOfInt8  = r.TypeOf(int8(0))
	TypeOfInt16 = r.TypeOf(int16(0))
	TypeOfInt32 = r.TypeOf(int32(0))
	TypeOfInt64 = r.TypeOf(int64(0))

	TypeOfUint    = r.TypeOf(uint(0))
	TypeOfUint8   = r.TypeOf(uint8(0))
	TypeOfUint16  = r.TypeOf(uint16(0))
	TypeOfUint32  = r.TypeOf(uint32(0))
	TypeOfUint64  = r.TypeOf(uint64(0))
	TypeOfUintptr = r.TypeOf(uintptr(0))

	TypeOfFloat32    = r.TypeOf(float32(0))
	TypeOfFloat64    = r.TypeOf(float64(0))
	TypeOfComplex64  = r.TypeOf(complex64(0))
	TypeOfComplex128 = r.TypeOf(complex128(0))

	TypeOfBool        = r.TypeOf(false)
	TypeOfByte        = r.TypeOf(byte(0))
	TypeOfRune        = r.TypeOf(rune(0))
	TypeOfString      = r.TypeOf("")
	TypeOfInterface   = r.TypeOf((*interface{})(nil)).Elem()
	TypeOfError       = r.TypeOf((*error)(nil)).Elem()
	TypeOfDeferFunc   = r.TypeOf(func() {})
	TypeOfReflectType = r.TypeOf((*r.Type)(nil)).Elem() // inception

	TypeOfPtrInt   = r.TypeOf((*int)(nil))
	TypeOfPtrInt8  = r.TypeOf((*int8)(nil))
	TypeOfPtrInt16 = r.TypeOf((*int16)(nil))
	TypeOfPtrInt32 = r.TypeOf((*int32)(nil))
	TypeOfPtrInt64 = r.TypeOf((*int64)(nil))

	TypeOfPtrUint    = r.TypeOf((*uint)(nil))
	TypeOfPtrUint8   = r.TypeOf((*uint8)(nil))
	TypeOfPtrUint16  = r.TypeOf((*uint16)(nil))
	TypeOfPtrUint32  = r.TypeOf((*uint32)(nil))
	TypeOfPtrUint64  = r.TypeOf((*uint64)(nil))
	TypeOfPtrUintptr = r.TypeOf((*uintptr)(nil))

	TypeOfPtrFloat32    = r.TypeOf((*float32)(nil))
	TypeOfPtrFloat64    = r.TypeOf((*float64)(nil))
	TypeOfPtrComplex64  = r.TypeOf((*complex64)(nil))
	TypeOfPtrComplex128 = r.TypeOf((*complex128)(nil))

	TypeOfPtrBool   = r.TypeOf((*bool)(nil))
	TypeOfPtrString = r.TypeOf((*string)(nil))

	ZeroStrings = []string{}
	ZeroTypes   = []r.Type{}
	ZeroValues  = []r.Value{}
)
