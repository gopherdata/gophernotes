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
 * api.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package disasm

import (
	"testing"
	"unsafe"

	"github.com/bnagy/gapstone"
	. "github.com/cosmos72/gomacro/jit"
	"github.com/cosmos72/gomacro/jit/asm"
)

type Engine = gapstone.Engine

func NewDisasm(archId ArchId) (Engine, error) {
	var arch uint = gapstone.CS_ARCH_X86
	var mode uint = gapstone.CS_MODE_64
	if archId == ARM64 {
		arch = gapstone.CS_ARCH_ARM64
		mode = gapstone.CS_MODE_ARM // | gapstone.CS_MODE_V8
	}
	engine, err := gapstone.New(
		int(arch),
		mode,
	)
	if err != nil {
		return engine, err
	}
	engine.SetOption(gapstone.CS_OPT_SYNTAX, gapstone.CS_OPT_SYNTAX_ATT)
	return engine, nil
}

func Disasm(code asm.MachineCode) ([]gapstone.Instruction, error) {
	engine, err := NewDisasm(code.ArchId)
	if err != nil {
		return nil, err
	}
	address := uint64(0x10000)
	if len(code.Bytes) != 0 {
		address = uint64(uintptr(unsafe.Pointer(&code.Bytes[0])))
	}
	return engine.Disasm(code.Bytes, address, 0)
}

func PrintDisasm(t *testing.T, code asm.MachineCode) {
	insns, err := Disasm(code)
	if err != nil {
		t.Error(err)
	} else {
		archId := code.ArchId
		t.Logf("----- %v -----", archId)
		for _, insn := range insns {
			Show(t, archId, insn)
		}
	}
}

func Show(t *testing.T, archId ArchId, insn gapstone.Instruction) {
	var prefix string
	bytes := insn.Bytes
	if archId == ARM64 && len(bytes) == 4 {
		// print high byte first
		prefix = "0x"
		bytes[0], bytes[1], bytes[2], bytes[3] = bytes[3], bytes[2], bytes[1], bytes[0]
	}
	t.Logf(" %s%x%s%s\t%s", prefix, bytes, spaces(2*len(insn.Bytes)), insn.Mnemonic, insn.OpStr)
}

func spaces(n int) string {
	return "                "[n%16:]
}
