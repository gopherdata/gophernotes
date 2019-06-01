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
 * debug.go
 *
 *  Created on Apr 04, 2018
 *      Author Massimiliano Ghilardi
 */

package xreflect

import "fmt"

func debugf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	fmt.Printf("// debug: %s\n", str)
}

func (v *Universe) debugf(format string, args ...interface{}) {
	depth := v.DebugDepth
	if depth == 0 {
		return
	}
	depth = depth*2 - 2
	const dots = ". . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . "
	pad := make([]byte, depth)
	for i := 0; i < depth; i += depth {
		copy(pad[i:], dots)
	}
	format = "// debug: %s" + format + "\n"
	args = append([]interface{}{pad}, args...)
	fmt.Printf(format, args...)
}

func (v *Universe) debug() bool {
	return v.DebugDepth != 0
}

func de(v *Universe) {
	v.DebugDepth--
}

func bug(v *Universe) *Universe {
	v.DebugDepth++
	return v
}
