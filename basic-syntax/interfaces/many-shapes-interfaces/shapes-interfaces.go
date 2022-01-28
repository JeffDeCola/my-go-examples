package main

import (
	"fmt"
	"math"
)

// INTERFACE TYPES - TIES INTERFACE WITH METHODS
type areaInterface interface {
	theArea() float64
}

type perimeterInterface interface {
	thePerimeter() float64
}

//  STRUCTS - SHAPE CHARACTERISTICS
type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

// INTERFACES - WILL PIC THE RIGHT METHOD
func area(a areaInterface) float64 {
	// CALL THE METHOD - WILL CHOSE RIGHT ONE
	area := a.theArea()
	return area
}

func perimeter(p perimeterInterface) float64 {
	// CALL THE METHOD - WILL CHOSE RIGHT ONE
	perimeter := p.thePerimeter()
	return perimeter
}

// MATH USING METHODS
func (r Rectangle) theArea() float64 {
	area := r.width * r.height
	return area
}

func (r Rectangle) thePerimeter() float64 {
	perimeter := 2 * (r.width + r.height)
	return perimeter
}

func (r Circle) theArea() float64 {
	area := math.Pi * math.Pow(r.radius, 2)
	return area
}

func (r Circle) thePerimeter() float64 {
	perimeter := 2 * math.Pi * r.radius
	return perimeter
}

func main() {

	// DEFINE
	rec := Rectangle{2.4, 34.4}
	circ := Circle{2.3}

	// CALCULATE AREA & PERIMETER - THE POWER OF THE INTERFACE
	recArea := area(rec)
	recPerimeter := perimeter(rec)
	circArea := area(circ)
	circPerimeter := perimeter(circ)

	// PRINT
	fmt.Printf("Rectangle (%.2f x %.2f): Area=%.2f, Perimeter=%.2f\n", rec.width, rec.height, recArea, recPerimeter)
	fmt.Printf("Circle (%.2f): Area=%.2f, Perimeter=%.2f\n", circ.radius, circArea, circPerimeter)
}
