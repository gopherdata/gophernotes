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
 * interface.go
 *
 *  Created on: Mar 29, 2017
 *      Author: Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) TypeInterface(node *ast.InterfaceType) xr.Type {
	if node.Methods == nil || len(node.Methods.List) == 0 {
		return c.TypeOfInterface()
	}
	methodtypes, methodnames := c.TypeFields(node.Methods)

	// TODO embedded interfaces
	return c.Universe.InterfaceOf(methodnames, methodtypes, nil)
}

func isInterfaceType(t xr.Type) bool {
	return t.Kind() == r.Interface
}
