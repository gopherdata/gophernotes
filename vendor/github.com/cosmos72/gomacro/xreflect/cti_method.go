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
 * cti_methods.go
 *
 *  Created on May 12, 2019
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	r "reflect"

	"github.com/cosmos72/gomacro/go/etoken"
)

var vtrue = r.ValueOf(true)
var vfalse = r.ValueOf(false)
var vvtruefalse = map[bool][]r.Value{true: {vtrue}, false: {vfalse}}

type cmp uint8

const (
	cmp_less cmp = iota
	cmp_eq
	cmp_gtr
)

var vvcmp = [...][]r.Value{cmp_less: {r.ValueOf(-1)}, cmp_eq: {r.ValueOf(0)}, cmp_gtr: {r.ValueOf(1)}}

// declare CTI methods on Array, Chan, Map, Slice

func (v *Universe) addTypeMethodsCTI(xt *xtype) {
	if !etoken.GENERICS_V2_CTI {
		return
	}
	rt := xt.rtype
	if rt == nil {
		return
	}
	k := xt.kind
	if k == r.Invalid {
		// forward-declared type?
		k = rt.Kind()
	}
	switch k {
	case r.Bool, r.Int, r.Int8, r.Int16, r.Int32, r.Int64,
		r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr,
		r.Float32, r.Float64, r.Complex64, r.Complex128, r.String:
		if xt.rtype == rbasictypes[k] {
			// use optimized implementations in cti_basic_method.go
			v.addBasicTypeMethodsCTI(xt)
		} else {
			v.addBasicTypeReflectMethodsCTI(xt)
		}
		return
	case r.Array, r.Chan, r.Map, r.Slice:
		break
	default:
		return
	}
	n := xt.NumExplicitMethod()
	if n == 0 {
		return
	}
	rbool := rbasictypes[r.Bool]
	rint := rbasictypes[r.Int]

	rkey := rint
	if k == r.Map {
		rkey = rt.Key()
	}

	var relem r.Type
	if k == r.Array || k == r.Chan || k == r.Map || k == r.Slice {
		relem = rt.Elem()
	} else if k == r.String {
		relem = rbasictypes[r.Uint8]
	}

	if k == r.Array {
		// methods on arrays have pointer receiver
		rt = r.PtrTo(rt)
	}
	vt := []r.Type{rt}
	vtkey := []r.Type{rt, rkey}
	vint := []r.Type{rint}

	m := xt.methodvalues
	if len(m) < n {
		xt.methodvalues = make([]r.Value, n)
		copy(xt.methodvalues, m)
		m = xt.methodvalues
	}
	if v.debug() {
		v.debugf("addTypeMethodsCTI: %s %v", k, xt.rtype)
		defer de(bug(v))
	}
	for i := 0; i < n; i++ {
		switch xt.method(i).Name {

		// array, slice, string methods
		case "Append":
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, r.SliceOf(relem)}, vt, true), ctiAppend)
		case "AppendString":
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, rbasictypes[r.String]}, vt, false), ctiAppendString)
		case "Copy":
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, rt}, nil, false), ctiCopy)
		case "CopyString":
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, rbasictypes[r.String]}, nil, false), ctiCopy)
		case "Cap":
			m[i] = r.MakeFunc(r.FuncOf(vt, vint, false), ctiCap)
		case "Len":
			m[i] = r.MakeFunc(r.FuncOf(vt, vint, false), ctiLen)
		case "Slice":
			vret := vt
			if k == r.Array {
				vret = []r.Type{r.SliceOf(relem)}
			}
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, rkey, rkey}, vret, false), ctiSlice)
		case "Slice3":
			vret := vt
			if k == r.Array {
				vret = []r.Type{r.SliceOf(relem)}
			}
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, rkey, rkey, rkey}, vret, false), ctiSlice3)

			// indexing
		case "AddrIndex":
			sig := r.FuncOf(vtkey, []r.Type{r.PtrTo(relem)}, false)
			m[i] = r.MakeFunc(sig, ctiAddrIndex)
		case "DelIndex":
			m[i] = r.MakeFunc(r.FuncOf(vtkey, nil, false), ctiDelMapIndex)
		case "Index":
			sig := r.FuncOf(vtkey, []r.Type{relem}, false)
			if k == r.Map {
				zero := r.Zero(relem)
				m[i] = r.MakeFunc(sig,
					func(v []r.Value) []r.Value {
						ret := v[0].MapIndex(v[1])
						if !ret.IsValid() {
							ret = zero
						}
						return []r.Value{ret}
					})
			} else {
				m[i] = r.MakeFunc(sig, ctiIndex)
			}
		case "SetIndex":
			sig := r.FuncOf([]r.Type{rt, rkey, relem}, nil, false)
			if k == r.Map {
				m[i] = r.MakeFunc(sig, ctiSetMapIndex)
			} else {
				m[i] = r.MakeFunc(sig, ctiSetIndex)
			}
		case "TryIndex":
			sig := r.FuncOf(vtkey, []r.Type{relem, rbool}, false)

			zero := r.Zero(relem)
			m[i] = r.MakeFunc(sig,
				func(v []r.Value) []r.Value {
					elem := v[0].MapIndex(v[1])
					ok := vtrue
					if !elem.IsValid() {
						elem = zero
						ok = vfalse
					}
					return []r.Value{elem, ok}
				})

			// chan methods
		case "Recv":
			m[i] = r.MakeFunc(r.FuncOf(vt, []r.Type{relem, rbool}, false), ctiRecv)
		case "Send":
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, relem}, nil, false), ctiSend)
		case "TryRecv":
			m[i] = r.MakeFunc(r.FuncOf(vt, []r.Type{relem, rbool}, false), ctiTryRecv)
		case "TrySend":
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, relem}, []r.Type{rbool}, false), ctiTrySend)
		case "Close":
			m[i] = r.MakeFunc(r.FuncOf(vt, nil, false), ctiClose)
		}
	}
}

