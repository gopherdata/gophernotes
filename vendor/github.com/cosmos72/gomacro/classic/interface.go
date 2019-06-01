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
