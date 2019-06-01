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
 * bench_test.go
 *
 *  Created on: Mar 06 2017
 *      Author: Massimiliano Ghilardi
 */
package main

import (
	"fmt"
	"os"
	r "reflect"
	"testing"

	"github.com/cosmos72/gomacro/classic"
	"github.com/cosmos72/gomacro/experiments/bytecode_interfaces"
	"github.com/cosmos72/gomacro/experiments/bytecode_values"
	"github.com/cosmos72/gomacro/experiments/closure_interfaces"
	"github.com/cosmos72/gomacro/experiments/closure_ints"
	"github.com/cosmos72/gomacro/experiments/closure_maps"
	"github.com/cosmos72/gomacro/experiments/closure_values"
	"github.com/cosmos72/gomacro/fast"
)

var (
	collatz_arg     = uint(837799) // sequence climbs to 1487492288, which also fits 32-bit ints
	collatz_arg_int = int(837799)
	sum_arg         = 1000
	fib_arg         = 12
	bigswitch_arg   = 100

	verbose = len(os.Args) == 0
)

/*
	--------- 2017-05-21: results on Intel Core i7 4770 ---------------

	BenchmarkFibonacciCompiler-8            	 3000000	       501 ns/op
	BenchmarkFibonacciFast-8                	  100000	     15774 ns/op
	BenchmarkFibonacciFast2-8               	  100000	     15141 ns/op
	BenchmarkFibonacciClassic-8             	    2000	    915990 ns/op
	BenchmarkFibonacciClassic2-8            	    2000	    912180 ns/op
	BenchmarkFibonacciClosureValues-8       	    5000	    259074 ns/op
	BenchmarkFibonacciClosureInterfaces-8   	   10000	    193098 ns/op
	BenchmarkFibonacciClosureMaps-8         	    5000	    358345 ns/op
	BenchmarkShellSortCompiler-8            	20000000	        74.0 ns/op
	BenchmarkShellSortFast-8                	  200000	      7790 ns/op
	BenchmarkShellSortClassic-8             	    5000	    276673 ns/op
	BenchmarkSwitchCompiler-8               	 1000000	      2363 ns/op
	BenchmarkSwitchFast-8                   	   50000	     37773 ns/op
	BenchmarkSwitchClassic-8                	     500	   3454461 ns/op
	BenchmarkArithCompiler1-8               	200000000	         8.41 ns/op
	BenchmarkArithCompiler2-8               	200000000	         8.41 ns/op
	BenchmarkArithFast-8                    	50000000	        30.8 ns/op
	BenchmarkArithFast2-8                   	30000000	        50.6 ns/op
	BenchmarkArithFastConst-8               	100000000	        15.2 ns/op
	BenchmarkArithFastCompileLoop-8         	  100000	     21442 ns/op
	BenchmarkArithClassic-8                 	 1000000	      1686 ns/op
	BenchmarkArithClassic2-8                	  500000	      2916 ns/op
	BenchmarkCollatzCompiler-8              	 5000000	       265 ns/op
	BenchmarkCollatzFast-8                  	  200000	     11812 ns/op
	BenchmarkCollatzClassic-8               	    2000	    654139 ns/op
	BenchmarkCollatzBytecodeInterfaces-8    	   50000	     30203 ns/op
	BenchmarkCollatzClosureValues-8         	  100000	     16570 ns/op
	BenchmarkSumCompiler-8                  	 5000000	       294 ns/op
	BenchmarkSumFast-8                      	  100000	     20789 ns/op
	BenchmarkSumFast2-8                     	  100000	     20720 ns/op
	BenchmarkSumClassic-8                   	    1000	   1223624 ns/op
	BenchmarkSumBytecodeValues-8            	   20000	     76201 ns/op
	BenchmarkSumBytecodeInterfaces-8        	   30000	     53031 ns/op
	BenchmarkSumClosureValues-8             	   30000	     41124 ns/op
	BenchmarkSumClosureInterfaces-8         	   10000	    147109 ns/op
	BenchmarkSumClosureMaps-8               	   20000	     93320 ns/op
*/

