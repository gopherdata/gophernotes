// empty file. stops "go build" from complaining that
// no buildable files are in the directory "examples"

package main

import (
	"fmt"
	"io"
	"os"
	r "reflect"
)

func main() {
	conv()
}

func main2() {
	var TypeOfInterface = r.TypeOf((*interface{})(nil)).Elem()
	p := r.ValueOf(new(interface{}))
	i := p.Elem()
	c := i.Convert(TypeOfInterface)
	fmt.Printf("%v %v\n", p, p.Type())
	fmt.Printf("%v %v\n", i, i.Type())
	fmt.Printf("%v %v\n", c, c.Type())
	/*
		defer func() {
			fmt.Println(recover())
		}()
		defer func() {
			fmt.Println("foo")
		}()
		panic("test panic")
	*/
	// m := [...]int{0x7ffffff: 3}
	// fmt.Println(m)
	// p := Pair{A: 1, B: true}
	// Pair{1, 2} = Pair{}
	// var f os.file
	// _ = bytes.Buffer{nil, 0}
}

func main1() {
	var x io.ReadWriteCloser = os.Stdin
	f := io.ReadWriteCloser.Close
	f(x)
	fmt.Printf("%T\n", f)
}
