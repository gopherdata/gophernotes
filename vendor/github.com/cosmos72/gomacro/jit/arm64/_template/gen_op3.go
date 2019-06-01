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
 * gen_op2.go
 *
 *  Created on Jan 28, 2019
 *      Author Massimiliano Ghilardi
 */

package main

import (
	"fmt"
	"io"
	"os"

	arch "github.com/cosmos72/gomacro/jit/arm64"
)

type genOp3 struct {
	opname, opName string
	w              io.Writer
}

func GenOp3() {
	for _, opname := range [...]string{
		"adc", "add", "sub", "sbc",
		"mul", "sdiv", "udiv",
		"and", "orr", "eor", "lsl", "lsr", "asr",
		"ldr", "str",
	} {
		f, err := os.Create("_gen_" + opname + ".s")
		if err != nil {
			panic(err)
		}
		g := newGenOp3(f, opname)
		g.generate()
		f.Close()
	}
}

func newGenOp3(w io.Writer, opname string) *genOp3 {
	return &genOp3{
		opname: opname,
		opName: string(opname[0]-'a'+'A') + opname[1:],
		w:      w,
	}
}

func (g *genOp3) generate() {
	g.fileHeader()
	switch g.opname {
	case "ldr", "str":
		g.opLoadStore()
	case "add", "sub":
		g.opRegRegReg()
		g.opAddSubRegRegConst()
	case "and", "orr", "eor":
		g.opRegRegReg()
		g.opBitwiseRegRegConst()
	case "lsl", "lsr", "asr":
		g.opRegRegReg()
		g.opShiftRegRegConst()
	default:
		g.opRegRegReg()
	}
}

func (g *genOp3) fileHeader() {
	fmt.Fprintf(g.w,
		`	.file	"%s.s"
	.text
`, g.opname)
}

func (g *genOp3) funcHeader(funcName string) {
	fmt.Fprintf(g.w,
		`
	.p2align 4,,15
	.globl	%s%s
	.type	%s%s, @function
%s%s:
	.cfi_startproc
`, g.opName, funcName, g.opName, funcName, g.opName, funcName)
}

func (g *genOp3) funcFooter() {
	fmt.Fprint(g.w, `	ret
	.cfi_endproc

`)
}

func (g *genOp3) opLoadStore() {
	g.funcHeader("RegRegReg")
	rlo := arch.MakeReg(arch.RLo, arch.Uint64)
	widths := [...]string{"b", "h", "", ""}
	shifts := [...]string{"", ", lsl #1", ", lsl #2", ", lsl #3"}
	for i, k := range [...]arch.Kind{arch.Uint8, arch.Uint16, arch.Uint32, arch.Uint64} {
		width := widths[i]
		shift := shifts[i]
		kbits := k.Size() * 8
		fmt.Fprintf(g.w, "\t// %s %d\n", g.opname, kbits)
		rlok := arch.MakeReg(arch.RLo, k)
		for id := arch.RLo; id < arch.RHi; id++ {
			fmt.Fprintf(g.w, "\t%s%s\t%v, [%v, %v%s]\n", g.opname, width,
				arch.MakeReg(id, k), rlo, rlo, shift)
		}
		fmt.Fprint(g.w, "\tnop\n")
		for id := arch.RLo; id < arch.RHi; id++ {
			fmt.Fprintf(g.w, "\t%s%s\t%v, [%v, %v%s]\n", g.opname, width,
				rlok, arch.MakeReg(id, arch.Uint64), rlo, shift)
		}
		fmt.Fprint(g.w, "\tnop\n")
		for id := arch.RLo; id <= arch.RHi; id++ {
			fmt.Fprintf(g.w, "\t%s%s\t%v, [%v, %v%s]\n", g.opname, width,
				rlok, rlo, arch.MakeReg(id, arch.Uint64), shift)
		}
		fmt.Fprint(g.w, "\tnop\n")
	}
	g.funcFooter()
}

func (g *genOp3) opRegRegReg() {
	g.funcHeader("RegRegReg")
	for _, k := range [...]arch.Kind{arch.Uint32, arch.Uint64} {
		kbits := k.Size() * 8
		fmt.Fprintf(g.w, "\t// reg%d OP= reg%d, reg%d\n", kbits, kbits, kbits)
		rlo := arch.MakeReg(arch.RLo, k)
		for id := arch.RLo; id < arch.RHi; id++ {
			fmt.Fprintf(g.w, "\t%s\t%v,%v,%v\n", g.opname, arch.MakeReg(id, k), rlo, rlo)
		}
		fmt.Fprint(g.w, "\tnop\n")
		for id := arch.RLo; id < arch.RHi; id++ {
			fmt.Fprintf(g.w, "\t%s\t%v,%v,%v\n", g.opname, rlo, arch.MakeReg(id, k), rlo)
		}
		fmt.Fprint(g.w, "\tnop\n")
		for id := arch.RLo; id < arch.RHi; id++ {
			fmt.Fprintf(g.w, "\t%s\t%v,%v,%v\n", g.opname, rlo, rlo, arch.MakeReg(id, k))
		}
		fmt.Fprint(g.w, "\tnop\n")
	}
	g.funcFooter()
}

// add|sub xn, xm, 12-bit-immediate shifted by 0|12
func (g *genOp3) opAddSubRegRegConst() {
	for _, k := range [...]arch.Kind{arch.Uint32, arch.Uint64} {
		g.opAddSubRegRegConstKind(k, k.Size()*8)
	}
}

