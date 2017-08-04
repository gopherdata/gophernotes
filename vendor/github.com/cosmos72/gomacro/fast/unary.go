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
	mt "github.com/cosmos72/gomacro/token"
)

func (c *Comp) UnaryExpr(node *ast.UnaryExpr) *Expr {
	switch node.Op {
	case mt.QUOTE:
		// surprisingly easy :)
		block := node.X.(*ast.FuncLit).Body
		node := base.SimplifyNodeForQuote(block, true)
		return c.exprValue(nil, node)

	case mt.QUASIQUOTE:
		return c.quasiquoteUnary(node)

	case mt.UNQUOTE, mt.UNQUOTE_SPLICE:
		c.Errorf("invalid %s outside %s: %v", mt.String(node.Op), mt.String(mt.QUASIQUOTE), node)

	case token.AND:
		// c.Expr(node.X) is useless here... skip it
		return c.AddressOf(node)
	}

	xe := c.Expr1(node.X)
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
		z.EvalConst(CompileKeepUntyped)
	}
	return z
}

func (c *Comp) UnaryExprUntyped(node *ast.UnaryExpr, xe *Expr) *Expr {
	op := node.Op
	switch op {
	case token.ADD, token.SUB, token.XOR, token.NOT:
		xlit := xe.Value.(UntypedLit)
		ret := constant.UnaryOp(op, xlit.Obj, 0)
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
	opstr := mt.String(node.Op)
	if xe != nil {
		c.Errorf("%s unary operation %s on <%v>: %s %v",
			reason, opstr, xe.Type, opstr, node.X)
	} else {
		c.Errorf("%s unary operation %s: %s %v",
			reason, opstr, opstr, node.X)
	}
	return nil
}
