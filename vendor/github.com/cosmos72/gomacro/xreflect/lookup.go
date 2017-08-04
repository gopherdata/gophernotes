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
 * lookup.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/types"
	"reflect"
)

// FieldByName returns the (possibly embedded) struct field with given name,
// and the number of fields found at the same (shallowest) depth: 0 if not found.
// Private fields are returned only if they were declared in pkgpath.
func (t *xtype) FieldByName(name, pkgpath string) (field StructField, count int) {
	if name == "_" || t.kind != reflect.Struct {
		return
	}
	// debugf("field cache for %v <%v> = %v", unsafe.Pointer(t), t, t.fieldcache)
	qname := QName2(name, pkgpath)

	v := t.universe
	if v.ThreadSafe {
		defer un(lock(v))
	}
	field, found := t.fieldcache[qname]
	if found {
		if field.Index == nil { // marker for ambiguous field names
			count = int(field.Offset) // reuse Offset as "number of ambiguous fields"
		} else {
			count = 1
		}
		return field, count
	}
	var tovisit []StructField

	field, count, tovisit = fieldByName(t, qname, 0, nil)

	// breadth-first recursion
	for count == 0 && len(tovisit) != 0 {
		var next []StructField
		for _, f := range tovisit {
			efield, ecount, etovisit := fieldByName(unwrap(f.Type), qname, f.Offset, f.Index)
			if count == 0 {
				if ecount > 0 {
					field = efield
				} else {
					// no recursion if we found something
					next = append(next, etovisit...)
				}
			}
			count += ecount
		}
		tovisit = next
	}
	if count > 0 {
		cacheFieldByName(t, qname, &field, count)
	}
	return field, count
}

func fieldByName(t *xtype, qname QName, offset uintptr, index []int) (field StructField, count int, tovisit []StructField) {
	// also support embedded fields: they can be named types or pointers to named types
	if t.kind == reflect.Ptr {
		t = unwrap(t.elem())
	}
	gtype, ok := t.gtype.Underlying().(*types.Struct)
	if !ok {
		debugf("fieldByName: type is %s, not struct. bailing out", t.kind)
		return
	}
	n := t.NumField()
	for i := 0; i < n; i++ {

		gfield := gtype.Field(i)
		if matchFieldByName(qname, gfield) {
			if count == 0 {
				field = t.field(i) // lock already held
				field.Offset += offset
				field.Index = concat(index, field.Index) // make a copy of index
				// debugf("fieldByName: %d-th field of <%v> matches: %#v", i, t.rtype, field)
			}
			count++
		} else if count == 0 && gfield.Anonymous() {
			efield := t.field(i) // lock already held
			efield.Offset += offset
			efield.Index = concat(index, efield.Index) // make a copy of index
			// debugf("fieldByName: %d-th field of <%v> is anonymous: %#v", i, t.rtype, efield)
			tovisit = append(tovisit, efield)
		}
	}
	return
}

// return true if gfield name matches given name, or if it's anonymous and its *type* name matches given name
func matchFieldByName(qname QName, gfield *types.Var) bool {
	// always check the field's package, not the type's package
	if qname == QNameGo(gfield) {
		return true
	}
	if gfield.Anonymous() {
		gtype := gfield.Type()
		if gptr, ok := gtype.(*types.Pointer); ok {
			// unnamed field has unnamed pointer type, as for example *Foo
			// check the element type
			gtype = gptr.Elem()
		}
		switch gtype := gtype.(type) {
		case *types.Basic:
			// is it possible to embed basic types?
			// yes, and they work as unexported embedded fields,
			// i.e. in the same package as the struct that includes them
			return qname == QNameGo2(gtype.Name(), gfield.Pkg())
		case *types.Named:
			// gtype.Obj().Pkg() and gfield.Pkg() should be identical for *unexported* fields
			// (they are ignored for exported fields)
			return qname == QNameGo2(gtype.Obj().Name(), gfield.Pkg())
		}
	}
	return false
}

// add field to type's fieldcache. used by Type.FieldByName after a successful lookup
func cacheFieldByName(t *xtype, qname QName, field *StructField, count int) {
	if t.fieldcache == nil {
		t.fieldcache = make(map[QName]StructField)
	}
	if count > 1 {
		field.Index = nil             // marker for ambiguous field names
		field.Offset = uintptr(count) // reuse Offset as "number of ambiguous fields"
	}
	t.fieldcache[qname] = *field
	t.universe.fieldcache = true
}

