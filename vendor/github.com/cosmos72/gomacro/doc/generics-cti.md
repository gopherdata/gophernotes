Generics in Go
==============

a wishlist
----------

This file contains experiences and desiderata for a Go implementation of generics,
i.e. polymorphic types and functions.

It is a feature present in many other languages with varying names. A few examples:
* C++ [templates](https://www.geeksforgeeks.org/templates-cpp)
* Java [generics](https://en.wikipedia.org/wiki/Generics_in_Java)
* Haskell [generics](https://wiki.haskell.org/Generics)

The author has experience using generics in the three languages listed above,
which will also be used for comparison and reference in the rest of this document.

In addition, the author personally added generics to three programming languages:
* Go: the unofficial interpreter [gomacro](https://github.com/cosmos72/gomacro)
  contains a Go implementation of generics, modeled after C++ templates.
* Common Lisp: the library [cl-parametric-types](https://github.com/cosmos72/cl-parametric-types)
  contains a Common Lisp implementation of generics, again modeled after C++ templates.
* the [lfyre](https://sourceforge.net/projects/lfyre) programming language,
  created by the author, used to contain an implementation of generics.
  It now has a different maintainer.

# Anti-goals

Things the author does not want from Go generics

* a compile-time sub-language:

  Go generics should be an extension of idiomatic Go, not a whole sub-language
  to be used for compile time operations.

  For example, we should avoid compile-time Turing completeness and "expression templates",
  two accidental features of C++ templates that together created a sub-language of C++
  made of template specializations and recursive templates.\
  Such sub-language also provides arbitrary computation at compile-time (possibly a good thing)
  with a terrible syntax and no alternative with cleaner syntax.\
  The much more recent C++ `constexpr` actually provides the desired alternative, clean syntax
  for compile-time Turing completeness, but it is more limited: it can only manipulate values,
  not types.

# Goals

The reasons to implement generics in Go can be many, and sometimes contradicting.
The author's personal list of reasons, which can also be interpreted as goals
that Go generics are expected to achieve, are:

* reusable, flexible algorithms and types. Examples:

  a single `sort#[T]` function that can sort any slice of any ordered type.\
  a single `cache#[K,V]` type that can cache key/value pairs of any type,
  provided that keys can be compared.\
  a single `sortedmap#[K,V]` type, similar to the existing `map[K]V`
  but keeps its entries sorted, like C++ `map<K,V>` or Java `TreeMap<K,V>`

* type-safety:

  generic functions and types should be instantiable on arbitrary,
  concrete types - for example `sort#[int]` would only accept `[]int` slices
  and `cache#[uint64, []byte]` would only accept `uint64` keys and `[]byte` values.
  In particular, generic functions and types should not need to use `interface{}`,
  either internally or in they exported API, and should not need type assertions at runtime.

* high runtime speed, low runtime overhead:

  generic functions and types should be reified in order to maximize code execution speed
  and have low or zero data representation overhead in memory.

  Reified means that `sort#[int]` and `sort#[uint]` will be two different and unrelated functions,
  one only dealing with `int` slices and the other only dealing with `uint` slices,
  and that `cache#[uint64, []byte]` and `cache#[uint32, []byte]` will be two different
  and unrelated types, with (possibly) different layout in memory.

  While reified generics also have disadvantages (see for example
  https://gbracha.blogspot.com/2018/10/reified-generics-search-for-cure.html)
  the author has extensive experience with both reified generics (C++, Haskell)
  and non reified generics (Java), and he is convinced that reified generics
  are a better fit for Go - the reasons can be explained if needed.

  One obvious disadvantage of reified generics is that each instantiation
  of a generic function must be compiled separately, for example `sort#[int]`
  and `sort#[uint]`, increasing build time.

  Luckily, Go `import`s compiled packages instead of `#include`-ing their source code,
  which is expected to contain build time for two reasons:

  1. each generic function will be parsed only once. Instead C++ `#include` mechanism
  typically needs to parse again the same generic function each time it is included
  by a different source file.

  2. each instantiation of a generic function - say `sort#[int]` - will be compiled
  only once, provided that Go implements a cache of instantiated functions and types,
  similarly to how it implements a cache of compiled packages.\
  Instead C++ `#include` mechanism typically needs to compile again
  the same generic function - say `sort<int>` even if it's instantiated with the same types
  from two different source files - for example `a.cpp` and `b.cpp` both use `sort<int>`.
  C++ compilers typically delegates to the linker the job of coalescing multiple,
  identical versions of the same generic function.

* reasonable build time:

  it is expected to be achieved / achievable even with reified generics, see the previous item

* type inference:

  Go extensively uses (and encourages to use) type inference instead
  of explicitly declaring the type of a variable.\
  Example: `a := foo()` rather than `var a int = foo()`.

  When an expression returns multiple values,
  Go actively pushes the programmer to use type inference. Example:
  ```
  n, err := fmt.Println("foo")
  ```
  becomes more verbose without type inference, because each `var`
  declaration can only reference one type:
  ```
  var n int
  var err error
  n, err = fmt.Println("foo")
  ```

  The goal for generics is to preserve and extend support for type inference,
  for example by allowing the syntax
  ```
  slice := make([]int, n)
  sort(slice)
  ```
  and automatically inferring that it means
  ```
  slice := make([]int, n)
  sort#[int](slice)
  ```

* contracts:

  when writing a generic function or type, it should be possible to specify contracts
  on their type arguments. This is an extensively discussed topic, for many reasons:

  1) contracts are expected to simplify compiler error messages, and make them
     more understandable. For example, a `sort#[T]` function would specify that values
	 of `T` must be ordered - the following syntax is just for illustration purposes:
	 ```Go
	 func sort#[T: Ordered](slice []T) {
		// ...
	 }
	 ```
	 Then, attempting to sort a non-ordered type as for example `func ()` could produce
	 an error message like `sort: type func() is not Ordered` instead
	 of some deeply nested error due to `a < b` used on `func()` values.

  2) contracts allow programmers writing generic code to specify explicitly
     the requirements of their code, i.e. on which types it can be used and why.

	 Without them, it is not always simple to understand if a complicated generic function
	 or type written by someone else can be used with a certain concrete type `T`,
	 and what are the requirements on such `T`:\
	 the author of generic code could document the requirements, for example in a comment,
	 but he/she may forget it, or the comment could become stale/erroneous if the generic
	 code gets updated.

	 A machine-readable, compiled information is less likely to become stale/erroneous,
	 especially if the compiler actually validates it.

  3) if the compiler assumes that contracts specify the **only** operations
     supported by the constrained types, it could detect immediately if a constrained
	 type is used improperly in generic code, without having to wait until it gets
	 instantiated (possibly by someone else) on concrete types - for example if methods
	 or arithmetic operations are used on a type that is only constrained as `T: Ordered`

	 For reference, Haskell does exactly that: a contract specifies the only operations
	 allowed on a type.\
	 Actually, Haskell does even more: if a contract for a type `T` is not specified,
	 the compiler infers it from the operations actually performed on `T` values
	 (it's not obvious whether such contract inference is appropriate for Go).

  It should also be possible to specify multiple contracts on a type.
  For example, if a type `T` must be both `Ordered` and `Printable`,
  one could imagine a syntax like:
  ```Go
  func foo#[T: Ordered, Printable](arg T) {
	// ...
  }
  ```

* contracts implementation:

  An important question is: what should a contract tell about a type?

  1) The signature of one or more methods?

  2) The signature of one or more functions and/or operators?

  3) The name and type of one or more fields?

  4) A combination of the above?

  It is surely tempting to answer 1. and reuse interfaces as contracts:
  this would spare us from inventing yet another language construct, but is it enough?

## Option 1. contracts declare type's methods

  Let's check with a relatively simple case: the `Ordered` contract.\
  It describes types that can be ordered, and there's immediately a difficulty:
  Go operator `<` only works on basic types (integers and floats), and cannot be overloaded
  i.e. cannot be extended to support further types.
  It will work on types whose underlying type is integer or float, as for example
  ```Go
    package time

    type Duration int64
  ```
  but even in such case you cannot define a custom implementation:\
  operator `<` compares `time.Duration` values as it would compare `int64`.

  So let's say that `Ordered` will instead use a **function** `Less()` to compare values.\
  Here we hit another Go (intentional) limitation: function overloading is not supported,\
  i.e. it's not possible to define multiple functions with the same name and different signatures.

  Ok, then let's say that `Ordered` will use a **method** `Less()` to compare values.\
  How do we express that a type must have a method `Less()` to be `Ordered`?\
  With an interface, of course:
  ```Go
    type Ordered interface {
	  Less(/*what here?*/) bool
    }
  ```
  We are getting close: we need to express that the single argument of `Less()`
  is the same as the receiver. Go does not support this either, but we are trying to
  extend it with generics, and the addition "we can give a name to the receiver type"
  feels quite minimal.\
  What about the following?
  ```Go
    type Ordered#[T] interface {
	  func (T) Less(T) bool
    }
  ```
  It's supposed to mean that `Ordered` is a generic interface, i.e. it's polymorphic,
  and has a single type argument `T`. To satisfy `Ordered`, a type must have a method
  `Less(T)` where `T` is also the receiver type (the `func (T)` part).
  I chose the syntax `func (T) Less ...` because that's exactly how we already declare
  methods, and the shorter `(T) Less ...` did not sound familiar enough.

  There are still a couple of issues.

  First issue: basic integers and floats do not have any method, so they cannot implement `Ordered`.
  This can only be solved with a Go language specs change which adds methods to basic types.
  On the other hand user-defined types, including standard library ones as `time.Duration`,
  could add a method `Less()`.

  Second issue: methods must be declared in the same package as their receiver.
  In other words, it's not possible to import a type `foo.Bar` and add a method `Less()` to it:
  either the method is already there because the author forecasted the need, or it's not there
  and there's no way to add it (unless you fork the package `foo` and modify it -
  something that should be a last resort, not the normal case).
  This cannot be solved reasonably - but it can become an intentional limitation.

