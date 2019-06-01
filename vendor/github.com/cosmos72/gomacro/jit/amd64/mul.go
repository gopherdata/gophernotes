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
 * mul.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package amd64

// amd64 has very constrained 8-bit multiply
// (it can only read/write %al), so use at least 16-bit
func (arch Amd64) mul2WidenReg(asm *Asm, r Reg) Reg {
	switch r.Kind() {
	case Bool, Uint8:
		w := MakeReg(r.RegId(), Uint16)
		arch.castRegReg(asm, r, w)
		return w
	case Int8:
		w := MakeReg(r.RegId(), Int16)
		arch.castRegReg(asm, r, w)
		return w
	}
	return r
}

func (arch Amd64) mul2ConstReg(asm *Asm, c Const, dst Reg) Amd64 {
	n := c.Val()
	switch n {
	case -1:
		return arch.op1Reg(asm, NEG1, dst)
	case 0:
		return arch.zeroReg(asm, dst)
	case 1:
		return arch
	case 4, 8:
		if SizeOf(dst) == 8 {
			return arch.lea4(asm, MakeMem(0, NoRegId, dst.Kind()), dst, n, dst)
		}
	case 2, 3, 5, 9:
		if SizeOf(dst) == 8 {
			return arch.lea4(asm, MakeMem(0, dst.RegId(), dst.Kind()), dst, n-1, dst)
		}
	}
	if shift, ok := log2uint(uint64(n)); ok && n > 0 {
		return arch.op2ConstReg(asm, SHL2, ConstUint8(shift), dst)
	}
	if n == int64(int8(n)) {
		return arch.mul3RegConst8Reg(asm, dst, int8(n), dst)
	}
	// constant is too large, must allocate a register
	dst = arch.mul2WidenReg(asm, dst)
	r := asm.RegAlloc(dst.Kind())
	arch.movConstReg64(asm, MakeConst(c.Val(), dst.Kind()), r)
	arch.mul2RegReg(asm, r, dst)
	asm.RegFree(r)
	return arch
}

func (arch Amd64) mul2RegReg(asm *Asm, src Reg, dst Reg) Amd64 {
	src = arch.mul2WidenReg(asm, src)
	dst = arch.mul2WidenReg(asm, dst)
	slo, shi := lohi(src)
	dlo, dhi := lohi(dst)
	switch dst.Kind().Size() {
	case 1:
		errorf("internal error, Asm.mul2WidenReg did not widen 8-bit registers: %v %v, %v", MUL2, src, dst)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi<<2|shi == 0 {
			asm.Bytes(0x0F, 0xAF, 0xC0|dlo<<3|slo)
		} else {
			asm.Bytes(0x40|dhi<<2|shi, 0x0F, 0xAF, 0xC0|dlo<<3|slo)
		}
	case 8:
		asm.Bytes(0x48|dhi<<2|shi, 0x0F, 0xAF, 0xC0|dlo<<3|slo)
	}
	return arch
}

