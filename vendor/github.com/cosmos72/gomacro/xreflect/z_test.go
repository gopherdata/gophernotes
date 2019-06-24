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
 * z_test.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/token"
	"io"
	"os"
	r "reflect"
	"testing"
	"time"

	"github.com/cosmos72/gomacro/go/etoken"

	"github.com/cosmos72/gomacro/go/types"
	"github.com/cosmos72/gomacro/go/typeutil"
)

var u = NewUniverse()

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
	if !r.DeepEqual(actual, expected) {
		fail2(t, actual, expected)
	}
}

func isfieldequal(t *testing.T, actual StructField, expected StructField) {
	is(t, actual.Name, expected.Name)
	is(t, actual.Pkg, expected.Pkg)
	if !actual.Type.IdenticalTo(expected.Type) {
		fail(t, actual.Type, expected.Type)
	}
	is(t, actual.Tag, expected.Tag)
	is(t, actual.Offset, expected.Offset)
	isdeepequal(t, actual.Index, expected.Index)
	is(t, actual.Anonymous, expected.Anonymous)
}

func isidenticalgotype(t *testing.T, actual types.Type, expected types.Type) {
	if !typeutil.Identical(actual, expected) {
		fail(t, actual, expected)
	}
}

func istypeof(t *testing.T, actual interface{}, expected interface{}) {
	is(t, r.TypeOf(actual), r.TypeOf(expected))
}

func TestCanonical(t *testing.T) {
	type testCase struct {
		expected, actual r.Type
	}
	tint := r.TypeOf(int(0))
	tints := []r.Type{tint}
	funcIntInt := func(x int) int {
		return x
	}
	fieldsAB := []r.StructField{
		{Name: "A", Type: tint},
		{Name: "B", Type: tint},
	}
	fieldsXY := []r.StructField{
		{Name: "X", Type: r.TypeOf((*error)(nil)).Elem()},
		{Name: "Y", Type: r.TypeOf((*Forward)(nil)).Elem()},
	}

	testCases := []testCase{
		{r.TypeOf([0]int{}), r.ArrayOf(0, tint)},
		{r.TypeOf(make(chan int)), r.ChanOf(r.BothDir, tint)},
		{r.TypeOf(funcIntInt), r.FuncOf(tints, tints, false)},
		{r.TypeOf(make(map[int]int)), r.MapOf(tint, tint)},
		{r.TypeOf((*int)(nil)), r.PtrTo(tint)},
		{r.TypeOf([]int{}), r.SliceOf(tint)},
		{r.TypeOf(struct{ A, B int }{}), r.StructOf(fieldsAB)},
		{r.TypeOf(struct {
			X error
			Y Forward
		}{}), r.StructOf(fieldsXY)},
	}

	for _, tc := range testCases {
		is(t, tc.actual, tc.expected)
	}
}

func TestBasic(t *testing.T) {
	for i, rtype := range rbasictypes {
		if rtype == nil {
			continue
		}
		kind := r.Kind(i)
		typ := u.BasicTypes[kind]
		is(t, typ.Kind(), rtype.Kind())
		is(t, typ.Name(), rtype.Name())
		is(t, typ.ReflectType(), rtype)
		istypeof(t, typ.GoType(), (*types.Basic)(nil))

		basic := typ.GoType().Underlying().(*types.Basic)
		k := ToReflectKind(basic.Kind())
		is(t, k, rtype.Kind())
	}
}

func TestArray(t *testing.T) {
	typ := u.ArrayOf(7, u.BasicTypes[r.Uint8])
	rtype := r.TypeOf([7]uint8{})
	is(t, typ.Kind(), r.Array)
	is(t, typ.Name(), "")
	is(t, typ.ReflectType(), rtype)
	istypeof(t, typ.GoType(), (*types.Array)(nil))
	is(t, typ.String(), "[7]uint8")
}

