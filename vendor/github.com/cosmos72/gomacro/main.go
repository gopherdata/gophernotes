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
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * main.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"fmt"
	"os"

	"github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/parser"
)

func main() {
	args := os.Args[1:]

	var cmd Cmd
	cmd.Init()

	cmd.Options |= base.OptFastInterpreter // use fast interpreter by default

	// cmd.Options |= base.OptShowTime // | base.OptTrapPanic // | base.OptShowAfterMacroExpansion // | base.OptShowAfterParse // | base.OptDebugMacroExpand // |  base.OptDebugQuasiquote

	cmd.ParserMode |= parser.Trace & 0

	err := cmd.Main(args)
	if err != nil {
		fmt.Fprintln(cmd.Stderr, err)
	}
}

/*
	miscellaneous annotations

	imports, generated with: find [a-u]* -type f -name \*.go | grep -v internal | grep -v testdata | grep -v cmd/ | grep -v builtin | xargs -d'\n' dirname | sort -u | while read i; do echo -n "_b \"$i\"; "; done
	plus some hand-made tweaks
	import ( _b "archive/tar"; _b "archive/zip"; _b "bufio"; _b "bytes"; _b "compress/bzip2"; _b "compress/flate"; _b "compress/gzip"; _b "compress/lzw"; _b "compress/zlib"; _b "container/heap"; _b "container/list"; _b "container/ring"; _b "context"; _b "crypto"; _b "crypto/aes"; _b "crypto/cipher"; _b "crypto/des"; _b "crypto/dsa"; _b "crypto/ecdsa"; _b "crypto/elliptic"; _b "crypto/hmac"; _b "crypto/md5"; _b "crypto/rand"; _b "crypto/rc4"; _b "crypto/rsa"; _b "crypto/sha1"; _b "crypto/sha256"; _b "crypto/sha512"; _b "crypto/subtle"; _b "crypto/tls"; _b "crypto/x509"; _b "crypto/x509/pkix"; _b "database/sql"; _b "database/sql/driver"; _b "debug/dwarf"; _b "debug/elf"; _b "debug/gosym"; _b "debug/macho"; _b "debug/pe"; _b "debug/plan9obj"; _b "encoding"; _b "encoding/ascii85"; _b "encoding/asn1"; _b "encoding/base32"; _b "encoding/base64"; _b "encoding/binary"; _b "encoding/csv"; _b "encoding/gob"; _b "encoding/hex"; _b "encoding/json"; _b "encoding/pem"; _b "encoding/xml"; _b "errors"; _b "expvar"; _b "flag"; _b "fmt"; _b "go/ast"; _b "go/build"; _b "go/constant"; _b "go/doc"; _b "go/format"; _b "go/importer"; _b "go/parser"; _b "go/printer"; _b "go/scanner"; _b "go/token"; _b "go/types"; _b "hash"; _b "hash/adler32"; _b "hash/crc32"; _b "hash/crc64"; _b "hash/fnv"; _b "html"; _b "html/template"; _b "image"; _b "image/color"; _b "image/color/palette"; _b "image/draw"; _b "image/gif"; _b "image/jpeg"; _b "image/png"; _b "index/suffixarray"; _b "io"; _b "io/ioutil"; _b "log"; _b "log/syslog"; _b "math"; _b "math/big"; _b "math/cmplx"; _b "math/rand"; _b "mime"; _b "mime/multipart"; _b "mime/quotedprintable"; _b "net"; _b "net/http"; _b "net/http/cgi"; _b "net/http/cookiejar"; _b "net/http/fcgi"; _b "net/http/httptest"; _b "net/http/httptrace"; _b "net/http/httputil"; _b "net/http/pprof"; _b "net/mail"; _b "net/rpc"; _b "net/rpc/jsonrpc"; _b "net/smtp"; _b "net/textproto"; _b "net/url"; _b "os"; _b "os/exec"; _b "os/signal"; _b "os/user"; _b "path"; _b "path/filepath"; _b "plugin"; _b "reflect"; _b "regexp"; _b "regexp/syntax"; _b "runtime"; _b "runtime/debug"; _b "runtime/pprof"; _b "runtime/trace"; _b "sort"; _b "strconv"; _b "strings"; _b "sync"; _b "sync/atomic"; _b "syscall"; _b "testing"; _b "testing/iotest"; _b "testing/quick"; _b "text/scanner"; _b "text/tabwriter"; _b "text/template"; _b "text/template/parse"; _b "time"; _b "unicode"; _b "unicode/utf16"; _b "unicode/utf8"; _b "unsafe" )

    // test interfaces:

	import ( "time"; "fmt" ); var s fmt.Stringer = time.Second  // ok
	s.String                                                    // ok, s is the interface fmt.Stringer

	import ( "os"; "io" );    var in io.Reader = os.Stdin       // ok
	import "reflect";         var t = reflect.TypeOf(os.Stdin)  // ok
	t.Elem                                                      // ok, t is the interface reflect.Type

	// test methods-to-functions:

	time.Duration.Seconds       // easy, time.Duration is a concrete type
	io.Stringer.String          // harder, io.Stringer is an interface

	// test methods:

	type Pair struct { A, B int }; var p Pair
	func (p *Pair) Lhs() int { return p.A }; func (p *Pair) SetLhs(a int) { p.A = a }
	p.SetLhs(2); p.Lhs()

	type Triple struct { Pair; C int }; var t Triple
	t.Pair = p
	t.SetLhs(3); t.Lhs()

	// test some valid methods:

	type Pair2 Pair; var p2 Pair2
	func (p Pair2) Foo() { }

	type SPair []Pair; var s SPair
	func (p SPair) Foo() { }; s.Foo()

	type Int int; var I Int
	func (x Int) Print() { println(x) }; I.Print; I.Print()
	func (x *Int) Print() { println(*x) }; I.Print; I.Print()

	// test some bogus methods:

	func (p **Pair) Foo() { }
	type PPair *Pair
	func (p PPair) Foo() { }
	func (p *PPair) Foo() { }

	// test type alias:

	type iint = int
	type Int int
	var i int = 1
	var ii iint = 2
	var I Int = 3
	func show(x int) { println(x) }
	func Show(x Int) { println(x) }

*/
