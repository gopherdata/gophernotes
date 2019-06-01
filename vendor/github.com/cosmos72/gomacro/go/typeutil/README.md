## typeutil - patched versions of `go/types.Identical` and `golang.org/x/tools/go/type/typeutil.Map`

typeutil contains patched version of some Go utilities to handle `go/types.Type`

1. an Identical() function with a stricter definition of type identity:

   * interfaces are considered identical only if they **print** equally,
     so embedding an interface is **different** from copying its methods
     (standard `go/types.Identical` intentionally does not distinguish
     these two cases). Also, the order of methods and embedded interfaces
	 is relevant.

   * methods are considered identical only if their receiver, parameters
     and results types are identical (standard `go/types.Identical`
     intentionally ignores the receiver type)

2. Map: a mapping from `go/types.Type` to `interface{}` values,
   using the stricter definition of type identity defined above.

   Since `go/types.Type` are not canonical, i.e. not unique, comparing
   them with == does not give the expected results, as explained in
   https://github.com/golang/example/tree/master/gotypes#types

   So a specialized map is needed to use them as keys - either
   `golang.org/x/tools/go/type/typeutil.Map`, or this patched version
   `github.com/cosmos72/gomacro/typeutil/Map`, or something analogous

They are useful as type canonicalizing tools for the Go interpreter gomacro,
and not necessarily suitable for other purpouses.

## License

BSD-3-Clause as the original, unpatched utilities.

