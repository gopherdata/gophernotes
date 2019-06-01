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
 * cache.go
 *
 *  Created on Feb 24, 2019
 *      Author Massimiliano Ghilardi
 */

package common

// map[len][crc32c][array of]executable machine code
type Cache map[int]map[uint32][]MemArea

func (cache Cache) Lookup(area MemArea) MemArea {
	// output.Debugf("cache lookup: %#04x %v", area.Checksum(), area)
	m := cache[area.Size()]
	if m != nil {
		v := m[area.Checksum()]
		if v != nil {
			for _, other := range v {
				if area.Equal(other) {
					// output.Debugf("cache hit:    %#04x %v", other.Checksum(), other)
					return other
				}
			}
		}
	}
	var ret MemArea // exploit zero value
	return ret
}

func (cache Cache) Add(area MemArea) {
	size := area.Size()
	m := cache[size]
	if m == nil {
		m = make(map[uint32][]MemArea)
		cache[size] = m
	}
	hash := area.Checksum()
	m[hash] = append(m[hash], area)
	// output.Debugf("cache add:    %#04x %v", hash, area)
}
