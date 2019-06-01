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
 * spinlock.go
 *
 *  Created on: Apr 30 2018
 *      Author: Massimiliano Ghilardi
 */

package atomic

import (
	"runtime"
	"sync/atomic"
)

type SpinLock int32

func (s *SpinLock) Lock() {
	for i := 0; i < 10; i++ {
		if atomic.CompareAndSwapInt32((*int32)(s), 0, 1) {
			return
		}
	}
	for !atomic.CompareAndSwapInt32((*int32)(s), 0, 1) {
		runtime.Gosched()
	}
}

func (s *SpinLock) Unlock() {
	atomic.StoreInt32((*int32)(s), 0)
}
