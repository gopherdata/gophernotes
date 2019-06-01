// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package parser implements a parser for Go source files. Input may be
// provided in a variety of forms (see the various Parse* functions); the
// output is an abstract syntax tree (AST) representing the Go source. The
// parser is invoked through one of the Parse* functions.
//
// The parser accepts a larger language than is syntactically permitted by
// the Go spec, for simplicity, and for improved robustness in the presence
// of syntax errors. For instance, in method declarations, the receiver is
// treated like an ordinary parameter list and thus may contain multiple
// entries where the spec permits exactly one. Consequently, the corresponding
// field in the AST (ast.FuncDecl.Recv) field is not restricted to one entry.
//
package parser

import (
	"go/ast"
	"go/token"

	etoken "github.com/cosmos72/gomacro/go/etoken"
)

// enable C++-style generics?
const GENERICS_V1_CXX = etoken.GENERICS_V1_CXX

// enable generics "contracts are interfaces" ?
const GENERICS_V2_CTI = etoken.GENERICS_V2_CTI

// do generics use Foo#[T1,T2...] syntax?
const _GENERICS_HASH = GENERICS_V1_CXX || GENERICS_V2_CTI

/*
 * used by GENERICS_V1_CXX and GENERICS_V2_CTI:
 *    parse prefix#[T1,T2...]
 *    as &ast.IndexExpr{X: prefix, Index: &ast.CompositeLit{Type: nil, Elts: [T1,T2...] } }
 * used by GENERICS_V2_CTI:
 *    parse prefix#[T1:C1,T2:C2...]
 *    as &ast.IndexExpr{X: prefix, Index: &ast.CompositeLit{Type: nil, Elts: [&KeyValueExpr{T1,C1}, &KeyValueExpr{T2,C2} ...] } }
 */
func (p *parser) parseHash(prefix ast.Expr) ast.Expr {
	if p.trace {
		defer un(trace(p, "Hash"))
	}
	p.expect(etoken.HASH)
	params := p.parseGenericParams()
	return &ast.IndexExpr{
		X:      prefix,
		Lbrack: params.Lbrace,
		Index:  params,
		Rbrack: params.Rbrace,
	}
}

// parse template[T1,T2...] type ...
// and template[T1,T2...] func ...
func (p *parser) parseTemplateDecl(sync func(*parser)) ast.Decl {
	if p.trace {
		defer un(trace(p, "TemplateDecl"))
	}
	p.expect(etoken.TEMPLATE)
	params := p.parseGenericParams()

	var specialize *ast.CompositeLit
	if p.tok == token.FOR {
		p.next()
		specialize = p.parseGenericParams()
		params.Elts = append(params.Elts, &ast.BadExpr{}, specialize)
	}
	switch tok := p.tok; tok {
	case token.TYPE:
		decl := p.parseGenDecl(tok, p.parseTypeSpec)
		return genericV1TypeDecl(params, decl)

	case token.FUNC, etoken.FUNCTION:
		decl := p.parseFuncDecl(tok)
		return genericFuncDecl(params, decl)

	default:
		pos := p.pos
		if specialize == nil {
			p.errorExpected(pos, "'type', 'func' or 'for' after 'template[...]'")
		} else {
			p.errorExpected(pos, "'type' or 'func' after 'template[...] for[...]'")
		}
		sync(p)
		return &ast.BadDecl{From: pos, To: p.pos}
	}
}

// parse [T1,T2...] in a generic declaration
func (p *parser) parseGenericParams() *ast.CompositeLit {
	if p.trace {
		defer un(trace(p, "GenericParams"))
	}
	var list []ast.Expr

	lbrack := p.expect(token.LBRACK)
	if p.tok != token.RBRACK {
		if GENERICS_V1_CXX {
			list = append(list, p.parseRhsOrType())
			for p.tok == token.COMMA {
				p.next()
				list = append(list, p.parseRhsOrType())
			}
		} else if GENERICS_V2_CTI {
			for {
				x := p.parseRhsOrType()
				if p.tok == token.COLON {
					colon := p.pos
					p.next()
					x = &ast.KeyValueExpr{Key: x, Colon: colon, Value: p.parseRhsOrType()}
				}
				list = append(list, x)
				if p.tok != token.COMMA {
					break
				}
				p.next()
			}
		}
	}
	rbrack := p.expect(token.RBRACK)

	return &ast.CompositeLit{
		Lbrace: lbrack,
		Elts:   list,
		Rbrace: rbrack,
	}
}

func genericV1TypeDecl(params *ast.CompositeLit, decl *ast.GenDecl) *ast.GenDecl {
	for _, spec := range decl.Specs {
		if typespec, ok := spec.(*ast.TypeSpec); ok {
			// hack: store template params in *ast.CompositeLit.
			// it is never used inside *ast.TypeSpec and has exacly the required fields
			typespec.Type = &ast.CompositeLit{
				Type:   typespec.Type,
				Lbrace: params.Lbrace,
				Elts:   params.Elts,
				Rbrace: params.Rbrace,
			}
		}
	}
	return decl
}

func genericFuncDecl(params *ast.CompositeLit, decl *ast.FuncDecl) *ast.FuncDecl {
	// hack: store generic types as second function receiver.
	// it's never used for functions and macros.
	recv := decl.Recv
	if recv == nil {
		recv = &ast.FieldList{Opening: params.Lbrace, Closing: params.Rbrace}
		decl.Recv = recv
	}
	list := []*ast.Field{
		nil,
		// add generic types as second receiver
		&ast.Field{Type: params},
	}
	if len(recv.List) != 0 {
		list[0] = recv.List[0]
	}
	recv.List = list
	return decl
}
