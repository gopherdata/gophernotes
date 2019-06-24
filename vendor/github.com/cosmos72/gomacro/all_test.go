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
 * all_test.go
 *
 *  Created on: Mar 06 2017
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"go/ast"
	"go/constant"
	"go/token"
	"math/big"
	r "reflect"
	"sync"
	"testing"
	"time"

	. "github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/reflect"
	"github.com/cosmos72/gomacro/base/untyped"
	"github.com/cosmos72/gomacro/classic"
	"github.com/cosmos72/gomacro/fast"
	"github.com/cosmos72/gomacro/go/etoken"
	"github.com/cosmos72/gomacro/go/parser"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type TestFor int

const (
	S  TestFor = 1 << iota // set option OptDebugSleepOnSwitch
	C                      // test for classic interpreter
	F                      // test for fast interpreter
	G1                     // test requires generics v1 (C++-style)
	G2                     // test requires generics v2 "contracts are interfaces"false
	U                      // test returns untyped constant (relevant only for fast interpreter)
	Z                      // temporary override: run only these tests, on fast interpreter only
	A  = C | F             // test for both interpreters
	G  = G1 | G2
)

type TestCase struct {
	testfor TestFor
	name    string
	program string
	result0 interface{}
	results []interface{}
}

func (tc *TestCase) shouldRun(interp TestFor) bool {
	if tc.testfor&interp == 0 {
		return false
	}
	if tc.testfor&G1 != 0 && etoken.GENERICS_V1_CXX {
		return true
	}
	if tc.testfor&G2 != 0 && etoken.GENERICS_V2_CTI {
		return true
	}
	return tc.testfor&(G1|G2) == 0
}

var foundZ bool

func init() {
	for i := range testcases {
		if testcases[i].testfor&Z != 0 {
			foundZ = true
		}
	}
}

func TestClassic(t *testing.T) {
	if foundZ {
		t.Skip("one or more tests marked with 'Z' i.e. run only those and only on fast interpreter")
	}
	ir := classic.New()
	// ir.Options |= OptDebugCallStack | OptDebugPanicRecover
	for i := range testcases {
		test := &testcases[i]
		if test.shouldRun(C) {
			t.Run(test.name, func(t *testing.T) { test.classic(t, ir) })
		}
	}
}

func TestFast(t *testing.T) {
	ir := fast.New()
	for i := range testcases {
		test := &testcases[i]
		if (!foundZ || test.testfor&Z != 0) && test.shouldRun(F) {
			t.Run(test.name, func(t *testing.T) { test.fast(t, ir) })
		}
	}
}

type shouldpanic struct{}

func (shouldpanic) String() string {
	return "shouldpanic"
}

// a value that the interpreter cannot produce.
// only matches if the interpreter panicked
var panics shouldpanic

var none = []interface{}{}

func (test *TestCase) classic(t *testing.T, ir *classic.Interp) {
	var rets []r.Value
	panicking := true
	if test.result0 == panics {
		defer func() {
			if panicking {
				recover()
			}
		}()
	}
	rets = reflect.PackValues(ir.Eval(test.program))
	panicking = false
	test.compareResults(t, rets)
}

func (test *TestCase) fast(t *testing.T, ir *fast.Interp) {
	if test.testfor&S != 0 {
		ir.Comp.Options |= OptDebugSleepOnSwitch
	} else {
		ir.Comp.Options &^= OptDebugSleepOnSwitch
	}
	if test.testfor&U != 0 {
		ir.Comp.Options |= OptKeepUntyped
	} else {
		ir.Comp.Options &^= OptKeepUntyped
	}

	var rets []r.Value
	panicking := true
	if test.result0 == panics {
		defer func() {
			if panicking {
				recover()
			}
		}()
	}

	rets, _ = ir.Eval(test.program)
	panicking = false
	test.compareResults(t, rets)
}

const sum_source_string = "func sum(n int) int { total := 0; for i := 1; i <= n; i++ { total += i }; return total }"
const fibonacci_source_string = "func fibonacci(n int) int { if n <= 2 { return 1 }; return fibonacci(n-1) + fibonacci(n-2) }"

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

const interface_interpreted_1_source_string = `
import (
	"errors"
	"fmt"
	"io"
	"os"
)

type R interface {
	Read([]uint8) (int, error)
}

type DevNull struct{}

func (d DevNull) Read(b []byte) (int, error) {
	return 0, io.EOF
}

type DevZero struct{}

func (d DevZero) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 0
	}
	return len(b), nil
}

true`

const interface_interpreted_2_source_string = `
(func() bool {

	fail := func(format string, args ...interface{}) {
		panic(errors.New(fmt.Sprintf(format, args...)))
	}

	f, _ := os.Open("README.md")
	bytes := make([]uint8, 80)

	rs := [3]R{f, DevNull{}, DevZero{}}
	lens := [3]int{80, 0, 80}
	errs := [3]error{nil, io.EOF, nil}

	for i, r := range rs {
		len, err := r.Read(bytes)
		if len != lens[i] || err != errs[i] {
			fail("Read(): expecting (%v, %v), returned (%v, %v)", lens[i], errs[i], len, err)
		}
		j := -1
		switch r := r.(type) {
		case *os.File:
			j = 0
			if r != rs[i] {
				fail("typeswitch: expecting %v, found %v", rs[i], r)
			}
		case DevNull:
			j = 1
			if r != rs[i] {
				fail("typeswitch: expecting %v, found %v", rs[i], r)
			}
		case DevZero:
			j = 2
			if r != rs[i] {
				fail("typeswitch: expecting %v, found %v", rs[i], r)
			}
		}
		if i != j {
			fail("typeswitch: expecting j=%d, found j=%d", i, j)
		}
	}
	return true
})()
`

var (
	classicInterfHeader = r.StructField{Name: StrGensymInterface, Type: r.TypeOf((*interface{})(nil)).Elem()}
	fastInterfHeader    = r.StructField{Name: StrGensymInterface, Type: r.TypeOf(xr.InterfaceHeader{})}

	classicTypStringer = r.StructOf(
		[]r.StructField{
			classicInterfHeader,
			r.StructField{Name: "String", Type: r.TypeOf((*func() string)(nil)).Elem()},
		},
	)
	fastTypeStringer = r.StructOf(
		[]r.StructField{
			fastInterfHeader,
			r.StructField{Name: "String", Type: r.TypeOf((*func() string)(nil)).Elem()},
		},
	)
	fastTypeEqint = r.StructOf(
		[]r.StructField{
			fastInterfHeader,
			r.StructField{Name: "Equal", Type: r.TypeOf((*func(int) bool)(nil)).Elem()},
		},
	)
	classicObjStringer = r.Zero(classicTypStringer).Interface()
	fastObjStringer    = r.Zero(r.PtrTo(fastTypeStringer)).Interface()
	fastObjEqint       = r.Zero(r.PtrTo(fastTypeEqint)).Interface()
)

var nil_map_int_string map[int]string

func for_range_string(s string) int32 {
	var v0 int32
	for i, r := range s {
		v0 += r << (uint8(i) * 8)
	}
	return v0
}

func makeQuote(node ast.Node) *ast.UnaryExpr {
	return makequote2(etoken.QUOTE, node)
}

func makeQUASIQUOTE(node ast.Node) *ast.UnaryExpr {
	return makequote2(etoken.QUASIQUOTE, node)
}

func makeUNQUOTE(node ast.Node) *ast.UnaryExpr {
	return makequote2(etoken.UNQUOTE, node)
}

func makeUNQUOTE_SPLICE(node ast.Node) *ast.UnaryExpr {
	return makequote2(etoken.UNQUOTE_SPLICE, node)
}

func makequote2(op token.Token, node ast.Node) *ast.UnaryExpr {
	unary, _ := parser.MakeQuote(nil, op, token.NoPos, node)
	return unary
}

type Pair = struct { // unnamed!
	A rune
	B string
}

type PairX2 = struct { // unnamed!
	First  complex64
	Second struct{}
}

type PairX3 = struct { // unnamed!
	First  bool
	Second interface{}
}

