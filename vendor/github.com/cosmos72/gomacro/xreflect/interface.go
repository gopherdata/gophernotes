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
 * interface.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/token"
	"go/types"
	"reflect"
	"sort"
)

func IsEmulatedInterface(t Type) bool {
	xt := unwrap(t)
	return xt.kind == reflect.Interface && xt.rtype.Kind() == reflect.Ptr
}

// extract the concrete value and type contained in an emulated interface
func FromEmulatedInterface(v reflect.Value) (reflect.Value, Type) {
	h := v.Elem().Field(0).Interface().(InterfaceHeader)
	return h.val, h.typ
}

// create an emulated interface from given value, type and method extractors
// (methods extractors are functions that, given a value, return one of its methods)
func ToEmulatedInterface(rtypeinterf reflect.Type, v reflect.Value,
	t Type, obj2methods []func(reflect.Value) reflect.Value) reflect.Value {

	addr := reflect.New(rtypeinterf.Elem())
	place := addr.Elem()
	place.Field(0).Set(reflect.ValueOf(InterfaceHeader{v, t}))
	for i := range obj2methods {
		mtd := obj2methods[i](v)
		place.Field(i + 1).Set(mtd)
	}
	return addr
}

// extract the already-made i-th closure from inside the emulated interface object.
func EmulatedInterfaceGetMethod(obj reflect.Value, index int) reflect.Value {
	return obj.Elem().Field(index + 1)
}

// create []*types.Func suitable for types.NewInterface.
// makes a copy of each methods[i].gunderlying().(*types.Signature)
// because types.NewInterface will destructively modify them!
func toGoFuncs(pkg *Package, names []string, methods []Type) []*types.Func {
	gfuns := make([]*types.Func, len(methods))
	for i, t := range methods {
		switch gsig := t.gunderlying().(type) {
		case *types.Signature:
			gsig = cloneGoSignature(gsig)
			gfuns[i] = types.NewFunc(token.NoPos, (*types.Package)(pkg), names[i], gsig)
		default:
			errorf(t, "interface contains non-function: %s %v", names[i], t)
		}
	}
	return gfuns
}

func cloneGoSignature(gsig *types.Signature) *types.Signature {
	return types.NewSignature(gsig.Recv(), gsig.Params(), gsig.Results(), gsig.Variadic())
}

func toGoNamedTypes(ts []Type) []*types.Named {
	gnameds := make([]*types.Named, len(ts))
	for i, t := range ts {
		if gt, ok := t.GoType().(*types.Named); ok {
			if t.Kind() == reflect.Interface {
				gnameds[i] = gt
			} else {
				errorf(t, "interface contains embedded non-interface: %v", t)
			}
		} else {
			errorf(t, "interface contains embedded interface without name: %v", t)
		}
	}
	return gnameds
}

type byQName struct {
	qname  []QName
	method []Type
}

func (a *byQName) Len() int { return len(a.qname) }

func (a *byQName) Less(i, j int) bool {
	return QLess(a.qname[i], a.qname[j])
}

func (a *byQName) Swap(i, j int) {
	a.qname[i], a.qname[j] = a.qname[j], a.qname[i]
	a.method[i], a.method[j] = a.method[j], a.method[i]
}

