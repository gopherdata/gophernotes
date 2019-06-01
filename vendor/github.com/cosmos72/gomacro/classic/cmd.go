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
 * cmd.go
 *
 *  Created on: Apr 11, 2018
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"fmt"
	"strings"

	. "github.com/cosmos72/gomacro/base"
	bstrings "github.com/cosmos72/gomacro/base/strings"
)

type Cmd struct {
	Name string
	Func func(ir *Interp, arg string, opt CmdOpt) (string, CmdOpt)
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

func (cmds Cmds) ShowHelp(g *Globals) {
	c := g.ReplCmdChar
	fmt.Fprintf(g.Stdout, `// type Go code to execute it. example: func add(x, y int) int { return x + y }

// interpreter commands:
%cdebug EXPR        debug expression or statement interactively
%cenv [NAME]        show available functions, variables and constants
                   in current package, or from imported package NAME
%chelp              show this help
%cinspect EXPR      inspect expression interactively
%coptions [OPTS]    show or toggle interpreter options
%cpackage "PKGPATH" switch to package PKGPATH, importing it if possible.
%cquit              quit the interpreter
%cunload "PKGPATH"  remove package PKGPATH from the list of known packages.
                   later attempts to import it will trigger a recompile
%cwrite [FILE]      write collected declarations and/or statements to standard output or to FILE
                   use %co Declarations and/or %co Statements to start collecting them
// abbreviations are allowed if unambiguous.
`, c, c, c, c, c, c, c, c, c, c, c)
}

var cmds Cmds

func init() {
	cmds = Cmds{
		'e': Cmd{"env", (*Interp).cmdEnv},
		'h': Cmd{"help", (*Interp).cmdHelp},
		'i': Cmd{"inspect", (*Interp).cmdInspect},
		'o': Cmd{"options", (*Interp).cmdOptions},
		'p': Cmd{"package", (*Interp).cmdPackage},
		'q': Cmd{"quit", (*Interp).cmdQuit},
		'u': Cmd{"unload", (*Interp).cmdUnload},
		'w': Cmd{"write", (*Interp).cmdWrite},
	}
}

// execute one of the REPL commands starting with ':'
// return any remainder string to be evaluated, and the options to evaluate it
func (ir *Interp) Cmd(src string) (string, CmdOpt) {
	g := ir.Env.Globals
	var opt CmdOpt

	src = strings.TrimSpace(src)
	n := len(src)
	if n > 0 && src[0] == g.ReplCmdChar {
		prefix, arg := bstrings.Split2(src[1:], ' ') // skip g.ReplCmdChar
		cmd, found := cmds.Lookup(prefix)
		if found {
			src, opt = cmd.Func(ir, arg, opt)
		} else {
			// ":<something>"
			// temporarily disable collection of declarations and statements,
			// and temporarily disable macroexpandonly (i.e. re-enable eval)
			opt |= CmdOptForceEval
			src = " " + src[1:] // slower than src = src[1:], but gives accurate column positions in error messages
		}
	}
	// :package and package are the same command
	if g.Options&OptMacroExpandOnly == 0 && (src == "package" || strings.HasPrefix(src, "package ")) {
		_, arg := bstrings.Split2(src, ' ')
		src, opt = ir.cmdPackage(arg, opt)
	}
	return src, opt
}

func (ir *Interp) cmdEnv(arg string, opt CmdOpt) (string, CmdOpt) {
	ir.Env.ShowPackage(arg)
	return "", opt
}

func (ir *Interp) cmdHelp(arg string, opt CmdOpt) (string, CmdOpt) {
	g := ir.Env.ThreadGlobals.Globals
	cmds.ShowHelp(g)
	return "", opt
}

func (ir *Interp) cmdInspect(arg string, opt CmdOpt) (string, CmdOpt) {
	env := ir.Env
	if len(arg) == 0 {
		fmt.Fprint(env.Stdout, "// inspect: missing argument\n")
	} else {
		env.Inspect(arg)
	}
	return "", opt
}

func (ir *Interp) cmdOptions(arg string, opt CmdOpt) (string, CmdOpt) {
	env := ir.Env
	g := env.Globals

	if len(arg) != 0 {
		g.Options ^= ParseOptions(arg)
	} else {
		fmt.Fprintf(env.Stdout, "// current options: %v\n", g.Options)
		fmt.Fprintf(env.Stdout, "// unset   options: %v\n", ^g.Options)
	}
	return "", opt
}

// change package. path can be empty or a package path WITH quotes
// 'package NAME' where NAME is without quotes has no effect.
func (ir *Interp) cmdPackage(path string, cmdopt CmdOpt) (string, CmdOpt) {
	env := ir.Env
	g := env.Globals
	path = strings.TrimSpace(path)
	n := len(path)
	if n == 0 {
		g.Fprintf(g.Stdout, "// current package: %s %q\n", env.Name, env.Path)
	} else if n > 2 && path[0] == '"' && path[n-1] == '"' {
		path = path[1 : n-1]
		n -= 2
		ir.ChangePackage(path)
	} else if g.Options&OptShowPrompt != 0 {
		g.Debugf(`package %s has no effect. To switch to a different package, use package "PACKAGE/FULL/PATH" - note the quotes`, path)
	}
	return "", cmdopt
}

func (ir *Interp) cmdQuit(_ string, opt CmdOpt) (string, CmdOpt) {
	return "", opt | CmdOptQuit
}

// remove package 'path' from the list of known packages
func (ir *Interp) cmdUnload(path string, opt CmdOpt) (string, CmdOpt) {
	if n := len(path); n >= 2 && path[0] == '"' && path[n-1] == '"' {
		path = path[1 : n-1]
	}
	if len(path) != 0 {
		ir.Env.Globals.UnloadPackage(path)
	}
	return "", opt
}

func (ir *Interp) cmdWrite(filepath string, opt CmdOpt) (string, CmdOpt) {
	env := ir.Env
	if len(filepath) == 0 {
		env.WriteDeclsToStream(env.Stdout)
	} else {
		env.WriteDeclsToFile(filepath)
	}
	return "", opt
}
