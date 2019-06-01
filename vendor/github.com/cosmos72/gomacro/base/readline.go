/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * readline.go
 *
 *  Created on: Apr 02, 2018
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	// "os/signal"

	"github.com/peterh/liner"
)

type Readline interface {
	Read(prompt string) ([]byte, error)
}

// -------------------- BufReadline --------------------

type BufReadline struct {
	in  *bufio.Reader
	out io.Writer
}

func MakeBufReadline(in *bufio.Reader, out io.Writer) BufReadline {
	return BufReadline{in, out}
}

var (
	paragraph_separator_bytes = []byte{0xe2, 0x80, 0xa9}
	nl_bytes                  = []byte{'\n'}
)

func (buf BufReadline) Read(prompt string) ([]byte, error) {
	line, err := buf.in.ReadBytes('\n')
	line = bytes.Replace(line, paragraph_separator_bytes, nl_bytes, -1)
	return line, err
}

// -------------------- TtyReadline --------------------

type TtyReadline struct {
	Term *liner.State
}

func MakeTtyReadline(historyfile string) (TtyReadline, error) {
	tty := TtyReadline{liner.NewLiner()}

	/*
		go func() {
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt, os.Kill)
			sig := <-c
			signal.Stop(c)

			fmt.Fprintf(os.Stderr, "\nreceived signal: %v\n", sig)
			tty.Close(historyfile)
		}()
	*/

	if len(historyfile) == 0 {
		return tty, nil
	}
	f, err := os.Open(historyfile)
	if err != nil {
		return tty, err
	}
	defer f.Close()
	_, err = tty.Term.ReadHistory(f)
	return tty, err
}

func (tty TtyReadline) Read(prompt string) ([]byte, error) {
	line, err := tty.Term.Prompt(prompt)
	if len(line) >= 3 {
		tty.Term.AppendHistory(line)
	}
	if n := len(line); n != 0 || err != io.EOF {
		b := make([]byte, n+1)
		copy(b, line)
		b[n] = '\n'
		b = bytes.Replace(b, paragraph_separator_bytes, nl_bytes, -1)
		return b, err
	}
	return nil, err
}

func (tty TtyReadline) Close(historyfile string) (err error) {
	if len(historyfile) == 0 {
		return tty.Term.Close()
	}
	f, err1 := os.OpenFile(historyfile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err1 != nil {
		err = fmt.Errorf("could not open %q to append history: %v", historyfile, err1)
	} else {
		defer f.Close()
		_, err2 := tty.Term.WriteHistory(f)
		if err2 != nil {
			err = fmt.Errorf("could not write history to %q: %v", historyfile, err2)
		}
	}
	err3 := tty.Term.Close()
	if err3 != nil {
		err = err3
	}
	return
}
