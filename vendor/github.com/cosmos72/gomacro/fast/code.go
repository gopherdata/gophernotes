/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * exec.go
 *
 *  Created on Apr 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/token"

	. "github.com/cosmos72/gomacro/base"
)

func (code *Code) Clear() {
	code.List = nil
	code.DebugPos = nil
	code.WithDefers = false
}

func (code *Code) Len() int {
	return len(code.List)
}

func (code *Code) Truncate(n int) {
	if len(code.List) > n {
		code.List = code.List[0:n]
	}
	if len(code.DebugPos) > n {
		code.DebugPos = code.DebugPos[0:n]
	}
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

// spinInterrupt is the statement executed while waiting for an interrupt to be serviced.
// To signal an interrupt, a statement must set env.ThreadGlobals.Signal to the desired signal,
// then return env.ThreadGlobals.Interrupt, env
func spinInterrupt(env *Env) (Stmt, *Env) {
	run := env.Run
	if run.Signals.IsEmpty() {
		run.Signals.Sync = SigReturn
	} else if sig := run.Signals.Async; sig != SigNone {
		run.applyAsyncSignal(sig)
	}
	return run.Interrupt, env
}

func (run *Run) applyAsyncSignal(sig Signal) {
	run.Signals.Async = SigNone
	switch sig {
	case SigNone:
		break
	case SigDebug:
		run.applyDebugOp(DebugOpStep)
	default:
		panic(SigInterrupt)
	}
}

func pushDefer(g *Run, deferOf *Env, panicking bool) (retg *Run, deferOf_ *Env, isDefer bool) {
	deferOf_ = g.DeferOfFun
	if panicking {
		g.PanicFun = deferOf
	}
	g.DeferOfFun = deferOf
	g.ExecFlags.SetStartDefer(true)
	return g, deferOf_, g.ExecFlags.IsDefer()
}

func popDefer(run *Run, deferOf *Env, isDefer bool) {
	run.DeferOfFun = deferOf
	run.ExecFlags.SetStartDefer(false)
	run.ExecFlags.SetDefer(isDefer)
}

func restore(run *Run, isDefer bool, interrupt Stmt, caller *Env) {
	run.ExecFlags.SetDefer(isDefer)
	run.Interrupt = interrupt
	run.CurrEnv = caller
	run.Signals.Sync = SigNone
	if sig := run.Signals.Async; sig == SigInterrupt {
		// do NOT handle async SigDebug here
		run.applyAsyncSignal(sig)
	}
}

func maybeRepanic(run *Run) bool {
	if run.PanicFun != nil {
		panic(run.Panic)
	}
	// either not panicking or recover() invoked, no longer panicking
	return false
}

func (run *Run) interrupt() {
	const CtrlCDebug = OptDebugger | OptCtrlCEnterDebugger
	var sig Signal
	if run.Options&CtrlCDebug == CtrlCDebug {
		sig = SigDebug
	} else {
		sig = SigInterrupt
	}
	run.Signals.Async = sig
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
	all = append(all, spinInterrupt)

	if defers {
		// code to support defer is slower... isolate it in a separate function
		return execWithFlags(all, pos)
	}
	return exec(all, pos)
}

func exec(all []Stmt, pos []token.Pos) func(*Env) {
	return func(env *Env) {
		run := env.Run
		run.Signals.Sync = SigNone
		if run.ExecFlags != 0 {
			// code to support defer and debugger is slower... isolate it in a separate function
			reExecWithFlags(env, all, pos, all[0], 0)
			return
		}
		if sig := run.Signals.Async; sig != SigNone {
			run.applyAsyncSignal(sig)
		}
		saveInterrupt := run.Interrupt
		run.Interrupt = nil

		stmt := all[0]
		env.IP = 0
		env.Code = all
		env.DebugPos = pos

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
																	if run.Signals.IsEmpty() {
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
			}
			goto finish
		}

		run.Interrupt = spinInterrupt
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

			if !run.Signals.IsEmpty() {
				break
			}
		}
	finish:
		// restore env.ThreadGlobals.Interrupt and Signal before returning
		run.Interrupt = saveInterrupt
		if sig := run.Signals.Async; sig != SigNone {
			run.applyAsyncSignal(sig) // may set run.Signals.Debug if OptCtrlCEnterDebugger is set
		}
		if run.Signals.Debug == SigNone {
			run.Signals.Sync = SigNone
		} else {
			reExecWithFlags(env, all, pos, stmt, env.IP)
		}
	}
}

