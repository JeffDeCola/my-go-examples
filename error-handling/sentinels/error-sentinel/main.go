// error-sentinel
//
// A sentinel is a named error that callers can check for by identity
// using the errors standard package.
//

package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrFileNotFound = errors.New("file not found")

func checkFilename(f string) error {
	if f == "" {
		return fmt.Errorf("checkFilename: %w", ErrFileNotFound)
	}
	return nil
}

func main() {

	// BAD
	filename := ""
	err := checkFilename(filename)
	if errors.Is(err, ErrFileNotFound) {
		// Recover
		fmt.Println("no file - using default data.txt")
		filename = "data.txt"
	} else if err != nil {
		// some OTHER error - can't handle it, bail
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("carrying on with:", filename)

	// GOOD
	filename = "something.txt"
	err = checkFilename(filename)
	if errors.Is(err, ErrFileNotFound) {
		// Recover
		fmt.Println("no file - using default data.txt")
		filename = "data.txt"
	} else if err != nil {
		// some OTHER error - can't handle it, bail
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("carrying on with:", filename)

}
