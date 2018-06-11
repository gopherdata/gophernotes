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
 * place_shifts.go
 *
 *  Created on May 17, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) placeShlConst(place *Place, val I) {
	if isLiteralNumber(val, 0) {
		c.placeForSideEffects(place)
		return
	}

	{
		var ret Stmt
		lhsfun := place.Fun
		keyfun := place.MapKey
		val := r.ValueOf(val).Uint()

		t := place.Type
		rt := t.ReflectType()
		cat := KindToCategory(t.Kind())
		if keyfun == nil {
			switch cat {
			case r.Int:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Int()
					lhs.SetInt(result <<
						val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Uint()
					lhs.SetUint(result <<
						val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		} else {
			switch cat {
			case r.Int:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Int()

					v := r.ValueOf(result <<
						val,
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Uint()

					v := r.ValueOf(result <<
						val,
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		}

		if ret == nil {
			c.Errorf("invalid operator %s= on <%v>", token.SHL, place.Type)
		}

		c.append(ret)
	}
}
func (c *Comp) placeShlExpr(place *Place, fun I) {
	var ret Stmt
	lhsfun := place.Fun
	keyfun := place.MapKey
	t := place.Type
	rt := t.ReflectType()
	cat := KindToCategory(t.Kind())
	if keyfun == nil {
		switch cat {
		case r.Int:

			switch fun := fun.(type) {
			case func(*Env) uint:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Int()
					lhs.SetInt(result <<
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint8:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Int()
					lhs.SetInt(result <<
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint16:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Int()
					lhs.SetInt(result <<
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint32:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Int()
					lhs.SetInt(result <<
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint64:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Int()
					lhs.SetInt(result <<
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uintptr:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Int()
					lhs.SetInt(result <<
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:

			switch fun := fun.(type) {
			case func(*Env) uint:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Uint()
					lhs.SetUint(result <<
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint8:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Uint()
					lhs.SetUint(result <<
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint16:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Uint()
					lhs.SetUint(result <<
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint32:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Uint()
					lhs.SetUint(result <<
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint64:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Uint()
					lhs.SetUint(result <<
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uintptr:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Uint()
					lhs.SetUint(result <<
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
	} else {
		switch cat {
		case r.Int:
			switch fun := fun.(type) {
			case func(*Env) uint:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Int()

					v := r.ValueOf(result <<
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint8:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Int()

					v := r.ValueOf(result <<
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint16:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Int()

					v := r.ValueOf(result <<
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint32:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Int()

					v := r.ValueOf(result <<
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint64:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Int()

					v := r.ValueOf(result <<
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uintptr:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Int()

					v := r.ValueOf(result <<
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			switch fun := fun.(type) {
			case func(*Env) uint:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Uint()

					v := r.ValueOf(result <<
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint8:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Uint()

					v := r.ValueOf(result <<
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint16:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Uint()

					v := r.ValueOf(result <<
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint32:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Uint()

					v := r.ValueOf(result <<
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint64:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Uint()

					v := r.ValueOf(result <<
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uintptr:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Uint()

					v := r.ValueOf(result <<
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
	}

	if ret == nil {
		c.Errorf("invalid operator %s= on <%v>", token.SHL, place.Type)
	}

	c.append(ret)
}
func (c *Comp) placeShrConst(place *Place, val I) {
	if isLiteralNumber(val, 0) {
		c.placeForSideEffects(place)
		return
	}

	{
		var ret Stmt
		lhsfun := place.Fun
		keyfun := place.MapKey
		val := r.ValueOf(val).Uint()

		t := place.Type
		rt := t.ReflectType()
		cat := KindToCategory(t.Kind())
		if keyfun == nil {
			switch cat {
			case r.Int:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Int()
					lhs.SetInt(result >>
						val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Uint()
					lhs.SetUint(result >>
						val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		} else {
			switch cat {
			case r.Int:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Int()

					v := r.ValueOf(result >>
						val,
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Uint()

					v := r.ValueOf(result >>
						val,
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		}

		if ret == nil {
			c.Errorf("invalid operator %s= on <%v>", token.SHR, place.Type)
		}

		c.append(ret)
	}
}
func (c *Comp) placeShrExpr(place *Place, fun I) {
	var ret Stmt
	lhsfun := place.Fun
	keyfun := place.MapKey
	t := place.Type
	rt := t.ReflectType()
	cat := KindToCategory(t.Kind())
	if keyfun == nil {
		switch cat {
		case r.Int:

			switch fun := fun.(type) {
			case func(*Env) uint:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Int()
					lhs.SetInt(result >>
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint8:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Int()
					lhs.SetInt(result >>
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint16:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Int()
					lhs.SetInt(result >>
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint32:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Int()
					lhs.SetInt(result >>
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint64:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Int()
					lhs.SetInt(result >>
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uintptr:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Int()
					lhs.SetInt(result >>
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:

			switch fun := fun.(type) {
			case func(*Env) uint:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Uint()
					lhs.SetUint(result >>
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint8:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Uint()
					lhs.SetUint(result >>
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint16:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Uint()
					lhs.SetUint(result >>
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint32:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Uint()
					lhs.SetUint(result >>
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint64:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Uint()
					lhs.SetUint(result >>
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uintptr:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Uint()
					lhs.SetUint(result >>
						fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
	} else {
		switch cat {
		case r.Int:
			switch fun := fun.(type) {
			case func(*Env) uint:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Int()

					v := r.ValueOf(result >>
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint8:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Int()

					v := r.ValueOf(result >>
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint16:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Int()

					v := r.ValueOf(result >>
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint32:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Int()

					v := r.ValueOf(result >>
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint64:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Int()

					v := r.ValueOf(result >>
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uintptr:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Int()

					v := r.ValueOf(result >>
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case r.Uint:
			switch fun := fun.(type) {
			case func(*Env) uint:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Uint()

					v := r.ValueOf(result >>
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint8:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Uint()

					v := r.ValueOf(result >>
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint16:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Uint()

					v := r.ValueOf(result >>
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint32:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Uint()

					v := r.ValueOf(result >>
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uint64:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Uint()

					v := r.ValueOf(result >>
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case func(*Env) uintptr:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Uint()

					v := r.ValueOf(result >>
						fun(env),
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
	}

	if ret == nil {
		c.Errorf("invalid operator %s= on <%v>", token.SHR, place.Type)
	}

	c.append(ret)
}
func (c *Comp) placeQuoPow2(place *Place, val I) bool {
	if isLiteralNumber(val, 0) {
		c.Errorf("division by %v <%v>", val, r.TypeOf(val))
		return false
	} else if isLiteralNumber(val, 1) {
		c.placeForSideEffects(place)
		return true
	}

	ypositive := true
	yv := r.ValueOf(val)
	ycat := KindToCategory(yv.Kind())
	var y uint64
	switch ycat {
	case r.Int:
		sy := yv.Int()
		if sy < 0 {
			ypositive = false
			y = uint64(-sy)
		} else {
			y = uint64(sy)
		}

	case r.Uint:
		y = yv.Uint()
	default:
		return false
	}
	if !isPowerOfTwo(y) {
		return false
	}

	shift := integerLen(y) - 1

	if !ypositive {
		return false
	}

	var roundup int64
	if ycat == r.Int {
		roundup = int64(y - 1)
	}
	{
		var ret Stmt
		lhsfun := place.Fun
		keyfun := place.MapKey
		val := shift

		t := place.Type
		rt := t.ReflectType()
		cat := KindToCategory(t.Kind())
		if keyfun == nil {
			switch cat {
			case r.Int:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Int()

					if result < 0 {
						result += roundup
					}
					lhs.SetInt(result >>
						val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					result := lhs.Uint()
					lhs.SetUint(result >>
						val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		} else {
			switch cat {
			case r.Int:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Int()

					if result < 0 {
						result += roundup
					}

					v := r.ValueOf(result >>
						val,
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			case r.Uint:

				ret = func(env *Env) (Stmt, *Env) {
					lhs := lhsfun(env)
					key := keyfun(env)
					result := lhs.MapIndex(key).Uint()

					v := r.ValueOf(result >>
						val,
					)
					if v.Type() != rt {
						v = convert(v, rt)
					}

					lhs.SetMapIndex(key, v)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		}

		if ret == nil {
			c.Errorf("invalid operator %s= on <%v>", token.QUO, place.Type)
		}

		c.append(ret)
	}
	return true
}
