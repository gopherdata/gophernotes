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
 * function.go
 *
 *  Created on Apr 02, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type funcMaker struct {
	nbinds      int
	nintbinds   int
	parambinds  []*Bind
	resultbinds []*Bind
	resultfuns  []I
	funcbody    func(*Env)
}

// DeclFunc compiles a function, macro or method declaration
// For closure declarations, use FuncLit()
func (c *Comp) FuncDecl(funcdecl *ast.FuncDecl) {
	var ismacro bool
	if funcdecl.Recv != nil {
		switch n := len(funcdecl.Recv.List); n {
		case 0:
			ismacro = true
		case 1:
			c.methodDecl(funcdecl)
			return
		default:
			c.Errorf("invalid function/method declaration: found %d receivers, expecting at most one: %v", n, funcdecl)
			return
		}
	}
	functype := funcdecl.Type
	t, paramnames, resultnames := c.TypeFunction(functype)

	// declare the function name and type before compiling its body: allows recursive functions/macros.
	funcname := funcdecl.Name.Name
	oldbind := c.Binds[funcname]
	panicking := true
	defer func() {
		// On compile error, restore pre-existing declaration
		if !panicking || c.Binds == nil {
			// nothing to do
		} else if oldbind != nil {
			c.Binds[funcname] = oldbind
		} else {
			delete(c.Binds, funcname)
		}
	}()
	var funcbind *Bind
	if ismacro {
		// use a ConstBind, as builtins do
		funcbind = c.AddBind(funcname, ConstBind, c.TypeOfMacro())
	} else {
		funcbind = c.AddBind(funcname, FuncBind, t)
	}
	cf := NewComp(c, nil)
	info, resultfuns := cf.funcBinds(functype, t, paramnames, resultnames)
	cf.Func = info

	if body := funcdecl.Body; body != nil {
		// in Go, function arguments/results and function body are in the same scope
		for _, node := range body.List {
			cf.Stmt(node)
		}
	}

	funcindex := funcbind.Desc.Index()
	if funcname == "_" || (!ismacro && funcindex == NoIndex) {
		// function/macro named "_". still compile it (to check for compile errors) but discard the compiled code
		panicking = false
		return
	}
	// do NOT keep a reference to compile environment!
	funcbody := cf.Code.Exec()

	var stmt Stmt
	if ismacro {
		// a macro declaration is a statement:
		// executing it stores the macro function into Comp.Binds[funcname].Value
		f := cf.macroCreate(t, info, resultfuns, funcbody)

		addr := &funcbind.Value
		argnum := t.NumIn()
		stmt = func(env *Env) (Stmt, *Env) {
			fun := f(env)
			*addr = Macro{fun, argnum}
			env.IP++
			return env.Code[env.IP], env
		}
	} else {
		// a function declaration is a statement:
		// executing it creates the function in the runtime environment
		f := cf.funcCreate(t, info, resultfuns, funcbody)

		stmt = func(env *Env) (Stmt, *Env) {
			fun := f(env)
			// Debugf("setting env.Binds[%d] = %v <%v>", funcindex, fun.Interface(), fun.Type())
			env.Binds[funcindex] = fun
			env.IP++
			return env.Code[env.IP], env
		}
	}
	c.Code.Append(stmt, funcdecl.Pos())
	panicking = false
}

func (c *Comp) methodAdd(funcdecl *ast.FuncDecl, t xr.Type) (methodindex int, methods *[]r.Value) {
	name := funcdecl.Name.Name
	trecv := t.In(0)
	if trecv.Kind() == r.Ptr && !trecv.Named() {
		// receiver is an unnamed pointer type. add the method to its element type
		trecv = trecv.Elem()
	}

	panicking := true
	defer func() {
		if panicking {
			rec := recover()
			c.Errorf("error adding method %s <%v> to type <%v>\n\t%v", name, t, trecv, rec)
		}
	}()
	n1 := trecv.NumMethod()
	methodindex = trecv.AddMethod(name, t)
	n2 := trecv.NumMethod()
	if n1 == n2 {
		c.Warnf("redefined method: %s.%s", trecv.Name(), name)
	}
	methods = trecv.GetMethods()
	panicking = false
	return
}

