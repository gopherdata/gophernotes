// +build gomacro_jit

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * jit_experimental.go
 *
 *  Created on Feb 16, 2019
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/token"
	"os"
	r "reflect"
	"strconv"

	"github.com/cosmos72/gomacro/base/output"
	"github.com/cosmos72/gomacro/jit"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// jit.Comp wrapper
type Jit struct {
	c               jit.Comp
	lastCode        jit.Code
	lastMachineCode jit.MachineCode
}

type jitExpr = jit.Expr

type jitField struct {
	index jit.Const
}

var (
	jit_verbose int  = 0
	jit_enabled bool = true

	envInts  jitField // description of Env.Ints struct field
	envIP    jitField // description of Env.IP   struct field
	envCode  jitField // description of Env.Code struct field
	envOuter jitField // description of Env.Outerstruct field
	envOk    bool
)

func init() {
	if s := os.Getenv("GOMACRO_JIT_V"); s != "" {
		jit_verbose, _ = strconv.Atoi(s)
	}

	jitExtractEnvFields()

	jitCheckSupported()
}

func jitExtractEnvFields() {
	var sizeofUintptr = uintptr(jit.Uintptr.Size())

	tenv := r.TypeOf((*Env)(nil)).Elem()
	f, ok := tenv.FieldByName("Ints")
	if !ok || f.Offset%sizeofUintptr != 0 {
		return
	}
	envInts = makeJitField(f.Offset, jit.Uintptr)

	f, ok = tenv.FieldByName("IP")
	if !ok || f.Offset%f.Type.Size() != 0 {
		return
	}
	envIP = makeJitField(f.Offset, jit.Kind(f.Type.Kind()))

	f, ok = tenv.FieldByName("Code")
	if !ok || f.Offset%sizeofUintptr != 0 {
		return
	}
	envCode = makeJitField(f.Offset, jit.Uintptr)

	f, ok = tenv.FieldByName("Outer")
	if !ok || f.Offset%sizeofUintptr != 0 {
		return
	}
	envOuter = makeJitField(f.Offset, jit.Uintptr)

	envOk = true
}

func makeJitField(offset uintptr, kind jit.Kind) jitField {
	return jitField{
		index: jit.ConstUintptr(offset / uintptr(kind.Size())),
	}
}

func jitCheckSupported() {
	if !envOk {
		if jit_verbose > 0 {
			output.Debugf("Jit: failed to extract *Env fields")
		}
		jit_enabled = false
		return
	}
	arch := jit.Archs[jit.ARCH_ID]
	if arch == nil || !jit.SUPPORTED {
		if jit_verbose > 0 {
			output.Debugf("Jit: unsupported architecture or operating system")
		}
		jit_enabled = false
		return
	}
	// stmtNop = jitMakeInterpNop()
}

func NewJit() *Jit {
	if !jit_enabled {
		return nil
	}
	arch := jit.Archs[jit.ARCH_ID]
	var j Jit
	j.InitArch(arch)
	j.preamble() // before first use
	if jit_verbose > 0 {
		output.Debugf("Jit supported and enabled")
	}
	return &j
}

func (j *Jit) InitArch(arch jit.Arch) *Jit {
	j.c.InitArch(arch)
	j.lastCode = nil
	j.lastMachineCode = jit.MachineCode{}
	return j
}

func (j *Jit) Comp() *jit.Comp {
	return &j.c
}

func (j *Jit) Asm() *jit.Asm {
	return j.c.Asm()
}

func (j *Jit) LastCode() jit.Code {
	return j.lastCode
}

func (j *Jit) LastMachineCode() jit.MachineCode {
	return j.lastMachineCode
}

func (j *Jit) RegIdConfig() jit.RegIdConfig {
	return j.c.RegIdConfig
}

func (j *Jit) Log(e *Expr) {
	if jit_verbose > 2 {
		if e.Const() {
			output.Debugf("jit const:      %+v => %v", e, e.Jit)
		} else {
			output.Debugf("jit expr:       %+v => %v", e, e.Jit)
		}
	}
}

// return true if e can be jit-compiled
func (j *Jit) Can(e *Expr) bool {
	if j != nil && e.Jit == nil && e.Const() {
		j.Const(e)
	}
	ret := e.Jit != nil
	if !ret && jit_verbose > 0 {
		output.Debugf("jit could compile expr: %v", e)
	}
	return ret
}

