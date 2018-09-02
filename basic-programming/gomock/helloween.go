// my-go-examples helloween.go

package main

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/gomock/creatures"
	"github.com/JeffDeCola/my-go-examples/gomock/laboratory"
)

func main() {

	Jeff := creatures.Warewolf{TimeofDay: "night"} // Get data
	fmt.Printf("Jeff is a %s\n", Jeff.Kind())
	sound := laboratory.Greet(Jeff)
	fmt.Println(sound)
	fly := laboratory.FlyAway(Jeff)
	fmt.Println(fly)
	fmt.Println("")

	Jeff.TimeofDay = "day"
	fmt.Printf("Jeff is a %s\n", Jeff.Kind())
	sound = laboratory.Greet(Jeff)
	fmt.Println(sound)
	fly = laboratory.FlyAway(Jeff)
	fmt.Println(fly)
	fmt.Println("")

	var Clif laboratory.Creatures = creatures.Vampire{Age: 132} // Interface - Access to behavior.
	fmt.Printf("Clif is a %s\n", Clif.Kind())
	sound = laboratory.Greet(Clif)
	fmt.Println(sound)
	fly = laboratory.FlyAway(Clif)
	fmt.Println(fly)
	fmt.Println("")

}
