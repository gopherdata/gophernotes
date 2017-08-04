// +build gomacro_fast_compact

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
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *
 * callnret0compact.go
 *
 *  Created on Jun 14, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
)

func (c *Comp) call0ret0(call *Call, maxdepth int) func(env *Env) {
	expr := call.Fun
	exprfun := expr.AsX1()
	return func(env *Env) {
		fun := exprfun(env).Interface().(func())
		fun()
	}
}

func (c *Comp) call1ret0(call *Call, maxdepth int) func(env *Env) {
	expr := call.Fun
	exprfun := expr.AsX1()

	argfun := call.MakeArgfunsX1()[0]

	return func(env *Env) {
		funv := exprfun(env)
		argv := []r.Value{argfun(env)}

		funv.Call(argv)
	}
}

func (c *Comp) call2ret0(call *Call, maxdepth int) func(env *Env) {
	expr := call.Fun
	exprfun := expr.AsX1()

	argfunsX1 := call.MakeArgfunsX1()
	argfuns := [2]func(*Env) r.Value{
		argfunsX1[0],
		argfunsX1[1],
	}

	return func(env *Env) {
		funv := exprfun(env)

		argv := []r.Value{
			argfuns[0](env),
			argfuns[1](env),
		}
		funv.Call(argv)
	}
}
