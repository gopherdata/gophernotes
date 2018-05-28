// +build !gomacro_classic_singlethread

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
