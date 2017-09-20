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
 * identifier.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
	"unsafe"
)

func (c *Comp) IsCompiledOuter(upn int) bool {
	for ; upn > 0; upn-- {
		for c.UpCost == 0 {
			c = c.Outer
		}
		c = c.Outer
	}
	for c.UpCost == 0 {
		c = c.Outer
	}
	return c.IsCompiled()
}

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
		bind := c.AddBind(name, VarBind, c.TypeOfInterface())
		return &Place{Var: *bind.AsVar(0, PlaceSettable)}
	}
	sym := c.Resolve(name)
	return &Place{Var: *sym.AsVar(opt)}
}

// Symbol compiles a read operation on a constant, variable or function
func (c *Comp) Symbol(sym *Symbol) *Expr {
	switch sym.Desc.Class() {
	case ConstBind:
		return exprLit(sym.Lit, sym)
	case VarBind, FuncBind:
		return c.symbol(sym)
	case IntBind:
		if c.IsCompiledOuter(sym.Upn) {
			return c.symbol(sym)
		} else {
			return c.intSymbol(sym)
		}
	default:
		c.Errorf("unknown symbol class %s", sym.Desc.Class())
	}
	return nil
}

func (c *Comp) symbol(sym *Symbol) *Expr {
	idx := sym.Desc.Index()
	upn := sym.Upn
	kind := sym.Type.Kind()
	var fun I
	switch upn {
	case 0:
		// if package is (partially) compiled, kind can also be one of Bool, Int*, Uint*, Float*, Complex64
		switch kind {
		case r.Bool:
			fun = func(env *Env) bool {
				return env.Binds[idx].Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				return int(env.Binds[idx].Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				return int8(env.Binds[idx].Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				return int16(env.Binds[idx].Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				return int32(env.Binds[idx].Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				return env.Binds[idx].Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				return uint(env.Binds[idx].Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				return uint8(env.Binds[idx].Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				return uint16(env.Binds[idx].Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				return uint32(env.Binds[idx].Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				return env.Binds[idx].Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				return uintptr(env.Binds[idx].Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				return float32(env.Binds[idx].Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				return env.Binds[idx].Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				return complex64(env.Binds[idx].Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				return env.Binds[idx].Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				return env.Binds[idx].String()
			}
		default:
			fun = func(env *Env) r.Value {
				return env.Binds[idx]
			}
		}
	case 1:
		switch kind {
		case r.Bool:
			fun = func(env *Env) bool {
				return env.Outer.Binds[idx].Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				return int(env.Outer.Binds[idx].Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				return int8(env.Outer.Binds[idx].Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				return int16(env.Outer.Binds[idx].Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				return int32(env.Outer.Binds[idx].Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				return env.Outer.Binds[idx].Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				return uint(env.Outer.Binds[idx].Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				return uint8(env.Outer.Binds[idx].Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				return uint16(env.Outer.Binds[idx].Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				return uint32(env.Outer.Binds[idx].Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				return env.Outer.Binds[idx].Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				return uintptr(env.Outer.Binds[idx].Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				return float32(env.Outer.Binds[idx].Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				return env.Outer.Binds[idx].Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				return complex64(env.Outer.Binds[idx].Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				return env.Outer.Binds[idx].Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				return env.Outer.Binds[idx].String()
			}
		default:
			fun = func(env *Env) r.Value {
				return env.Outer.Binds[idx]
			}
		}
	case 2:
		switch kind {
		case r.Bool:
			fun = func(env *Env) bool {
				return env.Outer.Outer.Binds[idx].Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				return int(env.Outer.Outer.Binds[idx].Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				return int8(env.Outer.Outer.Binds[idx].Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				return int16(env.Outer.Outer.Binds[idx].Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				return int32(env.Outer.Outer.Binds[idx].Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				return env.Outer.Outer.Binds[idx].Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				return uint(env.Outer.Outer.Binds[idx].Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				return uint8(env.Outer.Outer.Binds[idx].Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				return uint16(env.Outer.Outer.Binds[idx].Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				return uint32(env.Outer.Outer.Binds[idx].Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				return env.Outer.Outer.Binds[idx].Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				return uintptr(env.Outer.Outer.Binds[idx].Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				return float32(env.Outer.Outer.Binds[idx].Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				return env.Outer.Outer.Binds[idx].Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				return complex64(env.Outer.Outer.Binds[idx].Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				return env.Outer.Outer.Binds[idx].Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				return env.Outer.Outer.Binds[idx].String()
			}
		default:
			fun = func(env *Env) r.Value {
				return env.Outer.Outer.Binds[idx]
			}
		}
	default:
		switch kind {
		case r.Bool:
			fun = func(env *Env) bool {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx].Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return int(env.Outer.Outer.Outer.Binds[idx].Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return int8(env.Outer.Outer.Outer.Binds[idx].Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return int16(env.Outer.Outer.Outer.Binds[idx].Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return int32(env.Outer.Outer.Outer.Binds[idx].Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx].Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return uint(env.Outer.Outer.Outer.Binds[idx].Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return uint8(env.Outer.Outer.Outer.Binds[idx].Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return uint16(env.Outer.Outer.Outer.Binds[idx].Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return uint32(env.Outer.Outer.Outer.Binds[idx].Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx].Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return uintptr(env.Outer.Outer.Outer.Binds[idx].Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return float32(env.Outer.Outer.Outer.Binds[idx].Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx].Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return complex64(env.Outer.Outer.Outer.Binds[idx].Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx].Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx].String()
			}
		default:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx]
			}
		}
	case c.Depth - 1:
		switch kind {
		case r.Bool:
			fun = func(env *Env) bool {
				return env.ThreadGlobals.FileEnv.Binds[idx].Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				return int(env.ThreadGlobals.FileEnv.Binds[idx].Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				return int8(env.ThreadGlobals.FileEnv.Binds[idx].Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				return int16(env.ThreadGlobals.FileEnv.Binds[idx].Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				return int32(env.ThreadGlobals.FileEnv.Binds[idx].Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				return env.ThreadGlobals.FileEnv.Binds[idx].Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				return uint(env.ThreadGlobals.FileEnv.Binds[idx].Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				return uint8(env.ThreadGlobals.FileEnv.Binds[idx].Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				return uint16(env.ThreadGlobals.FileEnv.Binds[idx].Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				return uint32(env.ThreadGlobals.FileEnv.Binds[idx].Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				return env.ThreadGlobals.FileEnv.Binds[idx].Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				return uintptr(env.ThreadGlobals.FileEnv.Binds[idx].Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				return float32(env.ThreadGlobals.FileEnv.Binds[idx].Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				return env.ThreadGlobals.FileEnv.Binds[idx].Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				return complex64(env.ThreadGlobals.FileEnv.Binds[idx].Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				return env.ThreadGlobals.FileEnv.Binds[idx].Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				return env.ThreadGlobals.FileEnv.Binds[idx].String()
			}
		default:
			fun = func(env *Env) r.Value {
				return env.ThreadGlobals.FileEnv.Binds[idx]
			}
		}
	case c.Depth: // TopEnv should not contain variables or functions... but no harm
		switch kind {
		case r.Bool:
			fun = func(env *Env) bool {
				return env.ThreadGlobals.TopEnv.Binds[idx].Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				return int(env.ThreadGlobals.TopEnv.Binds[idx].Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				return int8(env.ThreadGlobals.TopEnv.Binds[idx].Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				return int16(env.ThreadGlobals.TopEnv.Binds[idx].Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				return int32(env.ThreadGlobals.TopEnv.Binds[idx].Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				return env.ThreadGlobals.TopEnv.Binds[idx].Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				return uint(env.ThreadGlobals.TopEnv.Binds[idx].Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				return uint8(env.ThreadGlobals.TopEnv.Binds[idx].Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				return uint16(env.ThreadGlobals.TopEnv.Binds[idx].Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				return uint32(env.ThreadGlobals.TopEnv.Binds[idx].Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				return env.ThreadGlobals.TopEnv.Binds[idx].Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				return uintptr(env.ThreadGlobals.TopEnv.Binds[idx].Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				return float32(env.ThreadGlobals.TopEnv.Binds[idx].Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				return env.ThreadGlobals.TopEnv.Binds[idx].Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				return complex64(env.ThreadGlobals.TopEnv.Binds[idx].Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				return env.ThreadGlobals.TopEnv.Binds[idx].Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				return env.ThreadGlobals.TopEnv.Binds[idx].String()
			}
		default:
			fun = func(env *Env) r.Value {
				return env.ThreadGlobals.TopEnv.Binds[idx]
			}
		}
	}
	return &Expr{Lit: Lit{Type: sym.Type}, Fun: fun, Sym: sym}
}

func (c *Comp) intSymbol(sym *Symbol) *Expr {
	k := sym.Type.Kind()
	idx := sym.Desc.Index()
	upn := sym.Upn
	var fun I
	switch upn {
	case 0:
		switch k {
		case r.Bool:
			fun = func(env *Env) bool {
				return *(*bool)(unsafe.Pointer(&env.IntBinds[idx]))
			}
		case r.Int:
			fun = func(env *Env) int {
				return *(*int)(unsafe.Pointer(&env.IntBinds[idx]))
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				return *(*int8)(unsafe.Pointer(&env.IntBinds[idx]))
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				return *(*int16)(unsafe.Pointer(&env.IntBinds[idx]))
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				return *(*int32)(unsafe.Pointer(&env.IntBinds[idx]))
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				return *(*int64)(unsafe.Pointer(&env.IntBinds[idx]))
			}
		case r.Uint:
			fun = func(env *Env) uint {
				return *(*uint)(unsafe.Pointer(&env.IntBinds[idx]))
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				return *(*uint8)(unsafe.Pointer(&env.IntBinds[idx]))
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				return *(*uint16)(unsafe.Pointer(&env.IntBinds[idx]))
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				return *(*uint32)(unsafe.Pointer(&env.IntBinds[idx]))
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				return env.IntBinds[idx]
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				return *(*uintptr)(unsafe.Pointer(&env.IntBinds[idx]))
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				return *(*float32)(unsafe.Pointer(&env.IntBinds[idx]))
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				return *(*float64)(unsafe.Pointer(&env.IntBinds[idx]))
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				return *(*complex64)(unsafe.Pointer(&env.IntBinds[idx]))
			}
		default:
			c.Errorf("unsupported symbol type, cannot use for optimized read: %s %s <%v>", sym.Desc.Class(), sym.Name, sym.Type)
			return nil
		}
	case 1:
		switch k {
		case r.Bool:
			fun = func(env *Env) bool {
				return *(*bool)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			}
		case r.Int:
			fun = func(env *Env) int {
				return *(*int)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				return *(*int8)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				return *(*int16)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				return *(*int32)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				return *(*int64)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			}
		case r.Uint:
			fun = func(env *Env) uint {
				return *(*uint)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				return *(*uint8)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				return *(*uint16)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				return *(*uint32)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				return env.Outer.IntBinds[idx]
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				return *(*uintptr)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				return *(*float32)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				return *(*float64)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				return *(*complex64)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			}
		default:
			c.Errorf("unsupported variable type, cannot use for optimized read: %s <%v>", sym.Name, sym.Type)
			return nil
		}
	case 2:
		switch k {
		case r.Bool:
			fun = func(env *Env) bool {
				return *(*bool)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Int:
			fun = func(env *Env) int {
				return *(*int)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				return *(*int8)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				return *(*int16)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				return *(*int32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				return *(*int64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Uint:
			fun = func(env *Env) uint {
				return *(*uint)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				return *(*uint8)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				return *(*uint16)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				return *(*uint32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				return env.Outer.Outer.IntBinds[idx]
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				return *(*uintptr)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				return *(*float32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				return *(*float64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				return *(*complex64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		default:
			c.Errorf("unsupported variable type, cannot use for optimized read: %s <%v>", sym.Name, sym.Type)
			return nil
		}
	default:
		switch k {
		case r.Bool:
			fun = func(env *Env) bool {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*bool)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Int:
			fun = func(env *Env) int {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*int)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*int8)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*int16)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*int32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*int64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Uint:
			fun = func(env *Env) uint {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*uint)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*uint8)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*uint16)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*uint32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.IntBinds[idx]
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*uintptr)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*float32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*float64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*complex64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			}
		default:
			c.Errorf("unsupported variable type, cannot use for optimized read: %s <%v>", sym.Name, sym.Type)
			return nil
		}
	}
	return &Expr{Lit: Lit{Type: sym.Type}, Fun: fun, Sym: sym}
}