// ---------------------- recursion: fibonacci ----------------------

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func BenchmarkFibonacciCompiler(b *testing.B) {
	var total int
	n := fib_arg
	for i := 0; i < b.N; i++ {
		total += fibonacci(n)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkFibonacciFast(b *testing.B) {
	ir := fast.New()
	ir.Eval(fibonacci_source_string)

	// compile the call to fibonacci(fib_n)
	expr := ir.Compile(fmt.Sprintf("fibonacci(%d)", fib_arg))
	fun := expr.Fun.(func(*fast.Env) int)
	env := ir.PrepareEnv()

	fun(env) // warm up

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fun(env)
	}
}

func BenchmarkFibonacciFast2(b *testing.B) {
	ir := fast.New()
	ir.Eval(fibonacci_source_string)

	// alternative: extract the function fibonacci, and call it ourselves
	//
	// ValueOf is the method to retrieve constants, functions and variables from the classic and fast interpreters
	// (if you are going to read or write the same interpreter variable repeatedly,
	// dereferencing the address returned by AddressOfVar is faster)
	fun := ir.ValueOf("fibonacci").Interface().(func(int) int)

	fun(fib_arg) // warm up

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fun(fib_arg)
	}
}

func BenchmarkFibonacciClassic(b *testing.B) {
	ir := classic.New()
	ir.Eval(fibonacci_source_string)

	// compile the call to fibonacci(fib_n)
	form := ir.Parse(fmt.Sprintf("fibonacci(%d)", fib_arg))

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += int(ir.EvalAst1(form).Int())
	}
}

func BenchmarkFibonacciClassic2(b *testing.B) {
	ir := classic.New()
	ir.Eval(fibonacci_source_string)

	// alternative: extract the function fibonacci, and call it ourselves
	fun := ir.ValueOf("fibonacci").Interface().(func(int) int)

	fun(fib_arg) // warm up

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fun(fib_arg)
	}
}

func off_TestFibonacciClosureInts(t *testing.T) {
	env := closure_ints.NewEnv(nil)
	f := closure_ints.DeclFibonacci(env)

	expected := fibonacci(fib_arg)
	actual := f(fib_arg)
	if actual != expected {
		t.Errorf("expecting %v, found %v\n", expected, actual)
	}
}

func BenchmarkFibonacciClosureInts(b *testing.B) {
	env := closure_ints.NewEnv(nil)
	fib := closure_ints.DeclFibonacci(env)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(fib_arg)
	}
}

func BenchmarkFibonacciClosureValues(b *testing.B) {
	env := closure_values.NewEnv(nil)
	fib := closure_values.DeclFibonacci(env, 0)
	n := r.ValueOf(fib_arg)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(n)
	}
}

func BenchmarkFibonacciClosureInterfaces(b *testing.B) {
	env := closure_interfaces.NewEnv(nil)
	fib := closure_interfaces.DeclFibonacci(env, 0)
	var n interface{} = fib_arg

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(n)
	}
}

func BenchmarkFibonacciClosureMaps(b *testing.B) {
	env := closure_maps.NewEnv(nil)
	fib := closure_maps.DeclFibonacci(env, "fib")
	n := r.ValueOf(fib_arg)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(n)
	}
}

// ---------------------- switch ------------------------

func bigswitch(n int) int {
	for i := 0; i < 1000; i++ {
		switch n & 15 {
		case 0:
			n++
		case 1:
			n += 2
		case 2:
			n += 3
		case 3:
			n += 4
		case 4:
			n += 5
		case 5:
			n += 6
		case 6:
			n += 7
		case 7:
			n += 8
		case 8:
			n += 9
		case 9:
			n += 10
		case 10:
			n += 11
		case 11:
			n += 12
		case 12:
			n += 13
		case 13:
			n += 14
		case 14:
			n += 15
		case 15:
			n--
		}
	}
	return n
}

func BenchmarkSwitchCompiler(b *testing.B) {
	var total int
	for i := 0; i < b.N; i++ {
		total += bigswitch(bigswitch_arg)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkSwitchFast(b *testing.B) {
	ir := fast.New()
	ir.Eval(switch_source_string)

	fun := ir.ValueOf("bigswitch").Interface().(func(int) int)
	fun(bigswitch_arg)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fun(bigswitch_arg)
	}
}

func BenchmarkSwitchClassic(b *testing.B) {
	ir := classic.New()
	ir.Eval(switch_source_string)

	fun := ir.ValueOf("bigswitch").Interface().(func(int) int)
	fun(bigswitch_arg)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fun(bigswitch_arg)
	}
}

// ---------------- simple arithmetic ------------------

//go:noinline
func arith(n int) int {
	return ((((n*2 + 3) | 4) &^ 5) ^ 6) - ((n & 2) | 1)
}

const arith_source = "((((n*2+3)|4) &^ 5) ^ 6) - ((n & 2) | 1)"

