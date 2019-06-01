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
 * op0.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arm64

// ============================================================================
// no-arg instruction

func op0val(op Op0) uint32 {
	var val uint32
	switch op {
	case BAD:
		val = 0x00000000
	case NOP:
		val = 0xD503201F
	case RET:
		val = 0xD65F03C0
	default:
		errorf("unknown Op0 instruction: %v", op)
	}
	return val
}

// ============================================================================
func (arch Arm64) Op0(asm *Asm, op Op0) *Asm {
	return asm.Uint32(op0val(op))
}
