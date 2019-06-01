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
 * macroexpand.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	etoken "github.com/cosmos72/gomacro/go/etoken"
)

type macroExpandCtx struct {
	env *Env
}

// MacroExpandCodewalk traverses the whole AST tree using pre-order traversal,
// and replaces each node with the result of MacroExpand(node).
// It implements the macroexpansion phase
func (env *Env) MacroExpandCodewalk(in ast.Node) (out ast.Node, anythingExpanded bool) {
	if in == nil {
		return nil, false
	}
	var form Ast = ToAst(in)
	form, anythingExpanded = env.MacroExpandAstCodewalk(form)
	out = ToNode(form)
	// if !anythingExpanded {
	//    env.Debugf("MacroExpand1() nothing to expand: %v <%v>", out, r.TypeOf(out))
	//}
	return out, anythingExpanded
}

func (env *Env) MacroExpandAstCodewalk(in Ast) (out Ast, anythingExpanded bool) {
	return env.macroExpandAstCodewalk(in, 0)
}

func (env *Env) macroExpandAstCodewalk(in Ast, quasiquoteDepth int) (out Ast, anythingExpanded bool) {
	if in == nil || in.Size() == 0 {
		return in, false
	}
	if quasiquoteDepth <= 0 {
		if env.Options&OptDebugMacroExpand != 0 {
			env.Debugf("MacroExpandCodewalk: qq = %d, macroexpanding %v", quasiquoteDepth, in.Interface())
		}
		in, anythingExpanded = env.macroExpandAst(in)
	}
	if in != nil {
		in = UnwrapTrivialAst(in)
	}
	if in == nil {
		return in, anythingExpanded
	}
	saved := in

	if expr, ok := in.(UnaryExpr); ok {
		isBlockWithinExpr := false
		switch expr.X.Op {
		case etoken.MACRO:
			isBlockWithinExpr = true
		case etoken.QUOTE:
			// QUOTE prevents macroexpansion only if found outside any QUASIQUOTE
			if quasiquoteDepth == 0 {
				return saved, anythingExpanded
			}
		case etoken.QUASIQUOTE:
			// extract the body of QUASIQUOTE
			quasiquoteDepth++
		case etoken.UNQUOTE, etoken.UNQUOTE_SPLICE:
			// extract the body of UNQUOTE or UNQUOTE_SPLICE
			quasiquoteDepth--
		default:
			goto Recurse
		}
		inChild := UnwrapTrivialAst(in.Get(0).Get(1))
		outChild, expanded := env.macroExpandAstCodewalk(inChild, quasiquoteDepth)
		if isBlockWithinExpr {
			return outChild, expanded
		} else {
			out := in
			if expanded {
				out = MakeQuote2(expr, outChild.(AstWithNode))
			}
			return out, expanded
		}
	}
Recurse:
	if in == nil {
		return saved, anythingExpanded
	}
	if env.Options&OptDebugMacroExpand != 0 {
		env.Debugf("MacroExpandCodewalk: qq = %d, recursing on %v", quasiquoteDepth, in)
	}
	out = in.New()
	n := in.Size()
	if outSlice, canAppend := out.(AstWithSlice); canAppend {
		// New() returns zero-length slice... resize it
		for i := 0; i < n; i++ {
			outSlice = outSlice.Append(nil)
		}
		out = outSlice
	}
	for i := 0; i < n; i++ {
		child := UnwrapTrivialAst(in.Get(i))
		if child != nil {
			expanded := false
			if child.Size() != 0 {
				child, expanded = env.macroExpandAstCodewalk(child, quasiquoteDepth)
			}
			if expanded {
				anythingExpanded = true
			}
		}
		out.Set(i, child)
	}
	if env.Options&OptDebugMacroExpand != 0 {
		env.Debugf("MacroExpandCodewalk: qq = %d, expanded to %v", quasiquoteDepth, out)
	}
	return out, anythingExpanded
}

// MacroExpand repeatedly invokes MacroExpand1
// as long as the node represents a macro call.
// it returns the resulting node.
func (env *Env) MacroExpand(in ast.Node) (out ast.Node, everExpanded bool) {
	if in == nil {
		return nil, false
	}
	inAst := ToAst(in)
	outAst, everExpanded := env.macroExpandAst(inAst)
	out = ToNode(outAst)
	// if !everExpanded {
	//    env.Debugf("MacroExpand1() not a macro: %v <%v>", out, r.TypeOf(out))
	//}
	return out, everExpanded
}

