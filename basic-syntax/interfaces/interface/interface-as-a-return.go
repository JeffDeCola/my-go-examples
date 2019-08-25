// my-go-examples interface-as-a-return.go

package main

import (
	"fmt"
)

type myInterfacer interface {
	doThis()
}

type myStructA struct {
	name string
}

type myStructB struct {
	x int
	y int
}

func (i *myStructA) doThis() {
	fmt.Printf("I'm in doThis() method with receiver myStructA - %v\n", i.name)
}

func (i *myStructB) doThis() {
	fmt.Printf("I'm in doThis() method with receiver myStructB - %v %v\n", i.x, i.y)
}

// INTERFACE AS A RETURN
func makemyStructA(name string) myInterfacer {
	return &myStructA{name}
}

// INTERFACE AS A RETURN
func makemyStructB(x, y int) myInterfacer {
	return &myStructB{x, y}
}

// INTERFACE AS A FUNCTION PARAMETER
func magic(i myInterfacer) {
	i.doThis()
}

func main() {

	// INTERFACE AS A RETURN
	// Get the interface via a function
	a := makemyStructA("jeff")
	b := makemyStructB(222, 333)

	// INTERFACE AS A FUNCTION PARAMETER
	// The interface figures out what method to use based on data type
	// It's magic.
	magic(a) // I'm in doThis() method with receiver myStructA - jeff
	magic(b) // I'm in doThis() method with receiver myStructB - 222 333

}
