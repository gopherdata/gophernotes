## gomacro - interactive Go interpreter with macros

gomacro is a fairly complete Go interpreter, implemented in pure Go. It offers both
an interactive REPL and a scripting mode, and does not require a Go toolchain at runtime
(except in one very specific case: import of a 3d-party package).

It has very few dependencies: go/ast, go/types, reflect, github.com/peterh/liner and,
for legacy reasons (no longer used by the default interpreter), golang.org/sync/syncmap.

Gomacro can be used as:
* a standalone executable with interactive Go REPL:
  just run `gomacro` from your command line.
  Available options:
    ```
    -c,   --collect          collect declarations and statements, to print them later
    -e,   --expr EXPR        evaluate expression
    -f,   --force-overwrite  option -w will overwrite existing files
    -h,   --help             show this help and exit
    -i,   --repl             interactive. start a REPL after evaluating expression, files and dirs.
                             default: start a REPL only if no expressions, files or dirs are specified
    -m,   --macro-only       do not execute code, only parse and macroexpand it.
                             useful to run gomacro as a Go preprocessor
    -n,   --no-trap          do not trap panics in the interpreter
    -t,   --trap             trap panics in the interpreter (default)
    -s,   --silent           silent. do NOT show startup message, prompt, and expressions results.
                             default when executing files and dirs.
    -v,   --verbose          verbose. show startup message, prompt, and expressions results.
                             default when executing an expression.
    -vv,  --very-verbose     as -v, and in addition show the type of expressions results.
                             default when executing a REPL
    -w,   --write-decls      write collected declarations and statements to *.go files.
                             implies -c
    -x,   --exec             execute parsed code (default). disabled by -m
    ```

    Options are processed in order, except for -i that is always processed as last.

    Collected declarations and statements can be also written to standard output
    or to a file with the REPL command :write

* an interactive tool to make science more productive and more fun.
  If you use compiled Go with scientific libraries (physics, bioinformatics, statistics...)
  you can import the same libraries from gomacro REPL (requires Go 1.8+ and Linux),
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

  Note: gomacro is currently [LGPL](LICENSE), which imposes some restrictions
  on programs that use gomacro. The author is currently looking for a license
  that has fewer (but not zero) restrictions. If you are sufficiently expert with licenses,
  you can help by reading [issue #14](https://github.com/cosmos72/gomacro/issues/14)
  and indicating a license that satisfies the requirements listed in it.

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

- Linux: amd64, 386, arm64, arm (on Linux it can also import 3rd party libraries at runtime)
- Mac OS X: amd64, 386
- Windows: amd64, 386
- FreeBSD: amd64, 386

### How to install

  The command
  ```
  go get -u github.com/cosmos72/gomacro
  ```
  downloads, compiles and installs gomacro

## Current Status

Almost complete.

The main missing features are:

* importing 3rd party libraries on non-Linux systems. Gomacro uses the Go 'plugin' package for this, and it currently works only on Linux.
* some corner cases using interpreted interfaces, as interface -> interface type assertions and type switches, are not implemented yet.
* out-of-order code. Types, constants, variables and functions must be declared **before** using them.
* switching to a different package (if you absolutely need it, the older and slower `gomacro.classic.Interp` supports switching to a different package)
* goto is partially implemented - can only jump back, not forward
* type inference in composite literals - see [github issue #9](https://github.com/cosmos72/gomacro/issues/9)

The [documentation](doc/) also contains the [full list of features and limitations](doc/features-and-limitations.md)


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
