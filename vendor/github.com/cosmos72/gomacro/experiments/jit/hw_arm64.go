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
 * hw_arm64.go
 *
 *  Created on May 26, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

const (
	noReg hwReg = iota
	x0
	x1
	x2
	x3
	x4
	x5
	x6
	x7
	x8
	x9
	x10
	x11
	x12
	x13
	x14
	x15
	x16
	x17
	x18
	x19
	x20
	x21
	x22
	x23
	x24
	x25
	x26
	x27
	x28
	x29
	x30
	rLo hwReg = x0
	rHi hwReg = x30
)

var alwaysLiveHwRegs = hwRegs{
	x28: 1, // pointer to goroutine-local data
	x29: 1, // jit *uint64 pointer-to-variables
	x30: 1, // link register?
}

func (r hwReg) Valid() bool {
	return r >= rLo && r <= rHi
}

func (r hwReg) Validate() {
	if !r.Valid() {
		errorf("invalid register: %d", r)
	}
}

func (r hwReg) lo() uint32 {
	r.Validate()
	return uint32(r) - 1
}

func (asm *Asm) lo(g Reg) uint32 {
	return asm.reg(g).lo()
}
