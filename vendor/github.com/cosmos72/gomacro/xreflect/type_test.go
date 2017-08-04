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
 * type_test.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/types"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
)

var v = universe

func fail(t *testing.T, actual interface{}, expected interface{}) {
	t.Errorf("expecting %v <%T>, found %v <%T>\n", expected, expected, actual, actual)
}

func fail2(t *testing.T, actual interface{}, expected interface{}) {
	t.Errorf("expecting %#v <%T>,\n\tfound %#v <%T>\n", expected, expected, actual, actual)
}

func is(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		fail(t, actual, expected)
	}
}

func istrue(t *testing.T, actual bool) {
	if !actual {
		fail(t, actual, true)
	}
}

func isdeepequal(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		fail2(t, actual, expected)
	}
}

func istype(t *testing.T, actual interface{}, expected interface{}) {
	is(t, reflect.TypeOf(actual), reflect.TypeOf(expected))
}

func TestBasic(t *testing.T) {
	for i, rtype := range rbasictypes {
		if rtype == nil {
			continue
		}
		kind := reflect.Kind(i)
		typ := v.BasicTypes[kind]
		is(t, typ.Kind(), rtype.Kind())
		is(t, typ.Name(), rtype.Name())
		is(t, typ.ReflectType(), rtype)
		istype(t, typ.GoType(), (*types.Basic)(nil))

		basic := typ.GoType().(*types.Basic)
		k := ToReflectKind(basic.Kind())
		is(t, k, rtype.Kind())
	}
}

func TestArray(t *testing.T) {
	typ := ArrayOf(7, v.BasicTypes[reflect.Uint8])
	rtype := reflect.TypeOf([7]uint8{})
	is(t, typ.Kind(), reflect.Array)
	is(t, typ.Name(), "")
	is(t, typ.ReflectType(), rtype)
	istype(t, typ.GoType(), (*types.Array)(nil))
	is(t, typ.String(), "[7]uint8")
}

func TestFunction(t *testing.T) {
	typ := FuncOf([]Type{v.BasicTypes[reflect.Bool], v.BasicTypes[reflect.Int16]}, []Type{v.BasicTypes[reflect.String]}, false)
	rtype := reflect.TypeOf(func(bool, int16) string { return "" })
	is(t, typ.Kind(), reflect.Func)
	is(t, typ.Name(), "")
	is(t, typ.ReflectType(), rtype)
	istype(t, typ.GoType(), (*types.Signature)(nil))
	is(t, typ.String(), "func(bool, int16) string")
}

func TestInterface1(t *testing.T) {
	methodtyp := FuncOf(nil, []Type{v.BasicTypes[reflect.Int]}, false)
	typ := InterfaceOf([]string{"Cap", "Len"}, []Type{methodtyp, methodtyp}, nil).Complete()

	is(t, typ.Kind(), reflect.Interface)
	is(t, typ.Name(), "")
	is(t, typ.NumMethod(), 2)
	actual := typ.Method(0)
	is(t, actual.Name, "Cap")
	is(t, true, types.Identical(methodtyp.GoType(), actual.Type.GoType()))
	actual = typ.Method(1)
	is(t, actual.Name, "Len")
	is(t, true, types.Identical(methodtyp.GoType(), actual.Type.GoType()))
	istype(t, typ.GoType(), (*types.Interface)(nil))

	rtype := reflect.PtrTo(
		reflect.StructOf([]reflect.StructField{
			approxInterfaceHeader(),
			reflect.StructField{Name: "Cap", Type: methodtyp.ReflectType()},
			reflect.StructField{Name: "Len", Type: methodtyp.ReflectType()},
		}))
	is(t, typ.ReflectType(), rtype)
	is(t, typ.String(), "interface{Cap() int; Len() int}")
}

// test implementing 'error' interface
func TestInterfaceError(t *testing.T) {
	methodtyp := FuncOf(nil, []Type{v.BasicTypes[reflect.String]}, false)
	typ := InterfaceOf([]string{"Error"}, []Type{methodtyp}, nil).Complete()

	is(t, typ.Kind(), reflect.Interface)
	is(t, typ.Name(), "")
	is(t, typ.NumMethod(), 1)

	is(t, typ.Implements(v.TypeOfError), true)
}

func TestMap(t *testing.T) {
	typ := MapOf(v.TypeOfInterface, v.BasicTypes[reflect.Bool])
	rtype := reflect.TypeOf(map[interface{}]bool{})
	is(t, typ.Kind(), reflect.Map)
	is(t, typ.Name(), "")
	is(t, typ.ReflectType(), rtype)
	istype(t, typ.GoType(), (*types.Map)(nil))
}

