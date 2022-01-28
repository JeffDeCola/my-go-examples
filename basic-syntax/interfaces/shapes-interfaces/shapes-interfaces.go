package main

import (
	"fmt"
)

// INTERFACE TYPE
type areaInterface interface {
	areaRec() float32
}

type perimeterInterface interface {
	perimeterRec() float32
}

// STRUCT TYPE
type Rectangle struct {
	width  float32
	height float32
}

func area(a areaInterface) float32 {
	return a.areaRec()
}

func perimeter(p perimeterInterface) float32 {
	return p.perimeterRec()
}

func (r Rectangle) areaRec() float32 {
	area := r.width * r.height
	return area
}

func (r Rectangle) perimeterRec() float32 {
	perimeter := 2 * (r.width + r.height)
	return perimeter
}

func main() {

	// DEFINE
	rec := Rectangle{2.4, 34.4}

	// CALCULATE AREA & PERIMETER
	recArea := area(rec)
	recPerimeter := perimeter(rec)

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f, Perimeter=%.2f\n", rec.width, rec.height, recArea, recPerimeter)
}