// anonymousFields returns the anonymous fields of a (named or unnamed) struct type
func anonymousFields(t *xtype, offset uintptr, index []int) []StructField {
	var tovisit []StructField
	gt := t.gtype.Underlying()
	if gptr, ok := gt.(*types.Pointer); ok {
		gt = gptr.Elem().Underlying()
	}
	gtype, ok := gt.(*types.Struct)
	if !ok {
		return tovisit
	}
	n := gtype.NumFields()
	for i := 0; i < n; i++ {
		gfield := gtype.Field(i)
		if gfield.Anonymous() {
			field := t.Field(i)
			field.Offset += offset
			field.Index = concat(index, field.Index) // make a copy of index
			tovisit = append(tovisit, field)
		}
	}
	return tovisit
}

// MethodByName returns the method with given name (including wrapper methods for embedded fields)
// and the number of methods found at the same (shallowest) depth: 0 if not found.
// Private methods are returned only if they were declared in pkgpath.
func (t *xtype) MethodByName(name, pkgpath string) (method Method, count int) {
	// debugf("method cache for %v <%v> = %v", unsafe.Pointer(t), t, t.methodcache)

	// only named types and interfaces can have methods
	if name == "_" || (!t.Named() && t.kind != reflect.Interface) {
		return
	}
	qname := QName2(name, pkgpath)
	v := t.universe
	if v.ThreadSafe {
		defer un(lock(v))
	}
	method, found := t.methodcache[qname]
	if found {
		index := method.Index
		if index < 0 { // marker for ambiguous method names
			count = -index
		} else {
			count = 1
		}
		return method, count
	}
	method, count = methodByName(t, qname, nil)
	if count == 0 {
		tovisit := anonymousFields(t, 0, nil)
		// breadth-first recursion
		for count == 0 && len(tovisit) != 0 {
			var next []StructField
			for _, f := range tovisit {
				et := unwrap(f.Type)
				emethod, ecount := methodByName(et, qname, f.Index)
				if count == 0 {
					if ecount > 0 {
						method = emethod
					} else {
						// no recursion if we found something
						next = append(next, anonymousFields(et, f.Offset, f.Index)...)
					}
				}
				count += ecount
			}
			tovisit = next
		}
	}
	if count > 0 {
		cacheMethodByName(t, qname, &method, count)
	}
	return method, count
}

func methodByName(t *xtype, qname QName, index []int) (method Method, count int) {
	// also support embedded fields: they can be named types or pointers to named types
	if t.kind == reflect.Ptr {
		t = unwrap(t.elem())
	}
	n := t.NumMethod()
	for i := 0; i < n; i++ {
		gmethod := t.gmethod(i)
		if matchMethodByName(qname, gmethod) {
			if count == 0 {
				method = t.method(i)                                 // lock already held
				method.FieldIndex = concat(index, method.FieldIndex) // make a copy of index
				// debugf("methodByName: %d-th method of <%v> matches: %#v", i, t.rtype, method)
			}
			count++
		}
	}
	return
}

// return true if gmethod name matches given name
func matchMethodByName(qname QName, gmethod *types.Func) bool {
	// always check the methods's package, not the type's package
	return qname == QNameGo(gmethod)
}

// add method to type's methodcache. used by Type.MethodByName after a successful lookup
func cacheMethodByName(t *xtype, qname QName, method *Method, count int) {
	if t.methodcache == nil {
		t.methodcache = make(map[QName]Method)
	}
	if count > 1 {
		method.Index = -count // marker for ambiguous method names
	}
	t.methodcache[qname] = *method
	t.universe.methodcache = true
}

func invalidateCache(gtype types.Type, t interface{}) {
	if t, ok := t.(Type); ok {
		t := unwrap(t)
		t.fieldcache = nil
		t.methodcache = nil
	}
}

func invalidateMethodCache(gtype types.Type, t interface{}) {
	if t, ok := t.(Type); ok {
		t := unwrap(t)
		t.methodcache = nil
	}
}

// clears all xtype.fieldcache and xtype.methodcache.
// invoked by NamedOf() when a type is redefined.
func (v *Universe) InvalidateCache() {
	if v.fieldcache || v.methodcache {
		v.gmap.Iterate(invalidateCache)
		v.fieldcache = false
		v.methodcache = false
	}
}

// clears all xtype.methodcache.
// invoked by AddMethod() when a method is redefined.
func (v *Universe) InvalidateMethodCache() {
	if v.methodcache {
		v.gmap.Iterate(invalidateMethodCache)
		v.methodcache = false
	}
}
