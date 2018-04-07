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
 * import.go
 *
 *  Created on Apr 02, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/types"
	r "reflect"
	"strconv"
	"strings"

	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// Import compiles an import statement
func (c *Comp) Import(node ast.Spec) {
	switch node := node.(type) {
	case *ast.ImportSpec:
		str := node.Path.Value
		path, err := strconv.Unquote(str)
		if err != nil {
			c.Errorf("error unescaping import path %q: %v", str, err)
		}
		path = c.sanitizeImportPath(path)
		var name string
		if node.Name != nil {
			name = node.Name.Name
		} else {
			name = FileName(path)
		}
		// yes, we support local imports
		// i.e. a function or block can import packages
		c.ImportPackage(name, path)
	default:
		c.Errorf("unimplemented import: %v", node)
	}
}

func (g *CompGlobals) sanitizeImportPath(path string) string {
	path = strings.Replace(path, "\\", "/", -1)
	l := len(path)
	if path == ".." || l >= 3 && (path[:3] == "../" || path[l-3:] == "/..") || strings.Contains(path, "/../") {
		g.Errorf("invalid import %q: contains \"..\"", path)
	}
	if path == "." || l >= 2 && (path[:2] == "./" || path[l-2:] == "/.") || strings.Contains(path, "/./") {
		g.Errorf("invalid import %q: contains \".\"", path)
	}
	return path
}

// ImportPackage imports a package. Usually invoked as Comp.FileComp().ImportPackage(name, path)
// because imports are usually top-level statements in a source file.
// But we also support local imports, i.e. import statements inside a function or block.
func (c *Comp) ImportPackage(name, path string) *Import {
	g := c.CompGlobals
	pkgref := g.ImportPackage(name, path)
	if pkgref == nil {
		return nil
	}
	binds, bindtypes := g.parseImportBinds(pkgref.Binds, pkgref.Untypeds)

	imp := Import{
		Binds:     binds,
		BindTypes: bindtypes,
		Types:     g.parseImportTypes(pkgref.Types, pkgref.Wrappers),
		Name:      name,
		Path:      path,
	}

	g.loadProxies(pkgref.Proxies, imp.Types)

	c.declImport0(name, imp)
	return &imp
}

// declImport0 compiles an import declaration.
// Note: does not loads proxies, use ImportPackage for that
func (c *Comp) declImport0(name string, imp Import) {
	// treat imported package as a constant,
	// because to compile code we need the declarations it contains:
	// importing them at runtime would be too late.
	t := c.TypeOfImport()
	bind := c.AddBind(name, ConstBind, t)
	bind.Value = imp // cfile.Binds[] is a map[string]*Bind => changes to *Bind propagate to the map
}

func (g *CompGlobals) parseImportBinds(binds map[string]r.Value, untypeds map[string]string) (map[string]r.Value, map[string]xr.Type) {
	retbinds := make(map[string]r.Value)
	rettypes := make(map[string]xr.Type)
	for name, bind := range binds {
		if untyped, ok := untypeds[name]; ok {
			value, typ, ok := g.parseImportUntyped(untyped)
			if ok {
				retbinds[name] = value
				rettypes[name] = typ
				continue
			}
		}
		retbinds[name] = bind
		rettypes[name] = g.Universe.FromReflectType(bind.Type())
	}
	return retbinds, rettypes
}

func (g *CompGlobals) parseImportUntyped(untyped string) (r.Value, xr.Type, bool) {
	gkind, value := UnmarshalUntyped(untyped)
	if gkind == types.Invalid {
		return Nil, nil, false
	}
	lit := UntypedLit{Kind: xr.ToReflectKind(gkind), Obj: value, BasicTypes: &g.Universe.BasicTypes}
	return r.ValueOf(lit), g.TypeOfUntypedLit(), true
}

func (g *CompGlobals) parseImportTypes(rtypes map[string]r.Type, wrappers map[string][]string) map[string]xr.Type {
	v := g.Universe
	xtypes := make(map[string]xr.Type)
	for name, rtype := range rtypes {
		// Universe.FromReflectType uses cached *types.Package if possible
		t := v.FromReflectType(rtype)
		if twrappers := wrappers[name]; len(twrappers) != 0 {
			t.RemoveMethods(twrappers, "")
		}
		xtypes[name] = t
	}
	return xtypes
}

// loadProxies adds to thread-global maps the proxies found in import
func (g *CompGlobals) loadProxies(proxies map[string]r.Type, xtypes map[string]xr.Type) {
	for name, proxy := range proxies {
		xtype := xtypes[name]
		if xtype == nil {
			g.Warnf("import %q: type not found for proxy <%v>", proxy.PkgPath(), proxy)
			continue
		}
		if xtype.Kind() != r.Interface {
			g.Warnf("import %q: type for proxy <%v> is not an interface: %v", proxy.PkgPath(), proxy, xtype)
			continue
		}
		rtype := xtype.ReflectType()
		g.interf2proxy[rtype] = proxy
		g.proxy2interf[proxy] = xtype
	}
}

// v is an imported variable. build a function that will return it.
// Do NOT expose its value while compiling, otherwise the fast interpreter
// will (incorrectly) assume that it's a constant and will perform constant propagation.
//
// mandatory optimization: for basic kinds, unwrap reflect.Value
func importedBindAsFun(t xr.Type, v r.Value) I {
	var fun I
	switch t.Kind() {
	case r.Bool:
		fun = func(env *Env) bool {
			return v.Bool()
		}
	case r.Int:
		fun = func(env *Env) int {
			return int(v.Int())
		}
	case r.Int8:
		fun = func(env *Env) int8 {
			return int8(v.Int())
		}
	case r.Int16:
		fun = func(env *Env) int16 {
			return int16(v.Int())
		}
	case r.Int32:
		fun = func(env *Env) int32 {
			return int32(v.Int())
		}
	case r.Int64:
		fun = func(env *Env) int64 {
			return v.Int()
		}
	case r.Uint:
		fun = func(env *Env) uint {
			return uint(v.Uint())
		}
	case r.Uint8:
		fun = func(env *Env) uint8 {
			return uint8(v.Uint())
		}
	case r.Uint16:
		fun = func(env *Env) uint16 {
			return uint16(v.Uint())
		}
	case r.Uint32:
		fun = func(env *Env) uint32 {
			return uint32(v.Uint())
		}
	case r.Uint64:
		fun = func(env *Env) uint64 {
			return v.Uint()
		}
	case r.Uintptr:
		fun = func(env *Env) uintptr {
			return uintptr(v.Uint())
		}
	case r.Float32:
		fun = func(env *Env) float32 {
			return float32(v.Float())
		}
	case r.Float64:
		fun = func(env *Env) float64 {
			return v.Float()
		}
	case r.Complex64:
		fun = func(env *Env) complex64 {
			return complex64(v.Complex())
		}
	case r.Complex128:
		fun = func(env *Env) complex128 {
			return v.Complex()
		}
	case r.String:
		fun = func(env *Env) string {
			return v.String()
		}
	default:
		fun = func(env *Env) r.Value {
			return v
		}
	}
	return fun
}
