// +build arm64

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
 * func_arm64.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

const SUPPORTED = true

func (asm *Asm) prologue() *Asm {
	return asm.Uint32(0xf94007fd) // ldr x29, [sp, #8]
}

func (asm *Asm) epilogue() *Asm {
	return asm.Uint32(0xd65f03c0) // ret
}