// if supported, set e.Jit to jit constant == e.Lit.Value
// always returns e.
func (j *Jit) Const(e *Expr) *Expr {
	if j != nil && e.Jit == nil && e.Const() {
		switch e.Lit.Type.Kind() {
		case r.Bool, r.Int, r.Int8, r.Int16, r.Int32, r.Int64,
			r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr,
			r.Float32, r.Float64: // r.Complex64, r.Complex128

			c, err := jit.ConstInterface(e.Lit.Value, e.Lit.Type.ReflectType())
			if err == nil {
				e.Jit = c
			}
		}
		j.Log(e)
	}
	return e
}

// if supported, set e.Jit to jit expression that will compute xe
// always returns e.
func (j *Jit) Identity(e *Expr, xe *Expr) *Expr {
	if e.Jit == nil && j.Can(xe) {
		e.Jit = xe.Jit
		j.Log(e)
	}
	return e
}

// if supported, set e.Jit to jit expression that will compute t(xe)
// always returns e.
func (j *Jit) Cast(e *Expr, t xr.Type, xe *Expr) *Expr {
	if e.Jit == nil && j.Can(xe) {
		jop, err := jit.KindOp1(t.Kind())
		if err == nil {
			e.Jit = jit.NewExpr1(jop, xe.Jit)
			j.Log(e)
		}
	}
	return e
}

// if supported, set e.Jit to jit expression that will compute *xe
// always returns e.
func (j *Jit) Deref(e *Expr, xe *Expr) *Expr {
	if j != nil && e.Jit == nil && xe.Jit != nil {
		kind := jit.Kind(e.Type.Elem().Kind())
		if kind.Size() != 0 {
			e.Jit = jit.NewExprIdx(xe.Jit, jit.ConstUint64(0), kind)
			output.Debugf("jit deref: %v", e.Jit)
		}
	}
	return e
}

// if supported, set e.Jit to jit expression that will compute op xe
// always returns e.
func (j *Jit) UnaryExpr(e *Expr, op token.Token, xe *Expr) *Expr {
	if e.Jit == nil && j.Can(xe) {
		jop, err := jit.TokenOp1(op)
		if err == nil {
			e.Jit = jit.NewExpr1(jop, xe.Jit)
			j.Log(e)
		}
	}
	return e
}

// if supported, set e.Jit to jit expression that will compute xe op ye
// always returns e.
func (j *Jit) BinaryExpr(e *Expr, op token.Token, xe *Expr, ye *Expr) *Expr {
	if e.Jit == nil && j.Can(xe) && j.Can(ye) {
		jop, err := jit.TokenOp2(op)
		if err == nil {
			e.Jit = jit.NewExpr2(jop, xe.Jit, ye.Jit)
			j.Log(e)
		}
	}
	return e
}

// if supported, set e.Jit to jit expression that will read local variable
// always returns e.
func (j *Jit) Symbol(e *Expr) *Expr {
	if j == nil || e.Jit != nil || e.Sym == nil || e.Sym.Desc.Class() != IntBind {
		return e
	}
	sym := e.Sym
	idx := sym.Desc.Index()
	kind := jit.Kind(sym.Type.Kind())
	size := int(kind.Size())
	if size == 0 || idx*8%size != 0 {
		// unaligned memory. not supported.
		return e
	}
	if sym.Upn == 0 {
		mem, err := jit.MakeVar(idx, kind, j.RegIdConfig())
		if err != nil {
			if jit_verbose > 0 {
				output.Debugf("jit symbol %v failed: %v", sym, err)
			}
		} else {
			e.Jit = mem
			if jit_verbose > 2 {
				output.Debugf("jit symbol:     %v => %v", e, e.Jit)
			}
		}
		return e
	}
	// on amd64 and arm64, in a func(env *Env) ...
	// the parameter env is on the stack at [RSP+8]
	// env = stack[env_param]
	var env jit.Expr = j.Comp().MakeParam(8, jit.Uintptr)
	for i := 0; i < sym.Upn; i++ {
		// env = env.Outer
		env = jit.NewExprIdx(env, envOuter.index, jit.Uintptr)
	}
	// binds = env.Ints equivalent to &env.Ints[0]
	binds := jit.NewExprIdx(env, envInts.index, jit.Uintptr)
	// binds[index]
	e.Jit = jit.NewExprIdx(binds, jit.ConstInt(idx*8/size), kind)
	// output.Debugf("jit Symbol %v => e.Jit = %v, kind = %v", sym, e.Jit, kind)
	return e
}

