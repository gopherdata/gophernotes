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
 * set_dummy.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

func (asm *Asm) load(dst hwReg, src Arg) *Asm {
	return asm
}

func (asm *Asm) loadConst(dst hwReg, val int64) *Asm {
	return asm
}

func (asm *Asm) mov(dst hwReg, src hwReg) *Asm {
	return asm
}

func (asm *Asm) store(dst *Var, src Arg) *Asm {
	return asm
}

func (asm *Asm) storeReg(dst *Var, src hwReg) *Asm {
	return asm
}
