// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

// +build !gomacro_fast_compact

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
 * call0ret1.go
 *
 *  Created on Apr 20, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
	"github.com/cosmos72/gomacro/base"
)

func (c *Comp) call0ret1(call *Call, maxdepth int) I {
	expr := call.Fun
	exprfun := expr.AsX1()
	funsym := expr.Sym
	funupn, funindex := -1, -1
	if funsym != nil {
		funupn = funsym.Upn
		funindex = funsym.Desc.Index()
		if funindex == NoIndex {
			base.Errorf("internal error: call0ret1() invoked for constant function %#v. use call_builtin() instead", expr)
		}

	}
	tret := expr.Type.Out(0)
	kret := tret.Kind()
	var cachedfunv r.Value
	var ret I
	switch kret {
	case r.Bool:

		{
			if !tret.IdenticalTo(c.TypeOfBool(),
			) {
				ret = func(env *Env) bool {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return ret.Bool()

				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() bool

					ret = func(env *Env) bool {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() bool)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) bool {
						fun := env.Binds[funindex].Interface().(func() bool)
						return fun()
					}
				case 1:
					ret = func(env *Env) bool {
						fun := env.Outer.Binds[funindex].Interface().(func() bool)
						return fun()
					}
				case 2:
					ret = func(env *Env) bool {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() bool)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) bool {
					fun := exprfun(env).Interface().(func() bool)
					return fun()
				}
			}

		}
	case r.Int:

		{
			if !tret.IdenticalTo(c.TypeOfInt(),
			) {
				ret = func(env *Env) int {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return int(ret.Int())
				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() int

					ret = func(env *Env) int {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) int {
						fun := env.Binds[funindex].Interface().(func() int)
						return fun()
					}
				case 1:
					ret = func(env *Env) int {
						fun := env.Outer.Binds[funindex].Interface().(func() int)
						return fun()
					}
				case 2:
					ret = func(env *Env) int {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() int)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) int {
					fun := exprfun(env).Interface().(func() int)
					return fun()
				}
			}

		}
	case r.Int8:

		{
			if !tret.IdenticalTo(c.TypeOfInt8(),
			) {
				ret = func(env *Env) int8 {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return int8(ret.Int())
				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() int8

					ret = func(env *Env) int8 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int8)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) int8 {
						fun := env.Binds[funindex].Interface().(func() int8)
						return fun()
					}
				case 1:
					ret = func(env *Env) int8 {
						fun := env.Outer.Binds[funindex].Interface().(func() int8)
						return fun()
					}
				case 2:
					ret = func(env *Env) int8 {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() int8)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) int8 {
					fun := exprfun(env).Interface().(func() int8)
					return fun()
				}
			}

		}
	case r.Int16:
		{
			if !tret.IdenticalTo(c.TypeOfInt16(),
			) {
				ret = func(env *Env) int16 {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return int16(ret.Int())
				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() int16

					ret = func(env *Env) int16 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int16)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) int16 {
						fun := env.Binds[funindex].Interface().(func() int16)
						return fun()
					}
				case 1:
					ret = func(env *Env) int16 {
						fun := env.Outer.Binds[funindex].Interface().(func() int16)
						return fun()
					}
				case 2:
					ret = func(env *Env) int16 {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() int16)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) int16 {
					fun := exprfun(env).Interface().(func() int16)
					return fun()
				}
			}

		}
	case r.Int32:
		{
			if !tret.IdenticalTo(c.TypeOfInt32(),
			) {
				ret = func(env *Env) int32 {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return int32(ret.Int())
				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() int32

					ret = func(env *Env) int32 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int32)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) int32 {
						fun := env.Binds[funindex].Interface().(func() int32)
						return fun()
					}
				case 1:
					ret = func(env *Env) int32 {
						fun := env.Outer.Binds[funindex].Interface().(func() int32)
						return fun()
					}
				case 2:
					ret = func(env *Env) int32 {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() int32)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) int32 {
					fun := exprfun(env).Interface().(func() int32)
					return fun()
				}
			}

		}
	case r.Int64:
		{
			if !tret.IdenticalTo(c.TypeOfInt64(),
			) {
				ret = func(env *Env) int64 {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return ret.Int()

				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() int64

					ret = func(env *Env) int64 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int64)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) int64 {
						fun := env.Binds[funindex].Interface().(func() int64)
						return fun()
					}
				case 1:
					ret = func(env *Env) int64 {
						fun := env.Outer.Binds[funindex].Interface().(func() int64)
						return fun()
					}
				case 2:
					ret = func(env *Env) int64 {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() int64)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) int64 {
					fun := exprfun(env).Interface().(func() int64)
					return fun()
				}
			}

		}
	case r.Uint:
		{
			if !tret.IdenticalTo(c.TypeOfUint(),
			) {
				ret = func(env *Env) uint {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return uint(ret.Uint())
				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() uint

					ret = func(env *Env) uint {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) uint {
						fun := env.Binds[funindex].Interface().(func() uint)
						return fun()
					}
				case 1:
					ret = func(env *Env) uint {
						fun := env.Outer.Binds[funindex].Interface().(func() uint)
						return fun()
					}
				case 2:
					ret = func(env *Env) uint {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() uint)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) uint {
					fun := exprfun(env).Interface().(func() uint)
					return fun()
				}
			}

		}
	case r.Uint8:
		{
			if !tret.IdenticalTo(c.TypeOfUint8(),
			) {
				ret = func(env *Env) uint8 {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return uint8(ret.Uint())
				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() uint8

					ret = func(env *Env) uint8 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint8)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) uint8 {
						fun := env.Binds[funindex].Interface().(func() uint8)
						return fun()
					}
				case 1:
					ret = func(env *Env) uint8 {
						fun := env.Outer.Binds[funindex].Interface().(func() uint8)
						return fun()
					}
				case 2:
					ret = func(env *Env) uint8 {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() uint8)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) uint8 {
					fun := exprfun(env).Interface().(func() uint8)
					return fun()
				}
			}

		}
	case r.Uint16:
		{
			if !tret.IdenticalTo(c.TypeOfUint16(),
			) {
				ret = func(env *Env) uint16 {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return uint16(ret.Uint())
				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() uint16

					ret = func(env *Env) uint16 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint16)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) uint16 {
						fun := env.Binds[funindex].Interface().(func() uint16)
						return fun()
					}
				case 1:
					ret = func(env *Env) uint16 {
						fun := env.Outer.Binds[funindex].Interface().(func() uint16)
						return fun()
					}
				case 2:
					ret = func(env *Env) uint16 {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() uint16)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) uint16 {
					fun := exprfun(env).Interface().(func() uint16)
					return fun()
				}
			}

		}
	case r.Uint32:
		{
			if !tret.IdenticalTo(c.TypeOfUint32(),
			) {
				ret = func(env *Env) uint32 {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return uint32(ret.Uint())
				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() uint32

					ret = func(env *Env) uint32 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint32)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) uint32 {
						fun := env.Binds[funindex].Interface().(func() uint32)
						return fun()
					}
				case 1:
					ret = func(env *Env) uint32 {
						fun := env.Outer.Binds[funindex].Interface().(func() uint32)
						return fun()
					}
				case 2:
					ret = func(env *Env) uint32 {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() uint32)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) uint32 {
					fun := exprfun(env).Interface().(func() uint32)
					return fun()
				}
			}

		}
	case r.Uint64:
		{
			if !tret.IdenticalTo(c.TypeOfUint64(),
			) {
				ret = func(env *Env) uint64 {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return ret.Uint()

				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() uint64

					ret = func(env *Env) uint64 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint64)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) uint64 {
						fun := env.Binds[funindex].Interface().(func() uint64)
						return fun()
					}
				case 1:
					ret = func(env *Env) uint64 {
						fun := env.Outer.Binds[funindex].Interface().(func() uint64)
						return fun()
					}
				case 2:
					ret = func(env *Env) uint64 {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() uint64)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) uint64 {
					fun := exprfun(env).Interface().(func() uint64)
					return fun()
				}
			}

		}
	case r.Uintptr:
		{
			if !tret.IdenticalTo(c.TypeOfUintptr(),
			) {
				ret = func(env *Env) uintptr {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return uintptr(ret.Uint())
				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() uintptr

					ret = func(env *Env) uintptr {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uintptr)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) uintptr {
						fun := env.Binds[funindex].Interface().(func() uintptr)
						return fun()
					}
				case 1:
					ret = func(env *Env) uintptr {
						fun := env.Outer.Binds[funindex].Interface().(func() uintptr)
						return fun()
					}
				case 2:
					ret = func(env *Env) uintptr {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() uintptr)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) uintptr {
					fun := exprfun(env).Interface().(func() uintptr)
					return fun()
				}
			}

		}
	case r.Float32:
		{
			if !tret.IdenticalTo(c.TypeOfFloat32(),
			) {
				ret = func(env *Env) float32 {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return float32(ret.Float())
				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() float32

					ret = func(env *Env) float32 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() float32)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) float32 {
						fun := env.Binds[funindex].Interface().(func() float32)
						return fun()
					}
				case 1:
					ret = func(env *Env) float32 {
						fun := env.Outer.Binds[funindex].Interface().(func() float32)
						return fun()
					}
				case 2:
					ret = func(env *Env) float32 {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() float32)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) float32 {
					fun := exprfun(env).Interface().(func() float32)
					return fun()
				}
			}

		}
	case r.Float64:
		{
			if !tret.IdenticalTo(c.TypeOfFloat64(),
			) {
				ret = func(env *Env) float64 {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return ret.Float()

				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() float64

					ret = func(env *Env) float64 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() float64)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) float64 {
						fun := env.Binds[funindex].Interface().(func() float64)
						return fun()
					}
				case 1:
					ret = func(env *Env) float64 {
						fun := env.Outer.Binds[funindex].Interface().(func() float64)
						return fun()
					}
				case 2:
					ret = func(env *Env) float64 {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() float64)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func() float64)
					return fun()
				}
			}

		}
	case r.Complex64:
		{
			if !tret.IdenticalTo(c.TypeOfComplex64(),
			) {
				ret = func(env *Env) complex64 {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return complex64(ret.Complex())
				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() complex64

					ret = func(env *Env) complex64 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() complex64)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) complex64 {
						fun := env.Binds[funindex].Interface().(func() complex64)
						return fun()
					}
				case 1:
					ret = func(env *Env) complex64 {
						fun := env.Outer.Binds[funindex].Interface().(func() complex64)
						return fun()
					}
				case 2:
					ret = func(env *Env) complex64 {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() complex64)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) complex64 {
					fun := exprfun(env).Interface().(func() complex64)
					return fun()
				}
			}

		}
	case r.Complex128:
		{
			if !tret.IdenticalTo(c.TypeOfComplex128(),
			) {
				ret = func(env *Env) complex128 {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return ret.Complex()

				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() complex128
					ret = func(env *Env) complex128 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() complex128)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) complex128 {
						fun := env.Binds[funindex].Interface().(func() complex128)
						return fun()
					}
				case 1:
					ret = func(env *Env) complex128 {
						fun := env.Outer.Binds[funindex].Interface().(func() complex128)
						return fun()
					}
				case 2:
					ret = func(env *Env) complex128 {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() complex128)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func() complex128)
					return fun()
				}
			}

		}
	case r.String:
		{
			if !tret.IdenticalTo(c.TypeOfString(),
			) {
				ret = func(env *Env) string {
					fun := exprfun(env)

					ret := fun.Call(base.ZeroValues)[0]
					return ret.String()

				}
			} else if funsym != nil {
				switch funupn {
				case maxdepth - 1:
					var cachedfun func() string
					ret = func(env *Env) string {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() string)
						}
						return cachedfun()
					}
				case 0:
					ret = func(env *Env) string {
						fun := env.Binds[funindex].Interface().(func() string)
						return fun()
					}
				case 1:
					ret = func(env *Env) string {
						fun := env.Outer.Binds[funindex].Interface().(func() string)
						return fun()
					}
				case 2:
					ret = func(env *Env) string {
						fun := env.Outer.Outer.Binds[funindex].Interface().(func() string)
						return fun()
					}
				}
			}

			if ret == nil {
				ret = func(env *Env) string {
					fun := exprfun(env).Interface().(func() string)
					return fun()
				}
			}

		}
	default:
		ret = func(env *Env) r.Value {
			funv := exprfun(env)
			return funv.Call(base.ZeroValues)[0]
		}

	}
	return ret
}
