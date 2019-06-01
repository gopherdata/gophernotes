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

func isMem(a Arg) bool {
	_, ok := a.(Mem)
	return ok
}

func isdivreg(a Arg) bool {
	id := a.RegId()
	return id == RAX || id == RDX
}

func checkdivrem(op Op3, a Arg, b Arg, dst Arg) {
	// amd64 div is hardcoded to use RAX and RDX as both wide dividend and wide result

	if isdivreg(b) {
		errorf("unimplemented %v with divisor that uses %v: %v %v, %v, %v", op, b.RegId(), op, a, b, dst)
	}
	// allow a == RAX/RDX only if we are overwriting it anyway because a == dst
	if isdivreg(a) && (isMem(a) || a != dst) {
		errorf("unimplemented %v with %T dividend that uses %v: %v %v, %v, %v", op, a, a.RegId(), op, a, b, dst)
	}
	if isMem(dst) && isdivreg(dst) {
		errorf("unimplemented %v with Mem result that uses %v: %v %v, %v, %v", op, dst.RegId(), op, a, b, dst)
	}
}

func (arch Amd64) divrem(asm *Asm, op Op3, a Arg, b Arg, dst Arg) Amd64 {

	checkdivrem(op, a, b, dst)

	if asm.RegIncUse(RAX) != 1 {
		errorf("unimplemented %v with RAX already in use: %v %v, %v, %v", op, op, a, b, dst)
	}
	defer asm.RegDecUse(RAX)
	if asm.RegIncUse(RDX) != 1 {
		errorf("unimplemented %v with RDX already in use: %v %v, %v, %v", op, op, a, b, dst)
	}
	defer asm.RegDecUse(RDX)

	// amd64 div cannot encode constants, and accepts Mem only as divisor
	var ra, rb Reg

	switch a := a.(type) {
	case Reg:
		ra = a
	case Const, Mem:
		ra = MakeReg(RAX, a.Kind())
		arch.mov(asm, a, ra)
	}

	switch b := b.(type) {
	case Reg:
		rb = b
	case Const:
		rb = asm.RegAlloc(b.Kind())
		arch.movConstReg(asm, b, rb)
		defer asm.RegFree(rb)
	case Mem:
		if dst.Kind().Signed() {
			return arch.sdivremMem(asm, op, ra, b, dst)
		} else {
			return arch.udivremMem(asm, op, ra, b, dst)
		}
	}

	if dst.Kind().Signed() {
		return arch.sdivremReg(asm, op, ra, rb, dst)
	} else {
		return arch.udivremReg(asm, op, ra, rb, dst)
	}
}

// signed division
func (arch Amd64) sdivremReg(asm *Asm, op Op3, a Reg, b Reg, dst Arg) Amd64 {
	kind := dst.Kind()
	size := kind.Size()
	if size <= 4 {
		kind = Int32
		size = 4
	}
	arch.cast(asm, a, MakeReg(RAX, kind))
	arch.cast(asm, b, MakeReg(b.RegId(), kind))

	blo, bhi := lohi(b)
	if size == 4 {
		// 0x99 == cltd, i.e. get sign bit from %eax and fill %edx with it
		if bhi == 0 {
			asm.Bytes(0x99, 0xF7, 0xF8|blo)
		} else {
			asm.Bytes(0x99, 0x40|bhi, 0xF7, 0xF8|blo)
		}
	} else {
		// 0x48 0x99 == cqto, i.e. get sign bit from %rax and fill %rdx with it
		asm.Bytes(0x48, 0x99, 0x48|bhi, 0xF7, 0xF8|blo)
	}
	var rid = RAX
	if op == REM3 {
		rid = RDX
	}
	return arch.mov(asm, MakeReg(rid, dst.Kind()), dst)
}

// signed division
func (arch Amd64) sdivremMem(asm *Asm, op Op3, a Reg, b Mem, dst Arg) Amd64 {
	kind := dst.Kind()
	size := kind.Size()
	if size < 4 {
		// do what gcc and clang do: copy b to Reg and use sdivremReg()
		rb := asm.RegAlloc(Int32)
		arch.castMemReg(asm, b, rb)
		defer asm.RegFree(rb)
		return arch.sdivremReg(asm, op, a, rb, dst)
	}
	arch.castRegReg(asm, a, MakeReg(RAX, kind))

	bregid := b.RegId()
	mofflen, offbit := offlen(b, bregid)
	blo, bhi := lohiId(bregid)
	if size == 4 {
		// 0x99 == cltd
		if bhi == 0 {
			asm.Bytes(0x99, 0xF7, offbit|0x38|blo)
		} else {
			asm.Bytes(0x99, 0x40|bhi, 0xF7, offbit|0x38|blo)
		}
	} else {
		// 0x48 0x99 == cqto
		asm.Bytes(0x48, 0x99, 0x48|bhi, 0xF7, offbit|0x38|blo)
	}
	quirk24(asm, bregid)
	switch mofflen {
	case 1:
		asm.Int8(int8(b.Offset()))
	case 4:
		asm.Int32(b.Offset())
	}
	var rid = RAX
	if op == REM3 {
		rid = RDX
	}
	return arch.mov(asm, MakeReg(rid, dst.Kind()), dst)
}

// unsigned division
func (arch Amd64) udivremReg(asm *Asm, op Op3, a Reg, b Reg, dst Arg) Amd64 {
	errorf("unimplemented: unsigned division: %v %v, %v, %v", op, a, b, dst)
	return arch
}

// unsigned division
func (arch Amd64) udivremMem(asm *Asm, op Op3, a Reg, b Mem, dst Arg) Amd64 {
	errorf("unimplemented: unsigned division: %v %v, %v, %v", op, a, b, dst)
	return arch
}
