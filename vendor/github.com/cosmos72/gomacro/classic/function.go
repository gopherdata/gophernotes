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
 * function.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"fmt"
	"go/ast"
	r "reflect"
)

func (env *Env) evalDeclFunction(decl *ast.FuncDecl, funcType *ast.FuncType, body *ast.BlockStmt) (r.Value, []r.Value) {
	isMacro := false
	var recv *ast.Field

	if decl != nil && decl.Recv != nil {
		recvList := decl.Recv.List
		if recvList != nil && len(recvList) == 0 {
			isMacro = true
		} else {
			recv = recvList[0]
		}
	}
	tFunc, tFuncOrMethod, argNames, resultNames := env.evalTypeFunctionOrMethod(recv, funcType)
	tret := tFuncOrMethod

	var funcName string
	if decl == nil {
		funcName = makeFuncNameForEnv(decl, isMacro)
	} else {
		funcName = decl.Name.Name
	}

	closure := func(args []r.Value) (results []r.Value) {
		return env.evalFuncCall(funcName, body, tFuncOrMethod, argNames, args, resultNames)
	}
	var ret r.Value
	if isMacro {
		// env.Debugf("defined macro %v, type %v, args (%v), returns (%v)", decl.Name.Name, t, strings.Join(argNames, ", "), strings.Join(resultNames, ", "))
		ret = r.ValueOf(Macro{closure: closure, argNum: len(argNames)})
		tret = ret.Type()
	} else {
		ret = r.MakeFunc(tFuncOrMethod, closure)

		if decl != nil && recv != nil {
			recvType := tFuncOrMethod.In(0)
			// register tFunc, i.e. without the receiver, to allow comparison with Interface methods
			env.registerMethod(recvType, funcName, tFunc, ret)
		}
	}
	if decl != nil && recv == nil {
		// register named functions and macros (NOT methods) in the current environment
		ret = env.DefineFunc(funcName, tret, ret)
	}
	return ret, nil
}

func makeFuncNameForEnv(decl *ast.FuncDecl, isMacro bool) string {
	var prefix, space, suffix string = "func", "", ""
	if isMacro {
		prefix = "macro"
	}
	if decl != nil {
		space = " "
		suffix = decl.Name.Name
	}
	return fmt.Sprintf("%s%s%s()", prefix, space, suffix)
}
