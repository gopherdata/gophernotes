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
 * stmt_4-5_test.go
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
	Env4 struct {
		Binds     []r.Value
		Outer     *Env4
		IP        int
		Code      []Stmt4
		Interrupt Stmt4
		Signal    int
	}
	Stmt4 func(env *Env4) Stmt4
)

func nop4(env *Env4) Stmt4 {
	env.IP++
	return env.Code[env.IP]
}

func interrupt4(env *Env4) Stmt4 {
	env.Signal = 1
	return env.Interrupt
}

func newEnv4() *Env4 {
	code := make([]Stmt4, n+1)
	for i := 0; i < n; i++ {
		code[i] = nop4
	}
	code[n] = nil
	return &Env4{
		Binds: make([]r.Value, 10),
		Code:  code,
	}
}

func BenchmarkStmt4(b *testing.B) {
	env := newEnv4()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := env.Code[0]
		for {
			if stmt = stmt(env); stmt == nil {
				break
			}
		}
	}
}

func BenchmarkStmt4Unroll(b *testing.B) {
	env := newEnv4()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := env.Code[0]
		for {
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
			break
		}
	}
}

func BenchmarkStmt4Spin(b *testing.B) {
	env := newEnv4()
	env.Code[n] = interrupt4
	env.Interrupt = interrupt4

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		env.Signal = 0
		stmt := env.Code[0]
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
			if env.Signal != 0 {
				break
			}
		}
	}
}

func BenchmarkStmt4Adaptive13(b *testing.B) {
	env := newEnv4()
	env.Code[n] = interrupt4

	b.ResetTimer()
outer:
	for i := 0; i < b.N; i++ {
		env.IP = 0
		env.Interrupt = nil
		env.Signal = 0
		stmt := env.Code[0]
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
			continue outer
		}
		env.Interrupt = interrupt4
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

			if env.Signal != 0 {
				continue outer
			}
		}
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

func BenchmarkStmt5(b *testing.B) {

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
		for {
			if stmt = stmt(&env); stmt == nil {
				break
			}
		}
	}
}
