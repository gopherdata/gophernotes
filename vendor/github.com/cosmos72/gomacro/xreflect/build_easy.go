// +build gomacro_xreflect_easy

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
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * build_easy.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/types"
	"reflect"
)

type Type interface {

	// Align returns the alignment in bytes of a value of
	// this type when allocated in memory.
	Align() int

	// FieldAlign returns the alignment in bytes of a value of
	// this type when used as a field in a struct.
	FieldAlign() int

	// AssignableTo reports whether a value of the type is assignable to type u.
	AssignableTo(u Type) bool

	// ConvertibleTo reports whether a value of the type is convertible to type u.
	ConvertibleTo(u Type) bool

	// Comparable reports whether values of this type are comparable.
	Comparable() bool

	// GoType returns the go/types.Type corresponding to the given type.
	GoType() types.Type

	// Implements reports whether the type implements the interface type u.
	// It panics if u's Kind is not Interface
	Implements(u Type) bool

	// Name returns the type's name within its package.
	// It returns an empty string for unnamed types.
	Name() string

	// Named returns whether the type is named.
	// It returns false for unnamed types.
	Named() bool

	// Pkg returns a named type's package, that is, the package where it was defined.
	// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
	// Pkg will return nil.
	Pkg() *Package

	// PkgName returns a named type's package name, that is,
	// the default name that the package provides when imported.
	// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
	// the package name will be the empty string.
	PkgName() string

	// PkgPath returns a named type's package path, that is, the import path
	// that uniquely identifies the package, such as "encoding/base64".
	// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
	// the package path will be the empty string.
	PkgPath() string

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
	ReflectType() reflect.Type

	UnsafeForceReflectType(rtype reflect.Type)

	// Size returns the number of bytes needed to store
	// a value of the given type; it is analogous to unsafe.Sizeof.
	Size() uintptr

	// String returns a string representation of a type.
	String() string

	// AddMethod adds method with given name and signature to type, unless it is already in the method list.
	// It panics if the type is unnamed, or if the signature is not a function-with-receiver type.
	// Returns the method index, or < 0 in case of errors
	AddMethod(name string, signature Type) int

	// Bits returns the size of the type in bits.
	// It panics if the type's Kind is not one of the
	// sized or unsized Int, Uint, Float, or Complex kinds.
	Bits() int

	// ChanDir returns a channel type's direction.
	// It panics if the type's Kind is not Chan.
	ChanDir() reflect.ChanDir

	// Complete marks an interface type as complete and computes wrapper methods for embedded fields.
	// It must be called by users of InterfaceOf after the interface's embedded types are fully defined
	// and before using the interface type in any way other than to form other types.
	// Complete returns a canonicalized (unique) version of the receiver.
	Complete() Type
	// Elem returns a type's element type.
	// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
	Elem() Type

	// Field returns a struct type's i-th field.
	// It panics if the type's Kind is not Struct.
	// It panics if i is not in the range [0, NumField()).
	Field(i int) StructField
	// FieldByName returns the (possibly embedded) struct field with the given name
	// and the number of fields found at the same (shallowest) depth: 0 if not found.
	// Private fields are returned only if they were declared in pkgpath.
	FieldByName(name, pkgpath string) (field StructField, count int)

	// IsMethod reports whether a function type's contains a receiver, i.e. is a method.
	// If IsMethod returns true, the actual receiver type is available as the first parameter, i.e. Type.In(0)
	// It panics if the type's Kind is not Func.
	IsMethod() bool

	// IsVariadic reports whether a function type's final input parameter is a "..." parameter.
	// If so, t.In(t.NumIn() - 1) returns the parameter's implicit actual type []T.
	// IsVariadic panics if the type's Kind is not Func.
	IsVariadic() bool

	// Key returns a map type's key type.
	// It panics if the type's Kind is not Map.
	Key() Type
	// Kind returns the specific kind of the type.
	Kind() reflect.Kind

	// Len returns an array type's length.
	// It panics if the type's Kind is not Array.
	Len() int

	// In returns the type of a function type's i'th input parameter.
	// It panics if the type's Kind is not Func.
	// It panics if i is not in the range [0, NumIn()).
	In(i int) Type
	// Method return the i-th explicitly declared method of named type or interface t.
	// Wrapper methods for embedded fields or embedded interfaces are not returned.
	// It panics if the type is unnamed, or if the type's Kind is not Interface
	Method(i int) Method
	// MethodByName returns the method with given name (including wrapper methods for embedded fields)
	// and the number of methods found at the same (shallowest) depth: 0 if not found.
	// Private methods are returned only if they were declared in pkgpath.
	MethodByName(name, pkgpath string) (method Method, count int)

	// NumMethod returns the number of explicitly declared methods of named type or interface t.
	// Wrapper methods for embedded fields or embedded interfaces are not counted.
	NumMethod() int
	// NumField returns a struct type's field count.
	// It panics if the type's Kind is not Struct.
	NumField() int

	// NumIn returns a function type's input parameter count.
	// It panics if the type's Kind is not Func.
	NumIn() int

	// NumOut returns a function type's output parameter count.
	// It panics if the type's Kind is not Func.
	NumOut() int

	// Out returns the type of a function type's i'th output parameter.
	// It panics if the type's Kind is not Func.
	// It panics if i is not in the range [0, NumOut()).
	Out(i int) Type

	// RemoveMethods removes given methods from type.
	// It panics if the type is unnamed, or if the signature is not a function type,
	RemoveMethods(names []string, pkgpath string)

	// SetUnderlying sets the underlying type of a named type and marks it as complete.
	// It panics if the type is unnamed, or if the underlying type is named,
	// or if SetUnderlying() was already invoked on the named type.
	SetUnderlying(underlying Type)

	// underlying returns the underlying types.Type of a type.
	// TODO implement Underlying() Type ?
	// Synthetizing the underlying reflect.Type is not possible for interface types,
	// or for struct types with embedded or unexported fields.
	underlying() types.Type

	elem() Type

	Universe() *Universe

	// GetMethods returns the pointer to the method values.
	// It panics if the type is unnamed
	GetMethods() *[]reflect.Value
}

