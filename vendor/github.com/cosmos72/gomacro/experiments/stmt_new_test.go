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

const (
	n int = 1000
)

/*
	benchmark results on Intel Core2 Duo P8400 @2.26GHz, Ubuntu 16.04.1 amd64, Linux 4.4.0-31-generic amd64, Go 1.8.1 linux/amd64

    -------- n = 10 --------
	BenchmarkThreadedStmtFunc6-2              	20000000	        64.6 ns/op
	BenchmarkThreadedStmtFunc6Unroll-2        	20000000	        59.8 ns/op
	BenchmarkThreadedStmtFunc6Terminate-2     	20000000	        98.5 ns/op
	BenchmarkThreadedStmtFunc6Adaptive-2      	20000000	        60.7 ns/op
	BenchmarkThreadedStmtStruct6-2            	20000000	        63.5 ns/op
	BenchmarkThreadedStmtStruct6Unroll-2      	30000000	        57.5 ns/op
	BenchmarkThreadedStmtStruct6Terminate-2   	20000000	        78.1 ns/op
	BenchmarkThreadedStmtStruct6Adaptive-2    	30000000	        52.4 ns/op
	BenchmarkThreadedStmtFunc0-2              	20000000	        76.3 ns/op
	BenchmarkThreadedStmtFunc1-2              	20000000	        80.0 ns/op
	BenchmarkThreadedStmtFunc2-2              	20000000	        70.4 ns/op
	BenchmarkThreadedStmtFunc3-2              	20000000	        70.4 ns/op
	BenchmarkThreadedStmtFunc4-2              	20000000	        66.3 ns/op
	BenchmarkThreadedStmtFunc4Unroll-2        	20000000	        59.9 ns/op
	BenchmarkThreadedStmtFunc4Terminate-2     	20000000	        83.1 ns/op
	BenchmarkThreadedStmtFunc4Adaptive-2      	20000000	        63.9 ns/op
	BenchmarkThreadedStmtFunc4Panic-2         	 5000000	       332 ns/op
	BenchmarkThreadedStmtFunc5-2              	20000000	        72.7 ns/op

    -------- n = 100 --------
	BenchmarkThreadedStmtFunc6-2              	 2000000	       665 ns/op
	BenchmarkThreadedStmtFunc6Unroll-2        	 2000000	       600 ns/op
	BenchmarkThreadedStmtFunc6Terminate-2     	 2000000	       634 ns/op
	BenchmarkThreadedStmtFunc6Adaptive-2      	 2000000	       631 ns/op
	BenchmarkThreadedStmtStruct6-2            	 2000000	       636 ns/op
	BenchmarkThreadedStmtStruct6Unroll-2      	 3000000	       581 ns/op
	BenchmarkThreadedStmtStruct6Terminate-2   	 2000000	       597 ns/op
	BenchmarkThreadedStmtStruct6Adaptive-2    	 3000000	       543 ns/op
	BenchmarkThreadedStmtFunc0-2              	 2000000	       777 ns/op
	BenchmarkThreadedStmtFunc1-2              	 2000000	       818 ns/op
	BenchmarkThreadedStmtFunc2-2              	 2000000	       701 ns/op
	BenchmarkThreadedStmtFunc3-2              	 2000000	       701 ns/op
	BenchmarkThreadedStmtFunc4-2              	 2000000	       654 ns/op
	BenchmarkThreadedStmtFunc4Unroll-2        	 3000000	       642 ns/op
	BenchmarkThreadedStmtFunc4Terminate-2     	 2000000	       643 ns/op
	BenchmarkThreadedStmtFunc4Adaptive-2      	 2000000	       606 ns/op
	BenchmarkThreadedStmtFunc4Panic-2         	 2000000	       902 ns/op
	BenchmarkThreadedStmtFunc5-2              	 2000000	       749 ns/op

    -------- n = 1000 --------
	BenchmarkThreadedStmtFunc6-2              	  200000	      6228 ns/op
	BenchmarkThreadedStmtFunc6Unroll-2        	  300000	      5896 ns/op
	BenchmarkThreadedStmtFunc6Terminate-2     	  200000	      5719 ns/op
	BenchmarkThreadedStmtFunc6Adaptive-2      	  300000	      5538 ns/op
	BenchmarkThreadedStmtStruct6-2            	  200000	      6227 ns/op
	BenchmarkThreadedStmtStruct6Unroll-2      	  200000	      5668 ns/op
	BenchmarkThreadedStmtStruct6Terminate-2   	  300000	      6068 ns/op
	BenchmarkThreadedStmtStruct6Adaptive-2    	  300000	      5584 ns/op
	BenchmarkThreadedStmtFunc0-2              	  200000	      7535 ns/op
	BenchmarkThreadedStmtFunc1-2              	  200000	      8109 ns/op
	BenchmarkThreadedStmtFunc2-2              	  200000	      7023 ns/op
	BenchmarkThreadedStmtFunc3-2              	  200000	      7078 ns/op
	BenchmarkThreadedStmtFunc4-2              	  200000	      6480 ns/op
	BenchmarkThreadedStmtFunc4Unroll-2        	  200000	      5987 ns/op
	BenchmarkThreadedStmtFunc4Terminate-2     	  200000	      5892 ns/op
	BenchmarkThreadedStmtFunc4Adaptive-2      	  200000	      5906 ns/op
	BenchmarkThreadedStmtFunc4Panic-2         	  200000	      6004 ns/op
	BenchmarkThreadedStmtFunc5-2              	  200000	      7016 ns/op
*/
type (
	Env6 struct {
		Binds     []r.Value
		Outer     *Env6
		IP        int
		Code      []Stmt6
		Interrupt Stmt6
	}
	Stmt6 func(*Env6) (Stmt6, *Env6)
	X6    func(*Env6)
)