func (arch Amd64) mul2MemReg(asm *Asm, src_m Mem, dst Reg) Amd64 {
	sregid := src_m.RegId()
	slo, shi := lohiId(sregid)
	dlo, dhi := lohi(dst)
	offlen, offbit := offlen(src_m, sregid)
	switch dst.Kind().Size() {
	case 1:
		// amd64 has very constrained 8-bit multiply
		// (it can only read/write %al), so copy 8-bit memory
		// to a register and use widening multiplication
		r := asm.RegAlloc(src_m.Kind())
		arch.load(asm, src_m, r).mul2RegReg(asm, r, dst)
		asm.RegFree(r)
		return arch
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi<<2|shi == 0 {
			asm.Bytes(0x0F, 0xAF, offbit|dlo<<3|slo)
		} else {
			asm.Bytes(0x40|dhi<<2|shi, 0x0F, 0xAF, offbit|dlo<<3|slo)
		}
	case 8:
		asm.Bytes(0x48|dhi<<2|shi, 0x0F, 0xAF, offbit|dlo<<3|slo)
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

func (arch Amd64) mul2ConstMem(asm *Asm, c Const, m Mem) Amd64 {
	switch c.Val() {
	case -1:
		return arch.op1Mem(asm, NEG1, m)
	case 0:
		return arch.zeroMem(asm, m)
	case 1:
		return arch
	default:
		r := asm.RegAlloc(m.Kind())
		arch.load(asm, m, r).mul2ConstReg(asm, c, r).store(asm, r, m)
		asm.RegFree(r)
		return arch
	}
}

func (arch Amd64) mul2RegMem(asm *Asm, src Reg, dst_m Mem) Amd64 {
	// must allocate a register
	r := asm.RegAlloc(dst_m.Kind())
	arch.load(asm, dst_m, r).mul2RegReg(asm, src, r).store(asm, r, dst_m)
	asm.RegFree(r)
	return arch
}

func (arch Amd64) mul2MemMem(asm *Asm, src_m Mem, dst_m Mem) Amd64 {
	// must allocate a register
	r := asm.RegAlloc(dst_m.Kind())
	arch.load(asm, dst_m, r).mul2MemReg(asm, src_m, r).store(asm, r, dst_m)
	asm.RegFree(r)
	return arch
}

// =============== 3-argument MUL3 ==================

func (arch Amd64) mul3(asm *Asm, a Arg, b Arg, dst Arg) Amd64 {
	if a.Const() && !b.Const() {
		a, b = b, a
	}
	if a == dst {
		return arch.op2(asm, MUL2, b, dst)
	} else if b == dst {
		return arch.op2(asm, MUL2, a, dst)
	}
	rdst, rokdst := dst.(Reg)
	if !a.Const() && b.Const() {
		bval := b.(Const).Val()
		if bval == int64(int8(bval)) {
			// use amd64 3-argument multiplication
			if !rokdst {
				rdst = asm.RegAlloc(dst.Kind())
			}
			switch a := a.(type) {
			case Reg:
				arch.mul3RegConst8Reg(asm, a, int8(bval), rdst)
			case Mem:
				if a.Kind().Size() == 1 {
					// to use 16-bit multiplication
					// we must widen Mem, so we need a register
					widekind := Uint16
					if a.Kind().Signed() {
						widekind = Int16
					}
					ra := asm.RegAlloc(widekind)
					arch.castMemReg(asm, a, ra)
					arch.mul3RegConst8Reg(asm, ra, int8(bval), rdst)
					asm.RegFree(ra)
				} else {
					arch.mul3MemConst8Reg(asm, a, int8(bval), rdst)
				}
			default:
				errorf("unknown argument type %T, expecting Const, Reg or Mem: %v %v, %v, %v", a, MUL3, a, b, dst)
			}
			if !rokdst {
				arch.store(asm, rdst, dst.(Mem))
				asm.RegFree(rdst)
			}
			return arch
		}
	}
	if rokdst && rdst.RegId() != b.RegId() {
		return arch.mov(asm, a, dst).op2(asm, MUL2, b, dst)
	}
	r := asm.RegAlloc(dst.Kind())
	arch.mov(asm, a, r).op2(asm, MUL2, b, r).mov(asm, r, dst)
	asm.RegFree(r)
	return arch
}

func (arch Amd64) mul3RegConst8Reg(asm *Asm, src Reg, cval int8, dst Reg) Amd64 {
	src = arch.mul2WidenReg(asm, src)
	dst = arch.mul2WidenReg(asm, dst)
	slo, shi := lohi(src)
	dlo, dhi := lohi(dst)
	switch dst.Kind().Size() {
	case 1:
		errorf("internal error, Asm.mul2WidenReg did not widen 8-bit registers: %v %v, %v, %v", MUL2, src, cval, dst)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi<<2|shi == 0 {
			asm.Bytes(0x6B, 0xC0|dlo<<3|slo, uint8(cval))
		} else {
			asm.Bytes(0x40|dhi<<2|shi, 0x6B, 0xC0|dlo<<3|slo, uint8(cval))
		}
	case 8:
		asm.Bytes(0x48|dhi<<2|shi, 0x6B, 0xC0|dlo<<3|slo, uint8(cval))
	}
	return arch
}

func (arch Amd64) mul3MemConst8Reg(asm *Asm, src_m Mem, cval int8, dst Reg) Amd64 {
	sregid := src_m.RegId()
	slo, shi := lohiId(sregid)
	dlo, dhi := lohi(dst)
	offlen, offbit := offlen(src_m, sregid)
	switch dst.Kind().Size() {
	case 1:
		errorf("internal error, missing call to Asm.mul2WidenReg to widen 8-bit registers: %v %v, %v, %v", MUL2, src_m, cval, dst)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi<<2|shi == 0 {
			asm.Bytes(0x6B, offbit|dlo<<3|slo)
		} else {
			asm.Bytes(0x40|dhi<<2|shi, 0x6B, offbit|dlo<<3|slo)
		}
	case 8:
		asm.Bytes(0x48|dhi<<2|shi, 0x6B, offbit|dlo<<3|slo)
	}
	quirk24(asm, sregid)
	switch offlen {
	case 1:
		asm.Int8(int8(src_m.Offset()))
	case 4:
		asm.Int32(src_m.Offset())
	}
	asm.Int8(cval)
	return arch
}
