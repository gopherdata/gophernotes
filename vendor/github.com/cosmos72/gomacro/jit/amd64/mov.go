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
 * mov.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package amd64

// ============================================================================
func (arch Amd64) mov(asm *Asm, src Arg, dst Arg) Amd64 {
	return arch.op2(asm, MOV, src, dst)
}

func (arch Amd64) load(asm *Asm, src Mem, dst Reg) Amd64 {
	return arch.op2MemReg(asm, MOV, src, dst)
}
func (arch Amd64) store(asm *Asm, src Reg, dst Mem) Amd64 {
	return arch.op2RegMem(asm, MOV, src, dst)
}

// %reg_dst = const
func (arch Amd64) movConstReg(asm *Asm, c Const, dst Reg) Amd64 {
	if c.Val() == 0 {
		return arch.zeroReg(asm, dst)
	}
	switch dst.Kind().Size() {
	case 1:
		arch.movConstReg8(asm, c, dst)
	case 2:
		arch.movConstReg16(asm, c, dst)
	case 4:
		arch.movConstReg32(asm, c, dst)
	case 8:
		arch.movConstReg64(asm, c, dst)
	}
	return arch
}

func (arch Amd64) movConstReg8(asm *Asm, c Const, dst Reg) Amd64 {
	dlo, dhi := lohi(dst)
	cval := c.Val()
	if cval == int64(int8(cval)) {
		if dst.RegId() < RSP {
			asm.Bytes(0xB0 | dlo)
		} else {
			asm.Bytes(0x40|dhi, 0xB0|dlo)
		}
		asm.Int8(int8(cval))
	} else {
		errorf("sign-extended constant overflows 8-bit destination: %v %v %v", MOV, c, dst)
	}
	return arch
}

func (arch Amd64) movConstReg16(asm *Asm, c Const, dst Reg) Amd64 {
	dlo, dhi := lohi(dst)
	cval := c.Val()
	if cval == int64(int16(cval)) {
		if dhi == 0 {
			asm.Bytes(0x66, 0xB8|dlo)
		} else {
			asm.Bytes(0x66, 0x40|dhi, 0xB8|dlo)
		}
		asm.Int16(int16(cval))
	} else {
		errorf("sign-extended constant overflows 16-bit destination: %v %v %v", MOV, c, dst)
	}
	return arch
}

func (arch Amd64) movConstReg32(asm *Asm, c Const, dst Reg) Amd64 {
	dlo, dhi := lohi(dst)
	cval := c.Val()
	if cval == int64(int32(cval)) {
		if dhi == 0 {
			asm.Byte(0xB8 | dlo)
		} else {
			asm.Bytes(40|dhi, 0xB8|dlo)
		}
		asm.Int32(int32(cval))
	} else {
		errorf("sign-extended constant overflows 16-bit destination: %v %v %v", MOV, c, dst)
	}
	return arch
}

func (arch Amd64) movConstReg64(asm *Asm, c Const, dst Reg) Amd64 {
	dlo, dhi := lohi(dst)
	cval := c.Val()
	// 32-bit signed immediate constants, use mov
	if cval == int64(int32(cval)) {
		asm.Bytes(0x48|dhi, 0xC7, 0xC0|dlo).Int32(int32(cval))
		return arch
	}
	// 64-bit constant, must use movabs
	asm.Bytes(0x48|dhi, 0xB8|dlo).Int64(cval)
	return arch
}

