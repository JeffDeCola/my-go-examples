// my-go-examples logging-error-handling.go

package main

import (
	"errors"
	"log"
	"os"
)

func addThis(x int, y int, logger *log.Logger) (int, error) {
	sum := x + y
	var err error = nil
	if sum > 4 {
		err = errors.New("math: The sum is too large")
		return sum, err
	}
	return sum, err
}

func addThese(x int, y int, logger *log.Logger) int {
	sum := x + y
	return sum
}

func main() {

	// Create Logger
	var logger = log.New(os.Stderr, "JeffLogging:", log.Lshortfile)

	logger.Println("Print Logging info")

	logger.Println("Calling function addThis() that does not returns error")
	sum, err := addThis(2, 2, logger)
	if err != nil {
		logger.Fatalf("Error: %s", err)
	}
	logger.Println("Sum is", sum)

	logger.Println("Calling function addThese() that does not return error")
	sum = addThese(2, 2, logger)
	logger.Println("Sum is", sum)

	logger.Println("Calling function addThis() again, that returns error")
	sum, err = addThis(2, 4, logger)
	if err != nil {
		logger.Fatalf("Error: %s", err)
	}
	logger.Println("Sum is", sum)

}
