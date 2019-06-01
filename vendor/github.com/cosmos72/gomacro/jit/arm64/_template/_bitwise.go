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
 * arith.go
 *
 *  Created on May 27, 2018
 *      Author Massimiliano Ghilardi
 */

package arm64

//go:nosplit
func and_1(z int64) int64 {
	return z & 1
}

//go:nosplit
func and_2(z int64) int64 {
	return z & 2
}

//go:nosplit
func and_3(z int64) int64 {
	return z & 3
}

//go:nosplit
func and_255(z int64) int64 {
	return z & 255
}

//go:nosplit
func and_256(z int64) int64 {
	return z & 256
}

//go:nosplit
func and_2048(z int64) int64 {
	return z & 2048
}

//go:nosplit
func and_4095(z int64) int64 {
	return z & 4095
}

//go:nosplit
func or_1(z int64) int64 {
	return z | 1
}

//go:nosplit
func or_2(z int64) int64 {
	return z | 2
}

//go:nosplit
func or_3(z int64) int64 {
	return z | 3
}

//go:nosplit
func or_255(z int64) int64 {
	return z | 255
}

//go:nosplit
func or_256(z int64) int64 {
	return z | 256
}

//go:nosplit
func or_2048(z int64) int64 {
	return z | 2048
}

//go:nosplit
func or_4095(z int64) int64 {
	return z | 4095
}

//go:nosplit
func xor_1(z int64) int64 {
	return z ^ 1
}

//go:nosplit
func xor_2(z int64) int64 {
	return z ^ 2
}

//go:nosplit
func xor_3(z int64) int64 {
	return z ^ 3
}

//go:nosplit
func xor_255(z int64) int64 {
	return z ^ 255
}

//go:nosplit
func xor_256(z int64) int64 {
	return z ^ 256
}

//go:nosplit
func xor_2048(z int64) int64 {
	return z ^ 2048
}

//go:nosplit
func xor_4095(z int64) int64 {
	return z ^ 4095
}
