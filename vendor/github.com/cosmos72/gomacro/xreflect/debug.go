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
