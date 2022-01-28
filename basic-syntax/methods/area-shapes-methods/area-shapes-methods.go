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
func (r Rectangle) areaRectangle() float64 {
	area := r.width * r.height
	return area
}

func (c Circle) areaCircle() float64 {
	area := math.Pi * math.Pow(c.radius, 2)
	return area
}

func main() {

	// DEFINE
	rec := Rectangle{2.4, 34.4}
	circ := Circle{2.3}

	// CALCULATE AREA OF SHAPES
	recArea := rec.areaRectangle()
	circArea := circ.areaCircle()

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f\n", rec.width, rec.height, recArea)
	fmt.Printf("Circle (%.2f): Area=%.2f\n", circ.radius, circArea)

}
