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
 * builtin.go
 *
 *  Created on: Apr 02, 2017
 *      Author: Massimiliano Ghilardi
 */

package fast

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"io"
	r "reflect"

	"github.com/cosmos72/gomacro/ast2"
	"github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/untyped"
	xr "github.com/cosmos72/gomacro/xreflect"
)

var (
	zeroTypes          = []xr.Type{}
	rtypeOfSliceOfByte = r.TypeOf([]byte{})
)

// =================================== iota ===================================

// returns the previous definition of iota - to be restored by Comp.endIota() below
func (top *Comp) beginIota() *Bind {
	return top.Binds["iota"]
}

func (top *Comp) endIota(orig *Bind) {
	if orig == nil {
		delete(top.Binds, "iota")
	} else {
		top.Binds["iota"] = orig
	}
}

func (top *Comp) setIota(iota int) {
	// https://golang.org/ref/spec#Constants
	// "Literal constants, true, false, iota, and certain constant expressions containing only untyped constant operands are untyped."

	// Binds are supposed to be immutable. to avoid issues, create a new Bind every time
	top.Binds["iota"] = top.BindUntyped(r.Int, constant.MakeInt64(int64(iota)))
}

// ============================== initialization ===============================

func (ce *Interp) addBuiltins() {
	basicTypes := &ce.Comp.Universe.BasicTypes

	// https://golang.org/ref/spec#Constants
	// "Literal constants, true, false, iota, and certain constant expressions containing only untyped constant operands are untyped."
	ce.DeclConst("false", nil, MakeUntypedLit(r.Bool, constant.MakeBool(false), basicTypes))
	ce.DeclConst("true", nil, MakeUntypedLit(r.Bool, constant.MakeBool(true), basicTypes))

	// https://golang.org/ref/spec#Variables : "[...] the predeclared identifier nil, which has no type"
	ce.DeclConst("nil", nil, nil)

	ce.DeclBuiltin("append", Builtin{compileAppend, 1, base.MaxUint16})
	ce.DeclBuiltin("cap", Builtin{compileCap, 1, 1})
	ce.DeclBuiltin("close", Builtin{compileClose, 1, 1})
	ce.DeclBuiltin("copy", Builtin{compileCopy, 2, 2})
	ce.DeclBuiltin("complex", Builtin{compileComplex, 2, 2})
	ce.DeclBuiltin("delete", Builtin{compileDelete, 2, 2})
	ce.DeclBuiltin("imag", Builtin{compileRealImag, 1, 1})
	ce.DeclBuiltin("len", Builtin{compileLen, 1, 1})
	ce.DeclBuiltin("make", Builtin{compileMake, 1, 3})
	ce.DeclBuiltin("new", Builtin{compileNew, 1, 1})
	ce.DeclBuiltin("panic", Builtin{compilePanic, 1, 1})
	ce.DeclBuiltin("print", Builtin{compilePrint, 0, base.MaxUint16})
	ce.DeclBuiltin("println", Builtin{compilePrint, 0, base.MaxUint16})
	ce.DeclBuiltin("real", Builtin{compileRealImag, 1, 1})
	ce.DeclBuiltin("recover", Builtin{compileRecover, 0, 0})
	// ce.DeclBuiltin("recover", Function{callRecover, ce.Comp.TypeOf((*func() interface{})(nil)).Elem()})

	tfunI2_Nb := ce.Comp.TypeOf(funI2_Nb)

	ce.DeclEnvFunc("Interp", Function{callIdentity, ce.Comp.TypeOf(funI_I)})
	ce.DeclEnvFunc("Eval", Function{callEval, ce.Comp.TypeOf(funI2_I)})
	ce.DeclEnvFunc("EvalKeepUntyped", Function{callEvalKeepUntyped, ce.Comp.TypeOf(funI2_I)})
	ce.DeclEnvFunc("EvalType", Function{callEvalType, ce.Comp.TypeOf(funI2_T)})
	ce.DeclEnvFunc("MacroExpand", Function{callMacroExpand, tfunI2_Nb})
	ce.DeclEnvFunc("MacroExpand1", Function{callMacroExpand1, tfunI2_Nb})
	ce.DeclEnvFunc("MacroExpandCodeWalk", Function{callMacroExpandCodeWalk, tfunI2_Nb})
	ce.DeclEnvFunc("Parse", Function{callParse, ce.Comp.TypeOf(funSI_I)})
	/*
		binds["Read"] = r.ValueOf(ReadString)
		binds["ReadDir"] = r.ValueOf(callReadDir)
		binds["ReadFile"] = r.ValueOf(callReadFile)
		binds["ReadMultiline"] = r.ValueOf(ReadMultiline)
		binds["Slice"] = r.ValueOf(callSlice)
		binds["String"] = r.ValueOf(func(args ...interface{}) string {
			return env.toString("", args...)
		})
		// return multiple values, extracting the concrete type of each interface
		binds["Values"] = r.ValueOf(Function{funcValues, -1})
	*/

	// --------- types ---------
	c := ce.Comp
	for _, t := range c.Universe.BasicTypes {
		ce.DeclType(t)
	}
	ce.DeclTypeAlias("byte", c.TypeOfUint8())
	ce.DeclTypeAlias("rune", c.TypeOfInt32())
	ce.DeclType(c.TypeOfError())

	/*
		// --------- proxies ---------
		if env.Proxies == nil {
			env.Proxies = make(map[string]Type)
		}
		proxies := env.Proxies

		proxies["error", TypeOf(*Error_builtin)(nil)).Elem()
	*/
}

