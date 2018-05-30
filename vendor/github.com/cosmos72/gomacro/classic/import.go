/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * declaration.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"go/ast"
	r "reflect"
	"strings"

	. "github.com/cosmos72/gomacro/base"
)

// eval a single import
func (env *Env) evalImportDecl(decl ast.Spec) (r.Value, []r.Value) {
	switch node := decl.(type) {
	case *ast.ImportSpec:
		return env.evalImport(node)
	default:
		return env.Errorf("unimplemented import: %v", decl)
	}
}

// eval a single import
func (env *Env) evalImport(imp *ast.ImportSpec) (r.Value, []r.Value) {
	path := UnescapeString(imp.Path.Value)
	path = env.sanitizeImportPath(path)
	var name string
	if imp.Name != nil {
		name = imp.Name.Name
	} else {
		name = FileName(path)
	}
	pkg := env.ImportPackage(name, path)
	if pkg != nil {
		// if import appears *inside* a block, it is local for that block
		if name == "." {
			// dot import, i.e. import . "the/package/path"
			env.MergePackage(pkg.Package)
		} else {
			env.DefineConst(name, r.TypeOf(pkg), r.ValueOf(pkg))
		}
	}
	return r.ValueOf(name), nil
}

func (ir *ThreadGlobals) sanitizeImportPath(path string) string {
	path = strings.Replace(path, "\\", "/", -1)
	l := len(path)
	if path == ".." || l >= 3 && (path[:3] == "../" || path[l-3:] == "/..") || strings.Contains(path, "/../") {
		ir.Errorf("invalid import %q: contains \"..\"", path)
	}
	if path == "." || l >= 2 && (path[:2] == "./" || path[l-2:] == "/.") || strings.Contains(path, "/./") {
		ir.Errorf("invalid import %q: contains \".\"", path)
	}
	return path
}
