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
 * dsl.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package common

func (asm *Asm) Assemble(args ...AsmCode) *Asm {
	n := len(args)
	for i := 0; i < n; i++ {
		i += asm.Op(args[i:]...)
	}
	return asm
}

func (asm *Asm) Op(args ...AsmCode) int {
	var n int
	switch op := args[0].(type) {
	case Op0:
		asm.Op0(op)
		n = 0
	case Op1:
		asm.Op1(op, asm.Arg(args[1]))
		n = 1
	case Op1Misc:
		asm.Op1Misc(op, args[1])
		n = 1
	case Op2Misc:
		asm.Op2Misc(op, args[1], args[2])
		n = 2
	case Op2:
		asm.Op2(op, asm.Arg(args[1]), asm.Arg(args[2]))
		n = 2
	case Op3:
		asm.Op3(op, asm.Arg(args[1]), asm.Arg(args[2]), asm.Arg(args[3]))
		n = 3
	case Op4:
		asm.Op4(op, asm.Arg(args[1]), asm.Arg(args[2]), asm.Arg(args[3]), asm.Arg(args[4]))
		n = 4
	default:
		errorf("syntax error: expecting Op0,Op1,Op1Misc,Op2Misc,Op2,Op3 or Op4 [args], found %v // %T", op, op)
	}
	return n
}
