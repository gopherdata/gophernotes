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
 * stmt_6_test.go
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
	Env6 struct {
		Binds     []r.Value
		Outer     *Env6
		Code      []Stmt6
		IP        int
		Signal    int
		Interrupt Stmt6
	}
	Stmt6 func(*Env6) (Stmt6, *Env6)
	X6    func(*Env6)
)

func nop6(env *Env6) (Stmt6, *Env6) {
	env.IP++
	return env.Code[env.IP], env
}

func interrupt6(env *Env6) (Stmt6, *Env6) {
	env.Signal = 1
	return env.Interrupt, env
}

func newEnv6() *Env6 {
	code := make([]Stmt6, n+1)
	for i := 0; i < n; i++ {
		code[i] = nop6
	}
	code[n] = nil
	return &Env6{
		Binds: make([]r.Value, 10),
		Code:  code,
	}
}

func BenchmarkStmt6(b *testing.B) {
	env := newEnv6()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := env.Code[0]
		for {
			if stmt, env = stmt(env); stmt == nil {
				break
			}
		}
	}
}

func BenchmarkStmt6Unroll(b *testing.B) {
	env := newEnv6()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := env.Code[0]
		for {
			if stmt, env = stmt(env); stmt != nil {
				if stmt, env = stmt(env); stmt != nil {
					if stmt, env = stmt(env); stmt != nil {
						if stmt, env = stmt(env); stmt != nil {
							if stmt, env = stmt(env); stmt != nil {
								if stmt, env = stmt(env); stmt != nil {
									if stmt, env = stmt(env); stmt != nil {
										if stmt, env = stmt(env); stmt != nil {
											if stmt, env = stmt(env); stmt != nil {
												if stmt, env = stmt(env); stmt != nil {
													if stmt, env = stmt(env); stmt != nil {
														if stmt, env = stmt(env); stmt != nil {
															if stmt, env = stmt(env); stmt != nil {
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

func BenchmarkStmt6Spin(b *testing.B) {

	env := newEnv6()
	env.Interrupt = interrupt6
	env.Code[n] = interrupt6

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		env.Signal = 0
		stmt := env.Code[0]
		for {
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			if env.Signal != 0 {
				break
			}
		}
	}
}

func BenchmarkStmt6Adaptive13(b *testing.B) {
	env := newEnv6()

	b.ResetTimer()
outer:
	for i := 0; i < b.N; i++ {
		env.IP = 0
		env.Signal = 0
		stmt := env.Code[0]
		if stmt == nil {
			continue outer
		}
		for j := 0; j < 5; j++ {
			if stmt, env = stmt(env); stmt != nil {
				if stmt, env = stmt(env); stmt != nil {
					if stmt, env = stmt(env); stmt != nil {
						if stmt, env = stmt(env); stmt != nil {
							if stmt, env = stmt(env); stmt != nil {
								if stmt, env = stmt(env); stmt != nil {
									if stmt, env = stmt(env); stmt != nil {
										if stmt, env = stmt(env); stmt != nil {
											if stmt, env = stmt(env); stmt != nil {
												if stmt, env = stmt(env); stmt != nil {
													if stmt, env = stmt(env); stmt != nil {
														if stmt, env = stmt(env); stmt != nil {
															if stmt, env = stmt(env); stmt != nil {
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

		env.Code[n] = interrupt6
		env.Interrupt = interrupt6
		for {
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)

			if env.Signal != 0 {
				continue outer
			}
		}
	}
}

type (
	EnvS6 struct {
		Binds     []r.Value
		Outer     *EnvS6
		IP        int
		Code      []StmtS6
		Interrupt StmtS6
	}
	StmtS6 struct {
		Exec func(env *EnvS6) (StmtS6, *EnvS6)
	}
)

func _BenchmarkStmtStruct6(b *testing.B) {

	var nop StmtS6 = StmtS6{func(env *EnvS6) (StmtS6, *EnvS6) {
		env.IP++
		return env.Code[env.IP], env
	}}
	env := &EnvS6{
		Binds: make([]r.Value, 10),
	}
	all := make([]StmtS6, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = StmtS6{}
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		for stmt.Exec != nil {
			stmt, env = stmt.Exec(env)
		}
	}
}

func _BenchmarkStmtStruct6Unroll(b *testing.B) {

	var nop StmtS6 = StmtS6{func(env *EnvS6) (StmtS6, *EnvS6) {
		env.IP++
		return env.Code[env.IP], env
	}}
	env := &EnvS6{
		Binds: make([]r.Value, 10),
	}
	all := make([]StmtS6, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = StmtS6{}
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		for stmt.Exec != nil {
			if stmt, env = stmt.Exec(env); stmt.Exec != nil {
				if stmt, env = stmt.Exec(env); stmt.Exec != nil {
					if stmt, env = stmt.Exec(env); stmt.Exec != nil {
						if stmt, env = stmt.Exec(env); stmt.Exec != nil {
							if stmt, env = stmt.Exec(env); stmt.Exec != nil {
								if stmt, env = stmt.Exec(env); stmt.Exec != nil {
									if stmt, env = stmt.Exec(env); stmt.Exec != nil {
										if stmt, env = stmt.Exec(env); stmt.Exec != nil {
											if stmt, env = stmt.Exec(env); stmt.Exec != nil {
												if stmt, env = stmt.Exec(env); stmt.Exec != nil {
													if stmt, env = stmt.Exec(env); stmt.Exec != nil {
														if stmt, env = stmt.Exec(env); stmt.Exec != nil {
															if stmt, env = stmt.Exec(env); stmt.Exec != nil {
																if stmt, env = stmt.Exec(env); stmt.Exec != nil {
																	stmt, env = stmt.Exec(env)
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

func _BenchmarkStmtStruct6Spin(b *testing.B) {

	var nop StmtS6 = StmtS6{func(env *EnvS6) (StmtS6, *EnvS6) {
		env.IP++
		return env.Code[env.IP], env
	}}
	var interrupt StmtS6 = StmtS6{func(env *EnvS6) (StmtS6, *EnvS6) {
		return env.Interrupt, env
	}}
	unsafeInterrupt := *(**uintptr)(unsafe.Pointer(&interrupt))

	env := &EnvS6{
		Binds: make([]r.Value, 10),
	}
	all := make([]StmtS6, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = interrupt
	env.Code = all
	env.Interrupt = interrupt

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		for {
			if x := stmt; *(**uintptr)(unsafe.Pointer(&x)) == unsafeInterrupt {
				break
			}
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
		}
	}
}

func _BenchmarkStmtStruct6Adaptive13(b *testing.B) {

	var nop StmtS6 = StmtS6{func(env *EnvS6) (StmtS6, *EnvS6) {
		env.IP++
		return env.Code[env.IP], env
	}}
	var interrupt StmtS6 = StmtS6{func(env *EnvS6) (StmtS6, *EnvS6) {
		return env.Interrupt, env
	}}
	unsafeInterrupt := *(**uintptr)(unsafe.Pointer(&interrupt))

	env := &EnvS6{
		Binds: make([]r.Value, 10),
	}
	all := make([]StmtS6, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = StmtS6{}
	env.Code = all

	b.ResetTimer()
outer:
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		if stmt.Exec == nil {
			continue outer
		}
		for j := 0; j < 5; j++ {
			if stmt, env = stmt.Exec(env); stmt.Exec != nil {
				if stmt, env = stmt.Exec(env); stmt.Exec != nil {
					if stmt, env = stmt.Exec(env); stmt.Exec != nil {
						if stmt, env = stmt.Exec(env); stmt.Exec != nil {
							if stmt, env = stmt.Exec(env); stmt.Exec != nil {
								if stmt, env = stmt.Exec(env); stmt.Exec != nil {
									if stmt, env = stmt.Exec(env); stmt.Exec != nil {
										if stmt, env = stmt.Exec(env); stmt.Exec != nil {
											if stmt, env = stmt.Exec(env); stmt.Exec != nil {
												if stmt, env = stmt.Exec(env); stmt.Exec != nil {
													if stmt, env = stmt.Exec(env); stmt.Exec != nil {
														if stmt, env = stmt.Exec(env); stmt.Exec != nil {
															if stmt, env = stmt.Exec(env); stmt.Exec != nil {
																if stmt, env = stmt.Exec(env); stmt.Exec != nil {
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
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)
			stmt, env = stmt.Exec(env)

			if x := stmt; *(**uintptr)(unsafe.Pointer(&x)) == unsafeInterrupt {
				continue outer
			}
		}
	}
}
