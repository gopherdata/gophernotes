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
 * _extend.c
 *
 *  Created on Feb 08, 2019
 *      Author Massimiliano Ghilardi
 */

#include <stdint.h>

void extend8_16(uint8_t a, uint16_t *b) {
	*b = a;
}

void extend8_32(uint8_t a, uint32_t *b) {
	*b = a;
}

void extend8_64(uint8_t a, uint64_t *b) {
	*b = a;
}

void extend16_32(uint16_t a, uint32_t *b) {
	*b = a;
}

void extend16_64(uint16_t a, uint64_t *b) {
	*b = a;
}

void extend32_64(uint32_t a, uint64_t *b) {
	*b = a;
}


void sextend8_16(int8_t a, int16_t *b) {
	*b = a;
}

void sextend8_32(int8_t a, int32_t *b) {
	*b = a;
}

void sextend8_64(int8_t a, int64_t *b) {
	*b = a;
}

void sextend16_32(int16_t a, int32_t *b) {
	*b = a;
}

void sextend16_64(int16_t a, int64_t *b) {
	*b = a;
}

void sextend32_64(int32_t a, int64_t *b) {
	*b = a;
}


