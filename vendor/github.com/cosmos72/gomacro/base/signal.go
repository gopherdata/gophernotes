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
 * signal.go
 *
 *  Created on: Apr 14, 2018
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"unsafe"
)

// =======================================================================

func StartSignalHandler(handler func(os.Signal)) chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go signalHandlerGoroutine(c, handler)
	return c
}

func StopSignalHandler(c chan os.Signal) {
	close(c)
}

func signalHandlerGoroutine(c chan os.Signal, handler func(os.Signal)) {
	for {
		sig, ok := <-c
		if !ok {
			break
		}
		if handler != nil {
			handler(sig)
		}
	}
}

// =======================================================================

type Signal uint8

const (
	SigDefer Signal = 1 << iota // request to install a defer function
	SigReturn
	SigInterrupt // user pressed Ctrl+C, process received SIGINT, or similar
	SigDebug     // debugger asked to execute in single-step mode

	SigNone = Signal(0) // no signal
	SigAll  = ^SigNone  // mask of all possible signals
)

func (sig Signal) String() string {
	var s string
	switch sig {
	case SigNone:
		s = "// signal: none"
	case SigDefer:
		s = "// signal: defer"
	case SigReturn:
		s = "// signal: return"
	case SigInterrupt:
		s = "// signal: interrupt"
	case SigDebug:
		s = "// signal: debug"
	default:
		s = fmt.Sprintf("// signal: unknown(%d)", uint16(sig))
	}
	return s
}

type Signals struct {
	Sync  Signal
	Debug Signal
	Async Signal
	_     Signal
}

func (s *Signals) IsEmpty() bool {
	return atomic.LoadUint32((*uint32)(unsafe.Pointer(s))) == 0
}
