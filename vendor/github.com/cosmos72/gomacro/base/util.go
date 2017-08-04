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
 * util.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	r "reflect"
)

func PackValues(val0 r.Value, vals []r.Value) []r.Value {
	if len(vals) == 0 && val0 != None {
		vals = []r.Value{val0}
	}
	return vals
}

func UnpackValues(vals []r.Value) (r.Value, []r.Value) {
	val0 := None
	if len(vals) > 0 {
		val0 = vals[0]
	}
	return val0, vals
}

// ValueInterface() is a zero-value-safe version of reflect.Value.Interface()
func ValueInterface(v r.Value) interface{} {
	if v == Nil || v == None || !v.IsValid() || !v.CanInterface() {
		return nil
	}
	return v.Interface()
}

// ValueType() is a zero-value-safe version of reflect.Value.Type()
func ValueType(v r.Value) r.Type {
	if v == Nil || v == None || !v.IsValid() {
		return nil
	}
	return v.Type()
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
