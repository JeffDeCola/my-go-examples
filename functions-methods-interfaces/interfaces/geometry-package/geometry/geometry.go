package geometry

import (
	"math"
)

// INTERFACE TYPES - TIES INTERFACE WITH METHODS
type Geometry interface {
	Area(*float64)
	Perimeter(*float64)
	Size(float64)
}

// RECTANGLE SHAPE CHARACTERISTICS
type Rectangle struct {
	Width  float64
	Height float64
}

// CIRCLE SHAPE CHARACTERISTICS
type Circle struct {
	Radius float64
}

// TRIANGLE SHAPE CHARACTERISTICS
type Triangle struct {
	A float64
	B float64
	C float64
}

// AREA FOR RECTANGLE
func (r Rectangle) Area(a *float64) {
	*a = r.Width * r.Height
}

// AREA FOR CIRCLE
func (c Circle) Area(a *float64) {
	*a = math.Pi * math.Pow(c.Radius, 2)
}

// AREA FOR TRIANGLE
func (t Triangle) Area(a *float64) {
	// Heron's Formula to get area from 3 sides
	s := ((t.A + t.B + t.C) / 2)
	*a = math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

// PERIMETER FOR RECTANGLE
func (r Rectangle) Perimeter(p *float64) {
	*p = 2 * (r.Width + r.Height)
}

// PERIMETER FOR CIRCLE
func (c Circle) Perimeter(p *float64) {
	*p = 2 * math.Pi * c.Radius
}

// Perimeter of a triangle
func (t Triangle) Perimeter(a *float64) {
	*a = t.A + t.B + t.C
}

// Size of a rectangle
func (r *Rectangle) Size(f float64) {
	r.Width = f * r.Width
	r.Height = f * r.Height
}

// Calculate Size of a Circle
func (c *Circle) Size(f float64) {
	c.Radius = f * c.Radius
}

// Calculate Size of a Triangle
func (t *Triangle) Size(f float64) {
	t.A = f * t.A
	t.B = f * t.B
	t.C = f * t.C
}
