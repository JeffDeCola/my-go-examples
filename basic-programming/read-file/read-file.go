// my-go-examples read-file.go

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fmt.Println("Method 1 (ioutil.ReadFile) - Read entire file contents into memory")
	data, err := ioutil.ReadFile("readthis.txt")
	check(err)
	fmt.Println(string(data))
	fmt.Println("")

	fmt.Println("Method 2 (os.Open) - Open the file and read what you want")
	f, err := os.Open("readthis.txt")
	check(err)
	// Read up to 10 bytes of file
	b1 := make([]byte, 10)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))
	f.Close()
	fmt.Println("")

	fmt.Println("Method 3 (bufio.NewReader) - Buffered reading using bufio package")
	g, err := os.Open("readthis.txt")
	check(err)
	r4 := bufio.NewReader(g)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	g.Close()

}
