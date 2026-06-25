// SLICES EXAMPLE
//
// _A flexible dynamically-sized array._
//

package main

import "fmt"

func main() {
	// DECLARE
	var a []float64
	fmt.Println("Declare:              ", a)

	// ASSIGN (append to grow)
	a = append(a, 1.1)
	a = append(a, 2.0)
	fmt.Println("Assign:               ", a)

	// DECLARE & INITIALIZE
	var b = []float32{1.1, 2.0}
	fmt.Println("Declare & Initialize: ", b)
	c := []float32{4.4, 5.0}
	fmt.Println("Declare & Initialize: ", c)

	// MAKE (preallocate length & capacity)
	d := make([]int, 2, 25)
	fmt.Printf("Make:                  %v  len=%d cap=%d\n", d, len(d), cap(d))
	d[0] = 6
	d[1] = 7
	fmt.Printf("Make:                  %v  len=%d cap=%d\n", d, len(d), cap(d))

}