// add|sub xn, xm, 12-bit-immediate shifted by 0|12
func (g *genOp3) opAddSubRegRegConstKind(k arch.Kind, kbits arch.Size) {
	g.funcHeader(fmt.Sprintf("Reg%dReg%dConst", kbits, kbits))
	rlo := arch.MakeReg(arch.RLo, k)
	conststr := "#0x0"
	for r := arch.RLo; r < arch.RHi; r++ {
		fmt.Fprintf(g.w, "\t%s\t%v,%v,%s\n", g.opname, arch.MakeReg(r, k), rlo, conststr)
	}
	for r := arch.RLo; r < arch.RHi; r++ {
		fmt.Fprintf(g.w, "\t%s\t%v,%v,%s\n", g.opname, rlo, arch.MakeReg(r, k), conststr)
	}
	for constval := 1; constval <= 0xFFFFFF; constval *= 2 {
		fmt.Fprintf(g.w, "\t%s\t%v,%v,#%#x\n", g.opname, rlo, rlo, constval)
	}
	g.funcFooter()
}

// lsl|lsr|asr xn, xm, {0..63}
func (g *genOp3) opShiftRegRegConst() {
	for _, k := range [...]arch.Kind{arch.Uint32, arch.Uint64} {
		g.opShiftRegRegConstKind(k, k.Size()*8)
	}
}

// lsl|lsr|asr xn, xm, {0..63}
func (g *genOp3) opShiftRegRegConstKind(k arch.Kind, kbits arch.Size) {
	g.funcHeader(fmt.Sprintf("Reg%dReg%dConst", kbits, kbits))
	rlo := arch.MakeReg(arch.RLo, k)
	conststr := "#31"
	for r := arch.RLo; r < arch.RHi; r++ {
		fmt.Fprintf(g.w, "\t%s\t%v,%v,%s\n", g.opname, arch.MakeReg(r, k), rlo, conststr)
	}
	for r := arch.RLo; r < arch.RHi; r++ {
		fmt.Fprintf(g.w, "\t%s\t%v,%v,%s\n", g.opname, rlo, arch.MakeReg(r, k), conststr)
	}
	for constval := arch.Size(0); constval < kbits; constval++ {
		fmt.Fprintf(g.w, "\t%s\t%v,%v,#%#x\n", g.opname, rlo, rlo, constval)
	}
	g.funcFooter()
}

// and|orr|eor xn, xm, complicated immediate constant
func (g *genOp3) opBitwiseRegRegConst() {
	for _, k := range [...]arch.Kind{arch.Uint32, arch.Uint64} {
		g.opBitwiseRegRegConstKind(k, k.Size()*8)
	}
}

// and|orr|eor xn, xm, complicated immediate constant
func (g *genOp3) opBitwiseRegRegConstKind(k arch.Kind, kbits arch.Size) {
	g.funcHeader(fmt.Sprintf("Reg%dReg%dConst", kbits, kbits))
	rlo := arch.MakeReg(arch.RLo, k)
	conststr := "#0x1"
	for r := arch.RLo; r < arch.RHi; r++ {
		fmt.Fprintf(g.w, "\t%s\t%v,%v,%s\n", g.opname, arch.MakeReg(r, k), rlo, conststr)
	}
	for r := arch.RLo; r < arch.RHi; r++ {
		fmt.Fprintf(g.w, "\t%s\t%v,%v,%s\n", g.opname, rlo, arch.MakeReg(r, k), conststr)
	}
	for _, constval := range opBitwiseConstants(kbits) {
		fmt.Fprintf(g.w, "\t%s\t%v,%v,#%#x\n", g.opname, rlo, rlo, constval)
	}
	g.funcFooter()
}

// list all possible immediate constants for bitwise operations
// on 32-bit or 64-bit registers
func opBitwiseConstants(kbits arch.Size) []uint64 {
	switch kbits {
	case 32:
		return opBitwiseConstants32()
	default:
		return opBitwiseConstants64()
	}
}

// list all possible immediate constants for bitwise operations
// on 64-bit registers
func opBitwiseConstants32() []uint64 {
	var result []uint64
	var bitmask uint32
	var size, length, e, rotation uint8
	for size = 2; size <= 32; size *= 2 {
		for length = 1; length < size; length++ {
			bitmask = 0xffffffff >> (32 - length)
			for e = size; e < 32; e *= 2 {
				bitmask |= bitmask << e
			}
			for rotation = 0; rotation < size; rotation++ {
				result = append(result, uint64(bitmask))
				bitmask = (bitmask >> 1) | (bitmask << 31)
			}
		}
	}
	return result
}

// list all possible immediate constants for bitwise operations
// on 64-bit registers
func opBitwiseConstants64() []uint64 {
	var result []uint64
	var bitmask uint64
	var size, length, e, rotation uint8
	for size = 2; size <= 64; size *= 2 {
		for length = 1; length < size; length++ {
			bitmask = 0xffffffffffffffff >> (64 - length)
			for e = size; e < 64; e *= 2 {
				bitmask |= bitmask << e
			}
			for rotation = 0; rotation < size; rotation++ {
				result = append(result, bitmask)
				bitmask = (bitmask >> 1) | (bitmask << 63)
			}
		}
	}
	return result
}
