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

package arm64

// ============================================================================
// three-arg instruction

var op3vals = map[Op3]uint32{
	AND3: 0x0A000000,
	ADD3: 0x0B000000,
	ADC3: 0x1A000000, // add with carry
	OR3:  0x2A000000,
	XOR3: 0x4A000000,
	SUB3: 0x4B000000,
	SBB3: 0x5A000000, // subtract with borrow

	SHL3: 0x1AC02000,
	// logical shr i.e. zero-extended right shift is 0x1AC02400
	// arithmetic shr i.e. sign-extended right shift is 0x1AC02800
	SHR3: 0x1AC02400,

	// MUL3 a,b,c is an alias for MADD4 xzr,a,b,c
	MUL3: 0x1B007C00,

	// unsigned division is 0x1AC00800
	// signed division is 0x1AC00C00
	DIV3: 0x1AC00800,

	// ldrb w0, [x0, x0]         is 0x38606800
	// ldrh w0, [x0, x0, lsl #1] is 0x78607800
	// ldr  w0, [x0, x0, lsl #2] is 0xB8607800
	// ldr  x0, [x0, x0, lsl #3] is 0xF8607800
	GETIDX: 0x38606800,

	// strb w0, [x0, x0]         is 0x38206800
	// strh w0, [x0, x0, lsl #1] is 0x78207800
	// str  w0, [x0, x0, lsl #2] is 0xB8207800
	// str  x0, [x0, x0, lsl #3] is 0xF8207800
	SETIDX: 0x38206800,
}

// return 32bit value used to encode operation on Reg,Reg,Reg
func op3val(op Op3) uint32 {
	var val uint32
	switch op {
	case REM3:
		errorf("internal error, operation %v needs to be implemented as {s|u}div followed by msub", op)
	default:
		val = op3vals[op]
		if val == 0 {
			errorf("unknown Op2 instruction: %v", op)
		}
	}
	return val
}

// return 32bit value used to encode operation on Reg,Const,Reg
func immval(op Op3) uint32 {
	switch op {
	case AND3:
		return 0x12 << 24
	case ADD3:
		return 0x11 << 24
	case SHL3, SHR3:
		// immediate constant is encoded differently
		return 0x53 << 24
	case OR3:
		return 0x32 << 24
	case XOR3:
		return 0x52 << 24
	case SUB3:
		return 0x51 << 24
	case GETIDX, SETIDX:
		return 1 // handled specially by caller
	default:
		errorf("cannot encode Op3 instruction %v with immediate constant", op)
		return 0
	}
}

// ============================================================================
func (arch Arm64) Op3(asm *Asm, op Op3, a Arg, b Arg, dst Arg) *Asm {
	arch.op3(asm, op, a, b, dst)
	return asm
}

