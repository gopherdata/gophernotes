/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * closure_maps.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package closure_maps

import (
	"errors"
	"fmt"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
)

type Env struct {
	Binds map[string]r.Value
	Outer *Env
}

type SParam struct {
	Name string
	Type r.Type
}

type SReturn struct {
	result0 r.Value
	results []r.Value
}

type X func(*Env) (r.Value, []r.Value)
type X1 func(*Env) r.Value
type XInt func(*Env) int
type XBool func(*Env) bool

type Func func(args ...r.Value) (r.Value, []r.Value)
type FuncInt func(args ...r.Value) int

type XFunc func(env *Env) Func
type XFuncInt func(env *Env) FuncInt

var typeOfInt = r.TypeOf(int(0))

func errorf(format string, args ...interface{}) (r.Value, []r.Value) {
	panic(errors.New(fmt.Sprintf(format, args...)))
}

func warnExtraValues(extraValues []r.Value) {
	fmt.Printf("// warning: expression returned %d values, using only the first one: %v",
		len(extraValues), extraValues)
}

func NewEnv(outer *Env) *Env {
	return &Env{Outer: outer}
}

func (env *Env) DefineVar(name string, t r.Type, value r.Value) r.Value {
	if env.Binds == nil {
		env.Binds = make(map[string]r.Value)
	}
	if t == nil {
		t = value.Type()
	} else {
		value = value.Convert(t)
	}
	place := r.New(t).Elem()
	place.Set(value)
	env.Binds[name] = place
	return value
}

func (env *Env) DefineFunc(name string, t r.Type, value r.Value) r.Value {
	if env.Binds == nil {
		env.Binds = make(map[string]r.Value)
	}
	if t == nil {
		t = value.Type()
	} else {
		value = value.Convert(t)
	}
	env.Binds[name] = value
	return value
}

func IntToX(f XInt) X {
	return func(env *Env) (r.Value, []r.Value) {
		return r.ValueOf(f(env)), nil
	}
}

func Const(value interface{}) X {
	v := r.ValueOf(value)
	return func(env *Env) (r.Value, []r.Value) {
		return v, nil
	}
}

func Int(n int) XInt {
	return func(env *Env) int {
		return n
	}
}

func Var(name string) X {
	return func(env *Env) (r.Value, []r.Value) {
		for e := env; e != nil; e = e.Outer {
			if v, ok := e.Binds[name]; ok {
				return v, nil
			}
		}
		return errorf("undefined identifier: %v", name)
	}
}

func VarInt(name string) XInt {
	return func(env *Env) int {
		return int(env.Binds[name].Int())
	}
}

func VarSetInt(name string, expr XInt) X {
	return func(env *Env) (r.Value, []r.Value) {
		val := expr(env)
		for e := env; e != nil; e = e.Outer {
			if v, ok := env.Binds[name]; ok {
				v.SetInt(int64(val))
				return base.None, nil
			}
		}
		return errorf("undefined identifier: %v", name)
	}
}

func VarIncInt(name string) X {
	return func(env *Env) (r.Value, []r.Value) {
		for e := env; e != nil; e = e.Outer {
			if v, ok := env.Binds[name]; ok {
				v.SetInt(v.Int() + 1)
				return base.None, nil
			}
		}
		return errorf("undefined identifier: %v", name)
	}
}

func AddIntInt(lhs, rhs XInt) XInt {
	return func(env *Env) int {
		return lhs(env) + rhs(env)
	}
}

func SubIntInt(lhs, rhs XInt) XInt {
	return func(env *Env) int {
		return lhs(env) - rhs(env)
	}
}

func LessIntInt(lhs, rhs XInt) XBool {
	return func(env *Env) bool {
		return lhs(env) < rhs(env)
	}
}

func LesseqIntInt(lhs, rhs XInt) XBool {
	return func(env *Env) bool {
		return lhs(env) <= rhs(env)
	}
}

func If(pred XBool, then, els X) X {
	return func(env *Env) (r.Value, []r.Value) {
		if pred(env) {
			return then(env)
		} else {
			return els(env)
		}
	}
}

func For(init X, pred XBool, post X, body X) X {
	if init == nil && post == nil {
		return func(env *Env) (r.Value, []r.Value) {
			for pred(env) {
				body(env)
			}
			return base.None, nil
		}

	} else {
		if init == nil || post == nil {
			panic("invalid for(): init and post must be both present, or both omitted")
		}
		return func(env *Env) (r.Value, []r.Value) {
			for init(env); pred(env); post(env) {
				body(env)
			}
			return base.None, nil
		}
	}
}

func Nop(env *Env) (r.Value, []r.Value) {
	return base.None, nil
}

