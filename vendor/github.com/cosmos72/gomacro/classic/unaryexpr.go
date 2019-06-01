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
 * unaryexpr.go
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
	etoken "github.com/cosmos72/gomacro/go/etoken"
)

func (env *Env) unsupportedUnaryExpr(xv r.Value, op token.Token) (r.Value, []r.Value) {
	opstr := etoken.String(op)
	return env.Errorf("unsupported unary expression %s on <%v>: %s %v", opstr, typeOf(xv), opstr, xv)
}

func (env *Env) warnOverflowSignedMinus(x interface{}, ret interface{}) {
	str := fmt.Sprintf("%d", x)
	if len(str) > 0 && str[0] == '-' {
		str = str[1:]
	}
	env.Warnf("value %s overflows <%v>, result truncated to %d", str, r.TypeOf(x), ret)
}

func (env *Env) warnUnderflowUnsignedMinus(x interface{}, ret interface{}) {
	env.Warnf("value -%d underflows <%v>, result truncated to %d", x, r.TypeOf(x), ret)
}

func (env *Env) evalUnaryExpr(node *ast.UnaryExpr) (r.Value, []r.Value) {
	op := node.Op
	switch op {
	case token.AND:
		place := env.evalExpr1(node.X)
		if place == Nil || !place.CanAddr() {
			return env.Errorf("cannot take the address of: %v = %v <%v>", node.X, place, typeOf(place))
		}
		return place.Addr(), nil

	// the various QUOTE special forms, the result of macroexpansion,
	// and our extension "block statement inside expression" are:
	// a block statements, wrapped in a closure, wrapped in a unary expression "MACRO", i.e.:
	// MACRO func() { /*block*/ }
	case etoken.MACRO:
		block := node.X.(*ast.FuncLit).Body
		return env.evalBlock(block)

	case etoken.QUOTE:
		block := node.X.(*ast.FuncLit).Body
		ret := env.evalQuote(block)
		return r.ValueOf(ret), nil

	case etoken.QUASIQUOTE:
		block := node.X.(*ast.FuncLit).Body
		ret := env.evalQuasiquote(block)
		return r.ValueOf(ret), nil

	case etoken.UNQUOTE, etoken.UNQUOTE_SPLICE:
		return env.Errorf("%s not inside quasiquote: %v <%v>", etoken.String(op), node, r.TypeOf(node))
	}

	xv, _ := env.EvalNode(node.X)

	if op == token.ADD {
		switch xv.Kind() {
		case r.Int, r.Int8, r.Int16, r.Int32, r.Int64,
			r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr,
			r.Float32, r.Float64, r.Complex64, r.Complex128:
			return xv, nil
		default:
			return env.unsupportedUnaryExpr(xv, op)
		}
	}
	var ret interface{}

	switch xv.Kind() {
	case r.Bool:
		if op == token.NOT {
			ret = !xv.Bool()
		}
	case r.Int:
		x := int(xv.Int())
		switch op {
		case token.SUB:
			ret = -x
			if x == -x && x != 0 {
				env.warnOverflowSignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case r.Int8:
		x := int8(xv.Int())
		switch op {
		case token.SUB:
			ret = -x
			if x == -x && x != 0 {
				env.warnOverflowSignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case r.Int16:
		x := int16(xv.Int())
		switch op {
		case token.SUB:
			ret = -x
			if x == -x && x != 0 {
				env.warnOverflowSignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case r.Int32:
		x := int32(xv.Int())
		switch op {
		case token.SUB:
			ret = -x
			if x == -x && x != 0 {
				env.warnOverflowSignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case r.Int64:
		x := xv.Int()
		switch op {
		case token.SUB:
			ret = -x
			if x == -x && x != 0 {
				env.warnOverflowSignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case r.Uint:
		x := uint(xv.Uint())
		switch op {
		case token.SUB:
			ret = -x
			if x != 0 {
				env.warnUnderflowUnsignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case r.Uint8:
		x := uint8(xv.Uint())
		switch op {
		case token.SUB:
			ret = -x
			if x != 0 {
				env.warnUnderflowUnsignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case r.Uint16:
		x := uint16(xv.Uint())
		switch op {
		case token.SUB:
			ret = -x
			if x != 0 {
				env.warnUnderflowUnsignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case r.Uint32:
		x := uint32(xv.Uint())
		switch op {
		case token.SUB:
			ret = -x
			if x != 0 {
				env.warnUnderflowUnsignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case r.Uint64:
		x := xv.Uint()
		switch op {
		case token.SUB:
			ret = -x
			if x != 0 {
				env.warnUnderflowUnsignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case r.Uintptr:
		x := uintptr(xv.Uint())
		switch op {
		case token.SUB:
			ret = -x
			if x != 0 {
				env.warnUnderflowUnsignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case r.Float32:
		x := float32(xv.Float())
		switch op {
		case token.SUB:
			ret = -x
		}
	case r.Float64:
		x := xv.Float()
		switch op {
		case token.SUB:
			ret = -x
		}
	case r.Complex64:
		x := complex64(xv.Complex())
		switch op {
		case token.SUB:
			ret = -x
		}
	case r.Complex128:
		x := xv.Complex()
		switch op {
		case token.SUB:
			ret = -x
		}
	case r.Chan:
		switch op {
		case token.ARROW:
			ret, ok := xv.Recv()
			return ret, []r.Value{ret, r.ValueOf(ok)}
		}
	}
	if ret == nil {
		return env.unsupportedUnaryExpr(xv, op)
	}
	retv := r.ValueOf(ret)
	xt := xv.Type()
	if retv.Type() != xt {
		retv = retv.Convert(xt)
	}
	return retv, nil
}
