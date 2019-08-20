package main

import "fmt"

func main() {

	x := 5

	// CLOSURE METHOD ONE
	// Assign anonymous function to a variable
	// This is also called a func literal.
	increment := func() int { // Got x = 5
		x++
		return x
	}

	// When you call function increment1, it uses x, which is outside function.
	fmt.Println(increment()) // 6
	fmt.Println(increment()) // 7
	fmt.Println(x)           // 7
	fmt.Println(increment()) // 8
	x = 30
	fmt.Println(x)           // 30
	fmt.Println(increment()) // 31
	fmt.Println(x)           // 31
}
