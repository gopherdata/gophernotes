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
 * util.go
 *
 *  Created on Feb 24, 2019
 *      Author Massimiliano Ghilardi
 */

package common

func sliceEqual(lhs []uint8, rhs []uint8) bool {
	size := len(lhs)
	if size != len(rhs) {
		return false
	}
	if size == 0 || &lhs[0] == &rhs[0] {
		return true
	}
	for i := 0; i < size; i++ {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}

func (code MachineCode) Equal(other MachineCode) bool {
	return code.ArchId == other.ArchId && sliceEqual(code.Bytes, other.Bytes)
}
