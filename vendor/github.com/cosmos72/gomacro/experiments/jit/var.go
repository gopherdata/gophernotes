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
 * var.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

import "reflect"

func NewVar(idx uint16) *Var {
	return &Var{desc: desc{idx: idx}}
}

// implement Arg interface
func (v *Var) reg(asm *Asm) hwReg {
	return noReg
}

func (v *Var) Const() bool {
	return false
}

func (v *Var) Kind() reflect.Kind {
	return v.kind
}
