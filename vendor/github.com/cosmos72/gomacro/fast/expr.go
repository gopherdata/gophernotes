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
		e := c.Expr(nodes[0], nil)
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
			inits[i] = c.Expr1(nodes[i], nil)
		}
	}
	return inits
}

// Expr compiles an expression that returns a single value
// t is optional and used for type inference on composite literals,
// see https://golang.org/ref/spec#Composite_literals
func (c *Comp) Expr1(in ast.Expr, t xr.Type) *Expr {
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
				xe := c.Expr1(node.X, nil)
				return c.Recv1(node, xe)
			} else {
				return c.UnaryExpr(node)
			}
		}
		break
	}
	e := c.Expr(in, t)
	nout := e.NumOut()
	switch nout {
	case 0:
		c.Errorf("expression returns no values, expecting one: %v", in)
		return nil
	case 1:
		return e
	default:
		return e.exprXVAsI()
	}
}

// Expr compiles an expression.
// t is optional and used for type inference on composite literals,
// see https://golang.org/ref/spec#Composite_literals
func (c *Comp) Expr(in ast.Expr, t xr.Type) *Expr {
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
			// propagate inferred type
			return c.CompositeLit(node, t)
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

// Expr1OrType compiles an single-valued expression or a type.
// looks up simultaneously for both types and expressions
func (c *Comp) Expr1OrType(expr ast.Expr) (e *Expr, t xr.Type) {
	node := expr
	for {
		switch n := node.(type) {
		case *ast.StarExpr:
			node = n.X
			continue
		case *ast.ParenExpr:
			node = n.X
			continue
		case *ast.Ident:
			name := n.Name
			for o := c; o != nil; o = o.Outer {
				if _, ok := o.Binds[name]; ok {
					return c.Expr1(expr, nil), nil
				} else if _, ok := o.Types[name]; ok {
					return nil, c.Type(expr)
				}
			}
		}
		break
	}
	panicking := true
	defer func() {
		if panicking {
			recover()
			t = c.Type(expr)
		}
	}()
	e = c.Expr1(expr, nil)
	panicking = false
	return
}
