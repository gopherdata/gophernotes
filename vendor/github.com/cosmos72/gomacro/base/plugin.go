// +build go1.8,linux,!android,!gccgo

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
	"plugin"
	"strings"
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

func (g *Globals) compilePlugin(filename string, stdout io.Writer, stderr io.Writer) string {
	gosrcdir := getGoSrcPath()
	gosrclen := len(gosrcdir)
	filelen := len(filename)
	if filelen < gosrclen || filename[0:gosrclen] != gosrcdir {
		g.Errorf("source %q is in unsupported directory, cannot compile it: should be inside %q", filename, gosrcdir)
	}

	cmd := exec.Command("go", "build", "-buildmode=plugin")
	cmd.Dir = filename[0 : 1+strings.LastIndexByte(filename, '/')]
	cmd.Stdin = nil
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	g.Debugf("compiling %q ...", filename)
	err := cmd.Run()
	if err != nil {
		g.Errorf("error executing \"go build -buildmode=plugin\" in directory %q: %v", cmd.Dir, err)
	}

	dirname := filename[:strings.LastIndexByte(filename, '/')]
	// go build uses innermost directory name as shared object name,
	// i.e.	foo/bar/main.go is compiled to foo/bar/bar.so
	filename = dirname[1+strings.LastIndexByte(dirname, '/'):]

	return fmt.Sprintf("%s/%s.so", dirname, filename)
}

func (g *Globals) loadPlugin(soname string, symbolName string) interface{} {
	pkg, err := plugin.Open(soname)
	if err != nil {
		g.Errorf("error loading plugin %q: %v", soname, err)
	}
	val, err := pkg.Lookup(symbolName)
	if err != nil {
		g.Errorf("error loading symbol %q from plugin %q: %v", symbolName, soname, err)
	}
	return val
}
