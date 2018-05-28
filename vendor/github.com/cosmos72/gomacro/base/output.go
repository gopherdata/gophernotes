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
 * output.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"io"
	r "reflect"
	"strings"
	"unsafe"

	. "github.com/cosmos72/gomacro/ast2"
	mt "github.com/cosmos72/gomacro/token"
)

type Stringer struct {
	Fileset    *mt.FileSet
	Pos        token.Pos
	Line       int
	NamedTypes map[r.Type]string
}

type Output struct {
	Stringer
	Stdout io.Writer
	Stderr io.Writer
}

type RuntimeError struct {
	st     *Stringer
	format string
	args   []interface{}
}

func (st *Stringer) Copy(other *Stringer) {
	st.Fileset = other.Fileset
	st.Pos = other.Pos
	st.Line = other.Line
}

func (err RuntimeError) Error() string {
	args := err.args
	var prefix string
	if st := err.st; st != nil {
		args = st.toPrintables(err.format, args)
		prefix = st.Position().String()
	}
	msg := fmt.Sprintf(err.format, args...)
	if prefix != "" && prefix != "-" {
		msg = fmt.Sprintf("%s: %s", prefix, msg)
	}
	return msg
}

func makeRuntimeError(format string, args ...interface{}) error {
	return RuntimeError{nil, format, args}
}

func (st *Stringer) MakeRuntimeError(format string, args ...interface{}) RuntimeError {
	return RuntimeError{st, format, args}
}

func Error(err error) interface{} {
	panic(err)
}

func (o *Output) Error(err error) interface{} {
	panic(err)
}

func Errorf(format string, args ...interface{}) {
	panic(RuntimeError{nil, format, args})
}

func (st *Stringer) Errorf(format string, args ...interface{}) (r.Value, []r.Value) {
	panic(RuntimeError{st, format, args})
}

func Warnf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	fmt.Printf("// warning: %s\n", str)
}

func (o *Output) Warnf(format string, args ...interface{}) {
	args = o.toPrintables(format, args)
	str := fmt.Sprintf(format, args...)
	fmt.Fprintf(o.Stderr, "// warning: %s\n", str)
}

var warnExtraValues = 5

func (o *Output) WarnExtraValues(extraValues []r.Value) {
	if warnExtraValues > 0 {
		o.Warnf("expression returned %d values, using only the first one: %v",
			len(extraValues), extraValues)
		warnExtraValues--
		if warnExtraValues == 0 {
			o.Warnf("suppressing further similar warnings")
		}
	}
}

func Debugf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	fmt.Printf("// debug: %s\n", str)
}

func (o *Output) Debugf(format string, args ...interface{}) {
	args = o.toPrintables(format, args)
	str := fmt.Sprintf(format, args...)
	fmt.Fprintf(o.Stdout, "// debug: %s\n", str)
}

func (st *Stringer) IncLine(src string) {
	st.Line += strings.Count(src, "\n")
}

func (st *Stringer) IncLineBytes(src []byte) {
	st.Line += bytes.Count(src, []byte("\n"))
}

func (st *Stringer) Position() token.Position {
	if st == nil || st.Fileset == nil {
		return token.Position{}
	}
	return st.Fileset.Position(st.Pos)
}

func (ref *PackageRef) String() string {
	return fmt.Sprintf("{%s %q, %d binds, %d types}", ref.Name, ref.Path, len(ref.Binds), len(ref.Types))
}

func ShowPackageHeader(out io.Writer, name string, path string, kind string) {
	if name == path {
		fmt.Fprintf(out, "// ----- %s %s -----\n", name, kind)
	} else if name == FileName(path) {
		fmt.Fprintf(out, "// ----- %q %s -----\n", path, kind)
	} else {
		fmt.Fprintf(out, "// ----- %s %q %s -----\n", name, path, kind)
	}
}

var typeOfReflectValue = r.TypeOf(r.Value{})

type unsafeType struct {
}

type unsafeFlag uintptr

type unsafeValue struct {
	typ *unsafeType
	ptr unsafe.Pointer
	unsafeFlag
}

func asUnsafeValue(v r.Value) unsafeValue {
	return *(*unsafeValue)(unsafe.Pointer(&v))
}

func (st *Stringer) Fprintf(out io.Writer, format string, values ...interface{}) (n int, err error) {
	values = st.toPrintables(format, values)
	return fmt.Fprintf(out, format, values...)
}

