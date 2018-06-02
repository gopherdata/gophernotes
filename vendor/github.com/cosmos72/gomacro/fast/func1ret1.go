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
	var debugC *Comp
	if c.Globals.Options&OptDebugger != 0 {
		debugC = c
	}

	karg0 := t.In(0).Kind()
	kret0 := t.Out(0).Kind()

	indexes := &[2]int{
		m.Param[0].Desc.Index(),
		m.Result[0].Desc.Index(),
	}
	var ret func(*Env) r.Value
	switch karg0 {
	case r.Bool:
		ret = func1ret1Bool(m, indexes, kret0, debugC)
	case r.Int:
		ret = func1ret1Int(m, indexes, kret0, debugC)
	case r.Int8:
		ret = func1ret1Int8(m, indexes, kret0, debugC)

	case r.Int16:
		ret = func1ret1Int16(m, indexes, kret0, debugC)

	case r.Int32:
		ret = func1ret1Int32(m, indexes, kret0, debugC)

	case r.Int64:
		ret = func1ret1Int64(m, indexes, kret0, debugC)

	case r.Uint:
		ret = func1ret1Uint(m, indexes, kret0, debugC)

	case r.Uint8:
		ret = func1ret1Uint8(m, indexes, kret0, debugC)

	case r.Uint16:
		ret = func1ret1Uint16(m, indexes, kret0, debugC)

	case r.Uint32:
		ret = func1ret1Uint32(m, indexes, kret0, debugC)

	case r.Uint64:
		ret = func1ret1Uint64(m, indexes, kret0, debugC)

	case r.Uintptr:
		ret = func1ret1Uintptr(m, indexes, kret0, debugC)

	case r.Float32:
		ret = func1ret1Float32(m, indexes, kret0, debugC)

	case r.Float64:
		ret = func1ret1Float64(m, indexes, kret0, debugC)

	case r.Complex64:
		ret = func1ret1Complex64(m, indexes, kret0, debugC)

	case r.Complex128:
		ret = func1ret1Complex128(m, indexes, kret0, debugC)

	case r.String:
		ret = func1ret1String(m, indexes, kret0, debugC)

	}
	return ret
}
func func1ret1Bool(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1Int(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1Int8(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1Int16(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1Int32(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1Int64(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1Uint(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1Uint8(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1Uint16(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1Uint32(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1Uint64(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1Uintptr(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1Float32(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1Float64(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1Complex64(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1Complex128(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
func func1ret1String(m *funcMaker, indexes *[2]int, kret0 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch kret0 {
	case r.Bool:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 bool,

				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = *(*bool)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 int,
				) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = *(*int)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 int8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = *(*int8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 int16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = *(*int16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 int32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = *(*int32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 int64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = *(*int64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 uint) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = *(*uint)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 uint8) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = *(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 uint16) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = *(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 uint32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = *(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 uint64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = env.Ints[indexes[1]]

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 uintptr) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = *(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 float32) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 float32,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = *(*float32)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 float64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = *(*float64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 complex64) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = *(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 complex128) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = *(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]]))

					env.freeEnv4Func()
					return

				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

				) (ret0 string) { return },
				)
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

				) (ret0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					funcbody(env)

					ret0 = env.Vals[indexes[1]].String()

					env.freeEnv4Func()
					return

				})
			}
		}
	}
	return ret
}