func BenchmarkArithCompiler1(b *testing.B) {
	total := 0
	for i := 0; i < b.N; i++ {
		n := b.N
		total += ((((n*2 + 3) | 4) &^ 5) ^ 6) - ((n & 2) | 1)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkArithCompiler2(b *testing.B) {
	total := 0
	for i := 0; i < b.N; i++ {
		total += arith(b.N)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkArithFast(b *testing.B) {
	ir := fast.New()
	ir.DeclVar("n", nil, int(0))

	addr := ir.AddressOfVar("n").Interface().(*int)

	expr := ir.Compile(arith_source)
	fun := expr.Fun.(func(*fast.Env) int)
	env := ir.PrepareEnv()
	fun(env)

	// interpreted code performs only arithmetic - iteration performed here
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		*addr = b.N
		total += fun(env)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkArithFast2(b *testing.B) {
	ir := fast.New()
	ir.Eval("var i, n, total int")

	n := ir.AddressOfVar("n").Interface().(*int)
	total := ir.AddressOfVar("total").Interface().(*int)

	// interpreted code performs iteration and arithmetic
	fun := ir.Compile("for i = 0; i < n; i++ { total += " + arith_source + " }").AsX()
	env := ir.PrepareEnv()
	fun(env)

	b.ResetTimer()

	*n = b.N
	*total = 0
	fun(env)

	if verbose {
		println(*total)
	}
}

func BenchmarkArithFastConst(b *testing.B) {
	ir := fast.New()
	// "cheat" a bit and declare n as a constant. checks if constant propagation works :)
	ir.DeclConst("n", nil, b.N)

	// interpreted code performs only arithmetic - iteration performed here
	expr := ir.Compile(arith_source)
	fun := expr.WithFun().(func(*fast.Env) int)
	env := ir.PrepareEnv()
	fun(env)

	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		total += fun(env)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkArithFastConst2(b *testing.B) {
	ir := fast.New()
	ir.Eval("var i, total int")
	// "cheat" a bit and declare n as a constant. checks if constant propagation works :)
	ir.DeclConst("n", nil, int(b.N))
	total := ir.AddressOfVar("total").Interface().(*int)

	// interpreted code performs iteration and arithmetic
	fun := ir.Compile("for i = 0; i < n; i++ { total += " + arith_source + " }").AsX()
	env := ir.PrepareEnv()
	fun(env)

	b.ResetTimer()

	*total = 0
	fun(env)

	if verbose {
		println(*total)
	}
}

func BenchmarkArithFastCompileLoop(b *testing.B) {
	ir := fast.New()
	ir.Eval("var i, n, total int")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ir.Compile("total = 0; for i = 0; i < n; i++ { total += " + arith_source + " }; total")
	}
}

func BenchmarkArithClassic(b *testing.B) {
	ir := classic.New()
	ir.Eval("n:=0")

	form := ir.Parse(arith_source)

	value := ir.ValueOf("n")
	var ret r.Value
	ir.EvalAst(form)

	// interpreted code performs only arithmetic - iteration performed here
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		value.SetInt(int64(b.N))
		ret, _ = ir.EvalAst(form)
		total += int(ret.Int())
	}
	if verbose {
		println(total)
	}
}

func BenchmarkArithClassic2(b *testing.B) {
	ir := classic.New()
	ir.Eval("var n, total int")

	// interpreted code performs iteration and arithmetic
	form := ir.Parse("total = 0; for i:= 0; i < n; i++ { total += " + arith_source + " }; total")

	value := ir.ValueOf("n")
	ir.EvalAst(form)

	b.ResetTimer()
	value.SetInt(int64(b.N))
	ret, _ := ir.EvalAst(form)

	if verbose {
		println(ret.Int())
	}
}

// ---------------- collatz conjecture --------------------

func collatz(n uint) uint {
	for n > 1 {
		if n&1 != 0 {
			n = ((n * 3) + 1) >> 1
		} else {
			n >>= 1
		}
	}
	return n
}

func BenchmarkCollatzCompiler(b *testing.B) {
	var n uint = collatz_arg
	for i := 0; i < b.N; i++ {
		collatz(n)
	}
}

func BenchmarkCollatzFast(b *testing.B) {
	ir := fast.New()
	ir.DeclVar("n", nil, uint(0))
	addr := ir.AddressOfVar("n").Interface().(*uint)

	fun := ir.Compile("for n > 1 { if n&1 != 0 { n = ((n * 3) + 1) >> 1 } else { n >>= 1 } }").AsX()
	env := ir.PrepareEnv()
	fun(env)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		*addr = collatz_arg
		fun(env)
	}
}

func BenchmarkCollatzClassic(b *testing.B) {
	ir := classic.New()
	ir.EvalAst(ir.Parse("var n uint"))
	addr := ir.ValueOf("n").Addr().Interface().(*uint)

	form := ir.Parse("for n > 1 { if n&1 != 0 { n = ((n * 3) + 1) >> 1 } else { n >>= 1 } }")
	ir.EvalAst(form)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		*addr = collatz_arg
		ir.EvalAst(form)
	}
}

func BenchmarkCollatzBytecodeInterfaces(b *testing.B) {
	coll := bytecode_interfaces.BytecodeCollatz()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Vars[0] = collatz_arg_int
		coll.Exec(0)
	}
}

func off_TestCollatzClosureInts(t *testing.T) {
	env := closure_ints.NewEnv(nil)
	f := closure_ints.DeclCollatz(env)

	expected := int(collatz(collatz_arg))
	actual := f(collatz_arg_int)
	if actual != expected {
		t.Errorf("expecting %v, found %v\n", expected, actual)
	}
}

func BenchmarkCollatzClosureInts(b *testing.B) {
	env := closure_ints.NewEnv(nil)
	coll := closure_ints.DeclCollatz(env)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += coll(collatz_arg_int)
	}
}

