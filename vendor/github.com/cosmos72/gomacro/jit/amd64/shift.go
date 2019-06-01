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
 * shift.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package amd64

// %reg_dst SHIFT= const
func (arch Amd64) shiftConstReg(asm *Asm, op Op2, c Const, dst Reg) Amd64 {
	n := c.Val()
	if n == 0 {
		// shift by 0 should be already optimized away,
		// but optimizing just in case does not hurt
		return arch
	}
	siz := SizeOf(dst)
	if n >= 8*int64(siz) {
		return arch.zeroReg(asm, dst)
	}
	var nbit uint8
	if n == 1 {
		nbit = 0x10
	}
	op_ := op2val(op)
	dlo, dhi := lohi(dst)
	switch siz {
	case 1:
		if dst.RegId() >= RSP {
			asm.Byte(0x40 | dhi)
		}
		asm.Bytes(0xC0|nbit, op_|dlo)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi != 0 {
			asm.Byte(0x40 | dhi)
		}
		asm.Bytes(0xC1|nbit, op_|dlo)
	case 8:
		asm.Bytes(0x48|dhi, 0xC1|nbit, op_|dlo)
	}
	if n != 1 {
		asm.Uint8(uint8(n))
	}
	return arch
}

// off_dst(%reg_dst) SHIFT= const
func (arch Amd64) shiftConstMem(asm *Asm, op Op2, c Const, m Mem) Amd64 {
	n := c.Val()
	assert(n > 0) // shift by 0 is optimized away
	size := SizeOf(m)
	if n >= 8*int64(size) {
		if m.Kind().Signed() {
			n = 8*int64(size) - 1
		} else {
			return arch.zeroMem(asm, m)
		}
	}
	dregid := m.RegId()
	dlo, dhi := lohiId(dregid)
	offlen, offbit := offlen(m, dregid)
	op_ := op2val(op) &^ 0xC0

	var nbit uint8
	if n == 1 {
		nbit = 0x10
	}
	switch size {
	case 1:
		if dregid >= RSP {
			asm.Byte(0x40 | dhi)
		}
		asm.Bytes(0xC0|nbit, offbit|op_|dlo)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi != 0 {
			asm.Byte(0x40 | dhi)
		}
		asm.Bytes(0xC1|nbit, offbit|op_|dlo)
	case 8:
		asm.Bytes(0x48|dhi, 0xC1|nbit, offbit|op_|dlo)
	}
	quirk24(asm, dregid)
	switch offlen {
	case 1:
		asm.Int8(int8(m.Offset()))
	case 4:
		asm.Int32(m.Offset())
	}
	if n != 1 {
		asm.Uint8(uint8(n))
	}
	return arch
}

// %reg_dst SHIFT= %reg_src
func (arch Amd64) shiftRegReg(asm *Asm, op Op2, src Reg, dst Reg) Amd64 {
	if dst.RegId() == RCX {
		errorf("unimplemented shift RCX by Reg: %v %v %v", op, src, dst)
	}
	if src.RegId() != RCX {
		arch.op2RegReg(asm, MOV, src, MakeReg(RCX, src.Kind()))
	}
	op_ := op2val(op)
	siz := SizeOf(dst)
	dlo, dhi := lohi(dst)

	switch siz {
	case 1:
		if dst.RegId() >= RSP {
			asm.Byte(0x40 | dhi)
		}
		asm.Bytes(0xD2, op_|dlo)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi != 0 {
			asm.Byte(0x40 | dhi)
		}
		asm.Bytes(0xD3, op_|dlo)
	case 8:
		asm.Bytes(0x48|dhi, 0xD3, op_|dlo)
	}
	return arch
}

// off_dst(%reg_dst) SHIFT= %reg_src
func (arch Amd64) shiftRegMem(asm *Asm, op Op2, src Reg, dst_m Mem) Amd64 {
	if dst_m.RegId() == RCX {
		errorf("unimplemented shift Mem[RCX] by Reg: %v %v %v", op, src, dst_m)
	}
	if src.RegId() != RCX {
		arch.op2RegReg(asm, MOV, src, MakeReg(RCX, src.Kind()))
	}
	siz := SizeOf(dst_m)
	dregid := dst_m.RegId()
	dlo, dhi := lohiId(dregid)
	offlen, offbit := offlen(dst_m, dregid)
	op_ := op2val(op) &^ 0xC0

	switch siz {
	case 1:
		if dregid >= RSP {
			asm.Byte(0x40 | dhi)
		}
		asm.Bytes(0xD2, offbit|op_|dlo)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi != 0 {
			asm.Byte(0x40 | dhi)
		}
		asm.Bytes(0xD3, offbit|op_|dlo)
	case 8:
		asm.Bytes(0x48|dhi, 0xD3, offbit|op_|dlo)
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

// %reg_dst SHIFT= off_src(%reg_src)
func (arch Amd64) shiftMemReg(asm *Asm, op Op2, src_m Mem, dst Reg) Amd64 {
	if dst.RegId() == RCX {
		errorf("unimplemented shift RCX by Mem: %v %v %v", op, src_m, dst)
	}
	r := MakeReg(RCX, src_m.Kind())
	arch.op2MemReg(asm, MOV, src_m, r)
	arch.shiftRegReg(asm, op, r, dst)
	return arch
}

// off_dst(%reg_dst) SHIFT= off_src(%reg_src)
func (arch Amd64) shiftMemMem(asm *Asm, op Op2, src_m Mem, dst_m Mem) Amd64 {
	if dst_m.RegId() == RCX {
		errorf("unimplemented shift Mem[RCX] by Mem: %v %v %v", op, src_m, dst_m)
	} else if src_m.RegId() == RCX {
		errorf("unimplemented shift Mem by Mem[RCX]: %v %v %v", op, src_m, dst_m)
	}
	r := MakeReg(RCX, src_m.Kind())
	arch.op2MemReg(asm, MOV, src_m, r)
	arch.shiftRegMem(asm, op, r, dst_m)
	return arch
}