func (env *Env) macroExpandAst(form Ast) (out Ast, everExpanded bool) {
	var expanded bool
	for {
		form, expanded = env.macroExpandAstOnce(form)
		if !expanded {
			return form, everExpanded
		}
		everExpanded = true
	}
}

// if node represents a macro call, MacroExpand1 executes it
// and returns the resulting node.
// Otherwise returns the node argument unchanged
func (env *Env) MacroExpand1(in ast.Node) (out ast.Node, expanded bool) {
	if in == nil {
		return nil, false
	}
	var form Ast = ToAst(in)
	form, expanded = env.macroExpandAstOnce(form)
	out = ToNode(form)
	// if !expanded {
	//    env.Debugf("MacroExpand1: not a macro: %v <%v>", out, r.TypeOf(out))
	//}
	return out, expanded
}

//
func (env *Env) extractMacroCall(form Ast) Macro {
	form = UnwrapTrivialAst(form)
	switch form := form.(type) {
	case Ident:
		bind, found := env.resolveIdentifier(form.X)
		if found && bind.Kind() == r.Struct {
			switch value := bind.Interface().(type) {
			case Macro:
				if env.Options&OptDebugMacroExpand != 0 {
					env.Debugf("MacroExpand1: found macro: %v", form.X.Name)
				}
				return value
			}
		}
	}
	return Macro{}
}

func (env *Env) macroExpandAstOnce(in Ast) (out Ast, expanded bool) {
	if in == nil {
		return nil, false
	}
	// unwrap trivial nodes: DeclStmt, ParenExpr, ExprStmt
	in = UnwrapTrivialAstKeepBlocks(in)
	ins, ok := in.(AstWithSlice)
	if !ok {
		return in, false
	}
	if env.Options&OptDebugMacroExpand != 0 {
		env.Debugf("MacroExpand1: found list: %v", ins.Interface())
	}
	outs := ins.New().(AstWithSlice)
	n := ins.Size()

	// since macro calls are sequences of statements,
	// we must scan the whole list,
	// consume it as needed by the macros we find,
	// and build a new list accumulating the results of macroexpansion
	for i := 0; i < n; i++ {
		elt := ins.Get(i)
		macro := env.extractMacroCall(elt)
		if macro.closure == nil {
			outs = outs.Append(elt)
			continue
		}
		argn := macro.argNum
		leftn := n - i - 1
		var args []r.Value
		if argn > leftn {
			args := make([]r.Value, leftn+1) // include the macro itself
			for j := 0; j <= leftn; j++ {
				args[j] = r.ValueOf(ins.Get(i + j).Interface())
			}
			env.Errorf("not enough arguments for macroexpansion of %v: expecting %d, found %d", args, macro.argNum, leftn)
			return in, false
		}
		if env.Options&OptDebugMacroExpand != 0 {
			env.Debugf("MacroExpand1: found macro call %v at %d-th position of %v", elt.Interface(), i, ins.Interface())
		}
		// wrap each ast.Node into a reflect.Value
		args = make([]r.Value, argn)
		for j := 0; j < argn; j++ {
			args[j] = r.ValueOf(ToNode(ins.Get(i + j + 1)))
		}
		// invoke the macro
		results := macro.closure(args)
		if env.Options&OptDebugMacroExpand != 0 {
			env.Debugf("MacroExpand1: macro expanded to: %v", results)
		}
		var out Ast
		switch len(results) {
		default:
			args = append([]r.Value{r.ValueOf(elt.Interface())}, args...)
			env.Warnf("macroexpansion returned %d values, using only the first one: %v %v returned %v",
				len(results), args, results)
			fallthrough
		case 1:
			any := results[0].Interface()
			if any != nil {
				out = AnyToAst(any, "macroexpansion")
				break
			}
			fallthrough
		case 0:
			// do not insert nil nodes... they would wreak havok, convert them to the identifier nil
			out = Ident{&ast.Ident{Name: "nil"}}
		}
		outs = outs.Append(out)
		i += argn
		expanded = true
	}
	if !expanded {
		return in, false
	}
	if outs.Size() == 0 {
		return EmptyStmt{&ast.EmptyStmt{}}, true
	}
	return UnwrapTrivialAst(outs), true
}
