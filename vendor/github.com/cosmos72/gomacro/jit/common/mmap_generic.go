// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!windows

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
 * mmap_generic.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"runtime"
)

const MMAP_SUPPORTED = false

type MemPool struct {
}

type MemArea struct {
}

func (pool *MemPool) Size() int {
	return 0
}

func NewMemPool(size int) *MemPool {
	errorf("MemPool: unsupported operating system %v, cannot create executable memory", runtime.GOOS)
	return nil
}

func (pool *MemPool) SetReadonly() {
}

func (pool *MemPool) SetReadWrite() {
}

func (pool *MemPool) Copy(area MemArea) MemArea {
	errorf("MemPool: unsupported operating system %v, cannot copy machine code to executable memory", runtime.GOOS)
	return MemArea{}
}

// convert MachineCode to MemArea
func (code MachineCode) MemArea() MemArea {
	return MemArea{}
}

func (area MemArea) Addr() *uint8 {
	return nil
}

func (area MemArea) Size() int {
	return 0
}

func (area MemArea) Equal(other MemArea) bool {
	return false
}

func (area MemArea) Checksum() uint32 {
	return 0
}
