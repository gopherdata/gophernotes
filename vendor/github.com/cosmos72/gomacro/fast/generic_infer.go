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
 * generic_infer.go
 *
 *  Created on Jun 06, 2018
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"fmt"
	"go/ast"
	"go/token"
	r "reflect"

	"github.com/cosmos72/gomacro/base/untyped"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type inferType struct {
	Type    xr.Type
	Untyped untyped.Kind // for untyped literals
	Value   I            // in case we infer a constant, not a type
	Exact   bool
}

func (inf *inferType) String() string {
	if inf.Value != nil {
		return fmt.Sprint(inf.Value)
	}
	var s string
	if inf.Type != nil {
		s = inf.Type.String()
	} else {
		s = inf.Untyped.String()
	}
	return "<" + s + ">"
}

// type inference on generic functions
type inferFuncType struct {
	comp     *Comp
	tfun     *GenericFunc
	funcname string
	inferred map[string]inferType
	patterns []ast.Expr
	targs    []inferType
	call     *ast.CallExpr // for error messages
}

func (inf *inferFuncType) String() string {
	return inf.tfun.Signature(inf.funcname)
}

func (c *Comp) inferGenericFunc(call *ast.CallExpr, fun *Expr, args []*Expr) *Expr {
	tfun, ok := fun.Value.(*GenericFunc)
	if !ok {
		c.Errorf("internal error: Comp.inferGenericFunc() invoked on non-generic function %v: %v", fun.Type, call.Fun)
	}
	var upc *Comp
	var funcname string
	{
		ident, ok := call.Fun.(*ast.Ident)
		if !ok {
			c.Errorf("unimplemented type inference on non-name generic function %v: %v", call.Fun, call)
		}
		if fun.Sym == nil {
			c.Errorf("unimplemented type inference on non-symbol generic function %v %#v: %v", call.Fun, fun, call)
		}
		// find the scope where fun is declared
		funcname = ident.Name
		fbind := &fun.Sym.Bind
		for upc = c; upc != nil; upc = upc.Outer {
			if bind, ok := upc.Binds[funcname]; ok && bind.Name == fbind.Name && bind.Desc == fbind.Desc && bind.Type.IdenticalTo(fbind.Type) {
				break
			}
		}
	}
	if upc == nil {
		c.Errorf("internal error: Comp.inferGenericFunc() failed to determine the scope containing generic function declaration: %v", call.Fun)
	}

	master := tfun.Master
	typ := master.Decl.Type

	var patterns []ast.Expr
	ellipsis := call.Ellipsis != token.NoPos
	variadic := false
	// collect generic function param types expressions
	if fields := typ.Params; fields != nil {
		if n := len(fields.List); n != 0 {
			_, variadic = fields.List[n-1].Type.(*ast.Ellipsis)
			for _, field := range fields.List {
				for _ = range field.Names {
					patterns = append(patterns, field.Type)
				}
			}
		}
	}
	if variadic && !ellipsis {
		c.Errorf("unimplemented type inference on variadic generic function: %v", call)
	} else if !variadic && ellipsis {
		c.Errorf("invalid use of ... in call to non-variadic generic function: %v", call)
	}

	// collect call arg types
	nargs := len(args)
	var targs []inferType
	if nargs == 1 {
		arg := args[0]
		nargs = arg.NumOut()
		targs = make([]inferType, nargs)
		for i := 0; i < nargs; i++ {
			targs[i] = inferType{Type: arg.Out(i)}
		}
	} else {
		targs = make([]inferType, nargs)
		for i, arg := range args {
			if kind := arg.UntypedKind(); kind != untyped.None {
				targs[i] = inferType{Untyped: kind}
			} else {
				targs[i] = inferType{Type: arg.Type}
			}
		}
	}
	if nargs != len(patterns) {
		c.Errorf("generic function %v has %d params, cannot call with %d values: %v", tfun, len(patterns), nargs, call)
	}
	inferred := make(map[string]inferType)
	for _, name := range master.Params {
		inferred[name] = inferType{}
	}
	inf := inferFuncType{comp: c, tfun: tfun, funcname: funcname, inferred: inferred, patterns: patterns, targs: targs, call: call}
	vals, types := inf.args()
	maker := &genericMaker{
		comp: upc, sym: fun.Sym, ifun: fun.Sym.Value,
		exprs: nil, vals: vals, types: types,
		ikey: GenericKey(vals, types),
		pos:  inf.call.Pos(),
	}
	return c.genericFunc(maker, call)
}

