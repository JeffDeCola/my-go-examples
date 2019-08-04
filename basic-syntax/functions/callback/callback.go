package main

import "fmt"

// Add all even numbers together and return the sum
// use callback to check for even numbers
func evenSum(numbers []int, evenCallback func(int) bool) int {
	sum := 0
	for _, n := range numbers {
		// Only get sum if even
		if evenCallback(n) {
			sum = sum + n
		}
	}
	return sum
}

func main() {

	numbers := []int{4, 5, 6, 7}

	// Get the total sum of the even numbers.
	// Pass the function (as an argument) a function
	// that tests for an even number.
	sum := evenSum(numbers, func(n int) bool {
		if n%2 == 0 {
			return true
		}
		return false
	})

	fmt.Println(sum)

}
