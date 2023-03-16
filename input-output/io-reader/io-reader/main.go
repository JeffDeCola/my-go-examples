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
		// Returns byte count till EOF
		n, err := r.Read(buffer)
		if err != nil {
			if err == io.EOF {
				bufferOutput := string(buffer[:n])
				fmt.Printf(bufferOutput + "\n")
				return nil
			}
			return fmt.Errorf("error reading io.Reader: %w", err)
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
	sourceString := "This is data in string.Reader"
	s := strings.NewReader(sourceString)

	// STREAM TO BUFFER (s is *strings.Reader)
	err := readToBufferandPrint(s)
	if err != nil {
		fmt.Printf("Error reading from string: %s\n", err)
	}

	// ---------------------------------------------------
	// READING FROM A BUFFER (*bytes.Reader)
	fmt.Printf("\nREADING FROM A BUFFER\n")

	// CREATE THE BUFFER
	b := new(bytes.Buffer)
	b.WriteString("This is data in byte.buffer")

	// STREAM TO BUFFER (b is *bytes.Reader)
	err = readToBufferandPrint(b)
	if err != nil {
		fmt.Printf("Error reading from buffer: %s\n", err)
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

	// CREATE THE STDIN
	i := os.Stdin

	// STREAM TO BUFFER (os.Stdin)
	err = readToBufferandPrint(i)
	if err != nil {
		fmt.Printf("Error reading from Stdin: %s\n", err)
	}

	// ---------------------------------------------------
	// READING FROM A PIPE (*io.PipeReader)
	fmt.Printf("\nREAD FROM A PIPE\n")

	// CREATE THE PIPE
	pipeReader, pipeWriter := io.Pipe()
	// WRITE INTO PIPE USING GO ROUTINE
	go func() {
		fmt.Fprint(pipeWriter, "This is data being put into a pipe")
		// Using Close method to close write
		pipeWriter.Close()
	}()

	// STREAM TO BUFFER (rpipe is *io.PipeReader)
	err = readToBufferandPrint(pipeReader)
	if err != nil {
		fmt.Printf("Error reading from pipe: %s\n", err)
	}

}
