// pointers

package main

import (
	"fmt"
)

type person struct {
	name   string
	gender string
	age    int
}

func changeName(p person) {
	p.name = "Fred"
}

func changeNamePtr(p *person) {
	p.name = "Fred"
}

func main() {
	a := person{"Larry", "male", 25}
	b := person{"Bob", "male", 27}

	fmt.Println("Pass struct to function - Makes Copy")
	fmt.Println(a.name, a.gender, a.age)
	changeName(a)
	fmt.Println(a.name, a.gender, a.age)

	fmt.Println("Pass struct Address to function - Original")
	fmt.Println(b.name, b.gender, b.age)
	changeNamePtr(&b)
	fmt.Println(b.name, b.gender, b.age)
}
