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
 * importer.go
 *
 *  Created on Feb 27, 2017
 *      Author Massimiliano Ghilardi
 */

package base

import (
	"bytes"
	"errors"
	"fmt"
	"go/importer"
	"go/types"
	"io/ioutil"
	"os"
	r "reflect"

	"github.com/cosmos72/gomacro/imports"
)

type ImportMode int

const (
	// ImBuiltin import mechanism is:
	// 1. write a file $GOPATH/src/github.com/cosmos72/gomacro/imports/$PKGPATH.go containing a single func init()
	//    i.e. *inside* gomacro sources
	// 2. tell the user to recompile gomacro
	ImBuiltin ImportMode = iota

	// ImThirdParty import mechanism is the same as ImBuiltin, except that files are created in a thirdparty/ subdirectory:
	// 1. write a file $GOPATH/src/github.com/cosmos72/gomacro/imports/thirdparty/$PKGPATH.go containing a single func init()
	//    i.e. *inside* gomacro sources
	// 2. tell the user to recompile gomacro
	ImThirdParty

	// ImPlugin import mechanism is:
	// 1. write a file $GOPATH/src/gomacro_imports/$PKGPATH/$PKGNAME.go containing a var Packages map[string]Package
	//    and a single func init() to populate it
	// 2. invoke "go build -buildmode=plugin" on the file to create a shared library
	// 3. load such shared library with plugin.Open().Lookup("Packages")
	ImPlugin

	// ImInception import mechanism is:
	// 1. write a file $GOPATH/src/$PKGPATH/x_package.go containing a single func init()
	//    i.e. *inside* the package to be imported
	// 2. tell the user to recompile $PKGPATH
	ImInception
)

type Importer struct {
	from       types.ImporterFrom
	compat     types.Importer
	srcDir     string
	mode       types.ImportMode
	PluginOpen r.Value // = reflect.ValueOf(plugin.Open)
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

func (imp *Importer) setPluginOpen() bool {
	if imp.PluginOpen == Nil {
		imp.PluginOpen = imports.Packages["plugin"].Binds["Open"]
		if imp.PluginOpen == Nil {
			imp.PluginOpen = None // cache the failure
		}
	}
	return imp.PluginOpen != None
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

// LookupPackage returns a package if already present in cache
func (g *Globals) LookupPackage(name, path string) *PackageRef {
	pkg, found := imports.Packages[path]
	if !found {
		return nil
	}
	if len(name) == 0 {
		name = TailIdentifier(FileName(path))
	}
	return &PackageRef{Package: pkg, Name: name, Path: path}
}

func (g *Globals) ImportPackage(name, path string) *PackageRef {
	ref, err := g.ImportPackageOrError(name, path)
	if err != nil {
		panic(err)
	}
	return ref
}

func (g *Globals) ImportPackageOrError(name, path string) (*PackageRef, error) {
	ref := g.LookupPackage(name, path)
	if ref != nil {
		return ref, nil
	}
	gpkg, err := g.Importer.Import(path) // loads names and types, not the values!
	if err != nil {
		return nil, g.MakeRuntimeError(
			"error loading package %q metadata, maybe you need to download (go get), compile (go build) and install (go install) it? %v",
			path, err)
	}
	var mode ImportMode
	switch name {
	case "_b":
		mode = ImBuiltin
	case "_i":
		mode = ImInception
	case "_3":
		mode = ImThirdParty
	default:
		if len(name) == 0 {
			name = gpkg.Name()
		}
		havePluginOpen := g.Importer.setPluginOpen()
		if havePluginOpen {
			mode = ImPlugin
		} else {
			mode = ImThirdParty
		}
	}
	file := g.createImportFile(path, gpkg, mode)
	ref = &PackageRef{Name: name, Path: path}
	if len(file) == 0 || mode != ImPlugin {
		// either the package exports nothing, or user must rebuild gomacro.
		// in both cases, still cache it to avoid recreating the file.
		imports.Packages[path] = ref.Package
		return ref, nil
	}
	soname := g.compilePlugin(file, g.Stdout, g.Stderr)
	ipkgs := g.loadPluginSymbol(soname, "Packages")
	pkgs := *ipkgs.(*map[string]imports.PackageUnderlying)

	// cache *all* found packages for future use
	imports.Packages.Merge(pkgs)

	// but return only requested one
	pkg, found := imports.Packages[path]
	if !found {
		return nil, g.MakeRuntimeError(
			"error loading package %q: the compiled plugin %q does not contain it! internal error? %v",
			path, soname)
	}
	ref.Package = pkg
	return ref, nil
}

func (g *Globals) createImportFile(path string, pkg *types.Package, mode ImportMode) string {
	file := g.computeImportFilename(path, mode)

	buf := bytes.Buffer{}
	isEmpty := g.writeImportFile(&buf, path, pkg, mode)
	if isEmpty {
		g.Warnf("package %q exports zero constants, functions, types and variables", path)
		return ""
	}

	err := ioutil.WriteFile(file, buf.Bytes(), os.FileMode(0666))
	if err != nil {
		g.Errorf("error writing file %q: %v", file, err)
	}
	if mode == ImPlugin {
		g.Debugf("created file %q...", file)
	} else {
		g.Warnf("created file %q, recompile gomacro to use it", file)
	}
	return file
}

func sanitizeIdentifier(str string) string {
	return sanitizeIdentifier2(str, '_')
}

func sanitizeIdentifier2(str string, replacement rune) string {
	runes := []rune(str)
	for i, ch := range runes {
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || ch == '_' ||
			(i != 0 && ch >= '0' && ch <= '9') {
			continue
		}
		runes[i] = replacement
	}
	return string(runes)
}

func (g *Globals) computeImportFilename(path string, mode ImportMode) string {
	switch mode {
	case ImBuiltin:
		// user will need to recompile gomacro
		return Subdir(GomacroDir, "imports", sanitizeIdentifier(path)+".go")
	case ImInception:
		// user will need to recompile gosrcdir / path
		return Subdir(GoSrcDir, path, "x_package.go")
	case ImThirdParty:
		// either plugin.Open is not available, or user explicitly requested import _3 "package".
		// In both cases, user will need to recompile gomacro
		return Subdir(GomacroDir, "imports", "thirdparty", sanitizeIdentifier(path)+".go")
	}

	file := FileName(path) + ".go"
	file = Subdir(GoSrcDir, "gomacro_imports", path, file)
	dir := DirName(file)
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		Errorf("error creating directory %q: %v", dir, err)
	}
	return file
}
