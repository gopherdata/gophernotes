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
	term *liner.State
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

	f, err := os.Open(historyfile)
	if err != nil {
		return tty, err
	}
	defer f.Close()
	_, err = tty.term.ReadHistory(f)
	return tty, err
}

func (tty TtyReadline) Read(prompt string) ([]byte, error) {
	line, err := tty.term.Prompt(prompt)
	if len(line) != 0 {
		tty.term.AppendHistory(line)
	}
	if n := len(line); n != 0 || err != io.EOF {
		bytes := make([]byte, n+1)
		copy(bytes, line)
		bytes[n] = '\n'
		return bytes, err
	}
	return nil, err
}

func (tty TtyReadline) Close(historyfile string) (err error) {
	if len(historyfile) == 0 {
		return tty.term.Close()
	}
	f, err1 := os.OpenFile(historyfile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err1 != nil {
		err = fmt.Errorf("could not open %q to append history: %v", historyfile, err1)
	} else {
		defer f.Close()
		_, err2 := tty.term.WriteHistory(f)
		if err2 != nil {
			err = fmt.Errorf("could not write history to %q: %v", historyfile, err2)
		}
	}
	err3 := tty.term.Close()
	if err3 != nil {
		err = err3
	}
	return
}
