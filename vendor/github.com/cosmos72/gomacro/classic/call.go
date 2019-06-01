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
 * call.go
 *
 *  Created on: Mar 30, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"go/ast"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/reflect"
)

// eval an interpreted function
func (env *Env) evalFuncCall(envName string, body *ast.BlockStmt, t r.Type, argNames []string, args []r.Value, resultNames []string) (results []r.Value) {
	if t.Kind() != r.Func {
		env.Errorf("call of non-function type %v", t)
		return nil
	}
	env = NewEnv(env, envName)
	// register this function call in the call stack
	env.CallStack.Frames = append(env.CallStack.Frames, CallFrame{FuncEnv: env})
	debugCall := env.Options&OptDebugCallStack != 0
	if debugCall {
		env.Debugf("func starting: %s, args = %v, call stack is:", envName, args)
		env.showStack()
	}

	panicking := true // use a flag to distinguish non-panic from panic(nil)
	defer func() {
		if debugCall {
			env.Debugf("func exiting:  %s, panicking = %v, stack length = %d",
				envName, panicking, len(env.CallStack.Frames))
		}
		frame := env.CurrentFrame()
		if panicking {
			pan := recover()
			switch p := pan.(type) {
			case eReturn:
				// return is implemented with a panic(eReturn{})
				results = env.convertFuncCallResults(t, p.results, true)
			default: // some interpreted or compiled code invoked panic()
				if env.Options&OptDebugRecover != 0 {
					env.Debugf("captured panic for defers: env = %v, panic = %#v", env.Name, p)
				}
				frame.panick = p
				frame.panicking = true
			}
		}
		if len(frame.defers) != 0 {
			frame.runDefers(env)
		}
		stack := env.CallStack
		stack.Frames = stack.Frames[0 : len(stack.Frames)-1]

		if debugCall {
			str := "is"
			if frame.panicking {
				str = "="
			}
			env.Debugf("func exited:   %s, panic     %s %v, stack length = %d",
				envName, str, frame.panick, len(stack.Frames))
		}

		if frame.panicking {
			panic(frame.panick)
		}
	}()

	for i, resultName := range resultNames {
		env.DefineVar(resultName, t.Out(i), r.Zero(t.Out(i)))
	}
	for i, argName := range argNames {
		env.DefineVar(argName, t.In(i), args[i])
	}
	// use evalStatements(), not evalBlock(): in Go, the function arguments and body are in the same scope
	rets := reflect.PackValues(env.evalStatements(body.List))
	results = env.convertFuncCallResults(t, rets, false)
	panicking = false
	return results
}

func (env *Env) convertFuncCallResults(t r.Type, rets []r.Value, warn bool) []r.Value {
	retsN := len(rets)
	expectedN := t.NumOut()
	if retsN < expectedN {
		if warn {
			env.Warnf("not enough return values: expecting %d, found %d: %v", expectedN, retsN, rets)
		}
		tmp := make([]r.Value, expectedN)
		copy(tmp, rets)
		rets = tmp
	} else if retsN > expectedN {
		if warn {
			env.Warnf("too many return values: expecting %d, found %d: %v", expectedN, retsN, rets)
		}
		rets = rets[:expectedN]
	}
	for i, ret := range rets {
		ti := t.Out(i)
		if ret == Nil || ret == None {
			rets[i] = r.Zero(ti)
		} else {
			rets[i] = ret.Convert(ti)
		}
	}
	return rets
}

func (frame *CallFrame) runDefers(env *Env) {
	// execute defers last-to-first
	frame.runningDefers = true
	if env.Options&OptDebugCallStack != 0 {
		str := "is"
		if frame.panicking {
			str = "="
		}
		env.Debugf("func defers:   %s, panic %s %v, stack length = %d",
			env.Name, str, frame.panick, len(env.CallStack.Frames))
	}
	defers := frame.defers
	for i := len(defers) - 1; i >= 0; i-- {
		frame.runDefer(defers[i])
	}
}

