// #!/usr/bin/env gomacro

package main

import "fmt"

type Person struct {
	Name, Surname string
}

type Driver struct {
	CanDrive []string
	Person
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}

var p = Person{"John", "Smith"}
var d = Driver{nil, p}

func pd() (Person, Driver) { return p, d }

type FmtPair struct{ a, b fmt.Stringer }

func fmtpair(a, b fmt.Stringer) FmtPair { return FmtPair{a, b} }

func conv() {
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", d)

	fmtpair(pd())

	var s fmt.Stringer
	var ch = make(chan fmt.Stringer, 2)

	s = p
	ch <- s
	fmt.Printf("%v\n", <-ch)

	s = d
	ch <- s
	fmt.Printf("%v\n", <-ch)

	ch <- p
	fmt.Printf("%v\n", <-ch)

	ch <- d
	fmt.Printf("%v\n", <-ch)

	fp := func() fmt.Stringer { return p }
	fmt.Printf("%v\n", fp())

	fd := func() fmt.Stringer { return d }
	fmt.Printf("%v\n", fd())
}
