// Copyright 2018 Massimiliano Ghilardi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gls

import (
	"testing"
)

var verbose bool = false

func AsyncGoID() <-chan uintptr {
	ch := make(chan uintptr)
	go func() {
		ch <- GoID()
	}()
	return ch
}

func TestGoID(t *testing.T) {
	id1 := GoID()
	id2 := GoID()
	if id1 == id2 {
		if verbose {
			t.Logf("TestGoID: 0x%x == 0x%x", id1, id2)
		}
	} else {
		t.Errorf("TestGoID: 0x%x != 0x%x", id1, id2)
	}
}

func TestAsyncGoID1(t *testing.T) {
	id1 := GoID()
	id2 := <-AsyncGoID()
	if id1 != id2 {
		if verbose {
			t.Logf("TestAsyncGoID1: 0x%x != 0x%x", id1, id2)
		}
	} else {
		t.Errorf("TestAsyncGoID1: 0x%x == 0x%x", id1, id2)
	}
}

func TestAsyncGoID2(t *testing.T) {
	ch1 := AsyncGoID()
	ch2 := AsyncGoID()
	id1 := <-ch1
	id2 := <-ch2
	if id1 != id2 {
		if verbose {
			t.Logf("TestAsyncGoID2: 0x%x != 0x%x", id1, id2)
		}
	} else {
		t.Errorf("TestAsyncGoID2: 0x%x == 0x%x", id1, id2)
	}
}
