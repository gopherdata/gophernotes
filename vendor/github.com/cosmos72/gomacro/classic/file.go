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
 * file.go
 *
 *  Created on: Feb 15, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"bufio"
	"go/ast"
	"os"

	. "github.com/cosmos72/gomacro/base"
)

func (ir *Interp) EvalFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		ir.Errorf("error opening file '%s': %v", filePath, err)
		return
	}
	defer file.Close()

	saveOpts := ir.Env.Options
	ir.Env.Options &^= OptShowEval

	defer func() {
		ir.Env.Options = saveOpts
	}()

	in := bufio.NewReader(file)
	ir.Repl(in)
}

func (env *Env) evalFile(node *ast.File) {
	env.Name = node.Name.Name
	env.Path = env.Name
	env.PackagePath = env.Name

	for _, imp := range node.Imports {
		env.evalImport(imp)
	}

	for _, decl := range node.Decls {
		env.evalDecl(decl)
	}
}