func TestMethod(t *testing.T) {
	typ := v.NamedOf("MyInt", "main")
	typ.SetUnderlying(v.BasicTypes[reflect.Int])
	rtype := reflect.TypeOf(int(0))
	is(t, typ.Kind(), reflect.Int)
	is(t, typ.Name(), "MyInt")
	is(t, typ.ReflectType(), rtype)
	istype(t, typ.GoType(), (*types.Named)(nil))
}

func TestNamed(t *testing.T) {
	typ := v.NamedOf("MyMap", "main")
	underlying := MapOf(v.TypeOfInterface, v.BasicTypes[reflect.Bool])
	typ.SetUnderlying(underlying)
	rtype := reflect.TypeOf(map[interface{}]bool{})
	is(t, typ.Kind(), reflect.Map)
	is(t, typ.Name(), "MyMap")
	is(t, typ.ReflectType(), rtype)
	istype(t, typ.GoType(), (*types.Named)(nil))
}

func TestSelfReference(t *testing.T) {
	typ := v.NamedOf("List", "main")
	underlying := StructOf([]StructField{
		StructField{Name: "First", Type: v.BasicTypes[reflect.Int]},
		StructField{Name: "Rest", Type: typ},
	})
	typ.SetUnderlying(underlying)
	rtype := reflect.TypeOf(struct {
		First int
		Rest  interface{}
	}{})
	is(t, typ.Kind(), reflect.Struct)
	is(t, typ.Name(), "List")
	is(t, typ.ReflectType(), rtype)
	is(t, true, types.Identical(typ.Field(1).Type.GoType(), typ.GoType()))
	istype(t, typ.GoType(), (*types.Named)(nil))

	is(t, typ.String(), "main.List")
	is(t, typ.underlying().String(), "struct{First int; Rest main.List}")
}

func TestStruct(t *testing.T) {
	typ := StructOf([]StructField{
		StructField{Name: "First", Type: v.BasicTypes[reflect.Int]},
		StructField{Name: "Rest", Type: v.TypeOfInterface},
	})
	rtype := reflect.TypeOf(struct {
		First int
		Rest  interface{}
	}{})
	is(t, typ.Kind(), reflect.Struct)
	is(t, typ.Name(), "")
	is(t, typ.ReflectType(), rtype)
	istype(t, typ.GoType(), (*types.Struct)(nil))
	is(t, typ.NumField(), rtype.NumField())
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		rfield1 := field.toReflectField(false)
		rfield2 := rtype.Field(i)
		isdeepequal(t, rfield1, rfield2)
	}
	is(t, typ.String(), "struct{First int; Rest interface{}}")
}

func TestEmbedded(t *testing.T) {
	etyp := v.NamedOf("Box", "")
	etyp.SetUnderlying(StructOf([]StructField{
		StructField{Name: "Value", Type: v.BasicTypes[reflect.Int]},
	}))
	ertype := reflect.TypeOf(struct {
		Value int
	}{})
	is(t, etyp.Kind(), reflect.Struct)
	is(t, etyp.Name(), "Box")
	is(t, etyp.ReflectType(), ertype)
	istype(t, etyp.GoType(), (*types.Named)(nil))
	istype(t, etyp.GoType().Underlying(), (*types.Struct)(nil))

	typ := StructOf([]StructField{
		StructField{Name: "Label", Type: v.BasicTypes[reflect.String]},
		StructField{Type: v.PtrTo(etyp)}, // empty name => anonymous, and autodetect name from type
	})
	is(t, typ.String(), "struct{Label string; *Box}")
	field1 := typ.Field(1)
	is(t, field1.Name, "Box") // autodetected
	is(t, field1.Anonymous, true)

	// access field Struct.Value - shorthand for Struct.Box.Value
	field, count := typ.FieldByName("Value", "")
	is(t, count, 1)
	isdeepequal(t, field.Index, []int{1, 0})

	efield := etyp.Field(0)
	field.Index = efield.Index
	field.Offset = efield.Offset
	isdeepequal(t, field, efield)

	// access anonymous field Struct.Box
	field, count = typ.FieldByName("Box", "")
	is(t, count, 1)
	isdeepequal(t, field, typ.Field(1))
}

func TestFromReflect0(t *testing.T) {
	rtype := reflect.TypeOf((*func(bool, int8, <-chan uint16, []float32, [2]float64, []complex64) map[interface{}]*string)(nil)).Elem()
	v := &Universe{RebuildDepth: MaxDepth}
	typ := v.FromReflectType(rtype)
	is(t, typ.ReflectType(), rtype) // recreated 100% accurately?
}

