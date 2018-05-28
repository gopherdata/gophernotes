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

	xr "github.com/cosmos72/gomacro/xreflect"
)

// Lit represents an untyped literal value, i.e. an untyped constant
type Lit struct {
	Kind       r.Kind // default type. matches Val.Kind() except for rune literals, where Kind == reflect.Int32
	Val        constant.Value
	basicTypes *[]xr.Type
}

func MakeLit(kind r.Kind, val constant.Value, basicTypes *[]xr.Type) Lit {
	return Lit{kind, val, basicTypes}
}

// pretty-print untyped constants
func (untyp Lit) String() string {
	val := untyp.Val
	var strkind, strobj interface{} = untyp.Kind, nil
	if untyp.Kind == r.Int32 {
		strkind = "rune"
		if val.Kind() == constant.Int {
			if i, exact := constant.Int64Val(val); exact {
				if i >= 0 && i <= 0x10FFFF {
					strobj = fmt.Sprintf("%q", i)
				}
			}
		}
	}
	if strobj == nil {
		strobj = val.ExactString()
	}
	return fmt.Sprintf("{%v %v}", strkind, strobj)
}

func ConstantKindToUntypedLitKind(ckind constant.Kind) r.Kind {
	ret := r.Invalid
	switch ckind {
	case constant.Bool:
		ret = r.Bool
	case constant.Int:
		ret = r.Int // actually ambiguous, could be a rune - thus r.Int32
	case constant.Float:
		ret = r.Float64
	case constant.Complex:
		ret = r.Complex128
	case constant.String:
		ret = r.String
	}
	return ret
}
