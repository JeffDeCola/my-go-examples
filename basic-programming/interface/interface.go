package main

import (
	"fmt"
)

// Interface has the method doThis for
// both a and b
type theInterface interface {
	doThis(n int)
}

// DEFINE a
type a struct {
	name string
}

func (i *a) doThis(n int) {
	fmt.Printf("I'm in A %v - %v\n", n, i.name)
}

// DEFINE b
type b struct {
	x int
	y int
}

func (i *b) doThis(n int) {
	fmt.Printf("I'm in B %v - %v %v\n", n, i.x, i.y)
}

// The interface accepts ANYTHING as long as that
// ANYTHING has a method attached to this interface
func magic(i theInterface, n int) {
	i.doThis(n)
}

func main() {

	// Declare assign a struct.
	var a1 = &a{"jeff"}
	var b1 = &b{222, 333}
	var b2 = &b{}

	// I can send magic an argument of ANYTHING
	// as long as that ANYTHING has the underlying method
	// attached to the interface.  It's magic.
	magic(a1, 44)  // I'm in A 44 - jeff
	magic(b1, 111) // I'm in B 111 - 222 333
	magic(b2, 121) // I'm in B 121 - 0 0

}