// ============================= builtin functions =============================

// --- append() ---

func compileAppend(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	n := len(node.Args)
	args := make([]*Expr, n)

	args[0] = c.Expr1(node.Args[0], nil)
	t0 := args[0].Type
	if t0.Kind() != r.Slice {
		c.Errorf("first argument to %s must be slice; have <%s>", sym.Name, t0)
		return nil
	}
	telem := t0.Elem()
	t1 := c.Universe.SliceOf(telem)

	if node.Ellipsis != token.NoPos {
		if n != 2 {
			return c.badBuiltinCallArgNum(sym.Name+"(arg1, arg2...)", 2, 2, node.Args)
		}
		telem = t1 // second argument is a slice too
	}
	for i := 1; i < n; i++ {
		argi := c.Expr1(node.Args[i], nil)
		if argi.Const() {
			argi.ConstTo(telem)
		} else if ti := argi.Type; ti == nil || !ti.AssignableTo(telem) {
			return c.badBuiltinCallArgType(sym.Name, node.Args[i], ti, telem)
		}
		args[i] = argi
	}
	t := c.Universe.FuncOf([]xr.Type{t0, t1}, []xr.Type{t0}, true) // compile as reflect.Append(), which is variadic
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: r.Append}, &sym)
	return &Call{
		Fun:      fun,
		Args:     args,
		OutTypes: []xr.Type{t0},
		Const:    false,
		Ellipsis: node.Ellipsis != token.NoPos,
	}
}

// --- cap() ---

func callCap(val r.Value) int {
	return val.Cap()
}

func compileCap(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	// argument of builtin cap() cannot be a literal
	arg := c.Expr1(node.Args[0], nil)
	tin := arg.Type
	tout := c.TypeOfInt()
	switch tin.Kind() {
	// no cap() on r.Map, see
	// https://golang.org/ref/spec#Length_and_capacity
	// and https://golang.org/pkg/reflect/#Value.Cap
	case r.Array, r.Chan, r.Slice:
		// ok
	case r.Ptr:
		if tin.Elem().Kind() == r.Array {
			// cap() on pointer to array
			arg = c.Deref(arg)
			tin = arg.Type
			break
		}
		fallthrough
	default:
		return c.badBuiltinCallArgType(sym.Name, node.Args[0], tin, "array, channel, slice, pointer to array")
	}
	t := c.Universe.FuncOf([]xr.Type{tin}, []xr.Type{tout}, false)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: callCap}, &sym)
	// capacity of arrays is part of their type: cannot change at runtime, we could optimize it.
	// TODO https://golang.org/ref/spec#Length_and_capacity specifies
	// when the array passed to cap() is evaluated and when is not...
	return newCall1(fun, arg, arg.Const(), tout)
}

// --- close() ---

func callClose(val r.Value) {
	val.Close()
}

func compileClose(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	arg := c.Expr1(node.Args[0], nil)
	tin := arg.Type
	if tin.Kind() != r.Chan {
		return c.badBuiltinCallArgType(sym.Name, node.Args[0], tin, "channel")
	}
	t := c.Universe.FuncOf([]xr.Type{tin}, zeroTypes, false)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: callClose}, &sym)
	return newCall1(fun, arg, false)
}

// --- complex() ---

func callComplex64(re float32, im float32) complex64 {
	return complex(re, im)
}

func callComplex128(re float64, im float64) complex128 {
	return complex(re, im)
}

func compileComplex(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	re := c.Expr1(node.Args[0], nil)
	im := c.Expr1(node.Args[1], nil)
	if re.Untyped() {
		if im.Untyped() {
			return compileComplexUntyped(c, sym, node, re.Value.(UntypedLit), im.Value.(UntypedLit))
		} else {
			re.ConstTo(im.Type)
		}
	} else if im.Untyped() {
		im.ConstTo(re.Type)
	}
	c.toSameFuncType(node, re, im)
	kre := base.KindToCategory(re.Type.Kind())
	if re.Const() && kre != r.Float64 {
		re.ConstTo(c.TypeOfFloat64())
		kre = r.Float64
	}
	kim := base.KindToCategory(im.Type.Kind())
	if im.Const() && kim != r.Float64 {
		im.ConstTo(c.TypeOfFloat64())
		kim = r.Float64
	}
	if kre != r.Float64 {
		c.Errorf("invalid operation: %v (arguments have type %v, expected integer or floating-point)",
			node, re.Type)
	}
	if kim != r.Float64 {
		c.Errorf("invalid operation: %v (arguments have type %v, expected integer or floating-point)",
			node, im.Type)
	}
	tin := re.Type
	k := re.Type.Kind()
	var tout xr.Type
	var call I
	switch k {
	case r.Float32:
		tout = c.TypeOfComplex64()
		call = callComplex64
	case r.Float64:
		tout = c.TypeOfComplex128()
		call = callComplex128
	default:
		return c.badBuiltinCallArgType(sym.Name, node.Args[0], tin, "floating point")
	}
	touts := []xr.Type{tout}
	tfun := c.Universe.FuncOf([]xr.Type{tin}, touts, false)
	sym.Type = tfun
	fun := exprLit(Lit{Type: tfun, Value: call}, &sym)
	// complex() of two constants is constant: it can be computed at compile time
	return &Call{Fun: fun, Args: []*Expr{re, im}, OutTypes: touts, Const: re.Const() && im.Const()}
}

var complexImagOne = constant.MakeFromLiteral("1i", token.IMAG, 0)

