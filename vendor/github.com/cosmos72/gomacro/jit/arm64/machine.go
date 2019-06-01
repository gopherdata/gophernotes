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
 *  Created on May 26, 2018
 *      Author Massimiliano Ghilardi
 */

package arm64

import (
	"fmt"
)

const (
	noregid = RegId(ARM64-1)<<8 + iota
	X0
	X1
	X2
	X3
	X4
	X5
	X6
	X7
	X8
	X9
	X10
	X11
	X12
	X13
	X14
	X15
	X16
	X17
	X18
	X19
	X20
	X21
	X22
	X23
	X24
	X25
	X26
	X27
	X28
	X29
	X30
	XZR, XSP = iota, iota // depending on context, zero register or stack pointer
	RLo      = X0
	RHi      = XZR
	// stack pointer
	RSP = XSP
	// suggested register to point to local variables
	RVAR = X29
)

var (
	regName4 = makeRegNames("w")
	regName8 = makeRegNames("x")
)

func makeRegNames(prefix string) []string {
	name := make([]string, RHi+1)
	for id := RLo; id < RHi; id++ {
		name[id] = fmt.Sprint(prefix, int(id)-1)
	}
	name[RHi] = prefix + "zr"
	return name
}

// return the bitmask to be or-ed to the instruction
// to specify the registers width
func kbit(r Reg) uint32 {
	return uint32(r.Kind().Size()) & 8 << 28
}

// validate and return uint32 mask representing r.id
// note that XSP/XZR is not considered valid
func val(r Reg) uint32 {
	r.Validate()
	return uint32(r.RegId()) - 1
}

// validate and return uint32 mask representing r.id
// if allowX31 is true, also allows r.id == XSP/XZR
func valOrX31(id RegId, allowX31 bool) uint32 {
	if !allowX31 || id != XZR {
		id.Validate()
	}
	return uint32(id) - 1
}
