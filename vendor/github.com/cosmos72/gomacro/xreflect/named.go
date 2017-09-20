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
 * named.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/ast"
	"go/token"
	"go/types"
	"reflect"
	"sort"
	"unsafe"
)

// NumMethod returns the number of explicitly declared methods of named type or interface t.
// Wrapper methods for embedded fields or embedded interfaces are not counted.
func (t *xtype) NumMethod() int {
	num := 0
	if gtype, ok := t.gtype.Underlying().(*types.Interface); ok {
		num = gtype.NumExplicitMethods()
	} else if gtype, ok := t.gtype.(*types.Named); ok {
		num = gtype.NumMethods()
	}
	return num
}

// Method return the i-th explicitly declared method of named type or interface t.
// Wrapper methods for embedded fields are not counted
func (t *xtype) Method(i int) Method {
	v := t.universe
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return t.method(i)
}

func (t *xtype) method(i int) Method {
	gfunc := t.gmethod(i)
	name := gfunc.Name()
	resizemethodvalues(t)

	rtype := t.rtype
	var rfunctype reflect.Type
	rfuncs := &t.methodvalues
	rfunc := t.methodvalues[i]
	if rfunc.Kind() == reflect.Func {
		// easy, method is cached already
		rfunctype = rfunc.Type()
	} else if gtype, ok := t.gtype.Underlying().(*types.Interface); ok {
		if rtype.Kind() == reflect.Ptr && isReflectInterfaceStruct(rtype.Elem()) {
			// rtype is our emulated interface type.
			// it's a pointer to a struct containing: InterfaceHeader, embeddeds, methods (without receiver)
			skip := gtype.NumEmbeddeds() + 1
			rfield := rtype.Elem().Field(i + skip)
			rfunctype = addreceiver(rtype, rfield.Type)
		} else if rtype.Kind() != reflect.Interface {
			xerrorf(t, "inconsistent interface type <%v>: expecting interface reflect.Type, found <%v>", t, rtype)
		} else if ast.IsExported(name) {
			// rtype is an interface type, and reflect only returns exported methods
			// rtype.MethodByName returns a Method with the following caveats
			// 1) Type == method signature, without a receiver
			// 2) Func == nil.
			rmethod, _ := rtype.MethodByName(name)
			if rmethod.Type == nil {
				xerrorf(t, "interface type <%v>: reflect method %q not found", t, name)
			} else if rmethod.Index != i {
				xerrorf(t, "inconsistent interface type <%v>: method %q has go/types.Func index=%d but reflect.Method index=%d",
					t, name, i, rmethod.Index)
			}
			rfunctype = addreceiver(rtype, rmethod.Type)
		}
	} else {
		rmethod, _ := rtype.MethodByName(gfunc.Name())
		rfunc = rmethod.Func
		if rfunc.Kind() != reflect.Func {
			if rtype.Kind() != reflect.Ptr {
				// also search in the method set of pointer-to-t
				rmethod, _ = reflect.PtrTo(rtype).MethodByName(gfunc.Name())
				rfunc = rmethod.Func
			}
		}
		if rfunc.Kind() != reflect.Func {
			if ast.IsExported(name) {
				xerrorf(t, "type <%v>: reflect method %q not found", t, gfunc.Name())
			}
		} else {
			t.methodvalues[i] = rfunc
			rfunctype = rmethod.Type
		}
	}
	return t.makemethod(i, gfunc, rfuncs, rfunctype) // lock already held
}

func addreceiver(recv reflect.Type, rtype reflect.Type) reflect.Type {
	nin := rtype.NumIn()
	rin := make([]reflect.Type, nin+1)
	rin[0] = recv
	for i := 0; i < nin; i++ {
		rin[i+1] = rtype.In(i)
	}
	nout := rtype.NumOut()
	rout := make([]reflect.Type, nout)
	for i := 0; i < nout; i++ {
		rout[i] = rtype.Out(i)
	}
	return reflect.FuncOf(rin, rout, rtype.IsVariadic())
}

func (t *xtype) gmethod(i int) *types.Func {
	var gfun *types.Func
	if gtype, ok := t.gtype.Underlying().(*types.Interface); ok {
		gfun = gtype.ExplicitMethod(i)
	} else if gtype, ok := t.gtype.(*types.Named); ok {
		gfun = gtype.Method(i)
	} else {
		xerrorf(t, "Method on invalid type %v", t)
	}
	return gfun
}

