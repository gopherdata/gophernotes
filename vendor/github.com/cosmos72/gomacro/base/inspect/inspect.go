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
 * inspect.go
 *
 *  Created on: Apr 20, 2018
 *      Author: Massimiliano Ghilardi
 */

package inspect

import (
	"errors"
	r "reflect"
	"strconv"
	"strings"

	"github.com/cosmos72/gomacro/base/reflect"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type Inspector struct {
	names   []string
	vals    []r.Value
	types   []r.Type
	xtypes  []xr.Type
	globals *base.Globals
}

func (ip *Inspector) Inspect(name string, val r.Value, typ r.Type, xtyp xr.Type, globals *base.Globals) {
	ip.Init(name, val, typ, xtyp, globals)
	ip.Show()
	ip.Repl()
}

func (ip *Inspector) Init(name string, val r.Value, typ r.Type, xtyp xr.Type, globals *base.Globals) {
	ip.names = []string{name}
	ip.vals = []r.Value{val}
	ip.types = []r.Type{typ}
	ip.xtypes = []xr.Type{xtyp}
	ip.globals = globals
}

func (ip *Inspector) ShowHelp() {
	g := ip.globals
	g.Fprintf(g.Stdout, "%s", `
// inspector commands:
NUMBER      enter n-th struct field, or n-th element of array, slice or string
.           show current expression
?           show this help
help        show this help
methods     show methods
quit        exit inspector
top         return to top-level expression
up          return to outer expression
// abbreviations are allowed if unambiguous.
`)
}

func (ip *Inspector) Show() {
	depth := len(ip.names)
	name := strings.Join(ip.names, ".")
	v := ip.vals[depth-1]
	t := ip.types[depth-1]
	ip.showVar(name, v, t)

	v = dereferenceValue(v) // dereference pointers on-the-fly
	switch v.Kind() {
	case r.Array, r.Slice, r.String:
		ip.showIndexes(v)
	case r.Struct:
		ip.showFields(v)
	}
}

func (ip *Inspector) Repl() error {
	g := ip.globals
	g.Fprintf(g.Stdout, "%s", "// type ? for inspector help\n")
	for len(ip.names) > 0 {
		prompt := g.Sprintf("inspect %s> ", strings.Join(ip.names, "."))
		bytes, err := g.Readline.Read(prompt)
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
		ip.ShowHelp()
	case strings.HasPrefix("methods", cmd):
		t := ip.types[len(ip.types)-1]
		xt := ip.xtypes[len(ip.types)-1]
		ip.showMethods(t, xt)
	case strings.HasPrefix("quit", cmd):
		return errors.New("user quit")
	case strings.HasPrefix("top", cmd):
		ip.Top()
		ip.Show()
	case cmd == "", cmd == ".":
		ip.Show()
	case cmd == "-", strings.HasPrefix("up", cmd):
		if len(ip.names) > 1 {
			ip.Leave()
		} else {
			ip.Show()
		}
	default:
		ip.Enter(cmd)
	}
	return nil
}

func (ip *Inspector) Top() {
	ip.names = ip.names[0:1]
	ip.vals = ip.vals[0:1]
	ip.types = ip.types[0:1]
}

func (ip *Inspector) Leave() {
	depth := len(ip.names)
	if depth <= 0 {
		return
	}
	depth--
	ip.names = ip.names[:depth]
	ip.vals = ip.vals[:depth]
	ip.types = ip.types[:depth]
	if depth > 0 {
		ip.Show()
	}
}

func (ip *Inspector) showVar(str string, v r.Value, t r.Type) {
	ip.globals.Fprintf(ip.globals.Stdout, "%s\t= %v\t// %v\n", str, v, t)
}

func (ip *Inspector) showFields(v r.Value) {
	g := ip.globals
	n := v.NumField()
	for i := 0; i < n; i++ {
		f := v.Field(i)
		t := reflect.Type(f)
		f = dereferenceValue(f)
		g.Fprintf(g.Stdout, "    %d. ", i)
		ip.showVar(v.Type().Field(i).Name, f, t)
	}
}

func (ip *Inspector) showIndexes(v r.Value) {
	g := ip.globals
	n := v.Len()
	for i := 0; i < n; i++ {
		f := v.Index(i)
		t := reflect.Type(f)
		f = dereferenceValue(f)
		g.Fprintf(g.Stdout, "    %d. ", i)
		ip.showVar("", f, t)
	}
}

func (ip *Inspector) showMethods(t r.Type, xt xr.Type) {
	g := ip.globals
	switch {
	case xt != nil:
		if xt.Kind() == r.Ptr {
			xt = xt.Elem()
		}
		n := xt.NumMethod()
		if n == 0 {
			g.Fprintf(g.Stdout, "no methods of %v\n", xt)
			return
		}
		g.Fprintf(g.Stdout, "methods of %v:\n", xt)
		for i := 0; i < n; i++ {
			g.Fprintf(g.Stdout, "    m%d. %v\n", i, xt.Method(i).GoFun)
		}

	case t != nil:
		n := t.NumMethod()
		if n == 0 {
			g.Fprintf(g.Stdout, "no methods of %v\n", t)
			return
		}
		g.Fprintf(g.Stdout, "methods of %v:\n", t)
		for i := 0; i < n; i++ {
			m := t.Method(i)
			g.Fprintf(g.Stdout, "    m%d. %s\t%v\n", i, m.Name, m.Type)
		}
	}
}

func (ip *Inspector) Enter(cmd string) {
	g := ip.globals
	i, err := strconv.Atoi(cmd)
	if err != nil {
		g.Fprintf(g.Stdout, "unknown inspector command \"%s\". Type ? for help\n", cmd)
		return
	}
	depth := len(ip.names)
	v := dereferenceValue(ip.vals[depth-1])
	var n int
	var fname string
	var f r.Value
	switch v.Kind() {
	case r.Array, r.Slice, r.String:
		n = v.Len()
		if !ip.validRange(i, n) {
			return
		}
		fname = "[" + cmd + "]"
		f = v.Index(i)
	case r.Struct:
		n = v.NumField()
		if !ip.validRange(i, n) {
			return
		}
		fname = v.Type().Field(i).Name
		f = v.Field(i)
	default:
		g.Fprintf(g.Stdout, "cannot enter <%v>: expecting array, slice, string or struct\n", reflect.Type(v))
		return
	}
	var t r.Type
	if f.IsValid() && f != base.None {
		if f.Kind() == r.Interface {
			f = f.Elem() // concrete type
		}
		t = f.Type()
	}

	switch dereferenceValue(f).Kind() { // dereference pointers on-the-fly
	case r.Array, r.Slice, r.String, r.Struct:
		ip.names = append(ip.names, fname)
		ip.vals = append(ip.vals, f)
		ip.types = append(ip.types, t)
		ip.Show()
	default:
		ip.showVar(fname, f, t)
	}
}

func dereferenceValue(v r.Value) r.Value {
	for {
		switch v.Kind() {
		case r.Interface, r.Ptr:
			v = v.Elem()
			continue
		}
		break
	}
	return v
}

func (ip *Inspector) validRange(i, n int) bool {
	if i < 0 || i >= n {
		g := ip.globals
		g.Fprintf(g.Stdout, "%s contains %d elements, cannot inspect element %d\n",
			strings.Join(ip.names, "."), n, i)
		return false
	}
	return true
}
