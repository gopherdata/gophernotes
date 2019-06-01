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
		e := c.expr(nodes[0], nil)
		if actualN := e.NumOut(); actualN != expectedValuesN {
			var plural string
			if actualN != 1 {
				plural = "s"
			}
			c.Errorf("expression returns %d value%s, expecting %d: %v", actualN, plural, expectedValuesN, nodes[0])
		}
		inits = []*Expr{e}
	} else {
		inits = c.exprs(nodes)
	}
	return inits
}

// Exprs compiles multiple expressions
func (c *Comp) Exprs(nodes []ast.Expr) []*Expr {
	es := c.exprs(nodes)
	for _, e := range es {
		c.Jit.Fun(e)
	}
	return es
}

// same as Exprs, but does not replace e[i].Fun with jit-compiled code
func (c *Comp) exprs(nodes []ast.Expr) []*Expr {
	var es []*Expr
	if n := len(nodes); n != 0 {
		es = make([]*Expr, n)
		for i := range nodes {
			es[i] = c.expr1(nodes[i], nil)
		}
	}
	return es
}

// Expr1 compiles an expression that returns a single value
// t is optional and used for type inference on composite literals,
// see https://golang.org/ref/spec#Composite_literals
func (c *Comp) Expr1(in ast.Expr, t xr.Type) *Expr {
	return c.expr1(in, t)
}

// same as Expr1, but does not replace e.Fun with jit-compiled code
func (c *Comp) expr1(in ast.Expr, t xr.Type) *Expr {
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
	e := c.expr(in, t)
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
	e := c.expr(in, t)
	return c.Jit.Fun(e)
}

// same as Expr, but does not replace e.Fun with jit-compiled code
func (c *Comp) expr(in ast.Expr, t xr.Type) *Expr {
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
// performs simultaneous lookup for type names, constants, variables and functions
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
				bind, okb := o.Binds[name]
				var okt bool
				if okb && (GENERICS_V1_CXX || GENERICS_V2_CTI) {
					_, okt = bind.Value.(*GenericType) // generic types are stored in Comp.Bind[]
					okb = !okt
				}
				if okb {
					return c.Expr1(expr, nil), nil
				} else if _, ok := o.Types[name]; ok || okt {
					return nil, c.Type(expr)
				}
			}
		case *ast.IndexExpr:
			if GENERICS_V1_CXX || GENERICS_V2_CTI {
				if lit, ok := n.Index.(*ast.CompositeLit); ok && lit.Type == nil {
					// foo#[a, b...] can be a generic function or a generic type
					node = n.X
					continue
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

// IndexExpr compiles a read operation on obj[idx]
// or a generic function name#[T1, T2...]
func (c *Comp) IndexExpr(node *ast.IndexExpr) *Expr {
	if GENERICS_V1_CXX || GENERICS_V2_CTI {
		if e := c.GenericFunc(node); e != nil {
			return e
		}
	}
	return c.indexExpr(node, true)
}

// IndexExpr1 compiles a single-valued read operation on obj[idx]
// or a generic function name#[T1, T2...]
func (c *Comp) IndexExpr1(node *ast.IndexExpr) *Expr {
	if GENERICS_V1_CXX || GENERICS_V2_CTI {
		if e := c.GenericFunc(node); e != nil {
			return e
		}
	}
	return c.indexExpr(node, false)
}