func BenchmarkCollatzClosureValues(b *testing.B) {
	env := closure_values.NewEnv(nil)
	coll := closure_values.DeclCollatz(env, 0)
	n := r.ValueOf(collatz_arg_int)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += coll(n)
	}
}

// ------------- looping: sum the integers from 1 to N -------------------

func sum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

func BenchmarkSumCompiler(b *testing.B) {
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(sum_arg)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkSumFast(b *testing.B) {
	ir := fast.New()
	ir.Eval("var i, total uint")
	ir.DeclConst("n", nil, uint(sum_arg))

	expr := ir.Compile("total = 0; for i = 1; i <= n; i++ { total += i }; total")
	fun := expr.Fun.(func(*fast.Env) uint)
	env := ir.PrepareEnv()
	fun(env)

	var total uint
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += fun(env)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkSumFast2(b *testing.B) {
	ir := fast.New()
	ir.Eval("var i, total uint")
	ir.DeclConst("n", nil, uint(sum_arg))

	fun := ir.Compile("for i = 1; i <= n; i++ { total += i }").AsX()
	env := ir.PrepareEnv()
	fun(env)
	total := ir.AddressOfVar("total").Interface().(*uint)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		*total = 0
		fun(env)
	}
	if verbose {
		println(*total)
	}
}

func BenchmarkSumClassic(b *testing.B) {
	ir := classic.New()
	ir.Eval("var i, n, total int")
	ir.ValueOf("n").SetInt(int64(sum_arg))
	form := ir.Parse("total = 0; for i = 1; i <= n; i++ { total += i }; total")

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += int(ir.EvalAst1(form).Int())
	}
}

func BenchmarkSumBytecodeValues(b *testing.B) {
	sum := bytecode_values.BytecodeSum(sum_arg)
	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += int(sum.Exec(0)[0].Int())
	}
}

func BenchmarkSumBytecodeInterfaces(b *testing.B) {
	p := bytecode_interfaces.BytecodeSum(sum_arg)
	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += p.Exec(0)[0].(int)
	}
}

func off_TestSumClosureInts(t *testing.T) {
	env := closure_ints.NewEnv(nil)
	f := closure_ints.DeclSum(env)

	expected := sum(sum_arg)
	actual := f(sum_arg)
	if actual != expected {
		t.Errorf("expecting %v, found %v\n", expected, actual)
	}
}

func BenchmarkSumClosureInts(b *testing.B) {
	env := closure_ints.NewEnv(nil)
	sum := closure_ints.DeclSum(env)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(sum_arg)
	}
}

func BenchmarkSumClosureValues(b *testing.B) {
	env := closure_values.NewEnv(nil)
	sum := closure_values.DeclSum(env, 0)
	n := r.ValueOf(sum_arg)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(n)
	}
}

func BenchmarkSumClosureInterfaces(b *testing.B) {
	env := closure_interfaces.NewEnv(nil)
	sum := closure_interfaces.DeclSum(env, 0)
	var n interface{} = sum_arg

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(n)
	}
}

func BenchmarkSumClosureMaps(b *testing.B) {
	env := closure_maps.NewEnv(nil)
	sum := closure_maps.DeclSum(env, "sum")
	n := r.ValueOf(sum_arg)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(n)
	}
}
