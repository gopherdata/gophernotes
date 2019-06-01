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
 * gen_op1.go
 *
 *  Created on Jan 28, 2019
 *      Author Massimiliano Ghilardi
 */

package main

import (
	"fmt"
	"io"
	"os"

	amd64 "github.com/cosmos72/gomacro/jit/amd64"
)

type genOp1 struct {
	opname, opName string
	w              io.Writer
}

func NewGenOp1(w io.Writer, opname string) *genOp1 {
	return &genOp1{
		opname: opname,
		opName: string(opname[0]-'a'+'A') + opname[1:],
		w:      w,
	}
}

func GenOp1() {
	for _, opname := range [...]string{"inc", "dec", "neg", "not", "div", "idiv"} {
		f, err := os.Create("_gen_" + opname + ".s")
		if err != nil {
			panic(err)
		}
		g := NewGenOp1(f, opname)
		g.generate()
		f.Close()
	}
}

func (g *genOp1) generate() {
	g.fileHeader()
	g.opReg()
	g.opMem()
}

func (g *genOp1) fileHeader() {
	fmt.Fprintf(g.w,
		`	.file	"%s.s"
	.text
`, g.opname)
}

func (g *genOp1) funcHeader(funcName string) {
	fmt.Fprintf(g.w,
		`
	.p2align 4,,15
	.globl	%s%s
	.type	%s%s, @function
%s%s:
	.cfi_startproc
`, g.opName, funcName, g.opName, funcName, g.opName, funcName)
}

func (g *genOp1) funcFooter() {
	fmt.Fprint(g.w, `	ret
	.cfi_endproc

`)
}

func (g *genOp1) opReg() {
	g.funcHeader("Reg")
	for _, k := range [...]amd64.Kind{amd64.Uint8, amd64.Uint16, amd64.Uint32, amd64.Uint64} {
		fmt.Fprintf(g.w, "\t// OP reg%d\n", k.Size()*8)
		for r := amd64.RLo; r <= amd64.RHi; r++ {
			fmt.Fprintf(g.w, "\t%s\t%v\n", g.opname, amd64.MakeReg(r, k))
		}
		fmt.Fprint(g.w, "\tnop\n")
	}
	g.funcFooter()
}

func (g *genOp1) opMem() {
	for _, k := range [...]amd64.Kind{amd64.Uint8, amd64.Uint16, amd64.Uint32, amd64.Uint64} {
		g.opMemKind(k)
	}
}

func (g *genOp1) opMemKind(k amd64.Kind) {
	ksuffix := map[amd64.Size]string{1: "b", 2: "w", 4: "l", 8: "q"}
	klen := k.Size() * 8
	g.funcHeader(fmt.Sprintf("Mem%d", klen))
	offstr := [...]string{"", "0x7F", "0x78563412"}
	for i, offlen := range [...]uint8{0, 8, 32} {
		fmt.Fprintf(g.w, "\t// OP mem%d[off%d]\n", klen, offlen)
		for r := amd64.RLo; r <= amd64.RHi; r++ {
			fmt.Fprintf(g.w, "\t%s%s\t%s(%v)\n", g.opname, ksuffix[k.Size()],
				offstr[i], amd64.MakeReg(r, amd64.Uintptr))
		}
		fmt.Fprint(g.w, "\tnop\n")
	}
	g.funcFooter()
}
