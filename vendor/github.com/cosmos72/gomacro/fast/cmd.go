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
 * cmd.go
 *
 *  Created on: Apr 20, 2018
 *      Author: Massimiliano Ghilardi
 */

package fast

import (
	"errors"
	"io"
	"sort"
	"strings"

	. "github.com/cosmos72/gomacro/base"
)

// ====================== Cmd ==============================

// Cmd is an interpreter special command.
//
// The following Interp methods look for special commands and execute them:
// Cmd, EvalFile, EvalReader, ParseEvalPrint, ReadParseEvalPrint, Repl, ReplStdin
// note that Interp.Eval() does **not** look for special commands!
//
// Cmd.Name is the command name **without** the initial ':'
//   it must be a valid Go identifier and must not be empty.
//   Using a reserved Go keyword (const, for, func, if, package, return, switch, type, var...)
//   or predefined identifier (bool, int, rune, true, false, nil...)
//   is a bad idea because it interferes with gomacro preprocessor mode.
//   Current limitation: Cmd.Name[0] must be ASCII.
//
// Cmd.Help is the help string that will be displayed by :help
//   please look at current :help output and use the same layout if possible.
//
// Cmd.Func is the command implementation. it receives as arguments:
//   - the current Interp object,
//   - the (possibly multi-line) argument string typed by the user
//     note: it will always have balanced amounts of {} [] () '' "" and ``
//   - the current command options
//
// Cmd.Func can perform any action desired by the implementor,
// including calls to Interp methods, and it must return:
//   - a string to be subsequently evaluated by the interpreter.
//     return the empty string if the command does not need any subsequent evaluation,
//     or if it performed the evaluation by itself.
//   - the updated command options.
//     return the received 'opt' argument unless you need to update it.
//
// If Cmd.Func needs to print something, it's recommended to use
//      g := &interp.Comp.Globals
//      g.Fprintf(g.Stdout, FORMAT, ARGS...)
//   instead of the various fmt.*Print* functions, in order to
//   pretty-print interpreter-generated objects (g.Fprintf)
//   and to honour configured redirections (g.Stdout)
//
// To register a new special command, use Commands.Add()
// To unregister an existing special command, use Commands.Del()
// To list existing special commands, use Commands.List()
type Cmd struct {
	Name string
	Func func(interp *Interp, arg string, opt CmdOpt) (string, CmdOpt)
	Help string
}

// if cmd.Name starts with prefix return 0;
// else if cmd.Name < prefix return -1;
// else return 1
func (cmd *Cmd) Match(prefix string) int {
	name := cmd.Name
	if strings.HasPrefix(name, prefix) {
		return 0
	} else if name < prefix {
		return -1
	} else {
		return 1
	}
}

func (cmd *Cmd) ShowHelp(g *Globals) {
	c := string(g.ReplCmdChar)

	help := strings.Replace(cmd.Help, "%c", c, -1)
	g.Fprintf(g.Stdout, "%s%s\n", c, help)
}

// ===================== Cmds ==============================

type Cmds struct {
	m map[byte][]Cmd
}

// search for a Cmd whose name starts with prefix.
// return (zero value, io.EOF) if no match.
// return (cmd, nil) if exactly one match.
// return (zero value, list of match names) if more than one match
func (cmds Cmds) Lookup(prefix string) (Cmd, error) {
	if len(prefix) != 0 {
		if vec, ok := cmds.m[prefix[0]]; ok {
			i, err := prefixSearch(vec, prefix)
			if err != nil {
				return Cmd{}, err
			}
			return vec[i], nil
		}
	}
	return Cmd{}, io.EOF
}

// prefix search: find all the Cmds whose name start with prefix.
// if there are none, return 0 and io.EOF
// if there is exactly one, return its index and nil.
// if there is more than one, return 0 and an error listing the matching ones
func prefixSearch(vec []Cmd, prefix string) (int, error) {
	lo, _ := binarySearch(vec, prefix)
	n := len(vec)
	for ; lo < n; lo++ {
		cmp := vec[lo].Match(prefix)
		if cmp < 0 {
			continue
		} else if cmp == 0 {
			break
		} else {
			return 0, io.EOF
		}
	}
	if lo == n {
		return 0, io.EOF
	}
	hi := lo + 1
	for ; hi < n; hi++ {
		if vec[hi].Match(prefix) > 0 {
			break
		}
	}
	if lo+1 == hi {
		return lo, nil
	}
	names := make([]string, hi-lo)
	for i := lo; i < hi; i++ {
		names[i-lo] = vec[i].Name
	}
	return 0, errors.New(strings.Join(names, " "))
}