## Option 2. contracts declare functions on a type

  Let's continue our thought experiment on the `Ordered` contract.\
  This time, contracts declare functions on a type, not its methods.

  Again, Go operator `<` cannot be overloaded, so we use a **function** `Less()`:
  ```Go
    type Ordered#[T] contract {
	  func Less(T, T) bool
    }
  ```
  which means that `Ordered` is a generic contract (is it still an interface?
  we can try to answer later) and has a single type argument `T`.\
  A concrete type `T` satisfies `Ordered` if there is a function `Less(T,T) bool`.\
  Since functions cannot be overloaded either, it's immediately evident that
  we can only declare one function `Less` per package.\
  That's not what we wanted, and it pushes us toward a much deeper language change:
  allow function overloading, i.e. multiple functions with the same name but different signatures.

  And once we allow function overloading, why not going the full way and allowing operator overloading too?

  The result would be something like:
  ```Go
    type Ordered#[T] contract {
	  operator<(T, T) bool
    }
  ```
  and an hypotetical type `Foo` could satisfy `Ordered` by declaring a function
  ```Go
    operator<(a, b Foo) bool {
	  // ...
    }
  ```

  A lot of design decisions would have to follow:\
  In which cases do we allow function overloading and/or operator overloading?\
  How do we select the function/operator to call when there are multiple candidates
  with the same name, differing only in their signature?

  And also more mundane questions, as whether we write `operator<(a, b Foo) bool { }`
  or `func operator<(a, b Foo) bool { }`.

  Although the author really likes Haskell generics, and they happen to go down this exact road,
  it still feels like a big language change and a hard sell to Go core team and Go community.

