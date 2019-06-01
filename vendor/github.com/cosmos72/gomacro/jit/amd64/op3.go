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
 * op3.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package amd64

import (
	"errors"
)

// ============================================================================
// tree-arg instruction

// dst = a OP b
func (arch Amd64) Op3(asm *Asm, op Op3, a Arg, b Arg, dst Arg) *Asm {
	arch.op3(asm, op, a, b, dst)
	return asm
}

var op3KindError = errors.New("Amd64.op3: arguments a, b, dst must have the same kind")

func (arch Amd64) op3(asm *Asm, op Op3, a Arg, b Arg, dst Arg) Amd64 {
	// validate kinds
	switch op {
	case SHL3, SHR3:
		assert(a.Kind() == dst.Kind())
		assert(!b.Kind().Signed())
	case SETIDX, GETIDX:
		assert(a.Kind().Size() == 8)
	default:
		if a.Kind() != dst.Kind() || b.Kind() != dst.Kind() {
			panic(op3KindError)
			// assert(a.Kind() == dst.Kind())
			// assert(b.Kind() == dst.Kind())
		}
	}
	// validate dst
	switch dst.(type) {
	case Reg, Mem:
		break
	case Const:
		if op != SETIDX {
			errorf("destination cannot be a constant: %v %v, %v, %v", op, a, b, dst)
		}
	default:
		errorf("unknown destination type %T, expecting Reg or Mem: %v %v, %v, %v", dst, op, a, b, dst)
	}
	if asm.Optimize3(op, a, b, dst) {
		return arch
	}
	switch op {
	case MUL3:
		return arch.mul3(asm, a, b, dst)
	case DIV3, REM3:
		return arch.divrem(asm, op, a, b, dst)
	case GETIDX, SETIDX:
		return arch.index(asm, op, a, b, dst)
	}
	op2 := Op2(op)
	if a == dst {
		arch.op2(asm, op2, b, dst)
	} else if op.IsCommutative() && b == dst {
		arch.op2(asm, op2, a, dst)
	} else if r, ok := dst.(Reg); ok && r.RegId() != b.RegId() {
		arch.mov(asm, a, dst).op2(asm, op2, b, dst)
	} else {
		r := asm.RegAlloc(dst.Kind())
		arch.mov(asm, a, r).op2(asm, op2, b, r).mov(asm, r, dst)
		asm.RegFree(r)
	}
	return arch
}

// either a[b] = val or val = a[b]
func (arch Amd64) index(asm *Asm, op Op3, a Arg, b Arg, val Arg) Amd64 {
	var ra, rb, rval Reg
	var rconst bool

	switch val := val.(type) {
	case Reg:
		rval = val
	case Const:
		// only SETIDX
		cval := val.Val()
		if cval == int64(int32(cval)) {
			rconst = true
		} else {
			rval = asm.RegAlloc(val.Kind())
			defer asm.RegFree(rval)
		}
	case Mem:
		rval = asm.RegAlloc(Uint64)
		defer asm.RegFree(rval)
	}

	switch a := a.(type) {
	case Reg:
		ra = a
	case Mem:
		ra = asm.RegAlloc(Uint64)
		arch.load(asm, a, ra)
		defer asm.RegFree(ra)
	case Const:
		// depending on b's type, could be optimized as MOV Mem Reg
		ra = asm.RegAlloc(Uint64)
		arch.movConstReg(asm, a, ra)
		defer asm.RegFree(ra)
	}

	// b.Kind().Size() could be < 8
	switch b := b.(type) {
	case Reg:
		rbx := b
		rb = MakeReg(rbx.RegId(), Uint64)
		arch.cast(asm, rbx, rb)
	case Mem:
		rb = asm.RegAlloc(Uint64)
		arch.cast(asm, b, rb)
		defer asm.RegFree(rb)
	case Const:
		k := val.Kind()
		idx := b.Val()
		off := idx * int64(k.Size())
		if idx == int64(int32(idx)) && off == int64(int32(off)) {
			if op == GETIDX {
				// optimize as MOV Mem Reg
				arch.load(asm, MakeMem(int32(off), ra.RegId(), k), rval)
				arch.mov(asm, rval, val)
			} else if rconst {
				// optimize as MOV Const Mem
				arch.movConstMem(asm, val.(Const), MakeMem(int32(off), ra.RegId(), k))
			} else {
				// optimize as MOV Reg Mem
				arch.mov(asm, val, rval)
				arch.store(asm, rval, MakeMem(int32(off), ra.RegId(), k))
			}
			return arch
		}
		rb = asm.RegAlloc(Uint64)
		arch.movConstReg(asm, b, rb)
		defer asm.RegFree(rb)
	}

	switch op {
	case SETIDX:
		if rconst {
			return arch.indexRegRegConst(asm, ra, rb, val.(Const))
		}
		arch.mov(asm, rval, rval)
		arch.indexRegRegReg(asm, op, ra, rb, rval)
	case GETIDX:
		arch.indexRegRegReg(asm, op, ra, rb, rval)
		arch.mov(asm, rval, val)
	}
	return arch
}

