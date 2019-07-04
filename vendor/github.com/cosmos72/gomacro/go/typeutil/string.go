// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements commonly used type predicates.

package typeutil

import (
	"bytes"
	"fmt"

	"github.com/cosmos72/gomacro/go/types"
)

// String returns a string representation of given type,
// including its receiver if present.
func String(x types.Type) string {
	return String2("", x)
}

func String2(name string, x types.Type) string {
	switch x.(type) {
	case nil:
		return "nil"
	case *types.Signature:
		break
	default:
		if name != "" {
			return name + " " + x.String()
		}
		return x.String()
	}

	var buf bytes.Buffer
	buf.WriteString("func")
	sig := x.(*types.Signature)
	writeRecv(&buf, sig.Recv())
	buf.WriteString(name)
	writeVars(&buf, sig.Params(), sig.Variadic())
	writeResults(&buf, sig.Results())
	return buf.String()
}

func writeRecv(buf *bytes.Buffer, recv *types.Var) {
	if recv != nil {
		buf.WriteString(" (")
		writeVar(buf, recv, false)
		buf.WriteString(").")
	}
}

func writeVars(buf *bytes.Buffer, vars *types.Tuple, variadic bool) {
	if vars == nil {
		buf.WriteString("()")
		return
	}
	buf.WriteByte('(')
	for i, n := 0, vars.Len(); i < n; i++ {
		if i != 0 {
			buf.WriteString(", ")
		}
		writeVar(buf, vars.At(i), variadic && i == n-1)
	}
	buf.WriteByte(')')
}

func writeResults(buf *bytes.Buffer, vars *types.Tuple) {
	if vars == nil {
		return
	}
	switch vars.Len() {
	case 0:
		break
	case 1:
		buf.WriteByte(' ')
		writeVar(buf, vars.At(0), false)
	default:
		buf.WriteByte(' ')
		writeVars(buf, vars, false)
	}
}

func writeVar(buf *bytes.Buffer, v *types.Var, variadic bool) {
	if v == nil {
		return
	}
	if v.Name() != "" {
		buf.WriteString(v.Name())
		buf.WriteByte(' ')
	}
	t := v.Type()
	if variadic {
		buf.WriteString("...")
		t = elemType(t)
	}
	buf.WriteString(t.String())
}

func elemType(t types.Type) types.Type {
	switch t := t.(type) {
	case *types.Array:
		return t.Elem()
	case *types.Slice:
		return t.Elem()
	case *types.Basic:
		if t.Kind() == types.String {
			return types.Typ[types.Uint8]
		}
	}
	panic(fmt.Errorf("type %v cannot be the last parameter of a variadic function", t))
}
