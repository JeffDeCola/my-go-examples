package main

import (
	"fmt"
	"math"
)

//  STRUCTS - SHAPE CHARACTERISTICS
type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

// MATH USING METHODS
func (r Rectangle) area(a *float64) {
	*a = r.width * r.height
}

func (c Circle) area(a *float64) {
	*a = math.Pi * math.Pow(c.radius, 2)
}

func (r *Rectangle) size(f float64) {
	r.width = r.width * f
	r.height = r.height * f
}

func (c *Circle) size(f float64) {
	c.radius = c.radius * f
}

func main() {

	// DEFINE
	rec := Rectangle{2.4, 34.4}
	circ := Circle{2.3}
	var recArea float64
	var circArea float64

	// CALCULATE AREA USING A STRUCT TYPE
	rec.area(&recArea)
	circ.area(&circArea)

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f\n", rec.width, rec.height, recArea)
	fmt.Printf("Circle (%.2f): Area=%.2f\n", circ.radius, circArea)

	// INCREASE SIZE BY A FACTOR OF 2 USING POINTER RECIEVER
	fmt.Println("Increase size by a factor of 2")
	rec.size(2)
	circ.size(2)

	// CALCULATE AREA USING A STRUCT TYPE (INCREASED BY 2)
	rec.area(&recArea)
	circ.area(&circArea)

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f\n", rec.width, rec.height, recArea)
	fmt.Printf("Circle (%.2f): Area=%.2f\n", circ.radius, circArea)

}
