// my-go-examples (methods-and-interfaces) no-structs.go

package main

import (
	"fmt"
	"math"
)

// Find the distance between two points
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// Compute the area of a rectagle
func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

// Compute the area of a circle
func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

func main() {

	// rect1
	var rx1, ry1 float64 = 1, 1
	var rx2, ry2 float64 = 11, 11

	// circ1
	var cx, cy, cr float64 = 1, 1, 5

	fmt.Println("Area of rect1 is: ", rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println("Area of circ1 is: ", circleArea(cx, cy, cr))

}