func compileComplexUntyped(c *Comp, sym Symbol, node *ast.CallExpr, re UntypedLit, im UntypedLit) *Call {
	checkComplexUntypedArg(c, node, re, "first")
	checkComplexUntypedArg(c, node, im, "second")
	rev := re.Val
	imv := constant.BinaryOp(im.Val, token.MUL, complexImagOne)
	val := MakeUntypedLit(r.Complex128, constant.BinaryOp(rev, token.ADD, imv), &c.Universe.BasicTypes)
	touts := []xr.Type{c.TypeOfUntypedLit()}
	tfun := c.Universe.FuncOf(nil, touts, false)
	sym.Type = tfun
	fun := exprLit(Lit{Type: tfun, Value: val}, &sym)
	// complex() of two untyped constants is both untyped and constant: it can be computed at compile time
	return &Call{Fun: fun, Args: nil, OutTypes: touts, Const: true}
}

func checkComplexUntypedArg(c *Comp, node *ast.CallExpr, arg UntypedLit, label string) {
	switch arg.Kind {
	case r.Int, r.Int32 /*rune*/, r.Float64:
		return
	case r.Complex128:
		im := constant.Imag(arg.Val)
		switch im.Kind() {
		case constant.Int:
			if x, exact := constant.Int64Val(im); x == 0 && exact {
				return
			}
		case constant.Float:
			if x, exact := constant.Float64Val(im); x == 0.0 && exact {
				return
			}
		}
	}
	c.Errorf("invalid operation: %v (first argument is untyped %v, expected untyped integer, untyped float, or untyped complex with zero imaginary part)",
		node, arg)
}

// --- copy() ---

func copyStringToBytes(dst []byte, src string) int {
	// reflect.Copy does not support this case... use the compiler support
	return copy(dst, src)
}

func compileCopy(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	args := []*Expr{
		c.Expr1(node.Args[0], nil),
		c.Expr1(node.Args[1], nil),
	}
	if args[1].Const() {
		// we also accept a string literal as second argument
		args[1].ConstTo(args[1].DefaultType())
	}
	t0, t1 := args[0].Type, args[1].Type
	var funCopy I = r.Copy
	if t0 == nil || t0.Kind() != r.Slice || !t0.AssignableTo(c.Universe.SliceOf(t0.Elem())) {
		// https://golang.org/ref/spec#Appending_and_copying_slices
		// copy [...] arguments must have identical element type T and must be assignable to a slice of type []T.
		c.Errorf("first argument to copy should be slice; have %v <%v>", node.Args[0], t0)
		return nil
	} else if t0.Elem().Kind() == r.Uint8 && t1.Kind() == r.String {
		// [...] As a special case, copy also accepts a destination argument assignable to type []byte
		// with a source argument of a string type. This form copies the bytes from the string into the byte slice.
		funCopy = copyStringToBytes
	} else if t1 == nil || t1.Kind() != r.Slice || !t1.AssignableTo(c.Universe.SliceOf(t1.Elem())) {
		c.Errorf("second argument to copy should be slice or string; have %v <%v>", node.Args[1], t1)
		return nil
	} else if !t0.Elem().IdenticalTo(t1.Elem()) {
		c.Errorf("arguments to copy have different element types: <%v> and <%v>", t0.Elem(), t1.Elem())
	}
	outtypes := []xr.Type{c.TypeOfInt()}
	t := c.Universe.FuncOf([]xr.Type{t0, t1}, outtypes, false)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: funCopy}, &sym)
	return &Call{Fun: fun, Args: args, OutTypes: outtypes, Const: false}
}

// --- delete() ---

// use whatever calling convention is convenient: reflect.Values, interface{}s, primitive types...
// as long as call_builtin supports it, we're fine
func callDelete(vmap r.Value, vkey r.Value) {
	vmap.SetMapIndex(vkey, base.Nil)
}

func compileDelete(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	emap := c.Expr1(node.Args[0], nil)
	ekey := c.Expr1(node.Args[1], nil)
	tmap := emap.Type
	if tmap.Kind() != r.Map {
		c.Errorf("first argument to delete must be map; have %v", tmap)
		return nil
	}
	tkey := tmap.Key()
	if ekey.Const() {
		ekey.ConstTo(tkey)
	} else if ekey.Type == nil || !ekey.Type.AssignableTo(tkey) {
		c.Errorf("cannot use %v <%v> as type <%v> in delete", node.Args[1], ekey.Type, tkey)
	}
	t := c.Universe.FuncOf([]xr.Type{tmap, tkey}, zeroTypes, false)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: callDelete}, &sym)
	return &Call{Fun: fun, Args: []*Expr{emap, ekey}, OutTypes: zeroTypes, Const: false}
}

// --- Env() ---

func funI_I(interface{}) interface{} {
	return nil
}

// we can use whatever signature we want, as long as call_builtin supports it
func callIdentity(v r.Value) r.Value {
	return v
}

// --- Eval() ---

func funI2_I(interface{}, interface{}) interface{} {
	return nil
}

func callEval(argv r.Value, interpv r.Value) r.Value {
	// always convert untyped constants to their default type.
	// To retrieve untyped constants, use EvalKeepUntyped()
	return callEval3(argv, interpv, COptDefaults)
}

func callEvalKeepUntyped(argv r.Value, interpv r.Value) r.Value {
	return callEval3(argv, interpv, COptKeepUntyped)
}

