// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!windows

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
 * sys_dummy.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

func nop(*uint64) {
}

func (asm *Asm) Func() func(*uint64) {
	return nop
}
