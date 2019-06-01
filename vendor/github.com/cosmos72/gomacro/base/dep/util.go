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
 * util.go
 *
 *  Created on: May 03, 2018
 *      Author: Massimiliano Ghilardi
 */

package dep

import (
	"sort"
)

// keep only items satisfying pred(item).
// destructively modifies list.
func filter_if_inplace(list []string, pred func(string) bool) []string {
	out := 0
	for _, e := range list {
		if pred(e) {
			list[out] = e
			out++
		}
	}
	return list[:out]
}

// remove all strings equal to 'str' from list
// destructively modifies list.
func remove_item_inplace(str string, list []string) []string {
	out := 0
	for _, e := range list {
		if e != str {
			list[out] = e
			out++
		}
	}
	return list[:out]
}

// make a copy of list
func dup(list []string) []string {
	if len(list) == 0 {
		return nil
	}
	ret := make([]string, len(list))
	copy(ret, list)
	return ret
}

// sort and remove duplicates from lists
func sort_unique_inplace(list []string) []string {
	if len(list) <= 1 {
		return list
	}
	sort.Strings(list)

	prev := list[0]
	out := 1

	// remove duplicates
	for _, e := range list[1:] {
		if e == prev {
			continue
		}
		prev = e
		list[out] = e
		out++
	}
	return list[:out]
}
