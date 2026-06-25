// ARRAYS EXAMPLE
//
// A fixed-length collection of elements of the same type.
//

package main

import "fmt"

func main() {

	// DECLARE
	var a [2]float32
	fmt.Println("Declare:              ", a)

	// ASSIGN
	a[0] = 1.1
	a[1] = 2.0
	fmt.Println("Assign:               ", a)

	// DECLARE & INITIALIZE
	var b = [2]float32{1.1, 2.0}
	fmt.Println("Declare & Initialize: ", b)
	c := [2]float32{4.4, 5.0}
	fmt.Println("Declare & Initialize: ", c)

}
