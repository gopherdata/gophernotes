/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * z_test.go
 *
 *  Created on May 26, 2018
 *      Author Massimiliano Ghilardi
 */

package zero

import (
	r "reflect"
	"testing"
	"unsafe"
)

var rbasictypes = []r.Type{
	r.Bool:          r.TypeOf(bool(false)),
	r.Int:           r.TypeOf(int(0)),
	r.Int8:          r.TypeOf(int8(0)),
	r.Int16:         r.TypeOf(int16(0)),
	r.Int32:         r.TypeOf(int32(0)),
	r.Int64:         r.TypeOf(int64(0)),
	r.Uint:          r.TypeOf(uint(0)),
	r.Uint8:         r.TypeOf(uint8(0)),
	r.Uint16:        r.TypeOf(uint16(0)),
	r.Uint32:        r.TypeOf(uint32(0)),
	r.Uint64:        r.TypeOf(uint64(0)),
	r.Uintptr:       r.TypeOf(uintptr(0)),
	r.Float32:       r.TypeOf(float32(0)),
	r.Float64:       r.TypeOf(float64(0)),
	r.Complex64:     r.TypeOf(complex64(0)),
	r.Complex128:    r.TypeOf(complex128(0)),
	r.String:        r.TypeOf(string("")),
	r.UnsafePointer: r.TypeOf(unsafe.Pointer(nil)),
}

func TestFunctionZero(t *testing.T) {
	var targs, trets = []r.Type{nil}, []r.Type{nil}
	lo, hi := r.Bool, r.UnsafePointer

	for karg := lo; karg <= hi; karg++ {
		targ := rbasictypes[karg]
		if targ == nil {
			continue
		}
		targs[0] = targ
		for kret := lo; kret <= hi; kret++ {
			tret := rbasictypes[kret]
			if tret == nil {
				continue
			}
			trets[0] = tret
			typ := r.FuncOf(targs, trets, false)
			f := FunctionZero(typ)
			if !f.IsValid() {
				t.Logf("FunctionZero(%v)\treturned nil function", typ)
				continue
			}
			rets := f.Call([]r.Value{r.Zero(targ)})
			if len(rets) != 1 {
				t.Errorf("f(%v) %v\treturned %d values, expecting one: %v", targ, tret, len(rets), rets)
				continue
			}
			if zero := r.Zero(tret); rets[0].Interface() != zero.Interface() {
				t.Errorf("f(%v) %v\treturned value %v, expecting zero value %v", targ, tret, rets[0], zero)
				continue
			}
			// t.Logf("f(%v) %v\treturned zero value %v", targ, tret, rets[0])
		}
	}
}
