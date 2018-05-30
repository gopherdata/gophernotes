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
 * util.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"fmt"
	"go/token"
	"go/types"
	"reflect"
)

func concat(a, b []int) []int {
	na := len(a)
	c := make([]int, na+len(b))
	copy(c, a)
	copy(c[na:], b)
	return c
}

type Error struct {
	Type   Type
	format string
	args   []interface{}
}

func (e *Error) Error() string {
	return fmt.Sprintf(e.format, e.args...)
}

func errorf(t Type, format string, args ...interface{}) {
	panic(&Error{t, format, args})
}

func xerrorf(t *xtype, format string, args ...interface{}) {
	panic(&Error{wrap(t), format, args})
}

func dirToGdir(dir reflect.ChanDir) types.ChanDir {
	var gdir types.ChanDir
	switch dir {
	case reflect.RecvDir:
		gdir = types.RecvOnly
	case reflect.SendDir:
		gdir = types.SendOnly
	case reflect.BothDir:
		gdir = types.SendRecv
	}
	return gdir
}

func gtypeToKind(t *xtype, gtype types.Type) reflect.Kind {
	gtype = gtype.Underlying()
	var kind reflect.Kind
	switch gtype := gtype.(type) {
	case *types.Array:
		kind = reflect.Array
	case *types.Basic:
		kind = ToReflectKind(gtype.Kind())
	case *types.Chan:
		kind = reflect.Chan
	case *types.Signature:
		kind = reflect.Func
	case *types.Interface:
		kind = reflect.Interface
	case *types.Map:
		kind = reflect.Map
	case *types.Pointer:
		kind = reflect.Ptr
	case *types.Slice:
		kind = reflect.Slice
	case *types.Struct:
		kind = reflect.Struct
	// case *types.Named: // impossible, handled above
	default:
		xerrorf(t, "unsupported types.Type: %v", gtype)
	}
	// debugf("gtypeToKind(%T) -> %v", gtype, kind)
	return kind
}

func IsGoUntypedKind(gkind types.BasicKind) bool {
	switch gkind {
	case types.UntypedBool, types.UntypedInt, types.UntypedRune,
		types.UntypedFloat, types.UntypedComplex, types.UntypedString, types.UntypedNil:
		return true
	default:
		return false
	}
}

func ToReflectKind(gkind types.BasicKind) reflect.Kind {
	var kind reflect.Kind
	switch gkind {
	case types.Bool, types.UntypedBool:
		kind = reflect.Bool
	case types.Int, types.UntypedInt:
		kind = reflect.Int
	case types.Int8:
		kind = reflect.Int8
	case types.Int16:
		kind = reflect.Int16
	case types.Int32, types.UntypedRune:
		kind = reflect.Int32
	case types.Int64:
		kind = reflect.Int64
	case types.Uint:
		kind = reflect.Uint
	case types.Uint8:
		kind = reflect.Uint8
	case types.Uint16:
		kind = reflect.Uint16
	case types.Uint32:
		kind = reflect.Uint32
	case types.Uint64:
		kind = reflect.Uint64
	case types.Uintptr:
		kind = reflect.Uintptr
	case types.Float32:
		kind = reflect.Float32
	case types.Float64, types.UntypedFloat:
		kind = reflect.Float64
	case types.Complex64:
		kind = reflect.Complex64
	case types.Complex128, types.UntypedComplex:
		kind = reflect.Complex128
	case types.String, types.UntypedString:
		kind = reflect.String
	case types.UnsafePointer:
		kind = reflect.UnsafePointer
	case types.UntypedNil:
		kind = reflect.Invalid
	default:
		errorf(nil, "unsupported types.BasicKind: %v", gkind)
	}
	return kind
}

func ToBasicKind(kind reflect.Kind, untyped bool) types.BasicKind {
	var gkind types.BasicKind
	switch kind {
	case reflect.Bool:
		if untyped {
			gkind = types.UntypedBool
		} else {
			gkind = types.Bool
		}
	case reflect.Int:
		if untyped {
			gkind = types.Int
		} else {
			gkind = types.UntypedInt
		}
	case reflect.Int8:
		gkind = types.Int8
	case reflect.Int16:
		gkind = types.Int16
	case reflect.Int32:
		if untyped {
			gkind = types.UntypedRune
		} else {
			gkind = types.Int32
		}
	case reflect.Int64:
		gkind = types.Int64
	case reflect.Uint:
		gkind = types.Uint
	case reflect.Uint8:
		gkind = types.Uint8
	case reflect.Uint16:
		gkind = types.Uint16
	case reflect.Uint32:
		gkind = types.Uint32
	case reflect.Uint64:
		gkind = types.Uint64
	case reflect.Uintptr:
		gkind = types.Uintptr
	case reflect.Float32:
		gkind = types.Float32
	case reflect.Float64:
		if untyped {
			gkind = types.UntypedFloat
		} else {
			gkind = types.Float64
		}
	case reflect.Complex64:
		gkind = types.Complex64
	case reflect.Complex128:
		if untyped {
			gkind = types.UntypedComplex
		} else {
			gkind = types.Complex128
		}
	case reflect.String:
		if untyped {
			gkind = types.UntypedString
		} else {
			gkind = types.String
		}
	case reflect.UnsafePointer:
		gkind = types.UnsafePointer
	case reflect.Invalid:
		gkind = types.UntypedNil
	default:
		errorf(nil, "unsupported refletc.Kind: %v", kind)
	}
	return gkind
}

func path(gpkg *types.Package) string {
	if gpkg == nil {
		return ""
	}
	return gpkg.Path()
}

func toReflectTypes(ts []Type) []reflect.Type {
	rts := make([]reflect.Type, len(ts))
	for i, t := range ts {
		rts[i] = t.ReflectType()
	}
	return rts
}

func toGoParam(t Type) *types.Var {
	return types.NewParam(token.NoPos, nil, "", t.GoType())
}

func toGoParams(ts []Type) []*types.Var {
	vars := make([]*types.Var, len(ts))
	for i, t := range ts {
		vars[i] = toGoParam(t)
	}
	return vars
}

func toGoTuple(ts []Type) *types.Tuple {
	vars := toGoParams(ts)
	return types.NewTuple(vars...)
}
