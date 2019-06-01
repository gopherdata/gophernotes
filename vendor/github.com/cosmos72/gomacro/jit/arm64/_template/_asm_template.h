#ifndef gomacro_asm_template_h
#define gomacro_asm_template_h

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
 * asm_template.h
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

#include <stdint.h>

typedef int8_t  i8;
typedef int16_t i16;
typedef int32_t i32;
typedef int64_t i64;

typedef uint8_t  u8;
typedef uint16_t u16;
typedef uint32_t u32;
typedef uint64_t u64;

#define Z (ints + 81)
#define A (ints + 82)
#define B (ints + 83)

#define z(l) (*(i##l *)Z)
#define a(l) (*(i##l *)A)
#define b(l) (*(i##l *)B)

#define uz(l) (*(u##l *)Z)
#define ua(l) (*(u##l *)A)
#define ub(l) (*(u##l *)B)

#if defined(__amd64) || defined(__amd64__) || defined(__i386) || defined(__i386__)
i64 _(i64 ax);
#else
# define _(arg) arg
#endif

#endif /* gomacro_asm_template_h */
