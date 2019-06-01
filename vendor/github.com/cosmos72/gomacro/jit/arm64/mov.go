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
 *  Created on Feb 02, 2019
 *      Author Massimiliano Ghilardi
 */

package arm64

// ============================================================================
type loadstore uint32

const (
	load  loadstore = 0x39400000
	store loadstore = 0x39000000
)

func (arch Arm64) load(asm *Asm, src Mem, dst Reg) Arm64 {
	return arch.loadstore(asm, load, src, dst)
}

func (arch Arm64) store(asm *Asm, src Reg, dst Mem) Arm64 {
	return arch.loadstore(asm, store, dst, src)
}

func (arch Arm64) mov(asm *Asm, src Arg, dst Arg) Arm64 {
	assert(SizeOf(src) == SizeOf(dst))

	if dst.Const() {
		errorf("destination cannot be a constant: %v %v, %v", MOV, src, dst)
	}
	if src == dst {
		return arch
	}

	switch dst := dst.(type) {
	case Reg:
		switch src := src.(type) {
		case Const:
			arch.movConstReg(asm, src, dst)
		case Reg:
			if src.RegId() != dst.RegId() {
				arch.movRegReg(asm, src, dst)
			}
		case Mem:
			arch.load(asm, src, dst)
		default:
			errorf("unknown source type %T, expecting Const, Reg or Mem: %v %v, %v", src, MOV, src, dst)
		}
	case Mem:
		switch src := src.(type) {
		case Const:
			arch.movConstMem(asm, src, dst)
		case Reg:
			arch.store(asm, src, dst)
		case Mem:
			arch.movMemMem(asm, src, dst)
		default:
			errorf("unknown source type %T, expecting Const, Reg or Mem: %v %v, %v", src, MOV, src, dst)
		}
	default:
		errorf("unknown destination type %T, expecting Reg or Mem: %v %v, %v", dst, MOV, src, dst)
	}
	return arch
}

func (arch Arm64) movRegReg(asm *Asm, src Reg, dst Reg) Arm64 {
	// arm64 implements "mov src,dst" as "orr xzr,src,dst"
	asm.Uint32(kbit(dst) | 0x2A0003E0 | valOrX31(src.RegId(), true)<<16 | val(dst))
	return arch
}

func (arch Arm64) movConstReg(asm *Asm, c Const, dst Reg) Arm64 {
	cval := c.Val()
	xzr := MakeReg(XZR, dst.Kind())
	var immcval uint32
	var movk bool
	if cval >= 0 && cval < 0x10000 {
		immcval = 0x40<<19 | uint32(cval)
	} else if cval < 0 && cval >= -0x10000 {
		immcval = uint32(^cval)
	} else if arch.tryOp3RegConstReg(asm, OR3, xzr, uint64(cval), dst) {
		return arch
	} else if arch.tryOp3RegConstReg(asm, OR3, xzr, uint64(uint32(cval)), dst) {
		if dst.Kind().Size() == 8 {
			arch.movk(asm, uint16(cval>>32), 32, dst)
			arch.movk(asm, uint16(cval>>48), 48, dst)
		}
		return arch
	} else {
		immcval = 0x40<<19 | uint32(cval&0xFFFF)
		movk = true
	}
	asm.Uint32(kbit(dst) | 0x12800000 | immcval<<5 | val(dst))
	if movk {
		arch.movk(asm, uint16(cval>>16), 16, dst)
		if dst.Kind().Size() == 8 {
			arch.movk(asm, uint16(cval>>32), 32, dst)
			arch.movk(asm, uint16(cval>>48), 48, dst)
		}
	}
	return arch
}

// set some bits of dst, preserving others
func (arch Arm64) movk(asm *Asm, cval uint16, shift uint8, dst Reg) Arm64 {
	if cval != 0 {
		asm.Uint32(kbit(dst) | 0xF2800000 | uint32(shift)<<17 | uint32(cval)<<5 | val(dst))
	}
	return arch
}

func (arch Arm64) loadstore(asm *Asm, op loadstore, m Mem, r Reg) Arm64 {
	assert(SizeOf(m) == SizeOf(r))
	off := m.Offset()
	var sizebit uint32
	mrid := m.RegId()
	rid := r.RegId()

	switch m.Kind().Size() {
	case 1:
		sizebit = 0
		if off >= 0 && off <= 4095 {
			asm.Uint32(sizebit | uint32(op) | uint32(off)<<10 | valOrX31(mrid, true)<<5 | valOrX31(rid, true))
			return arch
		}
	case 2:
		sizebit = 0x4 << 28
		if off >= 0 && off <= 8190 && off%2 == 0 {
			asm.Uint32(sizebit | uint32(op) | uint32(off)<<9 | valOrX31(mrid, true)<<5 | valOrX31(rid, true))
			return arch
		}
	case 4:
		sizebit = 0x8 << 28
		if off >= 0 && off <= 16380 && off%4 == 0 {
			asm.Uint32(sizebit | uint32(op) | uint32(off)<<8 | valOrX31(mrid, true)<<5 | valOrX31(rid, true))
			return arch
		}
	case 8:
		sizebit = 0xC << 28
		if off >= 0 && off <= 32760 && off%8 == 0 {
			asm.Uint32(sizebit | uint32(op) | uint32(off)<<7 | valOrX31(mrid, true)<<5 | valOrX31(rid, true))
			return arch
		}
	}
	// load offset in a register. we could also try "ldur" or "stur"...
	tmp := asm.RegAlloc(Uint64)
	arch.movConstReg(asm, ConstInt64(int64(off)), tmp)

	asm.Uint32(sizebit | uint32(op^0x1206800) | val(tmp)<<16 | valOrX31(mrid, true)<<5 | val(r))

	asm.RegFree(tmp)
	return arch
}

