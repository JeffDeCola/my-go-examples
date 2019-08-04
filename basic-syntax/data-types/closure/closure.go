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

	// CLOSURE METHOD ONE - Assign anonymous function to a variable
	// This is also called a func literal.
	increment1 := func() int { // Got x = 5
		x++
		return x
	}

	// When you call function increment1, it uses x, which is outside function.
	fmt.Println(increment1()) // 6
	fmt.Println(increment1()) // 7
	fmt.Println(x)            // 7
	fmt.Println(increment1()) // 8
	x = 30
	fmt.Println(x)            // 30
	fmt.Println(increment1()) // 31 <- NOTE THIS
	fmt.Println(x)            // 31

	// CLOSURE METHOD TWO - Return a function to a function
	// This is different because now the value is bound to the variable.
	// The function captures the scope it is in
	// Think of it as an assigned variable, BECAUSE it is.
	increment2 := inc(x) // Got x = 31
	increment3 := inc(x) // Got x = 31

	// When you call function increment2, it uses x, which is outside function.
	fmt.Println(increment2()) // 32
	fmt.Println(increment2()) // 33
	fmt.Println(increment2()) // 34
	fmt.Println(x)            // 31

	// increment2 has its own value of x
	// Treat the function like a variable, because it is.
	fmt.Println(increment3()) // 32
	fmt.Println(increment3()) // 33
	fmt.Println(increment3()) // 34
	x = 20
	fmt.Println(x) // 20
	// Nothing to do with x because we assigned inc() to to increment3
	fmt.Println(increment3()) // 35 <- NOTE THIS
	fmt.Println(increment3()) // 36

}
