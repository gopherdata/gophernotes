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
 * sys_unix.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

var PAGESIZE = unix.Getpagesize()

func nop(*uint64) {
}

func (asm *Asm) Func() func(*uint64) {
	if len(asm.code) == 0 {
		return nop
	}
	asm.epilogue()
	if VERBOSE {
		fmt.Printf("asm: %#v\n", asm.code)
	}
	mem, err := unix.Mmap(-1, 0, (len(asm.code)+PAGESIZE-1)&^(PAGESIZE-1),
		unix.PROT_READ|unix.PROT_WRITE, unix.MAP_ANON|unix.MAP_PRIVATE)
	if err != nil {
		errorf("sys/unix.Mmap failed: %v", err)
	}
	copy(mem, asm.code)
	err = unix.Mprotect(mem, unix.PROT_EXEC|unix.PROT_READ)
	if err != nil {
		unix.Munmap(mem)
		errorf("sys/unix.Mprotect failed: %v", err)
	}
	var f func(*uint64)
	*(**[]uint8)(unsafe.Pointer(&f)) = &mem
	// runtime.SetFinalizer(&f, munmap)
	return f
}

func munmap(obj interface{}) {
	f, ok := obj.(func(*uint64))
	if ok && f != nil {
		mem := **(**[]uint8)(unsafe.Pointer(&f))
		unix.Munmap(mem)
	}
}
