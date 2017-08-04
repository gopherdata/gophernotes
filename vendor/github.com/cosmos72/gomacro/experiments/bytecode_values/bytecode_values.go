/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * bytecode_values.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package bytecode_values

import (
	r "reflect"
	"time"
)

type Op uint8

const (
	storeInt16 Op = iota
	addIntInt
	subIntInt
	jmp
	jmpIfGtrInt
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
	Vars  []r.Value
	Outer *Prog
}

func (p *Prog) Exec(IP int) []r.Value {
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
			vars[i.Dst] = r.ValueOf(i.Src())
		case addIntInt:
			lhs := int(vars[i.Lhs].Int())
			rhs := int(vars[i.Rhs].Int())
			vars[i.Dst] = r.ValueOf(lhs + rhs)
		case subIntInt:
			lhs := int(vars[i.Lhs].Int())
			rhs := int(vars[i.Rhs].Int())
			vars[i.Dst] = r.ValueOf(lhs - rhs)
		case jmp:
			IP += Int(i.Dst)
		case jmpIfGtrInt:
			lhs := int(vars[i.Lhs].Int())
			rhs := int(vars[i.Rhs].Int())
			if lhs > rhs {
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
				n = (n + 3) >> 1
			} else {
				n = n >> 1
			}
		}
}
*/

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
	return &Prog{Vars: []r.Value{n: r.ValueOf(N), _1: r.ValueOf(1)},
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
