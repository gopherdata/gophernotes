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
 * func0ret0.go
 *
 *  Created on Apr 16, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"

	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) func0ret0(t xr.Type, m *funcMaker) func(env *Env) r.Value {
	funcbody := m.funcbody
	if funcbody == nil {
		return func(env *Env) r.Value {
			return valueOfNopFunc
		}
	}

	nbinds := m.nbinds
	nintbinds := m.nintbinds
	return func(env *Env) r.Value {
		// function is closed over the env used to DECLARE it
		env.MarkUsedByClosure()
		return r.ValueOf(func() {
			env := newEnv4Func(env, nbinds, nintbinds)
			// execute the body
			funcbody(env)

			env.FreeEnv()
		})
	}
}
