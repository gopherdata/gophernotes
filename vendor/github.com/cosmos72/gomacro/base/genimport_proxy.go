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
	"bytes"
	"fmt"
	"go/types"
	"strings"
)

type writeTypeOpts int

const (
	writeMethodsAsFields writeTypeOpts = 1 << iota
	writeForceParamNames
	writeIncludeParamTypes
)

func writeInterfaceProxy(out *bytes.Buffer, pkgPath string, pkgSuffix string, name string, t *types.Interface) {
	fmt.Fprintf(out, "\n// --------------- proxy for %s.%s ---------------\ntype %s%s struct {", pkgPath, name, name, pkgSuffix)
	writeInterfaceMethods(out, pkgSuffix, name, t, writeMethodsAsFields)
	out.WriteString("\n}\n")
	writeInterfaceMethods(out, pkgSuffix, name, t, writeForceParamNames)
}

func writeInterfaceMethods(out *bytes.Buffer, pkgSuffix string, name string, t *types.Interface, opts writeTypeOpts) {
	if opts&writeMethodsAsFields != 0 {
		fmt.Fprint(out, "\n\tObject\tinterface{}") // will be used to retrieve object wrapped in the proxy
	}
	n := t.NumMethods()
	for i := 0; i < n; i++ {
		writeInterfaceMethod(out, pkgSuffix, name, t.Method(i), opts)
	}
}

func writeInterfaceMethod(out *bytes.Buffer, pkgSuffix string, interfaceName string, method *types.Func, opts writeTypeOpts) {
	if !method.Exported() {
		return
	}
	sig, ok := method.Type().(*types.Signature)
	if !ok {
		return
	}
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
		fmt.Fprintf(out, "func (Proxy *%s%s) %s(", interfaceName, pkgSuffix, method.Name())
	}
	results := sig.Results()
	writeTypeTuple(out, params, opts|writeIncludeParamTypes)
	out.WriteString(") ")
	writeTypeTupleOut(out, results)
	if opts&writeMethodsAsFields != 0 {
		return
	}
	out.WriteString(" {\n\t")
	if results != nil && results.Len() > 0 {
		out.WriteString("return ")
	}
	fmt.Fprintf(out, "Proxy.%s_(Proxy.Object", method.Name())
	if params != nil && params.Len() != 0 {
		out.WriteString(", ")
	}
	writeTypeTuple(out, params, writeForceParamNames)
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

func writeTypeTupleOut(out *bytes.Buffer, tuple *types.Tuple) {
	if tuple == nil || tuple.Len() == 0 {
		return
	}
	ret0 := tuple.At(0)
	if tuple.Len() > 1 || len(ret0.Name()) > 0 {
		out.WriteString("(")
		writeTypeTuple(out, tuple, writeIncludeParamTypes)
		out.WriteString(")")
	} else {
		types.WriteType(out, ret0.Type(), packageNameQualifier)
	}
}

func writeTypeTuple(out *bytes.Buffer, tuple *types.Tuple, opts writeTypeOpts) {
	n := tuple.Len()
	for i := 0; i < n; i++ {
		if i != 0 {
			out.WriteString(", ")
		}
		writeTypeVar(out, tuple.At(i), i, opts)
	}
}

func writeTypeVar(out *bytes.Buffer, v *types.Var, index int, opts writeTypeOpts) {
	name := v.Name()
	if len(name) == 0 && opts&writeForceParamNames != 0 {
		name = fmt.Sprintf("unnamed%d", index)
	}
	out.WriteString(name)
	if opts&writeIncludeParamTypes != 0 {
		if len(name) != 0 {
			out.WriteString(" ")
		}
		types.WriteType(out, v.Type(), packageNameQualifier)
	}
}

func packageNameQualifier(pkg *types.Package) string {
	path := pkg.Path()
	return path[1+strings.LastIndexByte(path, '/'):]
}
