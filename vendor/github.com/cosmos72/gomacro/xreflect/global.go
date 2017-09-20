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
 * global.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/ast"
	"go/types"
	"reflect"
)

type Package types.Package

// InterfaceHeader is the internal header of interpreted interfaces
type InterfaceHeader struct {
	// val and typ must be private! otherwise interpreted code may mess with them and break type safety
	val reflect.Value
	typ Type
}

func MakeInterfaceHeader(val reflect.Value, typ Type) InterfaceHeader {
	if val.IsValid() && val.CanSet() {
		val = val.Convert(val.Type()) // make a copy
	}
	return InterfaceHeader{val, typ}
}

func (h InterfaceHeader) Value() reflect.Value {
	return h.val
}

func (h InterfaceHeader) Type() Type {
	return h.typ
}

type Method struct {
	Name       string
	Pkg        *Package
	Type       Type             // method type
	Funs       *[]reflect.Value // (*Funs)[Index] is the method, with receiver as first argument
	Index      int              // index for Type.Method
	FieldIndex []int            // embedded fields index sequence for reflect.Type.FieldByIndex or reflect.Value.FieldByIndex
	GoFun      *types.Func      // for completeness
}

type StructField struct {
	// Name is the field name.
	Name string
	// Pkg is the package that qualifies a lower case (unexported)
	// field name. It may be nil for upper case (exported) field names.
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	Pkg       *Package
	Type      Type              // field type
	Tag       reflect.StructTag // field tag string
	Offset    uintptr           // offset within struct, in bytes
	Index     []int             // index sequence for reflect.Type.FieldByIndex or reflect.Value.FieldByIndex
	Anonymous bool              // is an embedded field. Note: embedded field's name should be set to the type's name
}

type xtype struct {
	kind         reflect.Kind
	gtype        types.Type
	rtype        reflect.Type
	universe     *Universe
	methodvalues []reflect.Value
	fieldcache   map[QName]StructField
	methodcache  map[QName]Method
}

// QName is a replacement for go/types.Id and implements accurate comparison
// of type names, field names and method names.
// It recognizes unexported names, and names declared in different packages.
//
// To compare two names, build two QNames with the functions QName*
// then compare the two QName structs with ==
type QName struct {
	name, pkgpath string
}

func (q QName) Name() string {
	return q.name
}

func (q QName) PkgPath() string {
	return q.pkgpath
}

type QNameI interface {
	Name() string
	PkgPath() string
}

func QName2(name, pkgpath string) QName {
	if ast.IsExported(name) {
		pkgpath = ""
	}
	return QName{name, pkgpath}
}

func QName1(q QNameI) QName {
	return QName2(q.Name(), q.PkgPath())
}

func QNameGo2(name string, pkg *types.Package) QName {
	var pkgpath string
	if pkg != nil && !ast.IsExported(name) {
		pkgpath = pkg.Path()
	}
	return QName{name, pkgpath}
}

func QNameGo(obj types.Object) QName {
	return QNameGo2(obj.Name(), obj.Pkg())
}
