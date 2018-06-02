// Copyright 2018 Massimiliano Ghilardi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gc

#include "go_asm.h"
#include "textflag.h" // for NOSPLIT
#include "../../src/runtime/go_tls.h"

TEXT ·GoID(SB),NOSPLIT,$0-8
	MOVD g, goid+0(FP)
	RET
