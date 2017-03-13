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

func changeName(p *person) {
	p.name = "Fred"
}

func main() {
	UserID := person{"Larry", "male", 25}
	fmt.Println(UserID.name, UserID.gender, UserID.age)
	//Pass address of struct
	changeName(&UserID)
	fmt.Println(UserID.name, UserID.gender, UserID.age)
}