// methodDecl compiles a method declaration
func (c *Comp) methodDecl(funcdecl *ast.FuncDecl) {
	n := len(funcdecl.Recv.List)
	if n != 1 {
		c.Errorf("invalid function/method declaration: expecting one receiver or nil, found %d receivers: func %v %s(/*...*/)",
			n, funcdecl.Recv, funcdecl.Name)
		return
	}
	recvdecl := funcdecl.Recv.List[0]

	functype := funcdecl.Type
	t, paramnames, resultnames := c.TypeFunctionOrMethod(recvdecl, functype)

	// gtype := t.GoType().Underlying().(*types.Signature)
	// c.Debugf("declaring method (%v).%s%s %s\n\treflect.Type: <%v>", gtype.Recv().Type(), funcdecl.Name.Name, gtype.Params(), gtype.Results(), t.ReflectType())

	// declare the method name and type before compiling its body: allows recursive methods
	methodindex, methods := c.methodAdd(funcdecl, t)

	cf := NewComp(c, nil)
	info, resultfuns := cf.funcBinds(functype, t, paramnames, resultnames)
	cf.Func = info

	body := funcdecl.Body
	if body != nil && len(body.List) != 0 {
		// in Go, function arguments/results and function body are in the same scope
		cf.List(body.List)
	}
	// do NOT keep a reference to compile environment!
	funcbody := cf.Code.Exec()
	f := cf.funcCreate(t, info, resultfuns, funcbody)

	// a method declaration is a statement:
	// executing it sets the method value in the receiver type
	var stmt Stmt
	if c.Options&base.OptDebugMethod != 0 {
		trecv := t.In(0)
		tname := trecv.Name()
		if len(tname) == 0 && trecv.Kind() == r.Ptr {
			tname = trecv.Elem().Name()
		}
		methodname := funcdecl.Name
		stmt = func(env *Env) (Stmt, *Env) {
			(*methods)[methodindex] = f(env)
			env.ThreadGlobals.Debugf("implemented method %s.%s", tname, methodname)
			env.IP++
			return env.Code[env.IP], env
		}
	} else {
		stmt = func(env *Env) (Stmt, *Env) {
			(*methods)[methodindex] = f(env)
			env.IP++
			return env.Code[env.IP], env
		}
	}
	c.Code.Append(stmt, funcdecl.Pos())
}

// FuncLit compiles a function literal, i.e. a closure.
// For functions or methods declarations, use FuncDecl()
func (c *Comp) FuncLit(funclit *ast.FuncLit) *Expr {
	functype := funclit.Type
	t, paramnames, resultnames := c.TypeFunction(functype)

	cf := NewComp(c, nil)
	info, resultfuns := cf.funcBinds(functype, t, paramnames, resultnames)
	cf.Func = info

	body := funclit.Body
	if body != nil && len(body.List) != 0 {
		// in Go, function arguments/results and function body are in the same scope
		cf.List(body.List)
	}
	// do NOT keep a reference to compile environment!
	funcbody := cf.Code.Exec()

	f := cf.funcCreate(t, info, resultfuns, funcbody)

	// a function literal is an expression:
	// executing it returns the function
	return exprX1(t, f)
}

// prepare the function parameter binds, result binds and FuncInfo
func (c *Comp) funcBinds(functype *ast.FuncType, t xr.Type, paramnames, resultnames []string) (info *FuncInfo, resultfuns []I) {

	parambinds := c.funcParamBinds(functype, t, paramnames)

	resultbinds, resultfuns := c.funcResultBinds(functype, t, resultnames)
	namedresults := true
	for _, resultname := range resultnames {
		if len(resultname) == 0 {
			namedresults = false
		}
	}
	return &FuncInfo{
		Params:       parambinds,
		Results:      resultbinds,
		NamedResults: namedresults,
	}, resultfuns
}

// prepare the function parameter binds
func (c *Comp) funcParamBinds(functype *ast.FuncType, t xr.Type, names []string) []*Bind {
	nin := t.NumIn()
	binds := make([]*Bind, nin)
	var namedparams, unnamedparams bool
	ismethod := t.IsMethod()
	for i := 0; i < nin; i++ {
		// names[i] == "" means that argument is unnamed, and thus ignored inside the function.
		// change to "_" so that AddBind will not allocate a bind for it - correct optimization...
		// just remember to check for such case when creating the function
		name := names[i]
		if !ismethod || i != 0 {
			// method receiver can be named or unnamed, independently from other input parameters
			if name == "" {
				name = "_"
				unnamedparams = true
			} else {
				namedparams = true
			}
		}
		if namedparams && unnamedparams {
			c.Errorf("cannot mix named and unnamed parameters in function declaration: %v", functype)
		}
		bind := c.AddBind(name, VarBind, t.In(i))
		binds[i] = bind
	}
	return binds
}

// prepare the function result binds
func (c *Comp) funcResultBinds(functype *ast.FuncType, t xr.Type, names []string) (binds []*Bind, funs []I) {
	n := t.NumOut()
	binds = make([]*Bind, n)
	funs = make([]I, n)
	var namedresults, unnamedresults bool
	for i, n := 0, t.NumOut(); i < n; i++ {
		// names[i] == "" means that result is unnamed.
		// we must still allocate a bind for it.
		name := names[i]
		if name == "" {
			unnamedresults = true
		} else {
			namedresults = true
		}
		if namedresults && unnamedresults {
			c.Errorf("cannot mix named and unnamed results in function declaration: %v", functype)
		}
		c.Pos = functype.Pos()
		bind := c.DeclVar0(name, t.Out(i), nil)
		binds[i] = bind
		// compile the extraction of results from runtime env
		sym := bind.AsSymbol(0)
		funs[i] = c.Symbol(sym).WithFun()
	}
	return
}

