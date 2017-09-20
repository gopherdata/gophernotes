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
 * all_test.go
 *
 *  Created on: Mar 06 2017
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"go/ast"
	"go/token"
	r "reflect"
	"testing"
	"time"

	. "github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/classic"
	"github.com/cosmos72/gomacro/fast"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type TestFor int

const (
	I TestFor = 1 << iota
	F
	S // set option OptDebugSleepOnSwitch
	A = I | F
)

type TestCase struct {
	testfor TestFor
	name    string
	program string
	result0 interface{}
	results []interface{}
}

func TestClassic(t *testing.T) {
	ir := classic.New()
	// ir.Options |= OptDebugCallStack | OptDebugPanicRecover
	for _, test := range testcases {
		if test.testfor&I != 0 {
			test := test
			t.Run(test.name, func(t *testing.T) { test.classic(t, ir) })
		}
	}
}

func TestFast(t *testing.T) {
	ir := fast.New()
	for _, test := range testcases {
		if test.testfor&F != 0 {
			test := test
			t.Run(test.name, func(t *testing.T) { test.fast(t, ir) })
		}
	}
}

func (test *TestCase) classic(t *testing.T, ir *classic.Interp) {

	rets := PackValues(ir.Eval(test.program))

	test.compareResults(t, rets)
}

func (test *TestCase) fast(t *testing.T, ir *fast.Interp) {

	if test.testfor&S != 0 {
		ir.Comp.Options |= OptDebugSleepOnSwitch
	} else {
		ir.Comp.Options &^= OptDebugSleepOnSwitch
	}
	rets := PackValues(ir.Eval(test.program))

	test.compareResults(t, rets)
}

const sum_source_string = "func sum(n int) int { total := 0; for i := 1; i <= n; i++ { total += i }; return total }"
const fibonacci_source_string = "func fibonacci(n int) int { if n <= 2 { return 1 }; return fibonacci(n-1) + fibonacci(n-2) }"
const shellsort_source_string = `
var shellshort_gaps = &[...]int{701, 301, 132, 57, 23, 10, 4, 1}

func shellsort(v []int) {
	var i, j, n, gi, gapn, gap, temp int
	n = len(v)
	gapn = len(shellshort_gaps) // fast interpreter currently lacks for-range
	for gi = 0; gi < gapn; gi++ {
		gap = shellshort_gaps[gi]
		for i = gap; i < n; i++ {
			temp = v[i]
			for j = i; j >= gap && v[j-gap] > temp; j -= gap {
				v[j] = v[j-gap]
			}
			v[j] = temp
		}
	}
}
`
const switch_source_string = `func bigswitch(n int) int {
	for i := 0; i < 1000; i++ {
		switch n&15 {
		case 0: n++
		case 1: n+=2
		case 2: n+=3
		case 3: n+=4
		case 4: n+=5
		case 5: n+=6
		case 6: n+=7
		case 7: n+=8
		case 8: n+=9
		case 9: n+=10
		case 10: n+=11
		case 11: n+=12
		case 12: n+=13
		case 13: n+=14
		case 14: n+=15
		case 15: n--
		}
	}
	return n
}`

var (
	cti = r.StructOf(
		[]r.StructField{
			r.StructField{Name: StrGensymInterface, Type: r.TypeOf((*interface{})(nil)).Elem()},
			r.StructField{Name: "String", Type: r.TypeOf((*func() string)(nil)).Elem()},
		},
	)
	fti = r.StructOf(
		[]r.StructField{
			r.StructField{Name: StrGensymInterface, Type: r.TypeOf((*xr.InterfaceHeader)(nil)).Elem()},
			r.StructField{Name: "String", Type: r.TypeOf((*func() string)(nil)).Elem()},
		},
	)

	csi = r.Zero(cti).Interface()
	fsi = r.Zero(r.PtrTo(fti)).Interface()

	zeroValues = []r.Value{}
)

var nil_map_int_string map[int]string

func for_range_string(s string) int32 {
	var v0 int32
	for i, r := range s {
		v0 += r << (uint8(i) * 8)
	}
	return v0
}

