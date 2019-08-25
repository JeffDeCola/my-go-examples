// my-go-examples interface-as-a-return.go

package main

import (
	"github.com/JeffDeCola/my-go-examples/basic-syntax/interfaces/interface-package/example"
)

func main() {

	// INTERFACE AS A RETURN
	// Get the interface via a function
	a := example.MakemyStructA("jeff")
	b := example.MakemyStructB(222, 333)

	// INTERFACE AS A FUNCTION PARAMETER
	// The interface figures out what method to use based on data type
	// It's magic.
	example.Magic(a) // I'm in doThis() method with receiver myStructA - jeff
	example.Magic(b) // I'm in doThis() method with receiver myStructB - 222 333

	// UPDATE THE DATA
	example.UpdateName(a, "Jill")
	example.Magic(a) // I'm in doThis() method with receiver myStructA - Jill

}
