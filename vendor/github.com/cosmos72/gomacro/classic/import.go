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
func (env *Env) evalImport(node ast.Spec) (r.Value, []r.Value) {
	switch node := node.(type) {
	case *ast.ImportSpec:
		path := UnescapeString(node.Path.Value)
		path = env.sanitizeImportPath(path)
		var name string
		if node.Name != nil {
			name = node.Name.Name
		} else {
			name = FileName(path)
		}
		pkg := env.ImportPackage(name, path)
		if pkg != nil {
			fileEnv := env.FileEnv()
			fileEnv.DefineConst(name, r.TypeOf(pkg), r.ValueOf(pkg))
		}
		return r.ValueOf(path), nil
	default:
		return env.Errorf("unimplemented import: %v", node)
	}
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
