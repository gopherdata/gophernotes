// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

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
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *
 * place_set_value.go
 *
 *  Created on May 29, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

// placeSetValue compiles 'place = value' where value is a reflect.Value passed at runtime.
// Used to assign places with the result of multi-valued expressions,
// and to implement multiple assignment place1, place2... = expr1, expr2...
func (c *Comp) placeSetValue(place *Place) func(lhs, key, val r.Value) {
	rtype := place.Type.ReflectType()

	if place.MapKey != nil {
		zero := r.Zero(rtype)
		return func(lhs, key, val r.Value) {
			if val == Nil || val == None {
				val = zero
			} else if val.Type() != rtype {
				val = val.Convert(rtype)
			}
			lhs.SetMapIndex(key, val)
		}
	}
	var ret func(r.Value, r.Value, r.Value)
	switch KindToCategory(rtype.Kind()) {
	case r.Bool:
		ret = func(lhs, key, val r.Value) {
			lhs.SetBool(val.Bool())
		}
	case r.Int:
		ret = func(lhs, key, val r.Value) {
			lhs.SetInt(val.Int())
		}
	case r.Uint:
		ret = func(lhs, key, val r.Value) {
			lhs.SetUint(val.Uint())
		}
	case r.Float64:
		ret = func(lhs, key, val r.Value) {
			lhs.SetFloat(val.Float())
		}
	case r.Complex128:
		ret = func(lhs, key, val r.Value) {
			lhs.SetComplex(val.Complex())
		}
	case r.String:
		ret = func(lhs, key, val r.Value) {
			lhs.SetString(val.String())
		}
	default:
		zero := r.Zero(rtype)
		ret = func(lhs, key, val r.Value) {
			if val == Nil || val == None {
				val = zero
			} else if val.Type() != rtype {
				val = val.Convert(rtype)
			}
			lhs.Set(val)
		}
	}
	return ret
}
