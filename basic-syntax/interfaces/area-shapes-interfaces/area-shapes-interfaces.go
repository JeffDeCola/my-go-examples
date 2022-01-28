package main

import (
	"fmt"
	"math"
)

// INTERFACE TYPES - TIES INTERFACE WITH METHODS
type geometry interface {
	theArea() float64
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
func (r Rectangle) theArea() float64 {
	area := r.width * r.height
	return area
}

func (c Circle) theArea() float64 {
	area := math.Pi * math.Pow(c.radius, 2)
	return area
}

// FUNCTION USING AN INTERFACE - I FIND THIS CLEANER
func area(g geometry) float64 {
	// CALL THE METHOD - WILL CHOSE RIGHT ONE
	area := g.theArea()
	return area
}

func main() {

	// DEFINE
	rec := Rectangle{2.4, 34.4}
	circ := Circle{2.3}

	// CALCULATE AREA OF SHAPES
	recArea := area(rec)
	circArea := area(circ)

	// I COULD CALCULATE AREA WITHOUT USING FUNCTION
	// var g geometry
	// g = rec
	// recArea := g.theArea()
	// g = circ
	// circArea := g.thearea()

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f\n", rec.width, rec.height, recArea)
	fmt.Printf("Circle (%.2f): Area=%.2f\n", circ.radius, circArea)
}
