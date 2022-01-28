package main

import (
	"fmt"
	"math"
)

// MATH USING FUNCTIONS
func areaRectangle(w float64, h float64) float64 {
	area := w * h
	return area
}

func areaCircle(r float64) float64 {
	area := math.Pi * math.Pow(r, 2)
	return area
}

func main() {

	// DEFINE
	var recWidth float64 = 2.4
	var recHeight float64 = 34.4
	var circRadius float64 = 2.3

	// CALCULATE AREA OF SHAPES
	recArea := areaRectangle(recWidth, recHeight)
	circArea := areaCircle(circRadius)

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f\n", recWidth, recHeight, recArea)
	fmt.Printf("Circle (%.2f): Area=%.2f\n", circRadius, circArea)

}
