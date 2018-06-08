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
 * switch.go
 *
 *  Created on May 06, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/constant"
	"go/token"
	r "reflect"
	"sort"
	"time"

	. "github.com/cosmos72/gomacro/base"
)

type caseEntry struct {
	Pos token.Pos
	IP  int
}

type caseMap map[interface{}]caseEntry

type caseHelper struct {
	ConstMap caseMap // constains all case constants
	GotoMap  caseMap // contains only the constants appearing before any non-constant case expression
	AllConst bool
}

// keep track of constant expressions in cases. error on duplicates
func (seen *caseHelper) add(c *Comp, val interface{}, entry caseEntry) {
	prev, found := seen.ConstMap[val]
	if found {
		c.Errorf("duplicate case %v <%v> in switch\n\tprevious case at %s", val, r.TypeOf(val), c.Fileset.Position(prev.Pos))
		return
	}
	seen.ConstMap[val] = entry
	if seen.AllConst {
		seen.GotoMap[val] = entry
	}
}

func (c *Comp) Switch(node *ast.SwitchStmt, labels []string) {
	initLocals := false
	var initBinds [2]int
	c, initLocals = c.pushEnvIfLocalBinds(&initBinds, node.Init)
	if node.Init != nil {
		c.Stmt(node.Init)
	}
	var ibreak int
	sort.Strings(labels)
	c.Loop = &LoopInfo{
		Break:      &ibreak,
		ThisLabels: labels,
	}

	// tag.Value (if constant) or tag.Fun() will return the tag value at runtime
	var tag *Expr
	tagnode := node.Tag
	if tagnode == nil {
		// "switch { }" without an expression means "switch true { }"
		tag = c.exprUntypedLit(r.Bool, constant.MakeBool(true))
		tagnode = &ast.Ident{NamePos: node.Pos() + 6, Name: "true"} // only for error messages
	} else {
		tag = c.Expr1(tagnode, nil)
	}
	if !tag.Const() {
		// cannot invoke tag.Fun() multiple times because side effects must be applied only once!
		// switchTag saves the result of tag.Fun() in a runtime bind
		// and returns an expression that retrieves it
		tag = c.switchTag(tag)

		if c.Options&OptDebugSleepOnSwitch != 0 {
			c.append(func(env *Env) (Stmt, *Env) {
				Debugf("start sleeping on switch, env = %p", env)
				time.Sleep(time.Second / 30)
				Debugf("done  sleeping on switch, env = %p", env)
				env.IP++
				return env.Code[env.IP], env
			})
		}
	}
	if node.Body != nil {
		// reserve a code slot for switchGotoMap/switchGotoSlice optimizer
		ipswitchgoto := c.Code.Len()
		seen := &caseHelper{make(caseMap), make(caseMap), true} // keeps track of constant expressions in cases. errors on duplicates
		c.Append(stmtNop, node.Body.Pos())

		list := node.Body.List
		defaulti := -1
		var defaultpos token.Pos
		n := len(list)
		for i, stmt := range list {
			switch clause := stmt.(type) {
			case *ast.CaseClause:
				canfallthrough := i < n-1 // last case cannot fallthrough
				if clause.List == nil {
					if defaulti >= 0 {
						c.Errorf("multiple defaults in switch (first at %s)", c.Fileset.Position(defaultpos))
					}
					defaulti = c.Code.Len()
					defaultpos = clause.Pos()
					c.switchDefault(clause, canfallthrough)
				} else {
					c.switchCase(clause, tagnode, tag, canfallthrough, seen)
				}
			default:
				c.Errorf("invalid statement inside switch: expecting case or default, found: %v <%v>", stmt, r.TypeOf(stmt))
			}
		}
		// default is executed as last, if no other case matches
		if defaulti >= 0 {
			// +1 to skip its "never matches" header
			c.Append(func(env *Env) (Stmt, *Env) {
				ip := defaulti + 1
				env.IP = ip
				return env.Code[ip], env
			}, defaultpos)
		}
		// try to optimize
		c.switchGotoMap(tag, seen, ipswitchgoto)
	}
	// we finally know this
	ibreak = c.Code.Len()

	c = c.popEnvIfLocalBinds(initLocals, &initBinds, node.Init)
}

// switchTag takes the expression immediately following a switch,
// compiles it to a statement that evaluates it and saves its result
// in a runtime binding (an interpreter local variable),
// finally returns another expression that retrieves such runtime binding
func (c *Comp) switchTag(e *Expr) *Expr {
	var upn, o = 0, c
	// try to piggyback the binding to a Comp that already has some bindings,
	// but do not cross function boundaries
	for o.BindNum == 0 && o.IntBindNum == 0 && o.Func == nil && o.Outer != nil {
		upn += o.UpCost
		o = o.Outer
	}
	return c.Symbol(c.declUnnamedBind(e, o, upn))
}

