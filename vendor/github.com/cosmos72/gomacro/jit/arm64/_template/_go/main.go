/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * main.go
 *
 *  Created on Feb 02, 2019
 *      Author Massimiliano Ghilardi
 */

package main

import (
	"fmt"
	"reflect"
)

type EnvBinds struct {
	Vals []reflect.Value
	Ints []uint64
}

// simplified fast.Env
type Env struct {
	EnvBinds
	Outer   *Env
	IP      int
	Code    []Stmt
	Run     *Run
	FileEnv *Env
}

type Run struct {
}

type Stmt func(*Env) (Stmt, *Env)

func main() {
	fmt.Println(Add8, Add16, Add32, Add64,
		Zero0, Zero8, Zero16, Zero32, Zero64,
		Div8, Div16, Div32, Div64,
		UDiv8, UDiv16, UDiv32, UDiv64,
		Nop, Leave)

	var env Env
	t := reflect.TypeOf(env)
	showFields(t, "Ints", "Outer", "FileEnv")
}

func showFields(t reflect.Type, names ...string) {
	for _, name := range names {
		f, _ := t.FieldByName(name)
		fmt.Printf("%#v\n", f)
	}
}
