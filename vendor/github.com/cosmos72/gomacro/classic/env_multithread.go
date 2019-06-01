// +build !gomacro_classic_singlethread

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
 * env_multithread.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	r "reflect"

	"sync"
)

const MultiThread = true

type BindMap struct {
	l sync.RWMutex
	m map[string]r.Value
}

// ALWAYS use pointers to BindMap, because it contains a sync.RWMutex
// and https://golang.org/pkg/sync/#RWMutex states "A RWMutex must not be copied after first use."
func (x *BindMap) Ensure() *BindMap {
	x.l.RLock()
	m := x.m
	x.l.RUnlock()
	if m != nil {
		return x
	}
	m = make(map[string]r.Value)
	x.l.Lock()
	if x.m == nil {
		x.m = m
	}
	x.l.Unlock()
	return x
}

func (x *BindMap) Clear() {
	x.l.Lock()
	x.m = make(map[string]r.Value)
	x.l.Unlock()
}

func (x *BindMap) Merge(binds map[string]r.Value) {
	// make a copy. we do NOT want to modify binds!
	x.l.Lock()
	m := x.m
	for k, v := range binds {
		m[k] = v
	}
	x.l.Unlock()
}

func (x *BindMap) AsMap() map[string]r.Value {
	out := make(map[string]r.Value)
	x.l.RLock()
	for k, v := range x.m {
		out[k] = v
	}
	x.l.RUnlock()
	return out
}

func (x *BindMap) Get(key string) (r.Value, bool) {
	x.l.RLock()
	val, ok := x.m[key]
	x.l.RUnlock()
	return val, ok
}

func (x *BindMap) Get1(key string) r.Value {
	x.l.RLock()
	val := x.m[key]
	x.l.RUnlock()
	return val
}

func (x *BindMap) Set(key string, val r.Value) {
	x.l.Lock()
	x.m[key] = val
	x.l.Unlock()
}

func (x *BindMap) Del(key string) {
	x.l.Lock()
	delete(x.m, key)
	x.l.Unlock()
}

// -----------------------------------------

type TypeMap struct {
	l sync.RWMutex
	m map[string]r.Type
}

func (x *TypeMap) Ensure() *TypeMap {
	x.l.RLock()
	m := x.m
	x.l.RUnlock()
	if m != nil {
		return x
	}
	m = make(map[string]r.Type)
	x.l.Lock()
	if x.m == nil {
		x.m = m
	}
	x.l.Unlock()
	return x
}

func (x *TypeMap) Clear() {
	x.l.Lock()
	x.m = make(map[string]r.Type)
	x.l.Unlock()
}

func (x *TypeMap) Merge(types map[string]r.Type) {
	// make a copy. we do NOT want to modify types!
	x.l.Lock()
	m := x.m
	for k, v := range types {
		m[k] = v
	}
	x.l.Unlock()
}

func (x *TypeMap) AsMap() map[string]r.Type {
	out := make(map[string]r.Type)
	x.l.RLock()
	for k, t := range x.m {
		out[k] = t
	}
	x.l.RUnlock()
	return out
}

func (x *TypeMap) Get(key string) (r.Type, bool) {
	x.l.RLock()
	val, ok := x.m[key]
	x.l.RUnlock()
	return val, ok
}

func (x *TypeMap) Set(key string, val r.Type) {
	x.l.Lock()
	x.m[key] = val
	x.l.Unlock()
}

func (x *TypeMap) Del(key string) {
	x.l.Lock()
	delete(x.m, key)
	x.l.Unlock()
}