// InterfaceOf returns a new interface for the given methods and embedded types.
// After the methods and embeddeds are fully defined, call Complete() to mark
// the interface as complete and compute wrapper methods for embedded fields.
//
// unexported method names are created in 'pkg'.
//
// WARNING: the Type returned by InterfaceOf is not complete,
// i.e. its method set is not computed yet.
// Once you know that methods and embedded interfaces are complete,
// call Complete() to compute the method set and mark this Type as complete.
func (v *Universe) InterfaceOf(pkg *Package, methodnames []string, methodtypes []Type, embeddeds []Type) Type {
	methodnames = append(([]string)(nil), methodnames...) // dup before modifying
	methodtypes = append(([]Type)(nil), methodtypes...)   // dup before modifying
	embeddeds = append(([]Type)(nil), embeddeds...)       // dup before modifying

	// types.NewInterface() sorts methodtypes and embeddeds by Id().
	// We must do the same to keep the method and embedded order in sync.
	qnames := make([]QName, len(methodnames))
	for i, name := range methodnames {
		qnames[i] = QNameGo2(name, (*types.Package)(pkg))
	}
	sort.Sort(&byQName{qnames, methodtypes})
	sort.Slice(embeddeds, func(i, j int) bool {
		return embeddeds[i].GoType().(*types.Named).Obj().Id() < embeddeds[j].GoType().(*types.Named).Obj().Id()
	})
	for i, qname := range qnames {
		methodnames[i] = qname.name
	}

	gmethods := toGoFuncs(pkg, methodnames, methodtypes)
	gembeddeds := toGoNamedTypes(embeddeds)

	gtype := types.NewInterface(gmethods, gembeddeds)
	gtype.Complete()

	// for reflect.Type, approximate an interface as a pointer-to-struct:
	// one field for the wrapped object: type is interface{},
	// one field for each explicit method: type is the method type i.e. a function
	rfields := make([]reflect.StructField, 1+len(methodtypes), 1+gtype.NumMethods())
	rfields[0] = approxInterfaceHeader()

	for i, methodtype := range methodtypes {
		rfields[i+1] = approxInterfaceMethodAsField(methodnames[i], methodtype.ReflectType())
	}
	for _, e := range embeddeds {
		n := e.NumMethod()
		for i := 0; i < n; i++ {
			method := e.Method(i)
			rtype := rRemoveReceiver(method.Type.ReflectType()) // receiver is the embedded interface, remove it
			rfields = append(rfields, approxInterfaceMethodAsField(method.Name, rtype))
		}
	}
	// interfaces may have lots of methods, thus a lot of fields in the proxy struct.
	// Use a pointer to the proxy struct
	rtype := reflect.PtrTo(reflect.StructOf(rfields))
	t := v.maketype3(reflect.Interface, gtype, rtype)
	setInterfaceMethods(t)
	// debugf("InterfaceOf: new type %v", t)
	// debugf("           types.Type %v", gtype)
	// debugf("         reflect.Type %v", rtype)
	return t
}

// Complete marks an interface type as complete and computes wrapper methods for embedded fields.
// It must be called by users of InterfaceOf after the interface's embedded types are fully defined
// and before using the interface type in any way other than to form other types.
func (t *xtype) Complete() Type {
	if t.kind != reflect.Interface {
		xerrorf(t, "Complete of non-interface %v", t)
	}
	return wrap(t)
}

// return true if t is a named type that still waits for the caller to invoke SetUnderlying() on it
func (t *xtype) needSetUnderlying() bool {
	return t.Named() && t.kind != gtypeToKind(t, t.gtype)
}

// utilities for InterfaceOf()

func approxInterfaceHeader() reflect.StructField {
	return reflect.StructField{
		Name: StrGensymInterface,
		Type: rTypeOfInterfaceHeader,
	}
}

func approxInterfaceMethodAsField(name string, rtype reflect.Type) reflect.StructField {
	// interface methods cannot be anonymous
	if len(name) == 0 {
		name = "_"
	}
	return reflect.StructField{
		Name: toExportedFieldName(name, nil, false),
		Type: rtype,
	}
}

// fill t.methodvalues[] with wrappers that forward the call to closures stored in the emulated interface struct
func setInterfaceMethods(t Type) {
	xt := unwrap(t)
	n := xt.NumMethod()
	if n == 0 || xt.Named() || xt.kind != reflect.Interface || xt.methodvalues != nil {
		return
	}
	xt.methodvalues = make([]reflect.Value, n)
	rtype := xt.rtype
	for i := 0; i < n; i++ {
		xt.methodvalues[i] = interfaceMethod(rtype, i)
	}
}

// create and return a single wrapper function that forwards the call to the i-th closure
// stored in the emulated interface struct rtype (that will be received as first parameter)
func interfaceMethod(rtype reflect.Type, index int) reflect.Value {
	// rtype is *struct { InterfaceHeader; closures... }
	index++
	rclosure := rtype.Elem().Field(index).Type
	rfunc := rAddReceiver(rtype, rclosure)
	return reflect.MakeFunc(rfunc, func(args []reflect.Value) []reflect.Value {
		return args[0].Elem().Field(index).Call(args[1:])
	})
}
