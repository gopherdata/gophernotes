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
 * func_dummy.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

const SUPPORTED = false

func (asm *Asm) prologue() *Asm {
	return asm
}

func (asm *Asm) epilogue() *Asm {
	return asm
}
