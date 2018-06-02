// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the exported entry points for invoking the parser.

package parser

import (
	"fmt"
	"go/ast"
	"go/token"

	mt "github.com/cosmos72/gomacro/token"
)

// A Mode value is a set of flags (or 0).
// They control the amount of source code parsed and other optional
// parser functionality.
//
type Mode uint

const (
	PackageClauseOnly Mode             = 1 << iota // stop parsing after package clause
	ImportsOnly                                    // stop parsing after import declarations
	ParseComments                                  // parse comments and add them to AST
	Trace                                          // print a trace of parsed productions
	DeclarationErrors                              // report declaration errors
	SpuriousErrors                                 // same as AllErrors, for backward-compatibility
	CopySources                                    // copy source code to FileSet
	AllErrors         = SpuriousErrors             // report all errors (not just the first 10 on different lines)

)

type Parser struct {
	parser
}

func (p *parser) Configure(mode Mode, macroChar rune) {
	p.mode = mode
	p.macroChar = macroChar
}

func (p *parser) Init(fileset *mt.FileSet, filename string, lineOffset int, src []byte) {
	p.init(fileset, filename, lineOffset, src, p.mode)
}

func (p *parser) Parse() (list []ast.Node, err error) {
	if p.file == nil || p.pkgScope == nil {
		panic("Parser.Parse(): parser is not initialized, call Parser.Init() first")
	}

	defer func() {
		if e := recover(); e != nil {
			// resume same panic if it's not a bailout
			if _, ok := e.(bailout); !ok {
				panic(e)
			}
		}
		p.errors.Sort()
		err = p.errors.Err()
		p.file = nil
		p.pkgScope = nil
	}()

	topScope := p.topScope

	var lastpos1, lastpos2 token.Pos
	list = make([]ast.Node, 0)
	for p.tok != token.EOF && p.errors.Len() < 10 {
		list = append(list, p.parseAny())
		// fmt.Printf("// parser position is now %d (%s). parsed %#v\n", p.pos, p.file.Position(p.pos), list[len(list)-1])
		if p.pos == lastpos1 {
			p.error(p.pos, fmt.Sprintf("skipping '%s' to continue", mt.String(p.tok)))
			p.next()
		} else {
			lastpos1 = lastpos2
			lastpos2 = p.pos
		}
	}

	assert(topScope == p.topScope, "unbalanced scopes")

	if p.errors.Len() > 0 {
		p.errors.Sort()
		return list, p.errors.Err()
	}
	return list, nil
}

func (p *parser) parseAny() ast.Node {
	if p.tok == token.COMMENT {
		// advance to the next non-comment token
		p.next()
	}
	var node ast.Node
	switch p.tok {
	case token.PACKAGE:
		// not p.parseFile() because it does not support top-level statements and expressions
		node = p.parsePackage()
	case token.IMPORT:
		node = p.parseGenDecl(token.IMPORT, p.parseImportSpec)
	case token.CONST, token.TYPE, token.VAR, token.FUNC, mt.MACRO, mt.FUNCTION:
		// a "func" at top level can be either a function declaration: func foo(args) /*...*/
		// or a method declaration: func (receiver) foo(args) /*...*/
		// or a function literal, i.e. a closure: func(args) /*...*/
		// since method declaration and function literal are so similar,
		// there is no reasonable way to distinguish them here.
		//
		// decision: always parse as a declaration.
		// function literals at top level must either be written ~lambda(args) /*...*/
		// or come after some other token: a variable declaration, an expression,
		// or at least a '('
		node = p.parseDecl(syncDecl)
	default:
		node = p.parseStmt()
		if expr, ok := node.(*ast.ExprStmt); ok {
			// unwrap expressions
			node = expr.X
		}
	}
	return node
}

func (p *parser) parsePackage() ast.Node {
	if p.trace {
		defer un(trace(p, "Package"))
	}
	doc := p.leadComment
	pos := p.expect(token.PACKAGE)
	var path string

	switch p.tok {
	case token.IDENT:
		ident := p.parseIdent()
		path = ident.Name
	case token.STRING:
		path = p.lit
		p.next()
	default:
		p.expect(token.IDENT)
	}
	if path == "_" && p.mode&DeclarationErrors != 0 {
		p.error(p.pos, "invalid package name: _")
	}
	npos := p.pos
	p.expectSemi()

	return &ast.GenDecl{
		TokPos: pos,
		Tok:    token.PACKAGE,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Doc: doc,
				Values: []ast.Expr{
					&ast.BasicLit{
						ValuePos: npos,
						Kind:     token.STRING,
						Value:    path,
					},
				},
			},
		},
	}
}
