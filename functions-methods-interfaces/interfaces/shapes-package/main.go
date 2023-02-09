package main

import (
	"fmt"

	"github.com/JeffDeCola/my-go-packages/geometry/shapes"
)

func main() {

	// Create a Rectangle and Circle type
	r := shapes.Rectangle{Width: 10, Height: 10}
	c := shapes.Circle{Radius: 5}

	// Get the area (using the interface)
	//a := shapes.GetArea(r)
	a := r.Area()
	fmt.Println("Area of rectangle =", a)
	a = shapes.GetArea(c)
	fmt.Println("Area of circle =", a)

	// Change the size (x2)
	shapes.ChangeSize(&r, 2)
	shapes.ChangeSize(&c, 2)

	// Get the area (using the interface)
	a = shapes.GetArea(r)
	fmt.Println("Area of rectangle =", a)
	a = shapes.GetArea(c)
	fmt.Println("Area of circle =", a)

}