// switchCase compiles a case in a switch.
func (c *Comp) switchCase(node *ast.CaseClause, tagnode ast.Expr, tag *Expr, canfallthrough bool, seen *caseHelper) {
	cmpfuns := make([]func(*Env) bool, 0)
	cmpnode := &ast.BinaryExpr{Op: token.EQL, X: tagnode} // for error messages, and Comp.BinaryExpr1 dispatches on its Op

	ibody := c.Code.Len() + 1 // body will start here
	// compile a comparison of tag against each expression
	sometrue := false
	for _, enode := range node.List {
		e := c.Expr1(enode, nil)
		if e.Const() {
			e.ConstTo(tag.Type)
		}
		cmpnode.OpPos = enode.Pos()
		cmpnode.Y = enode
		cmp := c.BinaryExpr1(cmpnode, tag, e)
		if e.Const() {
			seen.add(c, e.Value, caseEntry{Pos: enode.Pos(), IP: ibody})
			if tag.Const() {
				// constant propagation
				flag := cmp.EvalConst(COptDefaults)
				if r.ValueOf(flag).Bool() {
					sometrue = true
					break // always matches, no need to check further expressions
				} else {
					// can never match, skip this expression
					continue
				}
			}
		} else {
			seen.AllConst = false
		}
		// constants are handled above. only add non-constant comparisons to cmpfuns
		cmpfuns = append(cmpfuns, cmp.Fun.(func(*Env) bool))
	}
	// compile like "if tag == e1 || tag == e2 ... { }"
	// and keep track of where to jump if no expression matches
	//
	// always occupy a Code slot for cmpfuns, even if nothing to do.
	// reason: both caseMap optimizer and fallthrough from previous case
	// skip such slot and jump to current body
	var iend int
	var stmt Stmt
	switch len(cmpfuns) {
	case 0:
		if sometrue {
			stmt = stmtNop
		} else {
			// compile anyway, a fallthrough from previous case may still reach the current body
			stmt = func(env *Env) (Stmt, *Env) {
				ip := iend
				env.IP = ip
				return env.Code[ip], env
			}
		}
	case 1:
		cmpfun := cmpfuns[0]
		if sometrue {
			stmt = func(env *Env) (Stmt, *Env) {
				// keep side effects
				cmpfun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		} else {
			stmt = func(env *Env) (Stmt, *Env) {
				var ip int
				if cmpfun(env) {
					ip = env.IP + 1
				} else {
					ip = iend
				}
				env.IP = ip
				return env.Code[ip], env
			}
		}
	case 2:
		cmpfuns := [...]func(*Env) bool{
			cmpfuns[0],
			cmpfuns[1],
		}
		if sometrue {
			stmt = func(env *Env) (Stmt, *Env) {
				// keep side effects
				_ = cmpfuns[0](env) || cmpfuns[1](env)
				env.IP++
				return env.Code[env.IP], env
			}
		} else {
			stmt = func(env *Env) (Stmt, *Env) {
				var ip int
				if cmpfuns[0](env) || cmpfuns[1](env) {
					ip = env.IP + 1
				} else {
					ip = iend
				}
				env.IP = ip
				return env.Code[ip], env
			}
		}
	default:
		if sometrue {
			stmt = func(env *Env) (Stmt, *Env) {
				for _, cmpfun := range cmpfuns {
					// keep side effects
					if cmpfun(env) {
						break
					}
				}
				env.IP++
				return env.Code[env.IP], env
			}
		} else {
			stmt = func(env *Env) (Stmt, *Env) {
				ip := iend
				for _, cmpfun := range cmpfuns {
					if cmpfun(env) {
						ip = env.IP + 1
						break
					}
				}
				env.IP = ip
				return env.Code[ip], env
			}
		}
	}
	c.Append(stmt, node.Pos())
	c.switchCaseBody(node.Body, canfallthrough)
	// we finally know where to jump if match fails
	iend = c.Code.Len()
}

// switchDefault compiles the default case in a switch
func (c *Comp) switchDefault(node *ast.CaseClause, canfallthrough bool) {
	var iend int
	c.Append(func(env *Env) (Stmt, *Env) {
		// jump to the next case. we must always add this statement for three reasons:
		// 1) if default is entered normally, it always fails to match and jumps to the next case
		// 2) if the previous case ends with fallthrough, it will skip this statement and jump to default's body
		// 3) if the switch ends without any matching case, it will manually jump to default's body (skipping this statement)
		ip := iend
		env.IP = ip
		return env.Code[ip], env
	}, node.Pos())
	c.switchCaseBody(node.Body, canfallthrough)
	// we finally know where to jump if match fails
	iend = c.Code.Len()
}

// switchCaseBody compiles the body of a case in a switch
func (c *Comp) switchCaseBody(list []ast.Stmt, canfallthrough bool) {
	var isfallthrough bool
	var endpos token.Pos
	n := len(list)
	if n != 0 {
		isfallthrough = isFallthrough(list[n-1])
		if isfallthrough {
			endpos = list[n-1].Pos()
			if canfallthrough {
				n--
				list = list[:n]
			} else {
				c.Errorf("cannot fallthrough final case in switch")
				return
			}
		} else {
			endpos = list[n-1].End()
		}

		// c.List creates a new scope... not accurate, compiled Go doesn't.
		// but at least isolates per-case variables, as compiled Go does
		if n != 0 {
			c.List(list)
		}
	}
	// after executing the case body, either break or fallthrough
	c.Pos = endpos
	if isfallthrough {
		c.append(stmtFallthrough)
	} else {
		c.jumpOut(0, c.Loop.Break)
	}
}

// stmtFallThrough executes a fallthrough statement - only works inside a switch,
// and cannot be used in the last switch of a case
func stmtFallthrough(env *Env) (Stmt, *Env) {
	env.IP += 2 // +2 to skip the comparisons in next case, and jump directly to its body
	return env.Code[env.IP], env
}

func isFallthrough(node ast.Stmt) bool {
	switch node := node.(type) {
	case *ast.BranchStmt:
		return node.Tok == token.FALLTHROUGH
	default:
		return false
	}
}
