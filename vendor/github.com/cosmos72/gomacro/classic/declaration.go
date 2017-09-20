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
 * declaration.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"fmt"
	"go/ast"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (env *Env) evalDecl(node ast.Decl) (r.Value, []r.Value) {
	switch node := node.(type) {
	case *ast.GenDecl:
		return env.evalDeclGen(node)
	case *ast.FuncDecl:
		return env.evalDeclFunction(node, node.Type, node.Body)
	default:
		return env.Errorf("unimplemented declaration: %v", node)
	}
}

func (env *Env) evalDeclGen(node *ast.GenDecl) (r.Value, []r.Value) {
	tok := node.Tok
	var ret r.Value
	var rets []r.Value
	switch tok {
	case token.IMPORT:
		for _, decl := range node.Specs {
			ret, rets = env.evalImport(decl)
		}
	case token.CONST:
		var defaultType ast.Expr
		var defaultExprs []ast.Expr
		top := env.TopEnv()
		top.addIota()
		defer top.removeIota()
		for _, decl := range node.Specs {
			ret, rets = env.evalDeclConsts(decl, defaultType, defaultExprs)
			if valueSpec, ok := decl.(*ast.ValueSpec); ok && valueSpec.Values != nil {
				defaultType = valueSpec.Type
				defaultExprs = valueSpec.Values
			}
			top.incrementIota()
		}
	case token.TYPE:
		for _, decl := range node.Specs {
			return env.evalDeclType(decl)
		}
	case token.VAR:
		for _, decl := range node.Specs {
			ret, rets = env.evalDeclVars(decl)
		}
	default:
		return env.Errorf("unimplemented declaration: %v", node)
	}
	return ret, rets
}

func (env *Env) evalDeclConsts(node ast.Spec, defaultType ast.Expr, defaultExprs []ast.Expr) (r.Value, []r.Value) {
	switch node := node.(type) {
	case *ast.ValueSpec:
		if node.Type != nil || node.Values != nil {
			defaultType = node.Type
			defaultExprs = node.Values
		}
		return env.evalDeclConstsOrVars(node.Names, defaultType, defaultExprs, true)
	default:
		return env.Errorf("unexpected constant declaration: expecting *ast.ValueSpec, found: %v <%v>", node, r.TypeOf(node))
	}
}

func (env *Env) evalDeclType(node ast.Spec) (r.Value, []r.Value) {
	switch node := node.(type) {
	case *ast.TypeSpec:
		name := node.Name.Name
		// PATCH: support type aliases
		if unary, ok := node.Type.(*ast.UnaryExpr); ok && unary.Op == token.ASSIGN {
			t := env.evalTypeAlias(name, unary.X)
			return r.ValueOf(&t).Elem(), nil // return a reflect.Type, not the concrete type
		}

		t := env.evalType(node.Type)
		if name != "_" {
			// never define bindings for "_"
			if _, ok := env.Types.Get(name); ok {
				env.Warnf("redefined type: %v", name)
			} else {
				env.Types.Ensure()
			}
			env.Types.Set(name, t)
			if _, ok := env.NamedTypes[t]; !ok {
				env.NamedTypes[t] = fmt.Sprintf("%s.%s", env.PackagePath, name)
			}
		}
		return r.ValueOf(&t).Elem(), nil // return a reflect.Type, not the concrete type

	default:
		return env.Errorf("unexpected type declaration: expecting *ast.TypeSpec, found: %v <%v>", node, r.TypeOf(node))
	}
}

func (env *Env) evalDeclVars(node ast.Spec) (r.Value, []r.Value) {
	switch node := node.(type) {
	case *ast.ValueSpec:
		return env.evalDeclConstsOrVars(node.Names, node.Type, node.Values, false)
	default:
		return env.Errorf("unexpected variable declaration: expecting *ast.ValueSpec, found: %v <%v>", node, r.TypeOf(node))
	}
}

func (env *Env) evalDeclConstsOrVars(idents []*ast.Ident, typ ast.Expr, exprs []ast.Expr, constant bool) (r.Value, []r.Value) {
	n := len(idents)
	names := make([]string, n)
	for i, ident := range idents {
		names[i] = ident.Name
	}
	t := env.evalType(typ)

	var values []r.Value
	if exprs != nil {
		values = env.evalExprsMultipleValues(exprs, n)
	}
	return env.defineConstsVarsOrFuncs(names, t, values, constant)
}

func (env *Env) defineConstsVarsOrFuncs(names []string, t r.Type, values []r.Value, constant bool) (r.Value, []r.Value) {
	n := len(names)
	if values == nil {
		if t == nil {
			return env.Errorf("no values and no type: cannot define %v", names)
		}
		values = make([]r.Value, n)
		zero := r.Zero(t)
		for i := 0; i < n; i++ {
			values[i] = env.defineConstVarOrFunc(names[i], t, zero, constant)
		}
	} else {
		for i := 0; i < n; i++ {
			values[i] = env.defineConstVarOrFunc(names[i], t, values[i], constant)
		}
	}
	return UnpackValues(values)
}

func (env *Env) DefineConst(name string, t r.Type, value r.Value) r.Value {
	return env.defineConstVarOrFunc(name, t, value, true)
}

func (env *Env) DefineVar(name string, t r.Type, value r.Value) r.Value {
	return env.defineConstVarOrFunc(name, t, value, false)
}

func (env *Env) DefineFunc(name string, t r.Type, value r.Value) r.Value {
	return env.defineConstVarOrFunc(name, t, value, true)
}

func (env *Env) defineConstVarOrFunc(name string, t r.Type, value r.Value, constant bool) r.Value {
	if name == "_" {
		// never define bindings for "_"
		if t != nil {
			value = env.valueToType(value, t)
		}
		return value
	}
	if t == nil {
		t = typeOf(value)
	}
	if _, found := env.Binds.Get(name); found {
		env.Warnf("redefined identifier: %v", name)
	} else {
		env.Binds.Ensure()
	}
	if constant {
		value = value.Convert(t)
		env.Binds.Set(name, value)
	} else {
		addr := r.New(t)
		value = env.assignPlace(placeType{addr.Elem(), Nil}, token.ASSIGN, value)
		env.Binds.Set(name, addr.Elem())
	}
	// Debugf("defineConstVarOrFunc() added %#v to %#v", name, env.Binds)
	return value
}
