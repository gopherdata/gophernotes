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
 * struct.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"reflect"
)

// Field returns a struct type's i'th field.
// It panics if the type's Kind is not Struct.
// It panics if i is not in the range [0, NumField()).
func (t *xtype) Field(i int) StructField {
	if t.kind != reflect.Struct {
		xerrorf(t, "Field of non-struct type %v", t)
	}
	v := t.universe
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return t.field(i)
}

func (t *xtype) field(i int) StructField {
	if t.kind != reflect.Struct {
		xerrorf(t, "Field of non-struct type %v", t)
	}
	gtype := t.gtype.Underlying().(*types.Struct)

	if i < 0 || i >= gtype.NumFields() {
		xerrorf(t, "Field(%v) out of bounds, struct type has %v fields: %v", i, gtype.NumFields(), t)
	}
	va := gtype.Field(i)
	var rf reflect.StructField
	if t.rtype != rTypeOfForward {
		rf = t.rtype.Field(i)
	} else {
		// cannot dig in a forward-declared type,
		// so try to resolve it
		it := t.universe.gmap.At(t.gtype)
		if it != nil {
			rtype := it.(Type).ReflectType()
			if rtype.Kind() != t.kind {
				debugf("mismatched Forward type: <%v> has reflect.Type <%v>", t, rtype)
			}
			rf = rtype.Field(i)
		} else {
			// populate  Field.Index and approximate Field.Type
			rf.Index = []int{i}
			rf.Type = rTypeOfForward
		}
	}

	return StructField{
		Name:      va.Name(),
		Pkg:       (*Package)(va.Pkg()),
		Type:      t.universe.maketype(va.Type(), rf.Type), // lock already held
		Tag:       rf.Tag,
		Offset:    rf.Offset,
		Index:     rf.Index,
		Anonymous: va.Anonymous(),
	}
}

// NumField returns a struct type's field count.
// It panics if the type's Kind is not Struct.
func (t *xtype) NumField() int {
	if t.kind != reflect.Struct {
		xerrorf(t, "NumField of non-struct type %v", t)
	}
	gtype := t.gunderlying().(*types.Struct)
	return gtype.NumFields()
}

func (field *StructField) toReflectField(forceExported bool) reflect.StructField {
	var pkgpath string
	if pkg := field.Pkg; pkg != nil && !forceExported {
		pkgpath = pkg.Path()
	}
	name := field.Name
	if forceExported {
		name = toExportedFieldName(name, field.Type, field.Anonymous)
	}
	return reflect.StructField{
		Name:    name,
		PkgPath: pkgpath,
		Type:    field.Type.ReflectType(),
		Tag:     field.Tag,
		Offset:  field.Offset,
		Index:   field.Index,
		// reflect.StructOf() has very limited support for anonymous fields,
		// do not even try to use it.
		Anonymous: false,
	}
}

func toReflectFields(fields []StructField, forceExported bool) []reflect.StructField {
	rfields := make([]reflect.StructField, len(fields))
	for i := range fields {
		rfields[i] = fields[i].toReflectField(forceExported)
	}
	return rfields
}

func (field *StructField) sanitize(i int) {
	if len(field.Name) != 0 {
		return
	}
	t := field.Type
	name := t.Name()
	if len(name) == 0 && t.Kind() == reflect.Ptr {
		name = t.elem().Name()
	}
	if len(name) == 0 {
		name = fmt.Sprintf("%s%d", StrGensymAnonymous, i)
	}
	field.Name = name
	field.Anonymous = true
}

func (field *StructField) toGoField(i int) *types.Var {
	field.sanitize(i)
	return types.NewField(token.NoPos, (*types.Package)(field.Pkg), field.Name, field.Type.GoType(), field.Anonymous)
}

func toGoFields(fields []StructField) []*types.Var {
	vars := make([]*types.Var, len(fields))
	for i := range fields {
		vars[i] = fields[i].toGoField(i)
	}
	return vars
}

func (field *StructField) toTag() string {
	return string(field.Tag)
}

func toTags(fields []StructField) []string {
	tags := make([]string, len(fields))
	for i := range fields {
		tags[i] = fields[i].toTag()
	}
	return tags
}

func toExportedFieldName(name string, t Type, anonymous bool) string {
	if len(name) == 0 && unwrap(t) != nil {
		if name = t.Name(); len(name) == 0 && t.Kind() == reflect.Ptr {
			name = t.elem().Name()
		}
	}
	if !ast.IsExported(name) {
		if anonymous {
			return GensymAnonymous(name)
		} else {
			return GensymPrivate(name)
		}
	}
	return name
}

func (v *Universe) StructOf(fields []StructField) Type {
	vars := toGoFields(fields)
	tags := toTags(fields)
	rfields := toReflectFields(fields, true)
	return v.MakeType(
		types.NewStruct(vars, tags),
		reflect.StructOf(rfields),
	)
}
