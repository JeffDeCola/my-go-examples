// my-go-examples recursion1.go

package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	switch {
	case n == 0:
		return 0
	case n == 1:
		return 1
	case n <= 45:
		fibN := fibonacci(n-2) + fibonacci(n-1)
		return fibN
	default:
		fmt.Printf("ERROR\n")
		return 0
	}
}

func main() {

	start := time.Now()
	var n int

	fmt.Printf("Enter the number (less than 45) of items in the series you want to calculate: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	fmt.Println("The total is", fibonacci(n))
	fmt.Printf("Finished this in %f seconds\n", time.Since(start).Seconds())

}