func (t *xtype) makemethod(index int, gfun *types.Func, rfuns *[]reflect.Value, rfunctype reflect.Type) Method {
	// sanity checks
	name := gfun.Name()
	gsig := gfun.Type().Underlying().(*types.Signature)
	if rfunctype != nil {
		nparams := 0
		if gsig.Params() != nil {
			nparams = gsig.Params().Len()
		}
		if gsig.Recv() != nil {
			if nparams+1 != rfunctype.NumIn() {
				xerrorf(t, `type <%v>: inconsistent %d-th method signature:
	go/types.Type has receiver <%v> and %d parameters: %v
	reflect.Type has %d parameters: %v`, t, index, gsig.Recv(), nparams, gsig, rfunctype.NumIn(), rfunctype)
			}
		} else if nparams != rfunctype.NumIn() {
			xerrorf(t, `type <%v>: inconsistent %d-th method signature:
	go/types.Type has no receiver and %d parameters: %v
	reflect.Type has %d parameters: %v`, t, index, nparams, gsig, rfunctype.NumIn(), rfunctype)
		}
	}
	var tmethod Type
	if rfunctype != nil {
		tmethod = t.universe.maketype(gsig, rfunctype) // lock already held
	}
	return Method{
		Name:  name,
		Pkg:   (*Package)(gfun.Pkg()),
		Type:  tmethod,
		Funs:  rfuns,
		Index: index,
		GoFun: gfun,
	}
}

func resizemethodvalues(t *xtype) {
	n := t.NumMethod()
	if cap(t.methodvalues) < n {
		slice := make([]reflect.Value, n, n+n/2+4)
		copy(slice, t.methodvalues)
		t.methodvalues = slice
	} else if len(t.methodvalues) < n {
		t.methodvalues = t.methodvalues[0:n]
	}
}

func (v *Universe) NamedOf(name, pkgpath string) Type {
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return v.namedOf(name, pkgpath)
}

func (v *Universe) namedOf(name, pkgpath string) Type {
	underlying := v.TypeOfInterface
	pkg := v.loadPackage(pkgpath)
	// typename := types.NewTypeName(token.NoPos, (*types.Package)(pkg), name, underlying.GoType())
	typename := types.NewTypeName(token.NoPos, (*types.Package)(pkg), name, nil)
	return v.maketype3(
		reflect.Invalid, // incomplete type! will be fixed by SetUnderlying
		types.NewNamed(typename, underlying.GoType(), nil),
		underlying.ReflectType(),
	)
}

// SetUnderlying sets the underlying type of a named type and marks t as complete.
// It panics if the type is unnamed, or if the underlying type is named,
// or if SetUnderlying() was already invoked on the named type.
func (t *xtype) SetUnderlying(underlying Type) {
	switch gtype := t.gtype.(type) {
	case *types.Named:
		v := t.universe
		if t.kind != reflect.Invalid || gtype.Underlying() != v.TypeOfInterface.GoType() || t.rtype != v.TypeOfInterface.ReflectType() {
			// redefined type. try really hard to support it.
			v.InvalidateCache()
			// xerrorf(t, "SetUnderlying invoked multiple times on named type %v", t)
		}
		gunderlying := underlying.GoType().Underlying() // in case underlying is named
		t.kind = gtypeToKind(t, gunderlying)
		gtype.SetUnderlying(gunderlying)
		t.rtype = underlying.ReflectType()
	default:
		xerrorf(t, "SetUnderlying of unnamed type %v", t)
	}
}

