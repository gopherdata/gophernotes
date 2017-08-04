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
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"
)

type Assign struct {
	placefun func(*Env) r.Value
	placekey func(*Env) r.Value
	setvar   func(*Env, r.Value)
	setplace func(r.Value, r.Value, r.Value)
}

func (a *Assign) init(c *Comp, place *Place) {
	if place.IsVar() {
		a.setvar = c.varSetValue(&place.Var)
	} else {
		a.placefun = place.Fun
		a.placekey = place.MapKey
		a.setplace = c.placeSetValue(place)
	}
}

// Assign compiles an *ast.AssignStmt into an assignment to one or more place
func (c *Comp) Assign(node *ast.AssignStmt) {
	c.Pos = node.Pos()

	lhs, rhs := node.Lhs, node.Rhs
	if node.Tok == token.DEFINE {
		c.DeclVarsShort(lhs, rhs)
		return
	}
	ln, rn := len(lhs), len(rhs)
	if node.Tok == token.ASSIGN {
		if ln < 1 || (rn != 1 && ln != rn) {
			c.Errorf("invalid assignment, cannot assign %d values to %d places: %v", rn, ln, node)
		}
	} else if ln != 1 || rn != 1 {
		c.Errorf("invalid assignment, operator %s does not support multiple parallel assignments: %v", node.Tok, node)
	}

	// the naive loop
	//   for i := range lhs { c.assign1(lhs[i], node.Tok, rhs[i]) }
	// is bugged. It breaks, among others, the common Go idiom to swap two values: a,b = b,a
	//
	// More accurately, Go states at: https://golang.org/ref/spec#Assignments
	//
	// "The assignment proceeds in two phases. First, the operands of index expressions
	// and pointer indirections (including implicit pointer indirections in selectors)
	// on the left and the expressions on the right are all evaluated in the usual order.
	// Second, the assignments are carried out in left-to-right order."
	//
	// A solution is to evaluate left-to-right all places on the left,
	// then all expressions on the right, then perform all the assignments

	places := make([]*Place, ln)
	exprs := make([]*Expr, rn)
	canreorder := true
	for i, li := range lhs {
		places[i] = c.Place(li)
		canreorder = canreorder && places[i].IsVar() // ach, needed. see for example i := 0; i, x[i] = 1, 2  // set i = 1, x[0] = 2
	}
	if rn == 1 && ln > 1 {
		exprs[0] = c.Expr(rhs[0])
		canreorder = false
	} else {
		for i, ri := range rhs {
			exprs[i] = c.Expr1(ri)
			canreorder = canreorder && exprs[i].Const()
		}
	}
	if ln == rn && (ln <= 1 || canreorder) {
		for i := range lhs {
			c.assign1(lhs[i], node.Tok, rhs[i], places[i], exprs[i])
		}
		return
	}
	// problem: we need to create temporary copies of the evaluations
	// before performing the assignments. Such temporary copies must be per-goroutine!
	//
	// so a technique like the following is bugged,
	// because it creates a *single* global location for the temporary copy:
	//   var tmp r.Value
	//   func set(env *Env) { tmp = places[i].Fun(env) }
	//   func get(env *Env) r.Value { return tmp }

	assign := make([]Assign, ln)
	for i, place := range places {
		assign[i].init(c, place)
	}

	exprfuns, exprxv := c.assignPrepareRhs(node, places, exprs)

	c.Pos = node.Pos()
	if ln == 2 && rn == 2 && assign[0].placekey == nil && assign[1].placekey == nil {
		c.assign2(assign, exprfuns)
	} else {
		c.assignMulti(assign, exprfuns, exprxv)
	}
}