func callEval3(argv r.Value, interpv r.Value, opt CompileOptions) r.Value {
	if !argv.IsValid() {
		return argv
	}
	form := anyToAst(argv.Interface(), "Eval")
	form = base.SimplifyAstForQuote(form, true)

	ir := interpv.Interface().(*Interp)

	// use Comp.Compile(), which always compiles, instead of Interp.CompileAst():
	// the latter compiles only if option MacroExpandOnly is unset
	e := ir.Comp.Compile(form)

	if e == nil {
		return base.None
	}
	e.CheckX1()

	if opt&COptKeepUntyped == 0 && e.Untyped() {
		e.ConstTo(e.DefaultType())
	}

	// do not use Interp.RunExpr() or Interp.RunExpr1()
	// because they convert untyped constants to their default type
	// if Interp.Comp.Globals.Options&OptKeepUntyped == 0
	env := ir.PrepareEnv()

	fun := e.AsXV(COptKeepUntyped)
	v, _ := fun(env)
	return v
}

// --- EvalType() ---

func funI2_T(interface{}, interface{}) r.Type {
	return nil
}

func callEvalType(argv r.Value, interpv r.Value) r.Value {
	if !argv.IsValid() {
		return zeroOfReflectType
	}
	form := anyToAst(argv.Interface(), "EvalType")
	form = base.UnwrapTrivialAst(form)
	node := form.Interface().(ast.Expr)

	interp := interpv.Interface().(*Interp)
	t := interp.Comp.compileTypeOrNil(node)
	if t == nil {
		return zeroOfReflectType
	}
	return r.ValueOf(t.ReflectType())
}

// --- len() ---

func callLenValue(val r.Value) int {
	return val.Len()
}

func callLenString(val string) int {
	return len(val)
}

func compileLen(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	arg := c.Expr1(node.Args[0], nil)
	if arg.Const() {
		arg.ConstTo(arg.DefaultType())
	}
	tin := arg.Type
	tout := c.TypeOfInt()
	switch tin.Kind() {
	case r.Array, r.Chan, r.Map, r.Slice, r.String:
		// ok
	case r.Ptr:
		if tin.Elem().Kind() == r.Array {
			// len() on pointer to array
			arg = c.Deref(arg)
			tin = arg.Type
			break
		}
		fallthrough
	default:
		return c.badBuiltinCallArgType(sym.Name, node.Args[0], tin, "array, channel, map, slice, string, pointer to array")
	}
	t := c.Universe.FuncOf([]xr.Type{tin}, []xr.Type{tout}, false)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: callLenValue}, &sym)
	if tin.Kind() == r.String {
		fun.Value = callLenString // optimization
	}
	// length of arrays is part of their type: cannot change at runtime,
	// so perform constant propagation on it.
	// TODO https://golang.org/ref/spec#Length_and_capacity specifies
	// when the array passed to len() is evaluated and when is not...
	isarray := tin.Kind() == r.Array
	if isarray {
		n := tin.Len()
		fun.Value = func(_ r.Value) int {
			return n
		}
		// since we currently optimize len() by evaluating it at compile time,
		// actual arg may not exist yet. optimize it away.
		arg = exprLit(Lit{Type: tin, Value: xr.Zero(tin).Interface()}, nil)
	}
	return newCall1(fun, arg, isarray || arg.Const(), tout)
}

// --- MacroExpand(), MacroExpand1(), MacroExpandCodeWalk() ---

func funI2_Nb(interface{}, interface{}) (ast.Node, bool) {
	return nil, false
}

func callMacroExpand(argv r.Value, interpv r.Value) (r.Value, r.Value) {
	return callMacroExpandDispatch(argv, interpv, "MacroExpand")
}

func callMacroExpand1(argv r.Value, interpv r.Value) (r.Value, r.Value) {
	return callMacroExpandDispatch(argv, interpv, "MacroExpand1")
}

func callMacroExpandCodeWalk(argv r.Value, interpv r.Value) (r.Value, r.Value) {
	return callMacroExpandDispatch(argv, interpv, "MacroExpandCodeWalk")
}

func callMacroExpandDispatch(argv r.Value, interpv r.Value, caller string) (r.Value, r.Value) {
	if !argv.IsValid() {
		return r.Zero(rtypeOfNode), base.False
	}
	form := anyToAst(argv.Interface(), caller)
	form = base.SimplifyAstForQuote(form, true)

	interp := interpv.Interface().(*Interp)
	c := interp.Comp

	var flag bool
	switch caller {
	default:
		form, flag = c.MacroExpand(form)
	case "MacroExpand1":
		form, flag = c.MacroExpand1(form)
	case "MacroExpandCodeWalk":
		form, flag = c.MacroExpandCodewalk(form)
	}
	flagv := base.False
	if flag {
		flagv = base.True
	}
	return r.ValueOf(form.Interface()).Convert(rtypeOfNode), flagv
}

// --- make() ---

func makeChan1(t r.Type) r.Value {
	return r.MakeChan(t, 0)
}

func makeMap2(t r.Type, n int) r.Value {
	// reflect.MakeMap cannot specify initial capacity
	return r.MakeMap(t)
}

func makeSlice2(t r.Type, n int) r.Value {
	// reflect.MakeSlice requires capacity
	return r.MakeSlice(t, n, n)
}

