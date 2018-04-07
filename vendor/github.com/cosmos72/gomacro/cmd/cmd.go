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
 * cmd.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package cmd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/classic"
)

type Cmd struct {
	Interp                                   *classic.Interp
	WriteDeclsAndStmtsToFile, OverwriteFiles bool
}

func New() *Cmd {
	cmd := Cmd{}
	cmd.Init()
	return &cmd
}

func (cmd *Cmd) Init() {
	ir := classic.New()
	ir.ParserMode = 0
	ir.Options = OptFastInterpreter | OptTrapPanic | OptShowPrompt | OptShowEval | OptShowEvalType // | OptShowAfterMacroExpansion // | OptDebugMacroExpand // |  OptDebugQuasiquote  // | OptShowEvalDuration // | OptShowAfterParse
	cmd.Interp = ir
	cmd.WriteDeclsAndStmtsToFile = false
	cmd.OverwriteFiles = false
}

func (cmd *Cmd) Main(args []string) (err error) {
	if cmd.Interp == nil {
		cmd.Init()
	}
	ir := cmd.Interp
	g := ir.Globals

	var set, clear Options
	var repl, forcerepl = true, false
	cmd.WriteDeclsAndStmtsToFile = false
	cmd.OverwriteFiles = false

	for len(args) > 0 {
		switch args[0] {
		case "-c", "--collect":
			g.Options |= OptCollectDeclarations | OptCollectStatements
		case "-e", "--expr":
			if len(args) > 1 {
				repl = false
				buf := bytes.NewBufferString(args[1])
				buf.WriteByte('\n')      // because ReadMultiLine() needs a final '\n'
				g.Options |= OptShowEval // set by default, overridden by -s, -v and -vv
				g.Options = (g.Options | set) &^ clear
				_, err := cmd.EvalReader(buf)
				if err != nil {
					return err
				}
				args = args[1:]
			}
		case "-f", "--force-overwrite":
			cmd.OverwriteFiles = true
		case "-h", "--help":
			return cmd.Usage()
		case "-i", "--repl":
			forcerepl = true
		case "-m", "--macro-only":
			set |= OptMacroExpandOnly
			clear &^= OptMacroExpandOnly
		case "-n", "--no-trap":
			set &^= OptTrapPanic | OptPanicStackTrace
			clear |= OptTrapPanic | OptPanicStackTrace
		case "-t", "--trap":
			set |= OptTrapPanic | OptPanicStackTrace
			clear &= OptTrapPanic | OptPanicStackTrace
		case "-s", "--silent":
			set &^= OptShowEval | OptShowEvalType
			clear |= OptShowEval | OptShowEvalType
		case "-v", "--verbose":
			set = (set | OptShowEval) &^ OptShowEvalType
			clear = (clear &^ OptShowEval) | OptShowEvalType
		case "-vv", "--very-verbose":
			set |= OptShowEval | OptShowEvalType
			clear &^= OptShowEval | OptShowEvalType
		case "-w", "--write-decls":
			cmd.WriteDeclsAndStmtsToFile = true
		case "-x", "--exec":
			clear |= OptMacroExpandOnly
			set &^= OptMacroExpandOnly
		default:
			arg := args[0]
			if len(arg) > 0 && arg[0] == '-' {
				fmt.Fprintf(g.Stderr, "gomacro: unrecognized option '%s'.\nTry 'gomacro --help' for more information\n", arg)
				return nil
			}
			repl = false
			if cmd.WriteDeclsAndStmtsToFile {
				g.Options |= OptCollectDeclarations | OptCollectStatements
			}
			g.Options &^= OptShowPrompt | OptShowEval | OptShowEvalType // cleared by default, overridden by -s, -v and -vv
			g.Options = (g.Options | set) &^ clear
			cmd.EvalFileOrDir(arg)

			g.Imports, g.Declarations, g.Statements = nil, nil, nil
		}
		args = args[1:]
	}
	if repl || forcerepl {
		g.Options |= OptShowPrompt | OptShowEval | OptShowEvalType // set by default, overridden by -s, -v and -vv
		g.Options = (g.Options | set) &^ clear
		ir.ReplStdin()
	}
	return nil
}

