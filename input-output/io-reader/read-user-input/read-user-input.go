package main

import (
	"fmt"
	"io"
	"os"
)

func readUsingReadMethod(r io.Reader) (string, error) {

	// CREATE SMALL BUFFER TO WRITE INTO
	buffer := make([]byte, 10)

	// READ METHOD
	// Returns byte count till EOF
	n, err := r.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("error reading io.Reader: %w", err)
	}

	bufferString := string(buffer[:n])
	return bufferString, nil
}

func readUsingFscan(r io.Reader) (string, error) {

	var input string

	_, err := fmt.Fscan(r, &input)
	if err != nil {
		return "", fmt.Errorf("unable to get string from user: %w", err)
	}

	return input, nil

}

func main() {

	// CREATE THE STDIN
	i := os.Stdin
	var userInput string

	fmt.Println("Read user intput using read method")
	userInput, err := readUsingReadMethod(i)
	if err != nil {
		fmt.Printf("Error reading from Stdin: %s\n", err)
	}
	fmt.Printf("User typed %s\n", userInput)

	fmt.Println("Read user intput using Fscan")
	userInput, err = readUsingFscan(i)
	if err != nil {
		fmt.Printf("Error reading from Stdin: %s\n", err)
	}
	fmt.Printf("User typed %s\n", userInput)

}
