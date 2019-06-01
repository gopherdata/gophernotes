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
 * generic_func.go
 *
 *  Created on Jun 06, 2018
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"bytes"
	"go/ast"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/output"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// an instantiated (and compiled) generic function.
type GenericFuncInstance struct {
	Func *func(*Env) r.Value
	Type xr.Type
}

// a generic function declaration.
// either general, or partially specialized or fully specialized
type GenericFuncDecl struct {
	Decl   *ast.FuncLit // generic function declaration. use a *ast.FuncLit because we will compile it with Comp.FuncLit()
	Params []string     // generic param names
	For    []ast.Expr   // partial or full specialization
}

// generic function
type GenericFunc struct {
	Master    GenericFuncDecl            // master (i.e. non specialized) declaration
	Special   map[string]GenericFuncDecl // partially or fully specialized declarations. key is GenericFuncDecl.For converted to string
	Instances map[I]*GenericFuncInstance // cache of instantiated functions. key is [N]interface{}{T1, T2...}
}

func (f *GenericFunc) String() string {
	return f.Signature("")
}

func (f *GenericFunc) Signature(name string) string {
	if f == nil {
		return "<nil>"
	}
	var buf bytes.Buffer // strings.Builder requires Go >= 1.10
	decl := f.Master
	if GENERICS_V1_CXX {
		buf.WriteString("template[")
		for i, param := range decl.Params {
			if i != 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(param)
		}
		buf.WriteString("] ")
		if len(name) == 0 {
			(*output.Stringer).Fprintf(nil, &buf, "%v", decl.Decl.Type)
		} else {
			(*output.Stringer).Fprintf(nil, &buf, "%v", &ast.FuncDecl{
				Name: &ast.Ident{Name: name},
				Type: decl.Decl.Type,
			})
		}
		return buf.String()
	}

	buf.WriteString(name)
	buf.WriteString("#[")
	for i, param := range decl.Params {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(param)
	}
	buf.WriteString("] ")
	gname := buf.String()
	buf.Reset()
	(*output.Stringer).Fprintf(nil, &buf, "%v", &ast.FuncDecl{
		Name: &ast.Ident{Name: gname},
		Type: decl.Decl.Type,
	})
	return buf.String()
}

// DeclGenericFunc stores a generic function or method declaration
// for later instantiation
func (c *Comp) DeclGenericFunc(decl *ast.FuncDecl) {
	n := 0
	if decl.Recv != nil {
		n = len(decl.Recv.List)
	}
	if n != 2 {
		c.Errorf("invalid generic function or method declaration: expecting exactly 2 receivers, found %d: %v", n, decl)
	}
	if decl.Recv.List[0] != nil {
		c.Errorf("generic method declaration not yet implemented: %v", decl)
	}
	lit, _ := decl.Recv.List[1].Type.(*ast.CompositeLit)
	if lit == nil {
		c.Errorf("invalid generic function or method declaration: the second receiver should be an *ast.CompositeLit, found %T: %v",
			decl.Recv.List[1].Type, decl)
	}

	params, fors := c.genericParams(lit.Elts, "function or method", decl)

	fdecl := GenericFuncDecl{
		Decl: &ast.FuncLit{
			Type: decl.Type,
			Body: decl.Body,
		},
		Params: params,
		For:    fors,
	}
	name := decl.Name.Name

	if len(fors) == 0 {
		// master (i.e. not specialized) declaration

		if len(params) == 0 {
			c.Errorf("cannot declare generic function with zero generic parameters: %v", decl.Type)
		}
		bind := c.NewBind(name, GenericFuncBind, c.TypeOfPtrGenericFunc())

		// a generic function declaration has no runtime effect:
		// it merely creates the bind for on-demand instantiation by other code
		bind.Value = &GenericFunc{
			Master:    fdecl,
			Special:   make(map[string]GenericFuncDecl),
			Instances: make(map[I]*GenericFuncInstance),
		}
		return
	}

	if !GENERICS_V1_CXX {
		c.Errorf("generic function partial/full specializations are only supported by C++-style generics: %v", decl)
	}
	// partially or fully specialized declaration
	bind := c.Binds[name]
	if bind == nil {
		c.Errorf("undefined identifier: %v", name)
	}
	fun, ok := bind.Value.(*GenericFunc)
	if !ok {
		c.Errorf("symbol is not a generic function, cannot declare function specializations on it: %s // %v", name, bind.Type)
	}
	key := c.Globals.Sprintf("%v", &ast.IndexExpr{X: decl.Name, Index: &ast.CompositeLit{Elts: fors}})
	if len(fun.Master.Params) != len(fors) {
		c.Errorf("generic function specialization for %d parameters, expecting %d: %s", len(fors), len(fun.Master.Params), key)
	}
	if _, ok := fun.Special[key]; ok {
		c.Warnf("redefined generic function specialization: %s", key)
	}
	fun.Special[key] = fdecl
}

