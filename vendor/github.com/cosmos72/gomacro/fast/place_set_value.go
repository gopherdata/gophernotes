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
