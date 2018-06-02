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
 * identifier.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
	"unsafe"

	"github.com/cosmos72/gomacro/base"
)

func (c *Comp) Resolve(name string) *Symbol {
	sym := c.TryResolve(name)
	if sym == nil {
		c.Errorf("undefined identifier: %v", name)
	}
	return sym
}

func (c *Comp) TryResolve(name string) *Symbol {
	upn := 0
	for ; c != nil; c = c.Outer {
		if bind, ok := c.Binds[name]; ok {
			// c.Debugf("TryResolve: %s is upn=%d %v", name, upn, bind)
			return bind.AsSymbol(upn)
		}
		upn += c.UpCost // c.UpCost is zero if *Comp has no local variables/functions so it will NOT have a corresponding *Env at runtime
	}
	return nil
}

// Ident compiles a read operation on a constant, variable or function
func (c *Comp) Ident(name string) *Expr {
	return c.Symbol(c.Resolve(name))
}

// IdentPlace compiles an assignment to a variable, or taking the address of a variable
func (c *Comp) IdentPlace(name string, opt PlaceOption) *Place {
	if name == "_" {
		if opt == PlaceAddress {
			c.Errorf("%s _", opt)
			return nil
		}
		// assignment to _ is allowed: it does nothing
		bind := c.NewBind(name, VarBind, c.TypeOfInterface())
		return &Place{Var: *bind.AsVar(0, PlaceSettable)}
	}
	sym := c.Resolve(name)
	return &Place{Var: *sym.AsVar(opt)}
}

// Bind compiles a read operation on a constant, variable or function declared in 'c'
func (c *Comp) Bind(bind *Bind) *Expr {
	return bind.Expr(&c.Globals.Stringer)
}

// Symbol compiles a read operation on a constant, variable or function
func (c *Comp) Symbol(sym *Symbol) *Expr {
	return sym.Expr(c.Depth, &c.Globals.Stringer)
}

// Expr returns an expression that will read the given Bind at runtime
func (bind *Bind) Expr(st *base.Stringer) *Expr {
	switch bind.Desc.Class() {
	case ConstBind:
		return exprLit(bind.Lit, bind.AsSymbol(0))
	case VarBind, FuncBind:
		return bind.expr(st)
	case IntBind:
		return bind.intExpr(st)
	default:
		st.Errorf("unknown symbol class %s", bind.Desc.Class())
	}
	return nil
}

// Expr returns an expression that will read the given Symbol at runtime
func (sym *Symbol) Expr(depth int, st *base.Stringer) *Expr {
	switch sym.Desc.Class() {
	case ConstBind:
		return exprLit(sym.Lit, sym)
	case VarBind, FuncBind:
		return sym.expr(depth, st)
	case IntBind:
		return sym.intExpr(depth, st)
	default:
		st.Errorf("unknown symbol class %s", sym.Desc.Class())
	}
	return nil
}

// upn must be >= 3
func outerEnv3(env *Env, upn int) *Env {
	for ; upn >= 3; upn -= 3 {
		env = env.Outer.Outer.Outer
	}
	switch upn {
	case 2:
		env = env.Outer
		fallthrough
	case 1:
		env = env.Outer
	}
	return env
}

// return an expression that will read Bind value at runtime
func (bind *Bind) expr(st *base.Stringer) *Expr {
	idx := bind.Desc.Index()
	var fun I

	// if package is (partially) compiled, kind can also be one of Bool, Int*, Uint*, Float*, Complex64
	switch bind.Type.Kind() {
	case r.Bool:
		fun = func(env *Env) bool {
			return env.Vals[idx].Bool()
		}
	case r.Int:
		fun = func(env *Env) int {
			return int(env.Vals[idx].Int())
		}
	case r.Int8:
		fun = func(env *Env) int8 {
			return int8(env.Vals[idx].Int())
		}
	case r.Int16:
		fun = func(env *Env) int16 {
			return int16(env.Vals[idx].Int())
		}
	case r.Int32:
		fun = func(env *Env) int32 {
			return int32(env.Vals[idx].Int())
		}
	case r.Int64:
		fun = func(env *Env) int64 {
			return env.Vals[idx].Int()
		}
	case r.Uint:
		fun = func(env *Env) uint {
			return uint(env.Vals[idx].Uint())
		}
	case r.Uint8:
		fun = func(env *Env) uint8 {
			return uint8(env.Vals[idx].Uint())
		}
	case r.Uint16:
		fun = func(env *Env) uint16 {
			return uint16(env.Vals[idx].Uint())
		}
	case r.Uint32:
		fun = func(env *Env) uint32 {
			return uint32(env.Vals[idx].Uint())
		}
	case r.Uint64:
		fun = func(env *Env) uint64 {
			return env.Vals[idx].Uint()
		}
	case r.Uintptr:
		fun = func(env *Env) uintptr {
			return uintptr(env.Vals[idx].Uint())
		}
	case r.Float32:
		fun = func(env *Env) float32 {
			return float32(env.Vals[idx].Float())
		}
	case r.Float64:
		fun = func(env *Env) float64 {
			return env.Vals[idx].Float()
		}
	case r.Complex64:
		fun = func(env *Env) complex64 {
			return complex64(env.Vals[idx].Complex())
		}
	case r.Complex128:
		fun = func(env *Env) complex128 {
			return env.Vals[idx].Complex()
		}
	case r.String:
		fun = func(env *Env) string {
			return env.Vals[idx].String()
		}
	default:
		fun = func(env *Env) r.Value {
			return env.Vals[idx]
		}
	}
	return &Expr{Lit: Lit{Type: bind.Type}, Fun: fun, Sym: bind.AsSymbol(0)}
}

