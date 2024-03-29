package main

import (
	"fmt"
	"math"
)

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

func main() {

	// DEFINE
	rec := rectangle{2.4, 34.4}
	circ := circle{2.3}

	// CALCULATE AREA USING A STRUCT TYPE
	recArea := rec.area()
	circArea := circ.area()

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f\n", rec.width, rec.height, recArea)
	fmt.Printf("Circle (%.2f): Area=%.2f\n", circ.radius, circArea)

}
