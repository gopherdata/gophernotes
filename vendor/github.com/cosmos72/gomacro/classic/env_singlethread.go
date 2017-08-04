// +build gomacro_classic_singlethread

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
 * env.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	r "reflect"
)

const MultiThread = false

type BindMap map[string]r.Value

func (m *BindMap) Ensure() BindMap {
	if *m == nil {
		*m = make(map[string]r.Value)
	}
	return *m
}

func (m *BindMap) Clear() {
	*m = make(map[string]r.Value)
}

func (m BindMap) Merge(binds map[string]r.Value) {
	// make a copy. we do NOT want to modify binds!
	for k, v := range binds {
		m[k] = v
	}
}

func (m BindMap) AsMap() map[string]r.Value {
	return m
}

func (m BindMap) Get(key string) (r.Value, bool) {
	val, ok := m[key]
	return val, ok
}

func (m BindMap) Get1(key string) r.Value {
	return m[key]
}

func (m BindMap) Set(key string, val r.Value) {
	m[key] = val
}

func (m BindMap) Del(key string) {
	delete(m, key)
}

// -----------------------------------------

type TypeMap map[string]r.Type

func (m *TypeMap) Ensure() TypeMap {
	if *m == nil {
		*m = make(map[string]r.Type)
	}
	return *m
}

func (m *TypeMap) Clear() {
	*m = make(map[string]r.Type)
}

func (m TypeMap) Merge(types map[string]r.Type) {
	// make a copy. we do NOT want to modify types!
	for k, v := range types {
		m[k] = v
	}
}

func (m TypeMap) AsMap() map[string]r.Type {
	return m
}

func (m TypeMap) Get(key string) (r.Type, bool) {
	val, ok := m[key]
	return val, ok
}

func (m TypeMap) Set(key string, val r.Type) {
	m[key] = val
}

func (m TypeMap) Del(key string) {
	delete(m, key)
}
