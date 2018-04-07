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
 * plugin.go
 *
 *  Created on Feb 27, 2017
 *      Author Massimiliano Ghilardi
 */

package base

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	r "reflect"

	"github.com/cosmos72/gomacro/imports"
)

func getGoPath() string {
	dir := os.Getenv("GOPATH")
	if len(dir) == 0 {
		dir = os.Getenv("HOME")
		if len(dir) == 0 {
			Errorf("cannot determine go source directory: both $GOPATH and $HOME are unset or empty")
		}
		dir += "/go"
	}
	return dir
}

func getGoSrcPath() string {
	return getGoPath() + "/src"
}

func (g *Globals) compilePlugin(filepath string, stdout io.Writer, stderr io.Writer) string {
	// panics if plugin.Open is not available
	// -> skip generating .go file and compiling it
	g.loadPlugin("", "")

	gosrcdir := getGoSrcPath()
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

	return fmt.Sprintf("%s/%s.so", dirname, filename)
}

func (g *Globals) loadPlugin(soname string, symbolName string) interface{} {
	// use imports.Packages["plugin"] and reflection instead of hard-coding call to plugin.Open()
	// reasons:
	// * import ( "plugin" ) does not work on all platforms (creates broken gomacro.exe on Windows/386)
	// * allow caller to provide us with a different implementation in Imported.PluginOpen

	imp := g.Importer
	if imp.PluginOpen == Nil {
		imp.PluginOpen = imports.Packages["plugin"].Binds["Open"]
		if imp.PluginOpen == Nil {
			imp.PluginOpen = None // cache the failure
		}
	}
	if imp.PluginOpen == None {
		g.Errorf("gomacro compiled without support to load plugins - requires Go 1.8+ and Linux / Mac OS X - cannot import packages at runtime")
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
