// +build darwin dragonfly freebsd linux netbsd openbsd

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
 * mmap_unix.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"hash/crc32"

	"golang.org/x/sys/unix"
)

const (
	MMAP_SUPPORTED = true
)

// allocate memory in 64k chunks
var minAllocSize = unix.Getpagesize() * 16

type MemPool struct {
	bytes  []byte
	offset int
}

type MemArea []byte

func NewMemPool(size int) *MemPool {
	bytes, err := unix.Mmap(-1, 0,
		(size+minAllocSize-1)&^(minAllocSize-1),
		unix.PROT_READ,
		unix.MAP_ANON|unix.MAP_PRIVATE)
	if err != nil {
		errorf("sys/unix.Mmap failed: %v", err)
	}
	return &MemPool{bytes, 0}
}

func (pool *MemPool) Addr() *uint8 {
	if pool.Size() <= 0 {
		return nil
	}
	return &pool.bytes[pool.offset]
}

func (pool *MemPool) Size() int {
	if pool == nil {
		return 0
	}
	return len(pool.bytes) - pool.offset
}

func (pool *MemPool) protect(prot int) {
	err := unix.Mprotect(pool.bytes, prot)
	if err != nil {
		errorf("sys/unix.Mprotect failed: %v", err)
	}
}

func (pool *MemPool) SetReadonly() {
	pool.protect(unix.PROT_READ | unix.PROT_EXEC)
}

func (pool *MemPool) SetReadWrite() {
	pool.protect(unix.PROT_READ | unix.PROT_WRITE | unix.PROT_EXEC)
}

func (pool *MemPool) Copy(area MemArea) MemArea {
	size := area.Size()
	avail := pool.Size()
	if size > avail {
		errorf("MemArea is %d bytes, cannot copy to %d bytes MemPool", size, avail)
	}
	if MMAP_VERBOSE {
		debugf("copying %d bytes MemArea to MemPool{addr:%p, size:%d, offset:%d}",
			size, &pool.bytes[0], len(pool.bytes), pool.offset)
	}
	pool.SetReadWrite()
	copy(pool.bytes[pool.offset:], area)
	pool.SetReadonly()
	used := (size + 15) &^ 15
	if used >= avail {
		used = avail
	}
	ret := pool.bytes[pool.offset : pool.offset+size]
	pool.offset += used
	return ret
}

// convert MachineCode to MemArea
func (code MachineCode) MemArea() MemArea {
	return code.Bytes
}

func (area MemArea) Size() int {
	return len(area)
}

func (area MemArea) Equal(other MemArea) bool {
	return sliceEqual(area, other)
}

var crcTable = crc32.MakeTable(crc32.Castagnoli)

func (area MemArea) Checksum() uint32 {
	return crc32.Checksum(area, crcTable)
}
