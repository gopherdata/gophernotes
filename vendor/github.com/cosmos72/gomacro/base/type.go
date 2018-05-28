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
 * type.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	"sort"
	"strings"

	"github.com/cosmos72/gomacro/imports"
)

type PackageRef struct {
	imports.Package
	Name, Path string
}

type Options uint
type WhichMacroExpand uint

const (
	OptCollectDeclarations Options = 1 << iota
	OptCollectStatements
	OptCtrlCEnterDebugger // Ctrl+C enters the debugger instead of injecting a panic. requires OptDebugger
	OptDebugger           // enable debugger support. "break" and _ = "break" are breakpoints and enter the debugger
	OptKeepUntyped
	OptMacroExpandOnly // do not compile or execute code, only parse and macroexpand it
	OptPanicStackTrace
	OptTrapPanic
	OptDebugCallStack
	OptDebugDebugger // print debug information related to the debugger
	OptDebugField
	OptDebugFromReflect
	OptDebugMacroExpand
	OptDebugMethod
	OptDebugParse
	OptDebugRecover
	OptDebugQuasiquote
	OptDebugSleepOnSwitch // to torture-test "switch" implementation for race conditions
	OptShowCompile
	OptShowEval
	OptShowEvalType
	OptShowMacroExpand
	OptShowParse
	OptShowPrompt
	OptShowTime
)

const (
	CMacroExpand1 WhichMacroExpand = iota
	CMacroExpand
	CMacroExpandCodewalk
)

var optNames = map[Options]string{
	OptCollectDeclarations: "Declarations.Collect",
	OptCollectStatements:   "Statements.Collect",
	OptCtrlCEnterDebugger:  "CtrlC.Debugger.Enter",
	OptDebugger:            "Debugger",
	OptKeepUntyped:         "Untyped.Keep",
	OptMacroExpandOnly:     "MacroExpandOnly",
	OptPanicStackTrace:     "StackTrace.OnPanic",
	OptTrapPanic:           "Trap.Panic",
	OptDebugCallStack:      "?CallStack.Debug",
	OptDebugDebugger:       "?Debugger.Debug",
	OptDebugField:          "?Field.Debug",
	OptDebugFromReflect:    "?FromReflect.Debug",
	OptDebugMacroExpand:    "?MacroExpand.Debug",
	OptDebugMethod:         "?Method.Debug",
	OptDebugParse:          "?Parse.Debug",
	OptDebugRecover:        "?Recover.Debug",
	OptDebugQuasiquote:     "?Quasiquote.Debug",
	OptDebugSleepOnSwitch:  "?SwitchSleep.Debug",
	OptShowCompile:         "Compile.Show",
	OptShowEval:            "Eval.Show",
	OptShowEvalType:        "Type.Eval.Show",
	OptShowMacroExpand:     "MacroExpand.Show",
	OptShowParse:           "Parse.Show",
	OptShowPrompt:          "Prompt.Show",
	OptShowTime:            "Time.Show",
}

var optValues = map[string]Options{}

func init() {
	for k, v := range optNames {
		optValues[v] = k
	}
}

func (o Options) String() string {
	names := make([]string, 0)
	for k, v := range optNames {
		if k&o != 0 {
			names = append(names, v)
		}
	}
	sort.Strings(names)
	return strings.Join(names, " ")
}

func ParseOptions(str string) Options {
	var opts Options
	for _, name := range strings.Split(str, " ") {
		if opt, ok := optValues[name]; ok {
			opts ^= opt
		} else if len(name) != 0 {
			for k, v := range optNames {
				if strings.HasPrefix(v, name) {
					opts ^= k
				}
			}
		}
	}
	return opts
}

func (m WhichMacroExpand) String() string {
	switch m {
	case CMacroExpand1:
		return "MacroExpand1"
	case CMacroExpandCodewalk:
		return "MacroExpandCodewalk"
	default:
		return "MacroExpand"
	}
}