var testcases = []TestCase{
	TestCase{A, "1+1", "1+1", 1 + 1, nil},
	TestCase{A, "1+'A'", "1+'A'", 'B', nil}, // rune i.e. int32 should win over untyped constant (or int)
	TestCase{A, "int8+1", "int8(1)+1", int8(1) + 1, nil},
	TestCase{A, "int8_overflow", "int8(64)+64", int8(-128), nil},
	TestCase{A, "string", "\"foobar\"", "foobar", nil},
	TestCase{A, "expr_and", "3 & 6", 3 & 6, nil},
	TestCase{A, "expr_or", "7 | 8", 7 | 8, nil},
	TestCase{A, "expr_xor", "0x1f ^ 0xf1", 0x1f ^ 0xf1, nil},
	TestCase{A, "expr_arith", "((1+2)*3^4|99)%112", ((1+2)*3 ^ 4 | 99) % 112, nil},
	TestCase{A, "expr_shift", "7<<(10>>1)", 7 << (10 >> 1), nil},

	TestCase{A, "complex_1", "7i", 7i, nil},
	TestCase{A, "complex_2", "0.5+1.75i", 0.5 + 1.75i, nil},
	TestCase{A, "complex_3", "1i * 2i", 1i * 2i, nil},
	TestCase{A, "const_0", "const c0 rune = 'x'; c0", 'x', nil},
	TestCase{A, "const_1", "const c1 = 11; c1", 11, nil},
	TestCase{A, "const_2", "const c2 = 0xff&555+23/12.2; c2", 0xff&555 + 23/12.2, nil},

	// the classic interpreter is not accurate in this cases... missing exact arithmetic on constants
	TestCase{I, "const_3", "const c3 = 0.1+0.2; c3", float64(0.1) + float64(0.2), nil},
	TestCase{I, "const_4", "const c4 = c3/3; c4", (float64(0.1) + float64(0.2)) / 3, nil},

	// the fast interpreter instead *IS* accurate, thanks to exact arithmetic on untyped constants
	TestCase{F, "const_3", "const c3 = 0.1+0.2; c3", 0.1 + 0.2, nil},
	TestCase{F, "const_4", "const c4 = c3/3; c4", (0.1 + 0.2) / 3, nil},

	TestCase{F, "untyped_1", "2.0 >> 1", 1, nil},
	TestCase{A, "untyped_2", "1/2", 0, nil},
	TestCase{A, "untyped_unary", "-+^6", -+^6, nil},

	TestCase{A, "iota_1", "const c5 = iota^7; c5", 7, nil},
	TestCase{A, "iota_2", "const ( c6 = iota+6; c7=iota+6 ); c6", 6, nil},
	TestCase{A, "iota_3", "c7", 7, nil},
	TestCase{A, "iota_implicit_1", "const ( c8 uint = iota+8; c9 ); c8", uint(8), nil},
	TestCase{A, "iota_implicit_2", "c9", uint(9), nil},

	TestCase{A, "var_0", "var v0 int = 11; v0", 11, nil},
	TestCase{A, "var_1", "var v1 bool; v1", false, nil},
	TestCase{A, "var_2", "var v2 uint8 = 7; v2", uint8(7), nil},
	TestCase{A, "var_3", "var v3 uint16 = 12; v3", uint16(12), nil},
	TestCase{A, "var_4", "var v uint32 = 99; v", uint32(99), nil},
	TestCase{A, "var_5", "var v5 string; v5", "", nil},
	TestCase{A, "var_6", "var v6 float32; v6", float32(0), nil},
	TestCase{A, "var_7", "var v7 complex64; v7", complex64(0), nil},
	TestCase{A, "var_8", "var err error; err", nil, nil},
	TestCase{A, "var_9", `ve, vf := "", 1.23; ve`, "", nil},
	TestCase{A, "var_pointer", "var vp *string; vp", (*string)(nil), nil},
	TestCase{A, "var_map", "var vm *map[error]bool; vm", (*map[error]bool)(nil), nil},
	TestCase{A, "var_slice", "var vs []byte; vs", ([]byte)(nil), nil},
	TestCase{A, "var_array", "var va [2]rune; va", [2]rune{}, nil},
	TestCase{A, "var_interface_1", "var vi interface{} = 1; vi", 1, nil},
	TestCase{A, "var_interface_2", "var vnil interface{}; vnil", nil, nil},
	TestCase{A, "var_shift_1", "7 << 8", 7 << 8, nil},
	TestCase{A, "var_shift_2", "-9 >> 2", -9 >> 2, nil},
	TestCase{A, "var_shift_3", "v2 << 3", uint8(7) << 3, nil},
	TestCase{A, "var_shift_4", "v2 >> 1", uint8(7) >> 1, nil},
	TestCase{A, "var_shift_5", "0xff << v2", 0xff << 7, nil},
	TestCase{A, "var_shift_6", "0x12345678 >> v2", 0x12345678 >> uint8(7), nil},
	TestCase{A, "var_shift_7", "v << v2", uint32(99) << uint8(7), nil},
	TestCase{A, "var_shift_8", "v3 << v3 >> v2", uint16(12) << 12 >> uint8(7), nil},
	TestCase{A, "var_shift_9", "v3 << 0", uint16(12), nil},
	TestCase{A, "var_shift_overflow", "v3 << 13", uint16(32768), nil},

	// test division by constant power-of-two
	TestCase{I, "var_div_1", "v3 = 11; v3 / 2", uint64(11) / 2, nil}, // classic interpreter is not type-accurate here
	TestCase{I, "var_div_2", "v3 = 63; v3 / 8", uint64(63) / 8, nil},
	TestCase{F, "var_div_1", "v3 = 11; v3 / 2", uint16(11) / 2, nil},
	TestCase{F, "var_div_2", "v3 = 63; v3 / 8", uint16(63) / 8, nil},

	TestCase{A, "var_div_3", "v3 = 11; v3 /=2; v3", uint16(11) / 2, nil},
	TestCase{A, "var_div_4", "v3 = 63; v3 /=8; v3", uint16(63) / 8, nil},

	TestCase{A, "var_div_5", "v0 =+7; v0 /-1", -7, nil},
	TestCase{A, "var_div_6", "v0 =-7; v0 /+1", -7, nil},
	TestCase{A, "var_div_7", "v0 =-7; v0 /-1", +7, nil},
	TestCase{A, "var_div_8", "v0 =+11; v0 /-2", +11 / -2, nil},
	TestCase{A, "var_div_9", "v0 =-11; v0 /+2", -11 / +2, nil},
	TestCase{A, "var_div_10", "v0 =-11; v0 /-2", -11 / -2, nil},
	TestCase{A, "var_div_11", "v0 =+63; v0 /-8", +63 / -8, nil},
	TestCase{A, "var_div_12", "v0 =-63; v0 /+8", -63 / +8, nil},
	TestCase{A, "var_div_13", "v0 =-63; v0 /-8", -63 / -8, nil},

	// test remainder by constant power-of-two
	TestCase{I, "var_rem_1", "v3 = 17; v3 % 4", uint64(17) % 4, nil}, // classic interpreter is not type-accurate here
	TestCase{I, "var_rem_2", "v3 = 61; v3 % 8", uint64(61) % 8, nil},
	TestCase{F, "var_rem_1", "v3 = 17; v3 % 4", uint16(17) % 4, nil},
	TestCase{F, "var_rem_2", "v3 = 61; v3 % 8", uint16(61) % 8, nil},

	TestCase{A, "var_rem_3", "v0 =+17; v0 %-4", +17 % -4, nil},
	TestCase{A, "var_rem_4", "v0 =-17; v0 %+4", -17 % +4, nil},
	TestCase{A, "var_rem_5", "v0 =-17; v0 %-4", -17 % -4, nil},
	TestCase{A, "var_rem_6", "v0 =+61; v0 %-8", +61 % -8, nil},
	TestCase{A, "var_rem_7", "v0 =-61; v0 %+8", -61 % +8, nil},
	TestCase{A, "var_rem_8", "v0 =-61; v0 %-8", -61 % -8, nil},

	TestCase{A, "eql_nil_1", "err == nil", true, nil},
	TestCase{A, "eql_nil_2", "vp == nil", true, nil},
	TestCase{A, "eql_nil_3", "vm == nil", true, nil},
	TestCase{A, "eql_nil_4", "vs == nil", true, nil},
	TestCase{A, "eql_nil_5", "vi == nil", false, nil},
	TestCase{A, "eql_nil_6", "vnil == nil", true, nil},
	TestCase{A, "eql_halfnil_1", "var vhalfnil interface{} = vm; vhalfnil == nil", false, nil},
	TestCase{A, "eql_halfnil_2", "vm = nil; vhalfnil = vm; vhalfnil == nil", false, nil},
	TestCase{A, "eql_interface", "vi == 1", true, nil},

	TestCase{A, "typed_unary_1", "!!!v1", true, nil},
	TestCase{A, "typed_unary_2", "+-^v2", uint8(8), nil},
	TestCase{A, "typed_unary_3", "v3 = 12; +^-v3", uint16(11), nil},
	TestCase{A, "typed_unary_4", "v7 = 2.5i; -v7", complex64(-2.5i), nil},

	TestCase{A, "type_int8", "type t8 int8; var v8 t8; v8", int8(0), nil},
	TestCase{A, "type_complicated", "type tfff func(int,int) func(error, func(bool)) string; var vfff tfff; vfff", (func(int, int) func(error, func(bool)) string)(nil), nil},
	TestCase{I, "type_interface", "type Stringer interface { String() string }; var s Stringer; s", csi, nil},
	TestCase{F, "type_interface", "type Stringer interface { String() string }; var s Stringer; s", fsi, nil},
	TestCase{A, "type_struct_1", "type Pair struct { A rune; B string }; var pair Pair; pair", struct {
		A rune
		B string
	}{}, nil},
	TestCase{A, "type_struct_2", "type Triple struct { Pair; C float32 }; var triple Triple; triple.C", float32(0), nil},
	TestCase{A, "field_get_1", "pair.A", rune(0), nil},
	TestCase{A, "field_get_2", "pair.B", "", nil},
	TestCase{F, "field_anonymous_1", "triple.Pair", struct {
		A rune
		B string
	}{}, nil},
	TestCase{F, "field_embedded_1", "triple.A", rune(0), nil},
	TestCase{F, "field_embedded_2", "triple.Pair.B", "", nil},

	TestCase{A, "address_0", "var vf = 1.25; *&vf == vf", true, nil},
	TestCase{A, "address_1", "var pvf = &vf; *pvf", 1.25, nil},
	TestCase{A, "address_2", "&*pvf == *&pvf", true, nil},
	TestCase{A, "address_3", "var pvs = &vs; v1 = (*pvs == nil); v1", true, nil},

	TestCase{A, "make_chan", "cx := make(chan interface{}, 2); cx", make(chan interface{}, 2), nil},
	TestCase{A, "make_map", "m := make(map[int]string); m", make(map[int]string), nil},
	TestCase{A, "make_slice", "y := make([]uint8, 7); y[0] = 100; y[3] = 103; y", []uint8{100, 0, 0, 103, 0, 0, 0}, nil},
	TestCase{A, "expr_index_string_1", `"abc"[2]`, byte('c'), nil},
	TestCase{A, "expr_index_string_2", `v5 = "foo"; v0 = 0; v5[v0]`, byte('f'), nil},
	TestCase{A, "expr_index_array_1", `va[1]`, rune(0), nil},
	TestCase{A, "expr_index_array_2", `(&va)[0]`, rune(0), nil},
	TestCase{A, "expr_index_map", `var m2 map[rune]string; m2['x']`, nil, []interface{}{"", false}},
	TestCase{A, "expr_slice_0", "y[:]", []uint8{100, 0, 0, 103, 0, 0, 0}, nil},
	TestCase{A, "expr_slice_1", "y[1:]", []uint8{0, 0, 103, 0, 0, 0}, nil},
	TestCase{A, "expr_slice_2", "y[2:4]", []uint8{0, 103}, nil},
	TestCase{A, "expr_slice_3", "y[:3]", []uint8{100, 0, 0}, nil},
	TestCase{A, "expr_slice_4", "y = y[:3:7]; y", []uint8{100, 0, 0}, nil},
	TestCase{A, "expr_slice_5", `"abc"[:]`, "abc", nil},
	TestCase{A, "expr_slice_6", `"abc"[1:]`, "bc", nil},
	TestCase{A, "expr_slice_7", `"abc"[1:2]`, "b", nil},

	TestCase{A, "set_const_1", "v1 = true;    v1", true, nil},
	TestCase{A, "set_const_2", "v2 = 9;       v2", uint8(9), nil},
	TestCase{A, "set_const_3", "v3 = 60000;   v3", uint16(60000), nil},
	TestCase{A, "set_const_4", "v  = 987;      v", uint32(987), nil},
	TestCase{A, "set_const_5", `v5 = "8y57r"; v5`, "8y57r", nil},
	TestCase{A, "set_const_6", "v6 = 0.12345678901234; v6", float32(0.12345678901234), nil},        // v6 is declared float32
	TestCase{A, "set_const_7", "v7 = 0.98765432109i; v7", complex(0, float32(0.98765432109)), nil}, // v7 is declared complex64

	TestCase{A, "set_expr_1", "v1 = v1 == v1;    v1", true, nil},
	TestCase{A, "set_expr_2", "v2 -= 7;      v2", uint8(2), nil},
	TestCase{A, "set_expr_3", "v3 %= 7;      v3", uint16(60000) % 7, nil},
	TestCase{A, "set_expr_4", "v  = v * 10;      v", uint32(9870), nil},
	TestCase{A, "set_expr_5", `v5 = v5 + "iuh";  v5`, "8y57riuh", nil},
	TestCase{A, "set_expr_6", "v6 = 1/v6;        v6", 1 / float32(0.12345678901234), nil},                              // v6 is declared float32
	TestCase{A, "set_expr_7", "v7 = v7 * v7;     v7", complex(-float32(0.98765432109)*float32(0.98765432109), 0), nil}, // v7 is declared complex64

	TestCase{A, "add_2", "v2 += 255;    v2", uint8(1), nil}, // overflow
	TestCase{A, "add_3", "v3 += 536;    v3", uint16(60000)%7 + 536, nil},
	TestCase{A, "add_4", "v  += 111;     v", uint32(9870 + 111), nil},
	TestCase{A, "add_5", `v5 += "@#$";  v5`, "8y57riuh@#$", nil},
	TestCase{A, "add_6", "v6 += 0.975319; v6", 1/float32(0.12345678901234) + float32(0.975319), nil}, // v6 is declared float32
	TestCase{A, "add_7", "v7 = 1; v7 += 0.999999i; v7", complex(float32(1), float32(0.999999)), nil}, // v7 is declared complex64

	TestCase{A, "mul_1", "v2 = 4;  v2 *= 3; v2", uint8(12), nil},
	TestCase{A, "rem_1", "v3 = 12; v3 %= 7; v3", uint16(5), nil},
	TestCase{A, "and_1", "v3 &= 9;          v3", uint16(1), nil},

	TestCase{A, "shift_1", "v3 <<= 7;       v3", uint16(128), nil},
	TestCase{A, "shift_2", "v3 >>= (v3>>5); v3", uint16(128 >> (128 >> 5)), nil},

	TestCase{A, "if_1", "v2 = 1; if v2 < 1 { v2 = v2-1 } else { v2 = v2+1 }; v2", uint8(2), nil},
	TestCase{A, "if_2", "if v2 < 5 { v2 = v2+2 } else { v2 = v2-2 }; v2", uint8(4), nil},
	TestCase{A, "for_1", "var i, j, k int; for i=1; i<=2; i=i+1 { if i<2 {j=i} else {k=i} }; i", 3, nil},
	TestCase{A, "for_2", "j", 1, nil},
	TestCase{A, "for_3", "k", 2, nil},
	TestCase{A, "continue_1", "j=0; k=0; for i:=1; i<=7; i=i+1 { if i==3 {j=i; continue}; k=k+i }; j", 3, nil},
	TestCase{A, "continue_2", "k", 25, nil},
	TestCase{A, "continue_3", "j=0; k=0; for i:=1; i<=7; i=i+1 { var ii = i; if ii==3 {j=ii; continue}; k=k+ii }; j", 3, nil},
	TestCase{A, "continue_4", "k", 25, nil},

	TestCase{A, "for_range_array", `v0 = 0; for _, s := range [2]string{"a", "bc"} { v0 += len(s) }; v0`, 3, nil},
	TestCase{A, "for_range_chan", `v0 = 0; c := make(chan int, 2); c <- 1; c <- 2; close(c); for e := range c { v0 += e }; v0`, 3, nil},
	TestCase{A, "for_range_map", `var vrune rune; m2 = map[rune]string{'x':"x", 'y':"y", 'z':"z"}; for k,v := range m2 { vrune += k + rune(v[0]) }; vrune`, ('x' + 'y' + 'z') * 2, nil},
	TestCase{A, "for_range_slice", `v0 = 0; for _, s := range [ ]string{"a", "bc"} { v0 += len(s) }; v0`, 3, nil},
	TestCase{A, "for_range_string", `vrune = 0; for i, r := range "abc\u00ff" { vrune += r << (uint8(i)*8) }; vrune`, for_range_string("abc\u00ff"), nil},

	TestCase{A, "function_0", "func nop() { }; nop()", nil, []interface{}{}},
	TestCase{A, "function_1", "func seven() int { return 7 }; seven()", 7, nil},
	TestCase{A, "function_2", "i=0; func seti(ii int) { i=ii }; seti(-493); i", -493, nil},
	TestCase{A, "function_3", "func ident(x uint) uint { return x }; ident(42)", uint(42), nil},
	TestCase{A, "function_4", "func equal(x, y float32) bool { return x == y }; equal(1.1, 1.1)", true, nil},
	TestCase{A, "function_5", "func swap(x, y int) (int,int) { return y, x }; swap(88,99)", nil, []interface{}{99, 88}},
	TestCase{A, "function_6", "i=0; func seti2() { i=2 }; seti2(); i", 2, nil},
	TestCase{A, "function_7", "i=0; func setiadd(x, y int) { i=x+y }; setiadd(7,8); i", 15, nil},

	// bool is the param name, NOT its type!
	TestCase{F, "function_named_return", "func add3(bool, x, y int) (z int) { z=bool+x+y; return }; add3(1,2,3)", 6, nil},

	TestCase{A, "function_variadic_1", "func list_args(args ...interface{}) []interface{} { return args }; list_args(1,2,3)", []interface{}{1, 2, 3}, nil},
	TestCase{A, "function_variadic_2", "si := make([]interface{}, 4); si[1]=1; si[2]=2; si[3]=3; list_args(si...)", []interface{}{nil, 1, 2, 3}, nil},
	TestCase{A, "fibonacci", fibonacci_source_string + "; fibonacci(13)", 233, nil},
	TestCase{A, "function_literal", "adder := func(a,b int) int { return a+b }; adder(-7,-9)", -16, nil},

	TestCase{A, "closure_1", `
		func test_closure_1() int {
			var x int
			func() {
				x = 1
			}()
			return x
		}
		test_closure_1()`, 1, nil},

	TestCase{F, "closure_2", `
		func test_closure_2() (x int) {
			func() {
				x = 2
			}()
		}
		test_closure_2()`, 2, nil},

	TestCase{A, "setvar_deref_1", `vstr := "foo"; pvstr := &vstr; *pvstr = "bar"; vstr`, "bar", nil},
	TestCase{A, "setvar_deref_2", `vint := 5; pvint := &vint; *pvint = 6; vint`, 6, nil},
	TestCase{A, "setplace_deref_1", `func vstr_addr() *string { return &vstr }; *vstr_addr() = "qwerty"; vstr`, "qwerty", nil},
	TestCase{A, "setplace_deref_2", `*vstr_addr() += "uiop"; vstr`, "qwertyuiop", nil},
	TestCase{A, "setplace_deref_3", `func vint_addr() *int { return &vint }; *vint_addr() = 7; vint`, 7, nil},
	TestCase{A, "setplace_deref_4", `*vint_addr() %= 4; vint`, 3, nil},

	TestCase{A, "setmap_1", `m[1]="x"; m[2]="y"; m`, map[int]string{1: "x", 2: "y"}, nil},
	TestCase{A, "setmap_2", `m[2]+="z"; m`, map[int]string{1: "x", 2: "yz"}, nil},
	TestCase{A, "setmap_3", `mi := make(map[rune]byte); mi['@']+=2; mi`, map[rune]byte{'@': 2}, nil},
	TestCase{A, "setmap_4", `mi['a'] |= 7; mi`, map[rune]byte{'@': 2, 'a': 7}, nil},
	TestCase{A, "getmap_1", `m[1]`, nil, []interface{}{"x", true}},
	TestCase{A, "getmap_2", `m1 := m[1]; m1`, "x", nil},
	TestCase{A, "getmap_3", `mi['b']`, nil, []interface{}{byte(0), false}},
	TestCase{A, "getmap_4", `v2 = mi['@']; v2`, byte(2), nil},

	TestCase{A, "divmap_1", "mi['@'] = 99; mi['@'] /= 3; v2 = mi['@']; v2", byte(33), nil},
	TestCase{A, "divmap_2", "mi['@'] /= 4; v2 = mi['@']; v2", byte(8), nil},

	TestCase{A, "swap_assignment", `i=1;j=2;  i,j=j,i;  list_args(i, j)`, []interface{}{2, 1}, nil},
	TestCase{A, "evil_assignment_1", `i=0; si[0]=7; si[1]=8
		i, si[i] = 1, 2
		list_args(i,si[0],si[1])`, []interface{}{1, 2, 8}, nil},
	TestCase{A, "evil_assignment_2", `i=0; m=make(map[int]string); mcopy:=m;
		i, m, m[i] = 1, nil, "foo"
		list_args(i,m,mcopy)`,
		[]interface{}{1, nil_map_int_string, map[int]string{0: "foo"}}, nil},

	TestCase{A, "field_set_1", `pair.A = 'k'; pair.B = "m"; pair`, struct {
		A rune
		B string
	}{'k', "m"}, nil},
	TestCase{A, "field_set_2", `pair.A, pair.B = 'x', "y"; pair`, struct {
		A rune
		B string
	}{'x', "y"}, nil},
	TestCase{F, "field_set_3", `triple.Pair.A, triple.C = 'a', 1.0; triple.Pair`, struct {
		A rune
		B string
	}{'a', ""}, nil},
	TestCase{F, "field_set_embedded_1", `triple.A, triple.B = 'b', "xy"; triple.Pair`, struct {
		A rune
		B string
	}{'b', "xy"}, nil},
	TestCase{F, "field_addr_1", "ppair := &triple.Pair; ppair.A", 'b', nil},
	TestCase{F, "field_addr_2", "ppair.A++; triple.Pair.A", 'c', nil},

	TestCase{I, "import", `import ( "fmt"; "time" )`, "time", nil},
	TestCase{F, "import", `import ( "fmt"; "time" )`, nil, []interface{}{}},

	TestCase{A, "goroutine_1", `go seti(9); time.Sleep(time.Second/50); i`, 9, nil},

	TestCase{A, "builtin_append", "append(vs,0,1,2)", []byte{0, 1, 2}, nil},
	TestCase{A, "builtin_cap", "cap(va)", 2, nil},
	TestCase{A, "builtin_len", "len(v5)", len("8y57riuh@#$"), nil},
	TestCase{A, "builtin_new", "new(int)", new(int), nil},
	TestCase{A, "builtin_make_1", "make(map[int]int)", make(map[int]int), nil},
	TestCase{A, "builtin_make_2", "make(map[int]int, 10)", make(map[int]int), nil}, // capacity is ignored
	TestCase{A, "builtin_make_4", "make([]*error, 2)", make([]*error, 2), nil},
	TestCase{A, "builtin_make_5", "make([]rune, 3, 4)", make([]rune, 3, 4), nil},
	TestCase{A, "builtin_make_6", "make(chan byte)", make(chan byte), nil},
	TestCase{A, "builtin_make_7", "make(chan byte, 2)", make(chan byte, 2), nil},
	TestCase{A, "builtin_make_8", "vs = make([]byte, 5); vs", make([]byte, 5), nil},
	TestCase{A, "builtin_copy_1", "copy(vs, v5)", 5, nil},
	TestCase{A, "builtin_copy_2", "vs", []byte("8y57r"), nil},
	TestCase{A, "builtin_delete_1", "delete(mi,64); mi", map[rune]byte{'a': 7}, nil},
	TestCase{A, "builtin_real_1", "real(0.5+1.75i)", real(0.5 + 1.75i), nil},
	TestCase{A, "builtin_real_2", "var cplx complex64 = 1.5+0.25i; real(cplx)", real(complex64(1.5 + 0.25i)), nil},
	TestCase{A, "builtin_imag_1", "imag(0.5+1.75i)", imag(0.5 + 1.75i), nil},
	TestCase{A, "builtin_imag_2", "imag(cplx)", imag(complex64(1.5 + 0.25i)), nil},
	TestCase{A, "builtin_complex_1", "complex(0,1)", complex(0, 1), nil},
	TestCase{A, "builtin_complex_2", "v6 = 0.1; complex(v6,-v6)", complex(float32(0.1), -float32(0.1)), nil},

	TestCase{A, "time_duration_0", `var td time.Duration = 1; td`, time.Duration(1), nil},
	TestCase{A, "time_duration_1", `- td`, time.Duration(-1), nil},
	TestCase{A, "time_duration_2", `td + 1`, time.Duration(2), nil},
	TestCase{A, "time_duration_3", `4 - td`, time.Duration(3), nil},
	TestCase{A, "time_duration_4", `td * 4`, time.Duration(4), nil},
	TestCase{A, "time_duration_5", `5 / td`, time.Duration(5), nil},
	TestCase{A, "time_duration_6", `&td`, func() *time.Duration { td := time.Duration(1); return &td }(), nil},
	TestCase{A, "time_duration_method", ` td = time.Second; td.String()`, "1s", nil},
	TestCase{A, "time_duration_closure", ` tds := td.String; tds()`, "1s", nil},
	TestCase{A, "time_utc", ` utc := time.UTC; utc.String()`, "UTC", nil},
	TestCase{A, "time_utc_addr", ` utcaddr := &time.UTC; *utcaddr == utc`, true, nil},
	TestCase{A, "time_utc_set_1", ` time.UTC = nil; time.UTC == nil`, true, nil},
	TestCase{A, "time_utc_set_2", ` time.UTC = utc; time.UTC.String()`, "UTC", nil},

	TestCase{A, "literal_array", "[3]int{1,2:3}", [3]int{1, 0, 3}, nil},
	TestCase{A, "literal_array_address", "&[...]int{3:4,5:6}", &[...]int{3: 4, 5: 6}, nil},
	TestCase{A, "literal_map", `map[int]string{1: "foo", 2: "bar"}`, map[int]string{1: "foo", 2: "bar"}, nil},
	TestCase{A, "literal_map_address", `&map[int]byte{6:7, 8:9}`, &map[int]byte{6: 7, 8: 9}, nil},
	TestCase{A, "literal_slice", "[]rune{'a','b','c'}", []rune{'a', 'b', 'c'}, nil},
	TestCase{A, "literal_slice_address", "&[]rune{'x','y'}", &[]rune{'x', 'y'}, nil},
	TestCase{A, "literal_struct", `Pair{A: 0x73, B: "\x94"}`, struct {
		A rune
		B string
	}{A: 0x73, B: "\x94"}, nil},
	TestCase{A, "literal_struct_address", `&Pair{1,"2"}`, &struct {
		A rune
		B string
	}{A: 1, B: "2"}, nil},

	TestCase{A, "method_decl_1", `func (p *Pair) SetA(a rune) { p.A = a }; func (p Pair) SetAV(a rune) { p.A = a }; nil`, nil, nil},
	TestCase{A, "method_decl_2", `func (p Pair) SetAV(a rune) { p.A = a }; nil`, nil, nil},
	TestCase{A, "method_decl_3", `func (p Pair) String() string { return fmt.Sprintf("%c %s", p.A, p.B) }; nil`, nil, nil},

	TestCase{A, "method_on_ptr", `pair.SetA(33); pair.A`, rune(33), nil},
	TestCase{A, "method_on_val_1", `pair.SetAV(11); pair.A`, rune(33), nil}, // method on value gets a copy of the receiver - changes to not propagate
	TestCase{A, "method_on_val_2", `pair.String()`, "! y", nil},
	TestCase{F, "method_embedded_on_ptr", `triple.SetA('+'); triple.A`, '+', nil},
	TestCase{F, "method_embedded_on_val", `triple.SetAV('*'); triple.A`, '+', nil},

	TestCase{A, "interface_1", "var st fmt.Stringer = time.Second; st", time.Second, nil},
	TestCase{A, "interface_method_1", "bind := st.String; bind()", "1s", nil},
	TestCase{F, "interface_2", "st = pair; nil", nil, nil},
	TestCase{F, "interface_method_2", "bind = st.String; bind()", "! y", nil},

	TestCase{F, "concrete_method_to_func", "f1 := time.Duration.Seconds; f1(time.Hour)", 3600.0, nil},
	TestCase{F, "interface_method_to_func", "f2 := fmt.Stringer.String; f2(time.Hour)", "1h0m0s", nil},

	TestCase{A, "multiple_values_1", "func twins(x float32) (float32,float32) { return x, x+1 }; twins(17.0)", nil, []interface{}{float32(17.0), float32(18.0)}},
	TestCase{A, "multiple_values_2", "func twins2(x float32) (float32,float32) { return twins(x) }; twins2(19.0)", nil, []interface{}{float32(19.0), float32(20.0)}},
	TestCase{A, "multiple_values_3", "f1, f2 := twins(23.0); list_args(f1, f2)", []interface{}{float32(23.0), float32(24.0)}, nil},
	TestCase{A, "multiple_values_4", "fm := make(map[int]float32); fm[1], fm[2] = twins(3.0); fm", map[int]float32{1: 3.0, 2: 4.0}, nil},
	TestCase{A, "multiple_values_5", "swap(swap(3,4))", nil, []interface{}{3, 4}},

	TestCase{A, "pred_bool_1", "false==false && true==true && true!=false", true, nil},
	TestCase{A, "pred_bool_2", "false!=false || true!=true || true==false", false, nil},
	TestCase{A, "pred_int", "1==1 && 1<=1 && 1>=1 && 1!=2 && 1<2 && 2>1 || 0==1", true, nil},
	TestCase{A, "pred_string_1", `""=="" && "">="" && ""<="" && ""<"a" && ""<="b" && "a">"" && "b">=""`, true, nil},
	TestCase{A, "pred_string_2", `ve=="" && ve>="" && ve<="" && ve<"a" && ve<="b" && "a">ve && "b">=ve`, true, nil},
	TestCase{A, "pred_string_3", `"x"=="x" && "x"<="x" && "x">="x" && "x"!="y" && "x"<"y" && "y">"x"`, true, nil},
	TestCase{A, "pred_string_4", `"x"!="x" || "y"!="y" || "x">="y" || "y"<="x"`, false, nil},

	TestCase{A, "defer_1", `
		vi = nil
		func test_defer_1() {
			defer func() {
				vi = 1
			}()
		}
		test_defer_1(); vi`, 1, nil},
	TestCase{F, "defer_2", `
		func test_defer_2() (x int) {
			defer func() {
				x = 2
			}()
		}
		test_defer_2()`, 2, nil},
	TestCase{A, "defer_3", "v = 0; func testdefer(x uint32) { if x != 0 { defer func() { v = x }() } }; testdefer(29); v", uint32(29), nil},
	TestCase{A, "defer_4", "v = 12; testdefer(0); v", uint32(12), nil},
	TestCase{A, "recover_1", `var vpanic interface{}
		func test_recover(rec bool, panick interface{}) {
			defer func() {
				if rec {
					vpanic = recover()
				}
			}()
			panic(panick)
		}
		test_recover(true, -1)
		vpanic`, -1, nil},
	TestCase{A, "recover_2", `var vpanic2, vpanic3 interface{}
		func test_nested_recover(repanic bool, panick interface{}) {
			defer func() {
				vpanic = recover()
			}()
			defer func() {
				func() {
					vpanic3 = recover()
				}()
				vpanic2 = recover()
				if repanic {
					panic(vpanic2)
				}
			}()
			panic(panick)
		}
		test_nested_recover(false, -2)
		list_args(vpanic, vpanic2, vpanic3)
		`, []interface{}{nil, -2, nil}, nil},
	TestCase{A, "recover_3", `vpanic, vpanic2, vpanic3 = nil, nil, nil
		test_nested_recover(true, -3)
		list_args(vpanic, vpanic2, vpanic3)
		`, []interface{}{-3, -3, nil}, nil},
	TestCase{A, "recover_4", `vpanic = nil
		func test_defer_panic(panick interface{}) {
			defer func() {
				vpanic = recover()
			}()
			defer func() {
				panic(panick)
			}()
		}
		test_defer_panic(-4)
		vpanic
		`, -4, nil},
	TestCase{A, "send_recv", `cx <- "x"; <-cx`, nil, []interface{}{"x", true}},
	TestCase{A, "sum", sum_source_string + "; sum(100)", 5050, nil},

	TestCase{A, "select_1", "vi = nil; cx <- 1; { var x interface{}; select { case x=<-cx: vi=x; default: } }; vi", 1, nil},
	TestCase{A, "select_2", "vi = nil; cx <- map[int]int{1:2}; select { case x:=<-cx: vi=x; default: }; vi", map[int]int{1: 2}, nil},
	TestCase{A, "select_3", "vi = nil; select { case cx<-1: vi=1; default: vi=0 }; vi", 1, nil},
	TestCase{A, "select_4", "vi = nil; select { case cx<-2: vi=2; default: vi=0 }; vi", 2, nil},
	TestCase{A, "select_5", "vi = nil; select { case cx<-3: vi=3; default: vi=0 }; vi", 0, nil},
	TestCase{A, "select_6", "vi = nil; select { case cx<-4: vi=4; case x:=<-cx: vi=x; default: vi=0 }; vi", 1, nil},

	TestCase{A, "switch_1", "vi=nil; switch { case false: ; default: vi='1' }; vi", '1', nil},
	TestCase{A, "switch_2", "vi=nil; switch v:=20; v { case 20: vi='2'; vi='3' }; vi", '3', nil},
	TestCase{A, "switch_3", `v3=0; vi=nil
		func inc_u16(addr *uint16, n uint16) uint16 { *addr += n; return *addr }
		switch v3++; inc_u16(&v3, 2) { case 1: ; case 2: ; case 3: vi='3'; default: }; vi`, '3', nil},
	TestCase{A, "switch_4", "v0=7; switch v:=7; v { default: vi=0; case 1: vi=1; case v0: vi=7; case 7: vi=-7 }; vi", 7, nil},
	TestCase{A, "switch_fallthrough", `vi=nil; switch v:=0; v {
		default:       fallthrough
		case 1: vi=10; fallthrough
		case 2: vi=20; break
		case 3: vi=30
	}; vi`, 20, nil},
	TestCase{A | S, "switch_multithread", `func doswitch(i, j int) { switch i { case 1: v0 = j; case 2: vi = j } }
		v0, vi = 0, nil
		go doswitch(1, 10)
		doswitch(2, 20)
		time.Sleep(time.Second/10)
		list_args(v0, vi)
	`, []interface{}{10, 20}, nil},

	TestCase{A, "typeswitch_1", `vi = nil; var x interface{} = "abc"; switch y := x.(type) { default: vi = 0; case string: vi = len(y) }; vi`, 3, nil},
	TestCase{A, "typeswitch_2", `vi = nil; switch x.(type) { default: vi = 0; case byte, bool: vi = 1; case interface{}: vi = 2 }; vi`, 2, nil},
	TestCase{A, "typeswitch_3", `vi = nil; switch x.(type) { default: vi = 0; case int:         vi = 3 }; vi`, 0, nil},
	TestCase{A, "typeswitch_4", `vi = nil; switch x.(type) { default: vi = 0; case string:      vi = 4 }; vi`, 4, nil},
	TestCase{A, "typeswitch_5", `vi,x = nil,nil; switch x.(type) { default: vi=0; case nil:     vi = 5 }; vi`, 5, nil},
	TestCase{A, "typeswitch_6", `var stringer fmt.Stringer = time.Minute
		switch s := stringer.(type) {
			case nil:           vi = 0
			default:            vi = 1
			case time.Duration: vi = 6
			case fmt.Stringer:  vi = 7
		}; vi`, 6, nil},

	TestCase{A, "typeassert_1", `var xi interface{} = "abc"; yi := xi.(string); yi`, "abc", nil},
	TestCase{A, "typeassert_2", `xi.(string)`, nil, []interface{}{"abc", true}},
	TestCase{A, "typeassert_3", `xi.(int)`, nil, []interface{}{0, false}},
	TestCase{A, "typeassert_4", `xi = nil; xi.(error)`, nil, []interface{}{error(nil), false}},

	TestCase{A, "quote_1", `~quote{7}`, &ast.BasicLit{Kind: token.INT, Value: "7"}, nil},
	TestCase{A, "quote_2", `~quote{x}`, &ast.Ident{Name: "x"}, nil},
	TestCase{A, "quote_3", `var ab = ~quote{a;b}; ab`, &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.Ident{Name: "a"}},
		&ast.ExprStmt{X: &ast.Ident{Name: "b"}},
	}}, nil},
	TestCase{A, "quote_4", `~'{"foo"+"bar"}`, &ast.BinaryExpr{
		Op: token.ADD,
		X:  &ast.BasicLit{Kind: token.STRING, Value: `"foo"`},
		Y:  &ast.BasicLit{Kind: token.STRING, Value: `"bar"`},
	}, nil},
	TestCase{A, "quasiquote_1", `~quasiquote{1 + ~unquote{2+3}}`, &ast.BinaryExpr{
		Op: token.ADD,
		X:  &ast.BasicLit{Kind: token.INT, Value: "1"},
		Y:  &ast.BasicLit{Kind: token.INT, Value: "5"},
	}, nil},
	TestCase{A, "quasiquote_2", `~"{2 * ~,{3<<1}}`, &ast.BinaryExpr{
		Op: token.MUL,
		X:  &ast.BasicLit{Kind: token.INT, Value: "2"},
		Y:  &ast.BasicLit{Kind: token.INT, Value: "6"},
	}, nil},
	TestCase{A, "quasiquote_3", `~"{func(int) {}}`, &ast.FuncLit{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					&ast.Field{
						Names: nil,
						Type: &ast.Ident{
							Name: "int",
						},
					},
				},
			},
		},
		Body: &ast.BlockStmt{},
	}, nil},
	TestCase{A, "quasiquote_func", `func qq_func(x interface{}) interface{} { y := ~"~,x; return y }; qq_func(123)`,
		&ast.BasicLit{Kind: token.INT, Value: "123"}, nil},
	TestCase{A, "quasiquote_case", `~"{case xy: nil}`,
		&ast.CaseClause{
			List: []ast.Expr{&ast.Ident{Name: "xy"}},
			Body: []ast.Stmt{&ast.ExprStmt{&ast.Ident{Name: "nil"}}},
		}, nil,
	},

	TestCase{A, "unquote_splice_1", `~quasiquote{~unquote_splice ab ; c}`, &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.Ident{Name: "a"}},
		&ast.ExprStmt{X: &ast.Ident{Name: "b"}},
		&ast.ExprStmt{X: &ast.Ident{Name: "c"}},
	}}, nil},
	TestCase{A, "unquote_splice_2", `~"{zero ; ~,@ab ; one}`, &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.Ident{Name: "zero"}},
		&ast.ExprStmt{X: &ast.Ident{Name: "a"}},
		&ast.ExprStmt{X: &ast.Ident{Name: "b"}},
		&ast.ExprStmt{X: &ast.Ident{Name: "one"}},
	}}, nil},
	TestCase{A, "macro", "~macro second_arg(a,b,c interface{}) interface{} { return b }; v = 98; v", uint32(98), nil},
	TestCase{A, "macro_call", "second_arg;1;v;3", uint32(98), nil},
	TestCase{A, "macro_nested", "second_arg;1;{second_arg;2;3;4};5", 3, nil},
	TestCase{I, "values", "Values(3,4,5)", nil, []interface{}{3, 4, 5}},
	TestCase{A, "eval", "Eval(~quote{1+2})", 3, nil},
	TestCase{I, "eval_quote", "Eval(~quote{Values(3,4,5)})", nil, []interface{}{3, 4, 5}},
}

