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
 * identifier.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (env *Env) evalIdentifier(ident *ast.Ident) r.Value {
	value, found := env.resolveIdentifier(ident)
	if !found {
		env.Errorf("undefined identifier: %s", ident.Name)
	}
	return value
}

func (env *Env) resolveIdentifier(ident *ast.Ident) (r.Value, bool) {
	name := ident.Name
	value := Nil
	found := false
	for e := env; e != nil; e = e.Outer {
		// Debugf("evalIdentifier() looking up %#v in %#v", name, env.Binds)
		if value, found = e.Binds.Get(name); found {
			break
		}
	}
	return value, found
}
