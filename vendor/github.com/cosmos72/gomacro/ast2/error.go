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
 * error.go
 *
 *  Created on: Mar 18, 2017
 *      Author: Massimiliano Ghilardi
 */

package ast2

import (
	"errors"
	"fmt"
)

func badIndex(index int, size int) AstWithNode {
	if size > 0 {
		errorf("index out of range: %d not in 0...%d", index, size-1)
	} else {
		errorf("index out of range: %d, slice is empty", index)
	}
	return nil
}

func errorf(format string, args ...interface{}) {
	panic(errors.New(fmt.Sprintf(format, args...)))
}
