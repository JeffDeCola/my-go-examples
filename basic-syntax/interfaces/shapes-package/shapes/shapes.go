package shapes

import (
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
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	A float64
	B float64
	C float64
}

// INTERFACES - WILL PIC THE RIGHT METHOD
func Area(a areaInterface) float64 {
	// CALL THE METHOD - WILL CHOSE RIGHT ONE
	area := a.theArea()
	return area
}

func Perimeter(p perimeterInterface) float64 {
	// CALL THE METHOD - WILL CHOSE RIGHT ONE
	perimeter := p.thePerimeter()
	return perimeter
}

// MATH USING METHODS
func (r Rectangle) theArea() float64 {
	area := r.Width * r.Height
	return area
}

func (r Rectangle) thePerimeter() float64 {
	perimeter := 2 * (r.Width + r.Height)
	return perimeter
}

func (c Circle) theArea() float64 {
	area := math.Pi * math.Pow(c.Radius, 2)
	return area
}

func (c Circle) thePerimeter() float64 {
	perimeter := 2 * math.Pi * c.Radius
	return perimeter
}

func (t Triangle) theArea() float64 {
	// Heron's Formula to get area from 3 sides
	s := ((t.A + t.B + t.C) / 2)
	perimeter := math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
	return perimeter
}

func (t Triangle) thePerimeter() float64 {
	perimeter := t.A + t.B + t.C
	return perimeter
}
