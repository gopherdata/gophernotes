## gomacro - interactive Go interpreter with macros

gomacro is a fairly complete Go interpreter, implemented in pure Go. It offers both
an interactive REPL and a scripting mode, and does not require a Go toolchain at runtime
(except in one very specific case: import of a non-standard package).

It has very few dependencies: go/ast, go/types, reflect and,
for goroutines support, golang.org/x/sync/syncmap.

Gomacro can be used as:
* a standalone executable with interactive Go REPL:  
  just run `gomacro` from your command line or, better, `rlwrap gomacro`
  (rlwrap is a wrapper that adds history and line editing to terminal-based
  programs - available on many platforms)
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

* a library that adds Eval() and scripting capabilities
  to your Go programs - provided you comply with its LGPL license

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

## Current Status

Fairly complete.

The intepreter supports:
* multiline input
* line comments starting with #! in addition to //
* basic types: booleans, integers, floats, complex numbers, strings (and iota)
* the empty interface, i.e. interface{} - other interfaces not implemented yet
* constant, variable and type declarations (untyped constants are emulated with typed constants)
* Go 1.9 type aliases (experimental)
* unary and binary operators
* assignment, i.e. operators = += -= *= /= %= &= |= ^= &^= <<= >>=
* composite types: arrays, channels, maps, pointers, slices, strings, structs
* composite literals
* type assertions
* function declarations (including variadic functions)
* method declarations (including variadic methods and methods with pointer receiver)
* seamless invocation of compiled functions from interpreter, and vice-versa
* channel send and receive
* goroutines, i.e. go function(args)
* function and method calls, including multiple return values
* if, for, for-range, break, continue, fallthrough, return (unimplemented: goto)
* select, switch, type switch, fallthrough
* all builtins: append, cap, close, comples, defer, delete, imag, len, make, new, panic, print, println, real, recover
* imports: Go standard packages "just work", importing other packages requires the "plugin" package (available only for Go 1.8+ on Linux)
* switching to a different package
* macro declarations, for example `macro foo(a, b, c interface{}) interface{} { return b }`
* macro calls, for example `foo; x; y; z`
* macroexpansion: code walker, MacroExpand and MacroExpand1
* ~quote and ~quasiquote. they take any number of arguments in curly braces, for example:
  `~quote { x; y; z }`
* ~unquote and ~unquote_splice
* ~func, ~lambda: specializations of "func".
  * ~lambda always starts a closure (lambda) or a function type
  * ~func always starts a function or method declaration
  useful to resolve a limitation in Go syntax that becomes significant for ~quote and ~quasiquote:
  * in declarations, "func" always declares a function or method - there is no way to declare a closure (lambda) or function type
  * in statements and expressions, including the body of ~quote and ~quasiquote,
    "func" always declares a closure (lambda) or a function type - there is no way to declare a function or method
* nesting macros, quotes and unquotes

Some features are still missing:
* interfaces. They can be declared, but nothing more: there is no way to implement them or call their methods
* extracting methods from types. For example `time.Duration.String` should return a `func(time.Duration) string`
  but currently gives an error.
  Instead extracting methods from objects is supported: `time.Duration(1s).String` correctly returns a func() string
* goto
* named return values
* history/readline (rlwrap does the job in most cases)

Limitations:
* no distinction between named and unnamed types created by interpreted code.
  For the interpreter, `struct { A, B int }` and `type Pair struct { A, B int }`
  are exactly the same type. This has subtle consequences, including the risk
  that two different packages define the same type and overwrite each other's methods.

  The reason for such limitation is simple: the interpreter uses `reflect.StructOf()`
  to define new types, which can only create unnamed types.
  The interpreter then defines named types as aliases for the underlying unnamed types.

* cannot create recursive types, as for example `type List struct { First interface{}; Rest *List}`
  The reason is the same as above: the interpreter uses `reflect.StructOf()` to define new types,
  which cannot create recursive types


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
