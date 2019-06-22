## gomacro - interactive Go interpreter and debugger with generics and macros

gomacro is an almost complete Go interpreter, implemented in pure Go. It offers both
an interactive REPL and a scripting mode, and does not require a Go toolchain at runtime
(except in one very specific case: import of a 3<sup>rd</sup> party package at runtime).

It has two dependencies beyond the Go standard library: github.com/peterh/liner and golang.org/x/sys

Gomacro can be used as:
* a standalone executable with interactive Go REPL, line editing and code completion:
  just run `gomacro` from your command line, then type Go code. Example:
    ```
    $ gomacro
    [greeting message...]

    gomacro> import "fmt"
    gomacro> fmt.Println("hello, world!")
    hello, world!
    14      // int
    <nil>   // error
    gomacro>
    ```
  press TAB to autocomplete a word, and press it again to cycle on possible completions.

  Line editing follows mostly Emacs: Ctrl+A or Home jumps to start of line,
  Ctrl+E or End jumps to end of line, Ald+D deletes word starting at cursor...
  For the full list of key bindings, see https://github.com/peterh/liner

* a tool to experiment with Go **generics**: see [Generics](#generics)

* a Go source code debugger: see [Debugger](#debugger)

* an interactive tool to make science more productive and more fun.
  If you use compiled Go with scientific libraries (physics, bioinformatics, statistics...)
  you can import the same libraries from gomacro REPL (immediate on Go 1.8+ and Linux
  or Go 1.10.2+ and Mac OS X, requires restarting on other platforms,
  see [Importing packages](#importing-packages) below), call them interactively,
  inspect the results, feed them to other functions/libraries, all in a single session.
  The imported libraries will be **compiled**, not interpreted,
  so they will be as fast as in compiled Go.

  For a graphical user interface on top of gomacro, see [Gophernotes](https://github.com/gopherdata/gophernotes).
  It is a Go kernel for Jupyter notebooks and nteract, and uses gomacro for Go code evaluation.

* a library that adds Eval() and scripting capabilities to your Go programs in few lines
  of code:
	```go
	package main
	import (
		"fmt"
		"reflect"
		"github.com/cosmos72/gomacro/fast"
	)
	func RunGomacro(toeval string) reflect.Value {
		interp := fast.New()
		vals, _ := interp.Eval(toeval)
		// for simplicity, only use the first returned value
		return vals[0]
	}
	func main() {
		fmt.Println(RunGomacro("1+1"))
	}
	```
  Also, [github issue #13](https://github.com/cosmos72/gomacro/issues/13) explains
  how to have your application's functions, variable, constants and types
  available in the interpreter.

  Note: gomacro license is [MPL 2.0](LICENSE), which imposes some restrictions
  on programs that use gomacro.
  See [MPL 2.0 FAQ](https://www.mozilla.org/en-US/MPL/2.0/FAQ/) for common questions
  regarding the license terms and conditions.

* a way to execute Go source code on-the-fly without a Go compiler:
  you can either run `gomacro FILENAME.go` (works on every supported platform)

  or you can insert a line `#!/usr/bin/env gomacro` at the beginning of a Go source file,
  then mark the file as executable with `chmod +x FILENAME.go` and finally execute it
  with `./FILENAME.go` (works only on Unix-like systems: Linux, *BSD, Mac OS X ...)

* a Go code generation tool:
  gomacro was started as an experiment to add Lisp-like macros to Go, and they are
  extremely useful (in the author's opinion) to simplify code generation.
  Macros are normal Go functions, they are special only in one aspect:
  they are executed **before** compiling code, and their input and output **is** code
  (abstract syntax trees, in the form of go/ast.Node)

  Don't confuse them with C preprocessor macros: in Lisp, Scheme and now in Go,
  macros are regular functions written in the same programming language
  as the rest of the source code. They can perform arbitrary computations
  and call any other function or library: they can even read and write files,
  open network connections, etc... as a normal Go function can do.

  See [doc/code_generation.pdf](https://github.com/cosmos72/gomacro/blob/master/doc/code_generation.pdf)
  for an introduction to the topic.

## Installation

### Prerequites

- [Go 1.9+](https://golang.org/doc/install)

### Supported platforms

Gomacro is pure Go, and in theory it should work on any platform supported by the Go compiler.
The following combinations are tested and known to work:

- Linux: amd64, 386, arm64, arm, mips, ppc64le
- Mac OS X: amd64, 386 (386 binaries running on amd64 system)
- Windows: amd64, 386
- FreeBSD: amd64, 386
- Android: arm64, arm (tested with [Termux](https://termux.com/) and the Go compiler distributed with it)

### How to install

  The command
  ```
  go get -u github.com/cosmos72/gomacro
  ```
  downloads, compiles and installs gomacro and its dependencies

## Current Status

Almost complete.

The main limitations and missing features are:

* importing 3<sup>rd</sup> party libraries at runtime currently only works on Linux and Mac OS X.
  On other systems as Windows, Android and *BSD it is cumbersome and requires recompiling - see [Importing packages](#importing-packages).
* some corner cases using interpreted interfaces, as interface -> interface type assertions and type switches, are not implemented yet.
* goto can only jump backward, not forward
* out-of-order code is under testing - some corner cases, as for example out-of-order declarations
  used in keys of composite literals, are not supported.
  Clearly, at REPL code is still executed as soon as possible, so it makes a difference mostly
  if you separate multiple declarations with ; on a single line. Example: `var a = b; var b = 42`\
  Support for "batch mode" is in progress - it reads as much source code as possible before executing it,
  and it's useful mostly to execute whole files or directories.

The [documentation](doc/) also contains the [full list of features and limitations](doc/features-and-limitations.md)

## Extensions

Compared to compiled Go, gomacro supports several extensions:

* generics (experimental) - see [Generics](#generics)

* an integrated debugger, see [Debugger](#debugger)

* configurable special commands. Type `:help` at REPL to list them,
  and see [cmd.go:37](https://github.com/cosmos72/gomacro/blob/master/fast/cmd.go#L37)
  for the documentation and API to define new ones.

* untyped constants can be manipulated directly at REPL. Examples:
    ```
	gomacro> 1<<100
	{int 1267650600228229401496703205376}	// untyped.Lit
	gomacro> const c = 1<<100; c * c / 100000000000
	{int 16069380442589902755419620923411626025222029937827}	// untyped.Lit
	```
  This provides a handy arbitrary-precision calculator.

  Note: operations on large untyped integer constants are always exact,
  while operations on large untyped float constants are implemented with `go/constant.Value`,
  and are exact as long as both numerator and denominator are <= 5e1232.

  Beyond that, `go/constant.Value` switches from `*big.Rat` to `*big.Float`
  with precision = 512, which can accumulate rounding errors.

  If you need **exact** results, convert the untyped float constant to `*big.Rat`
  (see next item) before exceeding 5e1232.

* untyped constants can be converted implicitly to `*big.Int`, `*big.Rat` and `*big.Float`. Examples:
    ```go
	import "math/big"
	var i *big.Int = 1<<1000                 // exact - would overflow int
	var r *big.Rat = 1.000000000000000000001 // exact - different from 1.0
	var s *big.Rat = 5e1232                  // exact - would overflow float64
	var t *big.Rat = 1e1234                  // approximate, exceeds 5e1232
	var f *big.Float = 1e646456992           // largest untyped float constant that is different from +Inf
    ```
  Note: every time such a conversion is evaluated, it creates a new value - no risk to modify the constant.

  Be aware that converting a huge value to string, as typing `f` at REPL would do, can be very slow.

* zero value constructors: for any type `T`, the expression `T()`
  returns the zero value of the type

* macros, quoting and quasiquoting: see
  [doc/code_generation.pdf](https://github.com/cosmos72/gomacro/blob/master/doc/code_generation.pdf)

and slightly relaxed checks:

* unused variables and unused return values never cause errors

## Examples

Some short, notable examples - to run them on non-Linux platforms, see [Importing packages](#importing-packages) first.

### plot mathematical functions

* install libraries: `go get gonum.org/v1/plot gonum.org/v1/plot/plotter gonum.org/v1/plot/vg`
* start the interpreter: `gomacro`
* at interpreter prompt, paste the whole Go code listed at https://github.com/gonum/plot/wiki/Example-plots#functions
  (the source code starts after the picture under the section "Functions", and ends just before the section "Histograms")
* still at interpreter prompt, enter `main()`
  If all goes well, it will create a file named "functions.png" in current directory containing the plotted functions.

### simple mandelbrot web server

* install libraries: `go get github.com/sverrirab/mandelbrot-go`
* chdir to mandelbrot-go source folder: `cd; cd go/src/github.com/sverrirab/mandelbrot-go`
* start interpreter with arguments: `gomacro -i mbrot.go`
* at interpreter prompt, enter `init(); main()`
* visit http://localhost:8090/
  Be patient, rendering and zooming mandelbrot set with an interpreter is a little slow.

Further examples are listed by [Gophernotes](https://github.com/gopherdata/gophernotes/#example-notebooks-dowload-and-run-them-locally-follow-the-links-to-view-in-github-or-use-the-jupyter-notebook-viewer)

## Importing packages

Gomacro supports the standard Go syntax `import`, including package renaming. Examples:
```go
import "fmt"
import (
    "io"
    "net/http"
    r "reflect"
)
```
Third party packages - i.e. packages not in Go standard library - can also be imported
with the same syntax, as long as the package is **already** installed.

To install a package, follow its installation procedure: quite often it is the command `go get PACKAGE-PATH`

The next steps depend on the system you are running gomacro on:

### Linux and Mac OS X

If you are running gomacro on Linux or Mac OS X, `import` will then just work. Example:
```
$ go get gonum.org/v1/plot
$ gomacro
[greeting message...]

gomacro> import "gonum.org/v1/plot"
// debug: created file "/home/max/src/gomacro_imports/gonum.org/v1/plot/plot.go"...
// debug: compiling "/home/max/go/src/gomacro_imports/gonum.org/v1/plot/plot.go" ...
gomacro> plot.New()
&{...} // *plot.Plot
<nil>  // error
```

Note: internally, gomacro will compile and load a Go plugin containing the package's exported declarations.
Go plugins require Go 1.8+ on Linux and Go 1.10.2+ on Mac OS X.

**WARNING** On Mac OS X, **never** execute `strip gomacro`: it breaks plugin support,
            and loading third party packages stops working.


### Other systems

On all other systems as Windows, Android and *BSD you can still use `import`, but there are some more steps.
Example:
```
$ go get gonum.org/v1/plot
$ gomacro
[greeting message...]

gomacro> import "gonum.org/v1/plot"
// warning: created file "/home/max/go/src/github.com/cosmos72/gomacro/imports/thirdparty/gonum_org_v1_plot.go", recompile gomacro to use it
```

Now quit gomacro, recompile and reinstall it:
```
gomacro> :quit
$ go install github.com/cosmos72/gomacro
```

Finally restart it. Your import is now linked **inside** gomacro and will work:
```
$ gomacro
[greeting message...]

gomacro> import "gonum.org/v1/plot"
gomacro> plot.New()
&{...} // *plot.Plot
<nil>  // error
```

Note: if you need several packages, you can first `import` all of them,
then quit and recompile gomacro only once.

## Generics

gomacro contains two alternative, experimental versions of Go generics:

* the first version is modeled after C++ templates, and is appropriately named "C++ style"\
  See [doc/generics-c++.md](doc/generics-c++.md) for how to enable and use them.

* the second version is named "contracts are interfaces" - or more briefly "CTI".
  It is modeled after several published proposals for Go generics,
  most notably Ian Lance Taylor's [Type Parameters in Go](https://github.com/golang/proposal/blob/master/design/15292/2013-12-type-params.md)
  It has some additions inspired from [Haskell generics](https://wiki.haskell.org/Generics)
  and original contributions from the author - in particular to create a simpler alternative to
  [Go 2 contracts](https://go.googlesource.com/proposal/+/master/design/go2draft-contracts.md)

  For their design document and reasoning behind some of the design choices, see [doc/generics-cti.md](doc/generics-cti.md)

The second version of generics "CTI" is enabled by default in gomacro.

They are in beta status, and at the moment only generic types and functions are supported.
Syntax and examples:
```go
// declare a generic type with two type arguments T and U
type Pair#[T,U] struct {
	First T
	Second U
}

// instantiate the generic type using explicit types for T and U,
// and create a variable of such type.
var pair Pair#[complex64, struct{}]

// equivalent:
pair := Pair#[complex64, struct{}] {}

// a more complex example, showing higher-order functions
func Transform#[T,U](slice []T, trans func(T) U) []U {
	ret := make([]U, len(slice))
	for i := range slice {
		ret[i] = trans(slice[i])
	}
	return ret
}
Transform#[string,int] // returns func([]string, func(string) int) []int

// returns []int{3, 2, 1} i.e. the len() of each string in input slice:
Transform#[string,int]([]string{"abc","xy","z"}, func(s string) int { return len(s) })
```
Contracts specify the available methods of a generic type.
For simplicity, they do not introduce a new syntax or new language concepts:
contracts are just (generic) interfaces.
With a tiny addition, actually: the ability to optionally indicate the receiver type.

For example, the contract specifying that values of type `T` can be compared with each other
to determine if the first is less, equal or greater than the second is:
```Go
type Comparable#[T] interface {
	// returns -1 if a is less than b
	// returns  0 if a is equal to b
	// returns  1 if a is greater than b
	func (a T) Cmp(b T) int
}
```
A type `T` implements `Comparable#[T]` if it has a method `func (T) Cmp(T) int`.
This interface is carefully chosen to match the existing methods of
`*math/big.Float`, `*math/big.Int` and `*math/big.Rat`.
In other words, `*math/big.Float`, `*math/big.Int` and `*math/big.Rat` already implement it.

What about basic types as `int8`, `int16`, `int32`, `uint`... `float*`, `complex*` ... ?
Gomacro extends them, adding many methods equivalent to the ones declared on `*math/big.Int`
to perform arithmetic and comparison, including `Cmp`:
```Go
func (a int) Cmp(b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}
```
If you do not specify the contract(s) satisfied by a type, generic functions
cannot access the fields and methods of a such type, which is then treated
as a "black box", similarly to `interface{}`
```Go
// declare a generic function with a single type argument T
func Sum#[T] (args ...T) T {
	var sum T // exploit zero value of T. this will be replaced by: sum := T().New()
	for _, elem := range args {
        // use operator += on T. this is currently accepted
		// as a temporary workaround until contracts are fully implemented.
		// the correct code would be: sum = sum.Add(sum, elem)
		sum += elem
	}
	return sum
}
Sum#[int]         // returns func(...int) int
Sum#[int] (1,2,3) // returns int(6)

Sum#[complex64]                 // returns func(...complex64) complex64
Sum#[complex64] (1.1+2.2i, 3.3) // returns complex64(4.4+2.2i)

Sum#[string]                         // returns func(...string) string
Sum#[string]("abc.","def.","xy","z") // returns "abc.def.xyz"

```
Partial and full specialization of generics is **not** supported in CTI generics,
both for simplicity and to avoid accidentally providing Turing completeness at compile-time.

Instantiation of generic types and functions is on-demand.

Current limitations:
* type inference on generic arguments #[...] is not yet implemented,
  thus generic arguments #[...] must be explicit.
* generic methods are not yet implemented.
* Contracts can be declared, but are not used.

## Debugger

Since version 2.6, gomacro also has an integrated debugger.
There are three ways to enter it:
* hit CTRL+C while interpreted code is running.
* type `:debug STATEMENT-OR-FUNCTION-CALL` at the prompt.
* add a statement (an expression is not enough) `"break"` or `_ = "break"` to your code, then execute it normally.

In all cases, execution will be suspended and you will get a `debug>` prompt, which accepts the following commands:\
`step`, `next`, `finish`, `continue`, `env [NAME]`, `inspect EXPR`, `list`, `print EXPR-OR-STATEMENT`

Also,
* commands can be abbreviated.
* `print` fully supports expressions or statements with side effects, including function calls and modifying local variables.
* `env` without arguments prints all global and local variables.
* an empty command (i.e. just pressing enter) repeats the last command.

Only interpreted statements can be debugged: expressions and compiled code will be executed, but you cannot step into them.

The debugger is quite new, and may have some minor glitches.

## Why it was created

First of all, to experiment with Go :)

Second, to simplify Go code generation tools (keep reading for the gory details)

---

Problem: "go generate" and many other Go tools automatically create
Go source code from some kind of description - usually an interface
specifications as WSDL, XSD, JSON...

Such specification may be written in Go, for example when creating JSON
marshallers/unmarshallers from Go structs, or in some other language,
for example when creating Go structs from JSON sample data.

In both cases, a variety of external programs are needed to
generate Go source code: such programs need to be installed
separately from the code being generated and compiled.

Also, Go is currently lacking generics (read: C++-like templates)
because of the rationale "we do not yet know how to do them right,
and once you do them wrong everybody is stuck with them"

The purpose of Lisp-like macros is to execute arbitrary code
while compiling, **in particular** to generate source code.

This makes them very well suited (although arguably a bit low level)
for both purposes: code generation and C++-like templates, which
are a special case of code generation - for a demonstration of how
to implement C++-like templates on top of Lisp-like macros,
see for example the project https://github.com/cosmos72/cl-parametric-types
from the same author.

Building a Go interpreter that supports Lisp-like macros,
allows to embed all these code-generation activities
into regular Go source code, without the need for external programs
(except for the interpreter itself).

As a free bonus, we get support for Eval()

## LEGAL

Gomacro is distributed under the terms of [Mozilla Public License 2.0](LICENSE)
or any later version.
