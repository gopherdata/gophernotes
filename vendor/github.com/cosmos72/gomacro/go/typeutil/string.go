// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements commonly used type predicates.

package typeutil

import (
	"strings"

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

	buf := []string{"func"}
	sig := x.(*types.Signature)
	if recv := sig.Recv(); recv != nil {
		buf = append(buf, " (", varString(recv), ").")
	}
	buf = append(buf, name)
	if vars := sig.Params(); vars != nil {
		buf = append(buf, vars.String())
	} else {
		buf = append(buf, "()")
	}
	if vars := sig.Results(); vars != nil {
		switch vars.Len() {
		case 0:
			break
		case 1:
			buf = append(buf, " ", varString(vars.At(0)))
		default:
			buf = append(buf, " ", vars.String())
		}
	}
	return strings.Join(buf, "")
}

func varString(v *types.Var) string {
	if v == nil {
		return ""
	} else if v.Name() == "" {
		return v.Type().String()
	} else {
		return v.Name() + " " + v.Type().String()
	}
}
