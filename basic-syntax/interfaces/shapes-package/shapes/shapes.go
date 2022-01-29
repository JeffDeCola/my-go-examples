package shapes

import (
	"math"
)

// INTERFACE TYPES - TIES INTERFACE WITH METHODS
type Geometry interface {
	Area() float64
	Perimeter() float64
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

// MATH USING METHODS
func (r Rectangle) Area() float64 {
	area := r.Width * r.Height
	return area
}

func (r Rectangle) Perimeter() float64 {
	perimeter := 2 * (r.Width + r.Height)
	return perimeter
}

func (c Circle) Area() float64 {
	area := math.Pi * math.Pow(c.Radius, 2)
	return area
}

func (c Circle) Perimeter() float64 {
	perimeter := 2 * math.Pi * c.Radius
	return perimeter
}

func (t Triangle) Area() float64 {
	// Heron's Formula to get area from 3 sides
	s := ((t.A + t.B + t.C) / 2)
	perimeter := math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
	return perimeter
}

func (t Triangle) Perimeter() float64 {
	perimeter := t.A + t.B + t.C
	return perimeter
}

// FUNCTION USING AN INTERFACE
func GetArea(g Geometry) float64 {
	// CALL THE METHOD - WILL CHOSE RIGHT ONE
	area := g.Area()
	return area
}

func GetPerimeter(g Geometry) float64 {
	// CALL THE METHOD - WILL CHOSE RIGHT ONE
	perimeter := g.Perimeter()
	return perimeter
}
