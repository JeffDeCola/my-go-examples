// error-simple
//
// Error handling using the errors standard package.
//

package main

import (
	"errors"
	"fmt"
	"os"
)

func checkFilename(filename string) error {

	if filename == "" {
		return errors.New("filename can not be empty")
	}
	return nil
}

func main() {

	// Good
	fmt.Println("PART 1 - GOOD - pass filename")
	err := checkFilename("data.txt")
	if err != nil {
		fmt.Println("CheckFilename failed:", err)
		os.Exit(1)
	}
	fmt.Println("    Everything is good")

	// Bad
	fmt.Println("PART 2 - BAD - pass no filename")
	err = checkFilename("")
	if err != nil {
		fmt.Println("CheckFilename failed:", err)
		os.Exit(1)
	}
	fmt.Println("    Everything is good")
}