// GenericFunc compiles a generic function name#[T1, T2...] instantiating it if needed.
func (c *Comp) GenericFunc(node *ast.IndexExpr) *Expr {
	maker := c.genericMaker(node, GenericFuncBind)
	return c.genericFunc(maker, node)
}

// genericFunc compiles a generic function name#[T1, T2...] instantiating it if needed.
// node is used only for error messages
func (c *Comp) genericFunc(maker *genericMaker, node ast.Node) *Expr {
	if maker == nil {
		return nil
	}
	fun := maker.ifun.(*GenericFunc)
	key := maker.ikey

	instance, _ := fun.Instances[key]
	g := &c.Globals
	debug := g.Options&base.OptDebugGenerics != 0
	if instance != nil {
		if debug {
			g.Debugf("found instantiated generic function %v", maker)
		}
	} else {
		if debug {
			g.Debugf("instantiating generic function %v", maker)
		}
		// hard part: instantiate the generic function.
		// must be instantiated in the same *Comp where it was declared!
		instance = maker.instantiateFunc(fun, node)
	}

	var efun, retfun func(*Env) r.Value
	eaddr := instance.Func
	if *eaddr == nil {
		// currently instantiating it, see comment in Comp.instantiateTemplateFunc() below.
		// We must try again later to dereference instance.Func.
		efun = func(env *Env) r.Value {
			return (*eaddr)(env)
		}
	} else {
		efun = *eaddr
	}
	upn := maker.sym.Upn
	if debug {
		g.Debugf("generic function: %v, upn = %v, instance = %v", maker, upn, instance)
	}
	// switch to the correct *Env before evaluating expr
	switch upn {
	case 0:
		retfun = efun
	case 1:
		retfun = func(env *Env) r.Value {
			return efun(env.Outer)
		}
	case 2:
		retfun = func(env *Env) r.Value {
			return efun(env.Outer.Outer)
		}
	case c.Depth - 1:
		retfun = func(env *Env) r.Value {
			return efun(env.FileEnv)
		}
	case c.Depth:
		retfun = func(env *Env) r.Value {
			return efun(env.FileEnv.Outer)
		}
	default:
		retfun = func(env *Env) r.Value {
			for i := upn; i > 0; i-- {
				env = env.Outer
			}
			return efun(env)
		}
	}
	// always return a new *Expr, in case caller modifies it
	return exprFun(instance.Type, retfun)
}

// instantiateFunc instantiates and compiles a generic function.
// node is used only for error messages
func (maker *genericMaker) instantiateFunc(fun *GenericFunc, node ast.Node) *GenericFuncInstance {

	// choose the specialization to use
	_, special := maker.chooseFunc(fun)

	// create a new nested Comp
	c := NewComp(maker.comp, nil)
	c.UpCost = 0
	c.Depth--

	// and inject generic arguments into it
	special.injectBinds(c)

	key := maker.ikey
	panicking := true
	defer func() {
		if panicking {
			delete(fun.Instances, key)
			c.ErrorAt(node.Pos(), "error instantiating generic function: %v\n\t%v", maker, recover())
		}
	}()

	if c.Globals.Options&base.OptDebugGenerics != 0 {
		c.Debugf("forward-declaring generic function before instantiation: %v", maker)
	}
	// support for generic recursive functions, as for example
	//   template[T] func fib(n T) T { if n <= 2 { return 1 }; return fib#[T](n-1) + fib#[T](n-2) }
	// requires to cache fib#[T] as instantiated **before** actually instantiating it.
	//
	// This is similar to the technique used for non-generic recursive function, as
	//    func fib(n int) int { if n <= 2 { return 1 }; return fib(n-1) + fib(n-2) }
	// with the difference that the cache is fun.Instances[key] instead of Comp.Binds[name]

	// for such trick to work, we must:
	// 1. compute in advance the instantiated function type
	// 2. check GenericFuncInstance.Func: if it's nil, take its address and dereference it later at runtime
	t, _, _ := c.TypeFunction(special.decl.Decl.Type)

	instance := &GenericFuncInstance{Type: t, Func: new(func(*Env) r.Value)}
	fun.Instances[key] = instance

	// compile an expression that, when evaluated at runtime in the *Env
	// where the generic function was declared, returns the instantiated function
	expr := c.FuncLit(special.decl.Decl)

	*instance.Func = expr.AsX1()
	instance.Type = expr.Type

	panicking = false
	return instance
}
