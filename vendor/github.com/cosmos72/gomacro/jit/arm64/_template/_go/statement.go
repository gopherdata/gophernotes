/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * statement.go
 *
 *  Created on Feb 16, 2019
 *      Author Massimiliano Ghilardi
 */

package main

func Nop(env *Env) (Stmt, *Env) {
	ip := env.IP + 1
	env.IP = ip
	return env.Code[ip], env
}

func Leave(env *Env) (Stmt, *Env) {
	env = env.Outer
	ip := env.IP + 1
	env.IP = ip
	return env.Code[ip], env
}
