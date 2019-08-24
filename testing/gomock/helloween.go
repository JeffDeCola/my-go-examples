// my-go-examples helloween.go

package main

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/testing/gomock/creatures"
	"github.com/JeffDeCola/my-go-examples/testing/gomock/laboratory"
)

// INTERFACE AS A RETURN
func makeVampire(age int) laboratory.Creatures {
	return creatures.Vampire{Age: age}
}

func main() {

	// CREATE A WEREWOLF STRUCT
	// Pass struct in Greet/Flyaway
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

	// CREATE AN INTERFACE VAMPIRE
	// Pass interface to Greet/Flyway
	// var Clif laboratory.Creatures = creatures.Vampire{Age: 132}
	Clif := makeVampire(132)

	// PRINT VAMPIRE INFO
	fmt.Printf("Clif is a %s\n", Clif.Kind())
	fmt.Println(laboratory.Greet(Clif))
	fmt.Println(laboratory.FlyAway(Clif))
	fmt.Println("")

}
