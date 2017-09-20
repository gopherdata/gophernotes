// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

import (
	"go/token"
)

// -----------------------------------------------------------------------------
// File

// A File is a handle for a file belonging to a FileSet.
// A File has a name, size, and line offset table.
//
type File struct {
	*token.File
	line int // starting line of this file
}

// PositionFor returns the Position value for the given file position p.
// If adjusted is set, the position may be adjusted by position-altering
// //line comments; otherwise those comments are ignored.
// p must be a Pos value in f or NoPos.
//
func (f *File) PositionFor(p token.Pos, adjusted bool) (pos token.Position) {
	pos = f.File.PositionFor(p, adjusted)
	if pos.IsValid() {
		pos.Line += f.line
	}
	return pos
}

// Position returns the Position value for the given file position p.
// Calling f.Position(p) is equivalent to calling f.PositionFor(p, true).
//
func (f *File) Position(p token.Pos) (pos token.Position) {
	return f.PositionFor(p, true)
}

// -----------------------------------------------------------------------------
// FileSet

// A FileSet represents a set of source files.
// This is a wrapper for go/token.FileSet that adds a starting line offset to each file in the set
//
type FileSet struct {
	token.FileSet
	filemap map[*token.File]*File
}

// NewFileSet creates a new file set.
func NewFileSet() *FileSet {
	return &FileSet{
		FileSet: *token.NewFileSet(),
		filemap: make(map[*token.File]*File),
	}
}

// AddFile adds a new file with a given filename, base offset, and file size
func (s *FileSet) AddFile(filename string, base, size, line int) *File {
	innerf := s.FileSet.AddFile(filename, base, size)
	f := &File{File: innerf, line: line}
	s.filemap[innerf] = f
	return f
}

// File returns the file that contains the position p.
// If no such file is found (for instance for p == NoPos),
// the result is nil.
//
func (s *FileSet) File(p token.Pos) (f *File) {
	if p != token.NoPos {
		innerf := s.FileSet.File(p)
		f = s.filemap[innerf]
	}
	return
}

// PositionFor converts a Pos p in the fileset into a Position value.
// If adjusted is set, the position may be adjusted by position-altering
// //line comments; otherwise those comments are ignored.
// p must be a Pos value in s or NoPos.
//
func (s *FileSet) PositionFor(p token.Pos, adjusted bool) (pos token.Position) {
	if f := s.File(p); f != nil {
		pos = f.PositionFor(p, adjusted)
	}
	return
}

// Position converts a Pos p in the fileset into a Position value.
// Calling s.Position(p) is equivalent to calling s.PositionFor(p, true).
//
func (s *FileSet) Position(p token.Pos) (pos token.Position) {
	return s.PositionFor(p, true)
}
