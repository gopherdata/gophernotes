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
 * constants.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	r "reflect"
)

type none struct{}

const (
	StrGensymInterface = "\u0080"     // name of extra struct field needed by the interpreter when creating interface proxies
	StrGensymPrivate   = "\u00AD"     // prefix to generate names for unexported struct fields
	StrGensymEmbedded  = "\u00BB"     // prefix to generate names for embedded struct fields
	StrGensym          = "\U000124AD" // prefix to generate names in macros - arbitrarily chosen U+124AD CUNEIFORM SIGN ERIN2 X - reasons:
	// * accepted by Go compiler identifier name in source code
	// * belongs to an ancient language no longer spoken, so hopefully low collision risk
	// * outside Unicode basic place, so hopefully lower collision risk
	// * relatively simple glyph picture

	MaxUint16 = ^uint16(0)
	MaxUint   = ^uint(0)
	MinUint   = 0
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

	TypeOfSliceOfByte      = r.TypeOf([]byte{})
	TypeOfSliceOfInterface = r.TypeOf([]interface{}{})

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
