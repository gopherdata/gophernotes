Generics
========

C++-style generics in Go
------------------------

gomacro contains two alternative, experimental implementation of Go generics.

* this document describes the first version of Go generics:
  it is modeled after C++ templates, and is appropriately named "C++ style"

* for the second version of Go generics, which is named "Contracts are Interfaces" (CTI) and is enabled by default,
  see the [Generics](../README.md#Generics) section in the main [README.md](../README.md)
  and, for more details, the design document [generics-cti.md](generics-cti.md).

Getting started
---------------

C++-style generics are **not** enabled by default in gomacro.

To enable them, edit the file [go/etoken/generics.go](../go/etoken/generics.go)
and set the constants
```Go
// enable C++-style generics?
const GENERICS_V1_CXX = true

// enable generics "contracts are interfaces" ?
const GENERICS_V2_CTI = false
```
then save the file and recompile gomacro.

Due to historical reasons, plus the fact that this version of generics are modeled after C++ templates,
Go generics are named 'templates' in this document.

They are in beta status, and at the moment only generic types and functions are supported.
Syntax and examples:
```go
template[T,U] type Pair struct { First T; Second U }

var pair Pair#[complex64, struct{}]

// equivalent:
pair := Pair#[complex64, struct{}] {}


template[T] func Sum(args ...T) T {
	var sum T // exploit zero value of T
	for _, elem := range args {
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

template[T,U] func Transform(slice []T, trans func(T) U) []U {
	ret := make([]U, len(slice))
	for i := range slice {
		ret[i] = trans(slice[i])
	}
	return ret
}
Transform#[string,int] // returns func([]string, func(string) int) []int

// returns []int{3, 2, 1} i.e. the len() of each string in input slice:

Transform#[string,int]([]string{"abc","xy","z"}, func(s string) int { return len(s) })

// Partial and full specialization of templates are supported.
// Together with recursive templates, they also (incidentally)
// provide Turing completeness at compile-time:

// The following example uses recursion and full specialization
// to compute fibonacci sequence at compile time.

// general case: encode Fib#[N] in the length of array type.
template[N] type Fib [
	len((*Fib#[N-1])(nil)) +
	len((*Fib#[N-2])(nil))   ]int

template[] for[2] type Fib [1]int // specialization for Fib#[2]
template[] for[1] type Fib [1]int // specialization for Fib#[1]

const Fib30 = len((*Fib#[30])(nil)) // compile-time constant

```
Current limitations:
* instantiation is on-demand, but template arguments #[...] must be explicit.
* template methods not supported yet.

Observation: the compile-time Turing completeness provided by these C++-style templates
is really poorly readable, for three reasons:
* iteration must be written as recursion
* `if` must be written as template specialization, outside the main template
* integers must be encoded inside types, for example in the length of array types

In the author's opinion, compile-time Turing completeness is a very enticing
feature for several use cases and for a non-trivial percentage of developers.

If the only way to get such feature is with poorly readable (ab)use of templates,
the result is a lot of poorly readable template code.

If Turing-complete templates are ever added to Go (or any other language)
it is thus very important to also provide an alternative, more natural syntax
to perform Turing-complete computation at compile-time. An example
could be: `const foo(args)` where the function `foo` must respect certain
constraints (to be defined) in order to be callable at compile time.

## History and details ##

The next sections contain observations, choices, difficulties and solutions
found while implementing C++ style generics in gomacro.

### Parser ###

#### Declaring templates ####

Adding a new syntax to declare template types, function and methods is easy:
it's just a matter of inventing the syntax and choosing a representation in
terms of `go/ast.Node`s

Current syntax is:
```
template [T1, T2...] type ...
template [T1, T2...] func ...
```

Template type declarations are represented with `*ast.TypeSpec` as usual,
with the difference that `TypeSpec.Type` now contains
`&ast.CompositeLit{Type: <type declaration>, Elts: [T1, T2 ...]}`

Template function and method declarations are represented with `*ast.FuncDecl`
as usual, with the difference that `FuncDecl.Recv.List` now has two elements:
the first element is nil for functions and non-nil for methods,
the second element is `&ast.Field{Names: nil, Type: &ast.CompositeLit{Elts: [T1, T2 ...]}}`

#### Using templates ####

The main idea is that template functions and methods will be used mostly
in the same ways non-template ones, i.e. `Func(args)` and `obj.Method(args)`
exploiting appropriate type inference (exact inference rules need to be defined).

In some cases, using template functions and methods will need to specify
the exact template arguments. Template types will need such explicit
qualification most of (or maybe all) the time.

For example, after a declaration
```
template [T1, T2] type Pair struct { First T1; Second T2 }
```
it is tempting to say that the syntax to specify the template arguments
(= to qualify the template name) is
```
Pair[int, string]
```
i.e. the template name is immediately followed by '[' and the comma-separated
list of template arguments.

Alas, such syntax is too ambiguous for current Go parser. Take for example the
code fragment
```
func Nop(Pair[int, int]) { }
```
By manual inspection, it's clear that `Pair` is a type name, not a parameter
name. But compare the fragment above with this:
```
func Nop(Pair []int) { }
```
where `Pair` is a parameter name with type `[]int`.

In both cases, the parser will encounter `Pair` followed by `[` and must
decide how to parse them without further look-ahead.

The current parser algorithm for this case assumes that `Pair` is an
identifier and that `[` starts a type expression to be parsed.

To avoid breaking lots of existing code, the current parser algorithm for
this case must be preserved. So we need a different, less ambiguous syntax to
qualify template names.

One of the suggestions in latest Ian Lance Taylor
[Type parameters (December 2013)](https://github.com/golang/proposal/blob/master/design/15292/2013-12-type-params.md)
proposal is "using double square brackets, as in `Vector[[int]]`, or perhaps
some other character(s)."

The authors' current decision - but it's trivial to change it - is to write
`Pair#[int, string]` and similarly `Vector#[int]`. The reason is twofold:

1. double square brackets look too "magical"
2. the hash character `#` is currently not used in Go syntax, and usually does not have
   strong connotations in programmers' minds. The alternatives are the other
   ASCII characters currently not used in Go syntax: `?` `@` `$` `~`
   * the question mark `?` is better suited for conditionals, as for example
     the C ternary operator `?:`
   * the at sign `@` already has several common meanings (at, email...).
   * the dollar sign `$` seems inappropriate, in the author's opinion, for
     this usage.
   * the tilde sign `~` is already used by gomacro for quasiquote and friends.

Implementation choice: `Pair#[int, string]` is represented by
```
&ast.IndexExpr{X: Pair, Index: &ast.CompositeLit{Elts: [T1, T2...]} }
```
The simpler `&ast.CompositeLit{Type: Pair, Elts: [T1, T2...]} }` would suffice
for the parser, but compiling it is much more ambiguous, since it could be
interpreted as the composite literal `Pair{T1, T2}`

#### Composite Literals ####

The parser had to be extended to recognize things like `Pair#[T1,T2] {}`
as a valid composite literal.

In practice, `isTypeName()` and `isLiteralType()` now return true for `*ast.IndexExpr`.

This solution should be better examined to understand whether the increased
syntax ambiguity is a problem, but an official implementation will surely create
new ast.Node types to hold template declarations and template uses, bypassing
this potential problem.

### Declaration ###

The declaration of template types and functions is straightforward.

For each template declaration found, the compiler now collects it in the same
`map[string]*Bind` used for const, var and func declarations.

Such declarations store in the *Bind their **source code** as an ast.Node, in order to
retrieve and compile it on-demand when the template type or function needs to be
instantiated.

This is easy for an interpreter, but more tricky for a compiler:
since a package A may use a template B.C declared in package B,
the compiler may need to instantiate B.C while compiling A.

There are at least two solutions:
1. for each compiled package, store in the compiled archive packagename.a
   the **source code** of each template, alongside with binary code.

   This may not play well with commercial, binary-only libraries since,
   with a little effort, the source code of templates could be extracted.

2. for each compiled package, require its source code to be available
   in order to instantiate its templates.

   This has the same problem as above, only in stronger form:
   the **full** source code of B must be available when compiling A.

Another question is: where to store the B.C instantiated while compiling A ?

For templates declared in the standard library and instantiated by non-root users,
$GOROOT may not be writeable, so it should probably be stored in
$GOPATH/pkg/$GOOS_$GOARCH/path/to/package, using a name like B.somehash.a

### Instantiation ###

Instantiantion is a regular compile, with some careful setup.

Since a template may access global symbols in the scope where it was declared,
it must be compiled in that **same** scope. Better yet, it can be compiled
in a new inner scope, that defines the template arguments to use for instantiation.

An example can help to understand the abstract sentence above:
suppose package B contains
```
package B

const N = 10

template[T] type Array [N]T
```

and is later used by package A as
```
package A

import "B"

var arr B.Array#[int]
```

the technique described abstractly above means: to compile `B.Array#[int]`,
pretend that package B contains the following (pseudo-code, it's not valid Go):
```
{ // open new scope

	type T = int // inject the template argument

	// inject the template declaration literally - no replacements needed
	type Array#[T] [N]T // finds T immediately above, and N in outer scope
}
```

There is a peculiarity in this approach that must be handled carefully:
`type Array#[T]` should not be taken too literally. It conveys the
intention, but the exact mechanics are more subtle:

1. the name `Array` is a template type. It must have an associated cache
   that keeps track of already-instantiated types based on it, otherwise
   each `Array#[Foo]` will trigger an instantiation (= a compile) even
   if it exists already.
2. such cache has the same role as the list of existing (non-template)
   types, functions, constants and variables: looks up identifiers and
   resolves them.
3. Go supports (non-template) recursive functions and types,
   and we want to also support recursive template functions and types,
   as for example `template[T] List { First T; Rest *List#[T] }`
   See the next paragraph for details.

### Recursive templates ###

Let's start with a non-template example for concreteness:
```
type IntList struct { First int; Rest *IntList }
```
Compiling it in Go is conceptually three-step process:
1. forward-declare `IntList`, i.e. create a new named type `IntList`
   with no underlying type (i.e. it's incomplete) and add it to the current
   scope.
2. compile the underlying type `struct { First int; Rest *IntList }`.
   It will find the **incomplete** type `IntList` in the current scope,
   but that's ok because it uses a **pointer** to `IntList`,
   not an `IntList` - Go, as C/C++/Java and many other languages,
   allow and can implement pointers to incomplete types because at the
   assembler level they are all implemented in the same way: a machine word
   (`void *`, `unsafe.Pointer`, etc.) with pointer semantics.
   For completeness: also slices, maps, channels and functions signatures
   of incomplete types are accepted in Go.
3. complete the forward-declared `IntList` by setting its underlying type to
   the result of step 2.

Recursive template types and functions can be implemented very similarly:
instantiating
```
template[T] List struct { First T; Rest *List#[T] }
```
as for example `List#[string]`, is almost the same process: it starts
with the technique described in the paragraph [Instantiation](#instantiation)
above:
```
{ // open new scope

	type T = string // inject the template argument

	// inject the template declaration literally - no replacements needed
	// except for conceptually replacing List -> List#[T] in the declaration
	// (not in the body)
	type List#[T] struct { First T; Rest *List#[T] }
}
```
and it continues with the analogous of the three-step process described above:
1. forward-declare `List#[string]` i.e. add to the cache of instantiated types
   a new named type `List#[string]` with no underlying type (i.e. it's
   incomplete)
2. compile the underlying type `struct { First T; Rest *List#[T] }` in the scope
   just prepared above for the instantiation.
   It will find the **incomplete** type `List#[string]` in the cache of
   instantiated types, but that's ok because its uses a **pointer** to
   `List#[string]`, not a `List#[string]`. As we said, pointers to incomplete
   types are accepted.
3. complete the forward-declared `List#[string]` by setting its underlying type
   to the result of step 2.

### Partial and full specialization ###

This is a desirable feature of C++ templates.
Although not overly difficult to implement, it introduces a lot of complexity:
C++ templates are Turing-complete because of it.

In extreme summary it means that, in addition to the general declaration of a template,
one can also declare special cases.

Example 1: given the template function declaration
```
template[T] func nonzero(a, b T) T { if a != 0 { return a }; return b }
```
one can declare the special case "T is a map of something" as:
```
template[K,V] for[map[K]V] func nonzero(a, b map[K]V) map[K]V { if a != nil { return a }; return b }
```
and the special case "T is struct{}" as:
```
template[] for[struct{}] func nonzero(a, b struct{}) struct{} { return struct{}{} }
```
Note that the number of template arguments **can** be different in each specialized declaration.

A specialized declaration with zero template arguments is named "full specialization"
or "fully specialized"; all other specialized declarations are named "partial
specialization" or "partially specialized".

The compiler is expected to automatically decide which specialization to use,
based on the criteria "use the most specialized declaration that is applicable".

In case there is no single "most specialized declaration", the compiler
is expected to produce a (hopefully descriptive) error.

Implementation note: choosing the "most specialized declaration" requires the
following steps:
1. keep a list of candidates, initially containing only the general declaration.
2. for each specialization, pattern-match it against the code to compile
   (for example `nonzero#[map[int]string]`).
   If it does not match, ignore it and repeat step 2. with the next specialization.
   It it matches, name it "new candidate" and continue to step 3.
3. compute the types and constants required to match the new candidate against the
   code to compile. For example, the candidate `template[K,V] for[map[K]V] func nonzero(...) ...`
   matches the code `nonzero#[map[int]string]` if `K = int` and `V = string`
4. perform a loop, comparing the new candidate selected at step 2. against each candidate
   currently in the list. If the new candidate is more is more specialized than a current one,
   the latter is removed from the candidate list.
5. add the new candidate to the candidate list, storing also the types and constants
   computed at step 3.
6. if there are more specializations, return to step 2. with the next specialization.

The comparison at step 4. "candidate A is more specialized than candidate B"
can be implemented as: B pattern-matches A, but A does not pattern-match B.

Pattern-matching compares the ast.Node tree structure and the contents
of each *ast.Ident and *ast.BasicList, but it should also expand type aliases
and compute the value of constant expressions before comparing them.

It is not yet clear whether it is feasible for pattern-matching to also expand
template types in case they are type aliases too.

### Turing completeness ###

If one has some familiarity with C++ templates, it is easy to see that
the partial and full specialization rules described above are Turing complete
at compile-time, just like C++ templates.

The reason is:
* partial and full specializations are a compile-time `if`
* instantiating a template from inside another one is a compile-time `while`
* compile-time computation on integers can be implemented with normal arithmetics
* intermediate results can be stored in the number of elements of an array type,
  and extracted with `len()`

For example, this is a compile-time computation of fibonacci numbers
using the rules proposed above:

```
template[N] type Fib [len((*Fib#[N-1])(nil)) + len((*Fib#[N-2])(nil))] int
template[] for[1] type Fib [1]int
template[] for[0] type Fib [0]int
const Fib10 = len((*Fib#[10])(nil))
```
arguably, the Go code above is even **less** readable than the already convoluted
C++ equivalent:
```
template<int N> struct Fib { enum { value = Fib<N-1>::value + Fib<N-1>::value }; };
template<> struct Fib<1> { enum { value = 1 }; };
template<> struct Fib<0> { enum { value = 0 }; };
enum { Fib10 = Fib<10>::value };
```

This seems to present a conundrum:
1. allow partial template specialization and, as a consequence, compile-time
   Turing-completeness, with the **extremely** unreadable syntax required to use it
2. or forbid partial template specialization, preserving readability as much
   as possible, but severely limiting the usefulness of templates?

If Go adds compile-time Turing-completeness, whatever its syntax,
it is such an enticing feature that many programmers will certainly use it.
Some programmers may **heavily** use it, and the result could be something
resembling the well-known C++ libraries STL and Boost:

professional code, that heavily uses templates, very useful and very used,
but written in a dialect very different from the basic language (C++ in this case),
almost unreadable for average programmers skilled mostly on non-template code,
and difficult to read even for experts.

In my opinion, there is only one solution to the conundrum:
<b>add another, readable syntax to perform compile-time computation.</b>

As minimum, such syntax would be used in most cases for compile-time
Turing-completeness **instead** of the extremely unreadable template
specializations, simply because it has the same features
(compile-time Turing-completeness) but is more readable.

Ideally, such syntax could also be used to simplify writing complex
template code.

To give some context, Go is not foreign to compile-time computation:
`//go:generate` allows to execute arbitrary commands at compile-time,
and Go code generation tools and techniques are accepted and
quite in widespread use (at least compared to many other languages).

### Compile-time function evaluation ###

Following the suggestion of the previous chapter, a very simple syntax
to perform compile-time computation could be `const EXPRESSION`,
as for example:
```
func fib(n int) int { if n <= 1 { return n }; return fib(n-1)+fib(n-2) }
const fib30 = const fib(30)
```
This is readable, and the programmer's intention is clear too:
invoke `fib(30)` and treat the result as a constant - which implies
`fib(30)` must be invoked at compile time.

Question: which functions can be invoked at compile time?\
Answer: a minimal set could be: all functions in current package,
provided they do not use imported packages, print() or println(),
or invoke other functions that (transitively) use them.

Question: global variables should be accessible by functions
invoked at compile time?\
Answer: tentatively no, because if such variables are modified at
compile-time, their value at program startup becomes difficult to
define unambiguously, and difficult to store in the compiled code.

So, a tentative definition of whether a function can be invoked
at compile time is:
1. is defined in the current package (so source code
   is available in order to check points 1. and 2. below)
2. does not use global variables, imported packages, print()
   or println()
3. calls only functions that (transitively) respect 1. and 2.
4. as a consequence, calls to closures are allowed, provided
   that the function creating the closure respects 1, 2 and 3.

An alternative, wider definition could be: only pure functions
can be invoked at compile time. A function is pure if:
1. does not use global variables, print() or println(), or assembler
2. either does not call other functions, or only calls pure functions
As a special case, all builtin functions except `print()` and `println()`
are considered pure.
This alternative definition allows calling function in other
packages at compile-time, provided they are pure.
Thus it requires storing in compiled packages a flag for each function,
indicating whether it is pure or not.
