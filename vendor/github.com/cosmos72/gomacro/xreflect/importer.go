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
	gotypes "go/types"

	"github.com/cosmos72/gomacro/go/types"
)

type Importer struct {
	// converts from go/types to github.com/cosmos72/gomacro/go/types
	Converter types.Converter
	from      gotypes.ImporterFrom
	compat    gotypes.Importer
	srcDir    string
	mode      gotypes.ImportMode
}

func DefaultImporter() *Importer {
	imp := Importer{}
	imp.Converter.Init(types.Universe)
	compat := importer.Default()
	if from, ok := compat.(gotypes.ImporterFrom); ok {
		imp.from = from
	} else {
		imp.compat = compat
	}
	return &imp
}

func (imp *Importer) Import(path string) (*types.Package, error) {
	return imp.ImportFrom(path, imp.srcDir, imp.mode)
}

func (imp *Importer) ImportFrom(path string, srcDir string, mode gotypes.ImportMode) (*types.Package, error) {
	var pkg *gotypes.Package
	var err error
	if imp.from != nil {
		pkg, err = imp.from.ImportFrom(path, srcDir, mode)
	} else if imp.compat != nil {
		pkg, err = imp.compat.Import(path)
	} else {
		return nil, errors.New(fmt.Sprintf("importer.Default() returned nil, cannot import %q", path))
	}
	return imp.Converter.Package(pkg), err
}
