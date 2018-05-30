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
 * literal.go
 *
 *  Created on Apr 11, 2017
 *      Author Massimiliano Ghilardi
 */

package base

import (
	r "reflect"
)

func KindToCategory(k r.Kind) r.Kind {
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
	k = KindToCategory(k)
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
	t := ValueType(v)
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
