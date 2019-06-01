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
 * generic_maker.go
 *
 *  Created on Jun 16, 2018
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	r "reflect"
	"sort"
	"strings"

	"github.com/cosmos72/gomacro/ast2"
	"github.com/cosmos72/gomacro/base"
	etoken "github.com/cosmos72/gomacro/go/etoken"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// enable C++-style generics?
const GENERICS_V1_CXX = etoken.GENERICS_V1_CXX

// enable "contracts are interfaces" generics?
const GENERICS_V2_CTI = etoken.GENERICS_V2_CTI

type genericMaker struct {
	comp  *Comp
	sym   *Symbol
	ifun  I
	exprs []ast.Expr
	vals  []I
	types []xr.Type
	ikey  I
	name  string
	pos   token.Pos
}

type genericTypeCandidate struct {
	decl   GenericTypeDecl
	params []string
	vals   []I
	types  []xr.Type
}

type genericFuncCandidate struct {
	decl  GenericFuncDecl
	vals  []I
	types []xr.Type
}

func (special *genericFuncCandidate) injectBinds(c *Comp) {
	for i, name := range special.decl.Params {
		t := special.types[i]
		if val := special.vals[i]; val != nil {
			c.DeclConst0(name, t, val)
		} else {
			c.declTypeAlias(name, t)
		}
	}
}

func (special *genericTypeCandidate) injectBinds(c *Comp) {
	for i, name := range special.decl.Params {
		t := special.types[i]
		if val := special.vals[i]; val != nil {
			c.DeclConst0(name, t, val)
		} else {
			c.declTypeAlias(name, t)
		}
	}
}

// return the qualified name of the function or type to instantiate, for example "Pair#[int,string]"
func (maker *genericMaker) String() string {
	if len(maker.name) != 0 {
		return maker.name
	}
	var buf bytes.Buffer
	buf.WriteString(maker.sym.Name)
	buf.WriteString("#[")

	for i, val := range maker.vals {
		if i != 0 {
			buf.WriteByte(',')
		}
		if val == nil {
			val = maker.types[i].ReflectType()
		}
		fmt.Fprint(&buf, val)
	}
	buf.WriteByte(']')
	maker.name = buf.String()
	return maker.name
}

func (c *Comp) genericMaker(node *ast.IndexExpr, which BindClass) *genericMaker {
	name, genericArgs, ok := splitGenericArgs(node)
	if !ok {
		return nil
	}
	sym, upc := c.tryResolve(name)
	if sym == nil {
		c.Errorf("undefined identifier: %v", name)
	}
	n := len(genericArgs)
	var params []string
	ifun := sym.Value
	ok = false
	if ifun != nil && sym.Desc.Class() == which {
		switch which {
		case GenericFuncBind:
			fun, _ := ifun.(*GenericFunc)
			ok = fun != nil
			if ok {
				params = fun.Master.Params
			}
		case GenericTypeBind:
			typ, _ := ifun.(*GenericType)
			ok = typ != nil
			if ok {
				params = typ.Master.Params
			}
		}
	}
	if !ok {
		c.Errorf("symbol is not a %v, cannot use #[...] on it: %s", which, name)
	}
	if n != len(params) {
		c.Errorf("%v expects exactly %d generic parameters %v, found %d: %v", which, len(params), params, n, node)
	}
	vals := make([]I, n)
	types := make([]xr.Type, n)

	// make a copy of genericArgs, then replace constant expressions with their values
	genericArgs = append([]ast.Expr(nil), genericArgs...)

	for i, genericArg := range genericArgs {
		e, t := c.Expr1OrType(genericArg)
		if e != nil {
			if !e.Const() {
				c.Errorf("argument of generic function %q is not a constant: %v", name, genericArg)
			}
			// UntypedLit is unsuitable as map key, because its == is not usable
			vals[i] = e.EvalConst(COptDefaults)
			types[i] = e.Type // also remember the type
			genericArgs[i] = c.constToAstExpr(vals[i], genericArg.Pos())
		} else {
			types[i] = t
		}
	}
	return &genericMaker{upc, sym, ifun, genericArgs, vals, types, GenericKey(vals, types), "", node.Pos()}
}

func GenericKey(vals []I, types []xr.Type) I {
	// slices cannot be used as map keys. use an array and reflection
	key := r.New(r.ArrayOf(len(types), rtypeOfInterface)).Elem()

	for i, t := range types {
		if val := vals[i]; val == nil {
			key.Index(i).Set(r.ValueOf(xr.MakeKey(t)))
		} else {
			key.Index(i).Set(r.ValueOf(val))
		}
	}
	return key.Interface()
}

