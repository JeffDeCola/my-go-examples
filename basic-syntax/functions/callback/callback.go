// my-go-examples callback.go

package main

import "fmt"

// Add all numbers together and return the sum
// Use callback to check for even/odd numbers
func sum(numbers []int, mycallbackfunction func(int) bool) int {
	sum := 0
	for _, n := range numbers {
		// Only add to sum on the results of this test
		if mycallbackfunction(n) {
			sum = sum + n
		}
	}
	return sum
}

func main() {

	numbers := []int{4, 5, 6, 7}
	fmt.Println("The numbers in the slice are", numbers)

	// Giving an int, determine if even
	anonFunc1 := func(n int) bool {
		if n%2 == 0 {
			return true
		}
		return false
	}
	// Giving an int, determine if odd
	anonFunc2 := func(n int) bool {
		if n%2 != 0 {
			return true
		}
		return false
	}

	// CALLBACK - PASS A FUNCTION TO A FUNCTION
	sum1 := sum(numbers, anonFunc1)
	sum2 := sum(numbers, anonFunc2)

	fmt.Println("The sum of the even numbers are", sum1)
	fmt.Println("The sum of the odd numbers are", sum2)

}
