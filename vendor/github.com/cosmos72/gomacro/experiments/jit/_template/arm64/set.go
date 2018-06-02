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
func load_16() uint64 {
	return 0xffff
}
//go:nosplit
func load_32() (uint64, uint64) {
	return 0x10000, 0xffffffff
}
//go:nosplit
func load_48() (uint64, uint64) {
	return 0x100000000, 0xffffffffffff
}
//go:nosplit
func load_64() (uint64,uint64) {
	return 0x1000000000000, 0xffffffffffffffff
}
