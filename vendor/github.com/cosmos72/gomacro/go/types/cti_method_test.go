// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"go/token"
	"testing"

	"github.com/cosmos72/gomacro/go/etoken"
)

func mktuple(ts ...Type) *Tuple {
	vs := make([]*Var, len(ts))
	for i := range ts {
		vs[i] = newVar(ts[i])
	}
	return NewTuple(vs...)
}

func mkfunc(name string, params *Tuple, results *Tuple) *Func {
	return NewFunc(token.NoPos, nil, name, NewSignature(nil, params, results, false))
}

func mkinterface(fs ...*Func) *Interface {
	return NewInterface(fs, nil).Complete()
}

/**
 * return
 * interface {
 *   Cap() int
 *   Len() int
 * }
 */
func mkInterfaceCapLen() *Interface {
	return mkinterface(
		mkfunc("Cap", nil, mktuple(Typ[Int])),
		mkfunc("Len", nil, mktuple(Typ[Int])),
	)
}

/**
 * return
 * interface {
 *   AddrIndex(k Key) *Value
 * }
 */
func mkInterfaceAddrIndex(key, value Type) *Interface {
	return mkinterface(
		mkfunc("AddrIndex", mktuple(key), mktuple(NewPointer(value))),
	)
}

/**
 * return
 * interface {
 *   Index(k Key) Value
 *   Len() int
 * }
 */
func mkInterfaceIndexLen(key, value Type) *Interface {
	return mkinterface(
		mkfunc("Index", mktuple(key), mktuple(value)),
		mkfunc("Len", nil, mktuple(Typ[Int])),
	)
}

/**
 * return
 * interface {
 *   Send(e Elem)
 *   Recv() (Elem, bool)
 * }
 */
func mkInterfaceSendRecv(elem Type) *Interface {
	return mkinterface(
		mkfunc("Send", mktuple(elem), nil),
		mkfunc("Recv", mktuple(elem, Typ[Bool]), nil),
	)
}

/**
 * return
 * interface {
 *   SetIndex(k Key, v Value)
 * }
 */
func mkInterfaceSetIndex(key, value Type) *Interface {
	return mkinterface(
		mkfunc("SetIndex", mktuple(key, value), nil),
	)
}

func mkNamed(name string, underlying Type) *Named {
	return NewNamed(NewTypeName(token.NoPos, nil, name, nil), underlying, nil)
}

type tcase struct {
	typ        Type
	interfaces []*Interface
}

func mkcase(typ Type, interfaces ...*Interface) tcase {
	return tcase{typ, interfaces}
}

func TestCTIMethods(t *testing.T) {
	if !etoken.GENERICS_V2_CTI {
		t.SkipNow()
		return
	}
	checkImplements := func(typ Type, v *Interface) {
		m, _ := MissingMethod(typ, v, true)
		if m != nil {
			t.Errorf("type %v does not implement %v: missing method %v", typ, v, m)
		}
	}
	checkNotImplements := func(typ Type, v *Interface) {
		m, _ := MissingMethod(typ, v, true)
		if m == nil {
			t.Errorf("type %v implements %v: this should not happen", typ, v)
		}
	}
	caplen := mkInterfaceCapLen()
	addrindex := mkInterfaceAddrIndex(Typ[Int], Typ[Uint8])
	getlen := mkInterfaceIndexLen(Typ[Int], Typ[Uint8])
	set := mkInterfaceSetIndex(Typ[Int], Typ[Uint8])
	sendrecv := mkInterfaceSendRecv(Typ[Int])
	allifaces := []*Interface{
		caplen, addrindex, getlen, set,
	}
	contains := func(slice []*Interface, key *Interface) bool {
		for _, elem := range slice {
			if elem == key {
				return true
			}
		}
		return false
	}

	tarray := NewArray(Typ[Uint8], 0)
	tchan := NewChan(SendRecv, Typ[Int])
	tmap := NewMap(Typ[Int], Typ[Uint8])
	tslice := NewSlice(Typ[Uint8])
	tstring := Typ[String]

	tchannamed := mkNamed("ChanInt", tchan)
	tmapnamed := mkNamed("MapIntUint8", tmap)
	tslicenamed := mkNamed("SliceUint8", tslice)
	tstringnamed := mkNamed("String", tstring)

	tcases := []tcase{
		mkcase(NewPointer(tarray), caplen, addrindex, getlen, set),
		mkcase(tchan, caplen, sendrecv),
		mkcase(tchannamed, caplen, sendrecv),
		mkcase(tmap, getlen, set),
		mkcase(tmapnamed, getlen, set),
		mkcase(tslice, caplen, addrindex, getlen, set),
		mkcase(tslicenamed, caplen, addrindex, getlen, set),
		mkcase(tstring, getlen),
		mkcase(tstringnamed, getlen),
	}

	for _, c := range tcases {
		t := c.typ
		for _, iface := range allifaces {
			if contains(c.interfaces, iface) {
				checkImplements(t, iface)
			} else {
				checkNotImplements(t, iface)
			}
		}
	}
}