func (st *Stringer) Sprintf(format string, values ...interface{}) string {
	values = st.toPrintables(format, values)
	return fmt.Sprintf(format, values...)
}

func (st *Stringer) ToString(separator string, values ...interface{}) string {
	if len(values) == 0 {
		return ""
	}
	values = st.toPrintables("", values)
	var buf bytes.Buffer
	for i, value := range values {
		if i != 0 {
			buf.WriteString(separator)
		}
		fmt.Fprint(&buf, value)
	}
	return buf.String()
}

func (st *Stringer) toPrintables(format string, values []interface{}) []interface{} {
	rets := make([]interface{}, len(values))
	for i, vi := range values {
		if percent := strings.IndexByte(format, '%'); percent >= 0 {
			format = format[percent:]
		}
		rets[i] = st.toPrintable(format, vi)
		switch len(format) {
		case 0:
		case 1, 2:
			format = ""
		default:
			format = format[2:] // skip %*
		}
	}
	return rets
}

func (st *Stringer) toPrintable(format string, value interface{}) (ret interface{}) {
	if value == nil {
		return nil
	}
	defer func() {
		if rec := recover(); rec != nil {
			ret = fmt.Sprintf("error pretty-printing %v", value)
		}
	}()

	switch v := value.(type) {
	case r.Value:
		return st.rvalueToPrintable(format, v)
	case fmt.Formatter:
		return v
	case fmt.GoStringer:
		if strings.HasPrefix(format, "%#v") {
			return v.GoString()
		}
	}

	usual := len(format) == 0 || strings.HasPrefix(format, "%v") || strings.HasPrefix(format, "%s")
	if usual {
		switch v := value.(type) {
		case AstWithNode:
			return st.nodeToPrintable(v.Node())
		case Ast:
			return st.toPrintable(format, v.Interface())
		case ast.Node:
			return st.nodeToPrintable(v)
		case r.Type:
			return st.typeToPrintable(v)
		case error:
			return v.Error()
		case fmt.Stringer:
			return v.String()
		}
	}

	v := r.ValueOf(value)
	switch k := v.Kind(); k {
	case r.Array, r.Slice:
		n := v.Len()
		values := make([]interface{}, n)
		converted := false
		for i := 0; i < n; i++ {
			vi := v.Index(i)
			if vi == Nil {
				values[i] = nil
			} else if !vi.CanInterface() {
				values[i] = vi
			} else {
				valuei := vi.Interface()
				values[i] = st.toPrintable(format, valuei)
				converted = converted || !vi.Type().Comparable() || valuei != values[i]
			}
		}
		// return []interface{} only if we actually converted some element
		if converted {
			return values
		} else {
			return value
		}
	case r.Struct:
		if usual {
			return st.structToPrintable(format, v)
		}
	case r.Func:
		return asUnsafeValue(v).ptr
	}
	return value
}

var config = printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}

func (st *Stringer) nodeToPrintable(node ast.Node) interface{} {
	if node == nil {
		return nil
	}
	var fset *mt.FileSet
	if st != nil {
		fset = st.Fileset
	}
	if fset == nil {
		fset = mt.NewFileSet()
	}
	var buf bytes.Buffer
	err := config.Fprint(&buf, &fset.FileSet, node)
	if err != nil {
		return err
	}
	return buf.String()
}

func (st *Stringer) rvalueToPrintable(format string, value r.Value) interface{} {
	var i interface{}
	if value == None {
		i = "/*no value*/"
	} else if value == Nil {
		i = nil
	} else if value.CanInterface() {
		i = st.toPrintable(format, value.Interface())
	} else {
		i = value
	}
	return i
}

func (st *Stringer) typeToPrintable(t r.Type) interface{} {
	if t == nil {
		return "nil" // because fmt.Printf("%v", nil) prints <nil> i.e adds extra <>
	}
	if st != nil {
		if name, ok := st.NamedTypes[t]; ok {
			return name
		}
	}
	return t
}

func (st *Stringer) structToPrintable(format string, v r.Value) string {
	n := v.NumField()
	if n == 0 {
		return "{}"
	}
	var buf bytes.Buffer
	t := v.Type()
	ch := '{'
	for i := 0; i < n; i++ {
		fmt.Fprintf(&buf, "%c%s:%v", ch, t.Field(i).Name, v.Field(i))
		ch = ' '
	}
	buf.WriteByte('}')
	return buf.String()
}