func (frame *CallFrame) runDefer(deferred func()) {
	// invoking panic() inside a deferred function exits it with a panic,
	// but the previously-installed deferred functions are still executed
	// and can recover() such panic

	panicking := true // use a flag to distinguish non-panic from panic(nil)
	defer func() {
		if panicking {
			frame.panick = recover()
			frame.panicking = true
		}
	}()
	deferred()
	panicking = false
}

func (env *Env) evalCall(node *ast.CallExpr) (r.Value, []r.Value) {
	var fun r.Value
	var t r.Type
	if len(node.Args) == 1 {
		// may be a type conversion
		fun, t = env.evalExpr1OrType(node.Fun)
	} else {
		fun = env.evalExpr1(node.Fun)
	}

	if t != nil {
		val := env.evalExpr1(node.Args[0])
		return env.valueToType(val, t), nil
	}

	{
		frames := env.CallStack.Frames
		frame := &frames[len(frames)-1]
		frame.CurrentCall = node
		frame.InnerEnv = env // leaks a bit... should be cleared after the call
	}

	switch fun.Kind() {
	case r.Struct:
		switch fun := fun.Interface().(type) {
		case Constructor:
			t, args := env.evalConstructorArgs(fun, node)
			return fun.exec(env, t, args)
		case Function:
			args := env.evalFunctionArgs(fun, node)
			return fun.exec(env, args)
		}
	case r.Func:
		args := env.evalFuncArgs(fun, node)
		var rets []r.Value

		if node.Ellipsis == token.NoPos {
			rets = fun.Call(args)
		} else {
			rets = fun.CallSlice(args)
		}
		return reflect.UnpackValues(rets)
	default:
		break
	}
	return env.Errorf("call of non-function %v <%v>: %v", reflect.Interface(fun), reflect.Type(fun), node)
}

func (env *Env) evalConstructorArgs(fun Constructor, node *ast.CallExpr) (r.Type, []r.Value) {
	args := node.Args
	if fun.argNum >= 0 && fun.argNum != len(args) {
		env.Errorf("builtin %v expects %d arguments, found %d",
			node.Fun, fun.argNum, len(args))
	} else if len(args) == 0 {
		env.Errorf("builtin %v expects at least one argument, found zero", node.Fun)
	}
	t := env.evalType(args[0])
	return t, env.evalExprs(args[1:])
}

func (env *Env) evalFunctionArgs(fun Function, node *ast.CallExpr) []r.Value {
	args := node.Args
	if fun.argNum >= 0 && fun.argNum != len(args) {
		env.Errorf("function %v expects %d arguments, found %d",
			node.Fun, fun.argNum, len(args))
	}
	return env.evalExprs(args)
}

func (env *Env) evalFuncArgs(fun r.Value, node *ast.CallExpr) []r.Value {
	funt := fun.Type()
	nin := funt.NumIn()
	var args []r.Value
	if len(node.Args) == 1 && nin > 1 {
		// special case fooAcceptsMultipleArgs( barReturnsMultipleValues() )
		args = env.evalExprsMultipleValues(node.Args, nin)
	} else {
		args = env.evalExprs(node.Args)
	}
	if funt.IsVariadic() == (node.Ellipsis != token.NoPos) {
		if len(args) != nin {
			env.Errorf("function %v expects %d arguments, found %d: %v", node.Fun, nin, len(args), args)
			return nil
		}
		for i, arg := range args {
			args[i] = env.valueToType(arg, funt.In(i))
		}
	}
	return args
}

func (env *Env) evalDefer(node *ast.CallExpr) (r.Value, []r.Value) {
	frame := env.CurrentFrame()
	if frame == nil {
		return env.Errorf("defer outside function: %v", node)
	}
	fun := env.evalExpr1(node.Fun)
	if fun.Kind() != r.Func {
		return env.Errorf("defer of non-function: %v", node)
	}
	args := env.evalFuncArgs(fun, node)
	closure := func() {
		var rets []r.Value
		if node.Ellipsis == token.NoPos {
			rets = fun.Call(args)
		} else {
			rets = fun.CallSlice(args)
		}
		if len(rets) != 0 {
			env.Warnf("call to deferred function %v returned %d values, expecting zero: %v", node, rets)
		}
	}
	frame.defers = append(frame.defers, closure)
	return None, nil
}
