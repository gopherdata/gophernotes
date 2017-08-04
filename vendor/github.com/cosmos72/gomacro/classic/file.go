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
 * file.go
 *
 *  Created on: Feb 15, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"go/ast"
	r "reflect"
)

func (env *Env) evalFile(node *ast.File) (r.Value, []r.Value) {
	env.PackagePath = node.Name.Name

	// TODO eval node.Imports
	var ret r.Value
	var rets []r.Value

	for _, decl := range node.Decls {
		ret, rets = env.evalDecl(decl)
	}
	return ret, rets
}
