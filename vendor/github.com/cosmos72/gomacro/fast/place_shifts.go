// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
						v = v.Convert(rt)
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
