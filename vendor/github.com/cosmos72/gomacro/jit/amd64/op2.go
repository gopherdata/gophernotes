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
 * op2.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package amd64

// ============================================================================
// two-arg instruction

var op2Val = [256]uint8{
	ADD2: 0x00,
	OR2:  0x08,
	ADC2: 0x10, // add with carry
	SBB2: 0x18, // subtract with borrow
	AND2: 0x20,
	SUB2: 0x28,
	XOR2: 0x30,
	// CMP: 0x38, // compare, set flags
	// XCHG: 0x86, // exchange. xchg %reg, %reg has different encoding
	MOV:  0x88,
	LEA2: 0x8D,
	CAST: 0xB6, // sign extend, zero extend or narrow
	SHL2: 0xE0, // shift left. has different encoding
	SHR2: 0xE8, // shift right. has different encoding
	MUL2: 0xF6,
	DIV2: 0xFE, // TODO divide
	REM2: 0xFF, // TODO remainder

	NEG2: 0x40,
	NOT2: 0x48,
}

func op2val(op Op2) uint8 {
	val := op2Val[op]
	// ADD.val() is zero
	if val == 0 && op != ADD2 {
		errorf("unknown Op2 instruction: %v", op)
	}
	return val
}

// ============================================================================
// dst OP= src
func (arch Amd64) Op2(asm *Asm, op Op2, src Arg, dst Arg) *Asm {
	arch.op2(asm, op, src, dst)
	return asm
}

func (arch Amd64) op2(asm *Asm, op Op2, src Arg, dst Arg) Amd64 {
	// validate kinds
	switch op {
	case CAST:
		if SizeOf(src) != SizeOf(dst) {
			return arch.cast(asm, src, dst)
		}
		op = MOV
		fallthrough
	case MOV:
		assert(SizeOf(src) == SizeOf(dst))
	case SHL2, SHR2:
		assert(!src.Kind().Signed())
	default:
		assert(src.Kind() == dst.Kind())
	}
	// validate dst
	switch dst.(type) {
	case Reg, Mem:
		break
	case Const:
		errorf("destination cannot be a constant: %v %v, %v", op, src, dst)
	default:
		errorf("unknown destination type %T, expecting Reg or Mem: %v %v, %v", dst, op, src, dst)
	}

	if asm.Optimize2(op, src, dst) {
		return arch
	}

	switch op {
	case DIV2, REM2:
		return arch.divrem(asm, Op3(op), dst, src, dst)
	case NEG2, NOT2:
		op1 := Op1(op) // NEG2 -> NEG, NOT2 -> NOT
		if src == dst {
			return arch.op1(asm, op1, dst)
		} else {
			return arch.mov(asm, src, dst).op1(asm, op1, dst)
		}
	case AND_NOT2:
		// no assembler instruction => emulate it
		if csrc, ok := src.(Const); ok {
			src = MakeConst(^csrc.Val(), csrc.Kind())
			op = AND2
		} else {
			r := asm.RegAlloc(src.Kind())
			arch.mov(asm, src, r).op1(asm, NEG1, r).op2(asm, AND2, r, dst)
			asm.RegFree(r)
			return arch
		}
	}
	switch dst := dst.(type) {
	case Reg:
		switch src := src.(type) {
		case Reg:
			arch.op2RegReg(asm, op, src, dst)
		case Mem:
			arch.op2MemReg(asm, op, src, dst)
		case Const:
			arch.op2ConstReg(asm, op, src, dst)
		default:
			errorf("unknown source type %T, expecting Reg, Mem or Const: %v %v, %v", src, op, src, dst)
		}
	case Mem:
		switch src := src.(type) {
		case Reg:
			arch.op2RegMem(asm, op, src, dst)
		case Mem:
			arch.op2MemMem(asm, op, src, dst)
		case Const:
			arch.op2ConstMem(asm, op, src, dst)
		default:
			errorf("unknown source type %T, expecting Reg, Mem or Const: %v %v, %v", src, op, src, dst)
		}
	}
	return arch
}

