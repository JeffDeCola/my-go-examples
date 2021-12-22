// my-go-examples example.go

package example

import (
	"fmt"
)

type MyInterfacer interface {
	doThis()
	changeName(string)
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

func (i *myStructA) changeName(name string) {
	i.name = name
	fmt.Printf("I'm in changeName() method with receiver myStructA - %v\n", i.name)
}

func (i *myStructB) changeName(name string) {
	fmt.Printf("I'm in changeName() method with receiver myStructB - %v %v\n", i.x, i.y)
}

// MakemyStructA create the interface for myStructA
// INTERFACE AS A RETURN
func MakemyStructA(name string) MyInterfacer {
	return &myStructA{name}
}

// MakemyStructB create the interface for myStructB
// INTERFACE AS A RETURN
func MakemyStructB(x, y int) MyInterfacer {
	return &myStructB{x, y}
}

// Magic is well, magic
// INTERFACE AS A FUNCTION PARAMETER
func Magic(i MyInterfacer) {
	i.doThis()
}

//UpdateName Ability to update data
func UpdateName(i MyInterfacer, name string) {
	i.changeName(name)
}