// execWithFlags returns a function that will execute the given compiled code, including support for defer() and debugger
func execWithFlags(all []Stmt, pos []token.Pos) func(*Env) {
	return func(env *Env) {
		env.Run.Signals.Sync = SigNone
		reExecWithFlags(env, all, pos, all[0], 0)
	}
}

func reExecWithFlags(env *Env, all []Stmt, pos []token.Pos, stmt Stmt, ip int) {
	run := env.Run

	ef := &run.ExecFlags
	trace := run.Options&OptDebugDebugger != 0
	if trace {
		run.Debugf("reExecWithFlags:  executing function   stmt = %p, env = %p, IP = %v, execFlags = %v, signals = %#v", stmt, env, ip, *ef, run.Signals)
	}
	if sig := run.Signals.Async; sig != SigNone {
		run.applyAsyncSignal(sig)
	}
	caller := run.CurrEnv
	// restore g.IsDefer, g.Signal, g.DebugCallDepth, g.Interrupt and g.Caller on return
	defer restore(run, run.ExecFlags.IsDefer(), run.Interrupt, caller)
	ef.SetDefer(ef.StartDefer())
	ef.SetStartDefer(false)
	ef.SetDebug(run.Signals.Debug != SigNone)

	funenv := env
	env.IP = ip
	env.Code = all
	env.DebugPos = pos

	panicking, panicking2 := true, false
	rundefer := func(fun func()) {
		if panicking || panicking2 {
			panicking = true
			panicking2 = false
			run.Panic = recover()
		}
		defer popDefer(pushDefer(run, funenv, panicking))
		panicking2 = true // detect panics inside defer
		fun()
		panicking2 = false
		if panicking {
			panicking = maybeRepanic(run)
		}
	}

	if stmt == nil || !run.Signals.IsEmpty() {
		goto signal
	}
again:
	run.Interrupt = nil
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
																if run.Signals.IsEmpty() {
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
		}
		for run.Signals.Sync == SigDefer {
			run.Signals.Sync = SigNone
			fun := run.InstallDefer
			run.InstallDefer = nil
			defer rundefer(fun)
			stmt = env.Code[env.IP]
			if stmt == nil {
				goto signal
			}
		}
		if !run.Signals.IsEmpty() {
			goto signal
		}
		continue
	}

	run.Interrupt = spinInterrupt
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

		for run.Signals.Sync == SigDefer {
			run.Signals.Sync = SigNone
			fun := run.InstallDefer
			run.InstallDefer = nil
			defer rundefer(fun)
			// single step
			stmt = env.Code[env.IP]
			stmt, env = stmt(env)
		}
		if !run.Signals.IsEmpty() {
			goto signal
		}
	}
signal:
	if sig := run.Signals.Async; sig != SigNone {
		// if OptCtrlCEnterDebugger is set, convert early
		// Signals.Async = SigDebug to Signals.Debug = SigDebug
		run.applyAsyncSignal(sig)
	}

	for run.Signals.Debug != SigNone {
		run.Interrupt = spinInterrupt
		stmt, env = singleStep(env)
		if trace {
			run.Debugf("singleStep returned stmt = %p, env = %p, IP = %v, execFlags = %v, signals = %#v", stmt, env, env.IP, run.ExecFlags, run.Signals)
		}
		// a Sync or Async signal may be pending.
		sig := run.Signals.Sync
		if run.Signals.IsEmpty() || sig == SigDefer {
			goto again
		} else if sig == SigReturn {
			break
		} else if sig = run.Signals.Async; sig != SigNone {
			run.applyAsyncSignal(sig)
		}
	}
	panicking = false
	// no need to restore g.IsDefer, g.Signal, g.Interrupt:
	// done by defer restore(g, g.IsDefer, interrupt) above
	return
}
