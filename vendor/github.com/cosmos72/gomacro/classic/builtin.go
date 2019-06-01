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
 * builtin.go
 *
 *  Created on: Feb 15, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"fmt"
	"go/ast"
	"io/ioutil"
	r "reflect"

	. "github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/reflect"
)

func funcAppend(env *Env, args []r.Value) (r.Value, []r.Value) {
	n := len(args)
	if n < 1 {
		return env.Errorf("builtin append() expects at least one argument, found %d", n)
	}
	t := args[0].Type().Elem()
	for i := 1; i < n; i++ {
		args[i] = args[i].Convert(t)
	}
	return r.Append(args[0], args[1:]...), nil
}

func callCap(arg interface{}) int {
	obj := r.ValueOf(arg)
	if obj.Kind() == r.Ptr {
		// cap() of pointer-to-array returns cap() of array
		obj = obj.Elem()
	}
	return obj.Cap()
}

func callClose(channel interface{}) {
	r.ValueOf(channel).Close()
}

func funcComplex(env *Env, args []r.Value) (r.Value, []r.Value) {
	rv, iv := args[0], args[1]
	r_, rok := env.toFloat(rv)
	i_, iok := env.toFloat(iv)
	if !rok {
		return env.Errorf("builtin complex(): not a float: %v <%v>", rv, typeOf(rv))
	}
	if !iok {
		return env.Errorf("builtin complex(): not a float: %v <%v>", iv, typeOf(iv))
	}
	cplx := complex(r_, i_)
	var ret interface{}
	if rv.Kind() == r.Float32 && iv.Kind() == r.Float32 {
		ret = complex64(cplx)
	} else {
		ret = cplx
	}
	return r.ValueOf(ret), nil
}

func callCopy(dst, src interface{}) int {
	if src, ok := src.(string); ok {
		if dst, ok := dst.([]byte); ok {
			// reflect.Copy does not support this case... use the compiler support
			return copy(dst, src)
		}
	}
	return r.Copy(r.ValueOf(dst), r.ValueOf(src))
}

func callDelete(m interface{}, key interface{}) {
	vmap := r.ValueOf(m)
	tkey := vmap.Type().Key()
	vkey := r.ValueOf(key)
	if key != nil && vkey.Type() != tkey {
		vkey = vkey.Convert(tkey)
	}
	vmap.SetMapIndex(vkey, Nil)
}

func funcEnv(env *Env, args []r.Value) (r.Value, []r.Value) {
	return r.ValueOf(env), nil
}

func funcEval(env *Env, args []r.Value) (r.Value, []r.Value) {
	arg := args[0]
	if arg == Nil || arg == None {
		return arg, nil
	}
	x := toInterface(arg)
	form := AnyToAst(x, "Eval")
	return env.EvalAst(form)
}

func funcEvalType(env *Env, args []r.Value) (r.Value, []r.Value) {
	arg := args[0]
	if arg == Nil || arg == None {
		return arg, nil
	}
	x := toInterface(arg)
	form := UnwrapTrivialAst(AnyToAst(x, "EvalType"))

	switch node := ToNode(form).(type) {
	case ast.Expr:
		// return nil for *ast.Ident{Name: "nil"}
		t := env.evalTypeOrNil(node)
		if t == nil {
			return Nil, nil
		}
		// return as reflect.Type, not as the concrete struct *reflect.type
		return r.ValueOf(&t).Elem(), nil
	default:
		return env.Errorf("EvalType: expecting <ast.Expr>, found: %v <%v>", node, r.TypeOf(node))
	}
}

func funcImag(env *Env, args []r.Value) (r.Value, []r.Value) {
	cv := args[0]
	c_, ok := env.toComplex(cv)
	if !ok {
		return env.Errorf("builtin imag(): not a complex: %v <%v>", cv, typeOf(cv))
	}
	i_ := imag(c_)
	var ret interface{}
	if cv.Kind() == r.Complex64 {
		ret = float32(i_)
	} else {
		ret = i_
	}
	return r.ValueOf(ret), nil
}

func callLen(arg interface{}) int {
	obj := r.ValueOf(arg)
	if obj.Kind() == r.Ptr {
		// len() of pointer-to-array returns len() of array
		obj = obj.Elem()
	}
	return obj.Len()
}

//
// --------- macroexpansion ----------
//

func funcMacroExpand(env *Env, args []r.Value) (r.Value, []r.Value) {
	return callMacroExpand(env, args, CMacroExpand)
}

func funcMacroExpand1(env *Env, args []r.Value) (r.Value, []r.Value) {
	return callMacroExpand(env, args, CMacroExpand1)
}

func funcMacroExpandCodewalk(env *Env, args []r.Value) (r.Value, []r.Value) {
	return callMacroExpand(env, args, CMacroExpandCodewalk)
}

