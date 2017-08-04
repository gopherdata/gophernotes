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
