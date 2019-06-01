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
 * inspect.go
 *
 *  Created on: Feb 11, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (env *Env) Inspect(str string) {
	inspector := env.Globals.Inspector
	if inspector == nil {
		env.Errorf("no inspector set: call Interp.SetInspector() first")
		return
	}

	form := env.Parse(str)
	v := env.EvalAst1(form)
	var t r.Type
	if v.IsValid() && v != None {
		if v.Kind() == r.Interface {
			v = v.Elem() // extract concrete type
		}
		t = v.Type()
	}
	inspector.Inspect(str, v, t, nil, env.Globals)
}
