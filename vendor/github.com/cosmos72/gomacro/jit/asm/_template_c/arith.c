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
 * arith.c
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */


#include "asm_template.h"

i64 Add_l(i64 ax) {
    return _(ax) + 0x55667788;
}
i64 Add_q(i64 ax) {
    return _(ax) + 0x5566778899aabbcc;
}
i64 Add(i64 ax) {
    return _(ax) + a(64);
}


i64 Sub_l(i64 ax) {
    return _(ax) - 0x55667788;
}
i64 Sub_q(i64 ax) {
    return _(ax) - 0x5566778899aabbcc;
}
i64 Sub(i64 ax) {
    return _(ax) - a(64);
}


i64 Mul_l(i64 ax) {
    return _(ax) * 0x55667788;
}
i64 Mul_q(i64 ax) {
    return _(ax) * 0x5566778899aabbcc;
}
i64 Mul(i64 ax) {
    return _(ax) * a(64);
}


i64 Quo_l(i64 ax) {
    return _(ax) / 0x55667788;
}
i64 Quo_q(i64 ax) {
    return _(ax) / 0x5566778899aabbcc;
}
i64 Quo(i64 ax) {
    return _(ax) / a(64);
}


u64 QuoU(u64 ax) {
    return _(ax) / ua(64);
}


i64 Rem(i64 ax) {
    return _(ax) % a(64);
}
u64 RemU(u64 ax) {
    return _(ax) % ua(64);
}
