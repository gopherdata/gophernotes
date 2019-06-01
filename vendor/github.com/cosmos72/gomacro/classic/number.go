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
 * number.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	r "reflect"
)

func (env *Env) toInt(xv r.Value) (int64, bool) {
	switch xv.Kind() {
	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		u := xv.Uint()
		i := int64(u)
		if uint64(i) != u {
			env.Warnf("value %d overflows int64, truncated to %d", u, i)
		}
		return i, true
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		return xv.Int(), true
	default:
		return 0, false
	}
}

func (env *Env) toFloat(xv r.Value) (float64, bool) {
	switch xv.Kind() {
	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		return float64(xv.Uint()), true
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		return float64(xv.Int()), true
	case r.Float32, r.Float64:
		return xv.Float(), true
	default:
		return 0.0, false
	}
}

func (env *Env) toComplex(xv r.Value) (complex128, bool) {
	switch xv.Kind() {
	case r.Complex64, r.Complex128:
		return xv.Complex(), true
	default:
		f, ok := env.toFloat(xv)
		return complex(f, 0.0), ok
	}
}