func (c *Comp) assignPrepareRhs(node *ast.AssignStmt, places []*Place, exprs []*Expr) ([]func(*Env) r.Value, func(*Env) (r.Value, []r.Value)) {
	lhs, rhs := node.Lhs, node.Rhs
	ln, rn := len(lhs), len(rhs)
	if ln == rn {
		exprfuns := make([]func(*Env) r.Value, rn)
		for i, expr := range exprs {
			tplace := places[i].Type
			if expr.Const() {
				expr.ConstTo(tplace)
			} else if !expr.Type.AssignableTo(tplace) {
				c.Pos = rhs[i].Pos()
				c.Errorf("cannot use <%v> as <%v> in assignment: %v %v %v", expr.Type, tplace, lhs[i], node.Tok, rhs[i])
			}
			exprfuns[i] = expr.AsX1()
		}
		return exprfuns, nil
	}
	if rn == 1 {
		expr := exprs[0]
		nexpr := expr.NumOut()
		if nexpr != ln {
			c.Pos = node.Pos()
			c.Errorf("invalid assignment: expression returns %d values, cannot assign them to %d places: %v", nexpr, ln, node)
		}
		for i := 0; i < nexpr; i++ {
			texpr := expr.Out(i)
			tplace := places[i].Type
			if !texpr.AssignableTo(tplace) {
				c.Pos = lhs[i].Pos()
				c.Errorf("cannot assign <%v> to %v <%v> in multiple assignment", texpr, lhs[i], tplace)
			}
		}
		return nil, expr.AsXV(CompileDefaults)
	}
	c.Pos = node.Pos()
	c.Errorf("invalid assignment, cannot assign %d values to %d places: %v", rn, ln, node)
	return nil, nil
}

// assign2 compiles multiple assignment to two places
func (c *Comp) assign2(assign []Assign, exprfuns []func(*Env) r.Value) {
	efuns := [2]func(*Env) r.Value{exprfuns[0], exprfuns[1]}
	var stmt Stmt
	if assign[0].placefun == nil {
		if assign[1].placefun == nil {
			setvars := [2]func(*Env, r.Value){assign[0].setvar, assign[1].setvar}
			stmt = func(env *Env) (Stmt, *Env) {
				val0 := efuns[0](env)
				val1 := efuns[1](env)
				setvars[0](env, val0)
				setvars[1](env, val1)
				env.IP++
				return env.Code[env.IP], env
			}
		} else {
			stmt = func(env *Env) (Stmt, *Env) {
				obj1 := assign[1].placefun(env)
				val0 := efuns[0](env)
				val1 := efuns[1](env)
				assign[0].setvar(env, val0)
				assign[1].setplace(obj1, obj1, val1)
				env.IP++
				return env.Code[env.IP], env
			}
		}
	} else {
		if assign[1].placefun == nil {
			stmt = func(env *Env) (Stmt, *Env) {
				obj0 := assign[0].placefun(env)
				val0 := efuns[0](env)
				val1 := efuns[1](env)
				assign[0].setplace(obj0, obj0, val0)
				assign[1].setvar(env, val1)
				env.IP++
				return env.Code[env.IP], env
			}
		} else {
			stmt = func(env *Env) (Stmt, *Env) {
				obj0 := assign[0].placefun(env)
				obj1 := assign[1].placefun(env)
				val0 := efuns[0](env)
				val1 := efuns[1](env)
				assign[0].setplace(obj0, obj0, val0)
				assign[1].setplace(obj1, obj1, val1)
				env.IP++
				return env.Code[env.IP], env
			}
		}
	}
	c.append(stmt)
}

// assignMulti compiles multiple assignment to places
func (c *Comp) assignMulti(assign []Assign, exprfuns []func(*Env) r.Value, exprxv func(*Env) (r.Value, []r.Value)) {
	stmt := func(env *Env) (Stmt, *Env) {
		n := len(assign)
		// these buffers must be allocated at runtime, per goroutine!
		objs := make([]r.Value, n)
		keys := make([]r.Value, n)
		var tmp r.Value
		var a *Assign
		// evaluate all lhs
		for i := range assign {
			if a = &assign[i]; a.placefun == nil {
				continue
			}
			objs[i] = a.placefun(env)
			if a.placekey == nil {
				continue
			}
			// assigning to obj[key] where obj is a map:
			// obj and key do NOT need to be settable,
			// and actually Go spec tell to make a copy of their values
			if tmp = objs[i]; tmp.CanSet() {
				objs[i] = tmp.Convert(tmp.Type())
			}
			if tmp = a.placekey(env); tmp.CanSet() {
				tmp = tmp.Convert(tmp.Type())
			}
			keys[i] = tmp
		}
		// evaluate all rhs
		var vals []r.Value
		if exprxv != nil {
			_, vals = exprxv(env)
		} else {
			vals = make([]r.Value, n)
			for i, exprfun := range exprfuns {
				vals[i] = exprfun(env)
			}
		}
		// execute assignments
		for i := range assign {
			a := &assign[i]
			if a.setvar != nil {
				a.setvar(env, vals[i])
			} else {
				a.setplace(objs[i], keys[i], vals[i])
			}
		}
		env.IP++
		return env.Code[env.IP], env
	}
	c.append(stmt)
}