func compileMake(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	nargs := len(node.Args)
	var nmin, nmax uint16 = 1, 2
	tin := c.Type(node.Args[0])
	var funMakes [4]I
	switch tin.Kind() {
	case r.Chan:
		funMakes[1] = makeChan1
		funMakes[2] = r.MakeChan
	case r.Map:
		funMakes[1] = r.MakeMap
		funMakes[2] = makeMap2
	case r.Slice:
		nmin, nmax = 2, 3
		funMakes[2] = makeSlice2
		funMakes[3] = r.MakeSlice
	default:
		return c.badBuiltinCallArgType(sym.Name, node.Args[0], tin, "channel, map, slice")
	}
	if nargs < int(nmin) || nargs > int(nmax) {
		return c.badBuiltinCallArgNum(sym.Name+"()", nmin, nmax, node.Args)
	}
	args := make([]*Expr, nargs)
	argtypes := make([]xr.Type, nargs)
	argtypes[0] = c.TypeOfInterface()
	args[0] = c.exprValue(argtypes[0], tin.ReflectType()) // no need to build TypeOfReflectType
	te := c.TypeOfInt()
	for i := 1; i < nargs; i++ {
		argi := c.Expr1(node.Args[i], nil)
		if argi.Const() {
			argi.ConstTo(te)
		} else if ti := argi.Type; ti == nil || (!ti.IdenticalTo(te) && !ti.AssignableTo(te)) {
			return c.badBuiltinCallArgType(sym.Name, node.Args[i], ti, te)
		}
		args[i] = argi
		argtypes[i] = te
	}
	outtypes := []xr.Type{tin}
	t := c.Universe.FuncOf(argtypes, outtypes, false)
	sym.Type = t
	funMake := funMakes[nargs]
	if funMake == nil {
		c.Errorf("internal error: no make() alternative to call for %v with %d arguments", tin, nargs)
		return nil
	}
	fun := exprLit(Lit{Type: t, Value: funMake}, &sym)
	return &Call{Fun: fun, Args: args, OutTypes: outtypes, Const: false}
}

// --- new() ---

func compileNew(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	tin := c.Type(node.Args[0])
	tout := c.Universe.PtrTo(tin)
	t := c.Universe.FuncOf([]xr.Type{c.TypeOfInterface()}, []xr.Type{tout}, false) // no need to build TypeOfReflectType
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: r.New}, &sym)
	arg := c.exprValue(c.TypeOfInterface(), tin.ReflectType())
	return newCall1(fun, arg, false, tout)
}

// --- panic() ---

func callPanic(arg interface{}) {
	panic(arg)
}

func compilePanic(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	arg := c.Expr1(node.Args[0], nil)
	arg.To(c, c.TypeOfInterface())
	t := c.TypeOf(callPanic)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: callPanic}, &sym)
	return newCall1(fun, arg, false)
}

// --- Parse() ---

func funSI_I(string, interface{}) interface{} {
	return nil
}

func callParse(argv r.Value, interpv r.Value) r.Value {
	if !argv.IsValid() {
		return argv
	}
	ir := interpv.Interface().(*Interp)

	if argv.Kind() == r.Interface {
		argv = argv.Elem()
	}
	if argv.Kind() != r.String {
		ir.Comp.Errorf("cannot convert %v to string: %v", argv.Type(), argv)
	}

	form := ir.Comp.Parse(argv.String())
	return r.ValueOf(&form).Elem() // always return type ast2.Ast
}

// --- print(), println() ---

func callPrint(out interface{}, args ...interface{}) {
	fmt.Fprint(out.(io.Writer), args...)
}

func callPrintln(out interface{}, args ...interface{}) {
	fmt.Fprintln(out.(io.Writer), args...)
}

func getStdout(env *Env) r.Value {
	return r.ValueOf(env.Run.Stdout)
}

func compilePrint(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	args := c.Exprs(node.Args)
	for _, arg := range args {
		arg.To(c, c.TypeOfInterface())
	}
	arg0 := exprFun(c.TypeOfInterface(), getStdout) // no need to build TypeOfIoWriter
	args = append([]*Expr{arg0}, args...)

	t := c.TypeOf(callPrint)
	sym.Type = t
	call := callPrint
	if sym.Name == "println" {
		call = callPrintln
	}
	fun := exprLit(Lit{Type: t, Value: call}, &sym)
	return &Call{Fun: fun, Args: args, OutTypes: zeroTypes, Const: false, Ellipsis: node.Ellipsis != token.NoPos}
}

// --- real() and imag() ---

func callReal32(val complex64) float32 {
	return real(val)
}

func callReal64(val complex128) float64 {
	return real(val)
}

func callImag32(val complex64) float32 {
	return imag(val)
}

func callImag64(val complex128) float64 {
	return imag(val)
}

func compileRealImag(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	arg := c.Expr1(node.Args[0], nil)
	if arg.Const() {
		if arg.Untyped() {
			return compileRealImagUntyped(c, sym, node, arg.Value.(UntypedLit))
		}
		arg.ConstTo(arg.DefaultType())
	}
	tin := arg.Type
	var tout xr.Type
	var call I
	switch tin.Kind() {
	case r.Complex64:
		tout = c.TypeOfFloat32()
		if sym.Name == "real" {
			call = callReal32
		} else {
			call = callImag32
		}
	case r.Complex128:
		tout = c.TypeOfFloat64()
		if sym.Name == "real" {
			call = callReal64
		} else {
			call = callImag64
		}
	default:
		return c.badBuiltinCallArgType(sym.Name, node.Args[0], tin, "complex")
	}
	t := c.Universe.FuncOf([]xr.Type{tin}, []xr.Type{tout}, false)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: call}, &sym)
	// real() and imag() of a constant are constants: they can be computed at compile time
	return newCall1(fun, arg, arg.Const(), tout)
}

