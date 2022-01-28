package main

import "fmt"

// STRUCT TYPE
type Rectangle struct {
	width  float32
	height float32
}

func (r Rectangle) areaRectangle() float32 {
	area := r.width * r.height
	return area
}

func (r Rectangle) perimeterRectangle() float32 {
	perimeter := 2 * (r.width + r.height)
	return perimeter
}

func main() {

	// DEFINE
	rec := Rectangle{2.4, 34.4}

	// CALCULATE AREA & PERIMETER
	recArea := rec.areaRectangle()
	recPerimeter := rec.perimeterRectangle()

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f, Perimeter=%.2f\n", rec.width, rec.height, recArea, recPerimeter)
}
