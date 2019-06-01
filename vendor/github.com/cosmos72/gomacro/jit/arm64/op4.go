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
 * op4.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arm64

// ============================================================================
// four-arg instruction

func (Arm64) Op4(asm *Asm, op Op4, a Arg, b Arg, c Arg, dst Arg) *Asm {
	errorf("unknown arm64 Op4 instruction: %v %v, %v, %v", a, b, c, dst)
	return asm
}
