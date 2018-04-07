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
 * package.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/types"
	"strings"
)

func (v *Universe) loadPackage(path string) *Package {
	if len(path) == 0 {
		// do not create unnamed packages
		return nil
	}
	// try the importer and its cache
	if pkg := v.importPackage(path); pkg != nil {
		return pkg
	}
	// no luck. create and return an empty Package
	if v.Packages == nil {
		v.Packages = make(map[string]*Package)
	}
	name := path[1+strings.LastIndexByte(path, '/'):]
	pkg := (*Package)(types.NewPackage(path, name))
	v.Packages[path] = pkg
	return pkg
}

func (v *Universe) LoadPackage(path string) *Package {
	if len(path) == 0 {
		// do not create unnamed packages
		return nil
	}
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return v.loadPackage(path)
}

func (pkg *Package) GoPackage() *types.Package {
	return (*types.Package)(pkg)
}

func (pkg *Package) Name() string {
	if pkg == nil {
		return ""
	}
	return (*types.Package)(pkg).Name()
}

func (pkg *Package) Path() string {
	if pkg == nil {
		return ""
	}
	return (*types.Package)(pkg).Path()
}
