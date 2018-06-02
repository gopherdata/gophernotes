// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

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
 * func0ret1.go
 *
 *  Created on Apr 16, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) func0ret1(t xr.Type, m *funcMaker) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody

	var debugC *Comp
	if c.Globals.Options&OptDebugger != 0 {
		debugC = c
	}

	tret0 := t.Out(0)
	kret0 := tret0.Kind()
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 bool) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) bool)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 int) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) int)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 int8) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) int8)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 int16) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) int16)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 int32) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) int32)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 int32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 int64) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) int64)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 uint) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) uint)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 uint8) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) uint8)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 uint16) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) uint16)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 uint32) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) uint32)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 uint64) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) uint64)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 uintptr) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) uintptr)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 float32) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) float32)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 float64) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) float64)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 complex64) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) complex64)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 complex128) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) complex128)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.ValueOf(func() (ret0 string) { return },
					)
				}
			}

			resultfun := m.resultfun[0].(func(*Env) string)
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					funcbody(env)

					ret0 = resultfun(env)
					env.freeEnv4Func()
					return

				})
			}
		}
	default:
		return nil
	}
}
