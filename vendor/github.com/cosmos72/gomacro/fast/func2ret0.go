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

	indexes := &[2]int{
		m.Param[0].Desc.Index(),
		m.Param[1].Desc.Index(),
	}
	var debugC *Comp
	if c.Globals.Options&OptDebugger != 0 {
		debugC = c
	}

	var ret func(*Env) r.Value
	switch karg0 {
	case r.Bool:
		ret = func2ret0Bool(m, indexes, karg1, debugC)

	case r.Int:
		ret = func2ret0Int(m, indexes, karg1, debugC)

	case r.Int8:
		ret = func2ret0Int8(m, indexes, karg1, debugC)

	case r.Int16:
		ret = func2ret0Int16(m, indexes, karg1, debugC)

	case r.Int32:
		ret = func2ret0Int32(m, indexes, karg1, debugC)

	case r.Int64:
		ret = func2ret0Int64(m, indexes, karg1, debugC)

	case r.Uint:
		ret = func2ret0Uint(m, indexes, karg1, debugC)

	case r.Uint8:
		ret = func2ret0Uint8(m, indexes, karg1, debugC)

	case r.Uint16:
		ret = func2ret0Uint16(m, indexes, karg1, debugC)

	case r.Uint32:
		ret = func2ret0Uint32(m, indexes, karg1, debugC)

	case r.Uint64:
		ret = func2ret0Uint64(m, indexes, karg1, debugC)

	case r.Uintptr:
		ret = func2ret0Uintptr(m, indexes, karg1, debugC)

	case r.Float32:
		ret = func2ret0Float32(m, indexes, karg1, debugC)

	case r.Float64:
		ret = func2ret0Float64(m, indexes, karg1, debugC)

	case r.Complex64:
		ret = func2ret0Complex64(m, indexes, karg1, debugC)

	case r.Complex128:
		ret = func2ret0Complex128(m, indexes, karg1, debugC)

	case r.String:
		ret = func2ret0String(m, indexes, karg1, debugC)

	}
	return ret
}
func func2ret0Bool(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(bool,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Int(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Int8(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int8,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Int16(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int16,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Int32(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int32,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Int64(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(int64,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Uint(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Uint8(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint8,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Uint16(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint16,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Uint32(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint32,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Uint64(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uint64,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					env.Ints[indexes[0]] = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Uintptr(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(uintptr,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Float32(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float32,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Float64(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(float64,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Complex64(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex64,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Complex128(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(complex128,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0String(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					bool,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 bool,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					int,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 int,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					int8,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 int8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					int16,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 int16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					int32,

				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 int32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					int64,
				) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 int64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					uint) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 uint,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					uint8) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 uint8,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					uint16) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 uint16,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					uint32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 uint32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					uint64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 uint64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					uintptr) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 uintptr,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					float32) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 float32,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					float64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 float64,

				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					complex64) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 complex64,
				) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					complex128) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*complex128)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				funv := r.ValueOf(func(string,

					string) {})
				ret = func(env *Env) r.Value { return funv }

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
