/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * fib_asm.go
 *
 *  Created on May 23, 2018
 *      Author Massimiliano Ghilardi
 */

// +build gc

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"
#include "../../src/runtime/go_tls.h"

// assembler will automatically save/restore BP and adjust SP on function entry/exit
// it will also call runtime.morestack() as needed
TEXT ·fib_asm(SB),0,$16-16
	NO_LOCAL_POINTERS

	MOVQ	n+0(FP), AX

	MOVQ	AX, 0(SP) // omit symbol name, otherwise assembler will add some stack offset
	CALL	·fib(SB)
	MOVQ	8(SP), AX // idem

	MOVQ	AX, ret+8(FP)
	RET
