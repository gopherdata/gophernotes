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
 * comp.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"bytes"

	"github.com/cosmos72/gomacro/jit/asm"
)

// argument of Comp.Compile()
type Source []interface{}

type Comp struct {
	code Code
	// code[toassemble] is the first AsmCode
	// not yet assembled
	toassemble  int
	nextSoftReg SoftRegId
	nextTempReg SoftRegId
	arch        Arch
	asm.RegIdConfig
	asm *Asm
}

func New() *Comp {
	var c Comp
	return c.Init()
}

func NewArchId(archId ArchId) *Comp {
	var c Comp
	return c.InitArchId(archId)
}

func NewArch(arch Arch) *Comp {
	var c Comp
	return c.InitArch(arch)
}

func (c *Comp) Init() *Comp {
	return c.InitArchId(ARCH_ID)
}

func (c *Comp) InitArchId(archId ArchId) *Comp {
	return c.InitArch(Archs[archId])
}

func (c *Comp) InitArch(arch Arch) *Comp {
	if arch == nil {
		errorf("unknown arch")
	}
	c.code = nil
	c.toassemble = 0
	c.nextSoftReg = FirstSoftRegId
	c.nextTempReg = FirstTempRegId
	c.arch = arch
	c.RegIdConfig = arch.RegIdConfig()
	if c.asm != nil {
		c.asm.InitArch(arch)
	}
	return c
}

func (c *Comp) Arch() Arch {
	return c.arch
}

func (c *Comp) ArchId() ArchId {
	if c.arch == nil {
		return asm.NOARCH
	}
	return c.arch.Id()
}

// return symbolic assembly code
func (c *Comp) Code() Code {
	return c.code
}

// call Comp.Assemble() followed by Comp.Asm().Epilogue()
func (c *Comp) Epilogue() {
	c.Assemble()
	c.Asm().Epilogue()
}

// discard assembly code and machine code
func (c *Comp) ClearCode() {
	c.code = nil
	c.toassemble = 0
	if c.asm != nil {
		c.asm.ClearCode()
	}
}

// forget all allocated registers
func (c *Comp) ClearRegs() {
	c.nextSoftReg = FirstSoftRegId
	c.nextTempReg = FirstTempRegId
	if c.asm != nil {
		c.asm.ClearRegs()
	}
}

// return assembler
func (c *Comp) Asm() *Asm {
	if c.asm == nil {
		// create asm.Asm on demand
		c.asm = asm.NewArch(c.arch)
	}
	return c.asm
}

// assemble the code compiled since previous call
// to Assemble(), and return machine code
func (c *Comp) Assemble() MachineCode {
	asm := c.Asm()
	if len(c.code) > c.toassemble {
		asm.Assemble(c.code[c.toassemble:]...)
		c.toassemble = len(c.code)
	}
	return asm.Code()
}

/*
 * call Assemble(), then set *funcaddr to assembled machine code.
 *
 * funcaddr must be a non-nil pointer to function.
 *
 * function type MUST match the code created by the programmer,
 * or BAD things will happen: crash, memory corruption, undefined behaviour...
 *
 * Obviously, code created by the programmer must be for the same architecture
 * the program is currently running on...
 *
 * Caller likely needs to call ClearCode() after this function returns
 */
func (c *Comp) Func(funcaddr interface{}) {
	c.Assemble()
	c.asm.Func(funcaddr)
}

func checkAssignable(e Expr) {
	switch e.(type) {
	case Reg, Mem, SoftReg:
		break
	default:
		errorf("cannot assign to %v", e)
	}
}

func (c *Comp) MakeParam(off int32, kind Kind) Mem {
	return MakeParam(off, kind, c.RegIdConfig)
}

func (c *Comp) MakeVar(idx int, kind Kind) Mem {
	mem, err := MakeVar(idx, kind, c.RegIdConfig)
	if err != nil {
		panic(err)
	}
	return mem
}

// compile statements and their arguments
func (c *Comp) Compile(s Source) {
	for i := 0; i < len(s); i++ {
		switch inst := s[i].(type) {
		case Inst1:
			c.Stmt1(inst, s[i+1].(Expr))
			i++
		case Inst1Misc:
			c.SoftReg(inst, s[i+1].(SoftReg))
			i++
		case Inst2:
			c.Stmt2(inst, s[i+1].(Expr), s[i+2].(Expr))
			i += 2
		case Inst3:
			c.Stmt3(inst, s[i+1].(Expr), s[i+2].(Expr), s[i+3].(Expr))
			i += 3
		default:
			errorf("unknown instruction type %T, expecting Inst1, Inst2, Inst3 or Inst1Misc", inst)
		}
	}
}

// pretty-print Source
func (s Source) String() string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if i != 0 {
			buf.WriteByte(' ')
		}
		switch inst := s[i].(type) {
		case Inst1:
			buf.WriteString(NewStmt1(inst, s[i+1].(Expr)).String())
			i++
		case Inst1Misc:
			buf.WriteString(inst.String())
			buf.WriteByte(' ')
			buf.WriteString(s[i+1].(SoftReg).String())
			buf.WriteByte(';')
			i++
		case Inst2:
			buf.WriteString(NewStmt2(inst, s[i+1].(Expr), s[i+2].(Expr)).String())
			i += 2
		case Inst3:
			buf.WriteString(NewStmt3(inst, s[i+1].(Expr), s[i+2].(Expr), s[i+3].(Expr)).String())
			i += 3
		default:
			errorf("unknown instruction type %T, expecting Inst1, Inst2 or Inst3", inst)
		}
	}
	return buf.String()
}
