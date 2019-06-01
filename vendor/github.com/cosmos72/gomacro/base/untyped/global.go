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
 * global.go
 *
 *  Created on Apr 25, 2018
 *      Author Massimiliano Ghilardi
 */

package untyped

import (
	"fmt"
	"go/constant"
	r "reflect"

	"github.com/cosmos72/gomacro/base/reflect"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// untyped kind. matches reflect.Kind except for rune literals, where Kind == reflect.Int32
type Kind r.Kind

const (
	None    = Kind(r.Invalid)
	Bool    = Kind(r.Bool)
	Int     = Kind(r.Int)
	Rune    = Kind(r.Int32)
	Float   = Kind(r.Float64)
	Complex = Kind(r.Complex128)
	String  = Kind(r.String)
)

func (k Kind) String() string {
	var s string
	switch k {
	case None:
		s = "nil"
	case Rune:
		s = "rune"
	default:
		s = r.Kind(k).String()
	}
	return s
}

func (k Kind) Category() r.Kind {
	return reflect.Category(r.Kind(k))
}

// Lit represents an untyped literal value, i.e. an untyped constant
type Lit struct {
	Kind       Kind // untyped constant's default type
	Val        constant.Value
	basicTypes *[]xr.Type
}

func MakeLit(kind Kind, val constant.Value, basicTypes *[]xr.Type) Lit {
	return Lit{kind, val, basicTypes}
}

// pretty-print untyped constants
func (untyp Lit) String() string {
	val := untyp.Val
	var strobj interface{}
	if untyp.Kind == Rune && val.Kind() == constant.Int {
		if i, exact := constant.Int64Val(val); exact {
			if i >= 0 && i <= 0x10FFFF {
				strobj = fmt.Sprintf("%q", i)
			}
		}
	}
	if strobj == nil {
		strobj = val.ExactString()
	}
	return fmt.Sprintf("{%v %v}", untyp.Kind, strobj)
}

func MakeKind(ckind constant.Kind) Kind {
	ret := None
	switch ckind {
	case constant.Bool:
		ret = Bool
	case constant.Int:
		ret = Int // actually ambiguous, could be a Rune
	case constant.Float:
		ret = Float
	case constant.Complex:
		ret = Complex
	case constant.String:
		ret = String
	}
	return ret
}
