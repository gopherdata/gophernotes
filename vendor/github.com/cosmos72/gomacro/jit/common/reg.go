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
 * reg.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"fmt"
)

// machine register
type RegId uint16

func (id RegId) ArchId() ArchId {
	return ArchId(1 + id>>8)
}

func (id RegId) Arch() Arch {
	return Archs[id.ArchId()]
}

func (id RegId) String() string {
	arch := id.Arch()
	if arch != nil {
		return arch.RegIdString(id)
	}
	return fmt.Sprintf("%%unknown_reg(%#x)", uint8(id))
}

func (id RegId) Valid() bool {
	return id.Arch().RegIdValid(id)
}

func (id RegId) Validate() {
	if !id.Valid() {
		errorf("invalid register: %v", id)
	}
}

// ===================================

type RegIdConfig struct {
	RLo, RHi, RSP, RVAR RegId
	/**
	 * first RegId to allocate.
	 * subsequent allocations will return progressively higher registers,
	 * eventually reach RHi, wrap around to RLo, and finally reach
	 * RAllocFirst again when all registers are allocated.
	 *
	 * used on amd64 to allocate RAX, RCX and RBX as last
	 * because some assembly instructions (shift, division)
	 * are hardcoded to use them
	 */
	RAllocFirst RegId
}

// register + kind
type Reg struct {
	id   RegId
	kind Kind // defines width and signedness
}

func MakeReg(id RegId, kind Kind) Reg {
	return Reg{id: id, kind: kind}
}

// implement Arg interface
func (r Reg) RegId() RegId {
	return r.id
}

func (r Reg) Kind() Kind {
	return r.kind
}

func (r Reg) Const() bool {
	return false
}

func (r Reg) asmcode() {
}

func (r Reg) String() string {
	arch := r.id.Arch()
	if arch != nil {
		return arch.RegString(r)
	}
	return fmt.Sprintf("%%unknown_reg(%#x,%v)", uint8(r.id), r.kind)
}

func (r Reg) Valid() bool {
	return r.id.Valid()
}

func (r Reg) Validate() {
	r.id.Validate()
}

// ===================================

type RegIds struct {
	inuse    map[RegId]uint32 // RegId -> use count
	first    RegId            // first RegId to allocate
	curr     RegId            // next RegId to allocate
	rlo, rhi RegId
}

func (rs *RegIds) Copy(other *RegIds) {
	if rs == other {
		return
	}
	rs.inuse = make(map[RegId]uint32)
	rs.first = other.first
	rs.curr = other.curr
	rs.rlo = other.rlo
	rs.rhi = other.rhi
	for id, use := range other.inuse {
		rs.inuse[id] = use
	}
}

func (rs *RegIds) IsUsed(id RegId) bool {
	return id.Valid() && rs.inuse[id] != 0
}

// return new use count
func (rs *RegIds) IncUse(id RegId) uint32 {
	if !id.Valid() {
		return 0
	}
	count := rs.inuse[id]
	if count < ^uint32(0) {
		count++
		rs.inuse[id] = count
	}
	return count
}

// return new use count
func (rs *RegIds) DecUse(id RegId) uint32 {
	if !id.Valid() {
		return 0
	}
	count := rs.inuse[id]
	switch count {
	case 0:
		return count
	case 1:
		delete(rs.inuse, id)
	default:
		rs.inuse[id] = count - 1
	}
	return count - 1
}

// return the RegId immediately after id,
// wrapping around after RHi.
// returned RegId may be used or not valid:
// it is caller's responsibility to check
// for valid and unused registers
func (rs *RegIds) Next(id RegId) RegId {
	if id >= rs.rhi {
		return rs.rlo
	}
	return id + 1
}

func (rs *RegIds) TryAlloc() RegId {
	id := rs.curr
	// fmt.Printf("TryAlloc: RegIds = %+v\n", *rs)
	for {
		next := rs.Next(id)
		// fmt.Printf("trying RegId = %d, next = %d\n", id, next)
		// time.Sleep(time.Second)
		if id.Valid() && rs.inuse[id] == 0 {
			rs.inuse[id] = 1
			rs.curr = next
			return id
		}
		id = next
		if id == rs.curr {
			// did a full circle,
			// no free register
			return NoRegId
		}
	}
}

func (rs *RegIds) Free(id RegId) {
	if id.Valid() && rs.DecUse(id) == 0 && id >= rs.first &&
		(id < rs.curr || rs.curr < rs.first) {

		rs.curr = id
	}
	// fmt.Printf("Free: RegIds = %+v, freed %d\n", *rs, id)
}

// ===================================

func (asm *Asm) RegIsUsed(id RegId) bool {
	return asm.regIds.IsUsed(id)
}

// return new use count
func (asm *Asm) RegIncUse(id RegId) uint32 {
	return asm.regIds.IncUse(id)
}

// return new use count
func (asm *Asm) RegDecUse(id RegId) uint32 {
	return asm.regIds.DecUse(id)
}

func (asm *Asm) TryRegAlloc(kind Kind) Reg {
	id := asm.regIds.TryAlloc()
	if !id.Valid() {
		return Reg{}
	}
	return Reg{id, kind}
}

func (asm *Asm) RegAlloc(kind Kind) Reg {
	r := asm.TryRegAlloc(kind)
	if !r.Valid() {
		errorf("no free registers")
	}
	return r
}

func (asm *Asm) RegFree(r Reg) *Asm {
	asm.regIds.Free(r.id)
	return asm
}
