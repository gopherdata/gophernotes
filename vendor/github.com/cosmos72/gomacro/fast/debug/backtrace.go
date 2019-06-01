/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * backtrace.go
 *
 *  Created on Apr 27, 2018
 *      Author Massimiliano Ghilardi
 */

package debug

import (
	"github.com/cosmos72/gomacro/fast"
)

func (d *Debugger) Backtrace(arg string) DebugOp {
	env := d.env
	var calls []*fast.Env
	for env != nil {
		if env.Caller != nil {
			// function body
			calls = append(calls, env)
			env = env.Caller
		} else {
			// nested env
			env = env.Outer
		}
	}
	d.showFunctionCalls(calls)
	return DebugOpRepl
}

func (d *Debugger) showFunctionCalls(calls []*fast.Env) {
	// show outermost stack frame first
	for i := len(calls) - 1; i >= 0; i-- {
		d.showFunctionCall(calls[i])
	}
}

func (d *Debugger) showFunctionCall(env *fast.Env) {
	g := d.globals
	c := env.DebugComp
	if c == nil || c.FuncMaker == nil {
		g.Fprintf(g.Stdout, "%p\tfunc (???) ???\n", env)
		return
	}
	m := c.FuncMaker

	g.Fprintf(g.Stdout, "%p\tfunc %s(", env, m.Name)
	d.showBinds(env, m.Param)
	g.Fprintf(g.Stdout, ") ")
	if len(m.Result) > 1 {
		g.Fprintf(g.Stdout, "(")
	}
	d.showBinds(env, m.Result)
	if len(m.Result) > 1 {
		g.Fprintf(g.Stdout, ")\n")
	} else {
		g.Fprintf(g.Stdout, "\n")
	}
}
