package main

import (
	"fmt"

	"github.com/JeffDeCola/my-go-packages/geometry/circle"
)

func main() {

	// Create a Circle type
	c := circle.Circle{Radius: 5}

	// Get the area
	a := c.Area()
	fmt.Println("Area =", a)

	// Get the circumference
	p := c.Circumference()
	fmt.Println("Circumference =", p)

}
