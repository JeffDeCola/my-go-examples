// functions
//
// Using functions to calculate the area of a rectangle and circle.
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

}