func (arch Arm64) op3(asm *Asm, op Op3, a Arg, b Arg, dst Arg) Arm64 {
	// validate kinds
	switch op {
	case SHL3, SHR3:
		assert(a.Kind() == dst.Kind())
		assert(!b.Kind().Signed())
	case GETIDX, SETIDX:
		assert(a.Kind().Size() == 8)
	default:
		assert(a.Kind() == dst.Kind())
		assert(b.Kind() == dst.Kind())
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

	var ra, rb, rdst Reg
	var ta, tdst bool // Reg is a temporary register?

	switch dst := dst.(type) {
	case Reg:
		rdst = dst
	case Mem:
		rdst = asm.RegAlloc(dst.Kind())
		defer asm.RegFree(rdst)
		tdst = true
		if op == SETIDX {
			arch.load(asm, dst, rdst)
		}
	case Const:
		// op == SETIDX
		if dst.Val() == 0 {
			rdst = MakeReg(XZR, dst.Kind())
		} else {
			rdst = asm.RegAlloc(dst.Kind())
			defer asm.RegFree(rdst)
			tdst = true
			arch.movConstReg(asm, dst, rdst)
		}
	}
	var not_dst bool
	if op == AND_NOT3 {
		// must be emulated
		not_dst = true
		op = AND3
	}
	if op.IsCommutative() && a.Const() && !b.Const() {
		a, b = b, a
	}
	switch xa := a.(type) {
	case Reg:
		ra = xa
	case Mem:
		if tdst && op != SETIDX {
			// reuse temporary register rdst
			ra = rdst
		} else {
			ra = asm.RegAlloc(xa.Kind())
			defer asm.RegFree(ra)
		}
		ta = true
		arch.load(asm, xa, ra)
	case Const:
		ra = asm.RegAlloc(xa.Kind())
		defer asm.RegFree(ra)
		arch.movConstReg(asm, xa, ra)
	default:
		errorf("unknown argument type %T, expecting Const, Reg or Mem: %v %v, %v, %v", a, op, a, b, dst)
	}
	switch xb := b.(type) {
	case Reg:
		arch.op3RegRegReg(asm, op, ra, xb, rdst)
	case Mem:
		if tdst && op != SETIDX && (!ta || ra != rdst) {
			// reuse temporary register rdst
			rb = rdst
		} else {
			rb = asm.RegAlloc(xb.Kind())
			defer asm.RegFree(rb)
		}
		arch.load(asm, xb, rb).op3RegRegReg(asm, op, ra, rb, rdst)
	case Const:
		arch.op3RegConstReg(asm, op, ra, xb, rdst)
	default:
		errorf("unknown argument type %T, expecting Const, Reg or Mem: %v %v, %v, %v", b, op, a, b, dst)
	}
	if not_dst {
		// operation was AND_NOT3: negate dst
		arch.op2RegReg(asm, NOT2, rdst, rdst)
	}
	if tdst && op != SETIDX {
		arch.store(asm, rdst, dst.(Mem))
	}
	return arch
}

func (arch Arm64) op3RegRegReg(asm *Asm, op Op3, a Reg, b Reg, dst Reg) Arm64 {
	var opbits uint32

	switch op {
	case SHR3:
		if dst.Kind().Signed() {
			// arithmetic right shift
			opbits = 0xC00
		}
	case DIV3:
		if dst.Kind().Signed() {
			// signed division
			opbits = 0x400
		}
	case GETIDX, SETIDX:
		// kbit(dst) below is redundant
		switch dst.Kind().Size() {
		case 2:
			opbits = 0x40001000
		case 4:
			opbits = 0x80001000
		case 8:
			opbits = 0xC0001000
		}
	}

	arch.extendHighBits(asm, op, a)
	arch.extendHighBits(asm, op, b)
	// TODO: on arm64, division by zero returns zero instead of panic
	asm.Uint32(kbit(dst) | (opbits ^ op3val(op)) | val(b)<<16 | val(a)<<5 | valOrX31(dst.RegId(), op == SETIDX))
	return arch
}

func (arch Arm64) op3RegConstReg(asm *Asm, op Op3, a Reg, cb Const, dst Reg) Arm64 {
	if arch.tryOp3RegConstReg(asm, op, a, uint64(cb.Val()), dst) {
		return arch
	}
	rb := asm.RegAlloc(cb.Kind())
	arch.movConstReg(asm, cb, rb).op3RegRegReg(asm, op, a, rb, dst)
	asm.RegFree(rb)
	return arch
}

// try to encode operation into a single instruction.
// return false if not possible because constant must be loaded in a register
func (arch Arm64) tryOp3RegConstReg(asm *Asm, op Op3, a Reg, cval uint64, dst Reg) bool {
	imm3 := immediate3(op)
	immcval, ok := imm3.Encode64(cval, dst.Kind())
	if !ok {
		return false
	}
	opval := immval(op)

	kbit := kbit(dst)

	arch.extendHighBits(asm, op, a)
	switch imm3 {
	case Imm3AddSub, Imm3Bitwise:
		// for op == OR3, also accept a == XZR
		asm.Uint32(kbit | opval | immcval | valOrX31(a.RegId(), op == OR3)<<5 | val(dst))
	case Imm3Shift:
		arch.shiftRegConstReg(asm, op, a, cval, dst)
	case Imm3Index:
		kind := dst.Kind()
		op_ := load
		if op == SETIDX {
			op_ = store
		}
		// index must be multiplied by element size
		off := int32(immcval) * int32(kind.Size())
		mem := MakeMem(off, a.RegId(), kind)
		arch.loadstore(asm, op_, mem, dst)
	default:
		cb := ConstInt64(int64(cval))
		errorf("unknown constant encoding style %v of %v: %v %v, %v, %v", imm3, op, op, a, cb, dst)
	}
	return true
}

func (arch Arm64) shiftRegConstReg(asm *Asm, op Op3, a Reg, cval uint64, dst Reg) {
	dsize := dst.Kind().Size()
	if cval >= 8*uint64(dsize) {
		cb := ConstInt64(int64(cval))
		errorf("constant is out of range for shift: %v %v, %v, %v", op, a, cb, dst)
	}
	switch op {
	case SHL3:
		switch dsize {
		case 1, 2, 4:
			asm.Uint32(0x53000000 | uint32(32-cval)<<16 | uint32(31-cval)<<10 | val(a)<<5 | val(dst))
		case 8:
			asm.Uint32(0xD3400000 | uint32(64-cval)<<16 | uint32(63-cval)<<10 | val(a)<<5 | val(dst))
		}
	case SHR3:
		var unsignedbit uint32
		if !dst.Kind().Signed() {
			unsignedbit = 0x40 << 24
		}
		switch dsize {
		case 1, 2, 4:
			asm.Uint32(unsignedbit | 0x13007C00 | uint32(cval)<<16 | val(a)<<5 | val(dst))
		case 8:
			asm.Uint32(unsignedbit | 0x9340FC00 | uint32(cval)<<16 | val(a)<<5 | val(dst))
		}
	}
}

// arm64 has no native operations to work on 8 bit and 16 bit registers.
// Actually, it only has ldr (load) and str (store), but no arithmetic
// or bitwise operations.
// So we emulate them similarly to what compilers do:
// use 32 bit registers and ignore high bits in operands and results.
// Exceptions:
// 1) right-shift, division and remainder move data from high bits to low bits,
// so we must zero-extend or sign-extend the operands
// 2) GETIDX and SETIDX i.e. get or set array element,
//    require address and offset to be 64 bits
func (arch Arm64) extendHighBits(asm *Asm, op Op3, r Reg) Arm64 {
	rkind := r.Kind()
	rsize := rkind.Size()
	if rsize > 2 {
		return arch
	}
	switch op {
	case SHR3, DIV3, REM3:
		if rkind.Signed() {
			arch.cast(asm, r, MakeReg(r.RegId(), Int32))
		} else {
			arch.cast(asm, r, MakeReg(r.RegId(), Uint32))
		}
	case GETIDX, SETIDX:
		arch.cast(asm, r, MakeReg(r.RegId(), Uint64))
	}
	return arch
}

// ============================================================================

// style of immediate constants
// embeddable in a single Op3 instruction
type Immediate3 uint8

const (
	Imm3None    Immediate3 = iota
	Imm3AddSub             // 12 bits wide, possibly shifted left by 12 bits
	Imm3Bitwise            // complicated
	Imm3Shift              // 0..63 for 64 bit registers; 0..31 for 32 bit registers
	Imm3Index              // 0..4095 index for GETIDX or SETIDX
)

// return the style of immediate constants
// embeddable in a single Op3 instruction
func immediate3(op Op3) Immediate3 {
	switch op {
	case ADD3, SUB3:
		return Imm3AddSub
	case AND3, OR3, XOR3:
		return Imm3Bitwise
	case SHL3, SHR3:
		return Imm3Shift
	case GETIDX, SETIDX:
		return Imm3Index
	default:
		return Imm3None
	}
}

// return false if val cannot be encoded using imm style
func (imm Immediate3) Encode64(val uint64, kind Kind) (e uint32, ok bool) {
	kbits := kind.Size() * 8
	switch imm {
	case Imm3AddSub:
		// 12 bits wide, possibly shifted left by 12 bits
		if val == val&0xFFF {
			return uint32(val << 10), true
		} else if val == val&0xFFF000 {
			return 0x400000 | uint32(val>>2), true
		}
	case Imm3Bitwise:
		// complicated
		if kbits <= 32 {
			e, ok = imm3Bitwise32[val]
		} else {
			e, ok = imm3Bitwise64[val]
		}
		return e, ok
	case Imm3Shift:
		if val >= 0 && val < uint64(kbits) {
			// actual encoding is complicated
			return uint32(val), true
		}
	case Imm3Index:
		if val >= 0 && val <= 4095 {
			return uint32(val), true
		}
	}
	return 0, false
}

var imm3Bitwise32 = makeImm3Bitwise32()
var imm3Bitwise64 = makeImm3Bitwise64()

// compute all immediate constants that can be encoded
// in and, orr, eor on 32-bit registers
func makeImm3Bitwise32() map[uint64]uint32 {
	result := make(map[uint64]uint32)
	var bitmask uint64
	var size, length, e, rotation uint32
	for size = 2; size <= 32; size *= 2 {
		for length = 1; length < size; length++ {
			bitmask = 0xffffffff >> (32 - length)
			for e = size; e < 32; e *= 2 {
				bitmask |= bitmask << e
			}
			for rotation = 0; rotation < size; rotation++ {
				result[bitmask] = (size&64|rotation)<<16 | (0x7800*size)&0xF000 | (length-1)<<10
				bitmask = (bitmask >> 1) | (bitmask << 31)
			}
		}
	}
	return result
}

// compute all immediate constants that can be encoded
// in and, orr, eor on 64-bit registers
func makeImm3Bitwise64() map[uint64]uint32 {
	result := make(map[uint64]uint32)
	var bitmask uint64
	var size, length, e, rotation uint32
	for size = 2; size <= 64; size *= 2 {
		for length = 1; length < size; length++ {
			bitmask = 0xffffffffffffffff >> (64 - length)
			for e = size; e < 64; e *= 2 {
				bitmask |= bitmask << e
			}
			for rotation = 0; rotation < size; rotation++ {
				// #0x5555555555555555 => size=2, length=1, rotation=0 => 0x00f000
				// #0xaaaaaaaaaaaaaaaa => size=2, length=1, rotation=1 => 0x01f000
				// #0x1111111111111111 => size=4, length=1, rotation=0 => 0x00e000
				// #0x8888888888888888 => size=4, length=1, rotation=1 => 0x01e000
				// #0x4444444444444444 => size=4, length=1, rotation=2 => 0x02e000
				// #0x2222222222222222 => size=4, length=1, rotation=3 => 0x03e000
				// #0x3333333333333333 => size=4, length=2, rotation=0 => 0x00e400
				// #0x7777777777777777 => size=4, length=3, rotation=0 => 0x00e800
				// #0x0101010101010101 => size=8, length=1, rotation=0 => 0x00c000
				// #0x0303030303030303 => size=8, length=2, rotation=0 => 0x00c400
				// #0x0707070707070707 => size=8, length=3, rotation=0 => 0x00c800
				// #0x0f0f0f0f0f0f0f0f => size=8, length=4, rotation=0 => 0x00cc00
				// #0x1f1f1f1f1f1f1f1f => size=8, length=5, rotation=0 => 0x00d000
				// #0x3f3f3f3f3f3f3f3f => size=8, length=6, rotation=0 => 0x00d400
				// #0x7f7f7f7f7f7f7f7f => size=8, length=7, rotation=0 => 0x00d800
				// ...
				result[bitmask] = (size&64|rotation)<<16 | (0x7800*size)&0xF000 | (length-1)<<10
				bitmask = (bitmask >> 1) | (bitmask << 63)
			}
		}
	}
	return result
}
