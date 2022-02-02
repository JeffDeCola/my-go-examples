package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(r io.Reader) error {

	// CREATE SMALL BUFFER TO WRITE INTO
	buffer := make([]byte, 10)

	// LETS STREAM FROM READER INTO BUFFER AND PRINT
	for {

		// READ METHOD
		// Returns byte count till EOF
		n, err := r.Read(buffer)
		if err != nil {
			if err == io.EOF {
				bufferString := string(buffer[:n])
				fmt.Printf(bufferString + "\n")
				return nil
			}
			return fmt.Errorf("error reading io.Reader: %w", err)
		}

		// PRINT THE BUFFER AND LOOP
		bufferString := string(buffer[:n])
		fmt.Printf(bufferString + "\n")

	}

}

func main() {

	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("Error reading file %s", err)
	}
	defer f.Close()

	err = readFile(f)
	if err != nil {
		fmt.Printf("Error reading file %s", err)
	}

}