func TestFromReflect1(t *testing.T) {
	rtype := reflect.TypeOf(time.Duration(0))
	typ := v.FromReflectType(rtype)
	is(t, typ.ReflectType(), rtype)
	is(t, typ.String(), "time.Duration")
	is(t, typ.underlying().String(), "int64")
}

func TestFromReflect2(t *testing.T) {
	type Bag struct {
		C <-chan bool
		I int32
		U uintptr
		F [3]float32
		G []float64
		M map[string]*complex64
	}
	in := reflect.TypeOf(Bag{})
	expected := reflect.TypeOf(struct {
		C <-chan bool
		I int32
		U uintptr
		F [3]float32
		G []float64
		M map[string]*complex64
	}{})
	v := &Universe{RebuildDepth: MaxDepth}
	typ := v.FromReflectType(in)
	actual := typ.ReflectType()
	is(t, typ.Kind(), reflect.Struct)
	is(t, typ.Name(), "Bag")
	is(t, actual, expected)
	is(t, actual.ConvertibleTo(in), true)
	is(t, in.ConvertibleTo(actual), true)
	is(t, actual.AssignableTo(in), true)
	is(t, in.AssignableTo(actual), true)
}

func TestFromReflect3(t *testing.T) {
	rtype := reflect.TypeOf((*io.Reader)(nil)).Elem()
	v := &Universe{RebuildDepth: 1}
	typ := v.FromReflectType(rtype)

	actual := typ.ReflectType()
	expected := reflect.PtrTo(
		reflect.StructOf([]reflect.StructField{
			approxInterfaceHeader(),
			reflect.StructField{Name: "Read", Type: reflect.TypeOf((*func([]uint8) (int, error))(nil)).Elem()},
		}))
	is(t, typ.Kind(), reflect.Interface)
	is(t, actual, expected)
	is(t, typ.String(), "io.Reader")
	is(t, typ.underlying().String(), "interface{Read([]uint8) (int, error)}")

	for depth := 0; depth <= 3; depth++ {
		v := &Universe{RebuildDepth: depth}
		typ = v.FromReflectType(rtype)
		// debugf("%v\t-> %v", typ, typ.ReflectType())
	}
}

func TestFromReflect4(t *testing.T) {
	type ToString func() string
	rtype := reflect.PtrTo(
		reflect.StructOf([]reflect.StructField{
			approxInterfaceHeader(),
			reflect.StructField{Name: "String", Type: reflect.TypeOf((*ToString)(nil)).Elem()},
		}))
	typ := v.NamedOf("Stringer", "io")
	v := &Universe{RebuildDepth: MaxDepth}
	underlying := v.FromReflectType(rtype)
	typ.SetUnderlying(underlying)

	actual := typ.ReflectType()
	expected := reflect.PtrTo(
		reflect.StructOf([]reflect.StructField{
			approxInterfaceHeader(),
			reflect.StructField{Name: "String", Type: reflect.TypeOf((*func() string)(nil)).Elem()},
		}))
	is(t, typ.Kind(), reflect.Interface)
	is(t, actual, expected)
	/*
		is(t, typ.String(), "io.Stringer")
		is(t, typ.underlying().String(), "interface{String() string}")

		for depth := 0; depth <= 3; depth++ {
			v := &Universe{RebuildDepth: depth}
			typ = v.FromReflectType(rtype)
			// debugf("%v\t-> %v", typ, typ.ReflectType())
		}
	*/
}

func TestFromReflect5(t *testing.T) {
	rtype := reflect.TypeOf((*reflect.Type)(nil)).Elem()
	typ := v.FromReflectType(rtype)

	is(t, typ.String(), "reflect.Type")

	// importer is more accurate and gives even function param names...
	s1 := "interface{Align() int; AssignableTo(reflect.Type) bool; Bits() int; ChanDir() reflect.ChanDir; Comparable() bool; ConvertibleTo(reflect.Type) bool; Elem() reflect.Type; Field(int) reflect.StructField; FieldAlign() int; FieldByIndex([]int) reflect.StructField; FieldByName(string) (reflect.StructField, bool); FieldByNameFunc(func(string) bool) (reflect.StructField, bool); Implements(reflect.Type) bool; In(int) reflect.Type; IsVariadic() bool; Key() reflect.Type; Kind() reflect.Kind; Len() int; Method(int) reflect.Method; MethodByName(string) (reflect.Method, bool); Name() string; NumField() int; NumIn() int; NumMethod() int; NumOut() int; Out(int) reflect.Type; PkgPath() string; Size() uintptr; String() string; common() *reflect.rtype; uncommon() *reflect.uncommonType}"
	s2 := "interface{Align() int; AssignableTo(u reflect.Type) bool; Bits() int; ChanDir() reflect.ChanDir; Comparable() bool; ConvertibleTo(u reflect.Type) bool; Elem() reflect.Type; Field(i int) reflect.StructField; FieldAlign() int; FieldByIndex(index []int) reflect.StructField; FieldByName(name string) (reflect.StructField, bool); FieldByNameFunc(match func(string) bool) (reflect.StructField, bool); Implements(u reflect.Type) bool; In(i int) reflect.Type; IsVariadic() bool; Key() reflect.Type; Kind() reflect.Kind; Len() int; Method(int) reflect.Method; MethodByName(string) (reflect.Method, bool); Name() string; NumField() int; NumIn() int; NumMethod() int; NumOut() int; Out(i int) reflect.Type; PkgPath() string; Size() uintptr; String() string; common() *reflect.rtype; uncommon() *reflect.uncommonType}"
	su := typ.underlying().String()

	if su != s1 && su != s2 {
		is(t, su, s1)
	}
}

