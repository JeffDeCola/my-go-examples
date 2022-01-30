package shapes

import (
	"math"
)

// INTERFACE TYPES - TIES INTERFACE WITH METHODS
type Geometry interface {
	Area(*float64)
	Perimeter(*float64)
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
func (r Rectangle) Area(a *float64) {
	*a = r.Width * r.Height
}

func (r Rectangle) Perimeter(p *float64) {
	*p = 2 * (r.Width + r.Height)
}

func (c Circle) Area(a *float64) {
	*a = math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Perimeter(p *float64) {
	*p = 2 * math.Pi * c.Radius
}

func (t Triangle) Area(a *float64) {
	// Heron's Formula to get area from 3 sides
	s := ((t.A + t.B + t.C) / 2)
	*a = math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter(a *float64) {
	*a = t.A + t.B + t.C
}

// FUNCTION USING AN INTERFACE
func GetArea(g Geometry) float64 {
	// CALL THE METHOD - WILL CHOSE RIGHT ONE
	var area float64
	g.Area(&area)
	return area
}

func GetPerimeter(g Geometry) float64 {
	// CALL THE METHOD - WILL CHOSE RIGHT ONE
	var perimeter float64
	g.Perimeter(&perimeter)
	return perimeter
}
