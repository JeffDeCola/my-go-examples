package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func writeFromBuffer(w io.Writer) error {

	// CREATE BUFFER
	buffer := []byte("This data is being put into a byte reader\n")

	// WRITE METHOD
	// Returns byte count
	n, err := w.Write(buffer)
	if err != nil {
		return fmt.Errorf("error writing io.Writer: %w", err)
	}

	fmt.Printf("Wrote %d bytes to io.Writer\n", n)
	return nil

}

func main() {

	// ---------------------------------------------------
	// WRITING TO A BUFFER (*bytes.Buffer)
	fmt.Printf("\nWRITING TO A BUFFER\n")

	// CREATE THE BUFFER
	b := new(bytes.Buffer)
	fmt.Printf("Buffer before io.Writer: %s \n", b)

	// STREAM FROM BUFFER (b is *bytes.Buffer)
	err := writeFromBuffer(b)
	if err != nil {
		fmt.Printf("Error writing to buffer: %s\n", err)
	}

	// PRINT BUFFER
	fmt.Printf("Buffer after io.Writer: %s \n", b)

	// ---------------------------------------------------
	// WRITING TO A FILE (*os.File)
	fmt.Printf("\nWRITING TO A FILE\n")

	// CREAT FILE
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Printf("error creating file: %s", err)
	}
	defer f.Close()

	// STREAM FROM BUFFER (f is *os.File)
	err = writeFromBuffer(f)
	if err != nil {
		fmt.Printf("Error with writing file: %s\n", err)
	}

	// CHECK FILE
	fmt.Printf("Remember to check test.txt\n")

	// ---------------------------------------------------
	// WRITING TO STDOUT(os.Stdout)
	fmt.Printf("\nWRITE TO STDIN\n")

	// STDOUT
	o := os.Stdout

	// STREAM FROM BUFFER (os.Stdout)
	err = writeFromBuffer(o)
	if err != nil {
		fmt.Printf("Error writing to Stdout: %s\n", err)
	}

	// OUTPUT WILL BE ON TERMINAL

	// ---------------------------------------------------
	// READING FROM A PIPE (*io.PipeReader)
	fmt.Printf("\nREAD FROM A PIPE\n")

	// CREATE THE PIPE
	pipeReader, pipeWriter := io.Pipe()

	// STREAM FROM BUFFER (wpipe is *io.PipeWriter)
	go func() {

		err = writeFromBuffer(pipeWriter)
		if err != nil {
			fmt.Printf("Error reading pipe: %s\n", err)
		}
		pipeWriter.Close()

	}()

	// READ FROM PIPE AND PRINT
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(pipeReader)

	// PRINT PIPE
	fmt.Printf("Pipe output: %s", buffer.String())

}