// off_dst(%reg_dst) = const
func (arch Amd64) movConstMem(asm *Asm, c Const, m Mem) Amd64 {
	dregid := m.RegId()
	dkind := m.Kind()
	dlo, dhi := lohiId(dregid)
	offlen, offbit := offlen(m, dregid)
	cval := c.Val()
	switch dkind.Size() {
	case 1:
		if dhi == 0 {
			asm.Bytes(0xC6, offbit|dlo)
		} else {
			asm.Bytes(0x40|dhi, 0xC6, offbit|dlo)
		}
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi == 0 {
			asm.Bytes(0xC7, offbit|dlo)
		} else {
			asm.Bytes(0x40|dhi, 0xC7, offbit|dlo)
		}
	case 8:
		if cval == int64(int32(cval)) {
			asm.Bytes(0x48|dhi, 0xC7, offbit|dlo)
		} else {
			r := asm.RegAlloc(dkind)
			arch.movConstReg64(asm, c, r).op2(asm, MOV, r, m)
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

	switch dkind.Size() {
	case 1:
		asm.Int8(int8(cval))
	case 2:
		asm.Int16(int16(cval))
	case 4, 8:
		asm.Int32(int32(cval))
	}
	return arch
}

// ============================================================================
// movsx, movzx or mov
func (arch Amd64) cast(asm *Asm, src Arg, dst Arg) Amd64 {
	if src == dst {
		return arch
	} else if SizeOf(src) == SizeOf(dst) {
		return arch.op2(asm, MOV, src, dst)
	}
	switch dst := dst.(type) {
	case Reg:
		switch src := src.(type) {
		case Reg:
			arch.castRegReg(asm, src, dst)
		case Mem:
			arch.castMemReg(asm, src, dst)
		case Const:
			src = src.Cast(dst.Kind())
			arch.movConstReg(asm, src, dst)
		default:
			errorf("unsupported source type %T, expecting Reg, Mem or Const: %v %v %v", src, CAST, src, dst)
		}
	case Mem:
		switch src := src.(type) {
		case Reg:
			arch.castRegMem(asm, src, dst)
		case Mem:
			arch.castMemMem(asm, src, dst)
		case Const:
			src = src.Cast(dst.Kind())
			arch.op2ConstMem(asm, MOV, src, dst)
		default:
			errorf("unsupported source type %T, expecting Reg, Mem or Const: %v %v %v", src, CAST, src, dst)
		}
	case Const:
		errorf("destination cannot be a constant: %v %v %v", CAST, src, dst)
	default:
		errorf("unsupported destination type %T, expecting Reg or Mem: %v %v %v", dst, CAST, src, dst)
	}
	return arch
}

func (arch Amd64) castRegReg(asm *Asm, src Reg, dst Reg) Amd64 {
	var op uint8 = 0xB6 // movzx
	if dst.Kind().Signed() {
		op = 0xBE // movsx
	}
	dlo, dhi := lohi(dst)
	slo, shi := lohi(src)
	switch SizeOf(src) {
	case 1:
		// movzbq, movsbq
	case 2:
		op++ // movzwq, movswq
	case 4:
		if dst.Kind().Signed() {
			// movsd i.e. movslq
			asm.Bytes(0x48|dhi<<2|shi, 0x63, 0xC0|dlo<<3|slo)
			return arch
		}
		// amd64 does not have zero-extend 32bit -> 64bit
		// because operations that write into 32bit registers
		// already zero the upper 32 bits.
		// So just compile as a regular MOV
		fallthrough
	case 8:
		return arch.op2RegReg(asm, MOV, src, MakeReg(dst.RegId(), src.Kind()))
	default:
		errorf("unsupported source register size %v, expecting 1, 2, 4 or 8: %v %v %v",
			SizeOf(src), CAST, src, dst)
	}
	// for simplicity, assume Sizeof(dst) == 8
	asm.Bytes(0x48|dhi<<2|shi, 0x0F, op, 0xC0|dlo<<3|slo)
	return arch
}

func (arch Amd64) castMemReg(asm *Asm, src_m Mem, dst Reg) Amd64 {
	sregid := src_m.RegId()
	skind := src_m.Kind()

	var op uint8 = 0xB6 // movzx
	if skind.Signed() {
		op = 0xBE // movsx
	}
	dlo, dhi := lohi(dst)
	slo, shi := lohiId(sregid)
	offlen, offbit := offlen(src_m, sregid)
	// debugf("castMemReg() src = %v, dst = %v", src, dst)
	switch skind.Size() {
	case 1:
		// movzbq, movsbq
	case 2:
		op++ // movzwq, movswq
	case 4:
		if skind.Signed() {
			// sign-extend 32bit -> 64bit
			// movsd i.e. movslq
			asm.Bytes(0x48|dhi<<2|shi, 0x63, offbit|dlo<<3|slo)
			quirk24(asm, sregid)
			switch offlen {
			case 1:
				asm.Int8(int8(src_m.Offset()))
			case 4:
				asm.Int32(src_m.Offset())
			}
			return arch
		}
		// amd64 does not have zero-extend 32bit -> 64bit
		// because operations that write into 32bit registers
		// already zero the upper 32 bits.
		// So just compile as a regular MOV
		// debugf("zero-extend 32bit -> 64bit: dst = %v, src = %v", dst, m)
		fallthrough
	case 8:
		arch.op2MemReg(asm, MOV, src_m, MakeReg(dst.RegId(), skind))
		return arch
	default:
		errorf("invalid source register size %v, expecting 1, 2, 4 or 8: %v %v %v",
			skind.Size(), CAST, src_m, dst)
	}
	// for simplicity, assume Sizeof(dst) == 8
	asm.Bytes(0x48|dhi<<2|shi, 0x0F, op, offbit|dlo<<3|slo)
	quirk24(asm, sregid)
	switch offlen {
	case 1:
		asm.Int8(int8(src_m.Offset()))
	case 4:
		asm.Int32(src_m.Offset())
	}
	return arch
}

func (arch Amd64) castRegMem(asm *Asm, src Reg, dst_m Mem) Amd64 {
	dkind := dst_m.Kind()
	// assume that user code cannot use the same register
	// multiple times with different kinds
	r := MakeReg(src.RegId(), dkind)
	arch.castRegReg(asm, src, r)
	return arch.op2RegMem(asm, MOV, r, dst_m)
}

func (arch Amd64) castMemMem(asm *Asm, src Mem, dst Mem) Amd64 {
	if SizeOf(src) > SizeOf(dst) && !src.Kind().IsFloat() {
		// just read the lowest bytes from src
		arch.op2MemMem(asm, MOV,
			MakeMem(src.Offset(), src.RegId(), dst.Kind()),
			dst)
	} else {
		r := asm.RegAlloc(dst.Kind())
		arch.castMemReg(asm, src, r)
		arch.op2RegMem(asm, MOV, r, dst)
		asm.RegFree(r)
	}
	return arch
}
