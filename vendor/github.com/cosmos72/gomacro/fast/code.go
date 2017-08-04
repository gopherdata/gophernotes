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
 * code.go
 *
 *  Created on Apr 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/token"
	"unsafe"
)

func (code *Code) Clear() {
	code.List = nil
	code.DebugPos = nil
	code.WithDefers = false
}

func (code *Code) Len() int {
	return len(code.List)
}

func (code *Code) Append(stmt Stmt, pos token.Pos) {
	if stmt != nil {
		code.List = append(code.List, stmt)
		code.DebugPos = append(code.DebugPos, pos)
	}
}

func (code *Code) AsExpr() *Expr {
	fun := code.Exec()
	if fun == nil {
		return nil
	}
	return expr0(fun)
}

// declare a var instead of function: Code.Exec() needs the address of Interrupt
var Interrupt Stmt = func(env *Env) (Stmt, *Env) {
	return env.ThreadGlobals.Interrupt, env
}

func pushDefer(g *ThreadGlobals, deferOf *Env, panicking bool) (retg *ThreadGlobals, deferOf_ *Env, isDefer bool) {
	deferOf_ = g.DeferOfFun
	if panicking {
		g.PanicFun = deferOf
	}
	g.DeferOfFun = deferOf
	g.StartDefer = true
	return g, deferOf_, g.IsDefer
}

func popDefer(g *ThreadGlobals, deferOf *Env, isDefer bool) {
	g.DeferOfFun = deferOf
	g.StartDefer = false
	g.IsDefer = isDefer
}

func restore(g *ThreadGlobals, flag bool, interrupt Stmt) {
	g.IsDefer = flag
	g.Signal = SigNone
	g.Interrupt = interrupt
}

func maybeRepanic(g *ThreadGlobals) bool {
	if g.PanicFun != nil {
		panic(g.Panic)
	}
	// either not panicking or recover() invoked, no longer panicking
	return false
}

// Exec returns a func(*Env) that will execute the compiled code
func (code *Code) Exec() func(*Env) {
	all := code.List
	pos := code.DebugPos
	defers := code.WithDefers

	code.Clear()
	if len(all) == 0 {
		return nil
	}
	all = append(all, Interrupt)

	if defers {
		// code to support defer is slower... isolate it in a separate function
		return func(env *Env) {
			execWithDefers(env, all, pos)
		}
	} else {
		return exec(all, pos)
	}
}

func exec(all []Stmt, pos []token.Pos) func(*Env) {
	return func(env *Env) {
		g := env.ThreadGlobals
		if g.IsDefer || g.StartDefer {
			// code to support defer is slower... isolate it in a separate function
			execWithDefers(env, all, pos)
			return
		}
		stmt := all[0]
		env.IP = 0
		env.Code = all
		env.DebugPos = pos

		interrupt := g.Interrupt
		g.Interrupt = nil
		var unsafeInterrupt *uintptr
		g.Signal = SigNone

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
			goto finish
		}

		unsafeInterrupt = *(**uintptr)(unsafe.Pointer(&Interrupt))
		env.ThreadGlobals.Interrupt = Interrupt
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

			if *(**uintptr)(unsafe.Pointer(&stmt)) == unsafeInterrupt {
				break
			}
		}
	finish:
		// restore env.ThreadGlobals.Interrupt and Signal before returning
		g.Interrupt = interrupt
		g.Signal = SigNone
		return
	}
}

// execWithDefers executes the given compiled code, including support for defer()
func execWithDefers(env *Env, all []Stmt, pos []token.Pos) {
	funenv := env
	stmt := all[0]
	env.IP = 0
	env.Code = all
	env.DebugPos = pos

	g := env.ThreadGlobals
	interrupt := g.Interrupt
	g.Interrupt = nil
	var unsafeInterrupt *uintptr

	defer restore(g, g.IsDefer, interrupt) // restore g.IsDefer, g.Signal and g.Interrupt on return
	g.Signal = SigNone
	g.IsDefer = g.StartDefer
	g.StartDefer = false
	panicking := true
	panicking2 := false

	rundefer := func(fun func()) {
		if panicking || panicking2 {
			panicking = true
			panicking2 = false
			g.Panic = recover()
		}
		defer popDefer(pushDefer(g, funenv, panicking))
		panicking2 = true // detect panics inside defer
		fun()
		panicking2 = false
		if panicking {
			panicking = maybeRepanic(g)
		}
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
		if g.Signal != SigDefer {
			goto finish
		}
		fun := g.InstallDefer
		g.Signal = SigNone
		g.InstallDefer = nil
		defer rundefer(fun)
		stmt = env.Code[env.IP]
		if stmt != nil {
			continue
		}
		break
	}

	unsafeInterrupt = *(**uintptr)(unsafe.Pointer(&Interrupt))
	env.ThreadGlobals.Interrupt = Interrupt
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

		if *(**uintptr)(unsafe.Pointer(&stmt)) == unsafeInterrupt {
			if g.Signal != SigDefer {
				goto finish
			}
			fun := g.InstallDefer
			g.Signal = SigNone
			g.InstallDefer = nil
			defer rundefer(fun)
			stmt = env.Code[env.IP]
			if *(**uintptr)(unsafe.Pointer(&stmt)) != unsafeInterrupt {
				continue
			}
			break
		}
	}
finish:
	panicking = false
	return
}