func TestFunction(t *testing.T) {
	typ := u.FuncOf([]Type{u.BasicTypes[r.Bool], u.BasicTypes[r.Int16]}, []Type{u.BasicTypes[r.String]}, false)
	rtype := r.TypeOf(func(bool, int16) string { return "" })
	is(t, typ.Kind(), r.Func)
	is(t, typ.Name(), "")
	is(t, typ.ReflectType(), rtype)
	istypeof(t, typ.GoType(), (*types.Signature)(nil))
	is(t, typ.String(), "func(bool, int16) string")
}

func TestInterface1(t *testing.T) {
	methodtyp := u.FuncOf(nil, []Type{u.BasicTypes[r.Int]}, false)
	typ := u.InterfaceOf(nil, []string{"Cap", "Len"}, []Type{methodtyp, methodtyp}, nil).Complete()

	is(t, typ.Kind(), r.Interface)
	is(t, typ.Name(), "")
	is(t, typ.NumExplicitMethod(), 2)
	actual := typ.Method(0)
	is(t, actual.Name, "Cap")
	is(t, true, types.Identical(methodtyp.GoType(), actual.Type.GoType()))
	actual = typ.Method(1)
	is(t, actual.Name, "Len")
	is(t, true, types.Identical(methodtyp.GoType(), actual.Type.GoType()))
	istypeof(t, typ.GoType(), (*types.Interface)(nil))

	rtype := r.PtrTo(
		r.StructOf([]r.StructField{
			approxInterfaceHeader(),
			r.StructField{Name: "Cap", Type: methodtyp.ReflectType()},
			r.StructField{Name: "Len", Type: methodtyp.ReflectType()},
		}))
	is(t, typ.ReflectType(), rtype)
	is(t, typ.String(), "interface{Cap() int; Len() int}")
}

// test implementing 'error' interface
func TestInterfaceError(t *testing.T) {
	methodtyp := u.FuncOf(nil, []Type{u.BasicTypes[r.String]}, false)
	typ := u.InterfaceOf(nil, []string{"Error"}, []Type{methodtyp}, nil).Complete()

	is(t, typ.Kind(), r.Interface)
	is(t, typ.Name(), "")
	is(t, typ.NumExplicitMethod(), 1)
	is(t, typ.NumAllMethod(), 1)

	methodtyp = typ.Method(0).Type
	is(t, methodtyp.NumIn(), 1) // one input parameter: the method receiver

	is(t, typ.Implements(u.TypeOfError), true)
}

func TestMap(t *testing.T) {
	typ := u.MapOf(u.TypeOfInterface, u.BasicTypes[r.Bool])
	rtype := r.TypeOf(map[interface{}]bool{})
	is(t, typ.Kind(), r.Map)
	is(t, typ.Name(), "")
	is(t, typ.ReflectType(), rtype)
	is(t, rtype.NumMethod(), 0)
	if etoken.GENERICS_V2_CTI {
		is(t, typ.NumMethod(), 5)
		is(t, typ.NumAllMethod(), 5)
	} else {
		is(t, typ.NumMethod(), 0)
		is(t, typ.NumAllMethod(), 0)
	}
	istypeof(t, typ.GoType(), (*types.Map)(nil))
}

func TestIntMethod(t *testing.T) {
	typ := u.NamedOf("MyInt", "main")
	typ.SetUnderlying(u.BasicTypes[r.Int])
	rtype := r.TypeOf(int(0))
	is(t, typ.Kind(), r.Int)
	is(t, typ.Name(), "MyInt")
	is(t, typ.ReflectType(), rtype)
	if etoken.GENERICS_V2_CTI {
		is(t, typ.NumMethod(), 16)
		is(t, typ.NumAllMethod(), 32)
	} else {
		is(t, typ.NumMethod(), 0)
		is(t, typ.NumAllMethod(), 0)
	}
	istypeof(t, typ.GoType(), (*types.Named)(nil))
}

