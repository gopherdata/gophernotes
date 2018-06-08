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
 * value.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"reflect"
)

type Value struct {
	reflect.Value
	XType Type
}

func (v Value) Kind() reflect.Kind {
	if v.XType == nil {
		return reflect.Invalid
	}
	return v.XType.Kind()
}

func (v Value) Type() Type {
	return v.XType
}

func (v Value) Convert(t Type) Value {
	return Value{
		v.Value.Convert(t.ReflectType()),
		t,
	}
}

func (v Value) FieldByName(name, pkgpath string) Value {
	field, count := v.XType.FieldByName(name, pkgpath)
	var w Value
	if count == 1 {
		w.Value = v.Value.FieldByIndex(field.Index)
		w.XType = field.Type
	}
	return w
}
