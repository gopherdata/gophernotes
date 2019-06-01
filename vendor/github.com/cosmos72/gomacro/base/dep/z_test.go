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
 * z_test.go
 *
 *  Created on: May 03, 2018
 *      Author: Massimiliano Ghilardi
 */

package dep

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/cosmos72/gomacro/go/etoken"
	"github.com/cosmos72/gomacro/go/parser"
)

func TestRemoveItem(t *testing.T) {
	list := []string{"Env", "Stmt"}
	out := remove_item_inplace("Stmt", list)
	expect := []string{"Env"}
	if !reflect.DeepEqual(out, expect) {
		t.Errorf("expected %v, actual %v", expect, out)
	}
}

func TestSortUnique1(t *testing.T) {
	in := []string{"c", "a", "c", "b", "a", "b", "x"}
	expect := []string{"a", "b", "c", "x"}
	_testSortUnique(t, in, expect)
}

func TestSortUnique2(t *testing.T) {
	in := []string{"Debugger", "Env", "IrGlobals", "Stmt", "Stmt", "poolCapacity"}
	expect := []string{"Debugger", "Env", "IrGlobals", "Stmt", "poolCapacity"}
	_testSortUnique(t, in, expect)
}

func _testSortUnique(t *testing.T, in []string, expect []string) {
	out := sort_unique_inplace(in)
	if !reflect.DeepEqual(out, expect) {
		t.Errorf("expected %v, actual %v", expect, out)
	}
}

func TestSorter(t *testing.T) {
	tests := []struct {
		Name string
		Path string
	}{
		{"api", "api.go"},
		{"z_test_data_1", "z_test_data_1.txt"},
		{"z_test_data_2", "z_test_data_2.txt"},
		{"z_test_data_3", "z_test_data_3.txt"},
		{"fast_global", "../../fast/global.go"},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			_testSorter(t, test.Path)
		})
	}
}

func _testSorter(t *testing.T, filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("read file %q failed: %v", filename, err)
		return
	}

	var p parser.Parser
	fset := etoken.NewFileSet()
	p.Init(fset, filename, 0, bytes)

	nodes, err := p.Parse()
	if err != nil {
		t.Errorf("parse file %q failed: %v", filename, err)
		return
	}
	s := NewSorter()
	s.LoadNodes(nodes)

	for {
		sorted := s.Some()
		if len(sorted) == 0 {
			break
		}
		fmt.Print("---- sorted decls ----\n")
		sorted.Print()
	}
}