func TestNamed(t *testing.T) {
	tkey := u.TypeOfInterface
	tval := u.BasicTypes[r.Bool]
	underlying := u.MapOf(tkey, tval)

	typ := u.NamedOf("MyMap", "main")
	typ.SetUnderlying(underlying)
	rtype := r.TypeOf(map[interface{}]bool{})
	is(t, typ.Kind(), r.Map)
	is(t, typ.Name(), "MyMap")
	is(t, typ.ReflectType(), rtype)
	is(t, rtype.NumMethod(), 0)
	if etoken.GENERICS_V2_CTI {
		is(t, typ.NumMethod(), 5)
		is(t, typ.NumAllMethod(), 10)

		m, count := typ.MethodByName("Index", "")
		is(t, count, 1)
		is(t, m.Name, "Index")
		is(t, m.Pkg, (*Package)(nil))
		is(t, m.Type.ReflectType(), r.TypeOf(func(map[interface{}]bool, interface{}) bool {
			return false
		}))

		newvar := func(t Type) *types.Var {
			return types.NewVar(token.NoPos, nil, "", t.GoType())
		}
		isidenticalgotype(t,
			m.Type.GoType(),
			types.NewSignature(
				newvar(typ),
				types.NewTuple(newvar(tkey)),
				types.NewTuple(newvar(tval)),
				false,
			),
		)

		m, count = typ.MethodByName("TryIndex", "")
		is(t, count, 1)
		is(t, m.Name, "TryIndex")
		is(t, m.Pkg, (*Package)(nil))
		is(t, m.Type.ReflectType(), r.TypeOf(func(map[interface{}]bool, interface{}) (bool, bool) {
			return false, false
		}))

	} else {
		is(t, typ.NumMethod(), 0)
		is(t, typ.NumAllMethod(), 0)
	}
	istypeof(t, typ.GoType(), (*types.Named)(nil))
}

func TestSelfReference(t *testing.T) {
	typ := u.NamedOf("List", "main")

	is(t, typ.Kind(), r.Invalid)
	isidenticalgotype(t, typ.gunderlying(), u.TypeOfForward.GoType())

	underlying := u.StructOf([]StructField{
		StructField{Name: "First", Type: u.BasicTypes[r.Int]},
		StructField{Name: "Rest", Type: typ},
	})
	typ.SetUnderlying(underlying)
	is(t, typ.Kind(), r.Struct)
	typ1 := typ.Field(1).Type
	rtype := r.TypeOf(struct {
		First int
		Rest  Forward
	}{})

	is(t, typ.Kind(), r.Struct)
	is(t, typ.Name(), "List")
	istypeof(t, typ.GoType(), (*types.Named)(nil))
	is(t, typ.ReflectType(), rtype)
	is(t, typ.NumAllMethod(), rtype.NumMethod())
	is(t, typ1.ReflectType(), rTypeOfForward)         // Rest is actually xreflect.Incomplete
	isidenticalgotype(t, typ1.GoType(), typ.GoType()) // but it must pretend to be a main.List

	is(t, typ.String(), "main.List")
	is(t, typ.gunderlying().String(), "struct{First int; Rest main.List}")
}

func TestStruct(t *testing.T) {
	typ := u.StructOf([]StructField{
		StructField{Name: "First", Type: u.BasicTypes[r.Int]},
		StructField{Name: "Rest", Type: u.TypeOfInterface},
	})
	rtype := r.TypeOf(struct {
		First int
		Rest  interface{}
	}{})
	is(t, typ.Kind(), r.Struct)
	is(t, typ.Name(), "")
	is(t, typ.ReflectType(), rtype)
	istypeof(t, typ.GoType(), (*types.Struct)(nil))
	is(t, typ.NumField(), rtype.NumField())
	is(t, typ.NumAllMethod(), rtype.NumMethod())
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		rfield1 := field.toReflectField(false)
		rfield2 := rtype.Field(i)
		isdeepequal(t, rfield1, rfield2)
	}
	is(t, typ.String(), "struct{First int; Rest interface{}}")
}

