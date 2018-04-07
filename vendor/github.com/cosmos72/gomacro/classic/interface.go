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

package classic

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)


func (env *Env) evalTypeInterface(node *ast.InterfaceType) r.Type {
	if node.Methods == nil || len(node.Methods.List) == 0 {
		return TypeOfInterface
	}
	types, names := env.evalTypeFields(node.Methods)

	types = append([]r.Type{TypeOfInterface}, types...)
	names = append([]string{StrGensymInterface}, names...)

	fields := makeStructFields(env.FileEnv().Path, names, types)
	return r.StructOf(fields)
}

func isInterfaceType(t r.Type) bool {
	if t.Kind() == r.Struct && t.NumField() > 0 {
		field := t.Field(0)
		return field.Name == StrGensymInterface && field.Type == TypeOfInterface
	}
	return false
}
