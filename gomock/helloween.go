package main

import (
	"fmt"

	"github.com/JeffDeCola/my-go-examples/gomock/creatures"
	"github.com/JeffDeCola/my-go-examples/gomock/laboratory"
)

func main() {

	creature := creatures.Jeff{TimeofDay: "day"}
	fmt.Printf("Jeff is a %s\n", creature.Kind())
	sound := laboratory.CreatureSound(creature)
	fmt.Println(sound)
	fly := laboratory.CreatureFly(creature)
	fmt.Println(fly)
	fmt.Println("")

	var creature2 laboratory.Creatures = creatures.Steven{Age: 132, Height: 56} // Another way to write this.
	fmt.Printf("Steven is a %s\n", creature2.Kind())
	sound = laboratory.CreatureSound(creature2)
	fmt.Println(sound)
	fly = laboratory.CreatureFly(creature2)
	fmt.Println(fly)
	fmt.Println("")

	creature3 := creatures.Doris("hello")
	fmt.Printf("Doris is a %s\n", creature3.Kind())
	//sound = laboratory.CreatureSound(creature3)
	fmt.Println(laboratory.CreatureSound(creature3))
	fly = laboratory.CreatureFly(creature3)
	fmt.Println(fly)

}
