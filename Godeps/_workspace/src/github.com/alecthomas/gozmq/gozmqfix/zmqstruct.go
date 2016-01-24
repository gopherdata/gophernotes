// Copyright 2013 Joshua Tacoma.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"go/ast"
)

func init() {
	register(zmqstructFix)
}

var zmqstructFix = fix{
	"zmqstruct",
	"2013-03-20",
	zmqstruct,
	`
	Make github.com/alecthomas/gozmq use structs instead of interfaces.
`,
}

func zmqstruct(f *ast.File) bool {
	spec := importSpec(f, "github.com/alecthomas/gozmq")
	if spec == nil {
		return false
	}
	zmq := "gozmq"
	if spec.Name != nil {
		zmq = spec.Name.Name
	}

	fixed := false
	walk(f, func(n interface{}) {
		switch node := n.(type) {
		case *ast.ArrayType:
			t := zmqstructtype(zmq, node.Elt)
			if t != nil {
				node.Elt = t
				fixed = true
			}
		case *ast.CompositeLit:
			// This is irrelevant only because the original type is an
			// interface i.e. cannot be the type of a composite literal.
		case *ast.Ellipsis:
			t := zmqstructtype(zmq, node.Elt)
			if t != nil {
				node.Elt = t
				fixed = true
			}
		case *ast.Field:
			t := zmqstructtype(zmq, node.Type)
			if t != nil {
				node.Type = t
				fixed = true
			}
		case *ast.MapType:
			t := zmqstructtype(zmq, node.Key)
			if t != nil {
				node.Key = t
				fixed = true
			}
			t = zmqstructtype(zmq, node.Value)
			if t != nil {
				node.Value = t
				fixed = true
			}
		case *ast.Object:
			// Does something need to be done here with node.Type?
			// What does it take to trigger this case?
		case *ast.TypeAssertExpr:
			t := zmqstructtype(zmq, node.Type)
			if t != nil {
				node.Type = t
				fixed = true
			}
		case *ast.TypeSpec:
			t := zmqstructtype(zmq, node.Type)
			if t != nil {
				node.Type = t
				fixed = true
			}
		case *ast.ValueSpec:
			t := zmqstructtype(zmq, node.Type)
			if t != nil {
				node.Type = t
				fixed = true
			}
		}
	})
	return fixed
}

func zmqstructtype(zmq string, n ast.Expr) ast.Expr {
	s, ok := n.(*ast.SelectorExpr)
	if ok {
		p, ok := s.X.(*ast.Ident)
		if ok && p.Name == zmq {
			if s.Sel.Name == "Context" || s.Sel.Name == "Socket" {
				return &ast.StarExpr{
					X: &ast.SelectorExpr{
						X:   ast.NewIdent(zmq),
						Sel: ast.NewIdent(s.Sel.Name),
					},
				}
			}
		}
	}
	return nil
}
