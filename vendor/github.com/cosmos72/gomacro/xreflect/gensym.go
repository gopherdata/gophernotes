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
	StrGensymAnonymous = "\U00012039" // prefix to generate names for anonymous struct fields.
)

var gensymn = 0

func GensymAnonymous(name string) string {
	if len(name) != 0 {
		return StrGensymAnonymous + name
	}
	n := gensymn
	gensymn++
	return fmt.Sprintf("%s%d", StrGensymAnonymous, n)
}

func GensymPrivate(name string) string {
	if len(name) != 0 {
		return StrGensymPrivate + name
	}
	n := gensymn
	gensymn++
	return fmt.Sprintf("%s%d", StrGensymPrivate, n)
}