type ListX2 = struct { // unnamed!
	First error
	Rest  xr.Forward
}

type ListX3 = struct { // unnamed!
	First interface{}
	Rest  xr.Forward
}

type TagPair = struct { // unnamed!
	A rune   `json:"foo"`
	B string `json:"bar"`
}

type TagTriple = struct { // unnamed!
	A    rune
	B, C string `json:"baz"`
}

var bigInt = new(big.Int)
var bigRat = new(big.Rat)
var bigFloat = new(big.Float)

func init() {
	bigInt.SetInt64(1)
	bigInt.Lsh(bigInt, 1000)

	bigRat.SetFrac64(1000000001, 1000000000)
	bigRat.Mul(bigRat, bigRat)
	bigRat.Mul(bigRat, bigRat)

	// use the same precision as constant.Value
	bigFloat.SetPrec(512)
	bigFloat.SetString("1e1234")
	bigFloat.Mul(bigFloat, bigFloat)
	bigFloat.Mul(bigFloat, bigFloat)
}

func decl_generic_type_pair_str() string {
	if etoken.GENERICS_V1_CXX {
		return "~quote{template [T1,T2] type Pair struct { First T1; Second T2 }}"
	} else if etoken.GENERICS_V2_CTI {
		return "~quote{type Pair#[T1,T2] struct { First T1; Second T2 }}"
	} else {
		return ""
	}
}

func decl_generic_func_sum_str() string {
	if etoken.GENERICS_V1_CXX {
		return "~quote{template [T] func Sum([]T) T { }}"
	} else if etoken.GENERICS_V2_CTI {
		return "~quote{~func Sum#[T] ([]T) T { }}"
	} else {
		return ""
	}
}

func decl_generic_method_rest_str() string {
	if etoken.GENERICS_V1_CXX {
		return "~quote{template [T] func (x Pair) Rest() T { }}"
	} else if etoken.GENERICS_V2_CTI {
		return "~quote{~func (x Pair) Rest#[T] () T { }}"
	} else {
		return ""
	}
}

func generic_func(name string, generic_args string) string {
	if etoken.GENERICS_V1_CXX {
		return "template[" + generic_args + "] func " + name + " "
	} else if etoken.GENERICS_V2_CTI {
		return "func " + name + "#[" + generic_args + "]"
	} else {
		return ""
	}
}

