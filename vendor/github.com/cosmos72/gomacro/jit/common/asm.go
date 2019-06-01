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

package common

type Asm struct {
	code          MachineCode
	softRegs      SoftRegs
	save          Save
	regIds        RegIds
	initialRegIds RegIds
	arch          Arch
	pool          *MemPool
	cache         Cache
	// map from indexes in code[] of 32-bit relative jumps offsets to be filled,
	// to absolute jump address
	jump map[int]uintptr
}

func New(id ArchId) *Asm {
	var asm Asm
	return asm.InitArchId(id)
}

func NewArch(arch Arch) *Asm {
	var asm Asm
	return asm.InitArch(arch)
}

func (asm *Asm) ArchId() ArchId {
	if asm.arch == nil {
		return NOARCH
	}
	return asm.arch.Id()
}

func (asm *Asm) Arch() Arch {
	return asm.arch
}

func (asm *Asm) InitArchId(archId ArchId) *Asm {
	return asm.InitArch2(Archs[archId], 0, 0)
}

func (asm *Asm) InitArchId2(archId ArchId, saveStart SaveSlot, saveEnd SaveSlot) *Asm {
	return asm.InitArch2(Archs[archId], saveStart, saveEnd)
}

func (asm *Asm) InitArch(arch Arch) *Asm {
	return asm.InitArch2(arch, 0, 0)
}

func (asm *Asm) InitArch2(arch Arch, saveStart SaveSlot, saveEnd SaveSlot) *Asm {
	if arch == nil {
		errorf("unknown arch")
	}
	id := arch.Id()
	if Archs[id] == nil {
		Archs[id] = arch
	}
	config := arch.RegIdConfig()
	asm.arch = arch
	asm.code = MachineCode{ArchId: id}
	asm.softRegs = make(SoftRegs)
	s := asm.save
	s.start, s.next, s.end = saveStart, saveStart, saveEnd
	s.reg = Reg{config.RSP, Uint64}
	s.bitmap = make([]bool, saveEnd-saveStart)
	asm.regIds.inuse = make(map[RegId]uint32)
	asm.regIds.first = config.RAllocFirst
	asm.regIds.curr = config.RAllocFirst
	asm.regIds.rlo = config.RLo
	asm.regIds.rhi = config.RHi
	asm.pool = nil
	asm.cache = nil
	asm.jump = nil
	arch.Init(asm, saveStart, saveEnd)
	asm.initialRegIds.Copy(&asm.regIds)
	arch.Prologue(asm)
	return asm
}

func (asm *Asm) Code() MachineCode {
	return asm.code
}

func (asm *Asm) ClearCode() *Asm {
	asm.code.Bytes = nil
	return asm
}

// forget all allocated registers
func (asm *Asm) ClearRegs() *Asm {
	asm.regIds.Copy(&asm.initialRegIds)
	return asm
}

// mark the last assembled 32 bits
// as a relative jump destination to be set to 'dst'
func (asm *Asm) AddJump(dst uintptr) *Asm {
	if asm.jump == nil {
		asm.jump = make(map[int]uintptr)
	}
	asm.jump[len(asm.code.Bytes)-4] = dst
	return asm
}

func (asm *Asm) Byte(b byte) *Asm {
	asm.code.Bytes = append(asm.code.Bytes, b)
	return asm
}

func (asm *Asm) Bytes(bytes ...byte) *Asm {
	asm.code.Bytes = append(asm.code.Bytes, bytes...)
	return asm
}

func (asm *Asm) Uint8(val uint8) *Asm {
	asm.code.Bytes = append(asm.code.Bytes, val)
	return asm
}

func (asm *Asm) Uint16(val uint16) *Asm {
	asm.code.Bytes = append(asm.code.Bytes, uint8(val), uint8(val>>8))
	return asm
}

func (asm *Asm) Uint32(val uint32) *Asm {
	asm.code.Bytes = append(asm.code.Bytes, uint8(val), uint8(val>>8), uint8(val>>16), uint8(val>>24))
	return asm
}

func (asm *Asm) Uint64(val uint64) *Asm {
	asm.code.Bytes = append(asm.code.Bytes, uint8(val), uint8(val>>8), uint8(val>>16), uint8(val>>24), uint8(val>>32), uint8(val>>40), uint8(val>>48), uint8(val>>56))
	return asm
}

func (asm *Asm) Int8(val int8) *Asm {
	return asm.Uint8(uint8(val))
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

// ===================================

// convert AsmCode to Arg
func (asm *Asm) Arg(x AsmCode) Arg {
	switch x := x.(type) {
	case SoftReg:
		return x.Arg(asm)
	case Arg:
		return x
	default:
		errorf("unknown argument type %T, expecting Const, Reg, Mem or SoftRegId", x)
		return nil
	}
}

// allocate a SoftRegId
func (asm *Asm) Alloc(s SoftReg) Arg {
	var rm regIdOrMem
	kind := s.Kind()
	if r := asm.TryRegAlloc(kind); r.Valid() {
		rm.regId = r.RegId()
		asm.softRegs[s.Id()] = rm
		return r
	}
	idx := asm.save.Alloc()
	if idx == InvalidSlot {
		errorf("no free register, and save area is full. Cannot allocate soft register %v", s)
	}
	rm.off = int32(idx) * 8
	rm.regId = asm.save.reg.id
	rm.ismem = true
	return MakeMem(rm.off, rm.regId, kind)
}

func (asm *Asm) Free(s SoftReg) {
	id := s.Id()
	rm, ok := asm.softRegs[id]
	if !ok {
		errorf("cannot free unallocated soft register %v", s)
	}
	if rm.ismem {
		asm.save.Free(SaveSlot(rm.off / 8))
	} else {
		asm.RegFree(MakeReg(rm.regId, s.Kind()))
	}
	delete(asm.softRegs, id)
}