// array, slice, string methods

func ctiAppend(v []r.Value) []r.Value {
	return []r.Value{
		r.AppendSlice(v[0], v[1]),
	}
}

var rTypeOfByteSlice = r.SliceOf(rbasictypes[r.Uint8])

func ctiAppendString(v []r.Value) []r.Value {
	vslice := v[0]
	t := vslice.Type()
	if t != rTypeOfByteSlice {
		vslice = vslice.Convert(rTypeOfByteSlice)
	}
	slice := vslice.Interface().([]byte)
	slice = append(slice, v[1].String()...)
	vslice = r.ValueOf(slice)
	if t != rTypeOfByteSlice {
		vslice = vslice.Convert(t)
	}
	return []r.Value{vslice}
}

func ctiCap(v []r.Value) []r.Value {
	return []r.Value{r.ValueOf(
		r.Indirect(v[0]).Cap(),
	)}
}

func ctiCopy(v []r.Value) []r.Value {
	r.Copy(r.Indirect(v[0]), v[1])
	return nil
}

func ctiLen(v []r.Value) []r.Value {
	return []r.Value{r.ValueOf(
		r.Indirect(v[0]).Len(),
	)}
}

func ctiSlice(v []r.Value) []r.Value {
	return []r.Value{
		r.Indirect(v[0]).Slice(int(v[1].Int()), int(v[2].Int())),
	}
}

func ctiSlice3(v []r.Value) []r.Value {
	return []r.Value{
		r.Indirect(v[0]).Slice3(int(v[1].Int()), int(v[2].Int()), int(v[3].Int())),
	}
}

// indexing

func ctiAddrIndex(v []r.Value) []r.Value {
	return []r.Value{
		r.Indirect(v[0]).Index(int(v[1].Int())).Addr(),
	}
}

func ctiDelMapIndex(v []r.Value) []r.Value {
	v[0].SetMapIndex(v[1], r.Value{})
	return nil
}

