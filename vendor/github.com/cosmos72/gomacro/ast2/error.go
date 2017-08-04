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
