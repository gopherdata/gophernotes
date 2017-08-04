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
 * type.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"fmt"
	"go/ast"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// DeclType compiles a type declaration.
func (c *Comp) DeclType(node ast.Spec) {
	switch node := node.(type) {
	case *ast.TypeSpec:
		name := node.Name.Name
		// PATCH: support type aliases
		if unary, ok := node.Type.(*ast.UnaryExpr); ok && unary.Op == token.ASSIGN {
			t := c.Type(unary.X)
			c.DeclTypeAlias(name, t)
		} else {
			// support self-referencing types, as for example: type List struct { First int; Rest *List }
			t := c.DeclNamedType(name)
			u := c.Type(node.Type)
			if t != nil { // t == nil means name == "_", discard the result of type declaration
				c.SetUnderlyingType(t, u)
			}
		}
	default:
		c.Errorf("Compile: unexpected type declaration, expecting <*ast.TypeSpec>, found: %v <%v>", node, r.TypeOf(node))
	}
}

// DeclTypeAlias compiles a typealias declaration, i.e. type Foo = /*...*/
// Returns the second argument.
func (c *Comp) DeclTypeAlias(name string, t xr.Type) xr.Type {
	if name == "_" {
		return t
	}
	if _, ok := c.Types[name]; ok {
		c.Warnf("redefined type alias: %v", name)
		c.Universe.InvalidateCache()
	} else if c.Types == nil {
		c.Types = make(map[string]xr.Type)
	}
	c.Types[name] = t
	return t
}

// DeclNamedType executes a named type forward declaration.
// Returns nil if name == "_"
// Otherwise it must be followed by Comp.SetUnderlyingType()
func (c *Comp) DeclNamedType(name string) xr.Type {
	if name == "_" {
		return nil
	}
	if t, ok := c.Types[name]; ok {
		c.Warnf("redefined type: %v", name)
		if xr.QName1(t) != xr.QName2(name, c.FileComp().Path) {
			// the current type "name" is an alias, discard it
			c.Universe.InvalidateCache()
		} else {
			// reuse t, change only its underlying type
			return t
		}
	} else if c.Types == nil {
		c.Types = make(map[string]xr.Type)
	}
	t := c.Universe.NamedOf(name, c.FileComp().Path)
	c.Types[name] = t
	return t
}

func (c *Comp) SetUnderlyingType(t, underlying xr.Type) {
	t.SetUnderlying(underlying)
}

// DeclType0 declares a type
// in Go, types are computed only at compile time - no need for a runtime *Env
func (c *Comp) DeclType0(t xr.Type) xr.Type {
	if t == nil {
		return nil
	}
	return c.DeclTypeAlias0(t.Name(), t)
}

// DeclTypeAlias0 declares a type alias
// in Go, types are computed only at compile time - no need for a runtime *Env
func (c *Comp) DeclTypeAlias0(alias string, t xr.Type) xr.Type {
	if alias == "" || alias == "_" {
		// never define bindings for "_"
		return t
	}
	if _, ok := c.Types[alias]; ok {
		c.Warnf("redefined type: %v", alias)
	} else if c.Types == nil {
		c.Types = make(map[string]xr.Type)
	}
	c.Types[alias] = t
	return t
}

// Type compiles a type expression.
func (c *Comp) Type(node ast.Expr) xr.Type {
	t, _ := c.compileType2(node, false)
	return t
}

// compileTypeOrNil compiles a type expression. as a special case used by type switch, compiles *ast.Ident{Name:"nil"} to nil
func (c *Comp) compileTypeOrNil(node ast.Expr) xr.Type {
	for {
		switch expr := node.(type) {
		case *ast.ParenExpr:
			node = expr.X
			continue
		case *ast.Ident:
			if expr.Name == "nil" {
				sym := c.TryResolve(expr.Name)
				if sym != nil && sym.Type == nil {
					return nil
				}
			}
		}
		break
	}
	t, _ := c.compileType2(node, false)
	return t
}