func (c *Comp) funcMaker(info *FuncInfo, resultfuns []I, funcbody func(*Env)) *funcMaker {
	return &funcMaker{
		nbinds:      c.BindNum,
		nintbinds:   c.IntBindNum,
		parambinds:  info.Params,
		resultbinds: info.Results,
		resultfuns:  resultfuns,
		funcbody:    funcbody,
	}
}

// actually create the function
func (c *Comp) funcCreate(t xr.Type, info *FuncInfo, resultfuns []I, funcbody func(*Env)) func(*Env) r.Value {
	c.ErrorIfCompiled(t)

	m := c.funcMaker(info, resultfuns, funcbody)

	rtype := t.ReflectType() // has receiver as first parameter
	nin := rtype.NumIn()
	nout := rtype.NumOut()

	// do not create optimized functions if arguments or results are named types
	optimize := true
	for i := 0; optimize && i < nin; i++ {
		rt := rtype.In(i)
		k := rt.Kind()
		optimize = base.IsOptimizedKind(k) && rt == c.Universe.BasicTypes[k].ReflectType()
	}
	for i := 0; optimize && i < nout; i++ {
		rt := rtype.Out(i)
		k := rt.Kind()
		optimize = base.IsOptimizedKind(k) && rt == c.Universe.BasicTypes[k].ReflectType()
	}

	var fun func(*Env) r.Value
	if optimize {
		switch nin {
		case 0:
			switch nout {
			case 0:
				fun = c.func0ret0(t, m)
			case 1:
				fun = c.func0ret1(t, m)
			}
		case 1:
			switch nout {
			case 0:
				fun = c.func1ret0(t, m)
			case 1:
				fun = c.func1ret1(t, m)
			}
		case 2:
			switch nout {
			case 0:
				fun = c.func2ret0(t, m)
			}
		}
	}
	if fun == nil {
		fun = c.funcGeneric(t, m)
	}
	return fun
}

// fallback: create a non-optimized function
func (c *Comp) funcGeneric(t xr.Type, m *funcMaker) func(*Env) r.Value {

	paramdecls := make([]func(*Env, r.Value), len(m.parambinds))
	for i, bind := range m.parambinds {
		if bind.Desc.Index() != NoIndex {
			paramdecls[i] = c.DeclBindRuntimeValue(bind)
		}
	}
	resultexprs := make([]func(*Env) r.Value, len(m.resultfuns))
	for i, resultfun := range m.resultfuns {
		resultexprs[i] = funAsX1(resultfun, m.resultbinds[i].Type)
	}

	// do NOT keep a reference to funcMaker
	nbinds := m.nbinds
	nintbinds := m.nintbinds
	funcbody := m.funcbody
	rtype := t.ReflectType()

	return func(env *Env) r.Value {
		// function is closed over the env used to DECLARE it
		env.MarkUsedByClosure()
		return r.MakeFunc(rtype, func(args []r.Value) []r.Value {
			env := NewEnv4Func(env, nbinds, nintbinds)

			if funcbody != nil {
				// copy runtime arguments into allocated binds
				for i, decl := range paramdecls {
					if decl != nil {
						// decl == nil means the argument is ignored inside the function
						decl(env, args[i])
					}
				}
				// execute the body
				funcbody(env)
			}
			// read results from allocated binds and return them
			rets := make([]r.Value, len(resultexprs))
			for i, expr := range resultexprs {
				rets[i] = expr(env)
			}
			env.FreeEnv()
			return rets
		})
	}
}

// create a macro
func (c *Comp) macroCreate(t xr.Type, info *FuncInfo, resultfuns []I, funcbody func(*Env)) func(*Env) func(args []r.Value) []r.Value {
	m := c.funcMaker(info, resultfuns, funcbody)

	paramdecls := make([]func(*Env, r.Value), len(m.parambinds))
	for i, bind := range m.parambinds {
		if bind.Desc.Index() != NoIndex {
			paramdecls[i] = c.DeclBindRuntimeValue(bind)
		}
	}
	resultexprs := make([]func(*Env) r.Value, len(m.resultfuns))
	for i, resultfun := range m.resultfuns {
		resultexprs[i] = funAsX1(resultfun, m.resultbinds[i].Type)
	}

	// do NOT keep a reference to funcMaker
	nbinds := m.nbinds
	nintbinds := m.nintbinds

	return func(env *Env) func(args []r.Value) []r.Value {
		// macro is closed over the env used to DECLARE it
		env.MarkUsedByClosure()
		return func(args []r.Value) []r.Value {
			env := NewEnv4Func(env, nbinds, nintbinds)

			if funcbody != nil {
				// copy runtime arguments into allocated binds
				for i, decl := range paramdecls {
					if decl != nil {
						// decl == nil means the argument is ignored inside the function
						decl(env, args[i])
					}
				}
				// execute the body
				funcbody(env)
			}
			// read results from allocated binds and return them
			rets := make([]r.Value, len(resultexprs))
			for i, expr := range resultexprs {
				rets[i] = expr(env)
			}
			env.FreeEnv()
			return rets
		}
	}
}

func declBindRuntimeValueNop(*Env, r.Value) {
}
