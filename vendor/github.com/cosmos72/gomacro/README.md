## gomacro - interactive Go interpreter and debugger with macros

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
  
  
* a Go source code debugger: see [Debugger](#debugger)

* an interactive tool to make science more productive and more fun.
  If you use compiled Go with scientific libraries (physics, bioinformatics, statistics...)
  you can import the same libraries from gomacro REPL (immediate on Go 1.8+ and Linux,
  requires restarting on other platforms, see [Importing packages](#importing-packages) below),
  call them interactively, inspect the results, feed them to other functions/libraries,
  all in a single session.
  The imported libraries will be **compiled**, not interpreted,
  so they will be as fast as in compiled Go.

  For a graphical user interface on top of gomacro, see [Gophernotes](https://github.com/gopherdata/gophernotes).
  It is a Go kernel for Jupyter notebooks and nteract, and uses gomacro for Go code evaluation.

* a library that adds Eval() and scripting capabilities to your Go programs in few lines
  of code:
	```
	package main
	import (
		"fmt"
		"reflect"
		"github.com/cosmos72/gomacro/fast"
	)
	func RunGomacro(toeval string) reflect.Value {
		interp := fast.New()
		// for simplicity, only collect the first returned value
		val, _ := interp.Eval(toeval)
		return val
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

  Run `gomacro -m -w FILENAMES` to parse and expand macros in one or more files.
  For each filename on the command line, gomacro will parse it, expand macros,
  then create a corresponding FILENAME.go with the parsed and macroexpanded
  imports, declarations and statements.

  To parse and macroexpand all *.gomacro files in a directory, run `gomacro -m -w DIRECTORY`

## Installation

### Prerequites

- [Go 1.9+](https://golang.org/doc/install)

### Supported platforms

Gomacro is pure Go, and in theory it should work on any platform supported by the Go compiler.
The following combinations are tested and known to work:

- Linux: amd64, 386, arm64, arm, mips, ppc64le (on Linux it can also import 3<sup>rd</sup> party libraries at runtime)
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

* importing 3<sup>rd</sup> party libraries on non-Linux systems is cumbersome - see [Importing packages](#importing-packages).
* some corner cases using interpreted interfaces, as interface -> interface type assertions and type switches, are not implemented yet.
* goto can only jump backward, not forward
* out-of-order code is under testing - some corner cases, as for example out-of-order declarations
  used in keys of composite literals, are not supported.  
  Clearly, at REPL code is still executed as soon as possible, so it makes a difference mostly
  if you separate multiple declarations with ; on a single line. Example: `var a = b; var b = 42`  
  Support for "batch mode" is in progress - it reads as much source code as possible before executing it,
  and it's useful mostly to execute whole files or directories.

The [documentation](doc/) also contains the [full list of features and limitations](doc/features-and-limitations.md)

## Extensions

Compared to compiled Go, gomacro supports several extensions:

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
    ```
	import "math/big"
	var i *big.Int = 1<<1000                 // exact - would overflow int
	var r *big.Rat = 1.000000000000000000001 // exact - different from 1.0
	var s *big.Rat = 5e1232                  // exact - would overflow float64
	var t *big.Rat = 1e1234                  // approximate, exceeds 5e1232
	var f *big.Float = 1e646456992           // largest untyped float constant that is different from +Inf
    ```
  Note: every time such a conversion is evaluated, it creates a new value - no risk to modify the constant.

  Be aware that converting a huge value to string, as typing `f` at REPL would do, can be very slow.

* macros, quoting and quasiquoting (to be documented)

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
```
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

### Linux

If you are running gomacro on Linux, `import` will then just work. Example:
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
Currently, Go plugins are fully functional only on Linux.


### Other systems

On Mac OS X, Windows, Android and *BSD you can still use `import`, but there are some more steps.
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

## Debugger

Since version 2.6, gomacro also has an integrated debugger.
There are two ways to use it:
* type `:debug STATEMENT-OR-FUNCTION-CALL` at the prompt.
* add a statement (an expression is not enough) `"break"` or `_ = "break"` to your code, then execute it normally.

In both cases, execution will be suspended and you will get a `debug>` prompt, which accepts the following commands:  
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
(except for the intepreter itself).

As a free bonus, we get support for Eval()

## LEGAL

Gomacro is distributed under the terms of [Mozilla Public License 2.0](LICENSE)
or any later version.
