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
 * binary_bitwise.c
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */


#include "asm_template.h"

i64 And_l_ax(u64 *ints, i64 ax) {
    return _(ax) & 0x55667788;
}
i64 And_q_ax(u64 *ints, i64 ax) {
    return _(ax) & 0x5566778899aabbccll;
}
i64 And_ax(u64 *ints, i64 ax) {
    return _(ax) & a(64);
}


i64 Or_l_ax(u64 *ints, i64 ax) {
    return _(ax) | 0x55667788;
}
i64 Or_q_ax(u64 *ints, i64 ax) {
    return _(ax) | 0x5566778899aabbccll;
}
i64 Or_ax(u64 *ints, i64 ax) {
    return _(ax) | a(64);
}


i64 Xor_l_ax(u64 *ints, i64 ax) {
    return _(ax) ^ 0x55667788;
}
i64 Xor_q_ax(u64 *ints, i64 ax) {
    return _(ax) ^ 0x5566778899aabbccll;
}
i64 Xor_ax(u64 *ints, i64 ax) {
    return _(ax) ^ a(64);
}



i64 Andnot_l_ax(u64 *ints, i64 ax) {
    return _(ax) & ~0x55667788;
}
i64 Andnot_q_ax(u64 *ints, i64 ax) {
    return _(ax) & ~0x5566778899aabbccll;
}
i64 Andnot_ax(u64 *ints, i64 ax) {
    return _(ax) & ~a(64);
}
