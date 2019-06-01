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
 * reg.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"reflect"
)

const (
	NoReg Reg = 0          // means "no register"
	RegLo Reg = 1          // first user-available register = 1
	RegHi Reg = 0x80000000 // last user-available register = 0x80000000
)

// implement Arg interface
func (g Reg) reg(asm *Asm) hwReg {
	return asm.reg(g)
}

func (g Reg) Const() bool {
	return false
}

func (g Reg) Kind() reflect.Kind {
	// update after implementing MMX and XMM
	return reflect.Int64
}
