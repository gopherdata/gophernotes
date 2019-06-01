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
 * debugger.go
 *
 *  Created on Apr 21, 2018
 *      Author Massimiliano Ghilardi
 */

package debug

import (
	"go/token"
	"reflect"
	"runtime/debug"

	"github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/xreflect"
)

func (d *Debugger) Help() {
	g := d.globals
	g.Fprintf(g.Stdout, "%s", `// debugger commands:
backtrace       show call stack
env [NAME]      show available functions, variables and constants
                in current scope, or from imported package NAME
?               show this help
help            show this help
inspect EXPR    inspect expression interactively
kill   [EXPR]   terminate execution with panic(EXPR)
print   EXPR    print expression, statement or declaration
list            show current source code
continue        resume normal execution
finish          run until the end of current function
next            execute a single statement, skipping functions
step            execute a single statement, entering functions
vars            show local variables
// abbreviations are allowed if unambiguous. enter repeats last command.
`)
	/*
		not implemented yet:

		backtrace [N] show function stack frames
	*/
}

func (d *Debugger) Show(breakpoint bool) bool {
	// d.env is the Env being debugged.
	// to execute code at debugger prompt, use d.interp
	env := d.env
	pos := env.DebugPos
	g := d.globals
	ip := env.IP

	var label string
	if breakpoint {
		label = "breakpoint"
	} else {
		label = "stopped"
	}
	if ip < len(pos) && g.Fileset != nil {
		p := pos[ip]
		if p == token.NoPos {
			return false
		}
		source, pos := g.Fileset.Source(p)
		g.Fprintf(g.Stdout, "// %s at %s IP=%d, call depth=%d. type ? for debugger help\n", label, pos, ip, env.CallDepth)
		if len(source) != 0 {
			g.Fprintf(g.Stdout, "%s\n", source)
			d.showCaret(source, pos.Column)
		}
	} else {
		g.Fprintf(g.Stdout, "// %s at IP=%d, call depth=%d. type ? for debugger help\n", label, ip, env.CallDepth)
	}
	return true
}

var spaces = []byte("                                                                      ")

func (d *Debugger) showCaret(source string, col int) {
	col--
	n := len(source)
	if col >= 0 && col < n && n >= 5 {
		out := d.globals.Stdout
		chunk := len(spaces)
		for col >= chunk {
			out.Write(spaces)
			col -= chunk
		}
		out.Write(spaces[:col])
		out.Write([]byte("^^^\n"))
	}
}

func (d *Debugger) Repl() DebugOp {
	g := d.globals
	var opts base.ReadOptions
	if g.Options&base.OptShowPrompt != 0 {
		opts |= base.ReadOptShowPrompt
	}
	op := DebugOpRepl
	for op == DebugOpRepl {
		src, firstToken := g.ReadMultiline(opts, "debug> ")
		empty := len(src) == 0
		if firstToken < 0 && empty {
			// EOF
			op = DebugOpContinue
			break
		}
		if empty || src == "\n" {
			// keyboard enter repeats last command
			src = d.lastcmd
		}
		if g.Options&base.OptDebugDebugger != 0 {
			g.Debugf("Debugger: command is %q", src)
		}
		op = d.Cmd(src)
	}
	return op
}

func (d *Debugger) Eval(src string) ([]reflect.Value, []xreflect.Type) {
	g := d.globals
	trap := g.Options&base.OptTrapPanic != 0

	// do NOT debug expression evaluated at debugger prompt!
	sig := &d.env.Run.Signals
	sigdebug := sig.Debug
	sig.Debug = base.SigNone

	defer func() {
		sig.Debug = sigdebug
		if trap {
			rec := recover()
			if g.Options&base.OptPanicStackTrace != 0 {
				g.Fprintf(g.Stderr, "%v\n%s", rec, debug.Stack())
			} else {
				g.Fprintf(g.Stderr, "%v\n", rec)
			}
		}
	}()
	vals, types := d.interp.Eval(src)

	trap = false // no panic happened
	return vals, types
}
