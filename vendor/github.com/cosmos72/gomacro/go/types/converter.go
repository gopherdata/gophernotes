// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file converts objects from go/types to github.com/cosmos72/go/types

package types

import (
	"fmt"
	"go/types"
	"reflect"
)

type Converter struct {
	pkg map[string]*Package
	// use pointer identity to compare types, not types.Identical.
	// faster although less accurate.
	cache        map[types.Type]Type
	toaddmethods map[*Named]*types.Named
	tocomplete   []*Interface
}

type funcOption bool

const (
	funcIgnoreRecv funcOption = false
	funcSetRecv    funcOption = true
)

// should be called with argument Universe
// to initialize basic types, constants true/false/iota and 'error'
func (c *Converter) Init(universe *Scope) {
	// create builtin package i.e. universe
	p := NewPackage("", "")
	if universe != nil {
		// fill package with contents of universe scope
		scope := p.Scope()
		for _, name := range universe.Names() {
			scope.Insert(universe.Lookup(name))
		}
	}
	c.pkg = map[string]*Package{
		"": p,
	}
}

// convert *go/types.Package -> *github.com/cosmos72/gomacro/go/types.Package
func (c *Converter) Package(g *types.Package) *Package {
	if g == nil {
		return nil
	}
	c.cache = nil
	p := c.mkpackage(g)
	scope := g.Scope()
	for _, name := range scope.Names() {
		obj := c.Object(scope.Lookup(name))
		if obj != nil {
			p.scope.Insert(obj)
		}
	}
	return p
}

// convert go/types.Object -> github.com/cosmos72/gomacro/go/types.Object
func (c *Converter) Object(g types.Object) Object {
	switch g := g.(type) {
	case *types.Const:
		return c.Const(g)
	case *types.Func:
		return c.Func(g)
	case *types.TypeName:
		return c.TypeName(g)
	case *types.Var:
		return c.Var(g)
	default:
		return nil
	}
}

// convert *go/types.Const -> *github.com/cosmos72/gomacro/go/types.Const
func (c *Converter) Const(g *types.Const) *Const {
	return NewConst(g.Pos(), c.mkpackage(g.Pkg()), g.Name(), c.Type(g.Type()), g.Val())
}

// convert *go/types.Func -> *github.com/cosmos72/gomacro/go/types.Func
func (c *Converter) Func(g *types.Func) *Func {
	return NewFunc(g.Pos(), c.mkpackage(g.Pkg()), g.Name(), c.Type(g.Type()).(*Signature))
}

// convert *go/types.TypeName -> *github.com/cosmos72/gomacro/go/types.TypeName
func (c *Converter) TypeName(g *types.TypeName) *TypeName {
	ret, _ := c.mktypename(g)
	if ret.typ == nil {
		ret.typ = c.typ(g.Type())
	}
	return ret
}

// convert *go/types.Var -> *github.com/cosmos72/gomacro/go/types.Var
func (c *Converter) Var(g *types.Var) *Var {
	return NewVar(g.Pos(), c.mkpackage(g.Pkg()), g.Name(), c.Type(g.Type()))
}

// convert go/types.Type -> github.com/cosmos72/gomacro/go/types.Type
func (c *Converter) Type(g types.Type) Type {
	ret := c.typ(g)
	for _, t := range c.tocomplete {
		t.Complete()
	}
	c.tocomplete = c.tocomplete[0:0:cap(c.tocomplete)]

	for t, g := range c.toaddmethods {
		c.addmethods(t, g)
		delete(c.toaddmethods, t)
	}
	return ret
}

func (c *Converter) typ(g types.Type) Type {
	t := c.cache[g]
	if t != nil {
		return t
	}
	switch g := g.(type) {
	case *types.Array:
		elem := c.typ(g.Elem())
		t = NewArray(elem, g.Len())
	case *types.Basic:
		return Typ[BasicKind(g.Kind())]
	case *types.Chan:
		elem := c.typ(g.Elem())
		t = NewChan(ChanDir(g.Dir()), elem)
	case *types.Interface:
		t = c.mkinterface(g)
	case *types.Map:
		t = c.mkmap(g)
	case *types.Named:
		t = c.mknamed(g)
	case *types.Pointer:
		elem := c.typ(g.Elem())
		t = NewPointer(elem)
	case *types.Signature:
		t = c.mksignature(g, funcSetRecv)
	case *types.Slice:
		elem := c.typ(g.Elem())
		t = NewSlice(elem)
	case *types.Struct:
		t = c.mkstruct(g)
	default:
		panic(fmt.Errorf("Converter.Type(): unsupported types.Type: %T", g))
	}
	if c.cache == nil {
		c.cache = make(map[types.Type]Type)
	}
	c.cache[g] = t
	return t
}

var getEmbeddedType func(*types.Interface, int) types.Type

