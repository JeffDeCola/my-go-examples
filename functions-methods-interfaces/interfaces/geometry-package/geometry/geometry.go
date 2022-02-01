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

func (c Circle) Area(a *float64) {
	*a = math.Pi * math.Pow(c.Radius, 2)
}

func (t Triangle) Area(a *float64) {
	// Heron's Formula to get area from 3 sides
	s := ((t.A + t.B + t.C) / 2)
	*a = math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (r Rectangle) Perimeter(p *float64) {
	*p = 2 * (r.Width + r.Height)
}

func (c Circle) Perimeter(p *float64) {
	*p = 2 * math.Pi * c.Radius
}

func (t Triangle) Perimeter(a *float64) {
	*a = t.A + t.B + t.C
}

func (r *Rectangle) Size(f float64) {
	r.Width = f * r.Width
	r.Height = f * r.Height
}

func (c *Circle) Size(f float64) {
	c.Radius = f * c.Radius
}

func (t *Triangle) Size(f float64) {
	t.A = f * t.A
	t.B = f * t.B
	t.C = f * t.C
}
