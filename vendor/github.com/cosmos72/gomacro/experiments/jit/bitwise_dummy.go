// +build !amd64,!arm64

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
 * bitwise_dummy.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

// %rax &= a
func (asm *Asm) And(z Reg, a Arg) *Asm {
	return asm
}

// %rax |= a
func (asm *Asm) Or(z Reg, a Arg) *Asm {
	return asm
}

// %rax ^= a
func (asm *Asm) Xor(z Reg, a Arg) *Asm {
	return asm
}

// %rax &^= a
func (asm *Asm) Andnot(z Reg, a Arg) *Asm {
	return asm
}

// %reg_z = ^ %reg_z
func (asm *Asm) Not(z Reg) *Asm {
	return asm
}
