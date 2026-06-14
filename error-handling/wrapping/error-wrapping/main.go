// error-wrapping
//
// Wrapping errors with %w to add context as they propagate up the call stack
// using the errors standard package.
//

package main

import (
	"errors"
	"fmt"
	"os"
)

func loadConfig(filename string) error {
	err := readConfigFile(filename)
	if err != nil {
		return fmt.Errorf("loadConfig: %w", err)
	}
	return nil
}

func readConfigFile(filename string) error {
	err := checkFilename(filename)
	if err != nil {
		return fmt.Errorf("readConfigFile: %w", err)
	}
	return nil
}

func checkFilename(filename string) error {
	if filename == "" {
		return errors.New("filename can not be empty")
	}
	return nil
}

func main() {

	// Good
	fmt.Println("PART 1 - GOOD - pass filename")
	err := loadConfig("data.txt")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("    Everything is good")

	// Bad
	fmt.Println("PART 2 - BAD - pass no filename")
	err = loadConfig("")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("    Everything is good")
}
