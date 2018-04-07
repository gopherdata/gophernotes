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
 * inspect.go
 *
 *  Created on: Feb 11, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"errors"
	"fmt"
	r "reflect"
	"strconv"
	"strings"

	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type Inspector struct {
	names []string
	vs    []r.Value
	ts    []r.Type
	xts   []xr.Type
	in    Readline
	env   *Env
}

func (env *Env) Inspect(in Readline, str string, fastInterpreter bool) {
	form := env.Parse(str)
	var v r.Value
	var xt xr.Type
	if fastInterpreter {
		v, _, xt, _ = env.fastEval(form)
	} else {
		v = env.EvalAst1(form)
	}
	var t r.Type
	if v != Nil && v != None {
		t = r.TypeOf(v.Interface()) // show concrete type
	}
	switch dereferenceValue(v).Kind() {
	case r.Array, r.Slice, r.String, r.Struct:
		break
	default:
		env.showVar(str, v, t)
		env.showMethods(t, xt)
		return
	}
	stack := Inspector{names: []string{str}, vs: []r.Value{v}, ts: []r.Type{t}, xts: []xr.Type{xt}, in: in, env: env}
	stack.Show()
	stack.Repl()
}

func (env *Env) showMethods(t r.Type, xt xr.Type) {
	switch {
	case xt != nil:
		if xt.Kind() == r.Ptr {
			xt = xt.Elem()
		}
		n := xt.NumMethod()
		if n == 0 {
			env.Fprintf(env.Stdout, "no methods of %v\n", xt)
			return
		}
		env.Fprintf(env.Stdout, "methods of %v:\n", xt)
		for i := 0; i < n; i++ {
			env.Fprintf(env.Stdout, "    m%d. %v\n", i, xt.Method(i).GoFun)
		}

	case t != nil:
		n := t.NumMethod()
		if n == 0 {
			env.Fprintf(env.Stdout, "no methods of %v\n", t)
			return
		}
		env.Fprintf(env.Stdout, "methods of %v:\n", t)
		for i := 0; i < n; i++ {
			m := t.Method(i)
			env.Fprintf(env.Stdout, "    m%d. %s\t%v\n", i, m.Name, m.Type)
		}
	}
}

func (env *Env) showVar(str string, v r.Value, t r.Type) {
	env.Fprintf(env.Stdout, "%s\t= %v\t// %v\n", str, v, t)
}

func (ip *Inspector) Help() {
	fmt.Fprint(ip.env.Stdout, "// inspect commands: <number> help methods quit top up\n")
}

func (ip *Inspector) Show() {
	depth := len(ip.names)
	name := strings.Join(ip.names, ".")
	v := ip.vs[depth-1]
	t := ip.ts[depth-1]
	ip.env.showVar(name, v, t)

	v = dereferenceValue(v) // dereference pointers on-the-fly
	switch v.Kind() {
	case r.Array, r.Slice, r.String:
		ip.showIndexes(v)
	case r.Struct:
		ip.showFields(v)
	}
}

func (ip *Inspector) Repl() error {
	for len(ip.names) > 0 {
		prompt := fmt.Sprintf("goinspect %s> ", strings.Join(ip.names, "."))
		bytes, err := ip.in.Read(prompt)
		if err != nil {
			return err
		}
		cmd := strings.TrimSpace(string(bytes))
		err = ip.Eval(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ip *Inspector) Eval(cmd string) error {
	switch {
	case cmd == "?", strings.HasPrefix("help", cmd):
		ip.Help()
	case strings.HasPrefix("methods", cmd):
		t := ip.ts[len(ip.ts)-1]
		xt := ip.xts[len(ip.xts)-1]
		ip.env.showMethods(t, xt)
	case strings.HasPrefix("quit", cmd):
		return errors.New("user quit")
	case strings.HasPrefix("top", cmd):
		ip.Top()
		ip.Show()
	case cmd == "", cmd == ".":
		ip.Show()
	case cmd == "-", strings.HasPrefix("up", cmd):
		ip.Leave()
	default:
		ip.Enter(cmd)
	}
	return nil
}

func (ip *Inspector) Top() {
	ip.names = ip.names[0:1]
	ip.vs = ip.vs[0:1]
	ip.ts = ip.ts[0:1]
}

func (ip *Inspector) Leave() {
	depth := len(ip.names)
	if depth <= 0 {
		return
	}
	depth--
	ip.names = ip.names[:depth]
	ip.vs = ip.vs[:depth]
	ip.ts = ip.ts[:depth]
	if depth > 0 {
		ip.Show()
	}
}

func (ip *Inspector) showFields(v r.Value) {
	n := v.NumField()
	for i := 0; i < n; i++ {
		f := v.Field(i)
		t := typeOf(f)
		f = dereferenceValue(f)
		fmt.Fprintf(ip.env.Stdout, "    %d. ", i)
		ip.env.showVar(v.Type().Field(i).Name, f, t)
	}
}

func (ip *Inspector) showIndexes(v r.Value) {
	n := v.Len()
	for i := 0; i < n; i++ {
		f := v.Index(i)
		t := typeOf(f)
		f = dereferenceValue(f)
		fmt.Fprintf(ip.env.Stdout, "    %d. ", i)
		ip.env.showVar("", f, t)
	}
}

func (ip *Inspector) Enter(cmd string) {
	i, err := strconv.Atoi(cmd)
	if err != nil {
		fmt.Fprintf(ip.env.Stdout, "unknown inspect command \"%s\". Type ? for help\n", cmd)
		return
	}
	depth := len(ip.names)
	v := dereferenceValue(ip.vs[depth-1])
	var n int
	var fname string
	var f r.Value
	switch v.Kind() {
	case r.Array, r.Slice, r.String:
		n = v.Len()
		if !ip.validRange(i, n) {
			return
		}
		fname = fmt.Sprintf("[%s]", cmd)
		f = v.Index(i)
	case r.Struct:
		n = v.NumField()
		if !ip.validRange(i, n) {
			return
		}
		fname = v.Type().Field(i).Name
		f = v.Field(i)
	default:
		fmt.Fprintf(ip.env.Stdout, "cannot enter <%v>: expecting array, slice, string or struct\n", typeOf(v))
		return
	}
	var t r.Type
	if f != Nil && f != None {
		if f.CanInterface() {
			t = r.TypeOf(f.Interface()) // concrete type
		} else {
			t = f.Type()
		}
	}

	switch dereferenceValue(f).Kind() { // dereference pointers on-the-fly
	case r.Array, r.Slice, r.String, r.Struct:
		ip.names = append(ip.names, fname)
		ip.vs = append(ip.vs, f)
		ip.ts = append(ip.ts, t)
		ip.Show()
	default:
		ip.env.showVar(fname, f, t)
	}
}

func dereferenceValue(v r.Value) r.Value {
	for {
		switch v.Kind() {
		case r.Ptr:
			v = v.Elem()
			continue
		case r.Interface:
			v = r.ValueOf(v.Interface())
			continue
		}
		break
	}
	return v
}

func (ip *Inspector) validRange(i, n int) bool {
	if i < 0 || i >= n {
		fmt.Fprintf(ip.env.Stdout, "%s contains %d elements, cannot inspect element %d\n",
			strings.Join(ip.names, "."), n, i)
		return false
	}
	return true
}
