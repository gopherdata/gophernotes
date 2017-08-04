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
 * importer.go
 *
 *  Created on May 14, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"errors"
	"fmt"
	"go/importer"
	"go/types"
)

type Importer struct {
	from   types.ImporterFrom
	compat types.Importer
	srcDir string
	mode   types.ImportMode
}

func DefaultImporter() *Importer {
	imp := Importer{}
	compat := importer.Default()
	if from, ok := compat.(types.ImporterFrom); ok {
		imp.from = from
	} else {
		imp.compat = compat
	}
	return &imp
}

func (imp *Importer) Import(path string) (*types.Package, error) {
	return imp.ImportFrom(path, imp.srcDir, imp.mode)
}

func (imp *Importer) ImportFrom(path string, srcDir string, mode types.ImportMode) (*types.Package, error) {
	if imp.from != nil {
		return imp.from.ImportFrom(path, srcDir, mode)
	} else if imp.compat != nil {
		return imp.compat.Import(path)
	} else {
		return nil, errors.New(fmt.Sprintf("importer.Default() returned nil, cannot import %q", path))
	}
}
