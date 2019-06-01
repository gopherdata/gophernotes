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
 * closure_interfaces.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package closure_interfaces

import (
	_ "errors"
	_ "fmt"

	"github.com/cosmos72/gomacro/base"
)

type Env struct {
	Binds []interface{}
	Outer *Env
}

func NewEnv(outer *Env) *Env {
	return &Env{
		Binds: make([]interface{}, 10),
		Outer: outer,
	}
}

type SParam struct {
	Name string
	// Type r.Type
}

type SReturn struct {
	result0 interface{}
	results []interface{}
}

type X func(*Env) (interface{}, []interface{})
type X1 func(*Env) interface{}
type XInt func(*Env) int
type XBool func(*Env) bool

type Func func(args ...interface{}) (interface{}, []interface{})
type FuncInt func(args ...interface{}) int

type XFunc func(env *Env) Func
type XFuncInt func(env *Env) FuncInt

func IntToX(f XInt) X {
	return func(env *Env) (interface{}, []interface{}) {
		return f(env), nil
	}
}

func Const(value interface{}) X {
	return func(env *Env) (interface{}, []interface{}) {
		return value, nil
	}
}

func Int(n int) XInt {
	return func(env *Env) int {
		return n
	}
}

func Var(upn, idx int) X {
	return func(env *Env) (interface{}, []interface{}) {
		for i := 0; i < upn; i++ {
			env = env.Outer
		}
		return env.Binds[idx], nil
	}
}

func VarInt0(env *Env) int {
	return env.Binds[0].(int)
}

func VarInt(idx int) XInt {
	return func(env *Env) int {
		return env.Binds[idx].(int)
		/*
			switch v := env.Binds[idx].(type) {
			case int:
				return v
			default:
				panic(errors.New(fmt.Sprintf("Binds[%v] = %#v <%T> is not an <int>", idx, v, v)))
			}
		*/
	}
}

func VarSetInt(idx int, expr XInt) X {
	return func(env *Env) (interface{}, []interface{}) {
		val := expr(env)
		env.Binds[idx] = val
		return base.None, nil
	}
}

func VarIncInt(idx int) X {
	return func(env *Env) (interface{}, []interface{}) {
		v := env.Binds[idx]
		env.Binds[idx] = v.(int) + 1
		return base.None, nil
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
	return func(env *Env) (interface{}, []interface{}) {
		if pred(env) {
			return then(env)
		} else {
			return els(env)
		}
	}
}

func For(init X, pred XBool, post X, body X) X {
	if init == nil && post == nil {
		return func(env *Env) (interface{}, []interface{}) {
			for pred(env) {
				body(env)
			}
			return base.None, nil
		}

	} else {
		if init == nil || post == nil {
			panic("invalid for(): init and post must be both present, or both omitted")
		}
		return func(env *Env) (interface{}, []interface{}) {
			for init(env); pred(env); post(env) {
				body(env)
			}
			return base.None, nil
		}
	}
}

func Nop(env *Env) (interface{}, []interface{}) {
	return base.None, nil
}

func Block(list ...X) X {
	switch len(list) {
	case 0:
		return Nop
	case 1:
		return list[0]
	case 2:
		return func(env *Env) (interface{}, []interface{}) {
			list[0](env)
			return list[1](env)
		}
	default:
		return func(env *Env) (interface{}, []interface{}) {
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
		expr0 := exprs[0]
		// return foo() returns *all* the values returned by foo, not just the first one
		return func(env *Env) (interface{}, []interface{}) {
			ret, rets := expr0(env)
			panic(SReturn{ret, rets})
		}
	default:
		return func(env *Env) (interface{}, []interface{}) {
			n := len(exprs)
			rets := make([]interface{}, n)
			for i, value := range exprs {
				rets[i], _ = value(env)
			}
			var ret0 interface{}
			if len(rets) > 0 {
				ret0 = rets[0]
			}
			panic(SReturn{ret0, rets})
		}
	}
}

func ReturnInt(expr XInt) X {
	return func(env *Env) (interface{}, []interface{}) {
		ret := expr(env)
		panic(SReturn{ret, nil})
	}
}

func DeclVar(idx int, expr X) X {
	return func(env *Env) (interface{}, []interface{}) {
		value, _ := expr(env)
		env.Binds[idx] = value
		return value, nil
	}
}

func DeclFuncInt(idx int, params []SParam, body X) XFuncInt {
	xf := MakeFuncInt(params, body)
	return func(env *Env) FuncInt {
		f := xf(env)
		env.Binds[idx] = f
		return f
	}
}

func MakeFuncInt(params []SParam, body X) XFuncInt {
	return func(env *Env) FuncInt {
		return func(args ...interface{}) (ret int) {
			fenv := NewEnv(env)
			panicking := true // use a flag to distinguish non-panic from panic(nil)
			defer func() {
				if panicking {
					pan := recover()
					switch p := pan.(type) {
					case SReturn:
						// return is implemented with a panic(cReturn{})
						ret = p.result0.(int)
					default:
						panic(pan)
					}
				}
			}()
			for i, _ := range params {
				fenv.Binds[i] = args[i]
			}
			ret0, _ := body(fenv)
			panicking = false
			return ret0.(int)
		}
	}
}

func CallInt(fun X, args ...X) XInt {
	return func(env *Env) int {
		fvalue, _ := fun(env)
		f := fvalue.(FuncInt)
		n := len(args)
		values := make([]interface{}, n)
		for i, arg := range args {
			values[i], _ = arg(env)
		}
		return f(values...)
	}
}

// var typeOfInt = r.TypeOf(int(0))

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
func DeclSum(env *Env, idx int) FuncInt {
	const (
		n     = 0
		total = 1
		i     = 2
	)
	return DeclFuncInt(
		idx, []SParam{{"n" /*typeOfInt*/}},
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
func DeclFibonacci(env *Env, idx int) FuncInt {
	return DeclFuncInt(
		idx, []SParam{{"n" /*typeOfInt*/}},
		Block(
			If(LessIntInt(VarInt0, Int(2)),
				ReturnInt(Int(1)),
				ReturnInt(
					AddIntInt(
						CallInt(Var(1, idx), IntToX(SubIntInt(VarInt0, Int(1)))),
						CallInt(Var(1, idx), IntToX(SubIntInt(VarInt0, Int(2)))),
					),
				),
			),
		),
	)(env)
}