// %reg_dst OP= const
func (arch Amd64) op2ConstReg(asm *Asm, op Op2, c Const, dst Reg) Amd64 {
	switch op {
	case MOV:
		return arch.movConstReg(asm, c, dst)
	case SHL2, SHR2:
		return arch.shiftConstReg(asm, op, c, dst)
	case MUL2:
		return arch.mul2ConstReg(asm, c, dst)
	}
	assert(op != LEA2)

	switch dst.Kind().Size() {
	case 1:
		arch.op2ConstReg8(asm, op, c, dst)
	case 2:
		arch.op2ConstReg16(asm, op, c, dst)
	case 4:
		arch.op2ConstReg32(asm, op, c, dst)
	case 8:
		arch.op2ConstReg64(asm, op, c, dst)
	}
	return arch
}

func (arch Amd64) op2ConstReg8(asm *Asm, op Op2, c Const, dst Reg) Amd64 {
	op_ := op2val(op)
	dlo, dhi := lohi(dst)
	cval := c.Val()
	if cval == int64(int8(cval)) {
		if dst.RegId() == RAX {
			asm.Bytes(0x04 | op_)
		} else if dst.RegId() < RSP {
			asm.Bytes(0x80, 0xC0|op_|dlo)
		} else {
			asm.Bytes(0x40|dhi, 0xC0|op_|dlo)
		}
		asm.Int8(int8(cval))
	} else {
		errorf("sign-extended constant overflows 8-bit destination: %v %v %v", op, c, dst)
	}
	return arch
}

func (arch Amd64) op2ConstReg16(asm *Asm, op Op2, c Const, dst Reg) Amd64 {
	op_ := op2val(op)
	dlo, dhi := lohi(dst)
	cval := c.Val()
	if cval == int64(int8(cval)) {
		if dhi == 0 {
			asm.Bytes(0x66, 0x83, 0xc0|op_|dlo)
		} else {
			asm.Bytes(0x66, 0x40|dhi, 0x83, 0xc0|op_|dlo)
		}
		asm.Int8(int8(cval))
	} else if cval == int64(int16(cval)) {
		if dst.RegId() == RAX {
			asm.Bytes(0x66, 0x05|op_)
		} else if dhi == 0 {
			asm.Bytes(0x66, 0x81, 0xc0|op_|dlo)
		} else {
			asm.Bytes(0x66, 0x40|dhi, 0x81, 0xc0|op_|dlo)
		}
		asm.Int16(int16(cval))
	} else {
		errorf("sign-extended constant overflows 16-bit destination: %v %v %v", op, c, dst)
	}
	return arch
}

func (arch Amd64) op2ConstReg32(asm *Asm, op Op2, c Const, dst Reg) Amd64 {
	op_ := op2val(op)
	dlo, dhi := lohi(dst)
	cval := c.Val()
	if cval == int64(int8(cval)) {
		if dhi == 0 {
			asm.Bytes(0x83, 0xc0|op_|dlo)
		} else {
			asm.Bytes(0x40|dhi, 0x83, 0xc0|op_|dlo)
		}
		asm.Int8(int8(cval))
	} else if cval == int64(int32(cval)) {
		if dst.RegId() == RAX {
			asm.Bytes(0x05 | op_)
		} else if dhi == 0 {
			asm.Bytes(0x81, 0xc0|op_|dlo)
		} else {
			asm.Bytes(0x40|dhi, 0x81, 0xc0|op_|dlo)
		}
		asm.Int32(int32(cval))
	} else {
		errorf("sign-extended constant overflows 32-bit destination: %v %v %v", op, c, dst)
	}
	return arch
}

func (arch Amd64) op2ConstReg64(asm *Asm, op Op2, c Const, dst Reg) Amd64 {
	op_ := op2val(op)
	dlo, dhi := lohi(dst)
	cval := c.Val()
	if cval == int64(int8(cval)) {
		asm.Bytes(0x48|dhi, 0x83, 0xC0|op_|dlo, uint8(int8(cval)))
	} else if cval == int64(int32(cval)) {
		if dst.RegId() == RAX {
			asm.Bytes(0x48|dhi, 0x05|op_)
		} else {
			asm.Bytes(0x48|dhi, 0x81, 0xC0|op_|dlo)
		}
		asm.Int32(int32(cval))
	} else {
		// constant is 64 bit wide, must load it in a register
		r := asm.RegAlloc(c.Kind())
		arch.movConstReg64(asm, c, r)
		arch.op2RegReg(asm, op, r, dst)
		asm.RegFree(r)
	}
	return arch
}

