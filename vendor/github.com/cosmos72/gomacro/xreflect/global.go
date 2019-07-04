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
 * global.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/ast"
	r "reflect"

	"github.com/cosmos72/gomacro/go/types"
)

type Package types.Package

type Forward interface{}

// InterfaceHeader is the internal header of interpreted interfaces
type InterfaceHeader struct {
	// val and typ must be private! otherwise interpreted code may mess with them and break type safety
	val r.Value
	typ Type
}

func MakeInterfaceHeader(val r.Value, typ Type) InterfaceHeader {
	if val.IsValid() && val.CanSet() {
		val = val.Convert(val.Type()) // make a copy
	}
	return InterfaceHeader{val, typ}
}

func (h InterfaceHeader) Value() r.Value {
	return h.val
}

func (h InterfaceHeader) Type() Type {
	return h.typ
}

type Method struct {
	Name       string
	Pkg        *Package
	Type       Type        // method type
	Funs       *[]r.Value  // (*Funs)[Index] is the method, with receiver as first argument
	Index      int         // index for Type.Method
	FieldIndex []int       // embedded fields index sequence for r.Type.FieldByIndex or r.Value.FieldByIndex
	GoFun      *types.Func // for completeness
}

type StructField struct {
	// Name is the field name. If empty, it will be computed from Type name, and Anonymous will be set to true
	Name string
	// Pkg is the package that qualifies a lower case (unexported)
	// field name. It may be nil for upper case (exported) field names.
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	Pkg       *Package
	Type      Type        // field type
	Tag       r.StructTag // field tag string
	Offset    uintptr     // offset within struct, in bytes. meaningful only if all Deref[] are false
	Index     []int       // index sequence for r.Type.FieldByIndex or r.Value.FieldByIndex
	Anonymous bool        // is an embedded field. If true, Name should be empty or equal to the type's name
}

type addmethods uint8

const (
	addmethodsNeeded addmethods = iota
	addmethodsDone
)

type xtype struct {
	kind         r.Kind
	gtype        types.Type
	rtype        r.Type
	universe     *Universe
	methodvalues []r.Value
	fieldcache   map[QName]StructField
	methodcache  map[QName]Method
	userdata     map[interface{}]interface{}
	addmethods   addmethods
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

func (q QName) String() string {
	if len(q.pkgpath) == 0 {
		return q.name
	}
	return q.pkgpath + "." + q.name
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

// Key is a Type wrapper suitable for use with operator == and as map[T1]T2 key
type Key struct {
	universe *Universe
	gtype    types.Type
}

func MakeKey(t Type) Key {
	xt := unwrap(t)
	if xt == nil {
		return Key{}
	}
	i := xt.universe.gmap.At(xt.gtype)
	if i != nil {
		xt = unwrap(i.(Type))
	}
	return Key{xt.universe, xt.gtype}
}

func (k Key) Type() Type {
	if k.universe == nil || k.gtype == nil {
		return nil
	}
	i := k.universe.gmap.At(k.gtype)
	if i == nil {
		return nil
	}
	return i.(Type)
}