// return an expression that will read Symbol value at runtime
func (sym *Symbol) expr(depth int, st *base.Stringer) *Expr {
	idx := sym.Desc.Index()
	upn := sym.Upn
	kind := sym.Type.Kind()
	var fun I
	switch upn {
	case 0:
		return sym.Bind.expr(st)
	case 1:
		switch kind {
		case r.Bool:
			fun = func(env *Env) bool {
				return env.Outer.Vals[idx].Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				return int(env.Outer.Vals[idx].Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				return int8(env.Outer.Vals[idx].Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				return int16(env.Outer.Vals[idx].Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				return int32(env.Outer.Vals[idx].Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				return env.Outer.Vals[idx].Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				return uint(env.Outer.Vals[idx].Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				return uint8(env.Outer.Vals[idx].Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				return uint16(env.Outer.Vals[idx].Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				return uint32(env.Outer.Vals[idx].Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				return env.Outer.Vals[idx].Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				return uintptr(env.Outer.Vals[idx].Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				return float32(env.Outer.Vals[idx].Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				return env.Outer.Vals[idx].Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				return complex64(env.Outer.Vals[idx].Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				return env.Outer.Vals[idx].Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				return env.Outer.Vals[idx].String()
			}
		default:
			fun = func(env *Env) r.Value {
				return env.Outer.Vals[idx]
			}
		}
	case 2:
		switch kind {
		case r.Bool:
			fun = func(env *Env) bool {
				return env.Outer.Outer.Vals[idx].Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				return int(env.Outer.Outer.Vals[idx].Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				return int8(env.Outer.Outer.Vals[idx].Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				return int16(env.Outer.Outer.Vals[idx].Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				return int32(env.Outer.Outer.Vals[idx].Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				return env.Outer.Outer.Vals[idx].Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				return uint(env.Outer.Outer.Vals[idx].Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				return uint8(env.Outer.Outer.Vals[idx].Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				return uint16(env.Outer.Outer.Vals[idx].Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				return uint32(env.Outer.Outer.Vals[idx].Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				return env.Outer.Outer.Vals[idx].Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				return uintptr(env.Outer.Outer.Vals[idx].Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				return float32(env.Outer.Outer.Vals[idx].Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				return env.Outer.Outer.Vals[idx].Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				return complex64(env.Outer.Outer.Vals[idx].Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				return env.Outer.Outer.Vals[idx].Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				return env.Outer.Outer.Vals[idx].String()
			}
		default:
			fun = func(env *Env) r.Value {
				return env.Outer.Outer.Vals[idx]
			}
		}
	case depth - 1:
		switch kind {
		case r.Bool:
			fun = func(env *Env) bool {
				return env.FileEnv.Vals[idx].Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				return int(env.FileEnv.Vals[idx].Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				return int8(env.FileEnv.Vals[idx].Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				return int16(env.FileEnv.Vals[idx].Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				return int32(env.FileEnv.Vals[idx].Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				return env.FileEnv.Vals[idx].Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				return uint(env.FileEnv.Vals[idx].Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				return uint8(env.FileEnv.Vals[idx].Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				return uint16(env.FileEnv.Vals[idx].Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				return uint32(env.FileEnv.Vals[idx].Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				return env.FileEnv.Vals[idx].Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				return uintptr(env.FileEnv.Vals[idx].Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				return float32(env.FileEnv.Vals[idx].Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				return env.FileEnv.Vals[idx].Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				return complex64(env.FileEnv.Vals[idx].Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				return env.FileEnv.Vals[idx].Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				return env.FileEnv.Vals[idx].String()
			}
		default:
			fun = func(env *Env) r.Value {
				return env.FileEnv.Vals[idx]
			}
		}
	case depth: // TopEnv should not contain variables or functions... but no harm
		switch kind {
		case r.Bool:
			fun = func(env *Env) bool {
				return env.FileEnv.Outer.Vals[idx].Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				return int(env.FileEnv.Outer.Vals[idx].Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				return int8(env.FileEnv.Outer.Vals[idx].Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				return int16(env.FileEnv.Outer.Vals[idx].Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				return int32(env.FileEnv.Outer.Vals[idx].Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				return env.FileEnv.Outer.Vals[idx].Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				return uint(env.FileEnv.Outer.Vals[idx].Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				return uint8(env.FileEnv.Outer.Vals[idx].Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				return uint16(env.FileEnv.Outer.Vals[idx].Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				return uint32(env.FileEnv.Outer.Vals[idx].Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				return env.FileEnv.Outer.Vals[idx].Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				return uintptr(env.FileEnv.Outer.Vals[idx].Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				return float32(env.FileEnv.Outer.Vals[idx].Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				return env.FileEnv.Outer.Vals[idx].Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				return complex64(env.FileEnv.Outer.Vals[idx].Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				return env.FileEnv.Outer.Vals[idx].Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				return env.FileEnv.Outer.Vals[idx].String()
			}
		default:
			fun = func(env *Env) r.Value {
				return env.FileEnv.Outer.Vals[idx]
			}
		}
	default:
		switch kind {
		case r.Bool:
			fun = func(env *Env) bool {
				env = env.Up(upn)
				return env.Vals[idx].Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				env = env.Up(upn)
				return int(env.Vals[idx].Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				env = env.Up(upn)
				return int8(env.Vals[idx].Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				env = env.Up(upn)
				return int16(env.Vals[idx].Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				env = env.Up(upn)
				return int32(env.Vals[idx].Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				env = env.Up(upn)
				return env.Vals[idx].Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				env = env.Up(upn)
				return uint(env.Vals[idx].Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				env = env.Up(upn)
				return uint8(env.Vals[idx].Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				env = env.Up(upn)
				return uint16(env.Vals[idx].Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				env = env.Up(upn)
				return uint32(env.Vals[idx].Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				env = env.Up(upn)
				return env.Vals[idx].Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				env = env.Up(upn)
				return uintptr(env.Vals[idx].Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				env = env.Up(upn)
				return float32(env.Vals[idx].Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				env = env.Up(upn)
				return env.Vals[idx].Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				env = env.Up(upn)
				return complex64(env.Vals[idx].Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				env = env.Up(upn)
				return env.Vals[idx].Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				env = env.Up(upn)
				return env.Vals[idx].String()
			}
		default:
			fun = func(env *Env) r.Value {
				env = env.Up(upn)
				return env.Vals[idx]
			}
		}
	}
	return &Expr{Lit: Lit{Type: sym.Type}, Fun: fun, Sym: sym}
}

// return an expression that will read Bind optimized value at runtime
func (bind *Bind) intExpr(st *base.Stringer) *Expr {
	idx := bind.Desc.Index()
	var fun I
	switch bind.Type.Kind() {
	case r.Bool:
		fun = func(env *Env) bool {
			return *(*bool)(unsafe.Pointer(&env.Ints[idx]))
		}
	case r.Int:
		fun = func(env *Env) int {
			return *(*int)(unsafe.Pointer(&env.Ints[idx]))
		}
	case r.Int8:
		fun = func(env *Env) int8 {
			return *(*int8)(unsafe.Pointer(&env.Ints[idx]))
		}
	case r.Int16:
		fun = func(env *Env) int16 {
			return *(*int16)(unsafe.Pointer(&env.Ints[idx]))
		}
	case r.Int32:
		fun = func(env *Env) int32 {
			return *(*int32)(unsafe.Pointer(&env.Ints[idx]))
		}
	case r.Int64:
		fun = func(env *Env) int64 {
			return *(*int64)(unsafe.Pointer(&env.Ints[idx]))
		}
	case r.Uint:
		fun = func(env *Env) uint {
			return *(*uint)(unsafe.Pointer(&env.Ints[idx]))
		}
	case r.Uint8:
		fun = func(env *Env) uint8 {
			return *(*uint8)(unsafe.Pointer(&env.Ints[idx]))
		}
	case r.Uint16:
		fun = func(env *Env) uint16 {
			return *(*uint16)(unsafe.Pointer(&env.Ints[idx]))
		}
	case r.Uint32:
		fun = func(env *Env) uint32 {
			return *(*uint32)(unsafe.Pointer(&env.Ints[idx]))
		}
	case r.Uint64:
		fun = func(env *Env) uint64 {
			return env.Ints[idx]
		}
	case r.Uintptr:
		fun = func(env *Env) uintptr {
			return *(*uintptr)(unsafe.Pointer(&env.Ints[idx]))
		}
	case r.Float32:
		fun = func(env *Env) float32 {
			return *(*float32)(unsafe.Pointer(&env.Ints[idx]))
		}
	case r.Float64:
		fun = func(env *Env) float64 {
			return *(*float64)(unsafe.Pointer(&env.Ints[idx]))
		}
	case r.Complex64:
		fun = func(env *Env) complex64 {
			return *(*complex64)(unsafe.Pointer(&env.Ints[idx]))
		}
	case r.Complex128:
		fun = func(env *Env) complex128 {
			return *(*complex128)(unsafe.Pointer(&env.Ints[idx]))
		}
	default:
		st.Errorf("unsupported symbol type, cannot use for optimized read: %s %s <%v>", bind.Desc.Class(), bind.Name, bind.Type)
		return nil
	}
	return &Expr{Lit: Lit{Type: bind.Type}, Fun: fun, Sym: bind.AsSymbol(0)}
}

// return an expression that will read Symbol optimized value at runtime
func (sym *Symbol) intExpr(depth int, st *base.Stringer) *Expr {
	upn := sym.Upn
	k := sym.Type.Kind()
	idx := sym.Desc.Index()
	var fun I
	switch upn {
	case 0:
		return sym.Bind.intExpr(st)
	case 1:
		switch k {
		case r.Bool:
			fun = func(env *Env) bool {
				return *(*bool)(unsafe.Pointer(&env.Outer.Ints[idx]))
			}
		case r.Int:
			fun = func(env *Env) int {
				return *(*int)(unsafe.Pointer(&env.Outer.Ints[idx]))
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				return *(*int8)(unsafe.Pointer(&env.Outer.Ints[idx]))
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				return *(*int16)(unsafe.Pointer(&env.Outer.Ints[idx]))
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				return *(*int32)(unsafe.Pointer(&env.Outer.Ints[idx]))
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				return *(*int64)(unsafe.Pointer(&env.Outer.Ints[idx]))
			}
		case r.Uint:
			fun = func(env *Env) uint {
				return *(*uint)(unsafe.Pointer(&env.Outer.Ints[idx]))
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				return *(*uint8)(unsafe.Pointer(&env.Outer.Ints[idx]))
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				return *(*uint16)(unsafe.Pointer(&env.Outer.Ints[idx]))
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				return *(*uint32)(unsafe.Pointer(&env.Outer.Ints[idx]))
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				return env.Outer.Ints[idx]
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				return *(*uintptr)(unsafe.Pointer(&env.Outer.Ints[idx]))
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				return *(*float32)(unsafe.Pointer(&env.Outer.Ints[idx]))
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				return *(*float64)(unsafe.Pointer(&env.Outer.Ints[idx]))
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				return *(*complex64)(unsafe.Pointer(&env.Outer.Ints[idx]))
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				return *(*complex128)(unsafe.Pointer(&env.Outer.Ints[idx]))
			}
		}
	case 2:
		switch k {
		case r.Bool:
			fun = func(env *Env) bool {
				return *(*bool)(unsafe.Pointer(&env.Outer.Outer.Ints[idx]))
			}
		case r.Int:
			fun = func(env *Env) int {
				return *(*int)(unsafe.Pointer(&env.Outer.Outer.Ints[idx]))
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				return *(*int8)(unsafe.Pointer(&env.Outer.Outer.Ints[idx]))
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				return *(*int16)(unsafe.Pointer(&env.Outer.Outer.Ints[idx]))
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				return *(*int32)(unsafe.Pointer(&env.Outer.Outer.Ints[idx]))
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				return *(*int64)(unsafe.Pointer(&env.Outer.Outer.Ints[idx]))
			}
		case r.Uint:
			fun = func(env *Env) uint {
				return *(*uint)(unsafe.Pointer(&env.Outer.Outer.Ints[idx]))
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				return *(*uint8)(unsafe.Pointer(&env.Outer.Outer.Ints[idx]))
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				return *(*uint16)(unsafe.Pointer(&env.Outer.Outer.Ints[idx]))
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				return *(*uint32)(unsafe.Pointer(&env.Outer.Outer.Ints[idx]))
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				return env.Outer.Outer.Ints[idx]
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				return *(*uintptr)(unsafe.Pointer(&env.Outer.Outer.Ints[idx]))
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				return *(*float32)(unsafe.Pointer(&env.Outer.Outer.Ints[idx]))
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				return *(*float64)(unsafe.Pointer(&env.Outer.Outer.Ints[idx]))
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				return *(*complex64)(unsafe.Pointer(&env.Outer.Outer.Ints[idx]))
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				return *(*complex128)(unsafe.Pointer(&env.Outer.Outer.Ints[idx]))
			}
		}
	case depth - 1:
		switch k {
		case r.Bool:
			fun = func(env *Env) bool {
				return *(*bool)(unsafe.Pointer(&env.FileEnv.Ints[idx]))
			}
		case r.Int:
			fun = func(env *Env) int {
				return *(*int)(unsafe.Pointer(&env.FileEnv.Ints[idx]))
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				return *(*int8)(unsafe.Pointer(&env.FileEnv.Ints[idx]))
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				return *(*int16)(unsafe.Pointer(&env.FileEnv.Ints[idx]))
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				return *(*int32)(unsafe.Pointer(&env.FileEnv.Ints[idx]))
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				return *(*int64)(unsafe.Pointer(&env.FileEnv.Ints[idx]))
			}
		case r.Uint:
			fun = func(env *Env) uint {
				return *(*uint)(unsafe.Pointer(&env.FileEnv.Ints[idx]))
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				return *(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[idx]))
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				return *(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[idx]))
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				return *(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[idx]))
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				return env.FileEnv.Ints[idx]
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				return *(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[idx]))
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				return *(*float32)(unsafe.Pointer(&env.FileEnv.Ints[idx]))
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				return *(*float64)(unsafe.Pointer(&env.FileEnv.Ints[idx]))
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				return *(*complex64)(unsafe.Pointer(&env.FileEnv.Ints[idx]))
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				return *(*complex128)(unsafe.Pointer(&env.FileEnv.Ints[idx]))
			}
		}
	default:
		switch k {
		case r.Bool:
			fun = func(env *Env) bool {
				env = env.Up(upn)
				return *(*bool)(unsafe.Pointer(&env.Ints[idx]))
			}
		case r.Int:
			fun = func(env *Env) int {
				env = env.Up(upn)
				return *(*int)(unsafe.Pointer(&env.Ints[idx]))
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				env = env.Up(upn)
				return *(*int8)(unsafe.Pointer(&env.Ints[idx]))
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				env = env.Up(upn)
				return *(*int16)(unsafe.Pointer(&env.Ints[idx]))
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				env = env.Up(upn)
				return *(*int32)(unsafe.Pointer(&env.Ints[idx]))
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				env = env.Up(upn)
				return *(*int64)(unsafe.Pointer(&env.Ints[idx]))
			}
		case r.Uint:
			fun = func(env *Env) uint {
				env = env.Up(upn)
				return *(*uint)(unsafe.Pointer(&env.Ints[idx]))
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				env = env.Up(upn)
				return *(*uint8)(unsafe.Pointer(&env.Ints[idx]))
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				env = env.Up(upn)
				return *(*uint16)(unsafe.Pointer(&env.Ints[idx]))
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				env = env.Up(upn)
				return *(*uint32)(unsafe.Pointer(&env.Ints[idx]))
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				env = env.Up(upn)
				return env.Outer.Outer.Ints[idx]
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				env = env.Up(upn)
				return *(*uintptr)(unsafe.Pointer(&env.Ints[idx]))
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				env = env.Up(upn)
				return *(*float32)(unsafe.Pointer(&env.Ints[idx]))
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				env = env.Up(upn)
				return *(*float64)(unsafe.Pointer(&env.Ints[idx]))
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				env = env.Up(upn)
				return *(*complex64)(unsafe.Pointer(&env.Ints[idx]))
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				env = env.Up(upn)
				return *(*complex128)(unsafe.Pointer(&env.Ints[idx]))
			}
		}
	}
	if fun == nil {
		st.Errorf("unsupported variable type, cannot use for optimized read: %s <%v>", sym.Name, sym.Type)
	}
	return &Expr{Lit: Lit{Type: sym.Type}, Fun: fun, Sym: sym}
}
