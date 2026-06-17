// functions-pointers-arguments
//
// Using functions to resize a rectangle and circle by passing pointers.
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

func areaRectangle(r rectangle) float64 {
	area := r.width * r.height
	return area
}

func areaCircle(c circle) float64 {
	area := math.Pi * c.radius * c.radius
	return area
}

func scaleRectangle(r *rectangle, factor float64) {
	r.width *= factor
	r.height *= factor
}

func scaleCircle(c *circle, factor float64) {
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

	// Pass by value - A complete copy is made
	recArea := areaRectangle(rec)
	fmt.Printf("The area of the rectangle (%.2f x %.2f) is %.2f\n", rec.width, rec.height, recArea)

	circArea := areaCircle(circ)
	fmt.Printf("The area of the circle (radius %.2f) is %.2f\n", circ.radius, circArea)

	// Pass by pointers
	scaleRectangle(&rec, 3)
	recArea = areaRectangle(rec)
	fmt.Printf("The area of the rectangle (%.2f x %.2f) is %.2f\n", rec.width, rec.height, recArea)

	scaleCircle(&circ, 4)
	circArea = areaCircle(circ)
	fmt.Printf("The area of the circle (radius %.2f) is %.2f\n", circ.radius, circArea)

}
