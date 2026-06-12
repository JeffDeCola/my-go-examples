// constructor-simple
//
// Constructors are functions that build and return a new instance of a struct,
// often with defaults.

package main

import "fmt"

type thing struct {
	a int
	s string
}

func newThing(a int) *thing {

	// build
	ss := "default"

	// assemble
	t := thing{
		a: a,
		s: ss,
	}

	// return
	return &t

}

func main() {

	t := newThing(4)
	fmt.Printf("The thing is %+v\n", *t)

	t.a = 5
	t.s = "happy"
	fmt.Printf("The thing is %+v\n", *t)

}
