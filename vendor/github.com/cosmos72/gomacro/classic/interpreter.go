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
)

type Interp struct {
	*Env
}

func New() *Interp {
	top := NewEnv(nil, "builtin")
	env := NewEnv(top, "main")
	return &Interp{env}
}

func (ir *Interp) ChangePackage(path string) {
	ir.Env = ir.Env.ChangePackage(path)
}

// cmd is expected to start with 'package'
func (ir *Interp) cmdPackage(cmd string) {
	cmd = strings.TrimSpace(cmd)
	space := strings.IndexByte(cmd, ' ')
	var arg string
	if space >= 0 {
		arg = strings.TrimSpace(cmd[1+space:])
	}
	n := len(arg)
	env := ir.Env
	switch {
	case n == 0:
		fmt.Fprintf(env.Stdout, "// current package: %s %q\n", env.Filename, env.PackagePath)
	case arg[0] == '"' && arg[n-1] == '"':
		arg := arg[1 : n-1]
		ir.Env = env.ChangePackage(arg)
	default:
		env.Filename = arg
	}
}

func (ir *Interp) ReplStdin() {
	if ir.Options&OptShowPrompt != 0 {
		fmt.Fprint(ir.Stdout, `// GOMACRO, an interactive Go interpreter with macros <https://github.com/cosmos72/gomacro>
// Copyright (C) 2017 Massimiliano Ghilardi
// License LGPL v3+: GNU Lesser GPL version 3 or later <https://gnu.org/licenses/lgpl>
// This is free software with ABSOLUTELY NO WARRANTY.
//
// Type :help for help
`)
	}
	in := bufio.NewReader(os.Stdin)

	ir.Line = 0
	for ir.ReadParseEvalPrint(in) {
		ir.Line = 0
	}
}

func (ir *Interp) Repl(in *bufio.Reader) {
	for ir.ReadParseEvalPrint(in) {
	}
}

func (ir *Interp) ReadParseEvalPrint(in *bufio.Reader) (callAgain bool) {
	var opts ReadOptions
	if ir.Options&OptShowPrompt != 0 {
		opts |= ReadOptShowPrompt
	}
	str, firstToken := ir.ReadMultiline(in, opts)
	if firstToken < 0 {
		// skip comments and continue, but fail on EOF or other errors
		ir.IncLine(str)
		return len(str) > 0
	} else if firstToken > 0 {
		ir.IncLine(str[0:firstToken])
	}
	return ir.ParseEvalPrint(str[firstToken:], in)
}

func (ir *Interp) ParseEvalPrint(str string, in *bufio.Reader) (callAgain bool) {
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
			delta := time.Now().Sub(t1)
			ir.Debugf("eval time %.6f s", float32(delta)/float32(time.Second))
		}
	}()
	callAgain = ir.parseEvalPrint(str, in)
	trap = false // no panic happened
	return callAgain
}