func unwrap(t Type) *xtype {
	if t == nil {
		return nil
	}
	return t.(*xtype)
}

func wrap(t *xtype) Type {
	return t
}

// Complete marks an interface type as complete and computes wrapper methods for embedded fields.
// It must be called by users of InterfaceOf after the interface's embedded types are fully defined
// and before using the interface type in any way other than to form other types.
func (t *xtype) Complete() Type {
	if t.kind != reflect.Interface {
		xerrorf(t, "Complete of non-interface %v", t)
	}
	gtype := t.gtype.Underlying().(*types.Interface)
	gtype.Complete()
	return wrap(t)
}

var (
	BasicTypes = universe.BasicTypes

	TypeOfBool          = BasicTypes[reflect.Bool]
	TypeOfInt           = BasicTypes[reflect.Int]
	TypeOfInt8          = BasicTypes[reflect.Int8]
	TypeOfInt16         = BasicTypes[reflect.Int16]
	TypeOfInt32         = BasicTypes[reflect.Int32]
	TypeOfInt64         = BasicTypes[reflect.Int64]
	TypeOfUint          = BasicTypes[reflect.Uint]
	TypeOfUint8         = BasicTypes[reflect.Uint8]
	TypeOfUint16        = BasicTypes[reflect.Uint16]
	TypeOfUint32        = BasicTypes[reflect.Uint32]
	TypeOfUint64        = BasicTypes[reflect.Uint64]
	TypeOfUintptr       = BasicTypes[reflect.Uintptr]
	TypeOfFloat32       = BasicTypes[reflect.Float32]
	TypeOfFloat64       = BasicTypes[reflect.Float64]
	TypeOfComplex64     = BasicTypes[reflect.Complex64]
	TypeOfComplex128    = BasicTypes[reflect.Complex128]
	TypeOfString        = BasicTypes[reflect.String]
	TypeOfUnsafePointer = BasicTypes[reflect.UnsafePointer]
	TypeOfError         = universe.TypeOfError
	TypeOfInterface     = universe.TypeOfInterface
)

// TypeOf creates a Type corresponding to reflect.TypeOf() of given value.
// Note: conversions from Type to reflect.Type and back are not exact,
// because of the reasons listed in Type.ReflectType()
// Conversions from reflect.Type to Type and back are not exact for the same reasons.
func TypeOf(rvalue interface{}) Type {
	return universe.FromReflectType(reflect.TypeOf(rvalue))
}

// FromReflectType creates a Type corresponding to given reflect.Type
// Note: conversions from Type to reflect.Type and back are not exact,
// because of the reasons listed in Type.ReflectType()
// Conversions from reflect.Type to Type and back are not exact for the same reasons.
func FromReflectType(rtype reflect.Type) Type {
	return universe.FromReflectType(rtype)
}

// NamedOf returns a new named type for the given type name and package.
// Initially, the underlying type is set to interface{} - use SetUnderlying to change it.
// These two steps are separate to allow creating self-referencing types,
// as for example type List struct { Elem int; Rest *List }
func NamedOf(name, pkgpath string) Type {
	return universe.NamedOf(name, pkgpath)
}

func NewPackage(path, name string) *Package {
	return universe.NewPackage(path, name)
}

func MakeType(gtype types.Type, rtype reflect.Type) Type {
	return universe.MakeType(gtype, rtype)
}
