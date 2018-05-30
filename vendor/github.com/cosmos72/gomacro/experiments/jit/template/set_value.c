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
 * set_value.c
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */


#include "asm_template.h"

void ZeroInt8(u64 *ints) {
    z(8) = 0;
}
void ZeroInt16(u64 *ints) {
    z(16) = 0;
}
void ZeroInt32(u64 *ints) {
    z(32) = 0;
}
void ZeroInt64(u64 *ints) {
    z(64) = 0;
}


void SetInt8(u64 *ints) {
    z(8) = 0x55;
}
void SetInt16(u64 *ints) {
    z(16) = 0x5566;
}
void SetInt32(u64 *ints) {
    z(32) = 0x55667788;
}
void SetInt64b(u64 *ints) {
    z(64) = 0x55;
}
void SetInt64l(u64 *ints) {
    z(64) = 0x55667788;
}
void SetInt64q(u64 *ints) {
    z(64) = 0x5566778899aabbcc;
}

void SetUint8(u64 *ints) {
    uz(8) = 0x55;
}
void SetUint16(u64 *ints) {
    uz(16) = 0x5566;
}
void SetUint32(u64 *ints) {
    uz(32) = 0x55667788;
}
void SetUint64b(u64 *ints) {
    uz(64) = 0x55;
}
void SetUint64l(u64 *ints) {
    uz(64) = 0x55667788;
}
void SetUint64q(u64 *ints) {
    uz(64) = 0x5566778899aabbcc;
}