func callMacroExpand(env *Env, args []r.Value, which WhichMacroExpand) (r.Value, []r.Value) {
	n := len(args)
	if n < 1 || n > 2 {
		return env.Errorf("builtin %v() expects one or two arguments, found %d: %v", which, n, args)
	}
	val := args[0]
	if val == Nil || val == None {
		return val, nil
	}
	form := AnyToAst(val.Interface(), which.String())
	if n == 2 {
		e := args[1]
		if e != Nil && e != None {
			env = e.Interface().(*Env)
		}
	}
	var expanded bool
	switch which {
	case CMacroExpand1:
		form, expanded = env.macroExpandAstOnce(form)
	case CMacroExpandCodewalk:
		form, expanded = env.MacroExpandAstCodewalk(form)
	default:
		form, expanded = env.macroExpandAst(form)
	}
	formv := r.ValueOf(form.Interface())
	return formv, []r.Value{formv, r.ValueOf(expanded)}
}

func funcMake(env *Env, t r.Type, args []r.Value) (r.Value, []r.Value) {
	n := len(args)
	if n > 2 {
		return env.Errorf("builtin make() expects one, two or three arguments, found %d", n+1)
	}
	ret := Nil
	switch t.Kind() {
	case r.Chan:
		buffer := 0
		if n > 0 {
			buffer = int(args[0].Int())
		}
		ret = r.MakeChan(t, buffer)
	case r.Map:
		ret = r.MakeMap(t)
	case r.Slice:
		length := 0
		if n > 0 {
			length = int(args[0].Int())
		}
		capacity := length
		if n > 1 {
			capacity = int(args[1].Int())
		}
		ret = r.MakeSlice(t, length, capacity)
	}
	return ret, nil
}

func funcNew(env *Env, t r.Type, args []r.Value) (r.Value, []r.Value) {
	return r.New(t), nil
}

func funcParse(env *Env, args []r.Value) (r.Value, []r.Value) {
	var in interface{}
	if arg := args[0]; arg != Nil && arg != None {
		in = arg.Interface()
	}
	out := env.Parse(in)
	if out != nil {
		return r.ValueOf(out.Interface()), nil
	}
	return Nil, nil
}

func callPanic(arg interface{}) {
	panic(arg)
}

func funcReal(env *Env, args []r.Value) (r.Value, []r.Value) {
	n := len(args)
	if n != 1 {
		return env.Errorf("builtin real() expects exactly one argument, found %d", n)
	}
	cv := args[0]
	c_, ok := env.toComplex(cv)
	if !ok {
		return env.Errorf("builtin real(): not a complex: %v <%v>", cv, typeOf(cv))
	}
	i_ := real(c_)
	var ret interface{}
	if cv.Kind() == r.Complex64 {
		ret = float32(i_)
	} else {
		ret = i_
	}
	return r.ValueOf(ret), nil
}

func callReadFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		callPanic(err)
	}
	return string(bytes)
}

func callReadDir(dirname string) []string {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		callPanic(err)
	}
	n := len(files)
	names := make([]string, n)
	for i := 0; i < n; i++ {
		names[i] = files[i].Name()
	}
	return names
}

func funcRecover(env *Env, args []r.Value) (r.Value, []r.Value) {
	// Go specs: "Executing a call to recover inside a deferred function
	// (but not any function called by it) stops the panicking sequence
	// by restoring normal execution and retrieves the error value passed to the call of panic"
	//
	// thus recover() is invoked inside deferred functions: find their caller's env
	ret := Nil

	trace := env.Options&OptDebugRecover != 0
	caller := env.CallerFrame()
	if trace {
		env.Debugf("recover(): env = %v, stack is:", env.Name)
		env.showStack()
		curr := env.CurrentFrame()
		if curr != nil {
			env.Debugf("           frame = %v, runningDefers = %v", curr.FuncEnv.Name, curr.runningDefers)
		} else {
			env.Debugf("           frame = nil")
		}
		if caller != nil {
			env.Debugf("           caller = %v, runningDefers = %v", caller.FuncEnv.Name, caller.runningDefers)
		} else {
			env.Debugf("           caller = nil")
		}
	}

	if caller != nil {
		if caller.runningDefers && caller.panicking {
			// consume current panic
			if trace {
				env.Debugf("           consuming current panic = %#v", caller.panick)
			}
			ret = r.ValueOf(caller.panick)
			caller.panick = nil
			caller.panicking = false
		} else if trace {
			env.Debugf("           no panic to consume: caller.runningDefers = %q, caller.panicking = %q",
				caller.runningDefers, caller.panicking)
		}
	}
	return ret, nil
}

func callSlice(args ...interface{}) []interface{} {
	return args
}

func funcValues(env *Env, args []r.Value) (r.Value, []r.Value) {
	for i, arg := range args {
		if arg != None && arg != Nil {
			if arg.Kind() == r.Interface {
				arg = arg.Elem() // extract concrete type
			}
			args[i] = arg
		}
	}
	return reflect.UnpackValues(args)
}

func (top *Env) addIota() {
	top.Binds.Set("iota", r.ValueOf(0))
}

