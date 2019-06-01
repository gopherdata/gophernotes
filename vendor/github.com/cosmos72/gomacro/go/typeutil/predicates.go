// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements commonly used type predicates.

package typeutil

import (
	"go/ast"

	"github.com/cosmos72/gomacro/go/types"
)

// Identical reports whether x and y are identical.
func Identical(x, y types.Type) bool {
	return identical(x, y, true, nil)
}

// IdenticalIgnoreTags reports whether x and y are identical if tags are ignored.
func IdenticalIgnoreTags(x, y types.Type) bool {
	return identical(x, y, false, nil)
}

func sameName(xname string, xpkg *types.Package, yname string, ypkg *types.Package) bool {
	// spec:
	// "Two identifiers are different if they are spelled differently,
	// or if they appear in different packages and are not exported.
	// Otherwise, they are the same."
	if xname != yname {
		return false
	}
	// xname == yname
	if ast.IsExported(xname) {
		return true
	}
	// not exported, so packages must be the same (pkg == nil for
	// fields in Universe scope; this can only happen for types
	// introduced via Eval)
	if xpkg == nil || ypkg == nil {
		return xpkg == ypkg
	}
	// xpkg != nil && ypkg != nil
	return xpkg.Path() == ypkg.Path()
}

func sameVarName(x, y *types.Var) bool {
	if x == nil || y == nil {
		return x == y
	}
	return x == y || sameName(x.Name(), x.Pkg(), y.Name(), y.Pkg())
}

func sameFuncName(x, y *types.Func) bool {
	if x == nil || y == nil {
		return x == y
	}
	return x == y || sameName(x.Name(), x.Pkg(), y.Name(), y.Pkg())
}

// An ifacePair is a node in a stack of interface type pairs compared for identity.
type ifacePair struct {
	x, y *types.Interface
	prev *ifacePair
}

func identicalVar(v, w *types.Var, cmpTags bool, p *ifacePair) bool {
	if v == nil || w == nil {
		return v == w
	}
	return v == w || identical(v.Type(), w.Type(), cmpTags, p)
}

func (p *ifacePair) identical(q *ifacePair) bool {
	return p.x == q.x && p.y == q.y || p.x == q.y && p.y == q.x
}

func identical(x, y types.Type, cmpTags bool, p *ifacePair) bool {
	if x == y {
		return true
	}

	switch x := x.(type) {
	case *types.Basic:
		// types.Basic types are singletons except for the rune and byte
		// aliases, thus we cannot solely rely on the x == y check
		// above.
		if y, ok := y.(*types.Basic); ok {
			return x.Kind() == y.Kind()
		}

	case *types.Array:
		// Two array types are identical if they have identical element types
		// and the same array length.
		if y, ok := y.(*types.Array); ok {
			return x.Len() == y.Len() && identical(x.Elem(), y.Elem(), cmpTags, p)
		}

	case *types.Slice:
		// Two slice types are identical if they have identical element types.
		if y, ok := y.(*types.Slice); ok {
			return identical(x.Elem(), y.Elem(), cmpTags, p)
		}

	case *types.Struct:
		// Two struct types are identical if they have the same sequence of fields,
		// and if corresponding fields have the same names, and identical types,
		// and identical tags. Two anonymous fields are considered to have the same
		// name. Lower-case field names from different packages are always different.
		if y, ok := y.(*types.Struct); ok {
			if x.NumFields() == y.NumFields() {
				for i, n := 0, x.NumFields(); i < n; i++ {
					f := x.Field(i)
					g := y.Field(i)
					if f.Anonymous() != g.Anonymous() ||
						cmpTags && x.Tag(i) != y.Tag(i) ||
						!sameVarName(f, g) ||
						!identical(f.Type(), g.Type(), cmpTags, p) {
						return false
					}
				}
				return true
			}
		}

	case *types.Pointer:
		// Two pointer types are identical if they have identical base types.
		if y, ok := y.(*types.Pointer); ok {
			return identical(x.Elem(), y.Elem(), cmpTags, p)
		}

	case *types.Tuple:
		// Two tuples types are identical if they have the same number of elements
		// and corresponding elements have identical types.
		if y, ok := y.(*types.Tuple); ok {
			if x.Len() == y.Len() {
				for i, n := 0, x.Len(); i < n; i++ {
					v := x.At(i)
					w := y.At(i)
					if !identical(v.Type(), w.Type(), cmpTags, p) {
						return false
					}
				}
				return true
			}
		}

	case *types.Signature:
		// Two function types are identical if they have the same number of parameters
		// and result values, corresponding parameter and result types are identical,
		// and either both functions are variadic or neither is. Parameter and result
		// names are not required to match.
		//
		// PATCH: also compare the receiver type
		if y, ok := y.(*types.Signature); ok {
			return x.Variadic() == y.Variadic() &&
				identicalVar(x.Recv(), y.Recv(), cmpTags, p) &&
				identical(x.Params(), y.Params(), cmpTags, p) &&
				identical(x.Results(), y.Results(), cmpTags, p)
		}

	case *types.Interface:
		// PATCH: two interface types are identical if they have the same explicit methods
		// and the same embedded interfaces. The order of methods and embeddeds is *relevant*.
		if y, ok := y.(*types.Interface); ok {
			na := x.NumMethods()
			nb := y.NumMethods()
			ne := x.NumEmbeddeds()
			nf := x.NumEmbeddeds()
			if na == nb && ne == nf {
				// this PATCHED definition of type identity sidesteps the type cycles
				// created via method parameter types that are anonymous interfaces
				// (directly or indirectly) embedding the current interface, as for example
				//
				//    type T interface {
				//        m() interface{T}
				//    }
				//
				// simply by *not* embedding the interfaces, and checking for identical
				// embedded interfaces - which are always named and thus compare trivially
				// without recursion, preventing any infinite cycle or recursion.
				q := &ifacePair{x, y, p}
				for p != nil {
					if p.identical(q) {
						return true // same pair was compared before
					}
					p = p.prev
				}
				for i := 0; i < na; i++ {
					a := x.Method(i)
					b := y.Method(i)
					if !sameFuncName(a, b) || !identical(a.Type(), b.Type(), cmpTags, q) {
						return false
					}
				}
				for i := 0; i < ne; i++ {
					e := x.Embedded(i)
					f := y.Embedded(i)
					if e.Obj() != f.Obj() {
						return false
					}
				}
				return true
			}
		}

	case *types.Map:
		// Two map types are identical if they have identical key and value types.
		if y, ok := y.(*types.Map); ok {
			return identical(x.Key(), y.Key(), cmpTags, p) && identical(x.Elem(), y.Elem(), cmpTags, p)
		}

	case *types.Chan:
		// Two channel types are identical if they have identical value types
		// and the same direction.
		if y, ok := y.(*types.Chan); ok {
			return x.Dir() == y.Dir() && identical(x.Elem(), y.Elem(), cmpTags, p)
		}

	case *types.Named:
		// Two named types are identical if their type names originate
		// in the same type declaration.
		if y, ok := y.(*types.Named); ok {
			return x.Obj() == y.Obj()
		}

	case nil:

	default:
		unreachable()
	}

	return false
}

func unreachable() {
	panic("unreachable")
}
