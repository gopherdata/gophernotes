// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements commonly used type predicates.

package typeutil

import (
	"bytes"
	"go/token"
	"strings"
	"testing"

	"github.com/cosmos72/gomacro/go/etoken"

	"github.com/cosmos72/gomacro/go/types"
)

func fail(t *testing.T, actual interface{}, expected interface{}) {
	t.Errorf("expecting %v <%T>, found %v <%T>\n", expected, expected, actual, actual)
}

func fail2(t *testing.T, actual interface{}, expected interface{}) {
	t.Errorf("expecting %#v <%T>,\n\tfound %#v <%T>\n", expected, expected, actual, actual)
}

func is(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		fail(t, actual, expected)
	}
}

func istrue(t *testing.T, actual bool) {
	if !actual {
		fail(t, actual, true)
	}
}

func isfalse(t *testing.T, actual bool) {
	if actual {
		fail(t, actual, false)
	}
}

func isidentical(t *testing.T, actual types.Type, expected types.Type) {
	if !Identical(actual, expected) {
		fail(t, actual, expected)
	}
}

func TestBasic(t *testing.T) {
	names := []string{
		types.Bool:       "bool",
		types.Int:        "int",
		types.Int8:       "int8",
		types.Int16:      "int16",
		types.Int32:      "int32",
		types.Int64:      "int64",
		types.Uint:       "uint",
		types.Uint8:      "uint8",
		types.Uint16:     "uint16",
		types.Uint32:     "uint32",
		types.Uint64:     "uint64",
		types.Uintptr:    "uintptr",
		types.Float32:    "float32",
		types.Float64:    "float64",
		types.Complex64:  "complex64",
		types.Complex128: "complex128",
		types.String:     "string",
		types.UntypedNil: "", // highest index
	}

	for i, typ := range types.Typ {
		name := names[i]
		if typ == nil || len(name) == 0 {
			continue
		}
		is(t, typ.Underlying(), typ)
		is(t, typ.Underlying(), typ)
		is(t, typ.Kind(), types.BasicKind(i))
		is(t, typ.Name(), name)
		is(t, typ.String(), name)
	}
}

func TestVar(t *testing.T) {
	newVar(t, "n", types.Typ[types.Int])
}

func TestTuple(t *testing.T) {
	newTuple(t,
		newVar(t, "x", types.Typ[types.Float32]),
		newVar(t, "y", types.Typ[types.Float64]),
	)
}

func TestNamed(t *testing.T) {
	newNamed(t, "MyInt", types.Typ[types.Int])
}

func TestFuncSignature(t *testing.T) {
	newSignature(t,
		nil,
		newTuple(t,
			newVar(t, "x", types.Typ[types.Float32]),
			newVar(t, "y", types.Typ[types.Float64]),
		),
		newTuple(t,
			newVar(t, "z", types.Typ[types.Complex64]),
			newVar(t, "w", types.Typ[types.Complex128]),
		),
		false,
	)
}

func TestMethodSignature(t *testing.T) {
	typ := newNamed(t, "MyInt", types.Typ[types.Int])
	newSignature(t,
		newVar(t, "obj", typ),
		newTuple(t,
			newVar(t, "x", types.Typ[types.Float32]),
			newVar(t, "y", types.Typ[types.Float64]),
		),
		newTuple(t,
			newVar(t, "z", types.Typ[types.Complex64]),
			newVar(t, "w", types.Typ[types.Complex128]),
		),
		false,
	)
}

func TestFunc(t *testing.T) {
	typ := newNamed(t, "MyInt", types.Typ[types.Int])
	sig := newSignature(t,
		newVar(t, "", typ), // types.Func.String() only shows receiver type, not its name
		nil,
		newTuple(t,
			newVar(t, "", types.Typ[types.String]),
		),
		false,
	)
	fun := types.NewFunc(token.NoPos, nil, "String", sig)
	is(t, fun.Name(), "String")
	is(t, fun.Type(), sig)
	is(t, fun.String(), String2(fun.Name(), fun.Type()))
}

func TestMap(t *testing.T) {
	typ := newNamed(t, "MyInt", types.Typ[types.Int])
	sig1 := newSignature(t,
		newVar(t, "obj", typ), // value reveicer
		nil,
		newTuple(t,
			newVar(t, "", types.Typ[types.String]),
		),
		false,
	)
	sig2 := newSignature(t,
		newVar(t, "obj", types.NewPointer(typ)), // pointer reveicer
		nil,
		newTuple(t,
			newVar(t, "", types.Typ[types.String]),
		),
		false,
	)
	m := Map{}
	m.Set(sig1, 1)
	m.Set(sig2, 2)
	is(t, m.Len(), 2)
	is(t, m.At(sig1), 1)
	is(t, m.At(sig2), 2)
}

func newVar(t *testing.T, namestr string, typ types.Type) *types.Var {
	v := types.NewVar(token.NoPos, nil, namestr, typ)
	isfalse(t, v.Anonymous())
	isfalse(t, v.Embedded())
	isfalse(t, v.IsField())
	is(t, v.Name(), namestr)
	is(t, v.Pkg(), (*types.Package)(nil))
	is(t, v.Type(), typ)
	is(t, v.String(), "var "+namestr+" "+typ.String())
	return v
}

func newTuple(t *testing.T, x ...*types.Var) *types.Tuple {
	tuple := types.NewTuple(x...)
	is(t, tuple.Len(), len(x))
	buf := []string{"("}
	for i, v := range x {
		is(t, tuple.At(i), v)
		if i != 0 {
			buf = append(buf, ", ")
		}
		if v.Name() != "" {
			buf = append(buf, v.Name(), " ")
		}
		buf = append(buf, v.Type().String())
	}
	buf = append(buf, ")")

	is(t, tuple.Underlying(), tuple)
	is(t, tuple.String(), strings.Join(buf, ""))
	return tuple
}

func newNamed(t *testing.T, namestr string, underlying types.Type) *types.Named {
	name := types.NewTypeName(token.NoPos, nil, "MyInt", nil)
	named := types.NewNamed(name, types.Typ[types.Int], nil)
	if !etoken.GENERICS_V2_CTI {
		is(t, named.NumMethods(), 0)
	}
	is(t, named.Obj(), name)
	is(t, named.Underlying(), types.Typ[types.Int])
	is(t, named.String(), "MyInt")
	return named
}

func newSignature(t *testing.T, recv *types.Var, params *types.Tuple, results *types.Tuple, variadic bool) *types.Signature {
	sig := types.NewSignature(recv, params, results, variadic)
	is(t, sig.Recv(), recv)
	is(t, sig.Params(), params)
	is(t, sig.Results(), results)
	is(t, sig.Variadic(), variadic)
	is(t, sig.Underlying(), sig)

	paramstr := "()"
	if params != nil {
		paramstr = params.String()
	}
	var resultstr string
	if results != nil {
		switch results.Len() {
		case 0:
			break
		case 1:
			resultstr = " " + varString(results.At(0))
		default:
			resultstr = " " + results.String()
		}
	}

	// types.Type.String() does NOT show method receiver
	is(t, sig.String(), "func"+paramstr+resultstr)
	// instead typeutil.String(types.Type) also shows it
	if recv == nil {
		is(t, String(sig), "func"+paramstr+resultstr)
	} else {
		is(t, String(sig), "func ("+varString(recv)+")."+paramstr+resultstr)
	}
	return sig
}

func varString(v *types.Var) string {
	var buf bytes.Buffer
	writeVar(&buf, v, false)
	return buf.String()
}
