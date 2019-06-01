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
 * add.c
 *
 *  Created on Feb 02, 2019
 *      Author Massimiliano Ghilardi
 */

package main

// go:nosplit
func Add8(a uint8, b uint8) uint8 {
	return a + b
}

// go:nosplit
func Add16(a uint16, b uint16) uint16 {
	return a + b
}

// go:nosplit
func Add32(a uint32, b uint32) uint32 {
	return a + b
}

// go:nosplit
func Add64(a uint64, b uint64) uint64 {
	return a + b
}
