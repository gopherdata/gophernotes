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
 * range.go
 *
 *  Created on Jun 04, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"
	"sort"
	"unicode/utf8"
	"unsafe"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type rangeJump struct {
	Start, Continue, Break int
}

// Range compiles a "for-range" statement
func (c *Comp) Range(node *ast.RangeStmt, labels []string) {
	var nbinds [2]int
	flag := true // node.Tok == token.DEFINE || (node.Body != nil && containLocalBinds(node.Body.List...))

	c, _ = c.pushEnvIfFlag(&nbinds, flag)
	erange := c.Expr1(node.X, nil)
	t := erange.Type
	if erange.Untyped() {
		t = erange.DefaultType()
		erange.ConstTo(t)
	}
	var jump rangeJump

	sort.Strings(labels)
	// we need a fresh Comp here... created above by c.pushEnvIfFlag()
	c.Loop = &LoopInfo{
		Continue:   &jump.Continue,
		Break:      &jump.Break,
		ThisLabels: labels,
	}

	switch t.Kind() {
	case r.Ptr:
		if t.Elem().Kind() != r.Array {
			c.Errorf("cannot range over %v <%v>", node.X, t)
		}
		// range on pointer to array: dereference it
		t = t.Elem()
		efun := erange.AsX1()
		erange = exprX1(t, func(env *Env) r.Value {
			return efun(env).Elem()
		})
		fallthrough
	case r.Array, r.Slice:
		c.rangeSlice(node, erange, &jump)
	case r.Chan:
		c.rangeChan(node, erange, &jump)
	case r.Map:
		c.rangeMap(node, erange, &jump)
	case r.String:
		c.rangeString(node, erange, &jump)
	default:
		c.Errorf("cannot range over %v <%v>", node.X, t)
	}

	jump.Break = c.Code.Len()

	c = c.popEnvIfFlag(&nbinds, flag)
}

func (c *Comp) rangeChan(node *ast.RangeStmt, erange *Expr, jump *rangeJump) {
	t := erange.Type
	telem := t.Elem()

	// unnamed bind, contains channel
	bindchan := c.DeclVar0("", nil, erange)
	idxchan := bindchan.Desc.Index()

	placekey, _ := c.rangeVars(node, telem, nil)

	jump.Start = c.Code.Len()

	if placekey == nil {
		c.append(func(env *Env) (Stmt, *Env) {
			_, ok := env.Vals[idxchan].Recv()
			var ip int
			if ok {
				ip = env.IP + 1
			} else {
				ip = jump.Break
			}
			env.IP = ip
			return env.Code[ip], env
		})
	} else {
		// unnamed bind, contains last received value
		bindrecv := c.NewBind("", VarBind, c.TypeOfInterface())
		idxrecv := bindrecv.Desc.Index()

		c.append(func(env *Env) (Stmt, *Env) {
			v, ok := env.Vals[idxchan].Recv()
			var ip int
			if ok {
				env.Vals[idxrecv] = v
				ip = env.IP + 1
			} else {
				ip = jump.Break
			}
			env.IP = ip
			return env.Code[ip], env
		})
		c.SetPlace(placekey, token.ASSIGN, unwrapBind(bindrecv, telem))
	}

	// compile the body
	c.Block(node.Body)

	// "continue" is a jump to loop beginning
	jump.Continue = jump.Start

	// jump back to start
	c.append(func(env *Env) (Stmt, *Env) {
		ip := jump.Start
		env.IP = ip
		return env.Code[ip], env
	})
}

