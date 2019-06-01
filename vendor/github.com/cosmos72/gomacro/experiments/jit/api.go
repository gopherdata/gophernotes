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
 * api.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"reflect"
)

// hardware register. implementation is architecture-dependent
type hwReg uint8

type hwRegs [rHi + 1]uint32 // hwReg -> use count

type hwRegCounter struct {
	hwReg
	count uint32
}

// hardware memory location.
type hwMem struct {
	off uint32
	siz uint8 // 1, 2, 4 or 8
	reg hwReg
}

// software-defined register. mapped to hardware register by Asm
type Reg uint32

type Const struct {
	kind reflect.Kind
	val  int64
}

type desc struct {
	kind reflect.Kind
	idx  uint16
	upn  uint16
}

type Var struct {
	desc
}

type Arg interface {
	reg(asm *Asm) hwReg // noReg if not a register
	Const() bool
	Kind() reflect.Kind
}

type Code []uint8

type Save struct {
	start, idx, end uint16 // memory area where spill registers can be saved
}

type Asm struct {
	code    Code
	hwRegs  hwRegs
	regs    map[Reg]hwRegCounter
	regNext Reg // first available register among jit-reserved ones
	save    Save
}
