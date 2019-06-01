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
 * constant.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	r "reflect"

	"github.com/cosmos72/gomacro/base/reflect"
)

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
	Nil = reflect.Nil

	None = reflect.None // used to indicate "no value"

	True  = r.ValueOf(true)
	False = r.ValueOf(false)

	One = r.ValueOf(1)

	TypeOfInt   = reflect.TypeOfInt
	TypeOfInt8  = reflect.TypeOfInt8
	TypeOfInt16 = reflect.TypeOfInt16
	TypeOfInt32 = reflect.TypeOfInt32
	TypeOfInt64 = reflect.TypeOfInt64

	TypeOfUint    = reflect.TypeOfUint
	TypeOfUint8   = reflect.TypeOfUint8
	TypeOfUint16  = reflect.TypeOfUint16
	TypeOfUint32  = reflect.TypeOfUint32
	TypeOfUint64  = reflect.TypeOfUint64
	TypeOfUintptr = reflect.TypeOfUintptr

	TypeOfFloat32    = reflect.TypeOfFloat32
	TypeOfFloat64    = reflect.TypeOfFloat64
	TypeOfComplex64  = reflect.TypeOfComplex64
	TypeOfComplex128 = reflect.TypeOfComplex128

	TypeOfBool   = reflect.TypeOfBool
	TypeOfString = reflect.TypeOfString

	TypeOfByte        = r.TypeOf(byte(0))
	TypeOfRune        = r.TypeOf(rune(0))
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
