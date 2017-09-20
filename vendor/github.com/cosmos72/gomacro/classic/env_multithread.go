// +build !gomacro_classic_singlethread

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

	"golang.org/x/sync/syncmap"
)

const MultiThread = true

type BindMap syncmap.Map

// ALWAYS use pointers to syncmap.Map, because https://godoc.org/golang.org/x/sync/syncmap#Map
// states "A Map must not be copied after first use."
func (m *BindMap) Ensure() *BindMap {
	return m
}

func (m *BindMap) Clear() {
	*m = BindMap{}
}

func (m *BindMap) Merge(binds map[string]r.Value) {
	for k, v := range binds {
		(*syncmap.Map)(m).Store(k, v)
	}
}

func (m *BindMap) AsMap() map[string]r.Value {
	ret := make(map[string]r.Value)
	(*syncmap.Map)(m).Range(func(k, v interface{}) bool {
		ret[k.(string)] = v.(r.Value)
		return true
	})
	return ret
}

func (m *BindMap) Get(key string) (r.Value, bool) {
	val, ok := (*syncmap.Map)(m).Load(key)
	if !ok {
		return r.Value{}, false
	}
	return val.(r.Value), ok
}

func (m *BindMap) Get1(key string) r.Value {
	val, _ := (*syncmap.Map)(m).Load(key)
	return val.(r.Value)
}

func (m *BindMap) Set(key string, val r.Value) {
	(*syncmap.Map)(m).Store(key, val)
}

func (m *BindMap) Del(key string) {
	(*syncmap.Map)(m).Delete(key)
}

// -----------------------------------------

type TypeMap syncmap.Map

// ALWAYS use pointers to TypeMap, because https://godoc.org/golang.org/x/sync/syncmap#Map
// states "A Map must not be copied after first use."
func (m *TypeMap) Ensure() *TypeMap {
	return m
}

func (m *TypeMap) Clear() {
	*m = TypeMap{}
}

func (m *TypeMap) Merge(binds map[string]r.Type) {
	for k, v := range binds {
		(*syncmap.Map)(m).Store(k, v)
	}
}

func (m *TypeMap) AsMap() map[string]r.Type {
	ret := make(map[string]r.Type)
	(*syncmap.Map)(m).Range(func(k, v interface{}) bool {
		ret[k.(string)] = v.(r.Type)
		return true
	})
	return ret
}

func (m *TypeMap) Get(key string) (r.Type, bool) {
	val, ok := (*syncmap.Map)(m).Load(key)
	if !ok {
		return nil, false
	}
	return val.(r.Type), ok
}

func (m *TypeMap) Get1(key string) r.Type {
	val, _ := (*syncmap.Map)(m).Load(key)
	return val.(r.Type)
}

func (m *TypeMap) Set(key string, val r.Type) {
	(*syncmap.Map)(m).Store(key, val)
}

func (m *TypeMap) Del(key string) {
	(*syncmap.Map)(m).Delete(key)
}
