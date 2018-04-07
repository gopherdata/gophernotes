// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

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
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *
 * func2ret0.go
 *
 *  Created on Apr 16, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
	"unsafe"

	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) func2ret0(t xr.Type, m *funcMaker) func(*Env) r.Value {
	karg0 := t.In(0).Kind()
	karg1 := t.In(1).Kind()

	if !IsOptimizedKind(karg0) || !IsOptimizedKind(karg1) {
		return nil
	}

	indexes := [2]int{
		m.parambinds[0].Desc.Index(),
		m.parambinds[1].Desc.Index(),
	}
	nbinds := m.nbinds
	nintbinds := m.nintbinds
	funcbody := m.funcbody
	{
		argdecls := [2]func(*Env, r.Value){nil, nil}
		for i, bind := range m.parambinds {
			argdecls[i] = c.DeclBindRuntimeValue(bind)
			if argdecls[i] == nil {
				argdecls[i] = declBindRuntimeValueNop
			}

		}
		switch karg0 {
		case r.Bool:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool,
								uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool,

								uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool,

								uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool,

								uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool,

								float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool,

								float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool,

								complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool,

								complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(bool,

								string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool,

							arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}
		case r.Int:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int,
								uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int,

								uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int,

								uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int,

								float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int,

								float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int,

								complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int,

								complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int,

								string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int,

							arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}
		case r.Int8:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8, uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8,
								uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8,

								uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8,

								float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8,

								float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8,

								complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8,

								complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int8,

								string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8,

							arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}
		case r.Int16:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16, uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16, uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16,
								uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16,

								float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16,

								float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16,

								complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16,

								complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int16,

								string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16,

							arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}
		case r.Int32:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32, uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32, uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32, uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32,
								float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32,

								float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32,

								complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32,

								complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int32,

								string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32,

							arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}
		case r.Int64:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64, uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64, uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64, uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64, float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64,
								float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64,

								complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64,

								complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(int64,

								string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64,

							arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}
		case r.Uint:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint, uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint, uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint, uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint, float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint, float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint,
								complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint,

								complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint,

								string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint,
							arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}
		case r.Uint8:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8, uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8, uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8, uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8, float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8, float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8, complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8,
								complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint8,

								string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8, arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}
		case r.Uint16:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16, complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint16,
								string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16, arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}
		case r.Uint32:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint32, string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32, arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}
		case r.Uint64:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uint64, string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64, arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)

							env.IntBinds[indexes[0]] = arg0

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}
		case r.Uintptr:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(uintptr, string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr, arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}

		case r.Float32:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float32, string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32, arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}

		case r.Float64:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(float64, string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64, arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}

		case r.Complex64:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex64, string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64, arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}

		case r.Complex128:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(complex128, string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128, arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}

		case r.String:
			switch karg1 {
			case r.Bool:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, bool) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 bool) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, int) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 int) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int8:

				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, int8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 int8) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, int16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 int16) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, int32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 int32) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Int64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, int64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 int64) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, uint) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 uint) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint8:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, uint8) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 uint8) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint16:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, uint16) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 uint16) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, uint32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 uint32) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uint64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, uint64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 uint64) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							env.IntBinds[indexes[1]] = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Uintptr:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, uintptr) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 uintptr) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float32:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, float32) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 float32) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Float64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, float64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 float64) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex64:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, complex64) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 complex64) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}

							*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]])) = arg1

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.Complex128:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, complex128) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 complex128) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg1,
								)
								env.Binds[indexes[1]] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			case r.String:
				{
					if funcbody == nil {
						return func(env *Env) r.Value {
							return r.ValueOf(func(string, string) {
							})
						}
					}
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string, arg1 string) {
							env := newEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[indexes[0]] = place
							}
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg1,
								)
								env.Binds[indexes[1]] = place
							}
							funcbody(env)

							env.FreeEnv()
						})
					}
				}
			}

		}
	}
	return nil
}
