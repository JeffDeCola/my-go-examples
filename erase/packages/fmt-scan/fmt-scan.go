// my-go-examples fmt-scan.go

package main

import "fmt"

func main() {

	word1 := ""

	fmt.Printf("What is your name? ")
	_, err := fmt.Scan(&word1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %v.\n", word1)

}
