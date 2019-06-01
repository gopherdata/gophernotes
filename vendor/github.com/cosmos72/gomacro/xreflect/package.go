/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * package.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"strings"

	"github.com/cosmos72/gomacro/go/types"
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