// either a[b] = val or val = a[b]
func (arch Amd64) indexRegRegReg(asm *Asm, op Op3, a Reg, b Reg, val Reg) Amd64 {
	assert(b.RegId() != RSP)

	alo, ahi := lohi(a)
	blo, bhi := lohi(b)
	vlo, vhi := lohi(val)
	hi := vhi<<2 | bhi<<1 | ahi

	kind := val.Kind()
	size := kind.Size()
	scalebit := map[Size]uint8{1: 0x00, 2: 0x40, 4: 0x80, 8: 0xC0}[size]

	offlen, offbit := offlen(MakeMem(0, a.RegId(), kind), a.RegId())

	op_ := uint8(0x88)
	if op == GETIDX {
		op_ = 0x8A
	}

	switch size {
	case 1:
		if hi == 0 {
			asm.Bytes(op_, offbit|vlo<<3|0x04, scalebit|blo<<3|alo)
		} else {
			asm.Bytes(0x40|hi, op_, offbit|vlo<<3|0x04, scalebit|blo<<3|alo)
		}
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if hi == 0 {
			asm.Bytes(op_|1, offbit|vlo<<3|0x04, scalebit|blo<<3|alo)
		} else {
			asm.Bytes(0x40|hi, op_|1, offbit|vlo<<3|0x04, scalebit|blo<<3|alo)
		}
	case 8:
		asm.Bytes(0x48|hi, op_|1, offbit|vlo<<3|0x04, scalebit|blo<<3|alo)
	}
	switch offlen {
	case 1:
		asm.Int8(0)
	case 4:
		asm.Int32(0)
	}
	return arch
}

// a[b] = const
func (arch Amd64) indexRegRegConst(asm *Asm, a Reg, b Reg, c Const) Amd64 {
	assert(b.RegId() != RSP)

	alo, ahi := lohi(a)
	blo, bhi := lohi(b)
	hi := bhi<<1 | ahi

	kind := c.Kind()
	size := kind.Size()
	scalebit := map[Size]uint8{1: 0x00, 2: 0x40, 4: 0x80, 8: 0xC0}[size]

	offlen, offbit := offlen(MakeMem(0, a.RegId(), kind), a.RegId())

	switch size {
	case 1:
		if hi == 0 {
			asm.Bytes(0xC6, offbit|0x04, scalebit|blo<<3|alo)
		} else {
			asm.Bytes(0x40|hi, 0xC6, offbit|0x04, scalebit|blo<<3|alo)
		}
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if hi == 0 {
			asm.Bytes(0xC7, offbit|0x04, scalebit|blo<<3|alo)
		} else {
			asm.Bytes(0x40|hi, 0xC7, offbit|0x04, scalebit|blo<<3|alo)
		}
	case 8:
		asm.Bytes(0x48|hi, 0xC7, offbit|0x04, scalebit|blo<<3|alo)
	}
	switch offlen {
	case 1:
		asm.Int8(0)
	case 4:
		asm.Int32(0)
	}
	switch size {
	case 1:
		asm.Int8(int8(c.Val()))
	case 2:
		asm.Int16(int16(c.Val()))
	case 4, 8:
		asm.Int32(int32(c.Val()))
	}
	return arch
}
