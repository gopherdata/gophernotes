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
 * set.c
 *
 *  Created on May 27, 2018
 *      Author Massimiliano Ghilardi
 */

#include "../asm_template.h"

u64 load_16_a(void) {
	return 0;
}
u64 load_16_b(void) {
	return 0xff;
}
u64 load_16_c(void) {
	return 0xfff;
}
u64 load_16_d(void) {
	return 0xffff;
}
u64 load_32_a(void) {
	return 0x55555;
}
u64 load_32_b(void) {
	return 0x555555;
}
u64 load_32_c(void) {
	return 0x5555555;
}
u64 load_32_d(void) {
	return 0x55555555;
}

u64 load_48_a(void) {
	return 0x955552222;
}
u64 load_48_b(void) {
	return 0x9955552222;
}
u64 load_48_c(void) {
	return 0x99955552222;
}
u64 load_48_d(void) {
	return 0x999955552222;
}

u64 load_64_a(void) {
	return 0xa999955552222;
}
u64 load_64_b(void) {
	return 0xaa999955552222;
}
u64 load_64_c(void) {
	return 0xaaa999955552222;
}
u64 load_64_d(void) {
	return 0xaaaa999955552222;
}
