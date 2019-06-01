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
 * optimize.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"math/bits"
)

// compute base-2 logarithm of integer n.
// return 0, false if argument is not a power of 2
// used to optimize multiplication by a constant power of two.
func Log2Uint(n uint64) (uint8, bool) {
	if n == 0 || n&(n-1) != 0 {
		return 0, false
	}
	return uint8(bits.Len64(n) - 1), true
}

func (asm *Asm) Optimize2(op Op2, src Arg, dst Arg) bool {
	if src == dst {
		switch op {
		case AND2, OR2, MOV, CAST:
			return true // operation is nop
		case SUB2, XOR2, AND_NOT2:
			asm.Zero(dst)
			return true
		}
	}
	// more optimizations on MOV dst, dst
	if op == MOV && src.Kind().Size() == dst.Kind().Size() {
		switch src := src.(type) {
		case Reg:
			dst, ok := dst.(Reg)
			return ok && src.RegId() == dst.RegId()
		case Mem:
			dst, ok := dst.(Mem)
			return ok && src.RegId() == dst.RegId() && src.Offset() == dst.Offset()
		}
	}
	c, ok := src.(Const)
	if !ok {
		return false
	}
	n := c.Cast(Int64).val
	src = MakeConst(n, dst.Kind())
	switch op {
	case ADD2:
		switch n {
		case 0:
			return true
		case 1:
			asm.Op1(INC, dst)
			return true
		case -1:
			asm.Op1(DEC, dst)
			return true
		}
	case OR2:
		switch n {
		case 0:
			return true
		case -1:
			asm.Op2(MOV, src, dst)
			return true
		}
	case AND2:
		switch n {
		case 0:
			asm.Op2(MOV, src, dst)
			return true
		case -1:
			return true
		}
	case SUB2:
		switch n {
		case 0:
			return true
		case 1:
			asm.Op1(DEC, dst)
			return true
		case -1:
			asm.Op1(INC, dst)
			return true
		}
	case XOR2:
		switch n {
		case 0:
			return true
		case -1:
			asm.Op1(NOT1, dst)
			return true
		}
	case AND_NOT2:
		switch n {
		case 0:
			return true
		case -1:
			asm.Op2(MOV, src, dst)
			return true
		}
	case CAST:
		asm.Op2(MOV, src, dst)
		return true
	case SHL2, SHR2:
		switch n {
		case 0:
			return true
		}
	case MUL2:
		switch n {
		case 0:
			asm.Op2(MOV, src, dst)
			return true
		case 1:
			return true
		case -1:
			asm.Op1(NEG1, dst)
			return true
		}
	case DIV2:
		switch n {
		case 1:
			return true
		case -1:
			asm.Op1(NEG1, dst)
			return true
		}
	case MOV:
		if n == 0 {
			asm.Zero(dst)
			return true
		}
	}
	return false
}

func (op Op3) IsCommutative() bool {
	switch op {
	case ADD3, OR3, ADC3, AND3, XOR3, MUL3:
		return true
	}
	return false
}

func (asm *Asm) Optimize3(op Op3, a Arg, b Arg, dst Arg) bool {
	if a == b {
		switch op {
		case AND3, OR3:
			if b == dst {
				// operation is NOP
				return true
			}
			asm.Mov(a, dst)
			return true
		case SUB3, XOR3, AND_NOT3:
			asm.Zero(dst)
			return true
		}
	}
	c, ok := b.(Const)
	if !ok {
		if op.IsCommutative() {
			a, b = b, a
			c, ok = b.(Const)
		}
		if !ok {
			return false
		}
	}
	n := c.Cast(Int64).val
	c = MakeConst(n, dst.Kind())
	switch op {
	case ADD3:
		switch n {
		case 0:
			asm.Mov(a, dst)
			return true
		}
	case OR3:
		switch n {
		case 0:
			asm.Mov(a, dst)
			return true
		case -1:
			asm.Mov(c, dst)
			return true
		}
	case AND3:
		switch n {
		case 0:
			asm.Zero(dst)
			return true
		case -1:
			asm.Mov(a, dst)
			return true
		}
	case SUB3:
		switch n {
		case 0:
			asm.Mov(a, dst)
			return true
		}
	case XOR3:
		switch n {
		case 0:
			asm.Mov(a, dst)
			return true
		case -1:
			asm.Op2(NOT2, a, dst)
			return true
		}
	case AND_NOT3:
		switch n {
		case 0:
			asm.Mov(a, dst)
			return true
		case -1:
			asm.Zero(dst)
			return true
		}
	case SHL3, SHR3:
		switch n {
		case 0:
			asm.Mov(a, dst)
			return true
		}
	case MUL3:
		switch n {
		case 0:
			asm.Zero(dst)
			return true
		case 1:
			asm.Mov(a, dst)
			return true
		case -1:
			asm.Op2(NEG2, a, dst)
			return true
		}
	case DIV3:
		switch n {
		case 1:
			asm.Mov(a, dst)
			return true
		case -1:
			asm.Op2(NEG2, a, dst)
			return true
		}
	}
	return false
}
