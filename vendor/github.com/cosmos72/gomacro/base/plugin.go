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
 * plugin.go
 *
 *  Created on Feb 27, 2017
 *      Author Massimiliano Ghilardi
 */

package base

import (
	"io"
	"os/exec"
	r "reflect"
)

func (g *Globals) compilePlugin(filepath string, stdout io.Writer, stderr io.Writer) string {
	gosrcdir := GoSrcDir
	gosrclen := len(gosrcdir)
	filelen := len(filepath)
	if filelen < gosrclen || filepath[0:gosrclen] != gosrcdir {
		g.Errorf("source %q is in unsupported directory, cannot compile it: should be inside %q", filepath, gosrcdir)
	}

	cmd := exec.Command("go", "build", "-buildmode=plugin")
	cmd.Dir = DirName(filepath)
	cmd.Stdin = nil
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	g.Debugf("compiling %q ...", filepath)
	err := cmd.Run()
	if err != nil {
		g.Errorf("error executing \"go build -buildmode=plugin\" in directory %q: %v", cmd.Dir, err)
	}

	dirname := RemoveLastByte(DirName(filepath))
	// go build uses innermost directory name as shared object name,
	// i.e.	foo/bar/main.go is compiled to foo/bar/bar.so
	filename := FileName(dirname)

	return Subdir(dirname, filename+".so")
}

func (g *Globals) loadPluginSymbol(soname string, symbolName string) interface{} {
	// use imports.Packages["plugin"].Binds["Open"] and reflection instead of hard-coding call to plugin.Open()
	// reasons:
	// * import ( "plugin" ) does not work on all platforms (creates broken gomacro.exe on Windows/386)
	// * allow caller to provide us with a different implementation,
	//   either in imports.Packages["plugin"].Binds["Open"]
	//   or in Globals.Importer.PluginOpen

	imp := g.Importer
	if !imp.setPluginOpen() {
		g.Errorf("gomacro compiled without support to load plugins - requires Go 1.8+ and Linux - cannot import packages at runtime")
	}
	if len(soname) == 0 || len(symbolName) == 0 {
		// caller is just checking whether PluginOpen() is available
		return nil
	}
	so, err := reflectcall(imp.PluginOpen, soname)
	if err != nil {
		g.Errorf("error loading plugin %q: %v", soname, err)
	}
	vsym, err := reflectcall(so.MethodByName("Lookup"), symbolName)
	if err != nil {
		g.Errorf("error loading symbol %q from plugin %q: %v", symbolName, soname, err)
	}
	return vsym.Interface()
}

func reflectcall(fun r.Value, arg interface{}) (r.Value, interface{}) {
	vs := fun.Call([]r.Value{r.ValueOf(arg)})
	return vs[0], vs[1].Interface()
}
