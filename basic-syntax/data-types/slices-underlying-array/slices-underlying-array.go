package main

import "fmt"

func main() {

	// Lets create and array and fill that array
	var x [10]int
	for i := 0; i < len(x); i++ {
		x[i] = i + 20
	}
	fmt.Println("x", x, len(x), cap(x)) // [21 22 23 24 25 26 27 28 29 30] 10 10

	// Lets create a slice and fill that slice (notice the underlying array is 16)
	var y []int
	for i := 0; i < 10; i++ {
		y = append(y, i+20)
	}
	fmt.Println("y", y, len(y), cap(y)) // [21 22 23 24 25 26 27 28 29 30] 10 16

	// Lets make the slice with an exact size (length 10, capacity 10)
	z := make([]int, 10)
	for i := 0; i < 10; i++ {
		z[i] = i + 20
	}
	fmt.Println("z", z, len(z), cap(z)) // [21 22 23 24 25 26 27 28 29 30] 10 10

	// Lets make the slice with an exact size (length 10, capacity 18)
	n := make([]int, 10, 18)
	for i := 0; i < 10; i++ {
		n[i] = i + 20
	}
	fmt.Println("n", n, len(n), cap(n)) // [21 22 23 24 25 26 27 28 29 30] 10 18

	// Lets make the slice from an array with an exact size (length 10, capacity 18)
	m := new([18]int)[0:10]
	for i := 0; i < 10; i++ {
		m[i] = i + 20
	}
	fmt.Println("m", m, len(m), cap(m)) // [21 22 23 24 25 26 27 28 29 30] 10 18

	// Same as above - Slice of an array two steps
	l := new([18]int)
	g := l[0:10]
	for i := 0; i < 10; i++ {
		g[i] = i + 20
	}
	fmt.Println("g", g, len(g), cap(g)) // [21 22 23 24 25 26 27 28 29 30] 10 18

	// Now lets demonstrate appending to a slice of length 0, capacity 5
	s := make([]int, 0, 5)
	for i := 0; i < 30; i++ {
		s = append(s, i+20)
		fmt.Println("s", s, len(s), cap(s))
	}

}