// compileType2 compiles a type expression.
// if allowEllipsis is true, it supports the special case &ast.Ellipsis{/*expression*/}
// that represents ellipsis in the last argument of a function declaration.
// The second return value is true both in the case above, and for array types whose length is [...]
func (c *Comp) compileType2(node ast.Expr, allowEllipsis bool) (t xr.Type, ellipsis bool) {
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
		c.Pos = node.Pos()
	}
	universe := c.Universe
	var ellipsisArray bool

	switch node := node.(type) {
	case *ast.ArrayType: // also for slices
		t, ellipsisArray = c.TypeArray(node)
	case *ast.ChanType:
		telem := c.Type(node.Value)
		dir := r.BothDir
		if node.Dir == ast.SEND {
			dir = r.SendDir
		} else if node.Dir == ast.RECV {
			dir = r.RecvDir
		}
		t = universe.ChanOf(dir, telem)
	case *ast.FuncType:
		t, _, _ = c.TypeFunction(node)
	case *ast.Ident:
		t = c.TypeIdent(node.Name)
	case *ast.InterfaceType:
		t = c.TypeInterface(node)
	case *ast.MapType:
		kt := c.Type(node.Key)
		vt := c.Type(node.Value)
		t = universe.MapOf(kt, vt)
	case *ast.SelectorExpr:
		ident, ok := node.X.(*ast.Ident)
		if !ok {
			c.Errorf("invalid qualified type, expecting packagename.identifier, found: %v <%v>", node, r.TypeOf(node))
		}
		// this could be Package.Type, or other non-type expressions: Type.Method, Value.Method, Struct.Field...
		// check for Package.Type
		name := ident.Name
		var bind *Bind
		for o := c; o != nil; o = o.Outer {
			if bind = o.Binds[name]; bind != nil {
				break
			}
		}
		if bind == nil {
			c.Errorf("undefined %q in %v <%v>", name, node, r.TypeOf(node))
		} else if !bind.Const() || bind.Type.ReflectType() != rtypeOfImport {
			c.Errorf("not a package: %q in %v <%v>", name, node, r.TypeOf(node))
		}
		imp, ok := bind.Value.(Import)
		if !ok {
			c.Errorf("not a package: %q in %v <%v>", name, node, r.TypeOf(node))
		}
		name = node.Sel.Name
		t, ok = imp.Types[name]
		if !ok || t == nil {
			c.Errorf("not a type: %v <%v>", node, r.TypeOf(node))
		}
		if !ast.IsExported(name) {
			c.Errorf("cannot refer to unexported name %v", node)
		}
	case *ast.StructType:
		// c.Debugf("evalType() struct declaration: %v <%v>", node, r.TypeOf(node))
		types, names := c.TypeFields(node.Fields)
		// c.Debugf("evalType() struct names and types: %v %v", types, names)
		pkg := universe.LoadPackage(c.FileComp().Path)
		fields := c.makeStructFields(pkg, names, types)
		// c.Debugf("compileType2() declaring struct type. fields=%#v", fields)
		t = universe.StructOf(fields)
	case nil:
		// type can be omitted in many case - then we must perform type inference
		break
	default:
		// TODO which types are still missing?
		c.Errorf("unimplemented type: %v <%v>", node, r.TypeOf(node))
	}
	if t != nil {
		for i := 0; i < stars; i++ {
			t = universe.PtrTo(t)
		}
		if allowEllipsis && ellipsis {
			// ellipsis in the last argument of a function declaration
			t = universe.SliceOf(t)
		}
	}
	return t, ellipsis || ellipsisArray
}

func (c *Comp) TypeArray(node *ast.ArrayType) (t xr.Type, ellipsis bool) {
	universe := c.Universe
	t = c.Type(node.Elt)
	n := node.Len
	switch n := n.(type) {
	case *ast.Ellipsis:
		t = universe.ArrayOf(0, t)
		ellipsis = true
	case nil:
		t = universe.SliceOf(t)
	default:
		// as stated by https://golang.org/ref/spec#Array_types
		// "The length is part of the array's type; it must evaluate to a non-negative constant
		// representable by a value of type int. "
		var count int
		init := c.Expr(n)
		if !init.Const() {
			c.Errorf("array length is not a constant: %v", node)
			return
		} else if init.Untyped() {
			count = init.ConstTo(c.TypeOfInt()).(int)
		} else {
			count = convertLiteralCheckOverflow(init.Value, c.TypeOfInt()).(int)
		}
		if count < 0 {
			c.Errorf("array length [%v] is negative: %v", count, node)
		}
		t = universe.ArrayOf(count, t)
	}
	return t, ellipsis
}

