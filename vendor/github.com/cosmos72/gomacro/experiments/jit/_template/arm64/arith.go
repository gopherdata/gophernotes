/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018 Massimiliano Ghilardi
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
func add_1(z int64) int64 {
	return z + 1
}
//go:nosplit
func add_2(z int64) int64 {
	return z + 2
}
//go:nosplit
func add_3(z int64) int64 {
	return z + 3
}
//go:nosplit
func add_255(z int64) int64 {
	return z + 255
}
//go:nosplit
func add_256(z int64) int64 {
	return z + 256
}
//go:nosplit
func add_2048(z int64) int64 {
	return z + 2048
}
//go:nosplit
func add_4095(z int64) int64 {
	return z + 4095
}



//go:nosplit
func sub_1(z int64) int64 {
	return z - 1
}
//go:nosplit
func sub_2(z int64) int64 {
	return z - 2
}
//go:nosplit
func sub_3(z int64) int64 {
	return z - 3
}
//go:nosplit
func sub_255(z int64) int64 {
	return z - 255
}
//go:nosplit
func sub_256(z int64) int64 {
	return z - 256
}
//go:nosplit
func sub_2048(z int64) int64 {
	return z - 2048
}
//go:nosplit
func sub_4095(z int64) int64 {
	return z - 4095
}




//go:nosplit
func mul_2(z int64) int64 {
	return z * 2
}
//go:nosplit
func mul_3(z int64) int64 {
	return z * 3
}
//go:nosplit
func mul_333(z int64) int64 {
	return z * 333
}
//go:nosplit
func mul_56789(z int64) int64 {
	return z * 56789
}




//go:nosplit
func quo(z int64, a int64) int64 {
	return z / a
}
//go:nosplit
func rem(z int64, a int64) int64 {
	return z % a
}
