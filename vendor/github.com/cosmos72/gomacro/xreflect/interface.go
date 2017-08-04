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

func toGoFuncs(names []string, methods []Type) []*types.Func {
	gfuns := make([]*types.Func, len(methods))
	for i, t := range methods {
		switch gsig := t.underlying().(type) {
		case *types.Signature:
			gfuns[i] = types.NewFunc(token.NoPos, nil, names[i], gsig)
		default:
			errorf(t, "InterfaceOf: %d-th 'method' argument is not a function type: %v", i, t)
		}
	}
	return gfuns
}

func toGoNamedTypes(ts []Type) []*types.Named {
	gnameds := make([]*types.Named, len(ts))
	for i, t := range ts {
		if gt, ok := t.GoType().(*types.Named); ok {
			gnameds[i] = gt
		} else {
			errorf(t, "InterfaceOf: %d-th 'embedded' argument is not a named type: %v", i, t)
		}
	}
	return gnameds
}

func (v *Universe) InterfaceOf(methodnames []string, methods []Type, embeddeds []Type) Type {
	gmethods := toGoFuncs(methodnames, methods)
	gembeddeds := toGoNamedTypes(embeddeds)

	// for reflect.Type, approximate an interface as a struct:
	// one field for the wrapped object: type is interface{},
	// one field for each embedded interface: type is the embedded interface type
	// one field for each method: type is the method type i.e. a function
	nemb := len(embeddeds)
	rfields := make([]reflect.StructField, 1+nemb+len(methods))
	rfields[0] = approxInterfaceHeader()
	for i, emb := range embeddeds {
		rfields[i+1] = approxInterfaceEmbedded(emb)
	}
	for i, method := range methods {
		rfields[i+nemb+1] = approxInterfaceMethod(methodnames[i], method.ReflectType())
	}
	return v.maketype3(
		reflect.Interface,
		types.NewInterface(gmethods, gembeddeds),
		// interfaces may have lots of methods, thus a lot of fields in the proxy struct.
		// Use a pointer to the proxy struct
		reflect.PtrTo(reflect.StructOf(rfields)),
	)
}

// InterfaceOf returns a new interface for the given methods and embedded types.
// After the methods and embeddeds are fully defined, call Complete() to mark
// the interface as complete and compute wrapper methods for embedded fields.
//
// WARNING: the Type returned by InterfaceOf is not complete,
// i.e. its method set is not computed yet.
// Once you know that methods and embedded interfaces are complete,
// call Complete() to compute the method set and mark this Type as complete.
func InterfaceOf(methodnames []string, methods []Type, embeddeds []Type) Type {
	v := universe
	if len(embeddeds) != 0 && embeddeds[0] != nil {
		v = embeddeds[0].Universe()
	} else if len(methods) != 0 && methods[0] != nil {
		v = methods[0].Universe()
	}
	return v.InterfaceOf(methodnames, methods, embeddeds)
}

// utilities for InterfaceOf()

func approxInterfaceHeader() reflect.StructField {
	return reflect.StructField{
		Name: StrGensymInterface,
		Type: reflectTypeOfInterfaceHeader,
	}
}

func approxInterfaceEmbedded(t Type) reflect.StructField {
	// embedded interfaces are always anonymous
	return reflect.StructField{
		Name: toExportedFieldName("", t, true),
		Type: t.ReflectType(),
	}
}

func approxInterfaceMethod(name string, rtype reflect.Type) reflect.StructField {
	// interface methods cannot be anonymous
	if len(name) == 0 {
		name = "_"
	}
	return reflect.StructField{
		Name: toExportedFieldName(name, nil, false),
		Type: rtype,
	}
}
