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
 * switch_type.go
 *
 *  Created on May 06, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	"go/types"
	r "reflect"
	"sort"
	"unsafe"

	"github.com/cosmos72/gomacro/typeutil"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type typecaseEntry struct {
	Type xr.Type
	Pos  token.Pos
	IP   int
}

type typecaseHelper struct {
	TypeMap     typeutil.Map // map types.Type -> typecaseEntry
	ConcreteMap typeutil.Map // only contains the initial segment of non-interface types
	AllConcrete bool
}

// keep track of types in type-switch. error on duplicates
func (seen *typecaseHelper) add(c *Comp, t xr.Type, entry typecaseEntry) {
	var gtype types.Type
	if t != nil {
		gtype = t.GoType()
	}
	prev := seen.TypeMap.At(gtype)
	if prev != nil {
		c.Errorf("duplicate case <%v> in switch\n\tprevious case at %s", t, c.Fileset.Position(prev.(typecaseEntry).Pos))
		return
	}
	entry.Type = t
	seen.TypeMap.Set(gtype, entry)
	if t != nil && t.Kind() == r.Interface {
		seen.AllConcrete = false
	} else if seen.AllConcrete {
		seen.ConcreteMap.Set(gtype, entry)
	}
}

/*
func (c *Comp) TypeSwitch(node *ast.TypeSwitchStmt, labels []string) {
	c.Errorf("unimplemented statement: %v <%v>", node, r.TypeOf(node))
}
*/

