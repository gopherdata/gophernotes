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
 * opmisc.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"fmt"
)

type SaveSlot uint16

const (
	InvalidSlot = ^SaveSlot(0)
)

// implement AsmCode interface
func (*SaveSlot) asmcode() {
}

type Op1Misc uint8
type Op2Misc uint8

const (
	ALLOC Op1Misc = 0x01 // allocate soft register
	FREE  Op1Misc = 0x02 // free soft register

	PUSH Op2Misc = 0x20
	POP  Op2Misc = 0x21
)

var op1MiscName = map[Op1Misc]string{
	ALLOC: "ALLOC",
	FREE:  "FREE",
}

var op2MiscName = map[Op2Misc]string{
	PUSH: "PUSH",
	POP:  "POP",
}

func (op Op1Misc) String() string {
	s, ok := op1MiscName[op]
	if !ok {
		s = fmt.Sprintf("Op1Misc(%d)", int(op))
	}
	return s
}

func (op Op2Misc) String() string {
	s, ok := op2MiscName[op]
	if !ok {
		s = fmt.Sprintf("Op2Misc(%d)", int(op))
	}
	return s
}

// implement AsmCode interface
func (op Op1Misc) asmcode() {
}

func (op Op2Misc) asmcode() {
}

func (asm *Asm) Op1Misc(op Op1Misc, arg1 AsmCode) *Asm {
	switch op {
	case ALLOC:
		asm.Alloc(arg1.(SoftReg))
	case FREE:
		asm.Free(arg1.(SoftReg))
	default:
		errorf("unknown Op1Misc: %v %v", op, arg1)
	}
	return asm
}

func (asm *Asm) Op2Misc(op Op2Misc, arg1 AsmCode, arg2 AsmCode) *Asm {
	switch op {
	case PUSH:
		asm.Push(arg1.(Reg), arg2.(*SaveSlot))
	case POP:
		asm.Pop(arg1.(Reg), arg2.(*SaveSlot))
	default:
		errorf("unknown Op2Misc: %v %v %v", op, arg1, arg2)
	}
	return asm
}

func (asm *Asm) Push(r Reg, index *SaveSlot) *Asm {
	r.Validate()
	if !asm.RegIsUsed(r.id) {
		// mark in use, caller wants this register
		asm.RegIncUse(r.id)
		*index = InvalidSlot
		return asm
	}
	idx := asm.save.Alloc()
	if idx == InvalidSlot {
		errorf("save area is full, cannot push register %v", r)
	}
	asm.archSave(r.id, idx)
	*index = idx
	return asm
}

func (asm *Asm) Pop(r Reg, index *SaveSlot) *Asm {
	r.Validate()
	idx := *index
	if idx == InvalidSlot {
		asm.RegDecUse(r.id)
		return asm
	}
	asm.save.Validate(idx)
	asm.archRestore(r.id, idx)
	asm.save.Free(idx)
	return asm
}

func (asm *Asm) archSave(id RegId, index SaveSlot) {
	asm.Store(
		Reg{id: id, kind: Uint64},
		Mem{off: int32(index) * 8, reg: asm.save.reg},
	)
}

func (asm *Asm) archRestore(id RegId, index SaveSlot) {
	asm.Load(
		Mem{off: int32(index) * 8, reg: asm.save.reg},
		Reg{id: id, kind: Uint64},
	)
}

/*
func (asm *Asm) pushRegs(rs *Regs) *Regs {
	var ret Regs
	v := &Var{}
	for r := Lo; r <= Hi; r++ {
		if !rs.Contains(r) {
			continue
		}
		if asm.Save.idx >= asm.Save.end {
			errorf("save area is full, cannot push registers")
		}
		v.idx = asm.save.idx
		asm.storeReg(v, r)
		asm.save.idx++
		ret.Set(r)
	}
	return &ret
}

func (asm *Asm) popRegs(rs *Regs) {
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
*/

// find a free slot and return it. return InvalidIndex on failure
func (s *Save) Alloc() SaveSlot {
	for ; s.next < s.end; s.next++ {
		if !s.bitmap[s.next-s.start] {
			idx := s.next
			s.next++
			return idx
		}
	}
	return InvalidSlot
}

// free a slot.
func (s *Save) Free(idx SaveSlot) {
	s.Validate(idx)
	s.bitmap[idx-s.start] = false
	for ; s.next > s.start; s.next-- {
		if s.bitmap[s.next-1-s.start] {
			break
		}
	}
}

// validate a slot
func (s *Save) Validate(idx SaveSlot) {
	if !s.Valid(idx) {
		errorf("invalid save area index %v", idx)
	}
}

// validate a slot
func (s *Save) Valid(idx SaveSlot) bool {
	return idx >= s.start && idx < s.end && s.bitmap[idx-s.start]
}
