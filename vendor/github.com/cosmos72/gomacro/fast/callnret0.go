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
 * callnret0.go
 *
 *  Created on Apr 15, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) call0ret0(call *Call, maxdepth int) func(env *Env) {
	expr := call.Fun
	funsym := expr.Sym
	if funsym == nil {
		exprfun := expr.AsX1()
		return func(env *Env) {
			fun := exprfun(env).Interface().(func())
			fun()
		}
	}

	var cachedfunv r.Value
	var cachedfun func()

	funupn := funsym.Upn
	funindex := funsym.Desc.Index()
	switch funupn {
	case maxdepth - 1:
		return func(env *Env) {
			funv := env.FileEnv.Vals[funindex]
			if cachedfunv != funv {
				cachedfunv = funv
				cachedfun = funv.Interface().(func())
			}
			cachedfun()
		}
	case 0:
		return func(env *Env) {
			fun := env.Vals[funindex].Interface().(func())
			fun()
		}
	case 1:
		return func(env *Env) {
			fun := env.Outer.Vals[funindex].Interface().(func())
			fun()
		}
	case 2:
		return func(env *Env) {
			fun := env.Outer.Outer.Vals[funindex].Interface().(func())
			fun()
		}
	default:
		return func(env *Env) {
			env = env.Outer.Outer.Outer.Outer
			for i := 3; i < funupn; i++ {
				env = env.Outer
			}

			fun := env.Vals[funindex].Interface().(func())
			fun()
		}
	}
}
func (c *Comp) call1ret0(call *Call, maxdepth int) func(env *Env) {
	expr := call.Fun
	exprfun := expr.AsX1()
	funsym := expr.Sym
	funupn, funindex := -1, -1
	if funsym != nil {
		funupn = funsym.Upn
		funindex = funsym.Desc.Index()
		if funindex == NoIndex {
			Errorf("internal error: call1ret0() invoked for constant function %#v. use call_builtin() instead", expr)
		}

	}
	arg := call.Args[0]
	argfun := call.MakeArgfunsX1()[0]

	var cachedfunv r.Value
	var ret func(env *Env)

	t := expr.Type.In(0)
	k := t.Kind()
	if KindToType(k) == t.ReflectType() {
		switch k {
		case r.Bool:

			if arg.Const() {
				argconst := r.ValueOf(arg.Value).Bool()
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(bool,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								bool))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(bool,

						))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) bool)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(bool,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								bool))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(bool,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.Int:

			if arg.Const() {
				argconst := int(r.ValueOf(arg.Value).Int())
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								int))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(int,

						))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) int)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								int))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(int,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.Int8:
			if arg.Const() {
				argconst := int8(r.ValueOf(arg.Value).Int())
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int8,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								int8))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(int8,

						))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) int8)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int8,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								int8))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(int8,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.Int16:
			if arg.Const() {
				argconst := int16(r.ValueOf(arg.Value).Int())
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int16,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								int16))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(int16,

						))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) int16)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int16,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								int16))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(int16,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.Int32:
			if arg.Const() {
				argconst := int32(r.ValueOf(arg.Value).Int())
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int32,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								int32))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(int32,

						))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) int32)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int32,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								int32))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(int32,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.Int64:
			if arg.Const() {
				argconst := r.ValueOf(arg.Value).Int()
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int64,
					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								int64))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(int64,

						))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) int64)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int64,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								int64))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(int64,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.Uint:
			if arg.Const() {
				argconst :=

					uint(r.ValueOf(arg.Value).Uint())
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(uint)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint,

						))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) uint)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(uint,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.Uint8:
			if arg.Const() {
				argconst :=

					uint8(r.ValueOf(arg.Value).Uint())
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(
						uint8)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint8))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint8,

						))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) uint8)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(uint8,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint8))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint8,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.Uint16:
			if arg.Const() {
				argconst :=

					uint16(r.ValueOf(arg.Value).Uint())
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						uint16)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint16))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint16,

						))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) uint16)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(uint16,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint16))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint16,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.Uint32:
			if arg.Const() {
				argconst :=

					uint32(r.ValueOf(arg.Value).Uint())
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						uint32)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint32))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint32,

						))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) uint32)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(uint32,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint32))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint32,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.Uint64:
			if arg.Const() {
				argconst := r.ValueOf(arg.Value).Uint()
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						uint64)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint64))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint64,

						))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) uint64)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(uint64,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint64))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint64,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.Uintptr:
			if arg.Const() {
				argconst :=

					uintptr(r.ValueOf(arg.Value).Uint())
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						uintptr)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uintptr))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uintptr,

						))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) uintptr)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(uintptr,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uintptr))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uintptr,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.Float32:
			if arg.Const() {
				argconst :=

					float32(r.ValueOf(arg.Value).Float())
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						float32)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								float32))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(float32,

						))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) float32)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(float32,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								float32))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(float32,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.Float64:
			if arg.Const() {
				argconst := r.ValueOf(arg.Value).Float()
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						float64)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								float64))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(float64,

						))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) float64)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(float64,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								float64))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(float64,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.Complex64:
			if arg.Const() {
				argconst :=

					complex64(r.ValueOf(arg.Value).Complex())
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						complex64)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								complex64))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(complex64,
						))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) complex64)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(complex64,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								complex64))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(complex64,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.Complex128:
			if arg.Const() {
				argconst := r.ValueOf(arg.Value).Complex()
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						complex128)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								complex128))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(complex128))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) complex128)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(complex128,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								complex128))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(complex128,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		case r.String:
			if arg.Const() {
				argconst := r.ValueOf(arg.Value).String()
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						string)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								string))
						}

						cachedfun(argconst)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(
							string))

						fun(argconst)
					}
				}

			} else {
				argfun := arg.Fun.(func(env *Env) string)
				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(string,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								string))
						}

						arg := argfun(env)

						cachedfun(arg)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(string,

						))
						arg := argfun(env)

						fun(arg)
					}
				}

			}
		}
	}

	if ret == nil {
		ret = func(env *Env) {
			funv := exprfun(env)

			argv := []r.Value{
				argfun(env),
			}
			callxr(funv, argv)
		}
	}
	return ret
}
func (c *Comp) call2ret0(call *Call, maxdepth int) func(env *Env) {
	expr := call.Fun
	exprfun := expr.AsX1()
	funsym := expr.Sym
	funupn, funindex := -1, -1
	if funsym != nil {
		funupn = funsym.Upn
		funindex = funsym.Desc.Index()
		if funindex == NoIndex {
			Errorf("internal error: call2ret0() invoked for constant function %#v. use call_builtin() instead", expr)
		}

	}
	args := call.Args
	argfunsX1 := call.MakeArgfunsX1()
	argfuns := [2]func(*Env) r.Value{
		argfunsX1[0],
		argfunsX1[1],
	}
	var cachedfunv r.Value
	var ret func(env *Env)

	t := expr.Type.In(0)
	rt := t.ReflectType()
	k := t.Kind()
	if KindToType(k) == rt && expr.Type.In(1).ReflectType() == rt {
		switch k {
		case r.Bool:
			{
				arg0fun := args[0].WithFun().(func(*Env) bool)
				arg1fun := args[1].WithFun().(func(*Env) bool)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(bool, bool,

					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								bool, bool))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(bool, bool,

						))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Int:
			{
				arg0fun := args[0].WithFun().(func(*Env) int)
				arg1fun := args[1].WithFun().(func(*Env) int)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int, int,
					)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								int, int))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(int, int,

						))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Int8:
			{
				arg0fun := args[0].WithFun().(func(*Env) int8)
				arg1fun := args[1].WithFun().(func(*Env) int8)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int8, int8)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								int8, int8))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(int8, int8,

						))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Int16:
			{
				arg0fun := args[0].WithFun().(func(*Env) int16)
				arg1fun := args[1].WithFun().(func(*Env) int16)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(
						int16, int16)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								int16, int16))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(int16, int16,

						))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Int32:
			{
				arg0fun := args[0].WithFun().(func(*Env) int32)
				arg1fun := args[1].WithFun().(func(*Env) int32)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						int32, int32)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								int32, int32))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(int32, int32,

						))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Int64:
			{
				arg0fun := args[0].WithFun().(func(*Env) int64)
				arg1fun := args[1].WithFun().(func(*Env) int64)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						int64, int64)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								int64, int64))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(int64, int64,

						))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Uint:
			{
				arg0fun := args[0].WithFun().(func(*Env) uint)
				arg1fun := args[1].WithFun().(func(*Env) uint)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						uint, uint)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint, uint))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint, uint,

						))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Uint8:
			{
				arg0fun := args[0].WithFun().(func(*Env) uint8)
				arg1fun := args[1].WithFun().(func(*Env) uint8)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						uint8, uint8)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint8, uint8))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint8, uint8,

						))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Uint16:
			{
				arg0fun := args[0].WithFun().(func(*Env) uint16)
				arg1fun := args[1].WithFun().(func(*Env) uint16)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						uint16, uint16)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint16, uint16))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint16, uint16,

						))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Uint32:
			{
				arg0fun := args[0].WithFun().(func(*Env) uint32)
				arg1fun := args[1].WithFun().(func(*Env) uint32)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						uint32, uint32)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint32, uint32))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint32, uint32,

						))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Uint64:
			{
				arg0fun := args[0].WithFun().(func(*Env) uint64)
				arg1fun := args[1].WithFun().(func(*Env) uint64)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						uint64, uint64)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint64, uint64))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint64, uint64,

						))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Uintptr:
			{
				arg0fun := args[0].WithFun().(func(*Env) uintptr)
				arg1fun := args[1].WithFun().(func(*Env) uintptr)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						uintptr, uintptr)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uintptr, uintptr))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(uintptr, uintptr,
						))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Float32:
			{
				arg0fun := args[0].WithFun().(func(*Env) float32)
				arg1fun := args[1].WithFun().(func(*Env) float32)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						float32, float32)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								float32, float32))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(float32, float32))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Float64:
			{
				arg0fun := args[0].WithFun().(func(*Env) float64)
				arg1fun := args[1].WithFun().(func(*Env) float64)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						float64, float64)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								float64, float64))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(
							float64, float64))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Complex64:
			{
				arg0fun := args[0].WithFun().(func(*Env) complex64)
				arg1fun := args[1].WithFun().(func(*Env) complex64)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						complex64, complex64)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								complex64, complex64))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(

							complex64, complex64))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Complex128:
			{
				arg0fun := args[0].WithFun().(func(*Env) complex128)
				arg1fun := args[1].WithFun().(func(*Env) complex128)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						complex128, complex128)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								complex128, complex128))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(

							complex128, complex128))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.String:
			{
				arg0fun := args[0].WithFun().(func(*Env) string)
				arg1fun := args[1].WithFun().(func(*Env) string)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(

						string, string)
					ret = func(env *Env) {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								string, string))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					ret = func(env *Env) {
						fun := exprfun(env).Interface().(func(

							string, string))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		}
	}

	if ret == nil {
		ret = func(env *Env) {
			funv := exprfun(env)

			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			callxr(funv, argv)
		}
	}
	return ret
}
