package main

import (
	"fmt"
)

type theInterface interface {
	doThis(n int)
}

// STRUCT A
type myStructA struct {
	name string
}

func (i *myStructA) doThis(n int) {
	fmt.Printf("I'm in myStructA %v - %v\n", n, i.name)
}

// STRUCT B
type myStructB struct {
	x int
	y int
}

func (i *myStructB) doThis(n int) {
	fmt.Printf("I'm in myStructB %v - %v %v\n", n, i.x, i.y)
}

// INTERFACE AS A PARAMETER FUNCTION
//   The interface accepts ANYTHING as long as that
//   ANYTHING has the method doThis()
func magic(i theInterface, n int) { // <- NOTE THIS
	i.doThis(n)
}

func main() {

	// Declare and assign the structs.
	var a1 = &myStructA{"jeff"}
	var b1 = &myStructB{222, 333}
	var b2 = &myStructB{}

	// I can send magic an argument of ANYTHING
	// as long as that ANYTHING has the underlying method
	// attached to the interface.  It's magic.
	magic(a1, 44)  // I'm in myStructA 44 - jeff
	magic(b1, 111) // I'm in myStructB 111 - 222 333
	magic(b2, 121) // I'm in myStructB 121 - 0 0

}
