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
 * stmt_0-3_test.go
 *
 *  Created on Apr 04, 2017
 *      Author Massimiliano Ghilardi
 */

package experiments

import (
	r "reflect"
	"testing"
)

type (
	Env0 struct {
		Binds []r.Value
		Outer *Env0
	}
	Stmt0 func(env *Env0, code []Stmt0) (Stmt0, *Env0)
)

func _BenchmarkStmt0(b *testing.B) {

	env := &Env0{
		Binds: make([]r.Value, 10),
	}
	code := make([]Stmt0, n+1)
	for i := 0; i < n; i++ {
		i := i
		code[i] = func(env *Env0, code []Stmt0) (Stmt0, *Env0) {
			return code[i+1], env
		}
	}
	code[n] = nil

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stmt := code[0]
		for stmt != nil {
			stmt, env = stmt(env, code)
		}
	}
}

type (
	Env1 struct {
		Binds []r.Value
		Outer *Env1
		IP    int
	}
	Stmt1 func(env *Env1, all []Stmt1) (Stmt1, *Env1)
)

func nop1(env *Env1, code []Stmt1) (Stmt1, *Env1) {
	env.IP++
	return code[env.IP], env
}

func _BenchmarkStmt1(b *testing.B) {

	env := &Env1{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt1, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop1
	}
	all[n] = nil

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		for stmt != nil {
			stmt, env = stmt(env, all)
		}
	}
}

type (
	Env2 struct {
		Binds []r.Value
		Outer *Env2
		Code  []Stmt2
	}
	Stmt2 func(env *Env2, ip int) (Stmt2, *Env2, int)
)

func nop2(env *Env2, ip int) (Stmt2, *Env2, int) {
	ip++
	return env.Code[ip], env, ip
}

func _BenchmarkStmt2(b *testing.B) {
	env := &Env2{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt2, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop2
	}
	all[n] = nil
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ip := 0
		stmt := all[ip]
		for stmt != nil {
			stmt, env, ip = stmt(env, ip)
		}
	}
}

type (
	Env3 struct {
		Binds []r.Value
		Outer *Env0
		Code  []Stmt3
	}
	Stmt3 func(env *Env3, ip int) (Stmt3, int)
)

func nop3(env *Env3, ip int) (Stmt3, int) {
	ip++
	return env.Code[ip], ip
}

func _BenchmarkStmt3(b *testing.B) {

	env := &Env3{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt3, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop3
	}
	all[n] = nil
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ip := 0
		stmt := all[ip]
		for stmt != nil {
			stmt, ip = stmt(env, ip)
		}
	}
}