func init() {
	t := reflect.TypeOf((*types.Interface)(nil))
	m, ok := t.MethodByName("EmbeddedType")
	if ok {
		getEmbeddedType = m.Func.Interface().(func(*types.Interface, int) types.Type)
	} else {
		// types.Interface.EmbeddedType() does not exist in go 1.9
		getEmbeddedType = func(g *types.Interface, i int) types.Type {
			return g.Embedded(i)
		}
	}
}

func (c *Converter) mkinterface(g *types.Interface) *Interface {
	n := g.NumExplicitMethods()
	fs := make([]*Func, n)
	for i := 0; i < n; i++ {
		fs[i] = c.mkfunc(g.ExplicitMethod(i), funcIgnoreRecv)
	}
	n = g.NumEmbeddeds()
	es := make([]Type, n)
	for i := 0; i < n; i++ {
		es[i] = c.typ(getEmbeddedType(g, i))
	}
	t := NewInterfaceType(fs, es)
	c.tocomplete = append(c.tocomplete, t)
	return t
}

func (c *Converter) mkmap(g *types.Map) *Map {
	key := c.typ(g.Key())
	elem := c.typ(g.Elem())
	return NewMap(key, elem)
}

func (c *Converter) mknamed(g *types.Named) *Named {
	typename, found := c.mktypename(g.Obj())
	if found && typename.Type() != nil {
		return typename.Type().(*Named)
	}
	t := NewNamed(typename, nil, nil)
	u := c.typ(g.Underlying())
	t.SetUnderlying(u)
	if g.NumMethods() != 0 {
		if c.toaddmethods == nil {
			c.toaddmethods = make(map[*Named]*types.Named)
		}
		c.toaddmethods[t] = g
	}
	return t
}

func (c *Converter) mksignature(g *types.Signature, opt funcOption) *Signature {
	var recv *Var
	if opt == funcSetRecv {
		recv = c.mkparam(g.Recv())
	}
	return NewSignature(
		recv,
		c.mkparams(g.Params()),
		c.mkparams(g.Results()),
		g.Variadic(),
	)
}

func (c *Converter) mkstruct(g *types.Struct) *Struct {
	n := g.NumFields()
	fields := make([]*Var, n)
	tags := make([]string, n)
	for i := 0; i < n; i++ {
		fields[i] = c.mkfield(g.Field(i))
		tags[i] = g.Tag(i)
	}
	return NewStruct(fields, tags)
}

func (c *Converter) mkpackage(g *types.Package) *Package {
	if g == nil {
		return nil
	}
	path := g.Path()
	if p := c.pkg[path]; p != nil {
		return p
	}
	p := NewPackage(path, g.Name())
	c.pkg[path] = p
	return p
}

func (c *Converter) universe() *Package {
	return c.pkg[""]
}

func (c *Converter) mktypename(g *types.TypeName) (*TypeName, bool) {
	pkg := c.mkpackage(g.Pkg())
	if pkg == nil {
		pkg = c.universe()
	}
	scope := pkg.Scope()
	obj := scope.Lookup(g.Name())
	// to preserve type identity, reuse existing typename if found
	if typename, ok := obj.(*TypeName); ok {
		return typename, true
	}
	typename := NewTypeName(g.Pos(), pkg, g.Name(), nil)
	pkg.Scope().Insert(typename)
	return typename, false
}

func (c *Converter) mkfield(g *types.Var) *Var {
	// g.Embedded() is a newer alias for g.Anonymous(),
	// but go 1.9 does not have it
	return NewField(g.Pos(), c.mkpackage(g.Pkg()), g.Name(), c.typ(g.Type()), g.Anonymous())
}

func (c *Converter) mkparam(g *types.Var) *Var {
	if g == nil {
		return nil
	}
	return NewParam(g.Pos(), c.mkpackage(g.Pkg()), g.Name(), c.typ(g.Type()))
}

func (c *Converter) mkparams(g *types.Tuple) *Tuple {
	if g == nil {
		return nil
	}
	n := g.Len()
	v := make([]*Var, n)
	for i := 0; i < n; i++ {
		v[i] = c.mkparam(g.At(i))
	}
	return NewTuple(v...)
}

func (c *Converter) mkvar(g *types.Var) *Var {
	if g == nil {
		return nil
	}
	return NewVar(g.Pos(), c.mkpackage(g.Pkg()), g.Name(), c.typ(g.Type()))
}

func (c *Converter) mkfunc(m *types.Func, opt funcOption) *Func {
	sig := c.mksignature(m.Type().(*types.Signature), opt)
	return NewFunc(m.Pos(), c.mkpackage(m.Pkg()), m.Name(), sig)
}

func (c *Converter) addmethods(t *Named, g *types.Named) {
	n := g.NumMethods()
	for i := 0; i < n; i++ {
		m := c.mkfunc(g.Method(i), funcSetRecv)
		t.AddMethod(m)
	}
}