// convert true to &ast.Ident{Name: "true"}, convert false similarly,
// convert integers to &ast.BasicLit{Kind: token.INT, Value: fmt.Sprint(val)}
// convert float32, float64 and strings analogously,
// convert complex64 and complex128 to &ast.BinaryExpr{X: real(...), Op: token.Add, Y: imag(...)}
func (c *Comp) constToAstExpr(val interface{}, pos token.Pos) ast.Expr {
	var kind token.Token
	var str string
	v := r.ValueOf(val)
	switch v.Kind() {
	case r.Bool:
		return &ast.Ident{NamePos: pos, Name: fmt.Sprint(val)}
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64,
		r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		kind = token.INT
		str = fmt.Sprint(val)
	case r.Float32, r.Float64:
		kind = token.FLOAT
		str = fmt.Sprintf("%g", val)
	case r.Complex64, r.Complex128:
		return &ast.BinaryExpr{
			X: &ast.BasicLit{
				Kind:     token.FLOAT,
				Value:    fmt.Sprintf("%g", real(v.Complex())),
				ValuePos: pos,
			},
			Op: token.ADD,
			Y: &ast.BasicLit{
				Kind:  token.IMAG,
				Value: fmt.Sprintf("%g", imag(v.Complex())),
			},
		}
	case r.String:
		kind = token.STRING
		str = fmt.Sprintf("%q", val)
	default:
		c.Errorf("unexpected const type, cannot convert to ast.Expr: %v // %T", val, val)
	}
	return &ast.BasicLit{
		Kind:     kind,
		Value:    str,
		ValuePos: pos,
	}
}

func splitGenericArgs(node *ast.IndexExpr) (string, []ast.Expr, bool) {
	if ident, _ := node.X.(*ast.Ident); ident != nil {
		cindex, _ := node.Index.(*ast.CompositeLit)
		if cindex != nil && cindex.Type == nil {
			return ident.Name, cindex.Elts, true
		}
	}
	return "", nil, false
}

func (c *Comp) genericParams(params []ast.Expr, errlabel string, node ast.Node) ([]string, []ast.Expr) {
	names := make([]string, 0, len(params))
	var exprs []ast.Expr
	for i, param := range params {
		switch param := param.(type) {
		case *ast.Ident:
			names = append(names, param.Name)
		case *ast.BadExpr:
		case *ast.CompositeLit:
			exprs = param.Elts
		default:
			c.Errorf("invalid generic %s declaration: generic parameter %d should be *ast.Ident or *ast.CompositeLit, found %T: %v",
				errlabel, i, param, node)
		}
	}
	return names, exprs
}

// return the most specialized function declaration applicable to used params.
// panics if there is no single most specialized declaration.
func (maker *genericMaker) chooseFunc(fun *GenericFunc) (string, *genericFuncCandidate) {
	candidates := map[string]*genericFuncCandidate{
		maker.sym.Name + "#[...]": &genericFuncCandidate{
			decl:  fun.Master,
			vals:  maker.vals,
			types: maker.types,
		},
	}
	g := &maker.comp.Globals
	debug := g.Options&base.OptDebugGenerics != 0
	var ok1, ok2 bool

	if debug {
		g.Debugf("choosing generic function for %s from %d specializations", maker.String(), 1+len(fun.Special))
	}

	for key, special := range fun.Special {
		vals, types, ok := maker.patternMatches(special.Params, special.For, maker.exprs)
		if !ok {
			continue
		}
		// check whether special is more specialized than all other candidates
		for declKey, candidate := range candidates {
			decl := candidate.decl
			if len(decl.For) == 0 {
				ok1, ok2 = false, true
			} else {
				_, _, ok1 = maker.patternMatches(special.Params, special.For, decl.For)
				_, _, ok2 = maker.patternMatches(decl.Params, decl.For, special.For)
			}
			if !ok1 && ok2 {
				// special is more specialized, remove the other
				if debug {
					g.Debugf("generic function %s is more specialized than %s, removing the latter", key, declKey)
				}
				delete(candidates, declKey)
			}
		}
		if debug {
			g.Debugf("adding   generic function specialization  %s to candidates", key)
		}
		candidates[key] = &genericFuncCandidate{
			decl:  special,
			vals:  vals,
			types: types,
		}
	}
	switch n := len(candidates); n {
	case 1:
		for key, candidate := range candidates {
			if debug {
				g.Debugf("chosen   generic function specialization: %v", key)
			}
			return key, candidate
		}
		fallthrough
	case 0:
		g.Errorf("no generic function specialization matches %v", maker.String())
	default:
		names := make([]string, n)
		var i int
		for name := range candidates {
			names[i] = name
			i++
		}
		sort.Strings(names)
		g.Errorf("multiple candidates match generic function %v:\n\t%s", maker.String(), strings.Join(names, "\n\t"))
	}
	return "", nil
}

