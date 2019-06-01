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
 * const.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

import "reflect"

func Int64(val int64) *Const {
	return &Const{val: val}
}

// implement Arg interface
func (c *Const) reg(asm *Asm) hwReg {
	return noReg
}

func (c *Const) Const() bool {
	return true
}

func (c *Const) Kind() reflect.Kind {
	return c.kind
}
