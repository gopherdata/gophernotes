/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * interface.go
 *
 *  Created on: Mar 29, 2017
 *      Author: Massimiliano Ghilardi
 */

package fast

import (
	"fmt"
	"go/ast"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// compile an interface definition
func (c *Comp) TypeInterface(node *ast.InterfaceType) xr.Type {
	if node.Methods == nil || len(node.Methods.List) == 0 {
		return c.TypeOfInterface()
	}
	types, names := c.TypeFields(node.Methods)

	// parser returns embedded interfaces as unnamed fields
	var methodnames []string
	var methodtypes, embeddedtypes []xr.Type
	for i, typ := range types {
		if i < len(names) && len(names[i]) != 0 {
			methodnames = append(methodnames, names[i])
			methodtypes = append(methodtypes, typ)
		} else {
			embeddedtypes = append(embeddedtypes, typ)
		}
	}
	return c.Universe.InterfaceOf(methodnames, methodtypes, embeddedtypes)
}

// InterfaceProxy returns the proxy struct that implements a compiled interface
func (c *Comp) InterfaceProxy(t xr.Type) r.Type {
	ret := c.interf2proxy[t.ReflectType()]
	if ret == nil {
		c.Errorf("internal error: proxy not found for %s type <%v>", t.Kind(), t)
	}
	return ret
}

// converterToProxy compiles a conversion from 'tin' into a proxy struct that implements the interface type 'tout'
// and returns a function that performs such conversion
func (c *Comp) converterToProxy(tin xr.Type, tout xr.Type) func(val r.Value) r.Value {
	rtout := tout.ReflectType()       // a compiled interface
	rtproxy := c.InterfaceProxy(tout) // one of our proxies that pre-implement the compiled interface

	vtable := r.New(rtproxy).Elem()
	n := rtout.NumMethod()
	for i := 0; i < n; i++ {
		mtdout := rtout.Method(i)
		mtdin, count := tin.MethodByName(mtdout.Name, mtdout.PkgPath)
		if count == 0 {
			c.Errorf("cannot convert type <%v> to interface <%v>: missing method %s %s", tin, rtout, mtdout.PkgPath, mtdout.Name)
		} else if count > 1 {
			c.Errorf("type <%v> has %d wrapper methods %s %s all at the same depth=%d - cannot convert to interface <%v>",
				tin, count, mtdout.PkgPath, mtdout.Name, len(mtdin.FieldIndex), tout)
		}
		e := c.compileMethodAsFunc(tin, mtdin)
		setProxyField(vtable.Field(i+1), r.ValueOf(e.Value))
	}
	extractor := c.extractor(tin)
	if extractor == nil {
		return func(val r.Value) r.Value {
			vaddr := r.New(rtproxy)
			vproxy := vaddr.Elem()
			vproxy.Set(vtable)
			vproxy.Field(0).Set(r.ValueOf(xr.MakeInterfaceHeader(val, tin)))
			return vaddr.Convert(rtout)
		}
	}
	// extract object from tin proxy or emulated interface (if any),
	// and wrap it in tout proxy
	return func(val r.Value) r.Value {
		v, t := extractor(val)
		vaddr := r.New(rtproxy)
		vproxy := vaddr.Elem()
		vproxy.Set(vtable)
		vproxy.Field(0).Set(r.ValueOf(xr.MakeInterfaceHeader(v, t)))
		return vaddr.Convert(rtout)
	}
}

func setProxyField(place r.Value, mtd r.Value) {
	rtin := mtd.Type()
	rtout := place.Type()
	if rtin == rtout {
		place.Set(mtd)
	} else if rtin.ConvertibleTo(rtout) {
		place.Set(mtd.Convert(rtout))
	} else {
		place.Set(r.MakeFunc(rtout, func(args []r.Value) []r.Value {
			args[0] = args[0].Interface().(xr.InterfaceHeader).Value()
			return mtd.Call(args)
		}))
	}
}

// extract a value from a proxy struct (one of the imports.* structs) that implements an interface
// this is the inverse of the function returned by Comp.converterToProxy() above
func (g *CompGlobals) extractFromProxy(v r.Value) (r.Value, xr.Type) {
	// base.Debugf("type assertion: value = %v <%v>", v, base.ValueType(v))
	if v == base.Nil || v == base.None || !v.IsValid() || !v.CanInterface() {
		// cannot rebuild with concrete type
		return v, nil
	}
	i := v.Interface()
	v = r.ValueOf(i) // rebuild with concrete type
	rt := r.TypeOf(i)
	var xt xr.Type
	// base.Debugf("type assertion: concrete value = %v <%v>", i, t)
	if rt != nil && rt.Kind() == r.Ptr && g.proxy2interf[rt.Elem()] != nil {
		v = v.Elem().Field(0)
		i = base.ValueInterface(v)
		if j, ok := i.(xr.InterfaceHeader); ok {
			// base.Debugf("type assertion: unwrapped value = %v <%T>", j, j)
			v = j.Value()
			xt = j.Type()
		} else {
			// base.Debugf("type assertion: failed to unwrap value = %v <%T>", i, i)
			v = r.ValueOf(i) // rebuild with concrete type
		}
	}
	return v, xt
}

// converterToProxy compiles a conversion from 'tin' into the emulated interface type 'tout'
// and returns a function that performs such conversion
func (c *Comp) converterToEmulatedInterface(tin, tout xr.Type) func(val r.Value) r.Value {
	if !tin.Implements(tout) {
		c.Errorf("cannot convert from <%v> to <%v>", tin, tout)
	}
	n := tout.NumMethod()
	obj2methodFuncs := make([]func(r.Value) r.Value, n)

	tsrc := tin
	if tin.Kind() == r.Ptr {
		// xr.Type.MethodByName wants T, not *T, even for methods with pointer receiver
		tsrc = tin.Elem()
	}
	for i := 0; i < n; i++ {
		mtdout := tout.Method(i)
		mtdin, count := tsrc.MethodByName(mtdout.Name, c.PackagePath) // pkgpath is ignored for exported names

		if count == 0 {
			c.Errorf("cannot convert from <%v> to <%v>: missing method %s %s", tin, tout, mtdout.Name, mtdout.Type)
		} else if count > 1 {
			c.Errorf("cannot convert from <%v> to <%v>: multiple methods match %s %s", tin, tout, mtdout.Name, mtdout.Type)
		}
		if !mtdin.Type.AssignableTo(mtdout.Type) {
			c.Errorf("cannot convert from <%v> to <%v>: mismatched method %s: expecting %s, found %s",
				tin, tout, mtdout.Name, mtdout.Type, mtdin.Type)
		}
		obj2methodFuncs[i] = c.compileObjGetMethod(tin, mtdin)
	}
	rtout := tout.ReflectType()

	extractor := c.extractor(tin)
	if extractor == nil {
		return func(obj r.Value) r.Value {
			return xr.ToEmulatedInterface(rtout, obj, tin, obj2methodFuncs)
		}
	}
	// extract object from tin proxy or emulated interface (if any),
	// and wrap it in tout emulated interface
	return func(obj r.Value) r.Value {
		v, t := extractor(obj)
		return xr.ToEmulatedInterface(rtout, v, t, obj2methodFuncs)
	}
}

// return a function that extracts value wrapped in a proxy or emulated interface
// returns nil if no extraction is needed
func (g *CompGlobals) extractor(tin xr.Type) func(r.Value) (r.Value, xr.Type) {
	if tin.Kind() != r.Interface {
		return nil
	} else if xr.IsEmulatedInterface(tin) {
		return xr.FromEmulatedInterface
	} else {
		return g.extractFromProxy
	}
}

// return the error "\n\treason: t does not implement tinterf: missing method <method>"
func interfaceMissingMethod(t, tinterf xr.Type) string {
	var s string
	if tinterf.Kind() == r.Interface {
		s = fmt.Sprintf("\n\treason: %v does not implement %v", t, tinterf)
		missingmtd := xr.MissingMethod(t, tinterf)
		if missingmtd != nil {
			s = fmt.Sprintf("%s: missing method %s", s, missingmtd.String())
		}
	}
	return s
}