## Option 3. contracts declare type's fields

  This would be likely frowned upon in many object-oriented languages as C++ or Java,
  where direct access to object's fields is strongly discouraged in favor of setter/getter methods.

  Yet Go composite literals are an extremely useful feature, and they rely on initializing
  exported struct fields to work. Thus maybe it could make sense. Let's see if it's also useful.

  One could say that a type `T` satisfies the contract `Ordered` if `T` has a certain field?\
  It does not seem very useful since fields contain values, they usually do not
  "do something" - that's for methods.

  Furthermore Go has the peculiar feature that methods can be declared on any named type,
  not just on structs. But requiring that a type has certain fields makes sense only
  for structs - quite limiting.

  In conclusion it seems to be usable only in some cases, and not useful enough even in those.

## Option 4. combination of the above

  The total complexity added to the language would be quite high: the sum of each complexity,
  plus all the interactions (intentional and accidental) among the proposals.

  If option 2. feels like a hard sell, this simply seems too much.

## Option summary

Among the three options analyzed above, the best one appears to be the first:
contracts declare type's methods.\
It allows to use generics in many scenarios, yet requires quite limited changes to the language:

* slightly extending `interface` syntax to optionally specify the receiver type
* adding methods to basic types - one method per supported operator (actually fewer,
  since `Less` can also cover `<=`, `>` and `>=`, while `Equal` can also cover `!=`, etc.)

In exchange it allows:

* declaring contracts with a familiar syntax - the same as method declaration
  and very similar to interface methods
