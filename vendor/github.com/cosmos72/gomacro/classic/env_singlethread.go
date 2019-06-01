// +build gomacro_classic_singlethread

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
 * env_singlethread.go
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
