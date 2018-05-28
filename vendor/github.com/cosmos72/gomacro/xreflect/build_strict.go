// +build gomacro_xreflect_strict

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 * build_strict.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/types"
	"reflect"
)

// Type:s must be compared with IdenticalTo, not with ==
// produce compile-time error on == between Type:s
type Type []*xtype

// Align returns the alignment in bytes of a value of
// this type when allocated in memory.
func (t Type) Align() int {
	return t[0].Align()
}

// FieldAlign returns the alignment in bytes of a value of
// this type when used as a field in a struct.
func (t Type) FieldAlign() int {
	return t[0].FieldAlign()
}

// Identical reports whether the type is identical to type u.
func (t Type) IdenticalTo(u Type) bool {
	return identicalType(t, u)
}

// AssignableTo reports whether a value of the type is assignable to type u.
func (t Type) AssignableTo(u Type) bool {
	return t[0].AssignableTo(u)
}

// ConvertibleTo reports whether a value of the type is convertible to type u.
func (t Type) ConvertibleTo(u Type) bool {
	return t[0].ConvertibleTo(u)
}

// Comparable reports whether values of this type are comparable.
func (t Type) Comparable() bool {
	return t[0].Comparable()
}

// GoType returns the go/types.Type corresponding to the given type.
func (t Type) GoType() types.Type {
	return t[0].GoType()
}

// Implements reports whether the type implements the interface type u.
// It panics if u's Kind is not Interface
func (t Type) Implements(u Type) bool {
	return t[0].Implements(u)
}

// Name returns the type's name within its package.
// It returns an empty string for unnamed types.
func (t Type) Name() string {
	if len(t) == 0 {
		return ""
	}
	return t[0].Name()
}

// Named returns whether the type is named.
// It returns false for unnamed types.
func (t Type) Named() bool {
	if len(t) == 0 {
		return false
	}
	return t[0].Named()
}

// Pkg returns a named type's package, that is, the package where it was defined.
// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
// Pkg will return nil.
func (t Type) Pkg() *Package {
	return t[0].Pkg()
}

// PkgName returns a named type's package name, that is,
// the default name that the package provides when imported.
// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
// the package name will be the empty string.
func (t Type) PkgName() string {
	return t[0].PkgName()
}

// PkgPath returns a named type's package path, that is, the import path
// that uniquely identifies the package, such as "encoding/base64".
// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
// the package path will be the empty string.
func (t Type) PkgPath() string {
	return t[0].PkgPath()
}

// ReflectType returns a best-effort reflect.Type that approximates the type.
// It may be inexact for the following reasons:
// 1) missing reflect.NamedOf(): no way to programmatically create named types, or to access the underlying type of a named type
// 2) missing reflect.InterfaceOf(): interface types created at runtime will be approximated by structs
// 3) missing reflect.MethodOf(): method types created at runtime will be approximated by functions
//    whose first parameter is the receiver
// 4) reflect.StructOf() does not support embedded or unexported fields
// 5) go/reflect lacks the ability to create self-referencing types:
//    references to the type itself will be replaced by interface{}.
//
// Examples:
//    after invoking at runtime type2.NewStruct() and type2.NewNamed()
//    to create the following type:
//        type List struct { Elem int; Rest *List }
//    ReflectType will return a reflect.Type equivalent to:
//        struct { Elem int; Rest interface{} }
//    i.e. the type name will be missing due to limitation 1 above,
//    and the field 'Rest' will have type interface{} instead of *List due to limitation 5.
func (t Type) ReflectType() reflect.Type {
	return t[0].ReflectType()
}

func (t Type) UnsafeForceReflectType(rtype reflect.Type) {
	t[0].UnsafeForceReflectType(rtype)
}

// Size returns the number of bytes needed to store
// a value of the given type; it is analogous to unsafe.Sizeof.
func (t Type) Size() uintptr {
	return t[0].Size()
}

// String returns a string representation of a type.
func (t Type) String() string {
	if t == nil {
		return "<nil>"
	}
	return t[0].String()
}

// AddMethod adds method with given name and signature to type, unless it is already in the method list.
// It panics if the type is unnamed, or if the signature is not a function-with-receiver type.
// Returns the method index, or < 0 in case of errors
func (t Type) AddMethod(name string, signature Type) int {
	return t[0].AddMethod(name, signature)
}

// Bits returns the size of the type in bits.
// It panics if the type's Kind is not one of the
// sized or unsized Int, Uint, Float, or Complex kinds.
func (t Type) Bits() int {
	return t[0].Bits()
}

// ChanDir returns a channel type's direction.
// It panics if the type's Kind is not Chan.
func (t Type) ChanDir() reflect.ChanDir {
	return t[0].ChanDir()
}

// Complete marks an interface type as complete and computes wrapper methods for embedded fields.
// It must be called by users of InterfaceOf after the interface's embedded types are fully defined
// and before using the interface type in any way other than to form other types.
// Complete returns the receiver.
func (t Type) Complete() Type {
	return t[0].Complete()
}

// Elem returns a type's element type.
// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
func (t Type) Elem() Type {
	return t[0].Elem()
}

func (t Type) elem() Type {
	return t[0].elem()
}

// Field returns a struct type's i-th field.
// It panics if the type's Kind is not Struct.
// It panics if i is not in the range [0, NumField()).
func (t Type) Field(i int) StructField {
	return t[0].Field(i)
}

