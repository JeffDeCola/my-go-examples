package main

import "fmt"

func fibonacci(n int) int {
	switch {
	case n == 0:
		return 0
	case n == 1:
		return 1
	case n <= 40:
		fibn := fibonacci(n-1) + fibonacci(n-2)
		return fibn
	default:
		return 0
	}
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	fmt.Println(fibonacci(n))
}
