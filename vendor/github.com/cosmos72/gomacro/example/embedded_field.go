package main

import (
	"fmt"
	"reflect"
	_ "reflect"
	"unsafe"
)

type Pair struct{ A, B int }

func pair(a, b int) Pair { var p Pair; p.A = a; p.B = b; return p }

type Triple struct {
	Pair
	C int
}

func (p Pair) First() int {
	return p.A
}

func (p Pair) Last() int {
	return p.B
}

func (t Triple) Last() int {
	return t.C
}

func embedded_field() {
	printChars()
	inspectTriple()
}

func printChars() {
	for i := 128; i <= 255; i++ {
		fmt.Printf("%x %c\n", i, i)
	}
}

func inspectTriple() {
	t := Triple{Pair{1, 2}, 3}
	inspect("declared:           ", t.Pair.First)
	inspect("declared2:          ", t.Pair.First)
	inspect("wrapped:            ", t.First)
	inspect("wrapped2:           ", t.First)
	inspect("declared  (reflect):", reflect.ValueOf(t.Pair).MethodByName("First").Interface().(func() int))
	inspect("declared2 (reflect):", reflect.ValueOf(t.Pair).MethodByName("First").Interface().(func() int))
	inspect("wrapped   (reflect):", reflect.ValueOf(t).MethodByName("First").Interface().(func() int))
	inspect("wrapped2  (reflect):", reflect.ValueOf(t).MethodByName("First").Interface().(func() int))
	fmt.Println()
	inspect("declared:             ", t.Pair.Last)
	inspect("declared2:            ", t.Pair.Last)
	inspect("overridden:           ", t.Last)
	inspect("overridden2:          ", t.Last)
	inspect("declared    (reflect):", reflect.ValueOf(t.Pair).MethodByName("Last").Interface().(func() int))
	inspect("declared2   (reflect):", reflect.ValueOf(t.Pair).MethodByName("Last").Interface().(func() int))
	inspect("overridden  (reflect):", reflect.ValueOf(t).MethodByName("Last").Interface().(func() int))
	inspect("overridden2 (reflect):", reflect.ValueOf(t).MethodByName("Last").Interface().(func() int))
	fmt.Println()
	/*
		inspectMethod1(t, "First")
		inspectMethod1(t.Pair, "First")
		inspectMethod1(t, "Last")
		inspectMethod1(t.Pair, "Last")
	*/
}

func inspect(name string, x func() int) {
	u := *(**uintptr)(unsafe.Pointer(&x))
	u4 := *(**[16]uintptr)(unsafe.Pointer(&x))
	fmt.Printf("%s %#v %#v %#v\n", name, x, u, *u4)
}

type UnsafeValue struct {
	typ  *uintptr
	ptr  unsafe.Pointer
	flag uintptr
}
type UnsafeInterface struct {
	typ *uintptr
	ptr unsafe.Pointer
}

func inspectMethod(x interface{}, y interface{}, name string) {
	mtd1 := reflect.ValueOf(x).MethodByName(name).Interface()
	mtd2 := reflect.ValueOf(y).MethodByName(name).Interface()

	inspectMethod1(mtd1, name)
	inspectMethod1(mtd2, name)
}

func inspectMethod1(mtd interface{}, name string) {
	fmt.Printf("%s:\t%#v\n", name, mtd)
	uf := *(*UnsafeInterface)(unsafe.Pointer(&mtd))
	fmt.Printf("%s:\tInterface = %#v\n", name, uf)
	fptr := (**uintptr)(uf.ptr)
	fmt.Printf("%s:\tInterface.ptr = %#v %#v\n", name, fptr, *fptr)
}

func inspectMethod2(x interface{}, name string) {
	t := reflect.TypeOf(x)
	mtd, _ := t.MethodByName(name)
	fmt.Printf("%s:\t%v\n", name, mtd)
	var f reflect.Value = mtd.Func
	uf := *(*UnsafeValue)(unsafe.Pointer(&f))
	fmt.Printf("%s:\tFunc = %#v\n", name, uf)
	fptr := (**uintptr)(uf.ptr)
	fmt.Printf("%s:\tFunc.ptr = %#v %#v %#v\n", name, fptr, *fptr, **fptr)
}