* creating generic algorithms as `sort#[T]` and generic types as `sortedmap#[K,V]`
  that work out of the box on both Go basic types and on user-defined types

## Option 1 deeper analysis

In option 1, contracts are interfaces, i.e. they declare the methods of a type.
With the small extension of allowing to specify also the receiver type,
they seem very useful and let programmers create very general generic types
and generic algorithms, yet they seem to have very few unintended side effects
on the language, and they do not introduce huge language changes.

Are there other downsides we did not consider yet?

Let's analyze more in detail the idea of adding methods on basic types.

To simplify the reasoning, we start with the concrete example `sort#[T]`,
which as we said requires a method `Less` on `T`.

So let's suppose that `int`, `int8`, `int16`, `int32`, `int64`,
`uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`, `float32` and `float64`
have such method.

Then a type such as `time.Duration`, which is declared as
```
package time
type Duration int64
```
will have the method `Less` or not?

### Underlying types

In Go, there is the rule
* a named type has the methods of its underlying type i.e. "wrapper methods",
  plus the methods declared on the named type
Following this rule, the question becomes: what's the underlying type of `time.Duration`?

* If the underlying type is `int64`, then `time.Duration` will have a wrapper method `Less`
* If the underlying type is something else (what?) then `time.Duration` will probably not have a wrapper method `Less`.

Now things get subtle. Usually, underlying types are **not** named types,
they are instead unnamed types: channels, maps, slices, arrays, functions, and very often structs.

If `int64` was the underlying type of `time.Duration`, then these two types
would be assignable to each other, as for example:
```
import "time"
var i int64 = 7
var d time.Duration = i
```
Instead, the above does not compile. It turns out that you need an explicit **conversion**, i.e.
```
import "time"
var i int64 = 7
var d time.Duration = time.Duration(i)
```
which is required when two named types have the same underlying type.

Then the underlying type of both `int64` and `time.Duration` is some unnamed type
that cannot be mentioned directly, and is not expected to have the method `Less`.

Thus `time.Duration` would **not** have a method `Less` either.

Small issue: if you ask to the package `go/types`, the underlying type of
`time.Duration` is `int64`. This is inconsistent, and I think Go specs explain
the inconsistency as an exception.

If we ignore this small issue, we get the following:

* `sort#[T]` works on `int64` because it declares the method `Less`
* `sort#[T]` does **not** work on `time.Duration` because it lacks the method `Less`

This is clearly annoying and cumbersome.

An alternative is to decide that the underlying type of both `int64` and `time.Duration`
(the unnamed type that cannot be mentioned directly) has the method `Less`,
thus both `int64` and `time.Duration` also have `Less` as wrapper method.

The situation becomes:

* `sort#[T]` works on `int64` because it has the wrapper method `Less`
* `sort#[T]` does **not** work on `time.Duration` because it has the wrapper method `Less`

Now this is good, but it has a subtle side effect: what happens if `time.Duration`
declares its own method `Less` for some reason?

Such method `Less` shadows (hides) the wrapper method, and `sort#[T]` will happily
use it for sorting, provided it has the expected signature.\
Thus we have a way to declare a custom ordering criterion for a type.

In essence, `time.Duration` can define its own ordering by declaring a method
`Less` - to be precise, a method `func (time.Duration) Less(time.Duration) bool`.

This is very similar to what C++ achieves using operator overloading:
if a C++ type has the `operator<`, such operator will be used by `std::sort()`
as default comparison operator.

So we have replaced operator overloading with a different but equivalent mechanism:\
"declare a method with a certain name and signature".\
The name `Less` becomes special, because `sort#[T]` looks for it.

But the function or method name `operator<` actualy looks special (it's not alphanumeric),
while the name `Less` does not look very special (it's alphanumeric).

Thus `time.Duration` may declare a method `Less` for its own purposes,
without realizing that `sort#[T]` will try to use it.

Worse, `time.Duration` or some other similar type may **already** declare a method `Less`,
and once we introduce generics, `sort#[T]` would use the method `Less` to compare values,
instead of comparing the underlying type (`int64` and friends).

This is an unwanted effect, and quite insidious too: an existing, innocent looking
method `Less` suddenly acquires a special meaning, and causes existing code
(the various sorting algorithms in package `sort`) to silently change their behaviour.


**TO BE CONTINUED**

# Design space

There are many possible ways to implement generics - one could say **too** many -
and they can be extremely different in usability, expressiveness, implementation complexity,
compile-time performance and run-time performance.

**TO BE CONTINUED**
