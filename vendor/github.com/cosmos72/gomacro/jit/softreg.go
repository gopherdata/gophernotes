/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * softreg.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"github.com/cosmos72/gomacro/jit/common"
)

func (c *Comp) NewSoftReg(kind Kind) SoftReg {
	id := c.nextSoftReg
	c.nextSoftReg++
	s := MakeSoftReg(id, kind)
	c.code.SoftReg(common.ALLOC, s)
	return s
}

func (c *Comp) newTempReg(kind Kind) SoftReg {
	id := c.nextTempReg
	c.nextTempReg++
	s := MakeSoftReg(id, kind)
	c.code.SoftReg(common.ALLOC, s)
	return s
}

func (c *Comp) FreeSoftReg(s SoftReg) {
	if s.Valid() && !s.IsTemp() {
		if s.Id()+1 == c.nextSoftReg {
			c.nextSoftReg--
		}
		c.code.SoftReg(common.FREE, s)
	}
}

func (c *Comp) freeTempReg(s SoftReg) {
	if s.Valid() && s.IsTemp() {
		if s.Id()+1 == c.nextTempReg {
			c.nextTempReg--
		}
		c.code.SoftReg(common.FREE, s)
	}
}

// alloc or free soft reg
func (c *Comp) SoftReg(inst Inst1Misc, s SoftReg) {
	if s.Valid() {
		c.code.SoftReg(inst.Asm(), s)
	}
}