// %reg_dst OP= %reg_src
func (arch Amd64) op2RegReg(asm *Asm, op Op2, src Reg, dst Reg) Amd64 {
	switch op {
	case MUL2:
		return arch.mul2RegReg(asm, src, dst)
	case SHL2, SHR2:
		return arch.shiftRegReg(asm, op, src, dst)
	}
	assert(op != LEA2)

	op_ := op2val(op)
	slo, shi := lohi(src)
	dlo, dhi := lohi(dst)

	switch SizeOf(src) { // == SizeOf(dst)
	case 1:
		if src.RegId() < RSP && dst.RegId() < RSP {
			asm.Bytes(op_, 0xC0|dlo|slo<<3)
		} else {
			asm.Bytes(0x40|dhi|shi<<2, op_, 0xC0|dlo|slo<<3)
		}
	case 2:
		if dhi|shi<<2 == 0 {
			asm.Bytes(0x66, 0x01|op_, 0xC0|dlo|slo<<3)
		} else {
			asm.Bytes(0x66, 0x40|dhi|shi<<2, 0x01|op_, 0xC0|dlo|slo<<3)
		}
	case 4:
		if dhi|shi<<2 == 0 {
			asm.Bytes(0x01|op_, 0xC0|dlo|slo<<3)
		} else {
			asm.Bytes(0x40|dhi|shi<<2, 0x01|op_, 0xC0|dlo|slo<<3)
		}
	case 8:
		asm.Bytes(0x48|dhi|shi<<2, 0x01|op_, 0xC0|dlo|slo<<3)
	}
	return arch
}

// off_m(%reg_m) OP= %reg_src
func (arch Amd64) op2RegMem(asm *Asm, op Op2, src Reg, dst_m Mem) Amd64 {
	switch op {
	case MUL2:
		return arch.mul2RegMem(asm, src, dst_m)
	case SHL2, SHR2:
		return arch.shiftRegMem(asm, op, src, dst_m)
	}
	assert(op != LEA2)
	assert(SizeOf(src) == SizeOf(dst_m))

	op_ := op2val(op)
	dregid := dst_m.RegId()
	dlo, dhi := lohiId(dregid)
	slo, shi := lohi(src)

	siz := SizeOf(dst_m)
	offlen, offbit := offlen(dst_m, dregid)

	switch siz {
	case 1:
		if src.RegId() < RSP && dhi == 0 {
			asm.Bytes(op_, offbit|dlo|slo<<3)
		} else {
			asm.Bytes(0x40|dhi|shi<<2, op_, offbit|dlo|slo<<3)
		}
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi|shi<<2 == 0 {
			asm.Bytes(0x01|op_, offbit|dlo|slo<<3)
		} else {
			asm.Bytes(0x40|dhi|shi<<2, 0x01|op_, offbit|dlo|slo<<3)
		}
	case 8:
		asm.Bytes(0x48|dhi|shi<<2, 0x01|op_, offbit|dlo|slo<<3)
	}
	quirk24(asm, dregid)
	switch offlen {
	case 1:
		asm.Int8(int8(dst_m.Offset()))
	case 4:
		asm.Int32(dst_m.Offset())
	}
	return arch
}