// if supported, return a jit-compiled statement that will perform va OP= init
// return nil on failure
func (j *Jit) SetVar(va *Var, op token.Token, init *Expr) Stmt {
	if j == nil {
		return nil
	}
	if !j.Can(init) {
		return nil
	}
	op_assign := tokenWithAssign(op)
	vkind := jit.Kind(va.Type.Kind())
	ekind := jit.Kind(init.Type.Kind())
	switch op_assign {
	case token.SHL_ASSIGN, token.SHR_ASSIGN:
		if vkind.IsFloat() || ekind.Signed() {
			if jit_verbose > 0 {
				output.Debugf("jit setvar: invalid kinds for shift: %v %v %v",
					vkind, op, ekind)
			}
			return nil
		}
	default:
		if vkind != ekind {
			if jit_verbose > 0 {
				output.Debugf("jit setvar: mismatched kinds: %v %v %v",
					vkind, op, ekind)
				return nil
			}
		}
	}
	if jit_verbose > 2 {
		output.Debugf("jit setvar:     %v %v %v", va, op_assign, init.Jit)
	}
	if va.Upn != 0 {
		return j.setvarupn(va, op, init)
	}
	inst, err := jit.TokenInst2(op_assign)
	if err != nil {
		if jit_verbose > 0 {
			output.Debugf("jit setvar: TokenInst2(%v) failed: %v", op_assign, err)
			return nil
		}
	}
	mem, err := jit.MakeVar(va.Desc.Index(), jit.Kind(va.Type.Kind()), j.RegIdConfig())
	if err != nil {
		if jit_verbose > 0 {
			output.Debugf("jit setvar: MakeVar failed: %v", err)
		}
		return nil
	}
	// output.Debugf("jit setvar on %v", va)
	// output.Debugf("jit setvar to compile:  %v %v %v", mem, op_assign, init.Jit)
	j.Comp().Stmt2(inst, mem, init.Jit)
	// output.Debugf("jit setvar compiled  to: %v", j.Comp().Code())
	ret := j.stmt0()
	// output.Debugf("jit setvar assembled to: %v", j.LastMachineCode())
	return ret
}

func (j *Jit) setvarupn(va *Var, op token.Token, init *Expr) Stmt {
	idx := va.Desc.Index()
	kind := jit.Kind(va.Type.Kind())
	size := int(kind.Size())
	if size == 0 || idx*8%size != 0 {
		if jit_verbose > 0 {
			if size == 0 {
				output.Debugf("jit setvarupn: unsupported kind: %v", kind)
			} else {
				output.Debugf("jit setvarupn: unaligned variable: %v", va)
			}
		}
		return nil
	}
	index := jit.ConstInt(idx * 8 / size)

	// on amd64 and arm64, in a func(env *Env) ...
	// the parameter env is on the stack at [RSP+8]
	// env = stack[env_param]
	jc := j.Comp()
	var env jit.Expr = jc.MakeParam(8, jit.Uintptr)
	for i := 0; i < va.Upn; i++ {
		// env = env.Outer
		env = jit.NewExprIdx(env, envOuter.index, jit.Uintptr)
	}
	// binds = env.Ints equivalent to &env.Ints[0]
	binds := jit.NewExprIdx(env, envInts.index, jit.Uintptr)

	if op == token.ASSIGN {
		// binds[index] = init
		ret := j.stmt0(jit.NewStmt3(jit.IDX_ASSIGN, binds, index, init.Jit))
		// output.Debugf("jit setvarupn source    is: %v %v %v", va, op, init.Jit)
		// output.Debugf("jit setvarupn compiled  to: %v", j.LastCode())
		// output.Debugf("jit setvarupn assembled to: %v", j.LastMachineCode())
		return ret
	}

	op = tokenWithoutAssign(op)
	inst, err := jit.TokenOp2(op)
	if err != nil {
		if jit_verbose > 0 {
			output.Debugf("jit setvarupn: TokenOp2(%v) failed: %v", op, err)
		}
		return nil
	}
	// softbinds = binds
	softbinds := jc.NewSoftReg(jit.Uintptr)
	stmt1 := jit.NewStmt2(jit.ASSIGN, softbinds, binds)
	// value = softbinds[index] OP init
	value := jit.NewExpr2(
		inst,
		jit.NewExprIdx(softbinds, index, kind),
		init.Jit)
	// softbinds[index] = value
	stmt2 := jit.NewStmt3(jit.IDX_ASSIGN, softbinds, index, value)
	ret := j.stmt0(stmt1, stmt2)
	// output.Debugf("jit setvarupn source    is: %v %v %v", va, op, init.Jit)
	// output.Debugf("jit setvarupn compiled  to: %v", j.LastCode())
	// output.Debugf("jit setvarupn assembled to: %v", j.LastMachineCode())
	return ret
}

