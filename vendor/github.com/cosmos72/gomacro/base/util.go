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
 * util.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	r "reflect"
	"strings"

	xr "github.com/cosmos72/gomacro/xreflect"
)

func PackValues(val0 r.Value, values []r.Value) []r.Value {
	if len(values) == 0 && val0 != None {
		values = []r.Value{val0}
	}
	return values
}

func PackTypes(typ0 xr.Type, types []xr.Type) []xr.Type {
	if len(types) == 0 && typ0 != nil {
		types = []xr.Type{typ0}
	}
	return types
}

func PackValuesAndTypes(val0 r.Value, values []r.Value, typ0 xr.Type, types []xr.Type) ([]r.Value, []xr.Type) {
	return PackValues(val0, values), PackTypes(typ0, types)
}

func UnpackValues(vals []r.Value) (r.Value, []r.Value) {
	val0 := None
	if len(vals) > 0 {
		val0 = vals[0]
	}
	return val0, vals
}

// ValueInterface() is a zero-value-safe version of reflect.Value.Interface()
func ValueInterface(v r.Value) interface{} {
	if !v.IsValid() || !v.CanInterface() || v == None {
		return nil
	}
	return v.Interface()
}

// ValueType() is a zero-value-safe version of reflect.Value.Type()
func ValueType(value r.Value) r.Type {
	if !value.IsValid() || value == None {
		return nil
	}
	return value.Type()
}

func IsNillableKind(k r.Kind) bool {
	switch k {
	case r.Invalid, // nil is nillable...
		r.Chan, r.Func, r.Interface, r.Map, r.Ptr, r.Slice:
		return true
	default:
		return false
	}
}

// split 's' into a prefix and suffix separated by 'separator'.
// suffix is trimmed with strings.TrimSpace() before returning it
func Split2(s string, separator rune) (string, string) {
	var prefix, suffix string
	if space := strings.IndexByte(s, ' '); space > 0 {
		prefix = s[:space]
		suffix = strings.TrimSpace(s[space+1:])
	} else {
		prefix = s
	}
	return prefix, suffix
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
	GoPkg = filepath.FromSlash("github.com/gopherdata/gophernotes/vendor/github.com/cosmos72/gomacro")

	GoSrcDir = Subdir(filepath.SplitList(build.Default.GOPATH)[0], "src")

	GomacroDir = findGomacroDir(GoPkg)
)

func findGomacroDir(pkg string) string {
	gopath := build.Default.GOPATH
	for _, dir := range filepath.SplitList(gopath) {
		path := filepath.Join(dir, "src", pkg)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	defaultDir := Subdir(GoSrcDir, pkg)
	fmt.Printf("// warning: could not find package %q in $GOPATH = %q, assuming package is located in %q\n", pkg, gopath, defaultDir)
	return defaultDir
}
