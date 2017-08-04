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
 * select.go
 *
 *  Created on Jun 05, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"
	"sort"
)

type selectEntry struct {
	Dir  r.SelectDir
	Chan func(*Env) r.Value
	Send func(*Env) r.Value
}

func (c *Comp) Select(node *ast.SelectStmt, labels []string) {
	if node.Body == nil || len(node.Body.List) == 0 {
		return
	}

	var ibreak int
	sort.Strings(labels)
	c.Loop = &LoopInfo{
		Break:      &ibreak,
		ThisLabels: labels,
	}

	// unnamed bind, contains received value. Nil means nothing received
	bindrecv := c.AddBind("", VarBind, c.TypeOfInterface())
	idxrecv := bindrecv.Desc.Index()

	list := node.Body.List
	n := len(list)
	entries := make([]selectEntry, n)
	ips := make([]int, n)
	defaultip := -1
	defaultpos := token.NoPos

	c.append(func(env *Env) (Stmt, *Env) {
		cases := make([]r.SelectCase, len(entries))
		for i := range entries {
			c := &cases[i]
			e := &entries[i]
			c.Dir = e.Dir
			if e.Chan != nil {
				c.Chan = e.Chan(env)
				if e.Send != nil {
					c.Send = e.Send(env)
				}
			}
		}
		chosen, recv, _ := r.Select(cases)
		env.Binds[idxrecv] = recv
		ip := ips[chosen]
		env.IP = ip
		return env.Code[ip], env
	})

	for i, stmt := range list {
		ips[i] = c.Code.Len()
		switch clause := stmt.(type) {
		case *ast.CommClause:
			if clause.Comm == nil {
				if defaultip >= 0 {
					c.Errorf("multiple defaults in select (first at %s)", c.Fileset.Position(defaultpos))
				}
				defaultip = c.Code.Len()
				defaultpos = clause.Pos()
				entries[i] = c.selectDefault(clause)
			} else {
				entries[i] = c.selectCase(clause, bindrecv)
			}
		default:
			c.Errorf("invalid statement inside select: expecting case or default, found: %v <%v>", stmt, r.TypeOf(stmt))
		}
	}
	// we finally know this
	ibreak = c.Code.Len()
}

// selectDefault compiles the default case in a switch
func (c *Comp) selectDefault(node *ast.CommClause) selectEntry {
	if len(node.Body) != 0 {
		c.List(node.Body)
	}
	c.jumpOut(0, c.Loop.Break)
	return selectEntry{Dir: r.SelectDefault}
}

// selectCase compiles a case in a select.
func (c *Comp) selectCase(clause *ast.CommClause, bind *Bind) selectEntry {

	var entry selectEntry
	var nbinds [2]int
	stmt := clause.Comm
	c2 := c
	locals := false

	switch node := stmt.(type) {
	case *ast.ExprStmt:
		// <-ch
		entry = selectEntry{
			Dir:  r.SelectRecv,
			Chan: c.selectRecv(stmt, node.X).AsX1(),
		}
	case *ast.AssignStmt:
		// v := <-ch or v = <-ch
		lhs := node.Lhs
		n := len(lhs)
		if (n != 1 && n != 2) || len(node.Rhs) != 1 {
			c.badSelectCase(stmt)
		}
		var l0, l1 ast.Expr = lhs[0], nil
		if n == 2 {
			l1 = lhs[1]
		}
		r0 := node.Rhs[0]
		switch node.Tok {
		case token.DEFINE:
			id0 := asIdent(l0)
			id1 := asIdent(l1)
			if (id0 == nil && l0 != nil) || (id1 == nil && l1 != nil) {
				c.badSelectCase(stmt)
			}
			echan := c.selectRecv(node, r0)
			entry = selectEntry{Dir: r.SelectRecv, Chan: echan.AsX1()}

			if id0 != nil && id0.Name != "_" || id1 != nil && id1.Name != "_" {
				c2, locals = c.pushEnvIfFlag(&nbinds, true)

				if id0 != nil && id0.Name != "_" {
					t := echan.Type.Elem()
					c2.DeclVar0(id0.Name, t, unwrapBindUp1(bind, t))
				}
				if id1 != nil && id1.Name != "_" {
					idx := bind.Desc.Index()
					c2.DeclVar0(id1.Name, c.TypeOfBool(), c.exprBool(func(env *Env) bool {
						return env.Outer.Binds[idx].IsValid()
					}))
				}
			} else if len(clause.Body) != 0 {
				c2, locals = c.pushEnvIfLocalBinds(&nbinds, clause.Body...)
			}

		case token.ASSIGN:
			echan := c.selectRecv(stmt, r0)
			entry = selectEntry{Dir: r.SelectRecv, Chan: echan.AsX1()}

			if l0 != nil {
				place := c.Place(l0)
				t := echan.Type.Elem()
				tplace := place.Type
				if !t.AssignableTo(tplace) {
					c.Errorf("cannot use <%v> as <%v> in assignment: %v = %v", t, tplace, l0, r0)
				}
				c.SetPlace(place, token.ASSIGN, unwrapBind(bind, t))
			}
			if l1 != nil {
				place := c.Place(l1)
				t := c.TypeOfBool()
				tplace := place.Type
				if !t.AssignableTo(tplace) {
					c.Errorf("cannot use <%v> as <%v> in assignment: _, %v = %v", t, tplace, l1, r0)
				}
				idx := bind.Desc.Index()
				c.SetPlace(place, token.ASSIGN, c.exprBool(func(env *Env) bool {
					return env.Binds[idx].IsValid()
				}))
			}

			if len(clause.Body) != 0 {
				c2, locals = c.pushEnvIfLocalBinds(&nbinds, clause.Body...)
			}
		}

	case *ast.SendStmt:
		// ch <- v
		echan := c.Expr1(node.Chan)
		if echan.Type.Kind() != r.Chan {
			c.Errorf("cannot use %v <%v> as channel in select case", node, echan.Type)
		}
		esend := c.Expr1(node.Value)
		tactual := esend.Type
		texpected := echan.Type.Elem()
		if !tactual.AssignableTo(texpected) {
			c.Errorf("cannot use %v <%v> as <%v> in channel send", node.Value, tactual, texpected)
		}
		entry = selectEntry{Dir: r.SelectSend, Chan: echan.AsX1(), Send: esend.AsX1()}

	default:
		c.badSelectCase(stmt)
	}

	if len(clause.Body) != 0 {
		c2.List(clause.Body)
	}
	if c2 != c {
		c2.popEnvIfFlag(&nbinds, locals)
	}
	c.jumpOut(0, c.Loop.Break)
	return entry
}

func (c *Comp) selectRecv(stmt ast.Stmt, node ast.Expr) *Expr {
	for {
		switch expr := node.(type) {
		case *ast.ParenExpr:
			node = expr.X
			continue
		case *ast.UnaryExpr:
			if expr.Op == token.ARROW {
				e := c.Expr1(expr.X)
				if e.Type.Kind() != r.Chan {
					c.Errorf("cannot use %v <%v> as channel in select case", node, e.Type)
				}
				return e
			}
		}
		c.badSelectCase(stmt)
		return nil
	}
}

func (c *Comp) badSelectCase(stmt ast.Stmt) {
	c.Errorf("invalid select case, expecting [ch <- val] or [<-ch] or [vars := <-ch] or [places = <-ch], found: %v <%v>",
		stmt, r.TypeOf(stmt))
}
