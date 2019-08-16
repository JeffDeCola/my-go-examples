// my-go-examples flag.go

package main

import (
	"flag"
	"fmt"
)

func main() {

	// STRING
	stringPtr := flag.String("s", "default", "This is the flag for a string")

	// INTEGER
	integerPtr := flag.Int("i", 1, "This is the flag for an integer")

	// BOOLEAN
	boolPtr := flag.Bool("b", false, "This is the flag for a boolean")

	// Parse the flags
	flag.Parse()

	// Print out the flags
	fmt.Printf("\n")
	fmt.Printf("String flag is %s\n", *stringPtr)
	fmt.Printf("Integer flag is %d\n", *integerPtr)
	fmt.Printf("Integer flag is %t\n\n", *boolPtr)

}
