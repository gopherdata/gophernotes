// +build gomacro_jit

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * bench_jit_test.go
 *
 *  Created on: Jun 06 2018
 *      Author: Massimiliano Ghilardi
 */
package main

import (
	"testing"
	"unsafe"

	"github.com/cosmos72/gomacro/experiments/jit"
)

func arithJitEmulate(uenv *uint64) {
	env := (*[3]int64)(unsafe.Pointer(uenv))
	a := env[0]
	a *= 2
	a += 3
	a |= 4
	a &^= 5
	a ^= 6
	b := env[0]
	b &= 2
	b |= 1
	a /= b
	env[1] = a
}

func BenchmarkArithJitEmul(b *testing.B) {
	benchArithJit(b, arithJitEmulate)
}

func BenchmarkArithJit(b *testing.B) {
	if !jit.SUPPORTED {
		b.SkipNow()
		return
	}
	benchArithJit(b, jit.DeclArith(5))
}

func benchArithJit(b *testing.B, f func(*uint64)) {
	total := 0
	var env [5]uint64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env[0] = uint64(b.N)
		f(&env[0])
		total += int(env[1])
	}
	if verbose {
		println(total)
	}
}

// --------------------------------------------------------------

func BenchmarkSumJit(b *testing.B) {
	if !jit.SUPPORTED {
		b.SkipNow()
		return
	}
	sum := jit.DeclSum()
	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(sum_arg)
	}
}
