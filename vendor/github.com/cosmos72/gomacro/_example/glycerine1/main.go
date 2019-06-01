package main

import (
	"fmt"
	"reflect"

	"github.com/cosmos72/gomacro/fast"
	"github.com/cosmos72/gomacro/imports"
)

// this example shows how to efficiently share a variable named "x" of a fictional package "github.com/path/y"
// between gomacro interpreter and compiled code.
//
// there are two techniques - pick your favorite.
func main() {
	technique1()
	technique2()
}

// this initialization is needed by both techniques.
//
// A more complete example would instead generate a file x_package.go alongside the sources of github.com/path/y,
// to collect all exported symbols of the package "github.com/path/y", by executing once the commands:
//
// go install github.com/path/y
// gomacro --expr 'import _i "github.com/path/y"'
//
// That cannot be done here because it does not work for main packages.
func init() {
	imports.Packages["github.com/path/y"] = imports.Package{
		Binds:    map[string]reflect.Value{},
		Types:    map[string]reflect.Type{},
		Proxies:  map[string]reflect.Type{},
		Untypeds: map[string]string{},
		Wrappers: map[string][]string{},
	}
}

func technique1() {
	fmt.Println("sharing a variable between gomacro interpreter and compiled code - technique 1")

	interp := fast.New()

	// before importing the package, add to it the declaration of a global variable "x"
	//
	// Note that the injected `reflect.Value` must be settable and addressable,
	// hence the expression `reflect.ValueOf(&x).Elem()`
	//
	// The naive `reflect.ValueOf(x)` would result in a read-only, non-shared variable.
	var x int
	imports.Packages["github.com/path/y"].Binds["x"] = reflect.ValueOf(&x).Elem()

	// import the package and switch to it
	interp.ImportPackage("lname", "github.com/path/y")
	interp.ChangePackage("lname", "github.com/path/y")

	// use the variable shared with the interpreter
	for i := 0; i < 10; i++ {
		// ...then to update x inside gomacro each time (still host code)
		x = i

		// and use "x" normally inside the interpreter
		v, _ := interp.Eval1("x*x")
		fmt.Println("x*x = ", v.Int())
	}

}

func technique2() {
	fmt.Println("sharing a variable between gomacro interpreter and compiled code - technique 2")

	interp := fast.New()

	// import the package and switch to it
	interp.ImportPackage("lname", "github.com/path/y")
	interp.ChangePackage("lname", "github.com/path/y")

	// declare and initialize in the interpreter a variable "x"
	// whose initial value is a copy of `int(0)`
	// and whose type is inferred from the initial value
	interp.DeclVar("x", nil, int(0))

	// retrieve "x" address from the interpreter
	vo := interp.ValueOf("x")
	pXinsideGomacro := vo.Addr().Interface().(*int)

	// use the variable shared with the interpreter
	for i := 0; i < 10; i++ {
		// ...then to update x inside gomacro each time (still host code)
		*pXinsideGomacro = i

		// and use "x" normally inside the interpreter
		v, _ := interp.Eval1("x*x")
		fmt.Println("x*x = ", v.Int())
	}

}
