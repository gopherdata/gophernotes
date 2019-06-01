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
 * hwregs.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

func newHwRegs(rs ...hwReg) *hwRegs {
	var ret hwRegs
	for _, r := range rs {
		ret.Set(r)
	}
	return &ret
}

func (rs *hwRegs) InitLive() {
	*rs = alwaysLiveHwRegs
}

func (rs *hwRegs) Contains(r hwReg) bool {
	return r >= rLo && r <= rHi && rs[r] != 0
}

func (rs *hwRegs) Set(r hwReg) {
	if r >= rLo && r <= rHi {
		rs[r]++
	}
}

func (rs *hwRegs) Unset(r hwReg) {
	if rs.Contains(r) {
		rs[r]--
	}
}

func (rs *hwRegs) Alloc() hwReg {
	for r := rLo; r <= rHi; r++ {
		if rs[r] == 0 {
			rs[r]++
			return r
		}
	}
	errorf("no free registers")
	return noReg
}

func (rs *hwRegs) Free(r hwReg) {
	rs.Unset(r)
}
