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
 * reflect.go
 *
 *  Created on: Apr 11, 2017
 *      Author: Massimiliano Ghilardi
 */

package reflect

import (
	r "reflect"

	xr "github.com/cosmos72/gomacro/xreflect"
)

type none struct{}

var (
	Nil = r.Value{}

	None = r.ValueOf(none{}) // used to indicate "no value"

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

	TypeOfBool   = r.TypeOf(false)
	TypeOfString = r.TypeOf("")
)

func Category(k r.Kind) r.Kind {
	switch k {
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		return r.Int
	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		return r.Uint
	case r.Float32, r.Float64:
		return r.Float64
	case r.Complex64, r.Complex128:
		return r.Complex128
	default:
		return k
	}
}

func IsCategory(k r.Kind, categories ...r.Kind) bool {
	k = Category(k)
	for _, c := range categories {
		if k == c {
			return true
		}
	}
	return false
}

// IsOptimizedKind returns true if fast interpreter expects optimized expressions for given Kind
func IsOptimizedKind(k r.Kind) bool {
	switch k {
	case r.Bool, r.Int, r.Int8, r.Int16, r.Int32, r.Int64,
		r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr,
		r.Float32, r.Float64, r.Complex64, r.Complex128, r.String:
		return true
	}
	return false
}

var kindToType = [...]r.Type{
	r.Bool:       TypeOfBool,
	r.Int:        TypeOfInt,
	r.Int8:       TypeOfInt8,
	r.Int16:      TypeOfInt16,
	r.Int32:      TypeOfInt32,
	r.Int64:      TypeOfInt64,
	r.Uint:       TypeOfUint,
	r.Uint8:      TypeOfUint8,
	r.Uint16:     TypeOfUint16,
	r.Uint32:     TypeOfUint32,
	r.Uint64:     TypeOfUint64,
	r.Uintptr:    TypeOfUintptr,
	r.Float32:    TypeOfFloat32,
	r.Float64:    TypeOfFloat64,
	r.Complex64:  TypeOfComplex64,
	r.Complex128: TypeOfComplex128,
	r.String:     TypeOfString,
}

func KindToType(k r.Kind) r.Type {
	if int(k) < len(kindToType) {
		return kindToType[k]
	}
	return nil
}

// ConvertValue converts a value to type t and returns the converted value.
// extends reflect.Value.Convert(t) by allowing conversions from/to complex numbers.
// does not check for overflows or truncation.
func ConvertValue(v r.Value, to r.Type) r.Value {
	t := Type(v)
	if t == to {
		return v
	}
	if !t.ConvertibleTo(to) {
		// reflect.Value does not allow conversions from/to complex types
		k := v.Kind()
		kto := to.Kind()
		if IsCategory(kto, r.Complex128) {
			if IsCategory(k, r.Int, r.Uint, r.Float64) {
				temp := v.Convert(TypeOfFloat64).Float()
				v = r.ValueOf(complex(temp, 0.0))
			}
		} else if IsCategory(k, r.Complex128) {
			if IsCategory(k, r.Int, r.Uint, r.Float64) {
				temp := real(v.Complex())
				v = r.ValueOf(temp)
			}
		}
	}
	return v.Convert(to)
}

func PackValues(val0 r.Value, values []r.Value) []r.Value {
	if len(values) == 0 && val0 != None {
		values = []r.Value{val0}
	}
	return values
}

func PackTypes(typ0 xr.Type, types []xr.Type) []xr.Type {
	if len(types) == 0 && typ0 != nil {
		types = []xr.Type{typ0}
	}
	return types
}

func UnpackValues(vals []r.Value) (r.Value, []r.Value) {
	val0 := None
	if len(vals) > 0 {
		val0 = vals[0]
	}
	return val0, vals
}

// Interface() is a zero-value-safe version of reflect.Value.Interface()
func Interface(v r.Value) interface{} {
	if !v.IsValid() || !v.CanInterface() || v == None {
		return nil
	}
	return v.Interface()
}

// Type() is a zero-value-safe version of reflect.Value.Type()
func Type(value r.Value) r.Type {
	if !value.IsValid() || value == None {
		return nil
	}
	return value.Type()
}

func IsNillableKind(k r.Kind) bool {
	switch k {
	case r.Invalid, // nil is nillable...
		r.Chan, r.Func, r.Interface, r.Map, r.Ptr, r.Slice:
		return true
	default:
		return false
	}
}
