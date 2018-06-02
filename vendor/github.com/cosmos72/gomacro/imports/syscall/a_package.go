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
 * a_package.go
 *
 *  Created on: Apr 09, 2018
 *      Author: Massimiliano Ghilardi
 */

package syscall

import (
	. "reflect"
)

type Package = struct { // unnamed
	Binds   map[string]Value
	Types   map[string]Type
	Proxies map[string]Type
	// Untypeds contains a string representation of untyped constants,
	// stored without loss of precision
	Untypeds map[string]string
	// Wrappers is the list of wrapper methods for named types.
	// Stored explicitly because reflect package cannot distinguish
	// between explicit methods and wrapper methods for embedded fields
	Wrappers map[string][]string
}

var Packages = make(map[string]Package)
