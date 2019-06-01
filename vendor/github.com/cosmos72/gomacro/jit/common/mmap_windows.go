// +build windows

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
 * mmap_windows.go
 *
 *  Created on May 25, 2018
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"hash/crc32"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	MMAP_SUPPORTED = true
)

var (
	// scavenge for win32 FlushInstructionCache()
	dllkernel32               = windows.NewLazySystemDLL("kernel32.dll")
	procFlushInstructionCache = dllkernel32.NewProc("FlushInstructionCache")

	// allocate memory in 4k chunks
	// because FlushInstructionCache() seems to have no effect
	minAllocSize = uintptr(windows.Getpagesize())
)

func flushInstructionCache(addr uintptr, size uintptr) {
	ret, _, err := syscall.Syscall(procFlushInstructionCache.Addr(), 3, ^uintptr(0), addr, size)
	if ret == 0 && err != 0 {
		errorf("win32 FlushInstructionCache() failed: %v", err)
	}
}

// use *uint8 instead of uintptr to avoid garbage collector
// freeing a MemArea created from Go-allocated memory
type ptr struct {
	x *uint8
}

func intptr(addr uintptr) ptr {
	return ptr{(*uint8)(unsafe.Pointer(addr))}
}

func (p ptr) int() uintptr {
	return uintptr(unsafe.Pointer(p.x))
}

func (p ptr) add(offset uintptr) ptr {
	return ptr{(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p.x)) + offset))}
}

func (p ptr) uint8(offset uintptr) *uint8 {
	return (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p.x)) + offset))
}

func (p ptr) uint64(offset uintptr) *uint64 {
	return (*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p.x)) + offset))
}

type MemPool struct {
	ptr          ptr
	size, offset uintptr
}

type MemArea struct {
	ptr  ptr
	size uintptr
}

func NewMemPool(size int) *MemPool {
	poolsize := (uintptr(size) + minAllocSize - 1) &^ (minAllocSize - 1)
	addr, err := windows.VirtualAlloc(0, poolsize,
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_READONLY)
	if err != nil {
		errorf("sys/windows.VirtualAlloc failed: %v", err)
	}
	return &MemPool{intptr(addr), poolsize, 0}
}

func (pool *MemPool) Addr() *uint8 {
	if pool == nil {
		return nil
	}
	return pool.ptr.uint8(pool.offset)
}

func (pool *MemPool) Size() int {
	if pool == nil {
		return 0
	}
	return int(pool.size - pool.offset)
}

func (pool *MemPool) protect(prot uint32) {
	var old uint32
	err := windows.VirtualProtect(pool.ptr.int(), pool.size, prot, &old)
	if err != nil {
		errorf("sys/windows.VirtualProtect failed: %v", err)
	}
}

func (pool *MemPool) SetReadonly() {
	pool.protect(windows.PAGE_EXECUTE_READ)
	flushInstructionCache(pool.ptr.int(), pool.size)
}

func (pool *MemPool) SetReadWrite() {
	pool.protect(windows.PAGE_EXECUTE_READWRITE)
}

func (pool *MemPool) Copy(area MemArea) MemArea {
	size := area.size
	avail := uintptr(pool.Size())
	if size > avail {
		errorf("MemArea is %d bytes, cannot copy to %d bytes MemPool", size, avail)
	}
	if MMAP_VERBOSE {
		debugf("copying %d bytes MemArea to MemPool{addr:%#x, size:%d, offset:%d}",
			size, pool.ptr.int(), pool.size, pool.offset)
	}
	pool.SetReadWrite()
	memcpy(pool.ptr, area.ptr, size)
	pool.SetReadonly()
	used := (size + 15) &^ 15
	if used >= avail {
		used = avail
	}
	ret := MemArea{pool.ptr.add(pool.offset), size}
	// consume all pool, because FlushInstructionCache
	// seems to have no effect
	// pool.offset += used
	pool.offset = pool.size
	return ret
}

// memory copy. a bit slow, but avoids depending on CGO
func memcpy(dst ptr, src ptr, size uintptr) {
	var i uintptr
	for ; i+32 <= size; i += 32 {
		*dst.uint64(i + 0) = *src.uint64(i + 0)
		*dst.uint64(i + 8) = *src.uint64(i + 8)
		*dst.uint64(i + 16) = *src.uint64(i + 16)
		*dst.uint64(i + 24) = *src.uint64(i + 24)
	}
	for ; i+8 <= size; i += 8 {
		*dst.uint64(i) = *src.uint64(i)
	}
	for ; i < size; i++ {
		*dst.uint8(i) = *src.uint8(i)
	}
}

// memory comparison. a bit slow, but avoids depending on CGO
func memcmp(lhs ptr, rhs ptr, size uintptr) int {
	if lhs == rhs || size == 0 {
		return 0
	}
	var i uintptr
	for ; i+8 <= size; i += 8 {
		l := *lhs.uint64(i)
		r := *rhs.uint64(i)
		if l < r {
			return -1
		} else if l > r {
			return 1
		}
	}
	for ; i < size; i++ {
		l := *lhs.uint8(i)
		r := *rhs.uint8(i)
		if l < r {
			return -1
		} else if l > r {
			return 1
		}
	}
	return 0
}

// convert MachineCode to MemArea
func (code MachineCode) MemArea() MemArea {
	size := uintptr(len(code.Bytes))
	var area MemArea
	if size != 0 {
		area.ptr = ptr{&code.Bytes[0]}
		area.size = size
	}
	return area
}

func (area MemArea) Size() int {
	return int(area.size)
}

func (area MemArea) Equal(other MemArea) bool {
	size := area.size
	if size != other.size {
		return false
	}
	if size == 0 {
		return true
	}
	return memcmp(area.ptr, other.ptr, size) == 0
}

var crcTable = crc32.MakeTable(crc32.Castagnoli)

func (area MemArea) Checksum() uint32 {
	// cannot use crc32.Checksum(): we do not have a []uint8 slice
	crc := ^uint32(0)
	p := area.ptr
	for i := uintptr(0); i < area.size; i++ {
		index := uint8(crc) ^ *p.uint8(i)
		crc = crcTable[index] ^ crc>>8
	}
	return ^crc
}
