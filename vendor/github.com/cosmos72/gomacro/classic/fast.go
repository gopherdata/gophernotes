/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * fast.go
 *
 *  Created on: Apr 02, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	r "reflect"

	"github.com/cosmos72/gomacro/ast2"
	"github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/fast"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// temporary helper to invoke the new fast interpreter.
// executes macroexpand + collect + compile + eval
func (env *Env) fastEval(form ast2.Ast) (r.Value, []r.Value, xr.Type, []xr.Type) {
	var f *fast.Interp
	if env.FastInterp == nil {
		f = fast.New()
		f.Comp.CompileOptions |= fast.OptKeepUntyped
		env.FastInterp = f
	} else {
		f = env.FastInterp.(*fast.Interp)
	}
	f.Comp.Stringer.Copy(&env.Stringer) // sync Fileset, Pos, Line
	f.Comp.Options = env.Options        // sync Options

	// macroexpand phase.
	// must be performed manually, because we used classic.Env.ParseOnly()
	// instead of fast.Comp.Parse()
	form, _ = f.Comp.MacroExpandCodewalk(form)
	if env.Options&base.OptShowMacroExpand != 0 {
		env.Debugf("after macroexpansion: %v", form.Interface())
	}

	// collect phase
	if env.Options&(base.OptCollectDeclarations|base.OptCollectStatements) != 0 {
		env.CollectAst(form)
	}

	if env.Options&base.OptMacroExpandOnly != 0 {
		x := form.Interface()
		return r.ValueOf(x), nil, f.Comp.TypeOf(x), nil
	}

	// compile phase
	expr := f.Comp.Compile(form)
	if env.Options&base.OptShowCompile != 0 {
		env.Fprintf(env.Stdout, "%v\n", expr)
	}

	// eval phase
	if expr == nil {
		return base.None, nil, nil, nil
	}
	value, values := f.RunExpr(expr)
	return value, values, expr.Type, expr.Types
}
