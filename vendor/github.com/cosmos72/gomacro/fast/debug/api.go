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
 * api.go
 *
 *  Created on Apr 21, 2018
 *      Author Massimiliano Ghilardi
 */

package debug

import (
	"github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/fast"
)

type DebugOp = fast.DebugOp

var (
	DebugOpContinue = fast.DebugOpContinue
	DebugOpStep     = fast.DebugOpStep
	DebugOpRepl     = DebugOp{-1, nil}
)

type Debugger struct {
	interp  *fast.Interp
	env     *fast.Env
	globals *base.Globals
	lastcmd string
}

func (d *Debugger) Breakpoint(interp *fast.Interp, env *fast.Env) DebugOp {
	return d.main(interp, env, true)
}

func (d *Debugger) At(interp *fast.Interp, env *fast.Env) DebugOp {
	return d.main(interp, env, false)
}

func (d *Debugger) main(interp *fast.Interp, env *fast.Env, breakpoint bool) DebugOp {
	// create an inner Interp to preserve existing Binds, compiled Code and IP
	//
	// this is needed to allow compiling and evaluating code at a breakpoint or single step
	// without disturbing the code being debugged
	d.interp = fast.NewInnerInterp(interp, "debug", "debug")
	d.env = env
	d.globals = &interp.Comp.Globals
	if !d.Show(breakpoint) {
		// skip synthetic statements
		return DebugOp{Depth: env.Run.DebugDepth}
	}
	return d.Repl()
}