// AddMethod adds method 'name' to type.
// It panics if the type is unnamed, or if the signature is not a function type,
// Returns the method index, or < 0 in case of errors
func (t *xtype) AddMethod(name string, signature Type) int {
	gtype, ok := t.gtype.(*types.Named)
	if !ok {
		xerrorf(t, "AddMethod on unnamed type %v", t)
	}
	kind := gtypeToKind(t, gtype.Underlying())
	if kind == reflect.Ptr || kind == reflect.Interface {
		xerrorf(t, "AddMethod: cannot add methods to named %s type: <%v>", kind, t)
	}
	if signature.Kind() != reflect.Func {
		xerrorf(t, "AddMethod on <%v> of non-function: %v", t, signature)
	}
	gsig := signature.underlying().(*types.Signature)
	// accept both signatures "non-nil receiver" and "nil receiver, use the first parameter as receiver"
	grecv := gsig.Recv()
	if grecv == nil && gsig.Params().Len() != 0 {
		grecv = gsig.Params().At(0)
	}
	if grecv == nil {
		xerrorf(t, "AddMethod on <%v> of function with no receiver and no parameters: %v", t, gsig)
	}
	if !types.IdenticalIgnoreTags(grecv.Type(), gtype) &&
		// !types.IdenticalIgnoreTags(grecv.Type(), gtype.Underlying()) &&
		!types.IdenticalIgnoreTags(grecv.Type(), types.NewPointer(gtype)) {

		label := "receiver"
		if gsig.Recv() == nil {
			label = "first parameter"
		}
		xerrorf(t, "AddMethod on <%v> of function <%v> with mismatched %s type: %v", t, gsig, label, grecv.Type())
	}

	gpkg := gtype.Obj().Pkg()
	gfun := types.NewFunc(token.NoPos, gpkg, name, gsig)

	n1 := gtype.NumMethods()
	index := unsafeAddMethod(gtype, gfun)
	n2 := gtype.NumMethods()

	// update the caches... be careful if the method was just redefined
	nilv := reflect.Value{}
	for len(t.methodvalues) < n2 {
		t.methodvalues = append(t.methodvalues, nilv)
	}
	t.methodvalues[index] = nilv
	if n1 == n2 {
		// an existing method was overwritten.
		// it may be cached in some other type's method cache.
		t.universe.InvalidateMethodCache()
	}
	return index
}

// RemoveMethods removes given methods from type.
// It panics if the type is unnamed, or if the signature is not a function type,
func (t *xtype) RemoveMethods(names []string, pkgpath string) {
	gtype, ok := t.gtype.(*types.Named)
	if !ok {
		xerrorf(t, "RemoveMethods on unnamed type %v", t)
	}
	if len(names) == 0 {
		return
	}
	n1 := gtype.NumMethods()
	unsafeRemoveMethods(gtype, names, pkgpath)
	n2 := gtype.NumMethods()
	if n1 != n2 {
		// some existing methods were removed.
		// they may be cached in some other type's method cache.
		t.universe.InvalidateMethodCache()
	}
}

// internal representation of go/types.Named
type unsafeNamed struct {
	obj        *types.TypeName
	underlying types.Type
	methods    []*types.Func
}

// patched version of go/types.Named.AddMethod() that *overwrites* matching methods
// (the original does not)
func unsafeAddMethod(gtype *types.Named, gfun *types.Func) int {
	if gfun.Name() == "_" {
		return -1
	}
	gt := (*unsafeNamed)(unsafe.Pointer(gtype))
	qname := QNameGo(gfun)
	for i, m := range gt.methods {
		if qname == QNameGo(m) {
			gt.methods[i] = gfun
			return i
		}
	}
	gt.methods = append(gt.methods, gfun)
	return len(gt.methods) - 1
}

func unsafeRemoveMethods(gtype *types.Named, names []string, pkgpath string) {
	names = append([]string{}, names...) // make a copy
	sort.Strings(names)                  // and sort it

	gt := (*unsafeNamed)(unsafe.Pointer(gtype))

	n1 := len(gt.methods)
	n2 := n1
	for i, j := 0, 0; i < n1; i++ {
		m := gt.methods[i]
		name := m.Name()
		pos := sort.SearchStrings(names, name)
		if pos < len(names) && names[pos] == name && (m.Exported() || m.Pkg().Path() == pkgpath) {
			// delete this method
			n2--
			continue
		}
		if i != j {
			gt.methods[j] = gt.methods[i]
		}
		j++
	}
	if n1 != n2 {
		gt.methods = gt.methods[:n2]
	}
}

// GetMethods returns the pointer to the method values.
// It panics if the type is unnamed
func (t *xtype) GetMethods() *[]reflect.Value {
	if !t.Named() {
		xerrorf(t, "GetMethods on unnamed type %v", t)
	}
	resizemethodvalues(t)
	return &t.methodvalues
}