// if supported, return a jit-compiled Stmt that will evaluate Expr.
// return nil on failure
func (j *Jit) AsStmt(e *Expr) Stmt {
	if j == nil || e.Jit == nil {
		return nil
	}
	var success bool

	defer j.cleanup(&success)

	// compile accumulated jit expression and discard the result.
	jc := j.Comp()
	jc.Stmt1(jit.NOP, e.Jit)

	stmt := j.makeStmt()
	success = true
	return stmt
}

// if supported, replace e.Fun with a jit-compiled equivalent function.
// always returns e.
func (j *Jit) Fun(e *Expr) *Expr {
	if jit_verbose > 2 && j != nil {
		output.Debugf("jit to compile: %v with e.Jit = %v", e, e.Jit)
	}
	if j == nil || e.Jit == nil {
		return e
	}
	kind := jit.Kind(e.Type.Kind())
	if kind.Size() == 0 {
		if jit_verbose > 0 {
			output.Debugf("jit failed to compile: e.Jit = %v has unsupported kind %v", e.Jit, kind)
		}
		return e
	}
	fun := j.fun0(e, kind)
	if fun != nil {
		e.Fun = fun
		e.Jit = nil // in case we are invoked again on the same Expr
	}
	return e
}

// implementation of Jit.Fun
func (j *Jit) fun0(e *Expr, kind jit.Kind) I {
	var success bool

	defer j.cleanup(&success)

	// compile accumulated jit expression and copy result to stack.
	// on amd64 and arm64, in a func(env *Env) ...
	// the return value is on the stack at [RSP+16]
	jc := j.Comp()
	jc.Stmt2(jit.ASSIGN, jc.MakeParam(16, e.Jit.Kind()), e.Jit)
	fun := j.makeFun(kind)
	success = true
	return fun
}

// implementation of Jit.Stmt
func (j *Jit) stmt0(ts ...jit.Stmt) Stmt {
	var success bool

	defer j.cleanup(&success)

	jc := j.Comp()
	for _, t := range ts {
		jc.Stmt(t)
	}

	stmt := j.makeStmt()
	success = true
	return stmt
}

func (j *Jit) preamble() {
	jc := j.Comp()
	// caller may have compiled some code already, do not clear it
	// jc.ClearCode()
	// jc.ClearRegs()
	jc.Asm().RegIncUse(jc.RegIdConfig.RVAR)
	// on amd64 and arm64, in a func(env *Env) ...
	// the parameter env is on the stack at [RSP+8]
	rvar := jit.MakeReg(jc.RegIdConfig.RVAR, jit.Uint64)
	// env = stack[env_param]
	jc.Stmt2(jit.ASSIGN, rvar, jc.MakeParam(8, jit.Uint64))
	// rvar = env.Ints equivalent to rvar = &env.Ints[0]
	jc.Stmt2(jit.ASSIGN, rvar, jit.NewExprIdx(rvar, envInts.index, jit.Uint64))
}

