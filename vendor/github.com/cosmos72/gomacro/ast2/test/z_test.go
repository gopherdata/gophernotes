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
 *  Created on: May 05, 2018
 *      Author: Massimiliano Ghilardi
 */

package test

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/cosmos72/gomacro/ast2"
	"github.com/cosmos72/gomacro/base/output"
	"github.com/cosmos72/gomacro/go/etoken"
	"github.com/cosmos72/gomacro/go/parser"
)

func TestToNodes(t *testing.T) {
	tests := []struct {
		Name string
		Path string
	}{
		{"z_test_data_2", "z_test_data_2.txt"},
		{"fast_global", "../../fast/global.go"},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			_testToNodes(t, test.Path)
		})
	}
}

func _testToNodes(t *testing.T, filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("read file %q failed: %v", filename, err)
		return
	}

	fset := etoken.NewFileSet()
	st := output.Stringer{Fileset: fset}

	var p parser.Parser
	p.Init(fset, filename, 0, bytes)

	nodes, err := p.Parse()
	if err != nil {
		t.Errorf("parse file %q failed: %v", filename, err)
		return
	}
	nodes = ToNodes(NodeSlice{nodes})

	for _, node := range nodes {
		st.Fprintf(os.Stdout, "%v\n", node)
	}
}
