package main

import "fmt"

func main() {

	// Lets create and array and fill that array
	var x [10]int
	for i := 0; i < len(x); i++ {
		x[i] = i + 20
	}
	fmt.Println(x, len(x), cap(x)) // [21 22 23 24 25 26 27 28 29 30] 10 10

	// Lets create a slice and fill that slice (notice the underlying array is 16)
	var y []int
	for i := 0; i < 10; i++ {
		y = append(y, i+20)
	}
	fmt.Println(y, len(y), cap(y)) // [21 22 23 24 25 26 27 28 29 30] 10 16

	// Lets make the slice with an exact size (capacity of array is 10)
	z := make([]int, 10)
	for i := 0; i < 10; i++ {
		z[i] = i + 20
	}
	fmt.Println(z, len(z), cap(z)) // [21 22 23 24 25 26 27 28 29 30] 10 10

	// Now lets demonstrate appending to a slice of length 0 and capacity 5
	s := make([]int, 0, 5)
	for i := 0; i < 30; i++ {
		s = append(s, i+20)
		fmt.Println(s, len(s), cap(s))
	}

}
