package main

import "fmt"

// Returns a function
func inc(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {

	x := 5

	// CLOSURE METHOD TWO - Return a function to a function
	// This is different because now the value is bound to the variable increment.
	// The function captures the scope it is in
	// Think of it as an assigned variable, BECAUSE it is.
	increment := inc(x)

	// When you call function increment1, it uses x, which is outside function.
	fmt.Println(increment()) // 6
	fmt.Println(increment()) // 7
	fmt.Println(x)           // 5 (NOT 7)
	fmt.Println(increment()) // 8
	x = 30
	fmt.Println(x)           // 30
	fmt.Println(increment()) // 9 (NOT 31)
	fmt.Println(x)           // 30 (NOT 31)

}
