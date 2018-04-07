Gomacro default interpreter supports:
* history/readline (uses https://github.com/peterh/liner)
* multiline input
* line comments starting with #! in addition to //
* basic types: booleans, integers, floats, complex numbers, strings (and iota)
* interfaces imported from compiled code, as for example `interface{}`, `error`, `fmt.Stringer`...
* new interfaces can be declared inside interpreted code
* constant, variable and type declaration, including untyped constants
* Go 1.9 type aliases (experimental)
* unary and binary operators
* assignment, i.e. operators = += -= *= /= %= &= |= ^= &^= <<= >>=
* composite types: arrays, channels, maps, pointers, slices, strings, structs
* composite literals
* type assertions
* seamless invocation of compiled functions from interpreter, and vice-versa
* channel send and receive
* goroutines, i.e. go function(args)
* function and method calls, including multiple return values
* function and method declarations (including variadic functions/methods,
  and methods with pointer receiver)
* named return values
* extracting methods from types and from instances.
  For example `time.Duration.String` returns a `func(time.Duration) string`
  and `time.Duration(1s).String` returns a `func() string`
* if, for, for-range, break, continue, fallthrough, return (goto is only partially implemented)
* select, switch, type switch, fallthrough
* all builtins: append, cap, close, comples, defer, delete, imag, len, make, new, panic, print, println, real, recover
* imports: Go standard packages "just work". Importing other packages requires either the "plugin" package
  (available only for Go 1.8+ on Linux) or, in alternative, adding an `init()` function to the package,
  then compiling it statically inside gomacro (see [github issue #13](https://github.com/cosmos72/gomacro/issues/13) for more details)
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

Some features are still missing or incomplete:
* interpreted interfaces are supported, but not extensively tested yet.
* out-of-order code. Types, variables and functions must be declared **before** using them.
* switching to a different package
  (if you absolutely need it, the older and slower `gomacro.classic.Interp` supports switching to a different package)
* goto is partially implemented, needs to be completed
* type inference in composite literals - see [github issue #9](https://github.com/cosmos72/gomacro/issues/9)

* incomplete interface -> interface type assertions and type switches:
  they do not support yet interpreted types stored in interfaces, and interpreted interfaces.

* unimplemented conversion typed constant -> interpreted interface (see fast/literal.go:207)
  Workaround: assign the constant to a variable, then convert the variable to the interpreted interface

* bug: on systems without Go toolchain installed, a few imports including "net/http" fail with xreflect panic "internal error"

* bug: import "gonum.org/v1/plot/plotter" and import "gonum.org/v1/plot/plotutil" fail with stack overflow due to infinite recursion

* bug: if gomacro is linked as a shared library (see https://stackoverflow.com/questions/1757090/shared-library-in-go)
  some method calls on constants do not work. example:
    import "os"
    os.ModeAppend.IsDir()
  interface conversion: interface {} is func() bool, not func() bool


Other limitations:
* named types created by interpreted code are emulated.
  When the interpreter is asked to create for example `type Pair struct { A, B int }`,
  it actually creates the unnamed type `struct { A, B int }`.
  Everything works as it should within the interpreter, but extracting the struct
  and using it in compiled code reveals the difference.

  Reason: gomacro relies on the Go reflect package to create new types,
  but there is no function `reflect.InterfaceOf()` or any other way to create new **named** types,
  so gomacro uses `reflect.StructOf` which can only create unnamed types.

* recursive types are emulated too.
  For example `type List struct { First interface{}; Rest *List}`
  is actually a `struct { First interface{}; Rest *interface{} }`.
  Again, everything works as it should within the interpreter, but extracting
  the struct and using it in compiled code reveals the difference.

  The reason is: the interpreter uses `reflect.StructOf()` to define new types,
  which cannot create recursive types

  Interestingly, this means the interpreter also accepts the following declaration,
  which is rejected by Go compiler: `type List2 struct { First int; Rest List2 }`
  Note that `Rest` is a `List2` **not** a pointer to `List2`

* interpreted interfaces are emulated too.
  New interface types created by interpreted code are actually anonymous structs.
  Also here, everything works as it should within the interpreter, but extracting
  the interface and using it in compiled code reveals the difference.

  Reason: gomacro relies on the Go reflect package to create new types,
  and there is no function `reflect.InterfaceOf()`, so the interpreter uses
  `reflect.StructOf()` and a lot of bookkeeping to emulate new interface types.

* operators << and >> on untyped constants do not follow the exact type deduction rules.
  The implemented behavior is:
  * an untyped constant shifted by a non-constant expression always returns an int
  * an untyped floating point constant shifted by a constant expression returns an untyped integer constant.
    the interpreter signals an error during the precompile phase
    if the left operand has a non-zero fractional or imaginary part,
    or it overflows both int64 and uint64.
  See [Go Language Specification](https://golang.org/ref/spec#Operators) for the correct behavior

* recover() does not support mixing interpreted and compiled code:

  recover() works normally if the function and its defer are either
  **both interpreted** or **both compiled**.

  but if an interpreted function invokes as defer a compiled function,
  or a compiled function invokes as defer an interpreted function,
  then, inside that defer, recover() will not work:
  it will return nil and will **not** stop panics.
