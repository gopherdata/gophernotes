/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * paths.go
 *
 *  Created on: Jun 24, 2018
 *      Author: Massimiliano Ghilardi
 */

package paths

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

// return the string after last '/' in path
func FileName(path string) string {
	return path[1+strings.LastIndexByte(path, '/'):]
}

// return the string up to (and including) last '/' in path
func DirName(path string) string {
	return path[0 : 1+strings.LastIndexByte(path, '/')]
}

// remove last byte from string
func RemoveLastByte(s string) string {
	if n := len(s); n != 0 {
		s = s[:n-1]
	}
	return s
}

// always use forward slashes. they work also on Windows...
func unixpath(path string) string {
	if os.PathSeparator != '/' && len(path) != 0 {
		path = strings.Replace(path, string(os.PathSeparator), "/", -1)
	}
	return path
}

// find user's home directory, see https://stackoverflow.com/questions/2552416/how-can-i-find-the-users-home-dir-in-a-cross-platform-manner-using-c
// without importing "os/user" - which requires cgo to work thus makes cross-compile difficult, see https://github.com/golang/go/issues/11797
func UserHomeDir() string {
	home := os.Getenv("HOME")
	if len(home) == 0 {
		home = os.Getenv("USERPROFILE")
		if len(home) == 0 {
			home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		}
	}
	return unixpath(home)
}

func Subdir(dirs ...string) string {
	// should use string(os.PathSeparator) instead of "/', but:
	// 1) package names use '/', not os.PathSeparator
	// 2) it would complicate DirName()
	return strings.Join(dirs, "/")
}

var (
	GoSrcDir = filepath.Join(filepath.SplitList(build.Default.GOPATH)[0], "src")

	// where to find the Go compiler used to compile gomacro.
	// needed to build compatible plugins
	GoRootDir = build.Default.GOROOT

	// also works for vendored or cloned copies of gomacro
	GomacroDir = findGomacroDir(getGomacroPkg())
)

func getGomacroPkg() string {
	type dummy struct{}
	path := strings.Split(reflect.TypeOf(dummy{}).PkgPath(), "/")
	return strings.Join(path[0:len(path)-2], "/") // skip .../base/paths
}

func findGomacroDir(pkg string) string {
	pkg = filepath.Join(strings.Split(pkg, "/")...)
	gopath := build.Default.GOPATH
	for _, dir := range filepath.SplitList(gopath) {
		path := filepath.Join(dir, "src", pkg)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	defaultDir := filepath.Join(GoSrcDir, pkg)
	fmt.Printf("// warning: could not find package %q in $GOPATH = %q, assuming package is located in %q\n", pkg, gopath, defaultDir)
	return defaultDir
}
