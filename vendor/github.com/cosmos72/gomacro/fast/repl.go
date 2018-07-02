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
 * repl.go
 *
 *  Created on: Apr 28, 2018
 *      Author: Massimiliano Ghilardi
 */

package fast

import (
	"bufio"
	"go/ast"
	"go/token"
	"os"
	r "reflect"
	"runtime/debug"
	"sort"
	"strings"
	"time"

	"github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// return read string and position of first non-comment token.
// return "", -1 on EOF
func (ir *Interp) Read() (string, int) {
	g := &ir.Comp.Globals
	var opts ReadOptions

	if g.Options&OptShowPrompt != 0 {
		opts |= ReadOptShowPrompt
	}
	src, firstToken := g.ReadMultiline(opts, ir.Comp.Prompt)
	if firstToken < 0 {
		g.IncLine(src)
	} else if firstToken > 0 {
		g.IncLine(src[0:firstToken])
	}
	return src, firstToken
}

// parse + macroexpansion + collect declarations & statements
func (ir *Interp) Parse(src string) ast2.Ast {
	if len(src) == 0 {
		return nil
	}
	form := ir.Comp.Parse(src)
	if form == nil {
		return nil
	}
	// collect phase
	g := &ir.Comp.Globals
	if g.Options&(OptCollectDeclarations|OptCollectStatements) != 0 {
		g.CollectAst(form)
	}
	return form
}

// combined Parse + Compile
func (ir *Interp) Compile(src string) *Expr {
	return ir.CompileAst(ir.Parse(src))
}

func (ir *Interp) CompileNode(node ast.Node) *Expr {
	return ir.CompileAst(ast2.ToAst(node))
}

func (ir *Interp) CompileAst(form ast2.Ast) *Expr {
	if form == nil {
		return nil
	}
	c := ir.Comp
	g := c.CompGlobals

	if g.Options&OptMacroExpandOnly != 0 {
		x := form.Interface()
		return c.exprValue(c.TypeOf(x), x)
	}

	// compile phase
	expr := c.Compile(form)

	if g.Options&OptKeepUntyped == 0 && expr != nil && expr.Untyped() {
		expr.ConstTo(expr.DefaultType())
	}
	if g.Options&OptShowCompile != 0 {
		g.Fprintf(g.Stdout, "%v\n", expr)
	}
	return expr
}

// run without debugging. to execute with single-step debugging, use Interp.DebugExpr() instead
func (ir *Interp) RunExpr1(e *Expr) (r.Value, xr.Type) {
	if e == nil {
		return None, nil
	}
	// do NOT use e.AsX1(), it converts untyped constants to their default type => may overflow
	e.CheckX1()
	vs, ts := ir.RunExpr(e)
	return vs[0], ts[0]
}

// run without debugging. to execute with single-step debugging, use Interp.DebugExpr() instead
func (ir *Interp) RunExpr(e *Expr) ([]r.Value, []xr.Type) {
	if e == nil {
		return nil, nil
	}
	env := ir.PrepareEnv()

	if ir.Comp.Globals.Options&OptKeepUntyped == 0 && e.Untyped() {
		e.ConstTo(e.DefaultType())
	}
	run := env.Run
	run.applyDebugOp(DebugOpContinue)

	defer run.setCurrEnv(run.setCurrEnv(env))

	fun := e.AsXV(COptKeepUntyped)
	v, vs := fun(env)
	return PackValues(v, vs), PackTypes(e.Type, e.Types)
}

// execute with single-step debugging. to run without debugging, use Interp.RunExpr() instead
func (ir *Interp) DebugExpr1(e *Expr) (r.Value, xr.Type) {
	if e == nil {
		return None, nil
	}
	// do NOT use e.AsX1(), it converts untyped constants to their default type => may overflow
	e.CheckX1()
	vs, ts := ir.DebugExpr(e)
	return vs[0], ts[0]
}

// execute with single-step debugging. to run without debugging, use Interp.RunExpr() instead
func (ir *Interp) DebugExpr(e *Expr) ([]r.Value, []xr.Type) {
	if e == nil {
		return nil, nil
	}
	env := ir.PrepareEnv()

	if ir.Comp.Globals.Options&OptKeepUntyped == 0 && e.Untyped() {
		e.ConstTo(e.DefaultType())
	}
	run := env.Run
	run.applyDebugOp(DebugOpStep)
	defer run.setCurrEnv(run.setCurrEnv(env))

	fun := e.AsXV(COptKeepUntyped)
	v, vs := fun(env)
	return PackValues(v, vs), PackTypes(e.Type, e.Types)
}

// combined Parse + Compile + DebugExpr
func (ir *Interp) Debug(src string) ([]r.Value, []xr.Type) {
	return ir.DebugExpr(ir.Compile(src))
}

// set CurrEnv, returns previous value
func (g *Run) setCurrEnv(env *Env) *Env {
	old := g.CurrEnv
	g.CurrEnv = env
	return old
}

// ================ PrepareEnv() ========================

func (ir *Interp) PrepareEnv() *Env {
	// allocate Env.Ints[] in large chunks while we can:
	// once an Env.Ints[idx] address is taken, we can no longer reallocate it
	return ir.prepareEnv(16, 1024)
}

func (ir *Interp) prepareEnv(minValDelta int, minIntDelta int) *Env {
	c := ir.Comp
	env := ir.env
	// usually we know at Env creation how many slots are needed in c.Env.Binds
	// but here we are modifying an existing Env...
	if minValDelta < 0 {
		minValDelta = 0
	}
	if minIntDelta < 0 {
		minIntDelta = 0
	}
	capacity, min := cap(env.Vals), c.BindNum
	// c.Debugf("prepareEnv() before: c.BindNum = %v, minValDelta = %v, len(env.Binds) = %v, cap(env.Binds) = %v, env = %p", c.BindNum, minValDelta, len(env.Binds), cap(env.Binds), env)

	if capacity < min {
		capacity *= 2
		if capacity < min {
			capacity = min
		}
		if capacity-cap(env.Vals) < minValDelta {
			capacity = cap(env.Vals) + minValDelta
		}
		binds := make([]r.Value, min, capacity)
		copy(binds, env.Vals)
		env.Vals = binds
	}
	if len(env.Vals) < min {
		env.Vals = env.Vals[0:min:cap(env.Vals)]
	}
	// c.Debugf("prepareEnv() after:  c.BindNum = %v, minDelta = %v, len(env.Binds) = %v, cap(env.Binds) = %v, env = %p", c.BindNum, minDelta, len(env.Binds), cap(env.Binds), env)

	capacity, min = cap(env.Ints), c.IntBindNum
	if capacity < min {
		if env.IntAddressTaken {
			c.Errorf("internal error: attempt to reallocate Env.Ints[] after one of its addresses was taken")
		}
		capacity *= 2
		if capacity < min {
			capacity = min
		}
		if capacity-cap(env.Ints) < minIntDelta {
			capacity = cap(env.Ints) + minIntDelta
		}
		binds := make([]uint64, min, capacity)
		copy(binds, env.Ints)
		env.Ints = binds
	}
	if len(env.Ints) < min {
		env.Ints = env.Ints[0:min:cap(env.Ints)] // does not reallocate
	}
	if env.IntAddressTaken {
		c.IntBindMax = cap(env.Ints)
	}
	g := env.Run
	// do NOT set g.CurrEnv = env, it messes up the call stack. done by Interp.RunExpr* and Interp.DebugExpr*
	// g.CurrEnv = env
	// in case we received a SigInterrupt in the meantime
	g.Signals.Sync = SigNone
	g.Signals.Async = SigNone
	if g.Options&OptDebugger != 0 {
		// for debugger
		env.DebugComp = c
	} else {
		env.DebugComp = nil
	}
	return env
}

// ====================== Repl() and friends =====================

var historyfile = Subdir(UserHomeDir(), ".gomacro_history")

func (ir *Interp) ReplStdin() {
	g := ir.Comp.CompGlobals

	if g.Options&OptShowPrompt != 0 {
		g.Fprintf(g.Stdout, `// GOMACRO, an interactive Go interpreter with generics and macros
// Copyright (C) 2017-2018 Massimiliano Ghilardi <https://github.com/cosmos72/gomacro>
// License MPL v2.0+: Mozilla Public License version 2.0 or later <http://mozilla.org/MPL/2.0/>
// This is free software with ABSOLUTELY NO WARRANTY.
//
// Type %chelp for help
`, g.ReplCmdChar)
	}
	tty, _ := MakeTtyReadline(historyfile)
	defer tty.Close(historyfile) // restore normal tty mode

	ch := StartSignalHandler(ir.Interrupt)
	defer StopSignalHandler(ch)

	savetty := g.Readline
	g.Readline = tty
	defer func() {
		g.Readline = savetty
	}()
	tty.Term.SetWordCompleter(ir.CompleteWords)

	g.Line = 0
	for ir.ReadParseEvalPrint() {
		g.Line = 0
	}
	os.Stdout.WriteString("\n")
}

func (ir *Interp) Repl(in *bufio.Reader) {
	g := ir.Comp.CompGlobals

	r := MakeBufReadline(in, g.Stdout)

	ch := StartSignalHandler(ir.Interrupt)
	defer StopSignalHandler(ch)

	savetty := g.Readline
	g.Readline = r
	defer func() {
		g.Readline = savetty
	}()

	for ir.ReadParseEvalPrint() {
	}
}

func (ir *Interp) ReadParseEvalPrint() (callAgain bool) {
	src, firstToken := ir.Read()
	if firstToken < 0 {
		// skip comment-only lines and continue, but fail on EOF or other errors
		return len(src) != 0
	}
	return ir.ParseEvalPrint(src)
}

func (ir *Interp) ParseEvalPrint(src string) (callAgain bool) {
	if len(src) == 0 || len(strings.TrimSpace(src)) == 0 {
		return true // no input => no form
	}

	t1, trap, duration := ir.beforeEval()
	defer ir.afterEval(src, &callAgain, &trap, t1, duration)

	src, opt := ir.Cmd(src)

	callAgain = opt&CmdOptQuit == 0
	if len(src) == 0 || !callAgain {
		trap = false // no panic happened
		return callAgain
	}

	g := &ir.Comp.Globals
	if toenable := cmdOptForceEval(g, opt); toenable != 0 {
		defer func() {
			g.Options |= toenable
		}()
	}

	ir.env.Run.CmdOpt = opt // store options where Interp.Interrupt() can find them

	// parse + macroexpansion
	form := ir.Parse(src)

	// compile
	expr := ir.CompileAst(form)

	// run expression
	values, types := ir.RunExpr(expr)

	// print phase
	g.Print(values, types)

	trap = false // no panic happened
	return callAgain
}

func (ir *Interp) beforeEval() (t1 time.Time, trap bool, duration bool) {
	g := &ir.Comp.Globals
	trap = g.Options&OptTrapPanic != 0
	duration = g.Options&OptShowTime != 0
	if duration {
		t1 = time.Now()
	}
	return t1, trap, duration
}

func (ir *Interp) afterEval(src string, callAgain *bool, trap *bool, t1 time.Time, duration bool) {
	g := &ir.Comp.Globals
	g.IncLine(src)
	if *trap {
		rec := recover()
		if g.Options&OptPanicStackTrace != 0 {
			g.Fprintf(g.Stderr, "%v\n%s", rec, debug.Stack())
		} else {
			g.Fprintf(g.Stderr, "%v\n", rec)
		}
		*callAgain = true
	}
	if duration {
		delta := time.Since(t1)
		g.Debugf("eval time %v", delta)
	}
}

func cmdOptForceEval(g *Globals, opt CmdOpt) (toenable Options) {
	if opt&CmdOptForceEval != 0 {
		// temporarily disable collection of declarations and statements,
		// and temporarily re-enable eval (i.e. disable macroexpandonly)
		const todisable = OptMacroExpandOnly | OptCollectDeclarations | OptCollectStatements
		if g.Options&todisable != 0 {
			g.Options &^= todisable
			return todisable
		}
	}
	return 0
}

// implement code completion API github.com/pererh/liner.WordCompleter
// Currently only supports global symbols and imported packages,
// optionally followed by a dot-separated sequence of field or method names,
// including embedded fields and wrapper methods.
func (ir *Interp) CompleteWords(line string, pos int) (head string, completions []string, tail string) {
	if pos > len(line) {
		pos = len(line)
	}
	head = line[:pos]
	tail = line[pos:]
	words := strings.Split(head, ".")
	n := len(words)
	// find the longest sequence of ident.ident.ident...

	for i := n - 1; i >= 0; i-- {
		// ignore spaces before and after identifiers
		words[i] = strings.TrimSpace(words[i])

		if i == n-1 && len(words[i]) == 0 {
			// last word can be empty: it means TAB immediately after '.'
			continue
		}
		word := TailIdentifier(words[i])
		if len(word) != len(words[i]) {
			if len(word) != 0 {
				words[i] = word
			} else {
				i++
			}
			words = words[i:]
			break
		}
	}
	completions = ir.Comp.CompleteWords(words)
	if len(completions) != 0 {
		fixed := len(head) - len(TailIdentifier(head))
		pos := strings.LastIndexByte(head, '.')
		if pos >= 0 && pos >= fixed {
			head = head[:pos+1]
		} else {
			head = head[:fixed]
		}
	}
	return head, completions, tail
}

// implement code completion on ident.ident.ident.ident...
func (c *Comp) CompleteWords(words []string) []string {
	var completions []string
	switch len(words) {
	case 0:
	case 1:
		completions = c.completeWord(words[0])
	default:
		var node interface{}
		if sym := c.TryResolve(words[0]); sym != nil {
			node = &sym.Bind
		} else if typ := c.TryResolveType(words[0]); typ != nil {
			node = typ
		} else {
			break
		}
		completions = c.completeWords(node, words[1:])
	}
	return completions
}

func (c *Comp) completeWords(node interface{}, words []string) []string {
	i, n := 0, len(words)
	for i+1 < n {
		switch obj := node.(type) {
		case *Bind:
			if obj.Const() {
				if imp, ok := obj.Value.(*Import); ok {
					// complete on imported package contents
					node = imp
					continue
				}
			}
			// complete on symbol type
			node = obj.Type
			continue
		case *Import:
			if i != 0 {
				break
			} else if bind := obj.Binds[words[i]]; bind != nil {
				// complete on imported package binds
				node = bind
				i++
				continue
			} else if typ := obj.Types[words[i]]; typ != nil {
				// complete on imported package types
				node = typ
				i++
				continue
			}
		case xr.Type:
			field, fieldok, _, _, err := c.TryLookupFieldOrMethod(obj, words[i])
			if err != nil {
				break
			} else if fieldok {
				node = field.Type
				i++
				continue
			}
			// {type,value}.method.anything will never compile
		}
		return nil
	}
	return c.completeLastWord(node, words[i])
}

var keywords []string

func init() {
	lo, hi := token.BREAK, token.VAR+1
	keywords = make([]string, hi-lo+1)
	for tok := lo; tok < hi; tok++ {
		keywords[tok-lo] = tok.String()
	}
	keywords[hi-lo] = "macro"
}

// complete a single, partial word
func (c *Comp) completeWord(word string) []string {
	var completions []string
	if size := len(word); size != 0 {
		// complete binds
		for co := c; co != nil; co = co.Outer {
			for name := range co.Binds {
				if len(name) >= size && name[:size] == word {
					completions = append(completions, name)
				}
			}
		}
		// complete types
		for co := c; co != nil; co = co.Outer {
			for name := range co.Types {
				if len(name) >= size && name[:size] == word {
					completions = append(completions, name)
				}
			}
		}
		// complete keywords
		for _, name := range keywords {
			if len(name) >= size && name[:size] == word {
				completions = append(completions, name)
			}
		}
	}
	return sortUnique(completions)
}

// complete the last partial word of a sequence ident.ident.ident...
func (c *Comp) completeLastWord(node interface{}, word string) []string {
	var completions []string
	size := len(word)
	for {
		switch obj := node.(type) {
		case *Bind:
			if obj.Const() {
				if imp, ok := obj.Value.(*Import); ok {
					// complete on imported package contents
					node = imp
					continue
				}
			}
			// complete on symbol type
			node = obj.Type
			continue
		case *Import:
			for name := range obj.Binds {
				if len(name) >= size && name[:size] == word {
					completions = append(completions, name)
				}
			}
			for name := range obj.Types {
				if len(name) >= size && name[:size] == word {
					completions = append(completions, name)
				}
			}
		case xr.Type:
			completions = c.listFieldsAndMethods(obj, word)
		}
		break
	}
	return sortUnique(completions)
}

func sortUnique(vec []string) []string {
	if n := len(vec); n > 1 {
		sort.Strings(vec)
		prev := vec[0]
		j := 1
		for i := 1; i < n; i++ {
			if s := vec[i]; s != prev {
				vec[j] = s
				prev = s
				j++
			}
		}
		vec = vec[:j]
	}
	return vec
}
