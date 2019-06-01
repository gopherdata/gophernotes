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
 * global.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"go/ast"
	r "reflect"
)

type CallStack struct {
	Frames []CallFrame
}

type CallFrame struct {
	FuncEnv       *Env
	InnerEnv      *Env          // innermost Env
	CurrentCall   *ast.CallExpr // call currently in progress
	defers        []func()
	panick        interface{} // current panic
	panicking     bool
	runningDefers bool
}

type Constructor struct {
	exec   func(env *Env, arg0 r.Type, args []r.Value) (r.Value, []r.Value)
	argNum int // if negative, do not check
}

type Function struct {
	exec   func(env *Env, args []r.Value) (r.Value, []r.Value)
	argNum int // if negative, do not check
}

type Macro struct {
	closure func(args []r.Value) (results []r.Value)
	argNum  int
}

type TypedValue struct {
	typ r.Type
	val r.Value
}

/**
 * inside Methods, each string is the method name
 * and each TypedValue is {
 *   Type: the method signature, i.e. the type of a func() *without* the receiver (to allow comparison with Interface methods)
 *   Value: the method implementation, i.e. a func() whose first argument is the receiver,
 * }
 */
type Methods map[string]TypedValue
