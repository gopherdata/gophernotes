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
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * type.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/types"
	"reflect"
)

// IsMethod reports whether a function type's contains a receiver, i.e. is a method.
// If IsMethod returns true, the actual receiver type is available as the first parameter, i.e. Type.In(0)
// It panics if the type's Kind is not Func.
func (t *xtype) IsMethod() bool {
	if t.Kind() != reflect.Func {
		xerrorf(t, "IsMethod of non-func type %v", t)
	}
	gtype := t.underlying().(*types.Signature)
	return gtype.Recv() != nil
}

// IsVariadic reports whether a function type's final input parameter is a "..." parameter.
// If so, t.In(t.NumIn() - 1) returns the parameter's implicit actual type []T.
// IsVariadic panics if the type's Kind is not Func.
func (t *xtype) IsVariadic() bool {
	if t.Kind() != reflect.Func {
		xerrorf(t, "In of non-func type %v", t)
	}
	return t.rtype.IsVariadic()
}

// In returns the type of a function type's i'th input parameter.
// It panics if the type's Kind is not Func.
// It panics if i is not in the range [0, NumIn()).
func (t *xtype) In(i int) Type {
	if t.Kind() != reflect.Func {
		xerrorf(t, "In of non-func type %v", t)
	}
	gtype := t.underlying().(*types.Signature)
	var va *types.Var
	if recv := gtype.Recv(); recv != nil {
		// include the receiver as first parameter
		if i == 0 {
			va = recv
		} else {
			va = gtype.Params().At(i - 1)
		}
	} else {
		va = gtype.Params().At(i)
	}
	return t.universe.MakeType(va.Type(), t.rtype.In(i))
}

// NumIn returns a function type's input parameter count.
// It panics if the type's Kind is not Func.
func (t *xtype) NumIn() int {
	if t.Kind() != reflect.Func {
		xerrorf(t, "NumIn of non-func type %v", t)
	}
	var nparams, nrecv int
	gtype := t.underlying().(*types.Signature)
	if gtype.Recv() != nil {
		nrecv = 1
	}
	if params := gtype.Params(); params != nil {
		nparams = params.Len()
	}
	return nparams + nrecv
}

// NumOut returns a function type's output parameter count.
// It panics if the type's Kind is not Func.
func (t *xtype) NumOut() int {
	if t.Kind() != reflect.Func {
		xerrorf(t, "NumOut of non-func type %v", t)
	}
	gtype := t.underlying().(*types.Signature)
	return gtype.Results().Len()
}

// Out returns the type of a function type's i'th output parameter.
// It panics if the type's Kind is not Func.
// It panics if i is not in the range [0, NumOut()).
func (t *xtype) Out(i int) Type {
	if t.Kind() != reflect.Func {
		xerrorf(t, "Out of non-func type %v", t)
	}
	gtype := t.underlying().(*types.Signature)
	va := gtype.Results().At(i)
	return t.universe.MakeType(va.Type(), t.rtype.Out(i))
}

func FuncOf(in []Type, out []Type, variadic bool) Type {
	return MethodOf(nil, in, out, variadic)
}

func (v *Universe) FuncOf(in []Type, out []Type, variadic bool) Type {
	return v.MethodOf(nil, in, out, variadic)
}

func MethodOf(recv Type, in []Type, out []Type, variadic bool) Type {
	v := universe
	if recv != nil {
		v = recv.Universe()
	} else if len(in) != 0 && in[0] != nil {
		v = in[0].Universe()
	} else if len(out) != 0 && out[0] != nil {
		v = out[0].Universe()
	}
	return v.MethodOf(recv, in, out, variadic)
}

func (v *Universe) MethodOf(recv Type, in []Type, out []Type, variadic bool) Type {
	gin := toGoTuple(in)
	gout := toGoTuple(out)
	rin := toReflectTypes(in)
	rout := toReflectTypes(out)
	var grecv *types.Var
	if recv != nil {
		rin = append([]reflect.Type{recv.ReflectType()}, rin...)
		grecv = toGoParam(recv)
	}
	return v.MakeType(
		types.NewSignature(grecv, gin, gout, variadic),
		reflect.FuncOf(rin, rout, variadic),
	)
}
