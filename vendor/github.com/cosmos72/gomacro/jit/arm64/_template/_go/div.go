/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * div.go
 *
 *  Created on Feb 08, 2019
 *      Author Massimiliano Ghilardi
 */

package main

// go:nosplit
func Div8(a int8, b int8) int8 {
	return a / b
}

// go:nosplit
func Div16(a int16, b int16) int16 {
	return a / b
}

// go:nosplit
func Div32(a int32, b int32) int32 {
	return a / b
}

// go:nosplit
func Div64(a int64, b int64) int64 {
	return a / b
}

// go:nosplit
func UDiv8(a uint8, b uint8) uint8 {
	return a / b
}

// go:nosplit
func UDiv16(a uint16, b uint16) uint16 {
	return a / b
}

// go:nosplit
func UDiv32(a uint32, b uint32) uint32 {
	return a / b
}

// go:nosplit
func UDiv64(a uint64, b uint64) uint64 {
	return a / b
}
