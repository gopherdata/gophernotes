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
 * interp.go
 *
 *  Created on: Jun 15, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"bufio"
	"fmt"
	"os"
	r "reflect"
	"runtime/debug"
	"strings"
	"time"

	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/paths"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type Interp struct {
	*Env
}

func New() *Interp {
	top := NewEnv(nil, "builtin")
	env := NewEnv(top, "main")
	return &Interp{Env: env}
}

func (ir *Interp) ChangePackage(path string) {
	ir.Env = ir.Env.ChangePackage(path)
}

var historyfile = paths.Subdir(paths.UserHomeDir(), ".gomacro_history")

func (ir *Interp) ReplStdin() {
	g := ir.Globals
	if g.Options&OptShowPrompt != 0 {
		fmt.Fprintf(ir.Stdout, `// GOMACRO, an interactive Go interpreter with macros <https://github.com/cosmos72/gomacro>
// Copyright (C) 2017-2019 Massimiliano Ghilardi
// License MPL v2.0+: Mozilla Public License version 2.0 or later <http://mozilla.org/MPL/2.0/>
// This is free software with ABSOLUTELY NO WARRANTY.
//
// Type %chelp for help
`, g.ReplCmdChar)
	}
	tty, _ := MakeTtyReadline(historyfile)
	defer tty.Close(historyfile) // restore normal tty mode

	c := StartSignalHandler(ir.Interrupt)
	defer StopSignalHandler(c)

	savetty := g.Readline
	g.Readline = tty
	defer func() {
		g.Readline = savetty
	}()

	ir.Line = 0
	for ir.ReadParseEvalPrint() {
		ir.Line = 0
	}
	os.Stdout.WriteString("\n")
}

func (ir *Interp) Repl(in *bufio.Reader) {
	r := MakeBufReadline(in, ir.Stdout)

	c := StartSignalHandler(ir.Interrupt)
	defer StopSignalHandler(c)

	g := ir.Globals
	savetty := g.Readline
	g.Readline = r
	defer func() {
		g.Readline = savetty
	}()

	for ir.ReadParseEvalPrint() {
	}
}

func (ir *Interp) ReadParseEvalPrint() (callAgain bool) {
	str, firstToken := ir.Read()
	if firstToken < 0 {
		// skip comment-only lines and continue, but fail on EOF or other errors
		return len(str) != 0
	}
	return ir.ParseEvalPrint(str[firstToken:])
}

// return read string and position of first non-comment token.
// return "", -1 on EOF
func (ir *Interp) Read() (string, int) {
	var opts ReadOptions
	if ir.Options&OptShowPrompt != 0 {
		opts |= ReadOptShowPrompt
	}
	str, firstToken := ir.Env.Globals.ReadMultiline(opts, "gomacro> ")
	if firstToken < 0 {
		ir.IncLine(str)
	} else if firstToken > 0 {
		ir.IncLine(str[0:firstToken])
	}
	return str, firstToken
}

func (ir *Interp) ParseEvalPrint(str string) (callAgain bool) {
	var t1 time.Time
	trap := ir.Options&OptTrapPanic != 0
	duration := ir.Options&OptShowTime != 0
	if duration {
		t1 = time.Now()
	}
	defer func() {
		ir.IncLine(str)
		if trap {
			rec := recover()
			if ir.Options&OptPanicStackTrace != 0 {
				fmt.Fprintf(ir.Stderr, "%v\n%s", rec, debug.Stack())
			} else {
				fmt.Fprintf(ir.Stderr, "%v\n", rec)
			}
			callAgain = true
		}
		if duration {
			delta := time.Since(t1)
			ir.Debugf("eval time %v", delta)
		}
	}()
	callAgain = ir.parseEvalPrint(str)
	trap = false // no panic happened
	return callAgain
}

func (ir *Interp) parseEvalPrint(src string) (callAgain bool) {
	if len(strings.TrimSpace(src)) == 0 {
		return true // no input. don't print anything
	}
	env := ir.Env
	g := env.Globals

	src, opt := ir.Cmd(src)

	callAgain = opt&CmdOptQuit == 0
	if len(src) == 0 || !callAgain {
		return callAgain
	}

	if opt&CmdOptForceEval != 0 {
		// temporarily disable collection of declarations and statements,
		// and temporarily re-enable eval (i.e. disable macroexpandonly)
		const todisable = OptMacroExpandOnly | OptCollectDeclarations | OptCollectStatements
		if g.Options&todisable != 0 {
			g.Options &^= todisable
			defer func() {
				g.Options |= todisable
			}()
		}
	}

	ir.currOpt = opt // store options where Interp.Interrupt() can find them

	// parse phase. no macroexpansion/collect yet
	form := env.ParseOnly(src)

	// macroexpand + collect + eval phase
	var values []r.Value
	var types []xr.Type
	if form != nil {
		values = env.classicEval(form)
	}

	// print phase
	g.Print(values, types)
	return true
}

func (ir *Interp) Interrupt(sig os.Signal) {
	// TODO not implemented
}
