/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * closure_ints.go
 *
 *  Created on Mar 28, 2018
 *      Author Massimiliano Ghilardi
 */

package closure_ints

type Env struct {
	Binds []int
	Outer *Env
}

func NewEnv(outer *Env) *Env {
	return &Env{
		Binds: make([]int, 10),
		Outer: outer,
	}
}

type SReturnInt struct {
	result int
}

type X0 func(*Env)
type XInt func(*Env) int
type XBool func(*Env) bool

type Func func(arg int) int

type XFunc func(env *Env) Func

func Const(n int) XInt {
	return func(env *Env) int {
		return n
	}
}

func Arg(env *Env) int {
	return env.Binds[0]
}

func Var(idx int) XInt {
	return func(env *Env) int {
		return env.Binds[idx]
	}
}

func VarSet(idx int, expr XInt) X0 {
	return func(env *Env) {
		env.Binds[idx] = expr(env)
	}
}

func VarInc(idx int) X0 {
	return func(env *Env) {
		env.Binds[idx]++
	}
}

func Bitand(lhs, rhs XInt) XInt {
	return func(env *Env) int {
		return lhs(env) & rhs(env)
	}
}

func Add(lhs, rhs XInt) XInt {
	return func(env *Env) int {
		return lhs(env) + rhs(env)
	}
}

func Sub(lhs, rhs XInt) XInt {
	return func(env *Env) int {
		return lhs(env) - rhs(env)
	}
}

func Mul(lhs, rhs XInt) XInt {
	return func(env *Env) int {
		return lhs(env) * rhs(env)
	}
}

func Rshift(lhs, rhs XInt) XInt {
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

func Less(lhs, rhs XInt) XBool {
	return func(env *Env) bool {
		return lhs(env) < rhs(env)
	}
}

func Lesseq(lhs, rhs XInt) XBool {
	return func(env *Env) bool {
		return lhs(env) <= rhs(env)
	}
}

func Noteq(lhs, rhs XInt) XBool {
	return func(env *Env) bool {
		return lhs(env) != rhs(env)
	}
}

func If(pred XBool, then, els X0) X0 {
	if els != nil {
		return func(env *Env) {
			if pred(env) {
				then(env)
			} else {
				els(env)
			}
		}
	} else {
		return func(env *Env) {
			if pred(env) {
				then(env)
			}
		}
	}
}

func For(init X0, pred XBool, post X0, body X0) X0 {
	if init == nil && post == nil {
		return func(env *Env) {
			for pred(env) {
				body(env)
			}
		}

	} else {
		if init == nil || post == nil {
			panic("invalid for(): init and post must be both present, or both omitted")
		}
		return func(env *Env) {
			for init(env); pred(env); post(env) {
				body(env)
			}
		}
	}
}

func Nop(env *Env) {
}

func Block(list ...X0) X0 {
	switch len(list) {
	case 0:
		return Nop
	case 1:
		return list[0]
	case 2:
		return func(env *Env) {
			list[0](env)
			list[1](env)
		}
	default:
		return func(env *Env) {
			for _, stmt := range list {
				stmt(env)
			}
		}
	}
}

func Return(expr XInt) X0 {
	return func(env *Env) {
		ret := expr(env)
		panic(SReturnInt{ret})
	}
}

func DeclVar(idx int, expr XInt) X0 {
	return func(env *Env) {
		env.Binds[idx] = expr(env)
	}
}

func MakeFunc(body X0) XFunc {
	return func(env *Env) Func {
		return func(arg int) (ret int) {
			fenv := NewEnv(env)
			panicking := true // use a flag to distinguish non-panic from panic(nil)
			defer func() {
				if panicking {
					pan := recover()
					switch p := pan.(type) {
					case SReturnInt:
						// return is implemented with a panic(cReturn{})
						ret = int(p.result)
					default:
						panic(pan)
					}
				}
			}()
			fenv.Binds[0] = arg
			body(fenv)
			panicking = false
			return 0
		}
	}
}

func Call(fun *Func, arg XInt) XInt {
	return func(env *Env) int {
		return (*fun)(arg(env))
	}
}

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
func DeclCollatz(env *Env) Func {
	const (
		n = 0
	)
	return MakeFunc(
		Block(
			For(nil, Less(Const(1), Var(n)), nil,
				If(Noteq(Bitand(Var(n), Const(1)), Const(0)),
					VarSet(n,
						Rshift(
							Add(
								Mul(Var(n), Const(3)),
								Const(1),
							),
							Const(1),
						),
					),
					VarSet(n,
						Rshift(
							Var(n),
							Const(1),
						),
					),
				),
			),
			Return(Var(n)),
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
func DeclSum(env *Env) Func {
	const (
		n     = 0
		total = 1
		i     = 2
	)
	return MakeFunc(
		Block(
			DeclVar(total, Const(0)),
			For(DeclVar(i, Const(1)), Lesseq(Var(i), Var(n)), VarInc(i),
				VarSet(total,
					Add(
						Var(total), Var(i),
					),
				),
			),
			Return(Var(total)),
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
func DeclFibonacci(env *Env) Func {
	const (
		n = 0
	)
	var fib Func
	fib = MakeFunc(
		Block(
			If(Lesseq(Arg, Const(2)),
				Return(Const(1)),
				Return(
					Add(
						Call(&fib, Sub(Arg, Const(1))),
						Call(&fib, Sub(Arg, Const(2))),
					),
				),
			),
		),
	)(env)
	return fib
}
