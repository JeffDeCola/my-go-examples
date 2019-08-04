package main

import (
	"fmt"
)

// Interface has the method doThis for
// both a and b
type theInterface interface {
	doThis(n int)
}

type myStructA struct {
	name string
}

func (i *myStructA) doThis(n int) {
	fmt.Printf("I'm in myStructA %v - %v\n", n, i.name)
}

type myStructB struct {
	x int
	y int
}

func (i *myStructB) doThis(n int) {
	fmt.Printf("I'm in myStructB %v - %v %v\n", n, i.x, i.y)
}

// The interface accepts ANYTHING as long as that
// ANYTHING has a method attached to this interface
func makemyStructA(name string) theInterface { // <- NOTE THIS
	return &myStructA{name}
}

func makemyStructB(x, y int) theInterface { // <- NOTE THIS
	return &myStructB{x, y}
}

// INTERFACE AS A PARAMETER FUNCTION
//   The interface accepts ANYTHING as long as that
//   ANYTHING has the method doThis()
func magic(i theInterface, n int) {
	i.doThis(n)
}

func main() {

	// Get the interface.
	a1 := makemyStructA("jeff")
	b1 := makemyStructB(222, 333)
	b2 := makemyStructB(0, 0)

	// I can see what's in the interface
	fmt.Printf("%+v\n", a1)
	fmt.Printf("%+v\n", b1)
	fmt.Printf("%+v\n", b2)

	// I can send magic an argument of ANYTHING
	// as long as that ANYTHING has the underlying method
	// attached to the interface.  It's magic.
	magic(a1, 44)  // I'm in myStructA 44 - jeff
	magic(b1, 111) // I'm in myStructB 111 - 222 333
	magic(b2, 121) // I'm in myStructB 121 - 0 0

}
