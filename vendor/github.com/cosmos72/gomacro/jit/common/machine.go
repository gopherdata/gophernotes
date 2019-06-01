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
 * machine.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package common

const (
	NoRegId RegId = 0
)

func (asm *Asm) RegIdConfig() RegIdConfig {
	return asm.arch.RegIdConfig()
}

func (asm *Asm) Op0(op Op0) *Asm {
	asm.arch.Op0(asm, op)
	return asm
}

func (asm *Asm) Op1(op Op1, dst Arg) *Asm {
	asm.arch.Op1(asm, op, dst)
	return asm
}

func (asm *Asm) Op2(op Op2, src Arg, dst Arg) *Asm {
	return asm.arch.Op2(asm, op, src, dst)
}

func (asm *Asm) Op3(op Op3, a Arg, b Arg, dst Arg) *Asm {
	return asm.arch.Op3(asm, op, a, b, dst)
}

func (asm *Asm) Op4(op Op4, a Arg, b Arg, c Arg, dst Arg) *Asm {
	return asm.arch.Op4(asm, op, a, b, c, dst)
}

func (asm *Asm) Zero(dst Arg) *Asm {
	return asm.arch.Op1(asm, ZERO, dst)
}

func (asm *Asm) Mov(src Arg, dst Arg) *Asm {
	return asm.arch.Op2(asm, MOV, src, dst)
}

func (asm *Asm) Load(src Mem, dst Reg) *Asm {
	return asm.arch.Op2(asm, MOV, src, dst)
}

func (asm *Asm) Store(src Reg, dst Mem) *Asm {
	return asm.arch.Op2(asm, MOV, src, dst)
}

func (asm *Asm) Cast(src Arg, dst Arg) *Asm {
	return asm.arch.Op2(asm, CAST, src, dst)
}

func (asm *Asm) Prologue() *Asm {
	return asm.arch.Prologue(asm)
}

func (asm *Asm) Epilogue() *Asm {
	return asm.arch.Epilogue(asm)
}