func (top *Env) removeIota() {
	top.Binds.Del("iota")
}

func (top *Env) incrementIota() {
	uIota := int(top.Binds.Get1("iota").Int())
	top.Binds.Set("iota", r.ValueOf(uIota+1))
}

func (env *Env) addBuiltins() {
	binds := env.Binds.Ensure()

	binds.Set("Env", r.ValueOf(Function{funcEnv, 0}))
	binds.Set("Eval", r.ValueOf(Function{funcEval, 1}))
	binds.Set("EvalType", r.ValueOf(Function{funcEvalType, 1}))
	binds.Set("MacroExpand", r.ValueOf(Function{funcMacroExpand, -1}))
	binds.Set("MacroExpand1", r.ValueOf(Function{funcMacroExpand1, -1}))
	binds.Set("MacroExpandCodewalk", r.ValueOf(Function{funcMacroExpandCodewalk, -1}))
	binds.Set("Parse", r.ValueOf(Function{funcParse, 1}))
	binds.Set("Read", r.ValueOf(ReadString))
	binds.Set("ReadDir", r.ValueOf(callReadDir))
	binds.Set("ReadFile", r.ValueOf(callReadFile))
	binds.Set("ReadMultiline", r.ValueOf(ReadMultiline))
	binds.Set("Slice", r.ValueOf(callSlice))
	binds.Set("String", r.ValueOf(func(args ...interface{}) string {
		return env.ToString("", args...)
	}))
	// return multiple values, extracting the concrete type of each interface
	binds.Set("Values", r.ValueOf(Function{funcValues, -1}))

	binds.Set("append", r.ValueOf(Function{funcAppend, -1}))
	binds.Set("cap", r.ValueOf(callCap))
	binds.Set("close", r.ValueOf(callClose))
	binds.Set("complex", r.ValueOf(Function{funcComplex, 2}))
	binds.Set("copy", r.ValueOf(callCopy))
	binds.Set("delete", r.ValueOf(callDelete))
	binds.Set("false", r.ValueOf(false))
	binds.Set("imag", r.ValueOf(Function{funcImag, 1}))
	binds.Set("len", r.ValueOf(callLen))
	binds.Set("make", r.ValueOf(Constructor{funcMake, -1}))
	binds.Set("new", r.ValueOf(Constructor{funcNew, 1}))
	binds.Set("nil", Nil)
	binds.Set("panic", r.ValueOf(callPanic))
	binds.Set("print", r.ValueOf(func(args ...interface{}) {
		fmt.Fprint(env.Stdout, args...)
	}))
	binds.Set("println", r.ValueOf(func(args ...interface{}) {
		fmt.Fprintln(env.Stdout, args...)
	}))
	binds.Set("real", r.ValueOf(Function{funcReal, 1}))
	binds.Set("recover", r.ValueOf(Function{funcRecover, 0}))
	binds.Set("true", r.ValueOf(true))

	// --------- types ---------
	types := env.Types.Ensure()

	types.Set("bool", r.TypeOf(false))
	types.Set("byte", r.TypeOf(byte(0)))
	types.Set("complex64", r.TypeOf(complex64(0)))
	types.Set("complex128", r.TypeOf(complex128(0)))
	types.Set("error", r.TypeOf((*error)(nil)).Elem())
	types.Set("float32", r.TypeOf(float32(0)))
	types.Set("float64", r.TypeOf(float64(0)))
	types.Set("int", r.TypeOf(int(0)))
	types.Set("int8", r.TypeOf(int8(0)))
	types.Set("int16", r.TypeOf(int16(0)))
	types.Set("int32", r.TypeOf(int32(0)))
	types.Set("int64", r.TypeOf(int64(0)))
	types.Set("rune", r.TypeOf(rune(0)))
	types.Set("string", r.TypeOf(""))
	types.Set("uint", r.TypeOf(uint(0)))
	types.Set("uint8", r.TypeOf(uint8(0)))
	types.Set("uint16", r.TypeOf(uint16(0)))
	types.Set("uint32", r.TypeOf(uint32(0)))
	types.Set("uint64", r.TypeOf(uint64(0)))
	types.Set("uintptr", r.TypeOf(uintptr(0)))

	// --------- proxies ---------
	proxies := env.Proxies.Ensure()

	proxies.Set("error", r.TypeOf((*Error_builtin)(nil)).Elem())
}

type Error_builtin struct {
	Obj    interface{}
	Error_ func() string
}

func (Proxy *Error_builtin) Error() string {
	return Proxy.Error_()
}

func (env *Env) addInterpretedBuiltins() {
	if false {
		line := "func not(flag bool) bool { if flag { return false } else { return true } }"
		env.EvalAst(env.Parse(line))
	}
	if false {
		// Factorial(1000000): eval() elapsed time: 1.233714899 s
		line := "func Factorial(n int) int { t := 1; for i := 2; i <= n; i=i+1 { t = t * i }; t }"
		env.EvalAst(env.Parse(line))
	}
}
