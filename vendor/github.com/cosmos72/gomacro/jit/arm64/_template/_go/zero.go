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
 * zero.go
 *
 *  Created on Feb 07, 2019
 *      Author Massimiliano Ghilardi
 */

package main

// go:nosplit
func Zero0(env *Env) {
	env.IP++
}

// go:nosplit
func Zero8(env *Env) uint8 {
	return uint8(env.Ints[0])
}

// go:nosplit
func Zero16(env *Env) uint16 {
	return uint16(env.Ints[0])
}

// go:nosplit
func Zero32(env *Env) uint32 {
	return uint32(env.Ints[0])
}

// go:nosplit
func Zero64(env *Env) uint64 {
	return uint64(env.Ints[0])
}
