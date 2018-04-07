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
