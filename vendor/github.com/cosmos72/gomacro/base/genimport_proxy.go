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
 * genimport_proxy.go
 *
 *  Created on Mar 06, 2017
 *      Author Massimiliano Ghilardi
 */

package base

import (
	"fmt"
	"go/types"
)

type writeTypeOpts int

const (
	writeMethodsAsFields writeTypeOpts = 1 << iota
	writeForceParamNames
	writeIncludeParamTypes
)

func (gen *genimport) writeInterfaceProxy(pkgPath string, name string, t *types.Interface) {
	fmt.Fprintf(gen.out, "\n// --------------- proxy for %s.%s ---------------\ntype %s%s struct {", pkgPath, name, gen.proxyprefix, name)
	gen.writeInterfaceMethods(name, t, writeMethodsAsFields)
	gen.out.WriteString("\n}\n")
	gen.writeInterfaceMethods(name, t, writeForceParamNames)
}

func (gen *genimport) writeInterfaceMethods(name string, t *types.Interface, opts writeTypeOpts) {
	if opts&writeMethodsAsFields != 0 {
		fmt.Fprint(gen.out, "\n\tObject\tinterface{}") // will be used to retrieve object wrapped in the proxy
	}
	n := t.NumMethods()
	for i := 0; i < n; i++ {
		gen.writeInterfaceMethod(name, t.Method(i), opts)
	}
}

func (gen *genimport) writeInterfaceMethod(interfaceName string, method *types.Func, opts writeTypeOpts) {
	if !method.Exported() {
		return
	}
	sig, ok := method.Type().(*types.Signature)
	if !ok {
		return
	}
	out := gen.out
	params := sig.Params()
	if opts&writeMethodsAsFields != 0 {
		var param0 string
		if opts&writeForceParamNames != 0 || isNamedTypeTuple(params) {
			param0 = "_proxy_obj_ "
		}
		fmt.Fprintf(out, "\n\t%s_\tfunc(%sinterface{}", method.Name(), param0)
		if params != nil && params.Len() != 0 {
			out.WriteString(", ")
		}
	} else {
		fmt.Fprintf(out, "func (P *%s%s) %s(", gen.proxyprefix, interfaceName, method.Name())
	}
	results := sig.Results()
	gen.writeTypeTuple(params, opts|writeIncludeParamTypes)
	out.WriteString(") ")
	gen.writeTypeTupleOut(results)
	if opts&writeMethodsAsFields != 0 {
		return
	}
	out.WriteString(" {\n\t")
	if results != nil && results.Len() > 0 {
		out.WriteString("return ")
	}
	fmt.Fprintf(out, "P.%s_(P.Object", method.Name())
	if params != nil && params.Len() != 0 {
		out.WriteString(", ")
	}
	gen.writeTypeTuple(params, writeForceParamNames)
	out.WriteString(")\n}\n")
}

func isNamedTypeTuple(tuple *types.Tuple) bool {
	if tuple == nil || tuple.Len() == 0 {
		return false
	}
	for i, n := 0, tuple.Len(); i < n; i++ {
		if len(tuple.At(i).Name()) != 0 {
			return true
		}
	}
	return false
}

func (gen *genimport) writeTypeTupleOut(tuple *types.Tuple) {
	if tuple == nil || tuple.Len() == 0 {
		return
	}
	out := gen.out
	ret0 := tuple.At(0)
	if tuple.Len() > 1 || len(ret0.Name()) > 0 {
		out.WriteString("(")
		gen.writeTypeTuple(tuple, writeIncludeParamTypes)
		out.WriteString(")")
	} else {
		types.WriteType(out, ret0.Type(), gen.packageNameQualifier)
	}
}

func (gen *genimport) writeTypeTuple(tuple *types.Tuple, opts writeTypeOpts) {
	n := tuple.Len()
	for i := 0; i < n; i++ {
		if i != 0 {
			gen.out.WriteString(", ")
		}
		gen.writeTypeVar(tuple.At(i), i, opts)
	}
}

func (gen *genimport) writeTypeVar(v *types.Var, index int, opts writeTypeOpts) {
	name := v.Name()
	if len(name) == 0 && opts&writeForceParamNames != 0 {
		name = fmt.Sprintf("unnamed%d", index)
	}
	out := gen.out
	out.WriteString(name)
	if opts&writeIncludeParamTypes != 0 {
		if len(name) != 0 {
			out.WriteString(" ")
		}
		types.WriteType(out, v.Type(), gen.packageNameQualifier)
	}
}

func (gen *genimport) packageNameQualifier(pkg *types.Package) string {
	path := pkg.Path()
	name, ok := gen.pkgrenames[path]
	if !ok {
		name = FileName(path)
	}
	return name
}
