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
 * binary_arith_const.c
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */


#include "asm_template.h"

void IncInt8c(u64 *ints) {
    z(8) += 0x55;
}
void IncInt16c(u64 *ints) {
    z(16) += 0x5566;
}
void IncInt32c(u64 *ints) {
    z(32) += 0x55667788;
}
void IncInt64bc(u64 *ints) {
    z(64) += 0x55;
}
void IncInt64lc(u64 *ints) {
    z(64) += 0x55667788;
}
void IncInt64qc(u64 *ints) {
    z(64) += 0x5566778899aabbcc;
}

void AddInt8c(u64 *ints) {
    z(8) = a(8) + 0x55;
}
void AddInt16c(u64 *ints) {
    z(16) = a(16) + 0x5566;
}
void AddInt32c(u64 *ints) {
    z(32) = a(32) + 0x55667788;
}
void AddInt64bc(u64 *ints) {
    z(64) = a(64) + 0x55;
}
void AddInt64lc(u64 *ints) {
    z(64) = a(64) + 0x55667788;
}
void AddInt64qc(u64 *ints) {
    z(64) = a(64) + 0x5566778899aabbcc;
}

void DecInt8c(u64 *ints) {
    z(8) -= 0x55;
}
void DecInt16c(u64 *ints) {
    z(16) -= 0x5566;
}
void DecInt32c(u64 *ints) {
    z(32) -= 0x55667788;
}
void DecInt64bc(u64 *ints) {
    z(64) -= 0x55;
}
void DecInt64lc(u64 *ints) {
    z(64) -= 0x55667788;
}
void DecInt64qc(u64 *ints) {
    z(64) -= 0x5566778899aabbcc;
}

void SubInt8c(u64 *ints) {
    z(8) = a(8) - 0x55;
}
void SubInt16c(u64 *ints) {
    z(16) = a(16) - 0x5566;
}
void SubInt32c(u64 *ints) {
    z(32) = a(32) - 0x55667788;
}
void SubInt64bc(u64 *ints) {
    z(64) = a(64) - 0x55;
}
void SubInt64lc(u64 *ints) {
    z(64) = a(64) - 0x55667788;
}
void SubInt64qc(u64 *ints) {
    z(64) = a(64) - 0x5566778899aabbcc;
}


void MulInt8c(u64 *ints) {
    z(8) = a(8) * 0x55;
}
void MulInt16c(u64 *ints) {
    z(16) = a(16) * 0x5566;
}
void MulInt32c(u64 *ints) {
    z(32) = a(32) * 0x55667788;
}
void MulInt64bc(u64 *ints) {
    z(64) = a(64) * 0x55;
}
void MulInt64lc(u64 *ints) {
    z(64) = a(64) * 0x55667788;
}
void MulInt64qc(u64 *ints) {
    z(64) = a(64) * 0x5566778899aabbcc;
}

