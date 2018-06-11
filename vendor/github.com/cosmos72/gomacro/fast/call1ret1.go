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
 * call1ret1.go
 *
 *  Created on Apr 15, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) call1ret1(call *Call, maxdepth int) I {
	expr := call.Fun
	exprfun := expr.AsX1()
	funsym := expr.Sym
	funupn, funindex := -1, -1
	if funsym != nil {
		funupn = funsym.Upn
		funindex = funsym.Desc.Index()
		if funindex == NoIndex {
			Errorf("internal error: call1ret1() invoked for constant function %v. use call_builtin() instead", expr)
		}

	}
	t := expr.Type
	targ, tret := t.In(0), t.Out(0)
	karg, kret := targ.Kind(), tret.Kind()
	var ret I
	if KindToType(karg) != targ.ReflectType() || KindToType(kret) != tret.ReflectType() {
		return c.call1ret1namedtype(call, maxdepth)
	}

	arg := call.Args[0]
	argfun := arg.AsX1()
	var cachedfunv r.Value

	switch kret {

	case r.Bool:
		switch karg {
		case r.Bool:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(bool,

				) bool

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Bool()

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(bool,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) bool)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(bool,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) bool)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(bool,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(bool,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(bool,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int,

				) bool

				if arg.Const() {
					argconst := int(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(int,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(int,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(int,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int8:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int8,

				) bool

				if arg.Const() {
					argconst := int8(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int8)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int8)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(int8,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(int8,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(int8,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int16:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int16,

				) bool

				if arg.Const() {
					argconst := int16(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int16)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int16)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(int16,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(int16,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(int16,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int32,

				) bool

				if arg.Const() {
					argconst := int32(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int32,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int32)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int32,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(int32,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(int32,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(int32,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int64,

				) bool

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Int()

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int64,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int64)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int64,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(int64,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(int64,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(int64,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint,

				) bool

				if arg.Const() {
					argconst := uint(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(uint,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(uint,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(uint,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint8:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint8,

				) bool

				if arg.Const() {
					argconst := uint8(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint8)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint8)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(uint8,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(uint8,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(uint8,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint16:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint16,

				) bool

				if arg.Const() {
					argconst := uint16(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint16,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint16)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint16,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint16)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(uint16,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(uint16,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(uint16,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint32,

				) bool

				if arg.Const() {
					argconst := uint32(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint32,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint32)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint32,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(uint32,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(uint32,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(uint32,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint64,

				) bool

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Uint()

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint64,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint64)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint64,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(uint64,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(uint64,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(uint64,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uintptr:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uintptr,

				) bool

				if arg.Const() {
					argconst := uintptr(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uintptr,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uintptr)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uintptr,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uintptr)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(uintptr,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(uintptr,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(uintptr,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Float32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float32,

				) bool

				if arg.Const() {
					argconst := float32(

						r.ValueOf(arg.Value).Float())

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float32,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) float32)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float32,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) float32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(float32,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(float32,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(float32,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Float64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float64,

				) bool

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Float()

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float64,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) float64)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float64,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) float64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(float64,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(float64,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(float64,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Complex64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex64,

				) bool

				if arg.Const() {
					argconst := complex64(

						r.ValueOf(arg.Value).Complex())

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex64,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) complex64)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex64,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) complex64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(complex64,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(complex64,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(complex64,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Complex128:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex128,

				) bool

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Complex()

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex128,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) complex128)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex128,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) complex128)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(complex128,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(complex128,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(complex128,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.String:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(string,

				) bool

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).String()

					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(string,

							) bool)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) string)
					ret = func(env *Env) bool {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(string,

							) bool)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) string)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) bool {
						fun := env.Vals[funindex].Interface().(func(string,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) bool {
						fun := env.Outer.Vals[funindex].Interface().(func(string,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(string,

						) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		default:
			ret = func(env *Env) bool {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return ret0.Bool()
			}
		}

	case r.Int:
		switch karg {
		case r.Bool:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(bool,

				) int

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Bool()

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(bool,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) bool)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(bool,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) bool)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(bool,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(bool,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(bool,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int,

				) int

				if arg.Const() {
					argconst := int(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(int,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(int,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(int,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int8:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int8,

				) int

				if arg.Const() {
					argconst := int8(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int8)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int8)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(int8,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(int8,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(int8,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int16:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int16,

				) int

				if arg.Const() {
					argconst := int16(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int16)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int16)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(int16,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(int16,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(int16,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int32,

				) int

				if arg.Const() {
					argconst := int32(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int32,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int32)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int32,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(int32,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(int32,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(int32,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int64,

				) int

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Int()

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int64,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int64)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int64,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(int64,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(int64,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(int64,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint,

				) int

				if arg.Const() {
					argconst := uint(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(uint,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(uint,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(uint,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint8:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint8,

				) int

				if arg.Const() {
					argconst := uint8(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint8)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint8)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(uint8,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(uint8,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(uint8,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint16:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint16,

				) int

				if arg.Const() {
					argconst := uint16(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint16,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint16)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint16,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint16)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(uint16,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(uint16,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(uint16,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint32,

				) int

				if arg.Const() {
					argconst := uint32(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint32,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint32)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint32,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(uint32,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(uint32,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(uint32,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint64,

				) int

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Uint()

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint64,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint64)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint64,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(uint64,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(uint64,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(uint64,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uintptr:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uintptr,

				) int

				if arg.Const() {
					argconst := uintptr(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uintptr,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uintptr)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uintptr,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uintptr)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(uintptr,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(uintptr,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(uintptr,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Float32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float32,

				) int

				if arg.Const() {
					argconst := float32(

						r.ValueOf(arg.Value).Float())

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float32,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) float32)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float32,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) float32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(float32,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(float32,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(float32,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Float64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float64,

				) int

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Float()

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float64,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) float64)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float64,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) float64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(float64,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(float64,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(float64,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Complex64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex64,

				) int

				if arg.Const() {
					argconst := complex64(

						r.ValueOf(arg.Value).Complex())

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex64,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) complex64)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex64,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) complex64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(complex64,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(complex64,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(complex64,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Complex128:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex128,

				) int

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Complex()

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex128,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) complex128)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex128,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) complex128)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(complex128,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(complex128,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(complex128,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.String:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(string,

				) int

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).String()

					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(string,

							) int)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) string)
					ret = func(env *Env) int {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(string,

							) int)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) string)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int {
						fun := env.Vals[funindex].Interface().(func(string,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int {
						fun := env.Outer.Vals[funindex].Interface().(func(string,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int {
						fun := exprfun(env).Interface().(func(string,

						) int)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		default:
			ret = func(env *Env) int {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return int(ret0.Int())
			}
		}

	case r.Int8:
		if karg == kret {
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int8,

				) int8

				if arg.Const() {
					argconst := int8(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) int8 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8,

							) int8)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int8)
					ret = func(env *Env) int8 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8,

							) int8)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int8)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int8 {
						fun := env.Vals[funindex].Interface().(func(int8,

						) int8)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int8 {
						fun := env.Outer.Vals[funindex].Interface().(func(int8,

						) int8)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(int8,

						) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		} else {
			ret = func(env *Env) int8 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return int8(ret0.Int())
			}
		}

	case r.Int16:
		if karg == kret {
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int16,

				) int16

				if arg.Const() {
					argconst := int16(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) int16 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16,

							) int16)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int16)
					ret = func(env *Env) int16 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16,

							) int16)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int16)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int16 {
						fun := env.Vals[funindex].Interface().(func(int16,

						) int16)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int16 {
						fun := env.Outer.Vals[funindex].Interface().(func(int16,

						) int16)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(int16,

						) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		} else {
			ret = func(env *Env) int16 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return int16(ret0.Int())
			}
		}

	case r.Int32:
		if karg == kret {
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int32,

				) int32

				if arg.Const() {
					argconst := int32(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) int32 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int32,

							) int32)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int32)
					ret = func(env *Env) int32 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int32,

							) int32)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int32 {
						fun := env.Vals[funindex].Interface().(func(int32,

						) int32)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int32 {
						fun := env.Outer.Vals[funindex].Interface().(func(int32,

						) int32)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int32 {
						fun := exprfun(env).Interface().(func(int32,

						) int32)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		} else {
			ret = func(env *Env) int32 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return int32(ret0.Int())
			}
		}

	case r.Int64:
		switch karg {
		case r.Bool:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(bool,

				) int64

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Bool()

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(bool,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) bool)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(bool,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) bool)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(bool,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(bool,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(bool,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int,

				) int64

				if arg.Const() {
					argconst := int(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(int,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(int,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(int,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int8:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int8,

				) int64

				if arg.Const() {
					argconst := int8(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int8)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int8)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(int8,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(int8,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(int8,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int16:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int16,

				) int64

				if arg.Const() {
					argconst := int16(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int16)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int16)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(int16,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(int16,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(int16,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int32,

				) int64

				if arg.Const() {
					argconst := int32(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int32,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int32)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int32,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(int32,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(int32,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(int32,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int64,

				) int64

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Int()

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int64,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int64)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int64,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(int64,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(int64,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(int64,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint,

				) int64

				if arg.Const() {
					argconst := uint(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(uint,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(uint,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(uint,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint8:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint8,

				) int64

				if arg.Const() {
					argconst := uint8(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint8)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint8)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(uint8,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(uint8,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(uint8,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint16:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint16,

				) int64

				if arg.Const() {
					argconst := uint16(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint16,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint16)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint16,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint16)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(uint16,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(uint16,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(uint16,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint32,

				) int64

				if arg.Const() {
					argconst := uint32(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint32,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint32)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint32,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(uint32,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(uint32,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(uint32,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint64,

				) int64

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Uint()

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint64,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint64)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint64,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(uint64,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(uint64,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(uint64,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uintptr:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uintptr,

				) int64

				if arg.Const() {
					argconst := uintptr(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uintptr,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uintptr)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uintptr,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uintptr)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(uintptr,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(uintptr,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(uintptr,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Float32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float32,

				) int64

				if arg.Const() {
					argconst := float32(

						r.ValueOf(arg.Value).Float())

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float32,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) float32)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float32,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) float32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(float32,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(float32,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(float32,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Float64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float64,

				) int64

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Float()

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float64,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) float64)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float64,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) float64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(float64,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(float64,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(float64,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Complex64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex64,

				) int64

				if arg.Const() {
					argconst := complex64(

						r.ValueOf(arg.Value).Complex())

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex64,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) complex64)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex64,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) complex64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(complex64,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(complex64,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(complex64,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Complex128:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex128,

				) int64

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Complex()

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex128,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) complex128)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex128,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) complex128)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(complex128,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(complex128,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(complex128,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.String:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(string,

				) int64

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).String()

					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(string,

							) int64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) string)
					ret = func(env *Env) int64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(string,

							) int64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) string)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) int64 {
						fun := env.Vals[funindex].Interface().(func(string,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) int64 {
						fun := env.Outer.Vals[funindex].Interface().(func(string,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(string,

						) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		default:
			ret = func(env *Env) int64 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return ret0.Int()
			}
		}

	case r.Uint:
		switch karg {
		case r.Bool:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(bool,

				) uint

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Bool()

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(bool,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) bool)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(bool,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) bool)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(bool,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(bool,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(bool,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int,

				) uint

				if arg.Const() {
					argconst := int(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(int,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(int,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(int,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int8:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int8,

				) uint

				if arg.Const() {
					argconst := int8(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int8)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int8)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(int8,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(int8,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(int8,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int16:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int16,

				) uint

				if arg.Const() {
					argconst := int16(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int16)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int16)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(int16,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(int16,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(int16,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int32,

				) uint

				if arg.Const() {
					argconst := int32(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int32,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int32)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int32,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(int32,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(int32,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(int32,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int64,

				) uint

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Int()

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int64,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int64)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int64,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(int64,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(int64,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(int64,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint,

				) uint

				if arg.Const() {
					argconst := uint(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(uint,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(uint,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(uint,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint8:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint8,

				) uint

				if arg.Const() {
					argconst := uint8(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint8)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint8)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(uint8,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(uint8,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(uint8,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint16:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint16,

				) uint

				if arg.Const() {
					argconst := uint16(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint16,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint16)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint16,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint16)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(uint16,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(uint16,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(uint16,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint32,

				) uint

				if arg.Const() {
					argconst := uint32(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint32,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint32)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint32,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(uint32,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(uint32,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(uint32,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint64,

				) uint

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Uint()

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint64,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint64)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint64,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(uint64,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(uint64,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(uint64,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uintptr:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uintptr,

				) uint

				if arg.Const() {
					argconst := uintptr(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uintptr,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uintptr)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uintptr,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uintptr)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(uintptr,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(uintptr,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(uintptr,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Float32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float32,

				) uint

				if arg.Const() {
					argconst := float32(

						r.ValueOf(arg.Value).Float())

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float32,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) float32)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float32,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) float32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(float32,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(float32,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(float32,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Float64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float64,

				) uint

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Float()

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float64,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) float64)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float64,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) float64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(float64,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(float64,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(float64,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Complex64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex64,

				) uint

				if arg.Const() {
					argconst := complex64(

						r.ValueOf(arg.Value).Complex())

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex64,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) complex64)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex64,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) complex64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(complex64,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(complex64,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(complex64,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Complex128:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex128,

				) uint

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Complex()

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex128,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) complex128)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex128,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) complex128)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(complex128,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(complex128,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(complex128,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.String:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(string,

				) uint

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).String()

					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(string,

							) uint)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) string)
					ret = func(env *Env) uint {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(string,

							) uint)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) string)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint {
						fun := env.Vals[funindex].Interface().(func(string,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint {
						fun := env.Outer.Vals[funindex].Interface().(func(string,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(string,

						) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		default:
			ret = func(env *Env) uint {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return uint(ret0.Uint())
			}
		}

	case r.Uint8:
		if karg == kret {
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint8,

				) uint8

				if arg.Const() {
					argconst := uint8(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) uint8 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8) uint8)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint8)
					ret = func(env *Env) uint8 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8) uint8)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint8)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint8 {
						fun := env.Vals[funindex].Interface().(func(uint8,

						) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint8 {
						fun := env.Outer.Vals[funindex].Interface().(func(uint8,

						) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(uint8,

						) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		} else {
			ret = func(env *Env) uint8 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return uint8(ret0.Uint())
			}
		}

	case r.Uint16:
		if karg == kret {
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint16,

				) uint16

				if arg.Const() {
					argconst := uint16(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) uint16 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(
								uint16) uint16)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint16)
					ret = func(env *Env) uint16 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(
								uint16) uint16)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint16)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint16 {
						fun := env.Vals[funindex].Interface().(func(uint16,

						) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint16 {
						fun := env.Outer.Vals[funindex].Interface().(func(uint16,

						) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(uint16,

						) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		} else {
			ret = func(env *Env) uint16 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return uint16(ret0.Uint())
			}
		}

	case r.Uint32:
		if karg == kret {
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint32,

				) uint32

				if arg.Const() {
					argconst := uint32(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) uint32 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint32) uint32)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint32)
					ret = func(env *Env) uint32 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uint32) uint32)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint32 {
						fun := env.Vals[funindex].Interface().(func(uint32,

						) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint32 {
						fun := env.Outer.Vals[funindex].Interface().(func(uint32,

						) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(uint32,

						) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		} else {
			ret = func(env *Env) uint32 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return uint32(ret0.Uint())
			}
		}

	case r.Uint64:
		switch karg {
		case r.Bool:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(bool,

				) uint64

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Bool()

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(bool,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) bool)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(bool,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) bool)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(bool,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(bool,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(bool,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int,

				) uint64

				if arg.Const() {
					argconst := int(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(int,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(int,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(int,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int8:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int8,

				) uint64

				if arg.Const() {
					argconst := int8(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int8)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int8)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(int8,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(int8,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(int8,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int16:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int16,

				) uint64

				if arg.Const() {
					argconst := int16(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int16)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int16)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(int16,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(int16,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(int16,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int32,

				) uint64

				if arg.Const() {
					argconst := int32(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int32,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int32)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int32,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(int32,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(int32,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(int32,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int64,

				) uint64

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Int()

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int64,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int64)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int64,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(int64,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(int64,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(int64,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint,

				) uint64

				if arg.Const() {
					argconst := uint(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(uint,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(uint,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(uint,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint8:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint8,

				) uint64

				if arg.Const() {
					argconst := uint8(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint8)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint8)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(uint8,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(uint8,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(uint8,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint16:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint16,

				) uint64

				if arg.Const() {
					argconst := uint16(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint16,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint16)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint16,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint16)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(uint16,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(uint16,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(uint16,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint32,

				) uint64

				if arg.Const() {
					argconst := uint32(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint32,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint32)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint32,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(uint32,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(uint32,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(uint32,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint64,

				) uint64

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Uint()

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint64,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint64)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint64,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(uint64,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(uint64,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(uint64,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uintptr:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uintptr,

				) uint64

				if arg.Const() {
					argconst := uintptr(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uintptr,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uintptr)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uintptr,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uintptr)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(uintptr,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(uintptr,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(uintptr,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Float32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float32,

				) uint64

				if arg.Const() {
					argconst := float32(

						r.ValueOf(arg.Value).Float())

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float32,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) float32)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float32,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) float32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(float32,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(float32,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(float32,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Float64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float64,

				) uint64

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Float()

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float64,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) float64)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float64,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) float64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(float64,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(float64,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(float64,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Complex64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex64,

				) uint64

				if arg.Const() {
					argconst := complex64(

						r.ValueOf(arg.Value).Complex())

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex64,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) complex64)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex64,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) complex64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(complex64,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(complex64,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(complex64,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Complex128:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex128,

				) uint64

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Complex()

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex128,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) complex128)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex128,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) complex128)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(complex128,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(complex128,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(complex128,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.String:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(string,

				) uint64

				if arg.Const() {
					argconst := r.ValueOf(arg.Value).String()

					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(string,

							) uint64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) string)
					ret = func(env *Env) uint64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(string,

							) uint64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) string)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uint64 {
						fun := env.Vals[funindex].Interface().(func(string,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uint64 {
						fun := env.Outer.Vals[funindex].Interface().(func(string,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(string,

						) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		default:
			ret = func(env *Env) uint64 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return ret0.Uint()
			}
		}

	case r.Uintptr:
		if karg == kret {
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uintptr,

				) uintptr

				if arg.Const() {
					argconst := uintptr(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) uintptr {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uintptr) uintptr)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uintptr)
					ret = func(env *Env) uintptr {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								uintptr) uintptr)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uintptr)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) uintptr {
						fun := env.Vals[funindex].Interface().(func(uintptr,

						) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) uintptr {
						fun := env.Outer.Vals[funindex].Interface().(func(uintptr,

						) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(uintptr,

						) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		} else {
			ret = func(env *Env) uintptr {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return uintptr(ret0.Uint())
			}
		}

	case r.Float32:
		if karg == kret {
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float32,

				) float32

				if arg.Const() {
					argconst := float32(

						r.ValueOf(arg.Value).Float())

					ret = func(env *Env) float32 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								float32) float32)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) float32)
					ret = func(env *Env) float32 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								float32) float32)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) float32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) float32 {
						fun := env.Vals[funindex].Interface().(func(float32,

						) float32)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) float32 {
						fun := env.Outer.Vals[funindex].Interface().(func(float32,

						) float32)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(float32,

						) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		} else {
			ret = func(env *Env) float32 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return float32(ret0.Float())
			}
		}

	case r.Float64:
		switch karg {
		case r.Bool:

			{
				argfun := arg.WithFun().(func(env *Env) bool)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(bool,

					) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Int:

			{
				argfun := arg.WithFun().(func(env *Env) int)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(int,

					) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Int8:
			{
				argfun := arg.WithFun().(func(env *Env) int8)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(int8,

					) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Int16:
			{
				argfun := arg.WithFun().(func(env *Env) int16)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(int16,

					) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Int32:
			{
				argfun := arg.WithFun().(func(env *Env) int32)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(int32,

					) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Int64:
			{
				argfun := arg.WithFun().(func(env *Env) int64)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(int64,
					) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Uint:
			{
				argfun := arg.WithFun().(func(env *Env) uint)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(uint) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Uint8:
			{
				argfun := arg.WithFun().(func(env *Env) uint8)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(
						uint8) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Uint16:
			{
				argfun := arg.WithFun().(func(env *Env) uint16)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(

						uint16) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Uint32:
			{
				argfun := arg.WithFun().(func(env *Env) uint32)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(

						uint32) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Uint64:
			{
				argfun := arg.WithFun().(func(env *Env) uint64)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(

						uint64) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Uintptr:
			{
				argfun := arg.WithFun().(func(env *Env) uintptr)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(

						uintptr) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Float32:
			{
				argfun := arg.WithFun().(func(env *Env) float32)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(

						float32) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}

		case r.Float64:
			{
				argfun := arg.WithFun().(func(env *Env) float64)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(

						float64) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}

		case r.Complex64:
			{
				argfun := arg.WithFun().(func(env *Env) complex64)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(

						complex64) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}

		case r.Complex128:
			{
				argfun := arg.WithFun().(func(env *Env) complex128)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(

						complex128) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}

		case r.String:
			{
				argfun := arg.WithFun().(func(env *Env) string)
				ret = func(env *Env) float64 {
					fun := exprfun(env).Interface().(func(

						string) float64)
					arg := argfun(env)
					return fun(arg)
				}
			}

		default:
			ret = func(env *Env) float64 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return ret0.Float()
			}
		}

	case r.Complex64:
		if karg == kret {
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex64,

				) complex64

				if arg.Const() {
					argconst := complex64(

						r.ValueOf(arg.Value).Complex())

					ret = func(env *Env) complex64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								complex64) complex64)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) complex64)
					ret = func(env *Env) complex64 {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(

								complex64) complex64)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) complex64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) complex64 {
						fun := env.Vals[funindex].Interface().(func(complex64,

						) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) complex64 {
						fun := env.Outer.Vals[funindex].Interface().(func(complex64,

						) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(complex64,

						) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		} else {
			ret = func(env *Env) complex64 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return complex64(ret0.Complex())
			}
		}

	case r.Complex128:
		switch karg {
		case r.Bool:

			{
				argfun := arg.WithFun().(func(env *Env) bool)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(bool,

					) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Int:

			{
				argfun := arg.WithFun().(func(env *Env) int)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(int,

					) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Int8:
			{
				argfun := arg.WithFun().(func(env *Env) int8)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(int8,

					) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Int16:
			{
				argfun := arg.WithFun().(func(env *Env) int16)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(int16,

					) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Int32:
			{
				argfun := arg.WithFun().(func(env *Env) int32)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(int32,

					) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Int64:
			{
				argfun := arg.WithFun().(func(env *Env) int64)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(int64,
					) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Uint:
			{
				argfun := arg.WithFun().(func(env *Env) uint)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(uint) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Uint8:
			{
				argfun := arg.WithFun().(func(env *Env) uint8)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(
						uint8) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Uint16:
			{
				argfun := arg.WithFun().(func(env *Env) uint16)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(

						uint16) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Uint32:
			{
				argfun := arg.WithFun().(func(env *Env) uint32)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(

						uint32) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Uint64:
			{
				argfun := arg.WithFun().(func(env *Env) uint64)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(

						uint64) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Uintptr:
			{
				argfun := arg.WithFun().(func(env *Env) uintptr)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(

						uintptr) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}
		case r.Float32:
			{
				argfun := arg.WithFun().(func(env *Env) float32)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(

						float32) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}

		case r.Float64:
			{
				argfun := arg.WithFun().(func(env *Env) float64)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(

						float64) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}

		case r.Complex64:
			{
				argfun := arg.WithFun().(func(env *Env) complex64)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(

						complex64) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}

		case r.Complex128:
			{
				argfun := arg.WithFun().(func(env *Env) complex128)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(

						complex128) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}

		case r.String:
			{
				argfun := arg.WithFun().(func(env *Env) string)
				ret = func(env *Env) complex128 {
					fun := exprfun(env).Interface().(func(

						string) complex128)
					arg := argfun(env)
					return fun(arg)
				}
			}

		default:
			ret = func(env *Env) complex128 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return ret0.Complex()
			}
		}

	case r.String:
		switch karg {
		case r.Bool:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(bool,

				) string
				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Bool()

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(bool,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) bool)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(bool,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) bool)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(bool,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(bool,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(bool,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int,

				) string
				if arg.Const() {
					argconst := int(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(int,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(int,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(int,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int8:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int8,

				) string
				if arg.Const() {
					argconst := int8(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int8)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int8)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(int8,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(int8,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(int8,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int16:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int16,

				) string
				if arg.Const() {
					argconst := int16(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int16)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int16)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(int16,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(int16,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(int16,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int32,

				) string
				if arg.Const() {
					argconst := int32(

						r.ValueOf(arg.Value).Int())

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int32,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int32)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int32,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(int32,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(int32,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(int32,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Int64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int64,

				) string
				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Int()

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int64,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) int64)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int64,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) int64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(int64,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(int64,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(int64,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint,

				) string
				if arg.Const() {
					argconst := uint(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(uint,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(uint,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(uint,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint8:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint8,

				) string
				if arg.Const() {
					argconst := uint8(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint8)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint8)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(uint8,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(uint8,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(uint8,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint16:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint16,

				) string
				if arg.Const() {
					argconst := uint16(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint16,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint16)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint16,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint16)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(uint16,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(uint16,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(uint16,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint32,

				) string
				if arg.Const() {
					argconst := uint32(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint32,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint32)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint32,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(uint32,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(uint32,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(uint32,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uint64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint64,

				) string
				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Uint()

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint64,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uint64)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint64,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uint64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(uint64,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(uint64,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(uint64,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Uintptr:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uintptr,

				) string
				if arg.Const() {
					argconst := uintptr(

						r.ValueOf(arg.Value).Uint())

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uintptr,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) uintptr)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uintptr,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) uintptr)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(uintptr,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(uintptr,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(uintptr,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Float32:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float32,

				) string
				if arg.Const() {
					argconst := float32(

						r.ValueOf(arg.Value).Float())

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float32,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) float32)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float32,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) float32)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(float32,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(float32,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(float32,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Float64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float64,

				) string
				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Float()

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float64,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) float64)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float64,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) float64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(float64,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(float64,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(float64,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Complex64:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex64,

				) string
				if arg.Const() {
					argconst := complex64(

						r.ValueOf(arg.Value).Complex())

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex64,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) complex64)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex64,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) complex64)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(complex64,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(complex64,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(complex64,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.Complex128:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex128,

				) string
				if arg.Const() {
					argconst := r.ValueOf(arg.Value).Complex()

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex128,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) complex128)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex128,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) complex128)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(complex128,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(complex128,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(complex128,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		case r.String:

			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(string,

				) string
				if arg.Const() {
					argconst := r.ValueOf(arg.Value).String()

					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(string,

							) string)
						}
						return cachedfun(argconst)
					}
				} else {
					argfun := arg.Fun.(func(env *Env) string)
					ret = func(env *Env) string {
						funv := env.FileEnv.Vals[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(string,

							) string)
						}

						arg := argfun(env)
						return cachedfun(arg)
					}
				}
			} else {
				argfun := arg.WithFun().(func(env *Env) string)
				if funsym != nil && funupn == 0 {
					ret = func(env *Env) string {
						fun := env.Vals[funindex].Interface().(func(string,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else if funsym != nil && funupn == 1 {
					ret = func(env *Env) string {
						fun := env.Outer.Vals[funindex].Interface().(func(string,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				} else {
					ret = func(env *Env) string {
						fun := exprfun(env).Interface().(func(string,

						) string)
						arg := argfun(env)
						return fun(arg)
					}
				}

			}
		default:
			ret = func(env *Env) string {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := callxr(funv, argv)[0]
				return ret0.String()
			}
		}

	}
	if ret == nil {
		ret = func(env *Env) r.Value {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			return callxr(funv, argv)[0]
		}
	}
	return ret
}
func (c *Comp) call1ret1namedtype(call *Call, maxdepth int) I {
	expr := call.Fun
	exprfun := expr.AsX1()
	t := expr.Type
	kret := t.Out(0).Kind()

	argfun := call.Args[0].AsX1()
	var ret I

	switch kret {
	case r.Bool:
		ret = func(env *Env) bool {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return retv.Bool()
		}
	case r.Int:
		ret = func(env *Env) int {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return int(retv.Int())
		}
	case r.Int8:
		ret = func(env *Env) int8 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return int8(retv.Int())
		}

	case r.Int16:
		ret = func(env *Env) int16 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return int16(retv.Int())
		}

	case r.Int32:
		ret = func(env *Env) int32 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return int32(retv.Int())
		}

	case r.Int64:
		ret = func(env *Env) int64 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return retv.Int()
		}

	case r.Uint:
		ret = func(env *Env) uint {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return uint(retv.Uint())
		}

	case r.Uint8:
		ret = func(env *Env) uint8 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return uint8(retv.Uint())
		}

	case r.Uint16:
		ret = func(env *Env) uint16 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return uint16(retv.Uint())
		}

	case r.Uint32:
		ret = func(env *Env) uint32 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return uint32(retv.Uint())
		}

	case r.Uint64:
		ret = func(env *Env) uint64 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return retv.Uint()
		}

	case r.Uintptr:
		ret = func(env *Env) uintptr {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return uintptr(retv.Uint())
		}

	case r.Float32:
		ret = func(env *Env) float32 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return float32(retv.Float())
		}

	case r.Float64:
		ret = func(env *Env) float64 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return retv.Float()
		}

	case r.Complex64:
		ret = func(env *Env) complex64 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return complex64(retv.Complex())
		}

	case r.Complex128:
		ret = func(env *Env) complex128 {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return retv.Complex()
		}

	case r.String:
		ret = func(env *Env) string {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := callxr(funv, argv)[0]
			return retv.String()
		}

	}
	if ret == nil {
		ret = func(env *Env) r.Value {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			return callxr(funv, argv)[0]
		}
	}
	return ret
}
