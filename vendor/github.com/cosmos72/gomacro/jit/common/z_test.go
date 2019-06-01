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
 * z_test.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"testing"
)

func EqUint8(t *testing.T, actual uint8, expected uint8) {
	if actual != expected {
		t.Errorf("expected %d,\tactual %d", expected, actual)
	}
}

func TestLog2(t *testing.T) {
	for shift := uint8(1); shift < 64; shift++ {
		n := uint64(1) << shift
		actual, _ := Log2Uint(n)
		EqUint8(t, actual, shift)
		actual, _ = Log2Uint(n - 1)
		EqUint8(t, actual, 0)
		actual, _ = Log2Uint(n + 1)
		EqUint8(t, actual, 0)
	}
}
