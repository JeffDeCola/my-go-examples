package main

import "fmt"

func areaRectangle(w float32, h float32) float32 {
	area := w * h
	return area
}

func perimeterRectangle(w float32, h float32) float32 {
	perimeter := 2 * (w + h)
	return perimeter
}

func main() {

	// DEFINE
	var recWidth float32 = 2.4
	var recHeight float32 = 34.4

	// CALCULATE AREA & PERIMETER
	recArea := areaRectangle(recWidth, recHeight)
	recPerimeter := perimeterRectangle(recWidth, recHeight)

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f, Perimeter=%.2f\n", recWidth, recHeight, recArea, recPerimeter)

}
