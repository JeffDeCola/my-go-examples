// methods-pointers-receivers
//
// Using methods to resize a rectangle and circle using pointer receivers.
//

package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	area := r.width * r.height
	return area
}

func (c circle) area() float64 {
	area := math.Pi * c.radius * c.radius
	return area
}

func (r *rectangle) scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func (c *circle) scale(factor float64) {
	c.radius *= factor
}

func main() {

	rec := rectangle{
		width:  10,
		height: 5,
	}

	circ := circle{
		radius: 5,
	}

	// Value receiver - reads a copy
	recArea := rec.area()
	fmt.Printf("The area of the rectangle (%.2f x %.2f) is %.2f\n", rec.width, rec.height, recArea)

	circArea := circ.area()
	fmt.Printf("The area of the circle (radius %.2f) is %.2f\n", circ.radius, circArea)

	// Pointer receiver - mutates the original
	rec.scale(3)
	recArea = rec.area()
	fmt.Printf("The area of the rectangle (%.2f x %.2f) is %.2f\n", rec.width, rec.height, recArea)

	circ.scale(4)
	circArea = circ.area()
	fmt.Printf("The area of the circle (radius %.2f) is %.2f\n", circ.radius, circArea)

}