// return the most specialized type declaration applicable to used params.
// panics if there is no single most specialized declaration.
func (maker *genericMaker) chooseType(typ *GenericType) (string, *genericTypeCandidate) {
	candidates := map[string]*genericTypeCandidate{
		maker.sym.Name + "#[...]": &genericTypeCandidate{
			decl:  typ.Master,
			vals:  maker.vals,
			types: maker.types,
		},
	}
	g := &maker.comp.Globals
	debug := g.Options&base.OptDebugGenerics != 0
	var ok1, ok2 bool

	if debug {
		g.Debugf("choosing generic type for %s from %d specializations", maker.String(), 1+len(typ.Special))
	}

	for key, special := range typ.Special {
		vals, types, ok := maker.patternMatches(special.Params, special.For, maker.exprs)
		if !ok {
			continue
		}
		// check whether special is more specialized than all other candidates
		for declKey, candidate := range candidates {
			decl := candidate.decl
			if len(decl.For) == 0 {
				ok1, ok2 = false, true
			} else {
				_, _, ok1 = maker.patternMatches(special.Params, special.For, decl.For)
				_, _, ok2 = maker.patternMatches(decl.Params, decl.For, special.For)
			}
			if !ok1 && ok2 {
				// special is more specialized, remove the other
				if debug {
					g.Debugf("generic type %s is more specialized than %s, removing the latter", key, declKey)
				}
				delete(candidates, declKey)
			}
		}
		if debug {
			g.Debugf("adding   generic type specialization  %s to candidates", key)
		}
		candidates[key] = &genericTypeCandidate{
			decl:  special,
			vals:  vals,
			types: types,
		}
	}
	switch n := len(candidates); n {
	case 1:
		for key, candidate := range candidates {
			if debug {
				g.Debugf("chosen   generic type specialization: %v", key)
			}
			return key, candidate
		}
		fallthrough
	case 0:
		g.Errorf("no generic type specialization matches %v", maker.String())
	default:
		names := make([]string, n)
		var i int
		for name := range candidates {
			names[i] = name
			i++
		}
		sort.Strings(names)
		g.Errorf("multiple candidates match generic type %v:\n\t%s", maker.String(), strings.Join(names, "\n\t"))
	}
	return "", nil
}

// if generic specialization 'patterns' parametrized on 'names' matches 'exprs',
// return the constants and types required for the match
func (maker *genericMaker) patternMatches(names []string, patterns []ast.Expr, exprs []ast.Expr) ([]interface{}, []xr.Type, bool) {
	vals := make([]interface{}, len(names))
	types := make([]xr.Type, len(names))
	ok := true

	for i, pattern := range patterns {
		ok = maker.patternMatch(names, vals, types, ast2.ToAst(pattern), ast2.ToAst(exprs[i]))
		if !ok {
			break
		}
	}
	return vals, types, ok
}

// if generic specialization 'pattern1' parametrized on 'names' matches 'expr1',
// fill 'vals' and 'types' with the constants and types required for the match
func (maker *genericMaker) patternMatch(names []string,
	vals []interface{}, types []xr.Type, pattern ast2.Ast, expr ast2.Ast) bool {

	switch node := pattern.Interface().(type) {
	case *ast.Ident:
		for i, name := range names {
			if name == node.Name {
				return maker.patternMatched(i, vals, types, expr)
			}
		}
		e, ok := expr.Interface().(*ast.Ident)
		return ok && node.Name == e.Name
	case *ast.BasicLit:
		e, ok := expr.Interface().(*ast.BasicLit)
		return ok && node.Kind == e.Kind && node.Value == e.Value
	default:
		if pattern.Op() == expr.Op() && pattern.Size() == expr.Size() {
			for i, n := 0, pattern.Size(); i < n; i++ {
				if !maker.patternMatch(names, vals, types, pattern.Get(i), expr.Get(i)) {
					return false
				}
			}
			return true
		}
		return false
	}
}

// if generic specialization 'pattern1' parametrized on 'names' matches 'expr1',
// fill 'vals' and 'types' with the constants and types required for the match
func (maker *genericMaker) patternMatched(i int, vals []interface{}, types []xr.Type, expr ast2.Ast) (ok bool) {
	expr1, eok := expr.Interface().(ast.Expr)
	if !eok {
		return false
	}
	panicking := true
	defer func() {
		if panicking {
			recover()
			ok = false
		}
	}()
	e, typ := maker.comp.Expr1OrType(expr1)
	panicking = false

	if e != nil {
		if e.Const() {
			val := e.EvalConst(COptDefaults)
			if vals[i] == nil {
				vals[i] = val
				ok = true
			} else {
				ok = vals[i] == val
			}
		}
	} else if typ != nil {
		if types[i] == nil {
			types[i] = typ
			ok = true
		} else {
			ok = typ.IdenticalTo(types[i])
		}
	}
	return ok
}
