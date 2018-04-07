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
	// run_for_nested()
	run_interface_method_to_closure()
	run_struct_method_to_closure()
}

type stringer interface{ String() string }
type Box struct{ value int }

func (b *Box) Value() int {
	return b.value
}

func run_struct_method_to_closure() {
	var b *Box
	fmt.Printf("%v %T\n", b, b)
	function := (*Box).Value
	fmt.Printf("%v %T\n", function, function)
	closure := b.Value
	fmt.Printf("%v %T\n", closure, closure)
}

func run_interface_method_to_closure() {
	var s stringer
	fmt.Printf("%v %T\n", s, s)
	function := stringer.String
	fmt.Printf("%v %T\n", function, function)
	closure := s.String
	fmt.Printf("%v %T\n", closure, closure)
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
