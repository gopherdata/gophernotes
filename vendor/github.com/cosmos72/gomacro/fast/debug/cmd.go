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
 * global.go
 *
 *  Created on Apr 21, 2018
 *      Author Massimiliano Ghilardi
 */

package debug

import (
	"strings"

	"github.com/cosmos72/gomacro/base"
	bstrings "github.com/cosmos72/gomacro/base/strings"
)

type Cmd struct {
	Name string
	Func func(d *Debugger, arg string) DebugOp
}

type Cmds map[byte]Cmd

func (cmd *Cmd) Match(prefix string) bool {
	return strings.HasPrefix(cmd.Name, prefix)
}

func (cmds Cmds) Lookup(prefix string) (Cmd, bool) {
	if len(prefix) != 0 {
		cmd, found := cmds[prefix[0]]
		if found && cmd.Match(prefix) {
			return cmd, true
		}
	}
	return Cmd{}, false
}

var cmds = Cmds{
	'b': Cmd{"backtrace", (*Debugger).cmdBacktrace},
	'c': Cmd{"continue", (*Debugger).cmdContinue},
	'e': Cmd{"env", (*Debugger).cmdEnv},
	'f': Cmd{"finish", (*Debugger).cmdFinish},
	'h': Cmd{"help", (*Debugger).cmdHelp},
	'?': Cmd{"?", (*Debugger).cmdHelp},
	'i': Cmd{"inspect", (*Debugger).cmdInspect},
	'k': Cmd{"kill", (*Debugger).cmdKill},
	'l': Cmd{"list", (*Debugger).cmdList},
	'n': Cmd{"next", (*Debugger).cmdNext},
	'p': Cmd{"print", (*Debugger).cmdPrint},
	's': Cmd{"step", (*Debugger).cmdStep},
	'v': Cmd{"vars", (*Debugger).cmdVars},
}

// execute one of the debugger commands
func (d *Debugger) Cmd(src string) DebugOp {
	op := DebugOpRepl
	src = strings.TrimSpace(src)
	n := len(src)
	if n > 0 {
		prefix, arg := bstrings.Split2(src, ' ')
		cmd, found := cmds.Lookup(prefix)
		if found {
			d.lastcmd = src
			op = cmd.Func(d, arg)
		} else {
			g := d.globals
			g.Fprintf(g.Stdout, "// unknown debugger command, type ? for help: %s\n", src)
		}
	}
	return op
}

func (d *Debugger) cmdBacktrace(arg string) DebugOp {
	d.Backtrace(arg)
	return DebugOpRepl
}

func (d *Debugger) cmdContinue(arg string) DebugOp {
	return DebugOpContinue
}

func (d *Debugger) cmdEnv(arg string) DebugOp {
	d.interp.ShowPackage(arg)
	return DebugOpRepl
}

func (d *Debugger) cmdFinish(arg string) DebugOp {
	return DebugOp{d.env.CallDepth, nil}
}

func (d *Debugger) cmdHelp(arg string) DebugOp {
	d.Help()
	return DebugOpRepl
}

func (d *Debugger) cmdInspect(arg string) DebugOp {
	if len(arg) == 0 {
		g := d.globals
		g.Fprintf(g.Stdout, "// inspect: missing argument\n")
	} else {
		d.interp.Inspect(arg)
	}
	return DebugOpRepl
}

func (d *Debugger) cmdKill(arg string) DebugOp {
	var panick interface{} = base.SigInterrupt
	if len(arg) != 0 {
		vals, _ := d.Eval(arg)
		if len(vals) != 0 {
			if !vals[0].IsValid() {
				panick = nil
			} else {
				val := vals[0]
				if val.CanInterface() {
					panick = val.Interface()
				} else {
					panick = val
				}
			}
		}
	}
	return DebugOp{0, &panick}
}

func (d *Debugger) cmdList(arg string) DebugOp {
	d.Show(false)
	return DebugOpRepl
}

func (d *Debugger) cmdNext(arg string) DebugOp {
	return DebugOp{d.env.CallDepth + 1, nil}
}

func (d *Debugger) cmdPrint(arg string) DebugOp {
	g := d.globals
	if len(arg) == 0 {
		g.Fprintf(g.Stdout, "// print: missing argument\n")
	} else {
		vals, types := d.Eval(arg)
		g.Print(vals, types)
	}
	return DebugOpRepl
}

func (d *Debugger) cmdStep(arg string) DebugOp {
	return DebugOpStep
}

func (d *Debugger) cmdVars(arg string) DebugOp {
	d.Vars()
	return DebugOpRepl
}
