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
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}

}

func withPipe() {

	// CREATE UNNAMED PIPE
	pr, pw := io.Pipe()

	// WRITE INTO PIPE USING IO.WRITER
	go func() {
		fmt.Fprint(pw, "I am some data using pipe\n")
		pw.Close()
	}()

	// READ FROM PIPE USING IO.READER
	_, err := io.Copy(os.Stdout, pr)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	noPipe()
	withPipe()

}
