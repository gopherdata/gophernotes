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
 * closure_values.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package closure_values

import (
	r "reflect"

	"github.com/cosmos72/gomacro/base"
)

type Env struct {
	Binds []r.Value
	Outer *Env
}

func NewEnv(outer *Env) *Env {
	return &Env{
		Binds: make([]r.Value, 10),
		Outer: outer,
	}
}

type SParam struct {
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

func Var(upn, idx int) X {
	return func(env *Env) (r.Value, []r.Value) {
		for i := 0; i < upn; i++ {
			env = env.Outer
		}
		return env.Binds[idx], nil
	}
}

func VarInt0(env *Env) int {
	return int(env.Binds[0].Int())
}

func VarInt(idx int) XInt {
	return func(env *Env) int {
		return int(env.Binds[idx].Int())
	}
}

func VarSetInt(idx int, expr XInt) X {
	return func(env *Env) (r.Value, []r.Value) {
		val := expr(env)
		env.Binds[idx].SetInt(int64(val))
		return base.None, nil
	}
}

func VarIncInt(idx int) X {
	return func(env *Env) (r.Value, []r.Value) {
		v := env.Binds[idx]
		v.SetInt(v.Int() + 1)
		return base.None, nil
	}
}

func BitandIntInt(lhs, rhs XInt) XInt {
	return func(env *Env) int {
		return lhs(env) & rhs(env)
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

func MulIntInt(lhs, rhs XInt) XInt {
	return func(env *Env) int {
		return lhs(env) * rhs(env)
	}
}

func RshiftIntInt(lhs, rhs XInt) XInt {
	if false {
		return func(env *Env) int {
			l, r := lhs(env), rhs(env)
			// Debugf("rshift: lhs=%v, rhs=%v\n", l, r)
			// time.Sleep(time.Second)
			return l >> uint(r)
		}
	} else {
		return func(env *Env) int {
			return lhs(env) >> uint(rhs(env))
		}
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

func NoteqIntInt(lhs, rhs XInt) XBool {
	return func(env *Env) bool {
		return lhs(env) != rhs(env)
	}
}

func If(pred XBool, then, els X) X {
	if els != nil {
		return func(env *Env) (r.Value, []r.Value) {
			if pred(env) {
				return then(env)
			} else {
				return els(env)
			}
		}
	} else {
		return func(env *Env) (r.Value, []r.Value) {
			if pred(env) {
				return then(env)
			} else {
				return base.None, nil
			}
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
		expr0 := exprs[0]
		// return foo() returns *all* the values returned by foo, not just the first one
		return func(env *Env) (r.Value, []r.Value) {
			ret, rets := expr0(env)
			panic(SReturn{ret, rets})
		}
	default:
		return func(env *Env) (r.Value, []r.Value) {
			n := len(exprs)
			rets := make([]r.Value, n)
			for i, value := range exprs {
				rets[i], _ = value(env)
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

func DeclVar(idx int, expr X) X {
	return func(env *Env) (r.Value, []r.Value) {
		value, _ := expr(env)
		place := r.New(value.Type()).Elem()
		place.Set(value)
		env.Binds[idx] = place
		return value, nil
	}
}

func DeclFuncInt(idx int, paramTypes []r.Type, body X) XFuncInt {
	xf := MakeFuncInt(paramTypes, body)
	return func(env *Env) FuncInt {
		f := xf(env)
		env.Binds[idx] = r.ValueOf(f)
		return f
	}
}

func MakeFuncInt(paramTypes []r.Type, body X) XFuncInt {
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
			for i, paramType := range paramTypes {
				place := r.New(paramType).Elem()
				place.Set(args[i])
				fenv.Binds[i] = place
			}
			ret0, _ := body(fenv)
			panicking = false
			return int(ret0.Int())
		}
	}
}

func CallInt(fun X, args ...X) XInt {
	return func(env *Env) int {
		fvalue, _ := fun(env)
		f := fvalue.Interface().(FuncInt)
		n := len(args)
		values := make([]r.Value, n)
		for i, arg := range args {
			values[i], _ = arg(env)
		}
		return f(values...)
	}
}

var typeOfInt = r.TypeOf(int(0))

/*
  interpreted version of:

	func collatz(n int) {
		for n > 1 {
			if n&1 != 0 {
				n = ((n * 3) + 1) / 2
			} else {
				n = n / 2
			}
		}
	}
*/
func DeclCollatz(env *Env, idx int) FuncInt {
	const (
		n = 0
	)
	return DeclFuncInt(
		idx, []r.Type{typeOfInt},
		Block(
			For(nil, LessIntInt(Int(1), VarInt(n)), nil,
				If(NoteqIntInt(BitandIntInt(VarInt(n), Int(1)), Int(0)),
					VarSetInt(n,
						RshiftIntInt(
							AddIntInt(
								MulIntInt(VarInt(n), Int(3)),
								Int(1),
							),
							Int(1),
						),
					),
					VarSetInt(n,
						RshiftIntInt(
							VarInt(n),
							Int(1),
						),
					),
				),
			),
			ReturnInt(VarInt(n)),
		),
	)(env)
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
func DeclSum(env *Env, idx int) FuncInt {
	const (
		n     = 0
		total = 1
		i     = 2
	)
	return DeclFuncInt(
		idx, []r.Type{typeOfInt},
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
	const (
		n = 0
	)
	return DeclFuncInt(
		idx, []r.Type{typeOfInt},
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
