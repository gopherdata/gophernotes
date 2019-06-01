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
 * set.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

func (asm *Asm) Load(dst Reg, src Arg) *Asm {
	return asm.load(asm.reg(dst), src)
}

func (asm *Asm) LoadConst(dst Reg, val int64) *Asm {
	return asm.loadConst(asm.reg(dst), val)
}

func (asm *Asm) Store(dst *Var, src Arg) *Asm {
	return asm.store(dst, src)
}

func (asm *Asm) Zero(dst *Var) *Asm {
	return asm.store(dst, Int64(0))
}