func (c *TestCase) compareResults(t *testing.T, actual []r.Value) {
	expected := c.results
	if expected == nil {
		expected = []interface{}{c.result0}
	}
	if len(actual) != len(expected) {
		c.fail(t, actual, expected)
		return
	}
	for i := range actual {
		c.compareResult(t, actual[i], expected[i])
	}
}

func (c *TestCase) compareResult(t *testing.T, actualv r.Value, expected interface{}) {
	if actualv == Nil || actualv == None {
		if expected != nil {
			c.fail(t, nil, expected)
		}
		return
	}
	actual := actualv.Interface()
	if !r.DeepEqual(actual, expected) {
		if r.TypeOf(actual) == r.TypeOf(expected) {
			if actualNode, ok := actual.(ast.Node); ok {
				if expectedNode, ok := expected.(ast.Node); ok {
					c.compareAst(t, ToAst(actualNode), ToAst(expectedNode))
					return
				}
			} else if actualv.Kind() == r.Chan {
				// for channels just check the type, length and capacity
				expectedv := r.ValueOf(expected)
				if actualv.Len() == expectedv.Len() && actualv.Cap() == expectedv.Cap() {
					return
				}
			}
		}
		c.fail(t, actual, expected)
	}
}

func (c *TestCase) compareAst(t *testing.T, actual Ast, expected Ast) {
	if actual == nil || expected == nil {
		if actual != nil || expected != nil {
			c.fail(t, actual, expected)
		}
		return
	}
	if r.TypeOf(actual) == r.TypeOf(expected) {
		switch actual := actual.(type) {
		case BadDecl, BadExpr, BadStmt:
			return
		case Ident:
			if actual.X.Name == expected.(Ident).X.Name {
				return
			}
		case BasicLit:
			actualp := actual.X
			expectedp := expected.(BasicLit).X
			if actualp == expectedp || (actualp != nil && expectedp != nil && actualp.Kind == expectedp.Kind && actualp.Value == expectedp.Value) {
				return
			}
		default:
			na := actual.Size()
			ne := expected.Size()
			if actual.Op() == expected.Op() && na == ne {
				for i := 0; i < na; i++ {
					c.compareAst(t, actual.Get(i), expected.Get(i))
				}
				return
			}
		}
	}
	c.fail(t, actual, expected)
}

func (c *TestCase) fail(t *testing.T, actual interface{}, expected interface{}) {
	t.Errorf("expecting %v <%T>, found %v <%T>\n", expected, expected, actual, actual)
}
