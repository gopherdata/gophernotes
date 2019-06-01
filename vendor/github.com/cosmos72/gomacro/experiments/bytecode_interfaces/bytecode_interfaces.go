/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * bytecode_interfaces.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package bytecode_interfaces

import (
	"time"
)

type Op uint8

const (
	storeInt16 Op = iota
	addIntInt
	subIntInt
	mulIntInt
	andIntInt
	rshiftIntInt
	jmp
	jmpIfGtrInt
	jmpIfLeqInt
	jmpIfEqlInt
	ret
)

type Inst struct {
	Op            Op
	Dst, Lhs, Rhs uint8
}

func Uint8(n int8) uint8 {
	return uint8(n)
}

func Int(n uint8) int {
	return int(int8(n))
}

func (i Inst) Src() int {
	return int(uint16(i.Lhs) | uint16(i.Rhs)<<8)
}

func Inst16(op Op, dst uint8, src int16) Inst {
	return Inst{
		Op:  op,
		Dst: dst,
		Lhs: uint8(src),
		Rhs: uint8(src >> 8),
	}
}

type Prog struct {
	IP    int
	Code  []Inst
	Vars  []interface{}
	Outer *Prog
}

func (p *Prog) Exec(IP int) []interface{} {
	if IP < 0 {
		IP = p.IP
	}
	code := p.Code
	vars := p.Vars
	for {
		if false {
			// Debugf("fetching IP=%v\n", IP)
		}
		i := code[IP]
		if false {
			// Debugf("IP=%v,\tinst=%v\tStack=%v\n", IP, i, vars)
			time.Sleep(time.Second)
		}
		IP++
		switch i.Op {

		case storeInt16:
			vars[i.Dst] = i.Src()
		case addIntInt:
			lhs := vars[i.Lhs].(int)
			rhs := vars[i.Rhs].(int)
			vars[i.Dst] = lhs + rhs
		case subIntInt:
			lhs := vars[i.Lhs].(int)
			rhs := vars[i.Rhs].(int)
			vars[i.Dst] = lhs - rhs
		case mulIntInt:
			lhs := vars[i.Lhs].(int)
			rhs := vars[i.Rhs].(int)
			vars[i.Dst] = lhs * rhs
		case andIntInt:
			lhs := vars[i.Lhs].(int)
			rhs := vars[i.Rhs].(int)
			vars[i.Dst] = lhs & rhs
		case rshiftIntInt:
			lhs := vars[i.Lhs].(int)
			rhs := vars[i.Rhs].(int)
			vars[i.Dst] = lhs >> uint(rhs)
		case jmp:
			IP += Int(i.Dst)
		case jmpIfGtrInt:
			lhs := vars[i.Lhs].(int)
			rhs := vars[i.Rhs].(int)
			if lhs > rhs {
				IP += Int(i.Dst)
			}
		case jmpIfLeqInt:
			lhs := vars[i.Lhs].(int)
			rhs := vars[i.Rhs].(int)
			if lhs <= rhs {
				IP += Int(i.Dst)
			}
		case jmpIfEqlInt:
			lhs := vars[i.Lhs].(int)
			rhs := vars[i.Rhs].(int)
			if lhs == rhs {
				IP += Int(i.Dst)
			}
		case ret:
			return vars[i.Lhs:i.Rhs]
		}
	}
}

/*
  interpreted version of:

	func collatz(n int) {
		for n > 1 {
			if n&1 != 0 {
				n = ((n * 3) + 1) / 2
			} else {
				n = n / 2
			}
		}
	}
*/
func BytecodeCollatz() *Prog {
	const (
		n = iota
		tmp
		_0
		_1
		_3
	)
	return &Prog{Vars: []interface{}{n: 0, _0: 0, _1: 1, _3: 3},
		Code: []Inst{
			{jmpIfLeqInt, 6, n, _1},
			{andIntInt, tmp, n, _1},
			{jmpIfEqlInt, 2, tmp, _0},

			{mulIntInt, tmp, n, _3},
			{addIntInt, n, tmp, _1},

			{rshiftIntInt, n, n, _1},
			{jmp, Uint8(-7), 0, 0},

			{ret, 0, 0, 0},
		}}
}

/*
  bytecode version of:

	func sum(n int) int {
		total := 0
		for i := 1; i <= n; i++ {
			total += i
		}
		return total
	}
*/
func BytecodeSum(N int) *Prog {
	const (
		i = iota
		n
		total
		_1
	)
	return &Prog{Vars: []interface{}{n: N, _1: 1},
		Code: []Inst{
			Inst16(storeInt16, i, 1),
			Inst16(storeInt16, total, 0),
			{jmpIfGtrInt, 3, i, n},
			{addIntInt, total, total, i},
			{addIntInt, i, i, _1},
			{jmp, Uint8(-4), 0, 0},
			{ret, 0, total, total + 1},
		}}
}
