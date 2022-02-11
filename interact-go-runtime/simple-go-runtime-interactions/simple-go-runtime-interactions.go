package main

import (
	"fmt"
	"runtime"
)

func main() {

	// NUMBER CPU
	numCPUs := runtime.NumCPU()
	fmt.Printf("Number of CPUS: %d\n", numCPUs)

	// NUMBER GOROUTINES
	numGoRoutines := runtime.NumGoroutine()
	fmt.Printf("Number of Go Routines that exist: %d\n", numGoRoutines)

}
