package main

import (
	"fmt"

	shapes "github.com/JeffDeCola/my-go-packages/geometry/circle"
)

func main() {

	c := shapes.Circle{R: 5}

	area := c.CircleArea()
	fmt.Println(area)

	circumference := c.CircleCircumference()
	fmt.Println(circumference)

}
