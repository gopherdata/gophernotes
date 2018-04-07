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
 * gensym.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"fmt"
        
)

// the following constants must match with github.com/cosmos72/gomacro/base/constants.go
const (
	StrGensymInterface = "\U0001202A" // name of extra struct field needed by the interpreter when creating interpreted interfaces
	StrGensymPrivate   = "\U00012038" // prefix to generate names for unexported struct fields.
	StrGensymEmbedded  = "\U00012039" // prefix to generate names for embedded struct fields.
)

var gensymn = 0

func GensymEmbedded(name string) string {
	if len(name) == 0 {
		n := gensymn
		gensymn++
		name = fmt.Sprintf("%d", n)
	}
	return StrGensymEmbedded + name
}

func GensymPrivate(name string) string {
	if len(name) == 0 {
		n := gensymn
		gensymn++
		name = fmt.Sprintf("%d", n)
	}
	return StrGensymPrivate + name
}
