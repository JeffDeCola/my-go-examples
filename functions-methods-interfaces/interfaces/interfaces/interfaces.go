package main

import (
	"fmt"
	"math"
)

// INTERFACE TYPES - TIES INTERFACE WITH METHODS
type geometry interface {
	area() float64
}

// STRUCTS - SHAPE CHARACTERISTICS
type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

// MATH USING METHODS
func (r rectangle) area() float64 {
	area := r.width * r.height
	return area
}

func (c circle) area() float64 {
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
	rec := rectangle{2.4, 34.4}
	circ := circle{2.3}

	// CALCULATE AREA USING AN INTERFACE TYPE
	var gRec geometry
	var gCirc geometry
	gRec = rec
	gCirc = circ
	recArea := gRec.area()
	circArea := gCirc.area()

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f\n", rec.width, rec.height, recArea)
	fmt.Printf("Circle (%.2f): Area=%.2f\n", circ.radius, circArea)

	// I COULD WRAP THE INTERFACE IN A FUNCTION
	recArea = getArea(rec)
	circArea = getArea(circ)

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f\n", rec.width, rec.height, recArea)
	fmt.Printf("Circle (%.2f): Area=%.2f\n", circ.radius, circArea)

}