func (c *Comp) rangeMap(node *ast.RangeStmt, erange *Expr, jump *rangeJump) {
	t := erange.Type
	tkey, tval := t.Key(), t.Elem()
	tkeyslice := c.Universe.SliceOf(tkey)
	rtkeyslice := tkeyslice.ReflectType()

	// unnamed bind, contains map
	bindmap := c.DeclVar0("", nil, erange)
	idxmap := bindmap.Desc.Index()

	// unnamed bind, contains map keys
	bindkeys := c.NewBind("", VarBind, tkeyslice)
	idxkeys := bindkeys.Desc.Index()
	c.append(func(env *Env) (Stmt, *Env) {
		// convert []r.Value slice into a []rtkey slice, to avoid reflect.Value.Interface() while iterating
		vkeys := env.Vals[idxmap].MapKeys()
		keys := r.MakeSlice(rtkeyslice, len(vkeys), len(vkeys))
		for i, vkey := range vkeys {
			keys.Index(i).Set(vkey)
		}
		env.Vals[idxkeys] = keys
		env.IP++
		return env.Code[env.IP], env
	})

	// unnamed bind, contains iteration index
	bindnext := c.DeclVar0("", c.TypeOfInt(), nil)
	idxnext := bindnext.Desc.Index()

	placekey, placeval := c.rangeVars(node, tkey, tval)

	var bindkey *Bind
	var ekey *Expr
	if placekey != nil || placeval != nil {
		// unnamed bind, contains iteration map key
		bindkey = c.DeclVar0("", c.TypeOfInterface(), nil)
		ekey = unwrapBind(bindkey, tkey)
	}

	jump.Start = c.Code.Len()

	if bindkey == nil {
		// check iteration index against # of keys
		c.append(func(env *Env) (Stmt, *Env) {
			n := env.Vals[idxkeys].Len()
			i := *(*int)(unsafe.Pointer(&env.Ints[idxnext]))
			var ip int
			if i < n {
				ip = env.IP + 1
			} else {
				ip = jump.Break
			}
			env.IP = ip
			return env.Code[ip], env
		})
	} else {
		// check iteration index against # of keys,
		// and copy current map key into bindkey
		idxkey := bindkey.Desc.Index()
		c.append(func(env *Env) (Stmt, *Env) {
			vkeys := env.Vals[idxkeys]
			n := vkeys.Len()
			i := *(*int)(unsafe.Pointer(&env.Ints[idxnext]))
			var ip int
			if i < n {
				env.Vals[idxkey] = vkeys.Index(i)
				ip = env.IP + 1
			} else {
				ip = jump.Break
			}
			env.IP = ip
			return env.Code[ip], env
		})
	}

	if placekey != nil {
		// copy current map key into placekey
		c.SetPlace(placekey, token.ASSIGN, ekey)
	}

	if placeval == nil {
		// nothing to do
	} else if placeval.IsVar() && !base.IsOptimizedKind(placeval.Type.Kind()) {
		idxkey := bindkey.Desc.Index()
		idxval := placeval.Var.Desc.Index()
		upval := placeval.Var.Upn
		rtype := tval.ReflectType()
		zero := r.Zero(rtype)
		c.append(func(env *Env) (Stmt, *Env) {
			vmap := env.Vals[idxmap]
			key := env.Vals[idxkey]
			o := env
			for j := 0; j < upval; j++ {
				o = o.Outer
			}
			val := vmap.MapIndex(key)
			if !val.IsValid() {
				val = zero
			} else if val.Type() != rtype {
				val = convert(val, rtype)
			}
			o.Vals[idxval].Set(val)
			env.IP++
			return env.Code[env.IP], env
		})
	} else {
		emap := c.Bind(bindmap)
		c.SetPlace(placeval, token.ASSIGN, c.mapIndex1(nil, emap, ekey))
	}

	// compile the body
	c.Block(node.Body)

	// "continue" is a jump to the last statement below
	jump.Continue = c.Code.Len()

	// increase iteration index and jump back to start
	c.append(func(env *Env) (Stmt, *Env) {
		(*(*int)(unsafe.Pointer(&env.Ints[idxnext])))++
		ip := jump.Start
		env.IP = ip
		return env.Code[ip], env
	})
}

func (c *Comp) rangeSlice(node *ast.RangeStmt, erange *Expr, jump *rangeJump) {
	t := erange.Type
	var constlen int
	var elen *Expr

	if node.Value != nil || t.Kind() != r.Array {
		// Go spec: one-variable range on array ONLY evaluates the array length, not the array itself
		// save range variable in an unnamed bind
		bind := c.DeclVar0("", nil, erange)
		erange = c.Bind(bind)
	}

	if t.Kind() == r.Array {
		constlen = t.Len()
	} else {
		// save range length in an unnamed bind
		rangefun := erange.AsX1()
		elen0 := exprFun(c.TypeOfInt(), func(env *Env) int {
			return rangefun(env).Len()
		})
		bindlen := c.DeclVar0("", nil, elen0)
		elen = c.Bind(bindlen)
	}

	placekey, placeval := c.rangeVars(node, c.TypeOfInt(), t.Elem())

	if placekey == nil {
		// we need an interation variable, even if user code ignores it
		placekey = c.DeclVar0("", c.TypeOfInt(), nil).AsVar(0, PlaceSettable).AsPlace()
	}
	if placekey.Desc.Class() != IntBind {
		c.Errorf("internal error: for-range counter variable allocated with class = %v, expecting class = %v",
			placekey.Desc.Class(), IntBind)
	}

	jump.Start = c.Code.Len()

	// compile comparison against range length
	ekey := c.GetPlace(placekey)
	funkey := ekey.WithFun().(func(*Env) int)

	if t.Kind() == r.Array {
		c.append(func(env *Env) (Stmt, *Env) {
			var ip int
			if funkey(env) < constlen {
				ip = env.IP + 1
			} else {
				ip = jump.Break
			}
			env.IP = ip
			return env.Code[ip], env
		})
	} else {
		funlen := elen.WithFun().(func(*Env) int)
		c.append(func(env *Env) (Stmt, *Env) {
			var ip int
			if funkey(env) < funlen(env) {
				ip = env.IP + 1
			} else {
				ip = jump.Break
			}
			env.IP = ip
			return env.Code[ip], env
		})
	}
	if placeval != nil {
		// for error messages
		indexnode := &ast.IndexExpr{X: node.X, Lbrack: node.X.Pos(), Index: node.Key, Rbrack: node.X.Pos()}
		eindex := c.vectorIndex(indexnode, erange, ekey)
		c.SetPlace(placeval, token.ASSIGN, eindex)
	}

	// compile the body
	c.Block(node.Body)

	// "continue" is a jump to the increment below
	jump.Continue = c.Code.Len()

	// increment key
	c.Pos = node.End() - 1
	one := c.exprValue(c.TypeOfInt(), 1)
	c.SetPlace(placekey, token.ADD_ASSIGN, one)

	// jump back to comparison
	c.append(func(env *Env) (Stmt, *Env) {
		ip := jump.Start
		env.IP = ip
		return env.Code[ip], env
	})
}

