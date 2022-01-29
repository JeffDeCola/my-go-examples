package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func readToBufferandPrint(r io.Reader) error {

	// CREATE SMALL BUFFER TO WRITE INTO
	buffer := make([]byte, 20)

	// LETS STREAM FROM READER INTO BUFFER AND PRINT
	for {

		// READ METHOD
		// Retruns count till EOF
		n, err := r.Read(buffer)
		if err != nil {
			if err == io.EOF {
				bufferOutput := string(buffer[:n])
				fmt.Printf(bufferOutput + "\n")
				return nil
			}
			return fmt.Errorf("error reading Stdin: %w", err)
		}

		// PRINT THE BUFFER AND LOOP
		bufferOutput := string(buffer[:n])
		if bufferOutput == "stop\n" {
			fmt.Printf("Stopping\n")
			return nil
		}
		fmt.Printf(bufferOutput + "\n")

	}

}

func main() {

	// ---------------------------------------------------
	// READING FROM A STRING (*string.Reader)
	fmt.Printf("\nREADING FROM A STRING\n")

	// CREATE THE STRING READER
	sourceString := "This data is being put into a string reader"
	s := strings.NewReader(sourceString)

	// STREAM TO BUFFER (s is *strings.Reader)
	err := readToBufferandPrint(s)
	if err != nil {
		fmt.Printf("Error reading string: %s\n", err)
	}

	// ---------------------------------------------------
	// READING FROM A BUFFER (*bytes.Reader)
	fmt.Printf("\nREADING FROM A BUFFER\n")

	// CREATE THE READER
	sourceBuffer := []byte("This data is being put into a byte reader")
	b := bytes.NewReader(sourceBuffer)

	// STREAM TO BUFFER (b is *bytes.Reader)
	err = readToBufferandPrint(b)
	if err != nil {
		fmt.Printf("Error reading bytes: %s\n", err)
	}

	// ---------------------------------------------------
	// READING FROM A FILE (*os.File)
	fmt.Printf("\nREAD FROM FILE\n")

	// OPEN THE FILE
	filename := "test.txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file: %s", err)
	}
	defer f.Close()

	// STREAM TO BUFFER (f is *os.File)
	err = readToBufferandPrint(f)
	if err != nil {
		fmt.Printf("Error with reading file: %s\n", err)
	}

	// ---------------------------------------------------
	// READING FROM STDIN (os.Stdin)
	fmt.Printf("\nREAD FROM STDIN - type 'stop' to continue\n")

	// STREAM TO BUFFER (os.Stdin)
	err = readToBufferandPrint(os.Stdin)
	if err != nil {
		fmt.Printf("Error reading Stdin: %s\n", err)
	}

	// ---------------------------------------------------
	// READING FROM A PIPE (*io.PipeReader)
	fmt.Printf("\nREAD FROM A PIPE\n")

	// CREATE THE PIPE
	rpipe, wpipe := io.Pipe()
	// WRITE INTO PIPE USING GO ROUTINE
	go func() {
		fmt.Fprint(wpipe, "This data is being put into a pipe")
		// Using Close method to close write
		wpipe.Close()
	}()

	// STREAM TO BUFFER (rpipe is *io.PipeReader)
	err = readToBufferandPrint(rpipe)
	if err != nil {
		fmt.Printf("Error reading pipe: %s\n", err)
	}

}
