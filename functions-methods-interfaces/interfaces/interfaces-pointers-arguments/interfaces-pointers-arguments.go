package main

import (
	"fmt"
	"math"
)

// INTERFACE TYPES - TIES INTERFACE WITH METHODS
type geometry interface {
	area(*float64)
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
func (r Rectangle) area(a *float64) {
	*a = r.width * r.height
}

func (c Circle) area(a *float64) {
	*a = math.Pi * math.Pow(c.radius, 2)
}

func main() {

	// DEFINE
	rec := Rectangle{2.4, 34.4}
	circ := Circle{2.3}

	// CALCULATE AREA USING AN INTERFACE TYPE
	var gRec geometry
	var gCirc geometry
	var recArea float64
	var circArea float64

	gRec = rec
	gCirc = circ
	gRec.area(&recArea)
	gCirc.area(&circArea)

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f\n", rec.width, rec.height, recArea)
	fmt.Printf("Circle (%.2f): Area=%.2f\n", circ.radius, circArea)

}
