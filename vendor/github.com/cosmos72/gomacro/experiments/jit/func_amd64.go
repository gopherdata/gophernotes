// +build amd64

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
 * func_amd64.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

const SUPPORTED = true

func (asm *Asm) prologue() *Asm {
	return asm.Bytes(0x48, 0x8b, 0x7c, 0x24, 0x08) // movq 0x8(%rsp), %rdi
}

func (asm *Asm) epilogue() *Asm {
	return asm.Bytes(0xc3) // ret
}
