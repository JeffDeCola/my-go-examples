// my-go-examples pointers1.go

package main

import (
	"fmt"
)

type person struct {
	name   string
	gender string
	age    int
}

// Pass by value
func changeName(p person) {
	p.name = "Fred"
}

// Pass by reference
func changeNamePtr(p *person) {
	p.name = "Lisa"
}

func main() {
	a := person{"Larry", "male", 25}
	b := person{"Jill", "female", 27}

	fmt.Println("Pass struct to function by Value - Makes Copy")
	fmt.Println(a)
	changeName(a)
	fmt.Println(a)

	fmt.Println("Pass struct Address to function by reference - Uses Original")
	fmt.Println(b)
	changeNamePtr(&b)
	fmt.Println(b)

}