// FieldByName returns the (possibly embedded) struct field with the given name
// and the number of fields found at the same (shallowest) depth: 0 if not found.
// Private fields are returned only if they were declared in pkgpath.
func (t Type) FieldByName(name, pkgpath string) (field StructField, count int) {
	return t[0].FieldByName(name, pkgpath)
}

// IsMethod reports whether a function type's contains a receiver, i.e. is a method.
// If IsMethod returns true, the actual receiver type is available as the first parameter, i.e. Type.In(0)
// It panics if the type's Kind is not Func.
func (t Type) IsMethod() bool {
	return t[0].IsMethod()
}

// IsVariadic reports whether a function type's final input parameter is a "..." parameter.
// If so, t.In(t.NumIn() - 1) returns the parameter's implicit actual type []T.
// IsVariadic panics if the type's Kind is not Func.
func (t Type) IsVariadic() bool {
	return t[0].IsVariadic()
}

// Key returns a map type's key type.
// It panics if the type's Kind is not Map.
func (t Type) Key() Type {
	return t[0].Key()
}

// Kind returns the specific kind of the type.
func (t Type) Kind() reflect.Kind {
	if len(t) == 0 {
		return reflect.Invalid
	}
	return t[0].Kind()
}

// Len returns an array type's length.
// It panics if the type's Kind is not Array.
func (t Type) Len() int {
	return t[0].Len()
}

// In returns the type of a function type's i'th input parameter.
// It panics if the type's Kind is not Func.
// It panics if i is not in the range [0, NumIn()).
func (t Type) In(i int) Type {
	return t[0].In(i)
}

// For interfaces, Method returns the i-th method, including methods from embedded interfaces.
// For all other named types, Method returns the i-th explicitly declared method, ignoring wrapper methods for embedded fields.
// It panics if i is outside the range 0 .. NumMethod()-1
func (t Type) Method(i int) Method {
	return t[0].Method(i)
}

// MethodByName returns the method with given name (including wrapper methods for embedded fields)
// and the number of methods found at the same (shallowest) depth: 0 if not found.
// Private methods are returned only if they were declared in pkgpath.
func (t Type) MethodByName(name, pkgpath string) (method Method, count int) {
	return t[0].MethodByName(name, pkgpath)
}

// For interfaces, NumMethod returns *total* number of methods for interface t,
// including wrapper methods for embedded interfaces.
// For all other named types, NumMethod returns the number of explicitly declared methods,
// ignoring wrapper methods for embedded fields.
// Returns 0 for other unnamed types.
func (t Type) NumMethod() int {
	return t[0].NumMethod()
}

// NumExplicitMethod returns the number of explicitly declared methods of named type or interface t.
// Wrapper methods for embedded fields or embedded interfaces are not counted.
func (t Type) NumExplicitMethod() int {
	return t[0].NumExplicitMethod()
}

// NumMethod returns the *total* number of methods for interface or named type t,
// including wrapper methods for embedded fields or embedded interfaces.
// Note: it has slightly different semantics from go/types.(*Named).NumMethods(),
//       since the latter returns 0 for named interfaces, and callers need to manually invoke
//       goNamedType.Underlying().NumMethods() to retrieve the number of methods
//       of a named interface
func (t Type) NumAllMethod() int {
	return t[0].NumAllMethod()
}

// NumField returns a struct type's field count.
// It panics if the type's Kind is not Struct.
func (t Type) NumField() int {
	return t[0].NumField()
}

// NumIn returns a function type's input parameter count.
// It panics if the type's Kind is not Func.
func (t Type) NumIn() int {
	return t[0].NumIn()
}

// NumOut returns a function type's output parameter count.
// It panics if the type's Kind is not Func.
func (t Type) NumOut() int {
	return t[0].NumOut()
}

// Out returns the type of a function type's i'th output parameter.
// It panics if the type's Kind is not Func.
// It panics if i is not in the range [0, NumOut()).
func (t Type) Out(i int) Type {
	return t[0].Out(i)
}

// RemoveMethods removes given methods from type.
// It panics if the type is unnamed.
func (t Type) RemoveMethods(names []string, pkgpath string) {
	t[0].RemoveMethods(names, pkgpath)
}

// SetUnderlying sets the underlying type of a named type and marks it as complete.
// It panics if the type is unnamed, or if the underlying type is named,
// or if SetUnderlying() was already invoked on the named type.
func (t Type) SetUnderlying(underlying Type) {
	t[0].SetUnderlying(underlying)
}

// gunderlying returns the underlying types.Type of a type.
// TODO implement Underlying() Type ?
// Synthetizing the underlying reflect.Type is not possible for interface types,
// or for struct types with embedded or unexported fields.
func (t Type) gunderlying() types.Type {
	return t[0].gunderlying()
}

func (t Type) Universe() *Universe {
	return t[0].Universe()
}

// GetMethods returns the pointer to the method values.
// It panics if the type is unnamed
func (t Type) GetMethods() *[]reflect.Value {
	return t[0].GetMethods()
}

func wrap(t *xtype) Type {
	if t != nil {
		return Type{t}
	}
	return nil
}

func unwrap(t Type) *xtype {
	if len(t) != 0 {
		return t[0]
	}
	return nil
}