// test implementing 'io.Reader' interface
func TestInterfaceIoReader(t *testing.T) {
	v.RebuildDepth = 0

	in := []Type{v.SliceOf(v.BasicTypes[reflect.Uint8])}
	out := []Type{v.BasicTypes[reflect.Int], v.TypeOfError}
	methodtyp := v.FuncOf(in, out, false)
	typ := InterfaceOf([]string{"Read"}, []Type{methodtyp}, nil).Complete()
	gtyp := typ.GoType()

	is(t, typ.Kind(), reflect.Interface)
	is(t, typ.Name(), "")
	is(t, typ.NumMethod(), 1)

	// ---------------------------
	treader := v.TypeOf((*io.Reader)(nil)).Elem()

	is(t, treader.Kind(), reflect.Interface)
	is(t, treader.Name(), "Reader")
	is(t, treader.NumMethod(), 1)

	istrue(t, typ.Implements(treader))
	istrue(t, typ.AssignableTo(treader))
	istrue(t, treader.AssignableTo(typ))
	istrue(t, types.Identical(gtyp, treader.GoType().Underlying()))

	// ---------------------------
	io, err := v.Importer.Import("io")
	istrue(t, err == nil)
	istrue(t, io != nil)

	reader := io.Scope().Lookup("Reader").Type().(*types.Named)
	ireader := reader.Underlying().(*types.Interface)

	is(t, reader.Obj().Name(), "Reader")
	is(t, reader.NumMethods(), 0) // method Read() is declared in the interface, not in the named type
	is(t, ireader.NumMethods(), 1)

	istrue(t, types.Implements(gtyp, ireader))
	istrue(t, types.Identical(gtyp, ireader))
	istrue(t, types.AssignableTo(gtyp, reader))
	istrue(t, types.AssignableTo(reader, gtyp))

	// ---------------------------
	t_file := v.TypeOf((*os.File)(nil))
	tfile := t_file.Elem()

	os, err := v.Importer.Import("os")
	istrue(t, err == nil)
	istrue(t, os != nil)

	file := os.Scope().Lookup("File").Type().(*types.Named)

	tfileRead := tfile.Method(6).Type.GoType().(*types.Signature)
	fileRead := file.Method(6).Type().(*types.Signature)
	ireaderRead := ireader.ExplicitMethod(0).Type().(*types.Signature)

	if false {
		inspect("error", types.Universe.Lookup("error").Type())
		inspect("Universe.TypeOfError.GoType()", v.TypeOfError.GoType())
		inspect("tfile.Read.Results.1.Type", tfileRead.Results().At(1).Type())
		inspect("file.Read.Results.1.Type", fileRead.Results().At(1).Type())
		inspect("ireader.Read.Results.1.Type", ireaderRead.Results().At(1).Type())
	}

	istrue(t, types.Identical(tfileRead, ireaderRead))
	istrue(t, types.Identical(fileRead, ireaderRead))
	istrue(t, types.Identical(tfileRead, fileRead))

	istrue(t, types.Implements(t_file.GoType(), ireader))
	istrue(t, types.AssignableTo(t_file.GoType(), reader))

}

func inspect(label string, t types.Type) {
	debugf("%s:\t%v", label, t)
	switch t := t.(type) {
	case *types.Named:
		debugf("  typename:\t%p\t%#v", t.Obj(), t.Obj())
		for i, n := 0, t.NumMethods(); i < n; i++ {
			debugf("    method %d:\t%s", i, t.Method(i))
		}
		debugf("  underlying:\t%v", t.Underlying())
	}
}
