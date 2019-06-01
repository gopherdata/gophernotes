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
 * output.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"errors"
	"fmt"
)

func debugf(format string, args ...interface{}) {
	fmt.Printf("// debug jit: "+format+"\n", args...)
}

var errorPrefix = "assembler error: "

func errorf(format string, args ...interface{}) {
	panic(errors.New(errorPrefix + fmt.Sprintf(format, args...)))
}
