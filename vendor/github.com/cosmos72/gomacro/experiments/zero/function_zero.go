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
 * function_zero.go
 *
 *  Created on May 26, 2018
 *      Author Massimiliano Ghilardi
 */

package zero

import (
	r "reflect"
	"unsafe"
)

// Functions that return the zero value of their return type

// Note: since Go GC relies on per-function stack maps to scavenge pointers,
// these functions can be used with pointer arguments and return types
// only because they do *NOT* actually access or create any pointer value.

// We use an integer-based 8-byte struct instead of complex128
// to avoid complications due to floating point registers:
// since IEE 754 represents floating-point zero as 'all bits are 0',
// this acceptable for zero values
type uint128 struct {
	a, b uint64
}

// return no values

//go:nosplit
func zeroArg0Ret0() {
}

//go:nosplit
func zeroArg1Ret0(uint8) {
}

//go:nosplit
func zeroArg2Ret0(uint16) {
}

//go:nosplit
func zeroArg4Ret0(uint32) {
}

//go:nosplit
func zeroArg8Ret0(uint64) {
}

//go:nosplit
func zeroArg16Ret0(uint128) {
}

// return a 1-byte zero value

//go:nosplit
func zeroArg0Ret1() (ret uint8) {
	return
}

//go:nosplit
func zeroArg1Ret1(uint8) (ret uint8) {
	return
}

//go:nosplit
func zeroArg2Ret1(uint16) (ret uint8) {
	return
}

//go:nosplit
func zeroArg4Ret1(uint32) (ret uint8) {
	return
}

//go:nosplit
func zeroArg8Ret1(uint64) (ret uint8) {
	return
}

//go:nosplit
func zeroArg16Ret1(uint128) (ret uint8) {
	return
}

// return a 2-byte zero value

//go:nosplit
func zeroArg0Ret2() (ret uint16) {
	return
}

//go:nosplit
func zeroArg1Ret2(uint8) (ret uint16) {
	return
}

//go:nosplit
func zeroArg2Ret2(uint16) (ret uint16) {
	return
}

//go:nosplit
func zeroArg4Ret2(uint32) (ret uint16) {
	return
}

//go:nosplit
func zeroArg8Ret2(uint64) (ret uint16) {
	return
}

//go:nosplit
func zeroArg16Ret2(uint128) (ret uint16) {
	return
}

// return a 4-byte zero value

//go:nosplit
func zeroArg0Ret4() (ret uint32) {
	return
}

//go:nosplit
func zeroArg1Ret4(uint8) (ret uint32) {
	return
}

//go:nosplit
func zeroArg2Ret4(uint16) (ret uint32) {
	return
}

//go:nosplit
func zeroArg4Ret4(uint32) (ret uint32) {
	return
}

//go:nosplit
func zeroArg8Ret4(uint64) (ret uint32) {
	return
}

//go:nosplit
func zeroArg16Ret4(uint128) (ret uint32) {
	return
}

// return a 8-byte zero value

//go:nosplit
func zeroArg0Ret8() (ret uint64) {
	return
}

//go:nosplit
func zeroArg1Ret8(uint8) (ret uint64) {
	return
}

//go:nosplit
func zeroArg2Ret8(uint16) (ret uint64) {
	return
}

//go:nosplit
func zeroArg4Ret8(uint32) (ret uint64) {
	return
}

//go:nosplit
func zeroArg8Ret8(uint64) (ret uint64) {
	return
}

//go:nosplit
func zeroArg16Ret8(uint128) (ret uint64) {
	return
}

// return a 16-byte zero value

//go:nosplit
func zeroArg0Ret16() (ret uint128) {
	return
}

//go:nosplit
func zeroArg1Ret16(uint8) (ret uint128) {
	return
}

//go:nosplit
func zeroArg2Ret16(uint16) (ret uint128) {
	return
}

//go:nosplit
func zeroArg4Ret16(uint32) (ret uint128) {
	return
}

//go:nosplit
func zeroArg8Ret16(uint64) (ret uint128) {
	return
}

//go:nosplit
func zeroArg16Ret16(uint128) (ret uint128) {
	return
}

var functionZero [6][6]uintptr