func (cmd *Cmd) Usage() error {
	fmt.Fprint(cmd.Interp.Stdout, `usage: gomacro [OPTIONS] [files-and-dirs]

  Recognized options:
    -c,   --collect          collect declarations and statements, to print them later
    -e,   --expr EXPR        evaluate expression
    -f,   --force-overwrite  option -w will overwrite existing files
    -h,   --help             show this help and exit
    -i,   --repl             interactive. start a REPL after evaluating expression, files and dirs.
                             default: start a REPL only if no expressions, files or dirs are specified
    -m,   --macro-only       do not execute code, only parse and macroexpand it.
                             useful to run gomacro as a Go preprocessor
    -n,   --no-trap          do not trap panics in the interpreter
    -t,   --trap             trap panics in the interpreter (default)
    -s,   --silent           silent. do NOT show startup message, prompt, and expressions results.
                             default when executing files and dirs.
    -v,   --verbose          verbose. show startup message, prompt, and expressions results.
                             default when executing an expression.
    -vv,  --very-verbose     as -v, and in addition show the type of expressions results.
                             default when executing a REPL
    -w,   --write-decls      write collected declarations and statements to *.go files.
                             implies -c
    -x,   --exec             execute parsed code (default). disabled by -m

    Options are processed in order, except for -i that is always processed as last.

    Collected declarations and statements can be also written to standard output
    or to a file with the REPL command :write
`)
	return nil
}

func (cmd *Cmd) EvalFilesAndDirs(filesAndDirs ...string) error {
	for _, fileOrDir := range filesAndDirs {
		err := cmd.EvalFileOrDir(fileOrDir)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cmd *Cmd) EvalFileOrDir(fileOrDir string) error {
	info, err := os.Stat(fileOrDir)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return cmd.EvalDir(fileOrDir)
	} else {
		return cmd.EvalFile(fileOrDir)
	}
}

func (cmd *Cmd) EvalDir(dirname string) error {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return err
	}
	for _, file := range files {
		filename := file.Name()
		if !file.IsDir() && strings.HasSuffix(filename, ".gomacro") {
			filename = fmt.Sprintf("%s%c%s", dirname, os.PathSeparator, filename)
			err := cmd.EvalFile(filename)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// use line comments for disclaimer: block comments prevent Go build tags from working
const disclaimer = `// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

`

func (cmd *Cmd) EvalFile(filename string) (err error) {
	g := cmd.Interp.Globals
	g.Declarations = nil
	g.Statements = nil
	saveFilename := g.Filename

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
		g.Filename = saveFilename
	}()
	g.Filename = filename
	var comments string
	comments, err = cmd.EvalReader(f)
	if err != nil {
		return err
	}

	if cmd.WriteDeclsAndStmtsToFile {
		outname := filename
		if dot := strings.LastIndexByte(outname, '.'); dot >= 0 {
			// sanity check: dot must be in the file name, NOT in its path
			if slash := strings.LastIndexByte(outname, os.PathSeparator); slash < dot {
				outname = outname[0:dot]
			}
		}
		outname += ".go"
		if !cmd.OverwriteFiles {
			_, err := os.Stat(outname)
			if err == nil {
				g.Warnf("file exists already, use -f to force overwriting: %v", outname)
				return nil
			}
		}
		g.WriteDeclsToFile(outname, disclaimer, comments)

		if g.Options&OptShowEval != 0 {
			fmt.Fprintf(g.Stdout, "// processed file: %v\t-> %v\n", filename, outname)
		}
	}
	return nil
}

func (cmd *Cmd) EvalReader(src io.Reader) (comments string, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			switch rec := rec.(type) {
			case error:
				err = rec
			default:
				err = errors.New(fmt.Sprint(rec))
			}
		}
	}()
	in := MakeBufReadline(bufio.NewReader(src), cmd.Interp.Stdout)
	ir := cmd.Interp
	env := ir.Env
	env.Options &^= OptShowPrompt // parsing a file: suppress prompt
	env.Line = 0

	// perform the first iteration manually, to collect comments
	str, firstToken := env.ReadMultiline(in, ReadOptCollectAllComments)
	if firstToken >= 0 {
		comments = str[0:firstToken]
		if firstToken > 0 {
			str = str[firstToken:]
			env.IncLine(comments)
		}
	}
	if ir.ParseEvalPrint(str, in) {
		for ir.ReadParseEvalPrint(in) {
		}
	}
	return comments, nil
}
