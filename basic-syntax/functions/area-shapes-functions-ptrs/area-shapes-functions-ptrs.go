package main

import (
	"fmt"
	"math"
)

// MATH USING FUNCTIONS
func areaRectangle(w float64, h float64, a *float64) {
	*a = w * h
}

func areaCircle(r float64, a *float64) {
	*a = math.Pi * math.Pow(r, 2)
}

func main() {

	// DEFINE
	var recWidth float64 = 2.4
	var recHeight float64 = 34.4
	var circRadius float64 = 2.3
	var recArea float64
	var circArea float64

	// CALCULATE AREA OF SHAPES
	areaRectangle(recWidth, recHeight, &recArea)
	areaCircle(circRadius, &circArea)

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f\n", recWidth, recHeight, recArea)
	fmt.Printf("Circle (%.2f): Area=%.2f\n", circRadius, circArea)

}
