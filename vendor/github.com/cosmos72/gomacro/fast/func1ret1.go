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
 * func_ret1.go
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

func (c *Comp) func1ret1(t xr.Type, m *funcMaker) func(*Env) r.Value {

	nbinds := m.nbinds
	nintbinds := m.nintbinds
	funcbody := m.funcbody

	targ0 := t.In(0)
	karg0 := targ0.Kind()
	kret0 := t.Out(0).Kind()

	indexes := [2]int{
		m.parambinds[0].Desc.Index(),
		m.resultbinds[0].Desc.Index(),
	}
	switch karg0 {
	case r.Bool:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 bool) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*bool)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.Int:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.Int8:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int8) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.Int16:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int16) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.Int32:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int32) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.Int64:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 int64) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*int64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.Uint:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.Uint8:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint8) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.Uint16:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint16) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.Uint32:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint32) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.Uint64:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uint64) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)

						env.IntBinds[indexes[0]] = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.Uintptr:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 uintptr) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.Float32:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float32) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float32)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.Float64:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 float64) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*float64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.Complex64:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex64) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)

						*(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[0]])) = arg0

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.Complex128:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 complex128) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfComplex128).Elem()
							place.SetComplex(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	case r.String:
		switch kret0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 bool) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 bool) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*bool)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 int) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 int) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*int)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 int8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 int8) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*int8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 int16) { return },
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 int16) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*int16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 int32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 int32) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*int32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 int64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 int64) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*int64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 uint) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 uint) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*uint)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 uint8) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 uint8) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*uint8)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 uint16) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 uint16) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*uint16)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 uint32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 uint32) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*uint32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 uint64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 uint64) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = env.IntBinds[indexes[1]]

						env.FreeEnv()
						return

					})
				}
			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 uintptr) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 uintptr) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*uintptr)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 float32) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 float32) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*float32)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 float64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 float64) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*float64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 complex64) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 complex64) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = *(*complex64)(unsafe.Pointer(&env.IntBinds[indexes[1]]))

						env.FreeEnv()
						return

					})
				}
			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 complex128) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 complex128) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = env.Binds[indexes[1]].Complex()

						env.FreeEnv()
						return

					})
				}
			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) (ret0 string) {
							return
						},
						)
					}
				}
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.ValueOf(func(arg0 string) (ret0 string) {
						env := newEnv4Func(env, nbinds, nintbinds)
						{
							place := r.New(TypeOfString).Elem()
							place.SetString(arg0,
							)
							env.Binds[indexes[0]] = place
						}

						funcbody(env)

						ret0 = env.Binds[indexes[1]].String()

						env.FreeEnv()
						return

					})
				}
			}
		}
	}
	return nil
}