func Block(list ...X) X {
	switch len(list) {
	case 0:
		return Nop
	case 1:
		return list[0]
	case 2:
		return func(env *Env) (r.Value, []r.Value) {
			list[0](env)
			return list[1](env)
		}
	default:
		return func(env *Env) (r.Value, []r.Value) {
			n_1 := len(list) - 1
			for i := 0; i < n_1; i++ {
				list[i](env)
			}
			return list[n_1](env)
		}
	}
}

func Return(exprs ...X) X {
	switch n := len(exprs); n {
	case 0:
		return Nop
	case 1:
		expr := exprs[0]
		// return foo() returns *all* the values returned by foo, not just the first one
		return func(env *Env) (r.Value, []r.Value) {
			ret, rets := expr(env)
			panic(SReturn{ret, rets})
		}
	default:
		return func(env *Env) (r.Value, []r.Value) {
			n := len(exprs)
			rets := make([]r.Value, n)
			var extra []r.Value
			for i, value := range exprs {
				rets[i], extra = value(env)
				if len(extra) > 1 {
					warnExtraValues(extra)
				}
			}
			ret0 := base.None
			if len(rets) > 0 {
				ret0 = rets[0]
			}
			panic(SReturn{ret0, rets})
		}
	}
}

func ReturnInt(expr XInt) X {
	return func(env *Env) (r.Value, []r.Value) {
		ret := expr(env)
		panic(SReturn{r.ValueOf(ret), nil})
	}
}

func DeclVar(name string, expr X) X {
	return func(env *Env) (r.Value, []r.Value) {
		value, extra := expr(env)
		if len(extra) > 1 {
			warnExtraValues(extra)
		}
		value = env.DefineVar(name, nil, value)
		return value, nil
	}
}

func DeclFuncInt(name string, params []SParam, body X) XFuncInt {
	xf := MakeFuncInt(name, params, body)
	return func(env *Env) FuncInt {
		f := xf(env)
		env.DefineFunc(name, nil, r.ValueOf(f))
		return f
	}
}

func MakeFuncInt(name string, params []SParam, body X) XFuncInt {
	return func(env *Env) FuncInt {
		return func(args ...r.Value) (ret int) {
			fenv := NewEnv(env)
			panicking := true // use a flag to distinguish non-panic from panic(nil)
			defer func() {
				if panicking {
					pan := recover()
					switch p := pan.(type) {
					case SReturn:
						// return is implemented with a panic(cReturn{})
						ret = int(p.result0.Int())
					default:
						panic(pan)
					}
				}
			}()
			for i, param := range params {
				fenv.DefineVar(param.Name, param.Type, args[i])
			}
			ret0, extra := body(fenv)
			if len(extra) > 1 {
				warnExtraValues(extra)
			}
			panicking = false
			return int(ret0.Int())
		}
	}
}

func CallInt(fun X, args ...X) XInt {
	return func(env *Env) int {
		var extra []r.Value
		fvalue, extra := fun(env)
		if len(extra) > 1 {
			warnExtraValues(extra)
		}
		if fvalue == base.Nil || fvalue == base.None {
			errorf("undefined identifier: %v", fun)
		}
		f := fvalue.Interface().(FuncInt)
		n := len(args)
		values := make([]r.Value, n)
		for i, arg := range args {
			values[i], extra = arg(env)
			if len(extra) > 1 {
				warnExtraValues(extra)
			}
		}
		return f(values...)
	}
}

/*
  interpreted version of:

	func sum(n int) int {
		total := 0
		for i := 1; i <= n; i++ {
			total += i
		}
		return total
	}
*/
func DeclSum(env *Env, funcName string) FuncInt {
	const (
		n     = "n"
		total = "total"
		i     = "i"
	)
	return DeclFuncInt(
		funcName, []SParam{{"n", typeOfInt}},
		Block(
			DeclVar(total, Const(0)),
			For(DeclVar(i, Const(1)), LesseqIntInt(VarInt(i), VarInt(n)), VarIncInt(i),
				VarSetInt(total,
					AddIntInt(
						VarInt(total), VarInt(i),
					),
				),
			),
			ReturnInt(VarInt(total)),
		),
	)(env)
}

/*
  interpreted version of:

	func fibonacci(n int) int {
		if (n <= 2) {
			return 1
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}
*/
func DeclFibonacci(env *Env, funcName string) FuncInt {
	return DeclFuncInt(
		funcName, []SParam{{"n", typeOfInt}},
		Block(
			If(LessIntInt(VarInt("n"), Int(2)),
				ReturnInt(Int(1)),
				ReturnInt(
					AddIntInt(
						CallInt(Var(funcName), IntToX(SubIntInt(VarInt("n"), Int(1)))),
						CallInt(Var(funcName), IntToX(SubIntInt(VarInt("n"), Int(2)))),
					),
				),
			),
		),
	)(env)
}