func (j *Jit) cleanup(success *bool) {
	jc := j.Comp()

	// save them before clearing
	j.lastCode = jc.Code()
	j.lastMachineCode = jc.Asm().Code() // not jc.Assemble(), may panic again

	jc.ClearCode()
	jc.ClearRegs()

	// reinit for next use
	j.preamble()

	if *success {
		if jit_verbose > 1 {
			output.Debugf("jit compiled:   %v", j.lastCode)
			output.Debugf("jit assembled:  %v", j.lastMachineCode)
		}
	} else {
		err := recover()
		if jit_verbose > 0 {
			output.Debugf("jit failed:     %v", err)
		}
	}
}

func (j *Jit) makeFun(kind jit.Kind) I {
	jc := j.Comp()
	switch kind {
	case jit.Bool:
		var fun func(*Env) bool
		jc.Func(&fun)
		return fun
	case jit.Int:
		var fun func(*Env) int
		jc.Func(&fun)
		return fun
	case jit.Int8:
		var fun func(*Env) int8
		jc.Func(&fun)
		return fun
	case jit.Int16:
		var fun func(*Env) int16
		jc.Func(&fun)
		return fun
	case jit.Int32:
		var fun func(*Env) int32
		jc.Func(&fun)
		return fun
	case jit.Int64:
		var fun func(*Env) int64
		jc.Func(&fun)
		return fun
	case jit.Uint:
		var fun func(*Env) uint
		jc.Func(&fun)
		return fun
	case jit.Uint8:
		var fun func(*Env) uint8
		jc.Func(&fun)
		return fun
	case jit.Uint16:
		var fun func(*Env) uint16
		jc.Func(&fun)
		return fun
	case jit.Uint32:
		var fun func(*Env) uint32
		jc.Func(&fun)
		return fun
	case jit.Uint64:
		var fun func(*Env) uint64
		jc.Func(&fun)
		return fun
	case jit.Uintptr:
		var fun func(*Env) uintptr
		jc.Func(&fun)
		return fun
	case jit.Float32:
		var fun func(*Env) float32
		jc.Func(&fun)
		return fun
	case jit.Float64:
		var fun func(*Env) float64
		jc.Func(&fun)
		return fun
	/*
		case jit.Complex64:
		case jit.Complex128:
	*/
	default:
		return nil
	}
}

func (j *Jit) makeStmt() Stmt {
	if false {
		// use a closure instead of jit-compiling the epilogue
		var fun func(*Env)
		j.Comp().Func(&fun)
		return func(env *Env) (Stmt, *Env) {
			fun(env)
			ip := env.IP + 1
			env.IP = ip
			return env.Code[ip], env
		}
	}
	// jit-compile the following
	/*
		func(env *Env) (Stmt, *Env) {
			fun(env)
			ip := env.IP + 1
			env.IP = ip
			return env.Code[ip], env
		}
	*/

	jc := j.Comp()
	renv := jc.NewSoftReg(jit.Uint64)
	s := jc.NewSoftReg(jit.Uint64)
	t := jc.NewSoftReg(jit.Uint64)
	// on amd64 and arm64, in a func(env *Env) ...
	// the parameter env is on the stack at [RSP+8]
	source := jit.Source{
		// renv = stack[env_param]
		jit.ASSIGN, renv, jc.MakeParam(8, jit.Uint64),
		// t = env.IP
		jit.ASSIGN, t, jit.NewExprIdx(renv, envIP.index, jit.Uint64),
		// t++
		jit.INC, t,
		// env.IP = t
		jit.IDX_ASSIGN, renv, envIP.index, t,
		// s = env.Code
		jit.ASSIGN, s, jit.NewExprIdx(renv, envCode.index, jit.Uint64),
		// s = s[t] i.e. s = env.Code[t] i.e. s = env.Code[env.IP+1]
		jit.ASSIGN, s, jit.NewExprIdx(s, t, jit.Uintptr),
		// stack[env_result] = renv
		jit.ASSIGN, jc.MakeParam(24, jit.Uint64), renv,
		// stack[stmt_result] = s, with s == env.Code[env.IP+1]
		jit.ASSIGN, jc.MakeParam(16, jit.Uint64), s,
		jit.FREE, renv,
		jit.FREE, s,
		jit.FREE, t,
	}
	jc.Compile(source)
	if jit_verbose > 1 {
		output.Debugf("jit compiled:   %v", jc.Code())
		output.Debugf("jit assembled:  %v", jc.Assemble())
	}
	var f func(*Env) (Stmt, *Env)
	jc.Func(&f)
	return f
}
