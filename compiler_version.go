// +build !go1.11

// cause a compile error if Go compiler version < 1.11

package main

var _ int = "error: Go >= 1.11 required to compile Gophernotes"