// infer type of generic function from arguments
func (inf *inferFuncType) args() (vals []I, types []xr.Type) {
	exact := false // allow implicit type conversions

	// first pass: types and typed constants
	for i, targ := range inf.targs {
		node := inf.patterns[i]
		if targ.Type != nil {
			inf.arg(node, targ.Type, exact)
		} else if targ.Untyped != untyped.None {
			// skip untyped constant, handled below
		} else if targ.Value != nil {
			inf.constant(node, targ.Value, exact)
		} else {
			inf.fail(node, targ)
		}
	}

	// second pass: untyped constants
	for i, targ := range inf.targs {
		if targ.Type == nil && targ.Untyped != untyped.None {
			inf.untyped(inf.patterns[i], targ.Untyped, exact)
		}
	}

	params := inf.tfun.Master.Params
	n := len(params)
	vals = make([]I, n)
	types = make([]xr.Type, n)
	for i, name := range params {
		inferred, ok := inf.inferred[name]
		if !ok || inferred.Type == nil {
			inf.comp.Errorf("failed to infer %v in call to generic function: %v", name, inf.call)
		}
		types[i] = inferred.Type
		vals[i] = inferred.Value
	}
	return vals, types
}

// partially infer type of generic function for a single parameter
func (inf *inferFuncType) arg(pattern ast.Expr, targ xr.Type, exact bool) {
	stars := 0
	for {
		if targ == nil {
			inf.fail(pattern, targ)
		}
		if node, ok := pattern.(*ast.Ident); ok {
			inf.ident(node, targ, exact)
			break
		}
		switch node := pattern.(type) {
		case *ast.ArrayType:
			pattern, targ, exact = inf.arrayType(node, targ, exact)
			continue
		case *ast.ChanType:
			pattern, targ, exact = inf.chanType(node, targ, exact)
			continue
		case *ast.FuncType:
			pattern, targ, exact = inf.funcType(node, targ, exact)
			if pattern != nil {
				continue
			}
		case *ast.IndexExpr:
			// function's parameter is itself a generic
			pattern, targ, exact = inf.genericType(node, targ, exact)
			if pattern != nil {
				continue
			}
		case *ast.InterfaceType:
			pattern, targ, exact = inf.interfaceType(node, targ, exact)
			if pattern != nil {
				continue
			}
		case *ast.MapType:
			pattern, targ, exact = inf.mapType(node, targ, exact)
			continue
		case *ast.ParenExpr:
			pattern = node.X
			continue
		case *ast.SelectorExpr:
			// packagename.typename
			pattern, targ, exact = inf.selector(node, targ, exact)
			if pattern != nil {
				continue
			}
		case *ast.StarExpr:
			inf.is(pattern, targ, r.Ptr)
			pattern, targ = node.X, targ.Elem()
			if stars != 0 {
				exact = true
			}
			stars++
			continue
		case *ast.StructType:
			pattern, targ, exact = inf.structType(node, targ, exact)
			if pattern != nil {
				continue
			}
		default:
			inf.unimplemented(node, targ)
		}
		break
	}
}