func (c *Comp) TypeFunction(node *ast.FuncType) (t xr.Type, paramNames []string, resultNames []string) {
	return c.TypeFunctionOrMethod(nil, node)
}

// TypeFunctionOrMethod compiles a function type corresponding to given receiver and function declaration
// If receiver is not null, the returned tFunc will have it as receiver.
func (c *Comp) TypeFunctionOrMethod(recv *ast.Field, node *ast.FuncType) (t xr.Type, paramNames []string, resultNames []string) {
	paramTypes, paramNames, variadic := c.typeFieldOrParamList(node.Params, true)
	resultTypes, resultNames := c.TypeFields(node.Results)

	var recvType xr.Type
	if recv != nil {
		// methods are functions with receiver. xreflect allows functions to be treated as methods
		// (using the first parameter as receiver), but go/types.Type loaded by go/importer.Default()
		// will have methods as functions with receivers.
		//
		// So be uniform with those.
		//
		// Alas, go/types.Type.String() does *not* print the receiver, making it cumbersome to debug.
		recvTypes, recvNames, _ := c.typeFieldsOrParams([]*ast.Field{recv}, false)
		recvType = recvTypes[0]

		// anyway, return the receiver *name* as first element of paramNames
		paramNames = append(recvNames, paramNames...)
	}
	t = c.Universe.MethodOf(recvType, paramTypes, resultTypes, variadic)
	return t, paramNames, resultNames
}

func (c *Comp) TypeFields(fields *ast.FieldList) (types []xr.Type, names []string) {
	types, names, _ = c.typeFieldOrParamList(fields, false)
	return types, names
}

func (c *Comp) typeFieldOrParamList(fields *ast.FieldList, allowEllipsis bool) (types []xr.Type, names []string, ellipsis bool) {
	var list []*ast.Field
	if fields != nil {
		list = fields.List
	}
	return c.typeFieldsOrParams(list, allowEllipsis)
}

