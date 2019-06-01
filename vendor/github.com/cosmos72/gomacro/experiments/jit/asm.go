/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * asm.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"unsafe"
)

const (
	S       = uint32(unsafe.Sizeof(uint64(0)))
	VERBOSE = false
)

func (s *Save) Init(start, end uint16) {
	s.start, s.idx, s.end = start, start, end
}

func (asm *Asm) Init() *Asm {
	return asm.Init2(0, 0)
}

func (asm *Asm) Init2(saveStart, saveEnd uint16) *Asm {
	asm.code = asm.code[:0]
	asm.hwRegs.InitLive()
	asm.regs = make(map[Reg]hwRegCounter)
	asm.regNext = RegHi + 1
	asm.save.Init(saveStart, saveEnd)
	return asm.prologue()
}

func (asm *Asm) Bytes(bytes ...uint8) *Asm {
	asm.code = append(asm.code, bytes...)
	return asm
}

func (asm *Asm) Uint16(val uint16) *Asm {
	asm.code = append(asm.code, uint8(val), uint8(val>>8))
	return asm
}

func (asm *Asm) Uint32(val uint32) *Asm {
	asm.code = append(asm.code, uint8(val), uint8(val>>8), uint8(val>>16), uint8(val>>24))
	return asm
}

func (asm *Asm) Uint64(val uint64) *Asm {
	asm.code = append(asm.code, uint8(val), uint8(val>>8), uint8(val>>16), uint8(val>>24), uint8(val>>32), uint8(val>>40), uint8(val>>48), uint8(val>>56))
	return asm
}

func (asm *Asm) Int16(val int16) *Asm {
	return asm.Uint16(uint16(val))
}

func (asm *Asm) Int32(val int32) *Asm {
	return asm.Uint32(uint32(val))
}

func (asm *Asm) Int64(val int64) *Asm {
	return asm.Uint64(uint64(val))
}

func (asm *Asm) Idx(a *Var) *Asm {
	return asm.Uint32(uint32(a.idx) * S)
}

func (asm *Asm) reg(g Reg) hwReg {
	return asm.regs[g].hwReg
}

func (asm *Asm) pushRegs(rs *hwRegs) *hwRegs {
	var ret hwRegs
	v := &Var{}
	for r := rLo; r <= rHi; r++ {
		if !rs.Contains(r) || !asm.hwRegs.Contains(r) {
			continue
		}
		if asm.save.idx >= asm.save.end {
			errorf("save area is full, cannot push registers")
		}
		v.idx = asm.save.idx
		asm.storeReg(v, r)
		asm.save.idx++
		ret.Set(r)
	}
	return &ret
}

func (asm *Asm) popRegs(rs *hwRegs) {
	v := &Var{}
	for r := rHi; r >= rLo; r-- {
		if !rs.Contains(r) {
			continue
		}
		if asm.save.idx <= asm.save.start {
			errorf("save area is empty, cannot pop registers")
		}
		asm.save.idx--
		v.idx = asm.save.idx
		asm.load(r, v)
	}
}

// allocate a jit-reserved register
func (asm *Asm) alloc() Reg {
	z := asm.regNext
	asm.regNext++
	asm.Alloc(z)
	return z
}

func (asm *Asm) Alloc(z Reg) *Asm {
	pair := asm.regs[z]
	if !pair.Valid() {
		pair.hwReg = asm.hwRegs.Alloc()
	}
	pair.count++
	asm.regs[z] = pair
	return asm
}

// combined Alloc + Load
func (asm *Asm) AllocLoad(z Reg, a Arg) *Asm {
	return asm.Alloc(z).Load(z, a)
}

func (asm *Asm) Free(z Reg) *Asm {
	pair, ok := asm.regs[z]
	if !ok {
		return asm
	}
	pair.count--
	if pair.count == 0 {
		asm.hwRegs.Free(pair.hwReg)
		delete(asm.regs, z)
	} else {
		asm.regs[z] = pair
	}
	return asm
}

// combined Store + Free
func (asm *Asm) StoreFree(z *Var, g Reg) *Asm {
	return asm.Store(z, g).Free(g)
}

func (asm *Asm) hwAlloc(a Arg) (r hwReg, allocated bool) {
	r = a.reg(asm)
	if r != noReg {
		return r, false
	}
	r = asm.hwRegs.Alloc()
	asm.load(r, a)
	return r, true
}

func (asm *Asm) hwAllocConst(val int64) hwReg {
	r := asm.hwRegs.Alloc()
	asm.loadConst(r, val)
	return r
}

func (asm *Asm) hwFree(r hwReg, allocated bool) *Asm {
	if r.Valid() && allocated {
		asm.hwRegs.Free(r)
	}
	return asm
}
