// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package etoken

// enable C++-style generics?
const GENERICS_V1_CXX = false

// enable generics "contracts are interfaces" ?
const GENERICS_V2_CTI = false

// can only enable one style of generics
func init() {
	if GENERICS_V1_CXX && GENERICS_V2_CTI {
		panic("github.com/cosmos72/gomacro/go/etoken: cannot enable both GENERICS_V1_CXX and GENERICS_V2_CTI. Please disable at least one of them.")
	}
}
