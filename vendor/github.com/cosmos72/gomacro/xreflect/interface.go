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
 * struct.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/token"
	"go/types"
	"reflect"
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
		place.Field(i + 2).Set(obj2methods[i](v))
	}
	return addr
}

// extract the already-made i-th closure from inside the emulated interface object.
func EmulatedInterfaceGetMethod(obj reflect.Value, index int) reflect.Value {
	return obj.Elem().Field(index + 2)
}

// create []*types.Func suitable for types.NewInterface.
// makes a copy of each methods[i].gunderlying().(*types.Signature)
// because types.NewInterface will destructively modify them!
func toGoFuncs(names []string, methods []Type) []*types.Func {
	gfuns := make([]*types.Func, len(methods))
	for i, t := range methods {
		switch gsig := t.gunderlying().(type) {
		case *types.Signature:
			gsig = cloneGoSignature(gsig)
			gfuns[i] = types.NewFunc(token.NoPos, nil, names[i], gsig)
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

// InterfaceOf returns a new interface for the given methods and embedded types.
// After the methods and embeddeds are fully defined, call Complete() to mark
// the interface as complete and compute wrapper methods for embedded fields.
//
// WARNING: the Type returned by InterfaceOf is not complete,
// i.e. its method set is not computed yet.
// Once you know that methods and embedded interfaces are complete,
// call Complete() to compute the method set and mark this Type as complete.
func (v *Universe) InterfaceOf(methodnames []string, methodtypes []Type, embeddeds []Type) Type {
	gmethods := toGoFuncs(methodnames, methodtypes)
	gembeddeds := toGoNamedTypes(embeddeds)

	gtype := types.NewInterface(gmethods, gembeddeds)
	gtype.Complete()

	// for reflect.Type, approximate an interface as a pointer-to-struct:
	// one field for the wrapped object: type is interface{},
	// one field for each explicit method: type is the method type i.e. a function
	rfields := make([]reflect.StructField, 2+len(methodtypes), gtype.NumMethods()+2)
	rfields[0] = approxInterfaceHeader()
	rfields[1] = approxInterfaceEmbeddeds(embeddeds)

	for i, methodtype := range methodtypes {
		rfields[i+2] = approxInterfaceMethodAsField(methodnames[i], methodtype.ReflectType())
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

// if true, produces more detailes reflect.Types for emulated interfaces,
// at the cost of an exponential explosion of their String() output
const ANNOTATE_EMULATED_INTERFACES_WITH_EMBEDDED_INTERFACE_TYPES = false

func approxInterfaceEmbeddeds(embeddeds []Type) reflect.StructField {
	var rtype reflect.Type
	if ANNOTATE_EMULATED_INTERFACES_WITH_EMBEDDED_INTERFACE_TYPES {
		fields := make([]reflect.StructField, len(embeddeds))
		for i, t := range embeddeds {
			fields[i] = approxInterfaceEmbedded(t)
		}
		rtype = reflect.ArrayOf(0, reflect.StructOf(fields))
	} else {
		rtype = reflect.ArrayOf(0, reflect.TypeOf(struct{}{}))
	}
	return reflect.StructField{Name: StrGensymEmbedded, Type: rtype}
}

func approxInterfaceEmbedded(t Type) reflect.StructField {
	return reflect.StructField{
		Name: toExportedFieldName("", t, true),
		Type: t.ReflectType(),
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
	// rtype is *struct { InterfaceHeader; [0]struct{ embeddeds.. }; closures... }
	index += 2
	rclosure := rtype.Elem().Field(index).Type
	rfunc := rAddReceiver(rtype, rclosure)
	return reflect.MakeFunc(rfunc, func(args []reflect.Value) []reflect.Value {
		return args[0].Elem().Field(index).Call(args[1:])
	})
}
