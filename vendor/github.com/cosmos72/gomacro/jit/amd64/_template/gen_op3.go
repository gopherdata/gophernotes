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
 * gen_op3.go
 *
 *  Created on Feb 28, 2019
 *      Author Massimiliano Ghilardi
 */

package main

import (
	"fmt"
	"io"
	"os"

	amd64 "github.com/cosmos72/gomacro/jit/amd64"
)

type genOp3 struct {
	opname, opName string
	w              io.Writer
}

func NewGenOp3(w io.Writer, opname string) *genOp3 {
	return &genOp3{
		opname: opname,
		opName: string(opname[0]-'a'+'A') + opname[1:],
		w:      w,
	}
}

func GenOp3() {
	for _, opname := range [...]string{"getidx", "setidx"} {
		f, err := os.Create("_gen_" + opname + ".s")
		if err != nil {
			panic(err)
		}
		g := NewGenOp3(f, opname)
		g.generate()
		f.Close()
	}
}

func (g *genOp3) generate() {
	g.fileHeader()

	if g.opname == "getidx" {
		g.opGetReg()
	} else {
		g.opSetConst()
		g.opSetReg()
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

func (g *genOp3) opGetReg() {
	for _, k := range [...]amd64.Kind{amd64.Uint8, amd64.Uint16, amd64.Uint32, amd64.Uint64} {
		g.opGetRegKind(k)
	}
}

func (g *genOp3) opGetRegKind(k amd64.Kind) {
	kbits := k.Size() * 8
	rbase := amd64.MakeReg(amd64.RLo, amd64.Uintptr)
	rscale := rbase
	rdst := amd64.MakeReg(amd64.RLo, k)

	g.funcHeader(fmt.Sprintf("Reg%d", kbits))
	fmt.Fprintf(g.w, "\t// reg%d = mem[reg]\n", kbits)
	for _, scale := range [...]uint8{1, 2, 4, 8} {
		for r1 := amd64.RLo; r1 <= amd64.RHi; r1++ {
			fmt.Fprintf(g.w, "\tmov\t(%v,%v,%d),%v\n",
				amd64.MakeReg(r1, amd64.Uintptr),
				rscale, scale,
				rdst)
		}
		fmt.Fprint(g.w, "\tnop\n")
		for r2 := amd64.RLo; r2 <= amd64.RHi; r2++ {
			if r2 == amd64.RSP {
				// not supported by amd64 assembly
				continue
			}
			fmt.Fprintf(g.w, "\tmov\t(%v,%v,%d),%v\n",
				rbase,
				amd64.MakeReg(r2, amd64.Uintptr), scale,
				rdst)
		}
		fmt.Fprint(g.w, "\tnop\n")
		for r3 := amd64.RLo; r3 <= amd64.RHi; r3++ {
			fmt.Fprintf(g.w, "\tmov\t(%v,%v,%d),%v\n",
				rbase,
				rscale, scale,
				amd64.MakeReg(r3, k))
		}
		fmt.Fprint(g.w, "\tnop\n")
	}
	fmt.Fprint(g.w, "\tnop\n")
	g.funcFooter()
}

func (g *genOp3) opSetReg() {
	for _, k := range [...]amd64.Kind{amd64.Uint8, amd64.Uint16, amd64.Uint32, amd64.Uint64} {
		g.opSetRegKind(k)
	}
}

func (g *genOp3) opSetRegKind(k amd64.Kind) {
	kbits := k.Size() * 8
	rsrc := amd64.MakeReg(amd64.RLo, k)
	rbase := amd64.MakeReg(amd64.RLo, amd64.Uintptr)
	rscale := rbase

	g.funcHeader(fmt.Sprintf("Reg%d", kbits))
	fmt.Fprintf(g.w, "\t// reg%d = mem[reg]\n", kbits)
	for _, scale := range [...]uint8{1, 2, 4, 8} {
		for r1 := amd64.RLo; r1 <= amd64.RHi; r1++ {
			fmt.Fprintf(g.w, "\tmov\t%v,(%v,%v,%d)\n",
				amd64.MakeReg(r1, k),
				rbase,
				rscale, scale)
		}
		fmt.Fprint(g.w, "\tnop\n")
		for r2 := amd64.RLo; r2 <= amd64.RHi; r2++ {
			fmt.Fprintf(g.w, "\tmov\t%v,(%v,%v,%d)\n",
				rsrc,
				amd64.MakeReg(r2, amd64.Uintptr),
				rscale, scale)
		}
		fmt.Fprint(g.w, "\tnop\n")
		for r3 := amd64.RLo; r3 <= amd64.RHi; r3++ {
			if r3 == amd64.RSP {
				// not supported by amd64 assembly
				continue
			}
			fmt.Fprintf(g.w, "\tmov\t%v,(%v,%v,%d)\n",
				rsrc,
				rbase,
				amd64.MakeReg(r3, amd64.Uintptr), scale)
		}
		fmt.Fprint(g.w, "\tnop\n")
	}
	fmt.Fprint(g.w, "\tnop\n")
	g.funcFooter()
}

func (g *genOp3) opSetConst() {
	for _, k := range [...]amd64.Kind{amd64.Uint8, amd64.Uint16, amd64.Uint32, amd64.Uint64} {
		g.opSetConstKind(k)
	}
}

func (g *genOp3) opSetConstKind(k amd64.Kind) {
	kbits := k.Size() * 8
	rbase := amd64.MakeReg(amd64.RLo, amd64.Uintptr)
	rscale := rbase
	g.funcHeader(fmt.Sprintf("Const%d", kbits))
	suffixstr := map[amd64.Size]string{1: "b", 2: "w", 4: "l", 8: "q"}[k.Size()]
	conststr := map[amd64.Size]string{8: "$0x33", 16: "$0xaabb", 32: "$0x11223344", 64: "$0x55667788"}[kbits]

	fmt.Fprintf(g.w, "\t// mem[reg] = const%d\n", kbits)
	for _, scale := range [...]uint8{1, 2, 4, 8} {
		for r1 := amd64.RLo; r1 <= amd64.RHi; r1++ {
			fmt.Fprintf(g.w, "\tmov%s\t%v,(%v,%v,%d)\n", suffixstr,
				conststr,
				amd64.MakeReg(r1, amd64.Uintptr),
				rscale, scale)
		}
		fmt.Fprint(g.w, "\tnop\n")
		for r2 := amd64.RLo; r2 <= amd64.RHi; r2++ {
			if r2 == amd64.RSP {
				// not supported by amd64 assembly
				continue
			}
			fmt.Fprintf(g.w, "\tmov%s\t%v,(%v,%v,%d)\n", suffixstr,
				conststr,
				rbase,
				amd64.MakeReg(r2, amd64.Uintptr), scale)
		}
		fmt.Fprint(g.w, "\tnop\n")
	}
	g.funcFooter()
}