func ctiIndex(v []r.Value) []r.Value {
	return []r.Value{
		r.Indirect(v[0]).Index(int(v[1].Int())),
	}
}

func ctiSetMapIndex(v []r.Value) []r.Value {
	v[0].SetMapIndex(v[1], v[2])
	return nil
}

func ctiSetIndex(v []r.Value) []r.Value {
	r.Indirect(v[0]).Index(int(v[1].Int())).Set(v[2])
	return nil
}

// chan methods

func ctiRecv(v []r.Value) []r.Value {
	ret, ok := v[0].Recv()
	return []r.Value{
		ret, r.ValueOf(ok),
	}
}

func ctiSend(v []r.Value) []r.Value {
	v[0].Send(v[1])
	return nil
}

func ctiTryRecv(v []r.Value) []r.Value {
	ret, ok := v[0].TryRecv()
	return []r.Value{
		ret, r.ValueOf(ok),
	}
}

func ctiTrySend(v []r.Value) []r.Value {
	return []r.Value{r.ValueOf(
		v[0].TrySend(v[1]),
	)}
}

func ctiClose(v []r.Value) []r.Value {
	v[0].Close()
	return nil
}

// add CTI methods to named type wrapping a basic type
func (v *Universe) addBasicTypeReflectMethodsCTI(xt *xtype) {
	if !etoken.GENERICS_V2_CTI {
		return
	}

	rt := xt.rtype
	vt := []r.Type{rt}
	vtt := []r.Type{rt, rt}
	vttt := []r.Type{rt, rt, rt}
	sig_tt_bool := r.FuncOf(vtt, []r.Type{rbasictypes[r.Bool]}, false)
	sig_unary := r.FuncOf(vtt, vt, false)
	sig_binary := r.FuncOf(vttt, vt, false)

	mvec := xt.GetMethods()
	switch xt.kind {
	case r.Bool:
		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.MakeFunc(
					sig_tt_bool,
					func(v []r.Value) []r.Value {
						flag := v[0].Bool() == v[1].Bool()
						return vvtruefalse[flag]
					},
				)
			case "Not":
				(*mvec)[i] = r.MakeFunc(
					r.FuncOf(vtt, vt, false),
					func(v []r.Value) []r.Value {
						rtyp := v[0].Type()
						ret := r.ValueOf(!v[1].Bool()).Convert(rtyp)
						return []r.Value{ret}
					},
				)
			}
		}
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:

		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.MakeFunc(
					sig_tt_bool,
					func(v []r.Value) []r.Value {
						flag := v[0].Int() == v[1].Int()
						return vvtruefalse[flag]
					},
				)
			case "Cmp":
				(*mvec)[i] = r.MakeFunc(
					r.FuncOf(vtt, []r.Type{rbasictypes[r.Int]}, false),
					func(v []r.Value) []r.Value {
						a, b := v[0].Int(), v[1].Int()
						ret := cmp_eq
						if a < b {
							ret = cmp_less
						} else if a > b {
							ret = cmp_gtr
						}
						return vvcmp[ret]
					},
				)
			case "Less":
				(*mvec)[i] = r.MakeFunc(
					sig_tt_bool,
					func(v []r.Value) []r.Value {
						flag := v[0].Int() < v[1].Int()
						return vvtruefalse[flag]
					},
				)
			case "Add":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Int() + v[2].Int()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Sub":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Int() - v[2].Int()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Mul":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Int() * v[2].Int()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Quo":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Int() / v[2].Int()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Neg":
				(*mvec)[i] = r.MakeFunc(
					sig_unary,
					func(v []r.Value) []r.Value {
						ret := -v[1].Int()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Rem":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Int() % v[2].Int()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "And":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Int() & v[2].Int()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "AndNot":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Int() &^ v[2].Int()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Or":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Int() | v[2].Int()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Xor":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Int() ^ v[2].Int()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Not":
				(*mvec)[i] = r.MakeFunc(
					sig_unary,
					func(v []r.Value) []r.Value {
						ret := ^v[1].Int()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Lsh":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Int() << uint8(v[2].Uint())
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Rsh":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Int() >> uint8(v[2].Uint())
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			}
		}
	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:

		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.MakeFunc(
					sig_tt_bool,
					func(v []r.Value) []r.Value {
						flag := v[0].Uint() == v[1].Uint()
						return vvtruefalse[flag]
					},
				)
			case "Cmp":
				(*mvec)[i] = r.MakeFunc(
					r.FuncOf(vtt, []r.Type{rbasictypes[r.Int]}, false),
					func(v []r.Value) []r.Value {
						a, b := v[0].Uint(), v[1].Uint()
						ret := cmp_eq
						if a < b {
							ret = cmp_less
						} else if a > b {
							ret = cmp_gtr
						}
						return vvcmp[ret]
					},
				)
			case "Less":
				(*mvec)[i] = r.MakeFunc(
					sig_tt_bool,
					func(v []r.Value) []r.Value {
						flag := v[0].Uint() < v[1].Uint()
						return vvtruefalse[flag]
					},
				)
			case "Add":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Uint() + v[2].Uint()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Sub":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Uint() - v[2].Uint()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Mul":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Uint() * v[2].Uint()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Quo":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Uint() / v[2].Uint()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Neg":
				(*mvec)[i] = r.MakeFunc(
					sig_unary,
					func(v []r.Value) []r.Value {
						ret := -v[1].Uint()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Rem":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Uint() % v[2].Uint()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "And":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Uint() & v[2].Uint()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "AndNot":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Uint() &^ v[2].Uint()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Or":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Uint() | v[2].Uint()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Xor":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Uint() ^ v[2].Uint()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Not":
				(*mvec)[i] = r.MakeFunc(
					sig_unary,
					func(v []r.Value) []r.Value {
						ret := ^v[1].Uint()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Lsh":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Uint() << uint8(v[2].Uint())
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Rsh":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Uint() >> uint8(v[2].Uint())
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			}
		}
	case r.Float32, r.Float64:

		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.MakeFunc(
					sig_tt_bool,
					func(v []r.Value) []r.Value {
						flag := v[0].Float() == v[1].Float()
						return vvtruefalse[flag]
					},
				)
			case "Cmp":
				(*mvec)[i] = r.MakeFunc(
					r.FuncOf(vtt, []r.Type{rbasictypes[r.Int]}, false),
					func(v []r.Value) []r.Value {
						a, b := v[0].Float(), v[1].Float()
						ret := cmp_eq
						if a < b {
							ret = cmp_less
						} else if a > b {
							ret = cmp_gtr
						}
						return vvcmp[ret]
					},
				)
			case "Less":
				(*mvec)[i] = r.MakeFunc(
					sig_tt_bool,
					func(v []r.Value) []r.Value {
						flag := v[0].Float() < v[1].Float()
						return vvtruefalse[flag]
					},
				)
			case "Add":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Float() + v[2].Float()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Sub":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Float() - v[2].Float()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Mul":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Float() * v[2].Float()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Quo":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Float() / v[2].Float()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Neg":
				(*mvec)[i] = r.MakeFunc(
					sig_unary,
					func(v []r.Value) []r.Value {
						ret := -v[1].Float()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			}
		}
	case r.Complex64, r.Complex128:

		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.MakeFunc(
					sig_tt_bool,
					func(v []r.Value) []r.Value {
						flag := v[0].Complex() == v[1].Complex()
						return vvtruefalse[flag]
					},
				)
			case "Add":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Complex() + v[2].Complex()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Sub":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Complex() - v[2].Complex()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Mul":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Complex() * v[2].Complex()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Quo":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].Complex() / v[2].Complex()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Neg":
				(*mvec)[i] = r.MakeFunc(
					sig_unary,
					func(v []r.Value) []r.Value {
						ret := -v[1].Complex()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Real":
				if xt.kind == r.Complex64 {
					(*mvec)[i] = r.MakeFunc(
						r.FuncOf(vt, []r.Type{rbasictypes[r.Float32]}, false),
						func(v []r.Value) []r.Value {
							ret := float32(real(v[0].Complex()))
							return []r.Value{r.ValueOf(ret)}
						},
					)
				} else {
					(*mvec)[i] = r.MakeFunc(
						r.FuncOf(vt, []r.Type{rbasictypes[r.Float64]}, false),
						func(v []r.Value) []r.Value {
							ret := real(v[0].Complex())
							return []r.Value{r.ValueOf(ret)}
						},
					)
				}
			case "Imag":
				if xt.kind == r.Complex64 {
					(*mvec)[i] = r.MakeFunc(
						r.FuncOf(vt, []r.Type{rbasictypes[r.Float32]}, false),
						func(v []r.Value) []r.Value {
							ret := float32(imag(v[0].Complex()))
							return []r.Value{r.ValueOf(ret)}
						},
					)
				} else {
					(*mvec)[i] = r.MakeFunc(
						r.FuncOf(vt, []r.Type{rbasictypes[r.Float64]}, false),
						func(v []r.Value) []r.Value {
							ret := imag(v[0].Complex())
							return []r.Value{r.ValueOf(ret)}
						},
					)
				}
			}
		}
	case r.String:

		for i, n := 0, xt.NumMethod(); i < n; i++ {
			switch xt.Method(i).Name {
			case "Equal":
				(*mvec)[i] = r.MakeFunc(
					sig_tt_bool,
					func(v []r.Value) []r.Value {
						flag := v[0].String() == v[1].String()
						return vvtruefalse[flag]
					},
				)
			case "Cmp":
				(*mvec)[i] = r.MakeFunc(
					r.FuncOf(vtt, []r.Type{rbasictypes[r.Int]}, false),
					func(v []r.Value) []r.Value {
						a, b := v[0].String(), v[1].String()
						ret := cmp_eq
						if a < b {
							ret = cmp_less
						} else if a > b {
							ret = cmp_gtr
						}
						return vvcmp[ret]
					},
				)
			case "Less":
				(*mvec)[i] = r.MakeFunc(
					sig_tt_bool,
					func(v []r.Value) []r.Value {
						flag := v[0].String() < v[1].String()
						return vvtruefalse[flag]
					},
				)
			case "Add":
				(*mvec)[i] = r.MakeFunc(
					sig_binary,
					func(v []r.Value) []r.Value {
						ret := v[1].String() + v[2].String()
						return []r.Value{r.ValueOf(ret).Convert(v[0].Type())}
					},
				)
			case "Index":
				(*mvec)[i] = r.MakeFunc(
					r.FuncOf([]r.Type{rt, rbasictypes[r.Int]}, []r.Type{rbasictypes[r.Uint8]}, false),
					func(v []r.Value) []r.Value {
						ret := v[0].String()[v[1].Int()]
						return []r.Value{r.ValueOf(ret)}
					},
				)
			case "Len":
				(*mvec)[i] = r.MakeFunc(
					r.FuncOf(vt, []r.Type{rt, rbasictypes[r.Int]}, false),
					func(v []r.Value) []r.Value {
						ret := len(v[0].String())
						return []r.Value{r.ValueOf(ret)}
					},
				)
			case "Slice":
				(*mvec)[i] = r.MakeFunc(
					r.FuncOf([]r.Type{rt, rbasictypes[r.Int], rbasictypes[r.Int]}, []r.Type{r.SliceOf(rbasictypes[r.Uint8])}, false),
					func(v []r.Value) []r.Value {
						ret := v[0].String()[v[1].Int():v[2].Int()]
						return []r.Value{r.ValueOf(ret)}
					},
				)
			}
		}
	}
}