func (c *Comp) TypeSwitch(node *ast.TypeSwitchStmt, labels []string) {
	initLocals := false
	var initBinds [2]int
	// TypeSwitch always allocates at least a bind "" in typeswitchTag()
	c, initLocals = c.pushEnvIfFlag(&initBinds, true)
	if node.Init != nil {
		c.Stmt(node.Init)
	}
	var ibreak int
	sort.Strings(labels)
	c.Loop = &LoopInfo{
		Break:      &ibreak,
		ThisLabels: labels,
	}

	tagnode, varname := c.typeswitchNode(node.Assign)
	tagexpr := c.Expr1(tagnode, nil)
	if tagexpr.Type.Kind() != r.Interface {
		c.Errorf("cannot type switch on non-interface type <%v>: %v", tagexpr.Type, tagnode)
	}
	if tagexpr.Const() {
		c.Warnf("type switch on constant!? %v = %v <%v>", tagnode, tagexpr.Value, tagexpr.Type)
	}
	// just like Comp.Switch, we cannot invoke tagexpr.Fun() multiple times because
	// side effects must be applied only once!
	// typeswitchTag saves the result of tagexpr.Fun() in a runtime bind
	// and returns the bind.
	bind := c.typeswitchTag(tagexpr)

	if node.Body != nil {
		// reserve a code slot for typeSwitchGotoMap optimizer
		ipswitchgoto := c.Code.Len()
		seen := &typecaseHelper{AllConcrete: true} // keeps track of types in cases. errors on duplicates
		c.Append(stmtNop, node.Body.Pos())

		list := node.Body.List
		defaulti := -1
		var defaultpos token.Pos
		for _, stmt := range list {
			c.Pos = stmt.Pos()
			switch clause := stmt.(type) {
			case *ast.CaseClause:
				if clause.List == nil {
					if defaulti >= 0 {
						c.Errorf("multiple defaults in switch (first at %s)", c.Fileset.Position(defaultpos))
					}
					defaulti = c.Code.Len()
					defaultpos = clause.Pos()
					c.typeswitchDefault(clause, varname, bind)
				} else {
					c.typeswitchCase(clause, varname, bind, seen)
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
		c.typeswitchGotoMap(bind, seen, ipswitchgoto)
	}
	// we finally know this
	ibreak = c.Code.Len()

	c = c.popEnvIfLocalBinds(initLocals, &initBinds, node.Init, node.Assign)
}

// typeswitchNode returns the expression to type-switch on.
// if such expression is used to declare a variable, the variable name is returned too
func (c *Comp) typeswitchNode(stmt ast.Stmt) (ast.Expr, string) {
	var varname string // empty, or name of variable in 'switch varname := expression.(type)'
	var tagnode ast.Expr
	switch stmt := stmt.(type) {
	case *ast.AssignStmt:
		if len(stmt.Lhs) == 1 && len(stmt.Rhs) == 1 && stmt.Tok == token.DEFINE {
			if lhs, ok := stmt.Lhs[0].(*ast.Ident); ok {
				varname = lhs.Name
				tagnode = stmt.Rhs[0]
			}
		}
	case *ast.ExprStmt:
		tagnode = stmt.X
	}

	for {
		switch e := tagnode.(type) {
		case *ast.ParenExpr:
			tagnode = e.X
			continue
		case *ast.TypeAssertExpr:
			if e.Type != nil {
				c.Errorf("invalid type switch: expecting '.(type)', found type assertion: %v", stmt)
			}
			tagnode = e.X
		default:
			tagnode = e
		}
		break
	}
	if tagnode == nil {
		c.Errorf("expected type-switch expression, found: %v", stmt)
	}
	return tagnode, varname
}

// typeswitchTag takes the expression immediately following a type-switch,
// compiles it to a statement that evaluates it and saves its result and its type
// in two runtime bindings (interpreter local variables),
// finally returns another expression that retrieves the expression value
// with its concrete type
func (c *Comp) typeswitchTag(e *Expr) *Bind {
	bind := c.NewBind("", VarBind, e.Type)               // e.Type must be an interface type...
	tbind := c.NewBind("", VarBind, c.TypeOfInterface()) // no need to store as xr.Type
	if tbind.Desc.Index() != bind.Desc.Index()+1 {
		c.Errorf("internal error: consecutive binds have non-consecutive indexes %d and %d",
			bind.Desc.Index()+1, tbind.Desc.Index())
	}

	extractor := c.extractor(e.Type)

	// c.Debugf("typeswitchTag: allocated bind %v", bind)
	switch bind.Desc.Class() {
	case VarBind:
		// cannot use c.DeclVar0 because the variable is declared in o
		// cannot use o.DeclVar0 because the initializer must be evaluated in c
		// so initialize the binding manually
		index := bind.Desc.Index()
		init := e.AsX1()
		c.append(func(env *Env) (Stmt, *Env) {
			v, xt := extractor(init(env)) // extract value with concrete type
			// Debugf("typeswitchTag = %v <%v>", v, ValueType(v))
			// no need to create a settable reflect.Value
			env.Vals[index] = v
			env.Vals[index+1] = r.ValueOf(xt)
			env.IP++
			return env.Code[env.IP], env
		})
	default:
		c.Errorf("internal error! Comp.NewBind(name=%q, class=VarBind, type=%v) returned class=%v, expecting VarBind",
			"", bind.Type, bind.Desc.Class())
		return nil
	}
	return bind
}

// typeswitchGotoMap tries to optimize the dispatching of a type-switch
func (c *Comp) typeswitchGotoMap(bind *Bind, seen *typecaseHelper, ip int) {
	if seen.ConcreteMap.Len() <= 1 {
		return
	}
	m := make(map[r.Type]int)
	seen.ConcreteMap.Iterate(func(k types.Type, v interface{}) {
		entry := v.(typecaseEntry)
		m[entry.Type.ReflectType()] = entry.IP
	})
	if len(m) != seen.ConcreteMap.Len() {
		// one or more interpreted types are implemented by the same reflect.Type.
		// cannot optimize typeswitch based on reflect.Type only.
		return
	}
	idx := bind.Desc.Index()

	stmt := func(env *Env) (Stmt, *Env) {
		var rtype r.Type
		if v := env.Vals[idx]; v.IsValid() {
			rtype = v.Type() // concrete reflect.Type already extracted by typeswitchTag
		}
		if ip, found := m[rtype]; found {
			env.IP = ip
		} else {
			env.IP++
		}
		return env.Code[env.IP], env
	}
	c.Code.List[ip] = stmt
}

// typeswitchCase compiles a case in a type-switch.
func (c *Comp) typeswitchCase(node *ast.CaseClause, varname string, bind *Bind, seen *typecaseHelper) {

	ibody := c.Code.Len() + 1 // body will start here
	ts := make([]xr.Type, len(node.List))
	rtypes := make([]r.Type, len(node.List))

	// compile a comparison of tag against each type
	for i, enode := range node.List {
		t := c.compileTypeOrNil(enode)
		if t != nil {
			rtypes[i] = t.ReflectType()
			if t.Kind() != r.Interface && !t.Implements(bind.Type) {
				c.Errorf("impossible typeswitch case: <%v> does not implement <%v>", t, bind.Type)
			}
		}
		ts[i] = t
		seen.add(c, t, typecaseEntry{Pos: enode.Pos(), IP: ibody})
	}
	// compile like "if r.TypeOf(bind) == t1 || r.TypeOf(bind) == t2 ... { }"
	// and keep track of where to jump if no expression matches
	//
	// always occupy a Code slot for type comparison, even if nothing to do.
	// reason: typeswitchGotoMap optimizer skips such slot and jumps to current body
	var iend int
	var stmt Stmt
	idx := bind.Desc.Index()
	switch len(node.List) {
	case 0:
		// compile anyway. reachable?
		stmt = func(env *Env) (Stmt, *Env) {
			// Debugf("typeswitchCase: comparing %v against zero types", tagfun(env))
			ip := iend
			env.IP = ip
			return env.Code[ip], env
		}
	case 1:
		t := ts[0]
		rtype := rtypes[0]
		if t == nil {
			// case nil:
			stmt = func(env *Env) (Stmt, *Env) {
				v := env.Vals[idx]
				// Debugf("typeswitchCase: comparing %v <%v> against nil type", v, ValueType(v))
				var ip int
				if v.IsValid() {
					ip = iend
				} else {
					ip = env.IP + 1
				}
				env.IP = ip
				return env.Code[ip], env
			}
		} else if t.Kind() == r.Interface && xr.IsEmulatedInterface(t) {
			// case emulated_interface:
			stmt = func(env *Env) (Stmt, *Env) {
				v := env.Vals[idx]
				// Debugf("typeswitchCase: comparing %v <%v> against interface type %v", v, ValueType(v), rtype)
				ip := iend
				if v.IsValid() {
					// rtype may be an interpreted type:
					// extract the concrete xr.Type and use it
					xtv := env.Vals[idx+1]
					if xtv.IsValid() && !xtv.IsNil() {
						xt := xtv.Interface().(xr.Type)
						if xt.Implements(t) {
							ip = env.IP + 1
							// need the compiler at run-time :(
							conv := c.converterToEmulatedInterface(xt, t)
							env.Vals[idx] = conv(v)
						}
					}
				}
				env.IP = ip
				return env.Code[ip], env
			}
		} else if t.Kind() == r.Interface {
			// case interface:
			stmt = func(env *Env) (Stmt, *Env) {
				v := env.Vals[idx]
				// Debugf("typeswitchCase: comparing %v <%v> against interface type %v", v, ValueType(v), rtype)
				ip := iend
				if v.IsValid() {
					if v.Type().Implements(rtype) {
						ip = env.IP + 1
					} else {
						// rtype may be an interpreted type:
						// extract the concrete xr.Type and use it
						xtv := env.Vals[idx+1]
						if xtv.IsValid() && !xtv.IsNil() {
							xt := xtv.Interface().(xr.Type)
							if xt.Implements(t) {
								ip = env.IP + 1
								// need the compiler at run-time :(
								conv := c.converterToProxy(xt, t)
								env.Vals[idx] = conv(v)
							}
						}
					}
				}
				env.IP = ip
				return env.Code[ip], env
			}
		} else {
			// case concrete_type:
			stmt = func(env *Env) (Stmt, *Env) {
				v := env.Vals[idx]
				ip := iend
				if v.IsValid() && v.Type() == rtype {
					// rtype may be an interpreted type:
					// extract the concrete xr.Type and use it
					xtv := env.Vals[idx+1]
					if xtv.IsValid() && !xtv.IsNil() {
						xt := xtv.Interface().(xr.Type)
						if xt.IdenticalTo(t) {
							ip = env.IP + 1
						}
					} else {
						// cannot check exactly...
						ip = env.IP + 1
					}
				}
				env.IP = ip
				return env.Code[ip], env
			}
		}
	default:
		stmt = func(env *Env) (Stmt, *Env) {
			v := env.Vals[idx]
			var vt r.Type
			if v.IsValid() {
				vt = v.Type()
			}
			// Debugf("typeswitchCase: comparing %v <%v> against types %v", v, vt, rtypes)
			ip := iend
			for _, rtype := range rtypes {
				switch {
				case vt == rtype:
				case rtype != nil:
					if rtype.Kind() != r.Interface || !vt.Implements(rtype) {
						continue
					}
				default: // rtype == nil
					if v.IsValid() {
						continue
					}
				}
				// Debugf("typeswitchCase: v <%v> matches type %v", v, vt, rtype)
				ip = env.IP + 1
				break
			}
			env.IP = ip
			return env.Code[ip], env
		}
	}
	c.Pos = node.Pos()
	c.append(stmt)
	var t xr.Type
	if len(ts) == 1 {
		t = ts[0]
	}
	c.typeswitchBody(node.Body, varname, t, bind)
	// we finally know where to jump if match fails
	iend = c.Code.Len()
}

// typeswitchDefault compiles the default case in a type-switch.
func (c *Comp) typeswitchDefault(node *ast.CaseClause, varname string, bind *Bind) {
	var iend int
	stmt := func(env *Env) (Stmt, *Env) {
		// Debugf("typeswitchDefault: default entered normally, skipping it")
		ip := iend
		env.IP = ip
		return env.Code[ip], env
	}
	c.Pos = node.Pos()
	c.append(stmt)
	c.typeswitchBody(node.Body, varname, nil, bind)
	iend = c.Code.Len()
}

// typeswitchBody compiles the body of a case in a type-switch.
func (c *Comp) typeswitchBody(list []ast.Stmt, varname string, t xr.Type, bind *Bind) {
	list1 := list
	if list1 == nil {
		list1 = []ast.Stmt{nil}
	}
	declvar := varname != "" && varname != "_"
	locals := declvar || containLocalBinds(list1...)
	var nbinds [2]int

	c2, locals2 := c.pushEnvIfFlag(&nbinds, locals)
	if declvar {
		sym := bind.AsSymbol(c2.UpCost)
		if t == nil {
			t = sym.Type
		}
		// cannot simply use sym as varname initializer: it returns the wrong type
		c2.typeswitchVar(varname, t, sym)
	}
	for _, stmt := range list {
		c2.Stmt(stmt)
	}
	c2.jumpOut(c2.UpCost, c.Loop.Break)
	c2.popEnvIfLocalBinds(locals2, &nbinds, list1...)
}

// typeswitchVar compiles the tag variable declaration in a type-switch.
func (c *Comp) typeswitchVar(varname string, t xr.Type, sym *Symbol) {
	sidx := sym.Bind.Desc.Index()

	bind := c.NewBind(varname, VarBind, t)
	idx := bind.Desc.Index()

	if sym.Upn != 1 {
		c.Errorf("typeswitchVar: impossible sym.Upn = %v", sym.Upn)
	}
	var stmt Stmt
	switch t.Kind() {
	case r.Bool:
		stmt = func(env *Env) (Stmt, *Env) {
			*(*bool)(unsafe.Pointer(&env.Ints[idx])) = env.Outer.Vals[sidx].Bool()
			env.IP++
			return env.Code[env.IP], env
		}
	case r.Int:
		stmt = func(env *Env) (Stmt, *Env) {
			*(*int)(unsafe.Pointer(&env.Ints[idx])) = int(env.Outer.Vals[sidx].Int())
			env.IP++
			return env.Code[env.IP], env
		}
	case r.Int8:
		stmt = func(env *Env) (Stmt, *Env) {
			*(*int8)(unsafe.Pointer(&env.Ints[idx])) = int8(env.Outer.Vals[sidx].Int())
			env.IP++
			return env.Code[env.IP], env
		}
	case r.Int16:
		stmt = func(env *Env) (Stmt, *Env) {
			*(*int16)(unsafe.Pointer(&env.Ints[idx])) = int16(env.Outer.Vals[sidx].Int())
			env.IP++
			return env.Code[env.IP], env
		}
	case r.Int32:
		stmt = func(env *Env) (Stmt, *Env) {
			*(*int32)(unsafe.Pointer(&env.Ints[idx])) = int32(env.Outer.Vals[sidx].Int())
			env.IP++
			return env.Code[env.IP], env
		}
	case r.Int64:
		stmt = func(env *Env) (Stmt, *Env) {
			*(*int64)(unsafe.Pointer(&env.Ints[idx])) = int64(env.Outer.Vals[sidx].Int())
			env.IP++
			return env.Code[env.IP], env
		}
	case r.Uint:
		stmt = func(env *Env) (Stmt, *Env) {
			*(*uint)(unsafe.Pointer(&env.Ints[idx])) = uint(env.Outer.Vals[sidx].Uint())
			env.IP++
			return env.Code[env.IP], env
		}
	case r.Uint8:
		stmt = func(env *Env) (Stmt, *Env) {
			*(*uint8)(unsafe.Pointer(&env.Ints[idx])) = uint8(env.Outer.Vals[sidx].Uint())
			env.IP++
			return env.Code[env.IP], env
		}
	case r.Uint16:
		stmt = func(env *Env) (Stmt, *Env) {
			*(*uint16)(unsafe.Pointer(&env.Ints[idx])) = uint16(env.Outer.Vals[sidx].Uint())
			env.IP++
			return env.Code[env.IP], env
		}
	case r.Uint32:
		stmt = func(env *Env) (Stmt, *Env) {
			*(*uint32)(unsafe.Pointer(&env.Ints[idx])) = uint32(env.Outer.Vals[sidx].Uint())
			env.IP++
			return env.Code[env.IP], env
		}
	case r.Uint64:
		stmt = func(env *Env) (Stmt, *Env) {
			env.Ints[idx] = env.Outer.Vals[sidx].Uint()
			env.IP++
			return env.Code[env.IP], env
		}
	case r.Uintptr:
		stmt = func(env *Env) (Stmt, *Env) {
			*(*uintptr)(unsafe.Pointer(&env.Ints[idx])) = uintptr(env.Outer.Vals[sidx].Uint())
			env.IP++
			return env.Code[env.IP], env
		}
	case r.Float32:
		stmt = func(env *Env) (Stmt, *Env) {
			*(*float32)(unsafe.Pointer(&env.Ints[idx])) = float32(env.Outer.Vals[sidx].Float())
			env.IP++
			return env.Code[env.IP], env
		}
	case r.Float64:
		stmt = func(env *Env) (Stmt, *Env) {
			*(*float64)(unsafe.Pointer(&env.Ints[idx])) = env.Outer.Vals[sidx].Float()
			env.IP++
			return env.Code[env.IP], env
		}
	case r.Complex64:
		stmt = func(env *Env) (Stmt, *Env) {
			*(*complex64)(unsafe.Pointer(&env.Ints[idx])) = complex64(env.Outer.Vals[sidx].Complex())
			env.IP++
			return env.Code[env.IP], env
		}
	case r.Complex128:
		stmt = func(env *Env) (Stmt, *Env) {
			*(*complex128)(unsafe.Pointer(&env.Ints[idx])) = env.Outer.Vals[sidx].Complex()
			env.IP++
			return env.Code[env.IP], env
		}
	default:
		rtype := t.ReflectType()
		zero := r.Zero(rtype)
		stmt = func(env *Env) (Stmt, *Env) {
			v := env.Outer.Vals[sidx]
			place := r.New(rtype).Elem()
			if !v.IsValid() {
				v = zero
			} else if v.Type() != rtype {
				v = convert(v, rtype)
			}
			place.Set(v)
			env.Vals[idx] = place
			env.IP++
			return env.Code[env.IP], env
		}
	}
	c.append(stmt)
}
