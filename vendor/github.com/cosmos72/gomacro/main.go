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
 * main.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"os"

	"github.com/cosmos72/gomacro/cmd"
)

func main() {
	args := os.Args[1:]

	cmd := cmd.New()

	err := cmd.Main(args)
	if err != nil {
		o := &cmd.Interp.Comp.Output
		o.Fprintf(o.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