func init() {
	v00, v01, v02, v03, v04, v05 := zeroArg0Ret0, zeroArg0Ret1, zeroArg0Ret2, zeroArg0Ret4, zeroArg0Ret8, zeroArg0Ret16

	functionZero[0][0] = *(*uintptr)(unsafe.Pointer(&v00))
	functionZero[0][1] = *(*uintptr)(unsafe.Pointer(&v01))
	functionZero[0][2] = *(*uintptr)(unsafe.Pointer(&v02))
	functionZero[0][3] = *(*uintptr)(unsafe.Pointer(&v03))
	functionZero[0][4] = *(*uintptr)(unsafe.Pointer(&v04))
	functionZero[0][5] = *(*uintptr)(unsafe.Pointer(&v05))

	v10, v11, v12, v13, v14, v15 := zeroArg1Ret0, zeroArg1Ret1, zeroArg1Ret2, zeroArg1Ret4, zeroArg1Ret8, zeroArg1Ret16

	functionZero[1][0] = *(*uintptr)(unsafe.Pointer(&v10))
	functionZero[1][1] = *(*uintptr)(unsafe.Pointer(&v11))
	functionZero[1][2] = *(*uintptr)(unsafe.Pointer(&v12))
	functionZero[1][3] = *(*uintptr)(unsafe.Pointer(&v13))
	functionZero[1][4] = *(*uintptr)(unsafe.Pointer(&v14))
	functionZero[1][5] = *(*uintptr)(unsafe.Pointer(&v15))

	v20, v21, v22, v23, v24, v25 := zeroArg2Ret0, zeroArg2Ret1, zeroArg2Ret2, zeroArg2Ret4, zeroArg2Ret8, zeroArg2Ret16

	functionZero[2][0] = *(*uintptr)(unsafe.Pointer(&v20))
	functionZero[2][1] = *(*uintptr)(unsafe.Pointer(&v21))
	functionZero[2][2] = *(*uintptr)(unsafe.Pointer(&v22))
	functionZero[2][3] = *(*uintptr)(unsafe.Pointer(&v23))
	functionZero[2][4] = *(*uintptr)(unsafe.Pointer(&v24))
	functionZero[2][5] = *(*uintptr)(unsafe.Pointer(&v25))

	v30, v31, v32, v33, v34, v35 := zeroArg4Ret0, zeroArg4Ret1, zeroArg4Ret2, zeroArg4Ret4, zeroArg4Ret8, zeroArg4Ret16

	functionZero[3][0] = *(*uintptr)(unsafe.Pointer(&v30))
	functionZero[3][1] = *(*uintptr)(unsafe.Pointer(&v31))
	functionZero[3][2] = *(*uintptr)(unsafe.Pointer(&v32))
	functionZero[3][3] = *(*uintptr)(unsafe.Pointer(&v33))
	functionZero[3][4] = *(*uintptr)(unsafe.Pointer(&v34))
	functionZero[3][5] = *(*uintptr)(unsafe.Pointer(&v35))

	v40, v41, v42, v43, v44, v45 := zeroArg8Ret0, zeroArg8Ret1, zeroArg8Ret2, zeroArg8Ret4, zeroArg8Ret8, zeroArg8Ret16

	functionZero[4][0] = *(*uintptr)(unsafe.Pointer(&v40))
	functionZero[4][1] = *(*uintptr)(unsafe.Pointer(&v41))
	functionZero[4][2] = *(*uintptr)(unsafe.Pointer(&v42))
	functionZero[4][3] = *(*uintptr)(unsafe.Pointer(&v43))
	functionZero[4][4] = *(*uintptr)(unsafe.Pointer(&v44))
	functionZero[4][5] = *(*uintptr)(unsafe.Pointer(&v45))

	v50, v51, v52, v53, v54, v55 := zeroArg16Ret0, zeroArg16Ret1, zeroArg16Ret2, zeroArg16Ret4, zeroArg16Ret8, zeroArg16Ret16

	functionZero[5][0] = *(*uintptr)(unsafe.Pointer(&v50))
	functionZero[5][1] = *(*uintptr)(unsafe.Pointer(&v51))
	functionZero[5][2] = *(*uintptr)(unsafe.Pointer(&v52))
	functionZero[5][3] = *(*uintptr)(unsafe.Pointer(&v53))
	functionZero[5][4] = *(*uintptr)(unsafe.Pointer(&v54))
	functionZero[5][5] = *(*uintptr)(unsafe.Pointer(&v55))
}

// if available, return the zero function matching function type 't'
func FunctionZero(t r.Type) r.Value {
	if t.NumIn() > 1 || t.NumOut() > 1 {
		return r.Value{}
	}
	var insize, outsize uintptr
	if t.NumIn() != 0 {
		insize = t.In(0).Size()
	}
	if t.NumOut() != 0 {
		outsize = t.Out(0).Size()
	}

	var i, o uint8
	ni, no := uint8(len(functionZero)), uint8(len(functionZero[0]))
	if insize != 0 {
		for i = 1; i < ni; i++ {
			if insize == 1<<(i-1) {
				break
			}
		}
	}
	if outsize != 0 {
		for o = 1; o < no; o++ {
			if outsize == 1<<(o-1) {
				break
			}
		}
	}
	if i >= ni || o >= no {
		return r.Value{}
	}

	ret := r.Zero(t)

	u := (*unsafeReflectValue)(unsafe.Pointer(&ret))
	u.ptr = functionZero[i][o]
	u.flag = uintptr(r.Func)

	return ret
}

type unsafeReflectValue struct {
	t    *struct{} // actually *reflect.rtype
	ptr  uintptr
	flag uintptr
}
