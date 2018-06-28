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

type Forward interface{}

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
	// Name is the field name. If empty, it will be computed from Type name, and Anonymous will be set to true
	Name string
	// Pkg is the package that qualifies a lower case (unexported)
	// field name. It may be nil for upper case (exported) field names.
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	Pkg       *Package
	Type      Type              // field type
	Tag       reflect.StructTag // field tag string
	Offset    uintptr           // offset within struct, in bytes. meaningful only if all Deref[] are false
	Index     []int             // index sequence for reflect.Type.FieldByIndex or reflect.Value.FieldByIndex
	Anonymous bool              // is an embedded field. If true, Name should be empty or equal to the type's name
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
	pkgpath, name string
}

func (q QName) Name() string {
	return q.name
}

func (q QName) PkgPath() string {
	return q.pkgpath
}

func QLess(p, q QName) bool {
	return p.pkgpath < q.pkgpath || (p.pkgpath == q.pkgpath && p.name < q.name)
}

type QNameI interface {
	Name() string
	PkgPath() string
}

func QName2(name, pkgpath string) QName {
	if ast.IsExported(name) {
		pkgpath = ""
	}
	return QName{pkgpath, name}
}

func QName1(q QNameI) QName {
	return QName2(q.Name(), q.PkgPath())
}

func QNameGo2(name string, pkg *types.Package) QName {
	var pkgpath string
	if !ast.IsExported(name) {
		if pkg != nil {
			pkgpath = pkg.Path()
		}
		if len(pkgpath) == 0 {
			pkgpath = "_"
		}
	}
	return QName{pkgpath, name}
}

func QNameGo(obj types.Object) QName {
	return QNameGo2(obj.Name(), obj.Pkg())
}