func compileRealImagUntyped(c *Comp, sym Symbol, node *ast.CallExpr, arg UntypedLit) *Call {
	val := arg.Val
	if sym.Name == "real" {
		val = constant.Real(val)
	} else {
		val = constant.Imag(val)
	}
	// convert constant.Value result to UntypedLit of appropriate kind
	kind := untyped.ConstantKindToUntypedLitKind(val.Kind())
	arg = MakeUntypedLit(kind, val, &c.Universe.BasicTypes)

	touts := []xr.Type{c.TypeOfUntypedLit()}
	tfun := c.Universe.FuncOf(nil, touts, false)
	sym.Type = tfun

	fun := exprLit(Lit{Type: tfun, Value: arg}, &sym)
	// real() and imag() of untyped constant is both untyped and constant: it can be computed at compile time
	return &Call{Fun: fun, Args: nil, OutTypes: touts, Const: true}
}

var nilInterface = r.Zero(base.TypeOfInterface)

// we can use whatever signature we want, as long as call_builtin supports it
func callRecover(v r.Value) r.Value {
	env := v.Interface().(*Env)
	g := env.Run
	debug := g.Options&base.OptDebugRecover != 0
	if !g.ExecFlags.IsDefer() {
		if debug {
			base.Debugf("recover() not directly inside a defer")
		}
		return nilInterface
	}
	if g.PanicFun == nil {
		if debug {
			base.Debugf("recover() no panic")
		}
		return nilInterface
	}
	if g.DeferOfFun != g.PanicFun {
		if debug {
			base.Debugf("recover() inside defer of function %p, not defer of the current panicking function %p", g.DeferOfFun, g.PanicFun)
		}
		return nilInterface
	}
	rec := g.Panic
	if rec == nil {
		if debug {
			base.Debugf("recover() consuming current panic: nil")
		}
		v = nilInterface
	} else {
		if debug {
			base.Debugf("recover() consuming current panic: %v <%v>", rec, r.TypeOf(rec))
		}
		v = r.ValueOf(rec).Convert(base.TypeOfInterface) // keep the interface{} type
	}
	// consume the current panic
	g.Panic = nil
	g.PanicFun = nil
	return v
}

func argEnv(env *Env) r.Value {
	return r.ValueOf(env)
}

func compileRecover(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	ti := c.TypeOfInterface()
	t := c.Universe.FuncOf([]xr.Type{ti}, []xr.Type{ti}, false)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: callRecover}, &sym)
	arg := exprX1(ti, argEnv)
	return newCall1(fun, arg, false, ti)
}

// ============================ support functions =============================

