Quasiquote
==========

implementing quasiquote, unquote and unquote_splice in Go
--------------------------------------------------------

One of the main motivations behind the creation of Go interpreter `gomacro`
was to add Lisp-like macros to Go.

This includes implementing Common Lisp `quote`, `quasiquote` and, more crucially,
`unquote` and `unquote_splice` i.e. Common Lisp macro characters `'` `` ` `` `,` and `,@`

Since Go language is not homoiconic, i.e. (source) code and (program) data
are not represented identically, this is a challenge.

### Parser ###

The first (moderate) difficulty is adding support for `'` `` ` `` `,` and `,@` to Go parser.
It was solved by forking Go standard packages https://golang.org/pkg/go/scanner/
and https://golang.org/pkg/go/parser/ and patching them.

Characters `'` `` ` `` and `,` are already reserved in Go,
so the author decided to replace them as follows:
* quote          `'`  must be written `~'`
* quasiquote `` ` ``  must be written `~"` (not ``~` `` because the latter messes up syntax hilighting in Go-aware editors and IDEs - starts a multiline raw string)
* unquote        `,`  must be written `~,`
* unquote_splice `,@` must be written `~,@`

the prefix `~` is configurabile when manually instantiating the modified parser.

Go parser produces as output an abstract syntax tree (AST) represented as a tree of `ast.Node`,
from the standard package https://golang.org/pkg/go/ast/

Defining new node types is deliberately impossible (`ast.Node` is an interface with unexported methods),
luckily the existing types are flexible enough to accommodate the new syntax.

The chosen representation is somewhat cumbersome but fully general: newly created constants `token.QUOTE`, `token.QUASIQUOTE`,
`token.UNQUOTE` and `token.UNQUOTE_SPLICE` are used as unary operators on a fictitious closure containing the quoted code.
Examples:
* `'x` must be written `~'x` and is parsed as if written `~' func() { x }`
* `` `{x = y}`` must be written `~"{x = y}` and is parsed as if written `~" func() { x = y }`
* `,{1 + 2}`  must be written `~,{1 + 2}` and is parsed as if written `~, func() { 1 + 2 }`
* `,@{foo()}` must be written `~,@{foo()}` and is parsed as if written `~,@ func() { foo() }`

The fictitious closures are necessary because `ast.UnaryExpr` only allows an expression as its operand - not arbitrary
statements or declarations.
In Go, the only expression that can contain arbitrary statements and declarations is a closure (in Go terms, a "function literal")

### Classic interpreter ###

`gomacro` contains two interpreters: "classic" and "fast".

The classic interpreter is compact (about 5k LOC) and directly executes the AST, producing `reflect.Value` objects as output.
It is also quite slow (1000-3000 times slower than compiled Go), due to the overhead of continuously dispatching on the type
and contents of `ast.Node` and working with `reflect.Value` instead of native Go types.

One significant advantage of directly executing the AST is the simplicity of quasiquote implementation:
it visits depth-first the whole AST, looking for `ast.UnaryExpr` whose operator is `token.QUOTE`, `token.QUASIQUOTE`,
`token.UNQUOTE` or `token.UNQUOTE_SPLICE`, and performs the corresponding operation (either return the quoted code literally or evaluate it)
while keeping track of the current quasiquotation depth (the number of entered `~"` minus the number of entered `~,` and `~,@`)

### Fast interpreter ###

#### Architecture ####

The second, "fast" interpreter included in `gomacro` is more sophisticated. Instead of directly executing the AST,
it splits the execution in two phases:
1. visits the AST depth-first and "compiles" i.e. transforms it into tree of closures - one for each expression to be executed.
   For example, `a + b` causes the interpreter to execute something like:
   ```
   var a = compile("a").(func(env *Env) int)
   var b = compile("b").(func(env *Env) int)
   var sum_ab = func(env *Env) int {
	   return a(env) + b(env)
   }
   ```
   which creates a closure that, when later executed, computes `a + b`.
   The fast interpreter also performs type checking and type inference while "compiling" this tree of closures.

   Statements (including declarations) are "compiled", i.e. transformed, a bit differently: each one becomes
   a closure executing the statement in the interpreter, and returning the next closure to be executed.
   For example, `if x { foo() } else { bar() }` causes the interpreter to execute something like:
   ```
   var x = compile("x").(func(env *Env) bool)
   var foo = compile("foo").(func(env *Env) func())
   var bar = compile("bar").(func(env *Env) func())
   var ip_then, ip_else, ip_finish int // will be set below
   Code.Append(func(env *Env) (Stmt, *Env) {
	  var ip int
	  if x(env) {
		 ip = ip_then // goto ip_then
	  } else {
		 ip = ip_else // goto ip_else
	  }
	  env.Code.IP = ip
	  return env.Code[ip], env
   })
   ip_then = Code.Len()
   Code.Append(func(env *Env) (Stmt, *Env) {
	   foo(env)()
	   env.Code.IP = ip_finish // goto ip_finish i.e. skip else branch
	   return env.Code[ip_finish], env
   })
   ip_else = Code.Len()
   Code.Append(func(env *Env) (Stmt, *Env) {
	   bar(env)()
	   env.Code.IP = ip_finish // can also be written env.Code.IP++
	   return env.Code[ip_finish], env
   })
   ip_finish = Code.Len()
   ```
   which creates a list of closures that, when later executed, computes `if x { foo() } else { bar() }`.

   Note the extensive use of closures, i.e. anonymous functions that access **mutable** variables
   of the surrounding scope: in this case, `x` `foo` `bar` `ip_then` `ip_else` and `ip_finish`.

2) executes the "compiled" code, i.e. calls the created closures

