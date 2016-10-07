package main

import "fmt"

// Takes in string and returns the count
func countLetters(word string) int {

	return len(word)

}

func main() {
	count := countLetters("happy")
	fmt.Println(count)
}
