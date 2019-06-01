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
 * unary.go
 *
 *  Created on Apr 07, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/constant"
	"go/token"

	"github.com/cosmos72/gomacro/base"
	etoken "github.com/cosmos72/gomacro/go/etoken"
)

func (c *Comp) UnaryExpr(node *ast.UnaryExpr) *Expr {
	switch node.Op {
	case etoken.QUOTE:
		// surprisingly easy :)
		block := node.X.(*ast.FuncLit).Body
		node := base.SimplifyNodeForQuote(block, true)
		return c.exprValue(nil, node)

	case etoken.QUASIQUOTE:
		return c.quasiquoteUnary(node)

	case etoken.UNQUOTE, etoken.UNQUOTE_SPLICE:
		c.Errorf("invalid %s outside %s: %v", etoken.String(node.Op), etoken.String(etoken.QUASIQUOTE), node)

	case token.AND:
		// c.Expr(node.X) is useless here... skip it
		return c.AddressOf(node)
	}

	xe := c.expr1(node.X, nil)
	if xe.Type == nil {
		return c.invalidUnaryExpr(node, xe)
	}
	if xe.Untyped() {
		return c.UnaryExprUntyped(node, xe)
	}
	isConst := xe.Const()
	xe.WithFun()
	var z *Expr

	switch node.Op {
	case token.ADD:
		z = c.UnaryPlus(node, xe) // only checks xe type, returns xe itself
	case token.SUB:
		z = c.UnaryMinus(node, xe)
	case token.NOT:
		z = c.UnaryNot(node, xe)
	case token.XOR:
		z = c.UnaryXor(node, xe)
	case token.ARROW:
		z = c.Recv(node, xe)
		// never returns a constant
		isConst = false
	// case token.MUL: // not seen, the parser produces *ast.StarExpr instead
	default:
		return c.invalidUnaryExpr(node, xe)
	}
	if isConst {
		// constant propagation
		z.EvalConst(COptKeepUntyped)
	} else {
		// create jit expression for z
		c.Jit.UnaryExpr(z, node.Op, xe)
	}
	return z
}

func (c *Comp) UnaryExprUntyped(node *ast.UnaryExpr, xe *Expr) *Expr {
	op := node.Op
	switch op {
	case token.ADD, token.SUB, token.XOR, token.NOT:
		xlit := xe.Value.(UntypedLit)
		ret := constant.UnaryOp(op, xlit.Val, 0)
		if ret == constant.MakeUnknown() {
			return c.invalidUnaryExpr(node, xe)
		}
		return c.exprUntypedLit(xlit.Kind, ret)
	}
	return c.invalidUnaryExpr(node, xe)
}

func (c *Comp) invalidUnaryExpr(node *ast.UnaryExpr, xe *Expr) *Expr {
	return c.badUnaryExpr("invalid", node, xe)
}

func (c *Comp) unimplementedUnaryExpr(node *ast.UnaryExpr, xe *Expr) *Expr {
	return c.badUnaryExpr("unimplemented", node, xe)
}

func (c *Comp) badUnaryExpr(reason string, node *ast.UnaryExpr, xe *Expr) *Expr {
	opstr := etoken.String(node.Op)
	if xe != nil {
		c.Errorf("%s unary operation %s on <%v>: %s %v",
			reason, opstr, xe.Type, opstr, node.X)
	} else {
		c.Errorf("%s unary operation %s: %s %v",
			reason, opstr, opstr, node.X)
	}
	return nil
}
