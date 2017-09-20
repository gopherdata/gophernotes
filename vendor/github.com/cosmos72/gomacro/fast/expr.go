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
 * expr.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"

	xr "github.com/cosmos72/gomacro/xreflect"
)

// ExprsMultipleValues either a single expression returning multiple values,
// or multiple expressions each returning a value.
func (c *Comp) ExprsMultipleValues(nodes []ast.Expr, expectedValuesN int) (inits []*Expr) {
	n := len(nodes)
	if n != expectedValuesN {
		if n != 1 {
			c.Errorf("value count mismatch: cannot assign %d values to %d places: %v",
				n, expectedValuesN, nodes)
			return nil
		}
		e := c.Expr(nodes[0])
		if actualN := e.NumOut(); actualN != expectedValuesN {
			var plural string
			if actualN != 1 {
				plural = "s"
			}
			c.Errorf("expression returns %d value%s, expecting %d: %v", actualN, plural, expectedValuesN, nodes[0])
		}
		inits = []*Expr{e}
	} else {
		inits = c.Exprs(nodes)
	}
	return inits
}

// Exprs compiles multiple expressions
func (c *Comp) Exprs(nodes []ast.Expr) []*Expr {
	var inits []*Expr
	if n := len(nodes); n != 0 {
		inits = make([]*Expr, n)
		for i := range nodes {
			inits[i] = c.Expr1(nodes[i])
		}
	}
	return inits
}

// Expr compiles an expression that returns a single value
func (c *Comp) Expr1(in ast.Expr) *Expr {
	for {
		if in != nil {
			c.Pos = in.Pos()
		}
		// env.Debugf("Expr1() %v", node)
		switch node := in.(type) {
		case *ast.ParenExpr:
			in = node.X
			continue
		case *ast.IndexExpr:
			return c.IndexExpr1(node)
		case *ast.TypeAssertExpr:
			return c.TypeAssert1(node)
		case *ast.UnaryExpr:
			if node.Op == token.ARROW {
				xe := c.Expr1(node.X)
				return c.Recv1(node, xe)
			} else {
				return c.UnaryExpr(node)
			}
		}
		break
	}
	e := c.Expr(in)
	nout := e.NumOut()
	switch nout {
	case 0:
		c.Errorf("expression returns no values, expecting one: %v", in)
		return nil
	default:
		c.Warnf("expression returns %d values %v, using only the first one: %v",
			len(e.Types), e.Types, in)
		fallthrough
	case 1:
		return e
	}
}

// Expr compiles an expression
func (c *Comp) Expr(in ast.Expr) *Expr {
	for {
		if in != nil {
			c.Pos = in.Pos()
		}
		// env.Debugf("Expr() %v", node)
		switch node := in.(type) {
		case *ast.BasicLit:
			return c.BasicLit(node)
		case *ast.BinaryExpr:
			return c.BinaryExpr(node)
		case *ast.CallExpr:
			return c.CallExpr(node)
		case *ast.CompositeLit:
			return c.CompositeLit(node)
		case *ast.FuncLit:
			return c.FuncLit(node)
		case *ast.Ident:
			return c.Ident(node.Name)
		case *ast.IndexExpr:
			return c.IndexExpr(node)
		case *ast.ParenExpr:
			in = node.X
			continue
		case *ast.UnaryExpr:
			return c.UnaryExpr(node)
		case *ast.SelectorExpr:
			return c.SelectorExpr(node)
		case *ast.SliceExpr:
			return c.SliceExpr(node)
		case *ast.StarExpr:
			return c.StarExpr(node)
		case *ast.TypeAssertExpr:
			return c.TypeAssert2(node)
		default:
		}
		c.Errorf("unimplemented Compile() for: %v <%v>", in, r.TypeOf(in))
		return nil
	}
}

// Expr1OrType compiles an single-valued expression or a type
// FIXME lookup simultaneously for both types and expressions
func (c *Comp) Expr1OrType(node ast.Expr) (e *Expr, t xr.Type) {
	panicking := true
	defer func() {
		if panicking {
			recover()
			t = c.Type(node)
		}
	}()
	e = c.Expr1(node)
	panicking = false
	return
}