// call_builtin compiles a call to a builtin function: append, cap, copy, delete, len, make, new...
func (c *Comp) call_builtin(call *Call) I {
	// builtin functions are always literals, i.e. funindex == NoIndex thus not stored in Env.Binds[]
	// we must retrieve them directly from c.Fun.Value
	if !call.Fun.Const() {
		base.Errorf("internal error: call_builtin() invoked for non-constant function %#v. use one of the callXretY() instead", call.Fun)
	}
	var name string
	if call.Fun.Sym != nil {
		name = call.Fun.Sym.Name
	}
	args := call.Args
	argfuns := make([]I, len(args))
	for i, arg := range args {
		argfuns[i] = arg.WithFun()
	}
	if false {
		argtypes := make([]xr.Type, len(args))
		for i, arg := range args {
			argtypes[i] = arg.Type
		}
		// Debugf("compiling builtin %s() <%v> with arg types %v", name, TypeOf(c.Fun.Value), argtypes)
	}
	var ret I
	switch fun := call.Fun.Value.(type) {
	case UntypedLit: // complex(), real(), imag() of untyped constants
		ret = fun
	case func(float32, float32) complex64: // complex
		arg0fun := argfuns[0].(func(*Env) float32)
		arg1fun := argfuns[1].(func(*Env) float32)
		if name == "complex" {
			if args[0].Const() {
				arg0 := args[0].Value.(float32)
				ret = func(env *Env) complex64 {
					arg1 := arg1fun(env)
					return complex(arg0, arg1)
				}
			} else if args[1].Const() {
				arg1 := args[1].Value.(float32)
				ret = func(env *Env) complex64 {
					arg0 := arg0fun(env)
					return complex(arg0, arg1)
				}
			} else {
				ret = func(env *Env) complex64 {
					arg0 := arg0fun(env)
					arg1 := arg1fun(env)
					return complex(arg0, arg1)
				}
			}
		} else {
			ret = func(env *Env) complex64 {
				arg0 := arg0fun(env)
				arg1 := arg1fun(env)
				return fun(arg0, arg1)
			}
		}
	case func(float64, float64) complex128: // complex()
		arg0fun := argfuns[0].(func(*Env) float64)
		arg1fun := argfuns[1].(func(*Env) float64)
		if name == "complex" {
			if args[0].Const() {
				arg0 := args[0].Value.(float64)
				ret = func(env *Env) complex128 {
					arg1 := arg1fun(env)
					return complex(arg0, arg1)
				}
			} else if args[1].Const() {
				arg1 := args[1].Value.(float64)
				ret = func(env *Env) complex128 {
					arg0 := arg0fun(env)
					return complex(arg0, arg1)
				}
			} else {
				ret = func(env *Env) complex128 {
					arg0 := arg0fun(env)
					arg1 := arg1fun(env)
					return complex(arg0, arg1)
				}
			}
		} else {
			ret = func(env *Env) complex128 {
				arg0 := arg0fun(env)
				arg1 := arg1fun(env)
				return fun(arg0, arg1)
			}
		}
	case func(complex64) float32: // real(), imag()
		argfun := argfuns[0].(func(*Env) complex64)
		if name == "real" {
			ret = func(env *Env) float32 {
				arg := argfun(env)
				return real(arg)
			}
		} else if name == "imag" {
			ret = func(env *Env) float32 {
				arg := argfun(env)
				return imag(arg)
			}
		} else {
			ret = func(env *Env) float32 {
				arg := argfun(env)
				return fun(arg)
			}
		}
	case func(complex128) float64: // real(), imag()
		argfun := argfuns[0].(func(*Env) complex128)
		if name == "real" {
			ret = func(env *Env) float64 {
				arg := argfun(env)
				return real(arg)
			}
		} else if name == "imag" {
			ret = func(env *Env) float64 {
				arg := argfun(env)
				return imag(arg)
			}
		} else {
			ret = func(env *Env) float64 {
				arg := argfun(env)
				return fun(arg)
			}
		}
	case func(string) int: // len(string)
		argfun := argfuns[0].(func(*Env) string)
		if name == "len" {
			ret = func(env *Env) int {
				arg := argfun(env)
				return len(arg)
			}
		} else {
			ret = func(env *Env) int {
				arg := argfun(env)
				return fun(arg)
			}
		}
	case func([]byte, string) int: // copy([]byte, string)
		arg0fun := args[0].AsX1()
		if args[1].Const() {
			// string is a literal
			arg1const := args[1].Value.(string)
			ret = func(env *Env) int {
				// arg0 is "assignable to []byte"
				arg0 := arg0fun(env)
				if arg0.Type() != rtypeOfSliceOfByte {
					arg0 = convert(arg0, rtypeOfSliceOfByte)
				}
				return fun(arg0.Interface().([]byte), arg1const)
			}
		} else {
			arg1fun := args[1].Fun.(func(*Env) string)
			ret = func(env *Env) int {
				// arg0 is "assignable to []byte"
				arg0 := arg0fun(env)
				if arg0.Type() != rtypeOfSliceOfByte {
					arg0 = convert(arg0, rtypeOfSliceOfByte)
				}
				arg1 := arg1fun(env)
				return fun(arg0.Interface().([]byte), arg1)
			}
		}
	case func(interface{}): // panic()
		argfunsX1 := call.MakeArgfunsX1()
		argfun := argfunsX1[0]
		if name == "panic" {
			ret = func(env *Env) {
				arg := argfun(env).Interface()
				panic(arg)
			}
		} else {
			ret = func(env *Env) {
				arg := argfun(env).Interface()
				fun(arg)
			}
		}
	case func(interface{}, ...interface{}): // print, println()
		argfunsX1 := call.MakeArgfunsX1()
		if call.Ellipsis {
			argfuns := [2]func(*Env) r.Value{
				argfunsX1[0],
				argfunsX1[1],
			}
			ret = func(env *Env) {
				arg := argfuns[0](env).Interface()
				argslice := argfuns[1](env).Interface().([]interface{})
				fun(arg, argslice...)
			}
		} else {
			ret = func(env *Env) {
				args := make([]interface{}, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					args[i] = argfun(env).Interface()
				}
				fun(args[0], args[1:]...)
			}
		}
	case func(r.Value): // close()
		argfun := call.MakeArgfunsX1()[0]
		if name == "close" {
			ret = func(env *Env) {
				arg := argfun(env)
				arg.Close()
			}
		} else {
			ret = func(env *Env) {
				arg := argfun(env)
				fun(arg)
			}
		}
	case func(r.Value) int: // cap(), len()
		argfun := call.MakeArgfunsX1()[0]
		ret = func(env *Env) int {
			arg := argfun(env)
			return fun(arg)
		}
	case func(r.Value) r.Value: // Env()
		argfun := call.MakeArgfunsX1()[0]
		if name == "Interp" {
			ret = func(env *Env) r.Value {
				return argfun(env)
			}
		} else {
			ret = func(env *Env) r.Value {
				arg0 := argfun(env)
				return fun(arg0)
			}
		}
	case func(r.Value, r.Value): // delete()
		argfunsX1 := call.MakeArgfunsX1()
		argfuns := [2]func(env *Env) r.Value{
			argfunsX1[0],
			argfunsX1[1],
		}
		ret = func(env *Env) {
			arg0 := argfuns[0](env)
			arg1 := argfuns[1](env)
			fun(arg0, arg1)
		}
	case func(r.Value, r.Value) int: // copy()
		argfunsX1 := call.MakeArgfunsX1()
		argfuns := [2]func(env *Env) r.Value{
			argfunsX1[0],
			argfunsX1[1],
		}
		ret = func(env *Env) int {
			arg0 := argfuns[0](env)
			arg1 := argfuns[1](env)
			return fun(arg0, arg1)
		}
	case func(r.Value, r.Value) r.Value: // Eval(), EvalType(), Parse()
		argfunsX1 := call.MakeArgfunsX1()
		argfuns := [2]func(env *Env) r.Value{
			argfunsX1[0],
			argfunsX1[1],
		}
		ret = func(env *Env) r.Value {
			arg0 := argfuns[0](env)
			arg1 := argfuns[1](env)
			return fun(arg0, arg1)
		}
	case func(r.Value, r.Value) (r.Value, r.Value): // MacroExpand*()
		argfunsX1 := call.MakeArgfunsX1()
		argfuns := [2]func(env *Env) r.Value{
			argfunsX1[0],
			argfunsX1[1],
		}
		ret = func(env *Env) (r.Value, []r.Value) {
			arg0 := argfuns[0](env)
			arg1 := argfuns[1](env)
			ret0, ret1 := fun(arg0, arg1)
			return ret0, []r.Value{ret0, ret1}
		}
	case func(r.Value, ...r.Value) r.Value: // append()
		argfunsX1 := call.MakeArgfunsX1()
		if call.Ellipsis {
			argfuns := [2]func(*Env) r.Value{
				argfunsX1[0],
				argfunsX1[1],
			}
			if name == "append" {
				ret = func(env *Env) r.Value {
					arg0 := argfuns[0](env)
					arg1 := argfuns[1](env)
					argslice := unwrapSlice(arg1)
					return r.Append(arg0, argslice...)
				}
			} else {
				ret = func(env *Env) r.Value {
					arg0 := argfuns[0](env)
					arg1 := argfuns[1](env)
					argslice := unwrapSlice(arg1)
					return fun(arg0, argslice...)
				}
			}
		} else {
			if name == "append" {
				ret = func(env *Env) r.Value {
					args := make([]r.Value, len(argfunsX1))
					for i, argfun := range argfunsX1 {
						args[i] = argfun(env)
					}
					return r.Append(args[0], args[1:]...)
				}
			} else {
				ret = func(env *Env) r.Value {
					args := make([]r.Value, len(argfunsX1))
					for i, argfun := range argfunsX1 {
						args[i] = argfun(env)
					}
					return fun(args[0], args[1:]...)
				}
			}
		}
	case func(r.Type) r.Value: // new(), make()
		arg0 := args[0].Value.(r.Type)
		if name == "new" {
			ret = func(env *Env) r.Value {
				return r.New(arg0)
			}
		} else {
			ret = func(env *Env) r.Value {
				return fun(arg0)
			}
		}
	case func(r.Type, int) r.Value: // make()
		arg0 := args[0].Value.(r.Type)
		arg1fun := argfuns[1].(func(*Env) int)
		ret = func(env *Env) r.Value {
			arg1 := arg1fun(env)
			return fun(arg0, arg1)
		}
	case func(r.Type, int, int) r.Value: // make()
		arg0 := args[0].Value.(r.Type)
		arg1fun := argfuns[1].(func(*Env) int)
		arg2fun := argfuns[2].(func(*Env) int)
		ret = func(env *Env) r.Value {
			arg1 := arg1fun(env)
			arg2 := arg2fun(env)
			return fun(arg0, arg1, arg2)
		}
	default:
		base.Errorf("unimplemented call_builtin() for function type %v", r.TypeOf(fun))
	}
	return ret
}