// %reg_dst OP= off_m(%reg_m)
func (arch Amd64) op2MemReg(asm *Asm, op Op2, src_m Mem, dst Reg) Amd64 {
	switch op {
	case MUL2:
		return arch.mul2MemReg(asm, src_m, dst)
	case SHL2, SHR2:
		return arch.shiftMemReg(asm, op, src_m, dst)
	}
	op_ := op2val(op)
	sregid := src_m.RegId()
	dlo, dhi := lohi(dst)
	slo, shi := lohiId(sregid)

	assert(SizeOf(src_m) == SizeOf(dst))
	siz := SizeOf(src_m)
	offlen, offbit := offlen(src_m, sregid)

	if op == LEA2 {
		assert(siz == 8)
	}

	switch siz {
	case 1:
		if dst.RegId() < RSP && shi == 0 {
			asm.Bytes(0x02|op_, offbit|dlo<<3|slo)
		} else {
			asm.Bytes(0x40|dhi<<2|shi, 0x02|op_, offbit|dlo<<3|slo)
		}
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi|shi<<2 == 0 {
			asm.Bytes(0x03|op_, offbit|dlo<<3|slo)
		} else {
			asm.Bytes(0x40|dhi<<2|shi, 0x03|op_, offbit|dlo<<3|slo)
		}
	case 8:
		if op != LEA2 {
			op_ |= 0x03
		}
		asm.Bytes(0x48|dhi<<2|shi, op_, offbit|dlo<<3|slo)
	}
	quirk24(asm, sregid)
	switch offlen {
	case 1:
		asm.Int8(int8(src_m.Offset()))
	case 4:
		asm.Int32(src_m.Offset())
	}
	return arch
}

// off_dst(%reg_dst) OP= off_src(%reg_src)
func (arch Amd64) op2MemMem(asm *Asm, op Op2, src_m Mem, dst_m Mem) Amd64 {
	switch op {
	case MUL2:
		return arch.mul2MemMem(asm, src_m, dst_m)
	case SHL2, SHR2:
		return arch.shiftMemMem(asm, op, src_m, dst_m)
	}
	assert(op != LEA2)
	// not natively supported by amd64,
	// must load src in a register
	r := asm.RegAlloc(src_m.Kind())
	arch.op2MemReg(asm, MOV, src_m, r)
	arch.op2RegMem(asm, op, r, dst_m)
	asm.RegFree(r)
	return arch
}

// off_dst(%reg_dst) OP= const
func (arch Amd64) op2ConstMem(asm *Asm, op Op2, c Const, m Mem) Amd64 {
	switch op {
	case MOV:
		return arch.movConstMem(asm, c, m)
	case SHL2, SHR2:
		return arch.shiftConstMem(asm, op, c, m)
	case MUL2:
		return arch.mul2ConstMem(asm, c, m)
	}
	assert(op != LEA2)
	op_ := op2val(op)
	dregid := m.RegId()
	dlo, dhi := lohiId(dregid)
	offlen, offbit := offlen(m, dregid)
	cval := c.Val()
	size := SizeOf(m)
	switch size {
	case 1:
		if dhi == 0 {
			asm.Bytes(0x80, offbit|op_|dlo)
		} else {
			asm.Bytes(0x40|dhi, 0x80, offbit|op_|dlo)
		}
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if cval == int64(int8(cval)) {
			if dhi == 0 {
				asm.Bytes(0x83, offbit|op_|dlo)
			} else {
				asm.Bytes(0x40|dhi, 0x83, offbit|op_|dlo)
			}
		} else {
			if dhi == 0 {
				asm.Bytes(0x81, offbit|op_|dlo)
			} else {
				asm.Bytes(0x40|dhi, 0x81, offbit|op_|dlo)
			}
		}
	case 8:
		if cval == int64(int8(cval)) {
			asm.Bytes(0x48|dhi, 0x83, offbit|op_|dlo)
		} else if cval == int64(int32(cval)) {
			asm.Bytes(0x48|dhi, 0x81, offbit|op_|dlo)
		} else {
			// constant too large,
			// must copy it in a register
			r := asm.RegAlloc(c.Kind())
			arch.movConstReg64(asm, c, r).op2RegMem(asm, op, r, m)
			asm.RegFree(r)
			return arch
		}
	}
	quirk24(asm, dregid)
	switch offlen {
	case 1:
		asm.Int8(int8(m.Offset()))
	case 4:
		asm.Int32(m.Offset())
	}

	if size == 1 || cval == int64(int8(cval)) {
		asm.Int8(int8(cval))
	} else if size == 2 {
		asm.Int16(int16(cval))
	} else {
		asm.Int32(int32(cval))
	}
	return arch
}
