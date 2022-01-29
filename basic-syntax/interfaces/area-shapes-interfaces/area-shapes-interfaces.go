package main

import (
	"fmt"
	"math"
)

// INTERFACE TYPES - TIES INTERFACE WITH METHODS
type geometry interface {
	area() float64
}

//  STRUCTS - SHAPE CHARACTERISTICS
type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

// MATH USING METHODS
func (r Rectangle) area() float64 {
	area := r.width * r.height
	return area
}

func (c Circle) area() float64 {
	area := math.Pi * math.Pow(c.radius, 2)
	return area
}

// FUNCTION USING AN INTERFACE - I FIND THIS CLEANER
func getArea(g geometry) float64 {
	// CALL THE METHOD - WILL CHOSE RIGHT ONE
	area := g.area()
	return area
}

func main() {

	// DEFINE
	rec := Rectangle{2.4, 34.4}
	circ := Circle{2.3}

	// CALCULATE AREA
	// USING AN INTERFACE THAT CALLS THE CORRECT METHOD
	var g geometry
	g = rec
	recArea := g.area()
	g = circ
	circArea := g.area()

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f\n", rec.width, rec.height, recArea)
	fmt.Printf("Circle (%.2f): Area=%.2f\n", circ.radius, circArea)

	// I COULD ALSO CALCULATE AREA USING A FUNCTION THAT USES THE INTERFACE
	recArea = getArea(rec)
	circArea = getArea(circ)

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f\n", rec.width, rec.height, recArea)
	fmt.Printf("Circle (%.2f): Area=%.2f\n", circ.radius, circArea)

}
