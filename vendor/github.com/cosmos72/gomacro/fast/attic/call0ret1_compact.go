// +build gomacro_fast_compact

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
 * call0ret1compact.go
 *
 *  Created on Jun 14, 2017
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

	tret := expr.Type.Out(0)
	kret := tret.Kind()
	var ret I

	switch kret {
	case r.Bool:
		ret = func(env *Env) bool {
			return exprfun(env).Call(base.ZeroValues)[0].Bool()
		}
	case r.Int:
		ret = func(env *Env) int {
			return int(exprfun(env).Call(base.ZeroValues)[0].Int())
		}
	case r.Int8:
		ret = func(env *Env) int8 {
			return int8(exprfun(env).Call(base.ZeroValues)[0].Int())
		}
	case r.Int16:
		ret = func(env *Env) int16 {
			return int16(exprfun(env).Call(base.ZeroValues)[0].Int())
		}
	case r.Int32:
		ret = func(env *Env) int32 {
			return int32(exprfun(env).Call(base.ZeroValues)[0].Int())
		}
	case r.Int64:
		ret = func(env *Env) int64 {
			return exprfun(env).Call(base.ZeroValues)[0].Int()
		}
	case r.Uint:
		ret = func(env *Env) uint {
			return uint(exprfun(env).Call(base.ZeroValues)[0].Uint())
		}
	case r.Uint8:
		ret = func(env *Env) uint8 {
			return uint8(exprfun(env).Call(base.ZeroValues)[0].Uint())
		}
	case r.Uint16:
		ret = func(env *Env) uint16 {
			return uint16(exprfun(env).Call(base.ZeroValues)[0].Uint())
		}
	case r.Uint32:
		ret = func(env *Env) uint32 {
			return uint32(exprfun(env).Call(base.ZeroValues)[0].Uint())
		}
	case r.Uint64:
		ret = func(env *Env) uint64 {
			return exprfun(env).Call(base.ZeroValues)[0].Uint()
		}
	case r.Uintptr:
		ret = func(env *Env) uintptr {
			return uintptr(exprfun(env).Call(base.ZeroValues)[0].Uint())
		}
	case r.Float32:
		ret = func(env *Env) float32 {
			return float32(exprfun(env).Call(base.ZeroValues)[0].Float())
		}
	case r.Float64:
		ret = func(env *Env) float64 {
			return exprfun(env).Call(base.ZeroValues)[0].Float()
		}
	case r.Complex64:
		ret = func(env *Env) complex64 {
			return complex64(exprfun(env).Call(base.ZeroValues)[0].Complex())
		}
	case r.Complex128:
		ret = func(env *Env) complex128 {
			return exprfun(env).Call(base.ZeroValues)[0].Complex()
		}
	case r.String:
		ret = func(env *Env) string {
			return exprfun(env).Call(base.ZeroValues)[0].String()
		}
	default:
		ret = func(env *Env) r.Value {
			return exprfun(env).Call(base.ZeroValues)[0]
		}
	}
	return ret
}
