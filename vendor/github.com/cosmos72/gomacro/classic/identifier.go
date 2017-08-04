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
 * identifier.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (env *Env) evalIdentifier(ident *ast.Ident) r.Value {
	value, found := env.resolveIdentifier(ident)
	if !found {
		env.Errorf("undefined identifier: %s", ident.Name)
	}
	return value
}

func (env *Env) resolveIdentifier(ident *ast.Ident) (r.Value, bool) {
	name := ident.Name
	value := Nil
	found := false
	for e := env; e != nil; e = e.Outer {
		// Debugf("evalIdentifier() looking up %#v in %#v", name, env.Binds)
		if value, found = e.Binds.Get(name); found {
			break
		}
	}
	return value, found
}
