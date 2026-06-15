// generics-function
//
// A generic function uses a type parameter so a and b are the same type,
// and so is the result - one function for many types instead of copy-pasting.
//

package main

import "fmt"

// Number is a constraint - T may be any of these types
type Number interface {
	int | float64
}

// max works for ANY type that satisfies Number
func biggerNumber[T Number](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {

	// T is inferred from the arguments - no need to spell it out
	fmt.Println("MAX WITH DIFFERENT TYPES")
	fmt.Printf("    biggerNumber(3, 5)        is  %v\n", biggerNumber(3, 5))
	fmt.Printf("    biggerNumber(2.5, 1.5)    is  %v\n", biggerNumber(2.5, 1.5))

	// You CAN spell out the type, but rarely need to
	fmt.Printf("    biggerNumber[int](7, 2)   is  %v\n", biggerNumber[int](7, 2))

}