// plain binary search for exact Cmd name
func binarySearch(vec []Cmd, exact string) (int, bool) {
	lo, hi := 0, len(vec)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		name := vec[mid].Name
		if name < exact {
			lo = mid + 1
		} else if name > exact {
			hi = mid - 1
		} else {
			return mid, true
		}
	}
	return lo, false
}

// return the list of currently registered special commands
func (cmds Cmds) List() []Cmd {
	var list []Cmd
	for _, vec := range cmds.m {
		for _, cmd := range vec {
			list = append(list, cmd)
		}
	}
	sortCmdList(list)
	return list
}

// order Cmd list by name
func sortCmdList(vec []Cmd) {
	sort.Slice(vec, func(i, j int) bool {
		return vec[i].Name < vec[j].Name
	})
}

// register a new Cmd.
// if cmd.Name is the empty string, do nothing and return false.
// overwrites any existing Cmd with the same name
func (cmds Cmds) Add(cmd Cmd) bool {
	name := cmd.Name
	if len(name) == 0 {
		return false
	}
	c := name[0]
	vec, _ := cmds.m[c]
	if pos, ok := binarySearch(vec, name); ok {
		vec[pos] = cmd
	} else {
		vec = append(vec, cmd)
		sortCmdList(vec)
		cmds.m[c] = vec
	}
	return true
}

// unregister an existing Cmd by name. return true if existed.
// Use with care!
func (cmds Cmds) Del(name string) bool {
	if len(name) != 0 {
		c := name[0]
		if vec, ok := cmds.m[c]; ok {
			if pos, ok := binarySearch(vec, name); ok {
				vec = removeCmd(vec, pos)
				if len(vec) == 0 {
					delete(cmds.m, c)
				} else {
					cmds.m[c] = vec
				}
				return true
			}
		}
	}
	return false
}

// remove Cmd at index 'pos' from slice.
// return updated slice.
func removeCmd(vec []Cmd, pos int) []Cmd {
	head := vec[:pos]
	n := len(vec)
	if pos == n-1 {
		return head
	}
	tail := vec[pos+1:]
	if pos == 0 {
		return tail
	}
	headn, tailn := pos, len(tail)
	if headn >= tailn {
		copy(vec[headn:], tail)
		vec = vec[:n-1]
	} else {
		copy(vec[1:], head)
		vec = vec[1:]
	}
	return vec
}

func (cmds Cmds) ShowHelp(g *Globals) {
	out := g.Stdout
	g.Fprintf(out, "%s",
		"// type Go code to execute it. example: func add(x, y int) int { return x + y }\n\n// interpreter commands:\n")

	for _, cmd := range cmds.List() {
		cmd.ShowHelp(g)
	}
	g.Fprintf(out, "%s", "// abbreviations are allowed if unambiguous.\n")
}

var Commands Cmds

func init() {
	Commands.m = map[byte][]Cmd{
		'd': []Cmd{{"debug", (*Interp).cmdDebug, `debug EXPR        debug expression or statement interactively`}},
		'e': []Cmd{{"env", (*Interp).cmdEnv, `env [NAME]        show available functions, variables and constants
                   in current package, or from imported package NAME`}},
		'h': []Cmd{{"help", (*Interp).cmdHelp, `help              show this help`}},
		'i': []Cmd{{"inspect", (*Interp).cmdInspect, `inspect EXPR      inspect expression interactively`}},
		'o': []Cmd{{"options", (*Interp).cmdOptions, `options [OPTS]    show or toggle interpreter options`}},
		'p': []Cmd{{"package", (*Interp).cmdPackage, `package "PKGPATH" switch to package PKGPATH, importing it if possible`}},
		'q': []Cmd{{"quit", (*Interp).cmdQuit, `quit              quit the interpreter`}},
		'u': []Cmd{{"unload", (*Interp).cmdUnload, `unload "PKGPATH"  remove package PKGPATH from the list of known packages.
                   later attempts to import it will trigger a recompile`}},
		'w': []Cmd{{"write", (*Interp).cmdWrite, `write [FILE]      write collected declarations and/or statements to standard output or to FILE
                   use %copt Declarations and/or %copt Statements to start collecting them`}},
	}
}

