// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * switch2.go
 *
 *  Created on May 06, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) switchGotoMap(tag *Expr, seen *caseHelper, ip int) {
	if len(seen.GotoMap) <= 1 {
		return
	}

	var stmt Stmt
	switch efun := tag.Fun.(type) {
	case func(*Env) bool:
		m := [2]int{-1, -1}
		for k, v := range seen.GotoMap {
			if r.ValueOf(k).Bool() {
				m[1] = v.IP
			} else {
				m[0] = v.IP
			}
		}

		stmt = func(env *Env) (Stmt, *Env) {
			val := efun(env)
			var ip int
			if val {
				ip = m[1]
			} else {
				ip = m[0]
			}

			if ip >= 0 {
				env.IP = ip
			} else {
				env.IP++
			}
			return env.Code[env.IP], env
		}
	case func(*Env) int:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[int]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[int(r.ValueOf(k).Int())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) int8:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[int8]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[int8(r.ValueOf(k).Int())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) int16:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[int16]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[int16(r.ValueOf(k).Int())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) int32:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[int32]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[int32(r.ValueOf(k).Int())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) int64:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[int64]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[r.ValueOf(k).Int()] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) uint:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[uint]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[uint(r.ValueOf(k).Uint())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) uint8:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[uint8]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[uint8(r.ValueOf(k).Uint())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) uint16:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[uint16]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[uint16(r.ValueOf(k).Uint())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) uint32:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[uint32]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[uint32(r.ValueOf(k).Uint())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) uint64:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[uint64]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[r.ValueOf(k).Uint()] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) uintptr:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[uintptr]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[uintptr(r.ValueOf(k).Uint())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) float32:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[float32]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[float32(r.ValueOf(k).Float())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) float64:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[float64]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[r.ValueOf(k).Float()] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) complex64:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[complex64]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[complex64(r.ValueOf(k).Complex())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) complex128:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[complex128]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[r.ValueOf(k).Complex()] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) string:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[string]int, len(seen.GotoMap))
			for k, v := range seen.GotoMap {
				m[r.ValueOf(k).String()] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) (r.Value, []r.Value):
		m := make(map[interface{}]int, len(seen.GotoMap))
		for k, v := range seen.GotoMap {
			m[k] = v.IP
		}

		stmt = func(env *Env) (Stmt, *Env) {
			v, _ := efun(env)
			if ip, ok := m[v.Interface()]; ok {
				env.IP = ip
			} else {
				env.IP++
			}
			return env.Code[env.IP], env
		}
	default:
		fun := tag.AsX1()
		m := make(map[interface{}]int, len(seen.GotoMap))
		for k, v := range seen.GotoMap {
			m[k] = v.IP
		}

		stmt = func(env *Env) (Stmt, *Env) {
			val := fun(env).Interface()
			if ip, ok := m[val]; ok {
				env.IP = ip
			} else {
				env.IP++
			}
			return env.Code[env.IP], env
		}
	}
	if stmt == nil {
		return
	}

	c.Code.List[ip] = stmt
}
func (c *Comp) switchGotoSlice(tag *Expr, seen *caseHelper) Stmt {
	var stmt Stmt
	switch efun := tag.Fun.(type) {
	case func(*Env) int:
		{
			var min, max int
			for k := range seen.GotoMap {
				key := int(r.ValueOf(k).Int())
				min = key
				max = key
				break
			}
			for k := range seen.GotoMap {
				key := int(r.ValueOf(k).Int())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.GotoMap) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.GotoMap) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.GotoMap))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.GotoMap {
				key := int(r.ValueOf(k).Int())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		{
			var min, max int8
			for k := range seen.GotoMap {
				key := int8(r.ValueOf(k).Int())
				min = key
				max = key
				break
			}
			for k := range seen.GotoMap {
				key := int8(r.ValueOf(k).Int())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.GotoMap) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.GotoMap) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.GotoMap))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.GotoMap {
				key := int8(r.ValueOf(k).Int())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		{
			var min, max int16
			for k := range seen.GotoMap {
				key := int16(r.ValueOf(k).Int())
				min = key
				max = key
				break
			}
			for k := range seen.GotoMap {
				key := int16(r.ValueOf(k).Int())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.GotoMap) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.GotoMap) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.GotoMap))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.GotoMap {
				key := int16(r.ValueOf(k).Int())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		{
			var min, max int32
			for k := range seen.GotoMap {
				key :=
					int32(r.ValueOf(k).Int())
				min = key
				max = key
				break
			}
			for k := range seen.GotoMap {
				key := int32(r.ValueOf(k).Int())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.GotoMap) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.GotoMap) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.GotoMap))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.GotoMap {
				key := int32(r.ValueOf(k).Int())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		{
			var min, max int64
			for k := range seen.GotoMap {
				key := r.ValueOf(k).Int()
				min = key
				max = key
				break
			}
			for k := range seen.GotoMap {
				key := r.ValueOf(k).Int()
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.GotoMap) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.GotoMap) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.GotoMap))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.GotoMap {
				key := r.ValueOf(k).Int()

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		{
			var min, max uint
			for k := range seen.GotoMap {
				key :=

					uint(r.ValueOf(k).Uint())
				min = key
				max = key
				break
			}
			for k := range seen.GotoMap {
				key := uint(r.ValueOf(k).Uint())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.GotoMap) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.GotoMap) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.GotoMap))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.GotoMap {
				key := uint(r.ValueOf(k).Uint())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		{
			var min, max uint8
			for k := range seen.GotoMap {
				key :=

					uint8(r.ValueOf(k).Uint())
				min = key
				max = key
				break
			}
			for k := range seen.GotoMap {
				key := uint8(r.ValueOf(k).Uint())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.GotoMap) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.GotoMap) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.GotoMap))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.GotoMap {
				key := uint8(r.ValueOf(k).Uint())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		{
			var min, max uint16
			for k := range seen.GotoMap {
				key :=

					uint16(r.ValueOf(k).Uint())
				min = key
				max = key
				break
			}
			for k := range seen.GotoMap {
				key := uint16(r.ValueOf(k).Uint())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.GotoMap) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.GotoMap) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.GotoMap))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.GotoMap {
				key := uint16(r.ValueOf(k).Uint())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		{
			var min, max uint32
			for k := range seen.GotoMap {
				key :=

					uint32(r.ValueOf(k).Uint())
				min = key
				max = key
				break
			}
			for k := range seen.GotoMap {
				key := uint32(r.ValueOf(k).Uint())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.GotoMap) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.GotoMap) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.GotoMap))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.GotoMap {
				key := uint32(r.ValueOf(k).Uint())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		{
			var min, max uint64
			for k := range seen.GotoMap {
				key := r.ValueOf(k).Uint()
				min = key
				max = key
				break
			}
			for k := range seen.GotoMap {
				key := r.ValueOf(k).Uint()
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.GotoMap) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.GotoMap) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.GotoMap))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.GotoMap {
				key := r.ValueOf(k).Uint()

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		{
			var min, max uintptr
			for k := range seen.GotoMap {
				key :=

					uintptr(r.ValueOf(k).Uint())
				min = key
				max = key
				break
			}
			for k := range seen.GotoMap {
				key :=

					uintptr(r.ValueOf(k).Uint())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.GotoMap) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.GotoMap) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.GotoMap))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.GotoMap {
				key := uintptr(r.ValueOf(k).Uint())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	}
	return stmt
}
