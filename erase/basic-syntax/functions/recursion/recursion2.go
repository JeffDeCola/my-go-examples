// my-go-examples recursion1.go

package main

import (
	"fmt"
	"time"
)

type fibS struct {
	series []int
	total  int
}

// Always going to call from 0 first
func (f *fibS) fib(i int, n int, getnumber bool) int {

	sum := 0
	if getnumber {
		return f.series[i]
	}

	switch {
	case i == 0:
		sum = 0
		f.total = f.total + sum
	case i == 1:
		sum = 1
		f.total = f.total + sum
	default:
		sum = f.fib(i-2, n, true) + f.fib(i-1, n, true)
		f.total = f.total + sum
	}

	f.series = append(f.series, sum)
	// fmt.Println(i, n, f.series, sum)

	// Call the next one
	if i != n {
		f.fib(i+1, n, false)
	}

	return f.total
}

func main() {

	start := time.Now()
	var n int
	var f fibS

	fmt.Printf("Enter the number of items in the series (less then 70) you want to calculate: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	total := f.fib(0, n-1, false)
	fmt.Printf("%v numbers of the series is\n", n)
	fmt.Printf("%v\n", f.series)
	fmt.Printf("The total is %v\n", total)
	fmt.Printf("Finished this in %f seconds\n", time.Since(start).Seconds())

}
