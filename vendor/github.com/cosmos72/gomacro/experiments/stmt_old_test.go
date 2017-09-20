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
 * stmt_test.go
 *
 *  Created on Apr 04, 2017
 *      Author Massimiliano Ghilardi
 */

package experiments

import (
	r "reflect"
	"testing"
	"unsafe"
)

type (
	Env0 struct {
		Binds []r.Value
		Outer *Env0
	}
	Stmt0 func(env *Env0, code []Stmt0) (Stmt0, *Env0)
)

func BenchmarkThreadedStmtFunc0(b *testing.B) {

	env := &Env0{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt0, n+1)
	for i := 0; i < n; i++ {
		i := i
		all[i] = func(env *Env0, code []Stmt0) (Stmt0, *Env0) {
			return code[i+1], env
		}
	}
	all[n] = nil

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stmt := all[0]
		for stmt != nil {
			stmt, env = stmt(env, all)
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

func BenchmarkThreadedStmtFunc1(b *testing.B) {

	var nop Stmt1 = func(env *Env1, code []Stmt1) (Stmt1, *Env1) {
		env.IP++
		return code[env.IP], env
	}

	env := &Env1{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt1, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
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

func BenchmarkThreadedStmtFunc2(b *testing.B) {

	var nop Stmt2 = func(env *Env2, ip int) (Stmt2, *Env2, int) {
		ip++
		return env.Code[ip], env, ip
	}

	env := &Env2{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt2, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
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

func BenchmarkThreadedStmtFunc3(b *testing.B) {

	var nop Stmt3 = func(env *Env3, ip int) (Stmt3, int) {
		ip++
		return env.Code[ip], ip
	}
	env := &Env3{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt3, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
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

type (
	Env4 struct {
		Binds     []r.Value
		Outer     *Env4
		IP        int
		Code      []Stmt4
		Interrupt Stmt4
	}
	Stmt4 func(env *Env4) Stmt4
)

func BenchmarkThreadedStmtFunc4(b *testing.B) {
	var nop Stmt4 = func(env *Env4) Stmt4 {
		env.IP++
		return env.Code[env.IP]
	}
	env := &Env4{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt4, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = nil
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		for stmt != nil {
			stmt = stmt(env)
		}
	}
}

func BenchmarkThreadedStmtFunc4Unroll(b *testing.B) {
	var nop Stmt4 = func(env *Env4) Stmt4 {
		env.IP++
		return env.Code[env.IP]
	}
	env := &Env4{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt4, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = nil
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		for stmt != nil {
			if stmt = stmt(env); stmt != nil {
				if stmt = stmt(env); stmt != nil {
					if stmt = stmt(env); stmt != nil {
						if stmt = stmt(env); stmt != nil {
							if stmt = stmt(env); stmt != nil {
								if stmt = stmt(env); stmt != nil {
									if stmt = stmt(env); stmt != nil {
										if stmt = stmt(env); stmt != nil {
											if stmt = stmt(env); stmt != nil {
												if stmt = stmt(env); stmt != nil {
													if stmt = stmt(env); stmt != nil {
														if stmt = stmt(env); stmt != nil {
															if stmt = stmt(env); stmt != nil {
																if stmt = stmt(env); stmt != nil {
																	stmt = stmt(env)
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func BenchmarkThreadedStmtFunc4Terminate(b *testing.B) {
	var interrupt Stmt4
	interrupt = func(env *Env4) Stmt4 {
		return interrupt
	}
	unsafeInterrupt := *(**uintptr)(unsafe.Pointer(&interrupt))

	var nop Stmt4 = func(env *Env4) Stmt4 {
		env.IP++
		return env.Code[env.IP]
	}
	env := &Env4{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt4, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = interrupt
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		for {
			if x := stmt; *(**uintptr)(unsafe.Pointer(&x)) == unsafeInterrupt {
				break
			}
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
		}
	}
}

func BenchmarkThreadedStmtFunc4Adaptive(b *testing.B) {
	var nop Stmt4 = func(env *Env4) Stmt4 {
		env.IP++
		return env.Code[env.IP]
	}
	var interrupt Stmt4 = func(env *Env4) Stmt4 {
		return env.Interrupt
	}
	unsafeInterrupt := *(**uintptr)(unsafe.Pointer(&interrupt))
	_ = unsafeInterrupt
	env := &Env4{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt4, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = nil
	env.Code = all

	b.ResetTimer()
outer:
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		if stmt == nil {
			continue outer
		}
		for j := 0; j < 5; j++ {
			if stmt = stmt(env); stmt != nil {
				if stmt = stmt(env); stmt != nil {
					if stmt = stmt(env); stmt != nil {
						if stmt = stmt(env); stmt != nil {
							if stmt = stmt(env); stmt != nil {
								if stmt = stmt(env); stmt != nil {
									if stmt = stmt(env); stmt != nil {
										if stmt = stmt(env); stmt != nil {
											if stmt = stmt(env); stmt != nil {
												if stmt = stmt(env); stmt != nil {
													if stmt = stmt(env); stmt != nil {
														if stmt = stmt(env); stmt != nil {
															if stmt = stmt(env); stmt != nil {
																if stmt = stmt(env); stmt != nil {
																	continue
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
			continue outer
		}
		all[n] = interrupt
		env.Interrupt = interrupt
		for {
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)

			if x := stmt; *(**uintptr)(unsafe.Pointer(&x)) == unsafeInterrupt {
				continue outer
			}
		}
	}
}

func BenchmarkThreadedStmtFunc4Panic(b *testing.B) {
	var terminate Stmt4 = func(env *Env4) Stmt4 {
		panic("end of code")
	}

	var nop Stmt4 = func(env *Env4) Stmt4 {
		env.IP++
		return env.Code[env.IP]
	}
	env := &Env4{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt4, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = terminate
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runThreadedStmtFunc4Panic(env)
	}
}

func runThreadedStmtFunc4Panic(env *Env4) {
	env.IP = 0
	stmt := env.Code[0]
	defer func() {
		recover()
	}()
	for {
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
	}
}

type (
	Env5 struct {
		Binds []r.Value
		IP    int
		Code  []Stmt5
		Outer *Env5
	}
	Stmt5 func(**Env5) Stmt5
)

func BenchmarkThreadedStmtFunc5(b *testing.B) {

	var nop Stmt5 = func(penv **Env5) Stmt5 {
		env := *penv
		env.IP++
		return env.Code[env.IP]
	}

	env := &Env5{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt5, n+1)
	for i := 0; i < n; i++ {
		i := i
		all[i] = nop
	}
	all[n] = nil
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		for stmt != nil {
			stmt = stmt(&env)
		}
	}
}
