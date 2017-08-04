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
 * genimport_wrapper.go
 *
 *  Created on May 26, 2017
 *      Author Massimiliano Ghilardi
 */

package base

import (
	"go/types"
	"sort"
)

// analyzer examines all methods of a named type and its embedded fields,
// and determines the wrapper methods for embedded fields.
type analyzer map[string][]int

func (a *analyzer) add(mtd *types.Func, depth int) {
	if *a == nil {
		(*a) = make(map[string][]int)
	}
	name := mtd.Name()
	(*a)[name] = append((*a)[name], depth)
}

func (a *analyzer) Analyze(t *types.Named) []string {
	a.analyze(t, 0)
	return a.wrappers()
}

func (a *analyzer) analyze(t *types.Named, depth int) {
	for i, n := 0, t.NumMethods(); i < n; i++ {
		mtd := t.Method(i)
		if mtd.Exported() {
			a.add(mtd, depth)
		}
	}

	if u, ok := t.Underlying().(*types.Struct); ok {
		for i, n := 0, u.NumFields(); i < n; i++ {
			if f := u.Field(i); f.Anonymous() {
				switch ft := f.Type().(type) {
				case *types.Named:
					a.analyze(ft, depth+1)
				case *types.Pointer:
					if ft, ok := ft.Elem().(*types.Named); ok {
						a.analyze(ft, depth+1)
					}
				}
			}
		}
	}
}

// listWrappers returns the names of wrapper methods for an analyzed type
func (a *analyzer) wrappers() []string {
	var wrappers []string
	for name, depths := range *a {
		if depths[0] == 0 {
			// explicit method declared in the outermost type. no ambiguity
			continue
		}
		wrappers = append(wrappers, name)
	}
	sort.Strings(wrappers)
	return wrappers
}
