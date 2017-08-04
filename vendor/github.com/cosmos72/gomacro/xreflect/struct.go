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

	va := gtype.Field(i)
	rf := t.rtype.Field(i)

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
	gtype := t.underlying().(*types.Struct)
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
		Name:      name,
		PkgPath:   pkgpath,
		Type:      field.Type.ReflectType(),
		Tag:       field.Tag,
		Offset:    field.Offset,
		Index:     field.Index,
		Anonymous: field.Anonymous,
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
		name = StrGensymEmbedded + fmt.Sprintf("%d", i)
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
	if len(name) == 0 && t != nil {
		if name = t.Name(); len(name) == 0 && t.Kind() == reflect.Ptr {
			name = t.elem().Name()
		}
	}
	if !ast.IsExported(name) {
		if anonymous {
			return GensymEmbedded(name)
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

func StructOf(fields []StructField) Type {
	v := universe
	if len(fields) != 0 && fields[0].Type != nil {
		v = fields[0].Type.Universe()
	}
	return v.StructOf(fields)
}
