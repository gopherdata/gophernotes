// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * cti_basic_method.go
 *
 *  Created on May 12, 2019
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	r "reflect"

	"github.com/cosmos72/gomacro/go/etoken"
)

func (v *Universe) addBasicTypesMethodsCTI() {
	if !etoken.GENERICS_V2_CTI {
		return
	}

	for _, t := range v.BasicTypes {
		if t != nil {
			v.addBasicTypeMethodsCTI(unwrap(t))
		}
	}

}
func (v *Universe) addBasicTypeMethodsCTI(xt *xtype) {
	if !etoken.GENERICS_V2_CTI {
		return
	}

	mvec := xt.GetMethods()
	switch xt.kind {
	case r.Bool:
		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a bool,
					b bool,
				) bool { return a == b })
			case "Not":
				(*mvec)[i] = r.ValueOf(func(z bool,

					a bool,

				) bool {
					return !a
				})
			}
		}
	case r.Int:

		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a int,

					b int,

				) bool { return a == b })
			case "Cmp":
				(*mvec)[i] = r.ValueOf(func(a int,

					b int,

				) int {
					if a < b {
						return -1
					}
					if a > b {
						return 1
					}
					return 0
				})
			case "Less":
				(*mvec)[i] = r.ValueOf(func(a int,

					b int,

				) bool { return a < b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z int,

					a int,

					b int,

				) int {
					return a + b
				})
			case "Sub":
				(*mvec)[i] = r.ValueOf(func(z int,

					a int,

					b int,

				) int {
					return a - b
				})
			case "Mul":
				(*mvec)[i] = r.ValueOf(func(z int,

					a int,

					b int,

				) int {
					return a * b
				})
			case "Quo":
				(*mvec)[i] = r.ValueOf(func(z int,

					a int,

					b int,

				) int {
					return a / b
				})
			case "Neg":
				(*mvec)[i] = r.ValueOf(func(z int,

					a int,

				) int {
					return -a
				})
			case "Rem":
				(*mvec)[i] = r.ValueOf(func(z int,

					a int,

					b int,

				) int {
					return a % b
				})
			case "And":
				(*mvec)[i] = r.ValueOf(func(z int,

					a int,

					b int,

				) int {
					return a & b
				})
			case "AndNot":
				(*mvec)[i] = r.ValueOf(func(z int,

					a int,

					b int,

				) int {
					return a &^ b
				})
			case "Or":
				(*mvec)[i] = r.ValueOf(func(z int,

					a int,

					b int,

				) int {
					return a | b
				})
			case "Xor":
				(*mvec)[i] = r.ValueOf(func(z int,

					a int,

					b int,

				) int {
					return a ^ b
				})
			case "Not":
				(*mvec)[i] = r.ValueOf(func(z int,

					a int,

				) int {
					return ^a
				})
			case "Lsh":
				(*mvec)[i] = r.ValueOf(func(z int,

					a int,

					b uint8) int {
					return a << b
				})
			case "Rsh":
				(*mvec)[i] = r.ValueOf(func(z int,

					a int,

					b uint8) int {
					return a >> b
				})
			}
		}
	case r.Int8:

		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a int8,

					b int8,

				) bool { return a == b })
			case "Cmp":
				(*mvec)[i] = r.ValueOf(func(a int8,

					b int8,

				) int {
					if a < b {
						return -1
					}
					if a > b {
						return 1
					}
					return 0
				})
			case "Less":
				(*mvec)[i] = r.ValueOf(func(a int8,

					b int8,

				) bool { return a < b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z int8,

					a int8,

					b int8,

				) int8 {
					return a + b
				})
			case "Sub":
				(*mvec)[i] = r.ValueOf(func(z int8,

					a int8,

					b int8,

				) int8 {
					return a - b
				})
			case "Mul":
				(*mvec)[i] = r.ValueOf(func(z int8,

					a int8,

					b int8,

				) int8 {
					return a * b
				})
			case "Quo":
				(*mvec)[i] = r.ValueOf(func(z int8,

					a int8,

					b int8,

				) int8 {
					return a / b
				})
			case "Neg":
				(*mvec)[i] = r.ValueOf(func(z int8,

					a int8,

				) int8 {
					return -a
				})
			case "Rem":
				(*mvec)[i] = r.ValueOf(func(z int8,

					a int8,

					b int8,

				) int8 {
					return a % b
				})
			case "And":
				(*mvec)[i] = r.ValueOf(func(z int8,

					a int8,

					b int8,

				) int8 {
					return a & b
				})
			case "AndNot":
				(*mvec)[i] = r.ValueOf(func(z int8,

					a int8,

					b int8,

				) int8 {
					return a &^ b
				})
			case "Or":
				(*mvec)[i] = r.ValueOf(func(z int8,

					a int8,

					b int8,

				) int8 {
					return a | b
				})
			case "Xor":
				(*mvec)[i] = r.ValueOf(func(z int8,

					a int8,

					b int8,

				) int8 {
					return a ^ b
				})
			case "Not":
				(*mvec)[i] = r.ValueOf(func(z int8,

					a int8,

				) int8 {
					return ^a
				})
			case "Lsh":
				(*mvec)[i] = r.ValueOf(func(z int8,

					a int8,

					b uint8) int8 {
					return a << b
				})
			case "Rsh":
				(*mvec)[i] = r.ValueOf(func(z int8,

					a int8,

					b uint8) int8 {
					return a >> b
				})
			}
		}
	case r.Int16:

		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a int16,

					b int16,

				) bool { return a == b })
			case "Cmp":
				(*mvec)[i] = r.ValueOf(func(a int16,

					b int16,

				) int {
					if a < b {
						return -1
					}
					if a > b {
						return 1
					}
					return 0
				})
			case "Less":
				(*mvec)[i] = r.ValueOf(func(a int16,

					b int16,

				) bool { return a < b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z int16,

					a int16,

					b int16,

				) int16 {
					return a + b
				})
			case "Sub":
				(*mvec)[i] = r.ValueOf(func(z int16,

					a int16,

					b int16,

				) int16 {
					return a - b
				})
			case "Mul":
				(*mvec)[i] = r.ValueOf(func(z int16,

					a int16,

					b int16,

				) int16 {
					return a * b
				})
			case "Quo":
				(*mvec)[i] = r.ValueOf(func(z int16,

					a int16,

					b int16,

				) int16 {
					return a / b
				})
			case "Neg":
				(*mvec)[i] = r.ValueOf(func(z int16,

					a int16,

				) int16 {
					return -a
				})
			case "Rem":
				(*mvec)[i] = r.ValueOf(func(z int16,

					a int16,

					b int16,

				) int16 {
					return a % b
				})
			case "And":
				(*mvec)[i] = r.ValueOf(func(z int16,

					a int16,

					b int16,

				) int16 {
					return a & b
				})
			case "AndNot":
				(*mvec)[i] = r.ValueOf(func(z int16,

					a int16,

					b int16,

				) int16 {
					return a &^ b
				})
			case "Or":
				(*mvec)[i] = r.ValueOf(func(z int16,

					a int16,

					b int16,

				) int16 {
					return a | b
				})
			case "Xor":
				(*mvec)[i] = r.ValueOf(func(z int16,

					a int16,

					b int16,

				) int16 {
					return a ^ b
				})
			case "Not":
				(*mvec)[i] = r.ValueOf(func(z int16,

					a int16,

				) int16 {
					return ^a
				})
			case "Lsh":
				(*mvec)[i] = r.ValueOf(func(z int16,

					a int16,

					b uint8) int16 {
					return a << b
				})
			case "Rsh":
				(*mvec)[i] = r.ValueOf(func(z int16,

					a int16,

					b uint8) int16 {
					return a >> b
				})
			}
		}
	case r.Int32:

		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a int32,

					b int32,

				) bool { return a == b })
			case "Cmp":
				(*mvec)[i] = r.ValueOf(func(a int32,

					b int32,

				) int {
					if a < b {
						return -1
					}
					if a > b {
						return 1
					}
					return 0
				})
			case "Less":
				(*mvec)[i] = r.ValueOf(func(a int32,

					b int32,

				) bool { return a < b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z int32,

					a int32,

					b int32,

				) int32 {
					return a + b
				})
			case "Sub":
				(*mvec)[i] = r.ValueOf(func(z int32,

					a int32,

					b int32,

				) int32 {
					return a - b
				})
			case "Mul":
				(*mvec)[i] = r.ValueOf(func(z int32,

					a int32,

					b int32,

				) int32 {
					return a * b
				})
			case "Quo":
				(*mvec)[i] = r.ValueOf(func(z int32,

					a int32,

					b int32,

				) int32 {
					return a / b
				})
			case "Neg":
				(*mvec)[i] = r.ValueOf(func(z int32,

					a int32,

				) int32 {
					return -a
				})
			case "Rem":
				(*mvec)[i] = r.ValueOf(func(z int32,

					a int32,

					b int32,

				) int32 {
					return a % b
				})
			case "And":
				(*mvec)[i] = r.ValueOf(func(z int32,

					a int32,

					b int32,

				) int32 {
					return a & b
				})
			case "AndNot":
				(*mvec)[i] = r.ValueOf(func(z int32,

					a int32,

					b int32,

				) int32 {
					return a &^ b
				})
			case "Or":
				(*mvec)[i] = r.ValueOf(func(z int32,

					a int32,

					b int32,

				) int32 {
					return a | b
				})
			case "Xor":
				(*mvec)[i] = r.ValueOf(func(z int32,

					a int32,

					b int32,

				) int32 {
					return a ^ b
				})
			case "Not":
				(*mvec)[i] = r.ValueOf(func(z int32,

					a int32,

				) int32 {
					return ^a
				})
			case "Lsh":
				(*mvec)[i] = r.ValueOf(func(z int32,

					a int32,

					b uint8) int32 {
					return a << b
				})
			case "Rsh":
				(*mvec)[i] = r.ValueOf(func(z int32,

					a int32,

					b uint8) int32 {
					return a >> b
				})
			}
		}
	case r.Int64:
		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a int64,

					b int64,

				) bool { return a == b })
			case "Cmp":
				(*mvec)[i] = r.ValueOf(func(a int64,

					b int64,

				) int {
					if a < b {
						return -1
					}
					if a > b {
						return 1
					}
					return 0
				})
			case "Less":
				(*mvec)[i] = r.ValueOf(func(a int64,

					b int64,

				) bool { return a < b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z int64,

					a int64,

					b int64,

				) int64 {
					return a + b
				})
			case "Sub":
				(*mvec)[i] = r.ValueOf(func(z int64,

					a int64,

					b int64,

				) int64 {
					return a - b
				})
			case "Mul":
				(*mvec)[i] = r.ValueOf(func(z int64,

					a int64,

					b int64,

				) int64 {
					return a * b
				})
			case "Quo":
				(*mvec)[i] = r.ValueOf(func(z int64,

					a int64,

					b int64,

				) int64 {
					return a / b
				})
			case "Neg":
				(*mvec)[i] = r.ValueOf(func(z int64,

					a int64,

				) int64 {
					return -a
				})
			case "Rem":
				(*mvec)[i] = r.ValueOf(func(z int64,

					a int64,

					b int64,

				) int64 {
					return a % b
				})
			case "And":
				(*mvec)[i] = r.ValueOf(func(z int64,

					a int64,

					b int64,

				) int64 {
					return a & b
				})
			case "AndNot":
				(*mvec)[i] = r.ValueOf(func(z int64,

					a int64,

					b int64,

				) int64 {
					return a &^ b
				})
			case "Or":
				(*mvec)[i] = r.ValueOf(func(z int64,

					a int64,

					b int64,

				) int64 {
					return a | b
				})
			case "Xor":
				(*mvec)[i] = r.ValueOf(func(z int64,

					a int64,

					b int64,

				) int64 {
					return a ^ b
				})
			case "Not":
				(*mvec)[i] = r.ValueOf(func(z int64,

					a int64,

				) int64 {
					return ^a
				})
			case "Lsh":
				(*mvec)[i] = r.ValueOf(func(z int64,

					a int64,

					b uint8) int64 {
					return a << b
				})
			case "Rsh":
				(*mvec)[i] = r.ValueOf(func(z int64,

					a int64,

					b uint8) int64 {
					return a >> b
				})
			}
		}
	case r.Uint:
		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a uint,
					b uint,
				) bool { return a == b })
			case "Cmp":
				(*mvec)[i] = r.ValueOf(func(a uint,

					b uint,

				) int {
					if a < b {
						return -1
					}
					if a > b {
						return 1
					}
					return 0
				})
			case "Less":
				(*mvec)[i] = r.ValueOf(func(a uint,

					b uint,

				) bool { return a < b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z uint,

					a uint,

					b uint,

				) uint {
					return a + b
				})
			case "Sub":
				(*mvec)[i] = r.ValueOf(func(z uint,

					a uint,

					b uint,

				) uint {
					return a - b
				})
			case "Mul":
				(*mvec)[i] = r.ValueOf(func(z uint,

					a uint,

					b uint,

				) uint {
					return a * b
				})
			case "Quo":
				(*mvec)[i] = r.ValueOf(func(z uint,

					a uint,

					b uint,

				) uint {
					return a / b
				})
			case "Neg":
				(*mvec)[i] = r.ValueOf(func(z uint,

					a uint,

				) uint {
					return -a
				})
			case "Rem":
				(*mvec)[i] = r.ValueOf(func(z uint,

					a uint,

					b uint,

				) uint {
					return a % b
				})
			case "And":
				(*mvec)[i] = r.ValueOf(func(z uint,

					a uint,

					b uint,

				) uint {
					return a & b
				})
			case "AndNot":
				(*mvec)[i] = r.ValueOf(func(z uint,

					a uint,

					b uint,

				) uint {
					return a &^ b
				})
			case "Or":
				(*mvec)[i] = r.ValueOf(func(z uint,

					a uint,

					b uint,

				) uint {
					return a | b
				})
			case "Xor":
				(*mvec)[i] = r.ValueOf(func(z uint,

					a uint,

					b uint,

				) uint {
					return a ^ b
				})
			case "Not":
				(*mvec)[i] = r.ValueOf(func(z uint,

					a uint,

				) uint {
					return ^a
				})
			case "Lsh":
				(*mvec)[i] = r.ValueOf(func(z uint,

					a uint,

					b uint8) uint {
					return a << b
				})
			case "Rsh":
				(*mvec)[i] = r.ValueOf(func(z uint,

					a uint,

					b uint8) uint {
					return a >> b
				})
			}
		}
	case r.Uint8:
		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a uint8, b uint8) bool { return a == b })
			case "Cmp":
				(*mvec)[i] = r.ValueOf(func(a uint8,
					b uint8,
				) int {
					if a < b {
						return -1
					}
					if a > b {
						return 1
					}
					return 0
				})
			case "Less":
				(*mvec)[i] = r.ValueOf(func(a uint8,

					b uint8,

				) bool { return a < b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z uint8,

					a uint8,

					b uint8,

				) uint8 {
					return a + b
				})
			case "Sub":
				(*mvec)[i] = r.ValueOf(func(z uint8,

					a uint8,

					b uint8,

				) uint8 {
					return a - b
				})
			case "Mul":
				(*mvec)[i] = r.ValueOf(func(z uint8,

					a uint8,

					b uint8,

				) uint8 {
					return a * b
				})
			case "Quo":
				(*mvec)[i] = r.ValueOf(func(z uint8,

					a uint8,

					b uint8,

				) uint8 {
					return a / b
				})
			case "Neg":
				(*mvec)[i] = r.ValueOf(func(z uint8,

					a uint8,

				) uint8 {
					return -a
				})
			case "Rem":
				(*mvec)[i] = r.ValueOf(func(z uint8,

					a uint8,

					b uint8,

				) uint8 {
					return a % b
				})
			case "And":
				(*mvec)[i] = r.ValueOf(func(z uint8,

					a uint8,

					b uint8,

				) uint8 {
					return a & b
				})
			case "AndNot":
				(*mvec)[i] = r.ValueOf(func(z uint8,

					a uint8,

					b uint8,

				) uint8 {
					return a &^ b
				})
			case "Or":
				(*mvec)[i] = r.ValueOf(func(z uint8,

					a uint8,

					b uint8,

				) uint8 {
					return a | b
				})
			case "Xor":
				(*mvec)[i] = r.ValueOf(func(z uint8,

					a uint8,

					b uint8,

				) uint8 {
					return a ^ b
				})
			case "Not":
				(*mvec)[i] = r.ValueOf(func(z uint8,

					a uint8,

				) uint8 {
					return ^a
				})
			case "Lsh":
				(*mvec)[i] = r.ValueOf(func(z uint8,

					a uint8,

					b uint8) uint8 {
					return a << b
				})
			case "Rsh":
				(*mvec)[i] = r.ValueOf(func(z uint8,

					a uint8,

					b uint8) uint8 {
					return a >> b
				})
			}
		}
	case r.Uint16:
		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a uint16, b uint16) bool { return a == b })
			case "Cmp":
				(*mvec)[i] = r.ValueOf(func(a uint16, b uint16) int {
					if a < b {
						return -1
					}
					if a > b {
						return 1
					}
					return 0
				})
			case "Less":
				(*mvec)[i] = r.ValueOf(func(a uint16,
					b uint16,
				) bool { return a < b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z uint16,

					a uint16,

					b uint16,

				) uint16 {
					return a + b
				})
			case "Sub":
				(*mvec)[i] = r.ValueOf(func(z uint16,

					a uint16,

					b uint16,

				) uint16 {
					return a - b
				})
			case "Mul":
				(*mvec)[i] = r.ValueOf(func(z uint16,

					a uint16,

					b uint16,

				) uint16 {
					return a * b
				})
			case "Quo":
				(*mvec)[i] = r.ValueOf(func(z uint16,

					a uint16,

					b uint16,

				) uint16 {
					return a / b
				})
			case "Neg":
				(*mvec)[i] = r.ValueOf(func(z uint16,

					a uint16,

				) uint16 {
					return -a
				})
			case "Rem":
				(*mvec)[i] = r.ValueOf(func(z uint16,

					a uint16,

					b uint16,

				) uint16 {
					return a % b
				})
			case "And":
				(*mvec)[i] = r.ValueOf(func(z uint16,

					a uint16,

					b uint16,

				) uint16 {
					return a & b
				})
			case "AndNot":
				(*mvec)[i] = r.ValueOf(func(z uint16,

					a uint16,

					b uint16,

				) uint16 {
					return a &^ b
				})
			case "Or":
				(*mvec)[i] = r.ValueOf(func(z uint16,

					a uint16,

					b uint16,

				) uint16 {
					return a | b
				})
			case "Xor":
				(*mvec)[i] = r.ValueOf(func(z uint16,

					a uint16,

					b uint16,

				) uint16 {
					return a ^ b
				})
			case "Not":
				(*mvec)[i] = r.ValueOf(func(z uint16,

					a uint16,

				) uint16 {
					return ^a
				})
			case "Lsh":
				(*mvec)[i] = r.ValueOf(func(z uint16,

					a uint16,

					b uint8) uint16 {
					return a << b
				})
			case "Rsh":
				(*mvec)[i] = r.ValueOf(func(z uint16,

					a uint16,

					b uint8) uint16 {
					return a >> b
				})
			}
		}
	case r.Uint32:
		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a uint32, b uint32) bool { return a == b })
			case "Cmp":
				(*mvec)[i] = r.ValueOf(func(a uint32, b uint32) int {
					if a < b {
						return -1
					}
					if a > b {
						return 1
					}
					return 0
				})
			case "Less":
				(*mvec)[i] = r.ValueOf(func(a uint32, b uint32) bool { return a < b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z uint32,
					a uint32,
					b uint32,
				) uint32 {
					return a + b
				})
			case "Sub":
				(*mvec)[i] = r.ValueOf(func(z uint32,

					a uint32,

					b uint32,

				) uint32 {
					return a - b
				})
			case "Mul":
				(*mvec)[i] = r.ValueOf(func(z uint32,

					a uint32,

					b uint32,

				) uint32 {
					return a * b
				})
			case "Quo":
				(*mvec)[i] = r.ValueOf(func(z uint32,

					a uint32,

					b uint32,

				) uint32 {
					return a / b
				})
			case "Neg":
				(*mvec)[i] = r.ValueOf(func(z uint32,

					a uint32,

				) uint32 {
					return -a
				})
			case "Rem":
				(*mvec)[i] = r.ValueOf(func(z uint32,

					a uint32,

					b uint32,

				) uint32 {
					return a % b
				})
			case "And":
				(*mvec)[i] = r.ValueOf(func(z uint32,

					a uint32,

					b uint32,

				) uint32 {
					return a & b
				})
			case "AndNot":
				(*mvec)[i] = r.ValueOf(func(z uint32,

					a uint32,

					b uint32,

				) uint32 {
					return a &^ b
				})
			case "Or":
				(*mvec)[i] = r.ValueOf(func(z uint32,

					a uint32,

					b uint32,

				) uint32 {
					return a | b
				})
			case "Xor":
				(*mvec)[i] = r.ValueOf(func(z uint32,

					a uint32,

					b uint32,

				) uint32 {
					return a ^ b
				})
			case "Not":
				(*mvec)[i] = r.ValueOf(func(z uint32,

					a uint32,

				) uint32 {
					return ^a
				})
			case "Lsh":
				(*mvec)[i] = r.ValueOf(func(z uint32,

					a uint32,

					b uint8) uint32 {
					return a << b
				})
			case "Rsh":
				(*mvec)[i] = r.ValueOf(func(z uint32,

					a uint32,

					b uint8) uint32 {
					return a >> b
				})
			}
		}
	case r.Uint64:
		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a uint64, b uint64) bool { return a == b })
			case "Cmp":
				(*mvec)[i] = r.ValueOf(func(a uint64, b uint64) int {
					if a < b {
						return -1
					}
					if a > b {
						return 1
					}
					return 0
				})
			case "Less":
				(*mvec)[i] = r.ValueOf(func(a uint64, b uint64) bool { return a < b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z uint64, a uint64, b uint64) uint64 {
					return a + b
				})
			case "Sub":
				(*mvec)[i] = r.ValueOf(func(z uint64,
					a uint64,
					b uint64,
				) uint64 {
					return a - b
				})
			case "Mul":
				(*mvec)[i] = r.ValueOf(func(z uint64,

					a uint64,

					b uint64,

				) uint64 {
					return a * b
				})
			case "Quo":
				(*mvec)[i] = r.ValueOf(func(z uint64,

					a uint64,

					b uint64,

				) uint64 {
					return a / b
				})
			case "Neg":
				(*mvec)[i] = r.ValueOf(func(z uint64,

					a uint64,

				) uint64 {
					return -a
				})
			case "Rem":
				(*mvec)[i] = r.ValueOf(func(z uint64,

					a uint64,

					b uint64,

				) uint64 {
					return a % b
				})
			case "And":
				(*mvec)[i] = r.ValueOf(func(z uint64,

					a uint64,

					b uint64,

				) uint64 {
					return a & b
				})
			case "AndNot":
				(*mvec)[i] = r.ValueOf(func(z uint64,

					a uint64,

					b uint64,

				) uint64 {
					return a &^ b
				})
			case "Or":
				(*mvec)[i] = r.ValueOf(func(z uint64,

					a uint64,

					b uint64,

				) uint64 {
					return a | b
				})
			case "Xor":
				(*mvec)[i] = r.ValueOf(func(z uint64,

					a uint64,

					b uint64,

				) uint64 {
					return a ^ b
				})
			case "Not":
				(*mvec)[i] = r.ValueOf(func(z uint64,

					a uint64,

				) uint64 {
					return ^a
				})
			case "Lsh":
				(*mvec)[i] = r.ValueOf(func(z uint64,

					a uint64,

					b uint8) uint64 {
					return a << b
				})
			case "Rsh":
				(*mvec)[i] = r.ValueOf(func(z uint64,

					a uint64,

					b uint8) uint64 {
					return a >> b
				})
			}
		}
	case r.Uintptr:
		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a uintptr, b uintptr) bool { return a == b })
			case "Cmp":
				(*mvec)[i] = r.ValueOf(func(a uintptr, b uintptr) int {
					if a < b {
						return -1
					}
					if a > b {
						return 1
					}
					return 0
				})
			case "Less":
				(*mvec)[i] = r.ValueOf(func(a uintptr, b uintptr) bool { return a < b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z uintptr, a uintptr, b uintptr) uintptr {
					return a + b
				})
			case "Sub":
				(*mvec)[i] = r.ValueOf(func(z uintptr, a uintptr, b uintptr) uintptr {
					return a - b
				})
			case "Mul":
				(*mvec)[i] = r.ValueOf(func(z uintptr,
					a uintptr,
					b uintptr,
				) uintptr {
					return a * b
				})
			case "Quo":
				(*mvec)[i] = r.ValueOf(func(z uintptr,

					a uintptr,

					b uintptr,

				) uintptr {
					return a / b
				})
			case "Neg":
				(*mvec)[i] = r.ValueOf(func(z uintptr,

					a uintptr,

				) uintptr {
					return -a
				})
			case "Rem":
				(*mvec)[i] = r.ValueOf(func(z uintptr,

					a uintptr,

					b uintptr,

				) uintptr {
					return a % b
				})
			case "And":
				(*mvec)[i] = r.ValueOf(func(z uintptr,

					a uintptr,

					b uintptr,

				) uintptr {
					return a & b
				})
			case "AndNot":
				(*mvec)[i] = r.ValueOf(func(z uintptr,

					a uintptr,

					b uintptr,

				) uintptr {
					return a &^ b
				})
			case "Or":
				(*mvec)[i] = r.ValueOf(func(z uintptr,

					a uintptr,

					b uintptr,

				) uintptr {
					return a | b
				})
			case "Xor":
				(*mvec)[i] = r.ValueOf(func(z uintptr,

					a uintptr,

					b uintptr,

				) uintptr {
					return a ^ b
				})
			case "Not":
				(*mvec)[i] = r.ValueOf(func(z uintptr,

					a uintptr,

				) uintptr {
					return ^a
				})
			case "Lsh":
				(*mvec)[i] = r.ValueOf(func(z uintptr,

					a uintptr,

					b uint8) uintptr {
					return a << b
				})
			case "Rsh":
				(*mvec)[i] = r.ValueOf(func(z uintptr,

					a uintptr,

					b uint8) uintptr {
					return a >> b
				})
			}
		}
	case r.Float32:

		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a float32,

					b float32,

				) bool { return a == b })
			case "Cmp":
				(*mvec)[i] = r.ValueOf(func(a float32,

					b float32,

				) int {
					if a < b {
						return -1
					}
					if a > b {
						return 1
					}
					return 0
				})
			case "Less":
				(*mvec)[i] = r.ValueOf(func(a float32,

					b float32,

				) bool { return a < b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z float32,

					a float32,

					b float32,

				) float32 {
					return a + b
				})
			case "Sub":
				(*mvec)[i] = r.ValueOf(func(z float32,

					a float32,

					b float32,

				) float32 {
					return a - b
				})
			case "Mul":
				(*mvec)[i] = r.ValueOf(func(z float32,

					a float32,

					b float32,

				) float32 {
					return a * b
				})
			case "Quo":
				(*mvec)[i] = r.ValueOf(func(z float32,

					a float32,

					b float32,

				) float32 {
					return a / b
				})
			case "Neg":
				(*mvec)[i] = r.ValueOf(func(z float32,

					a float32,

				) float32 {
					return -a
				})
			}
		}
	case r.Float64:

		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a float64,

					b float64,

				) bool { return a == b })
			case "Cmp":
				(*mvec)[i] = r.ValueOf(func(a float64,

					b float64,

				) int {
					if a < b {
						return -1
					}
					if a > b {
						return 1
					}
					return 0
				})
			case "Less":
				(*mvec)[i] = r.ValueOf(func(a float64,

					b float64,

				) bool { return a < b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z float64,

					a float64,

					b float64,

				) float64 {
					return a + b
				})
			case "Sub":
				(*mvec)[i] = r.ValueOf(func(z float64,

					a float64,

					b float64,

				) float64 {
					return a - b
				})
			case "Mul":
				(*mvec)[i] = r.ValueOf(func(z float64,

					a float64,

					b float64,

				) float64 {
					return a * b
				})
			case "Quo":
				(*mvec)[i] = r.ValueOf(func(z float64,

					a float64,

					b float64,

				) float64 {
					return a / b
				})
			case "Neg":
				(*mvec)[i] = r.ValueOf(func(z float64,

					a float64,

				) float64 {
					return -a
				})
			}
		}
	case r.Complex64:

		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a complex64,

					b complex64,

				) bool { return a == b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z complex64,

					a complex64,

					b complex64,

				) complex64 {
					return a + b
				})
			case "Sub":
				(*mvec)[i] = r.ValueOf(func(z complex64,

					a complex64,

					b complex64,

				) complex64 {
					return a - b
				})
			case "Mul":
				(*mvec)[i] = r.ValueOf(func(z complex64,

					a complex64,

					b complex64,

				) complex64 {
					return a * b
				})
			case "Quo":
				(*mvec)[i] = r.ValueOf(func(z complex64,

					a complex64,

					b complex64,

				) complex64 {
					return a / b
				})
			case "Neg":
				(*mvec)[i] = r.ValueOf(func(z complex64,

					a complex64,

				) complex64 {
					return -a
				})
			case "Real":
				(*mvec)[i] = r.ValueOf(func(a complex64,

				) float32 {
					return real(a)
				})
			case "Imag":
				(*mvec)[i] = r.ValueOf(func(a complex64,

				) float32 {
					return imag(a)
				})
			}
		}
	case r.Complex128:

		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a complex128,

					b complex128,

				) bool { return a == b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z complex128,

					a complex128,

					b complex128,

				) complex128 {
					return a + b
				})
			case "Sub":
				(*mvec)[i] = r.ValueOf(func(z complex128,

					a complex128,

					b complex128,

				) complex128 {
					return a - b
				})
			case "Mul":
				(*mvec)[i] = r.ValueOf(func(z complex128,

					a complex128,

					b complex128,

				) complex128 {
					return a * b
				})
			case "Quo":
				(*mvec)[i] = r.ValueOf(func(z complex128,

					a complex128,

					b complex128,

				) complex128 {
					return a / b
				})
			case "Neg":
				(*mvec)[i] = r.ValueOf(func(z complex128,

					a complex128,

				) complex128 {
					return -a
				})
			case "Real":
				(*mvec)[i] = r.ValueOf(func(a complex128,

				) float64 {
					return real(a)
				})
			case "Imag":
				(*mvec)[i] = r.ValueOf(func(a complex128,

				) float64 {
					return imag(a)
				})
			}
		}
	case r.String:

		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.ValueOf(func(a string,

					b string,

				) bool { return a == b })
			case "Cmp":
				(*mvec)[i] = r.ValueOf(func(a string,

					b string,

				) int {
					if a < b {
						return -1
					}
					if a > b {
						return 1
					}
					return 0
				})
			case "Less":
				(*mvec)[i] = r.ValueOf(func(a string,

					b string,

				) bool { return a < b })
			case "Add":
				(*mvec)[i] = r.ValueOf(func(z string,

					a string,

					b string,

				) string {
					return a + b
				})
			case "Index":
				(*mvec)[i] = r.ValueOf(func(a string,

					b int) uint8 { return a[b] })
			case "Len":
				(*mvec)[i] = r.ValueOf(func(a string,

				) int { return len(a) })
			case "Slice":
				(*mvec)[i] = r.ValueOf(func(a string,

					b int, c int) string {
					return a[b:c]
				})
			}
		}
	}
}
