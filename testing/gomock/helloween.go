// my-go-examples helloween.go

package main

import (
	"fmt"

	"./creatures"
	"./laboratory"
)

func main() {

	// CREATE A WEREWOLF
	Jeff := creatures.Werewolf{TimeofDay: "night"}

	// PRINT WEREWOLF INFO
	fmt.Printf("Jeff is a %s\n", Jeff.Kind())
	fmt.Println(laboratory.Greet(Jeff))
	fmt.Println(laboratory.FlyAway(Jeff))
	fmt.Println("")
	// Change the time of day
	Jeff.TimeofDay = "day"
	fmt.Printf("Jeff is a %s\n", Jeff.Kind())
	fmt.Println(laboratory.Greet(Jeff))
	fmt.Println(laboratory.FlyAway(Jeff))
	fmt.Println("")

	// CREATE A VAMPIRE
	var Clif laboratory.Creatures = creatures.Vampire{Age: 132}

	// PRINT VAMPIRE INFO
	fmt.Printf("Clif is a %s\n", Clif.Kind())
	fmt.Println(laboratory.Greet(Clif))
	fmt.Println(laboratory.FlyAway(Clif))
	fmt.Println("")

}