// unwrapSlice accepts a reflect.Value with kind == reflect.Array, Slice or String
// and returns slice of its elements, each wrapped in a reflect.Value
func unwrapSlice(arg r.Value) []r.Value {
	n := arg.Len()
	slice := make([]r.Value, n)
	for i := range slice {
		slice[i] = arg.Index(i)
	}
	return slice
}

// callBuiltin invokes the appropriate compiler for a call to a builtin function: cap, copy, len, make, new...
func (c *Comp) callBuiltin(node *ast.CallExpr, fun *Expr) *Call {
	builtin := fun.Value.(Builtin)
	if fun.Sym == nil {
		c.Errorf("invalid call to non-name builtin: %v", node)
		return nil
	}
	nmin := builtin.ArgMin
	nmax := builtin.ArgMax
	n := len(node.Args)
	if n < int(nmin) || n > int(nmax) {
		return c.badBuiltinCallArgNum(fun.Sym.Name+"()", nmin, nmax, node.Args)
	}
	call := builtin.Compile(c, *fun.Sym, node)
	if call != nil {
		call.Builtin = true
	}
	return call
}

// callFunction compiles a call to a function that accesses interpreter's *CompEnv
func (c *Comp) callFunction(node *ast.CallExpr, fun *Expr) (newfun *Expr, lastarg *Expr) {
	function := fun.Value.(Function)
	t := function.Type
	var sym *Symbol
	if fun.Sym != nil {
		symcopy := *fun.Sym
		symcopy.Type = t
		sym = &symcopy
	}
	newfun = exprLit(Lit{Type: t, Value: function.Fun}, sym)
	if len(node.Args) < t.NumIn() {
		lastarg = exprX1(c.TypeOfInterface(), func(env *Env) r.Value {
			return r.ValueOf(&Interp{Comp: c, env: env})
		})
	}
	return newfun, lastarg
}

func (c *Comp) badBuiltinCallArgNum(name interface{}, nmin uint16, nmax uint16, args []ast.Expr) *Call {
	prefix := "not enough"
	nargs := len(args)
	if nargs > int(nmax) {
		prefix = "too many"
	}
	str := fmt.Sprintf("%d", nmin)
	if nmax <= nmin {
	} else if nmax == nmin+1 {
		str = fmt.Sprintf("%s or %d", str, nmax)
	} else if nmax < base.MaxUint16 {
		str = fmt.Sprintf("%s to %d", str, nmax)
	} else {
		str = fmt.Sprintf("%s or more", str)
	}
	c.Errorf("%s arguments in call to builtin %v: expecting %s, found %d: %v", prefix, name, str, nargs, args)
	return nil
}

func (c *Comp) badBuiltinCallArgType(name string, arg ast.Expr, tactual xr.Type, texpected interface{}) *Call {
	c.Errorf("cannot use %v <%v> as %v in builtin %s()", arg, tactual, texpected, name)
	return nil
}

func anyToAst(any interface{}, caller interface{}) ast2.Ast {
	if untyped, ok := any.(UntypedLit); ok {
		any = untyped.Convert(untyped.DefaultType())
	}
	return ast2.AnyToAst(any, caller)
}