func (c *Comp) typeFieldsOrParams(list []*ast.Field, allowEllipsis bool) (types []xr.Type, names []string, ellipsis bool) {
	types = make([]xr.Type, 0)
	names = ZeroStrings
	n := len(list)
	if n == 0 {
		return types, names, ellipsis
	}
	var t xr.Type
	for i, f := range list {
		t, ellipsis = c.compileType2(f.Type, i == n-1)
		if len(f.Names) == 0 {
			types = append(types, t)
			names = append(names, "")
			// c.Debugf("evalTypeFields() %v -> %v", f.Type, t)
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

func (c *Comp) TypeIdent(name string) xr.Type {
	for co := c; co != nil; co = co.Outer {
		if t, ok := co.Types[name]; ok {
			return t
		}
	}
	c.Errorf("undefined identifier: %v", name)
	return nil
}

func (c *Comp) makeStructFields(pkg *xr.Package, names []string, types []xr.Type) []xr.StructField {
	// pkgIdentifier := sanitizeIdentifier(pkgPath)
	fields := make([]xr.StructField, len(names))
	for i, name := range names {
		t := types[i]
		fields[i] = xr.StructField{
			Name:      name,
			Pkg:       pkg,
			Type:      t,
			Tag:       "",
			Anonymous: len(name) == 0,
		}
	}
	return fields
}

// TypeAssert2 compiles a multi-valued type assertion
func (c *Comp) TypeAssert2(node *ast.TypeAssertExpr) *Expr {
	val := c.Expr1(node.X)
	tin := val.Type
	tout := c.Type(node.Type)
	rtout := tout.ReflectType()
	if tin == nil || tin.Kind() != r.Interface {
		c.Errorf("invalid type assertion: %v (non-interface type <%v> on left)", node, tin)
		return nil
	}
	kin := tin.Kind()
	kout := tout.Kind()
	if kout != r.Interface && !tout.Implements(tin) {
		c.Errorf("impossible type assertion: <%v> does not implement <%v>", tout, tin)
	}
	fun := val.Fun.(func(*Env) r.Value)     // val returns an interface... must be already wrapped in a reflect.Value
	fail := []r.Value{xr.Zero(tout), False} // returned by type assertion in case of failure

	var ret func(env *Env) (r.Value, []r.Value)

	if IsOptimizedKind(kout) {
		ret = func(env *Env) (r.Value, []r.Value) {
			v := fun(env)
			v = r.ValueOf(v.Interface()) // rebuild reflect.Value with concrete type
			if v.Type() != rtout {
				return fail[0], fail
			}
			return v, []r.Value{v, True}
		}
	} else if tout.ReflectType() == TypeOfInterface {
		// special case, nil is a valid interface{}
		ret = func(env *Env) (r.Value, []r.Value) {
			v := fun(env).Convert(TypeOfInterface)
			return v, []r.Value{v, True}
		}
	} else if kout == r.Interface && tin.Implements(tout) {
		ret = func(env *Env) (r.Value, []r.Value) {
			v := fun(env)
			// nil is not a valid tout, check for it.
			// IsNil() can be invoked only on nillable types...
			// but v.Type().Kind() should be r.Interface, which is nillable :)
			if v.IsNil() {
				return fail[0], fail
			}
			v = v.Convert(rtout)
			return v, []r.Value{v, True}
		}
	} else if kout == r.Interface {
		ret = func(env *Env) (r.Value, []r.Value) {
			v := fun(env)
			// nil is not a valid tout, check for it.
			// IsNil() can be invoked only on nillable types...
			// but v.Type().Kind() should be r.Interface, which is nillable :)
			if v.IsNil() {
				return fail[0], fail
			}
			v = r.ValueOf(v.Interface()) // rebuild reflect.Value with concrete type
			rtconcr := v.Type()
			if rtconcr != rtout && !rtconcr.Implements(rtout) {
				return fail[0], fail
			}
			v = v.Convert(rtout)
			return v, []r.Value{v, True}
		}
	} else if IsNillableKind(kin) {
		ret = func(env *Env) (r.Value, []r.Value) {
			v := fun(env)
			// nil is not a valid tout, check for it.
			// IsNil() can be invoked only on nillable types...
			// but we just checked IsNillableKind(kin)
			if v.IsNil() {
				return fail[0], fail
			}
			v = r.ValueOf(v.Interface()) // rebuild reflect.Value with concrete type
			rtconcr := v.Type()
			if rtconcr != rtout {
				return fail[0], fail
			}
			return v, []r.Value{v, True}
		}
	} else {
		ret = func(env *Env) (r.Value, []r.Value) {
			v := fun(env)
			v = r.ValueOf(v.Interface()) // rebuild reflect.Value with concrete type
			rtconcr := v.Type()
			if rtconcr != rtout {
				return fail[0], fail
			}
			return v, []r.Value{v, True}
		}
	}
	return exprXV([]xr.Type{tout, c.TypeOfBool()}, ret)
}

// TypeAssert1 compiles a single-valued type assertion
func (c *Comp) TypeAssert1(node *ast.TypeAssertExpr) *Expr {
	if node.Type == nil {
		c.Errorf("invalid type assertion: expecting actual type, found type switch: %v", node)
	}
	val := c.Expr1(node.X)
	tin := val.Type
	tout := c.Type(node.Type)
	kout := tout.Kind()
	if tin == nil || tin.Kind() != r.Interface {
		c.Errorf("invalid type assertion: %v (non-interface type <%v> on left)", node, tin)
		return nil
	}
	if tout.Kind() != r.Interface && !tout.Implements(tin) {
		c.Errorf("impossible type assertion: <%v> does not implement <%v>", tout, tin)
	}
	fun := val.Fun.(func(*Env) r.Value) // val returns an interface... must be already wrapped in a reflect.Value
	rtin := tin.ReflectType()
	rtout := tout.ReflectType()
	var ret I
	switch kout {
	case r.Bool:
		ret = func(env *Env) bool {
			v := typeassert(fun(env), rtin, rtout)
			return v.Bool()
		}
	case r.Int:
		ret = func(env *Env) int {
			v := typeassert(fun(env), rtin, rtout)
			return int(v.Int())
		}
	case r.Int8:
		ret = func(env *Env) int8 {
			v := typeassert(fun(env), rtin, rtout)
			return int8(v.Int())
		}
	case r.Int16:
		ret = func(env *Env) int16 {
			v := typeassert(fun(env), rtin, rtout)
			return int16(v.Int())
		}
	case r.Int32:
		ret = func(env *Env) int32 {
			v := typeassert(fun(env), rtin, rtout)
			return int32(v.Int())
		}
	case r.Int64:
		ret = func(env *Env) int64 {
			v := typeassert(fun(env), rtin, rtout)
			return v.Int()
		}
	case r.Uint:
		ret = func(env *Env) uint {
			v := typeassert(fun(env), rtin, rtout)
			return uint(v.Uint())
		}
	case r.Uint8:
		ret = func(env *Env) uint8 {
			v := typeassert(fun(env), rtin, rtout)
			return uint8(v.Uint())
		}
	case r.Uint16:
		ret = func(env *Env) uint16 {
			v := typeassert(fun(env), rtin, rtout)
			return uint16(v.Uint())
		}
	case r.Uint32:
		ret = func(env *Env) uint32 {
			v := typeassert(fun(env), rtin, rtout)
			return uint32(v.Uint())
		}
	case r.Uint64:
		ret = func(env *Env) uint64 {
			v := typeassert(fun(env), rtin, rtout)
			return v.Uint()
		}
	case r.Uintptr:
		ret = func(env *Env) uintptr {
			v := typeassert(fun(env), rtin, rtout)
			return uintptr(v.Uint())
		}
	case r.Float32:
		ret = func(env *Env) float32 {
			v := typeassert(fun(env), rtin, rtout)
			return float32(v.Float())
		}
	case r.Float64:
		ret = func(env *Env) float64 {
			v := typeassert(fun(env), rtin, rtout)
			return v.Float()
		}
	case r.Complex64:
		ret = func(env *Env) complex64 {
			v := typeassert(fun(env), rtin, rtout)
			return complex64(v.Complex())
		}
	case r.Complex128:
		ret = func(env *Env) complex128 {
			v := typeassert(fun(env), rtin, rtout)
			return v.Convert(rtout).Complex()
		}
	case r.String:
		ret = func(env *Env) string {
			v := typeassert(fun(env), rtin, rtout)
			return v.String()
		}
	default:
		if tout.ReflectType() == TypeOfInterface {
			// special case, nil is a valid interface{}
			ret = func(env *Env) r.Value {
				return fun(env).Convert(TypeOfInterface)
			}
			break
		}
		if tout.Kind() == r.Interface && tin.Implements(tout) {
			ret = func(env *Env) r.Value {
				v := fun(env)
				// nil is not a valid tout, check for it.
				// IsNil() can be invoked only on nillable types...
				// but v.Type().Kind() should be r.Interface, which is nillable :)
				if v.IsNil() {
					panic(&TypeAssertionError{
						Interface: rtin,
						Concrete:  nil,
						Asserted:  rtout,
					})
				}
				return v.Convert(rtout)
			}
			break
		}
		ret = func(env *Env) r.Value {
			v := fun(env)
			// nil is not a valid tout, check for it.
			// IsNil() can be invoked only on nillable types...
			// but v.Type().Kind() should be r.Interface, which is nillable :)
			if v.IsNil() {
				panic(&TypeAssertionError{
					Interface: rtin,
					Concrete:  nil,
					Asserted:  rtout,
				})
			}
			v = r.ValueOf(v.Interface()) // rebuild reflect.Value with concrete type
			rtconcr := v.Type()
			if rtconcr != rtout && (rtout.Kind() != r.Interface || !rtconcr.Implements(rtout)) {
				panic(&TypeAssertionError{
					Interface: rtin,
					Concrete:  rtconcr,
					Asserted:  rtout,
				})
			}
			return v.Convert(rtout)
		}
	}
	return exprFun(tout, ret)
}

func typeassert(v r.Value, rtin r.Type, rtout r.Type) r.Value {
	v = r.ValueOf(v.Interface()) // extract concrete type
	rtconcr := v.Type()
	if rtconcr != rtout {
		panic(&TypeAssertionError{
			Interface: rtin,
			Concrete:  rtconcr,
			Asserted:  rtout,
		})
	}
	return v
}

func (g *CompThreadGlobals) TypeOfBool() xr.Type {
	return g.Universe.BasicTypes[r.Bool]
}

func (g *CompThreadGlobals) TypeOfInt() xr.Type {
	return g.Universe.BasicTypes[r.Int]
}

func (g *CompThreadGlobals) TypeOfInt8() xr.Type {
	return g.Universe.BasicTypes[r.Int8]
}

func (g *CompThreadGlobals) TypeOfInt16() xr.Type {
	return g.Universe.BasicTypes[r.Int16]
}

func (g *CompThreadGlobals) TypeOfInt32() xr.Type {
	return g.Universe.BasicTypes[r.Int32]
}

func (g *CompThreadGlobals) TypeOfInt64() xr.Type {
	return g.Universe.BasicTypes[r.Int64]
}

func (g *CompThreadGlobals) TypeOfUint() xr.Type {
	return g.Universe.BasicTypes[r.Uint]
}

func (g *CompThreadGlobals) TypeOfUint8() xr.Type {
	return g.Universe.BasicTypes[r.Uint8]
}

func (g *CompThreadGlobals) TypeOfUint16() xr.Type {
	return g.Universe.BasicTypes[r.Uint16]
}

func (g *CompThreadGlobals) TypeOfUint32() xr.Type {
	return g.Universe.BasicTypes[r.Uint32]
}

func (g *CompThreadGlobals) TypeOfUint64() xr.Type {
	return g.Universe.BasicTypes[r.Uint64]
}

func (g *CompThreadGlobals) TypeOfUintptr() xr.Type {
	return g.Universe.BasicTypes[r.Uintptr]
}

func (g *CompThreadGlobals) TypeOfFloat32() xr.Type {
	return g.Universe.BasicTypes[r.Float32]
}

func (g *CompThreadGlobals) TypeOfFloat64() xr.Type {
	return g.Universe.BasicTypes[r.Float64]
}

func (g *CompThreadGlobals) TypeOfComplex64() xr.Type {
	return g.Universe.BasicTypes[r.Complex64]
}

func (g *CompThreadGlobals) TypeOfComplex128() xr.Type {
	return g.Universe.BasicTypes[r.Complex128]
}

func (g *CompThreadGlobals) TypeOfString() xr.Type {
	return g.Universe.BasicTypes[r.String]
}

func (g *CompThreadGlobals) TypeOfError() xr.Type {
	return g.Universe.TypeOfError
}

func (g *CompThreadGlobals) TypeOfInterface() xr.Type {
	return g.Universe.TypeOfInterface
}

var (
	rtypeOfBuiltin     = r.TypeOf(Builtin{})
	rtypeOfFunction    = r.TypeOf(Function{})
	rtypeOfImport      = r.TypeOf(Import{})
	rtypeOfMacro       = r.TypeOf(Macro{})
	rtypeOfUntypedLit  = r.TypeOf(UntypedLit{})
	rtypeOfReflectType = r.TypeOf((*r.Type)(nil)).Elem()

	zeroOfReflectType = r.Zero(rtypeOfReflectType)
)

func (g *CompThreadGlobals) TypeOfBuiltin() xr.Type {
	return g.Universe.ReflectTypes[rtypeOfBuiltin]
}

func (g *CompThreadGlobals) TypeOfFunction() xr.Type {
	return g.Universe.ReflectTypes[rtypeOfFunction]
}

func (g *CompThreadGlobals) TypeOfImport() xr.Type {
	return g.Universe.ReflectTypes[rtypeOfImport]
}

func (g *CompThreadGlobals) TypeOfMacro() xr.Type {
	return g.Universe.ReflectTypes[rtypeOfMacro]
}

func (g *CompThreadGlobals) TypeOfUntypedLit() xr.Type {
	return g.Universe.ReflectTypes[rtypeOfUntypedLit]
}

// A TypeAssertionError explains a failed type assertion.
type TypeAssertionError struct {
	Interface     r.Type
	Concrete      r.Type
	Asserted      r.Type
	MissingMethod string // one method needed by Interface, missing from Concrete
}

func (*TypeAssertionError) RuntimeError() {}

func (e *TypeAssertionError) Error() string {
	in := e.Interface
	concr := e.Concrete
	if concr == nil {
		return fmt.Sprintf("interface conversion: <%v> is nil, not <%v>", in, e.Asserted)
	}
	if len(e.MissingMethod) == 0 {
		return fmt.Sprintf("interface conversion: <%v> is <%v>, not <%v>", in, concr, e.Asserted)
	}
	return fmt.Sprintf("interface conversion: <%v> is not <%v>: missing method ", concr, e.Asserted, e.MissingMethod)
}