func generic_type(name string, generic_args string) string {
	if etoken.GENERICS_V1_CXX {
		return "template[" + generic_args + "] type " + name + " "
	} else if etoken.GENERICS_V2_CTI {
		return "type " + name + "#[" + generic_args + "]"
	} else {
		return ""
	}
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

	// the classic interpreter is not accurate in these cases... missing exact arithmetic on constants
	TestCase{C, "const_3", "const c3 = 0.1+0.2; c3", float64(0.1) + float64(0.2), nil},
	TestCase{C, "const_4", "const c4 = c3/3; c4", (float64(0.1) + float64(0.2)) / 3, nil},

	// the fast interpreter instead *IS* accurate, thanks to exact arithmetic on untyped constants
	TestCase{F, "const_3", "const c3 = 0.1+0.2; c3", 0.1 + 0.2, nil},
	TestCase{F, "const_4", "const c4 = c3/3; c4", (0.1 + 0.2) / 3, nil},

	TestCase{F, "const_complex_1", "const c5 = complex(c3, c4); c5", 0.3 + 0.1i, nil},
	TestCase{F | U, "untyped_const_complex_1", "c5",
		untyped.MakeLit(
			untyped.Complex,
			constant.BinaryOp(
				constant.MakeFromLiteral("0.3", token.FLOAT, 0),
				token.ADD,
				constant.MakeFromLiteral("0.1i", token.IMAG, 0)),
			nil),
		nil,
	},

	TestCase{F, "untyped_1", "2.0 >> 1", 1, nil},
	TestCase{A, "untyped_2", "1/2", 0, nil},
	TestCase{A, "untyped_unary", "-+^6", -+^6, nil},
	TestCase{F | U, "untyped_const_large", "1<<100",
		untyped.MakeLit(untyped.Int, constant.Shift(constant.MakeInt64(1), token.SHL, 100), nil),
		nil,
	},

	TestCase{A, "iota_1", "const c5 = iota^7; c5", 7, nil},
	TestCase{A, "iota_2", "const ( c6 = iota+6; c7=iota+6 ); c6", 6, nil},
	TestCase{A, "iota_3", "c7", 7, nil},
	TestCase{A, "iota_implicit_1", "const ( c8 uint = iota+8; c9 ); c8", uint(8), nil},
	TestCase{A, "iota_implicit_2", "c9", uint(9), nil},

	TestCase{F, "zero_value_constructor_1", "int()", int(0), nil},
	TestCase{F, "zero_value_constructor_2", "uint16()", uint16(0), nil},
	TestCase{F, "zero_value_constructor_3", "float32()", float32(0), nil},
	TestCase{F, "zero_value_constructor_4", "complex128()", complex128(0), nil},
	TestCase{F, "zero_value_constructor_5", "string()", "", nil},
	TestCase{F, "zero_value_constructor_6", "[]int()", ([]int)(nil), nil},
	TestCase{F, "zero_value_constructor_7", "[2]int()", [2]int{0, 0}, nil},
	TestCase{F, "zero_value_constructor_8", "map[int]int()", (map[int]int)(nil), nil},
	TestCase{F, "zero_value_constructor_9", "chan string()", (chan string)(nil), nil},
	TestCase{F, "zero_value_constructor_10", "(*bool)()", (*bool)(nil), nil},
	TestCase{F, "zero_value_constructor_11", "struct{Foo int}()", struct{ Foo int }{}, nil},

	TestCase{A, "var_0", "var v0 int = 11; v0", 11, nil},
	TestCase{A, "var_1", "var v1 bool; v1", false, nil},
	TestCase{A, "var_2", "var v2 uint8 = 7; v2", uint8(7), nil},
	TestCase{A, "var_3", "var v3 uint16 = 12; v3", uint16(12), nil},
	TestCase{A, "var_4", "var v uint32 = 99; v", uint32(99), nil},
	TestCase{A, "var_5", "var vs string; vs", "", nil},
	TestCase{A, "var_6", "var v6 float32; v6", float32(0), nil},
	TestCase{A, "var_7", "var v7 complex64; v7", complex64(0), nil},
	TestCase{A, "var_8", "var v8 complex128; v8", complex128(0), nil},
	TestCase{A, "var_9", "var err error; err", nil, nil},
	TestCase{A, "var_10", `ve, vf := "", 1.23; ve`, "", nil},
	TestCase{A, "var_pointer", "var vp *string; vp", (*string)(nil), nil},
	TestCase{A, "var_map", "var vm *map[error]bool; vm", (*map[error]bool)(nil), nil},
	TestCase{A, "var_slice", "var vbs []byte; vbs", ([]byte)(nil), nil},
	TestCase{A, "var_named_slice", "type Bytes []byte; var vns Bytes; vns", ([]byte)(nil), nil},
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
	TestCase{C, "var_div_1", "v3 = 11; v3 / 2", uint64(11) / 2, nil}, // classic interpreter is not type-accurate here
	TestCase{C, "var_div_2", "v3 = 63; v3 / 8", uint64(63) / 8, nil},
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
	TestCase{C, "var_rem_1", "v3 = 17; v3 % 4", uint64(17) % 4, nil}, // classic interpreter is not type-accurate here
	TestCase{C, "var_rem_2", "v3 = 61; v3 % 8", uint64(61) % 8, nil},
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
	TestCase{A, "eql_nil_4", "vbs == nil", true, nil},
	TestCase{A, "eql_nil_5", "vi == nil", false, nil},
	TestCase{A, "eql_nil_6", "vnil == nil", true, nil},
	TestCase{A, "eql_halfnil_1", "var vhalfnil interface{} = vm; vhalfnil == nil", false, nil},
	TestCase{A, "eql_halfnil_2", "vm = nil; vhalfnil = vm; vhalfnil == nil", false, nil},
	TestCase{A, "eql_interface", "vi == 1", true, nil},

	TestCase{A, "typed_unary_1", "!!!v1", true, nil},
	TestCase{A, "typed_unary_2", "+-^v2", uint8(8), nil},
	TestCase{A, "typed_unary_3", "v3 = 12; +^-v3", uint16(11), nil},
	TestCase{A, "typed_unary_4", "v7 = 2.5i; -v7", complex64(-2.5i), nil},
	TestCase{A, "typed_unary_5", "v8 = 3.75i; -v8", complex128(-3.75i), nil},

	TestCase{A, "type_int8", "type t8 int8; var u8 t8; u8", int8(0), nil},
	TestCase{A, "type_complicated", "type tfff func(int,int) func(error, func(bool)) string; var vfff tfff; vfff", (func(int, int) func(error, func(bool)) string)(nil), nil},
	TestCase{C, "type_interface_1", "type Stringer interface { String() string }; var s Stringer; s", classicObjStringer, nil},
	TestCase{F, "type_interface_1", "type Stringer interface { String() string }; var s Stringer; s", fastObjStringer, nil},
	TestCase{F, "type_struct_0", "type PairPrivate struct { a, b rune }; var pp PairPrivate; pp.a+pp.b", rune(0), nil},
	TestCase{A, "type_struct_1", "type Pair struct { A rune; B string}; var pair Pair; pair", Pair{}, nil},
	TestCase{A, "type_struct_2", "type Triple struct { Pair; C float32 }; var triple Triple; triple.C", float32(0), nil},
	TestCase{A, "type_struct_3", "type TripleP struct { *Pair; D float64 }; var tp TripleP; tp.D", float64(0), nil},
	TestCase{F, "tagged_struct_1", "type TagPair struct { A rune `json:\"foo\"`; B string `json:\"bar\"`}; var tagpair TagPair; tagpair", TagPair{}, nil},
	TestCase{F, "tagged_struct_2", "type TagTriple struct { A rune; B, C string `json:\"baz\"`}; TagTriple{}", TagTriple{}, nil},

	TestCase{A, "field_get_1", "pair.A", rune(0), nil},
	TestCase{A, "field_get_2", "pair.B", "", nil},
	TestCase{F, "field_anonymous_1", "triple.Pair", Pair{}, nil},
	TestCase{F, "field_anonymous_2", "type Z struct { *Z }; Z{}", struct {
		Z xr.Forward
	}{}, nil},
	TestCase{F, "field_embedded_1", "triple.A", rune(0), nil},
	TestCase{F, "field_embedded_2", "triple.B", "", nil},
	TestCase{F, "field_embedded_3", "triple.Pair.A", rune(0), nil},
	TestCase{F, "field_embedded_4", "triple.Pair.B", "", nil},
	TestCase{F, "field_embedded_5", "tp.A", panics, nil},
	TestCase{F, "field_embedded_6", "tp.Pair = &triple.Pair; tp.B", "", nil},

	TestCase{F, "self_embedded_1", "type X struct { *X }; X{}.X", (xr.Forward)(nil), nil},
	TestCase{F, "self_embedded_2", "var x X; x.X = &x; x.X.X.X.X.X.X.X.X == &x", true, nil},
	TestCase{F, "self_embedded_3", "x.X.X.X == x.X.X.X.X.X", true, nil},

	TestCase{A, "address_0", "var vf = 1.25; *&vf == vf", true, nil},
	TestCase{A, "address_1", "var pvf = &vf; *pvf", 1.25, nil},
	TestCase{A, "address_2", "&*pvf == *&pvf", true, nil},
	TestCase{A, "address_3", "var pvs = &vbs; v1 = (*pvs == nil); v1", true, nil},

	TestCase{A, "make_chan", "cx := make(chan interface{}, 2); cx", make(chan interface{}, 2), nil},
	TestCase{A, "make_map", "m := make(map[int]string); m", make(map[int]string), nil},
	TestCase{A, "make_slice", "y := make([]uint8, 7); y[0] = 100; y[3] = 103; y", []uint8{100, 0, 0, 103, 0, 0, 0}, nil},
	TestCase{A, "expr_index_string_1", `"abc"[2]`, byte('c'), nil},
	TestCase{A, "expr_index_string_2", `vs = "foo"; v0 = 0; vs[v0]`, byte('f'), nil},
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
	TestCase{A, "set_const_5", `vs = "8y57r"; vs`, "8y57r", nil},
	TestCase{A, "set_const_6", "v6 = 0.12345678901234; v6", float32(0.12345678901234), nil},  // v6 is declared float32
	TestCase{A, "set_const_7", "v7 = 0.98765432109i;   v7", complex64(0.98765432109i), nil},  // v7 is declared complex64
	TestCase{A, "set_const_8", "v8 = 0.98765432109i;   v8", complex128(0.98765432109i), nil}, // v8 is declared complex128

	TestCase{A, "set_expr_1", "v1 = v1 == v1;    v1", true, nil},
	TestCase{A, "set_expr_2", "v2 -= 7;          v2", uint8(2), nil},
	TestCase{A, "set_expr_3", "v3 %= 7;          v3", uint16(60000) % 7, nil},
	TestCase{A, "set_expr_4", "v  = v * 10;       v", uint32(9870), nil},
	TestCase{A, "set_expr_5", `vs = vs + "iuh";  vs`, "8y57riuh", nil},
	TestCase{A, "set_expr_6", "v6 = 1/v6;        v6", 1 / float32(0.12345678901234), nil},                          // v6 is declared float32
	TestCase{A, "set_expr_7", "v7 = v7 * v7;     v7", -complex64(0.98765432109) * complex64(0.98765432109), nil},   // v7 is declared complex64
	TestCase{A, "set_expr_8", "v8 = v8 * v8;     v8", -complex128(0.98765432109) * complex128(0.98765432109), nil}, // v8 is declared complex64
	TestCase{A, "set_expr_9", `v9 := 0; { a := 1; { b := a+1; { c := b+1; { v9 = c+1 } } } }; v9`, int(4), nil},

	TestCase{A, "add_2", "v2 += 255;    v2", uint8(1), nil}, // overflow
	TestCase{A, "add_3", "v3 += 536;    v3", uint16(60000)%7 + 536, nil},
	TestCase{A, "add_4", "v  += 111;     v", uint32(9870 + 111), nil},
	TestCase{A, "add_5", `vs += "@#$";  vs`, "8y57riuh@#$", nil},
	TestCase{A, "add_6", "v6 += 0.975319; v6", 1/float32(0.12345678901234) + float32(0.975319), nil}, // v6 is declared float32
	TestCase{A, "add_7", "v7 = 1; v7 += 0.999999i; v7", complex(float32(1), float32(0.999999)), nil}, // v7 is declared complex64
	TestCase{A, "add_8", "v8 = 1; v8 += 0.999999i; v8", complex(1, 0.999999), nil},                   // v8 is declared complex128

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
	TestCase{A, "for_nested", `x := 0
		{
			var n1, n2, n3 = 2, 3, 5
			for i := 0; i < n1; i++ {
				for j := 0; j < n2; j++ {
					for k := 0; k < n3; k++ {
						x++
					}
				}
			}
		}
		x`, 2 * 3 * 5, nil},
	TestCase{A, "continue_1", "j=0; k=0; for i:=1; i<=7; i=i+1 { if i==3 {j=i; continue}; k=k+i }; j", 3, nil},
	TestCase{A, "continue_2", "k", 25, nil},
	TestCase{A, "continue_3", "j=0; k=0; for i:=1; i<=7; i=i+1 { var ii = i; if ii==3 {j=ii; continue}; k=k+ii }; j", 3, nil},
	TestCase{A, "continue_4", "k", 25, nil},

	TestCase{A, "for_range_array", `v0 = 0; for _, s := range [2]string{"a", "bc"} { v0 += len(s); continue }; v0`, 3, nil},
	TestCase{A, "for_range_ptr_array", `v0 = 0; var vis string; for _, vis = range &[...]string{"999", "1234"} { v0 += len(vis); continue }; v0`, 7, nil},
	TestCase{A, "for_range_chan", `v0 = 0; c := make(chan int, 2); c <- 1; c <- 2; close(c); for e := range c { v0 += e; continue }; v0`, 3, nil},
	TestCase{A, "for_range_map", `var vrune rune; m2 = map[rune]string{'x':"x", 'y':"y", 'z':"z"}; for k,v := range m2 { vrune += k + rune(v[0]); continue }; vrune`,
		('x' + 'y' + 'z') * 2, nil},
	TestCase{A, "for_range_slice", `v0 = 0; for _, s := range [ ]string{"a", "bc"} { v0 += len(s); continue }; v0`, 3, nil},
	TestCase{A, "for_range_string", `vrune = 0; for i, r := range "abc\u00ff" { vrune += r << (uint8(i)*8); continue }; vrune`, for_range_string("abc\u00ff"), nil},

	TestCase{A, "function_0", "func nop() { }; nop()", nil, none},
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

	TestCase{F, "y_combinator_1", "type F func(F); var f F; &f", new(xr.Forward), nil}, // xr.Forward is contagious
	TestCase{F, "y_combinator_2", "func Y(f F) { }; Y", func(xr.Forward) {}, nil},      // avoid the infinite recursion, only check the types
	TestCase{F, "y_combinator_3", "Y(Y)", nil, none},                                   // also check actual invokations
	TestCase{F, "y_combinator_4", "f=Y; f(Y)", nil, none},
	TestCase{F, "y_combinator_5", "Y(f)", nil, none},
	TestCase{F, "y_combinator_6", "f(f)", nil, none},

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
	TestCase{F, "multi_assignment_1", "v7, v8 = func () (complex64, complex128) { return 1.0, 2.0 }(); v7", complex64(1.0), nil},
	TestCase{F, "multi_assignment_2", "v8 ", complex128(2.0), nil},
	// gophernotes issue 175
	TestCase{F, "multi_assignment_3", `
		arr := [2]struct{X int}{{3},{4}}
		arr[0], arr[1] = arr[1], arr[0]
		arr`, [2]struct{ X int }{{4}, {3}}, nil},

	TestCase{A, "field_set_1", `pair.A = 'k'; pair.B = "m"; pair`, Pair{'k', "m"}, nil},
	TestCase{A, "field_set_2", `pair.A, pair.B = 'x', "y"; pair`, Pair{'x', "y"}, nil},
	TestCase{F, "field_set_3", `triple.Pair.A, triple.C = 'a', 1.0; triple.Pair`, Pair{'a', ""}, nil},
	TestCase{F, "field_set_embedded_1", `triple.A, triple.B = 'b', "xy"; triple.Pair`, Pair{'b', "xy"}, nil},
	TestCase{F, "field_addr_1", "ppair := &triple.Pair; ppair.A", 'b', nil},
	TestCase{F, "field_addr_2", "ppair.A++; triple.Pair.A", 'c', nil},

	TestCase{F, "infer_type_compositelit_1", `[]Pair{{'a', "b"}, {'c', "d"}}`, []Pair{{'a', "b"}, {'c', "d"}}, nil},
	TestCase{F, "infer_type_compositelit_2", `[]*Pair{{'a', "b"}, {'c', "d"}}`, []*Pair{{'a', "b"}, {'c', "d"}}, nil},
	TestCase{F, "infer_type_compositelit_3", `[...]Pair{{'e', "f"}, {'g', "h"}}`, [...]Pair{{'e', "f"}, {'g', "h"}}, nil},
	TestCase{F, "infer_type_compositelit_4", `map[int]Pair{1:{'i', "j"}, 2:{}}`, map[int]Pair{1: {'i', "j"}, 2: {}}, nil},
	TestCase{F, "infer_type_compositelit_5", `map[int]map[int]int{1:{2:3}}`, map[int]map[int]int{1: {2: 3}}, nil},
	TestCase{F, "infer_type_compositelit_6", `map[int]*map[int]int{1:{2:3}}`, map[int]*map[int]int{1: {2: 3}}, nil},

	TestCase{A, "import", `import ( "errors"; "fmt"; "io"; "math/big"; "math/rand"; "net/http"; "reflect"; "time" )`, nil, none},
	TestCase{A, "import_constant", `const micro = time.Microsecond; micro`, time.Microsecond, nil},
	TestCase{A, "dot_import_1", `import . "errors"`, nil, none},
	TestCase{A, "dot_import_2", `reflect.ValueOf(New) == reflect.ValueOf(errors.New)`, true, nil}, // a small but very strict check... good

	TestCase{A, "goroutine_1", `go seti(9); time.Sleep(time.Second/50); i`, 9, nil},

	TestCase{F, "big.Int", `(func() *big.Int { return 1<<1000 })()`, bigInt, nil},
	TestCase{F, "big.Rat", `(func() *big.Rat { var x *big.Rat = 1.000000001; x.Mul(x,x); x.Mul(x,x); return x })()`, bigRat, nil},
	TestCase{F, "big.Float", `(func() *big.Float { var x *big.Float = 1e1234; x.Mul(x,x); x.Mul(x,x); return x })()`, bigFloat, nil},

	TestCase{A, "builtin_append_1", "append(vbs,0,1,2)", []byte{0, 1, 2}, nil},
	TestCase{A, "builtin_append_2", "append(vns,3,4)", []byte{3, 4}, nil},
	TestCase{A, "builtin_cap", "cap(va)", 2, nil},
	TestCase{A, "builtin_len_1", "len(vs)", len("8y57riuh@#$"), nil},
	TestCase{A, "builtin_len_2", "{ a := [...]int{1,2,3}; len(a) }", nil, none},
	TestCase{A, "builtin_new", "new(int)", new(int), nil},
	TestCase{A, "builtin_make_1", "make(map[int]int)", make(map[int]int), nil},
	TestCase{A, "builtin_make_2", "make(map[int]int, 10)", make(map[int]int), nil}, // capacity is ignored
	TestCase{A, "builtin_make_4", "make([]*error, 2)", make([]*error, 2), nil},
	TestCase{A, "builtin_make_5", "make([]rune, 3, 4)", make([]rune, 3, 4), nil},
	TestCase{A, "builtin_make_6", "make(chan byte)", make(chan byte), nil},
	TestCase{A, "builtin_make_7", "make(chan byte, 2)", make(chan byte, 2), nil},
	TestCase{A, "builtin_make_8", "vbs = make([]byte, 5); vbs", make([]byte, 5), nil},
	TestCase{A, "builtin_copy_1", "copy(vbs, vs)", 5, nil},
	TestCase{A, "builtin_copy_2", "vbs", []byte("8y57r"), nil},
	TestCase{A, "builtin_delete_1", "delete(mi,64); mi", map[rune]byte{'a': 7}, nil},
	TestCase{A, "builtin_real_1", "real(0.5+1.75i)", real(0.5 + 1.75i), nil},
	TestCase{A, "builtin_real_2", "const cplx complex64 = 1.5+0.25i; real(cplx)", real(complex64(1.5 + 0.25i)), nil},
	TestCase{A, "builtin_imag_1", "imag(0.5+1.75i)", imag(0.5 + 1.75i), nil},
	TestCase{A, "builtin_imag_2", "imag(cplx)", imag(complex64(1.5 + 0.25i)), nil},
	TestCase{A, "builtin_complex_1", "complex(0,1)", complex(0, 1), nil},
	TestCase{A, "builtin_complex_2", "v6 = 0.1; complex(v6,-v6)", complex(float32(0.1), -float32(0.1)), nil},

	TestCase{F | U, "untyped_builtin_real_1", "real(0.5+1.75i)",
		untyped.MakeLit(untyped.Float, constant.MakeFloat64(0.5), nil), // 0.5 is exactly representable by float64
		nil},
	TestCase{F | U, "untyped_builtin_imag_1", "imag(1.5+0.25i)",
		untyped.MakeLit(untyped.Float, constant.MakeFloat64(0.25), nil), // 0.25 is exactly representable by float64
		nil},
	TestCase{F | U, "untyped_builtin_complex_1", "complex(1, 2)",
		untyped.MakeLit(
			untyped.Complex,
			constant.BinaryOp(
				constant.MakeInt64(1),
				token.ADD,
				constant.MakeFromLiteral("2i", token.IMAG, 0)),
			nil),
		nil,
	},

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

	TestCase{A, "index_is_named_type", `"abc"[time.Nanosecond]`, uint8('b'), nil},
	TestCase{A, "panic_at_runtime", `"abc"[micro]`, panics, nil},
	TestCase{F, "panic_oob_at_compile", `(func() uint8 { return "abc"[micro] })`, panics, nil}, // string index out of range
	TestCase{F, "panic_non_const_initialization", `const _ = rand.Int()`, panics, nil},         // const initializer is not a constant

	TestCase{A, "literal_array", "[3]int{1,2:3}", [3]int{1, 0, 3}, nil},
	TestCase{A, "literal_array_address", "&[...]int{3:4,5:6}", &[...]int{3: 4, 5: 6}, nil},
	TestCase{A, "literal_map", `map[int]string{1: "foo", 2: "bar"}`, map[int]string{1: "foo", 2: "bar"}, nil},
	TestCase{A, "literal_map_address", `&map[int]byte{6:7, 8:9}`, &map[int]byte{6: 7, 8: 9}, nil},
	TestCase{A, "literal_slice", "[]rune{'a','b','c'}", []rune{'a', 'b', 'c'}, nil},
	TestCase{A, "literal_slice_address", "&[]rune{'x','y'}", &[]rune{'x', 'y'}, nil},
	TestCase{A, "literal_struct", `Pair{A: 0x73, B: "\x94"}`, Pair{A: 0x73, B: "\x94"}, nil},
	TestCase{A, "literal_struct_address", `&Pair{1,"2"}`, &Pair{A: 1, B: "2"}, nil},

	TestCase{A, "named_func_type_1", `import "context"
		         _, cancel := context.WithCancel(context.Background())
		         cancel()`, nil, none},

	TestCase{A, "method_decl_1", `func (p *Pair) SetA(a rune) { p.A = a }; nil`, nil, nil},
	TestCase{A, "method_decl_2", `func (p Pair) SetAV(a rune) { p.A = a }; nil`, nil, nil},
	TestCase{A, "method_decl_3", `func (p Pair) String() string { return fmt.Sprintf("%c %s", p.A, p.B) }; nil`, nil, nil},

	TestCase{A, "method_on_ptr", `pair.SetA(33); pair.A`, rune(33), nil},
	TestCase{A, "method_on_val_1", `pair.SetAV(11); pair.A`, rune(33), nil}, // method on value gets a copy of the receiver - changes to not propagate
	TestCase{A, "method_on_val_2", `pair.String()`, "! y", nil},
	// gophernotes issue 174
	TestCase{F, "method_decl_and_use", `type person struct{}; func (p person) speak() {}; person.speak`, func(struct{}) {}, nil},
	TestCase{F, "method_embedded=val_recv=ptr", `triple.SetA('1'); triple.A`, '1', nil},
	TestCase{F, "method_embedded=val_recv=val", `triple.SetAV('2'); triple.A`, '1', nil},
	TestCase{F, "method_embedded=ptr_recv=val", `tp.SetAV('3'); tp.A`, '1', nil}, // set by triple.SetA('1') above
	TestCase{F, "method_embedded=ptr_recv=ptr", `tp.SetA('4'); tp.A`, '4', nil},

	TestCase{F, "concrete_method_to_func_1", `cf0 := time.Duration.Seconds; cf0(time.Hour)`, 3600.0, nil},
	TestCase{F, "concrete_method_to_closure_1", `cl1 := time.Hour.Seconds; cl1()`, 3600.0, nil},
	TestCase{F, "concrete_method_to_func_2", `import "sync"; sync.WaitGroup.Done`, (*sync.WaitGroup).Done, nil},
	TestCase{F, "concrete_method_to_closure_2", `var wg sync.WaitGroup; wg.Done`, func() {}, nil},

	// tricky because Comp.compileObjGetMethod() asks for the package path of 'error', which has nil package
	TestCase{A, "interface_0", `errors.New("abc").Error()`, "abc", nil},

	TestCase{A, "interface_1", "var st fmt.Stringer = time.Second; st", time.Second, nil},
	TestCase{A, "interface_method_to_closure_1", "bind := st.String; bind()", "1s", nil},
	TestCase{F, "interface_2", "st = pair; nil", nil, nil},
	TestCase{F, "interface_method_to_closure_2", "bind = st.String; bind()", "! y", nil},
	TestCase{F, "interface_method_to_closure_3", "st = nil; bind = st.String", panics, nil},
	// interpreted interface
	TestCase{F, "interface_3", "type IStringer interface { String() string }; nil", nil, nil},
	TestCase{F, "interface_method_to_closure_4", "var ist IStringer; nil", nil, nil},
	TestCase{F, "interface_method_to_closure_5", "ist.String", panics, nil},
	TestCase{F, "interface_4", `
		type IncAdd interface { Inc(int); Add(int) int }
		type Int int
		func (i Int)  Add(j int) int { return int(i) + j }
		func (i *Int) Inc(j int) { *i += Int(j) }
		var ia IncAdd
		var ii Int = 7`, nil, none},
	TestCase{F, "interface_5", `ia = &ii`, nil, none},

	TestCase{F, "interface_method_to_func_1", "f1 := fmt.Stringer.String; f1(time.Hour)", "1h0m0s", nil},
	TestCase{F, "interface_method_to_func_2", "f2 := io.ReadWriter.Read; f2 != nil", true, nil},
	TestCase{F, "interface_method_to_func_3", "type Fooer interface { Foo() }; Fooer.Foo != nil", true, nil},
	TestCase{F, "interface_method_to_func_4", "type RW interface { io.Reader; io.Writer }; RW.Read != nil && RW.Write != nil", true, nil},

	TestCase{F, "interface_interpreted_1", interface_interpreted_1_source_string, true, nil},
	TestCase{F, "interface_interpreted_2", interface_interpreted_2_source_string, true, nil},
	// gophernotes issue 176
	TestCase{F, "interface_interpreted_3", `
		type xerror struct { }
		func (x xerror) Error() string {
		  return "some error"
		}
		var xe error = xerror{}
		xe.Error()`, "some error", nil},

	TestCase{A, "multiple_values_1", "func twins(x float32) (float32,float32) { return x, x+1 }; twins(17.0)", nil, []interface{}{float32(17.0), float32(18.0)}},
	TestCase{A, "multiple_values_2", "func twins2(x float32) (float32,float32) { return twins(x) }; twins2(19.0)", nil, []interface{}{float32(19.0), float32(20.0)}},
	TestCase{A, "multiple_values_3", "f1, f2 := twins(23.0); list_args(f1, f2)", []interface{}{float32(23.0), float32(24.0)}, nil},
	TestCase{A, "multiple_values_4", "fm := make(map[int]float32); fm[1], fm[2] = twins(3.0); fm", map[int]float32{1: 3.0, 2: 4.0}, nil},
	TestCase{A, "multiple_values_5", "swap(swap(3,4))", nil, []interface{}{3, 4}},
	TestCase{F, "multiple_values_6", `fmt.Sprintf("foo")`, "foo", nil},
	TestCase{A, "multiple_values_7", `func args() (string, interface{}, interface{}) { return "%v %v", 5, 6 }; nil`, nil, nil},
	TestCase{A, "multiple_values_8", `fmt.Sprintf(args())`, "5 6", nil},

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
	// classic does not fully support named return types
	TestCase{F, "defer_2", `
		func test_defer_2() (x int) {
			defer func() {
				x = 2
			}()
		}
		test_defer_2()`, 2, nil},
	TestCase{A, "defer_3", `
		v = 0
		func test_defer_3(x uint32) {
			if x != 0 {
				defer func(y uint32) {
					 v = y
				}(x)
			}
		}
		test_defer_3(3); v`, uint32(3), nil},
	TestCase{A, "defer_4", "v = 4; test_defer_3(0); v", uint32(4), nil},
	TestCase{A, "defer_5", `
		v = 0
		func test_defer_5(x uint32) {
			if x != 0 {
				defer func() {
					 v = x
				}()
			}
		}
		test_defer_5(5); v`, uint32(5), nil},
	TestCase{A, "defer_6", "v = 6; test_defer_5(0); v", uint32(6), nil},
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
	TestCase{A, "for_select_1", "for { select { }; break }", nil, none},
	// FIXME: break is confused by select { default: }
	TestCase{C, "for_select_2", "for { select { default: }; break }", nil, none},
	// non-empty 'select' needs a local bind, and 'for' must know it
	TestCase{A, "for_select_3", "_ = func() { for { select { }; return } }", nil, none},
	TestCase{A, "for_select_4", "_ = func() { for { select { case <-cx: default: return } } }", nil, none},

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
	TestCase{A, "for_switch_1", "for { switch { }; break }", nil, none},
	TestCase{A, "for_switch_2", "for { switch { default: }; break }", nil, none},

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
	TestCase{A, "typeassert_5", `xi = 7; xi.(int)+2`, 9, nil},
	TestCase{F, "typeassert_6", `type T struct { Val int }; func (t T) String() string { return "T" }`, nil, none},
	TestCase{F, "typeassert_7", `stringer = T{}; nil`, nil, nil},
	TestCase{F, "typeassert_8", `st1 := stringer.(T); st1`, struct{ Val int }{0}, nil},
	TestCase{F, "typeassert_9", `stringer.(T)`, nil, []interface{}{struct{ Val int }{0}, true}},
	// can interpreted type assertions distinguish between emulated named types with identical underlying type?
	TestCase{F, "typeassert_10", `type U struct { Val int }; func (u U) String() string { return "U" }; nil`, nil, nil},
	TestCase{F, "typeassert_11", `stringer.(U)`, nil, []interface{}{struct{ Val int }{0}, false}},

	TestCase{A, "quote_1", `~quote{7}`, &ast.BasicLit{Kind: token.INT, Value: "7"}, nil},
	TestCase{A, "quote_2", `~quote{x}`, &ast.Ident{Name: "x"}, nil},
	TestCase{A, "quote_3", `var ab = ~quote{a;b}; ab`, &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.Ident{Name: "a"}},
		&ast.ExprStmt{X: &ast.Ident{Name: "b"}},
	}}, nil},
	TestCase{A, "quote_4", `~'{"foo"+"bar"}`, &ast.BinaryExpr{
		X:  &ast.BasicLit{Kind: token.STRING, Value: `"foo"`},
		Op: token.ADD,
		Y:  &ast.BasicLit{Kind: token.STRING, Value: `"bar"`},
	}, nil},
	TestCase{A, "quasiquote_1", `~quasiquote{1 + ~unquote{2+3}}`, &ast.BinaryExpr{
		X:  &ast.BasicLit{Kind: token.INT, Value: "1"},
		Op: token.ADD,
		Y:  &ast.BasicLit{Kind: token.INT, Value: "5"},
	}, nil},
	TestCase{A, "quasiquote_2", `~"{2 * ~,{3<<1}}`, &ast.BinaryExpr{
		X:  &ast.BasicLit{Kind: token.INT, Value: "2"},
		Op: token.MUL,
		Y:  &ast.BasicLit{Kind: token.INT, Value: "6"},
	}, nil},
	TestCase{A, "quasiquote_3", `~"{func(int) {}}`, &ast.FuncLit{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
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
	TestCase{A, "unquote_splice_3", `~"~"{zero ; ~,~,@ab ; one}`,
		makeQUASIQUOTE(&ast.BlockStmt{List: []ast.Stmt{
			&ast.ExprStmt{X: &ast.Ident{Name: "zero"}},
			&ast.ExprStmt{X: makeUNQUOTE(&ast.Ident{Name: "a"})},
			&ast.ExprStmt{X: makeUNQUOTE(&ast.Ident{Name: "b"})},
			&ast.ExprStmt{X: &ast.Ident{Name: "one"}},
		}}), nil},
	TestCase{A, "unquote_splice_4", `~"~"{zero ; ~,@~,@ab ; one}`,
		makeQUASIQUOTE(&ast.BlockStmt{List: []ast.Stmt{
			&ast.ExprStmt{X: &ast.Ident{Name: "zero"}},
			&ast.ExprStmt{X: makeUNQUOTE_SPLICE(&ast.Ident{Name: "a"})},
			&ast.ExprStmt{X: makeUNQUOTE_SPLICE(&ast.Ident{Name: "b"})},
			&ast.ExprStmt{X: &ast.Ident{Name: "one"}},
		}}), nil},
	TestCase{A, "macro", "~macro second_arg(a,b,c interface{}) interface{} { return b }; v = 98; v", uint32(98), nil},
	TestCase{A, "macro_call", "second_arg;1;v;3", uint32(98), nil},
	TestCase{A, "macro_nested", "second_arg;1;{second_arg;2;3;4};5", 3, nil},
	TestCase{C, "values", "Values(3,4,5)", nil, []interface{}{3, 4, 5}},
	TestCase{A, "eval", "Eval(~quote{1+2})", 3, nil},
	TestCase{C, "eval_quote", "Eval(~quote{Values(3,4,5)})", nil, []interface{}{3, 4, 5}},

	TestCase{A | G1 | G2, "parse_decl_generic_type_1", decl_generic_type_pair_str(),
		&ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{
				&ast.TypeSpec{
					Name: &ast.Ident{Name: "Pair"},
					Type: &ast.CompositeLit{
						Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									{
										Names: []*ast.Ident{{Name: "First"}},
										Type:  &ast.Ident{Name: "T1"},
									},
									{
										Names: []*ast.Ident{{Name: "Second"}},
										Type:  &ast.Ident{Name: "T2"},
									},
								},
							},
						},
						Elts: []ast.Expr{
							&ast.Ident{Name: "T1"},
							&ast.Ident{Name: "T2"},
						},
					},
				},
			},
		}, nil},

	TestCase{A | G1 | G2, "parse_decl_generic_func_1", decl_generic_func_sum_str(),
		&ast.FuncDecl{
			Recv: &ast.FieldList{
				List: []*ast.Field{
					nil,
					{
						Type: &ast.CompositeLit{
							Elts: []ast.Expr{
								&ast.Ident{Name: "T"},
							},
						},
					},
				},
			},
			Name: &ast.Ident{Name: "Sum"},
			Type: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Type: &ast.ArrayType{
								Elt: &ast.Ident{Name: "T"},
							},
						},
					},
				},
				Results: &ast.FieldList{
					List: []*ast.Field{
						{
							Type: &ast.Ident{Name: "T"},
						},
					},
				},
			},
			Body: &ast.BlockStmt{},
		}, nil},

	TestCase{A | G1 | G2, "parse_decl_generic_method", decl_generic_method_rest_str(),
		&ast.FuncDecl{
			Recv: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{{Name: "x"}},
						Type:  &ast.Ident{Name: "Pair"},
					},
					{
						Type: &ast.CompositeLit{
							Elts: []ast.Expr{
								&ast.Ident{Name: "T"},
							},
						},
					},
				},
			},
			Name: &ast.Ident{Name: "Rest"},
			Type: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{},
				},
				Results: &ast.FieldList{
					List: []*ast.Field{
						{
							Type: &ast.Ident{Name: "T"},
						},
					},
				},
			},
			Body: &ast.BlockStmt{},
		}, nil},

	TestCase{A | G1 | G2, "parse_qual_generic_name_1", "~quote{Pair#[]}",
		&ast.IndexExpr{
			X:     &ast.Ident{Name: "Pair"},
			Index: &ast.CompositeLit{},
		}, nil},

	TestCase{A | G1 | G2, "parse_qual_generic_name_2", "~quote{Pair#[x + 1]}",
		&ast.IndexExpr{
			X: &ast.Ident{Name: "Pair"},
			Index: &ast.CompositeLit{
				Elts: []ast.Expr{
					&ast.BinaryExpr{
						X:  &ast.Ident{Name: "x"},
						Op: token.ADD,
						Y: &ast.BasicLit{
							Kind:  token.INT,
							Value: "1",
						},
					},
				},
			},
		}, nil},

	TestCase{A | G1 | G2, "parse_qual_generic_name_3", "~quote{Pair#[T1, T2]}",
		&ast.IndexExpr{
			X: &ast.Ident{Name: "Pair"},
			Index: &ast.CompositeLit{
				Elts: []ast.Expr{
					&ast.Ident{Name: "T1"},
					&ast.Ident{Name: "T2"},
				},
			},
		}, nil},

	TestCase{F | G1 | G2, "generic_func_1",
		generic_func("Identity", "T") + `(arg T) T {
			return arg
		}`, nil, none,
	},
	TestCase{F | G1 | G2, "generic_func_2",
		`Identity#[float64](1.5)`,
		float64(1.5), nil,
	},
	TestCase{F | G1, "generic_func_3",
		`Identity#[func()]`,
		func(func()) func() { return nil }, nil,
	},
	TestCase{F | G1 | G2, "generic_func_4",
		generic_func("Sum", "T") + `(args ...T) T {
			var sum T
			for _, elem := range args {
				sum += elem
			}
			return sum
		}`, nil, none,
	},
	TestCase{F | G1 | G2, "generic_func_5", `Sum#[int]`, func(...int) int { return 0 }, nil},
	TestCase{F | G1 | G2, "generic_func_6", `Sum#[complex64]`, func(...complex64) complex64 { return 0 }, nil},
	TestCase{F | G1 | G2, "generic_func_7", `Sum#[int](1, 2, 3)`, 6, nil},
	TestCase{F | G1 | G2, "generic_func_8", `Sum#[complex64](1.1+2.2i, 3.3)`, complex64(1.1+2.2i) + complex64(3.3), nil},
	TestCase{F | G1 | G2, "generic_func_9", `Sum#[string]("abc","def","xy","z")`, "abcdefxyz", nil},

	TestCase{F | G1 | G2, "generic_func_10",
		generic_func("Transform", "T,U") + ` (slice []T, trans func(T) U) []U {
			ret := make([]U, len(slice))
			for i := range slice {
				ret[i] = trans(slice[i])
			}
			return ret
		}
		func stringLen(s string) int { return len(s) }`, nil, none,
	},
	TestCase{F | G1 | G2, "generic_func_11", `Transform#[string,int]([]string{"abc","xy","z"}, stringLen)`,
		[]int{3, 2, 1}, nil,
	},
	TestCase{F | G1 | G2, "generic_func_12",
		generic_func("SwapArgs", "A,B,C") + ` (f func(A, B) C) func(B,A) C {
			return func (b B, a A) C {
				return f(a, b)
			}
		}`, nil, none,
	},
	TestCase{F | G1 | G2, "generic_func_13", `
		SwapArgs#[float64,float64,float64](func (a float64, b float64) float64 { return a/b })(2.0, 3.0)
	    `, 1.5, nil,
	},
	TestCase{F | G1 | G2, "generic_func_curry_1",
		generic_func("Curry", "A,B,C") + ` (f func(A, B) C) func(A) func(B) C {
			return func (a A) func (B) C {
				return func (b B) C {
					return f(a, b)
				}
			}
		}
		` + generic_func("add2", "T") + ` (a,b T) T { return a+b }
		Curry#[int,int,int](add2#[int])(2)(3)
	`,
		5, nil},

	TestCase{F | G2, "generic_func_curry_2",
		generic_func("add2m", "T") + ` (a,b T) T { return T().Add(a,b) }
		Curry#[uint,uint,uint](add2m#[uint])(5)(6)
	`,
		uint(11), nil},

	TestCase{F | G1 | G2, "generic_func_lift_1",
		generic_func("Lift1", "A,B") + ` (trans func(A) B) func([]A) []B {
			return func(slice []A) []B {
				ret := make([]B, len(slice))
				for i := range slice {
					ret[i] = trans(slice[i])
				}
				return ret
			}
		}
		Lift1#[string,int](stringLen)([]string{"a","bc","def"})
	`,
		[]int{1, 2, 3}, nil},

	// quite a convoluted test
	TestCase{F | G1 | G2, "generic_func_lift_2",
		generic_func("Lift2", "A,B") + ` (trans func(A) B) func([]A) []B {
			return Curry#[func(A)B, []A, []B](
				SwapArgs#[[]A, func(A)B, []B](Transform#[A,B]),
			)(trans)
		}
		Lift2#[string,int](stringLen)([]string{"xy","z",""})
	`,
		[]int{2, 1, 0}, nil},

	TestCase{F | G1 | G2, "recursive_generic_func_1",
		generic_func("count", "T") + ` (a, b T) T { if a <= 0 { return b }
		return count#[T](a-1,b+1) }`,
		nil, none,
	},
	TestCase{F | G1 | G2, "recursive_generic_func_2", `count#[uint16]`, func(uint16, uint16) uint16 { return 0 }, nil},
	TestCase{F | G1 | G2, "recursive_generic_func_3", `count#[uint32](2,3)`, uint32(5), nil},

	TestCase{F | G1, "specialized_generic_func_1", `template[] for[bool] func count(a, b bool) bool { return a || b }`, nil, none},
	TestCase{F | G1, "specialized_generic_func_2", `count#[bool]`, func(bool, bool) bool { return false }, nil},
	TestCase{F | G1, "specialized_generic_func_3", `count#[bool](false, true)`, true, nil},
	TestCase{F | G1, "specialized_generic_func_4", `template[T] for[*T] func count(a, b *T) *T { return a }`, nil, none},
	TestCase{F | G1, "specialized_generic_func_5", `count#[*int]`, func(*int, *int) *int { return nil }, nil},

	TestCase{F | G1 | G2, "generic_type_1",
		generic_type("PairX", "T1,T2") + `struct { First T1; Second T2 }`,
		nil, none,
	},
	TestCase{F | G1 | G2, "generic_type_2", `var px PairX#[complex64, struct{}]; px`, PairX2{}, nil},
	TestCase{F | G1 | G2, "generic_type_3", `PairX#[bool, interface{}] {true, "foo"}`, PairX3{true, "foo"}, nil},

	TestCase{F | G1 | G2, "recursive_generic_type_1",
		generic_type("ListX", "T") + `struct { First T; Rest *ListX#[T] }
		var lx ListX#[error]; lx`, ListX2{}, nil},
	TestCase{F | G1 | G2, "recursive_generic_type_2", `ListX#[interface{}]{}`, ListX3{}, nil},

	TestCase{F | G1, "specialized_generic_type_1", `
		template[] for[struct{}] type ListX struct { }
		template[T] for[T,T] type PairX struct { Left, Right T }
		PairX#[bool,bool]{false,true}`, struct{ Left, Right bool }{false, true}, nil},

	TestCase{F | G1, "turing_complete_generic_1", `
		template[N] type Fib [len((*Fib#[N-1])(nil)) + len((*Fib#[N-2])(nil))] int
		template[] for[1] type Fib [1]int
		template[] for[0] type Fib [0]int
		const Fib30 = len((*Fib#[30])(nil)); Fib30`, 832040, nil},

	TestCase{F | G2, "cti_basic_method_1", `1.Add(2, 3)`, 2 + 3, nil},
	TestCase{F | G2, "cti_basic_method_2", `1.2.Mul(2.3, 3.4)`, float64(2.3) * float64(3.4), nil},
	TestCase{F | G2, "cti_basic_method_3", `false.Not(true)`, false, nil},
	TestCase{F | G2, "cti_basic_method_4", `uint64(7).Less(7)`, false, nil},
	TestCase{F | G2, "cti_basic_method_5", `int.Cmp(1, 2)`, -1, nil},
	TestCase{F | G2, "cti_basic_method_6", `8.Equal(8)`, true, nil},
	TestCase{F | G2, "cti_basic_method_7", `8.9i.Imag()`, 8.9, nil},
	TestCase{F | G2, "cti_basic_method_8", `"abc".Index(2)`, "abc"[2], nil},
	TestCase{F | G2, "cti_basic_method_9", `"abcdefgh".Len()`, len("abcdefgh"), nil},
	TestCase{F | G2, "cti_basic_method_10", `"wxyz".Slice(1,2)`, "wxyz"[1:2], nil},

	TestCase{F | G2, "cti_method_array_len", `[...]int{1,2}.Len()`, len([...]int{1, 2}), nil},
	TestCase{F | G2, "cti_method_array_index", `[...]int{999:1}.Index(999)`, 1, nil},
	TestCase{F | G2, "cti_method_array_slice", `[...]int{0,1,2,3,4,5}.Slice(2,5)`, []int{2, 3, 4}, nil},
	TestCase{F | G2, "cti_method_chan_cap", `make(chan int).Cap()`, cap(make(chan int)), nil},
	TestCase{F | G2, "cti_method_slice_len", `[]int{3,4,5}.Len()`, len([]int{3, 4, 5}), nil},
	TestCase{F | G2, "cti_method_slice_slice", `[]int{0,1,2,3,4,5}.Slice(1,4)`, []int{1, 2, 3}, nil},
	TestCase{F | G2, "cti_method_map_index", `map[int]uint{1:1,-2:2}.Index(-2)`, map[int]uint{1: 1, -2: 2}[-2], nil},

	TestCase{A | G2, "parse_constrained_generic_1", "~quote{Set#[T: Eq]}",
		&ast.IndexExpr{
			X: &ast.Ident{Name: "Set"},
			Index: &ast.CompositeLit{
				Elts: []ast.Expr{
					&ast.KeyValueExpr{
						Key:   &ast.Ident{Name: "T"},
						Value: &ast.Ident{Name: "Eq"},
					},
				},
			},
		}, nil},
	TestCase{A | G2, "parse_constrained_generic_2", "~quote{Set#[T: Eq && Ord]}",
		&ast.IndexExpr{
			X: &ast.Ident{Name: "Set"},
			Index: &ast.CompositeLit{
				Elts: []ast.Expr{
					&ast.KeyValueExpr{
						Key: &ast.Ident{Name: "T"},
						Value: &ast.BinaryExpr{
							X:  &ast.Ident{Name: "Eq"},
							Op: token.LAND,
							Y:  &ast.Ident{Name: "Ord"},
						},
					},
				},
			},
		}, nil},
	TestCase{A | G2, "parse_constrained_generic_3", "~quote{Set#[T: Eq#[T] && Ord#[T]]}",
		&ast.IndexExpr{
			X: &ast.Ident{Name: "Set"},
			Index: &ast.CompositeLit{
				Elts: []ast.Expr{
					&ast.KeyValueExpr{
						Key: &ast.Ident{Name: "T"},
						Value: &ast.BinaryExpr{
							X: &ast.IndexExpr{
								X: &ast.Ident{Name: "Eq"},
								Index: &ast.CompositeLit{
									Elts: []ast.Expr{
										&ast.Ident{Name: "T"},
									},
								},
							},
							Op: token.LAND,
							Y: &ast.IndexExpr{
								X: &ast.Ident{Name: "Ord"},
								Index: &ast.CompositeLit{
									Elts: []ast.Expr{
										&ast.Ident{Name: "T"},
									},
								},
							},
						},
					},
				},
			},
		}, nil},
	TestCase{A | G2, "parse_constrained_generic_4", "~quote{SortedMap#[K: Ord, V: Container#[SortedMap#[K,V],K,V]]}",
		&ast.IndexExpr{
			X: &ast.Ident{Name: "SortedMap"},
			Index: &ast.CompositeLit{
				Elts: []ast.Expr{
					&ast.KeyValueExpr{
						Key:   &ast.Ident{Name: "K"},
						Value: &ast.Ident{Name: "Ord"},
					},
					&ast.KeyValueExpr{
						Key: &ast.Ident{Name: "V"},
						Value: &ast.IndexExpr{
							X: &ast.Ident{Name: "Container"},
							Index: &ast.CompositeLit{
								Elts: []ast.Expr{
									&ast.IndexExpr{
										X: &ast.Ident{Name: "SortedMap"},
										Index: &ast.CompositeLit{
											Elts: []ast.Expr{
												&ast.Ident{Name: "K"},
												&ast.Ident{Name: "V"},
											},
										},
									},
									&ast.Ident{Name: "K"},
									&ast.Ident{Name: "V"},
								},
							},
						},
					},
				},
			},
		}, nil},
	TestCase{A | G2, "parse_generic_contract_1", `~quote{
		type Eq#[T] interface{
			func (T) Equal(T) bool
		}}`,
		&ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{
				&ast.TypeSpec{
					Name: &ast.Ident{Name: "Eq"},
					Type: &ast.CompositeLit{
						Type: &ast.InterfaceType{
							Methods: &ast.FieldList{
								List: []*ast.Field{
									&ast.Field{
										Names: []*ast.Ident{
											&ast.Ident{Name: "Equal"},
										},
										Type: &ast.MapType{
											Key: &ast.Ident{Name: "T"}, // receiver
											Value: &ast.FuncType{
												Params: &ast.FieldList{
													List: []*ast.Field{
														&ast.Field{
															Type: &ast.Ident{Name: "T"},
														},
													},
												},
												Results: &ast.FieldList{
													List: []*ast.Field{
														&ast.Field{
															Type: &ast.Ident{Name: "bool"},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						Elts: []ast.Expr{
							&ast.Ident{Name: "T"},
						},
					},
				},
			},
		}, nil},
	TestCase{F | G2, "generic_contract_1", `
		type Eq#[T] interface{
			func (T) Equal(T) bool
		}
		var xg1 Eq#[int]
		xg1`, fastObjEqint, nil},
	TestCase{F | G2, "generic_contract_2", `
		type UInt uint
		func (i UInt) Equal(j UInt) bool {
			return i == j
		}`, nil, none},
	TestCase{F | G2, "generic_contract_3", `
		xg2 := UInt(9)
		var xg3 Eq#[UInt]
		xg3 = xg2
		xg2`, uint(9), nil},
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
	if actual == nil || expected == nil {
		if actual != nil || expected != nil {
			c.fail(t, actual, expected)
		}
		return
	}
	if !r.DeepEqual(actual, expected) {
		if r.TypeOf(actual) == r.TypeOf(expected) {
			if actualNode, ok := actual.(ast.Node); ok {
				if expectedNode, ok := expected.(ast.Node); ok {
					c.compareAst(t, ToAst(actualNode), ToAst(expectedNode))
					return
				}
			} else if actualUntyped, ok := actual.(untyped.Lit); ok {
				if expectedUntyped, ok := expected.(untyped.Lit); ok {
					c.compareUntyped(t, actualUntyped, expectedUntyped)
					return
				}
			} else if actualv.Kind() == r.Chan {
				// for channels just check the type, length and capacity
				expectedv := r.ValueOf(expected)
				if actualv.Len() == expectedv.Len() && actualv.Cap() == expectedv.Cap() {
					return
				}
			} else if actualv.Kind() == r.Func {
				// for functions just check the type
				return
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
	c.fail(t, actual.Interface(), expected.Interface())
}

func (c *TestCase) compareUntyped(t *testing.T, actual untyped.Lit, expected untyped.Lit) {
	if actual.Kind == expected.Kind && actual.Val.Kind() == expected.Val.Kind() && constant.Compare(actual.Val, token.EQL, expected.Val) {
		return
	}
	c.fail(t, actual, expected)
}

func (c *TestCase) fail(t *testing.T, actual interface{}, expected interface{}) {
	t.Errorf("expecting %v <%T>, found %v <%T>\n", expected, expected, actual, actual)
}
