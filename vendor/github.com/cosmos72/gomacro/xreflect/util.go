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
 * util.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"fmt"
	"go/token"
	r "reflect"

	"github.com/cosmos72/gomacro/go/types"
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

func dirToGdir(dir r.ChanDir) types.ChanDir {
	var gdir types.ChanDir
	switch dir {
	case r.RecvDir:
		gdir = types.RecvOnly
	case r.SendDir:
		gdir = types.SendOnly
	case r.BothDir:
		gdir = types.SendRecv
	}
	return gdir
}

func gtypeToKind(t *xtype, gtype types.Type) r.Kind {
	gtype = gtype.Underlying()
	var kind r.Kind
	switch gtype := gtype.(type) {
	case *types.Array:
		kind = r.Array
	case *types.Basic:
		kind = ToReflectKind(gtype.Kind())
	case *types.Chan:
		kind = r.Chan
	case *types.Signature:
		kind = r.Func
	case *types.Interface:
		kind = r.Interface
	case *types.Map:
		kind = r.Map
	case *types.Pointer:
		kind = r.Ptr
	case *types.Slice:
		kind = r.Slice
	case *types.Struct:
		kind = r.Struct
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

func ToReflectKind(gkind types.BasicKind) r.Kind {
	var kind r.Kind
	switch gkind {
	case types.Bool, types.UntypedBool:
		kind = r.Bool
	case types.Int, types.UntypedInt:
		kind = r.Int
	case types.Int8:
		kind = r.Int8
	case types.Int16:
		kind = r.Int16
	case types.Int32, types.UntypedRune:
		kind = r.Int32
	case types.Int64:
		kind = r.Int64
	case types.Uint:
		kind = r.Uint
	case types.Uint8:
		kind = r.Uint8
	case types.Uint16:
		kind = r.Uint16
	case types.Uint32:
		kind = r.Uint32
	case types.Uint64:
		kind = r.Uint64
	case types.Uintptr:
		kind = r.Uintptr
	case types.Float32:
		kind = r.Float32
	case types.Float64, types.UntypedFloat:
		kind = r.Float64
	case types.Complex64:
		kind = r.Complex64
	case types.Complex128, types.UntypedComplex:
		kind = r.Complex128
	case types.String, types.UntypedString:
		kind = r.String
	case types.UnsafePointer:
		kind = r.UnsafePointer
	case types.UntypedNil:
		kind = r.Invalid
	default:
		errorf(nil, "unsupported types.BasicKind: %v", gkind)
	}
	return kind
}

func ToBasicKind(kind r.Kind, untyped bool) types.BasicKind {
	var gkind types.BasicKind
	switch kind {
	case r.Bool:
		if untyped {
			gkind = types.UntypedBool
		} else {
			gkind = types.Bool
		}
	case r.Int:
		if untyped {
			gkind = types.Int
		} else {
			gkind = types.UntypedInt
		}
	case r.Int8:
		gkind = types.Int8
	case r.Int16:
		gkind = types.Int16
	case r.Int32:
		if untyped {
			gkind = types.UntypedRune
		} else {
			gkind = types.Int32
		}
	case r.Int64:
		gkind = types.Int64
	case r.Uint:
		gkind = types.Uint
	case r.Uint8:
		gkind = types.Uint8
	case r.Uint16:
		gkind = types.Uint16
	case r.Uint32:
		gkind = types.Uint32
	case r.Uint64:
		gkind = types.Uint64
	case r.Uintptr:
		gkind = types.Uintptr
	case r.Float32:
		gkind = types.Float32
	case r.Float64:
		if untyped {
			gkind = types.UntypedFloat
		} else {
			gkind = types.Float64
		}
	case r.Complex64:
		gkind = types.Complex64
	case r.Complex128:
		if untyped {
			gkind = types.UntypedComplex
		} else {
			gkind = types.Complex128
		}
	case r.String:
		if untyped {
			gkind = types.UntypedString
		} else {
			gkind = types.String
		}
	case r.UnsafePointer:
		gkind = types.UnsafePointer
	case r.Invalid:
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

func toReflectTypes(ts []Type) []r.Type {
	rts := make([]r.Type, len(ts))
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