func (ir *Interp) parseEvalPrint(src string, in *bufio.Reader) (callAgain bool) {

	src = strings.TrimSpace(src)
	n := len(src)
	if n == 0 {
		return true // no input. don't print anything
	}
	env := ir.Env
	fast := env.Options&OptFastInterpreter != 0 // use the fast interpreter?

	if n > 0 && src[0] == ':' {
		args := strings.SplitN(src, " ", 2)
		cmd := args[0]
		switch {
		case strings.HasPrefix(":classic", cmd):
			if len(args) <= 1 {
				if env.Options&OptFastInterpreter != 0 {
					env.Debugf("switched to classic interpreter")
				}
				env.Options &^= OptFastInterpreter
				return true
			}
			// temporary override
			src = strings.TrimSpace(args[1])
			fast = false
		case strings.HasPrefix(":env", cmd):
			if len(args) <= 1 {
				env.ShowPackage("")
			} else {
				env.ShowPackage(args[1])
			}
			return true
		case strings.HasPrefix(":fast", cmd):
			if len(args) <= 1 {
				if env.Options&OptFastInterpreter == 0 {
					env.Debugf("switched to fast interpreter")
				}
				env.Options |= OptFastInterpreter
				return true
			}
			// temporary override
			src = strings.TrimSpace(args[1])
			fast = true
		case strings.HasPrefix(":help", cmd):
			env.ShowHelp(env.Stdout)
			return true
		case strings.HasPrefix(":inspect", cmd):
			if in == nil {
				fmt.Fprint(env.Stdout, "// not connected to user input, cannot :inspect\n")
			} else if len(args) == 1 {
				fmt.Fprint(env.Stdout, "// inspect: missing argument\n")
			} else {
				env.Inspect(in, args[1], fast)
			}
			return true
		case strings.HasPrefix(":options", cmd):
			if len(args) > 1 {
				env.Options ^= ParseOptions(args[1])
			}
			fmt.Fprintf(env.Stdout, "// current options: %v\n", env.Options)
			fmt.Fprintf(env.Stdout, "// unset   options: %v\n", ^env.Options)
			return true
		case strings.HasPrefix(":quit", cmd):
			return false
		case strings.HasPrefix(":write", cmd):
			if len(args) <= 1 {
				env.WriteDeclsToStream(env.Stdout)
			} else {
				env.WriteDeclsToFile(args[1])
			}
			return true
		default:
			// temporarily disable collection of declarations and statements,
			// and temporarily disable macroexpandonly (i.e. re-enable eval)
			saved := env.Options
			todisable := OptMacroExpandOnly | OptCollectDeclarations | OptCollectStatements
			if saved&todisable != 0 {
				env.Options &^= todisable
				defer func() {
					env.Options = saved
				}()
			}
			src = " " + src[1:] // slower than src = src[1:], but gives accurate column positions in error messages
		}
	}
	if !fast {
		if src == "package" || src == " package" || strings.HasPrefix(src, "package ") || strings.HasPrefix(src, " package ") {
			ir.cmdPackage(src)
			return true
		}
	}

	// parse phase. no macroexpansion/collect yet
	form := env.ParseOnly(src)

	var value r.Value
	var values []r.Value
	var typ interface{}
	var types []interface{}

	// eval phase
	if form != nil {
		if fast {
			// macroexpand + collect + eval
			xvalue, xvalues, xtype, xtypes := env.fastEval(form)
			value, values, typ = xvalue, xvalues, xtype
			types := make([]interface{}, len(xtypes))
			for i, xt := range xtypes {
				types[i] = xt
			}
		} else {
			// macroexpand + collect + eval
			value, values = env.classicEval(form)
		}
	}
	// print phase
	opts := env.Options
	if opts&OptShowEval != 0 {
		if len(values) != 0 {
			if opts&OptShowEvalType != 0 {
				for i, vi := range values {
					var ti interface{}
					if types != nil && len(types) > i {
						ti = types[i]
					} else {
						ti = ValueType(vi)
					}
					env.Fprintf(env.Stdout, "%v\t// %v\n", vi, ti)
				}
			} else {
				for _, vi := range values {
					env.Fprintf(env.Stdout, "%v\n", vi)
				}
			}
		} else if value != None {
			if opts&OptShowEvalType != 0 {
				if typ == nil {
					typ = ValueType(value)
				}
				env.Fprintf(env.Stdout, "%v\t// %v\n", value, typ)
			} else {
				env.Fprintf(env.Stdout, "%v\n", value)
			}
		}
	}
	return true
}

func (ir *Interp) EvalFile(filePath string, pkgPath string) {
	file, err := os.Open(filePath)
	if err != nil {
		ir.Errorf("error opening file '%s': %v", filePath, err)
		return
	}
	defer file.Close()

	savePath := ir.Env.ThreadGlobals.PackagePath
	saveOpts := ir.Env.Options

	ir.ChangePackage(pkgPath)
	ir.Env.Options &^= OptShowEval

	defer func() {
		ir.ChangePackage(savePath)
		ir.Env.Options = saveOpts
	}()

	in := bufio.NewReader(file)
	ir.Repl(in)
}