func (arch Arm64) movConstMem(asm *Asm, c Const, dst Mem) Arm64 {
	if c.Val() == 0 {
		return arch.zeroMem(asm, dst)
	}
	r := asm.RegAlloc(dst.Kind())
	arch.movConstReg(asm, c, r).store(asm, r, dst)
	asm.RegFree(r)
	return arch
}

func (arch Arm64) movMemMem(asm *Asm, src Mem, dst Mem) Arm64 {
	r := asm.RegAlloc(src.Kind())
	arch.load(asm, src, r).store(asm, r, dst)
	asm.RegFree(r)
	return arch
}

// ============================================================================
func (arch Arm64) Cast(asm *Asm, src Arg, dst Arg) {
	arch.Cast(asm, src, dst)
}

func (arch Arm64) cast(asm *Asm, src Arg, dst Arg) Arm64 {
	if src == dst {
		return arch
	} else if SizeOf(src) == SizeOf(dst) {
		return arch.mov(asm, src, dst)
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
			errorf("unsupported source type %T, expecting Const, Reg or Mem: %v %v %v", src, CAST, src, dst)
		}
	case Mem:
		switch src := src.(type) {
		case Reg:
			arch.castRegMem(asm, src, dst)
		case Mem:
			arch.castMemMem(asm, src, dst)
		case Const:
			src = src.Cast(dst.Kind())
			arch.movConstMem(asm, src, dst)
		default:
			errorf("unsupported source type %T, expecting Const, Reg or Mem: %v %v %v", src, CAST, src, dst)
		}
	case Const:
		errorf("destination cannot be a constant: %v %v %v", CAST, src, dst)
	default:
		errorf("unsupported destination type %T, expecting Reg or Mem: %v %v %v", dst, CAST, src, dst)
	}
	return arch
}

func (arch Arm64) castMemMem(asm *Asm, src Mem, dst Mem) Arm64 {
	r1 := asm.RegAlloc(src.Kind())
	r2 := MakeReg(r1.RegId(), dst.Kind())
	arch.load(asm, src, r1).castRegReg(asm, r1, r2).store(asm, r2, dst)
	asm.RegFree(r1)
	return arch
}

func (arch Arm64) castMemReg(asm *Asm, src Mem, dst Reg) Arm64 {
	r := MakeReg(dst.RegId(), src.Kind())
	return arch.load(asm, src, r).castRegReg(asm, r, dst)
}

func (arch Arm64) castRegMem(asm *Asm, src Reg, dst Mem) Arm64 {
	r := MakeReg(src.RegId(), dst.Kind())
	if SizeOf(src) < SizeOf(dst) {
		// extend src. we can safely overwrite its high bits: they are junk
		return arch.castRegReg(asm, src, r).store(asm, r, dst)
	} else {
		// just ignore src high bits
		return arch.store(asm, r, dst)
	}
}

func (arch Arm64) castRegReg(asm *Asm, src Reg, dst Reg) Arm64 {
	skind := src.Kind()
	dkind := dst.Kind()
	ssize := skind.Size()
	dsize := dkind.Size()
	if ssize >= dsize {
		// truncate. easy, just ignore src high bits
		return arch.mov(asm, MakeReg(src.RegId(), dst.Kind()), dst)
	} else if skind.Signed() {
		// sign-extend. use one of:
		// "sxtb	src, dst"
		// "sxth	src, dst"
		// "sxtw	src, dst"
		kbit := uint32(dsize&8) * 0x10080000
		op := 0x13000C00 | uint32(ssize*2-1)<<12
		asm.Uint32(kbit | op | val(src)<<5 | val(dst))
		return arch
	} else {
		// zero-extend
		if ssize == 4 {
			// zero-extend 32 bit -> 64 bit: use
			// "mov dst, src"
			// must be kept even if src == dst to zero high bits,
			// so use Asm.movRegReg() instead of too smart Asm.Mov()
			return arch.movRegReg(asm, src, MakeReg(dst.RegId(), skind))
		}
		// zero-extend, src is 8 bit or 16 bit. use one of:
		// "and dst, src, #0xff"
		// "and dst, src, #0xffff"
		if dsize <= 4 {
			dkind = Uint32
		}
		r := MakeReg(src.RegId(), dkind)
		c := MakeConst(int64(0xffff)>>(16-ssize*8), dkind)
		return arch.op3RegConstReg(asm, AND3, r, c, dst)
	}
}
