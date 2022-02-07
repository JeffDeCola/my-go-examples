package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func noPipe() {

	var w io.Writer
	var r io.Reader

	// INTERMEDIATE BUFFER
	b := new(bytes.Buffer)

	// I USE IO.WRITER
	w = b
	fmt.Fprint(w, "I am some data without pipe\n")

	// I USE IO.READER
	r = b
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}

func withPipe() {

	// Pipes in go can be used to connect code expecting an io.Reader with
	// code expecting an io.Writer.

	pr, pw := io.Pipe()

	// I USE IO.WRITER
	go func() {
		fmt.Fprint(pw, "I am some data using pipe\n")
		pw.Close()
	}()

	// I USE IO.READER
	if _, err := io.Copy(os.Stdout, pr); err != nil {
		log.Fatal(err)
	}

}

func main() {

	noPipe()
	withPipe()

}
