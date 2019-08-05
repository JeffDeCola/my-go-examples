package jeffshapes

import (
	"math"
)

// Circle defines a circle
type Circle struct {
	R float64
}

// CircleArea is the area of a circle
func (c Circle) CircleArea() float64 {
	return math.Pi * math.Pow(c.R, 2)
}

// CircleCircumference is the circumference of a circle
func (c Circle) CircleCircumference() float64 {
	return 2 * math.Pi * c.R
}
