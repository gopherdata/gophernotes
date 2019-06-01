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
 * type.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"fmt"
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/genimport"
	"github.com/cosmos72/gomacro/base/output"
	"github.com/cosmos72/gomacro/base/reflect"
)

func typeOf(value r.Value) r.Type {
	if value == None || value == Nil {
		return TypeOfInterface
	}
	return value.Type()
}

func (env *Env) evalExpr1OrType(node ast.Expr) (val r.Value, t r.Type) {
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case output.RuntimeError:
				t = env.evalType(node)
			default:
				panic(r)
			}
		}
	}()
	val = env.evalExpr1(node)
	return val, nil
}

// evalTypeAlias evaluates a type alias declaration, i.e. type Foo = /*...*/
func (env *Env) evalTypeAlias(name string, node ast.Expr) r.Type {
	t := env.evalType(node)
	// never define bindings for "_"
	if name != "_" {
		if _, ok := env.Types.Get(name); ok {
			env.Warnf("redefined type alias: %v", name)
		} else {
			env.Types.Ensure()
		}
		env.Types.Set(name, t)
	}
	return t
}

// evalType evaluates a type
func (env *Env) evalType(node ast.Expr) r.Type {
	t, _ := env.evalType2(node, false)
	return t
}

// evalTypeOrNil evaluates a type. as a special case used by type switch, evaluates *ast.Ident{Name:"nil"} to nil
func (env *Env) evalTypeOrNil(node ast.Expr) r.Type {
	for {
		switch expr := node.(type) {
		case *ast.ParenExpr:
			node = expr.X
			continue
		case *ast.Ident:
			if expr.Name == "nil" {
				val, found := env.resolveIdentifier(expr)
				if found && val == Nil {
					return nil
				}
			}
		}
		break
	}
	t, _ := env.evalType2(node, false)
	return t
}

// evalType0 evaluates a type expression.
// if allowEllipsis is true, it supports the special case &ast.Ellipsis{/*expression*/}
// that represents ellipsis in the last argument of a function declaration.
// The second return value is true both in the case above, and for array types whose length is [...]
func (env *Env) evalType2(node ast.Expr, allowEllipsis bool) (t r.Type, ellipsis bool) {
	stars := 0
	for {
		switch expr := node.(type) {
		case *ast.StarExpr:
			stars++
			node = expr.X
			continue
		case *ast.ParenExpr:
			node = expr.X
			continue
		case *ast.Ellipsis:
			if allowEllipsis {
				node = expr.Elt
				ellipsis = true
				continue
			}
		}
		break
	}
	if node != nil {
		env.Pos = node.Pos()
	}

	switch node := node.(type) {
	case *ast.ArrayType: // also for slices
		var ellipsis2 bool
		t, ellipsis2 = env.evalTypeArray(node)
		if !ellipsis {
			ellipsis = ellipsis2
		}
	case *ast.ChanType:
		t = env.evalType(node.Value)
		dir := r.BothDir
		if node.Dir == ast.SEND {
			dir = r.SendDir
		} else if node.Dir == ast.RECV {
			dir = r.RecvDir
		}
		t = r.ChanOf(dir, t)
	case *ast.FuncType:
		t, _, _ = env.evalTypeFunction(node)
	case *ast.Ident:
		t = env.evalTypeIdentifier(node.Name)
	case *ast.InterfaceType:
		t = env.evalTypeInterface(node)
	case *ast.MapType:
		kt := env.evalType(node.Key)
		vt := env.evalType(node.Value)
		t = r.MapOf(kt, vt)
	case *ast.SelectorExpr:
		if pkgIdent, ok := node.X.(*ast.Ident); ok {
			pkgv := env.evalIdentifier(pkgIdent)
			if pkg, ok := pkgv.Interface().(*genimport.PackageRef); ok {
				name := node.Sel.Name
				if t, ok = pkg.Types[name]; !ok {
					env.Errorf("not a type: %v <%v>", node, r.TypeOf(node))
				}
			} else {
				env.Errorf("not a package: %v = %v <%v>", pkgIdent, pkgv, typeOf(pkgv))
			}
		} else {
			env.Errorf("unimplemented qualified type, expecting packageName.identifier: %v <%v>", node, r.TypeOf(node))
		}
	case *ast.StructType:
		// env.Debugf("evalType() struct declaration: %v <%v>", node, r.TypeOf(node))
		types, names := env.evalTypeFields(node.Fields)
		// env.Debugf("evalType() struct names and types: %v %v", types, names)
		fields := makeStructFields(env.FileEnv().Path, names, types)
		// env.Debugf("evalType() struct fields: %#v", fields)
		t = r.StructOf(fields)
	case nil:
		// type can be omitted in many case - then we must perform type inference
		break
	default:
		// TODO which types are still missing?
		env.Errorf("unimplemented type: %v <%v>", node, r.TypeOf(node))
	}
	for i := 0; i < stars; i++ {
		t = r.PtrTo(t)
	}
	if allowEllipsis && ellipsis {
		t = r.SliceOf(t)
	}
	return t, ellipsis
}

func (env *Env) evalTypeArray(node *ast.ArrayType) (t r.Type, ellipsis bool) {
	t = env.evalType(node.Elt)
	n := node.Len
	switch n := n.(type) {
	case *ast.Ellipsis:
		t = r.SliceOf(t)
		ellipsis = true
	case nil:
		t = r.SliceOf(t)
	default:
		count := env.evalExpr1(n).Int()
		t = r.ArrayOf(int(count), t)
	}
	return t, ellipsis
}

