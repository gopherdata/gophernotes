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
 * hw_amd64.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

type hwOp = uint8

const (
	hwADD hwOp = 0
	hwOR  hwOp = 0x08
	// hwADC hwOp = 0x10 // add with carry
	// hwSBB hwOp = 0x18 // subtract with borrow
	hwAND hwOp = 0x20
	hwSUB hwOp = 0x28
	hwXOR hwOp = 0x30
	// hwCMP hwOp = 0x38 // compare, set flags
	// hwXCHG hwOp = 0x86 // exchange. xchg %reg, %reg has different encoding
	hwMOV hwOp = 0x88
)

const (
	noReg hwReg = iota
	rAX
	rCX
	rDX
	rBX
	rSP
	rBP
	rSI
	rDI
	rR8
	rR9
	rR10
	rR11
	rR12
	rR13
	rR14
	rR15
	rLo hwReg = rAX
	rHi hwReg = rR15
)

var alwaysLiveHwRegs = hwRegs{rSP: 1, rBP: 1, rDI: 1}

func (r hwReg) Valid() bool {
	return r >= rLo && r <= rHi
}

func (r hwReg) Validate() {
	if !r.Valid() {
		errorf("invalid register: %d", r)
	}
}

func (r hwReg) bits() uint8 {
	r.Validate()
	return uint8(r) - 1
}

func (r hwReg) lo() uint8 {
	return r.bits() & 0x7
}

func (r hwReg) hi() uint8 {
	return (r.bits() & 0x8) >> 3
}

func (r hwReg) lohi() (uint8, uint8) {
	bits := r.bits()
	return bits & 0x7, (bits & 0x8) >> 3
}

func (asm *Asm) lohi(g Reg) (uint8, uint8) {
	return asm.reg(g).lohi()
}