// ==================== Interp =============================

// execute one of the REPL commands starting with ':'
// return any remainder string to be evaluated, and the options to evaluate it
func (ir *Interp) Cmd(src string) (string, CmdOpt) {
	g := &ir.Comp.Globals
	var opt CmdOpt

	trim := strings.TrimSpace(src)
	n := len(trim)
	if n > 0 && trim[0] == g.ReplCmdChar {
		prefix, arg := Split2(trim[1:], ' ') // skip g.ReplCmdChar
		cmd, err := Commands.Lookup(prefix)
		if err == nil {
			src, opt = cmd.Func(ir, arg, opt)
		} else if err == io.EOF {
			// ":<something>"
			// temporarily disable collection of declarations and statements,
			// and temporarily disable macroexpandonly (i.e. re-enable eval)
			opt |= CmdOptForceEval
			src = " " + src[1:] // slower than src = src[1:], but gives accurate column positions in error messages
		} else {
			g.Warnf("ambiguous command %q matches: %s", prefix, err)
			return "", opt
		}
	} else if g.Options&OptMacroExpandOnly == 0 && (trim == "package" || strings.HasPrefix(trim, "package ")) {
		_, arg := Split2(trim, ' ')
		src, opt = ir.cmdPackage(arg, opt)
	}
	return src, opt
}

func (ir *Interp) cmdDebug(arg string, opt CmdOpt) (string, CmdOpt) {
	g := &ir.Comp.Globals
	if len(arg) == 0 {
		g.Fprintf(g.Stdout, "// debug: missing argument\n")
	} else {
		g.Print(ir.Debug(arg))
	}
	return "", opt
}

func (ir *Interp) cmdEnv(arg string, opt CmdOpt) (string, CmdOpt) {
	ir.ShowPackage(arg)
	return "", opt
}

func (ir *Interp) cmdHelp(arg string, opt CmdOpt) (string, CmdOpt) {
	Commands.ShowHelp(&ir.Comp.Globals)
	return "", opt
}

func (ir *Interp) cmdInspect(arg string, opt CmdOpt) (string, CmdOpt) {
	g := &ir.Comp.Globals
	if len(arg) == 0 {
		g.Fprintf(g.Stdout, "// inspect: missing argument\n")
	} else {
		ir.Inspect(arg)
	}
	return "", opt
}

func (ir *Interp) cmdOptions(arg string, opt CmdOpt) (string, CmdOpt) {
	c := ir.Comp
	g := &c.Globals

	if len(arg) != 0 {
		g.Options ^= ParseOptions(arg)

		debugdepth := 0
		if g.Options&OptDebugFromReflect != 0 {
			debugdepth = 1
		}
		c.CompGlobals.Universe.DebugDepth = debugdepth

	} else {
		g.Fprintf(g.Stdout, "// current options: %v\n", g.Options)
		g.Fprintf(g.Stdout, "// unset   options: %v\n", ^g.Options)
	}
	return "", opt
}

// change package. path can be empty or a package path WITH quotes
// 'package NAME' where NAME is without quotes has no effect.
func (ir *Interp) cmdPackage(path string, cmdopt CmdOpt) (string, CmdOpt) {
	c := ir.Comp
	g := &c.Globals
	path = strings.TrimSpace(path)
	n := len(path)
	if len(path) == 0 {
		g.Fprintf(g.Stdout, "// current package: %s %q\n", c.Name, c.Path)
	} else if n > 2 && path[0] == '"' && path[n-1] == '"' {
		path = path[1 : n-1]
		ir.ChangePackage(FileName(path), path)
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
	if len(path) != 0 {
		ir.Comp.UnloadPackage(path)
	}
	return "", opt
}

func (ir *Interp) cmdWrite(filepath string, opt CmdOpt) (string, CmdOpt) {
	g := &ir.Comp.Globals
	if len(filepath) == 0 {
		g.WriteDeclsToStream(g.Stdout)
	} else {
		g.WriteDeclsToFile(filepath)
	}
	return "", opt
}
