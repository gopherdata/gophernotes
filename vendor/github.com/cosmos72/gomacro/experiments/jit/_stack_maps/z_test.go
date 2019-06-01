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
 * fib_asm.go
 *
 *  Created on May 23, 2018
 *      Author Massimiliano Ghilardi
 */

package stack_maps

import "testing"

func TestFib(t *testing.T) {
	const in = 10
	out1, out2 := fib(in), fib_asm(in)
	if out1 != out2 {
		t.Errorf("fib_asm(%d) = %d, expecting %d", in, out2, out1)
	}
}