func BenchmarkThreadedFuncX6(b *testing.B) {

	var nop X6 = func(env *Env6) {
	}
	env := &Env6{
		Binds: make([]r.Value, 10),
	}
	all := make([]X6, n)
	for i := 0; i < n; i++ {
		all[i] = nop
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, x := range all {
			x(env)
		}
	}
}

func BenchmarkThreadedStmtFuncX6(b *testing.B) {

	var xnop X6 = func(env *Env6) {
	}
	var nop Stmt6 = func(env *Env6) (Stmt6, *Env6) {
		xnop(env)
		env.IP++
		return env.Code[env.IP], env
	}
	env := &Env6{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt6, n+1)
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
			stmt, env = stmt(env)
		}
	}
}

func BenchmarkThreadedStmtFunc6(b *testing.B) {

	var nop Stmt6 = func(env *Env6) (Stmt6, *Env6) {
		env.IP++
		return env.Code[env.IP], env
	}

	env := &Env6{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt6, n+1)
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
			stmt, env = stmt(env)
		}
	}
}

func BenchmarkThreadedStmtFunc6Unroll(b *testing.B) {
	var nop Stmt6 = func(env *Env6) (Stmt6, *Env6) {
		env.IP++
		return env.Code[env.IP], env
	}
	env := &Env6{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt6, n+1)
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
																if stmt, env = stmt(env); stmt != nil {
																	stmt, env = stmt(env)
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

func BenchmarkThreadedStmtFunc6Terminate(b *testing.B) {
	var interrupt Stmt6
	interrupt = func(env *Env6) (Stmt6, *Env6) {
		return interrupt, env
	}
	unsafeInterrupt := *(**uintptr)(unsafe.Pointer(&interrupt))

	var nop Stmt6 = func(env *Env6) (Stmt6, *Env6) {
		env.IP++
		return env.Code[env.IP], env
	}
	env := &Env6{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt6, n+1)
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
			stmt, env = stmt(env)
			stmt, env = stmt(env)
		}
	}
}

func BenchmarkThreadedStmtFunc6Adaptive(b *testing.B) {
	var nop Stmt6 = func(env *Env6) (Stmt6, *Env6) {
		env.IP++
		return env.Code[env.IP], env
	}
	var interrupt Stmt6 = func(env *Env6) (Stmt6, *Env6) {
		return env.Interrupt, env
	}
	unsafeInterrupt := *(**uintptr)(unsafe.Pointer(&interrupt))

	env := &Env6{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt6, n+1)
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
			}
			continue outer
		}

		all[n] = interrupt
		env.Interrupt = interrupt
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
			stmt, env = stmt(env)
			stmt, env = stmt(env)

			if x := stmt; *(**uintptr)(unsafe.Pointer(&x)) == unsafeInterrupt {
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

func init() {
	println("sizeof(*uintptr) =", unsafe.Sizeof((*uintptr)(nil)))
	println("sizeof(Stmt6) =", unsafe.Sizeof(func(env *Env6) (Stmt6, *Env6) { return nil, env }))
	println("sizeof(StmtS6) =", unsafe.Sizeof(StmtS6{}))
}

func BenchmarkThreadedStmtStruct6(b *testing.B) {

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

func BenchmarkThreadedStmtStruct6Unroll(b *testing.B) {

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

func BenchmarkThreadedStmtStruct6Terminate(b *testing.B) {

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

func BenchmarkThreadedStmtStruct6Adaptive(b *testing.B) {

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
