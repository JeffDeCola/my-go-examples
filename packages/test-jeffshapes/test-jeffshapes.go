package main

import (
	"fmt"

	"github.com/JeffDeCola/my-go-packages/jeffshapes"
)

func main() {

	c1 := jeffshapes.Circle{R: 5.0}

	fmt.Printf("The Area of Circle of radius %v is %v\n", c1.R, c1.CircleArea())
	fmt.Printf("The Circumference of Circle of radius %v is %v\n", c1.R, c1.CircleCircumference())

}