// assign1 compiles a single assignment to a place
func (c *Comp) assign1(lhs ast.Expr, op token.Token, rhs ast.Expr, place *Place, init *Expr) {
	panicking := true
	defer func() {
		if !panicking {
			return
		}
		rec := recover()
		node := &ast.AssignStmt{Lhs: []ast.Expr{lhs}, Tok: op, Rhs: []ast.Expr{rhs}} // for nice error messages
		c.Errorf("error compiling assignment: %v\n\t%v", node, rec)
	}()
	if place.IsVar() {
		c.SetVar(&place.Var, op, init)
	} else {
		c.SetPlace(place, op, init)
	}
	panicking = false
}

// LookupVar compiles the left-hand-side of an assignment, in case it's an identifier (i.e. a variable name)
func (c *Comp) LookupVar(name string) *Var {
	if name == "_" {
		return &Var{}
	}
	sym := c.Resolve(name)
	return sym.AsVar(PlaceSettable)
}

// Place compiles the left-hand-side of an assignment
func (c *Comp) Place(node ast.Expr) *Place {
	return c.placeOrAddress(node, false)
}

// PlaceOrAddress compiles the left-hand-side of an assignment or the location of an address-of
func (c *Comp) placeOrAddress(in ast.Expr, opt PlaceOption) *Place {
	for {
		if in != nil {
			c.Pos = in.Pos()
		}
		switch node := in.(type) {
		case *ast.CompositeLit:
			// composite literals are addressable but not settable
			if opt == PlaceSettable {
				c.Errorf("%s composite literal", opt)
			}
			e := c.Expr1(node)
			fun := e.AsX1()
			var addr func(*Env) r.Value
			switch e.Type.Kind() {
			case r.Array, r.Struct:
				// array and struct composite literals are directly addressable
				// because they are created with reflect.New(t).Elem()
				addr = func(env *Env) r.Value {
					return fun(env).Addr()
				}
			default:
				// other composite literals (maps, slices) are not directly addressable:
				// the result of reflect.MakeMap and reflect.MakeSlice is not addressable,
				// so implement a workaround to behave as compiled Go.
				//
				// 'addr' below creates a new pointer-to-t at each execution,
				// but since the map or slice is freshly created each time
				// and 'addr' below is the only one code accessing it,
				// it's not a problem
				addr = func(env *Env) r.Value {
					obj := fun(env)
					place := r.New(obj.Type())
					place.Elem().Set(obj)
					return place
				}
			}
			return &Place{Var: Var{Type: e.Type}, Fun: fun, Addr: addr}
		case *ast.Ident:
			return c.IdentPlace(node.Name, opt)
		case *ast.IndexExpr:
			return c.IndexPlace(node, opt)
		case *ast.ParenExpr:
			in = node.X
			continue
		case *ast.StarExpr:
			e := c.Expr1(node.X)
			if e.Const() {
				c.Errorf("%s a constant: %v <%v>", opt, node, e.Type)
				return nil
			}
			// we cannot optimize the case "node.X is a variable" because we are compiling *variable, not variable
			// e.Fun is already the address we want, dereference its type
			t := e.Type.Elem()
			// c.Debugf("placeOrAddress: %v has type %v, transformed into: %v has type %v", node.X, e.Type, node, t)
			addr := e.AsX1()
			fun := func(env *Env) r.Value {
				return addr(env).Elem()
			}
			return &Place{Var: Var{Type: t}, Fun: fun, Addr: addr}
		case *ast.SelectorExpr:
			return c.SelectorPlace(node, opt)
		default:
			c.Errorf("%s: %v", opt, in)
			return nil
		}
	}
}

// placeForSideEffects compiles the left-hand-side of a do-nothing assignment,
// as for example *addressOfInt() += 0, in order to apply its side effects
func (c *Comp) placeForSideEffects(place *Place) {
	if place.IsVar() {
		return
	}
	var ret Stmt
	fun := place.Fun
	if mapkey := place.MapKey; mapkey != nil {
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			mapkey(env)
			// no need to call obj.MapIndex(key): it has no side effects and cannot panic.
			// obj := fun(env)
			// key := mapkey(env)
			// obj.MapIndex(key)
			env.IP++
			return env.Code[env.IP], env
		}
	} else {
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	}
	c.append(ret)
}