func (c *Comp) rangeString(node *ast.RangeStmt, erange *Expr, jump *rangeJump) {
	// save string in an unnamed bind
	bindrange := c.DeclVar0("", nil, erange)
	idxrange := bindrange.Desc.Index()

	placekey, placeval := c.rangeVars(node, c.TypeOfInt(), c.TypeOfInt32())
	bindnext := c.DeclVar0("", c.TypeOfInt(), nil)
	idxnext := bindnext.Desc.Index()

	var bindrune *Bind
	if placeval != nil && !placeval.IsVar() {
		bindrune = c.DeclVar0("", c.TypeOfInt32(), nil)
	}

	jump.Start = c.Code.Len()

	if placekey != nil {
		c.SetPlace(placekey, token.ASSIGN, c.Bind(bindnext))
	}
	if placeval == nil {
		c.append(func(env *Env) (Stmt, *Env) {
			s := env.Vals[idxrange].String()
			pnext := (*int)(unsafe.Pointer(&env.Ints[idxnext]))
			next := *pnext

			_, size := utf8.DecodeRuneInString(s[next:])
			var ip int
			if size != 0 {
				next += size
				*pnext = next
				ip = env.IP + 1
			} else {
				ip = jump.Break
			}
			env.IP = ip
			return env.Code[ip], env
		})
	} else if placeval.IsVar() {
		idxval := placeval.Var.Desc.Index()
		upval := placeval.Var.Upn
		c.append(func(env *Env) (Stmt, *Env) {
			s := env.Vals[idxrange].String()
			pnext := (*int)(unsafe.Pointer(&env.Ints[idxnext]))
			next := *pnext

			r, size := utf8.DecodeRuneInString(s[next:])
			var ip int
			if size != 0 {
				next += size
				*pnext = next
				o := env
				for i := 0; i < upval; i++ {
					o = o.Outer
				}
				*(*int32)(unsafe.Pointer(&env.Ints[idxval])) = r
				ip = env.IP + 1
			} else {
				ip = jump.Break
			}
			env.IP = ip
			return env.Code[ip], env
		})
	} else {
		idxrune := bindrune.Desc.Index()
		c.append(func(env *Env) (Stmt, *Env) {
			s := env.Vals[idxrange].String()
			pnext := (*int)(unsafe.Pointer(&env.Ints[idxnext]))
			next := *pnext

			r, size := utf8.DecodeRuneInString(s[next:])
			var ip int
			if size != 0 {
				next += size
				*pnext = next
				*(*int32)(unsafe.Pointer(&env.Ints[idxrune])) = r
				ip = env.IP + 1
			} else {
				ip = jump.Break
			}
			env.IP = ip
			return env.Code[ip], env
		})
		c.SetPlace(placeval, token.ASSIGN, c.Bind(bindrune))
	}

	// compile the body
	c.Block(node.Body)

	// "continue" is a jump to the iteration above
	jump.Continue = jump.Start

	// jump back to iteration
	c.append(func(env *Env) (Stmt, *Env) {
		ip := jump.Start
		env.IP = ip
		return env.Code[ip], env
	})
}

// rangeVars compiles the key and value iteration variables in a for-range
func (c *Comp) rangeVars(node *ast.RangeStmt, tkey xr.Type, tval xr.Type) (*Place, *Place) {
	place := [2]*Place{nil, nil}
	t := [2]xr.Type{tkey, tval}

	for i, expr := range [2]ast.Expr{node.Key, node.Value} {
		if expr == nil {
			continue
		} else if t[i] == nil {
			c.Pos = expr.Pos()
			c.Errorf("too many variables in range")
		}
		c.Pos = expr.Pos()
		if node.Tok == token.DEFINE {
			switch expr := expr.(type) {
			case *ast.Ident:
				name := expr.Name
				if name != "_" {
					place[i] = c.DeclVar0(name, t[i], nil).AsVar(0, PlaceSettable).AsPlace()
				}
			default:
				c.Errorf("non-name %v on left side of :=", expr)
			}
		} else {
			if ident, ok := expr.(*ast.Ident); ok && ident.Name == "_" {
				// ignore range variable "_"
				continue
			}
			place[i] = c.Place(expr)
			if !t[i].AssignableTo(place[i].Type) {
				c.Errorf("cannot assign type <%v> to %v <%v> in range", t[i], expr, place[i].Type)
			}
		}
	}
	return place[0], place[1]
}