func (env *Env) evalTypeFunction(node *ast.FuncType) (t r.Type, argNames []string, resultNames []string) {
	tFunc, _, argNames, resultNames := env.evalTypeFunctionOrMethod(nil, node)
	return tFunc, argNames, resultNames
}

func (env *Env) evalTypeFunctionOrMethod(recv *ast.Field, node *ast.FuncType) (tFunc r.Type, tFuncOrMethod r.Type, argNames []string, resultNames []string) {
	argTypes, argNames, variadic := env.evalTypeFieldOrParamList(node.Params, true)
	resultTypes, resultNames := env.evalTypeFields(node.Results)
	tFunc = r.FuncOf(argTypes, resultTypes, variadic)

	if recv != nil {
		recvTypes, recvNames, _ := env.evalTypeFieldsOrParams([]*ast.Field{recv}, false)
		argTypes = append(recvTypes, argTypes...)
		argNames = append(recvNames, argNames...)
		tFuncOrMethod = r.FuncOf(argTypes, resultTypes, variadic)
	} else {
		tFuncOrMethod = tFunc
	}
	return tFunc, tFuncOrMethod, argNames, resultNames
}

func (env *Env) evalTypeFields(fields *ast.FieldList) (types []r.Type, names []string) {
	types, names, _ = env.evalTypeFieldOrParamList(fields, false)
	return types, names
}

func (env *Env) evalTypeFieldOrParamList(fields *ast.FieldList, allowEllipsis bool) (types []r.Type, names []string, ellipsis bool) {
	var list []*ast.Field
	if fields != nil {
		list = fields.List
	}
	return env.evalTypeFieldsOrParams(list, allowEllipsis)
}

func (env *Env) evalTypeFieldsOrParams(list []*ast.Field, allowEllipsis bool) (types []r.Type, names []string, ellipsis bool) {
	types = make([]r.Type, 0)
	names = ZeroStrings
	n := len(list)
	if n == 0 {
		return types, names, ellipsis
	}
	var t r.Type
	for i, f := range list {
		t, ellipsis = env.evalType2(f.Type, i == n-1)
		if len(f.Names) == 0 {
			types = append(types, t)
			names = append(names, "_")
			// env.Debugf("evalTypeFields() %v -> %v", f.Type, t)
		} else {
			for _, ident := range f.Names {
				types = append(types, t)
				names = append(names, ident.Name)
				// Debugf("evalTypeFields() %v %v -> %v", ident.Name, f.Type, t)
			}
		}
	}
	return types, names, ellipsis
}

func (env *Env) evalTypeIdentifier(name string) r.Type {
	for e := env; e != nil; e = e.Outer {
		if t, ok := e.Types.Get(name); ok {
			return t
		}
	}
	env.Errorf("undefined identifier: %v", name)
	return nil
}

func makeStructFields(pkgPath string, names []string, types []r.Type) []r.StructField {
	// pkgIdentifier := sanitizeIdentifier(pkgPath)
	fields := make([]r.StructField, len(names))
	for i, name := range names {
		fields[i] = r.StructField{
			Name:      toExportedName(name), // Go 1.8 reflect.StructOf() supports *only* exported fields
			Type:      types[i],
			Tag:       "",
			Anonymous: false,
		}
	}
	return fields
}

func toExportedName(name string) string {
	if len(name) == 0 {
		return name
	}
	ch := name[0]
	if ch >= 'a' && ch <= 'z' {
		ch -= 'a' - 'A'
	} else if ch == '_' {
		ch = 'X'
	} else {
		return name
	}
	return fmt.Sprintf("%c%s", ch, name[1:])
}

func (env *Env) valueToType(value r.Value, t r.Type) r.Value {
	if value == None || value == Nil {
		switch t.Kind() {
		case r.Chan, r.Func, r.Interface, r.Map, r.Ptr, r.Slice:
			return r.Zero(t)
		}
	}
	newValue := reflect.ConvertValue(value, t)
	if differentIntegerValues(value, newValue) {
		env.Warnf("value %d overflows <%v>, truncated to %d", value, t, newValue)
	}
	return newValue
}

func differentIntegerValues(v1 r.Value, v2 r.Value) bool {
	k1, k2 := v1.Kind(), v2.Kind()
	switch k1 {
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		n1 := v1.Int()
		switch k2 {
		case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
			return n1 != v2.Int()
		case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
			return n1 < 0 || uint64(n1) != v2.Uint()
		default:
			return false
		}
	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		n1 := v1.Uint()
		switch k2 {
		case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
			n2 := v2.Int()
			return n2 < 0 || uint64(n2) != n1
		case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
			return n1 != v2.Uint()
		default:
			return false
		}
	default:
		return false
	}
}

func toValues(args []interface{}) []r.Value {
	n := len(args)
	values := make([]r.Value, n)
	for i := 0; i < n; i++ {
		values[i] = r.ValueOf(args[i])
	}
	return values
}

func toInterfaces(values []r.Value) []interface{} {
	n := len(values)
	rets := make([]interface{}, n)
	for i := 0; i < n; i++ {
		rets[i] = toInterface(values[i])
	}
	return rets
}

func toInterface(value r.Value) interface{} {
	if value != Nil && value != None {
		return value.Interface()
	}
	return nil
}
