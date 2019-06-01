// +build gomacro_fast_compact

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
 * callnret0_compact.go
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
