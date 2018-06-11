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
 * type.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"fmt"
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
	gtype := t.gunderlying().(*types.Signature)
	return gtype.Recv() != nil
}

// IsVariadic reports whether a function type's final input parameter is a "..." parameter.
// If so, t.In(t.NumIn() - 1) returns the parameter's implicit actual type []T.
// IsVariadic panics if the type's Kind is not Func.
func (t *xtype) IsVariadic() bool {
	if t.Kind() != reflect.Func {
		xerrorf(t, "In of non-func type %v", t)
	}
	gtype := t.gunderlying().(*types.Signature)
	return gtype.Variadic()
}

// In returns the type of a function type's i'th input parameter.
// It panics if the type's Kind is not Func.
// It panics if i is not in the range [0, NumIn()).
func (t *xtype) In(i int) Type {
	if t.Kind() != reflect.Func {
		xerrorf(t, "In of non-func type %v", t)
	}
	gtype := t.gunderlying().(*types.Signature)
	recv := gtype.Recv()
	var va *types.Var
	if recv != nil {
		// include the receiver as first parameter
		if i == 0 {
			va = recv
		} else {
			va = gtype.Params().At(i - 1)
		}
	} else {
		va = gtype.Params().At(i)
	}
	t.NumIn() // for consistency check
	// contagion: if func reflect.Type is Forward, set whole function type to Forward
	rt := t.rtype
	if rt != rTypeOfForward {
		rt = rt.In(i)
	}
	return t.universe.MakeType(va.Type(), rt)
}

// NumIn returns a function type's input parameter count.
// It panics if the type's Kind is not Func.
func (t *xtype) NumIn() int {
	if t.Kind() != reflect.Func {
		xerrorf(t, "NumIn of non-func type %v", t)
	}
	n := 0
	gtype, ok := t.gunderlying().(*types.Signature)
	if !ok {
		xerrorf(t, "NumIn of non-func type %v (gotype = %v)", t, t.gunderlying())
	}
	if gtype.Recv() != nil {
		n++
	}
	if params := gtype.Params(); params != nil {
		n += params.Len()
	}
	if t.rtype != rTypeOfForward && t.rtype.NumIn() != n {
		var srecv string
		if gtype.Recv() != nil {
			srecv = fmt.Sprintf(" - including receiver type %v", gtype.Recv().Type())
		}
		xerrorf(t, `inconsistent function type: %v has %d params%s
      but its reflect.Type: %v has %d params`, t, n, srecv, t.rtype, t.rtype.NumIn())
	}
	return n
}

// NumOut returns a function type's output parameter count.
// It panics if the type's Kind is not Func.
func (t *xtype) NumOut() int {
	if t.Kind() != reflect.Func {
		xerrorf(t, "NumOut of non-func type %v", t)
	}
	gtype := t.gunderlying().(*types.Signature)
	return gtype.Results().Len()
}

// Out returns the type of a function type's i'th output parameter.
// It panics if the type's Kind is not Func.
// It panics if i is not in the range [0, NumOut()).
func (t *xtype) Out(i int) Type {
	if t.Kind() != reflect.Func {
		xerrorf(t, "Out of non-func type %v", t)
	}
	gtype := t.gunderlying().(*types.Signature)
	va := gtype.Results().At(i)
	// contagion: if func reflect.Type is Forward, return Forward
	rt := t.rtype
	if rt != rTypeOfForward {
		rt = rt.Out(i)
	}
	return t.universe.MakeType(va.Type(), rt)
}

func (v *Universe) FuncOf(in []Type, out []Type, variadic bool) Type {
	return v.MethodOf(nil, in, out, variadic)
}

/*
func FuncOf(in []Type, out []Type, variadic bool) Type {
	return MethodOf(nil, in, out, variadic)
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
*/

func (v *Universe) MethodOf(recv Type, in []Type, out []Type, variadic bool) Type {
	gin := toGoTuple(in)
	gout := toGoTuple(out)
	rin := toReflectTypes(in)
	rout := toReflectTypes(out)
	var grecv *types.Var
	if unwrap(recv) != nil {
		rin = append([]reflect.Type{recv.ReflectType()}, rin...)
		grecv = toGoParam(recv)
	}
	// contagion: if one or more in/out reflect.Type is Forward,
	// set the whole func reflect.Type to Forward
	var rfunc reflect.Type
loop:
	for {
		for _, rt := range rin {
			if rt == rTypeOfForward {
				rfunc = rTypeOfForward
				break loop
			}
		}
		for _, rt := range rout {
			if rt == rTypeOfForward {
				rfunc = rTypeOfForward
				break loop
			}
		}
		rfunc = reflect.FuncOf(rin, rout, variadic)
		break
	}
	return v.MakeType(
		types.NewSignature(grecv, gin, gout, variadic),
		rfunc,
	)
}