"fast" interpreter also uses native Go types where possible, to further speed up execution
and reduce the reliance on `reflect.Value` and the overhead it imposes.

The result is a much larger interpreter:
* 20k LOC written manually
* plus further 80k LOC, generated from 8k LOC of macros, by using the "classic" interpreter and its quote,
  quasiquote and macros as a code generation tool

It is also significantly faster than the "classic" interpreter:
on most microbenchmarks, "fast" interpreter is 10-100 times slower than compiled code, instead of 1000-3000 times slower.

Interestingly, the "fast" interpreter appears to be faster than [python](https://www.python.org/) at least
on the fibonacci and collatz microbenchmarks - see [examples](../examples/)

#### Quasiquotation difficulties ####

The main difficulty in implementing quasiquotation in the "fast" interpreter is the "compile" phase:
code containing quasiquote must be type checked, and code fragments that must be evaluated should be transformed
into closures returning the result of evaluation. This is a problem similar to what Common Lisp compilers face
when compiling quasiquotations, with the difference that Go is not homoiconic.

In practice, the lack of homoiconicity means that standard textbook quasiquotation algorithms for Common Lisp
are not directly applicable to Go. Some examples will clarify this last statement:

In Common Lisp the textbook quasiquotation algorithms, as for example http://www.lispworks.com/documentation/HyperSpec/Body/02_df.htm
recursively visit the input AST, producing an output AST that does **not** contain `` ` `` `,` or `,@`
at the price of typically producing output significantly different from the input. Examples:

* `` `(+ x ,y)`` is typically expanded to the equivalent source code: `(list '+ 'x y)` - to verify it, try ``(macroexpand '`(+ x ,y))`` in a Common Lisp REPL
* `` `(x ,y)`` is typically expanded to the equivalent `(list 'x y)`
* `` `(x ,@y)`` is typically expanded to the equivalent `(list* 'x y)`
* `` `(x ,@y ,@z)`` is typically expanded to the equivalent `(list* 'x (append y z))`
* and so on...

#### First approach: quasiquotation is source transformation ####

Any attempt to translate (almost) one-to-one the same algorithm in Go, and thus also the resulting examples,
would create an algorithm able to perform the following expansions:

* `~"{x + ,y}` would be expanded to the equivalent source code: `&ast.BinaryExpr{Op: token.ADD, X: &ast.BasicLit{Name: "x"}, Y: y}`
  i.e. an ast.Node representing addition between a literal "x" and an ast.Expr equal to the result of evaluating y
* `~"{x; ,y}` would be expanded to the source code: `[]ast.Stmt{&ast.ExprStmt{X: &ast.BasicLit{Name: "x"}}, y}`
  i.e. a list of two ast.Stmt: a literal "x" wrapped in a statement, and an ast.Stmt equal to the result of evaluating y.
* `~"{x; ,@y}` would be expanded to the source code: `append([]ast.Stmt{&ast.ExprStmt{X: &ast.BasicLit{Name: "x"}}}, y...)`
  where y must be an expression that, once type-checked and transformed into a closure, will return a []ast.Stmt
* `~"{x; ,@y; ,@z}` would be expanded to the source code: `append(append([]ast.Stmt{&ast.ExprStmt{X: &ast.BasicLit{Name: "x"}}}, y...), z...)`
  where x and y must be expressions that, once type-checked and transformed into closures, will return []ast.Stmt

Note the differences between the various expansions, and the dependency on "go/ast" package of the expanded source code.

Some possible simplifications are:
1. allow `~'` in the expanded source code
2. work on `ast2.Ast` instead of `ast.Node` and the dozens of concrete types
   implementing `ast.Node`: the formers is an `ast.Node` wrapper with an uniform API that greatly simplifies
   manipulating Go abstract syntax trees.
3. implement dedicated functions and methods operating on `ast2.Ast`: `Quote`, `Unquote`, `Init`, `Concat`
   and whatever else is needed to simplify the expansion

* `~"{x + ~,y}` would be expanded to the equivalent source code: `in.New().Init(~'x, y)`
  where `in` is an `ast2.Ast` containing `~"{x + ~,y}`
* `~"{x; ~,y}` would be expanded to the source code: `in.New().Init(~'x, y)`
  where `in` is an `ast2.Ast` containing `~"{x; ~,y}`
* `~"{x; ~,@y}` would be expanded to the source code: `in.New().Init(~'x).Concat(y)`
  where `in` is an `ast2.Ast` containing `~"{x; ~,@y}`
* `~"{x; ~,@y; ~,@z}` would be expanded to the source code: `in.New().Init(~'x).Concat(y).Concat(z)`
  where `in` is an `ast2.Ast` containing `~"{x; ~,@y; ~,@z}`

Even with such simplifications, this first approach looks tricky to implement correctly in all cases,
and also fragile: expanded source code depends on external libraries, which could be shadowed or unavailable.

#### Second approach: quasiquotation merged with compile ####

Since quasiquotation must be executed on the output of macroexpansion (quasiquote could be even considered a macro),
it is the last phase before "compile" i.e. before transformation of source code to a tree of closures.

Thus an alternative approach is to merge quasiquotation with the compile phase:
while transforming AST nodes to closures, the "fast" interpreter could detect quasiquotes
and expand them - possibly not to source code, but directly to a tree of closures.

In other words, quasiquotation could directly produce executable code, without going through
the intermediate phase of expanding it to source code.

Is it easier to implement and/or more robust? Let's see.

* `~"{x + ~,y}` would be transformed into a closure, by executing something like (`node` is an `ast.Node`
  containing `~"{x + ~,y}`):
  ```
  var x = quasiquote(node.X).(func(*Env) reflect.Value) // compile to a closure that returns a copy of the &ast.BasicLit wrapped in reflect.Value 
  var y = compile("y").(func (*Env) reflect.Value) // compile to a closure that returns an ast.Node wrapped in reflect.Value
  var in = ToAst(node) // wrap into ast2.Ast
  var form = in.New()  // empty ast2.Ast with same type, operator and source position as 'in'

  var closure = func(env *Env) ast.Node {
	  var out = form.New() // create a new, empty ast2.Ast at each invokation
	  var xform = ToAst(x(env).Interface())
	  var yform = ToAst(y(env).Interface())

	  out.Set(0, xform)
	  out.Set(1, yform)
	  return ToNode(out)
  }
  ```
* `~"{x; ~,y}` would be transformed into a closure, by executing something like (`node` is an `ast.Node`
  containing `~"{x; ~,y}`):
  ```
  var x = quasiquote(node.X).(func(*Env) reflect.Value) // compile to a closure that returns a copy of the &ast.BasicLit wrapped in reflect.Value 
  var y = compile("y").(func (env *Env) reflect.Value) // compile to a closure that returns ast.Node wrapped in reflect.Value
  var in = ToAst(node) // wrap into ast2.Ast
  var form = in.New()  // empty ast2.Ast with same type, operator and source position as 'in'

  var closure = func(env *Env) ast.Node {
	  var out = form.New() // create a new, empty ast2.Ast at each invokation
	  var xform = ToAst(x(env).Interface())
	  var yform = ToAst(y(env).Interface())

	  out.Append(xform)
	  out.Append(yform)
	  return ToNode(out)
  }
  ```
* `~"{x; ~,@y}` would be transformed into a closure, by executing something like (`node` is an `ast.Node`
  containing `~"{x; ~,@y}`):
  ```
  var x = quasiquote(node.X).(func(*Env) reflect.Value) // compile to a closure that returns a copy of the &ast.BasicLit wrapped in reflect.Value 
  var y = compile("y").(func (env *Env) reflect.Value) // compile to a closure that returns ast.Node wrapped in reflect.Value
  var in = ToAst(node) // wrap into ast2.Ast
  var form = in.New()  // empty ast2.Ast with same type, operator and source position as 'in'

  var closure = func(env *Env) ast.Node {
	  var out = form.New() // create a new, empty ast2.Ast at each invokation
	  var xform = ToAst(x(env).Interface())
	  var yform = ToAst(y(env).Interface())

	  out.Append(xform)
	  for i := 0; i < y.Len(); i++ {
	      out.Append(y.Get(i))
      }
	  return ToNode(out)
  }
  ```

While the above looks somewhat complicated, it changes very little from one case to the other,
and it is actually the **implementation** of quasiquote, not its output!

Such implementation depends on `ast2.Ast` and related functions, but it will be part of the interpreter
itself - which already has such dependency - while macroexpanded source code would remain free of such dependencies.

It seems this second approach only has advantages... the only evident disadvantage is the lack
of user-available mechanisms to expand quasiquotations, i.e. an eventual "Macroexpand" function
available at the REPL and to interpreted code, would **not** expand quasiquotes.