func TestEmbedded(t *testing.T) {
	etyp := u.NamedOf("Box", "")
	etyp.SetUnderlying(u.StructOf([]StructField{
		StructField{Name: "Value", Type: u.BasicTypes[r.Int]},
	}))
	ertype := r.TypeOf(struct {
		Value int
	}{})
	is(t, etyp.Kind(), r.Struct)
	is(t, etyp.Name(), "Box")
	is(t, etyp.ReflectType(), ertype)
	istypeof(t, etyp.GoType(), (*types.Named)(nil))
	istypeof(t, etyp.GoType().Underlying(), (*types.Struct)(nil))

	typ := u.StructOf([]StructField{
		StructField{Name: "Label", Type: u.BasicTypes[r.String]},
		StructField{Type: u.PtrTo(etyp)}, // empty name => anonymous, and autodetect name from type
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
	isfieldequal(t, field, efield)

	// access anonymous field Struct.Box
	field, count = typ.FieldByName("Box", "")
	is(t, count, 1)
	isfieldequal(t, field, typ.Field(1))
}

func TestFromReflect0(t *testing.T) {
	rtype := r.TypeOf((*func(bool, int8, <-chan uint16, []float32, [2]float64, []complex64) map[interface{}]*string)(nil)).Elem()
	v := NewUniverse()
	v.RebuildDepth = MaxDepth
	typ := v.FromReflectType(rtype)
	is(t, typ.ReflectType(), rtype) // recreated 100% accurately?
}

func TestFromReflect1(t *testing.T) {
	rtype := r.TypeOf(time.Duration(0))
	typ := u.FromReflectType(rtype)
	is(t, typ.ReflectType(), rtype)
	is(t, typ.String(), "time.Duration")
	is(t, typ.gunderlying().String(), "int64")
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
	in := r.TypeOf(Bag{})
	expected := r.TypeOf(struct {
		C <-chan bool
		I int32
		U uintptr
		F [3]float32
		G []float64
		M map[string]*complex64
	}{})
	v := NewUniverse()
	v.RebuildDepth = MaxDepth
	typ := v.FromReflectType(in)
	actual := typ.ReflectType()
	is(t, typ.Kind(), r.Struct)
	is(t, typ.Name(), "Bag")
	is(t, actual, expected)
	is(t, actual.ConvertibleTo(in), true)
	is(t, in.ConvertibleTo(actual), true)
	is(t, actual.AssignableTo(in), true)
	is(t, in.AssignableTo(actual), true)
}

func TestFromReflect3(t *testing.T) {
	rtype := r.TypeOf((*io.Reader)(nil)).Elem()
	v := NewUniverse()
	v.RebuildDepth = 2
	typ := v.FromReflectType(rtype)

	actual := typ.ReflectType()
	expected := r.PtrTo(
		r.StructOf([]r.StructField{
			approxInterfaceHeader(),
			r.StructField{Name: "Read", Type: r.TypeOf((*func([]uint8) (int, error))(nil)).Elem()},
		}))
	is(t, typ.Kind(), r.Interface)
	is(t, actual, expected)
	is(t, typ.String(), "io.Reader")
	is(t, typ.gunderlying().String(), "interface{Read([]uint8) (int, error)}")
	is(t, typ.NumExplicitMethod(), 1)
	is(t, typ.NumAllMethod(), 1)
	is(t, rtype.NumMethod(), 1)

	for depth := 0; depth <= 3; depth++ {
		v := NewUniverse()
		v.RebuildDepth = depth
		typ = v.FromReflectType(rtype)
		// debugf("%v\t-> %v", typ, typ.ReflectType())
	}
}

func TestFromReflect4(t *testing.T) {
	type ToString func() string
	rtype := r.PtrTo(
		r.StructOf([]r.StructField{
			approxInterfaceHeader(),
			r.StructField{Name: "String", Type: r.TypeOf((*ToString)(nil)).Elem()},
		}))
	typ := u.NamedOf("Stringer", "io")
	v := NewUniverse()
	v.RebuildDepth = MaxDepth
	underlying := v.FromReflectType(rtype)
	typ.SetUnderlying(underlying)

	actual := typ.ReflectType()
	expected := r.PtrTo(
		r.StructOf([]r.StructField{
			approxInterfaceHeader(),
			r.StructField{Name: "String", Type: r.TypeOf((*func() string)(nil)).Elem()},
		}))
	is(t, typ.Kind(), r.Interface)
	is(t, actual, expected)
	is(t, typ.NumExplicitMethod(), 1)
	is(t, typ.NumAllMethod(), 1)
	is(t, typ.String(), "io.Stringer")
	is(t, typ.gunderlying().String(), "interface{String() string}")
	/*
		for depth := 0; depth <= 3; depth++ {
			v := NewUniverse()
			v.RebuildDepth = depth
			typ = v.FromReflectType(rtype)
			// debugf("%v\t-> %v", typ, typ.ReflectType())
		}
	*/
}

type myType interface {
	Elem() myType
	Field(int) myStructField
}
type myStructField struct {
	Name  string
	Type  myType
	Index []int
}

var myPkg = r.TypeOf(myStructField{}).PkgPath()

func TestFromReflect5(t *testing.T) {
	if testing.Verbose() {
		defer de(bug(u))
	}
	rtype := r.TypeOf((*myType)(nil)).Elem()
	typ := u.FromReflectType(rtype)

	is(t, typ.String(), myPkg+".myType")

	is(t, typ.NumExplicitMethod(), rtype.NumMethod())
	is(t, typ.NumAllMethod(), rtype.NumMethod())
}

func TestFromReflect6(t *testing.T) {
	if testing.Verbose() {
		defer de(bug(u))
	}
	rtype := r.TypeOf((*r.Type)(nil)).Elem()
	typ := u.FromReflectType(rtype)

	is(t, typ.String(), "reflect.Type")

	// importer is more accurate and gives even function param names... accept both variants
	s1 := "interface{Align() int; AssignableTo(reflect.Type) bool; Bits() int; ChanDir() reflect.ChanDir; Comparable() bool; ConvertibleTo(reflect.Type) bool; Elem() reflect.Type; Field(int) reflect.StructField; FieldAlign() int; FieldByIndex([]int) reflect.StructField; FieldByName(string) (reflect.StructField, bool); FieldByNameFunc(func(string) bool) (reflect.StructField, bool); Implements(reflect.Type) bool; In(int) reflect.Type; IsVariadic() bool; Key() reflect.Type; Kind() reflect.Kind; Len() int; Method(int) reflect.Method; MethodByName(string) (reflect.Method, bool); Name() string; NumField() int; NumIn() int; NumMethod() int; NumOut() int; Out(int) reflect.Type; PkgPath() string; Size() uintptr; String() string; common() *reflect.rtype; uncommon() *reflect.uncommonType}"
	s2 := "interface{Align() int; AssignableTo(u reflect.Type) bool; Bits() int; ChanDir() reflect.ChanDir; Comparable() bool; ConvertibleTo(u reflect.Type) bool; Elem() reflect.Type; Field(i int) reflect.StructField; FieldAlign() int; FieldByIndex(index []int) reflect.StructField; FieldByName(name string) (reflect.StructField, bool); FieldByNameFunc(match func(string) bool) (reflect.StructField, bool); Implements(u reflect.Type) bool; In(i int) reflect.Type; IsVariadic() bool; Key() reflect.Type; Kind() reflect.Kind; Len() int; Method(int) reflect.Method; MethodByName(string) (reflect.Method, bool); Name() string; NumField() int; NumIn() int; NumMethod() int; NumOut() int; Out(i int) reflect.Type; PkgPath() string; Size() uintptr; String() string; common() *reflect.rtype; uncommon() *reflect.uncommonType}"
	su := typ.gunderlying().String()

	if su != s1 && su != s2 {
		is(t, su, s1)
	}
	is(t, typ.NumExplicitMethod(), rtype.NumMethod())
	is(t, typ.NumAllMethod(), rtype.NumMethod())
}

func TestFromReflect7(t *testing.T) {
	tfunc := u.FuncOf(nil, []Type{u.BasicTypes[r.Int]}, false)
	rtfunc := r.TypeOf((*func() int)(nil)).Elem()
	is(t, tfunc.String(), "func() int")
	is(t, tfunc.ReflectType(), rtfunc)

	rtype := r.TypeOf((*interface{ Len() int })(nil)).Elem()
	typ := u.FromReflectType(rtype)
	is(t, typ.String(), "interface{Len() int}")

	is(t, tfunc.ReflectType(), rtfunc)
}

type Request4Test struct {
	Header   map[string]string
	Response *Response4Test
}
type Response4Test struct {
	HttpStatus int
	Request    *Request4Test
}

func TestFromReflectMutualRecursion(t *testing.T) {
	if testing.Verbose() {
		defer de(bug(u))
	}
	rtype1 := r.TypeOf(Request4Test{})
	rtype2 := r.TypeOf(Response4Test{})

	typ1 := u.FromReflectType(rtype1)
	typ2 := typ1.Field(1).Type.Elem()
	typ1_loop := typ2.Field(1).Type.Elem()

	is(t, typ1.ReflectType(), rtype1)
	is(t, typ2.ReflectType(), rtype2)
	is(t, typ1_loop.ReflectType(), rtype1)
	is(t, typ1.Name(), "Request4Test")
	is(t, typ2.Name(), "Response4Test")
	isidenticalgotype(t, typ1.GoType(), typ1_loop.GoType())

	is(t, typ1.gunderlying().String(), "struct{Header map[string]string; Response *"+myPkg+".Response4Test}")
	is(t, typ2.gunderlying().String(), "struct{HttpStatus int; Request *"+myPkg+".Request4Test}")
}

// test implementing 'io.Reader' interface
func TestInterfaceIoReader(t *testing.T) {
	if testing.Verbose() {
		defer de(bug(u))
	}
	typ := makeIoReaderType()
	gtyp := typ.GoType()

	is(t, typ.Kind(), r.Interface)
	is(t, typ.Name(), "")
	is(t, typ.NumExplicitMethod(), 1)
	is(t, typ.NumAllMethod(), 1)

	// ---------------------------
	treader := u.TypeOf((*io.Reader)(nil)).Elem()

	is(t, treader.Kind(), r.Interface)
	is(t, treader.Name(), "Reader")
	is(t, treader.NumExplicitMethod(), 1)
	is(t, treader.NumAllMethod(), 1)

	istrue(t, typ.Implements(treader))
	istrue(t, typ.AssignableTo(treader))
	istrue(t, treader.AssignableTo(typ))
	istrue(t, types.Identical(gtyp, treader.GoType().Underlying()))
}

// return the Type equivalent to "type io.Reader interface { Read([]uint8) (int, error) }"
func makeIoReaderType() Type {
	u.RebuildDepth = 0

	in := []Type{u.SliceOf(u.BasicTypes[r.Uint8])}
	out := []Type{u.BasicTypes[r.Int], u.TypeOfError}
	methodtyp := u.FuncOf(in, out, false)
	return u.InterfaceOf(nil, []string{"Read"}, []Type{methodtyp}, nil).Complete()
}

func TestImporter(t *testing.T) {
	// ---------------------------
	io, err := u.Importer.Import("io")
	if err != nil {
		t.Skipf("Go toolchain not detected or not available, skipping Importer test: %v", err)
		return
	}
	istrue(t, io != nil)

	reader := io.Scope().Lookup("Reader").Type().(*types.Named)
	ireader := reader.Underlying().(*types.Interface)

	is(t, reader.Obj().Name(), "Reader")
	is(t, reader.NumMethods(), 0) // method Read() is declared in the interface, not in the named type
	is(t, ireader.NumMethods(), 1)

	gtyp := makeIoReaderType().GoType()

	istrue(t, types.Implements(gtyp, ireader))
	istrue(t, types.Identical(gtyp, ireader))
	istrue(t, types.AssignableTo(gtyp, reader))
	istrue(t, types.AssignableTo(reader, gtyp))

	// ---------------------------
	t_file := u.TypeOf((*os.File)(nil))
	tfile := t_file.Elem()

	os, err := u.Importer.Import("os")
	if err != nil {
		t.Skipf("Go toolchain not detected or not available, skipping Importer test: %v", err)
		return
	}
	istrue(t, os != nil)

	file := os.Scope().Lookup("File").Type().(*types.Named)

	tfileMethod, _ := tfile.MethodByName("Read", "")
	var fileMethod *types.Func
	for i, n := 0, file.NumMethods(); i < n; i++ {
		fileMethod = file.Method(i)
		if fileMethod.Name() == "Read" {
			break
		}
	}
	is(t, fileMethod.Name(), "Read")

	tfileRead := tfileMethod.Type.GoType().(*types.Signature)
	fileRead := fileMethod.Type().(*types.Signature)
	ireaderRead := ireader.ExplicitMethod(0).Type().(*types.Signature)

	if false {
		inspect("error", types.Universe.Lookup("error").Type())
		inspect("Universe.TypeOfError.GoType()", u.TypeOfError.GoType())
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

// return the Type equivalent to "type io.ReadWriter interface { io.Reader, io.Writer }"
func makeIoReaderWriterType() Type {
	in := []Type{u.SliceOf(u.BasicTypes[r.Uint8])}
	out := []Type{u.BasicTypes[r.Int], u.TypeOfError}
	method := u.FuncOf(in, out, false)
	read_interf := u.InterfaceOf(nil, []string{"Read"}, []Type{method}, nil).Complete()
	reader := u.NamedOf("Reader", "io")
	reader.SetUnderlying(read_interf)
	write_interf := u.InterfaceOf(nil, []string{"Write"}, []Type{method}, nil).Complete()
	writer := u.NamedOf("Writer", "io")
	writer.SetUnderlying(write_interf)
	rw_interf := u.InterfaceOf(nil, nil, nil, []Type{reader, writer}).Complete()
	readwriter := u.NamedOf("ReadWriter", "io")
	readwriter.SetUnderlying(rw_interf)
	return readwriter
}

// test implementing 'io.ReadWriter' interface
func TestInterfaceIoReadWriter(t *testing.T) {
	rw := makeIoReaderWriterType()

	is(t, rw.NumExplicitMethod(), 0)
	is(t, rw.NumAllMethod(), 2)

	m, count := rw.MethodByName("Read", "")
	is(t, count, 1)
	is(t, m.Name, "Read")
	is(t, m.Type.NumIn(), 2) // receiver and []uint8
	is(t, m.Type.NumOut(), 2)

	// go/types.Type.String() does not show the receiver
	is(t, m.Type.GoType().String(), "func([]uint8) (int, error)")
	// instead xreflect.Type.String() also shows it
	is(t, m.Type.String(), "func (interface{Read([]uint8) (int, error)}).([]uint8) (int, error)")

	// isidenticalgotype(t, m.Type.In(0).GoType(), rw.gunderlying()) // changed in Go 1.12.beta1

	m, count = rw.MethodByName("Write", "")
	is(t, count, 1)
	is(t, m.Name, "Write")
	is(t, m.Type.NumIn(), 2) // receiver and []uint8
	is(t, m.Type.NumOut(), 2)

	// go/types.Type.String() does not show the receiver
	is(t, m.Type.GoType().String(), "func([]uint8) (int, error)")
	// instead xreflect.Type.String() also shows it
	is(t, m.Type.String(), "func (interface{Write([]uint8) (int, error)}).([]uint8) (int, error)")

	// isidenticalgotype(t, m.Type.In(0).GoType(), rw.gunderlying()) // changed in Go 1.12.beta1

	trw := u.TypeOf((*io.ReadWriter)(nil)).Elem()

	is(t, rw.ConvertibleTo(trw), true)
	is(t, trw.ConvertibleTo(rw), true)
	is(t, rw.AssignableTo(trw), true)
	is(t, trw.AssignableTo(rw), true)
	is(t, rw.Implements(trw), true)
	is(t, trw.Implements(rw), true)
	// named types have been redeclared... they cannot be identical
	is(t, rw.IdenticalTo(trw), false)
	is(t, trw.IdenticalTo(rw), false)
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
