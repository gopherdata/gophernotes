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

package amd64

// ============================================================================
// no-arg instruction

func (arch Amd64) Op0(asm *Asm, op Op0) *Asm {
	switch op {
	case BAD:
		asm.Bytes(0x0F, 0x0B) // UD2
	case RET:
		asm.Byte(0xC3)
	case NOP:
		asm.Byte(0x90)
	default:
		errorf("unknown Op0 instruction: %v", op)
	}
	return asm
}