// partially infer type of generic function from an array or slice parameter
func (inf *inferFuncType) arrayType(node *ast.ArrayType, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {
	if node.Len == nil {
		inf.is(node, targ, r.Slice)
	} else {
		inf.is(node, targ, r.Array)
		if _, ok := node.Len.(*ast.Ellipsis); !ok {
			// [n]array
			inf.constant(node.Len, targ.Len(), exact)
		}
	}
	return node.Elt, targ.Elem(), true
}

// partially infer type of generic function for a channel parameter
func (inf *inferFuncType) chanType(node *ast.ChanType, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {
	inf.is(node, targ, r.Chan)
	tdir := targ.ChanDir()
	dir := reflectChanDir(node.Dir)
	if dir&tdir == 0 || (exact && dir != tdir) {
		inf.fail(node, targ)
	}
	return node.Value, targ.Elem(), true
}

// partially infer type of generic function for a constant parameter
func (inf *inferFuncType) constant(node ast.Expr, val I, exact bool) {
	// TODO
	inf.comp.ErrorAt(node.Pos(), "unimplemented type inference: generic function with parameter type %v and argument %v: %v",
		node, val, inf.call)
}

// partially infer type of generic function for a func parameter
func (inf *inferFuncType) funcType(node *ast.FuncType, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {
	// TODO
	return inf.unimplemented(node, targ)
}

// partially infer type of generic function for an identifier parameter
func (inf *inferFuncType) ident(node *ast.Ident, targ xr.Type, exact bool) {
	c := inf.comp
	name := node.Name
	inferred, ok := inf.inferred[name]
	if !ok {
		// name must be an existing type
		t := c.TryResolveType(name)
		if t != nil {
			if !targ.AssignableTo(t) {
				inf.comp.ErrorAt(node.Pos(),
					"type inference: in %v, mismatched types for %v: %v cannot be assigned to %v: %v",
					inf, name, targ, t, inf.call)
			}
		}
		return
	}

	// inferring one of the function generic parameters
	inf.combine(node, &inferred, inferType{Type: targ, Exact: exact})
	inf.inferred[name] = inferred

}

func (inf *inferFuncType) untyped(node ast.Expr, kind untyped.Kind, exact bool) {
	ident, ok := node.(*ast.Ident)
	if !ok {
		inf.fail(node, kind)
	}
	inf.unimplemented(ident, kind)
}

func (inf *inferFuncType) combine(node ast.Expr, inferred *inferType, with inferType) {
	targ := with.Type
	exact := with.Exact
	if inferred.Type == nil {
		inferred.Type = targ
	} else if !inferred.Type.IdenticalTo(targ) {
		if exact && inferred.Exact {
			inf.fail3(node, inferred, targ)
		}
		fwd := targ.AssignableTo(inferred.Type)
		rev := inferred.Type.AssignableTo(targ)
		if inferred.Exact {
			if fwd {
				inf.fail3(node, inferred, targ)
			}
		} else if exact {
			if rev {
				inferred.Type = targ
			} else {
				inf.fail3(node, inferred, targ)
			}
		} else {
			if fwd && rev {
				if !targ.Named() {
					inferred.Type = targ
				}
			} else if fwd {
			} else if rev {
				inferred.Type = targ
			} else {
				inf.fail3(node, inferred, targ)
			}
		}
	}
	if exact {
		inferred.Exact = true
	}
}

// partially infer type of generic function for an interface parameter
func (inf *inferFuncType) interfaceType(node *ast.InterfaceType, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {
	// TODO
	return inf.unimplemented(node, targ)
}

// partially infer type of generic function for a map parameter
func (inf *inferFuncType) mapType(node *ast.MapType, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {
	inf.is(node, targ, r.Map)
	inf.arg(node.Key, targ.Key(), true)
	return node.Value, targ.Elem(), true
}

// partially infer type of generic function for an imported type
func (inf *inferFuncType) selector(node *ast.SelectorExpr, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {
	// TODO
	return inf.unimplemented(node, targ)
}

// partially infer type of generic function for a struct parameter
func (inf *inferFuncType) structType(node *ast.StructType, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {
	// TODO
	return inf.unimplemented(node, targ)
}

// partially infer type of generic function for a generic parameter
func (inf *inferFuncType) genericType(node *ast.IndexExpr, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {
	// TODO
	return inf.unimplemented(node, targ)
}

func (inf *inferFuncType) is(node ast.Expr, targ xr.Type, kind r.Kind) {
	if targ.Kind() != kind {
		inf.fail(node, targ)
	}
}

func (inf *inferFuncType) fail(node ast.Expr, targ I) {
	inf.comp.ErrorAt(node.Pos(),
		"type inference: in %v, parameter %v cannot match argument type %v: %v",
		inf, node, targ, inf.call)
}

func (inf *inferFuncType) fail3(node ast.Expr, tinferred *inferType, targ xr.Type) {
	inf.comp.ErrorAt(node.Pos(),
		"type inference: in %v, parameter %v cannot match both %v and <%v>: %v",
		inf, node, tinferred, targ, inf.call)
}

func (inf *inferFuncType) unimplemented(node ast.Expr, targ I) (ast.Expr, xr.Type, bool) {
	inf.comp.ErrorAt(node.Pos(), "unimplemented type inference: in %v, parameter type %v with argument type %v: %v",
		inf, node, targ, inf.call)
	return nil, nil, false
}

var chandirs = map[ast.ChanDir]r.ChanDir{
	ast.RECV:            r.RecvDir,
	ast.SEND:            r.SendDir,
	ast.RECV | ast.SEND: r.BothDir,
}

func reflectChanDir(dir ast.ChanDir) r.ChanDir {
	return chandirs[dir]
}
